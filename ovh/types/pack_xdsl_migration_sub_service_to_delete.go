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

// PackXdslMigrationSubServiceToDelete Sub service to delete
type PackXdslMigrationSubServiceToDelete struct {

	// NumberToDelete Number of services to be deleted
	NumberToDelete int64 `json:"numberToDelete,omitempty"`

	// Services List of domains of sub services
	Services []string `json:"services,omitempty"`

	// TType Type of service to be deleted
	TType string `json:"type,omitempty"`
}