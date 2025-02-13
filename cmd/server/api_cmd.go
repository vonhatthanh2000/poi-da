package cmd

import (
	"layerg-poi-da/internal/initialize"

	"github.com/spf13/cobra"
)

func startApi(cmd *cobra.Command, args []string) {
	initialize.Run()
}
