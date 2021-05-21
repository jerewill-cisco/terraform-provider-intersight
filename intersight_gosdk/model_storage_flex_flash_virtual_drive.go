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

// StorageFlexFlashVirtualDrive Virtual Drive repersenting a SD Card.
type StorageFlexFlashVirtualDrive struct {
	EquipmentBase
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// The drive scope of the flex flash virtual drive.
	DriveScope *string `json:"DriveScope,omitempty"`
	// Status of virtual drive on the flex controller.
	DriveStatus *string `json:"DriveStatus,omitempty"`
	// The partition Id of the flex flash virtual Drive.
	PartitionId *string `json:"PartitionId,omitempty"`
	// The resident image on the flex flash virtual Drive.
	ResidentImage *string `json:"ResidentImage,omitempty"`
	// Size of virtual drive on the flex controller.
	Size *string `json:"Size,omitempty"`
	// Virtual drive on the flex flash controller.
	VirtualDrive               *string                                 `json:"VirtualDrive,omitempty"`
	InventoryDeviceInfo        *InventoryDeviceInfoRelationship        `json:"InventoryDeviceInfo,omitempty"`
	RegisteredDevice           *AssetDeviceRegistrationRelationship    `json:"RegisteredDevice,omitempty"`
	StorageFlexFlashController *StorageFlexFlashControllerRelationship `json:"StorageFlexFlashController,omitempty"`
	AdditionalProperties       map[string]interface{}
}

type _StorageFlexFlashVirtualDrive StorageFlexFlashVirtualDrive

// NewStorageFlexFlashVirtualDrive instantiates a new StorageFlexFlashVirtualDrive object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStorageFlexFlashVirtualDrive(classId string, objectType string) *StorageFlexFlashVirtualDrive {
	this := StorageFlexFlashVirtualDrive{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewStorageFlexFlashVirtualDriveWithDefaults instantiates a new StorageFlexFlashVirtualDrive object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStorageFlexFlashVirtualDriveWithDefaults() *StorageFlexFlashVirtualDrive {
	this := StorageFlexFlashVirtualDrive{}
	var classId string = "storage.FlexFlashVirtualDrive"
	this.ClassId = classId
	var objectType string = "storage.FlexFlashVirtualDrive"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *StorageFlexFlashVirtualDrive) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *StorageFlexFlashVirtualDrive) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *StorageFlexFlashVirtualDrive) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *StorageFlexFlashVirtualDrive) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *StorageFlexFlashVirtualDrive) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *StorageFlexFlashVirtualDrive) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDriveScope returns the DriveScope field value if set, zero value otherwise.
func (o *StorageFlexFlashVirtualDrive) GetDriveScope() string {
	if o == nil || o.DriveScope == nil {
		var ret string
		return ret
	}
	return *o.DriveScope
}

// GetDriveScopeOk returns a tuple with the DriveScope field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageFlexFlashVirtualDrive) GetDriveScopeOk() (*string, bool) {
	if o == nil || o.DriveScope == nil {
		return nil, false
	}
	return o.DriveScope, true
}

// HasDriveScope returns a boolean if a field has been set.
func (o *StorageFlexFlashVirtualDrive) HasDriveScope() bool {
	if o != nil && o.DriveScope != nil {
		return true
	}

	return false
}

// SetDriveScope gets a reference to the given string and assigns it to the DriveScope field.
func (o *StorageFlexFlashVirtualDrive) SetDriveScope(v string) {
	o.DriveScope = &v
}

// GetDriveStatus returns the DriveStatus field value if set, zero value otherwise.
func (o *StorageFlexFlashVirtualDrive) GetDriveStatus() string {
	if o == nil || o.DriveStatus == nil {
		var ret string
		return ret
	}
	return *o.DriveStatus
}

// GetDriveStatusOk returns a tuple with the DriveStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageFlexFlashVirtualDrive) GetDriveStatusOk() (*string, bool) {
	if o == nil || o.DriveStatus == nil {
		return nil, false
	}
	return o.DriveStatus, true
}

// HasDriveStatus returns a boolean if a field has been set.
func (o *StorageFlexFlashVirtualDrive) HasDriveStatus() bool {
	if o != nil && o.DriveStatus != nil {
		return true
	}

	return false
}

// SetDriveStatus gets a reference to the given string and assigns it to the DriveStatus field.
func (o *StorageFlexFlashVirtualDrive) SetDriveStatus(v string) {
	o.DriveStatus = &v
}

// GetPartitionId returns the PartitionId field value if set, zero value otherwise.
func (o *StorageFlexFlashVirtualDrive) GetPartitionId() string {
	if o == nil || o.PartitionId == nil {
		var ret string
		return ret
	}
	return *o.PartitionId
}

