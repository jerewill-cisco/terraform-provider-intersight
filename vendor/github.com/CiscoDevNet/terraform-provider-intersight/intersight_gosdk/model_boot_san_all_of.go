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

// BootSanAllOf Definition of the list of properties defined in 'boot.San', excluding properties defined in parent classes.
type BootSanAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string                 `json:"ObjectType"`
	Bootloader NullableBootBootloader `json:"Bootloader,omitempty"`
	// The name of the underlying vHBA interface to be used by the SAN boot device.
	InterfaceName *string `json:"InterfaceName,omitempty"`
	// The Logical Unit Number (LUN) of the device.
	Lun *int64 `json:"Lun,omitempty"`
	// Slot ID of the device. Supported values are ( 1 - 255, \"MLOM\", \"L1\", \"L2\" ).
	Slot *string `json:"Slot,omitempty"`
	// The WWPN Address of the underlying fibre channel interface used by the SAN boot device. Value must be in hexadecimal format xx:xx:xx:xx:xx:xx:xx:xx.
	Wwpn                 *string `json:"Wwpn,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _BootSanAllOf BootSanAllOf

// NewBootSanAllOf instantiates a new BootSanAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBootSanAllOf(classId string, objectType string) *BootSanAllOf {
	this := BootSanAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	var lun int64 = 0
	this.Lun = &lun
	return &this
}

// NewBootSanAllOfWithDefaults instantiates a new BootSanAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBootSanAllOfWithDefaults() *BootSanAllOf {
	this := BootSanAllOf{}
	var classId string = "boot.San"
	this.ClassId = classId
	var objectType string = "boot.San"
	this.ObjectType = objectType
	var lun int64 = 0
	this.Lun = &lun
	return &this
}

// GetClassId returns the ClassId field value
func (o *BootSanAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *BootSanAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *BootSanAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *BootSanAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *BootSanAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *BootSanAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetBootloader returns the Bootloader field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BootSanAllOf) GetBootloader() BootBootloader {
	if o == nil || o.Bootloader.Get() == nil {
		var ret BootBootloader
		return ret
	}
	return *o.Bootloader.Get()
}

// GetBootloaderOk returns a tuple with the Bootloader field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BootSanAllOf) GetBootloaderOk() (*BootBootloader, bool) {
	if o == nil {
		return nil, false
	}
	return o.Bootloader.Get(), o.Bootloader.IsSet()
}

// HasBootloader returns a boolean if a field has been set.
func (o *BootSanAllOf) HasBootloader() bool {
	if o != nil && o.Bootloader.IsSet() {
		return true
	}

	return false
}

// SetBootloader gets a reference to the given NullableBootBootloader and assigns it to the Bootloader field.
func (o *BootSanAllOf) SetBootloader(v BootBootloader) {
	o.Bootloader.Set(&v)
}

// SetBootloaderNil sets the value for Bootloader to be an explicit nil
func (o *BootSanAllOf) SetBootloaderNil() {
	o.Bootloader.Set(nil)
}

// UnsetBootloader ensures that no value is present for Bootloader, not even an explicit nil
func (o *BootSanAllOf) UnsetBootloader() {
	o.Bootloader.Unset()
}

// GetInterfaceName returns the InterfaceName field value if set, zero value otherwise.
func (o *BootSanAllOf) GetInterfaceName() string {
	if o == nil || o.InterfaceName == nil {
		var ret string
		return ret
	}
	return *o.InterfaceName
}

// GetInterfaceNameOk returns a tuple with the InterfaceName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BootSanAllOf) GetInterfaceNameOk() (*string, bool) {
	if o == nil || o.InterfaceName == nil {
		return nil, false
	}
	return o.InterfaceName, true
}

// HasInterfaceName returns a boolean if a field has been set.
func (o *BootSanAllOf) HasInterfaceName() bool {
	if o != nil && o.InterfaceName != nil {
		return true
	}

	return false
}

// SetInterfaceName gets a reference to the given string and assigns it to the InterfaceName field.
func (o *BootSanAllOf) SetInterfaceName(v string) {
	o.InterfaceName = &v
}

// GetLun returns the Lun field value if set, zero value otherwise.
func (o *BootSanAllOf) GetLun() int64 {
	if o == nil || o.Lun == nil {
		var ret int64
		return ret
	}
	return *o.Lun
}

// GetLunOk returns a tuple with the Lun field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BootSanAllOf) GetLunOk() (*int64, bool) {
	if o == nil || o.Lun == nil {
		return nil, false
	}
	return o.Lun, true
}

// HasLun returns a boolean if a field has been set.
func (o *BootSanAllOf) HasLun() bool {
	if o != nil && o.Lun != nil {
		return true
	}

	return false
}

// SetLun gets a reference to the given int64 and assigns it to the Lun field.
func (o *BootSanAllOf) SetLun(v int64) {
	o.Lun = &v
}

// GetSlot returns the Slot field value if set, zero value otherwise.
func (o *BootSanAllOf) GetSlot() string {
	if o == nil || o.Slot == nil {
		var ret string
		return ret
	}
	return *o.Slot
}

// GetSlotOk returns a tuple with the Slot field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BootSanAllOf) GetSlotOk() (*string, bool) {
	if o == nil || o.Slot == nil {
		return nil, false
	}
	return o.Slot, true
}

// HasSlot returns a boolean if a field has been set.
func (o *BootSanAllOf) HasSlot() bool {
	if o != nil && o.Slot != nil {
		return true
	}

	return false
}

// SetSlot gets a reference to the given string and assigns it to the Slot field.
func (o *BootSanAllOf) SetSlot(v string) {
	o.Slot = &v
}

// GetWwpn returns the Wwpn field value if set, zero value otherwise.
func (o *BootSanAllOf) GetWwpn() string {
	if o == nil || o.Wwpn == nil {
		var ret string
		return ret
	}
	return *o.Wwpn
}

// GetWwpnOk returns a tuple with the Wwpn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BootSanAllOf) GetWwpnOk() (*string, bool) {
	if o == nil || o.Wwpn == nil {
		return nil, false
	}
	return o.Wwpn, true
}

// HasWwpn returns a boolean if a field has been set.
func (o *BootSanAllOf) HasWwpn() bool {
	if o != nil && o.Wwpn != nil {
		return true
	}

	return false
}

// SetWwpn gets a reference to the given string and assigns it to the Wwpn field.
func (o *BootSanAllOf) SetWwpn(v string) {
	o.Wwpn = &v
}

func (o BootSanAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Bootloader.IsSet() {
		toSerialize["Bootloader"] = o.Bootloader.Get()
	}
	if o.InterfaceName != nil {
		toSerialize["InterfaceName"] = o.InterfaceName
	}
	if o.Lun != nil {
		toSerialize["Lun"] = o.Lun
	}
	if o.Slot != nil {
		toSerialize["Slot"] = o.Slot
	}
	if o.Wwpn != nil {
		toSerialize["Wwpn"] = o.Wwpn
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *BootSanAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varBootSanAllOf := _BootSanAllOf{}

	if err = json.Unmarshal(bytes, &varBootSanAllOf); err == nil {
		*o = BootSanAllOf(varBootSanAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Bootloader")
		delete(additionalProperties, "InterfaceName")
		delete(additionalProperties, "Lun")
		delete(additionalProperties, "Slot")
		delete(additionalProperties, "Wwpn")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableBootSanAllOf struct {
	value *BootSanAllOf
	isSet bool
}

func (v NullableBootSanAllOf) Get() *BootSanAllOf {
	return v.value
}

func (v *NullableBootSanAllOf) Set(val *BootSanAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableBootSanAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableBootSanAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBootSanAllOf(val *BootSanAllOf) *NullableBootSanAllOf {
	return &NullableBootSanAllOf{value: val, isSet: true}
}

func (v NullableBootSanAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBootSanAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
