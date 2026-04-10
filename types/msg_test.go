package types

import (
	"testing"

	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

func testAccAddr(t *testing.T) string {
	t.Helper()
	priv := secp256k1.GenPrivKey()
	return sdk.AccAddress(priv.PubKey().Address()).String()
}

func TestMsgConvertNFT_ValidateBasic(t *testing.T) {
	t.Parallel()
	sender := testAccAddr(t)
	valid := MsgConvertNFT{
		CosmosSender: sender,
		EvmReceiver:  "0x1234567890123456789012345678901234567890",
	}
	require.NoError(t, valid.ValidateBasic())

	invalidSender := valid
	invalidSender.CosmosSender = "not-bech32"
	require.Error(t, invalidSender.ValidateBasic())

	invalidRecv := valid
	invalidRecv.EvmReceiver = "0x123"
	require.Error(t, invalidRecv.ValidateBasic())
}

func TestMsgConvertERC721_ValidateBasic(t *testing.T) {
	t.Parallel()
	sender := testAccAddr(t)
	receiver := testAccAddr(t)
	valid := MsgConvertERC721{
		EvmContractAddress: "0x1234567890123456789012345678901234567890",
		CosmosReceiver:     receiver,
		CosmosSender:       sender,
	}
	require.NoError(t, valid.ValidateBasic())

	badContract := valid
	badContract.EvmContractAddress = "nope"
	require.Error(t, badContract.ValidateBasic())

	badRecv := valid
	badRecv.CosmosReceiver = "xx"
	require.Error(t, badRecv.ValidateBasic())
}

func TestMsgTransferERC721_ValidateBasic(t *testing.T) {
	t.Parallel()
	sender := testAccAddr(t)
	valid := MsgTransferERC721{
		EvmContractAddress: "0x1234567890123456789012345678901234567890",
		CosmosSender:       sender,
	}
	require.NoError(t, valid.ValidateBasic())

	badContract := valid
	badContract.EvmContractAddress = ""
	require.Error(t, badContract.ValidateBasic())

	badSender := valid
	badSender.CosmosSender = ""
	require.Error(t, badSender.ValidateBasic())
}

func TestMsgRoutesAndTypes(t *testing.T) {
	t.Parallel()
	var (
		m0 MsgConvertNFT
		m1 MsgConvertERC721
		m2 MsgTransferERC721
	)
	require.Equal(t, RouterKey, m0.Route())
	require.Equal(t, TypeMsgConvertNFT, m0.Type())
	require.Equal(t, TypeMsgConvertERC721, m1.Type())
	require.Equal(t, TypeMsgTransferERC721, m2.Type())
}
