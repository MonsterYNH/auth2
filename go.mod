module github.com/MonsterYNH/auth2

go 1.16

require (
	github.com/MonsterYNH/api v1.0.0
	github.com/MonsterYNH/athena v1.0.1
	google.golang.org/grpc v1.38.0
	gorm.io/gorm v1.21.11
)

replace (
	github.com/coreos/bbolt => go.etcd.io/bbolt v1.3.5
	github.com/golang/protobuf => github.com/golang/protobuf v1.3.2
	google.golang.org/grpc => google.golang.org/grpc v1.27.0
)
