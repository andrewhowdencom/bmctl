package phantompower

import (
"github.com/andrewhowdencom/bmctl/cmd/errors"
"github.com/spf13/cobra"
)

// PhantomPowerCmd represents the audio phantom-power command
var PhantomPowerCmd = &cobra.Command{
Use:   "phantom-power",
Short: "Commands associated with managing the audio phantom power",
RunE: func(cmd *cobra.Command, args []string) error {
return errors.ErrNoSubcommand
},
}

func init() {
PhantomPowerCmd.AddCommand(GetPhantomPowerCmd)
PhantomPowerCmd.AddCommand(SetPhantomPowerCmd)
}
