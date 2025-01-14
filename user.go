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

// Returns the current authenticate user's profile
func (r *UserService) Get(ctx context.Context, opts ...option.RequestOption) (res *UserGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "users/@me"
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
