protoc Math.proto --proto_path="IDL" \
                  --go-grpc_out='gen/math' \
                  --go_out='gen/math' \
                  --go_opt=paths=source_relative \
                  --go-grpc_opt=paths=source_relative 
