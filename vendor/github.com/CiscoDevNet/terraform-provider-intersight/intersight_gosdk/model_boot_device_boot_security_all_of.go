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

// BootDeviceBootSecurityAllOf Definition of the list of properties defined in 'boot.DeviceBootSecurity', excluding properties defined in parent classes.
type BootDeviceBootSecurityAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// The user desired BIOS secure boot as configured in the boot policy.
	SecureBoot           *string                              `json:"SecureBoot,omitempty"`
	ComputePhysical      *ComputePhysicalRelationship         `json:"ComputePhysical,omitempty"`
	InventoryDeviceInfo  *InventoryDeviceInfoRelationship     `json:"InventoryDeviceInfo,omitempty"`
	RegisteredDevice     *AssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _BootDeviceBootSecurityAllOf BootDeviceBootSecurityAllOf

// NewBootDeviceBootSecurityAllOf instantiates a new BootDeviceBootSecurityAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBootDeviceBootSecurityAllOf(classId string, objectType string) *BootDeviceBootSecurityAllOf {
	this := BootDeviceBootSecurityAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewBootDeviceBootSecurityAllOfWithDefaults instantiates a new BootDeviceBootSecurityAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBootDeviceBootSecurityAllOfWithDefaults() *BootDeviceBootSecurityAllOf {
	this := BootDeviceBootSecurityAllOf{}
	var classId string = "boot.DeviceBootSecurity"
	this.ClassId = classId
	var objectType string = "boot.DeviceBootSecurity"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *BootDeviceBootSecurityAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *BootDeviceBootSecurityAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *BootDeviceBootSecurityAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *BootDeviceBootSecurityAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *BootDeviceBootSecurityAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *BootDeviceBootSecurityAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetSecureBoot returns the SecureBoot field value if set, zero value otherwise.
func (o *BootDeviceBootSecurityAllOf) GetSecureBoot() string {
	if o == nil || o.SecureBoot == nil {
		var ret string
		return ret
	}
	return *o.SecureBoot
}

// GetSecureBootOk returns a tuple with the SecureBoot field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BootDeviceBootSecurityAllOf) GetSecureBootOk() (*string, bool) {
	if o == nil || o.SecureBoot == nil {
		return nil, false
	}
	return o.SecureBoot, true
}

// HasSecureBoot returns a boolean if a field has been set.
func (o *BootDeviceBootSecurityAllOf) HasSecureBoot() bool {
	if o != nil && o.SecureBoot != nil {
		return true
	}

	return false
}

// SetSecureBoot gets a reference to the given string and assigns it to the SecureBoot field.
func (o *BootDeviceBootSecurityAllOf) SetSecureBoot(v string) {
	o.SecureBoot = &v
}

// GetComputePhysical returns the ComputePhysical field value if set, zero value otherwise.
func (o *BootDeviceBootSecurityAllOf) GetComputePhysical() ComputePhysicalRelationship {
	if o == nil || o.ComputePhysical == nil {
		var ret ComputePhysicalRelationship
		return ret
	}
	return *o.ComputePhysical
}

// GetComputePhysicalOk returns a tuple with the ComputePhysical field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BootDeviceBootSecurityAllOf) GetComputePhysicalOk() (*ComputePhysicalRelationship, bool) {
	if o == nil || o.ComputePhysical == nil {
		return nil, false
	}
	return o.ComputePhysical, true
}

// HasComputePhysical returns a boolean if a field has been set.
func (o *BootDeviceBootSecurityAllOf) HasComputePhysical() bool {
	if o != nil && o.ComputePhysical != nil {
		return true
	}

	return false
}

// SetComputePhysical gets a reference to the given ComputePhysicalRelationship and assigns it to the ComputePhysical field.
func (o *BootDeviceBootSecurityAllOf) SetComputePhysical(v ComputePhysicalRelationship) {
	o.ComputePhysical = &v
}

