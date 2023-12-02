package solve

import (
	"fmt"
	"github.com/kkcaZ/advent-2024/pkg/days"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

type Options struct {
	day  int
	part int
}

func NewCmd() *cobra.Command {
	o := Options{}

	cmd := &cobra.Command{
		Use:   "solve",
		Short: "Solve an advent of code day",
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
	cmd.Flags().IntVarP(&o.day, "day", "d", 1, "Day to run")
	cmd.Flags().IntVarP(&o.part, "part", "p", 1, "Part of the day to run")
}

func (o *Options) Validate() error {
	if o.part < 1 || o.part > 2 {
		return fmt.Errorf("part must be 1 or 2")
	}

	return nil
}

func (o *Options) Run() error {
	err := o.Validate()
	if err != nil {
		return errors.Wrap(err, "failed to validate options")
	}

	dayService, err := days.NewDayService(o.day)
	if err != nil {
		return errors.Wrap(err, "failed whilst initialising day service")
	}

	if o.part == 1 {
		err := dayService.PartOne()
		if err != nil {
			return errors.Wrap(err, "failed whilst running part one")
		}
	} else if o.part == 2 {
		err := dayService.PartTwo()
		if err != nil {
			return errors.Wrap(err, "failed whilst running part two")
		}
	}

	return nil
}
