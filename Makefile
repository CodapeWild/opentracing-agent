## Prerequisites for compile this repo from source
# install golang
# install protoc
# intall go-swagger

protoc        := $(shell which protoc)
proto_version := v1
proto_dir     := ./idl/proto/${proto_version}
proto_files   = $(wildcard ${proto_dir}/*.proto)
proto_out     := ${GOPATH}/src

.PHONY: gen-network
gen-network:
	${protoc} ${proto_files} --proto_path=. --go_opt=paths=import --go_out=${proto_out} --go-grpc_out=${proto_out}
	@echo generate network ok