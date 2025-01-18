// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package nestri_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/nestrilabs/nestri-go-sdk"
	"github.com/nestrilabs/nestri-go-sdk/internal/testutil"
	"github.com/nestrilabs/nestri-go-sdk/option"
)

func TestSubscriptionNew(t *testing.T) {
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
	_, err := client.Subscriptions.New(context.TODO(), nestri.SubscriptionNewParams{
		CheckoutID: nestri.F("0bfcb712-df13-4454-81a8-fbee66eddca4"),
	})
	if err != nil {
		var apierr *nestri.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestSubscriptionList(t *testing.T) {
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
	_, err := client.Subscriptions.List(context.TODO())
	if err != nil {
		var apierr *nestri.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestSubscriptionDelete(t *testing.T) {
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
	_, err := client.Subscriptions.Delete(context.TODO(), "0bfcb712-df13-4454-81a8-fbee66eddca4")
	if err != nil {
		var apierr *nestri.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
