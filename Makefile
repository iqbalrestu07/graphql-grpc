PACKAGE = $(shell head -1 go.mod | awk '{print $$2}')
SERVICE_NAME = services/$@
CHECK_DIR_CMD = test -d ${SERVICE_NAME} || (echo "Directory ${SERVICE_NAME} doesn't exist" && false)
RM_F_CMD = rm -f
RM_RF_CMD = ${RM_F_CMD} -r
HELP_CMD = grep -E '^[a-zA-Z_-]+:.*?\#\# .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?\#\# "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

.DEFAULT_GOAL := help
.PHONY: post all test clean api help

services := epul

all: $(services) ## Generate Pbs and build

$(services): 
	@${CHECK_DIR_CMD}
	echo "package :${PACKAGE}"
	echo "service name: ${SERVICE_NAME}"
	echo "service : ${service}"
	protoc -I${SERVICE_NAME}/proto --experimental_allow_proto3_optional --go_out=paths=source_relative:./${SERVICE_NAME}/grpc --go-grpc_out=paths=source_relative:./${SERVICE_NAME}/grpc ${SERVICE_NAME}/proto/*.proto

clean: ## Clean proto and bin "with service=<serviceName>"
	test ${service} || (echo "service must be defined" && false)
	${RM_RF_CMD} bin/${service}
	${RM_RF_CMD} services/${service}/proto/*.pb.go


help: ## Show this help
	@${HELP_CMD}