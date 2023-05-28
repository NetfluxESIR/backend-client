package cmd

import (
	"github.com/NetfluxESIR/backend-client/internal"
	"github.com/spf13/cobra"
)

var (
	statusCmd = &cobra.Command{
		Use:   "status",
		Short: "status the backend",
		Long:  `status the backend`,
		RunE: func(cmd *cobra.Command, args []string) error {
			c := internal.NewClient(cmd.Context(), url, email, password, &role)
			return c.SetStatus(cmd.Context(), videoId, status)
		},
	}
	status string
)

func init() {
	statusCmd.PersistentFlags().StringVarP(&status, "status", "s", "", "status")
}
