// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package sfcnodes

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"

	"github.com/sfcompute/nodes-go/internal/apijson"
	"github.com/sfcompute/nodes-go/internal/apiquery"
	"github.com/sfcompute/nodes-go/internal/requestconfig"
	"github.com/sfcompute/nodes-go/option"
	"github.com/sfcompute/nodes-go/packages/param"
	"github.com/sfcompute/nodes-go/packages/respjson"
	"github.com/sfcompute/nodes-go/shared/constant"
)

// Custom machine images for instances.
//
// VMImageService contains methods and other services that help with interacting
// with the sfc-nodes API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewVMImageService] method instead.
type VMImageService struct {
	Options []option.RequestOption
}

// NewVMImageService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewVMImageService(opts ...option.RequestOption) (r VMImageService) {
	r = VMImageService{}
	r.Options = opts
	return
}

// > ⚠️ This endpoint is in [public preview](/preview/roadmap).
//
// List images in the specified workspace. Pass `sfc:workspace:sfcompute:public` as
// the workspace to list sfc-provided public images instead.
func (r *VMImageService) List(ctx context.Context, query VMImageListParams, opts ...option.RequestOption) (res *VMImageListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithBaseURL("https://api.sfcompute.com/")}, opts...)
	path := "preview/v2/images"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// > ⚠️ This endpoint is in [public preview](/preview/roadmap).
//
// Retrieve an image by ID. Returns both user-owned and public images.
func (r *VMImageService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *VMImageGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithBaseURL("https://api.sfcompute.com/")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("preview/v2/images/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

type VMImageListResponse struct {
	Data    []VMImageListResponseData `json:"data" api:"required"`
	HasMore bool                      `json:"has_more" api:"required"`
	Object  constant.List             `json:"object" default:"list"`
	Cursor  string                    `json:"cursor" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		HasMore     respjson.Field
		Object      respjson.Field
		Cursor      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r VMImageListResponse) RawJSON() string { return r.JSON.raw }
func (r *VMImageListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type VMImageListResponseData struct {
	ID string `json:"id" api:"required"`
	// Unix timestamp.
	CreatedAt int64          `json:"created_at" api:"required"`
	Name      string         `json:"name" api:"required"`
	Object    constant.Image `json:"object" default:"image"`
	Owner     string         `json:"owner" api:"required"`
	// A resource path for a image resource. Format:
	// sfc:image:<account>:<workspace>:<name>.
	ResourcePath string `json:"resource_path" api:"required"`
	// Any of "started", "uploading", "completed", "failed", "revoked".
	UploadStatus string `json:"upload_status" api:"required"`
	Workspace    string `json:"workspace" api:"required"`
	Provider     string `json:"provider" api:"nullable"`
	Sha256       string `json:"sha256" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID           respjson.Field
		CreatedAt    respjson.Field
		Name         respjson.Field
		Object       respjson.Field
		Owner        respjson.Field
		ResourcePath respjson.Field
		UploadStatus respjson.Field
		Workspace    respjson.Field
		Provider     respjson.Field
		Sha256       respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r VMImageListResponseData) RawJSON() string { return r.JSON.raw }
func (r *VMImageListResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type VMImageGetResponse struct {
	ID string `json:"id" api:"required"`
	// Unix timestamp.
	CreatedAt int64          `json:"created_at" api:"required"`
	Name      string         `json:"name" api:"required"`
	Object    constant.Image `json:"object" default:"image"`
	Owner     string         `json:"owner" api:"required"`
	// A resource path for a image resource. Format:
	// sfc:image:<account>:<workspace>:<name>.
	ResourcePath string `json:"resource_path" api:"required"`
	// Any of "started", "uploading", "completed", "failed", "revoked".
	UploadStatus VMImageGetResponseUploadStatus `json:"upload_status" api:"required"`
	Workspace    string                         `json:"workspace" api:"required"`
	Provider     string                         `json:"provider" api:"nullable"`
	Sha256       string                         `json:"sha256" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID           respjson.Field
		CreatedAt    respjson.Field
		Name         respjson.Field
		Object       respjson.Field
		Owner        respjson.Field
		ResourcePath respjson.Field
		UploadStatus respjson.Field
		Workspace    respjson.Field
		Provider     respjson.Field
		Sha256       respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r VMImageGetResponse) RawJSON() string { return r.JSON.raw }
func (r *VMImageGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type VMImageGetResponseUploadStatus string

const (
	VMImageGetResponseUploadStatusStarted   VMImageGetResponseUploadStatus = "started"
	VMImageGetResponseUploadStatusUploading VMImageGetResponseUploadStatus = "uploading"
	VMImageGetResponseUploadStatusCompleted VMImageGetResponseUploadStatus = "completed"
	VMImageGetResponseUploadStatusFailed    VMImageGetResponseUploadStatus = "failed"
	VMImageGetResponseUploadStatusRevoked   VMImageGetResponseUploadStatus = "revoked"
)

type VMImageListParams struct {
	// Filter by workspace. Pass `sfc:workspace:sfcompute:public` to list sfc-provided
	// public images.
	Workspace string `query:"workspace" api:"required" json:"-"`
	// Cursor for backward pagination.
	EndingBefore param.Opt[string] `query:"ending_before,omitzero" json:"-"`
	// Maximum number of results to return (1-200, default 50).
	Limit param.Opt[int64] `query:"limit,omitzero" format:"u-int32" json:"-"`
	// Cursor for forward pagination (from a previous response's `cursor` field).
	StartingAfter param.Opt[string] `query:"starting_after,omitzero" json:"-"`
	// Filter by image ID (repeatable).
	ID []string `query:"id,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [VMImageListParams]'s query parameters as `url.Values`.
func (r VMImageListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
