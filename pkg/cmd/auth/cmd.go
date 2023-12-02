package auth

import (
	"fmt"
	"github.com/kkcaZ/advent-2024/pkg/util"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

type Options struct {
	session string
}

func NewCmd() *cobra.Command {
	o := Options{}

	cmd := &cobra.Command{
		Use:   "auth",
		Short: "Sets up authentication for input retrieval",
		Run: func(cmd *cobra.Command, args []string) {
			err := o.Run()
			if err != nil {
				fmt.Println(errors.Wrap(err, "failed to run command"))
			}
		},
	}

	o.AddFlags(cmd)
	return cmd
}

func (o *Options) AddFlags(cmd *cobra.Command) {
	cmd.Flags().StringVarP(&o.session, "session", "s", "", "Session cookie for adventofcode.com")
	_ = cmd.MarkFlagRequired("session")
}

func (o *Options) Run() error {
	err := util.StoreToken(o.session)
	if err != nil {
		return errors.Wrap(err, "failed to store token")
	}

	return nil
}
