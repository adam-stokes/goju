package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"strings"
)

func main(){
	var cmdPrint = &cobra.Command{
		Use: "print [string]",
		Short: "Print anything",
		Long: "stfu",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Print:" + strings.Join(args, ""))},
	}
	var rootCmd = &cobra.Command{Use: "goju"}
	rootCmd.AddCommand(cmdPrint)
	rootCmd.Execute()
}
