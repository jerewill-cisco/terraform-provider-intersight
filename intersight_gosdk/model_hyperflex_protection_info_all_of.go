/*
Cisco Intersight

Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document. This document was created on 2021-08-10T21:48:06Z.

API version: 1.0.9-4430
Contact: intersight@cisco.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package intersight

import (
	"encoding/json"
)

// HyperflexProtectionInfoAllOf Definition of the list of properties defined in 'hyperflex.ProtectionInfo', excluding properties defined in parent classes.
type HyperflexProtectionInfoAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType                     string                                  `json:"ObjectType"`
	VmCurrentProtectionInfo        NullableHyperflexSnapshotInfoBrief      `json:"VmCurrentProtectionInfo,omitempty"`
	VmLastSuccessfulProtectionInfo NullableHyperflexSnapshotInfoBrief      `json:"VmLastSuccessfulProtectionInfo,omitempty"`
	VmSpaceUsage                   NullableHyperflexVmProtectionSpaceUsage `json:"VmSpaceUsage,omitempty"`
	AdditionalProperties           map[string]interface{}
}

type _HyperflexProtectionInfoAllOf HyperflexProtectionInfoAllOf

// NewHyperflexProtectionInfoAllOf instantiates a new HyperflexProtectionInfoAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHyperflexProtectionInfoAllOf(classId string, objectType string) *HyperflexProtectionInfoAllOf {
	this := HyperflexProtectionInfoAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewHyperflexProtectionInfoAllOfWithDefaults instantiates a new HyperflexProtectionInfoAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHyperflexProtectionInfoAllOfWithDefaults() *HyperflexProtectionInfoAllOf {
	this := HyperflexProtectionInfoAllOf{}
	var classId string = "hyperflex.ProtectionInfo"
	this.ClassId = classId
	var objectType string = "hyperflex.ProtectionInfo"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *HyperflexProtectionInfoAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *HyperflexProtectionInfoAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *HyperflexProtectionInfoAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *HyperflexProtectionInfoAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *HyperflexProtectionInfoAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *HyperflexProtectionInfoAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetVmCurrentProtectionInfo returns the VmCurrentProtectionInfo field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HyperflexProtectionInfoAllOf) GetVmCurrentProtectionInfo() HyperflexSnapshotInfoBrief {
	if o == nil || o.VmCurrentProtectionInfo.Get() == nil {
		var ret HyperflexSnapshotInfoBrief
		return ret
	}
	return *o.VmCurrentProtectionInfo.Get()
}

// GetVmCurrentProtectionInfoOk returns a tuple with the VmCurrentProtectionInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HyperflexProtectionInfoAllOf) GetVmCurrentProtectionInfoOk() (*HyperflexSnapshotInfoBrief, bool) {
	if o == nil {
		return nil, false
	}
	return o.VmCurrentProtectionInfo.Get(), o.VmCurrentProtectionInfo.IsSet()
}

// HasVmCurrentProtectionInfo returns a boolean if a field has been set.
func (o *HyperflexProtectionInfoAllOf) HasVmCurrentProtectionInfo() bool {
	if o != nil && o.VmCurrentProtectionInfo.IsSet() {
		return true
	}

	return false
}

// SetVmCurrentProtectionInfo gets a reference to the given NullableHyperflexSnapshotInfoBrief and assigns it to the VmCurrentProtectionInfo field.
func (o *HyperflexProtectionInfoAllOf) SetVmCurrentProtectionInfo(v HyperflexSnapshotInfoBrief) {
	o.VmCurrentProtectionInfo.Set(&v)
}

// SetVmCurrentProtectionInfoNil sets the value for VmCurrentProtectionInfo to be an explicit nil
func (o *HyperflexProtectionInfoAllOf) SetVmCurrentProtectionInfoNil() {
	o.VmCurrentProtectionInfo.Set(nil)
}

// UnsetVmCurrentProtectionInfo ensures that no value is present for VmCurrentProtectionInfo, not even an explicit nil
func (o *HyperflexProtectionInfoAllOf) UnsetVmCurrentProtectionInfo() {
	o.VmCurrentProtectionInfo.Unset()
}

// GetVmLastSuccessfulProtectionInfo returns the VmLastSuccessfulProtectionInfo field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HyperflexProtectionInfoAllOf) GetVmLastSuccessfulProtectionInfo() HyperflexSnapshotInfoBrief {
	if o == nil || o.VmLastSuccessfulProtectionInfo.Get() == nil {
		var ret HyperflexSnapshotInfoBrief
		return ret
	}
	return *o.VmLastSuccessfulProtectionInfo.Get()
}

// GetVmLastSuccessfulProtectionInfoOk returns a tuple with the VmLastSuccessfulProtectionInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HyperflexProtectionInfoAllOf) GetVmLastSuccessfulProtectionInfoOk() (*HyperflexSnapshotInfoBrief, bool) {
	if o == nil {
		return nil, false
	}
	return o.VmLastSuccessfulProtectionInfo.Get(), o.VmLastSuccessfulProtectionInfo.IsSet()
}

// HasVmLastSuccessfulProtectionInfo returns a boolean if a field has been set.
func (o *HyperflexProtectionInfoAllOf) HasVmLastSuccessfulProtectionInfo() bool {
	if o != nil && o.VmLastSuccessfulProtectionInfo.IsSet() {
		return true
	}

	return false
}

// SetVmLastSuccessfulProtectionInfo gets a reference to the given NullableHyperflexSnapshotInfoBrief and assigns it to the VmLastSuccessfulProtectionInfo field.
func (o *HyperflexProtectionInfoAllOf) SetVmLastSuccessfulProtectionInfo(v HyperflexSnapshotInfoBrief) {
	o.VmLastSuccessfulProtectionInfo.Set(&v)
}

// SetVmLastSuccessfulProtectionInfoNil sets the value for VmLastSuccessfulProtectionInfo to be an explicit nil
func (o *HyperflexProtectionInfoAllOf) SetVmLastSuccessfulProtectionInfoNil() {
	o.VmLastSuccessfulProtectionInfo.Set(nil)
}

// UnsetVmLastSuccessfulProtectionInfo ensures that no value is present for VmLastSuccessfulProtectionInfo, not even an explicit nil
func (o *HyperflexProtectionInfoAllOf) UnsetVmLastSuccessfulProtectionInfo() {
	o.VmLastSuccessfulProtectionInfo.Unset()
}

// GetVmSpaceUsage returns the VmSpaceUsage field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HyperflexProtectionInfoAllOf) GetVmSpaceUsage() HyperflexVmProtectionSpaceUsage {
	if o == nil || o.VmSpaceUsage.Get() == nil {
		var ret HyperflexVmProtectionSpaceUsage
		return ret
	}
	return *o.VmSpaceUsage.Get()
}

// GetVmSpaceUsageOk returns a tuple with the VmSpaceUsage field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HyperflexProtectionInfoAllOf) GetVmSpaceUsageOk() (*HyperflexVmProtectionSpaceUsage, bool) {
	if o == nil {
		return nil, false
	}
	return o.VmSpaceUsage.Get(), o.VmSpaceUsage.IsSet()
}

// HasVmSpaceUsage returns a boolean if a field has been set.
func (o *HyperflexProtectionInfoAllOf) HasVmSpaceUsage() bool {
	if o != nil && o.VmSpaceUsage.IsSet() {
		return true
	}

	return false
}

// SetVmSpaceUsage gets a reference to the given NullableHyperflexVmProtectionSpaceUsage and assigns it to the VmSpaceUsage field.
func (o *HyperflexProtectionInfoAllOf) SetVmSpaceUsage(v HyperflexVmProtectionSpaceUsage) {
	o.VmSpaceUsage.Set(&v)
}

// SetVmSpaceUsageNil sets the value for VmSpaceUsage to be an explicit nil
func (o *HyperflexProtectionInfoAllOf) SetVmSpaceUsageNil() {
	o.VmSpaceUsage.Set(nil)
}

// UnsetVmSpaceUsage ensures that no value is present for VmSpaceUsage, not even an explicit nil
func (o *HyperflexProtectionInfoAllOf) UnsetVmSpaceUsage() {
	o.VmSpaceUsage.Unset()
}

func (o HyperflexProtectionInfoAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.VmCurrentProtectionInfo.IsSet() {
		toSerialize["VmCurrentProtectionInfo"] = o.VmCurrentProtectionInfo.Get()
	}
	if o.VmLastSuccessfulProtectionInfo.IsSet() {
		toSerialize["VmLastSuccessfulProtectionInfo"] = o.VmLastSuccessfulProtectionInfo.Get()
	}
	if o.VmSpaceUsage.IsSet() {
		toSerialize["VmSpaceUsage"] = o.VmSpaceUsage.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *HyperflexProtectionInfoAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varHyperflexProtectionInfoAllOf := _HyperflexProtectionInfoAllOf{}

	if err = json.Unmarshal(bytes, &varHyperflexProtectionInfoAllOf); err == nil {
		*o = HyperflexProtectionInfoAllOf(varHyperflexProtectionInfoAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "VmCurrentProtectionInfo")
		delete(additionalProperties, "VmLastSuccessfulProtectionInfo")
		delete(additionalProperties, "VmSpaceUsage")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableHyperflexProtectionInfoAllOf struct {
	value *HyperflexProtectionInfoAllOf
	isSet bool
}

func (v NullableHyperflexProtectionInfoAllOf) Get() *HyperflexProtectionInfoAllOf {
	return v.value
}

func (v *NullableHyperflexProtectionInfoAllOf) Set(val *HyperflexProtectionInfoAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableHyperflexProtectionInfoAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableHyperflexProtectionInfoAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHyperflexProtectionInfoAllOf(val *HyperflexProtectionInfoAllOf) *NullableHyperflexProtectionInfoAllOf {
	return &NullableHyperflexProtectionInfoAllOf{value: val, isSet: true}
}

func (v NullableHyperflexProtectionInfoAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHyperflexProtectionInfoAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
