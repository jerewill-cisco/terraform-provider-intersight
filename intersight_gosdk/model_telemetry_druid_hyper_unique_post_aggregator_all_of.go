/*
Cisco Intersight

Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document.

API version: 1.0.9-4663
Contact: intersight@cisco.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package intersight

import (
	"encoding/json"
)

// TelemetryDruidHyperUniquePostAggregatorAllOf struct for TelemetryDruidHyperUniquePostAggregatorAllOf
type TelemetryDruidHyperUniquePostAggregatorAllOf struct {
	// Output name for the post-aggregator.
	Name *string `json:"name,omitempty"`
	// The name field value of the hyperUnique aggregator.
	FieldName            *string `json:"fieldName,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _TelemetryDruidHyperUniquePostAggregatorAllOf TelemetryDruidHyperUniquePostAggregatorAllOf

// NewTelemetryDruidHyperUniquePostAggregatorAllOf instantiates a new TelemetryDruidHyperUniquePostAggregatorAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTelemetryDruidHyperUniquePostAggregatorAllOf() *TelemetryDruidHyperUniquePostAggregatorAllOf {
	this := TelemetryDruidHyperUniquePostAggregatorAllOf{}
	return &this
}

// NewTelemetryDruidHyperUniquePostAggregatorAllOfWithDefaults instantiates a new TelemetryDruidHyperUniquePostAggregatorAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTelemetryDruidHyperUniquePostAggregatorAllOfWithDefaults() *TelemetryDruidHyperUniquePostAggregatorAllOf {
	this := TelemetryDruidHyperUniquePostAggregatorAllOf{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *TelemetryDruidHyperUniquePostAggregatorAllOf) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TelemetryDruidHyperUniquePostAggregatorAllOf) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *TelemetryDruidHyperUniquePostAggregatorAllOf) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *TelemetryDruidHyperUniquePostAggregatorAllOf) SetName(v string) {
	o.Name = &v
}

// GetFieldName returns the FieldName field value if set, zero value otherwise.
func (o *TelemetryDruidHyperUniquePostAggregatorAllOf) GetFieldName() string {
	if o == nil || o.FieldName == nil {
		var ret string
		return ret
	}
	return *o.FieldName
}

// GetFieldNameOk returns a tuple with the FieldName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TelemetryDruidHyperUniquePostAggregatorAllOf) GetFieldNameOk() (*string, bool) {
	if o == nil || o.FieldName == nil {
		return nil, false
	}
	return o.FieldName, true
}

// HasFieldName returns a boolean if a field has been set.
func (o *TelemetryDruidHyperUniquePostAggregatorAllOf) HasFieldName() bool {
	if o != nil && o.FieldName != nil {
		return true
	}

	return false
}

// SetFieldName gets a reference to the given string and assigns it to the FieldName field.
func (o *TelemetryDruidHyperUniquePostAggregatorAllOf) SetFieldName(v string) {
	o.FieldName = &v
}

func (o TelemetryDruidHyperUniquePostAggregatorAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.FieldName != nil {
		toSerialize["fieldName"] = o.FieldName
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *TelemetryDruidHyperUniquePostAggregatorAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varTelemetryDruidHyperUniquePostAggregatorAllOf := _TelemetryDruidHyperUniquePostAggregatorAllOf{}

	if err = json.Unmarshal(bytes, &varTelemetryDruidHyperUniquePostAggregatorAllOf); err == nil {
		*o = TelemetryDruidHyperUniquePostAggregatorAllOf(varTelemetryDruidHyperUniquePostAggregatorAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "name")
		delete(additionalProperties, "fieldName")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTelemetryDruidHyperUniquePostAggregatorAllOf struct {
	value *TelemetryDruidHyperUniquePostAggregatorAllOf
	isSet bool
}

func (v NullableTelemetryDruidHyperUniquePostAggregatorAllOf) Get() *TelemetryDruidHyperUniquePostAggregatorAllOf {
	return v.value
}

func (v *NullableTelemetryDruidHyperUniquePostAggregatorAllOf) Set(val *TelemetryDruidHyperUniquePostAggregatorAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableTelemetryDruidHyperUniquePostAggregatorAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableTelemetryDruidHyperUniquePostAggregatorAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTelemetryDruidHyperUniquePostAggregatorAllOf(val *TelemetryDruidHyperUniquePostAggregatorAllOf) *NullableTelemetryDruidHyperUniquePostAggregatorAllOf {
	return &NullableTelemetryDruidHyperUniquePostAggregatorAllOf{value: val, isSet: true}
}

func (v NullableTelemetryDruidHyperUniquePostAggregatorAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTelemetryDruidHyperUniquePostAggregatorAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
