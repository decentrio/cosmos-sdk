package simapp

import (
	"fmt"
	"os"

	"github.com/spf13/viper"

	"cosmossdk.io/core/log"
	"cosmossdk.io/core/transaction"

	serverv2 "cosmossdk.io/server/v2"
	networkv2 "github.com/cosmos/cosmos-sdk/testutil/network/v2"
	"github.com/cosmos/cosmos-sdk/types/module/testutil"
)

// NewTestNetworkFixture returns a new simapp AppConstructor for network simulation tests
func NewTestNetworkFixture() networkv2.TestFixture {
	dir, err := os.MkdirTemp("", "simapp")
	if err != nil {
		panic(fmt.Sprintf("failed creating temporary directory: %v", err))
	}
	defer os.RemoveAll(dir)

	v := viper.New()
	v.Set(serverv2.FlagHome, dir)

	app := NewSimApp[transaction.Tx](log.NewNopLogger(), v)

	appCtr := func(val networkv2.ValidatorI) serverv2.AppI[transaction.Tx] {
		return NewSimApp[transaction.Tx](
			val.GetLogger(), viper.New(),
		)
	}

	return networkv2.TestFixture{
		AppConstructor: appCtr,
		GenesisState:   app.DefaultGenesis(),
		EncodingConfig: testutil.TestEncodingConfig{
			InterfaceRegistry: app.interfaceRegistry,
			Codec:             app.AppCodec(),
			TxConfig:          app.TxConfig(),
		},
	}
}
