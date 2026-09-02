// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package sfcnodes

import (
	"github.com/sfcompute/nodes-go/option"
)

// VMService contains methods and other services that help with interacting with
// the sfc-nodes API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewVMService] method instead.
type VMService struct {
	Options []option.RequestOption
	Script  VMScriptService
	// Custom machine images for instances.
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
