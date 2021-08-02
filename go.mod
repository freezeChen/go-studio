module github.com/freezeChen/go-studio

go 1.16

require (
	github.com/coreos/bbolt v1.3.2 // indirect
	github.com/coreos/etcd v3.3.13+incompatible // indirect
	github.com/coreos/go-systemd v0.0.0-20190321100706-95778dfbb74e // indirect
	github.com/coreos/pkg v0.0.0-20180928190104-399ea9e2e55f // indirect
	github.com/gin-gonic/gin v1.7.2
	github.com/go-logfmt/logfmt v0.4.0 // indirect
	github.com/go-sql-driver/mysql v1.6.0
	github.com/golang/protobuf v1.5.2
	github.com/google/uuid v1.2.0
	github.com/gorilla/websocket v1.4.2 // indirect
	github.com/kr/pretty v0.2.1 // indirect
	github.com/leodido/go-urn v1.2.1 // indirect
	github.com/prometheus/tsdb v0.7.1 // indirect
	github.com/spf13/viper v1.8.1
	github.com/tmc/grpc-websocket-proxy v0.0.0-20190109142713-0ad062ec5ee5 // indirect
	github.com/ugorji/go v1.2.4 // indirect
	github.com/urfave/cli v1.20.0
	go.etcd.io/etcd v0.0.0-20200402134248-51bdeb39e698
	golang.org/x/crypto v0.0.0-20210314154223-e6e6c4f2bb5b // indirect
	golang.org/x/sync v0.0.0-20210220032951-036812b2e83c
	google.golang.org/grpc v1.38.0
	gopkg.in/yaml.v2 v2.4.0 // indirect
	xorm.io/xorm v1.1.2
)

replace xorm.io/xorm v1.1.2 => ./core/xorm
