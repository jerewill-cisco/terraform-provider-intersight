/*
 * Cisco Intersight
 *
 * Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document. This document was created on 2021-06-30T12:14:04Z.
 *
 * API version: 1.0.9-4375
 * Contact: intersight@cisco.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package intersight

import (
	"encoding/json"
)

// VnicCdnAllOf Definition of the list of properties defined in 'vnic.Cdn', excluding properties defined in parent classes.
type VnicCdnAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Source of the CDN. It can either be user specified or be the same as the vNIC name. * `vnic` - Source of the CDN is the same as the vNIC name. * `user` - Source of the CDN is specified by the user.
	Source *string `json:"Source,omitempty"`
	// The CDN value entered in case of user defined mode.
	Value                *string `json:"Value,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _VnicCdnAllOf VnicCdnAllOf

// NewVnicCdnAllOf instantiates a new VnicCdnAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVnicCdnAllOf(classId string, objectType string) *VnicCdnAllOf {
	this := VnicCdnAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	var source string = "vnic"
	this.Source = &source
	return &this
}

// NewVnicCdnAllOfWithDefaults instantiates a new VnicCdnAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVnicCdnAllOfWithDefaults() *VnicCdnAllOf {
	this := VnicCdnAllOf{}
	var classId string = "vnic.Cdn"
	this.ClassId = classId
	var objectType string = "vnic.Cdn"
	this.ObjectType = objectType
	var source string = "vnic"
	this.Source = &source
	return &this
}

// GetClassId returns the ClassId field value
func (o *VnicCdnAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *VnicCdnAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *VnicCdnAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *VnicCdnAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *VnicCdnAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *VnicCdnAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetSource returns the Source field value if set, zero value otherwise.
func (o *VnicCdnAllOf) GetSource() string {
	if o == nil || o.Source == nil {
		var ret string
		return ret
	}
	return *o.Source
}

// GetSourceOk returns a tuple with the Source field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VnicCdnAllOf) GetSourceOk() (*string, bool) {
	if o == nil || o.Source == nil {
		return nil, false
	}
	return o.Source, true
}

// HasSource returns a boolean if a field has been set.
func (o *VnicCdnAllOf) HasSource() bool {
	if o != nil && o.Source != nil {
		return true
	}

	return false
}

// SetSource gets a reference to the given string and assigns it to the Source field.
func (o *VnicCdnAllOf) SetSource(v string) {
	o.Source = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *VnicCdnAllOf) GetValue() string {
	if o == nil || o.Value == nil {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VnicCdnAllOf) GetValueOk() (*string, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *VnicCdnAllOf) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *VnicCdnAllOf) SetValue(v string) {
	o.Value = &v
}

func (o VnicCdnAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Source != nil {
		toSerialize["Source"] = o.Source
	}
	if o.Value != nil {
		toSerialize["Value"] = o.Value
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *VnicCdnAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varVnicCdnAllOf := _VnicCdnAllOf{}

	if err = json.Unmarshal(bytes, &varVnicCdnAllOf); err == nil {
		*o = VnicCdnAllOf(varVnicCdnAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Source")
		delete(additionalProperties, "Value")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableVnicCdnAllOf struct {
	value *VnicCdnAllOf
	isSet bool
}

func (v NullableVnicCdnAllOf) Get() *VnicCdnAllOf {
	return v.value
}

func (v *NullableVnicCdnAllOf) Set(val *VnicCdnAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableVnicCdnAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableVnicCdnAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVnicCdnAllOf(val *VnicCdnAllOf) *NullableVnicCdnAllOf {
	return &NullableVnicCdnAllOf{value: val, isSet: true}
}

func (v NullableVnicCdnAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVnicCdnAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
