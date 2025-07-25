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

// NodeService contains methods and other services that help with interacting with
// the sfc-nodes API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewNodeService] method instead.
type NodeService struct {
	Options []option.RequestOption
}

// NewNodeService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewNodeService(opts ...option.RequestOption) (r NodeService) {
	r = NodeService{}
	r.Options = opts
	return
}

// Create VM nodes
func (r *NodeService) New(ctx context.Context, body NodeNewParams, opts ...option.RequestOption) (res *NodeNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/nodes"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// List all VM nodes for the authenticated account
func (r *NodeService) List(ctx context.Context, opts ...option.RequestOption) (res *NodeListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/nodes"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Purchase additional time to extend the end time of a reserved VM node
func (r *NodeService) Extend(ctx context.Context, id string, body NodeExtendParams, opts ...option.RequestOption) (res *Node, err error) {
	opts = append(r.Options[:], opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("v1/nodes/%s/extend", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// Release an on-demand VM node from its procurement, reducing the procurement's
// desired quantity by 1
func (r *NodeService) Release(ctx context.Context, id string, opts ...option.RequestOption) (res *Node, err error) {
	opts = append(r.Options[:], opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("v1/nodes/%s/release", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, nil, &res, opts...)
	return
}

type Node struct {
	ID string `json:"id,required"`
	// Any of "H100", "H200".
	GPUType NodeGPUType `json:"gpu_type,required"`
	Name    string      `json:"name,required"`
	// Any of "on_demand", "reserved".
	NodeType NodeType `json:"node_type,required"`
	Object   string   `json:"object,required"`
	Owner    string   `json:"owner,required"`
	// Node Status
	//
	// Any of "pending", "running", "terminated", "failed", "unknown".
	Status NodeStatus `json:"status,required"`
	// Creation time as Unix timestamp in seconds
	CreatedAt int64 `json:"created_at,nullable"`
	// End time as Unix timestamp in seconds
	EndAt int64 `json:"end_at,nullable"`
	// Max price per hour you're willing to pay for a node in cents
	MaxPricePerNodeHour int64  `json:"max_price_per_node_hour,nullable"`
	ProcurementID       string `json:"procurement_id,nullable"`
	// Any of "uninitialized", "active", "ended", "awaiting_capacity".
	ProcurementStatus NodeProcurementStatus `json:"procurement_status,nullable"`
	// Start time as Unix timestamp in seconds
	StartAt int64 `json:"start_at,nullable"`
	// Last updated time as Unix timestamp in seconds
	UpdatedAt int64 `json:"updated_at,nullable"`
	// Choose from these zones when creating a node
	//
	// Any of "hayesvalley".
	Zone NodeZone `json:"zone,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                  respjson.Field
		GPUType             respjson.Field
		Name                respjson.Field
		NodeType            respjson.Field
		Object              respjson.Field
		Owner               respjson.Field
		Status              respjson.Field
		CreatedAt           respjson.Field
		EndAt               respjson.Field
		MaxPricePerNodeHour respjson.Field
		ProcurementID       respjson.Field
		ProcurementStatus   respjson.Field
		StartAt             respjson.Field
		UpdatedAt           respjson.Field
		Zone                respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Node) RawJSON() string { return r.JSON.raw }
func (r *Node) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type NodeGPUType string

const (
	NodeGPUTypeH100 NodeGPUType = "H100"
	NodeGPUTypeH200 NodeGPUType = "H200"
)

// Node Status
type NodeStatus string

const (
	NodeStatusPending    NodeStatus = "pending"
	NodeStatusRunning    NodeStatus = "running"
	NodeStatusTerminated NodeStatus = "terminated"
	NodeStatusFailed     NodeStatus = "failed"
	NodeStatusUnknown    NodeStatus = "unknown"
)

type NodeProcurementStatus string

const (
	NodeProcurementStatusUninitialized    NodeProcurementStatus = "uninitialized"
	NodeProcurementStatusActive           NodeProcurementStatus = "active"
	NodeProcurementStatusEnded            NodeProcurementStatus = "ended"
	NodeProcurementStatusAwaitingCapacity NodeProcurementStatus = "awaiting_capacity"
)

// Choose from these zones when creating a node
type NodeZone string

const (
	NodeZoneHayesvalley NodeZone = "hayesvalley"
)

type NodeType string

const (
	NodeTypeOnDemand NodeType = "on_demand"
	NodeTypeReserved NodeType = "reserved"
)

type NodeNewResponse struct {
	Data   []NodeNewResponseData `json:"data,required"`
	Object string                `json:"object,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Object      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NodeNewResponse) RawJSON() string { return r.JSON.raw }
func (r *NodeNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type NodeNewResponseData struct {
	ID string `json:"id,required"`
	// Any of "H100", "H200".
	GPUType string `json:"gpu_type,required"`
	Name    string `json:"name,required"`
	// Any of "on_demand", "reserved".
	NodeType NodeType `json:"node_type,required"`
	Object   string   `json:"object,required"`
	Owner    string   `json:"owner,required"`
	// Node Status
	//
	// Any of "pending", "running", "terminated", "failed", "unknown".
	Status string `json:"status,required"`
	// Creation time as Unix timestamp in seconds
	CreatedAt int64 `json:"created_at,nullable"`
	// End time as Unix timestamp in seconds
	EndAt int64 `json:"end_at,nullable"`
	// Max price per hour you're willing to pay for a node in cents
	MaxPricePerNodeHour int64  `json:"max_price_per_node_hour,nullable"`
	ProcurementID       string `json:"procurement_id,nullable"`
	// Any of "uninitialized", "active", "ended", "awaiting_capacity".
	ProcurementStatus string `json:"procurement_status,nullable"`
	// Start time as Unix timestamp in seconds
	StartAt int64 `json:"start_at,nullable"`
	// Last updated time as Unix timestamp in seconds
	UpdatedAt int64 `json:"updated_at,nullable"`
	// Choose from these zones when creating a node
	//
	// Any of "hayesvalley".
	Zone string `json:"zone,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                  respjson.Field
		GPUType             respjson.Field
		Name                respjson.Field
		NodeType            respjson.Field
		Object              respjson.Field
		Owner               respjson.Field
		Status              respjson.Field
		CreatedAt           respjson.Field
		EndAt               respjson.Field
		MaxPricePerNodeHour respjson.Field
		ProcurementID       respjson.Field
		ProcurementStatus   respjson.Field
		StartAt             respjson.Field
		UpdatedAt           respjson.Field
		Zone                respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NodeNewResponseData) RawJSON() string { return r.JSON.raw }
func (r *NodeNewResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type NodeListResponse struct {
	Data   []NodeListResponseData `json:"data,required"`
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
func (r NodeListResponse) RawJSON() string { return r.JSON.raw }
func (r *NodeListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type NodeListResponseData struct {
	ID string `json:"id,required"`
	// Any of "H100", "H200".
	GPUType string `json:"gpu_type,required"`
	Name    string `json:"name,required"`
	// Any of "on_demand", "reserved".
	NodeType NodeType `json:"node_type,required"`
	Object   string   `json:"object,required"`
	Owner    string   `json:"owner,required"`
	// Node Status
	//
	// Any of "pending", "running", "terminated", "failed", "unknown".
	Status string `json:"status,required"`
	// Creation time as Unix timestamp in seconds
	CreatedAt int64 `json:"created_at,nullable"`
	// End time as Unix timestamp in seconds
	EndAt int64 `json:"end_at,nullable"`
	// Max price per hour you're willing to pay for a node in cents
	MaxPricePerNodeHour int64  `json:"max_price_per_node_hour,nullable"`
	ProcurementID       string `json:"procurement_id,nullable"`
	// Any of "uninitialized", "active", "ended", "awaiting_capacity".
	ProcurementStatus string `json:"procurement_status,nullable"`
	// Start time as Unix timestamp in seconds
	StartAt int64 `json:"start_at,nullable"`
	// Last updated time as Unix timestamp in seconds
	UpdatedAt int64 `json:"updated_at,nullable"`
	// Choose from these zones when creating a node
	//
	// Any of "hayesvalley".
	Zone string `json:"zone,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                  respjson.Field
		GPUType             respjson.Field
		Name                respjson.Field
		NodeType            respjson.Field
		Object              respjson.Field
		Owner               respjson.Field
		Status              respjson.Field
		CreatedAt           respjson.Field
		EndAt               respjson.Field
		MaxPricePerNodeHour respjson.Field
		ProcurementID       respjson.Field
		ProcurementStatus   respjson.Field
		StartAt             respjson.Field
		UpdatedAt           respjson.Field
		Zone                respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NodeListResponseData) RawJSON() string { return r.JSON.raw }
func (r *NodeListResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type NodeNewParams struct {
	DesiredCount int64 `json:"desired_count,required"`
	// Max price per hour for a node in cents
	MaxPricePerNodeHour int64 `json:"max_price_per_node_hour,required"`
	// Zone to create the nodes in. See Zone enum for valid values.
	Zone string `json:"zone,required"`
	// End time as Unix timestamp in seconds. If provided, end time must be aligned to
	// the hour. If not provided, the node will be created as an on-demand node.
	EndAt param.Opt[int64] `json:"end_at,omitzero"`
	// Start time as Unix timestamp in seconds
	StartAt param.Opt[int64] `json:"start_at,omitzero"`
	// Custom node names. Names cannot follow the vm\_{alpha_numeric_chars} as this is
	// reserved for system-generated IDs. Names cannot be numeric strings.
	Names []string `json:"names,omitzero"`
	// Any of "on_demand", "reserved".
	NodeType NodeType `json:"node_type,omitzero"`
	paramObj
}

func (r NodeNewParams) MarshalJSON() (data []byte, err error) {
	type shadow NodeNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *NodeNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type NodeExtendParams struct {
	// Duration in seconds to extend the node Must be at least 1 hour (3600 seconds)
	// and a multiple of 1 hour.
	DurationSeconds int64 `json:"duration_seconds,required"`
	// Max price per hour for the extension in cents
	MaxPricePerNodeHour int64 `json:"max_price_per_node_hour,required"`
	paramObj
}

func (r NodeExtendParams) MarshalJSON() (data []byte, err error) {
	type shadow NodeExtendParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *NodeExtendParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
