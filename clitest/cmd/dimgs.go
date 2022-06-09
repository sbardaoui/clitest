package cmd

import (
	"fmt"
	"context"

	
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"github.com/spf13/cobra"
)

var dimgsCmd = &cobra.Command{
	Use:   "dimgs",
	Short: "Lists all images",
	Long: `Lists all available images in your current engine`,
	Run: func(cmd *cobra.Command, args []string) {
	
	fmt.Println("dimgs called")

	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		panic(err)
	}

	images, err := cli.ImageList(ctx, types.ImageListOptions{})
	if err != nil {
		panic(err)
	}

	for _, image := range images {
		fmt.Println(image.ID)
		fmt.Println(image)
	}
	},
}

func init() {
	rootCmd.AddCommand(dimgsCmd)

	}
