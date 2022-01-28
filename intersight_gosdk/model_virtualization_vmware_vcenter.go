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

// VirtualizationVmwareVcenter VMware vCenter entity. The vCenter has a name assigned by user in Intersight.
type VirtualizationVmwareVcenter struct {
	VirtualizationBaseHypervisorManager
	AdditionalProperties map[string]interface{}
}

type _VirtualizationVmwareVcenter VirtualizationVmwareVcenter

// NewVirtualizationVmwareVcenter instantiates a new VirtualizationVmwareVcenter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVirtualizationVmwareVcenter(classId string, objectType string) *VirtualizationVmwareVcenter {
	this := VirtualizationVmwareVcenter{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewVirtualizationVmwareVcenterWithDefaults instantiates a new VirtualizationVmwareVcenter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVirtualizationVmwareVcenterWithDefaults() *VirtualizationVmwareVcenter {
	this := VirtualizationVmwareVcenter{}
	return &this
}

func (o VirtualizationVmwareVcenter) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedVirtualizationBaseHypervisorManager, errVirtualizationBaseHypervisorManager := json.Marshal(o.VirtualizationBaseHypervisorManager)
	if errVirtualizationBaseHypervisorManager != nil {
		return []byte{}, errVirtualizationBaseHypervisorManager
	}
	errVirtualizationBaseHypervisorManager = json.Unmarshal([]byte(serializedVirtualizationBaseHypervisorManager), &toSerialize)
	if errVirtualizationBaseHypervisorManager != nil {
		return []byte{}, errVirtualizationBaseHypervisorManager
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *VirtualizationVmwareVcenter) UnmarshalJSON(bytes []byte) (err error) {
	type VirtualizationVmwareVcenterWithoutEmbeddedStruct struct {
	}

	varVirtualizationVmwareVcenterWithoutEmbeddedStruct := VirtualizationVmwareVcenterWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varVirtualizationVmwareVcenterWithoutEmbeddedStruct)
	if err == nil {
		varVirtualizationVmwareVcenter := _VirtualizationVmwareVcenter{}
		*o = VirtualizationVmwareVcenter(varVirtualizationVmwareVcenter)
	} else {
		return err
	}

	varVirtualizationVmwareVcenter := _VirtualizationVmwareVcenter{}

	err = json.Unmarshal(bytes, &varVirtualizationVmwareVcenter)
	if err == nil {
		o.VirtualizationBaseHypervisorManager = varVirtualizationVmwareVcenter.VirtualizationBaseHypervisorManager
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {

		// remove fields from embedded structs
		reflectVirtualizationBaseHypervisorManager := reflect.ValueOf(o.VirtualizationBaseHypervisorManager)
		for i := 0; i < reflectVirtualizationBaseHypervisorManager.Type().NumField(); i++ {
			t := reflectVirtualizationBaseHypervisorManager.Type().Field(i)

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

type NullableVirtualizationVmwareVcenter struct {
	value *VirtualizationVmwareVcenter
	isSet bool
}

func (v NullableVirtualizationVmwareVcenter) Get() *VirtualizationVmwareVcenter {
	return v.value
}

func (v *NullableVirtualizationVmwareVcenter) Set(val *VirtualizationVmwareVcenter) {
	v.value = val
	v.isSet = true
}

func (v NullableVirtualizationVmwareVcenter) IsSet() bool {
	return v.isSet
}

func (v *NullableVirtualizationVmwareVcenter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVirtualizationVmwareVcenter(val *VirtualizationVmwareVcenter) *NullableVirtualizationVmwareVcenter {
	return &NullableVirtualizationVmwareVcenter{value: val, isSet: true}
}

func (v NullableVirtualizationVmwareVcenter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVirtualizationVmwareVcenter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
