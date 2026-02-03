// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package minimaltodo

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"

	"github.com/nlsandler/minimal-todo-go/internal/requestconfig"
	"github.com/nlsandler/minimal-todo-go/option"
)

// TaskTagService contains methods and other services that help with interacting
// with the minimal-todo API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewTaskTagService] method instead.
type TaskTagService struct {
	Options []option.RequestOption
}

// NewTaskTagService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewTaskTagService(opts ...option.RequestOption) (r TaskTagService) {
	r = TaskTagService{}
	r.Options = opts
	return
}

func (r *TaskTagService) List(ctx context.Context, todoID int64, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := fmt.Sprintf("v1/tasks/%v/tags", todoID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, nil, opts...)
	return
}

func (r *TaskTagService) Add(ctx context.Context, tagID string, body TaskTagAddParams, opts ...option.RequestOption) (res *Task, err error) {
	opts = slices.Concat(r.Options, opts)
	if body.TodoID == "" {
		err = errors.New("missing required todoId parameter")
		return
	}
	if tagID == "" {
		err = errors.New("missing required tagId parameter")
		return
	}
	path := fmt.Sprintf("v1/tasks/%s/tags/%s", body.TodoID, tagID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

func (r *TaskTagService) Remove(ctx context.Context, tagID string, body TaskTagRemoveParams, opts ...option.RequestOption) (res *Task, err error) {
	opts = slices.Concat(r.Options, opts)
	if body.TodoID == "" {
		err = errors.New("missing required todoId parameter")
		return
	}
	if tagID == "" {
		err = errors.New("missing required tagId parameter")
		return
	}
	path := fmt.Sprintf("v1/tasks/%s/tags/%s", body.TodoID, tagID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

type TaskTagAddParams struct {
	TodoID string `path:"todoId,required" json:"-"`
	paramObj
}

type TaskTagRemoveParams struct {
	TodoID string `path:"todoId,required" json:"-"`
	paramObj
}
