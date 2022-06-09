package cmd

import (
	"fmt"
	"context"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"github.com/spf13/cobra"
)

var killallCmd = &cobra.Command{
	Use:   "killall",
	Short: "Kills all running containers",
	Long: `this kills all running containers using their ID`,
	Run: func(cmd *cobra.Command, args []string) {
	
	fmt.Println("killall called")
	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		panic(err)
	}

	containers, err := cli.ContainerList(ctx, types.ContainerListOptions{})
	if err != nil {
		panic(err)
	}

	for _, container := range containers {
		fmt.Print("Stopping container ", container.ID[:10], "... ")
		if err := cli.ContainerStop(ctx, container.ID, nil); err != nil {
			panic(err)
		}
		fmt.Println("Success")
	}
	},
}

func init() {
	rootCmd.AddCommand(killallCmd)
}
