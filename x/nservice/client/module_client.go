package client

import (
	"github.com/cosmos/cosmos-sdk/client"
	nameservicecmd "github.com/kidinamoto01/nameservice/x/nservice/client/cli"
	"github.com/spf13/cobra"
	amino "github.com/tendermint/go-amino"
)

// ModuleClient exports all client functionality from this module
type ModuleClient struct {
	storeKey string
	cdc      *amino.Codec
}

func NewModuleClient(storeKey string, cdc *amino.Codec) ModuleClient {
	return ModuleClient{storeKey, cdc}
}

// GetQueryCmd returns the cli query commands for this module
func (mc ModuleClient) GetQueryCmd() *cobra.Command {
	// Group nameservice queries under a subcommand
	nsQueryCmd := &cobra.Command{
		Use:   "nameservice",
		Short: "Querying commands for the nameservice module",
	}

	nsQueryCmd.AddCommand(client.GetCommands(
		nameservicecmd.GetCmdResolveName(mc.storeKey, mc.cdc),
		nameservicecmd.GetCmdWhois(mc.storeKey, mc.cdc),
	)...)

	return nsQueryCmd
}

// GetTxCmd returns the transaction commands for this module
func (mc ModuleClient) GetTxCmd() *cobra.Command {
	nsTxCmd := &cobra.Command{
		Use:   "nameservice",
		Short: "Nameservice transactions subcommands",
	}

	nsTxCmd.AddCommand(client.PostCommands(
		nameservicecmd.GetCmdBuyName(mc.cdc),
		nameservicecmd.GetCmdSetName(mc.cdc),
	)...)

	return nsTxCmd
}