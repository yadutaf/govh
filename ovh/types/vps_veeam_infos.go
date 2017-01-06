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

// VpsVeeamInfos A structure describing a Veeam backup's access informations
type VpsVeeamInfos struct {

	// Nfs NFS URL of the backup
	Nfs string `json:"nfs,omitempty"`

	// Smb SMB URL of the backup
	Smb string `json:"smb,omitempty"`
}