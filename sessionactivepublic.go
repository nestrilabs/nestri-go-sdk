// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package nestri

import (
	"context"
	"net/http"
	"reflect"

	"github.com/nestrilabs/nestri-go-sdk/internal/apijson"
	"github.com/nestrilabs/nestri-go-sdk/internal/requestconfig"
	"github.com/nestrilabs/nestri-go-sdk/option"
	"github.com/nestrilabs/nestri-go-sdk/shared"
	"github.com/tidwall/gjson"
)

// SessionActivePublicService contains methods and other services that help with
// interacting with the nestri API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSessionActivePublicService] method instead.
type SessionActivePublicService struct {
	Options []option.RequestOption
}

// NewSessionActivePublicService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewSessionActivePublicService(opts ...option.RequestOption) (r *SessionActivePublicService) {
	r = &SessionActivePublicService{}
	r.Options = opts
	return
}

// Returns a list of all publicly active gaming sessions associated
func (r *SessionActivePublicService) List(ctx context.Context, opts ...option.RequestOption) (res *SessionActivePublicListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "sessions/active/public"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type SessionActivePublicListResponse struct {
	// A list of publicly active gaming sessions
	Data []SessionActivePublicListResponseData `json:"data,required"`
	JSON sessionActivePublicListResponseJSON   `json:"-"`
}

// sessionActivePublicListResponseJSON contains the JSON metadata for the struct
// [SessionActivePublicListResponse]
type sessionActivePublicListResponseJSON struct {
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SessionActivePublicListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sessionActivePublicListResponseJSON) RawJSON() string {
	return r.raw
}

// Represents a single game play session, tracking its lifetime and accessibility
// settings.
type SessionActivePublicListResponseData struct {
	// Unique object identifier. The format and length of IDs may change over time.
	ID string `json:"id,required"`
	// A human-readable name for the session to help identify it
	Name string `json:"name,required"`
	// If true, the session is publicly viewable by all users. If false, only
	// authorized users can access it
	Public bool `json:"public,required"`
	// The timestamp indicating when this session started.
	StartedAt SessionActivePublicListResponseDataStartedAtUnion `json:"startedAt,required"`
	// The timestamp indicating when this session was completed or terminated. Null if
	// session is still active.
	EndedAt SessionActivePublicListResponseDataEndedAtUnion `json:"endedAt"`
	JSON    sessionActivePublicListResponseDataJSON         `json:"-"`
}

// sessionActivePublicListResponseDataJSON contains the JSON metadata for the
// struct [SessionActivePublicListResponseData]
type sessionActivePublicListResponseDataJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	Public      apijson.Field
	StartedAt   apijson.Field
	EndedAt     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SessionActivePublicListResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sessionActivePublicListResponseDataJSON) RawJSON() string {
	return r.raw
}

// The timestamp indicating when this session started.
//
// Union satisfied by [shared.UnionString] or [shared.UnionFloat].
type SessionActivePublicListResponseDataStartedAtUnion interface {
	ImplementsSessionActivePublicListResponseDataStartedAtUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*SessionActivePublicListResponseDataStartedAtUnion)(nil)).Elem(),
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

// The timestamp indicating when this session was completed or terminated. Null if
// session is still active.
//
// Union satisfied by [shared.UnionString] or [shared.UnionFloat].
type SessionActivePublicListResponseDataEndedAtUnion interface {
	ImplementsSessionActivePublicListResponseDataEndedAtUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*SessionActivePublicListResponseDataEndedAtUnion)(nil)).Elem(),
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
