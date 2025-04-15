package keeper_test

import (
	"testing"

	"cosmossdk.io/depinject"
	"cosmossdk.io/log"

	"cosmossdk.io/x/feegrant"
	"cosmossdk.io/x/feegrant/keeper"
	"github.com/cosmos/cosmos-sdk/testutil/configurator"
	simtestutil "github.com/cosmos/cosmos-sdk/testutil/sims"
	sdk "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/cosmos/cosmos-sdk/x/auth"           // import as blank for app wiring
	_ "github.com/cosmos/cosmos-sdk/x/auth/tx/config" // import as blank for app wiring
	_ "github.com/cosmos/cosmos-sdk/x/auth/vesting"   // import as blank for app wiring
	_ "github.com/cosmos/cosmos-sdk/x/bank"           // import as blank for app wiring
	_ "github.com/cosmos/cosmos-sdk/x/consensus"      // import as blank for app wiring
	_ "github.com/cosmos/cosmos-sdk/x/genutil"        // import as blank for app wiring
	_ "github.com/cosmos/cosmos-sdk/x/params"         // import as blank for app wiring
	_ "github.com/cosmos/cosmos-sdk/x/staking"        // import as blank for app wiring
	"github.com/stretchr/testify/require"
)

var appConfig = configurator.NewAppConfig(
	configurator.FeegrantModule(),
	configurator.AuthModule(),
	configurator.BankModule(),
	configurator.VestingModule(),
	configurator.StakingModule(),
	configurator.TxModule(),
	configurator.ConsensusModule(),
	configurator.ParamsModule(),
	configurator.GenutilModule(),
)

func BenchmarkGetAllowance(b *testing.B) {
	b.ReportAllocs()
	var keeper keeper.Keeper
	app, err := simtestutil.Setup(
		depinject.Configs(
			depinject.Supply(log.NewNopLogger()),
			appConfig,
		),
		&keeper,
	)
	require.NoError(b, err)

	ctx := app.NewContext(false)
	granter := sdk.AccAddress([]byte("granter________________"))
	grantee := sdk.AccAddress([]byte("grantee________________"))

	now := ctx.BlockTime()
	if err := keeper.GrantAllowance(ctx, granter, grantee, &feegrant.BasicAllowance{
		SpendLimit: sdk.NewCoins(sdk.NewInt64Coin("atom", 555)),
		Expiration: &now,
	}); err != nil {
		b.Fatal(err)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := keeper.GetAllowance(ctx, granter, grantee)
		if err != nil {
			b.Fatal(err)
		}
	}
}
