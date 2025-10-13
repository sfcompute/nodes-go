// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package sfcnodes

import (
	"context"
	"encoding/json"
	"net/http"
	"slices"

	"github.com/sfcompute/nodes-go/internal/apijson"
	"github.com/sfcompute/nodes-go/internal/requestconfig"
	"github.com/sfcompute/nodes-go/option"
	"github.com/sfcompute/nodes-go/packages/param"
	"github.com/sfcompute/nodes-go/packages/respjson"
)

// VMScriptService contains methods and other services that help with interacting
// with the sfc-nodes API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewVMScriptService] method instead.
type VMScriptService struct {
	Options []option.RequestOption
}

// NewVMScriptService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewVMScriptService(opts ...option.RequestOption) (r VMScriptService) {
	r = VMScriptService{}
	r.Options = opts
	return
}

func (r *VMScriptService) New(ctx context.Context, body VMScriptNewParams, opts ...option.RequestOption) (res *VMScriptNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v0/vms/script"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

func (r *VMScriptService) Get(ctx context.Context, opts ...option.RequestOption) (res *VMScriptGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v0/vms/script"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// UserDataUnion contains all possible properties and values from [string],
// [[]int64].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfIntArray]
type UserDataUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [[]int64] instead of an object.
	OfIntArray []int64 `json:",inline"`
	JSON       struct {
		OfString   respjson.Field
		OfIntArray respjson.Field
		raw        string
	} `json:"-"`
}

func (u UserDataUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UserDataUnion) AsIntArray() (v []int64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u UserDataUnion) RawJSON() string { return u.JSON.raw }

func (r *UserDataUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this UserDataUnion to a UserDataUnionParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// UserDataUnionParam.Overrides()
func (r UserDataUnion) ToParam() UserDataUnionParam {
	return param.Override[UserDataUnionParam](json.RawMessage(r.RawJSON()))
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type UserDataUnionParam struct {
	OfString   param.Opt[string] `json:",omitzero,inline"`
	OfIntArray []int64           `json:",omitzero,inline"`
	paramUnion
}

func (u UserDataUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfIntArray)
}
func (u *UserDataUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *UserDataUnionParam) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfIntArray) {
		return &u.OfIntArray
	}
	return nil
}

type VMScriptNewResponse struct {
	// if the script is valid utf8 then the response may be in either string, or byte
	// form and the client must handle both
	Script UserDataUnion `json:"script,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Script      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r VMScriptNewResponse) RawJSON() string { return r.JSON.raw }
func (r *VMScriptNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type VMScriptGetResponse struct {
	// if the script is valid utf8 then the response may be in either string, or byte
	// form and the client must handle both
	Script UserDataUnion `json:"script,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Script      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r VMScriptGetResponse) RawJSON() string { return r.JSON.raw }
func (r *VMScriptGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type VMScriptNewParams struct {
	// if the script is valid utf8 then the response may be in either string, or byte
	// form and the client must handle both
	Script UserDataUnionParam `json:"script,omitzero,required"`
	paramObj
}

func (r VMScriptNewParams) MarshalJSON() (data []byte, err error) {
	type shadow VMScriptNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *VMScriptNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
