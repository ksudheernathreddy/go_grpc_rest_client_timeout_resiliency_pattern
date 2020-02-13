#! /bin/bash
echo "Installing prerequisites for building application binary..."
go get -u "github.com/golang/protobuf/protoc-gen-go"
go get -u "github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway"
go get -u "github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger"
export PATH=$PATH:/usr/local/include/bin
if [[ ! -f "/usr/local/include/bin/protoc" ]]
then
 PROTOC_ZIP=protoc-3.7.1-linux-x86_64.zip
  curl -OL https://github.com/protocolbuffers/protobuf/releases/download/v3.7.1/$PROTOC_ZIP
   unzip -o $PROTOC_ZIP -d /usr/local/include
    rm -f $PROTOC_ZIP
    else
     echo "Protoc already present, its version: $(protoc --version)"
     fi
     echo "Building proto files..."
     protoc -I${GOPATH}/src/github.com/interfaces -I${GOPATH}/pkg/mod/github.com/grpc-ecosystem/grpc-gateway\@v1.12.1/third_party/googleapis --go_out=plugins=grpc:${GOPATH}/src test.proto
     protoc -I${GOPATH}/src/github.com/interfaces -I${GOPATH}/src -I${GOPATH}/pkg/mod/github.com/grpc-ecosystem/grpc-gateway\@v1.12.1/third_party/googleapis --grpc-gateway_out=logtostderr=true:${GOPATH}/src  test.proto
     protoc -I${GOPATH}/src/github.com/interfaces -I${GOPATH}/src -I${GOPATH}/pkg/mod/github.com/grpc-ecosystem/grpc-gateway\@v1.12.1/third_party/googleapis --swagger_out=logtostderr=true:${GOPATH}/src/github.com/interfaces test.proto
