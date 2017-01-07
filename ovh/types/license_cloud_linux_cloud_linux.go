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

import (
	"time"
)

// LicenseCloudLinux Your CloudLinux license
type LicenseCloudLinux struct {

	// Creation This license creation date
	Creation *time.Time `json:"creation,omitempty"`

	// Domain The internal name of your license
	Domain string `json:"domain,omitempty"`

	// IP The ip on which this license is attached
	IP string `json:"ip,omitempty"`

	// LicenseID The license id on license provider side
	LicenseID string `json:"licenseId,omitempty"`

	// Status This license state
	Status string `json:"status,omitempty"`

	// Version This license version
	Version string `json:"version,omitempty"`
}
