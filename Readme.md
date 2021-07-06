# GPRC-Bidirectional stream example

Bidirectional stream client-server example.

Server response a summary value through the stream if the summary of received number is multiple of 3.

Client send a number using stream with opening the receiver stream in goroutine.



```
    ./compileIDL.sh
    go run server/main.go
    go run client/main.go
```

