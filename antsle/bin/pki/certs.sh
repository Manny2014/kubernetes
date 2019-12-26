# Download 
# curl -o cfssl https://storage.googleapis.com/kubernetes-the-hard-way/cfssl/darwin/cfssl
# curl -o cfssljson https://storage.googleapis.com/kubernetes-the-hard-way/cfssl/darwin/cfssljson

# Generate CA Certs
cfssl gencert -initca conf/pki/ca-csr.json | cfssljson -bare ca

# Generate Admin Client Certs
cfssl gencert -ca=certs/ca/ca.pem -ca-key=certs/ca/ca-key.pem -config=conf/pki/ca-config.json -profile=kubernetes conf/pki/admin-csr.json | cfssljson -bare certs/admin/admin

# Generate Certs for Worker Nodes (REDO!)
cfssl gencert -ca=certs/ca/ca.pem -ca-key=certs/ca/ca-key.pem -config=conf/pki/ca-config.json -hostname=10.10.1.20,192.168.1.122 -profile=kubernetes conf/pki/antlet20-csr.json | cfssljson -bare certs/workers/antlet20
cfssl gencert -ca=certs/ca/ca.pem -ca-key=certs/ca/ca-key.pem -config=conf/pki/ca-config.json -hostname=10.10.1.21,192.168.1.122 -profile=kubernetes conf/pki/antlet21-csr.json | cfssljson -bare certs/workers/antlet21
cfssl gencert -ca=certs/ca/ca.pem -ca-key=certs/ca/ca-key.pem -config=conf/pki/ca-config.json -hostname=10.10.1.22,192.168.1.122 -profile=kubernetes conf/pki/antlet22-csr.json | cfssljson -bare certs/workers/antlet22

# Generate Kube-Controller Certs
cfssl gencert -ca=certs/ca/ca.pem -ca-key=certs/ca/ca-key.pem -config=conf/pki/ca-config.json -profile=kubernetes conf/pki/kube-controller-manager-csr.json | cfssljson -bare certs/controller/kube-controller-manager

# Generate Kube-Proxy Certs
cfssl gencert -ca=certs/ca/ca.pem -ca-key=certs/ca/ca-key.pem -config=conf/pki/ca-config.json -profile=kubernetes conf/pki/kube-proxy-csr.json | cfssljson -bare certs/proxy/kube-proxy

# Generate Kube-Scheduler Certs
cfssl gencert -ca=certs/ca/ca.pem -ca-key=certs/ca/ca-key.pem -config=conf/pki/ca-config.json -profile=kubernetes conf/pki/kube-scheduler-csr.json | cfssljson -bare certs/scheduler/kube-scheduler

# Generate Kube-API Server Certs
KUBERNETES_HOSTNAMES=kubernetes,kubernetes.default,kubernetes.default.svc,kubernetes.default.svc.cluster,kubernetes.svc.cluster.local
cfssl gencert -ca=certs/ca/ca.pem -ca-key=certs/ca/ca-key.pem -config=conf/pki/ca-config.json -hostname=192.168.1.122,127.0.0.1,10.10.1.10,10.10.1.11,10.10.1.12,${KUBERNETES_HOSTNAMES} -profile=kubernetes conf/pki/kubernetes-csr.json | cfssljson -bare certs/api/kubernetes

# Generate Service Account Certs
cfssl gencert -ca=certs/ca/ca.pem -ca-key=certs/ca/ca-key.pem -config=conf/pki/ca-config.json -profile=kubernetes conf/pki/service-account-csr.json | cfssljson -bare certs/sa/service-account