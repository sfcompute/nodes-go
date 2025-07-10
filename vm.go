// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package sfcnodes

import (
	"context"
	"net/http"
	"net/url"

	"github.com/sfcompute/nodes-go/internal/apijson"
	"github.com/sfcompute/nodes-go/internal/apiquery"
	"github.com/sfcompute/nodes-go/internal/requestconfig"
	"github.com/sfcompute/nodes-go/option"
	"github.com/sfcompute/nodes-go/packages/param"
	"github.com/sfcompute/nodes-go/packages/respjson"
)

// VmService contains methods and other services that help with interacting with
// the sfc-nodes API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewVmService] method instead.
type VmService struct {
	Options []option.RequestOption
	Script  VmScriptService
}

// NewVmService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewVmService(opts ...option.RequestOption) (r VmService) {
	r = VmService{}
	r.Options = opts
	r.Script = NewVmScriptService(opts...)
	return
}

func (r *VmService) List(ctx context.Context, opts ...option.RequestOption) (res *VmListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v0/vms/instances"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

func (r *VmService) Logs(ctx context.Context, query VmLogsParams, opts ...option.RequestOption) (res *VmLogsResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v0/vms/logs2"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

func (r *VmService) Replace(ctx context.Context, body VmReplaceParams, opts ...option.RequestOption) (res *VmReplaceResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v0/vms/replace"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

func (r *VmService) SSH(ctx context.Context, query VmSSHParams, opts ...option.RequestOption) (res *VmSSHResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v0/vms/ssh"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type VmListResponse struct {
	Data []VmListResponseData `json:"data,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r VmListResponse) RawJSON() string { return r.JSON.raw }
func (r *VmListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type VmListResponseData struct {
	ID              string `json:"id,required"`
	ClusterID       string `json:"cluster_id,required"`
	CurrentStatus   string `json:"current_status,required"`
	InstanceGroupID string `json:"instance_group_id,required"`
	LastUpdatedAt   string `json:"last_updated_at,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID              respjson.Field
		ClusterID       respjson.Field
		CurrentStatus   respjson.Field
		InstanceGroupID respjson.Field
		LastUpdatedAt   respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r VmListResponseData) RawJSON() string { return r.JSON.raw }
func (r *VmListResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type VmLogsResponse struct {
	Data []VmLogsResponseData `json:"data,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r VmLogsResponse) RawJSON() string { return r.JSON.raw }
func (r *VmLogsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type VmLogsResponseData struct {
	Data                      []int64 `json:"data,required"`
	InstanceID                string  `json:"instance_id,required"`
	MonotonicTimestampNanoSec int64   `json:"monotonic_timestamp_nano_sec,required"`
	MonotonicTimestampSec     int64   `json:"monotonic_timestamp_sec,required"`
	// In RFC 3339 format
	RealtimeTimestamp string `json:"realtime_timestamp,required"`
	Seqnum            int64  `json:"seqnum,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data                      respjson.Field
		InstanceID                respjson.Field
		MonotonicTimestampNanoSec respjson.Field
		MonotonicTimestampSec     respjson.Field
		RealtimeTimestamp         respjson.Field
		Seqnum                    respjson.Field
		ExtraFields               map[string]respjson.Field
		raw                       string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r VmLogsResponseData) RawJSON() string { return r.JSON.raw }
func (r *VmLogsResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type VmReplaceResponse struct {
	Replaced   string `json:"replaced,required"`
	ReplacedBy string `json:"replaced_by,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Replaced    respjson.Field
		ReplacedBy  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r VmReplaceResponse) RawJSON() string { return r.JSON.raw }
func (r *VmReplaceResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type VmSSHResponse struct {
	SSHHostname string                    `json:"ssh_hostname,required"`
	SSHPort     int64                     `json:"ssh_port,required"`
	SSHHostKeys []VmSSHResponseSSHHostKey `json:"ssh_host_keys,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		SSHHostname respjson.Field
		SSHPort     respjson.Field
		SSHHostKeys respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r VmSSHResponse) RawJSON() string { return r.JSON.raw }
func (r *VmSSHResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type VmSSHResponseSSHHostKey struct {
	Base64EncodedKey string `json:"base64_encoded_key,required"`
	KeyType          string `json:"key_type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Base64EncodedKey respjson.Field
		KeyType          respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r VmSSHResponseSSHHostKey) RawJSON() string { return r.JSON.raw }
func (r *VmSSHResponseSSHHostKey) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type VmLogsParams struct {
	InstanceID string `query:"instance_id,required" json:"-"`
	// Any of "seqnum_asc", "seqnum_desc".
	OrderBy                 VmLogsParamsOrderBy `query:"order_by,omitzero,required" json:"-"`
	BeforeRealtimeTimestamp param.Opt[string]   `query:"before_realtime_timestamp,omitzero" json:"-"`
	BeforeSeqnum            param.Opt[int64]    `query:"before_seqnum,omitzero" json:"-"`
	Limit                   param.Opt[int64]    `query:"limit,omitzero" json:"-"`
	SinceRealtimeTimestamp  param.Opt[string]   `query:"since_realtime_timestamp,omitzero" json:"-"`
	SinceSeqnum             param.Opt[int64]    `query:"since_seqnum,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [VmLogsParams]'s query parameters as `url.Values`.
func (r VmLogsParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type VmLogsParamsOrderBy string

const (
	VmLogsParamsOrderBySeqnumAsc  VmLogsParamsOrderBy = "seqnum_asc"
	VmLogsParamsOrderBySeqnumDesc VmLogsParamsOrderBy = "seqnum_desc"
)

type VmReplaceParams struct {
	VmID string `json:"vm_id,required"`
	paramObj
}

func (r VmReplaceParams) MarshalJSON() (data []byte, err error) {
	type shadow VmReplaceParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *VmReplaceParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type VmSSHParams struct {
	VmID string `query:"vm_id,required" json:"-"`
	paramObj
}

// URLQuery serializes [VmSSHParams]'s query parameters as `url.Values`.
func (r VmSSHParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
