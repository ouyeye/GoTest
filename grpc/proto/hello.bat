@echo off
protoc --proto_path=. --go_out=. --go-grpc_out=. hello.proto
echo Proto files generated successfully.
pause