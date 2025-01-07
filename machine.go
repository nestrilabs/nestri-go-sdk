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

// MachineService contains methods and other services that help with interacting
// with the nestri API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewMachineService] method instead.
type MachineService struct {
	Options []option.RequestOption
}

// NewMachineService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewMachineService(opts ...option.RequestOption) (r *MachineService) {
	r = &MachineService{}
	r.Options = opts
	return
}

// Associates a machine with the currently authenticated user's account, enabling
// them to manage and control the machine
func (r *MachineService) New(ctx context.Context, fingerprint string, opts ...option.RequestOption) (res *MachineNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	if fingerprint == "" {
		err = errors.New("missing required fingerprint parameter")
		return
	}
	path := fmt.Sprintf("machines/%s", fingerprint)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

// Fetches detailed information about a specific machine using its unique
// fingerprint derived from the Linux machine ID
func (r *MachineService) Get(ctx context.Context, fingerprint string, opts ...option.RequestOption) (res *MachineGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	if fingerprint == "" {
		err = errors.New("missing required fingerprint parameter")
		return
	}
	path := fmt.Sprintf("machines/%s", fingerprint)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Returns a list of all machines registered to the authenticated user in the
// Nestri network
func (r *MachineService) List(ctx context.Context, opts ...option.RequestOption) (res *MachineListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "machines"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Removes the association between a machine and the authenticated user's account.
// This does not delete the machine itself, but removes the user's ability to
// manage it
func (r *MachineService) Delete(ctx context.Context, fingerprint string, opts ...option.RequestOption) (res *MachineDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	if fingerprint == "" {
		err = errors.New("missing required fingerprint parameter")
		return
	}
	path := fmt.Sprintf("machines/%s", fingerprint)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

type MachineNewResponse struct {
	Data MachineNewResponseData `json:"data,required"`
	JSON machineNewResponseJSON `json:"-"`
}

// machineNewResponseJSON contains the JSON metadata for the struct
// [MachineNewResponse]
type machineNewResponseJSON struct {
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MachineNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r machineNewResponseJSON) RawJSON() string {
	return r.raw
}

type MachineNewResponseData string

const (
	MachineNewResponseDataOk MachineNewResponseData = "ok"
)

func (r MachineNewResponseData) IsKnown() bool {
	switch r {
	case MachineNewResponseDataOk:
		return true
	}
	return false
}

type MachineGetResponse struct {
	// Represents a physical or virtual machine connected to the Nestri network..
	Data MachineGetResponseData `json:"data,required"`
	JSON machineGetResponseJSON `json:"-"`
}

// machineGetResponseJSON contains the JSON metadata for the struct
// [MachineGetResponse]
type machineGetResponseJSON struct {
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MachineGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r machineGetResponseJSON) RawJSON() string {
	return r.raw
}

// Represents a physical or virtual machine connected to the Nestri network..
type MachineGetResponseData struct {
	// Unique object identifier. The format and length of IDs may change over time.
	ID string `json:"id,required"`
	// Represents a machine running on the Nestri network, containing its identifying
	// information and metadata.
	CreatedAt MachineGetResponseDataCreatedAtUnion `json:"createdAt,required"`
	// A unique identifier derived from the machine's Linux machine ID.
	Fingerprint string `json:"fingerprint,required"`
	// The Linux hostname that identifies this machine
	Hostname string                     `json:"hostname,required"`
	JSON     machineGetResponseDataJSON `json:"-"`
}

// machineGetResponseDataJSON contains the JSON metadata for the struct
// [MachineGetResponseData]
type machineGetResponseDataJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	Fingerprint apijson.Field
	Hostname    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MachineGetResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r machineGetResponseDataJSON) RawJSON() string {
	return r.raw
}

// Represents a machine running on the Nestri network, containing its identifying
// information and metadata.
//
// Union satisfied by [shared.UnionString] or [shared.UnionFloat].
type MachineGetResponseDataCreatedAtUnion interface {
	ImplementsMachineGetResponseDataCreatedAtUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*MachineGetResponseDataCreatedAtUnion)(nil)).Elem(),
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

type MachineListResponse struct {
	// A list of machines associated with the user
	Data []MachineListResponseData `json:"data,required"`
	JSON machineListResponseJSON   `json:"-"`
}

// machineListResponseJSON contains the JSON metadata for the struct
// [MachineListResponse]
type machineListResponseJSON struct {
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MachineListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r machineListResponseJSON) RawJSON() string {
	return r.raw
}

// Represents a physical or virtual machine connected to the Nestri network..
type MachineListResponseData struct {
	// Unique object identifier. The format and length of IDs may change over time.
	ID string `json:"id,required"`
	// Represents a machine running on the Nestri network, containing its identifying
	// information and metadata.
	CreatedAt MachineListResponseDataCreatedAtUnion `json:"createdAt,required"`
	// A unique identifier derived from the machine's Linux machine ID.
	Fingerprint string `json:"fingerprint,required"`
	// The Linux hostname that identifies this machine
	Hostname string                      `json:"hostname,required"`
	JSON     machineListResponseDataJSON `json:"-"`
}

// machineListResponseDataJSON contains the JSON metadata for the struct
// [MachineListResponseData]
type machineListResponseDataJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	Fingerprint apijson.Field
	Hostname    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MachineListResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r machineListResponseDataJSON) RawJSON() string {
	return r.raw
}

// Represents a machine running on the Nestri network, containing its identifying
// information and metadata.
//
// Union satisfied by [shared.UnionString] or [shared.UnionFloat].
type MachineListResponseDataCreatedAtUnion interface {
	ImplementsMachineListResponseDataCreatedAtUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*MachineListResponseDataCreatedAtUnion)(nil)).Elem(),
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

type MachineDeleteResponse struct {
	Data MachineDeleteResponseData `json:"data,required"`
	JSON machineDeleteResponseJSON `json:"-"`
}

// machineDeleteResponseJSON contains the JSON metadata for the struct
// [MachineDeleteResponse]
type machineDeleteResponseJSON struct {
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MachineDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r machineDeleteResponseJSON) RawJSON() string {
	return r.raw
}

type MachineDeleteResponseData string

const (
	MachineDeleteResponseDataOk MachineDeleteResponseData = "ok"
)

func (r MachineDeleteResponseData) IsKnown() bool {
	switch r {
	case MachineDeleteResponseDataOk:
		return true
	}
	return false
}
