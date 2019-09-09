# E2E Knative Provisioner

## Requirements
| CLI                     | Required | Notes | 
|-------------------------|----------|---------|
| eksctl                   | no       | used for AWS EKS prosivioning   |
| kubectl                 | yes      | used for deployment of knative and istio|
| helm                    | yes      | used for deployment of istio|
| openssl                 | no       | only required of certs are not provided|



## Known Issues
- The example broke namespaces does not deploy the broker initially. A delete and re-run fixes it. (Race condition, I believe...)


## Execution

**No Args**
```
git clone https://github.com/Manny2014/kubernetes.git
cd kubernetes/ansible
ansible-playbook main.yaml 
```


**With Var overrides**
```
git clone https://github.com/Manny2014/kubernetes.git
cd kubernetes/ansible
ansible-playbook main.yaml -extra-vars "cluster_name=E2E"
```


## TODO
- Add full support for eksctl options
- Add additional provisioners for k8s cluster
- Remove source controlled version of istio, and download based on version