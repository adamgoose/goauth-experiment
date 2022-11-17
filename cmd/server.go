package cmd

import (
	"fmt"
	"net/http"

	"github.com/adamgoose/goauth-experiment/lib"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Runs a simple API server with both public and auth-protected endpoints",
	RunE: func(cmd *cobra.Command, args []string) error {
		g, err := lib.NewGoauthServer(cmd.Context())
		if err != nil {
			return err
		}

		return http.ListenAndServe(
			fmt.Sprintf(":%d", viper.GetInt("port")),
			g,
		)
	},
}

func init() {
	rootCmd.AddCommand(serverCmd)

	serverCmd.Flags().Int("port", 9090, "Listen port")
	viper.BindPFlag("port", serverCmd.Flags().Lookup("port"))

	serverCmd.Flags().String("jwks-url", "", "JWKs URL")
	viper.BindPFlag("jwks_url", serverCmd.Flags().Lookup("jwks-url"))
}
