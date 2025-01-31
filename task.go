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

// TaskService contains methods and other services that help with interacting with
// the nestri API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewTaskService] method instead.
type TaskService struct {
	Options []option.RequestOption
}

// NewTaskService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewTaskService(opts ...option.RequestOption) (r *TaskService) {
	r = &TaskService{}
	r.Options = opts
	return
}

// Create a task
func (r *TaskService) New(ctx context.Context, opts ...option.RequestOption) (res *TaskNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "tasks"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

// Get a task by its id
func (r *TaskService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *TaskGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("tasks/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Updates the metadata about a task by querying remote task
func (r *TaskService) Update(ctx context.Context, id string, opts ...option.RequestOption) (res *TaskUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("tasks/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, nil, &res, opts...)
	return
}

// List all tasks by this user
func (r *TaskService) List(ctx context.Context, opts ...option.RequestOption) (res *TaskListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "tasks"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Stop a running task by its id
func (r *TaskService) Delete(ctx context.Context, id string, opts ...option.RequestOption) (res *TaskDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("tasks/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Get a task by its id
func (r *TaskService) Session(ctx context.Context, id string, opts ...option.RequestOption) (res *TaskSessionResponse, err error) {
	opts = append(r.Options[:], opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("tasks/%s/session", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type TaskNewResponse struct {
	// The id of the task created
	Data string              `json:"data,required"`
	JSON taskNewResponseJSON `json:"-"`
}

// taskNewResponseJSON contains the JSON metadata for the struct [TaskNewResponse]
type taskNewResponseJSON struct {
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TaskNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r taskNewResponseJSON) RawJSON() string {
	return r.raw
}

type TaskGetResponse struct {
	// Subscription to a Nestri product.
	Data TaskGetResponseData `json:"data,required"`
	JSON taskGetResponseJSON `json:"-"`
}

// taskGetResponseJSON contains the JSON metadata for the struct [TaskGetResponse]
type taskGetResponseJSON struct {
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TaskGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r taskGetResponseJSON) RawJSON() string {
	return r.raw
}

// Subscription to a Nestri product.
type TaskGetResponseData struct {
	// Unique object identifier. The format and length of IDs may change over time.
	ID string `json:"id,required"`
	// The polar.sh checkout id
	CheckoutID string `json:"checkoutID,required"`
	// Cancelled date for the subscription.
	CanceledAt TaskGetResponseDataCanceledAtUnion `json:"canceledAt"`
	JSON       taskGetResponseDataJSON            `json:"-"`
}

// taskGetResponseDataJSON contains the JSON metadata for the struct
// [TaskGetResponseData]
type taskGetResponseDataJSON struct {
	ID          apijson.Field
	CheckoutID  apijson.Field
	CanceledAt  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TaskGetResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r taskGetResponseDataJSON) RawJSON() string {
	return r.raw
}

// Cancelled date for the subscription.
//
// Union satisfied by [shared.UnionString] or [shared.UnionFloat].
type TaskGetResponseDataCanceledAtUnion interface {
	ImplementsTaskGetResponseDataCanceledAtUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*TaskGetResponseDataCanceledAtUnion)(nil)).Elem(),
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

type TaskUpdateResponse struct {
	// Subscription to a Nestri product.
	Data TaskUpdateResponseData `json:"data,required"`
	JSON taskUpdateResponseJSON `json:"-"`
}

// taskUpdateResponseJSON contains the JSON metadata for the struct
// [TaskUpdateResponse]
type taskUpdateResponseJSON struct {
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TaskUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r taskUpdateResponseJSON) RawJSON() string {
	return r.raw
}

// Subscription to a Nestri product.
type TaskUpdateResponseData struct {
	// Unique object identifier. The format and length of IDs may change over time.
	ID string `json:"id,required"`
	// The polar.sh checkout id
	CheckoutID string `json:"checkoutID,required"`
	// Cancelled date for the subscription.
	CanceledAt TaskUpdateResponseDataCanceledAtUnion `json:"canceledAt"`
	JSON       taskUpdateResponseDataJSON            `json:"-"`
}

// taskUpdateResponseDataJSON contains the JSON metadata for the struct
// [TaskUpdateResponseData]
type taskUpdateResponseDataJSON struct {
	ID          apijson.Field
	CheckoutID  apijson.Field
	CanceledAt  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TaskUpdateResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r taskUpdateResponseDataJSON) RawJSON() string {
	return r.raw
}

// Cancelled date for the subscription.
//
// Union satisfied by [shared.UnionString] or [shared.UnionFloat].
type TaskUpdateResponseDataCanceledAtUnion interface {
	ImplementsTaskUpdateResponseDataCanceledAtUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*TaskUpdateResponseDataCanceledAtUnion)(nil)).Elem(),
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

type TaskListResponse struct {
	// Subscription to a Nestri product.
	Data TaskListResponseData `json:"data,required"`
	JSON taskListResponseJSON `json:"-"`
}

// taskListResponseJSON contains the JSON metadata for the struct
// [TaskListResponse]
type taskListResponseJSON struct {
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TaskListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r taskListResponseJSON) RawJSON() string {
	return r.raw
}

// Subscription to a Nestri product.
type TaskListResponseData struct {
	// Unique object identifier. The format and length of IDs may change over time.
	ID string `json:"id,required"`
	// The polar.sh checkout id
	CheckoutID string `json:"checkoutID,required"`
	// Cancelled date for the subscription.
	CanceledAt TaskListResponseDataCanceledAtUnion `json:"canceledAt"`
	JSON       taskListResponseDataJSON            `json:"-"`
}

// taskListResponseDataJSON contains the JSON metadata for the struct
// [TaskListResponseData]
type taskListResponseDataJSON struct {
	ID          apijson.Field
	CheckoutID  apijson.Field
	CanceledAt  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TaskListResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r taskListResponseDataJSON) RawJSON() string {
	return r.raw
}

// Cancelled date for the subscription.
//
// Union satisfied by [shared.UnionString] or [shared.UnionFloat].
type TaskListResponseDataCanceledAtUnion interface {
	ImplementsTaskListResponseDataCanceledAtUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*TaskListResponseDataCanceledAtUnion)(nil)).Elem(),
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

type TaskDeleteResponse struct {
	Data TaskDeleteResponseData `json:"data,required"`
	JSON taskDeleteResponseJSON `json:"-"`
}

// taskDeleteResponseJSON contains the JSON metadata for the struct
// [TaskDeleteResponse]
type taskDeleteResponseJSON struct {
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TaskDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r taskDeleteResponseJSON) RawJSON() string {
	return r.raw
}

type TaskDeleteResponseData string

const (
	TaskDeleteResponseDataOk TaskDeleteResponseData = "ok"
)

func (r TaskDeleteResponseData) IsKnown() bool {
	switch r {
	case TaskDeleteResponseDataOk:
		return true
	}
	return false
}

type TaskSessionResponse struct {
	// Represents a single game play session, tracking its lifetime and accessibility
	// settings.
	Data TaskSessionResponseData `json:"data,required"`
	JSON taskSessionResponseJSON `json:"-"`
}

// taskSessionResponseJSON contains the JSON metadata for the struct
// [TaskSessionResponse]
type taskSessionResponseJSON struct {
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TaskSessionResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r taskSessionResponseJSON) RawJSON() string {
	return r.raw
}

// Represents a single game play session, tracking its lifetime and accessibility
// settings.
type TaskSessionResponseData struct {
	// Unique object identifier. The format and length of IDs may change over time.
	ID string `json:"id,required"`
	// If true, the session is publicly viewable by all users. If false, only
	// authorized users can access it
	Public bool `json:"public,required"`
	// The timestamp indicating when this session started.
	StartedAt TaskSessionResponseDataStartedAtUnion `json:"startedAt,required"`
	// The timestamp indicating when this session was completed or terminated. Null if
	// session is still active.
	EndedAt TaskSessionResponseDataEndedAtUnion `json:"endedAt"`
	JSON    taskSessionResponseDataJSON         `json:"-"`
}

// taskSessionResponseDataJSON contains the JSON metadata for the struct
// [TaskSessionResponseData]
type taskSessionResponseDataJSON struct {
	ID          apijson.Field
	Public      apijson.Field
	StartedAt   apijson.Field
	EndedAt     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TaskSessionResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r taskSessionResponseDataJSON) RawJSON() string {
	return r.raw
}

// The timestamp indicating when this session started.
//
// Union satisfied by [shared.UnionString] or [shared.UnionFloat].
type TaskSessionResponseDataStartedAtUnion interface {
	ImplementsTaskSessionResponseDataStartedAtUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*TaskSessionResponseDataStartedAtUnion)(nil)).Elem(),
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
type TaskSessionResponseDataEndedAtUnion interface {
	ImplementsTaskSessionResponseDataEndedAtUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*TaskSessionResponseDataEndedAtUnion)(nil)).Elem(),
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
