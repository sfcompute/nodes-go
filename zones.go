package sfcnodes

import (
	"context"
	"net/http"
	"slices"

	"github.com/sfcompute/nodes-go/internal/requestconfig"
	"github.com/sfcompute/nodes-go/option"
)

func (r *NodeService) Zones(ctx context.Context, body NodeNewParams, opts ...option.RequestOption) (res *ListResponseNode, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v0/zones"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, body, &res, opts...)
	return
}
