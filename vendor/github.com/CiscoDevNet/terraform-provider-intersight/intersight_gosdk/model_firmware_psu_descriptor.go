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
	"reflect"
	"strings"
)

// FirmwarePsuDescriptor Descriptor to uniquely identify a PSU component.
type FirmwarePsuDescriptor struct {
	FirmwareComponentDescriptor
	AdditionalProperties map[string]interface{}
}

type _FirmwarePsuDescriptor FirmwarePsuDescriptor

// NewFirmwarePsuDescriptor instantiates a new FirmwarePsuDescriptor object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFirmwarePsuDescriptor(classId string, objectType string) *FirmwarePsuDescriptor {
	this := FirmwarePsuDescriptor{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewFirmwarePsuDescriptorWithDefaults instantiates a new FirmwarePsuDescriptor object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFirmwarePsuDescriptorWithDefaults() *FirmwarePsuDescriptor {
	this := FirmwarePsuDescriptor{}
	return &this
}

func (o FirmwarePsuDescriptor) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedFirmwareComponentDescriptor, errFirmwareComponentDescriptor := json.Marshal(o.FirmwareComponentDescriptor)
	if errFirmwareComponentDescriptor != nil {
		return []byte{}, errFirmwareComponentDescriptor
	}
	errFirmwareComponentDescriptor = json.Unmarshal([]byte(serializedFirmwareComponentDescriptor), &toSerialize)
	if errFirmwareComponentDescriptor != nil {
		return []byte{}, errFirmwareComponentDescriptor
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *FirmwarePsuDescriptor) UnmarshalJSON(bytes []byte) (err error) {
	type FirmwarePsuDescriptorWithoutEmbeddedStruct struct {
	}

	varFirmwarePsuDescriptorWithoutEmbeddedStruct := FirmwarePsuDescriptorWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varFirmwarePsuDescriptorWithoutEmbeddedStruct)
	if err == nil {
		varFirmwarePsuDescriptor := _FirmwarePsuDescriptor{}
		*o = FirmwarePsuDescriptor(varFirmwarePsuDescriptor)
	} else {
		return err
	}

	varFirmwarePsuDescriptor := _FirmwarePsuDescriptor{}

	err = json.Unmarshal(bytes, &varFirmwarePsuDescriptor)
	if err == nil {
		o.FirmwareComponentDescriptor = varFirmwarePsuDescriptor.FirmwareComponentDescriptor
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {

		// remove fields from embedded structs
		reflectFirmwareComponentDescriptor := reflect.ValueOf(o.FirmwareComponentDescriptor)
		for i := 0; i < reflectFirmwareComponentDescriptor.Type().NumField(); i++ {
			t := reflectFirmwareComponentDescriptor.Type().Field(i)

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

type NullableFirmwarePsuDescriptor struct {
	value *FirmwarePsuDescriptor
	isSet bool
}

func (v NullableFirmwarePsuDescriptor) Get() *FirmwarePsuDescriptor {
	return v.value
}

func (v *NullableFirmwarePsuDescriptor) Set(val *FirmwarePsuDescriptor) {
	v.value = val
	v.isSet = true
}

func (v NullableFirmwarePsuDescriptor) IsSet() bool {
	return v.isSet
}

func (v *NullableFirmwarePsuDescriptor) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFirmwarePsuDescriptor(val *FirmwarePsuDescriptor) *NullableFirmwarePsuDescriptor {
	return &NullableFirmwarePsuDescriptor{value: val, isSet: true}
}

func (v NullableFirmwarePsuDescriptor) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFirmwarePsuDescriptor) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
