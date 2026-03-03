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

// Manage your Virtual Machines.
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

// Response body for listing images
type VMImageListResponse struct {
	Data    []VMImageListResponseData `json:"data" api:"required"`
	HasMore bool                      `json:"has_more" api:"required"`
	// Any of "list".
	Object VMImageListResponseObject `json:"object" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		HasMore     respjson.Field
		Object      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r VMImageListResponse) RawJSON() string { return r.JSON.raw }
func (r *VMImageListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Response body for individual image info (used in lists)
type VMImageListResponseData struct {
	// Creation timestamp as Unix timestamp in seconds
	CreatedAt int64 `json:"created_at" api:"required"`
	// The image ID
	ImageID string `json:"image_id" api:"required"`
	// Client given name of the image. Must be unique per account.
	Name string `json:"name" api:"required"`
	// Any of "image".
	Object string `json:"object" api:"required"`
	// Upload status of the image
	UploadStatus string `json:"upload_status" api:"required"`
	// SHA256 hash of the image file for integrity verification
	Sha256Hash string `json:"sha256_hash" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CreatedAt    respjson.Field
		ImageID      respjson.Field
		Name         respjson.Field
		Object       respjson.Field
		UploadStatus respjson.Field
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

// Response body for image download presigned URL generation
type VMImageGetResponse struct {
	// The presigned URL that can be used to download the image
	DownloadURL string `json:"download_url" api:"required"`
	// Timestamp when the presigned URL expires (RFC 3339 format)
	ExpiresAt string `json:"expires_at" api:"required"`
	// The image ID
	ImageID string `json:"image_id" api:"required"`
	// Human readable name of the image. Must be unique per account.
	Name string `json:"name" api:"required"`
	// Any of "image".
	Object VMImageGetResponseObject `json:"object" api:"required"`
	// Size of the image file in bytes
	ObjectSize int64 `json:"object_size" api:"required"`
	// SHA256 hash of the image file for integrity verification
	Sha256Hash string `json:"sha256_hash" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DownloadURL respjson.Field
		ExpiresAt   respjson.Field
		ImageID     respjson.Field
		Name        respjson.Field
		Object      respjson.Field
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

type VMImageGetResponseObject string

const (
	VMImageGetResponseObjectImage VMImageGetResponseObject = "image"
)
