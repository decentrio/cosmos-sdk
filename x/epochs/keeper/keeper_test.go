package keeper_test

import (
	"testing"
	"time"

	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/testutil"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"

	"github.com/stretchr/testify/suite"

	"cosmossdk.io/core/appmodule"
	"cosmossdk.io/core/header"
	"cosmossdk.io/log"
	storetypes "cosmossdk.io/store/types"
	epochskeeper "cosmossdk.io/x/epochs/keeper"
	"cosmossdk.io/x/epochs/types"
	codectestutil "github.com/cosmos/cosmos-sdk/codec/testutil"
	"github.com/cosmos/cosmos-sdk/runtime"
	moduletestutil "github.com/cosmos/cosmos-sdk/types/module/testutil"
)

type KeeperTestSuite struct {
	suite.Suite
	Ctx          sdk.Context
	environment  appmodule.Environment
	EpochsKeeper epochskeeper.Keeper
	queryClient  types.QueryClient
}

func (s *KeeperTestSuite) SetupTest() {
	ctx, epochsKeeper, environment := Setup(s.T())

	s.Ctx = ctx
	s.EpochsKeeper = epochsKeeper
	s.environment = environment
	queryRouter := baseapp.NewGRPCQueryRouter()
	cfg := module.NewConfigurator(nil, nil, queryRouter)
	types.RegisterQueryServer(cfg.QueryServer(), epochskeeper.NewQuerier(s.EpochsKeeper))
}

func Setup(t *testing.T) (sdk.Context, epochskeeper.Keeper, appmodule.Environment) {
	key := storetypes.NewKVStoreKey(types.StoreKey)
	storeService := runtime.NewKVStoreService(key)
	environment := runtime.NewEnvironment(storeService, log.NewNopLogger())
	testCtx := testutil.DefaultContextWithDB(t, key, storetypes.NewTransientStoreKey("transient_test"))
	ctx := testCtx.Ctx.WithHeaderInfo(header.Info{Time: time.Now()})
	encCfg := moduletestutil.MakeTestEncodingConfig(codectestutil.CodecOptions{})

	epochsKeeper := epochskeeper.NewKeeper(
		environment,
		encCfg.Codec,
	)
	epochsKeeper = epochsKeeper.SetHooks(types.NewMultiEpochHooks())
	ctx.WithHeaderInfo(header.Info{Height: 1, Time: time.Now().UTC(), ChainID: "epochs"})
	epochsKeeper.InitGenesis(ctx, *types.DefaultGenesis())
	SetEpochStartTime(ctx, epochsKeeper)

	return ctx, epochsKeeper, environment
}

func TestKeeperTestSuite(t *testing.T) {
	suite.Run(t, new(KeeperTestSuite))
}

func SetEpochStartTime(ctx sdk.Context, epochsKeeper epochskeeper.Keeper) {
	for _, epoch := range epochsKeeper.AllEpochInfos(ctx) {
		epoch.StartTime = ctx.BlockTime()
		epochsKeeper.DeleteEpochInfo(ctx, epoch.Identifier)
		err := epochsKeeper.AddEpochInfo(ctx, epoch)
		if err != nil {
			panic(err)
		}
	}
}
