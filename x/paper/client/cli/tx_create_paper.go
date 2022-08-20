package cli

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cobra"
	"paper/x/paper/types"
)

var _ = strconv.Itoa(0)

func CmdCreatePaper() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-paper [host] [paper-name] [owner] [price]",
		Short: "Broadcast message createPaper",
		Args:  cobra.ExactArgs(4),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argHost := args[0]
			argPaperName := args[1]
			argOwner := args[2]
			argPrice := args[3]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgCreatePaper(
				clientCtx.GetFromAddress().String(),
				argHost,
				argPaperName,
				argOwner,
				argPrice,
			)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
