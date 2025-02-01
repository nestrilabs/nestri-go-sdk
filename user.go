// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package nestri

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"reflect"

	"github.com/nestrilabs/nestri-go-sdk/internal/apijson"
	"github.com/nestrilabs/nestri-go-sdk/internal/requestconfig"
	"github.com/nestrilabs/nestri-go-sdk/option"
	"github.com/nestrilabs/nestri-go-sdk/shared"
	"github.com/tidwall/gjson"
)

// UserService contains methods and other services that help with interacting with
// the nestri API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewUserService] method instead.
type UserService struct {
	Options []option.RequestOption
}

// NewUserService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewUserService(opts ...option.RequestOption) (r *UserService) {
	r = &UserService{}
	r.Options = opts
	return
}

// Gets a user's profile by their id
func (r *UserService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *UserGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("users/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Returns all user profiles
func (r *UserService) List(ctx context.Context, opts ...option.RequestOption) (res *UserListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "users"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Get a user's gaming session details by their id
func (r *UserService) Session(ctx context.Context, id string, opts ...option.RequestOption) (res *UserSessionResponse, err error) {
	opts = append(r.Options[:], opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("users/%s/session", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type UserGetResponse struct {
	// Represents a profile of a user on Nestri
	Data UserGetResponseData `json:"data,required"`
	JSON userGetResponseJSON `json:"-"`
}

// userGetResponseJSON contains the JSON metadata for the struct [UserGetResponse]
type userGetResponseJSON struct {
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userGetResponseJSON) RawJSON() string {
	return r.raw
}

// Represents a profile of a user on Nestri
type UserGetResponseData struct {
	// Unique object identifier. The format and length of IDs may change over time.
	ID string `json:"id,required"`
	// The time when this profile was first created
	CreatedAt UserGetResponseDataCreatedAtUnion `json:"createdAt,required"`
	// The number discriminator for each username
	Discriminator UserGetResponseDataDiscriminatorUnion `json:"discriminator,required"`
	// Whether the user is active, idle or offline
	Status UserGetResponseDataStatus `json:"status,required"`
	// The time when this profile was last edited
	UpdatedAt UserGetResponseDataUpdatedAtUnion `json:"updatedAt,required"`
	// The user's unique username
	Username string `json:"username,required"`
	// The url to the profile picture.
	AvatarURL string                  `json:"avatarUrl"`
	JSON      userGetResponseDataJSON `json:"-"`
}

// userGetResponseDataJSON contains the JSON metadata for the struct
// [UserGetResponseData]
type userGetResponseDataJSON struct {
	ID            apijson.Field
	CreatedAt     apijson.Field
	Discriminator apijson.Field
	Status        apijson.Field
	UpdatedAt     apijson.Field
	Username      apijson.Field
	AvatarURL     apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *UserGetResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userGetResponseDataJSON) RawJSON() string {
	return r.raw
}

// The time when this profile was first created
//
// Union satisfied by [shared.UnionString] or [shared.UnionFloat].
type UserGetResponseDataCreatedAtUnion interface {
	ImplementsUserGetResponseDataCreatedAtUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*UserGetResponseDataCreatedAtUnion)(nil)).Elem(),
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

// The number discriminator for each username
//
// Union satisfied by [shared.UnionString] or [shared.UnionFloat].
type UserGetResponseDataDiscriminatorUnion interface {
	ImplementsUserGetResponseDataDiscriminatorUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*UserGetResponseDataDiscriminatorUnion)(nil)).Elem(),
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

// Whether the user is active, idle or offline
type UserGetResponseDataStatus string

const (
	UserGetResponseDataStatusActive  UserGetResponseDataStatus = "active"
	UserGetResponseDataStatusIdle    UserGetResponseDataStatus = "idle"
	UserGetResponseDataStatusOffline UserGetResponseDataStatus = "offline"
)

func (r UserGetResponseDataStatus) IsKnown() bool {
	switch r {
	case UserGetResponseDataStatusActive, UserGetResponseDataStatusIdle, UserGetResponseDataStatusOffline:
		return true
	}
	return false
}

// The time when this profile was last edited
//
// Union satisfied by [shared.UnionString] or [shared.UnionFloat].
type UserGetResponseDataUpdatedAtUnion interface {
	ImplementsUserGetResponseDataUpdatedAtUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*UserGetResponseDataUpdatedAtUnion)(nil)).Elem(),
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

type UserListResponse struct {
	// Represents a profile of a user on Nestri
	Data UserListResponseData `json:"data,required"`
	JSON userListResponseJSON `json:"-"`
}

// userListResponseJSON contains the JSON metadata for the struct
// [UserListResponse]
type userListResponseJSON struct {
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userListResponseJSON) RawJSON() string {
	return r.raw
}

// Represents a profile of a user on Nestri
type UserListResponseData struct {
	// Unique object identifier. The format and length of IDs may change over time.
	ID string `json:"id,required"`
	// The time when this profile was first created
	CreatedAt UserListResponseDataCreatedAtUnion `json:"createdAt,required"`
	// The number discriminator for each username
	Discriminator UserListResponseDataDiscriminatorUnion `json:"discriminator,required"`
	// Whether the user is active, idle or offline
	Status UserListResponseDataStatus `json:"status,required"`
	// The time when this profile was last edited
	UpdatedAt UserListResponseDataUpdatedAtUnion `json:"updatedAt,required"`
	// The user's unique username
	Username string `json:"username,required"`
	// The url to the profile picture.
	AvatarURL string                   `json:"avatarUrl"`
	JSON      userListResponseDataJSON `json:"-"`
}

// userListResponseDataJSON contains the JSON metadata for the struct
// [UserListResponseData]
type userListResponseDataJSON struct {
	ID            apijson.Field
	CreatedAt     apijson.Field
	Discriminator apijson.Field
	Status        apijson.Field
	UpdatedAt     apijson.Field
	Username      apijson.Field
	AvatarURL     apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *UserListResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userListResponseDataJSON) RawJSON() string {
	return r.raw
}

// The time when this profile was first created
//
// Union satisfied by [shared.UnionString] or [shared.UnionFloat].
type UserListResponseDataCreatedAtUnion interface {
	ImplementsUserListResponseDataCreatedAtUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*UserListResponseDataCreatedAtUnion)(nil)).Elem(),
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

// The number discriminator for each username
//
// Union satisfied by [shared.UnionString] or [shared.UnionFloat].
type UserListResponseDataDiscriminatorUnion interface {
	ImplementsUserListResponseDataDiscriminatorUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*UserListResponseDataDiscriminatorUnion)(nil)).Elem(),
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

// Whether the user is active, idle or offline
type UserListResponseDataStatus string

const (
	UserListResponseDataStatusActive  UserListResponseDataStatus = "active"
	UserListResponseDataStatusIdle    UserListResponseDataStatus = "idle"
	UserListResponseDataStatusOffline UserListResponseDataStatus = "offline"
)

func (r UserListResponseDataStatus) IsKnown() bool {
	switch r {
	case UserListResponseDataStatusActive, UserListResponseDataStatusIdle, UserListResponseDataStatusOffline:
		return true
	}
	return false
}

// The time when this profile was last edited
//
// Union satisfied by [shared.UnionString] or [shared.UnionFloat].
type UserListResponseDataUpdatedAtUnion interface {
	ImplementsUserListResponseDataUpdatedAtUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*UserListResponseDataUpdatedAtUnion)(nil)).Elem(),
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

