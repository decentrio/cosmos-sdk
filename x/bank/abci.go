package bank

import (
	"github.com/cosmos/cosmos-sdk/telemetry"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/bank/keeper"
	"github.com/cosmos/cosmos-sdk/x/bank/types"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
)

// EndBlocker called every block, process transfer coins from EVM dead account to gov.
// Then burn these coins from gov module.
func EndBlocker(ctx sdk.Context, keeper keeper.Keeper) error {
	defer telemetry.ModuleMeasureSince(types.ModuleName, telemetry.Now(), telemetry.MetricKeyEndBlocker)

	deadBalances := keeper.GetAllBalances(ctx, types.EvmDeadAddr)
	err := keeper.SendCoinsFromAccountToModule(ctx, types.EvmDeadAddr, govtypes.ModuleName, deadBalances)
	if err != nil {
		return err
	}

	err = keeper.BurnCoins(ctx, govtypes.ModuleName, deadBalances)
	if err != nil {
		return err
	}
	return nil
}
