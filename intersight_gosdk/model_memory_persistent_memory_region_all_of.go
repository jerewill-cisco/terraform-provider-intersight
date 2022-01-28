/*
Cisco Intersight

Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document.

API version: 1.0.9-5208
Contact: intersight@cisco.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package intersight

import (
	"encoding/json"
)

// MemoryPersistentMemoryRegionAllOf Definition of the list of properties defined in 'memory.PersistentMemoryRegion', excluding properties defined in parent classes.
type MemoryPersistentMemoryRegionAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Free capacity in GiB of the Persistent Memory Region.
	FreeCapacity *string `json:"FreeCapacity,omitempty"`
	// Health state of the Persistent Memory Region.
	HealthState *string `json:"HealthState,omitempty"`
	// ID of the Interleaved Set formed for this Persistent Memory Region.
	InterleavedSetId *string `json:"InterleavedSetId,omitempty"`
	// Set of locator IDs that are included in the Persistent Memory Region.
	LocaterIds *string `json:"LocaterIds,omitempty"`
	// Persistent Memory type of the Persistent Memory Region.
	PersistentMemoryType *string `json:"PersistentMemoryType,omitempty"`
	// ID of the Persistent Memory Region.
	RegionId *string `json:"RegionId,omitempty"`
	// Socket ID of the Persistent Memory Region.
	SocketId *string `json:"SocketId,omitempty"`
	// Socket Memory ID of the Persistent Memory Region.
	SocketMemoryId *string `json:"SocketMemoryId,omitempty"`
	// Total capacity in GiB of the Persistent Memory Region.
	TotalCapacity                       *string                                          `json:"TotalCapacity,omitempty"`
	InventoryDeviceInfo                 *InventoryDeviceInfoRelationship                 `json:"InventoryDeviceInfo,omitempty"`
	MemoryPersistentMemoryConfiguration *MemoryPersistentMemoryConfigurationRelationship `json:"MemoryPersistentMemoryConfiguration,omitempty"`
	// An array of relationships to memoryPersistentMemoryNamespace resources.
	PersistentMemoryNamespaces []MemoryPersistentMemoryNamespaceRelationship `json:"PersistentMemoryNamespaces,omitempty"`
	RegisteredDevice           *AssetDeviceRegistrationRelationship          `json:"RegisteredDevice,omitempty"`
	AdditionalProperties       map[string]interface{}
}

type _MemoryPersistentMemoryRegionAllOf MemoryPersistentMemoryRegionAllOf

// NewMemoryPersistentMemoryRegionAllOf instantiates a new MemoryPersistentMemoryRegionAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMemoryPersistentMemoryRegionAllOf(classId string, objectType string) *MemoryPersistentMemoryRegionAllOf {
	this := MemoryPersistentMemoryRegionAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewMemoryPersistentMemoryRegionAllOfWithDefaults instantiates a new MemoryPersistentMemoryRegionAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMemoryPersistentMemoryRegionAllOfWithDefaults() *MemoryPersistentMemoryRegionAllOf {
	this := MemoryPersistentMemoryRegionAllOf{}
	var classId string = "memory.PersistentMemoryRegion"
	this.ClassId = classId
	var objectType string = "memory.PersistentMemoryRegion"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *MemoryPersistentMemoryRegionAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *MemoryPersistentMemoryRegionAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *MemoryPersistentMemoryRegionAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *MemoryPersistentMemoryRegionAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *MemoryPersistentMemoryRegionAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *MemoryPersistentMemoryRegionAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetFreeCapacity returns the FreeCapacity field value if set, zero value otherwise.
func (o *MemoryPersistentMemoryRegionAllOf) GetFreeCapacity() string {
	if o == nil || o.FreeCapacity == nil {
		var ret string
		return ret
	}
	return *o.FreeCapacity
}

// GetFreeCapacityOk returns a tuple with the FreeCapacity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MemoryPersistentMemoryRegionAllOf) GetFreeCapacityOk() (*string, bool) {
	if o == nil || o.FreeCapacity == nil {
		return nil, false
	}
	return o.FreeCapacity, true
}

// HasFreeCapacity returns a boolean if a field has been set.
func (o *MemoryPersistentMemoryRegionAllOf) HasFreeCapacity() bool {
	if o != nil && o.FreeCapacity != nil {
		return true
	}

	return false
}

// SetFreeCapacity gets a reference to the given string and assigns it to the FreeCapacity field.
func (o *MemoryPersistentMemoryRegionAllOf) SetFreeCapacity(v string) {
	o.FreeCapacity = &v
}

// GetHealthState returns the HealthState field value if set, zero value otherwise.
func (o *MemoryPersistentMemoryRegionAllOf) GetHealthState() string {
	if o == nil || o.HealthState == nil {
		var ret string
		return ret
	}
	return *o.HealthState
}

// GetHealthStateOk returns a tuple with the HealthState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MemoryPersistentMemoryRegionAllOf) GetHealthStateOk() (*string, bool) {
	if o == nil || o.HealthState == nil {
		return nil, false
	}
	return o.HealthState, true
}

// HasHealthState returns a boolean if a field has been set.
func (o *MemoryPersistentMemoryRegionAllOf) HasHealthState() bool {
	if o != nil && o.HealthState != nil {
		return true
	}

	return false
}

// SetHealthState gets a reference to the given string and assigns it to the HealthState field.
func (o *MemoryPersistentMemoryRegionAllOf) SetHealthState(v string) {
	o.HealthState = &v
}

// GetInterleavedSetId returns the InterleavedSetId field value if set, zero value otherwise.
func (o *MemoryPersistentMemoryRegionAllOf) GetInterleavedSetId() string {
	if o == nil || o.InterleavedSetId == nil {
		var ret string
		return ret
	}
	return *o.InterleavedSetId
}

// GetInterleavedSetIdOk returns a tuple with the InterleavedSetId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MemoryPersistentMemoryRegionAllOf) GetInterleavedSetIdOk() (*string, bool) {
	if o == nil || o.InterleavedSetId == nil {
		return nil, false
	}
	return o.InterleavedSetId, true
}

// HasInterleavedSetId returns a boolean if a field has been set.
func (o *MemoryPersistentMemoryRegionAllOf) HasInterleavedSetId() bool {
	if o != nil && o.InterleavedSetId != nil {
		return true
	}

	return false
}

// SetInterleavedSetId gets a reference to the given string and assigns it to the InterleavedSetId field.
func (o *MemoryPersistentMemoryRegionAllOf) SetInterleavedSetId(v string) {
	o.InterleavedSetId = &v
}

// GetLocaterIds returns the LocaterIds field value if set, zero value otherwise.
func (o *MemoryPersistentMemoryRegionAllOf) GetLocaterIds() string {
	if o == nil || o.LocaterIds == nil {
		var ret string
		return ret
	}
	return *o.LocaterIds
}

// GetLocaterIdsOk returns a tuple with the LocaterIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MemoryPersistentMemoryRegionAllOf) GetLocaterIdsOk() (*string, bool) {
	if o == nil || o.LocaterIds == nil {
		return nil, false
	}
	return o.LocaterIds, true
}

// HasLocaterIds returns a boolean if a field has been set.
func (o *MemoryPersistentMemoryRegionAllOf) HasLocaterIds() bool {
	if o != nil && o.LocaterIds != nil {
		return true
	}

	return false
}

// SetLocaterIds gets a reference to the given string and assigns it to the LocaterIds field.
func (o *MemoryPersistentMemoryRegionAllOf) SetLocaterIds(v string) {
	o.LocaterIds = &v
}

// GetPersistentMemoryType returns the PersistentMemoryType field value if set, zero value otherwise.
func (o *MemoryPersistentMemoryRegionAllOf) GetPersistentMemoryType() string {
	if o == nil || o.PersistentMemoryType == nil {
		var ret string
		return ret
	}
	return *o.PersistentMemoryType
}

// GetPersistentMemoryTypeOk returns a tuple with the PersistentMemoryType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MemoryPersistentMemoryRegionAllOf) GetPersistentMemoryTypeOk() (*string, bool) {
	if o == nil || o.PersistentMemoryType == nil {
		return nil, false
	}
	return o.PersistentMemoryType, true
}

// HasPersistentMemoryType returns a boolean if a field has been set.
func (o *MemoryPersistentMemoryRegionAllOf) HasPersistentMemoryType() bool {
	if o != nil && o.PersistentMemoryType != nil {
		return true
	}

	return false
}

// SetPersistentMemoryType gets a reference to the given string and assigns it to the PersistentMemoryType field.
func (o *MemoryPersistentMemoryRegionAllOf) SetPersistentMemoryType(v string) {
	o.PersistentMemoryType = &v
}

// GetRegionId returns the RegionId field value if set, zero value otherwise.
func (o *MemoryPersistentMemoryRegionAllOf) GetRegionId() string {
	if o == nil || o.RegionId == nil {
		var ret string
		return ret
	}
	return *o.RegionId
}

// GetRegionIdOk returns a tuple with the RegionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MemoryPersistentMemoryRegionAllOf) GetRegionIdOk() (*string, bool) {
	if o == nil || o.RegionId == nil {
		return nil, false
	}
	return o.RegionId, true
}

// HasRegionId returns a boolean if a field has been set.
func (o *MemoryPersistentMemoryRegionAllOf) HasRegionId() bool {
	if o != nil && o.RegionId != nil {
		return true
	}

	return false
}

// SetRegionId gets a reference to the given string and assigns it to the RegionId field.
func (o *MemoryPersistentMemoryRegionAllOf) SetRegionId(v string) {
	o.RegionId = &v
}

// GetSocketId returns the SocketId field value if set, zero value otherwise.
func (o *MemoryPersistentMemoryRegionAllOf) GetSocketId() string {
	if o == nil || o.SocketId == nil {
		var ret string
		return ret
	}
	return *o.SocketId
}

// GetSocketIdOk returns a tuple with the SocketId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MemoryPersistentMemoryRegionAllOf) GetSocketIdOk() (*string, bool) {
	if o == nil || o.SocketId == nil {
		return nil, false
	}
	return o.SocketId, true
}

// HasSocketId returns a boolean if a field has been set.
func (o *MemoryPersistentMemoryRegionAllOf) HasSocketId() bool {
	if o != nil && o.SocketId != nil {
		return true
	}

	return false
}

// SetSocketId gets a reference to the given string and assigns it to the SocketId field.
func (o *MemoryPersistentMemoryRegionAllOf) SetSocketId(v string) {
	o.SocketId = &v
}

// GetSocketMemoryId returns the SocketMemoryId field value if set, zero value otherwise.
func (o *MemoryPersistentMemoryRegionAllOf) GetSocketMemoryId() string {
	if o == nil || o.SocketMemoryId == nil {
		var ret string
		return ret
	}
	return *o.SocketMemoryId
}

// GetSocketMemoryIdOk returns a tuple with the SocketMemoryId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MemoryPersistentMemoryRegionAllOf) GetSocketMemoryIdOk() (*string, bool) {
	if o == nil || o.SocketMemoryId == nil {
		return nil, false
	}
	return o.SocketMemoryId, true
}

// HasSocketMemoryId returns a boolean if a field has been set.
func (o *MemoryPersistentMemoryRegionAllOf) HasSocketMemoryId() bool {
	if o != nil && o.SocketMemoryId != nil {
		return true
	}

	return false
}

// SetSocketMemoryId gets a reference to the given string and assigns it to the SocketMemoryId field.
func (o *MemoryPersistentMemoryRegionAllOf) SetSocketMemoryId(v string) {
	o.SocketMemoryId = &v
}

// GetTotalCapacity returns the TotalCapacity field value if set, zero value otherwise.
func (o *MemoryPersistentMemoryRegionAllOf) GetTotalCapacity() string {
	if o == nil || o.TotalCapacity == nil {
		var ret string
		return ret
	}
	return *o.TotalCapacity
}

// GetTotalCapacityOk returns a tuple with the TotalCapacity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MemoryPersistentMemoryRegionAllOf) GetTotalCapacityOk() (*string, bool) {
	if o == nil || o.TotalCapacity == nil {
		return nil, false
	}
	return o.TotalCapacity, true
}

// HasTotalCapacity returns a boolean if a field has been set.
func (o *MemoryPersistentMemoryRegionAllOf) HasTotalCapacity() bool {
	if o != nil && o.TotalCapacity != nil {
		return true
	}

	return false
}

// SetTotalCapacity gets a reference to the given string and assigns it to the TotalCapacity field.
func (o *MemoryPersistentMemoryRegionAllOf) SetTotalCapacity(v string) {
	o.TotalCapacity = &v
}

// GetInventoryDeviceInfo returns the InventoryDeviceInfo field value if set, zero value otherwise.
func (o *MemoryPersistentMemoryRegionAllOf) GetInventoryDeviceInfo() InventoryDeviceInfoRelationship {
	if o == nil || o.InventoryDeviceInfo == nil {
		var ret InventoryDeviceInfoRelationship
		return ret
	}
	return *o.InventoryDeviceInfo
}

// GetInventoryDeviceInfoOk returns a tuple with the InventoryDeviceInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MemoryPersistentMemoryRegionAllOf) GetInventoryDeviceInfoOk() (*InventoryDeviceInfoRelationship, bool) {
	if o == nil || o.InventoryDeviceInfo == nil {
		return nil, false
	}
	return o.InventoryDeviceInfo, true
}

// HasInventoryDeviceInfo returns a boolean if a field has been set.
func (o *MemoryPersistentMemoryRegionAllOf) HasInventoryDeviceInfo() bool {
	if o != nil && o.InventoryDeviceInfo != nil {
		return true
	}

	return false
}

// SetInventoryDeviceInfo gets a reference to the given InventoryDeviceInfoRelationship and assigns it to the InventoryDeviceInfo field.
func (o *MemoryPersistentMemoryRegionAllOf) SetInventoryDeviceInfo(v InventoryDeviceInfoRelationship) {
	o.InventoryDeviceInfo = &v
}

// GetMemoryPersistentMemoryConfiguration returns the MemoryPersistentMemoryConfiguration field value if set, zero value otherwise.
func (o *MemoryPersistentMemoryRegionAllOf) GetMemoryPersistentMemoryConfiguration() MemoryPersistentMemoryConfigurationRelationship {
	if o == nil || o.MemoryPersistentMemoryConfiguration == nil {
		var ret MemoryPersistentMemoryConfigurationRelationship
		return ret
	}
	return *o.MemoryPersistentMemoryConfiguration
}

// GetMemoryPersistentMemoryConfigurationOk returns a tuple with the MemoryPersistentMemoryConfiguration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MemoryPersistentMemoryRegionAllOf) GetMemoryPersistentMemoryConfigurationOk() (*MemoryPersistentMemoryConfigurationRelationship, bool) {
	if o == nil || o.MemoryPersistentMemoryConfiguration == nil {
		return nil, false
	}
	return o.MemoryPersistentMemoryConfiguration, true
}

// HasMemoryPersistentMemoryConfiguration returns a boolean if a field has been set.
func (o *MemoryPersistentMemoryRegionAllOf) HasMemoryPersistentMemoryConfiguration() bool {
	if o != nil && o.MemoryPersistentMemoryConfiguration != nil {
		return true
	}

	return false
}

// SetMemoryPersistentMemoryConfiguration gets a reference to the given MemoryPersistentMemoryConfigurationRelationship and assigns it to the MemoryPersistentMemoryConfiguration field.
func (o *MemoryPersistentMemoryRegionAllOf) SetMemoryPersistentMemoryConfiguration(v MemoryPersistentMemoryConfigurationRelationship) {
	o.MemoryPersistentMemoryConfiguration = &v
}

// GetPersistentMemoryNamespaces returns the PersistentMemoryNamespaces field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MemoryPersistentMemoryRegionAllOf) GetPersistentMemoryNamespaces() []MemoryPersistentMemoryNamespaceRelationship {
	if o == nil {
		var ret []MemoryPersistentMemoryNamespaceRelationship
		return ret
	}
	return o.PersistentMemoryNamespaces
}

// GetPersistentMemoryNamespacesOk returns a tuple with the PersistentMemoryNamespaces field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MemoryPersistentMemoryRegionAllOf) GetPersistentMemoryNamespacesOk() (*[]MemoryPersistentMemoryNamespaceRelationship, bool) {
	if o == nil || o.PersistentMemoryNamespaces == nil {
		return nil, false
	}
	return &o.PersistentMemoryNamespaces, true
}

// HasPersistentMemoryNamespaces returns a boolean if a field has been set.
func (o *MemoryPersistentMemoryRegionAllOf) HasPersistentMemoryNamespaces() bool {
	if o != nil && o.PersistentMemoryNamespaces != nil {
		return true
	}

	return false
}

// SetPersistentMemoryNamespaces gets a reference to the given []MemoryPersistentMemoryNamespaceRelationship and assigns it to the PersistentMemoryNamespaces field.
func (o *MemoryPersistentMemoryRegionAllOf) SetPersistentMemoryNamespaces(v []MemoryPersistentMemoryNamespaceRelationship) {
	o.PersistentMemoryNamespaces = v
}

// GetRegisteredDevice returns the RegisteredDevice field value if set, zero value otherwise.
func (o *MemoryPersistentMemoryRegionAllOf) GetRegisteredDevice() AssetDeviceRegistrationRelationship {
	if o == nil || o.RegisteredDevice == nil {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.RegisteredDevice
}

// GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MemoryPersistentMemoryRegionAllOf) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil || o.RegisteredDevice == nil {
		return nil, false
	}
	return o.RegisteredDevice, true
}

// HasRegisteredDevice returns a boolean if a field has been set.
func (o *MemoryPersistentMemoryRegionAllOf) HasRegisteredDevice() bool {
	if o != nil && o.RegisteredDevice != nil {
		return true
	}

	return false
}

// SetRegisteredDevice gets a reference to the given AssetDeviceRegistrationRelationship and assigns it to the RegisteredDevice field.
func (o *MemoryPersistentMemoryRegionAllOf) SetRegisteredDevice(v AssetDeviceRegistrationRelationship) {
	o.RegisteredDevice = &v
}

func (o MemoryPersistentMemoryRegionAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.FreeCapacity != nil {
		toSerialize["FreeCapacity"] = o.FreeCapacity
	}
	if o.HealthState != nil {
		toSerialize["HealthState"] = o.HealthState
	}
	if o.InterleavedSetId != nil {
		toSerialize["InterleavedSetId"] = o.InterleavedSetId
	}
	if o.LocaterIds != nil {
		toSerialize["LocaterIds"] = o.LocaterIds
	}
	if o.PersistentMemoryType != nil {
		toSerialize["PersistentMemoryType"] = o.PersistentMemoryType
	}
	if o.RegionId != nil {
		toSerialize["RegionId"] = o.RegionId
	}
	if o.SocketId != nil {
		toSerialize["SocketId"] = o.SocketId
	}
	if o.SocketMemoryId != nil {
		toSerialize["SocketMemoryId"] = o.SocketMemoryId
	}
	if o.TotalCapacity != nil {
		toSerialize["TotalCapacity"] = o.TotalCapacity
	}
	if o.InventoryDeviceInfo != nil {
		toSerialize["InventoryDeviceInfo"] = o.InventoryDeviceInfo
	}
	if o.MemoryPersistentMemoryConfiguration != nil {
		toSerialize["MemoryPersistentMemoryConfiguration"] = o.MemoryPersistentMemoryConfiguration
	}
	if o.PersistentMemoryNamespaces != nil {
		toSerialize["PersistentMemoryNamespaces"] = o.PersistentMemoryNamespaces
	}
	if o.RegisteredDevice != nil {
		toSerialize["RegisteredDevice"] = o.RegisteredDevice
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *MemoryPersistentMemoryRegionAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varMemoryPersistentMemoryRegionAllOf := _MemoryPersistentMemoryRegionAllOf{}

	if err = json.Unmarshal(bytes, &varMemoryPersistentMemoryRegionAllOf); err == nil {
		*o = MemoryPersistentMemoryRegionAllOf(varMemoryPersistentMemoryRegionAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "FreeCapacity")
		delete(additionalProperties, "HealthState")
		delete(additionalProperties, "InterleavedSetId")
		delete(additionalProperties, "LocaterIds")
		delete(additionalProperties, "PersistentMemoryType")
		delete(additionalProperties, "RegionId")
		delete(additionalProperties, "SocketId")
		delete(additionalProperties, "SocketMemoryId")
		delete(additionalProperties, "TotalCapacity")
		delete(additionalProperties, "InventoryDeviceInfo")
		delete(additionalProperties, "MemoryPersistentMemoryConfiguration")
		delete(additionalProperties, "PersistentMemoryNamespaces")
		delete(additionalProperties, "RegisteredDevice")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableMemoryPersistentMemoryRegionAllOf struct {
	value *MemoryPersistentMemoryRegionAllOf
	isSet bool
}

func (v NullableMemoryPersistentMemoryRegionAllOf) Get() *MemoryPersistentMemoryRegionAllOf {
	return v.value
}

func (v *NullableMemoryPersistentMemoryRegionAllOf) Set(val *MemoryPersistentMemoryRegionAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableMemoryPersistentMemoryRegionAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableMemoryPersistentMemoryRegionAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMemoryPersistentMemoryRegionAllOf(val *MemoryPersistentMemoryRegionAllOf) *NullableMemoryPersistentMemoryRegionAllOf {
	return &NullableMemoryPersistentMemoryRegionAllOf{value: val, isSet: true}
}

func (v NullableMemoryPersistentMemoryRegionAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMemoryPersistentMemoryRegionAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
