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

// StorageDiskGroup Group of one or more Spans to configure virtual drive.
type StorageDiskGroup struct {
	InventoryBase
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Name to identity this disk group in the controller.
	Name *string `json:"Name,omitempty"`
	// Raid level of the virtual drives in this diskgroup.
	RaidType *string `json:"RaidType,omitempty"`
	// An array of relationships to storagePhysicalDisk resources.
	DedicatedHotSpares []StoragePhysicalDiskRelationship    `json:"DedicatedHotSpares,omitempty"`
	RegisteredDevice   *AssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	// An array of relationships to storageSpan resources.
	Spans             []StorageSpanRelationship      `json:"Spans,omitempty"`
	StorageController *StorageControllerRelationship `json:"StorageController,omitempty"`
	// An array of relationships to storageVirtualDrive resources.
	VirtualDrives        []StorageVirtualDriveRelationship `json:"VirtualDrives,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _StorageDiskGroup StorageDiskGroup

// NewStorageDiskGroup instantiates a new StorageDiskGroup object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStorageDiskGroup(classId string, objectType string) *StorageDiskGroup {
	this := StorageDiskGroup{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewStorageDiskGroupWithDefaults instantiates a new StorageDiskGroup object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStorageDiskGroupWithDefaults() *StorageDiskGroup {
	this := StorageDiskGroup{}
	var classId string = "storage.DiskGroup"
	this.ClassId = classId
	var objectType string = "storage.DiskGroup"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *StorageDiskGroup) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *StorageDiskGroup) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *StorageDiskGroup) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *StorageDiskGroup) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *StorageDiskGroup) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *StorageDiskGroup) SetObjectType(v string) {
	o.ObjectType = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *StorageDiskGroup) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageDiskGroup) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *StorageDiskGroup) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *StorageDiskGroup) SetName(v string) {
	o.Name = &v
}

// GetRaidType returns the RaidType field value if set, zero value otherwise.
func (o *StorageDiskGroup) GetRaidType() string {
	if o == nil || o.RaidType == nil {
		var ret string
		return ret
	}
	return *o.RaidType
}

// GetRaidTypeOk returns a tuple with the RaidType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageDiskGroup) GetRaidTypeOk() (*string, bool) {
	if o == nil || o.RaidType == nil {
		return nil, false
	}
	return o.RaidType, true
}

// HasRaidType returns a boolean if a field has been set.
func (o *StorageDiskGroup) HasRaidType() bool {
	if o != nil && o.RaidType != nil {
		return true
	}

	return false
}

// SetRaidType gets a reference to the given string and assigns it to the RaidType field.
func (o *StorageDiskGroup) SetRaidType(v string) {
	o.RaidType = &v
}

// GetDedicatedHotSpares returns the DedicatedHotSpares field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StorageDiskGroup) GetDedicatedHotSpares() []StoragePhysicalDiskRelationship {
	if o == nil {
		var ret []StoragePhysicalDiskRelationship
		return ret
	}
	return o.DedicatedHotSpares
}

// GetDedicatedHotSparesOk returns a tuple with the DedicatedHotSpares field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StorageDiskGroup) GetDedicatedHotSparesOk() (*[]StoragePhysicalDiskRelationship, bool) {
	if o == nil || o.DedicatedHotSpares == nil {
		return nil, false
	}
	return &o.DedicatedHotSpares, true
}

// HasDedicatedHotSpares returns a boolean if a field has been set.
func (o *StorageDiskGroup) HasDedicatedHotSpares() bool {
	if o != nil && o.DedicatedHotSpares != nil {
		return true
	}

	return false
}

// SetDedicatedHotSpares gets a reference to the given []StoragePhysicalDiskRelationship and assigns it to the DedicatedHotSpares field.
func (o *StorageDiskGroup) SetDedicatedHotSpares(v []StoragePhysicalDiskRelationship) {
	o.DedicatedHotSpares = v
}

// GetRegisteredDevice returns the RegisteredDevice field value if set, zero value otherwise.
func (o *StorageDiskGroup) GetRegisteredDevice() AssetDeviceRegistrationRelationship {
	if o == nil || o.RegisteredDevice == nil {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.RegisteredDevice
}

// GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageDiskGroup) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil || o.RegisteredDevice == nil {
		return nil, false
	}
	return o.RegisteredDevice, true
}

// HasRegisteredDevice returns a boolean if a field has been set.
func (o *StorageDiskGroup) HasRegisteredDevice() bool {
	if o != nil && o.RegisteredDevice != nil {
		return true
	}

	return false
}

// SetRegisteredDevice gets a reference to the given AssetDeviceRegistrationRelationship and assigns it to the RegisteredDevice field.
func (o *StorageDiskGroup) SetRegisteredDevice(v AssetDeviceRegistrationRelationship) {
	o.RegisteredDevice = &v
}

// GetSpans returns the Spans field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StorageDiskGroup) GetSpans() []StorageSpanRelationship {
	if o == nil {
		var ret []StorageSpanRelationship
		return ret
	}
	return o.Spans
}

// GetSpansOk returns a tuple with the Spans field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StorageDiskGroup) GetSpansOk() (*[]StorageSpanRelationship, bool) {
	if o == nil || o.Spans == nil {
		return nil, false
	}
	return &o.Spans, true
}

// HasSpans returns a boolean if a field has been set.
func (o *StorageDiskGroup) HasSpans() bool {
	if o != nil && o.Spans != nil {
		return true
	}

	return false
}

// SetSpans gets a reference to the given []StorageSpanRelationship and assigns it to the Spans field.
func (o *StorageDiskGroup) SetSpans(v []StorageSpanRelationship) {
	o.Spans = v
}

// GetStorageController returns the StorageController field value if set, zero value otherwise.
func (o *StorageDiskGroup) GetStorageController() StorageControllerRelationship {
	if o == nil || o.StorageController == nil {
		var ret StorageControllerRelationship
		return ret
	}
	return *o.StorageController
}

// GetStorageControllerOk returns a tuple with the StorageController field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageDiskGroup) GetStorageControllerOk() (*StorageControllerRelationship, bool) {
	if o == nil || o.StorageController == nil {
		return nil, false
	}
	return o.StorageController, true
}

// HasStorageController returns a boolean if a field has been set.
func (o *StorageDiskGroup) HasStorageController() bool {
	if o != nil && o.StorageController != nil {
		return true
	}

	return false
}

// SetStorageController gets a reference to the given StorageControllerRelationship and assigns it to the StorageController field.
func (o *StorageDiskGroup) SetStorageController(v StorageControllerRelationship) {
	o.StorageController = &v
}

// GetVirtualDrives returns the VirtualDrives field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StorageDiskGroup) GetVirtualDrives() []StorageVirtualDriveRelationship {
	if o == nil {
		var ret []StorageVirtualDriveRelationship
		return ret
	}
	return o.VirtualDrives
}

// GetVirtualDrivesOk returns a tuple with the VirtualDrives field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StorageDiskGroup) GetVirtualDrivesOk() (*[]StorageVirtualDriveRelationship, bool) {
	if o == nil || o.VirtualDrives == nil {
		return nil, false
	}
	return &o.VirtualDrives, true
}

// HasVirtualDrives returns a boolean if a field has been set.
func (o *StorageDiskGroup) HasVirtualDrives() bool {
	if o != nil && o.VirtualDrives != nil {
		return true
	}

	return false
}

// SetVirtualDrives gets a reference to the given []StorageVirtualDriveRelationship and assigns it to the VirtualDrives field.
func (o *StorageDiskGroup) SetVirtualDrives(v []StorageVirtualDriveRelationship) {
	o.VirtualDrives = v
}

func (o StorageDiskGroup) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedInventoryBase, errInventoryBase := json.Marshal(o.InventoryBase)
	if errInventoryBase != nil {
		return []byte{}, errInventoryBase
	}
	errInventoryBase = json.Unmarshal([]byte(serializedInventoryBase), &toSerialize)
	if errInventoryBase != nil {
		return []byte{}, errInventoryBase
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Name != nil {
		toSerialize["Name"] = o.Name
	}
	if o.RaidType != nil {
		toSerialize["RaidType"] = o.RaidType
	}
	if o.DedicatedHotSpares != nil {
		toSerialize["DedicatedHotSpares"] = o.DedicatedHotSpares
	}
	if o.RegisteredDevice != nil {
		toSerialize["RegisteredDevice"] = o.RegisteredDevice
	}
	if o.Spans != nil {
		toSerialize["Spans"] = o.Spans
	}
	if o.StorageController != nil {
		toSerialize["StorageController"] = o.StorageController
	}
	if o.VirtualDrives != nil {
		toSerialize["VirtualDrives"] = o.VirtualDrives
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *StorageDiskGroup) UnmarshalJSON(bytes []byte) (err error) {
	type StorageDiskGroupWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Name to identity this disk group in the controller.
		Name *string `json:"Name,omitempty"`
		// Raid level of the virtual drives in this diskgroup.
		RaidType *string `json:"RaidType,omitempty"`
		// An array of relationships to storagePhysicalDisk resources.
		DedicatedHotSpares []StoragePhysicalDiskRelationship    `json:"DedicatedHotSpares,omitempty"`
		RegisteredDevice   *AssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
		// An array of relationships to storageSpan resources.
		Spans             []StorageSpanRelationship      `json:"Spans,omitempty"`
		StorageController *StorageControllerRelationship `json:"StorageController,omitempty"`
		// An array of relationships to storageVirtualDrive resources.
		VirtualDrives []StorageVirtualDriveRelationship `json:"VirtualDrives,omitempty"`
	}

	varStorageDiskGroupWithoutEmbeddedStruct := StorageDiskGroupWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varStorageDiskGroupWithoutEmbeddedStruct)
	if err == nil {
		varStorageDiskGroup := _StorageDiskGroup{}
		varStorageDiskGroup.ClassId = varStorageDiskGroupWithoutEmbeddedStruct.ClassId
		varStorageDiskGroup.ObjectType = varStorageDiskGroupWithoutEmbeddedStruct.ObjectType
		varStorageDiskGroup.Name = varStorageDiskGroupWithoutEmbeddedStruct.Name
		varStorageDiskGroup.RaidType = varStorageDiskGroupWithoutEmbeddedStruct.RaidType
		varStorageDiskGroup.DedicatedHotSpares = varStorageDiskGroupWithoutEmbeddedStruct.DedicatedHotSpares
		varStorageDiskGroup.RegisteredDevice = varStorageDiskGroupWithoutEmbeddedStruct.RegisteredDevice
		varStorageDiskGroup.Spans = varStorageDiskGroupWithoutEmbeddedStruct.Spans
		varStorageDiskGroup.StorageController = varStorageDiskGroupWithoutEmbeddedStruct.StorageController
		varStorageDiskGroup.VirtualDrives = varStorageDiskGroupWithoutEmbeddedStruct.VirtualDrives
		*o = StorageDiskGroup(varStorageDiskGroup)
	} else {
		return err
	}

	varStorageDiskGroup := _StorageDiskGroup{}

	err = json.Unmarshal(bytes, &varStorageDiskGroup)
	if err == nil {
		o.InventoryBase = varStorageDiskGroup.InventoryBase
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Name")
		delete(additionalProperties, "RaidType")
		delete(additionalProperties, "DedicatedHotSpares")
		delete(additionalProperties, "RegisteredDevice")
		delete(additionalProperties, "Spans")
		delete(additionalProperties, "StorageController")
		delete(additionalProperties, "VirtualDrives")

		// remove fields from embedded structs
		reflectInventoryBase := reflect.ValueOf(o.InventoryBase)
		for i := 0; i < reflectInventoryBase.Type().NumField(); i++ {
			t := reflectInventoryBase.Type().Field(i)

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

type NullableStorageDiskGroup struct {
	value *StorageDiskGroup
	isSet bool
}

func (v NullableStorageDiskGroup) Get() *StorageDiskGroup {
	return v.value
}

func (v *NullableStorageDiskGroup) Set(val *StorageDiskGroup) {
	v.value = val
	v.isSet = true
}

func (v NullableStorageDiskGroup) IsSet() bool {
	return v.isSet
}

func (v *NullableStorageDiskGroup) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStorageDiskGroup(val *StorageDiskGroup) *NullableStorageDiskGroup {
	return &NullableStorageDiskGroup{value: val, isSet: true}
}

func (v NullableStorageDiskGroup) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStorageDiskGroup) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
