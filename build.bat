go build -v -o ./stats.exe ./cmd/stats-codegen
./stats.exe
./glide install
go build -v -o ./fesl-backend.exe ./cmd/fesl-backend
go build -v -o ./xml.exe ./cmd/tpl-bindata
./xml.exe
go build -v -o ./magma.exe ./cmd/heroes-api
./fesl-backend.exe 
./magma.exe