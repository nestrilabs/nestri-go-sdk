// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package nestri

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"reflect"

	"github.com/nestrilabs/nestri-go-sdk/internal/apijson"
	"github.com/nestrilabs/nestri-go-sdk/internal/param"
	"github.com/nestrilabs/nestri-go-sdk/internal/requestconfig"
	"github.com/nestrilabs/nestri-go-sdk/option"
	"github.com/nestrilabs/nestri-go-sdk/shared"
	"github.com/tidwall/gjson"
)

// SessionService contains methods and other services that help with interacting
// with the nestri API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSessionService] method instead.
type SessionService struct {
	Options []option.RequestOption
	Active  *SessionActiveService
}

// NewSessionService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewSessionService(opts ...option.RequestOption) (r *SessionService) {
	r = &SessionService{}
	r.Options = opts
	r.Active = NewSessionActiveService(opts...)
	return
}

// Create a new gaming session for the currently authenticated user, enabling them
// to play a game
func (r *SessionService) New(ctx context.Context, body SessionNewParams, opts ...option.RequestOption) (res *SessionNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "sessions"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Fetches detailed information about a specific gaming session using its unique id
func (r *SessionService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *SessionGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("sessions/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Returns a list of all gaming sessions associated with the authenticated user
func (r *SessionService) List(ctx context.Context, opts ...option.RequestOption) (res *SessionListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "sessions"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// This endpoint allows a user to terminate an active gaming session by providing
// the session's unique ID
func (r *SessionService) Delete(ctx context.Context, id string, opts ...option.RequestOption) (res *SessionDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("sessions/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

type SessionNewResponse struct {
	Data SessionNewResponseData `json:"data,required"`
	JSON sessionNewResponseJSON `json:"-"`
}

// sessionNewResponseJSON contains the JSON metadata for the struct
// [SessionNewResponse]
type sessionNewResponseJSON struct {
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SessionNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sessionNewResponseJSON) RawJSON() string {
	return r.raw
}

type SessionNewResponseData string

const (
	SessionNewResponseDataOk SessionNewResponseData = "ok"
)

func (r SessionNewResponseData) IsKnown() bool {
	switch r {
	case SessionNewResponseDataOk:
		return true
	}
	return false
}

type SessionGetResponse struct {
	// Represents a single game play session, tracking its lifetime and accessibility
	// settings.
	Data SessionGetResponseData `json:"data,required"`
	JSON sessionGetResponseJSON `json:"-"`
}

// sessionGetResponseJSON contains the JSON metadata for the struct
// [SessionGetResponse]
type sessionGetResponseJSON struct {
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SessionGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sessionGetResponseJSON) RawJSON() string {
	return r.raw
}

// Represents a single game play session, tracking its lifetime and accessibility
// settings.
type SessionGetResponseData struct {
	// Unique object identifier. The format and length of IDs may change over time.
	ID string `json:"id,required"`
	// A human-readable name for the session to help identify it
	Name string `json:"name,required"`
	// If true, the session is publicly viewable by all users. If false, only
	// authorized users can access it
	Public bool `json:"public,required"`
	// The timestamp indicating when this session started.
	StartedAt SessionGetResponseDataStartedAtUnion `json:"startedAt,required"`
	// The timestamp indicating when this session was completed or terminated. Null if
	// session is still active.
	EndedAt SessionGetResponseDataEndedAtUnion `json:"endedAt"`
	JSON    sessionGetResponseDataJSON         `json:"-"`
}

// sessionGetResponseDataJSON contains the JSON metadata for the struct
// [SessionGetResponseData]
type sessionGetResponseDataJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	Public      apijson.Field
	StartedAt   apijson.Field
	EndedAt     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SessionGetResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sessionGetResponseDataJSON) RawJSON() string {
	return r.raw
}

// The timestamp indicating when this session started.
//
// Union satisfied by [shared.UnionString] or [shared.UnionFloat].
type SessionGetResponseDataStartedAtUnion interface {
	ImplementsSessionGetResponseDataStartedAtUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*SessionGetResponseDataStartedAtUnion)(nil)).Elem(),
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
type SessionGetResponseDataEndedAtUnion interface {
	ImplementsSessionGetResponseDataEndedAtUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*SessionGetResponseDataEndedAtUnion)(nil)).Elem(),
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

type SessionListResponse struct {
	// A list of gaming sessions associated with the user
	Data []SessionListResponseData `json:"data,required"`
	JSON sessionListResponseJSON   `json:"-"`
}

// sessionListResponseJSON contains the JSON metadata for the struct
// [SessionListResponse]
type sessionListResponseJSON struct {
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SessionListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sessionListResponseJSON) RawJSON() string {
	return r.raw
}

// Represents a single game play session, tracking its lifetime and accessibility
// settings.
type SessionListResponseData struct {
	// Unique object identifier. The format and length of IDs may change over time.
	ID string `json:"id,required"`
	// A human-readable name for the session to help identify it
	Name string `json:"name,required"`
	// If true, the session is publicly viewable by all users. If false, only
	// authorized users can access it
	Public bool `json:"public,required"`
	// The timestamp indicating when this session started.
	StartedAt SessionListResponseDataStartedAtUnion `json:"startedAt,required"`
	// The timestamp indicating when this session was completed or terminated. Null if
	// session is still active.
	EndedAt SessionListResponseDataEndedAtUnion `json:"endedAt"`
	JSON    sessionListResponseDataJSON         `json:"-"`
}

// sessionListResponseDataJSON contains the JSON metadata for the struct
// [SessionListResponseData]
type sessionListResponseDataJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	Public      apijson.Field
	StartedAt   apijson.Field
	EndedAt     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SessionListResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sessionListResponseDataJSON) RawJSON() string {
	return r.raw
}

// The timestamp indicating when this session started.
//
// Union satisfied by [shared.UnionString] or [shared.UnionFloat].
type SessionListResponseDataStartedAtUnion interface {
	ImplementsSessionListResponseDataStartedAtUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*SessionListResponseDataStartedAtUnion)(nil)).Elem(),
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
type SessionListResponseDataEndedAtUnion interface {
	ImplementsSessionListResponseDataEndedAtUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*SessionListResponseDataEndedAtUnion)(nil)).Elem(),
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

type SessionDeleteResponse struct {
	Data SessionDeleteResponseData `json:"data,required"`
	JSON sessionDeleteResponseJSON `json:"-"`
}

// sessionDeleteResponseJSON contains the JSON metadata for the struct
// [SessionDeleteResponse]
type sessionDeleteResponseJSON struct {
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SessionDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sessionDeleteResponseJSON) RawJSON() string {
	return r.raw
}

type SessionDeleteResponseData string

const (
	SessionDeleteResponseDataOk SessionDeleteResponseData = "ok"
)

func (r SessionDeleteResponseData) IsKnown() bool {
	switch r {
	case SessionDeleteResponseDataOk:
		return true
	}
	return false
}

type SessionNewParams struct {
	// The unique fingerprint of the machine to play on, derived from its Linux machine
	// ID
	Fingerprint param.Field[string] `json:"fingerprint,required"`
	// The human readable name to give this session
	Name param.Field[string] `json:"name,required"`
	// Whether the session is publicly viewable by all users. If false, only authorized
	// users can access it
	Public param.Field[bool] `json:"public,required"`
	// The Steam ID of the game the user wants to play
	SteamID param.Field[float64] `json:"steamID,required"`
}

func (r SessionNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
