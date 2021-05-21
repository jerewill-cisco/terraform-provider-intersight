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

// StorageHitachiInitiator An initiator is the consumer of Hitachi Array, typically a server with an adapter card in it called a Host Bus Adapter (HBA). The initiator \"initiates\" a connection over the fabric to one or more ports on Hitachi Array ports.
type StorageHitachiInitiator struct {
	StorageBaseInitiator
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// World wide port name, 64 bit address represented in hexa decimal notation.
	Wwpn                 *string `json:"Wwpn,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _StorageHitachiInitiator StorageHitachiInitiator

// NewStorageHitachiInitiator instantiates a new StorageHitachiInitiator object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStorageHitachiInitiator(classId string, objectType string) *StorageHitachiInitiator {
	this := StorageHitachiInitiator{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewStorageHitachiInitiatorWithDefaults instantiates a new StorageHitachiInitiator object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStorageHitachiInitiatorWithDefaults() *StorageHitachiInitiator {
	this := StorageHitachiInitiator{}
	var classId string = "storage.HitachiInitiator"
	this.ClassId = classId
	var objectType string = "storage.HitachiInitiator"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *StorageHitachiInitiator) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *StorageHitachiInitiator) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *StorageHitachiInitiator) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *StorageHitachiInitiator) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *StorageHitachiInitiator) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *StorageHitachiInitiator) SetObjectType(v string) {
	o.ObjectType = v
}

// GetWwpn returns the Wwpn field value if set, zero value otherwise.
func (o *StorageHitachiInitiator) GetWwpn() string {
	if o == nil || o.Wwpn == nil {
		var ret string
		return ret
	}
	return *o.Wwpn
}

// GetWwpnOk returns a tuple with the Wwpn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageHitachiInitiator) GetWwpnOk() (*string, bool) {
	if o == nil || o.Wwpn == nil {
		return nil, false
	}
	return o.Wwpn, true
}

// HasWwpn returns a boolean if a field has been set.
func (o *StorageHitachiInitiator) HasWwpn() bool {
	if o != nil && o.Wwpn != nil {
		return true
	}

	return false
}

// SetWwpn gets a reference to the given string and assigns it to the Wwpn field.
func (o *StorageHitachiInitiator) SetWwpn(v string) {
	o.Wwpn = &v
}

func (o StorageHitachiInitiator) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedStorageBaseInitiator, errStorageBaseInitiator := json.Marshal(o.StorageBaseInitiator)
	if errStorageBaseInitiator != nil {
		return []byte{}, errStorageBaseInitiator
	}
	errStorageBaseInitiator = json.Unmarshal([]byte(serializedStorageBaseInitiator), &toSerialize)
	if errStorageBaseInitiator != nil {
		return []byte{}, errStorageBaseInitiator
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Wwpn != nil {
		toSerialize["Wwpn"] = o.Wwpn
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *StorageHitachiInitiator) UnmarshalJSON(bytes []byte) (err error) {
	type StorageHitachiInitiatorWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// World wide port name, 64 bit address represented in hexa decimal notation.
		Wwpn *string `json:"Wwpn,omitempty"`
	}

	varStorageHitachiInitiatorWithoutEmbeddedStruct := StorageHitachiInitiatorWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varStorageHitachiInitiatorWithoutEmbeddedStruct)
	if err == nil {
		varStorageHitachiInitiator := _StorageHitachiInitiator{}
		varStorageHitachiInitiator.ClassId = varStorageHitachiInitiatorWithoutEmbeddedStruct.ClassId
		varStorageHitachiInitiator.ObjectType = varStorageHitachiInitiatorWithoutEmbeddedStruct.ObjectType
		varStorageHitachiInitiator.Wwpn = varStorageHitachiInitiatorWithoutEmbeddedStruct.Wwpn
		*o = StorageHitachiInitiator(varStorageHitachiInitiator)
	} else {
		return err
	}

	varStorageHitachiInitiator := _StorageHitachiInitiator{}

	err = json.Unmarshal(bytes, &varStorageHitachiInitiator)
	if err == nil {
		o.StorageBaseInitiator = varStorageHitachiInitiator.StorageBaseInitiator
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Wwpn")

		// remove fields from embedded structs
		reflectStorageBaseInitiator := reflect.ValueOf(o.StorageBaseInitiator)
		for i := 0; i < reflectStorageBaseInitiator.Type().NumField(); i++ {
			t := reflectStorageBaseInitiator.Type().Field(i)

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

type NullableStorageHitachiInitiator struct {
	value *StorageHitachiInitiator
	isSet bool
}

func (v NullableStorageHitachiInitiator) Get() *StorageHitachiInitiator {
	return v.value
}

func (v *NullableStorageHitachiInitiator) Set(val *StorageHitachiInitiator) {
	v.value = val
	v.isSet = true
}

func (v NullableStorageHitachiInitiator) IsSet() bool {
	return v.isSet
}

func (v *NullableStorageHitachiInitiator) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStorageHitachiInitiator(val *StorageHitachiInitiator) *NullableStorageHitachiInitiator {
	return &NullableStorageHitachiInitiator{value: val, isSet: true}
}

func (v NullableStorageHitachiInitiator) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStorageHitachiInitiator) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
