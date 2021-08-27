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
)

// TelemetryDruidMinMaxAggregatorAllOf struct for TelemetryDruidMinMaxAggregatorAllOf
type TelemetryDruidMinMaxAggregatorAllOf struct {
	// Output name for the min/max value.
	Name string `json:"name"`
	// Name of the metric column.
	FieldName            string `json:"fieldName"`
	AdditionalProperties map[string]interface{}
}

type _TelemetryDruidMinMaxAggregatorAllOf TelemetryDruidMinMaxAggregatorAllOf

// NewTelemetryDruidMinMaxAggregatorAllOf instantiates a new TelemetryDruidMinMaxAggregatorAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTelemetryDruidMinMaxAggregatorAllOf(name string, fieldName string) *TelemetryDruidMinMaxAggregatorAllOf {
	this := TelemetryDruidMinMaxAggregatorAllOf{}
	this.Name = name
	this.FieldName = fieldName
	return &this
}

// NewTelemetryDruidMinMaxAggregatorAllOfWithDefaults instantiates a new TelemetryDruidMinMaxAggregatorAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTelemetryDruidMinMaxAggregatorAllOfWithDefaults() *TelemetryDruidMinMaxAggregatorAllOf {
	this := TelemetryDruidMinMaxAggregatorAllOf{}
	return &this
}

// GetName returns the Name field value
func (o *TelemetryDruidMinMaxAggregatorAllOf) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *TelemetryDruidMinMaxAggregatorAllOf) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *TelemetryDruidMinMaxAggregatorAllOf) SetName(v string) {
	o.Name = v
}

// GetFieldName returns the FieldName field value
func (o *TelemetryDruidMinMaxAggregatorAllOf) GetFieldName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FieldName
}

// GetFieldNameOk returns a tuple with the FieldName field value
// and a boolean to check if the value has been set.
func (o *TelemetryDruidMinMaxAggregatorAllOf) GetFieldNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FieldName, true
}

// SetFieldName sets field value
func (o *TelemetryDruidMinMaxAggregatorAllOf) SetFieldName(v string) {
	o.FieldName = v
}

func (o TelemetryDruidMinMaxAggregatorAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["fieldName"] = o.FieldName
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *TelemetryDruidMinMaxAggregatorAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varTelemetryDruidMinMaxAggregatorAllOf := _TelemetryDruidMinMaxAggregatorAllOf{}

	if err = json.Unmarshal(bytes, &varTelemetryDruidMinMaxAggregatorAllOf); err == nil {
		*o = TelemetryDruidMinMaxAggregatorAllOf(varTelemetryDruidMinMaxAggregatorAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "name")
		delete(additionalProperties, "fieldName")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTelemetryDruidMinMaxAggregatorAllOf struct {
	value *TelemetryDruidMinMaxAggregatorAllOf
	isSet bool
}

func (v NullableTelemetryDruidMinMaxAggregatorAllOf) Get() *TelemetryDruidMinMaxAggregatorAllOf {
	return v.value
}

func (v *NullableTelemetryDruidMinMaxAggregatorAllOf) Set(val *TelemetryDruidMinMaxAggregatorAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableTelemetryDruidMinMaxAggregatorAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableTelemetryDruidMinMaxAggregatorAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTelemetryDruidMinMaxAggregatorAllOf(val *TelemetryDruidMinMaxAggregatorAllOf) *NullableTelemetryDruidMinMaxAggregatorAllOf {
	return &NullableTelemetryDruidMinMaxAggregatorAllOf{value: val, isSet: true}
}

func (v NullableTelemetryDruidMinMaxAggregatorAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTelemetryDruidMinMaxAggregatorAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
