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

// ComputeStoragePhysicalDriveOperationAllOf Definition of the list of properties defined in 'compute.StoragePhysicalDriveOperation', excluding properties defined in parent classes.
type ComputeStoragePhysicalDriveOperationAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Administrative actions that can be performed on the Storage Physical Drives. * `None` - No action on the selected Storage Physical Drives. * `SetJbod` - Set Jbod action state on the selected Storage Physical Drives. * `SetUnconfiguredGood` - Set Unconfigured Good action state on the selected Storage Physical Drives.
	AdminAction *string `json:"AdminAction,omitempty"`
	// Storage Controller Id of the storage Physical Drives of the server.
	ControllerId         *string                       `json:"ControllerId,omitempty"`
	PhysicalDrives       []ComputeStoragePhysicalDrive `json:"PhysicalDrives,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ComputeStoragePhysicalDriveOperationAllOf ComputeStoragePhysicalDriveOperationAllOf

// NewComputeStoragePhysicalDriveOperationAllOf instantiates a new ComputeStoragePhysicalDriveOperationAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewComputeStoragePhysicalDriveOperationAllOf(classId string, objectType string) *ComputeStoragePhysicalDriveOperationAllOf {
	this := ComputeStoragePhysicalDriveOperationAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	var adminAction string = "None"
	this.AdminAction = &adminAction
	return &this
}

// NewComputeStoragePhysicalDriveOperationAllOfWithDefaults instantiates a new ComputeStoragePhysicalDriveOperationAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewComputeStoragePhysicalDriveOperationAllOfWithDefaults() *ComputeStoragePhysicalDriveOperationAllOf {
	this := ComputeStoragePhysicalDriveOperationAllOf{}
	var classId string = "compute.StoragePhysicalDriveOperation"
	this.ClassId = classId
	var objectType string = "compute.StoragePhysicalDriveOperation"
	this.ObjectType = objectType
	var adminAction string = "None"
	this.AdminAction = &adminAction
	return &this
}

// GetClassId returns the ClassId field value
func (o *ComputeStoragePhysicalDriveOperationAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *ComputeStoragePhysicalDriveOperationAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *ComputeStoragePhysicalDriveOperationAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *ComputeStoragePhysicalDriveOperationAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *ComputeStoragePhysicalDriveOperationAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *ComputeStoragePhysicalDriveOperationAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetAdminAction returns the AdminAction field value if set, zero value otherwise.
func (o *ComputeStoragePhysicalDriveOperationAllOf) GetAdminAction() string {
	if o == nil || o.AdminAction == nil {
		var ret string
		return ret
	}
	return *o.AdminAction
}

// GetAdminActionOk returns a tuple with the AdminAction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputeStoragePhysicalDriveOperationAllOf) GetAdminActionOk() (*string, bool) {
	if o == nil || o.AdminAction == nil {
		return nil, false
	}
	return o.AdminAction, true
}

// HasAdminAction returns a boolean if a field has been set.
func (o *ComputeStoragePhysicalDriveOperationAllOf) HasAdminAction() bool {
	if o != nil && o.AdminAction != nil {
		return true
	}

	return false
}

// SetAdminAction gets a reference to the given string and assigns it to the AdminAction field.
func (o *ComputeStoragePhysicalDriveOperationAllOf) SetAdminAction(v string) {
	o.AdminAction = &v
}

// GetControllerId returns the ControllerId field value if set, zero value otherwise.
func (o *ComputeStoragePhysicalDriveOperationAllOf) GetControllerId() string {
	if o == nil || o.ControllerId == nil {
		var ret string
		return ret
	}
	return *o.ControllerId
}

// GetControllerIdOk returns a tuple with the ControllerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputeStoragePhysicalDriveOperationAllOf) GetControllerIdOk() (*string, bool) {
	if o == nil || o.ControllerId == nil {
		return nil, false
	}
	return o.ControllerId, true
}

// HasControllerId returns a boolean if a field has been set.
func (o *ComputeStoragePhysicalDriveOperationAllOf) HasControllerId() bool {
	if o != nil && o.ControllerId != nil {
		return true
	}

	return false
}

// SetControllerId gets a reference to the given string and assigns it to the ControllerId field.
func (o *ComputeStoragePhysicalDriveOperationAllOf) SetControllerId(v string) {
	o.ControllerId = &v
}

// GetPhysicalDrives returns the PhysicalDrives field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ComputeStoragePhysicalDriveOperationAllOf) GetPhysicalDrives() []ComputeStoragePhysicalDrive {
	if o == nil {
		var ret []ComputeStoragePhysicalDrive
		return ret
	}
	return o.PhysicalDrives
}

// GetPhysicalDrivesOk returns a tuple with the PhysicalDrives field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ComputeStoragePhysicalDriveOperationAllOf) GetPhysicalDrivesOk() (*[]ComputeStoragePhysicalDrive, bool) {
	if o == nil || o.PhysicalDrives == nil {
		return nil, false
	}
	return &o.PhysicalDrives, true
}

// HasPhysicalDrives returns a boolean if a field has been set.
func (o *ComputeStoragePhysicalDriveOperationAllOf) HasPhysicalDrives() bool {
	if o != nil && o.PhysicalDrives != nil {
		return true
	}

	return false
}

// SetPhysicalDrives gets a reference to the given []ComputeStoragePhysicalDrive and assigns it to the PhysicalDrives field.
func (o *ComputeStoragePhysicalDriveOperationAllOf) SetPhysicalDrives(v []ComputeStoragePhysicalDrive) {
	o.PhysicalDrives = v
}

func (o ComputeStoragePhysicalDriveOperationAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.AdminAction != nil {
		toSerialize["AdminAction"] = o.AdminAction
	}
	if o.ControllerId != nil {
		toSerialize["ControllerId"] = o.ControllerId
	}
	if o.PhysicalDrives != nil {
		toSerialize["PhysicalDrives"] = o.PhysicalDrives
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ComputeStoragePhysicalDriveOperationAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varComputeStoragePhysicalDriveOperationAllOf := _ComputeStoragePhysicalDriveOperationAllOf{}

	if err = json.Unmarshal(bytes, &varComputeStoragePhysicalDriveOperationAllOf); err == nil {
		*o = ComputeStoragePhysicalDriveOperationAllOf(varComputeStoragePhysicalDriveOperationAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "AdminAction")
		delete(additionalProperties, "ControllerId")
		delete(additionalProperties, "PhysicalDrives")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableComputeStoragePhysicalDriveOperationAllOf struct {
	value *ComputeStoragePhysicalDriveOperationAllOf
	isSet bool
}

func (v NullableComputeStoragePhysicalDriveOperationAllOf) Get() *ComputeStoragePhysicalDriveOperationAllOf {
	return v.value
}

func (v *NullableComputeStoragePhysicalDriveOperationAllOf) Set(val *ComputeStoragePhysicalDriveOperationAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableComputeStoragePhysicalDriveOperationAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableComputeStoragePhysicalDriveOperationAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableComputeStoragePhysicalDriveOperationAllOf(val *ComputeStoragePhysicalDriveOperationAllOf) *NullableComputeStoragePhysicalDriveOperationAllOf {
	return &NullableComputeStoragePhysicalDriveOperationAllOf{value: val, isSet: true}
}

func (v NullableComputeStoragePhysicalDriveOperationAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableComputeStoragePhysicalDriveOperationAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
