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
	"reflect"
	"strings"
)

// StorageHitachiVolume A volume entity in Hitachi storage array.
type StorageHitachiVolume struct {
	StorageBaseVolume
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string   `json:"ObjectType"`
	Attributes []string `json:"Attributes,omitempty"`
	// CLPR (Cache Logical Partition) number of this volume.
	ClprId *int64 `json:"ClprId,omitempty"`
	// Setting of the capacity saving function (dedupe and compression). * `N/A` - The capacity saving function is not available. * `Compression` - The capacity saving function (compression) is enabled. * `Compression Deduplication` - The capacity saving function (compression and deduplication) is enabled. * `Disabled` - The capacity saving function (compression and deduplication) is disabled.
	DataReductionMode *string `json:"DataReductionMode,omitempty"`
	// Status of the capacity saving function. * `N/A` - The capacity saving function is not available. * `Enabled` - The capacity saving function is enabled. * `Disabled` - The capacity saving function is disabled. * `Enabling` - The capacity saving function is being enabled. * `Rehydrating` - The capacity saving function is being disabled. * `Deleting` - The volumes for which the capacity saving function is enabled are being deleted. * `Failed` - An attempt to enable the capacity saving function failed.
	DataReductionStatus *string `json:"DataReductionStatus,omitempty"`
	// Code indicating the drive type of the drive belonging to the volume.
	DriveType *string `json:"DriveType,omitempty"`
	// The volume emulation type or the volume status information. * `N/A` - Not available. * `NOT DEFINED` - The volume is not implemented. * `DEFINING` - The volume is being created. * `REMOVING` - The volume is being removed. * `OPEN-V` - To be provided by Hitachi.
	EmulationType *string `json:"EmulationType,omitempty"`
	// Whether pages are reserved by the FullAllocation functionality.
	IsFullAllocationEnabled *bool `json:"IsFullAllocationEnabled,omitempty"`
	// Label of the volume, as configured in the storage array.
	Label *string `json:"Label,omitempty"`
	// Number of paths set for the volume.
	NumOfPaths     *int64   `json:"NumOfPaths,omitempty"`
	ParityGroupIds []string `json:"ParityGroupIds,omitempty"`
	// ID of the pool with which the volume is associated.
	PoolId *string `json:"PoolId,omitempty"`
	// RAID level for the volume. * `N/A` - RAID level is unknown or multiple RAID levels are being used. * `RAID1` - RAID1. * `RAID5` - RAID5. * `RAID6` - RAID6.
	RaidLevel *string `json:"RaidLevel,omitempty"`
	// RAID type drive configuration.
	RaidType *string `json:"RaidType,omitempty"`
	// Status information of the volume. * `N/A` - The volume status is not available. * `NML` - The volume is in normal status. * `BLK` - The volume is in blocked state. * `BSY` - The volume status is being changed. * `Unknown` - The volume status is unknown (not supported).
	Status *string                          `json:"Status,omitempty"`
	Array  *StorageHitachiArrayRelationship `json:"Array,omitempty"`
	// An array of relationships to storageHitachiParityGroup resources.
	ParityGroups         []StorageHitachiParityGroupRelationship `json:"ParityGroups,omitempty"`
	Pool                 *StorageHitachiPoolRelationship         `json:"Pool,omitempty"`
	RegisteredDevice     *AssetDeviceRegistrationRelationship    `json:"RegisteredDevice,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _StorageHitachiVolume StorageHitachiVolume

// NewStorageHitachiVolume instantiates a new StorageHitachiVolume object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStorageHitachiVolume(classId string, objectType string) *StorageHitachiVolume {
	this := StorageHitachiVolume{}
	this.ClassId = classId
	this.ObjectType = objectType
	var dataReductionMode string = "N/A"
	this.DataReductionMode = &dataReductionMode
	var dataReductionStatus string = "N/A"
	this.DataReductionStatus = &dataReductionStatus
	var emulationType string = "N/A"
	this.EmulationType = &emulationType
	var raidLevel string = "N/A"
	this.RaidLevel = &raidLevel
	var status string = "N/A"
	this.Status = &status
	return &this
}

// NewStorageHitachiVolumeWithDefaults instantiates a new StorageHitachiVolume object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStorageHitachiVolumeWithDefaults() *StorageHitachiVolume {
	this := StorageHitachiVolume{}
	var classId string = "storage.HitachiVolume"
	this.ClassId = classId
	var objectType string = "storage.HitachiVolume"
	this.ObjectType = objectType
	var dataReductionMode string = "N/A"
	this.DataReductionMode = &dataReductionMode
	var dataReductionStatus string = "N/A"
	this.DataReductionStatus = &dataReductionStatus
	var emulationType string = "N/A"
	this.EmulationType = &emulationType
	var raidLevel string = "N/A"
	this.RaidLevel = &raidLevel
	var status string = "N/A"
	this.Status = &status
	return &this
}

// GetClassId returns the ClassId field value
func (o *StorageHitachiVolume) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *StorageHitachiVolume) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *StorageHitachiVolume) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *StorageHitachiVolume) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *StorageHitachiVolume) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *StorageHitachiVolume) SetObjectType(v string) {
	o.ObjectType = v
}

// GetAttributes returns the Attributes field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StorageHitachiVolume) GetAttributes() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StorageHitachiVolume) GetAttributesOk() (*[]string, bool) {
	if o == nil || o.Attributes == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *StorageHitachiVolume) HasAttributes() bool {
	if o != nil && o.Attributes != nil {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given []string and assigns it to the Attributes field.
func (o *StorageHitachiVolume) SetAttributes(v []string) {
	o.Attributes = v
}

// GetClprId returns the ClprId field value if set, zero value otherwise.
func (o *StorageHitachiVolume) GetClprId() int64 {
	if o == nil || o.ClprId == nil {
		var ret int64
		return ret
	}
	return *o.ClprId
}

// GetClprIdOk returns a tuple with the ClprId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageHitachiVolume) GetClprIdOk() (*int64, bool) {
	if o == nil || o.ClprId == nil {
		return nil, false
	}
	return o.ClprId, true
}

// HasClprId returns a boolean if a field has been set.
func (o *StorageHitachiVolume) HasClprId() bool {
	if o != nil && o.ClprId != nil {
		return true
	}

	return false
}

// SetClprId gets a reference to the given int64 and assigns it to the ClprId field.
func (o *StorageHitachiVolume) SetClprId(v int64) {
	o.ClprId = &v
}

// GetDataReductionMode returns the DataReductionMode field value if set, zero value otherwise.
func (o *StorageHitachiVolume) GetDataReductionMode() string {
	if o == nil || o.DataReductionMode == nil {
		var ret string
		return ret
	}
	return *o.DataReductionMode
}

// GetDataReductionModeOk returns a tuple with the DataReductionMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageHitachiVolume) GetDataReductionModeOk() (*string, bool) {
	if o == nil || o.DataReductionMode == nil {
		return nil, false
	}
	return o.DataReductionMode, true
}

// HasDataReductionMode returns a boolean if a field has been set.
func (o *StorageHitachiVolume) HasDataReductionMode() bool {
	if o != nil && o.DataReductionMode != nil {
		return true
	}

	return false
}

// SetDataReductionMode gets a reference to the given string and assigns it to the DataReductionMode field.
func (o *StorageHitachiVolume) SetDataReductionMode(v string) {
	o.DataReductionMode = &v
}

// GetDataReductionStatus returns the DataReductionStatus field value if set, zero value otherwise.
func (o *StorageHitachiVolume) GetDataReductionStatus() string {
	if o == nil || o.DataReductionStatus == nil {
		var ret string
		return ret
	}
	return *o.DataReductionStatus
}

// GetDataReductionStatusOk returns a tuple with the DataReductionStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageHitachiVolume) GetDataReductionStatusOk() (*string, bool) {
	if o == nil || o.DataReductionStatus == nil {
		return nil, false
	}
	return o.DataReductionStatus, true
}

// HasDataReductionStatus returns a boolean if a field has been set.
func (o *StorageHitachiVolume) HasDataReductionStatus() bool {
	if o != nil && o.DataReductionStatus != nil {
		return true
	}

	return false
}

// SetDataReductionStatus gets a reference to the given string and assigns it to the DataReductionStatus field.
func (o *StorageHitachiVolume) SetDataReductionStatus(v string) {
	o.DataReductionStatus = &v
}

// GetDriveType returns the DriveType field value if set, zero value otherwise.
func (o *StorageHitachiVolume) GetDriveType() string {
	if o == nil || o.DriveType == nil {
		var ret string
		return ret
	}
	return *o.DriveType
}

// GetDriveTypeOk returns a tuple with the DriveType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageHitachiVolume) GetDriveTypeOk() (*string, bool) {
	if o == nil || o.DriveType == nil {
		return nil, false
	}
	return o.DriveType, true
}

// HasDriveType returns a boolean if a field has been set.
func (o *StorageHitachiVolume) HasDriveType() bool {
	if o != nil && o.DriveType != nil {
		return true
	}

	return false
}

// SetDriveType gets a reference to the given string and assigns it to the DriveType field.
func (o *StorageHitachiVolume) SetDriveType(v string) {
	o.DriveType = &v
}

// GetEmulationType returns the EmulationType field value if set, zero value otherwise.
func (o *StorageHitachiVolume) GetEmulationType() string {
	if o == nil || o.EmulationType == nil {
		var ret string
		return ret
	}
	return *o.EmulationType
}

// GetEmulationTypeOk returns a tuple with the EmulationType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageHitachiVolume) GetEmulationTypeOk() (*string, bool) {
	if o == nil || o.EmulationType == nil {
		return nil, false
	}
	return o.EmulationType, true
}

// HasEmulationType returns a boolean if a field has been set.
func (o *StorageHitachiVolume) HasEmulationType() bool {
	if o != nil && o.EmulationType != nil {
		return true
	}

	return false
}

// SetEmulationType gets a reference to the given string and assigns it to the EmulationType field.
func (o *StorageHitachiVolume) SetEmulationType(v string) {
	o.EmulationType = &v
}

// GetIsFullAllocationEnabled returns the IsFullAllocationEnabled field value if set, zero value otherwise.
func (o *StorageHitachiVolume) GetIsFullAllocationEnabled() bool {
	if o == nil || o.IsFullAllocationEnabled == nil {
		var ret bool
		return ret
	}
	return *o.IsFullAllocationEnabled
}

// GetIsFullAllocationEnabledOk returns a tuple with the IsFullAllocationEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageHitachiVolume) GetIsFullAllocationEnabledOk() (*bool, bool) {
	if o == nil || o.IsFullAllocationEnabled == nil {
		return nil, false
	}
	return o.IsFullAllocationEnabled, true
}

// HasIsFullAllocationEnabled returns a boolean if a field has been set.
func (o *StorageHitachiVolume) HasIsFullAllocationEnabled() bool {
	if o != nil && o.IsFullAllocationEnabled != nil {
		return true
	}

	return false
}

// SetIsFullAllocationEnabled gets a reference to the given bool and assigns it to the IsFullAllocationEnabled field.
func (o *StorageHitachiVolume) SetIsFullAllocationEnabled(v bool) {
	o.IsFullAllocationEnabled = &v
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *StorageHitachiVolume) GetLabel() string {
	if o == nil || o.Label == nil {
		var ret string
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageHitachiVolume) GetLabelOk() (*string, bool) {
	if o == nil || o.Label == nil {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *StorageHitachiVolume) HasLabel() bool {
	if o != nil && o.Label != nil {
		return true
	}

	return false
}

// SetLabel gets a reference to the given string and assigns it to the Label field.
func (o *StorageHitachiVolume) SetLabel(v string) {
	o.Label = &v
}

// GetNumOfPaths returns the NumOfPaths field value if set, zero value otherwise.
func (o *StorageHitachiVolume) GetNumOfPaths() int64 {
	if o == nil || o.NumOfPaths == nil {
		var ret int64
		return ret
	}
	return *o.NumOfPaths
}

// GetNumOfPathsOk returns a tuple with the NumOfPaths field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageHitachiVolume) GetNumOfPathsOk() (*int64, bool) {
	if o == nil || o.NumOfPaths == nil {
		return nil, false
	}
	return o.NumOfPaths, true
}

// HasNumOfPaths returns a boolean if a field has been set.
func (o *StorageHitachiVolume) HasNumOfPaths() bool {
	if o != nil && o.NumOfPaths != nil {
		return true
	}

	return false
}

// SetNumOfPaths gets a reference to the given int64 and assigns it to the NumOfPaths field.
func (o *StorageHitachiVolume) SetNumOfPaths(v int64) {
	o.NumOfPaths = &v
}

// GetParityGroupIds returns the ParityGroupIds field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StorageHitachiVolume) GetParityGroupIds() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.ParityGroupIds
}

// GetParityGroupIdsOk returns a tuple with the ParityGroupIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StorageHitachiVolume) GetParityGroupIdsOk() (*[]string, bool) {
	if o == nil || o.ParityGroupIds == nil {
		return nil, false
	}
	return &o.ParityGroupIds, true
}

// HasParityGroupIds returns a boolean if a field has been set.
func (o *StorageHitachiVolume) HasParityGroupIds() bool {
	if o != nil && o.ParityGroupIds != nil {
		return true
	}

	return false
}

// SetParityGroupIds gets a reference to the given []string and assigns it to the ParityGroupIds field.
func (o *StorageHitachiVolume) SetParityGroupIds(v []string) {
	o.ParityGroupIds = v
}

// GetPoolId returns the PoolId field value if set, zero value otherwise.
func (o *StorageHitachiVolume) GetPoolId() string {
	if o == nil || o.PoolId == nil {
		var ret string
		return ret
	}
	return *o.PoolId
}

// GetPoolIdOk returns a tuple with the PoolId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageHitachiVolume) GetPoolIdOk() (*string, bool) {
	if o == nil || o.PoolId == nil {
		return nil, false
	}
	return o.PoolId, true
}

// HasPoolId returns a boolean if a field has been set.
func (o *StorageHitachiVolume) HasPoolId() bool {
	if o != nil && o.PoolId != nil {
		return true
	}

	return false
}

// SetPoolId gets a reference to the given string and assigns it to the PoolId field.
func (o *StorageHitachiVolume) SetPoolId(v string) {
	o.PoolId = &v
}

// GetRaidLevel returns the RaidLevel field value if set, zero value otherwise.
func (o *StorageHitachiVolume) GetRaidLevel() string {
	if o == nil || o.RaidLevel == nil {
		var ret string
		return ret
	}
	return *o.RaidLevel
}

// GetRaidLevelOk returns a tuple with the RaidLevel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageHitachiVolume) GetRaidLevelOk() (*string, bool) {
	if o == nil || o.RaidLevel == nil {
		return nil, false
	}
	return o.RaidLevel, true
}

// HasRaidLevel returns a boolean if a field has been set.
func (o *StorageHitachiVolume) HasRaidLevel() bool {
	if o != nil && o.RaidLevel != nil {
		return true
	}

	return false
}

// SetRaidLevel gets a reference to the given string and assigns it to the RaidLevel field.
func (o *StorageHitachiVolume) SetRaidLevel(v string) {
	o.RaidLevel = &v
}

// GetRaidType returns the RaidType field value if set, zero value otherwise.
func (o *StorageHitachiVolume) GetRaidType() string {
	if o == nil || o.RaidType == nil {
		var ret string
		return ret
	}
	return *o.RaidType
}

// GetRaidTypeOk returns a tuple with the RaidType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageHitachiVolume) GetRaidTypeOk() (*string, bool) {
	if o == nil || o.RaidType == nil {
		return nil, false
	}
	return o.RaidType, true
}

// HasRaidType returns a boolean if a field has been set.
func (o *StorageHitachiVolume) HasRaidType() bool {
	if o != nil && o.RaidType != nil {
		return true
	}

	return false
}

// SetRaidType gets a reference to the given string and assigns it to the RaidType field.
func (o *StorageHitachiVolume) SetRaidType(v string) {
	o.RaidType = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *StorageHitachiVolume) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageHitachiVolume) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *StorageHitachiVolume) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *StorageHitachiVolume) SetStatus(v string) {
	o.Status = &v
}

// GetArray returns the Array field value if set, zero value otherwise.
func (o *StorageHitachiVolume) GetArray() StorageHitachiArrayRelationship {
	if o == nil || o.Array == nil {
		var ret StorageHitachiArrayRelationship
		return ret
	}
	return *o.Array
}

// GetArrayOk returns a tuple with the Array field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageHitachiVolume) GetArrayOk() (*StorageHitachiArrayRelationship, bool) {
	if o == nil || o.Array == nil {
		return nil, false
	}
	return o.Array, true
}

// HasArray returns a boolean if a field has been set.
func (o *StorageHitachiVolume) HasArray() bool {
	if o != nil && o.Array != nil {
		return true
	}

	return false
}

// SetArray gets a reference to the given StorageHitachiArrayRelationship and assigns it to the Array field.
func (o *StorageHitachiVolume) SetArray(v StorageHitachiArrayRelationship) {
	o.Array = &v
}

// GetParityGroups returns the ParityGroups field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StorageHitachiVolume) GetParityGroups() []StorageHitachiParityGroupRelationship {
	if o == nil {
		var ret []StorageHitachiParityGroupRelationship
		return ret
	}
	return o.ParityGroups
}

// GetParityGroupsOk returns a tuple with the ParityGroups field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StorageHitachiVolume) GetParityGroupsOk() (*[]StorageHitachiParityGroupRelationship, bool) {
	if o == nil || o.ParityGroups == nil {
		return nil, false
	}
	return &o.ParityGroups, true
}

// HasParityGroups returns a boolean if a field has been set.
func (o *StorageHitachiVolume) HasParityGroups() bool {
	if o != nil && o.ParityGroups != nil {
		return true
	}

	return false
}

// SetParityGroups gets a reference to the given []StorageHitachiParityGroupRelationship and assigns it to the ParityGroups field.
func (o *StorageHitachiVolume) SetParityGroups(v []StorageHitachiParityGroupRelationship) {
	o.ParityGroups = v
}

// GetPool returns the Pool field value if set, zero value otherwise.
func (o *StorageHitachiVolume) GetPool() StorageHitachiPoolRelationship {
	if o == nil || o.Pool == nil {
		var ret StorageHitachiPoolRelationship
		return ret
	}
	return *o.Pool
}

// GetPoolOk returns a tuple with the Pool field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageHitachiVolume) GetPoolOk() (*StorageHitachiPoolRelationship, bool) {
	if o == nil || o.Pool == nil {
		return nil, false
	}
	return o.Pool, true
}

// HasPool returns a boolean if a field has been set.
func (o *StorageHitachiVolume) HasPool() bool {
	if o != nil && o.Pool != nil {
		return true
	}

	return false
}

// SetPool gets a reference to the given StorageHitachiPoolRelationship and assigns it to the Pool field.
func (o *StorageHitachiVolume) SetPool(v StorageHitachiPoolRelationship) {
	o.Pool = &v
}

// GetRegisteredDevice returns the RegisteredDevice field value if set, zero value otherwise.
func (o *StorageHitachiVolume) GetRegisteredDevice() AssetDeviceRegistrationRelationship {
	if o == nil || o.RegisteredDevice == nil {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.RegisteredDevice
}

// GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageHitachiVolume) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil || o.RegisteredDevice == nil {
		return nil, false
	}
	return o.RegisteredDevice, true
}

// HasRegisteredDevice returns a boolean if a field has been set.
func (o *StorageHitachiVolume) HasRegisteredDevice() bool {
	if o != nil && o.RegisteredDevice != nil {
		return true
	}

	return false
}

// SetRegisteredDevice gets a reference to the given AssetDeviceRegistrationRelationship and assigns it to the RegisteredDevice field.
func (o *StorageHitachiVolume) SetRegisteredDevice(v AssetDeviceRegistrationRelationship) {
	o.RegisteredDevice = &v
}

func (o StorageHitachiVolume) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedStorageBaseVolume, errStorageBaseVolume := json.Marshal(o.StorageBaseVolume)
	if errStorageBaseVolume != nil {
		return []byte{}, errStorageBaseVolume
	}
	errStorageBaseVolume = json.Unmarshal([]byte(serializedStorageBaseVolume), &toSerialize)
	if errStorageBaseVolume != nil {
		return []byte{}, errStorageBaseVolume
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Attributes != nil {
		toSerialize["Attributes"] = o.Attributes
	}
	if o.ClprId != nil {
		toSerialize["ClprId"] = o.ClprId
	}
	if o.DataReductionMode != nil {
		toSerialize["DataReductionMode"] = o.DataReductionMode
	}
	if o.DataReductionStatus != nil {
		toSerialize["DataReductionStatus"] = o.DataReductionStatus
	}
	if o.DriveType != nil {
		toSerialize["DriveType"] = o.DriveType
	}
	if o.EmulationType != nil {
		toSerialize["EmulationType"] = o.EmulationType
	}
	if o.IsFullAllocationEnabled != nil {
		toSerialize["IsFullAllocationEnabled"] = o.IsFullAllocationEnabled
	}
	if o.Label != nil {
		toSerialize["Label"] = o.Label
	}
	if o.NumOfPaths != nil {
		toSerialize["NumOfPaths"] = o.NumOfPaths
	}
	if o.ParityGroupIds != nil {
		toSerialize["ParityGroupIds"] = o.ParityGroupIds
	}
	if o.PoolId != nil {
		toSerialize["PoolId"] = o.PoolId
	}
	if o.RaidLevel != nil {
		toSerialize["RaidLevel"] = o.RaidLevel
	}
	if o.RaidType != nil {
		toSerialize["RaidType"] = o.RaidType
	}
	if o.Status != nil {
		toSerialize["Status"] = o.Status
	}
	if o.Array != nil {
		toSerialize["Array"] = o.Array
	}
	if o.ParityGroups != nil {
		toSerialize["ParityGroups"] = o.ParityGroups
	}
	if o.Pool != nil {
		toSerialize["Pool"] = o.Pool
	}
	if o.RegisteredDevice != nil {
		toSerialize["RegisteredDevice"] = o.RegisteredDevice
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *StorageHitachiVolume) UnmarshalJSON(bytes []byte) (err error) {
	type StorageHitachiVolumeWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string   `json:"ObjectType"`
		Attributes []string `json:"Attributes,omitempty"`
		// CLPR (Cache Logical Partition) number of this volume.
		ClprId *int64 `json:"ClprId,omitempty"`
		// Setting of the capacity saving function (dedupe and compression). * `N/A` - The capacity saving function is not available. * `Compression` - The capacity saving function (compression) is enabled. * `Compression Deduplication` - The capacity saving function (compression and deduplication) is enabled. * `Disabled` - The capacity saving function (compression and deduplication) is disabled.
		DataReductionMode *string `json:"DataReductionMode,omitempty"`
		// Status of the capacity saving function. * `N/A` - The capacity saving function is not available. * `Enabled` - The capacity saving function is enabled. * `Disabled` - The capacity saving function is disabled. * `Enabling` - The capacity saving function is being enabled. * `Rehydrating` - The capacity saving function is being disabled. * `Deleting` - The volumes for which the capacity saving function is enabled are being deleted. * `Failed` - An attempt to enable the capacity saving function failed.
		DataReductionStatus *string `json:"DataReductionStatus,omitempty"`
		// Code indicating the drive type of the drive belonging to the volume.
		DriveType *string `json:"DriveType,omitempty"`
		// The volume emulation type or the volume status information. * `N/A` - Not available. * `NOT DEFINED` - The volume is not implemented. * `DEFINING` - The volume is being created. * `REMOVING` - The volume is being removed. * `OPEN-V` - To be provided by Hitachi.
		EmulationType *string `json:"EmulationType,omitempty"`
		// Whether pages are reserved by the FullAllocation functionality.
		IsFullAllocationEnabled *bool `json:"IsFullAllocationEnabled,omitempty"`
		// Label of the volume, as configured in the storage array.
		Label *string `json:"Label,omitempty"`
		// Number of paths set for the volume.
		NumOfPaths     *int64   `json:"NumOfPaths,omitempty"`
		ParityGroupIds []string `json:"ParityGroupIds,omitempty"`
		// ID of the pool with which the volume is associated.
		PoolId *string `json:"PoolId,omitempty"`
		// RAID level for the volume. * `N/A` - RAID level is unknown or multiple RAID levels are being used. * `RAID1` - RAID1. * `RAID5` - RAID5. * `RAID6` - RAID6.
		RaidLevel *string `json:"RaidLevel,omitempty"`
		// RAID type drive configuration.
		RaidType *string `json:"RaidType,omitempty"`
		// Status information of the volume. * `N/A` - The volume status is not available. * `NML` - The volume is in normal status. * `BLK` - The volume is in blocked state. * `BSY` - The volume status is being changed. * `Unknown` - The volume status is unknown (not supported).
		Status *string                          `json:"Status,omitempty"`
		Array  *StorageHitachiArrayRelationship `json:"Array,omitempty"`
		// An array of relationships to storageHitachiParityGroup resources.
		ParityGroups     []StorageHitachiParityGroupRelationship `json:"ParityGroups,omitempty"`
		Pool             *StorageHitachiPoolRelationship         `json:"Pool,omitempty"`
		RegisteredDevice *AssetDeviceRegistrationRelationship    `json:"RegisteredDevice,omitempty"`
	}

	varStorageHitachiVolumeWithoutEmbeddedStruct := StorageHitachiVolumeWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varStorageHitachiVolumeWithoutEmbeddedStruct)
	if err == nil {
		varStorageHitachiVolume := _StorageHitachiVolume{}
		varStorageHitachiVolume.ClassId = varStorageHitachiVolumeWithoutEmbeddedStruct.ClassId
		varStorageHitachiVolume.ObjectType = varStorageHitachiVolumeWithoutEmbeddedStruct.ObjectType
		varStorageHitachiVolume.Attributes = varStorageHitachiVolumeWithoutEmbeddedStruct.Attributes
		varStorageHitachiVolume.ClprId = varStorageHitachiVolumeWithoutEmbeddedStruct.ClprId
		varStorageHitachiVolume.DataReductionMode = varStorageHitachiVolumeWithoutEmbeddedStruct.DataReductionMode
		varStorageHitachiVolume.DataReductionStatus = varStorageHitachiVolumeWithoutEmbeddedStruct.DataReductionStatus
		varStorageHitachiVolume.DriveType = varStorageHitachiVolumeWithoutEmbeddedStruct.DriveType
		varStorageHitachiVolume.EmulationType = varStorageHitachiVolumeWithoutEmbeddedStruct.EmulationType
		varStorageHitachiVolume.IsFullAllocationEnabled = varStorageHitachiVolumeWithoutEmbeddedStruct.IsFullAllocationEnabled
		varStorageHitachiVolume.Label = varStorageHitachiVolumeWithoutEmbeddedStruct.Label
		varStorageHitachiVolume.NumOfPaths = varStorageHitachiVolumeWithoutEmbeddedStruct.NumOfPaths
		varStorageHitachiVolume.ParityGroupIds = varStorageHitachiVolumeWithoutEmbeddedStruct.ParityGroupIds
		varStorageHitachiVolume.PoolId = varStorageHitachiVolumeWithoutEmbeddedStruct.PoolId
		varStorageHitachiVolume.RaidLevel = varStorageHitachiVolumeWithoutEmbeddedStruct.RaidLevel
		varStorageHitachiVolume.RaidType = varStorageHitachiVolumeWithoutEmbeddedStruct.RaidType
		varStorageHitachiVolume.Status = varStorageHitachiVolumeWithoutEmbeddedStruct.Status
		varStorageHitachiVolume.Array = varStorageHitachiVolumeWithoutEmbeddedStruct.Array
		varStorageHitachiVolume.ParityGroups = varStorageHitachiVolumeWithoutEmbeddedStruct.ParityGroups
		varStorageHitachiVolume.Pool = varStorageHitachiVolumeWithoutEmbeddedStruct.Pool
		varStorageHitachiVolume.RegisteredDevice = varStorageHitachiVolumeWithoutEmbeddedStruct.RegisteredDevice
		*o = StorageHitachiVolume(varStorageHitachiVolume)
	} else {
		return err
	}

	varStorageHitachiVolume := _StorageHitachiVolume{}

	err = json.Unmarshal(bytes, &varStorageHitachiVolume)
	if err == nil {
		o.StorageBaseVolume = varStorageHitachiVolume.StorageBaseVolume
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Attributes")
		delete(additionalProperties, "ClprId")
		delete(additionalProperties, "DataReductionMode")
		delete(additionalProperties, "DataReductionStatus")
		delete(additionalProperties, "DriveType")
		delete(additionalProperties, "EmulationType")
		delete(additionalProperties, "IsFullAllocationEnabled")
		delete(additionalProperties, "Label")
		delete(additionalProperties, "NumOfPaths")
		delete(additionalProperties, "ParityGroupIds")
		delete(additionalProperties, "PoolId")
		delete(additionalProperties, "RaidLevel")
		delete(additionalProperties, "RaidType")
		delete(additionalProperties, "Status")
		delete(additionalProperties, "Array")
		delete(additionalProperties, "ParityGroups")
		delete(additionalProperties, "Pool")
		delete(additionalProperties, "RegisteredDevice")

		// remove fields from embedded structs
		reflectStorageBaseVolume := reflect.ValueOf(o.StorageBaseVolume)
		for i := 0; i < reflectStorageBaseVolume.Type().NumField(); i++ {
			t := reflectStorageBaseVolume.Type().Field(i)

			if jsonTag := t.Tag.Get("json"); jsonTag != "" {
				fieldName := ""
				if commaIdx := strings.Index(jsonTag, ","); commaIdx > 0 {
					fieldName = jsonTag[:commaIdx]
				} else {
					fieldName = jsonTag
				}
				if fieldName != "AdditionalProperties" {
					delete(additionalProperties, fieldName)
				}
			}
		}

		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableStorageHitachiVolume struct {
	value *StorageHitachiVolume
	isSet bool
}

func (v NullableStorageHitachiVolume) Get() *StorageHitachiVolume {
	return v.value
}

func (v *NullableStorageHitachiVolume) Set(val *StorageHitachiVolume) {
	v.value = val
	v.isSet = true
}

func (v NullableStorageHitachiVolume) IsSet() bool {
	return v.isSet
}

func (v *NullableStorageHitachiVolume) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStorageHitachiVolume(val *StorageHitachiVolume) *NullableStorageHitachiVolume {
	return &NullableStorageHitachiVolume{value: val, isSet: true}
}

func (v NullableStorageHitachiVolume) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStorageHitachiVolume) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
