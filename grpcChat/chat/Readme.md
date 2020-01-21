# Chat Applicaiton using GRPC

1. Create the .protofile

## Compiling service

Convert protofile into a go package. 
protoc one level above the proto file

protoc -I chat chat/chat.proto --go_out=chat

protoc -I chat chat/chat.proto --go_out=plugins=grpc:chat
