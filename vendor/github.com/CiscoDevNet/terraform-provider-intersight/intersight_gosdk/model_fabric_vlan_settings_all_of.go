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

// FabricVlanSettingsAllOf Definition of the list of properties defined in 'fabric.VlanSettings', excluding properties defined in parent classes.
type FabricVlanSettingsAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Allowed VLAN IDs of the virtual interface. A list of comma seperated VLAN ids and/or VLAN id ranges.
	AllowedVlans *string `json:"AllowedVlans,omitempty"`
	// Native VLAN ID of the virtual interface or the corresponding vethernet on the peer Fabric Interconnect to which the virtual interface is connected. If the native VLAN is not a part of the allowed VLANs, it will automatically be added to the list of allowed VLANs.
	NativeVlan           *int64 `json:"NativeVlan,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _FabricVlanSettingsAllOf FabricVlanSettingsAllOf

// NewFabricVlanSettingsAllOf instantiates a new FabricVlanSettingsAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFabricVlanSettingsAllOf(classId string, objectType string) *FabricVlanSettingsAllOf {
	this := FabricVlanSettingsAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	var nativeVlan int64 = 1
	this.NativeVlan = &nativeVlan
	return &this
}

// NewFabricVlanSettingsAllOfWithDefaults instantiates a new FabricVlanSettingsAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFabricVlanSettingsAllOfWithDefaults() *FabricVlanSettingsAllOf {
	this := FabricVlanSettingsAllOf{}
	var classId string = "fabric.VlanSettings"
	this.ClassId = classId
	var objectType string = "fabric.VlanSettings"
	this.ObjectType = objectType
	var nativeVlan int64 = 1
	this.NativeVlan = &nativeVlan
	return &this
}

// GetClassId returns the ClassId field value
func (o *FabricVlanSettingsAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *FabricVlanSettingsAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *FabricVlanSettingsAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *FabricVlanSettingsAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *FabricVlanSettingsAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *FabricVlanSettingsAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetAllowedVlans returns the AllowedVlans field value if set, zero value otherwise.
func (o *FabricVlanSettingsAllOf) GetAllowedVlans() string {
	if o == nil || o.AllowedVlans == nil {
		var ret string
		return ret
	}
	return *o.AllowedVlans
}

// GetAllowedVlansOk returns a tuple with the AllowedVlans field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FabricVlanSettingsAllOf) GetAllowedVlansOk() (*string, bool) {
	if o == nil || o.AllowedVlans == nil {
		return nil, false
	}
	return o.AllowedVlans, true
}

// HasAllowedVlans returns a boolean if a field has been set.
func (o *FabricVlanSettingsAllOf) HasAllowedVlans() bool {
	if o != nil && o.AllowedVlans != nil {
		return true
	}

	return false
}

// SetAllowedVlans gets a reference to the given string and assigns it to the AllowedVlans field.
func (o *FabricVlanSettingsAllOf) SetAllowedVlans(v string) {
	o.AllowedVlans = &v
}

// GetNativeVlan returns the NativeVlan field value if set, zero value otherwise.
func (o *FabricVlanSettingsAllOf) GetNativeVlan() int64 {
	if o == nil || o.NativeVlan == nil {
		var ret int64
		return ret
	}
	return *o.NativeVlan
}

// GetNativeVlanOk returns a tuple with the NativeVlan field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FabricVlanSettingsAllOf) GetNativeVlanOk() (*int64, bool) {
	if o == nil || o.NativeVlan == nil {
		return nil, false
	}
	return o.NativeVlan, true
}

// HasNativeVlan returns a boolean if a field has been set.
func (o *FabricVlanSettingsAllOf) HasNativeVlan() bool {
	if o != nil && o.NativeVlan != nil {
		return true
	}

	return false
}

// SetNativeVlan gets a reference to the given int64 and assigns it to the NativeVlan field.
func (o *FabricVlanSettingsAllOf) SetNativeVlan(v int64) {
	o.NativeVlan = &v
}

func (o FabricVlanSettingsAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.AllowedVlans != nil {
		toSerialize["AllowedVlans"] = o.AllowedVlans
	}
	if o.NativeVlan != nil {
		toSerialize["NativeVlan"] = o.NativeVlan
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *FabricVlanSettingsAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varFabricVlanSettingsAllOf := _FabricVlanSettingsAllOf{}

	if err = json.Unmarshal(bytes, &varFabricVlanSettingsAllOf); err == nil {
		*o = FabricVlanSettingsAllOf(varFabricVlanSettingsAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "AllowedVlans")
		delete(additionalProperties, "NativeVlan")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableFabricVlanSettingsAllOf struct {
	value *FabricVlanSettingsAllOf
	isSet bool
}

func (v NullableFabricVlanSettingsAllOf) Get() *FabricVlanSettingsAllOf {
	return v.value
}

func (v *NullableFabricVlanSettingsAllOf) Set(val *FabricVlanSettingsAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableFabricVlanSettingsAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableFabricVlanSettingsAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFabricVlanSettingsAllOf(val *FabricVlanSettingsAllOf) *NullableFabricVlanSettingsAllOf {
	return &NullableFabricVlanSettingsAllOf{value: val, isSet: true}
}

func (v NullableFabricVlanSettingsAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFabricVlanSettingsAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
