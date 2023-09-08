// SPDX-FileCopyrightText: The Cobra authors
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"fmt"

	"github.com/spf13/cobra"
)

func main() {
	root := &cobra.Command{
		Use: "kubectl-cobraplugin",
		Annotations: map[string]string{
			cobra.CommandDisplayNameAnnotation: "kubectl cobraplugin",
		},
	}

	sub := &cobra.Command{
		Use: "subcmd",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("kubectl cobraplugin subcmd")
		},
	}

	root.AddCommand(sub)
	root.Execute()
}
