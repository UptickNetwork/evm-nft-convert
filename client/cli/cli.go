package cli

import (
	"github.com/UptickNetwork/evm-nft-convert/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/spf13/cobra"
)

// GetQueryCmd returns the parent command for all erc721 CLI query commands.
func GetQueryCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "Querying commands for the erc721 module",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(
		GetTokenPairsCmd(),
		GetTokenPairCmd(),
		GetParamsCmd(),
		GetEvmAddressFromIBCCmd(),
	)
	return cmd
}

// NewTxCmd returns a root CLI command handler for erc721 transaction commands
func NewTxCmd() *cobra.Command {
	txCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "erc721 subcommands",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	txCmd.AddCommand(
		NewConvertNFTCmd(),
		NewConvertERC721Cmd(),
	)
	return txCmd
}
