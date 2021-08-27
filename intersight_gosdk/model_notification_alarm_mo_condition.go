/*
Cisco Intersight

Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document. This document was created on 2021-08-10T21:48:06Z.

API version: 1.0.9-4430
Contact: intersight@cisco.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package intersight

import (
	"encoding/json"
	"reflect"
	"strings"
)

// NotificationAlarmMoCondition This condition type gives the ability to subscribe to the cond.Alarm ObjectType and provide a list of Severities you're interested in.
type NotificationAlarmMoCondition struct {
	NotificationAbstractMoCondition
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType           string   `json:"ObjectType"`
	Severity             []string `json:"Severity,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _NotificationAlarmMoCondition NotificationAlarmMoCondition

// NewNotificationAlarmMoCondition instantiates a new NotificationAlarmMoCondition object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNotificationAlarmMoCondition(classId string, objectType string) *NotificationAlarmMoCondition {
	this := NotificationAlarmMoCondition{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewNotificationAlarmMoConditionWithDefaults instantiates a new NotificationAlarmMoCondition object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNotificationAlarmMoConditionWithDefaults() *NotificationAlarmMoCondition {
	this := NotificationAlarmMoCondition{}
	var classId string = "notification.AlarmMoCondition"
	this.ClassId = classId
	var objectType string = "notification.AlarmMoCondition"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *NotificationAlarmMoCondition) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *NotificationAlarmMoCondition) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *NotificationAlarmMoCondition) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *NotificationAlarmMoCondition) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *NotificationAlarmMoCondition) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *NotificationAlarmMoCondition) SetObjectType(v string) {
	o.ObjectType = v
}

// GetSeverity returns the Severity field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *NotificationAlarmMoCondition) GetSeverity() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Severity
}

// GetSeverityOk returns a tuple with the Severity field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NotificationAlarmMoCondition) GetSeverityOk() (*[]string, bool) {
	if o == nil || o.Severity == nil {
		return nil, false
	}
	return &o.Severity, true
}

// HasSeverity returns a boolean if a field has been set.
func (o *NotificationAlarmMoCondition) HasSeverity() bool {
	if o != nil && o.Severity != nil {
		return true
	}

	return false
}

// SetSeverity gets a reference to the given []string and assigns it to the Severity field.
func (o *NotificationAlarmMoCondition) SetSeverity(v []string) {
	o.Severity = v
}

func (o NotificationAlarmMoCondition) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedNotificationAbstractMoCondition, errNotificationAbstractMoCondition := json.Marshal(o.NotificationAbstractMoCondition)
	if errNotificationAbstractMoCondition != nil {
		return []byte{}, errNotificationAbstractMoCondition
	}
	errNotificationAbstractMoCondition = json.Unmarshal([]byte(serializedNotificationAbstractMoCondition), &toSerialize)
	if errNotificationAbstractMoCondition != nil {
		return []byte{}, errNotificationAbstractMoCondition
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Severity != nil {
		toSerialize["Severity"] = o.Severity
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *NotificationAlarmMoCondition) UnmarshalJSON(bytes []byte) (err error) {
	type NotificationAlarmMoConditionWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string   `json:"ObjectType"`
		Severity   []string `json:"Severity,omitempty"`
	}

	varNotificationAlarmMoConditionWithoutEmbeddedStruct := NotificationAlarmMoConditionWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varNotificationAlarmMoConditionWithoutEmbeddedStruct)
	if err == nil {
		varNotificationAlarmMoCondition := _NotificationAlarmMoCondition{}
		varNotificationAlarmMoCondition.ClassId = varNotificationAlarmMoConditionWithoutEmbeddedStruct.ClassId
		varNotificationAlarmMoCondition.ObjectType = varNotificationAlarmMoConditionWithoutEmbeddedStruct.ObjectType
		varNotificationAlarmMoCondition.Severity = varNotificationAlarmMoConditionWithoutEmbeddedStruct.Severity
		*o = NotificationAlarmMoCondition(varNotificationAlarmMoCondition)
	} else {
		return err
	}

	varNotificationAlarmMoCondition := _NotificationAlarmMoCondition{}

	err = json.Unmarshal(bytes, &varNotificationAlarmMoCondition)
	if err == nil {
		o.NotificationAbstractMoCondition = varNotificationAlarmMoCondition.NotificationAbstractMoCondition
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Severity")

		// remove fields from embedded structs
		reflectNotificationAbstractMoCondition := reflect.ValueOf(o.NotificationAbstractMoCondition)
		for i := 0; i < reflectNotificationAbstractMoCondition.Type().NumField(); i++ {
			t := reflectNotificationAbstractMoCondition.Type().Field(i)

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

type NullableNotificationAlarmMoCondition struct {
	value *NotificationAlarmMoCondition
	isSet bool
}

func (v NullableNotificationAlarmMoCondition) Get() *NotificationAlarmMoCondition {
	return v.value
}

func (v *NullableNotificationAlarmMoCondition) Set(val *NotificationAlarmMoCondition) {
	v.value = val
	v.isSet = true
}

func (v NullableNotificationAlarmMoCondition) IsSet() bool {
	return v.isSet
}

func (v *NullableNotificationAlarmMoCondition) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNotificationAlarmMoCondition(val *NotificationAlarmMoCondition) *NullableNotificationAlarmMoCondition {
	return &NullableNotificationAlarmMoCondition{value: val, isSet: true}
}

func (v NullableNotificationAlarmMoCondition) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNotificationAlarmMoCondition) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
