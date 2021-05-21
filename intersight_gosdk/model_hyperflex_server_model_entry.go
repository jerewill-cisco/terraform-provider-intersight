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

// HyperflexServerModelEntry A constraint specifying supported server models in regex format.
type HyperflexServerModelEntry struct {
	HyperflexAbstractAppSetting
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType           string                                `json:"ObjectType"`
	Constraint           NullableHyperflexAppSettingConstraint `json:"Constraint,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _HyperflexServerModelEntry HyperflexServerModelEntry

// NewHyperflexServerModelEntry instantiates a new HyperflexServerModelEntry object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHyperflexServerModelEntry(classId string, objectType string) *HyperflexServerModelEntry {
	this := HyperflexServerModelEntry{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewHyperflexServerModelEntryWithDefaults instantiates a new HyperflexServerModelEntry object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHyperflexServerModelEntryWithDefaults() *HyperflexServerModelEntry {
	this := HyperflexServerModelEntry{}
	var classId string = "hyperflex.ServerModelEntry"
	this.ClassId = classId
	var objectType string = "hyperflex.ServerModelEntry"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *HyperflexServerModelEntry) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *HyperflexServerModelEntry) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *HyperflexServerModelEntry) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *HyperflexServerModelEntry) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *HyperflexServerModelEntry) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *HyperflexServerModelEntry) SetObjectType(v string) {
	o.ObjectType = v
}

// GetConstraint returns the Constraint field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HyperflexServerModelEntry) GetConstraint() HyperflexAppSettingConstraint {
	if o == nil || o.Constraint.Get() == nil {
		var ret HyperflexAppSettingConstraint
		return ret
	}
	return *o.Constraint.Get()
}

// GetConstraintOk returns a tuple with the Constraint field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HyperflexServerModelEntry) GetConstraintOk() (*HyperflexAppSettingConstraint, bool) {
	if o == nil {
		return nil, false
	}
	return o.Constraint.Get(), o.Constraint.IsSet()
}

// HasConstraint returns a boolean if a field has been set.
func (o *HyperflexServerModelEntry) HasConstraint() bool {
	if o != nil && o.Constraint.IsSet() {
		return true
	}

	return false
}

// SetConstraint gets a reference to the given NullableHyperflexAppSettingConstraint and assigns it to the Constraint field.
func (o *HyperflexServerModelEntry) SetConstraint(v HyperflexAppSettingConstraint) {
	o.Constraint.Set(&v)
}

// SetConstraintNil sets the value for Constraint to be an explicit nil
func (o *HyperflexServerModelEntry) SetConstraintNil() {
	o.Constraint.Set(nil)
}

// UnsetConstraint ensures that no value is present for Constraint, not even an explicit nil
func (o *HyperflexServerModelEntry) UnsetConstraint() {
	o.Constraint.Unset()
}

func (o HyperflexServerModelEntry) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedHyperflexAbstractAppSetting, errHyperflexAbstractAppSetting := json.Marshal(o.HyperflexAbstractAppSetting)
	if errHyperflexAbstractAppSetting != nil {
		return []byte{}, errHyperflexAbstractAppSetting
	}
	errHyperflexAbstractAppSetting = json.Unmarshal([]byte(serializedHyperflexAbstractAppSetting), &toSerialize)
	if errHyperflexAbstractAppSetting != nil {
		return []byte{}, errHyperflexAbstractAppSetting
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Constraint.IsSet() {
		toSerialize["Constraint"] = o.Constraint.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *HyperflexServerModelEntry) UnmarshalJSON(bytes []byte) (err error) {
	type HyperflexServerModelEntryWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string                                `json:"ObjectType"`
		Constraint NullableHyperflexAppSettingConstraint `json:"Constraint,omitempty"`
	}

	varHyperflexServerModelEntryWithoutEmbeddedStruct := HyperflexServerModelEntryWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varHyperflexServerModelEntryWithoutEmbeddedStruct)
	if err == nil {
		varHyperflexServerModelEntry := _HyperflexServerModelEntry{}
		varHyperflexServerModelEntry.ClassId = varHyperflexServerModelEntryWithoutEmbeddedStruct.ClassId
		varHyperflexServerModelEntry.ObjectType = varHyperflexServerModelEntryWithoutEmbeddedStruct.ObjectType
		varHyperflexServerModelEntry.Constraint = varHyperflexServerModelEntryWithoutEmbeddedStruct.Constraint
		*o = HyperflexServerModelEntry(varHyperflexServerModelEntry)
	} else {
		return err
	}

	varHyperflexServerModelEntry := _HyperflexServerModelEntry{}

	err = json.Unmarshal(bytes, &varHyperflexServerModelEntry)
	if err == nil {
		o.HyperflexAbstractAppSetting = varHyperflexServerModelEntry.HyperflexAbstractAppSetting
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Constraint")

		// remove fields from embedded structs
		reflectHyperflexAbstractAppSetting := reflect.ValueOf(o.HyperflexAbstractAppSetting)
		for i := 0; i < reflectHyperflexAbstractAppSetting.Type().NumField(); i++ {
			t := reflectHyperflexAbstractAppSetting.Type().Field(i)

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

type NullableHyperflexServerModelEntry struct {
	value *HyperflexServerModelEntry
	isSet bool
}

func (v NullableHyperflexServerModelEntry) Get() *HyperflexServerModelEntry {
	return v.value
}

func (v *NullableHyperflexServerModelEntry) Set(val *HyperflexServerModelEntry) {
	v.value = val
	v.isSet = true
}

func (v NullableHyperflexServerModelEntry) IsSet() bool {
	return v.isSet
}

func (v *NullableHyperflexServerModelEntry) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHyperflexServerModelEntry(val *HyperflexServerModelEntry) *NullableHyperflexServerModelEntry {
	return &NullableHyperflexServerModelEntry{value: val, isSet: true}
}

func (v NullableHyperflexServerModelEntry) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHyperflexServerModelEntry) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
