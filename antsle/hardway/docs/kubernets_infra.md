# Kubernetes the Hard way

## Education
- Minikube
- Single node
- Kubeadm
## Production cluster
- Limits
    - 5,000 nodes
    - 150,000 PODs in a cluster
    - 300,000 Total Containers
    - 100 Pods per node
- Add taint to master nodes

- Minikube
    - Deploys VM's
- Kubeadm
    - Expects vm's to be provisioned already

# Installation
- Controller manager
    - Leader elect option by default it's set to true
    - Try's to optain a lock on leader object
- etcd
    - Stacked Topology
        - Same nodes in the control plane
    - External ETCD Topology
        - etcd is external to the cluster
    - API server is the only thing that talks to this
    - Writes
        - only persisted once all nodes agree
        - Writes are complete when it can be completed on the "majority" of the nodes, better known as, quarum.
            - Quorum = N / 2 + 1
                - If there's a decimal, consider the whole number only.
                

## TLS Node bootstrapping
- Let nodes manage their own certificates
- Steps
    - Generate token per node (or one for all)
    - Associate it to group ```system:bootstrappers```
    - Assign Role(s) to ```system:bootstrappers```:
        - ```system:node-bootstarpper```
        - ```system:certificates.k8s.io:certificatesigningrequests:nodeclient```
        - ```system:certificates.k8s.io:certificatesigningrequests:selfnodeclient```

## Networking
- Deploy Weave
    - ```kubectl apply -f https://cloud.weave.works/k8s/net?k8s-version=$(kubectl version | base64 | tr -d '\n')```
- IPAM (IP Address Management)
    - it's the responsibility of CNI plugin
- CoreDNS
    - ```/etc/coredns/Corefile```