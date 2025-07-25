package keeper

import (
	sdkerrors "cosmossdk.io/errors"
	"github.com/UptickNetwork/evm-nft-convert/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// GetPair checks that:
//   - the global parameter for erc721 conversion is enabled
//   - minting is enabled for the given (erc721,nft) token pair
func (k Keeper) GetPair(
	ctx sdk.Context,
	token string,
) (types.TokenPair, error) {

	id := k.GetTokenPairID(ctx, token)
	if len(id) == 0 {

		return types.TokenPair{}, sdkerrors.Wrapf(
			types.ErrTokenPairNotFound, "token '%s' not registered by id", token,
		)
	}

	pair, found := k.GetTokenPair(ctx, id)
	if !found {
		return types.TokenPair{}, sdkerrors.Wrapf(
			types.ErrTokenPairNotFound, "token '%s' not registered", token,
		)
	}

	return pair, nil
}
