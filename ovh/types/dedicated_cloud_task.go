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

// DedicatedCloudTask Operation on a Dedicated Cloud component
type DedicatedCloudTask struct {

	// CreatedBy Creator of the task
	CreatedBy string `json:"createdBy,omitempty"`

	// CreatedFrom Origin of the task
	CreatedFrom string `json:"createdFrom,omitempty"`

	// DatacenterID datacenterId of the associated dedicatedCloud.Datacenter object
	DatacenterID int64 `json:"datacenterId,omitempty"`

	// Description Current progress description
	Description string `json:"description,omitempty"`

	// EndDate Task end date
	EndDate *time.Time `json:"endDate,omitempty"`

	// ExecutionDate Task execution date
	ExecutionDate *time.Time `json:"executionDate,omitempty"`

	// FilerID filerId of the associated dedicatedCloud.Filer object
	FilerID int64 `json:"filerId,omitempty"`

	// HostID hostId of the associated dedicatedCloud.Host object
	HostID int64 `json:"hostId,omitempty"`

	// LastModificationDate Task last modification date
	LastModificationDate *time.Time `json:"lastModificationDate,omitempty"`

	// MaintenanceDateFrom Maintenance task min allowed execution date
	MaintenanceDateFrom *time.Time `json:"maintenanceDateFrom,omitempty"`

	// MaintenanceDateTo Maintenance task max allowed execution date
	MaintenanceDateTo *time.Time `json:"maintenanceDateTo,omitempty"`

	// Name Task name
	Name string `json:"name,omitempty"`

	// Network network of the associated dedicatedCloud.Ip object
	Network string `json:"network,omitempty"`

	// NetworkAccessID networkAccessId of the associated dedicatedCloud.AllowedNetwork object
	NetworkAccessID int64 `json:"networkAccessId,omitempty"`

	// OrderID orderId of the associated billing.Order object
	OrderID int64 `json:"orderId,omitempty"`

	// ParentTaskID taskId of the parent dedicatedCloud.Task object
	ParentTaskID int64 `json:"parentTaskId,omitempty"`

	// Progress Current progress
	Progress int64 `json:"progress,omitempty"`

	// State Current Task state
	State string `json:"state,omitempty"`

	// TaskID Task id
	TaskID int64 `json:"taskId,omitempty"`

	// TType Task type
	TType string `json:"type,omitempty"`

	// UserID userId of the associated dedicatedCloud.User object
	UserID int64 `json:"userId,omitempty"`

	// VlanID vlanId of the parent dedicatedCloud.Vlan object
	VlanID int64 `json:"vlanId,omitempty"`
}
