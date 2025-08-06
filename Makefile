IMAGE ?= snowzach/dns-noerror:latest

default: build

buildx-setup:
	docker run --rm --privileged multiarch/qemu-user-static --reset -p yes
	docker buildx create --name crossplat --use

.PHONY: build # Build the container image
build:
	docker buildx build \
		--output "type=docker,push=false" \
		--tag $(IMAGE) \
		.

.PHONY: publish # Push the image to the remote registry
publish:
	docker buildx build --platform linux/amd64,linux/arm64,linux/arm/v7 -t $(IMAGE) --push .

