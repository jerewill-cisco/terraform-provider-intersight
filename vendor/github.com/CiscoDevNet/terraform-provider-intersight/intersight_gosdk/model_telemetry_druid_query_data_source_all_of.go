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

// TelemetryDruidQueryDataSourceAllOf struct for TelemetryDruidQueryDataSourceAllOf
type TelemetryDruidQueryDataSourceAllOf struct {
	Query                TelemetryDruidGroupByRequest `json:"query"`
	AdditionalProperties map[string]interface{}
}

type _TelemetryDruidQueryDataSourceAllOf TelemetryDruidQueryDataSourceAllOf

// NewTelemetryDruidQueryDataSourceAllOf instantiates a new TelemetryDruidQueryDataSourceAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTelemetryDruidQueryDataSourceAllOf(query TelemetryDruidGroupByRequest) *TelemetryDruidQueryDataSourceAllOf {
	this := TelemetryDruidQueryDataSourceAllOf{}
	this.Query = query
	return &this
}

// NewTelemetryDruidQueryDataSourceAllOfWithDefaults instantiates a new TelemetryDruidQueryDataSourceAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTelemetryDruidQueryDataSourceAllOfWithDefaults() *TelemetryDruidQueryDataSourceAllOf {
	this := TelemetryDruidQueryDataSourceAllOf{}
	return &this
}

// GetQuery returns the Query field value
func (o *TelemetryDruidQueryDataSourceAllOf) GetQuery() TelemetryDruidGroupByRequest {
	if o == nil {
		var ret TelemetryDruidGroupByRequest
		return ret
	}

	return o.Query
}

// GetQueryOk returns a tuple with the Query field value
// and a boolean to check if the value has been set.
func (o *TelemetryDruidQueryDataSourceAllOf) GetQueryOk() (*TelemetryDruidGroupByRequest, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Query, true
}

// SetQuery sets field value
func (o *TelemetryDruidQueryDataSourceAllOf) SetQuery(v TelemetryDruidGroupByRequest) {
	o.Query = v
}

func (o TelemetryDruidQueryDataSourceAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["query"] = o.Query
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *TelemetryDruidQueryDataSourceAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varTelemetryDruidQueryDataSourceAllOf := _TelemetryDruidQueryDataSourceAllOf{}

	if err = json.Unmarshal(bytes, &varTelemetryDruidQueryDataSourceAllOf); err == nil {
		*o = TelemetryDruidQueryDataSourceAllOf(varTelemetryDruidQueryDataSourceAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "query")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTelemetryDruidQueryDataSourceAllOf struct {
	value *TelemetryDruidQueryDataSourceAllOf
	isSet bool
}

func (v NullableTelemetryDruidQueryDataSourceAllOf) Get() *TelemetryDruidQueryDataSourceAllOf {
	return v.value
}

func (v *NullableTelemetryDruidQueryDataSourceAllOf) Set(val *TelemetryDruidQueryDataSourceAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableTelemetryDruidQueryDataSourceAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableTelemetryDruidQueryDataSourceAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTelemetryDruidQueryDataSourceAllOf(val *TelemetryDruidQueryDataSourceAllOf) *NullableTelemetryDruidQueryDataSourceAllOf {
	return &NullableTelemetryDruidQueryDataSourceAllOf{value: val, isSet: true}
}

func (v NullableTelemetryDruidQueryDataSourceAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTelemetryDruidQueryDataSourceAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
