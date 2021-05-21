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

// CapabilitySwitchingModeCapability Combines and lists the Switching mode capability for each of the Fabric/Switch Platforms.
type CapabilitySwitchingModeCapability struct {
	MoBaseComplexType
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Switching mode type (endhost, switch) of the switch. * `end-host` - In end-host mode, the fabric interconnects appear to the upstream devices as end hosts with multiple links.In this mode, the switch does not run Spanning Tree Protocol and avoids loops by following a set of rules for traffic forwarding.In case of ethernet switching mode - Ethernet end-host mode is also known as Ethernet host virtualizer. * `switch` - In switch mode, the switch runs Spanning Tree Protocol to avoid loops, and broadcast and multicast packets are handled in the traditional way.This is the traditional switch mode.
	SwitchingMode *string `json:"SwitchingMode,omitempty"`
	// VP Compression support on this switch.
	VpCompressionSupported *bool `json:"VpCompressionSupported,omitempty"`
	AdditionalProperties   map[string]interface{}
}

type _CapabilitySwitchingModeCapability CapabilitySwitchingModeCapability

// NewCapabilitySwitchingModeCapability instantiates a new CapabilitySwitchingModeCapability object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCapabilitySwitchingModeCapability(classId string, objectType string) *CapabilitySwitchingModeCapability {
	this := CapabilitySwitchingModeCapability{}
	this.ClassId = classId
	this.ObjectType = objectType
	var switchingMode string = "end-host"
	this.SwitchingMode = &switchingMode
	return &this
}

// NewCapabilitySwitchingModeCapabilityWithDefaults instantiates a new CapabilitySwitchingModeCapability object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCapabilitySwitchingModeCapabilityWithDefaults() *CapabilitySwitchingModeCapability {
	this := CapabilitySwitchingModeCapability{}
	var classId string = "capability.SwitchingModeCapability"
	this.ClassId = classId
	var objectType string = "capability.SwitchingModeCapability"
	this.ObjectType = objectType
	var switchingMode string = "end-host"
	this.SwitchingMode = &switchingMode
	return &this
}

// GetClassId returns the ClassId field value
func (o *CapabilitySwitchingModeCapability) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *CapabilitySwitchingModeCapability) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *CapabilitySwitchingModeCapability) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *CapabilitySwitchingModeCapability) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *CapabilitySwitchingModeCapability) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *CapabilitySwitchingModeCapability) SetObjectType(v string) {
	o.ObjectType = v
}

// GetSwitchingMode returns the SwitchingMode field value if set, zero value otherwise.
func (o *CapabilitySwitchingModeCapability) GetSwitchingMode() string {
	if o == nil || o.SwitchingMode == nil {
		var ret string
		return ret
	}
	return *o.SwitchingMode
}

// GetSwitchingModeOk returns a tuple with the SwitchingMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CapabilitySwitchingModeCapability) GetSwitchingModeOk() (*string, bool) {
	if o == nil || o.SwitchingMode == nil {
		return nil, false
	}
	return o.SwitchingMode, true
}

// HasSwitchingMode returns a boolean if a field has been set.
func (o *CapabilitySwitchingModeCapability) HasSwitchingMode() bool {
	if o != nil && o.SwitchingMode != nil {
		return true
	}

	return false
}

// SetSwitchingMode gets a reference to the given string and assigns it to the SwitchingMode field.
func (o *CapabilitySwitchingModeCapability) SetSwitchingMode(v string) {
	o.SwitchingMode = &v
}

// GetVpCompressionSupported returns the VpCompressionSupported field value if set, zero value otherwise.
func (o *CapabilitySwitchingModeCapability) GetVpCompressionSupported() bool {
	if o == nil || o.VpCompressionSupported == nil {
		var ret bool
		return ret
	}
	return *o.VpCompressionSupported
}

// GetVpCompressionSupportedOk returns a tuple with the VpCompressionSupported field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CapabilitySwitchingModeCapability) GetVpCompressionSupportedOk() (*bool, bool) {
	if o == nil || o.VpCompressionSupported == nil {
		return nil, false
	}
	return o.VpCompressionSupported, true
}

// HasVpCompressionSupported returns a boolean if a field has been set.
func (o *CapabilitySwitchingModeCapability) HasVpCompressionSupported() bool {
	if o != nil && o.VpCompressionSupported != nil {
		return true
	}

	return false
}

