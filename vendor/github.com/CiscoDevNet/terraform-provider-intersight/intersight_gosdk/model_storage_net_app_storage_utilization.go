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

// StorageNetAppStorageUtilization Storage capacity information which includes savings realized.
type StorageNetAppStorageUtilization struct {
	StorageBaseCapacity
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Total logical used capacity, represented in bytes.
	LogicalUsed *int64 `json:"LogicalUsed,omitempty"`
	// Total savings capacity, represented in bytes.
	Savings              *int64 `json:"Savings,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _StorageNetAppStorageUtilization StorageNetAppStorageUtilization

// NewStorageNetAppStorageUtilization instantiates a new StorageNetAppStorageUtilization object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStorageNetAppStorageUtilization(classId string, objectType string) *StorageNetAppStorageUtilization {
	this := StorageNetAppStorageUtilization{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewStorageNetAppStorageUtilizationWithDefaults instantiates a new StorageNetAppStorageUtilization object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStorageNetAppStorageUtilizationWithDefaults() *StorageNetAppStorageUtilization {
	this := StorageNetAppStorageUtilization{}
	var classId string = "storage.NetAppStorageUtilization"
	this.ClassId = classId
	var objectType string = "storage.NetAppStorageUtilization"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *StorageNetAppStorageUtilization) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *StorageNetAppStorageUtilization) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *StorageNetAppStorageUtilization) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *StorageNetAppStorageUtilization) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *StorageNetAppStorageUtilization) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *StorageNetAppStorageUtilization) SetObjectType(v string) {
	o.ObjectType = v
}

// GetLogicalUsed returns the LogicalUsed field value if set, zero value otherwise.
func (o *StorageNetAppStorageUtilization) GetLogicalUsed() int64 {
	if o == nil || o.LogicalUsed == nil {
		var ret int64
		return ret
	}
	return *o.LogicalUsed
}

// GetLogicalUsedOk returns a tuple with the LogicalUsed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNetAppStorageUtilization) GetLogicalUsedOk() (*int64, bool) {
	if o == nil || o.LogicalUsed == nil {
		return nil, false
	}
	return o.LogicalUsed, true
}

// HasLogicalUsed returns a boolean if a field has been set.
func (o *StorageNetAppStorageUtilization) HasLogicalUsed() bool {
	if o != nil && o.LogicalUsed != nil {
		return true
	}

	return false
}

// SetLogicalUsed gets a reference to the given int64 and assigns it to the LogicalUsed field.
func (o *StorageNetAppStorageUtilization) SetLogicalUsed(v int64) {
	o.LogicalUsed = &v
}

// GetSavings returns the Savings field value if set, zero value otherwise.
func (o *StorageNetAppStorageUtilization) GetSavings() int64 {
	if o == nil || o.Savings == nil {
		var ret int64
		return ret
	}
	return *o.Savings
}

// GetSavingsOk returns a tuple with the Savings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNetAppStorageUtilization) GetSavingsOk() (*int64, bool) {
	if o == nil || o.Savings == nil {
		return nil, false
	}
	return o.Savings, true
}

// HasSavings returns a boolean if a field has been set.
func (o *StorageNetAppStorageUtilization) HasSavings() bool {
	if o != nil && o.Savings != nil {
		return true
	}

	return false
}

// SetSavings gets a reference to the given int64 and assigns it to the Savings field.
func (o *StorageNetAppStorageUtilization) SetSavings(v int64) {
	o.Savings = &v
}

func (o StorageNetAppStorageUtilization) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedStorageBaseCapacity, errStorageBaseCapacity := json.Marshal(o.StorageBaseCapacity)
	if errStorageBaseCapacity != nil {
		return []byte{}, errStorageBaseCapacity
	}
	errStorageBaseCapacity = json.Unmarshal([]byte(serializedStorageBaseCapacity), &toSerialize)
	if errStorageBaseCapacity != nil {
		return []byte{}, errStorageBaseCapacity
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.LogicalUsed != nil {
		toSerialize["LogicalUsed"] = o.LogicalUsed
	}
	if o.Savings != nil {
		toSerialize["Savings"] = o.Savings
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *StorageNetAppStorageUtilization) UnmarshalJSON(bytes []byte) (err error) {
	type StorageNetAppStorageUtilizationWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Total logical used capacity, represented in bytes.
		LogicalUsed *int64 `json:"LogicalUsed,omitempty"`
		// Total savings capacity, represented in bytes.
		Savings *int64 `json:"Savings,omitempty"`
	}

	varStorageNetAppStorageUtilizationWithoutEmbeddedStruct := StorageNetAppStorageUtilizationWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varStorageNetAppStorageUtilizationWithoutEmbeddedStruct)
	if err == nil {
		varStorageNetAppStorageUtilization := _StorageNetAppStorageUtilization{}
		varStorageNetAppStorageUtilization.ClassId = varStorageNetAppStorageUtilizationWithoutEmbeddedStruct.ClassId
		varStorageNetAppStorageUtilization.ObjectType = varStorageNetAppStorageUtilizationWithoutEmbeddedStruct.ObjectType
		varStorageNetAppStorageUtilization.LogicalUsed = varStorageNetAppStorageUtilizationWithoutEmbeddedStruct.LogicalUsed
		varStorageNetAppStorageUtilization.Savings = varStorageNetAppStorageUtilizationWithoutEmbeddedStruct.Savings
		*o = StorageNetAppStorageUtilization(varStorageNetAppStorageUtilization)
	} else {
		return err
	}

	varStorageNetAppStorageUtilization := _StorageNetAppStorageUtilization{}

	err = json.Unmarshal(bytes, &varStorageNetAppStorageUtilization)
	if err == nil {
		o.StorageBaseCapacity = varStorageNetAppStorageUtilization.StorageBaseCapacity
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "LogicalUsed")
		delete(additionalProperties, "Savings")

		// remove fields from embedded structs
		reflectStorageBaseCapacity := reflect.ValueOf(o.StorageBaseCapacity)
		for i := 0; i < reflectStorageBaseCapacity.Type().NumField(); i++ {
			t := reflectStorageBaseCapacity.Type().Field(i)

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

type NullableStorageNetAppStorageUtilization struct {
	value *StorageNetAppStorageUtilization
	isSet bool
}

func (v NullableStorageNetAppStorageUtilization) Get() *StorageNetAppStorageUtilization {
	return v.value
}

func (v *NullableStorageNetAppStorageUtilization) Set(val *StorageNetAppStorageUtilization) {
	v.value = val
	v.isSet = true
}

func (v NullableStorageNetAppStorageUtilization) IsSet() bool {
	return v.isSet
}

func (v *NullableStorageNetAppStorageUtilization) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStorageNetAppStorageUtilization(val *StorageNetAppStorageUtilization) *NullableStorageNetAppStorageUtilization {
	return &NullableStorageNetAppStorageUtilization{value: val, isSet: true}
}

func (v NullableStorageNetAppStorageUtilization) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStorageNetAppStorageUtilization) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
