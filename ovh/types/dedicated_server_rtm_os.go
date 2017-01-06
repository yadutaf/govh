/*
 * OVH API - EU
 *
 * Build your own OVH world.
 *
 * OpenAPI spec version: 1.0.0
 * Contact: api-subscribe@ml.ovh.net
 * Generated by: https://github.com/swagger-api/swagger-codegen.git
 */

package types

// DedicatedServerRtmOs A structure describing informations about Rtm os
type DedicatedServerRtmOs struct {

	// KernelRelease OS kernel release
	KernelRelease string `json:"kernelRelease,omitempty"`

	// KernelVersion OS kernel version
	KernelVersion string `json:"kernelVersion,omitempty"`

	// Release OS release
	Release string `json:"release,omitempty"`
}