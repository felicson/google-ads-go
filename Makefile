ADS_VERSION=v9
PROTO_ROOT_DIR=googleapis/
PROTO_CORE_DIR=protobuf/src
PROTO_SRC_DIR=/google/ads/googleads/$(ADS_VERSION)/
PROTO_OUT_DIR:=$(shell mktemp -d)
PKG_PATH=paths=source_relative
PROTOC_GO_ARGS=--go_out=$(PROTO_OUT_DIR) --go_opt=paths=source_relative --go-grpc_out=$(PROTO_OUT_DIR) --go-grpc_opt=paths=source_relative
GITHUB_USER=felicson

MATCH=google.golang.org/genproto/googleapis/ads/googleads/$(ADS_VERSION)/
REPLACE=github.com/$(GITHUB_USER)/google-ads-go/

ENTRY=cmd/main.go
BIN=gads

build:
	go build -o $$GOPATH/bin/$(BIN) $(ENTRY)

run:
	go run $(ENTRY)

run-debug:
	GODEBUG=http2debug=2 GRPC_GO_LOG_SEVERITY_LEVEL=info GRPC_GO_LOG_VERBOSITY_LEVEL=2 go run $(ENTRY)

test:
	go test -v -cover ./...

.SILENT protos: clean-protos clean-gen-protos
	echo "converting protos for version $(ADS_VERSION)"
	for file in $(PROTO_ROOT_DIR)$(PROTO_SRC_DIR)/**/*.proto; do \
		echo "converting proto $$(basename $$file)"; \
		sed -i "s|experiment_arm\.proto|experiment_arm0.proto|g" $$file; \
		protoc -I /usr/include --proto_path=$(PROTO_ROOT_DIR) --proto_path=$(PROTO_CORE_DIR) $(PROTOC_GO_ARGS) $$file; \
	done; \
	for file in ${PROTO_OUT_DIR}$(PROTO_SRC_DIR)/**/*.pb.go; do \
		sed -i "s|$(MATCH)|$(REPLACE)|g" $$file; \
	done; \
	mv $(PROTO_OUT_DIR)$(PROTO_SRC_DIR)/* ./
	rm -rf ${PROTO_OUT_DIR}
	echo "built proto files to $$(basename $(PROTO_OUT_DIR))"

clean-protos:
	rm -rf common/
	rm -rf enums/
	rm -rf errors/
	rm -rf resources/
	rm -rf services/

clean-gen-protos:
	rm -rf google/

clone: clone-googleapis clone-protobuf
clone-googleapis:
	cd $(PROTO_ROOT_DIR)
	git submodule update --init --recursive

clone-protocore:
	cd $(PROTO_CORE_DIR)
	git submodule update --init --recursive

update-googleapis:
	cd $(PROTO_ROOT_DIR)
	git submodule update --recursive --remote

update-protocore:
	cd $(PROTO_CORE_DIR)
	git submodule update --recursive --remote

.PHONY: protos clone-googleapis update-googleapis
