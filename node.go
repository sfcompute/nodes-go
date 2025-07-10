// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package sfcnodes

import (
	"context"
	"encoding/json"
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
func (r *NodeService) New(ctx context.Context, body NodeNewParams, opts ...option.RequestOption) (res *[]Node, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/nodes"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// List all VM nodes for the authenticated account
func (r *NodeService) List(ctx context.Context, opts ...option.RequestOption) (res *[]Node, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/nodes"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Extend the end time of a reservation-based VM node by purchasing additional time
func (r *NodeService) Extend(ctx context.Context, id string, body NodeExtendParams, opts ...option.RequestOption) (res *UpdateNode, err error) {
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
func (r *NodeService) Release(ctx context.Context, id string, body NodeReleaseParams, opts ...option.RequestOption) (res *UpdateNode, err error) {
	opts = append(r.Options[:], opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("v1/nodes/%s/release", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

type Node struct {
	ID string `json:"id,required"`
	// Any of "H100", "H200".
	GPUType NodeGPUType `json:"gpu_type,required"`
	Name    string      `json:"name,required"`
	Node    string      `json:"node,required"`
	// Any of "OnDemand", "Reserved".
	NodeType NodeType `json:"node_type,required"`
	Owner    string   `json:"owner,required"`
	StartAt  int64    `json:"start_at,required"`
	// Any of "Pending", "Running", "Terminated", "Failed", "Unknown".
	Status          NodeStatus `json:"status,required"`
	CreatedAt       int64      `json:"created_at,nullable"`
	EndAt           int64      `json:"end_at,nullable"`
	MaxPricePerHour int64      `json:"max_price_per_hour,nullable"`
	ProcurementID   string     `json:"procurement_id,nullable"`
	// Any of "Uninitialized", "Active", "Ended", "AwaitingCapacity".
	ProcurementStatus NodeProcurementStatus `json:"procurement_status,nullable"`
	UpdatedAt         int64                 `json:"updated_at,nullable"`
	// Possible zones to choose from when creating a node.
	//
	// TODO (ENG-1976): Support dynamic zones through an endpoint.
	//
	// Any of "HayesValley".
	Zone NodeZone `json:"zone,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                respjson.Field
		GPUType           respjson.Field
		Name              respjson.Field
		Node              respjson.Field
		NodeType          respjson.Field
		Owner             respjson.Field
		StartAt           respjson.Field
		Status            respjson.Field
		CreatedAt         respjson.Field
		EndAt             respjson.Field
		MaxPricePerHour   respjson.Field
		ProcurementID     respjson.Field
		ProcurementStatus respjson.Field
		UpdatedAt         respjson.Field
		Zone              respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
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

type NodeStatus string

const (
	NodeStatusPending    NodeStatus = "Pending"
	NodeStatusRunning    NodeStatus = "Running"
	NodeStatusTerminated NodeStatus = "Terminated"
	NodeStatusFailed     NodeStatus = "Failed"
	NodeStatusUnknown    NodeStatus = "Unknown"
)

type NodeProcurementStatus string

const (
	NodeProcurementStatusUninitialized    NodeProcurementStatus = "Uninitialized"
	NodeProcurementStatusActive           NodeProcurementStatus = "Active"
	NodeProcurementStatusEnded            NodeProcurementStatus = "Ended"
	NodeProcurementStatusAwaitingCapacity NodeProcurementStatus = "AwaitingCapacity"
)

// Possible zones to choose from when creating a node.
//
// TODO (ENG-1976): Support dynamic zones through an endpoint.
type NodeZone string

const (
	NodeZoneHayesValley NodeZone = "HayesValley"
)

type NodeType string

const (
	NodeTypeOnDemand NodeType = "OnDemand"
	NodeTypeReserved NodeType = "Reserved"
)

type UpdateNode struct {
	Node Node `json:"node,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Node        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UpdateNode) RawJSON() string { return r.JSON.raw }
func (r *UpdateNode) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type NodeNewParams struct {
	DesiredCount    int64 `json:"desired_count,required"`
	MaxPricePerHour int64 `json:"max_price_per_hour,required"`
	// End time as Unix timestamp in seconds
	EndAt param.Opt[int64] `json:"end_at,omitzero"`
	// Start time as Unix timestamp in seconds
	StartAt param.Opt[int64]  `json:"start_at,omitzero"`
	Zone    param.Opt[string] `json:"zone,omitzero"`
	// Custom node names. Names cannot follow the vm*id pattern vm*{16_hex_chars} as
	// this is reserved for system-generated IDs.
	Names []string `json:"names,omitzero"`
	// Any of "OnDemand", "Reserved".
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
	// Duration in seconds to extend the node by
	DurationSeconds int64 `json:"duration_seconds,required"`
	// Max price per hour for the extension in cents
	MaxPricePerHour int64 `json:"max_price_per_hour,required"`
	paramObj
}

func (r NodeExtendParams) MarshalJSON() (data []byte, err error) {
	type shadow NodeExtendParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *NodeExtendParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type NodeReleaseParams struct {
	Body any
	paramObj
}

func (r NodeReleaseParams) MarshalJSON() (data []byte, err error) {
	return json.Marshal(r.Body)
}
func (r *NodeReleaseParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.Body)
}
