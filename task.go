// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package minimaltodo

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/nlsandler/minimal-todo-go/internal/apijson"
	"github.com/nlsandler/minimal-todo-go/internal/apiquery"
	"github.com/nlsandler/minimal-todo-go/internal/requestconfig"
	"github.com/nlsandler/minimal-todo-go/option"
	"github.com/nlsandler/minimal-todo-go/packages/param"
	"github.com/nlsandler/minimal-todo-go/packages/respjson"
	"github.com/nlsandler/minimal-todo-go/shared"
)

// TaskService contains methods and other services that help with interacting with
// the minimal-todo API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewTaskService] method instead.
type TaskService struct {
	Options []option.RequestOption
	Tags    TaskTagService
}

// NewTaskService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewTaskService(opts ...option.RequestOption) (r TaskService) {
	r = TaskService{}
	r.Options = opts
	r.Tags = NewTaskTagService(opts...)
	return
}

func (r *TaskService) New(ctx context.Context, body TaskNewParams, opts ...option.RequestOption) (res *Task, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/tasks"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

func (r *TaskService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *Task, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("v1/tasks/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

func (r *TaskService) Update(ctx context.Context, id string, body TaskUpdateParams, opts ...option.RequestOption) (res *Task, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("v1/tasks/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

func (r *TaskService) List(ctx context.Context, query TaskListParams, opts ...option.RequestOption) (res *TaskListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/tasks"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

func (r *TaskService) Delete(ctx context.Context, id string, opts ...option.RequestOption) (res *TaskDeleteResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("v1/tasks/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

func (r *TaskService) Complete(ctx context.Context, id string, opts ...option.RequestOption) (res *Task, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("v1/tasks/%s/complete", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

type Task struct {
	Deadline    time.Time        `json:"deadline,required" format:"date"`
	Name        string           `json:"name,required"`
	Tags        []TaskTag        `json:"tags,required"`
	ExtraFields map[string]int64 `json:",extras"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Deadline    respjson.Field
		Name        respjson.Field
		Tags        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Task) RawJSON() string { return r.JSON.raw }
func (r *Task) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TaskTag struct {
	ID        string `json:"id,required"`
	CreatedAt string `json:"created_at"`
	Label     string `json:"label"`
	OwnerID   string `json:"owner_id"`
	UpdatedAt string `json:"updated_at"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		CreatedAt   respjson.Field
		Label       respjson.Field
		OwnerID     respjson.Field
		UpdatedAt   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TaskTag) RawJSON() string { return r.JSON.raw }
func (r *TaskTag) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TaskListResponse struct {
	Data       []Task `json:"data,required"`
	HasMore    bool   `json:"has_more,required"`
	NextCursor string `json:"next_cursor,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		HasMore     respjson.Field
		NextCursor  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TaskListResponse) RawJSON() string { return r.JSON.raw }
func (r *TaskListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TaskDeleteResponse struct {
	ID      string `json:"id,required"`
	Deleted bool   `json:"deleted,required"`
	Note    string `json:"note,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Deleted     respjson.Field
		Note        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TaskDeleteResponse) RawJSON() string { return r.JSON.raw }
func (r *TaskDeleteResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TaskNewParams struct {
	Deadline param.Opt[time.Time] `json:"deadline,omitzero,required" format:"date"`
	Name     param.Opt[string]    `json:"name,omitzero"`
	paramObj
}

func (r TaskNewParams) MarshalJSON() (data []byte, err error) {
	type shadow TaskNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TaskNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TaskUpdateParams struct {
	CompletedAt param.Opt[string] `json:"completed_at,omitzero"`
	Description param.Opt[string] `json:"description,omitzero"`
	Name        param.Opt[string] `json:"name,omitzero"`
	Title       param.Opt[string] `json:"title,omitzero"`
	TagIDs      []string          `json:"tag_ids,omitzero"`
	Tags        []shared.TagParam `json:"tags,omitzero"`
	paramObj
}

func (r TaskUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow TaskUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TaskUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TaskListParams struct {
	Limit  param.Opt[float64] `query:"limit,omitzero" json:"-"`
	Cursor param.Opt[string]  `query:"cursor,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [TaskListParams]'s query parameters as `url.Values`.
func (r TaskListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
