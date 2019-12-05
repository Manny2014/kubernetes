
VERSION=$(shell date +"%y.%m.%d.%H.%M.%S")
NAME="knative-consumer-producer"

build:
	$(MAKE) docker-do

docker-do:
	docker build --no-cache -t manny87/${NAME}:${VERSION} .
	docker push manny87/${NAME}:${VERSION}
