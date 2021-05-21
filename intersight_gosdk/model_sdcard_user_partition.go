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

// SdcardUserPartition This partition is available for user to use. This drive is available under utility partition.
type SdcardUserPartition struct {
	SdcardVirtualDrive
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Name of virtual drive for user partition.
	Name                 *string `json:"Name,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SdcardUserPartition SdcardUserPartition

// NewSdcardUserPartition instantiates a new SdcardUserPartition object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSdcardUserPartition(classId string, objectType string) *SdcardUserPartition {
	this := SdcardUserPartition{}
	this.ClassId = classId
	this.ObjectType = objectType
	var name string = "UserPartition"
	this.Name = &name
	return &this
}

// NewSdcardUserPartitionWithDefaults instantiates a new SdcardUserPartition object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSdcardUserPartitionWithDefaults() *SdcardUserPartition {
	this := SdcardUserPartition{}
	var classId string = "sdcard.UserPartition"
	this.ClassId = classId
	var objectType string = "sdcard.UserPartition"
	this.ObjectType = objectType
	var name string = "UserPartition"
	this.Name = &name
	return &this
}

// GetClassId returns the ClassId field value
func (o *SdcardUserPartition) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *SdcardUserPartition) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *SdcardUserPartition) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *SdcardUserPartition) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *SdcardUserPartition) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *SdcardUserPartition) SetObjectType(v string) {
	o.ObjectType = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *SdcardUserPartition) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SdcardUserPartition) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *SdcardUserPartition) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *SdcardUserPartition) SetName(v string) {
	o.Name = &v
}

func (o SdcardUserPartition) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedSdcardVirtualDrive, errSdcardVirtualDrive := json.Marshal(o.SdcardVirtualDrive)
	if errSdcardVirtualDrive != nil {
		return []byte{}, errSdcardVirtualDrive
	}
	errSdcardVirtualDrive = json.Unmarshal([]byte(serializedSdcardVirtualDrive), &toSerialize)
	if errSdcardVirtualDrive != nil {
		return []byte{}, errSdcardVirtualDrive
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Name != nil {
		toSerialize["Name"] = o.Name
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *SdcardUserPartition) UnmarshalJSON(bytes []byte) (err error) {
	type SdcardUserPartitionWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Name of virtual drive for user partition.
		Name *string `json:"Name,omitempty"`
	}

	varSdcardUserPartitionWithoutEmbeddedStruct := SdcardUserPartitionWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varSdcardUserPartitionWithoutEmbeddedStruct)
	if err == nil {
		varSdcardUserPartition := _SdcardUserPartition{}
		varSdcardUserPartition.ClassId = varSdcardUserPartitionWithoutEmbeddedStruct.ClassId
		varSdcardUserPartition.ObjectType = varSdcardUserPartitionWithoutEmbeddedStruct.ObjectType
		varSdcardUserPartition.Name = varSdcardUserPartitionWithoutEmbeddedStruct.Name
		*o = SdcardUserPartition(varSdcardUserPartition)
	} else {
		return err
	}

	varSdcardUserPartition := _SdcardUserPartition{}

	err = json.Unmarshal(bytes, &varSdcardUserPartition)
	if err == nil {
		o.SdcardVirtualDrive = varSdcardUserPartition.SdcardVirtualDrive
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Name")

		// remove fields from embedded structs
		reflectSdcardVirtualDrive := reflect.ValueOf(o.SdcardVirtualDrive)
		for i := 0; i < reflectSdcardVirtualDrive.Type().NumField(); i++ {
			t := reflectSdcardVirtualDrive.Type().Field(i)

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

type NullableSdcardUserPartition struct {
	value *SdcardUserPartition
	isSet bool
}

func (v NullableSdcardUserPartition) Get() *SdcardUserPartition {
	return v.value
}

func (v *NullableSdcardUserPartition) Set(val *SdcardUserPartition) {
	v.value = val
	v.isSet = true
}

func (v NullableSdcardUserPartition) IsSet() bool {
	return v.isSet
}

func (v *NullableSdcardUserPartition) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSdcardUserPartition(val *SdcardUserPartition) *NullableSdcardUserPartition {
	return &NullableSdcardUserPartition{value: val, isSet: true}
}

func (v NullableSdcardUserPartition) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSdcardUserPartition) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
