// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package shared

import (
	"encoding/json"

	"github.com/nlsandler/minimal-todo-go/internal/apijson"
	"github.com/nlsandler/minimal-todo-go/packages/param"
	"github.com/nlsandler/minimal-todo-go/packages/respjson"
)

// aliased to make [param.APIUnion] private when embedding
type paramUnion = param.APIUnion

// aliased to make [param.APIObject] private when embedding
type paramObj = param.APIObject

type Tag struct {
	ID          string         `json:"id,required"`
	CreatedAt   string         `json:"created_at,required"`
	Label       string         `json:"label,required"`
	OwnerID     string         `json:"owner_id,required"`
	UpdatedAt   string         `json:"updated_at,required"`
	ExtraFields map[string]any `json:",extras"`
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

// ToParam converts this Tag to a TagParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// TagParam.Overrides()
func (r Tag) ToParam() TagParam {
	return param.Override[TagParam](json.RawMessage(r.RawJSON()))
}

// The properties ID, CreatedAt, Label, OwnerID, UpdatedAt are required.
type TagParam struct {
	ID          string         `json:"id,required"`
	CreatedAt   string         `json:"created_at,required"`
	Label       string         `json:"label,required"`
	OwnerID     string         `json:"owner_id,required"`
	UpdatedAt   string         `json:"updated_at,required"`
	ExtraFields map[string]any `json:"-"`
	paramObj
}

func (r TagParam) MarshalJSON() (data []byte, err error) {
	type shadow TagParam
	return param.MarshalWithExtras(r, (*shadow)(&r), r.ExtraFields)
}
func (r *TagParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
