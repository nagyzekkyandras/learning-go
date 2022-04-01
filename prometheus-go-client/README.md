# Prometheus Go Client

### Dependencies
```sh
go mod download github.com/prometheus/client_golang
go get github.com/prometheus/client_golang/prometheus@v1.5.1
go get github.com/prometheus/common/log@v0.9.1
go get -u golang.org/x/sys
```

### Build
```sh
go build main.go 
```

### Run
```sh
./main
```

### Check
```sh
curl http://localhost:5000/metrics
```