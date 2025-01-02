// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package nestrisdk_test

import (
	"context"
	"os"
	"testing"

	"github.com/stainless-sdks/nestri-sdk-go"
	"github.com/stainless-sdks/nestri-sdk-go/internal/testutil"
	"github.com/stainless-sdks/nestri-sdk-go/option"
)

func TestUsage(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := nestrisdk.NewClient(
		option.WithBaseURL(baseURL),
		option.WithBearerToken("My Bearer Token"),
	)
	machine, err := client.Machines.Get(context.TODO(), "REPLACE_ME")
	if err != nil {
		t.Error(err)
	}
	t.Logf("%+v\n", machine.Data)
}
