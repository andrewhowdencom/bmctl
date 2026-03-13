package phantompower

import (
	"fmt"
	"strconv"

	"github.com/andrewhowdencom/bmctl/client"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var SetPhantomPowerCmd = &cobra.Command{
	Use:   "set [true|false]",
	Short: "Set the audio phantom power for a channel",
	Long: `Enables or disables 48V Phantom Power for the specified audio channel.
Must be set to true when using active or condenser microphones via XLR connections that require external power.
WARNING: Ensure this is set to false before hot-plugging sensitive equipment like ribbon mics to prevent electrical damage.
Use the '--all' flag to bulk apply this power boolean to all active channels.`,
	Example: `  bmctl audio phantom-power set true --channel 1
  bmctl audio phantom-power set false -c 2
  bmctl audio phantom-power set true --all`,
	Args:  cobra.ExactArgs(1),
	RunE:  DoSetPhantomPower,
}

func DoSetPhantomPower(cmd *cobra.Command, args []string) error {
	c, err := client.New(viper.GetString("api.server"))
	if err != nil {
		return err
	}

	power, err := strconv.ParseBool(args[0])
	if err != nil {
		return fmt.Errorf("invalid boolean argument: %w", err)
	}

	all, _ := cmd.Flags().GetBool("all")
	if all {
		err = c.AudioSetAllPhantomPower(power)
		if err != nil {
			return err
		}
		cmd.Println("Audio phantom power set for all channels successfully")
		return nil
	}

	channelIndex, _ := cmd.Flags().GetInt("channel")
	err = c.AudioSetPhantomPower(channelIndex, power)
	if err != nil {
		return err
	}

	cmd.Printf("Audio phantom power for channel %d set successfully\n", channelIndex)
	return nil
}

func init() {
	SetPhantomPowerCmd.Flags().IntP("channel", "c", 0, "The channel index to set the phantom power for")
	SetPhantomPowerCmd.Flags().Bool("all", false, "Set the phantom power for all channels")
}
