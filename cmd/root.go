/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	b64 "encoding/base64"
	"fmt"
	"math/rand/v2"
	"os"

	guiApp "github.com/MartinJindra/terry/gui"
	"github.com/MartinJindra/terry/quotes"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "terry",
	Short: "Get quotes from computer genius Terry A. Davis",
	Long:  `Get real quotes from computer genius and creator of TempleOS Terry A. Davis.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {

		cat, _ := cmd.Flags().GetString("category")
		big, _ := cmd.Flags().GetBool("big")
		gui, _ := cmd.Flags().GetBool("gui")

		if gui {
			guiApp.Execute()
			return
		}

		quoteList := quotes.GetQuotes()

		if cat != "" && !catValid(cat) {
			fmt.Println("Category:", cat, "doesn't exist")
			return
		}

		if catValid(cat) {
			var filtered []quotes.Quote
			for _, quote := range quoteList {
				if quote.Category == cat {
					filtered = append(filtered, quote)
				}
			}
			quoteList = filtered
		}

		if len(quoteList) == 0 {
			fmt.Println("No quotes found")
			return
		}

		index := rand.Int() % len(quoteList)

		fmt.Print("\"")
		fmt.Print(quoteList[index].Text)
		fmt.Println("\"")

		if !big {
			fmt.Println("― Terry A. Davis")
			return
		}
		fmt.Println(beautify())
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().StringP("category", "c", "", "Filter quotes by category")
	rootCmd.Flags().BoolP("big", "b", false, "Write Terry A. Davis' Name in big ASCII art")
	rootCmd.Flags().BoolP("gui", "g", false, "Open the Terry A. Davis GUI")
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.terry.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	//rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	cobra.OnInitialize()
}

// Helper function:
func catValid(cat string) bool {
	if cat == "crazy" || cat == "cia" || cat == "confused" || cat == "racist" || cat == "deep" || cat == "technical" {
		return true
	}
	return false
}

func beautify() string {
	const enc = "ICBfX19fX19fICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgIF9fX19fICAgICAgICAgICAgICBfICAgICAKIHxfXyAgIF9ffCAgICAgICAgICAgICAgICAgICAgICAgL1wgICAgICAgfCAgX18gXCAgICAgICAgICAgIChfKSAgICAKICAgIHwgfCBfX18gXyBfXyBfIF9fIF8gICBfICAgICAvICBcICAgICAgfCB8ICB8IHwgX18gX19fICAgX19fIF9fXyAKICAgIHwgfC8gXyBcICdfX3wgJ19ffCB8IHwgfCAgIC8gL1wgXCAgICAgfCB8ICB8IHwvIF9gIFwgXCAvIC8gLyBfX3wKICAgIHwgfCAgX18vIHwgIHwgfCAgfCB8X3wgfCAgLyBfX19fIFwgXyAgfCB8X198IHwgKF98IHxcIFYgL3wgXF9fIFwKICAgIHxffFxfX198X3wgIHxffCAgIFxfXywgfCAvXy8gICAgXF8oXykgfF9fX19fLyBcX18sX3wgXF8vIHxffF9fXy8KICAgICAgICAgICAgICAgICAgICAgICBfXy8gfCAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAKICAgICAgICAgICAgICAgICAgICAgIHxfX18vICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICA="
	dec, _ := b64.StdEncoding.DecodeString(enc)
	return string(dec)
}
