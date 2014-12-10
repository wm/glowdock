// Basic OAuth2 command line helper.
// This code is based on the example found at https://code.google.com/p/goauth2/source/browse/oauth/example/oauthreq.go
// See LICENSE-GOOGLE

package auth

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"syscall"

	"code.google.com/p/goauth2/oauth"
)

var (
	scope       = "flow"
	redirectURL = "urn:ietf:wg:oauth:2.0:oob"
	authURL     = "https://api.flowdock.com/oauth/authorize"
	tokenURL    = "https://api.flowdock.com/oauth/token"
	home, _     = syscall.Getenv("HOME")
	cachefile   = home + "/.godock.json"
	code        = flag.String("code", "", "Authorization Code")
)

func AuthenticationRequest(clientSecret string, clientId string) *http.Client {
	flag.Parse()

	// Set up a configuration.
	config := &oauth.Config{
		ClientId:     clientId,
		ClientSecret: clientSecret,
		RedirectURL:  redirectURL,
		Scope:        scope,
		AuthURL:      authURL,
		TokenURL:     tokenURL,
		TokenCache:   oauth.CacheFile(cachefile),
	}

	// Set up a Transport using the config.
	transport := &oauth.Transport{Config: config}

	// Try to pull the token from the cache; if this fails, we need to get one.
	token, err := config.TokenCache.Token()
	if err != nil {
		if *code == "" {
			// Get an authorization code from the data provider.
			// ("Please ask the user if I can access this resource.")
			url := config.AuthCodeURL("")
			fmt.Println("Visit this URL to get a code, then run again with -code=YOUR_CODE\n")
			fmt.Println(url)
			fmt.Printf("\nThis is a one time request - your credentials will be stored in %s\n", cachefile)
			os.Exit(0)
		}
		// Exchange the authorization code for an access token.
		// ("Here's the code you gave the user, now give me a token!")
		token, err = transport.Exchange(*code)
		if err != nil {
			log.Fatal("Exchange:", err)
		}
		// (The Exchange method will automatically cache the token.)
		fmt.Printf("Token is cached in %v\n", config.TokenCache)
	}

	// Make the actual request using the cached token to authenticate.
	// ("Here's the token, let me in!")
	transport.Token = token
	client := transport.Client()

	return client
}
