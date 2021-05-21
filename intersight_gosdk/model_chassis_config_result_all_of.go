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

// ChassisConfigResultAllOf Definition of the list of properties defined in 'chassis.ConfigResult', excluding properties defined in parent classes.
type ChassisConfigResultAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string                      `json:"ObjectType"`
	Profile    *ChassisProfileRelationship `json:"Profile,omitempty"`
	// An array of relationships to chassisConfigResultEntry resources.
	ResultEntries        []ChassisConfigResultEntryRelationship `json:"ResultEntries,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ChassisConfigResultAllOf ChassisConfigResultAllOf

// NewChassisConfigResultAllOf instantiates a new ChassisConfigResultAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewChassisConfigResultAllOf(classId string, objectType string) *ChassisConfigResultAllOf {
	this := ChassisConfigResultAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewChassisConfigResultAllOfWithDefaults instantiates a new ChassisConfigResultAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewChassisConfigResultAllOfWithDefaults() *ChassisConfigResultAllOf {
	this := ChassisConfigResultAllOf{}
	var classId string = "chassis.ConfigResult"
	this.ClassId = classId
	var objectType string = "chassis.ConfigResult"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *ChassisConfigResultAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *ChassisConfigResultAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *ChassisConfigResultAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *ChassisConfigResultAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *ChassisConfigResultAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *ChassisConfigResultAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetProfile returns the Profile field value if set, zero value otherwise.
func (o *ChassisConfigResultAllOf) GetProfile() ChassisProfileRelationship {
	if o == nil || o.Profile == nil {
		var ret ChassisProfileRelationship
		return ret
	}
	return *o.Profile
}

// GetProfileOk returns a tuple with the Profile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChassisConfigResultAllOf) GetProfileOk() (*ChassisProfileRelationship, bool) {
	if o == nil || o.Profile == nil {
		return nil, false
	}
	return o.Profile, true
}

// HasProfile returns a boolean if a field has been set.
func (o *ChassisConfigResultAllOf) HasProfile() bool {
	if o != nil && o.Profile != nil {
		return true
	}

	return false
}

// SetProfile gets a reference to the given ChassisProfileRelationship and assigns it to the Profile field.
func (o *ChassisConfigResultAllOf) SetProfile(v ChassisProfileRelationship) {
	o.Profile = &v
}

// GetResultEntries returns the ResultEntries field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ChassisConfigResultAllOf) GetResultEntries() []ChassisConfigResultEntryRelationship {
	if o == nil {
		var ret []ChassisConfigResultEntryRelationship
		return ret
	}
	return o.ResultEntries
}

// GetResultEntriesOk returns a tuple with the ResultEntries field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ChassisConfigResultAllOf) GetResultEntriesOk() (*[]ChassisConfigResultEntryRelationship, bool) {
	if o == nil || o.ResultEntries == nil {
		return nil, false
	}
	return &o.ResultEntries, true
}

// HasResultEntries returns a boolean if a field has been set.
func (o *ChassisConfigResultAllOf) HasResultEntries() bool {
	if o != nil && o.ResultEntries != nil {
		return true
	}

	return false
}

// SetResultEntries gets a reference to the given []ChassisConfigResultEntryRelationship and assigns it to the ResultEntries field.
func (o *ChassisConfigResultAllOf) SetResultEntries(v []ChassisConfigResultEntryRelationship) {
	o.ResultEntries = v
}

func (o ChassisConfigResultAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Profile != nil {
		toSerialize["Profile"] = o.Profile
	}
	if o.ResultEntries != nil {
		toSerialize["ResultEntries"] = o.ResultEntries
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ChassisConfigResultAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varChassisConfigResultAllOf := _ChassisConfigResultAllOf{}

	if err = json.Unmarshal(bytes, &varChassisConfigResultAllOf); err == nil {
		*o = ChassisConfigResultAllOf(varChassisConfigResultAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Profile")
		delete(additionalProperties, "ResultEntries")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableChassisConfigResultAllOf struct {
	value *ChassisConfigResultAllOf
	isSet bool
}

func (v NullableChassisConfigResultAllOf) Get() *ChassisConfigResultAllOf {
	return v.value
}

func (v *NullableChassisConfigResultAllOf) Set(val *ChassisConfigResultAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableChassisConfigResultAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableChassisConfigResultAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableChassisConfigResultAllOf(val *ChassisConfigResultAllOf) *NullableChassisConfigResultAllOf {
	return &NullableChassisConfigResultAllOf{value: val, isSet: true}
}

func (v NullableChassisConfigResultAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableChassisConfigResultAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
