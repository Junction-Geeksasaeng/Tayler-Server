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

func CmdUpdatePaper() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update-paper [id] [new-owner] [new-price]",
		Short: "Broadcast message updatePaper",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argId := args[0]
			argNewOwner := args[1]
			argNewPrice := args[2]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgUpdatePaper(
				clientCtx.GetFromAddress().String(),
				argId,
				argNewOwner,
				argNewPrice,
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
