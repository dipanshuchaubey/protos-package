FILES = $(shell find *_service -name "*proto")
SERVICE = $(if $(service),$(service),all)

# Replace _ with - in service name
BUILD_NAME = $(echo $$SERVICE | tr '_' '-')

.PHONY: protos
protos:
	protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
	$(FILES)

python-protos:
	python -m grpc_tools.protoc -I. --python_out=. \
	--pyi_out=. --grpc_python_out=. \
	$(FILES)