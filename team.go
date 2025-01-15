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

// TeamService contains methods and other services that help with interacting with
// the nestri API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewTeamService] method instead.
type TeamService struct {
	Options []option.RequestOption
}

// NewTeamService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewTeamService(opts ...option.RequestOption) (r *TeamService) {
	r = &TeamService{}
	r.Options = opts
	return
}

// Create a new team for the currently authenticated user, enabling them to invite
// and play a game together with friends
func (r *TeamService) New(ctx context.Context, body TeamNewParams, opts ...option.RequestOption) (res *TeamNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "teams"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Fetch detailed information about a specific team using its unique slug
func (r *TeamService) Get(ctx context.Context, slug string, opts ...option.RequestOption) (res *TeamGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	if slug == "" {
		err = errors.New("missing required slug parameter")
		return
	}
	path := fmt.Sprintf("teams/%s", slug)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Returns a list of all teams which the authenticated user is part of
func (r *TeamService) List(ctx context.Context, opts ...option.RequestOption) (res *TeamListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "teams"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// This endpoint allows a user to delete a team, by providing it's unique slug
func (r *TeamService) Delete(ctx context.Context, slug string, opts ...option.RequestOption) (res *TeamDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	if slug == "" {
		err = errors.New("missing required slug parameter")
		return
	}
	path := fmt.Sprintf("teams/%s", slug)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Invite a user to a team owned by the current user
func (r *TeamService) Invite(ctx context.Context, slug string, email string, opts ...option.RequestOption) (res *TeamInviteResponse, err error) {
	opts = append(r.Options[:], opts...)
	if slug == "" {
		err = errors.New("missing required slug parameter")
		return
	}
	if email == "" {
		err = errors.New("missing required email parameter")
		return
	}
	path := fmt.Sprintf("teams/%s/invite/%s", slug, email)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

type TeamNewResponse struct {
	Data TeamNewResponseData `json:"data,required"`
	JSON teamNewResponseJSON `json:"-"`
}

// teamNewResponseJSON contains the JSON metadata for the struct [TeamNewResponse]
type teamNewResponseJSON struct {
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TeamNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r teamNewResponseJSON) RawJSON() string {
	return r.raw
}

type TeamNewResponseData string

const (
	TeamNewResponseDataOk TeamNewResponseData = "ok"
)

func (r TeamNewResponseData) IsKnown() bool {
	switch r {
	case TeamNewResponseDataOk:
		return true
	}
	return false
}

type TeamGetResponse struct {
	// A group of users sharing the same machines for gaming.
	Data TeamGetResponseData `json:"data,required"`
	JSON teamGetResponseJSON `json:"-"`
}

// teamGetResponseJSON contains the JSON metadata for the struct [TeamGetResponse]
type teamGetResponseJSON struct {
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TeamGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r teamGetResponseJSON) RawJSON() string {
	return r.raw
}

// A group of users sharing the same machines for gaming.
type TeamGetResponseData struct {
	// Unique object identifier. The format and length of IDs may change over time.
	ID string `json:"id,required"`
	// The time when this team was first created
	CreatedAt TeamGetResponseDataCreatedAtUnion `json:"createdAt,required"`
	// Name of the team
	Name string `json:"name,required"`
	// Whether this team is owned by this user
	Owner bool `json:"owner,required"`
	// This is the unique name identifier for the team
	Slug string `json:"slug,required"`
	// The time when this team was last edited
	UpdatedAt TeamGetResponseDataUpdatedAtUnion `json:"updatedAt,required"`
	JSON      teamGetResponseDataJSON           `json:"-"`
}

// teamGetResponseDataJSON contains the JSON metadata for the struct
// [TeamGetResponseData]
type teamGetResponseDataJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	Name        apijson.Field
	Owner       apijson.Field
	Slug        apijson.Field
	UpdatedAt   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TeamGetResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r teamGetResponseDataJSON) RawJSON() string {
	return r.raw
}

// The time when this team was first created
//
// Union satisfied by [shared.UnionString] or [shared.UnionFloat].
type TeamGetResponseDataCreatedAtUnion interface {
	ImplementsTeamGetResponseDataCreatedAtUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*TeamGetResponseDataCreatedAtUnion)(nil)).Elem(),
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

// The time when this team was last edited
//
// Union satisfied by [shared.UnionString] or [shared.UnionFloat].
type TeamGetResponseDataUpdatedAtUnion interface {
	ImplementsTeamGetResponseDataUpdatedAtUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*TeamGetResponseDataUpdatedAtUnion)(nil)).Elem(),
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

type TeamListResponse struct {
	// A list of teams associated with the user
	Data []TeamListResponseData `json:"data,required"`
	JSON teamListResponseJSON   `json:"-"`
}

// teamListResponseJSON contains the JSON metadata for the struct
// [TeamListResponse]
type teamListResponseJSON struct {
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TeamListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r teamListResponseJSON) RawJSON() string {
	return r.raw
}

// A group of users sharing the same machines for gaming.
type TeamListResponseData struct {
	// Unique object identifier. The format and length of IDs may change over time.
	ID string `json:"id,required"`
	// The time when this team was first created
	CreatedAt TeamListResponseDataCreatedAtUnion `json:"createdAt,required"`
	// Name of the team
	Name string `json:"name,required"`
	// Whether this team is owned by this user
	Owner bool `json:"owner,required"`
	// This is the unique name identifier for the team
	Slug string `json:"slug,required"`
	// The time when this team was last edited
	UpdatedAt TeamListResponseDataUpdatedAtUnion `json:"updatedAt,required"`
	JSON      teamListResponseDataJSON           `json:"-"`
}

// teamListResponseDataJSON contains the JSON metadata for the struct
// [TeamListResponseData]
type teamListResponseDataJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	Name        apijson.Field
	Owner       apijson.Field
	Slug        apijson.Field
	UpdatedAt   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TeamListResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r teamListResponseDataJSON) RawJSON() string {
	return r.raw
}

// The time when this team was first created
//
// Union satisfied by [shared.UnionString] or [shared.UnionFloat].
type TeamListResponseDataCreatedAtUnion interface {
	ImplementsTeamListResponseDataCreatedAtUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*TeamListResponseDataCreatedAtUnion)(nil)).Elem(),
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

