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
	"time"
)

// TelemetryDruidGroupByResult A time series result based on a provided groupBy query. Result can contain multiple dimensions and values.
type TelemetryDruidGroupByResult struct {
	// The ISO 8601 timestamp.
	Timestamp *time.Time `json:"timestamp,omitempty"`
	// The version of the Druid GroupBy Engine used in query
	Version *string `json:"version,omitempty"`
	// Grouped aggregate dimension(s) with values
	Event                *map[string]interface{} `json:"event,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _TelemetryDruidGroupByResult TelemetryDruidGroupByResult

// NewTelemetryDruidGroupByResult instantiates a new TelemetryDruidGroupByResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTelemetryDruidGroupByResult() *TelemetryDruidGroupByResult {
	this := TelemetryDruidGroupByResult{}
	return &this
}

// NewTelemetryDruidGroupByResultWithDefaults instantiates a new TelemetryDruidGroupByResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTelemetryDruidGroupByResultWithDefaults() *TelemetryDruidGroupByResult {
	this := TelemetryDruidGroupByResult{}
	return &this
}

// GetTimestamp returns the Timestamp field value if set, zero value otherwise.
func (o *TelemetryDruidGroupByResult) GetTimestamp() time.Time {
	if o == nil || o.Timestamp == nil {
		var ret time.Time
		return ret
	}
	return *o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TelemetryDruidGroupByResult) GetTimestampOk() (*time.Time, bool) {
	if o == nil || o.Timestamp == nil {
		return nil, false
	}
	return o.Timestamp, true
}

// HasTimestamp returns a boolean if a field has been set.
func (o *TelemetryDruidGroupByResult) HasTimestamp() bool {
	if o != nil && o.Timestamp != nil {
		return true
	}

	return false
}

// SetTimestamp gets a reference to the given time.Time and assigns it to the Timestamp field.
func (o *TelemetryDruidGroupByResult) SetTimestamp(v time.Time) {
	o.Timestamp = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *TelemetryDruidGroupByResult) GetVersion() string {
	if o == nil || o.Version == nil {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TelemetryDruidGroupByResult) GetVersionOk() (*string, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *TelemetryDruidGroupByResult) HasVersion() bool {
	if o != nil && o.Version != nil {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *TelemetryDruidGroupByResult) SetVersion(v string) {
	o.Version = &v
}

// GetEvent returns the Event field value if set, zero value otherwise.
func (o *TelemetryDruidGroupByResult) GetEvent() map[string]interface{} {
	if o == nil || o.Event == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.Event
}

// GetEventOk returns a tuple with the Event field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TelemetryDruidGroupByResult) GetEventOk() (*map[string]interface{}, bool) {
	if o == nil || o.Event == nil {
		return nil, false
	}
	return o.Event, true
}

// HasEvent returns a boolean if a field has been set.
func (o *TelemetryDruidGroupByResult) HasEvent() bool {
	if o != nil && o.Event != nil {
		return true
	}

	return false
}

// SetEvent gets a reference to the given map[string]interface{} and assigns it to the Event field.
func (o *TelemetryDruidGroupByResult) SetEvent(v map[string]interface{}) {
	o.Event = &v
}

func (o TelemetryDruidGroupByResult) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Timestamp != nil {
		toSerialize["timestamp"] = o.Timestamp
	}
	if o.Version != nil {
		toSerialize["version"] = o.Version
	}
	if o.Event != nil {
		toSerialize["event"] = o.Event
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *TelemetryDruidGroupByResult) UnmarshalJSON(bytes []byte) (err error) {
	varTelemetryDruidGroupByResult := _TelemetryDruidGroupByResult{}

	if err = json.Unmarshal(bytes, &varTelemetryDruidGroupByResult); err == nil {
		*o = TelemetryDruidGroupByResult(varTelemetryDruidGroupByResult)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "timestamp")
		delete(additionalProperties, "version")
		delete(additionalProperties, "event")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTelemetryDruidGroupByResult struct {
	value *TelemetryDruidGroupByResult
	isSet bool
}

func (v NullableTelemetryDruidGroupByResult) Get() *TelemetryDruidGroupByResult {
	return v.value
}

func (v *NullableTelemetryDruidGroupByResult) Set(val *TelemetryDruidGroupByResult) {
	v.value = val
	v.isSet = true
}

func (v NullableTelemetryDruidGroupByResult) IsSet() bool {
	return v.isSet
}

func (v *NullableTelemetryDruidGroupByResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTelemetryDruidGroupByResult(val *TelemetryDruidGroupByResult) *NullableTelemetryDruidGroupByResult {
	return &NullableTelemetryDruidGroupByResult{value: val, isSet: true}
}

func (v NullableTelemetryDruidGroupByResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTelemetryDruidGroupByResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
