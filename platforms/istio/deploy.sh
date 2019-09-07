

helm template helm/istio-init istio-system -f istio-values.yaml --namespace=istio-system | kubectl apply -f -

helm template helm/istio istio-system -f istio-values.yaml --namespace=istio-system | kubectl apply -f -


kubectl apply -f policies/mesh.yaml