module github.com/packethost/hegel

go 1.13

require (
	github.com/golang/protobuf v1.3.5
	github.com/grpc-ecosystem/go-grpc-middleware v1.2.0
	github.com/grpc-ecosystem/go-grpc-prometheus v1.2.0
	github.com/packethost/cacher v0.0.0-20200319200613-5dc1cac4fd33
	github.com/packethost/pkg v0.0.0-20190715213007-7c3a64b4b5e3
	github.com/pkg/errors v0.9.1
	github.com/prometheus/client_golang v1.5.1
	github.com/tinkerbell/tink v0.0.0-20200527081417-d9d6b637de27
	go.mongodb.org/mongo-driver v1.1.2
	golang.org/x/net v0.0.0-20200324143707-d3edc9973b7e
	google.golang.org/grpc v1.28.0
)

replace github.com/tinkerbell/tink v0.0.0-20200527081417-d9d6b637de27 => github.com/kdeng3849/tink v0.0.0-20200527231431-56a0d8e7117e
