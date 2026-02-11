// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package sfcnodes

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"

	"github.com/sfcompute/nodes-go/internal/apijson"
	"github.com/sfcompute/nodes-go/internal/requestconfig"
	"github.com/sfcompute/nodes-go/option"
	"github.com/sfcompute/nodes-go/packages/respjson"
)

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

// List all VM Images for the authenticated account
func (r *VMImageService) List(ctx context.Context, opts ...option.RequestOption) (res *VMImageListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/vms/images"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Get the download URL for a VM image by ID
func (r *VMImageService) Get(ctx context.Context, imageID string, opts ...option.RequestOption) (res *VMImageGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if imageID == "" {
		err = errors.New("missing required image_id parameter")
		return
	}
	path := fmt.Sprintf("v1/vms/images/%s", imageID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type VMImageListResponse struct {
	Data    []VMImageListResponseData `json:"data,required"`
	HasMore bool                      `json:"has_more,required"`
	// Any of "list".
	Object VMImageListResponseObject `json:"object,required"`
	// Opaque cursor for pagination. Pass as `starting_after` to get the next page.
	Cursor string `json:"cursor,nullable"`
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
	// Unique identifier with prefix 'image\_'.
	ID string `json:"id,required"`
	// Unix timestamp in seconds since epoch
	CreatedAt int64 `json:"created_at,required"`
	// A validated resource name. Must start with alphanumeric, followed by
	// alphanumeric, '.', '\_', or '-'. Max 255 characters.
	Name string `json:"name,required"`
	// Any of "started", "uploading", "completed", "failed".
	UploadStatus string `json:"upload_status,required"`
	// Any of "image".
	Object     string `json:"object"`
	Sha256Hash string `json:"sha256_hash,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID           respjson.Field
		CreatedAt    respjson.Field
		Name         respjson.Field
		UploadStatus respjson.Field
		Object       respjson.Field
		Sha256Hash   respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r VMImageListResponseData) RawJSON() string { return r.JSON.raw }
func (r *VMImageListResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type VMImageListResponseObject string

const (
	VMImageListResponseObjectList VMImageListResponseObject = "list"
)

type VMImageGetResponse struct {
	DownloadURL string `json:"download_url,required"`
	// Unix timestamp in seconds since epoch
	ExpiresAt  int64  `json:"expires_at,required"`
	ObjectSize int64  `json:"object_size,required"`
	Sha256Hash string `json:"sha256_hash,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DownloadURL respjson.Field
		ExpiresAt   respjson.Field
		ObjectSize  respjson.Field
		Sha256Hash  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r VMImageGetResponse) RawJSON() string { return r.JSON.raw }
func (r *VMImageGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
