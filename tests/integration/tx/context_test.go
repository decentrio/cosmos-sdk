package tx

import (
	"testing"

	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/proto"

	"cosmossdk.io/depinject"
	"cosmossdk.io/log"
	_ "cosmossdk.io/x/accounts"
	"cosmossdk.io/x/tx/signing"

	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/tests/integration/tx/internal/pulsar/testpb"
	"github.com/cosmos/cosmos-sdk/testutil/configurator"
	simtestutil "github.com/cosmos/cosmos-sdk/testutil/sims"
)

func ProvideCustomGetSigners() signing.CustomGetSigner {
	return signing.CustomGetSigner{
		MsgType: proto.MessageName(&testpb.TestRepeatedFields{}),
		Fn: func(msg proto.Message) ([][]byte, error) {
			testMsg := msg.(*testpb.TestRepeatedFields)
			// arbitrary logic
			signer := testMsg.NullableDontOmitempty[1].Value
			return [][]byte{[]byte(signer)}, nil
		},
	}
}

func TestDefineCustomGetSigners(t *testing.T) {
	var interfaceRegistry codectypes.InterfaceRegistry
	_, err := simtestutil.SetupAtGenesis(
		depinject.Configs(
			configurator.NewAppConfig(
				configurator.AccountsModule(),
				configurator.AuthModule(),
				configurator.StakingModule(),
				configurator.BankModule(),
				configurator.ConsensusModule(),
			),
			depinject.Supply(log.NewNopLogger()),
			depinject.Provide(ProvideCustomGetSigners),
		),
		&interfaceRegistry,
	)
	require.NoError(t, err)
	require.NotNil(t, interfaceRegistry)

	msg := &testpb.TestRepeatedFields{
		NullableDontOmitempty: []*testpb.Streng{
			{Value: "foo"},
			{Value: "bar"},
		},
	}
	signers, err := interfaceRegistry.SigningContext().GetSigners(msg)
	require.NoError(t, err)
	require.Equal(t, [][]byte{[]byte("bar")}, signers)

	// Provide no CustomGetSigners. Consequently, validation will fail and depinject will return an error
	require.PanicsWithError(t, "no cosmos.msg.v1.signer option found for message testpb.TestRepeatedFields; use DefineCustomGetSigners to specify a custom getter", func() {
		_, _ = simtestutil.SetupAtGenesis(
			depinject.Configs(
				configurator.NewAppConfig(
					configurator.AuthModule(),
					configurator.StakingModule(),
					configurator.BankModule(),
					configurator.ConsensusModule(),
				),
				depinject.Supply(log.NewNopLogger()),
			),
			&interfaceRegistry,
		)
	})
}
