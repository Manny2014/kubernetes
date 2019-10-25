
build:
	$(MAKE) docker-build
	$(MAKE) docker-push

docker-build-push:
	$(MAKE) docker-build
	$(MAKE) docker-push

docker-build:
	docker build -t manny87/knative-py:${VERSION} .

docker-push:
	docker push manny87/knative-py:${VERSION}