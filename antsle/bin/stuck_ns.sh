NAMESPACE=knative-monitoring

curl -k -H "Content-Type: application/json" -X PUT --data-binary @ns.json http://127.0.0.1:8001/api/v1/namespaces/${NAMESPACE}/finalize