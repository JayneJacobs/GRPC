# GRPC

## [Installation Mac](https://grpc.io/docs/quickstart/go/)

√ install PROTOC program
√ the protoc compiler is used to generate gRPC service code
√ take the executable from PROTOC, store it some location, have that location on your PATH environment variable
√ echo $PATH
√ for example:
√ mv protoc /usr/local/bin
√ import `http://github.com/golang/protobuf/protoc-gen-go`
√ using go modules
√ create a file [like this in your project](tools.go)

```sh
mkdir grpc
git init
git add .; git commit -m "first commit"; git remote add origin https://github.com/JayneJacobs/GRPC.git;git push -u origin master
git add .; git commit -m "first commit"; git remote add origin https://github.com/JayneJacobs/GRPC.git;git push -u origin master
git init github.com/JayneJacobs/GRPC
go mod init
go run .
go mod tidy
```

√ Set GOBIN in bash_profile


```GOBIN="/usr/local/go1.12.9.darwin-amd64/bin"```

√ in project directory:

```go install github.com/golang/protobuf/protoc-gen-go```

## Protobuffers

Create Protobuf
Do not reuse numbers even if the item is deleted.

DSL Domain Specific Language.
Helps you write code automatically


IDL Interface Definition Language. (DSL)

## Compiling service

Convert protofile into a go package. 
protoc one level above the proto file

protoc -I echo echo/echo.proto --go_out=echo

protoc -I echo echo/echo.proto --go_out=plugins=grpc:echo
