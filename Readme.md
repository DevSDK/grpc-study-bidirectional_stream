# GRPC-Bidirectional stream example

Bidirectional stream client-server example.

Server response a sum value through the stream if the sum of the received number is multiple of 3.

The client sends a number using stream with opening the receiver stream in a goroutine.



```
    ./compileIDL.sh
    go run server/main.go
    go run client/main.go
```

