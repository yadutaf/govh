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

// PaasDatabaseInstanceDatabaseExtension Databases extension
type PaasDatabaseInstanceDatabaseExtension struct {

	// Description Extension description
	Description string `json:"description,omitempty"`

	// ExtensionName Extension name
	ExtensionName string `json:"extensionName,omitempty"`

	// RequiredExtensions Name of required extensions to enable this one
	RequiredExtensions []string `json:"requiredExtensions,omitempty"`

	// Status Extension status
	Status string `json:"status,omitempty"`
}