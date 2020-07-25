protoc --proto_path=api/proto --proto_path=third_party --go_out=plugins=grpc:pkg/api cosmos-service.proto
protoc --proto_path=api/proto --proto_path=third_party --grpc-gateway_out=logtostderr=true:pkg/api cosmos-service.proto
protoc --proto_path=api/proto --proto_path=third_party --swagger_out=logtostderr=true:api/swagger cosmos-service.proto
spectacle .\api\swagger\specs.swagger.json -t .\api\swagger\public