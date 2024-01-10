package vesting

import (
	"context"
	"time"

	"cosmossdk.io/math"
	"cosmossdk.io/x/accounts/accountstd"
	vestingtypes "cosmossdk.io/x/accounts/vesting/types/v1"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// Compile-time type assertions
var (
	_ accountstd.Interface = (*PermanentLockedAccount)(nil)
)

// Permernant Vesting Account

// NewPermanentLockedAccount creates a new PermanentLockedAccount object.
func NewPermanentLockedAccount(d accountstd.Dependencies) (*PermanentLockedAccount, error) {
	baseVestingAccount, err := NewBaseVestingAccount(d)

	return &PermanentLockedAccount{baseVestingAccount}, err
}

type PermanentLockedAccount struct {
	*BaseVestingAccount
}

// --------------- Init -----------------

func (plva PermanentLockedAccount) Init(ctx context.Context, msg *vestingtypes.MsgInitVestingAccount) (*vestingtypes.MsgInitVestingAccountResponse, error) {
	resp, err := plva.BaseVestingAccount.Init(ctx, msg)
	plva.EndTime.Set(ctx, math.ZeroInt())

	return resp, err
}

// --------------- execute -----------------

func (plva *PermanentLockedAccount) ExecuteMessages(ctx context.Context, msg *vestingtypes.MsgExecuteMessages) (
	*vestingtypes.MsgExecuteMessagesResponse, error,
) {
	return plva.BaseVestingAccount.ExecuteMessages(ctx, msg, func(_ context.Context, _ time.Time) (sdk.Coins, error) {
		var originalVesting sdk.Coins
		plva.IterateEntries(ctx, plva.OriginalVesting, func(key string, value math.Int) (stop bool) {
			originalVesting = append(originalVesting, sdk.NewCoin(key, value))
			return false
		})
		return originalVesting, nil
	})
}

// --------------- Query -----------------

func (plva PermanentLockedAccount) QueryVestedCoins(ctx context.Context, msg *vestingtypes.QueryVestedCoinsRequest) (
	*vestingtypes.QueryVestedCoinsResponse, error,
) {
	return &vestingtypes.QueryVestedCoinsResponse{
		VestedVesting: nil,
	}, nil
}

func (plva PermanentLockedAccount) QueryVestingCoins(ctx context.Context, msg *vestingtypes.QueryVestingCoinsRequest) (
	*vestingtypes.QueryVestingCoinsResponse, error,
) {
	var originalVesting sdk.Coins
	plva.IterateEntries(ctx, plva.OriginalVesting, func(key string, value math.Int) (stop bool) {
		originalVesting = append(originalVesting, sdk.NewCoin(key, value))
		return false
	})
	return &vestingtypes.QueryVestingCoinsResponse{
		VestingCoins: originalVesting,
	}, nil
}

// Implement smart account interface
func (plva PermanentLockedAccount) RegisterInitHandler(builder *accountstd.InitBuilder) {
	accountstd.RegisterInitHandler(builder, plva.Init)
}

func (plva PermanentLockedAccount) RegisterExecuteHandlers(builder *accountstd.ExecuteBuilder) {
	accountstd.RegisterExecuteHandler(builder, plva.ExecuteMessages)
}

func (plva PermanentLockedAccount) RegisterQueryHandlers(builder *accountstd.QueryBuilder) {
	accountstd.RegisterQueryHandler(builder, plva.QueryVestedCoins)
	accountstd.RegisterQueryHandler(builder, plva.QueryVestingCoins)
	plva.BaseVestingAccount.RegisterQueryHandlers(builder)
}
