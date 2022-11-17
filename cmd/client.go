package cmd

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/davecgh/go-spew/spew"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var clientCmd = &cobra.Command{
	Use:   "client",
	Short: "Initiates a login flow, obtains an access token, and tests it against the protected endpoint of the API",
	RunE: func(cmd *cobra.Command, args []string) error {

		// Insert your code here!

		// TODO: Obtain an Access Token!
		accessToken := ""

		resp, err := TestAccessToken(accessToken)
		if err != nil {
			return err
		}
		defer resp.Body.Close()

		spew.Dump(resp.Status)
		io.Copy(os.Stdout, resp.Body)

		return nil
	},
}

// TestAccessToken assumes a locally running server, and calls an auth-protected
// endpoint using the given access token.
func TestAccessToken(accessToken string) (*http.Response, error) {
	req, err := http.NewRequest(
		http.MethodGet,
		fmt.Sprintf("http://localhost:%d/api/secret", viper.GetInt("port")),
		nil,
	)
	if err != nil {
		return nil, errors.Wrap(err, "unable to create scenario request")
	}

	req.Header.Set(
		"Authorization",
		fmt.Sprintf("Bearer %s", accessToken),
	)

	return http.DefaultClient.Do(req)
}

func init() {
	rootCmd.AddCommand(clientCmd)

	clientCmd.Flags().String("issuer-url", "", "OAuth2 Issuer URL")
	viper.BindPFlag("issuer_url", clientCmd.Flags().Lookup("issuer-url"))

	clientCmd.Flags().String("client-id", "", "OAuth2 Client ID")
	viper.BindPFlag("client_id", clientCmd.Flags().Lookup("client-id"))

	clientCmd.Flags().String("client-secret", "", "OAuth2 Client Secret")
	viper.BindPFlag("client_secret", clientCmd.Flags().Lookup("client-secret"))
}
