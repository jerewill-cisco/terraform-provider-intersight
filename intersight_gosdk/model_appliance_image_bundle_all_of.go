/*
 * Cisco Intersight
 *
 * Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document. This document was created on 2021-05-12T14:10:48Z.
 *
 * API version: 1.0.9-4289
 * Contact: intersight@cisco.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package intersight

import (
	"encoding/json"
	"time"
)

// ApplianceImageBundleAllOf Definition of the list of properties defined in 'appliance.ImageBundle', excluding properties defined in parent classes.
type ApplianceImageBundleAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType      string               `json:"ObjectType"`
	AnsiblePackages []OnpremImagePackage `json:"AnsiblePackages,omitempty"`
	// Indicates that the software upgrade was automatically initiated by the Intersight Appliance.
	AutoUpgrade   *bool                `json:"AutoUpgrade,omitempty"`
	DcPackages    []OnpremImagePackage `json:"DcPackages,omitempty"`
	DebugPackages []OnpremImagePackage `json:"DebugPackages,omitempty"`
	// Short description of the software upgrade bundle.
	Description      *string              `json:"Description,omitempty"`
	EndpointPackages []OnpremImagePackage `json:"EndpointPackages,omitempty"`
	// Fingerprint of the software manifest from which this bundle is created. Fingerprint is calculated using the SHA256 algorithm.
	Fingerprint *string `json:"Fingerprint,omitempty"`
	// Indicates that the ImageBundle has errors. The upgrade service sets this field when it encounters errors during the manifest processing.
	HasError      *bool                `json:"HasError,omitempty"`
	InfraPackages []OnpremImagePackage `json:"InfraPackages,omitempty"`
	InitPackages  []OnpremImagePackage `json:"InitPackages,omitempty"`
	// Name of the software upgrade bundle.
	Name *string `json:"Name,omitempty"`
	// Detailed description of the software upgrade bundle.
	Notes *string `json:"Notes,omitempty"`
	// Software upgrade manifest's upgrade priority. The upgrade service supports two priorities, Normal and Critical. Normal priority is used for regular software upgrades, and the upgrade service uses the Upgrade Policy to compute upgrade start time. Critical priority is used for the critical software security patches, and the upgrade service ignores the Upgrade Policy when it computes the upgrade start time. * `Normal` - Normal upgrade priority is used for all the software upgrades except for the critical security updates. The upgrade service of Intersight Appliance uses the Software Upgrade Policy settings to start the upgrade process. * `Critical` - Critical upgrade priority is used for critical updates such as security patches. The upgrade service of the Intersight Appliance starts the upgrade as specified by the upgrade properties in the software manifest file. The upgrade service will not use the settings specified in the Software Upgrade Policy.
	Priority *string `json:"Priority,omitempty"`
	// Software upgrade manifest's release date and time.
	ReleaseTime     *time.Time           `json:"ReleaseTime,omitempty"`
	ServicePackages []OnpremImagePackage `json:"ServicePackages,omitempty"`
	// Status message set during the manifest processing.
	StatusMessage  *string              `json:"StatusMessage,omitempty"`
	SystemPackages []OnpremImagePackage `json:"SystemPackages,omitempty"`
	UiPackages     []OnpremImagePackage `json:"UiPackages,omitempty"`
	// End date of the software upgrade process.
	UpgradeEndTime *time.Time `json:"UpgradeEndTime,omitempty"`
	// Grace period in seconds before the automatic upgrade is initiated. The upgrade service uses the grace period to compute the upgrade start time when it receives an upgrade notfication from the Intersight. If there is an Upgrade Policy configured for the Intersight Appliance, then the upgrade service uses the policy to compute the upgrade start time. However, the upgrade start time cannot not exceed the limit enforced by the grace period.
	UpgradeGracePeriod *int64 `json:"UpgradeGracePeriod,omitempty"`
	// Duration (in minutes) for which services will be disrupted.
	UpgradeImpactDuration *int64 `json:"UpgradeImpactDuration,omitempty"`
	// UpgradeImpactEnum is used to indicate the kind of impact the upgrade has on currently running services on the appliance. * `None` - The upgrade has no effect on the system. * `Disruptive` - The services will not be functional during the upgrade. * `Disruptive-reboot` - The upgrade needs a reboot.
	UpgradeImpactEnum *string `json:"UpgradeImpactEnum,omitempty"`
	// Start date of the software upgrade process.
	UpgradeStartTime *time.Time `json:"UpgradeStartTime,omitempty"`
	// Software upgrade manifest's version.
	Version              *string `json:"Version,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ApplianceImageBundleAllOf ApplianceImageBundleAllOf

// NewApplianceImageBundleAllOf instantiates a new ApplianceImageBundleAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApplianceImageBundleAllOf(classId string, objectType string) *ApplianceImageBundleAllOf {
	this := ApplianceImageBundleAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	var priority string = "Normal"
	this.Priority = &priority
	var upgradeImpactDuration int64 = 0
	this.UpgradeImpactDuration = &upgradeImpactDuration
	var upgradeImpactEnum string = "None"
	this.UpgradeImpactEnum = &upgradeImpactEnum
	return &this
}

// NewApplianceImageBundleAllOfWithDefaults instantiates a new ApplianceImageBundleAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApplianceImageBundleAllOfWithDefaults() *ApplianceImageBundleAllOf {
	this := ApplianceImageBundleAllOf{}
	var classId string = "appliance.ImageBundle"
	this.ClassId = classId
	var objectType string = "appliance.ImageBundle"
	this.ObjectType = objectType
	var priority string = "Normal"
	this.Priority = &priority
	var upgradeImpactDuration int64 = 0
	this.UpgradeImpactDuration = &upgradeImpactDuration
	var upgradeImpactEnum string = "None"
	this.UpgradeImpactEnum = &upgradeImpactEnum
	return &this
}

// GetClassId returns the ClassId field value
func (o *ApplianceImageBundleAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *ApplianceImageBundleAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *ApplianceImageBundleAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *ApplianceImageBundleAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *ApplianceImageBundleAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *ApplianceImageBundleAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetAnsiblePackages returns the AnsiblePackages field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApplianceImageBundleAllOf) GetAnsiblePackages() []OnpremImagePackage {
	if o == nil {
		var ret []OnpremImagePackage
		return ret
	}
	return o.AnsiblePackages
}

// GetAnsiblePackagesOk returns a tuple with the AnsiblePackages field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApplianceImageBundleAllOf) GetAnsiblePackagesOk() (*[]OnpremImagePackage, bool) {
	if o == nil || o.AnsiblePackages == nil {
		return nil, false
	}
	return &o.AnsiblePackages, true
}

// HasAnsiblePackages returns a boolean if a field has been set.
func (o *ApplianceImageBundleAllOf) HasAnsiblePackages() bool {
	if o != nil && o.AnsiblePackages != nil {
		return true
	}

	return false
}

// SetAnsiblePackages gets a reference to the given []OnpremImagePackage and assigns it to the AnsiblePackages field.
func (o *ApplianceImageBundleAllOf) SetAnsiblePackages(v []OnpremImagePackage) {
	o.AnsiblePackages = v
}

// GetAutoUpgrade returns the AutoUpgrade field value if set, zero value otherwise.
func (o *ApplianceImageBundleAllOf) GetAutoUpgrade() bool {
	if o == nil || o.AutoUpgrade == nil {
		var ret bool
		return ret
	}
	return *o.AutoUpgrade
}

// GetAutoUpgradeOk returns a tuple with the AutoUpgrade field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceImageBundleAllOf) GetAutoUpgradeOk() (*bool, bool) {
	if o == nil || o.AutoUpgrade == nil {
		return nil, false
	}
	return o.AutoUpgrade, true
}

// HasAutoUpgrade returns a boolean if a field has been set.
func (o *ApplianceImageBundleAllOf) HasAutoUpgrade() bool {
	if o != nil && o.AutoUpgrade != nil {
		return true
	}

	return false
}

// SetAutoUpgrade gets a reference to the given bool and assigns it to the AutoUpgrade field.
func (o *ApplianceImageBundleAllOf) SetAutoUpgrade(v bool) {
	o.AutoUpgrade = &v
}

// GetDcPackages returns the DcPackages field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApplianceImageBundleAllOf) GetDcPackages() []OnpremImagePackage {
	if o == nil {
		var ret []OnpremImagePackage
		return ret
	}
	return o.DcPackages
}

// GetDcPackagesOk returns a tuple with the DcPackages field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApplianceImageBundleAllOf) GetDcPackagesOk() (*[]OnpremImagePackage, bool) {
	if o == nil || o.DcPackages == nil {
		return nil, false
	}
	return &o.DcPackages, true
}

// HasDcPackages returns a boolean if a field has been set.
func (o *ApplianceImageBundleAllOf) HasDcPackages() bool {
	if o != nil && o.DcPackages != nil {
		return true
	}

	return false
}

// SetDcPackages gets a reference to the given []OnpremImagePackage and assigns it to the DcPackages field.
func (o *ApplianceImageBundleAllOf) SetDcPackages(v []OnpremImagePackage) {
	o.DcPackages = v
}

// GetDebugPackages returns the DebugPackages field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApplianceImageBundleAllOf) GetDebugPackages() []OnpremImagePackage {
	if o == nil {
		var ret []OnpremImagePackage
		return ret
	}
	return o.DebugPackages
}

// GetDebugPackagesOk returns a tuple with the DebugPackages field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApplianceImageBundleAllOf) GetDebugPackagesOk() (*[]OnpremImagePackage, bool) {
	if o == nil || o.DebugPackages == nil {
		return nil, false
	}
	return &o.DebugPackages, true
}

// HasDebugPackages returns a boolean if a field has been set.
func (o *ApplianceImageBundleAllOf) HasDebugPackages() bool {
	if o != nil && o.DebugPackages != nil {
		return true
	}

	return false
}

// SetDebugPackages gets a reference to the given []OnpremImagePackage and assigns it to the DebugPackages field.
func (o *ApplianceImageBundleAllOf) SetDebugPackages(v []OnpremImagePackage) {
	o.DebugPackages = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ApplianceImageBundleAllOf) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceImageBundleAllOf) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ApplianceImageBundleAllOf) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ApplianceImageBundleAllOf) SetDescription(v string) {
	o.Description = &v
}

// GetEndpointPackages returns the EndpointPackages field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApplianceImageBundleAllOf) GetEndpointPackages() []OnpremImagePackage {
	if o == nil {
		var ret []OnpremImagePackage
		return ret
	}
	return o.EndpointPackages
}

// GetEndpointPackagesOk returns a tuple with the EndpointPackages field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApplianceImageBundleAllOf) GetEndpointPackagesOk() (*[]OnpremImagePackage, bool) {
	if o == nil || o.EndpointPackages == nil {
		return nil, false
	}
	return &o.EndpointPackages, true
}

// HasEndpointPackages returns a boolean if a field has been set.
func (o *ApplianceImageBundleAllOf) HasEndpointPackages() bool {
	if o != nil && o.EndpointPackages != nil {
		return true
	}

	return false
}

// SetEndpointPackages gets a reference to the given []OnpremImagePackage and assigns it to the EndpointPackages field.
func (o *ApplianceImageBundleAllOf) SetEndpointPackages(v []OnpremImagePackage) {
	o.EndpointPackages = v
}

// GetFingerprint returns the Fingerprint field value if set, zero value otherwise.
func (o *ApplianceImageBundleAllOf) GetFingerprint() string {
	if o == nil || o.Fingerprint == nil {
		var ret string
		return ret
	}
	return *o.Fingerprint
}

// GetFingerprintOk returns a tuple with the Fingerprint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceImageBundleAllOf) GetFingerprintOk() (*string, bool) {
	if o == nil || o.Fingerprint == nil {
		return nil, false
	}
	return o.Fingerprint, true
}

// HasFingerprint returns a boolean if a field has been set.
func (o *ApplianceImageBundleAllOf) HasFingerprint() bool {
	if o != nil && o.Fingerprint != nil {
		return true
	}

	return false
}

// SetFingerprint gets a reference to the given string and assigns it to the Fingerprint field.
func (o *ApplianceImageBundleAllOf) SetFingerprint(v string) {
	o.Fingerprint = &v
}

// GetHasError returns the HasError field value if set, zero value otherwise.
func (o *ApplianceImageBundleAllOf) GetHasError() bool {
	if o == nil || o.HasError == nil {
		var ret bool
		return ret
	}
	return *o.HasError
}

// GetHasErrorOk returns a tuple with the HasError field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceImageBundleAllOf) GetHasErrorOk() (*bool, bool) {
	if o == nil || o.HasError == nil {
		return nil, false
	}
	return o.HasError, true
}

// HasHasError returns a boolean if a field has been set.
func (o *ApplianceImageBundleAllOf) HasHasError() bool {
	if o != nil && o.HasError != nil {
		return true
	}

	return false
}

// SetHasError gets a reference to the given bool and assigns it to the HasError field.
func (o *ApplianceImageBundleAllOf) SetHasError(v bool) {
	o.HasError = &v
}

// GetInfraPackages returns the InfraPackages field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApplianceImageBundleAllOf) GetInfraPackages() []OnpremImagePackage {
	if o == nil {
		var ret []OnpremImagePackage
		return ret
	}
	return o.InfraPackages
}

// GetInfraPackagesOk returns a tuple with the InfraPackages field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApplianceImageBundleAllOf) GetInfraPackagesOk() (*[]OnpremImagePackage, bool) {
	if o == nil || o.InfraPackages == nil {
		return nil, false
	}
	return &o.InfraPackages, true
}

// HasInfraPackages returns a boolean if a field has been set.
func (o *ApplianceImageBundleAllOf) HasInfraPackages() bool {
	if o != nil && o.InfraPackages != nil {
		return true
	}

	return false
}

// SetInfraPackages gets a reference to the given []OnpremImagePackage and assigns it to the InfraPackages field.
func (o *ApplianceImageBundleAllOf) SetInfraPackages(v []OnpremImagePackage) {
	o.InfraPackages = v
}

// GetInitPackages returns the InitPackages field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApplianceImageBundleAllOf) GetInitPackages() []OnpremImagePackage {
	if o == nil {
		var ret []OnpremImagePackage
		return ret
	}
	return o.InitPackages
}

// GetInitPackagesOk returns a tuple with the InitPackages field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApplianceImageBundleAllOf) GetInitPackagesOk() (*[]OnpremImagePackage, bool) {
	if o == nil || o.InitPackages == nil {
		return nil, false
	}
	return &o.InitPackages, true
}

// HasInitPackages returns a boolean if a field has been set.
func (o *ApplianceImageBundleAllOf) HasInitPackages() bool {
	if o != nil && o.InitPackages != nil {
		return true
	}

	return false
}

// SetInitPackages gets a reference to the given []OnpremImagePackage and assigns it to the InitPackages field.
func (o *ApplianceImageBundleAllOf) SetInitPackages(v []OnpremImagePackage) {
	o.InitPackages = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ApplianceImageBundleAllOf) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceImageBundleAllOf) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ApplianceImageBundleAllOf) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ApplianceImageBundleAllOf) SetName(v string) {
	o.Name = &v
}

// GetNotes returns the Notes field value if set, zero value otherwise.
func (o *ApplianceImageBundleAllOf) GetNotes() string {
	if o == nil || o.Notes == nil {
		var ret string
		return ret
	}
	return *o.Notes
}

// GetNotesOk returns a tuple with the Notes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceImageBundleAllOf) GetNotesOk() (*string, bool) {
	if o == nil || o.Notes == nil {
		return nil, false
	}
	return o.Notes, true
}

// HasNotes returns a boolean if a field has been set.
func (o *ApplianceImageBundleAllOf) HasNotes() bool {
	if o != nil && o.Notes != nil {
		return true
	}

	return false
}

// SetNotes gets a reference to the given string and assigns it to the Notes field.
func (o *ApplianceImageBundleAllOf) SetNotes(v string) {
	o.Notes = &v
}

// GetPriority returns the Priority field value if set, zero value otherwise.
func (o *ApplianceImageBundleAllOf) GetPriority() string {
	if o == nil || o.Priority == nil {
		var ret string
		return ret
	}
	return *o.Priority
}

// GetPriorityOk returns a tuple with the Priority field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceImageBundleAllOf) GetPriorityOk() (*string, bool) {
	if o == nil || o.Priority == nil {
		return nil, false
	}
	return o.Priority, true
}

// HasPriority returns a boolean if a field has been set.
func (o *ApplianceImageBundleAllOf) HasPriority() bool {
	if o != nil && o.Priority != nil {
		return true
	}

	return false
}

// SetPriority gets a reference to the given string and assigns it to the Priority field.
func (o *ApplianceImageBundleAllOf) SetPriority(v string) {
	o.Priority = &v
}

// GetReleaseTime returns the ReleaseTime field value if set, zero value otherwise.
func (o *ApplianceImageBundleAllOf) GetReleaseTime() time.Time {
	if o == nil || o.ReleaseTime == nil {
		var ret time.Time
		return ret
	}
	return *o.ReleaseTime
}

// GetReleaseTimeOk returns a tuple with the ReleaseTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceImageBundleAllOf) GetReleaseTimeOk() (*time.Time, bool) {
	if o == nil || o.ReleaseTime == nil {
		return nil, false
	}
	return o.ReleaseTime, true
}

// HasReleaseTime returns a boolean if a field has been set.
func (o *ApplianceImageBundleAllOf) HasReleaseTime() bool {
	if o != nil && o.ReleaseTime != nil {
		return true
	}

	return false
}

// SetReleaseTime gets a reference to the given time.Time and assigns it to the ReleaseTime field.
func (o *ApplianceImageBundleAllOf) SetReleaseTime(v time.Time) {
	o.ReleaseTime = &v
}

// GetServicePackages returns the ServicePackages field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApplianceImageBundleAllOf) GetServicePackages() []OnpremImagePackage {
	if o == nil {
		var ret []OnpremImagePackage
		return ret
	}
	return o.ServicePackages
}

// GetServicePackagesOk returns a tuple with the ServicePackages field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApplianceImageBundleAllOf) GetServicePackagesOk() (*[]OnpremImagePackage, bool) {
	if o == nil || o.ServicePackages == nil {
		return nil, false
	}
	return &o.ServicePackages, true
}

// HasServicePackages returns a boolean if a field has been set.
func (o *ApplianceImageBundleAllOf) HasServicePackages() bool {
	if o != nil && o.ServicePackages != nil {
		return true
	}

	return false
}

// SetServicePackages gets a reference to the given []OnpremImagePackage and assigns it to the ServicePackages field.
func (o *ApplianceImageBundleAllOf) SetServicePackages(v []OnpremImagePackage) {
	o.ServicePackages = v
}

// GetStatusMessage returns the StatusMessage field value if set, zero value otherwise.
func (o *ApplianceImageBundleAllOf) GetStatusMessage() string {
	if o == nil || o.StatusMessage == nil {
		var ret string
		return ret
	}
	return *o.StatusMessage
}

// GetStatusMessageOk returns a tuple with the StatusMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceImageBundleAllOf) GetStatusMessageOk() (*string, bool) {
	if o == nil || o.StatusMessage == nil {
		return nil, false
	}
	return o.StatusMessage, true
}

// HasStatusMessage returns a boolean if a field has been set.
func (o *ApplianceImageBundleAllOf) HasStatusMessage() bool {
	if o != nil && o.StatusMessage != nil {
		return true
	}

	return false
}

// SetStatusMessage gets a reference to the given string and assigns it to the StatusMessage field.
func (o *ApplianceImageBundleAllOf) SetStatusMessage(v string) {
	o.StatusMessage = &v
}

// GetSystemPackages returns the SystemPackages field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApplianceImageBundleAllOf) GetSystemPackages() []OnpremImagePackage {
	if o == nil {
		var ret []OnpremImagePackage
		return ret
	}
	return o.SystemPackages
}

// GetSystemPackagesOk returns a tuple with the SystemPackages field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApplianceImageBundleAllOf) GetSystemPackagesOk() (*[]OnpremImagePackage, bool) {
	if o == nil || o.SystemPackages == nil {
		return nil, false
	}
	return &o.SystemPackages, true
}

// HasSystemPackages returns a boolean if a field has been set.
func (o *ApplianceImageBundleAllOf) HasSystemPackages() bool {
	if o != nil && o.SystemPackages != nil {
		return true
	}

	return false
}

// SetSystemPackages gets a reference to the given []OnpremImagePackage and assigns it to the SystemPackages field.
func (o *ApplianceImageBundleAllOf) SetSystemPackages(v []OnpremImagePackage) {
	o.SystemPackages = v
}

// GetUiPackages returns the UiPackages field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApplianceImageBundleAllOf) GetUiPackages() []OnpremImagePackage {
	if o == nil {
		var ret []OnpremImagePackage
		return ret
	}
	return o.UiPackages
}

// GetUiPackagesOk returns a tuple with the UiPackages field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApplianceImageBundleAllOf) GetUiPackagesOk() (*[]OnpremImagePackage, bool) {
	if o == nil || o.UiPackages == nil {
		return nil, false
	}
	return &o.UiPackages, true
}

// HasUiPackages returns a boolean if a field has been set.
func (o *ApplianceImageBundleAllOf) HasUiPackages() bool {
	if o != nil && o.UiPackages != nil {
		return true
	}

	return false
}

// SetUiPackages gets a reference to the given []OnpremImagePackage and assigns it to the UiPackages field.
func (o *ApplianceImageBundleAllOf) SetUiPackages(v []OnpremImagePackage) {
	o.UiPackages = v
}

// GetUpgradeEndTime returns the UpgradeEndTime field value if set, zero value otherwise.
func (o *ApplianceImageBundleAllOf) GetUpgradeEndTime() time.Time {
	if o == nil || o.UpgradeEndTime == nil {
		var ret time.Time
		return ret
	}
	return *o.UpgradeEndTime
}

// GetUpgradeEndTimeOk returns a tuple with the UpgradeEndTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceImageBundleAllOf) GetUpgradeEndTimeOk() (*time.Time, bool) {
	if o == nil || o.UpgradeEndTime == nil {
		return nil, false
	}
	return o.UpgradeEndTime, true
}

// HasUpgradeEndTime returns a boolean if a field has been set.
func (o *ApplianceImageBundleAllOf) HasUpgradeEndTime() bool {
	if o != nil && o.UpgradeEndTime != nil {
		return true
	}

	return false
}

// SetUpgradeEndTime gets a reference to the given time.Time and assigns it to the UpgradeEndTime field.
func (o *ApplianceImageBundleAllOf) SetUpgradeEndTime(v time.Time) {
	o.UpgradeEndTime = &v
}

// GetUpgradeGracePeriod returns the UpgradeGracePeriod field value if set, zero value otherwise.
func (o *ApplianceImageBundleAllOf) GetUpgradeGracePeriod() int64 {
	if o == nil || o.UpgradeGracePeriod == nil {
		var ret int64
		return ret
	}
	return *o.UpgradeGracePeriod
}

// GetUpgradeGracePeriodOk returns a tuple with the UpgradeGracePeriod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceImageBundleAllOf) GetUpgradeGracePeriodOk() (*int64, bool) {
	if o == nil || o.UpgradeGracePeriod == nil {
		return nil, false
	}
	return o.UpgradeGracePeriod, true
}

// HasUpgradeGracePeriod returns a boolean if a field has been set.
func (o *ApplianceImageBundleAllOf) HasUpgradeGracePeriod() bool {
	if o != nil && o.UpgradeGracePeriod != nil {
		return true
	}

	return false
}

// SetUpgradeGracePeriod gets a reference to the given int64 and assigns it to the UpgradeGracePeriod field.
func (o *ApplianceImageBundleAllOf) SetUpgradeGracePeriod(v int64) {
	o.UpgradeGracePeriod = &v
}

// GetUpgradeImpactDuration returns the UpgradeImpactDuration field value if set, zero value otherwise.
func (o *ApplianceImageBundleAllOf) GetUpgradeImpactDuration() int64 {
	if o == nil || o.UpgradeImpactDuration == nil {
		var ret int64
		return ret
	}
	return *o.UpgradeImpactDuration
}

// GetUpgradeImpactDurationOk returns a tuple with the UpgradeImpactDuration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceImageBundleAllOf) GetUpgradeImpactDurationOk() (*int64, bool) {
	if o == nil || o.UpgradeImpactDuration == nil {
		return nil, false
	}
	return o.UpgradeImpactDuration, true
}

// HasUpgradeImpactDuration returns a boolean if a field has been set.
func (o *ApplianceImageBundleAllOf) HasUpgradeImpactDuration() bool {
	if o != nil && o.UpgradeImpactDuration != nil {
		return true
	}

	return false
}

// SetUpgradeImpactDuration gets a reference to the given int64 and assigns it to the UpgradeImpactDuration field.
func (o *ApplianceImageBundleAllOf) SetUpgradeImpactDuration(v int64) {
	o.UpgradeImpactDuration = &v
}

// GetUpgradeImpactEnum returns the UpgradeImpactEnum field value if set, zero value otherwise.
func (o *ApplianceImageBundleAllOf) GetUpgradeImpactEnum() string {
	if o == nil || o.UpgradeImpactEnum == nil {
		var ret string
		return ret
	}
	return *o.UpgradeImpactEnum
}

// GetUpgradeImpactEnumOk returns a tuple with the UpgradeImpactEnum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceImageBundleAllOf) GetUpgradeImpactEnumOk() (*string, bool) {
	if o == nil || o.UpgradeImpactEnum == nil {
		return nil, false
	}
	return o.UpgradeImpactEnum, true
}

// HasUpgradeImpactEnum returns a boolean if a field has been set.
func (o *ApplianceImageBundleAllOf) HasUpgradeImpactEnum() bool {
	if o != nil && o.UpgradeImpactEnum != nil {
		return true
	}

	return false
}

// SetUpgradeImpactEnum gets a reference to the given string and assigns it to the UpgradeImpactEnum field.
func (o *ApplianceImageBundleAllOf) SetUpgradeImpactEnum(v string) {
	o.UpgradeImpactEnum = &v
}

// GetUpgradeStartTime returns the UpgradeStartTime field value if set, zero value otherwise.
func (o *ApplianceImageBundleAllOf) GetUpgradeStartTime() time.Time {
	if o == nil || o.UpgradeStartTime == nil {
		var ret time.Time
		return ret
	}
	return *o.UpgradeStartTime
}

// GetUpgradeStartTimeOk returns a tuple with the UpgradeStartTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceImageBundleAllOf) GetUpgradeStartTimeOk() (*time.Time, bool) {
	if o == nil || o.UpgradeStartTime == nil {
		return nil, false
	}
	return o.UpgradeStartTime, true
}

// HasUpgradeStartTime returns a boolean if a field has been set.
func (o *ApplianceImageBundleAllOf) HasUpgradeStartTime() bool {
	if o != nil && o.UpgradeStartTime != nil {
		return true
	}

	return false
}

// SetUpgradeStartTime gets a reference to the given time.Time and assigns it to the UpgradeStartTime field.
func (o *ApplianceImageBundleAllOf) SetUpgradeStartTime(v time.Time) {
	o.UpgradeStartTime = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *ApplianceImageBundleAllOf) GetVersion() string {
	if o == nil || o.Version == nil {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceImageBundleAllOf) GetVersionOk() (*string, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *ApplianceImageBundleAllOf) HasVersion() bool {
	if o != nil && o.Version != nil {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *ApplianceImageBundleAllOf) SetVersion(v string) {
	o.Version = &v
}

func (o ApplianceImageBundleAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.AnsiblePackages != nil {
		toSerialize["AnsiblePackages"] = o.AnsiblePackages
	}
	if o.AutoUpgrade != nil {
		toSerialize["AutoUpgrade"] = o.AutoUpgrade
	}
	if o.DcPackages != nil {
		toSerialize["DcPackages"] = o.DcPackages
	}
	if o.DebugPackages != nil {
		toSerialize["DebugPackages"] = o.DebugPackages
	}
	if o.Description != nil {
		toSerialize["Description"] = o.Description
	}
	if o.EndpointPackages != nil {
		toSerialize["EndpointPackages"] = o.EndpointPackages
	}
	if o.Fingerprint != nil {
		toSerialize["Fingerprint"] = o.Fingerprint
	}
	if o.HasError != nil {
		toSerialize["HasError"] = o.HasError
	}
	if o.InfraPackages != nil {
		toSerialize["InfraPackages"] = o.InfraPackages
	}
	if o.InitPackages != nil {
		toSerialize["InitPackages"] = o.InitPackages
	}
	if o.Name != nil {
		toSerialize["Name"] = o.Name
	}
	if o.Notes != nil {
		toSerialize["Notes"] = o.Notes
	}
	if o.Priority != nil {
		toSerialize["Priority"] = o.Priority
	}
	if o.ReleaseTime != nil {
		toSerialize["ReleaseTime"] = o.ReleaseTime
	}
	if o.ServicePackages != nil {
		toSerialize["ServicePackages"] = o.ServicePackages
	}
	if o.StatusMessage != nil {
		toSerialize["StatusMessage"] = o.StatusMessage
	}
	if o.SystemPackages != nil {
		toSerialize["SystemPackages"] = o.SystemPackages
	}
	if o.UiPackages != nil {
		toSerialize["UiPackages"] = o.UiPackages
	}
	if o.UpgradeEndTime != nil {
		toSerialize["UpgradeEndTime"] = o.UpgradeEndTime
	}
	if o.UpgradeGracePeriod != nil {
		toSerialize["UpgradeGracePeriod"] = o.UpgradeGracePeriod
	}
	if o.UpgradeImpactDuration != nil {
		toSerialize["UpgradeImpactDuration"] = o.UpgradeImpactDuration
	}
	if o.UpgradeImpactEnum != nil {
		toSerialize["UpgradeImpactEnum"] = o.UpgradeImpactEnum
	}
	if o.UpgradeStartTime != nil {
		toSerialize["UpgradeStartTime"] = o.UpgradeStartTime
	}
	if o.Version != nil {
		toSerialize["Version"] = o.Version
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ApplianceImageBundleAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varApplianceImageBundleAllOf := _ApplianceImageBundleAllOf{}

	if err = json.Unmarshal(bytes, &varApplianceImageBundleAllOf); err == nil {
		*o = ApplianceImageBundleAllOf(varApplianceImageBundleAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "AnsiblePackages")
		delete(additionalProperties, "AutoUpgrade")
		delete(additionalProperties, "DcPackages")
		delete(additionalProperties, "DebugPackages")
		delete(additionalProperties, "Description")
		delete(additionalProperties, "EndpointPackages")
		delete(additionalProperties, "Fingerprint")
		delete(additionalProperties, "HasError")
		delete(additionalProperties, "InfraPackages")
		delete(additionalProperties, "InitPackages")
		delete(additionalProperties, "Name")
		delete(additionalProperties, "Notes")
		delete(additionalProperties, "Priority")
		delete(additionalProperties, "ReleaseTime")
		delete(additionalProperties, "ServicePackages")
		delete(additionalProperties, "StatusMessage")
		delete(additionalProperties, "SystemPackages")
		delete(additionalProperties, "UiPackages")
		delete(additionalProperties, "UpgradeEndTime")
		delete(additionalProperties, "UpgradeGracePeriod")
		delete(additionalProperties, "UpgradeImpactDuration")
		delete(additionalProperties, "UpgradeImpactEnum")
		delete(additionalProperties, "UpgradeStartTime")
		delete(additionalProperties, "Version")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableApplianceImageBundleAllOf struct {
	value *ApplianceImageBundleAllOf
	isSet bool
}

func (v NullableApplianceImageBundleAllOf) Get() *ApplianceImageBundleAllOf {
	return v.value
}

func (v *NullableApplianceImageBundleAllOf) Set(val *ApplianceImageBundleAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableApplianceImageBundleAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableApplianceImageBundleAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApplianceImageBundleAllOf(val *ApplianceImageBundleAllOf) *NullableApplianceImageBundleAllOf {
	return &NullableApplianceImageBundleAllOf{value: val, isSet: true}
}

func (v NullableApplianceImageBundleAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApplianceImageBundleAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
