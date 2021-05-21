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
)

// CapabilitySiocModuleCapabilityDefAllOf Definition of the list of properties defined in 'capability.SiocModuleCapabilityDef', excluding properties defined in parent classes.
type CapabilitySiocModuleCapabilityDefAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Device connector support on SIOC.
	DcSupported          *bool `json:"DcSupported,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CapabilitySiocModuleCapabilityDefAllOf CapabilitySiocModuleCapabilityDefAllOf

// NewCapabilitySiocModuleCapabilityDefAllOf instantiates a new CapabilitySiocModuleCapabilityDefAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCapabilitySiocModuleCapabilityDefAllOf(classId string, objectType string) *CapabilitySiocModuleCapabilityDefAllOf {
	this := CapabilitySiocModuleCapabilityDefAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewCapabilitySiocModuleCapabilityDefAllOfWithDefaults instantiates a new CapabilitySiocModuleCapabilityDefAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCapabilitySiocModuleCapabilityDefAllOfWithDefaults() *CapabilitySiocModuleCapabilityDefAllOf {
	this := CapabilitySiocModuleCapabilityDefAllOf{}
	var classId string = "capability.SiocModuleCapabilityDef"
	this.ClassId = classId
	var objectType string = "capability.SiocModuleCapabilityDef"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *CapabilitySiocModuleCapabilityDefAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *CapabilitySiocModuleCapabilityDefAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *CapabilitySiocModuleCapabilityDefAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *CapabilitySiocModuleCapabilityDefAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *CapabilitySiocModuleCapabilityDefAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *CapabilitySiocModuleCapabilityDefAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDcSupported returns the DcSupported field value if set, zero value otherwise.
func (o *CapabilitySiocModuleCapabilityDefAllOf) GetDcSupported() bool {
	if o == nil || o.DcSupported == nil {
		var ret bool
		return ret
	}
	return *o.DcSupported
}

// GetDcSupportedOk returns a tuple with the DcSupported field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CapabilitySiocModuleCapabilityDefAllOf) GetDcSupportedOk() (*bool, bool) {
	if o == nil || o.DcSupported == nil {
		return nil, false
	}
	return o.DcSupported, true
}

// HasDcSupported returns a boolean if a field has been set.
func (o *CapabilitySiocModuleCapabilityDefAllOf) HasDcSupported() bool {
	if o != nil && o.DcSupported != nil {
		return true
	}

	return false
}

// SetDcSupported gets a reference to the given bool and assigns it to the DcSupported field.
func (o *CapabilitySiocModuleCapabilityDefAllOf) SetDcSupported(v bool) {
	o.DcSupported = &v
}

func (o CapabilitySiocModuleCapabilityDefAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.DcSupported != nil {
		toSerialize["DcSupported"] = o.DcSupported
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *CapabilitySiocModuleCapabilityDefAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varCapabilitySiocModuleCapabilityDefAllOf := _CapabilitySiocModuleCapabilityDefAllOf{}

	if err = json.Unmarshal(bytes, &varCapabilitySiocModuleCapabilityDefAllOf); err == nil {
		*o = CapabilitySiocModuleCapabilityDefAllOf(varCapabilitySiocModuleCapabilityDefAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "DcSupported")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCapabilitySiocModuleCapabilityDefAllOf struct {
	value *CapabilitySiocModuleCapabilityDefAllOf
	isSet bool
}

func (v NullableCapabilitySiocModuleCapabilityDefAllOf) Get() *CapabilitySiocModuleCapabilityDefAllOf {
	return v.value
}

func (v *NullableCapabilitySiocModuleCapabilityDefAllOf) Set(val *CapabilitySiocModuleCapabilityDefAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableCapabilitySiocModuleCapabilityDefAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableCapabilitySiocModuleCapabilityDefAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCapabilitySiocModuleCapabilityDefAllOf(val *CapabilitySiocModuleCapabilityDefAllOf) *NullableCapabilitySiocModuleCapabilityDefAllOf {
	return &NullableCapabilitySiocModuleCapabilityDefAllOf{value: val, isSet: true}
}

func (v NullableCapabilitySiocModuleCapabilityDefAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCapabilitySiocModuleCapabilityDefAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