// SetVpCompressionSupported gets a reference to the given bool and assigns it to the VpCompressionSupported field.
func (o *CapabilitySwitchingModeCapability) SetVpCompressionSupported(v bool) {
	o.VpCompressionSupported = &v
}

func (o CapabilitySwitchingModeCapability) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedMoBaseComplexType, errMoBaseComplexType := json.Marshal(o.MoBaseComplexType)
	if errMoBaseComplexType != nil {
		return []byte{}, errMoBaseComplexType
	}
	errMoBaseComplexType = json.Unmarshal([]byte(serializedMoBaseComplexType), &toSerialize)
	if errMoBaseComplexType != nil {
		return []byte{}, errMoBaseComplexType
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.SwitchingMode != nil {
		toSerialize["SwitchingMode"] = o.SwitchingMode
	}
	if o.VpCompressionSupported != nil {
		toSerialize["VpCompressionSupported"] = o.VpCompressionSupported
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *CapabilitySwitchingModeCapability) UnmarshalJSON(bytes []byte) (err error) {
	type CapabilitySwitchingModeCapabilityWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Switching mode type (endhost, switch) of the switch. * `end-host` - In end-host mode, the fabric interconnects appear to the upstream devices as end hosts with multiple links.In this mode, the switch does not run Spanning Tree Protocol and avoids loops by following a set of rules for traffic forwarding.In case of ethernet switching mode - Ethernet end-host mode is also known as Ethernet host virtualizer. * `switch` - In switch mode, the switch runs Spanning Tree Protocol to avoid loops, and broadcast and multicast packets are handled in the traditional way.This is the traditional switch mode.
		SwitchingMode *string `json:"SwitchingMode,omitempty"`
		// VP Compression support on this switch.
		VpCompressionSupported *bool `json:"VpCompressionSupported,omitempty"`
	}

	varCapabilitySwitchingModeCapabilityWithoutEmbeddedStruct := CapabilitySwitchingModeCapabilityWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varCapabilitySwitchingModeCapabilityWithoutEmbeddedStruct)
	if err == nil {
		varCapabilitySwitchingModeCapability := _CapabilitySwitchingModeCapability{}
		varCapabilitySwitchingModeCapability.ClassId = varCapabilitySwitchingModeCapabilityWithoutEmbeddedStruct.ClassId
		varCapabilitySwitchingModeCapability.ObjectType = varCapabilitySwitchingModeCapabilityWithoutEmbeddedStruct.ObjectType
		varCapabilitySwitchingModeCapability.SwitchingMode = varCapabilitySwitchingModeCapabilityWithoutEmbeddedStruct.SwitchingMode
		varCapabilitySwitchingModeCapability.VpCompressionSupported = varCapabilitySwitchingModeCapabilityWithoutEmbeddedStruct.VpCompressionSupported
		*o = CapabilitySwitchingModeCapability(varCapabilitySwitchingModeCapability)
	} else {
		return err
	}

	varCapabilitySwitchingModeCapability := _CapabilitySwitchingModeCapability{}

	err = json.Unmarshal(bytes, &varCapabilitySwitchingModeCapability)
	if err == nil {
		o.MoBaseComplexType = varCapabilitySwitchingModeCapability.MoBaseComplexType
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "SwitchingMode")
		delete(additionalProperties, "VpCompressionSupported")

		// remove fields from embedded structs
		reflectMoBaseComplexType := reflect.ValueOf(o.MoBaseComplexType)
		for i := 0; i < reflectMoBaseComplexType.Type().NumField(); i++ {
			t := reflectMoBaseComplexType.Type().Field(i)

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

type NullableCapabilitySwitchingModeCapability struct {
	value *CapabilitySwitchingModeCapability
	isSet bool
}

func (v NullableCapabilitySwitchingModeCapability) Get() *CapabilitySwitchingModeCapability {
	return v.value
}

func (v *NullableCapabilitySwitchingModeCapability) Set(val *CapabilitySwitchingModeCapability) {
	v.value = val
	v.isSet = true
}

func (v NullableCapabilitySwitchingModeCapability) IsSet() bool {
	return v.isSet
}

func (v *NullableCapabilitySwitchingModeCapability) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCapabilitySwitchingModeCapability(val *CapabilitySwitchingModeCapability) *NullableCapabilitySwitchingModeCapability {
	return &NullableCapabilitySwitchingModeCapability{value: val, isSet: true}
}

func (v NullableCapabilitySwitchingModeCapability) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCapabilitySwitchingModeCapability) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
