// Command quickstart is a compile-checkable example of the openapi-generator
// based client. It exercises the public API surface documented in the README
// and MIGRATION.md; it is built in CI but not run (it would need a live API
// key). Keep it in sync with the README Quick Start.
package main

import (
	"context"
	"fmt"
	"log"

	zernio "github.com/zernio-dev/zernio-go/zernio"
)

func main() {
	// Defaults to the https://zernio.com/api base URL from the spec.
	client := zernio.NewAPIClient(zernio.NewConfiguration())

	// Authenticate with your Bearer API key via the request context.
	ctx := context.WithValue(context.Background(), zernio.ContextAccessToken, "YOUR_API_KEY")

	// List accounts: service field -> method -> Execute, returning the decoded
	// model, the raw *http.Response, and an error.
	accounts, httpResp, err := client.AccountsAPI.ListAccounts(ctx).Execute()
	if err != nil {
		log.Fatalf("list accounts: %v (status %d)", err, httpResp.StatusCode)
	}
	for _, acct := range accounts.GetAccounts() {
		fmt.Printf("account: %+v\n", acct)
	}

	// Create a post: parameters and body are builder methods on the request.
	body := zernio.NewCreatePostRequest()
	post, _, err := client.PostsAPI.CreatePost(ctx).CreatePostRequest(*body).Execute()
	if err != nil {
		log.Fatalf("create post: %v", err)
	}
	fmt.Printf("created post: %+v\n", post)
}
