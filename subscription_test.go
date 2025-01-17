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
	"github.com/nestrilabs/nestri-go-sdk/shared"
)

func TestSubscriptionNewWithOptionalParams(t *testing.T) {
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
		ID:         nestri.F("0bfcb712-df13-4454-81a8-fbee66eddca4"),
		Frequency:  nestri.F(nestri.SubscriptionNewParamsFrequencyFixed),
		Next:       nestri.F[nestri.SubscriptionNewParamsNextUnion](shared.UnionString("2025-01-09T01:56:23.902Z")),
		ProductID:  nestri.F("0bfcb712-df43-4454-81a8-fbee66eddca4"),
		Quantity:   nestri.F(int64(1)),
		CanceledAt: nestri.F[nestri.SubscriptionNewParamsCanceledAtUnion](shared.UnionString("2025-02-09T01:56:23.902Z")),
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
