module sample-source

go 1.13

require (
	github.com/cloudevents/sdk-go v0.10.0
	github.com/go-logr/logr v0.1.0
	github.com/knative/eventing-sources v0.9.0
	github.com/onsi/ginkgo v1.7.0
	github.com/onsi/gomega v1.4.3
	gopkg.in/go-playground/webhooks.v3 v3.13.0
	k8s.io/api v0.0.0-20190409021203-6e4e0e4f393b
	k8s.io/apimachinery v0.0.0-20190404173353-6a84e37a896d
	k8s.io/client-go v11.0.1-0.20190409021438-1a26190bd76a+incompatible
	sigs.k8s.io/controller-runtime v0.2.2
)
