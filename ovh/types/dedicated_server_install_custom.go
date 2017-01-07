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

// DedicatedServerInstallCustom A structure describing informations about installation custom
type DedicatedServerInstallCustom struct {

	// CustomHostname Personnal hostname to use in server reinstallation
	CustomHostname string `json:"customHostname,omitempty"`

	// DiskGroupID Disk group id to process install on (only available for some templates)
	DiskGroupID int64 `json:"diskGroupId,omitempty"`

	// InstallSQLServer true if you want to install windows with sql feature
	InstallSQLServer bool `json:"installSqlServer,omitempty"`

	// Language install language for ovh install choice
	Language string `json:"language,omitempty"`

	// NoRaid true if you want to install only on the first disk
	NoRaid bool `json:"noRaid,omitempty"`

	// PostInstallationScriptLink the url to your custom install script
	PostInstallationScriptLink string `json:"postInstallationScriptLink,omitempty"`

	// PostInstallationScriptReturn  the return of your script if everythings ok. Advice: your script should return a unique validation string in case of succes. A good example is \"loh1Xee7eo OK OK OK UGh8Ang1Gu
	PostInstallationScriptReturn string `json:"postInstallationScriptReturn,omitempty"`

	// SoftRaidDevices Number of devices to use for system's software RAID
	SoftRaidDevices int64 `json:"softRaidDevices,omitempty"`

	// SSHKeyName The name of ssh key to install
	SSHKeyName string `json:"sshKeyName,omitempty"`

	// UseDistribKernel true if you want to install with distrib kernel
	UseDistribKernel bool `json:"useDistribKernel,omitempty"`

	// UseSpla true if you want to install windows with your spla license
	UseSpla bool `json:"useSpla,omitempty"`
}
