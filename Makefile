.PHONY: build, serve, serve-cfg, linkedlist, generic-linkedlist

build:
	go build ./...

serve:
	go run ./... serve --port 80

serve-cfg:
	go run ./... serve --config ./configuration.yaml

linkedlist:
	go run ./... linkedlist

generic-linkedlist:
	go run ./... genericLinkedlist