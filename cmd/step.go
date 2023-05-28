package cmd

import (
	"github.com/NetfluxESIR/backend-client/internal"
	"github.com/spf13/cobra"
)

var (
	stepCmd = &cobra.Command{
		Use:   "step",
		Short: "notify the backend of a step",
		Long:  `notify the backend of a step`,
		RunE: func(cmd *cobra.Command, args []string) error {
			c := internal.NewClient(cmd.Context(), url, email, password, &role)
			if previousStep != "" {
				err := c.SetSteps(cmd.Context(), videoId, previousStep, previousStepStatus, previousStepLog)
				if err != nil {
					return err
				}
			}
			if currentStep != "" {
				return c.SetSteps(cmd.Context(), videoId, currentStep, currentStepStatus, "")
			}
			return nil
		},
	}
	previousStep       string
	previousStepLog    string
	previousStepStatus string
	currentStep        string
	currentStepStatus  string
)

func init() {
	stepCmd.PersistentFlags().StringVarP(&previousStep, "previous-step", "P", "", "previous step")
	stepCmd.PersistentFlags().StringVarP(&currentStep, "current-step", "c", "", "current step")
	stepCmd.PersistentFlags().StringVarP(&previousStepLog, "previous-step-log", "l", "", "previous step log")
	stepCmd.PersistentFlags().StringVarP(&previousStepStatus, "previous-step-status", "s", "", "previous step status")
	stepCmd.PersistentFlags().StringVarP(&currentStepStatus, "current-step-status", "S", "", "current step status")
}
