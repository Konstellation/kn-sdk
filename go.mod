module github.com/konstellation/kn-sdk

go 1.15

require (
	github.com/cosmos/cosmos-sdk v0.41.0
	github.com/ethereum/go-ethereum v1.9.11
	github.com/golang/protobuf v1.4.3
	github.com/google/uuid v1.1.2
	github.com/gorilla/mux v1.8.0
	github.com/mitchellh/mapstructure v1.1.2
	github.com/rs/cors v1.7.0 // indirect
	github.com/spf13/cast v1.3.1
	github.com/spf13/cobra v1.1.1
	github.com/spf13/viper v1.7.1
	github.com/tendermint/tendermint v0.34.3
	github.com/tendermint/tm-db v0.6.3
)

replace github.com/gogo/protobuf => github.com/regen-network/protobuf v1.3.3-alpha.regen.1

replace google.golang.org/grpc => google.golang.org/grpc v1.33.2