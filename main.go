// SPDX-FileCopyrightText: The Cobra authors
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"fmt"

	"github.com/spf13/cobra"
)

func main() {
	rootCmd := &cobra.Command{
		Use: "kubectl-cobraplugin",
		Annotations: map[string]string{
			cobra.CommandDisplayNameAnnotation: "kubectl cobraplugin",
		},
	}
	subCmd := &cobra.Command{
		Use: "subcmd",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("kubectl cobraplugin subcmd")
		},
	}
	rootCmd.AddCommand(subCmd)
	rootCmd.Execute()
}
