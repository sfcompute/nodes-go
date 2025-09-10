// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package sfcnodes

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/sfcompute/nodes-go/internal/apijson"
	"github.com/sfcompute/nodes-go/internal/requestconfig"
	"github.com/sfcompute/nodes-go/option"
	"github.com/sfcompute/nodes-go/packages/param"
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

func (r *VMImageService) CompleteUpload(ctx context.Context, imageID string, opts ...option.RequestOption) (res *VMImageCompleteUploadResponse, err error) {
	opts = append(r.Options[:], opts...)
	if imageID == "" {
		err = errors.New("missing required image_id parameter")
		return
	}
	path := fmt.Sprintf("v1/vms/images/%s/complete_upload", imageID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, nil, &res, opts...)
	return
}

// Get the download URL for a VM image by ID
func (r *VMImageService) Get(ctx context.Context, imageID string, opts ...option.RequestOption) (res *VMImageGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	if imageID == "" {
		err = errors.New("missing required image_id parameter")
		return
	}
	path := fmt.Sprintf("v1/vms/images/%s", imageID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

func (r *VMImageService) StartUpload(ctx context.Context, body VMImageStartUploadParams, opts ...option.RequestOption) (res *VMImageStartUploadResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/vms/images/start_upload"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

func (r *VMImageService) Upload(ctx context.Context, imageID string, body VMImageUploadParams, opts ...option.RequestOption) (res *VMImageUploadResponse, err error) {
	opts = append(r.Options[:], opts...)
	if imageID == "" {
		err = errors.New("missing required image_id parameter")
		return
	}
	path := fmt.Sprintf("v1/vms/images/%s/upload", imageID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Response body for completing a multipart upload
type VMImageCompleteUploadResponse struct {
	// The image ID
	ImageID string `json:"image_id,required"`
	// Any of "image".
	Object VMImageCompleteUploadResponseObject `json:"object,required"`
	// Status of the upload verification
	UploadStatus string `json:"upload_status,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ImageID      respjson.Field
		Object       respjson.Field
		UploadStatus respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r VMImageCompleteUploadResponse) RawJSON() string { return r.JSON.raw }
func (r *VMImageCompleteUploadResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type VMImageCompleteUploadResponseObject string

const (
	VMImageCompleteUploadResponseObjectImage VMImageCompleteUploadResponseObject = "image"
)

// Response body for image download presigned URL generation
type VMImageGetResponse struct {
	// The presigned URL that can be used to download the image
	DownloadURL string `json:"download_url,required"`
	// Timestamp when the presigned URL expires (RFC 3339 format)
	ExpiresAt string `json:"expires_at,required"`
	// The image ID
	ImageID string `json:"image_id,required"`
	// Human readable name of the image
	Name string `json:"name,required"`
	// Any of "image".
	Object VMImageGetResponseObject `json:"object,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DownloadURL respjson.Field
		ExpiresAt   respjson.Field
		ImageID     respjson.Field
		Name        respjson.Field
		Object      respjson.Field
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

// Response body for starting a multipart upload
type VMImageStartUploadResponse struct {
	// The image ID for the created image
	ImageID string `json:"image_id,required"`
	// Any of "image".
	Object VMImageStartUploadResponseObject `json:"object,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ImageID     respjson.Field
		Object      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r VMImageStartUploadResponse) RawJSON() string { return r.JSON.raw }
func (r *VMImageStartUploadResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type VMImageStartUploadResponseObject string

const (
	VMImageStartUploadResponseObjectImage VMImageStartUploadResponseObject = "image"
)

// Response body for image upload presigned URL generation
type VMImageUploadResponse struct {
	// Timestamp when the presigned URL expires (RFC 3339 format)
	ExpiresAt string `json:"expires_at,required"`
	// The presigned URL that can be used to upload the image part
	UploadURL string `json:"upload_url,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ExpiresAt   respjson.Field
		UploadURL   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r VMImageUploadResponse) RawJSON() string { return r.JSON.raw }
func (r *VMImageUploadResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type VMImageStartUploadParams struct {
	// Name of the image file
	Name string `json:"name,required"`
	paramObj
}

func (r VMImageStartUploadParams) MarshalJSON() (data []byte, err error) {
	type shadow VMImageStartUploadParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *VMImageStartUploadParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type VMImageUploadParams struct {
	// part idx (1-based)
	PartID int64 `json:"part_id,required"`
	paramObj
}

func (r VMImageUploadParams) MarshalJSON() (data []byte, err error) {
	type shadow VMImageUploadParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *VMImageUploadParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
