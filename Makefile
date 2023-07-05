
protoc        := $(shell which protoc)
proto_dir     := ./proto
proto_files   = $(wildcard ${proto_dir}/*.proto)
proto_out     := ${GOPATH}/src

.PHONY: gen-network
gen-network:
	${protoc} --proto_path=. --go_opt=paths=import --go_out=${proto_out} --go-grpc_out=${proto_out} ${proto_files}