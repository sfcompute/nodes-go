// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package sfcnodes

import (
	"context"
	"net/http"
	"net/url"
	"slices"

	"github.com/sfcompute/nodes-go/internal/apijson"
	"github.com/sfcompute/nodes-go/internal/apiquery"
	"github.com/sfcompute/nodes-go/internal/requestconfig"
	"github.com/sfcompute/nodes-go/option"
	"github.com/sfcompute/nodes-go/packages/param"
	"github.com/sfcompute/nodes-go/packages/respjson"
)

// Manage your Virtual Machines.
//
// VMService contains methods and other services that help with interacting with
// the sfc-nodes API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewVMService] method instead.
type VMService struct {
	Options []option.RequestOption
	// Manage your Virtual Machines.
	Script VMScriptService
	// Manage your Virtual Machines.
	Images VMImageService
}

// NewVMService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewVMService(opts ...option.RequestOption) (r VMService) {
	r = VMService{}
	r.Options = opts
	r.Script = NewVMScriptService(opts...)
	r.Images = NewVMImageService(opts...)
	return
}

func (r *VMService) Logs(ctx context.Context, query VMLogsParams, opts ...option.RequestOption) (res *VMLogsResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v0/vms/logs2"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

func (r *VMService) SSH(ctx context.Context, query VMSSHParams, opts ...option.RequestOption) (res *VmsshResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v0/vms/ssh"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type VMLogsResponse struct {
	Data []VMLogsResponseData `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r VMLogsResponse) RawJSON() string { return r.JSON.raw }
func (r *VMLogsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type VMLogsResponseData struct {
	Data                      []int64 `json:"data" api:"required"`
	InstanceID                string  `json:"instance_id" api:"required"`
	MonotonicTimestampNanoSec int64   `json:"monotonic_timestamp_nano_sec" api:"required"`
	MonotonicTimestampSec     int64   `json:"monotonic_timestamp_sec" api:"required"`
	// In RFC 3339 format
	RealtimeTimestamp string `json:"realtime_timestamp" api:"required"`
	Seqnum            int64  `json:"seqnum" api:"required"`
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
func (r VMLogsResponseData) RawJSON() string { return r.JSON.raw }
func (r *VMLogsResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type VmsshResponse struct {
	SSHHostname string `json:"ssh_hostname" api:"required"`
	SSHPort     int64  `json:"ssh_port" api:"required"`
	// Unix timestamp.
	LastAttemptedKeyUpdate int64 `json:"last_attempted_key_update" api:"nullable"`
	// Unix timestamp.
	LastSuccessfulKeyUpdate int64                     `json:"last_successful_key_update" api:"nullable"`
	SSHHostKeys             []VmsshResponseSSHHostKey `json:"ssh_host_keys" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		SSHHostname             respjson.Field
		SSHPort                 respjson.Field
		LastAttemptedKeyUpdate  respjson.Field
		LastSuccessfulKeyUpdate respjson.Field
		SSHHostKeys             respjson.Field
		ExtraFields             map[string]respjson.Field
		raw                     string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r VmsshResponse) RawJSON() string { return r.JSON.raw }
func (r *VmsshResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type VmsshResponseSSHHostKey struct {
	Base64EncodedKey string `json:"base64_encoded_key" api:"required" format:"byte"`
	KeyType          string `json:"key_type" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Base64EncodedKey respjson.Field
		KeyType          respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r VmsshResponseSSHHostKey) RawJSON() string { return r.JSON.raw }
func (r *VmsshResponseSSHHostKey) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type VMLogsParams struct {
	InstanceID string `query:"instance_id" api:"required" json:"-"`
	// Any of "seqnum_asc", "seqnum_desc".
	OrderBy                 VMLogsParamsOrderBy `query:"order_by,omitzero" api:"required" json:"-"`
	BeforeRealtimeTimestamp param.Opt[string]   `query:"before_realtime_timestamp,omitzero" json:"-"`
	BeforeSeqnum            param.Opt[int64]    `query:"before_seqnum,omitzero" json:"-"`
	Limit                   param.Opt[int64]    `query:"limit,omitzero" json:"-"`
	SinceRealtimeTimestamp  param.Opt[string]   `query:"since_realtime_timestamp,omitzero" json:"-"`
	SinceSeqnum             param.Opt[int64]    `query:"since_seqnum,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [VMLogsParams]'s query parameters as `url.Values`.
func (r VMLogsParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type VMLogsParamsOrderBy string

const (
	VMLogsParamsOrderBySeqnumAsc  VMLogsParamsOrderBy = "seqnum_asc"
	VMLogsParamsOrderBySeqnumDesc VMLogsParamsOrderBy = "seqnum_desc"
)

type VMSSHParams struct {
	VMID string `query:"vm_id" api:"required" json:"-"`
	paramObj
}

// URLQuery serializes [VMSSHParams]'s query parameters as `url.Values`.
func (r VMSSHParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
