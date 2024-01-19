package vesting

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	vestingtypes "cosmossdk.io/x/vesting/types/v1"
)

func validateAmount(amount sdk.Coins) error {
	if !amount.IsValid() {
		return sdkerrors.ErrInvalidCoins.Wrap(amount.String())
	}

	if !amount.IsAllPositive() {
		return sdkerrors.ErrInvalidCoins.Wrap(amount.String())
	}

	return nil
}

func validateOwnerVestingAccountRequest(request *vestingtypes.QueryOwnerVestingAccountRequest) error {
	if request.Owner == "" {
		return sdkerrors.ErrInvalidRequest.Wrap("owner must not be empty")
	}

	if request.VestingType == "" {
		return sdkerrors.ErrInvalidRequest.Wrap("vesting type must not be empty")
	}

	if _, ok := vestingTypeMap[request.VestingType]; !ok {
		return sdkerrors.ErrInvalidRequest.Wrapf("vesting type is invalid. Should be the following: %s, %s, %s or %s",
			CONTINUOS_VESTING_ACCOUNT,
			PERIODIC_VESTING_ACCOUNT,
			DELAYED_VESTING_ACCOUNT,
			PERMERNANT_VESTING_ACCOUNT,
		)
	}

	return nil
}
