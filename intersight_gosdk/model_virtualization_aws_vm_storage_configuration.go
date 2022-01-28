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

// VirtualizationAwsVmStorageConfiguration Specify Amazon Web Service VM memory configuration.
type VirtualizationAwsVmStorageConfiguration struct {
	VirtualizationCloudVmStorageConfiguration
	AdditionalProperties map[string]interface{}
}

type _VirtualizationAwsVmStorageConfiguration VirtualizationAwsVmStorageConfiguration

// NewVirtualizationAwsVmStorageConfiguration instantiates a new VirtualizationAwsVmStorageConfiguration object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVirtualizationAwsVmStorageConfiguration(classId string, objectType string) *VirtualizationAwsVmStorageConfiguration {
	this := VirtualizationAwsVmStorageConfiguration{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewVirtualizationAwsVmStorageConfigurationWithDefaults instantiates a new VirtualizationAwsVmStorageConfiguration object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVirtualizationAwsVmStorageConfigurationWithDefaults() *VirtualizationAwsVmStorageConfiguration {
	this := VirtualizationAwsVmStorageConfiguration{}
	return &this
}

func (o VirtualizationAwsVmStorageConfiguration) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedVirtualizationCloudVmStorageConfiguration, errVirtualizationCloudVmStorageConfiguration := json.Marshal(o.VirtualizationCloudVmStorageConfiguration)
	if errVirtualizationCloudVmStorageConfiguration != nil {
		return []byte{}, errVirtualizationCloudVmStorageConfiguration
	}
	errVirtualizationCloudVmStorageConfiguration = json.Unmarshal([]byte(serializedVirtualizationCloudVmStorageConfiguration), &toSerialize)
	if errVirtualizationCloudVmStorageConfiguration != nil {
		return []byte{}, errVirtualizationCloudVmStorageConfiguration
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *VirtualizationAwsVmStorageConfiguration) UnmarshalJSON(bytes []byte) (err error) {
	type VirtualizationAwsVmStorageConfigurationWithoutEmbeddedStruct struct {
	}

	varVirtualizationAwsVmStorageConfigurationWithoutEmbeddedStruct := VirtualizationAwsVmStorageConfigurationWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varVirtualizationAwsVmStorageConfigurationWithoutEmbeddedStruct)
	if err == nil {
		varVirtualizationAwsVmStorageConfiguration := _VirtualizationAwsVmStorageConfiguration{}
		*o = VirtualizationAwsVmStorageConfiguration(varVirtualizationAwsVmStorageConfiguration)
	} else {
		return err
	}

	varVirtualizationAwsVmStorageConfiguration := _VirtualizationAwsVmStorageConfiguration{}

	err = json.Unmarshal(bytes, &varVirtualizationAwsVmStorageConfiguration)
	if err == nil {
		o.VirtualizationCloudVmStorageConfiguration = varVirtualizationAwsVmStorageConfiguration.VirtualizationCloudVmStorageConfiguration
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {

		// remove fields from embedded structs
		reflectVirtualizationCloudVmStorageConfiguration := reflect.ValueOf(o.VirtualizationCloudVmStorageConfiguration)
		for i := 0; i < reflectVirtualizationCloudVmStorageConfiguration.Type().NumField(); i++ {
			t := reflectVirtualizationCloudVmStorageConfiguration.Type().Field(i)

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

type NullableVirtualizationAwsVmStorageConfiguration struct {
	value *VirtualizationAwsVmStorageConfiguration
	isSet bool
}

func (v NullableVirtualizationAwsVmStorageConfiguration) Get() *VirtualizationAwsVmStorageConfiguration {
	return v.value
}

func (v *NullableVirtualizationAwsVmStorageConfiguration) Set(val *VirtualizationAwsVmStorageConfiguration) {
	v.value = val
	v.isSet = true
}

func (v NullableVirtualizationAwsVmStorageConfiguration) IsSet() bool {
	return v.isSet
}

func (v *NullableVirtualizationAwsVmStorageConfiguration) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVirtualizationAwsVmStorageConfiguration(val *VirtualizationAwsVmStorageConfiguration) *NullableVirtualizationAwsVmStorageConfiguration {
	return &NullableVirtualizationAwsVmStorageConfiguration{value: val, isSet: true}
}

func (v NullableVirtualizationAwsVmStorageConfiguration) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVirtualizationAwsVmStorageConfiguration) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
