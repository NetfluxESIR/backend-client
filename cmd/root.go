package cmd

import (
	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:   "backend-client",
		Short: "backend-client is a client for the backend",
		Long:  `backend-client is a client for the backend`,
		RunE: func(cmd *cobra.Command, args []string) error {
			return nil
		},
	}
	url      string
	email    string
	password string
	role     string
	videoId  string
)

func init() {
	rootCmd.PersistentFlags().StringVarP(&url, "url", "U", "", "url")
	rootCmd.PersistentFlags().StringVarP(&email, "email", "e", "", "email")
	rootCmd.PersistentFlags().StringVarP(&password, "password", "p", "", "password")
	rootCmd.PersistentFlags().StringVarP(&role, "role", "r", "", "role")
	rootCmd.PersistentFlags().StringVarP(&videoId, "id", "i", "", "video id")
	rootCmd.AddCommand(
		stepCmd,
		statusCmd,
	)
}

func Execute() error {
	return rootCmd.Execute()
}
