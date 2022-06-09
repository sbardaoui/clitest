package cmd

import (
	"fmt"
	"context"
	"io"
	"os"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"github.com/spf13/cobra"
)

// dpullCmd represents the dpull command
var dpullCmd = &cobra.Command{
	Use:   "dpull",
	Short: "Pulls a docker image",
	Long: `Pull a docker image that is passed as an argument example Alpine`,
	Run: func(cmd *cobra.Command, args []string) {
	
	fmt.Println("dpull called")
		
	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		panic(err)
	}

	out, err := cli.ImagePull(ctx, args[0], types.ImagePullOptions{})
	if err != nil {
		panic(err)
	}

	defer out.Close()

	io.Copy(os.Stdout, out)
	},
}

func init() {
	rootCmd.AddCommand(dpullCmd)
}
