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

// HyperflexReplicationClusterReferenceToScheduleAllOf Definition of the list of properties defined in 'hyperflex.ReplicationClusterReferenceToSchedule', excluding properties defined in parent classes.
type HyperflexReplicationClusterReferenceToScheduleAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType                   string                               `json:"ObjectType"`
	Schedule                     NullableHyperflexReplicationSchedule `json:"Schedule,omitempty"`
	TargetClusterEntityReference NullableHyperflexEntityReference     `json:"TargetClusterEntityReference,omitempty"`
	AdditionalProperties         map[string]interface{}
}

type _HyperflexReplicationClusterReferenceToScheduleAllOf HyperflexReplicationClusterReferenceToScheduleAllOf

// NewHyperflexReplicationClusterReferenceToScheduleAllOf instantiates a new HyperflexReplicationClusterReferenceToScheduleAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHyperflexReplicationClusterReferenceToScheduleAllOf(classId string, objectType string) *HyperflexReplicationClusterReferenceToScheduleAllOf {
	this := HyperflexReplicationClusterReferenceToScheduleAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewHyperflexReplicationClusterReferenceToScheduleAllOfWithDefaults instantiates a new HyperflexReplicationClusterReferenceToScheduleAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHyperflexReplicationClusterReferenceToScheduleAllOfWithDefaults() *HyperflexReplicationClusterReferenceToScheduleAllOf {
	this := HyperflexReplicationClusterReferenceToScheduleAllOf{}
	var classId string = "hyperflex.ReplicationClusterReferenceToSchedule"
	this.ClassId = classId
	var objectType string = "hyperflex.ReplicationClusterReferenceToSchedule"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *HyperflexReplicationClusterReferenceToScheduleAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *HyperflexReplicationClusterReferenceToScheduleAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *HyperflexReplicationClusterReferenceToScheduleAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *HyperflexReplicationClusterReferenceToScheduleAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *HyperflexReplicationClusterReferenceToScheduleAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *HyperflexReplicationClusterReferenceToScheduleAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetSchedule returns the Schedule field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HyperflexReplicationClusterReferenceToScheduleAllOf) GetSchedule() HyperflexReplicationSchedule {
	if o == nil || o.Schedule.Get() == nil {
		var ret HyperflexReplicationSchedule
		return ret
	}
	return *o.Schedule.Get()
}

// GetScheduleOk returns a tuple with the Schedule field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HyperflexReplicationClusterReferenceToScheduleAllOf) GetScheduleOk() (*HyperflexReplicationSchedule, bool) {
	if o == nil {
		return nil, false
	}
	return o.Schedule.Get(), o.Schedule.IsSet()
}

// HasSchedule returns a boolean if a field has been set.
func (o *HyperflexReplicationClusterReferenceToScheduleAllOf) HasSchedule() bool {
	if o != nil && o.Schedule.IsSet() {
		return true
	}

	return false
}

// SetSchedule gets a reference to the given NullableHyperflexReplicationSchedule and assigns it to the Schedule field.
func (o *HyperflexReplicationClusterReferenceToScheduleAllOf) SetSchedule(v HyperflexReplicationSchedule) {
	o.Schedule.Set(&v)
}

// SetScheduleNil sets the value for Schedule to be an explicit nil
func (o *HyperflexReplicationClusterReferenceToScheduleAllOf) SetScheduleNil() {
	o.Schedule.Set(nil)
}

// UnsetSchedule ensures that no value is present for Schedule, not even an explicit nil
func (o *HyperflexReplicationClusterReferenceToScheduleAllOf) UnsetSchedule() {
	o.Schedule.Unset()
}

// GetTargetClusterEntityReference returns the TargetClusterEntityReference field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HyperflexReplicationClusterReferenceToScheduleAllOf) GetTargetClusterEntityReference() HyperflexEntityReference {
	if o == nil || o.TargetClusterEntityReference.Get() == nil {
		var ret HyperflexEntityReference
		return ret
	}
	return *o.TargetClusterEntityReference.Get()
}

// GetTargetClusterEntityReferenceOk returns a tuple with the TargetClusterEntityReference field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HyperflexReplicationClusterReferenceToScheduleAllOf) GetTargetClusterEntityReferenceOk() (*HyperflexEntityReference, bool) {
	if o == nil {
		return nil, false
	}
	return o.TargetClusterEntityReference.Get(), o.TargetClusterEntityReference.IsSet()
}

// HasTargetClusterEntityReference returns a boolean if a field has been set.
func (o *HyperflexReplicationClusterReferenceToScheduleAllOf) HasTargetClusterEntityReference() bool {
	if o != nil && o.TargetClusterEntityReference.IsSet() {
		return true
	}

	return false
}

// SetTargetClusterEntityReference gets a reference to the given NullableHyperflexEntityReference and assigns it to the TargetClusterEntityReference field.
func (o *HyperflexReplicationClusterReferenceToScheduleAllOf) SetTargetClusterEntityReference(v HyperflexEntityReference) {
	o.TargetClusterEntityReference.Set(&v)
}

// SetTargetClusterEntityReferenceNil sets the value for TargetClusterEntityReference to be an explicit nil
func (o *HyperflexReplicationClusterReferenceToScheduleAllOf) SetTargetClusterEntityReferenceNil() {
	o.TargetClusterEntityReference.Set(nil)
}

// UnsetTargetClusterEntityReference ensures that no value is present for TargetClusterEntityReference, not even an explicit nil
func (o *HyperflexReplicationClusterReferenceToScheduleAllOf) UnsetTargetClusterEntityReference() {
	o.TargetClusterEntityReference.Unset()
}

func (o HyperflexReplicationClusterReferenceToScheduleAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Schedule.IsSet() {
		toSerialize["Schedule"] = o.Schedule.Get()
	}
	if o.TargetClusterEntityReference.IsSet() {
		toSerialize["TargetClusterEntityReference"] = o.TargetClusterEntityReference.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *HyperflexReplicationClusterReferenceToScheduleAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varHyperflexReplicationClusterReferenceToScheduleAllOf := _HyperflexReplicationClusterReferenceToScheduleAllOf{}

	if err = json.Unmarshal(bytes, &varHyperflexReplicationClusterReferenceToScheduleAllOf); err == nil {
		*o = HyperflexReplicationClusterReferenceToScheduleAllOf(varHyperflexReplicationClusterReferenceToScheduleAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Schedule")
		delete(additionalProperties, "TargetClusterEntityReference")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableHyperflexReplicationClusterReferenceToScheduleAllOf struct {
	value *HyperflexReplicationClusterReferenceToScheduleAllOf
	isSet bool
}

func (v NullableHyperflexReplicationClusterReferenceToScheduleAllOf) Get() *HyperflexReplicationClusterReferenceToScheduleAllOf {
	return v.value
}

func (v *NullableHyperflexReplicationClusterReferenceToScheduleAllOf) Set(val *HyperflexReplicationClusterReferenceToScheduleAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableHyperflexReplicationClusterReferenceToScheduleAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableHyperflexReplicationClusterReferenceToScheduleAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHyperflexReplicationClusterReferenceToScheduleAllOf(val *HyperflexReplicationClusterReferenceToScheduleAllOf) *NullableHyperflexReplicationClusterReferenceToScheduleAllOf {
	return &NullableHyperflexReplicationClusterReferenceToScheduleAllOf{value: val, isSet: true}
}

func (v NullableHyperflexReplicationClusterReferenceToScheduleAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHyperflexReplicationClusterReferenceToScheduleAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
