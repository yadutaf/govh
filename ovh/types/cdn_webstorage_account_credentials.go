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

// CDNWebstorageAccountCredentials A structure with credentials for using openstack account
type CDNWebstorageAccountCredentials struct {
	Endpoint string `json:"endpoint,omitempty"`

	Login string `json:"login,omitempty"`

	Password string `json:"password,omitempty"`

	Tenant string `json:"tenant,omitempty"`
}
