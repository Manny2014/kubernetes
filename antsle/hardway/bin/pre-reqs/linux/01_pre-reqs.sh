# K8s installation pre-reqs

yum install wget -y

## PKI required tools
wget -q --show-progress --https-only --timestamping \
  https://storage.googleapis.com/kubernetes-the-hard-way/cfssl/linux/cfssl \
  https://storage.googleapis.com/kubernetes-the-hard-way/cfssl/linux/cfssljson


chmod +x cfssl cfssljson
mv cfssl cfssljson /usr/local/bin/

## Install kubectl
wget https://storage.googleapis.com/kubernetes-release/release/v1.15.3/bin/linux/amd64/kubectl
chmod +x kubectl
mv kubectl /usr/local/bin/


