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

// ConnectorStreamAcknowledge Acknowledge a set of messages read from the device, on receipt device will drop the acknowledged messages from its cache.
type ConnectorStreamAcknowledge struct {
	ConnectorStreamMessage
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// The latest message sequence processed in the cloud. Device connector will drop all messages up to this sequence from its cache.
	AckSequence          *int64 `json:"AckSequence,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ConnectorStreamAcknowledge ConnectorStreamAcknowledge

// NewConnectorStreamAcknowledge instantiates a new ConnectorStreamAcknowledge object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConnectorStreamAcknowledge(classId string, objectType string) *ConnectorStreamAcknowledge {
	this := ConnectorStreamAcknowledge{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewConnectorStreamAcknowledgeWithDefaults instantiates a new ConnectorStreamAcknowledge object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConnectorStreamAcknowledgeWithDefaults() *ConnectorStreamAcknowledge {
	this := ConnectorStreamAcknowledge{}
	var classId string = "connector.StreamAcknowledge"
	this.ClassId = classId
	var objectType string = "connector.StreamAcknowledge"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *ConnectorStreamAcknowledge) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *ConnectorStreamAcknowledge) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *ConnectorStreamAcknowledge) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *ConnectorStreamAcknowledge) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *ConnectorStreamAcknowledge) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *ConnectorStreamAcknowledge) SetObjectType(v string) {
	o.ObjectType = v
}

// GetAckSequence returns the AckSequence field value if set, zero value otherwise.
func (o *ConnectorStreamAcknowledge) GetAckSequence() int64 {
	if o == nil || o.AckSequence == nil {
		var ret int64
		return ret
	}
	return *o.AckSequence
}

// GetAckSequenceOk returns a tuple with the AckSequence field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectorStreamAcknowledge) GetAckSequenceOk() (*int64, bool) {
	if o == nil || o.AckSequence == nil {
		return nil, false
	}
	return o.AckSequence, true
}

// HasAckSequence returns a boolean if a field has been set.
func (o *ConnectorStreamAcknowledge) HasAckSequence() bool {
	if o != nil && o.AckSequence != nil {
		return true
	}

	return false
}

// SetAckSequence gets a reference to the given int64 and assigns it to the AckSequence field.
func (o *ConnectorStreamAcknowledge) SetAckSequence(v int64) {
	o.AckSequence = &v
}

func (o ConnectorStreamAcknowledge) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedConnectorStreamMessage, errConnectorStreamMessage := json.Marshal(o.ConnectorStreamMessage)
	if errConnectorStreamMessage != nil {
		return []byte{}, errConnectorStreamMessage
	}
	errConnectorStreamMessage = json.Unmarshal([]byte(serializedConnectorStreamMessage), &toSerialize)
	if errConnectorStreamMessage != nil {
		return []byte{}, errConnectorStreamMessage
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.AckSequence != nil {
		toSerialize["AckSequence"] = o.AckSequence
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ConnectorStreamAcknowledge) UnmarshalJSON(bytes []byte) (err error) {
	type ConnectorStreamAcknowledgeWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// The latest message sequence processed in the cloud. Device connector will drop all messages up to this sequence from its cache.
		AckSequence *int64 `json:"AckSequence,omitempty"`
	}

	varConnectorStreamAcknowledgeWithoutEmbeddedStruct := ConnectorStreamAcknowledgeWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varConnectorStreamAcknowledgeWithoutEmbeddedStruct)
	if err == nil {
		varConnectorStreamAcknowledge := _ConnectorStreamAcknowledge{}
		varConnectorStreamAcknowledge.ClassId = varConnectorStreamAcknowledgeWithoutEmbeddedStruct.ClassId
		varConnectorStreamAcknowledge.ObjectType = varConnectorStreamAcknowledgeWithoutEmbeddedStruct.ObjectType
		varConnectorStreamAcknowledge.AckSequence = varConnectorStreamAcknowledgeWithoutEmbeddedStruct.AckSequence
		*o = ConnectorStreamAcknowledge(varConnectorStreamAcknowledge)
	} else {
		return err
	}

	varConnectorStreamAcknowledge := _ConnectorStreamAcknowledge{}

	err = json.Unmarshal(bytes, &varConnectorStreamAcknowledge)
	if err == nil {
		o.ConnectorStreamMessage = varConnectorStreamAcknowledge.ConnectorStreamMessage
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "AckSequence")

		// remove fields from embedded structs
		reflectConnectorStreamMessage := reflect.ValueOf(o.ConnectorStreamMessage)
		for i := 0; i < reflectConnectorStreamMessage.Type().NumField(); i++ {
			t := reflectConnectorStreamMessage.Type().Field(i)

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

type NullableConnectorStreamAcknowledge struct {
	value *ConnectorStreamAcknowledge
	isSet bool
}

func (v NullableConnectorStreamAcknowledge) Get() *ConnectorStreamAcknowledge {
	return v.value
}

func (v *NullableConnectorStreamAcknowledge) Set(val *ConnectorStreamAcknowledge) {
	v.value = val
	v.isSet = true
}

func (v NullableConnectorStreamAcknowledge) IsSet() bool {
	return v.isSet
}

func (v *NullableConnectorStreamAcknowledge) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConnectorStreamAcknowledge(val *ConnectorStreamAcknowledge) *NullableConnectorStreamAcknowledge {
	return &NullableConnectorStreamAcknowledge{value: val, isSet: true}
}

func (v NullableConnectorStreamAcknowledge) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConnectorStreamAcknowledge) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