// GetPartitionIdOk returns a tuple with the PartitionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageFlexFlashVirtualDrive) GetPartitionIdOk() (*string, bool) {
	if o == nil || o.PartitionId == nil {
		return nil, false
	}
	return o.PartitionId, true
}

// HasPartitionId returns a boolean if a field has been set.
func (o *StorageFlexFlashVirtualDrive) HasPartitionId() bool {
	if o != nil && o.PartitionId != nil {
		return true
	}

	return false
}

// SetPartitionId gets a reference to the given string and assigns it to the PartitionId field.
func (o *StorageFlexFlashVirtualDrive) SetPartitionId(v string) {
	o.PartitionId = &v
}

// GetResidentImage returns the ResidentImage field value if set, zero value otherwise.
func (o *StorageFlexFlashVirtualDrive) GetResidentImage() string {
	if o == nil || o.ResidentImage == nil {
		var ret string
		return ret
	}
	return *o.ResidentImage
}

// GetResidentImageOk returns a tuple with the ResidentImage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageFlexFlashVirtualDrive) GetResidentImageOk() (*string, bool) {
	if o == nil || o.ResidentImage == nil {
		return nil, false
	}
	return o.ResidentImage, true
}

// HasResidentImage returns a boolean if a field has been set.
func (o *StorageFlexFlashVirtualDrive) HasResidentImage() bool {
	if o != nil && o.ResidentImage != nil {
		return true
	}

	return false
}

// SetResidentImage gets a reference to the given string and assigns it to the ResidentImage field.
func (o *StorageFlexFlashVirtualDrive) SetResidentImage(v string) {
	o.ResidentImage = &v
}

// GetSize returns the Size field value if set, zero value otherwise.
func (o *StorageFlexFlashVirtualDrive) GetSize() string {
	if o == nil || o.Size == nil {
		var ret string
		return ret
	}
	return *o.Size
}

// GetSizeOk returns a tuple with the Size field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageFlexFlashVirtualDrive) GetSizeOk() (*string, bool) {
	if o == nil || o.Size == nil {
		return nil, false
	}
	return o.Size, true
}

// HasSize returns a boolean if a field has been set.
func (o *StorageFlexFlashVirtualDrive) HasSize() bool {
	if o != nil && o.Size != nil {
		return true
	}

	return false
}

// SetSize gets a reference to the given string and assigns it to the Size field.
func (o *StorageFlexFlashVirtualDrive) SetSize(v string) {
	o.Size = &v
}

// GetVirtualDrive returns the VirtualDrive field value if set, zero value otherwise.
func (o *StorageFlexFlashVirtualDrive) GetVirtualDrive() string {
	if o == nil || o.VirtualDrive == nil {
		var ret string
		return ret
	}
	return *o.VirtualDrive
}

// GetVirtualDriveOk returns a tuple with the VirtualDrive field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageFlexFlashVirtualDrive) GetVirtualDriveOk() (*string, bool) {
	if o == nil || o.VirtualDrive == nil {
		return nil, false
	}
	return o.VirtualDrive, true
}

// HasVirtualDrive returns a boolean if a field has been set.
func (o *StorageFlexFlashVirtualDrive) HasVirtualDrive() bool {
	if o != nil && o.VirtualDrive != nil {
		return true
	}

	return false
}

// SetVirtualDrive gets a reference to the given string and assigns it to the VirtualDrive field.
func (o *StorageFlexFlashVirtualDrive) SetVirtualDrive(v string) {
	o.VirtualDrive = &v
}

// GetInventoryDeviceInfo returns the InventoryDeviceInfo field value if set, zero value otherwise.
func (o *StorageFlexFlashVirtualDrive) GetInventoryDeviceInfo() InventoryDeviceInfoRelationship {
	if o == nil || o.InventoryDeviceInfo == nil {
		var ret InventoryDeviceInfoRelationship
		return ret
	}
	return *o.InventoryDeviceInfo
}

// GetInventoryDeviceInfoOk returns a tuple with the InventoryDeviceInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageFlexFlashVirtualDrive) GetInventoryDeviceInfoOk() (*InventoryDeviceInfoRelationship, bool) {
	if o == nil || o.InventoryDeviceInfo == nil {
		return nil, false
	}
	return o.InventoryDeviceInfo, true
}

// HasInventoryDeviceInfo returns a boolean if a field has been set.
func (o *StorageFlexFlashVirtualDrive) HasInventoryDeviceInfo() bool {
	if o != nil && o.InventoryDeviceInfo != nil {
		return true
	}

	return false
}