type UserSessionResponse struct {
	// Represents a single game play session, tracking its lifetime and accessibility
	// settings.
	Data UserSessionResponseData `json:"data,required"`
	JSON userSessionResponseJSON `json:"-"`
}

// userSessionResponseJSON contains the JSON metadata for the struct
// [UserSessionResponse]
type userSessionResponseJSON struct {
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserSessionResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userSessionResponseJSON) RawJSON() string {
	return r.raw
}

// Represents a single game play session, tracking its lifetime and accessibility
// settings.
type UserSessionResponseData struct {
	// Unique object identifier. The format and length of IDs may change over time.
	ID string `json:"id,required"`
	// If true, the session is publicly viewable by all users. If false, only
	// authorized users can access it
	Public bool `json:"public,required"`
	// The timestamp indicating when this session started.
	StartedAt UserSessionResponseDataStartedAtUnion `json:"startedAt,required"`
	// The timestamp indicating when this session was completed or terminated. Null if
	// session is still active.
	EndedAt UserSessionResponseDataEndedAtUnion `json:"endedAt"`
	JSON    userSessionResponseDataJSON         `json:"-"`
}

// userSessionResponseDataJSON contains the JSON metadata for the struct
// [UserSessionResponseData]
type userSessionResponseDataJSON struct {
	ID          apijson.Field
	Public      apijson.Field
	StartedAt   apijson.Field
	EndedAt     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserSessionResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userSessionResponseDataJSON) RawJSON() string {
	return r.raw
}

// The timestamp indicating when this session started.
//
// Union satisfied by [shared.UnionString] or [shared.UnionFloat].
type UserSessionResponseDataStartedAtUnion interface {
	ImplementsUserSessionResponseDataStartedAtUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*UserSessionResponseDataStartedAtUnion)(nil)).Elem(),
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
type UserSessionResponseDataEndedAtUnion interface {
	ImplementsUserSessionResponseDataEndedAtUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*UserSessionResponseDataEndedAtUnion)(nil)).Elem(),
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
