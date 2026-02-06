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

// ZoneService contains methods and other services that help with interacting with
// the sfc-nodes API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewZoneService] method instead.
type ZoneService struct {
	Options []option.RequestOption
}

// NewZoneService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewZoneService(opts ...option.RequestOption) (r ZoneService) {
	r = ZoneService{}
	r.Options = opts
	return
}

// List all available zones
func (r *ZoneService) List(ctx context.Context, opts ...option.RequestOption) (res *ZoneListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v0/zones"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Get detailed information about a specific zone
func (r *ZoneService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *ZoneGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("v0/zones/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type ZoneListResponse struct {
	Data   []ZoneListResponseData `json:"data,required"`
	Object string                 `json:"object,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Object      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ZoneListResponse) RawJSON() string { return r.JSON.raw }
func (r *ZoneListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneListResponseData struct {
	// The available capacity on this cluster, in the shape of consecutive
	// "availability rectangles".
	AvailableCapacity []ZoneListResponseDataAvailableCapacity `json:"available_capacity,required"`
	// Any of "K8s", "VM".
	DeliveryType string `json:"delivery_type,required"`
	// Any of "H100", "H200".
	HardwareType AcceleratorType `json:"hardware_type,required"`
	// Any of "Infiniband", "None".
	InterconnectType string `json:"interconnect_type,required"`
	Name             string `json:"name,required"`
	Object           string `json:"object,required"`
	// Any of "NorthAmerica", "AsiaPacific", "EuropeMiddleEastAfrica".
	Region string `json:"region,required"`
	// User-facing zone name (e.g., "Hayes Valley", "Land's End")
	DisplayName string `json:"display_name,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AvailableCapacity respjson.Field
		DeliveryType      respjson.Field
		HardwareType      respjson.Field
		InterconnectType  respjson.Field
		Name              respjson.Field
		Object            respjson.Field
		Region            respjson.Field
		DisplayName       respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ZoneListResponseData) RawJSON() string { return r.JSON.raw }
func (r *ZoneListResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneListResponseDataAvailableCapacity struct {
	// Unix timestamp in seconds since epoch
	EndTimestamp int64 `json:"end_timestamp,required"`
	// The number of nodes available during this time period
	Quantity int64 `json:"quantity,required"`
	// Unix timestamp in seconds since epoch
	StartTimestamp int64 `json:"start_timestamp,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		EndTimestamp   respjson.Field
		Quantity       respjson.Field
		StartTimestamp respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ZoneListResponseDataAvailableCapacity) RawJSON() string { return r.JSON.raw }
func (r *ZoneListResponseDataAvailableCapacity) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneGetResponse struct {
	// The available capacity on this cluster, in the shape of consecutive
	// "availability rectangles".
	AvailableCapacity []ZoneGetResponseAvailableCapacity `json:"available_capacity,required"`
	// Any of "K8s", "VM".
	DeliveryType ZoneGetResponseDeliveryType `json:"delivery_type,required"`
	// Any of "H100", "H200".
	HardwareType AcceleratorType `json:"hardware_type,required"`
	// Any of "Infiniband", "None".
	InterconnectType ZoneGetResponseInterconnectType `json:"interconnect_type,required"`
	Name             string                          `json:"name,required"`
	Object           string                          `json:"object,required"`
	// Any of "NorthAmerica", "AsiaPacific", "EuropeMiddleEastAfrica".
	Region ZoneGetResponseRegion `json:"region,required"`
	// User-facing zone name (e.g., "Hayes Valley", "Land's End")
	DisplayName string `json:"display_name,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AvailableCapacity respjson.Field
		DeliveryType      respjson.Field
		HardwareType      respjson.Field
		InterconnectType  respjson.Field
		Name              respjson.Field
		Object            respjson.Field
		Region            respjson.Field
		DisplayName       respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ZoneGetResponse) RawJSON() string { return r.JSON.raw }
func (r *ZoneGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneGetResponseAvailableCapacity struct {
	// Unix timestamp in seconds since epoch
	EndTimestamp int64 `json:"end_timestamp,required"`
	// The number of nodes available during this time period
	Quantity int64 `json:"quantity,required"`
	// Unix timestamp in seconds since epoch
	StartTimestamp int64 `json:"start_timestamp,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		EndTimestamp   respjson.Field
		Quantity       respjson.Field
		StartTimestamp respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ZoneGetResponseAvailableCapacity) RawJSON() string { return r.JSON.raw }
func (r *ZoneGetResponseAvailableCapacity) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneGetResponseDeliveryType string

const (
	ZoneGetResponseDeliveryTypeK8s ZoneGetResponseDeliveryType = "K8s"
	ZoneGetResponseDeliveryTypeVM  ZoneGetResponseDeliveryType = "VM"
)

type ZoneGetResponseInterconnectType string

const (
	ZoneGetResponseInterconnectTypeInfiniband ZoneGetResponseInterconnectType = "Infiniband"
	ZoneGetResponseInterconnectTypeNone       ZoneGetResponseInterconnectType = "None"
)

type ZoneGetResponseRegion string

const (
	ZoneGetResponseRegionNorthAmerica           ZoneGetResponseRegion = "NorthAmerica"
	ZoneGetResponseRegionAsiaPacific            ZoneGetResponseRegion = "AsiaPacific"
	ZoneGetResponseRegionEuropeMiddleEastAfrica ZoneGetResponseRegion = "EuropeMiddleEastAfrica"
)
