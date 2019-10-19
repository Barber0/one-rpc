PROTOC_DIR		:=	tools/protoc-gen-go
CORE_PB_DIR		:= 	protocol/proto
CORE_PB_DEST	:= 	protocol/res

.PHONY:gen-core-proto,buildprotoc

buildprotoc:
	cd $(PROTOC_DIR) && go install

gen-core-proto:
	./pb2go.sh $(CORE_PB_DIR) $(CORE_PB_DEST)