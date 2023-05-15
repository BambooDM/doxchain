package cli

import (
	"strconv"

	"github.com/be-heroes/doxchain/x/did/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdCreateDidDocument() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "did-document [did-document-json]",
		Short: "Broadcast message DidDocument",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			didDocument := types.DidDocument{}
			clientCtx, err := client.GetClientTxContext(cmd)

			if err != nil {
				return err
			}

			err = clientCtx.Codec.UnmarshalJSON([]byte(args[0]), &didDocument)

			if err != nil {
				return err
			}

			msg := types.NewMsgCreateDidDocumentRequest(
				clientCtx.GetFromAddress().String(),
				didDocument,
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