/*
Cisco Intersight

Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document.

API version: 1.0.9-5208
Contact: intersight@cisco.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package intersight

import (
	"encoding/json"
)

// TelemetryDruidInvertedTopNMetricSpecAllOf struct for TelemetryDruidInvertedTopNMetricSpecAllOf
type TelemetryDruidInvertedTopNMetricSpecAllOf struct {
	Metric               TelemetryDruidTopNMetricSpec `json:"metric"`
	AdditionalProperties map[string]interface{}
}

type _TelemetryDruidInvertedTopNMetricSpecAllOf TelemetryDruidInvertedTopNMetricSpecAllOf

// NewTelemetryDruidInvertedTopNMetricSpecAllOf instantiates a new TelemetryDruidInvertedTopNMetricSpecAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTelemetryDruidInvertedTopNMetricSpecAllOf(metric TelemetryDruidTopNMetricSpec) *TelemetryDruidInvertedTopNMetricSpecAllOf {
	this := TelemetryDruidInvertedTopNMetricSpecAllOf{}
	this.Metric = metric
	return &this
}

// NewTelemetryDruidInvertedTopNMetricSpecAllOfWithDefaults instantiates a new TelemetryDruidInvertedTopNMetricSpecAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTelemetryDruidInvertedTopNMetricSpecAllOfWithDefaults() *TelemetryDruidInvertedTopNMetricSpecAllOf {
	this := TelemetryDruidInvertedTopNMetricSpecAllOf{}
	return &this
}

// GetMetric returns the Metric field value
func (o *TelemetryDruidInvertedTopNMetricSpecAllOf) GetMetric() TelemetryDruidTopNMetricSpec {
	if o == nil {
		var ret TelemetryDruidTopNMetricSpec
		return ret
	}

	return o.Metric
}

// GetMetricOk returns a tuple with the Metric field value
// and a boolean to check if the value has been set.
func (o *TelemetryDruidInvertedTopNMetricSpecAllOf) GetMetricOk() (*TelemetryDruidTopNMetricSpec, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Metric, true
}

// SetMetric sets field value
func (o *TelemetryDruidInvertedTopNMetricSpecAllOf) SetMetric(v TelemetryDruidTopNMetricSpec) {
	o.Metric = v
}

func (o TelemetryDruidInvertedTopNMetricSpecAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["metric"] = o.Metric
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *TelemetryDruidInvertedTopNMetricSpecAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varTelemetryDruidInvertedTopNMetricSpecAllOf := _TelemetryDruidInvertedTopNMetricSpecAllOf{}

	if err = json.Unmarshal(bytes, &varTelemetryDruidInvertedTopNMetricSpecAllOf); err == nil {
		*o = TelemetryDruidInvertedTopNMetricSpecAllOf(varTelemetryDruidInvertedTopNMetricSpecAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "metric")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTelemetryDruidInvertedTopNMetricSpecAllOf struct {
	value *TelemetryDruidInvertedTopNMetricSpecAllOf
	isSet bool
}

func (v NullableTelemetryDruidInvertedTopNMetricSpecAllOf) Get() *TelemetryDruidInvertedTopNMetricSpecAllOf {
	return v.value
}

func (v *NullableTelemetryDruidInvertedTopNMetricSpecAllOf) Set(val *TelemetryDruidInvertedTopNMetricSpecAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableTelemetryDruidInvertedTopNMetricSpecAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableTelemetryDruidInvertedTopNMetricSpecAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTelemetryDruidInvertedTopNMetricSpecAllOf(val *TelemetryDruidInvertedTopNMetricSpecAllOf) *NullableTelemetryDruidInvertedTopNMetricSpecAllOf {
	return &NullableTelemetryDruidInvertedTopNMetricSpecAllOf{value: val, isSet: true}
}

func (v NullableTelemetryDruidInvertedTopNMetricSpecAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTelemetryDruidInvertedTopNMetricSpecAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
