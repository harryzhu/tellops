package cmd

import (
	"os"

	"github.com/harryzhu/util"
	"github.com/spf13/cobra"
)

var (
	SMTPHost     string
	SMTPPort     string
	SMTPStartTLS string
	AccessKey    string
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "tellme via email",
	Short: "tellme gossip | encrypt",
	Long: `
	general send-mail tool.
	env vars: 
	TELLOPSSMTPHOST="", 
	TELLOPSSMTPPORT="", 
	TELLOPSSMTPSTARTTLS="yes" or "no", 
	TELLOPSACCESSKEY=""`,
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
	SMTPHost = util.GetEnv("TELLOPSSMTPHOST", "smtp.office365.com")
	SMTPPort = util.GetEnv("TELLOPSSMTPPORT", "587")
	SMTPStartTLS = util.GetEnv("TELLOPSSMTPSTARTTLS", "yes")
	AccessKey = util.GetEnv("TELLOPSACCESSKEY", "kxooRczff0V1L8C81uBrcg==")
}
