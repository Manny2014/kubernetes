
build:
	$(MAKE) docker-build
	$(MAKE) docker-push

docker-build:
	docker build --no-cache -t manny87/knative-go:${VERSION} .

docker-push:
	docker push manny87/knative-go:${VERSION}