// The time when this team was last edited
//
// Union satisfied by [shared.UnionString] or [shared.UnionFloat].
type TeamListResponseDataUpdatedAtUnion interface {
	ImplementsTeamListResponseDataUpdatedAtUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*TeamListResponseDataUpdatedAtUnion)(nil)).Elem(),
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

type TeamDeleteResponse struct {
	Data TeamDeleteResponseData `json:"data,required"`
	JSON teamDeleteResponseJSON `json:"-"`
}

// teamDeleteResponseJSON contains the JSON metadata for the struct
// [TeamDeleteResponse]
type teamDeleteResponseJSON struct {
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TeamDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r teamDeleteResponseJSON) RawJSON() string {
	return r.raw
}

type TeamDeleteResponseData string

const (
	TeamDeleteResponseDataOk TeamDeleteResponseData = "ok"
)

func (r TeamDeleteResponseData) IsKnown() bool {
	switch r {
	case TeamDeleteResponseDataOk:
		return true
	}
	return false
}

type TeamInviteResponse struct {
	Data TeamInviteResponseData `json:"data,required"`
	JSON teamInviteResponseJSON `json:"-"`
}

// teamInviteResponseJSON contains the JSON metadata for the struct
// [TeamInviteResponse]
type teamInviteResponseJSON struct {
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TeamInviteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r teamInviteResponseJSON) RawJSON() string {
	return r.raw
}

type TeamInviteResponseData string

const (
	TeamInviteResponseDataOk TeamInviteResponseData = "ok"
)

func (r TeamInviteResponseData) IsKnown() bool {
	switch r {
	case TeamInviteResponseDataOk:
		return true
	}
	return false
}

type TeamNewParams struct {
	// The human readable name to give this team
	Name param.Field[string] `json:"name,required"`
	// The unique name to be used with this team
	Slug param.Field[string] `json:"slug,required"`
}

func (r TeamNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
