// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package nestri

import (
	"context"
	"net/http"
	"reflect"

	"github.com/nestrilabs/nestri-go-sdk/internal/apijson"
	"github.com/nestrilabs/nestri-go-sdk/internal/param"
	"github.com/nestrilabs/nestri-go-sdk/internal/requestconfig"
	"github.com/nestrilabs/nestri-go-sdk/option"
	"github.com/nestrilabs/nestri-go-sdk/shared"
	"github.com/tidwall/gjson"
)

// SubscriptionService contains methods and other services that help with
// interacting with the nestri API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSubscriptionService] method instead.
type SubscriptionService struct {
	Options []option.RequestOption
}

// NewSubscriptionService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewSubscriptionService(opts ...option.RequestOption) (r *SubscriptionService) {
	r = &SubscriptionService{}
	r.Options = opts
	return
}

// Create a subscription for the current user.
func (r *SubscriptionService) New(ctx context.Context, body SubscriptionNewParams, opts ...option.RequestOption) (res *SubscriptionNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "subscriptions"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// List the subscriptions associated with the current user.
func (r *SubscriptionService) List(ctx context.Context, opts ...option.RequestOption) (res *SubscriptionListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "subscriptions"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type SubscriptionNewResponse struct {
	Data SubscriptionNewResponseData `json:"data,required"`
	JSON subscriptionNewResponseJSON `json:"-"`
}

// subscriptionNewResponseJSON contains the JSON metadata for the struct
// [SubscriptionNewResponse]
type subscriptionNewResponseJSON struct {
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SubscriptionNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionNewResponseJSON) RawJSON() string {
	return r.raw
}

type SubscriptionNewResponseData string

const (
	SubscriptionNewResponseDataOk SubscriptionNewResponseData = "ok"
)

func (r SubscriptionNewResponseData) IsKnown() bool {
	switch r {
	case SubscriptionNewResponseDataOk:
		return true
	}
	return false
}

type SubscriptionListResponse struct {
	// List of subscriptions.
	Data []SubscriptionListResponseData `json:"data,required"`
	JSON subscriptionListResponseJSON   `json:"-"`
}

// subscriptionListResponseJSON contains the JSON metadata for the struct
// [SubscriptionListResponse]
type subscriptionListResponseJSON struct {
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SubscriptionListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionListResponseJSON) RawJSON() string {
	return r.raw
}

// Subscription to a Nestri product.
type SubscriptionListResponseData struct {
	// Unique object identifier. The format and length of IDs may change over time.
	ID string `json:"id,required"`
	// Frequency of the subscription.
	Frequency SubscriptionListResponseDataFrequency `json:"frequency,required"`
	// Next billing date for the subscription.
	Next SubscriptionListResponseDataNextUnion `json:"next,required"`
	// ID of the product being subscribed to.
	ProductID string `json:"productID,required"`
	// Quantity of the subscription.
	Quantity int64 `json:"quantity,required"`
	// Cancelled date for the subscription.
	CanceledAt SubscriptionListResponseDataCanceledAtUnion `json:"canceledAt"`
	JSON       subscriptionListResponseDataJSON            `json:"-"`
}

// subscriptionListResponseDataJSON contains the JSON metadata for the struct
// [SubscriptionListResponseData]
type subscriptionListResponseDataJSON struct {
	ID          apijson.Field
	Frequency   apijson.Field
	Next        apijson.Field
	ProductID   apijson.Field
	Quantity    apijson.Field
	CanceledAt  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SubscriptionListResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionListResponseDataJSON) RawJSON() string {
	return r.raw
}

// Frequency of the subscription.
type SubscriptionListResponseDataFrequency string

const (
	SubscriptionListResponseDataFrequencyFixed   SubscriptionListResponseDataFrequency = "fixed"
	SubscriptionListResponseDataFrequencyDaily   SubscriptionListResponseDataFrequency = "daily"
	SubscriptionListResponseDataFrequencyWeekly  SubscriptionListResponseDataFrequency = "weekly"
	SubscriptionListResponseDataFrequencyMonthly SubscriptionListResponseDataFrequency = "monthly"
	SubscriptionListResponseDataFrequencyYearly  SubscriptionListResponseDataFrequency = "yearly"
)

func (r SubscriptionListResponseDataFrequency) IsKnown() bool {
	switch r {
	case SubscriptionListResponseDataFrequencyFixed, SubscriptionListResponseDataFrequencyDaily, SubscriptionListResponseDataFrequencyWeekly, SubscriptionListResponseDataFrequencyMonthly, SubscriptionListResponseDataFrequencyYearly:
		return true
	}
	return false
}

// Next billing date for the subscription.
//
// Union satisfied by [shared.UnionString] or [shared.UnionFloat].
type SubscriptionListResponseDataNextUnion interface {
	ImplementsSubscriptionListResponseDataNextUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*SubscriptionListResponseDataNextUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
	)
}

// Cancelled date for the subscription.
//
// Union satisfied by [shared.UnionString] or [shared.UnionFloat].
type SubscriptionListResponseDataCanceledAtUnion interface {
	ImplementsSubscriptionListResponseDataCanceledAtUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*SubscriptionListResponseDataCanceledAtUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
	)
}

type SubscriptionNewParams struct {
	// Unique object identifier. The format and length of IDs may change over time.
	ID param.Field[string] `json:"id,required"`
	// Frequency of the subscription.
	Frequency param.Field[SubscriptionNewParamsFrequency] `json:"frequency,required"`
	// Next billing date for the subscription.
	Next param.Field[SubscriptionNewParamsNextUnion] `json:"next,required"`
	// ID of the product being subscribed to.
	ProductID param.Field[string] `json:"productID,required"`
	// Quantity of the subscription.
	Quantity param.Field[int64] `json:"quantity,required"`
	// Cancelled date for the subscription.
	CanceledAt param.Field[SubscriptionNewParamsCanceledAtUnion] `json:"canceledAt"`
}

func (r SubscriptionNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Frequency of the subscription.
type SubscriptionNewParamsFrequency string

const (
	SubscriptionNewParamsFrequencyFixed   SubscriptionNewParamsFrequency = "fixed"
	SubscriptionNewParamsFrequencyDaily   SubscriptionNewParamsFrequency = "daily"
	SubscriptionNewParamsFrequencyWeekly  SubscriptionNewParamsFrequency = "weekly"
	SubscriptionNewParamsFrequencyMonthly SubscriptionNewParamsFrequency = "monthly"
	SubscriptionNewParamsFrequencyYearly  SubscriptionNewParamsFrequency = "yearly"
)

func (r SubscriptionNewParamsFrequency) IsKnown() bool {
	switch r {
	case SubscriptionNewParamsFrequencyFixed, SubscriptionNewParamsFrequencyDaily, SubscriptionNewParamsFrequencyWeekly, SubscriptionNewParamsFrequencyMonthly, SubscriptionNewParamsFrequencyYearly:
		return true
	}
	return false
}

// Next billing date for the subscription.
//
// Satisfied by [shared.UnionString], [shared.UnionFloat].
type SubscriptionNewParamsNextUnion interface {
	ImplementsSubscriptionNewParamsNextUnion()
}

// Cancelled date for the subscription.
//
// Satisfied by [shared.UnionString], [shared.UnionFloat].
type SubscriptionNewParamsCanceledAtUnion interface {
	ImplementsSubscriptionNewParamsCanceledAtUnion()
}
