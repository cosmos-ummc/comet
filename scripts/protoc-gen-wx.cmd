protoc --proto_path=api/proto --proto_path=$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis --proto_path=$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway --go_out=plugins=grpc:pkg/api cosmos-service.proto
protoc --proto_path=api/proto --proto_path=$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis --proto_path=$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway --grpc-gateway_out=logtostderr=true:pkg/api cosmos-service.proto
protoc --proto_path=api/proto --proto_path=$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis --proto_path=$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway --swagger_out=logtostderr=true:api/swagger cosmos-service.proto
spectacle ./api/swagger/specs.swagger.json -t ./api/swagger/public