// SetInventoryDeviceInfo gets a reference to the given InventoryDeviceInfoRelationship and assigns it to the InventoryDeviceInfo field.
func (o *StorageFlexFlashVirtualDrive) SetInventoryDeviceInfo(v InventoryDeviceInfoRelationship) {
	o.InventoryDeviceInfo = &v
}

// GetRegisteredDevice returns the RegisteredDevice field value if set, zero value otherwise.
func (o *StorageFlexFlashVirtualDrive) GetRegisteredDevice() AssetDeviceRegistrationRelationship {
	if o == nil || o.RegisteredDevice == nil {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.RegisteredDevice
}

// GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageFlexFlashVirtualDrive) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil || o.RegisteredDevice == nil {
		return nil, false
	}
	return o.RegisteredDevice, true
}

// HasRegisteredDevice returns a boolean if a field has been set.
func (o *StorageFlexFlashVirtualDrive) HasRegisteredDevice() bool {
	if o != nil && o.RegisteredDevice != nil {
		return true
	}

	return false
}

// SetRegisteredDevice gets a reference to the given AssetDeviceRegistrationRelationship and assigns it to the RegisteredDevice field.
func (o *StorageFlexFlashVirtualDrive) SetRegisteredDevice(v AssetDeviceRegistrationRelationship) {
	o.RegisteredDevice = &v
}

// GetStorageFlexFlashController returns the StorageFlexFlashController field value if set, zero value otherwise.
func (o *StorageFlexFlashVirtualDrive) GetStorageFlexFlashController() StorageFlexFlashControllerRelationship {
	if o == nil || o.StorageFlexFlashController == nil {
		var ret StorageFlexFlashControllerRelationship
		return ret
	}
	return *o.StorageFlexFlashController
}

// GetStorageFlexFlashControllerOk returns a tuple with the StorageFlexFlashController field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageFlexFlashVirtualDrive) GetStorageFlexFlashControllerOk() (*StorageFlexFlashControllerRelationship, bool) {
	if o == nil || o.StorageFlexFlashController == nil {
		return nil, false
	}
	return o.StorageFlexFlashController, true
}

// HasStorageFlexFlashController returns a boolean if a field has been set.
func (o *StorageFlexFlashVirtualDrive) HasStorageFlexFlashController() bool {
	if o != nil && o.StorageFlexFlashController != nil {
		return true
	}

	return false
}

// SetStorageFlexFlashController gets a reference to the given StorageFlexFlashControllerRelationship and assigns it to the StorageFlexFlashController field.
func (o *StorageFlexFlashVirtualDrive) SetStorageFlexFlashController(v StorageFlexFlashControllerRelationship) {
	o.StorageFlexFlashController = &v
}