// GetInventoryDeviceInfo returns the InventoryDeviceInfo field value if set, zero value otherwise.
func (o *BootDeviceBootSecurityAllOf) GetInventoryDeviceInfo() InventoryDeviceInfoRelationship {
	if o == nil || o.InventoryDeviceInfo == nil {
		var ret InventoryDeviceInfoRelationship
		return ret
	}
	return *o.InventoryDeviceInfo
}

// GetInventoryDeviceInfoOk returns a tuple with the InventoryDeviceInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BootDeviceBootSecurityAllOf) GetInventoryDeviceInfoOk() (*InventoryDeviceInfoRelationship, bool) {
	if o == nil || o.InventoryDeviceInfo == nil {
		return nil, false
	}
	return o.InventoryDeviceInfo, true
}

// HasInventoryDeviceInfo returns a boolean if a field has been set.
func (o *BootDeviceBootSecurityAllOf) HasInventoryDeviceInfo() bool {
	if o != nil && o.InventoryDeviceInfo != nil {
		return true
	}

	return false
}

// SetInventoryDeviceInfo gets a reference to the given InventoryDeviceInfoRelationship and assigns it to the InventoryDeviceInfo field.
func (o *BootDeviceBootSecurityAllOf) SetInventoryDeviceInfo(v InventoryDeviceInfoRelationship) {
	o.InventoryDeviceInfo = &v
}

// GetRegisteredDevice returns the RegisteredDevice field value if set, zero value otherwise.
func (o *BootDeviceBootSecurityAllOf) GetRegisteredDevice() AssetDeviceRegistrationRelationship {
	if o == nil || o.RegisteredDevice == nil {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.RegisteredDevice
}

// GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BootDeviceBootSecurityAllOf) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil || o.RegisteredDevice == nil {
		return nil, false
	}
	return o.RegisteredDevice, true
}

// HasRegisteredDevice returns a boolean if a field has been set.
func (o *BootDeviceBootSecurityAllOf) HasRegisteredDevice() bool {
	if o != nil && o.RegisteredDevice != nil {
		return true
	}

	return false
}

// SetRegisteredDevice gets a reference to the given AssetDeviceRegistrationRelationship and assigns it to the RegisteredDevice field.
func (o *BootDeviceBootSecurityAllOf) SetRegisteredDevice(v AssetDeviceRegistrationRelationship) {
	o.RegisteredDevice = &v
}

func (o BootDeviceBootSecurityAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.SecureBoot != nil {
		toSerialize["SecureBoot"] = o.SecureBoot
	}
	if o.ComputePhysical != nil {
		toSerialize["ComputePhysical"] = o.ComputePhysical
	}
	if o.InventoryDeviceInfo != nil {
		toSerialize["InventoryDeviceInfo"] = o.InventoryDeviceInfo
	}
	if o.RegisteredDevice != nil {
		toSerialize["RegisteredDevice"] = o.RegisteredDevice
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *BootDeviceBootSecurityAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varBootDeviceBootSecurityAllOf := _BootDeviceBootSecurityAllOf{}

	if err = json.Unmarshal(bytes, &varBootDeviceBootSecurityAllOf); err == nil {
		*o = BootDeviceBootSecurityAllOf(varBootDeviceBootSecurityAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "SecureBoot")
		delete(additionalProperties, "ComputePhysical")
		delete(additionalProperties, "InventoryDeviceInfo")
		delete(additionalProperties, "RegisteredDevice")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableBootDeviceBootSecurityAllOf struct {
	value *BootDeviceBootSecurityAllOf
	isSet bool
}

func (v NullableBootDeviceBootSecurityAllOf) Get() *BootDeviceBootSecurityAllOf {
	return v.value
}

func (v *NullableBootDeviceBootSecurityAllOf) Set(val *BootDeviceBootSecurityAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableBootDeviceBootSecurityAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableBootDeviceBootSecurityAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBootDeviceBootSecurityAllOf(val *BootDeviceBootSecurityAllOf) *NullableBootDeviceBootSecurityAllOf {
	return &NullableBootDeviceBootSecurityAllOf{value: val, isSet: true}
}

func (v NullableBootDeviceBootSecurityAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBootDeviceBootSecurityAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
