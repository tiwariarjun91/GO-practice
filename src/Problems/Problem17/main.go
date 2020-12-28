package main

import(
	"os"
	"github.com/spf13/cobra"
	// command to get the cobra package go get -u -v github.com/spf13/cobra/cobra
	
)


func main(){
	cmd := &cobra.Command{
		Use: "gifm",
		Short: "Hello Arjun",
		SilenceUsage: true, 
	} 

	cmd.AddCommand(printTimeCmd())

	if err:= cmd.Execute(); err != nil{
		os.Exit(1)
	}
}