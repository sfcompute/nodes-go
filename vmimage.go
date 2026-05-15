// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package sfcnodes

import (
	"github.com/sfcompute/nodes-go/option"
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
