package cmd

import (
	"fmt"
	"context"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"github.com/spf13/cobra"
)

var dpsCmd = &cobra.Command{
	Use:   "dps",
	Short: "Lists running containers",
	Long: `Lists the ID and image name of all running containers`,
	Run: func(cmd *cobra.Command, args []string) {
	
	fmt.Println("dps called")

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
		fmt.Println(container.ID)
		fmt.Println(container.Image)
	}
	},
}

func init() {
	rootCmd.AddCommand(dpsCmd)
}
