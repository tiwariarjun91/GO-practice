package main

import(
	"time"
	"github.com/spf13/cobra"
)

func printTimeCmd() *cobra.Command{
	return &cobra.Command{
		Use : "curtime",
		RunE : func(cmd *cobra.Command, args []string) error{
			now := time.Now()
			newtime := now.Format(time.RubyDate)
			cmd.Println("Hi Arjun the current time is :-", newtime)
			return nil
		},
	}
}