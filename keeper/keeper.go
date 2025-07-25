package keeper

import (
	"cosmossdk.io/log"
	storetypes "cosmossdk.io/store/types"
	"fmt"
	nftkeeper "github.com/UptickNetwork/uptick/x/collection/keeper"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	porttypes "github.com/cosmos/ibc-go/v8/modules/core/05-port/types"

	"github.com/UptickNetwork/evm-nft-convert/types"
	ibcnfttransferkeeper "github.com/bianjieai/nft-transfer/keeper"
	ibcnfttransfertypes "github.com/bianjieai/nft-transfer/types"
)

// Keeper of this module maintains collections of erc721.
type Keeper struct {
	storeKey   storetypes.StoreKey
	cdc        codec.BinaryCodec
	paramstore paramtypes.Subspace

	accountKeeper types.AccountKeeper
	nftKeeper     nftkeeper.Keeper
	evmKeeper     types.EVMKeeper
	ics4Wrapper   porttypes.ICS4Wrapper
	ibcKeeper     ibcnfttransferkeeper.Keeper
}

// NewKeeper creates new instances of the erc721 Keeper
func NewKeeper(storeKey storetypes.StoreKey,
	cdc codec.BinaryCodec,
	ps paramtypes.Subspace,
	ak types.AccountKeeper,
	nk nftkeeper.Keeper,
	ek types.EVMKeeper,
	ik ibcnfttransferkeeper.Keeper,
) Keeper {
	// set KeyTable if it has not already been set
	if !ps.HasKeyTable() {
		ps = ps.WithKeyTable(types.ParamKeyTable())
	}

	return Keeper{
		storeKey:      storeKey,
		cdc:           cdc,
		paramstore:    ps,
		accountKeeper: ak,
		nftKeeper:     nk,
		evmKeeper:     ek,
		ibcKeeper:     ik,
	}
}

// Logger returns a module-specific logger.
func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}

// SetICS4Wrapper sets the ICS4 wrapper to the keeper.
// It panics if already set
func (k *Keeper) SetICS4Wrapper(ics4Wrapper porttypes.ICS4Wrapper) {
	if k.ics4Wrapper != nil {
		panic("ICS4 wrapper already set")
	}

	k.ics4Wrapper = ics4Wrapper
}

func (k *Keeper) GetVoucherClassID(port string, channel string, classId string) string {
	// since SendPacket did not prefix the classID, we must prefix classID here
	classPrefix := ibcnfttransfertypes.GetClassPrefix(port, channel)
	// NOTE: sourcePrefix contains the trailing "/"
	prefixedClassID := classPrefix + classId

	// construct the class trace from the full raw classID
	classTrace := ibcnfttransfertypes.ParseClassTrace(prefixedClassID)
	voucherClassID := classTrace.IBCClassID()

	return voucherClassID
}
