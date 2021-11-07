.PHONY: default help

default: help
help: ## display make targets
	@echo  "\n \033[36m make all -> for quick start\n\033[0m"
	@echo  "\n \033[36m make clean -> to reset\n\033[0m"
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(word 1, $(MAKEFILE_LIST)) | \
		awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m make %-20s -> %s\n\033[0m", $$1, $$2}'


.PHONY: all
all: certs start-rabbitmq

.PHONY: clean
clean: stop-rabbitmq delete-kind


.PHONY: start-rabbitmq
start-rabbitmq: ## start rabbitmq
	@bash -c "docker run -d \
  -it \
  --name rabbitmq \
  --mount type=bind,source=$(shell pwd)/config,target=/etc/rabbitmq/conf.d \
  -p 15672:15672 -p 5671:5671 \
  rabbitmq:3-management"


.PHONY: stop-rabbitmq
stop-rabbitmq: ## start rabbitmq
	@bash -c "docker stop rabbitmq"
	@bash -c "docker rm rabbitmq"



.PHONY: certs
certs: ## Create certs
	@bash -c "rm -rf tls-gen config/certs etc/certs"
	@bash -c "git clone https://github.com/michaelklishin/tls-gen tls-gen"
	@bash -c "cd tls-gen/basic && make PASSWORD=bunnies"
	@bash -c "cd tls-gen/basic/result && openssl rsa -in client_key.pem -out key.unencrypted.pem -passin pass:bunnies"
	@bash -c "mv tls-gen/basic/result config/certs"
	@bash -c "mkdir -p etc"
	@bash -c "cp -r config/certs etc/certs"
	@bash -c "sudo chown 999:999 config/certs/*"

.PHONY: kind
kind: ## setup local kind cluster 	
	@bash -c "./infra/local/scripts/create-kind-config.sh"
	@bash -c "kind create cluster --name rabbit --config infra/local/scripts/kind-config-with-mounts-and-ingress.yaml"
	@bash -c "echo -n 'k port-forward  deployment/rabbitmq  5671:5671'"


.PHONY: delete-kind
delete-kind: ## delete kind cluster
	@bash -c "kind delete cluster --name rabbit"


