package level

import (
	"errors"
	"math"

	"github.com/andrewhowdencom/bmctl/client"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// SetLevelCmd represents the set audio level command
var SetLevelCmd = &cobra.Command{
	Use:   "set",
	Short: "Set the audio level for a channel",
	Long:  `Set the audio level for a channel using either the --gain flag or the --normalised flag. If both are provided, gain will be prioritized and level would be ignored according to the API specification.`,
	PreRunE: func(cmd *cobra.Command, args []string) error {
		hasGain := cmd.Flags().Changed("gain")
		hasNorm := cmd.Flags().Changed("normalised")
		
		if !hasGain && !hasNorm {
			return errors.New("must provide either --gain or --normalised")
		}
		return nil
	},
	RunE: DoSetLevel,
}

func DoSetLevel(cmd *cobra.Command, args []string) error {
	c, err := client.New(viper.GetString("api.server"))
	if err != nil {
		return err
	}

	channelIndex, _ := cmd.Flags().GetInt("channel")
	
	req := client.AudioLevel{}

	hasGain := cmd.Flags().Changed("gain")
	hasNorm := cmd.Flags().Changed("normalised")
	
	if hasGain {
		gain, _ := cmd.Flags().GetFloat64("gain")
		req.Gain = &gain
	} else if hasNorm {
		norm, _ := cmd.Flags().GetFloat64("normalised")
		req.Normalised = &norm
	}

	all, _ := cmd.Flags().GetBool("all")
	if all {
		err = c.AudioSetAllLevels(req)
		if err != nil {
			return err
		}
		cmd.Println("Audio level for all channels set successfully")
		return nil
	}

	err = c.AudioSetLevel(channelIndex, req)
	if err != nil {
		return err
	}

	cmd.Printf("Audio level for channel %d set successfully\n", channelIndex)
	return nil
}

func init() {
	SetLevelCmd.Flags().IntP("channel", "c", 0, "The channel index to set the level for")
	SetLevelCmd.Flags().Float64P("gain", "g", math.NaN(), "The gain value in dB")
	SetLevelCmd.Flags().Float64P("normalised", "n", math.NaN(), "The normalised value (0.0 to 1.0)")
	SetLevelCmd.Flags().Bool("all", false, "Set the level for all channels")
}
