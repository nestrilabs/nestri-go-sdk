// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package nestri

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/nestrilabs/nestri-go-sdk/internal/apijson"
	"github.com/nestrilabs/nestri-go-sdk/internal/requestconfig"
	"github.com/nestrilabs/nestri-go-sdk/option"
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

// Get the machine with the given ID.
func (r *MachineService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *MachineGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("machine/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// List the current user's machines.
func (r *MachineService) List(ctx context.Context, opts ...option.RequestOption) (res *MachineListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "machine"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Delete the machine with the given ID.
func (r *MachineService) Delete(ctx context.Context, id string, opts ...option.RequestOption) (res *MachineDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("machine/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

type MachineGetResponse struct {
	// A machine running on the Nestri network.
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

// A machine running on the Nestri network.
type MachineGetResponseData struct {
	// Unique object identifier. The format and length of IDs may change over time.
	ID string `json:"id,required"`
	// The machine's fingerprint, derived from the machine's Linux machine ID.
	Fingerprint string `json:"fingerprint,required"`
	// Hostname of the machine
	Hostname string `json:"hostname,required"`
	// The machine's approximate location; country and continent.
	Location string                     `json:"location,required"`
	JSON     machineGetResponseDataJSON `json:"-"`
}

// machineGetResponseDataJSON contains the JSON metadata for the struct
// [MachineGetResponseData]
type machineGetResponseDataJSON struct {
	ID          apijson.Field
	Fingerprint apijson.Field
	Hostname    apijson.Field
	Location    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MachineGetResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r machineGetResponseDataJSON) RawJSON() string {
	return r.raw
}

type MachineListResponse struct {
	// List of machines.
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

// A machine running on the Nestri network.
type MachineListResponseData struct {
	// Unique object identifier. The format and length of IDs may change over time.
	ID string `json:"id,required"`
	// The machine's fingerprint, derived from the machine's Linux machine ID.
	Fingerprint string `json:"fingerprint,required"`
	// Hostname of the machine
	Hostname string `json:"hostname,required"`
	// The machine's approximate location; country and continent.
	Location string                      `json:"location,required"`
	JSON     machineListResponseDataJSON `json:"-"`
}

// machineListResponseDataJSON contains the JSON metadata for the struct
// [MachineListResponseData]
type machineListResponseDataJSON struct {
	ID          apijson.Field
	Fingerprint apijson.Field
	Hostname    apijson.Field
	Location    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MachineListResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r machineListResponseDataJSON) RawJSON() string {
	return r.raw
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
