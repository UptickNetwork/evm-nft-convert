package cli

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/common"
	ethermint "github.com/evmos/ethermint/types"

	"github.com/UptickNetwork/evm-nft-convert/types"
)

// NewConvertNFTCmd returns a CLI command handler for converting a Cosmos coin
func NewConvertNFTCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "convert-nft [class_id] [cosmos_token_ids] [evm_contract_address] [evm_token_ids] [receiver_hex]",
		Short: "Convert a Cosmos nft to erc721. When the receiver [optional] is omitted, the erc721 tokens are transferred to the sender.",
		Args:  cobra.RangeArgs(4, 5),
		RunE: func(cmd *cobra.Command, args []string) error {

			cliCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			classId := args[0]
			if len(classId) == 0 {
				return fmt.Errorf("classId can not be empty")
			}

			cosmosTokenIds := strings.Split(args[1], ",")
			if len(cosmosTokenIds) == 0 {
				return fmt.Errorf("nftID can not be empty")
			}

			evmContractAddress := args[2]
			evmTokenIds := strings.Split(args[3], ",")

			var evmReceiver string
			cosmosSender := cliCtx.GetFromAddress()
			if len(args) == 5 {
				evmReceiver = args[4]
				if err := ethermint.ValidateAddress(evmReceiver); err != nil {
					return fmt.Errorf("invalid receiver hex address %w", err)
				}
			} else {
				evmReceiver = common.BytesToAddress(cosmosSender).Hex()
			}

			msg := &types.MsgConvertNFT{
				EvmContractAddress: evmContractAddress,
				CosmosTokenIds:     cosmosTokenIds,
				ClassId:            classId,
				EvmTokenIds:        evmTokenIds,
				EvmReceiver:        evmReceiver,
				CosmosSender:       cosmosSender.String(),
			}

			if err := msg.ValidateBasic(); err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(cliCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)
	return cmd
}

// NewConvertERC721Cmd returns a CLI command handler for converting an erc721
func NewConvertERC721Cmd() *cobra.Command {
	cmd := &cobra.Command{
		Use: "convert-erc721 [evm_contract_address] [evm_token_ids] [class_id] [cosmos_token_ids] [cosmos_receiver]",
		Short: "Convert an erc721 token to Cosmos coin.  " +
			"When the receiver [optional] is omitted, the Cosmos coins are transferred to the sender.",
		Args: cobra.RangeArgs(4, 5),
		RunE: func(cmd *cobra.Command, args []string) error {

			cliCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			evmContractAddress := args[0]
			if err := ethermint.ValidateAddress(evmContractAddress); err != nil {
				return fmt.Errorf("invalid erc721 contract address %w", err)
			}

			evmTokenIds := strings.Split(args[1], ",")
			if len(evmTokenIds) == 0 {
				return fmt.Errorf("tokenID can not be empty")
			}

			evmSender := common.BytesToAddress(cliCtx.GetFromAddress().Bytes())

			classId := args[2]
			cosmosTokenIds := strings.Split(args[3], ",")

			cosmosReceiver := cliCtx.GetFromAddress()
			if len(args) == 5 {
				cosmosReceiver, err = sdk.AccAddressFromBech32(args[4])
				if err != nil {
					return err
				}
			}

			msg := &types.MsgConvertERC721{
				EvmContractAddress: evmContractAddress,
				EvmTokenIds:        evmTokenIds,
				EvmSender:          evmSender.Hex(),
				CosmosReceiver:     cosmosReceiver.String(),
				ClassId:            classId,
				CosmosTokenIds:     cosmosTokenIds,
			}

			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(cliCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)
	return cmd
}
