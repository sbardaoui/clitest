/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// drunCmd represents the drun command
var drunCmd = &cobra.Command{
	Use:   "drun",
	Short: "Quickly prompt a hello world !",
	Long: `Runs a prompt hello world for testing purposes `,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("drun called")
	},
}

func init() {
	rootCmd.AddCommand(drunCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// drunCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// drunCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
