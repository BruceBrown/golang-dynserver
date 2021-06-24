# golang-dynserver
POC of server with dynamic load of go shared libraries

build and run via:
```
cd server
go build

cd ../plugins/test
go build -buildmode=plugin

cd ../other
go build -buildmode=plugin

cd ../../server
./server
```

The server.json file configures the .so to load.