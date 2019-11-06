package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	appsv1 "k8s.io/api/apps/v1"
	apiv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

func int32Ptr(i int32) *int32 { return &i }

type Controller struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata"`
	Spec              ControllerSpec   `json:"spec"`
	Status            ControllerStatus `json:"status"`
}

type ControllerSpec struct {
	Deployments []*appsv1.Deployment `json:"deployments"`
}

type ControllerStatus struct {
	Deployments int `json:"deployments"`
	Succeeded   int `json:"succeeded"`
}

type SyncRequest struct {
	Parent   Controller          `json:"parent"`
	Children SyncRequestChildren `json:"children"`
}

type SyncRequestChildren struct {
	Deployments map[string]*appsv1.Deployment `json:"Deployment.v1`
}

type SyncResponse struct {
	Status   ControllerStatus `json:"status"`
	Children []runtime.Object `json:"children"`
}

func sync(request *SyncRequest) (*SyncResponse, error) {
	r := SyncResponse{}
	counter := 0

	log.Print("Running sync")
	for _, d := range request.Parent.Spec.Deployments {
		log.Print("Adding deployment")
		r.Children = append(r.Children, getDeployment(request.Parent.Name, counter, request.Parent.Namespace, request.Parent.Labels, d))
		counter += 1
	}

	return &r, nil
}

func getDeployment(parentName string, count int, namespace string, labels map[string]string, d *appsv1.Deployment) *appsv1.Deployment {

	log.Print("Creating deployment", d)

	deployment := &appsv1.Deployment{
		TypeMeta: metav1.TypeMeta{
			APIVersion: "apps/v1",
			Kind:       "Deployment",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      parentName + "-" + strconv.Itoa(count),
			Namespace: namespace,
		},
		Spec: appsv1.DeploymentSpec{
			Replicas: d.Spec.Replicas,
			Selector: &metav1.LabelSelector{
				MatchLabels: d.Spec.Selector.MatchLabels, // TODO : SHOULD BE PROVIDED
			},
			Template: apiv1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: d.Spec.Selector.MatchLabels, // TODO : SHOULD BE PROVIDED
				},
				Spec: apiv1.PodSpec{
					Containers: []apiv1.Container{
						{
							Name:  "web",
							Image: "nginx:1.12",
							Ports: []apiv1.ContainerPort{
								{
									Name:          "http",
									Protocol:      apiv1.ProtocolTCP,
									ContainerPort: 80,
								},
							},
						},
					},
				},
			},
		},
	}

	return deployment
}

func syncHandler(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	request := &SyncRequest{}
	if err := json.Unmarshal(body, request); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	response, err := sync(request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	body, err = json.Marshal(&response)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(body)
}

func main() {
	log.Print("Serving on port 8080")
	http.HandleFunc("/sync", syncHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
