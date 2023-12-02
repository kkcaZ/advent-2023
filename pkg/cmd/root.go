package cmd

import (
	"github.com/kkcaZ/advent-2024/pkg/cmd/auth"
	"github.com/kkcaZ/advent-2024/pkg/cmd/solve"
	"github.com/spf13/cobra"
	"log/slog"
)

func Main() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "advent",
		Short: "Advent of Code 2024",
		Run: func(cmd *cobra.Command, args []string) {
			err := cmd.Help()
			if err != nil {
				slog.Error(err.Error())
			}
		},
	}
	cmd.AddCommand(solve.NewCmd())
	cmd.AddCommand(auth.NewCmd())
	return cmd
}
