glide install
go build -v  ./cmd/stats-codegen && stats-codegen.exe
go build -v -o fesl-backend.exe ./cmd/fesl-backend
go build -v ./cmd/tpl-bindata && tpl-bindata.exe
go build -v -o magma.exe ./cmd/heroes-api
start fesl-backend.exe && start magma.exe


