module k8s-project

go 1.13

require (
	github.com/aws/aws-sdk-go v1.25.27
	github.com/go-logr/logr v0.1.0
	github.com/hashicorp/consul/api v1.2.0
	github.com/mdempsky/gocode v0.0.0-20190203001940-7fb65232883f // indirect
	github.com/onsi/ginkgo v1.6.0
	github.com/onsi/gomega v1.4.2
	github.com/prometheus/client_golang v0.9.0
	golang.org/x/tools v0.0.0-20191105231337-689d0f08e67a // indirect
	k8s.io/api v0.0.0-20190409021203-6e4e0e4f393b
	k8s.io/apimachinery v0.0.0-20190404173353-6a84e37a896d
	k8s.io/client-go v11.0.1-0.20190409021438-1a26190bd76a+incompatible
	sigs.k8s.io/controller-runtime v0.2.2
)
