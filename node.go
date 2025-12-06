// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package sfcnodes

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"

	"github.com/sfcompute/nodes-go/internal/apijson"
	"github.com/sfcompute/nodes-go/internal/apiquery"
	shimjson "github.com/sfcompute/nodes-go/internal/encoding/json"
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
func (r *NodeService) New(ctx context.Context, body NodeNewParams, opts ...option.RequestOption) (res *ListResponseNode, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/nodes"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// List all nodes for the authenticated account
func (r *NodeService) List(ctx context.Context, query NodeListParams, opts ...option.RequestOption) (res *ListResponseNode, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/nodes"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Delete a node by id. The node cannot be deleted if it has active or pending VMs.
func (r *NodeService) Delete(ctx context.Context, id string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("v1/nodes/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

// Purchase additional time to extend the end time of a reserved VM node
func (r *NodeService) Extend(ctx context.Context, id string, body NodeExtendParams, opts ...option.RequestOption) (res *Node, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("v1/nodes/%s/extend", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// Retrieve details of a specific node by its ID or name
func (r *NodeService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *Node, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("v1/nodes/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Redeploy a node by replacing its current VM with a new one. Optionally update
// the VM image and cloud init user data.
func (r *NodeService) Redeploy(ctx context.Context, id string, body NodeRedeployParams, opts ...option.RequestOption) (res *Node, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("v1/nodes/%s/redeploy", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// Release an auto reserved VM node from its procurement, reducing the
// procurement's desired quantity by 1
func (r *NodeService) Release(ctx context.Context, id string, opts ...option.RequestOption) (res *Node, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("v1/nodes/%s/release", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, nil, &res, opts...)
	return
}

type AcceleratorType string

const (
	AcceleratorTypeH100 AcceleratorType = "H100"
	AcceleratorTypeH200 AcceleratorType = "H200"
)

// The properties DesiredCount, MaxPricePerNodeHour are required.
type CreateNodesRequestParam struct {
	DesiredCount int64 `json:"desired_count,required"`
	// Max price per hour for a node in cents
	MaxPricePerNodeHour int64 `json:"max_price_per_node_hour,required"`
	// End time as Unix timestamp in seconds If provided, end time must be aligned to
	// the hour If not provided, the node will be created as an autoreserved node
	EndAt param.Opt[int64] `json:"end_at,omitzero"`
	// Allow auto reserved nodes to be created in any zone that meets the requirements
	AnyZone param.Opt[bool] `json:"any_zone,omitzero"`
	// User script to be executed during the VM's boot process Data should be base64
	// encoded
	CloudInitUserData param.Opt[string] `json:"cloud_init_user_data,omitzero" format:"byte"`
	// Custom image ID to use for the VM instances
	ImageID param.Opt[string] `json:"image_id,omitzero"`
	// Start time as Unix timestamp in seconds Optional for reserved nodes. If not
	// provided, defaults to now
	StartAt param.Opt[int64] `json:"start_at,omitzero"`
	// Zone to create the nodes in. Required for auto reserved nodes if any_zone is
	// false.
	Zone param.Opt[string] `json:"zone,omitzero"`
	// Custom node names Names cannot begin with 'vm*' or 'n*' as this is reserved for
	// system-generated IDs Names cannot be numeric strings Names cannot exceed 128
	// characters
	Names []string `json:"names,omitzero"`
	// Any of "autoreserved", "reserved".
	NodeType NodeType `json:"node_type,omitzero"`
	paramObj
}

func (r CreateNodesRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow CreateNodesRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CreateNodesRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties DurationSeconds, MaxPricePerNodeHour are required.
type ExtendNodeRequestParam struct {
	// Duration in seconds to extend the node Must be at least 1 hour (3600 seconds)
	// and a multiple of 1 hour.
	DurationSeconds int64 `json:"duration_seconds,required"`
	// Max price per hour for the extension in cents
	MaxPricePerNodeHour int64 `json:"max_price_per_node_hour,required"`
	paramObj
}

func (r ExtendNodeRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow ExtendNodeRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ExtendNodeRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ListResponseNode struct {
	Data   []ListResponseNodeData `json:"data,required"`
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
func (r ListResponseNode) RawJSON() string { return r.JSON.raw }
func (r *ListResponseNode) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ListResponseNodeData struct {
	ID string `json:"id,required"`
	// Any of "H100", "H200".
	GPUType AcceleratorType `json:"gpu_type,required"`
	Name    string          `json:"name,required"`
	// Any of "autoreserved", "reserved".
	NodeType NodeType `json:"node_type,required"`
	Object   string   `json:"object,required"`
	Owner    string   `json:"owner,required"`
	// Node Status
	//
	// Any of "pending", "awaitingcapacity", "running", "released", "terminated",
	// "deleted", "failed", "unknown".
	Status Status `json:"status,required"`
	// Creation time as Unix timestamp in seconds
	CreatedAt int64                         `json:"created_at,nullable"`
	CurrentVM ListResponseNodeDataCurrentVM `json:"current_vm,nullable"`
	// Deletion time as Unix timestamp in seconds
	DeletedAt int64 `json:"deleted_at,nullable"`
	// End time as Unix timestamp in seconds
	EndAt int64 `json:"end_at,nullable"`
	// Max price per hour you're willing to pay for a node in cents
	MaxPricePerNodeHour int64  `json:"max_price_per_node_hour,nullable"`
	ProcurementID       string `json:"procurement_id,nullable"`
	// Start time as Unix timestamp in seconds
	StartAt int64 `json:"start_at,nullable"`
	// Last updated time as Unix timestamp in seconds
	UpdatedAt int64                   `json:"updated_at,nullable"`
	VMs       ListResponseNodeDataVMs `json:"vms,nullable"`
	Zone      string                  `json:"zone,nullable"`
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
		CurrentVM           respjson.Field
		DeletedAt           respjson.Field
		EndAt               respjson.Field
		MaxPricePerNodeHour respjson.Field
		ProcurementID       respjson.Field
		StartAt             respjson.Field
		UpdatedAt           respjson.Field
		VMs                 respjson.Field
		Zone                respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ListResponseNodeData) RawJSON() string { return r.JSON.raw }
func (r *ListResponseNodeData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ListResponseNodeDataCurrentVM struct {
	ID        string `json:"id,required"`
	CreatedAt int64  `json:"created_at,required"`
	EndAt     int64  `json:"end_at,required"`
	Object    string `json:"object,required"`
	StartAt   int64  `json:"start_at,required"`
	// Any of "Pending", "Running", "Destroyed", "NodeFailure", "Unspecified".
	Status    string `json:"status,required"`
	UpdatedAt int64  `json:"updated_at,required"`
	Zone      string `json:"zone,required"`
	ImageID   string `json:"image_id,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		CreatedAt   respjson.Field
		EndAt       respjson.Field
		Object      respjson.Field
		StartAt     respjson.Field
		Status      respjson.Field
		UpdatedAt   respjson.Field
		Zone        respjson.Field
		ImageID     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ListResponseNodeDataCurrentVM) RawJSON() string { return r.JSON.raw }
func (r *ListResponseNodeDataCurrentVM) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ListResponseNodeDataVMs struct {
	Data   []ListResponseNodeDataVMsData `json:"data,required"`
	Object string                        `json:"object,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Object      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ListResponseNodeDataVMs) RawJSON() string { return r.JSON.raw }
func (r *ListResponseNodeDataVMs) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ListResponseNodeDataVMsData struct {
	ID        string `json:"id,required"`
	CreatedAt int64  `json:"created_at,required"`
	EndAt     int64  `json:"end_at,required"`
	Object    string `json:"object,required"`
	StartAt   int64  `json:"start_at,required"`
	// Any of "Pending", "Running", "Destroyed", "NodeFailure", "Unspecified".
	Status    string `json:"status,required"`
	UpdatedAt int64  `json:"updated_at,required"`
	Zone      string `json:"zone,required"`
	ImageID   string `json:"image_id,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		CreatedAt   respjson.Field
		EndAt       respjson.Field
		Object      respjson.Field
		StartAt     respjson.Field
		Status      respjson.Field
		UpdatedAt   respjson.Field
		Zone        respjson.Field
		ImageID     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ListResponseNodeDataVMsData) RawJSON() string { return r.JSON.raw }
func (r *ListResponseNodeDataVMsData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type Node struct {
	ID string `json:"id,required"`
	// Any of "H100", "H200".
	GPUType AcceleratorType `json:"gpu_type,required"`
	Name    string          `json:"name,required"`
	// Any of "autoreserved", "reserved".
	NodeType NodeType `json:"node_type,required"`
	Object   string   `json:"object,required"`
	Owner    string   `json:"owner,required"`
	// Node Status
	//
	// Any of "pending", "awaitingcapacity", "running", "released", "terminated",
	// "deleted", "failed", "unknown".
	Status Status `json:"status,required"`
	// Creation time as Unix timestamp in seconds
	CreatedAt int64         `json:"created_at,nullable"`
	CurrentVM NodeCurrentVM `json:"current_vm,nullable"`
	// Deletion time as Unix timestamp in seconds
	DeletedAt int64 `json:"deleted_at,nullable"`
	// End time as Unix timestamp in seconds
	EndAt int64 `json:"end_at,nullable"`
	// Max price per hour you're willing to pay for a node in cents
	MaxPricePerNodeHour int64  `json:"max_price_per_node_hour,nullable"`
	ProcurementID       string `json:"procurement_id,nullable"`
	// Start time as Unix timestamp in seconds
	StartAt int64 `json:"start_at,nullable"`
	// Last updated time as Unix timestamp in seconds
	UpdatedAt int64   `json:"updated_at,nullable"`
	VMs       NodeVMs `json:"vms,nullable"`
	Zone      string  `json:"zone,nullable"`
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
		CurrentVM           respjson.Field
		DeletedAt           respjson.Field
		EndAt               respjson.Field
		MaxPricePerNodeHour respjson.Field
		ProcurementID       respjson.Field
		StartAt             respjson.Field
		UpdatedAt           respjson.Field
		VMs                 respjson.Field
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

type NodeCurrentVM struct {
	ID        string `json:"id,required"`
	CreatedAt int64  `json:"created_at,required"`
	EndAt     int64  `json:"end_at,required"`
	Object    string `json:"object,required"`
	StartAt   int64  `json:"start_at,required"`
	// Any of "Pending", "Running", "Destroyed", "NodeFailure", "Unspecified".
	Status    string `json:"status,required"`
	UpdatedAt int64  `json:"updated_at,required"`
	Zone      string `json:"zone,required"`
	ImageID   string `json:"image_id,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		CreatedAt   respjson.Field
		EndAt       respjson.Field
		Object      respjson.Field
		StartAt     respjson.Field
		Status      respjson.Field
		UpdatedAt   respjson.Field
		Zone        respjson.Field
		ImageID     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NodeCurrentVM) RawJSON() string { return r.JSON.raw }
func (r *NodeCurrentVM) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type NodeVMs struct {
	Data   []NodeVMsData `json:"data,required"`
	Object string        `json:"object,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Object      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NodeVMs) RawJSON() string { return r.JSON.raw }
func (r *NodeVMs) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type NodeVMsData struct {
	ID        string `json:"id,required"`
	CreatedAt int64  `json:"created_at,required"`
	EndAt     int64  `json:"end_at,required"`
	Object    string `json:"object,required"`
	StartAt   int64  `json:"start_at,required"`
	// Any of "Pending", "Running", "Destroyed", "NodeFailure", "Unspecified".
	Status    string `json:"status,required"`
	UpdatedAt int64  `json:"updated_at,required"`
	Zone      string `json:"zone,required"`
	ImageID   string `json:"image_id,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		CreatedAt   respjson.Field
		EndAt       respjson.Field
		Object      respjson.Field
		StartAt     respjson.Field
		Status      respjson.Field
		UpdatedAt   respjson.Field
		Zone        respjson.Field
		ImageID     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NodeVMsData) RawJSON() string { return r.JSON.raw }
func (r *NodeVMsData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type NodeType string

const (
	NodeTypeAutoreserved NodeType = "autoreserved"
	NodeTypeReserved     NodeType = "reserved"
)

// Node Status
type Status string

const (
	StatusPending          Status = "pending"
	StatusAwaitingcapacity Status = "awaitingcapacity"
	StatusRunning          Status = "running"
	StatusReleased         Status = "released"
	StatusTerminated       Status = "terminated"
	StatusDeleted          Status = "deleted"
	StatusFailed           Status = "failed"
	StatusUnknown          Status = "unknown"
)

type NodeNewParams struct {
	CreateNodesRequest CreateNodesRequestParam
	paramObj
}

func (r NodeNewParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.CreateNodesRequest)
}
func (r *NodeNewParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.CreateNodesRequest)
}

type NodeListParams struct {
	// Filter nodes by node_id Use ?id=n_b1dc52505c6db142&id=n_b1dc52505c6db133 to
	// specify multiple IDs. Cannot combine with name or node_type
	ID []string `query:"id,omitzero" json:"-"`
	// Filter nodes by their names Use ?name=val1&name=val2 to specify multiple names.
	// Cannot combine with id or node_type
	Name []string `query:"name,omitzero" json:"-"`
	// Filter nodes by their type Cannot combine with id or name
	//
	// Any of "autoreserved", "reserved".
	Type NodeType `query:"type,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [NodeListParams]'s query parameters as `url.Values`.
func (r NodeListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type NodeExtendParams struct {
	ExtendNodeRequest ExtendNodeRequestParam
	paramObj
}

func (r NodeExtendParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.ExtendNodeRequest)
}
func (r *NodeExtendParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.ExtendNodeRequest)
}

type NodeRedeployParams struct {
	// Update the cloud init user data for VMs running on this node Data should be
	// base64 encoded
	CloudInitUserData param.Opt[string] `json:"cloud_init_user_data,omitzero" format:"byte"`
	// Redeploy node with this VM image ID
	ImageID param.Opt[string] `json:"image_id,omitzero"`
	// If false, then the new VM will inherit any configuration (like image_id,
	// cloud_init_user_data) that is left empty in this request from the current VM.
	//
	// If true, then any configuration left empty will be set as empty in the new VM.
	// E.g if cloud_init_user_data is left unset and override_empty is true, then the
	// new VM will not have any cloud init user data. override_empty defaults to false.
	OverrideEmpty param.Opt[bool] `json:"override_empty,omitzero"`
	paramObj
}

func (r NodeRedeployParams) MarshalJSON() (data []byte, err error) {
	type shadow NodeRedeployParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *NodeRedeployParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
