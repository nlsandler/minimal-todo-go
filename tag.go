// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package minimaltodo

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"

	"github.com/nlsandler/minimal-todo-go/internal/apijson"
	"github.com/nlsandler/minimal-todo-go/internal/apiquery"
	"github.com/nlsandler/minimal-todo-go/internal/requestconfig"
	"github.com/nlsandler/minimal-todo-go/option"
	"github.com/nlsandler/minimal-todo-go/packages/param"
	"github.com/nlsandler/minimal-todo-go/packages/respjson"
)

// TagService contains methods and other services that help with interacting with
// the minimal-todo API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewTagService] method instead.
type TagService struct {
	Options []option.RequestOption
}

// NewTagService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewTagService(opts ...option.RequestOption) (r TagService) {
	r = TagService{}
	r.Options = opts
	return
}

func (r *TagService) New(ctx context.Context, body TagNewParams, opts ...option.RequestOption) (res *Tag, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/tags"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

func (r *TagService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *Tag, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("v1/tags/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

func (r *TagService) List(ctx context.Context, query TagListParams, opts ...option.RequestOption) (res *TagListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/tags"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

func (r *TagService) Delete(ctx context.Context, id string, opts ...option.RequestOption) (res *TagDeleteResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("v1/tags/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

type Tag struct {
	ID        string `json:"id,required"`
	CreatedAt string `json:"created_at,required"`
	Label     string `json:"label,required"`
	OwnerID   string `json:"owner_id,required"`
	UpdatedAt string `json:"updated_at,required"`
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
func (r Tag) RawJSON() string { return r.JSON.raw }
func (r *Tag) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TagListResponse struct {
	Data       []Tag  `json:"data,required"`
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
func (r TagListResponse) RawJSON() string { return r.JSON.raw }
func (r *TagListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TagDeleteResponse struct {
	ID      string `json:"id,required"`
	Deleted bool   `json:"deleted,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Deleted     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TagDeleteResponse) RawJSON() string { return r.JSON.raw }
func (r *TagDeleteResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TagNewParams struct {
	Label   string `json:"label,required"`
	OwnerID string `json:"owner_id,required"`
	paramObj
}

func (r TagNewParams) MarshalJSON() (data []byte, err error) {
	type shadow TagNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TagNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TagListParams struct {
	Limit  param.Opt[float64] `query:"limit,omitzero" json:"-"`
	Cursor param.Opt[string]  `query:"cursor,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [TagListParams]'s query parameters as `url.Values`.
func (r TagListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
