// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package nestri

import (
	"context"
	"fmt"
	"net/http"
	"reflect"

	"github.com/nestrilabs/nestri-go-sdk/internal/apijson"
	"github.com/nestrilabs/nestri-go-sdk/internal/requestconfig"
	"github.com/nestrilabs/nestri-go-sdk/option"
	"github.com/nestrilabs/nestri-go-sdk/shared"
	"github.com/tidwall/gjson"
)

// GameSessionService contains methods and other services that help with
// interacting with the nestri API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewGameSessionService] method instead.
type GameSessionService struct {
	Options []option.RequestOption
}

// NewGameSessionService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewGameSessionService(opts ...option.RequestOption) (r *GameSessionService) {
	r = &GameSessionService{}
	r.Options = opts
	return
}

// Fetches active and public game sessions associated with a specific game using
// its Steam ID
func (r *GameSessionService) List(ctx context.Context, steamID float64, opts ...option.RequestOption) (res *GameSessionListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("games/%v/sessions", steamID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type GameSessionListResponse struct {
	// Publicly active sessions associated with the game
	Data []GameSessionListResponseData `json:"data,required"`
	JSON gameSessionListResponseJSON   `json:"-"`
}

// gameSessionListResponseJSON contains the JSON metadata for the struct
// [GameSessionListResponse]
type gameSessionListResponseJSON struct {
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GameSessionListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r gameSessionListResponseJSON) RawJSON() string {
	return r.raw
}

// Represents a single game play session, tracking its lifetime and accessibility
// settings.
type GameSessionListResponseData struct {
	// Unique object identifier. The format and length of IDs may change over time.
	ID string `json:"id,required"`
	// A human-readable name for the session to help identify it
	Name string `json:"name,required"`
	// If true, the session is publicly viewable by all users. If false, only
	// authorized users can access it
	Public bool `json:"public,required"`
	// The timestamp indicating when this session started.
	StartedAt GameSessionListResponseDataStartedAtUnion `json:"startedAt,required"`
	// The timestamp indicating when this session was completed or terminated. Null if
	// session is still active.
	EndedAt GameSessionListResponseDataEndedAtUnion `json:"endedAt"`
	JSON    gameSessionListResponseDataJSON         `json:"-"`
}

// gameSessionListResponseDataJSON contains the JSON metadata for the struct
// [GameSessionListResponseData]
type gameSessionListResponseDataJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	Public      apijson.Field
	StartedAt   apijson.Field
	EndedAt     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GameSessionListResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r gameSessionListResponseDataJSON) RawJSON() string {
	return r.raw
}

// The timestamp indicating when this session started.
//
// Union satisfied by [shared.UnionString] or [shared.UnionFloat].
type GameSessionListResponseDataStartedAtUnion interface {
	ImplementsGameSessionListResponseDataStartedAtUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*GameSessionListResponseDataStartedAtUnion)(nil)).Elem(),
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
type GameSessionListResponseDataEndedAtUnion interface {
	ImplementsGameSessionListResponseDataEndedAtUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*GameSessionListResponseDataEndedAtUnion)(nil)).Elem(),
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
