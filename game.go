// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package nestri

import (
	"context"
	"fmt"
	"net/http"

	"github.com/nestrilabs/nestri-go-sdk/internal/apijson"
	"github.com/nestrilabs/nestri-go-sdk/internal/param"
	"github.com/nestrilabs/nestri-go-sdk/internal/requestconfig"
	"github.com/nestrilabs/nestri-go-sdk/option"
)

// GameService contains methods and other services that help with interacting with
// the nestri API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewGameService] method instead.
type GameService struct {
	Options  []option.RequestOption
	Sessions *GameSessionService
}

// NewGameService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewGameService(opts ...option.RequestOption) (r *GameService) {
	r = &GameService{}
	r.Options = opts
	r.Sessions = NewGameSessionService(opts...)
	return
}

// Adds a game to the currently authenticated user's library. Once added, the user
// can play the game and share their progress with others
func (r *GameService) New(ctx context.Context, steamID float64, opts ...option.RequestOption) (res *GameNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("games/%v", steamID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

// Fetches detailed metadata about a specific game using its Steam ID
func (r *GameService) Get(ctx context.Context, steamID float64, opts ...option.RequestOption) (res *GameGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("games/%v", steamID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Updates the metadata about a specific game using its Steam ID
func (r *GameService) Update(ctx context.Context, body GameUpdateParams, opts ...option.RequestOption) (res *GameUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "games"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// Returns a list of all (known) games associated with the authenticated user
func (r *GameService) List(ctx context.Context, opts ...option.RequestOption) (res *GameListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "games"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Removes a game from the authenticated user's library. The game remains in the
// system but will no longer be accessible to the user
func (r *GameService) Delete(ctx context.Context, steamID float64, opts ...option.RequestOption) (res *GameDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("games/%v", steamID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

type GameNewResponse struct {
	Data GameNewResponseData `json:"data,required"`
	JSON gameNewResponseJSON `json:"-"`
}

// gameNewResponseJSON contains the JSON metadata for the struct [GameNewResponse]
type gameNewResponseJSON struct {
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GameNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r gameNewResponseJSON) RawJSON() string {
	return r.raw
}

type GameNewResponseData string

const (
	GameNewResponseDataOk GameNewResponseData = "ok"
)

func (r GameNewResponseData) IsKnown() bool {
	switch r {
	case GameNewResponseDataOk:
		return true
	}
	return false
}

type GameGetResponse struct {
	// Represents a Steam game that can be installed and played on a machine.
	Data GameGetResponseData `json:"data,required"`
	JSON gameGetResponseJSON `json:"-"`
}

// gameGetResponseJSON contains the JSON metadata for the struct [GameGetResponse]
type gameGetResponseJSON struct {
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GameGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r gameGetResponseJSON) RawJSON() string {
	return r.raw
}

// Represents a Steam game that can be installed and played on a machine.
type GameGetResponseData struct {
	// Unique object identifier. The format and length of IDs may change over time.
	ID string `json:"id,required"`
	// A human-readable name for the game, used for easy identification.
	Name string `json:"name,required"`
	// The Steam ID of the game, used to identify it during installation and runtime.
	SteamID float64                 `json:"steamID,required"`
	JSON    gameGetResponseDataJSON `json:"-"`
}

// gameGetResponseDataJSON contains the JSON metadata for the struct
// [GameGetResponseData]
type gameGetResponseDataJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	SteamID     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GameGetResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r gameGetResponseDataJSON) RawJSON() string {
	return r.raw
}

type GameUpdateResponse struct {
	Data GameUpdateResponseData `json:"data,required"`
	JSON gameUpdateResponseJSON `json:"-"`
}

// gameUpdateResponseJSON contains the JSON metadata for the struct
// [GameUpdateResponse]
type gameUpdateResponseJSON struct {
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GameUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r gameUpdateResponseJSON) RawJSON() string {
	return r.raw
}

type GameUpdateResponseData string

const (
	GameUpdateResponseDataOk GameUpdateResponseData = "ok"
)

func (r GameUpdateResponseData) IsKnown() bool {
	switch r {
	case GameUpdateResponseDataOk:
		return true
	}
	return false
}

type GameListResponse struct {
	// A list of games owned by the user
	Data []GameListResponseData `json:"data,required"`
	JSON gameListResponseJSON   `json:"-"`
}

// gameListResponseJSON contains the JSON metadata for the struct
// [GameListResponse]
type gameListResponseJSON struct {
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GameListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r gameListResponseJSON) RawJSON() string {
	return r.raw
}

// Represents a Steam game that can be installed and played on a machine.
type GameListResponseData struct {
	// Unique object identifier. The format and length of IDs may change over time.
	ID string `json:"id,required"`
	// A human-readable name for the game, used for easy identification.
	Name string `json:"name,required"`
	// The Steam ID of the game, used to identify it during installation and runtime.
	SteamID float64                  `json:"steamID,required"`
	JSON    gameListResponseDataJSON `json:"-"`
}

// gameListResponseDataJSON contains the JSON metadata for the struct
// [GameListResponseData]
type gameListResponseDataJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	SteamID     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GameListResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r gameListResponseDataJSON) RawJSON() string {
	return r.raw
}

type GameDeleteResponse struct {
	Data GameDeleteResponseData `json:"data,required"`
	JSON gameDeleteResponseJSON `json:"-"`
}

// gameDeleteResponseJSON contains the JSON metadata for the struct
// [GameDeleteResponse]
type gameDeleteResponseJSON struct {
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GameDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r gameDeleteResponseJSON) RawJSON() string {
	return r.raw
}

type GameDeleteResponseData string

const (
	GameDeleteResponseDataOk GameDeleteResponseData = "ok"
)

func (r GameDeleteResponseData) IsKnown() bool {
	switch r {
	case GameDeleteResponseDataOk:
		return true
	}
	return false
}

type GameUpdateParams struct {
	// A human-readable name for the game, used for easy identification.
	Name param.Field[string] `json:"name,required"`
	// The Steam ID of the game, used to identify it during installation and runtime.
	SteamID param.Field[float64] `json:"steamID,required"`
}

func (r GameUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
