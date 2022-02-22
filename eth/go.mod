module github.com/hyperbench/hyperbench-plugins/eth

go 1.15

require (
	github.com/btcsuite/btcd v0.21.0-beta // indirect
	github.com/ethereum/go-ethereum v1.10.9
	github.com/jackpal/go-nat-pmp v1.0.2 // indirect
	github.com/hyperbench/hyperbench-common v0.0.0-20220128060413-2f16193421b0
	github.com/spf13/cast v1.4.1
	github.com/spf13/viper v1.10.1
	github.com/stretchr/testify v1.7.0
        golang.org/x/sys v0.0.0-20220209214540-3681064d5158 // indirect
)

replace google.golang.org/grpc => google.golang.org/grpc v1.29.1
