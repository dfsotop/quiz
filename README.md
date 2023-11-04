# quizz
a simple quizz build using grpc

## generatig binding files
### pre-required tools
Install the necessary tools:

- Protocol Buffers compiler (protoc)
- Go plugin for Protocol Buffers (protoc-gen-go)
- Go plugin for gRPC (protoc-gen-go-grpc)

On a terminal, situated at the project's root folder, execute:

```go
protoc --go_out=. --go-grpc_out=.  proto/quiz.proto
```