func (o StorageFlexFlashVirtualDrive) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedEquipmentBase, errEquipmentBase := json.Marshal(o.EquipmentBase)
	if errEquipmentBase != nil {
		return []byte{}, errEquipmentBase
	}
	errEquipmentBase = json.Unmarshal([]byte(serializedEquipmentBase), &toSerialize)
	if errEquipmentBase != nil {
		return []byte{}, errEquipmentBase
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.DriveScope != nil {
		toSerialize["DriveScope"] = o.DriveScope
	}
	if o.DriveStatus != nil {
		toSerialize["DriveStatus"] = o.DriveStatus
	}
	if o.PartitionId != nil {
		toSerialize["PartitionId"] = o.PartitionId
	}
	if o.ResidentImage != nil {
		toSerialize["ResidentImage"] = o.ResidentImage
	}
	if o.Size != nil {
		toSerialize["Size"] = o.Size
	}
	if o.VirtualDrive != nil {
		toSerialize["VirtualDrive"] = o.VirtualDrive
	}
	if o.InventoryDeviceInfo != nil {
		toSerialize["InventoryDeviceInfo"] = o.InventoryDeviceInfo
	}
	if o.RegisteredDevice != nil {
		toSerialize["RegisteredDevice"] = o.RegisteredDevice
	}
	if o.StorageFlexFlashController != nil {
		toSerialize["StorageFlexFlashController"] = o.StorageFlexFlashController
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *StorageFlexFlashVirtualDrive) UnmarshalJSON(bytes []byte) (err error) {
	type StorageFlexFlashVirtualDriveWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// The drive scope of the flex flash virtual drive.
		DriveScope *string `json:"DriveScope,omitempty"`
		// Status of virtual drive on the flex controller.
		DriveStatus *string `json:"DriveStatus,omitempty"`
		// The partition Id of the flex flash virtual Drive.
		PartitionId *string `json:"PartitionId,omitempty"`
		// The resident image on the flex flash virtual Drive.
		ResidentImage *string `json:"ResidentImage,omitempty"`
		// Size of virtual drive on the flex controller.
		Size *string `json:"Size,omitempty"`
		// Virtual drive on the flex flash controller.
		VirtualDrive               *string                                 `json:"VirtualDrive,omitempty"`
		InventoryDeviceInfo        *InventoryDeviceInfoRelationship        `json:"InventoryDeviceInfo,omitempty"`
		RegisteredDevice           *AssetDeviceRegistrationRelationship    `json:"RegisteredDevice,omitempty"`
		StorageFlexFlashController *StorageFlexFlashControllerRelationship `json:"StorageFlexFlashController,omitempty"`
	}

	varStorageFlexFlashVirtualDriveWithoutEmbeddedStruct := StorageFlexFlashVirtualDriveWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varStorageFlexFlashVirtualDriveWithoutEmbeddedStruct)
	if err == nil {
		varStorageFlexFlashVirtualDrive := _StorageFlexFlashVirtualDrive{}
		varStorageFlexFlashVirtualDrive.ClassId = varStorageFlexFlashVirtualDriveWithoutEmbeddedStruct.ClassId
		varStorageFlexFlashVirtualDrive.ObjectType = varStorageFlexFlashVirtualDriveWithoutEmbeddedStruct.ObjectType
		varStorageFlexFlashVirtualDrive.DriveScope = varStorageFlexFlashVirtualDriveWithoutEmbeddedStruct.DriveScope
		varStorageFlexFlashVirtualDrive.DriveStatus = varStorageFlexFlashVirtualDriveWithoutEmbeddedStruct.DriveStatus
		varStorageFlexFlashVirtualDrive.PartitionId = varStorageFlexFlashVirtualDriveWithoutEmbeddedStruct.PartitionId
		varStorageFlexFlashVirtualDrive.ResidentImage = varStorageFlexFlashVirtualDriveWithoutEmbeddedStruct.ResidentImage
		varStorageFlexFlashVirtualDrive.Size = varStorageFlexFlashVirtualDriveWithoutEmbeddedStruct.Size
		varStorageFlexFlashVirtualDrive.VirtualDrive = varStorageFlexFlashVirtualDriveWithoutEmbeddedStruct.VirtualDrive
		varStorageFlexFlashVirtualDrive.InventoryDeviceInfo = varStorageFlexFlashVirtualDriveWithoutEmbeddedStruct.InventoryDeviceInfo
		varStorageFlexFlashVirtualDrive.RegisteredDevice = varStorageFlexFlashVirtualDriveWithoutEmbeddedStruct.RegisteredDevice
		varStorageFlexFlashVirtualDrive.StorageFlexFlashController = varStorageFlexFlashVirtualDriveWithoutEmbeddedStruct.StorageFlexFlashController
		*o = StorageFlexFlashVirtualDrive(varStorageFlexFlashVirtualDrive)
	} else {
		return err
	}

	varStorageFlexFlashVirtualDrive := _StorageFlexFlashVirtualDrive{}

	err = json.Unmarshal(bytes, &varStorageFlexFlashVirtualDrive)
	if err == nil {
		o.EquipmentBase = varStorageFlexFlashVirtualDrive.EquipmentBase
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "DriveScope")
		delete(additionalProperties, "DriveStatus")
		delete(additionalProperties, "PartitionId")
		delete(additionalProperties, "ResidentImage")
		delete(additionalProperties, "Size")
		delete(additionalProperties, "VirtualDrive")
		delete(additionalProperties, "InventoryDeviceInfo")
		delete(additionalProperties, "RegisteredDevice")
		delete(additionalProperties, "StorageFlexFlashController")

		// remove fields from embedded structs
		reflectEquipmentBase := reflect.ValueOf(o.EquipmentBase)
		for i := 0; i < reflectEquipmentBase.Type().NumField(); i++ {
			t := reflectEquipmentBase.Type().Field(i)

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

type NullableStorageFlexFlashVirtualDrive struct {
	value *StorageFlexFlashVirtualDrive
	isSet bool
}

func (v NullableStorageFlexFlashVirtualDrive) Get() *StorageFlexFlashVirtualDrive {
	return v.value
}

func (v *NullableStorageFlexFlashVirtualDrive) Set(val *StorageFlexFlashVirtualDrive) {
	v.value = val
	v.isSet = true
}

func (v NullableStorageFlexFlashVirtualDrive) IsSet() bool {
	return v.isSet
}

func (v *NullableStorageFlexFlashVirtualDrive) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStorageFlexFlashVirtualDrive(val *StorageFlexFlashVirtualDrive) *NullableStorageFlexFlashVirtualDrive {
	return &NullableStorageFlexFlashVirtualDrive{value: val, isSet: true}
}

func (v NullableStorageFlexFlashVirtualDrive) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStorageFlexFlashVirtualDrive) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
