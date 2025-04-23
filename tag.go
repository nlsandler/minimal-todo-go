// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package minimaltodo

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"

	"github.com/stainless-sdks/minimal-todo-go/internal/apijson"
	"github.com/stainless-sdks/minimal-todo-go/internal/apiquery"
	"github.com/stainless-sdks/minimal-todo-go/internal/requestconfig"
	"github.com/stainless-sdks/minimal-todo-go/option"
	"github.com/stainless-sdks/minimal-todo-go/packages/param"
	"github.com/stainless-sdks/minimal-todo-go/packages/resp"
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
	opts = append(r.Options[:], opts...)
	path := "v1/tags"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

func (r *TagService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *Tag, err error) {
	opts = append(r.Options[:], opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("v1/tags/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

func (r *TagService) List(ctx context.Context, query TagListParams, opts ...option.RequestOption) (res *TagListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/tags"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

func (r *TagService) Delete(ctx context.Context, id string, opts ...option.RequestOption) (res *TagDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
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
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		ID          resp.Field
		CreatedAt   resp.Field
		Label       resp.Field
		OwnerID     resp.Field
		UpdatedAt   resp.Field
		ExtraFields map[string]resp.Field
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
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		Data        resp.Field
		HasMore     resp.Field
		NextCursor  resp.Field
		ExtraFields map[string]resp.Field
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
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		ID          resp.Field
		Deleted     resp.Field
		ExtraFields map[string]resp.Field
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

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f TagNewParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }

func (r TagNewParams) MarshalJSON() (data []byte, err error) {
	type shadow TagNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}

type TagListParams struct {
	Limit  param.Opt[float64] `query:"limit,omitzero" json:"-"`
	Cursor param.Opt[string]  `query:"cursor,omitzero" json:"-"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f TagListParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }

// URLQuery serializes [TagListParams]'s query parameters as `url.Values`.
func (r TagListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
