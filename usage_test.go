// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package nestri_test

import (
	"context"
	"os"
	"testing"

	"github.com/nestrilabs/nestri-go-sdk"
	"github.com/nestrilabs/nestri-go-sdk/internal/testutil"
	"github.com/nestrilabs/nestri-go-sdk/option"
)

func TestUsage(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := nestri.NewClient(
		option.WithBaseURL(baseURL),
		option.WithBearerToken("My Bearer Token"),
	)
	game, err := client.Games.Get(context.TODO(), 870780.000000)
	if err != nil {
		t.Error(err)
	}
	t.Logf("%+v\n", game.Data)
}
