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

// ImapcopyTaskIDs Ids of ImapCopy Task
type ImapcopyTaskIDs struct {

	// ID Id of task
	ID int64 `json:"id,omitempty"`

	// SecretKey Secret key of task
	SecretKey string `json:"secretKey,omitempty"`
}