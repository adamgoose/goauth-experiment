# Go Auth Experiment

Welcome to the super simple Golang Authentication Experiment!

In this experiment, you'll write Golang code to fetch an Access Token from an
OAuth2/OIDC Server. Once you've fetched an access token, you'll be able to run
a local API server which will verify your token.

> It's important to note that this is a flow typically executed from a browser.
> It can also be implemented from a CLI like this, but it's awkward to have to
> open a browser tab. You'll recognize this behavior from `tsh login` and
> `vault login -method=oidc`.

Read the instructions below fully before beginning. It'll be useful for you to
understand the entire process before diving in.

1. Clone the repo, install Go, and make sure you can run the program with `go run .`
1. Copy the provided `.envrc` file and make sure the values are loaded with `direnv allow`
1. Open the `cmd/client.go` file. Notice that it's your job to populate `accessToken`
1. Use the `golang.org/x/oauth2` library to generate an `AuthCodeURL`
  - Use `http://localhost:9000/oauth/callback` as your `RedirectURL` value
  - Use the `github.com/coreos/go-oidc` library to fetch `Endpoint` value
  - Access config values with `viper.GetString("issuer_url")` for example
1. Run an HTTP server on `localhost:9000`. Upon receiving a request, extract the `code` Query Parameter from the Request URL
1. Print your Auth Code URL to the console so that you can open it in your browser
1. Login, and you'll be redirected back to your HTTP server on `localhost:9000` with a code
1. Once your HTTP server receives a request and you have a code, use the `golang.org/x/oauth2` library again to `Exchange` the code for an access token

Use `go run . client` to iteratively test your changes. You can `return nil`
anywhere in the command to bail early. You can use `spew.Dump(something)` to
spit out a debug representation of a variable.

Once you have a valid `accessToken`, pass it to `TestAccessToken`. To test it,
you'll have to first run a server.

- In one shell, run `go run . server`
- In another, run `go run . client`

## OAuth2 Flow Diagram

![](https://miro.medium.com/max/1400/1*ULF38OTiNJNQZ4lHQZqRwQ.png)

1. This is equivalent to you choosing to execute `go run . client`
2. This is you clicking on the `AuthCodeURL` that your code prints out
3. This is your `localhost:9000` HTTP server extracting the `code`
4. Authorization is implicit in our use case.
5. Authorization is implicit in our use case.
6. This happens inside Authentic during the `Exchange`
7. This is you `Exchange`ing the `code` for an access token
8. This is you `Exchange`ing the `code` for an access token

A. This happens in `/cmd/client.go:TestAccesToken`
B. We don't actually do this. We just trust the IdP's JWKs (public keys)
C. We don't actually do this. We just trust the IdP's JWKs (public keys)
D. This happens in `/lib/auth.go:Middleware`

## Server API Reference

When you run `go run . server`, the fake API server will be available at
http://localhost:9090.

- `GET /` - Public - `{"challenge": "string"}`
- `GET /api/secret` - Protected - `{"prize":"string","recipient":"string"}`
