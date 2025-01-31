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

// SessionActiveService contains methods and other services that help with
// interacting with the nestri API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSessionActiveService] method instead.
type SessionActiveService struct {
	Options []option.RequestOption
	Public  *SessionActivePublicService
}

// NewSessionActiveService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewSessionActiveService(opts ...option.RequestOption) (r *SessionActiveService) {
	r = &SessionActiveService{}
	r.Options = opts
	r.Public = NewSessionActivePublicService(opts...)
	return
}

// Returns a list of all active gaming sessions associated with the authenticated
// user
func (r *SessionActiveService) List(ctx context.Context, opts ...option.RequestOption) (res *SessionActiveListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "sessions/active"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type SessionActiveListResponse struct {
	// A list of active gaming sessions associated with the user
	Data []SessionActiveListResponseData `json:"data,required"`
	JSON sessionActiveListResponseJSON   `json:"-"`
}

// sessionActiveListResponseJSON contains the JSON metadata for the struct
// [SessionActiveListResponse]
type sessionActiveListResponseJSON struct {
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SessionActiveListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sessionActiveListResponseJSON) RawJSON() string {
	return r.raw
}

// Represents a single game play session, tracking its lifetime and accessibility
// settings.
type SessionActiveListResponseData struct {
	// Unique object identifier. The format and length of IDs may change over time.
	ID string `json:"id,required"`
	// If true, the session is publicly viewable by all users. If false, only
	// authorized users can access it
	Public bool `json:"public,required"`
	// The timestamp indicating when this session started.
	StartedAt SessionActiveListResponseDataStartedAtUnion `json:"startedAt,required"`
	// The timestamp indicating when this session was completed or terminated. Null if
	// session is still active.
	EndedAt SessionActiveListResponseDataEndedAtUnion `json:"endedAt"`
	JSON    sessionActiveListResponseDataJSON         `json:"-"`
}

// sessionActiveListResponseDataJSON contains the JSON metadata for the struct
// [SessionActiveListResponseData]
type sessionActiveListResponseDataJSON struct {
	ID          apijson.Field
	Public      apijson.Field
	StartedAt   apijson.Field
	EndedAt     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SessionActiveListResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sessionActiveListResponseDataJSON) RawJSON() string {
	return r.raw
}

// The timestamp indicating when this session started.
//
// Union satisfied by [shared.UnionString] or [shared.UnionFloat].
type SessionActiveListResponseDataStartedAtUnion interface {
	ImplementsSessionActiveListResponseDataStartedAtUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*SessionActiveListResponseDataStartedAtUnion)(nil)).Elem(),
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
type SessionActiveListResponseDataEndedAtUnion interface {
	ImplementsSessionActiveListResponseDataEndedAtUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*SessionActiveListResponseDataEndedAtUnion)(nil)).Elem(),
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
