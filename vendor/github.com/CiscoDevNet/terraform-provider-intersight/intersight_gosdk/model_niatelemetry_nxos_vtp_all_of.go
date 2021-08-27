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

// NiatelemetryNxosVtpAllOf Definition of the list of properties defined in 'niatelemetry.NxosVtp', excluding properties defined in parent classes.
type NiatelemetryNxosVtpAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Returns the status of operational mode of vtp.
	OperMode *string `json:"OperMode,omitempty"`
	// Returns the status pruning mode of vtp.
	PruningMode *string `json:"PruningMode,omitempty"`
	// Returns the running version of vtp.
	RunningVersion *string `json:"RunningVersion,omitempty"`
	// Returns the status of trap in vtp.
	TrapEnabled *string `json:"TrapEnabled,omitempty"`
	// Returns the status of v2 mode of vtp.
	V2Mode *string `json:"V2Mode,omitempty"`
	// Returns version of vtp running.
	Version              *int64 `json:"Version,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _NiatelemetryNxosVtpAllOf NiatelemetryNxosVtpAllOf

// NewNiatelemetryNxosVtpAllOf instantiates a new NiatelemetryNxosVtpAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNiatelemetryNxosVtpAllOf(classId string, objectType string) *NiatelemetryNxosVtpAllOf {
	this := NiatelemetryNxosVtpAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewNiatelemetryNxosVtpAllOfWithDefaults instantiates a new NiatelemetryNxosVtpAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNiatelemetryNxosVtpAllOfWithDefaults() *NiatelemetryNxosVtpAllOf {
	this := NiatelemetryNxosVtpAllOf{}
	var classId string = "niatelemetry.NxosVtp"
	this.ClassId = classId
	var objectType string = "niatelemetry.NxosVtp"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *NiatelemetryNxosVtpAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *NiatelemetryNxosVtpAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *NiatelemetryNxosVtpAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *NiatelemetryNxosVtpAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *NiatelemetryNxosVtpAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *NiatelemetryNxosVtpAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetOperMode returns the OperMode field value if set, zero value otherwise.
func (o *NiatelemetryNxosVtpAllOf) GetOperMode() string {
	if o == nil || o.OperMode == nil {
		var ret string
		return ret
	}
	return *o.OperMode
}

// GetOperModeOk returns a tuple with the OperMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryNxosVtpAllOf) GetOperModeOk() (*string, bool) {
	if o == nil || o.OperMode == nil {
		return nil, false
	}
	return o.OperMode, true
}

// HasOperMode returns a boolean if a field has been set.
func (o *NiatelemetryNxosVtpAllOf) HasOperMode() bool {
	if o != nil && o.OperMode != nil {
		return true
	}

	return false
}

// SetOperMode gets a reference to the given string and assigns it to the OperMode field.
func (o *NiatelemetryNxosVtpAllOf) SetOperMode(v string) {
	o.OperMode = &v
}

// GetPruningMode returns the PruningMode field value if set, zero value otherwise.
func (o *NiatelemetryNxosVtpAllOf) GetPruningMode() string {
	if o == nil || o.PruningMode == nil {
		var ret string
		return ret
	}
	return *o.PruningMode
}

// GetPruningModeOk returns a tuple with the PruningMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryNxosVtpAllOf) GetPruningModeOk() (*string, bool) {
	if o == nil || o.PruningMode == nil {
		return nil, false
	}
	return o.PruningMode, true
}

// HasPruningMode returns a boolean if a field has been set.
func (o *NiatelemetryNxosVtpAllOf) HasPruningMode() bool {
	if o != nil && o.PruningMode != nil {
		return true
	}

	return false
}

// SetPruningMode gets a reference to the given string and assigns it to the PruningMode field.
func (o *NiatelemetryNxosVtpAllOf) SetPruningMode(v string) {
	o.PruningMode = &v
}

// GetRunningVersion returns the RunningVersion field value if set, zero value otherwise.
func (o *NiatelemetryNxosVtpAllOf) GetRunningVersion() string {
	if o == nil || o.RunningVersion == nil {
		var ret string
		return ret
	}
	return *o.RunningVersion
}

// GetRunningVersionOk returns a tuple with the RunningVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryNxosVtpAllOf) GetRunningVersionOk() (*string, bool) {
	if o == nil || o.RunningVersion == nil {
		return nil, false
	}
	return o.RunningVersion, true
}

// HasRunningVersion returns a boolean if a field has been set.
func (o *NiatelemetryNxosVtpAllOf) HasRunningVersion() bool {
	if o != nil && o.RunningVersion != nil {
		return true
	}

	return false
}

// SetRunningVersion gets a reference to the given string and assigns it to the RunningVersion field.
func (o *NiatelemetryNxosVtpAllOf) SetRunningVersion(v string) {
	o.RunningVersion = &v
}

// GetTrapEnabled returns the TrapEnabled field value if set, zero value otherwise.
func (o *NiatelemetryNxosVtpAllOf) GetTrapEnabled() string {
	if o == nil || o.TrapEnabled == nil {
		var ret string
		return ret
	}
	return *o.TrapEnabled
}

// GetTrapEnabledOk returns a tuple with the TrapEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryNxosVtpAllOf) GetTrapEnabledOk() (*string, bool) {
	if o == nil || o.TrapEnabled == nil {
		return nil, false
	}
	return o.TrapEnabled, true
}

// HasTrapEnabled returns a boolean if a field has been set.
func (o *NiatelemetryNxosVtpAllOf) HasTrapEnabled() bool {
	if o != nil && o.TrapEnabled != nil {
		return true
	}

	return false
}

// SetTrapEnabled gets a reference to the given string and assigns it to the TrapEnabled field.
func (o *NiatelemetryNxosVtpAllOf) SetTrapEnabled(v string) {
	o.TrapEnabled = &v
}

// GetV2Mode returns the V2Mode field value if set, zero value otherwise.
func (o *NiatelemetryNxosVtpAllOf) GetV2Mode() string {
	if o == nil || o.V2Mode == nil {
		var ret string
		return ret
	}
	return *o.V2Mode
}

// GetV2ModeOk returns a tuple with the V2Mode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryNxosVtpAllOf) GetV2ModeOk() (*string, bool) {
	if o == nil || o.V2Mode == nil {
		return nil, false
	}
	return o.V2Mode, true
}

// HasV2Mode returns a boolean if a field has been set.
func (o *NiatelemetryNxosVtpAllOf) HasV2Mode() bool {
	if o != nil && o.V2Mode != nil {
		return true
	}

	return false
}

// SetV2Mode gets a reference to the given string and assigns it to the V2Mode field.
func (o *NiatelemetryNxosVtpAllOf) SetV2Mode(v string) {
	o.V2Mode = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *NiatelemetryNxosVtpAllOf) GetVersion() int64 {
	if o == nil || o.Version == nil {
		var ret int64
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryNxosVtpAllOf) GetVersionOk() (*int64, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *NiatelemetryNxosVtpAllOf) HasVersion() bool {
	if o != nil && o.Version != nil {
		return true
	}

	return false
}

// SetVersion gets a reference to the given int64 and assigns it to the Version field.
func (o *NiatelemetryNxosVtpAllOf) SetVersion(v int64) {
	o.Version = &v
}

func (o NiatelemetryNxosVtpAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.OperMode != nil {
		toSerialize["OperMode"] = o.OperMode
	}
	if o.PruningMode != nil {
		toSerialize["PruningMode"] = o.PruningMode
	}
	if o.RunningVersion != nil {
		toSerialize["RunningVersion"] = o.RunningVersion
	}
	if o.TrapEnabled != nil {
		toSerialize["TrapEnabled"] = o.TrapEnabled
	}
	if o.V2Mode != nil {
		toSerialize["V2Mode"] = o.V2Mode
	}
	if o.Version != nil {
		toSerialize["Version"] = o.Version
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *NiatelemetryNxosVtpAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varNiatelemetryNxosVtpAllOf := _NiatelemetryNxosVtpAllOf{}

	if err = json.Unmarshal(bytes, &varNiatelemetryNxosVtpAllOf); err == nil {
		*o = NiatelemetryNxosVtpAllOf(varNiatelemetryNxosVtpAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "OperMode")
		delete(additionalProperties, "PruningMode")
		delete(additionalProperties, "RunningVersion")
		delete(additionalProperties, "TrapEnabled")
		delete(additionalProperties, "V2Mode")
		delete(additionalProperties, "Version")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableNiatelemetryNxosVtpAllOf struct {
	value *NiatelemetryNxosVtpAllOf
	isSet bool
}

func (v NullableNiatelemetryNxosVtpAllOf) Get() *NiatelemetryNxosVtpAllOf {
	return v.value
}

func (v *NullableNiatelemetryNxosVtpAllOf) Set(val *NiatelemetryNxosVtpAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableNiatelemetryNxosVtpAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableNiatelemetryNxosVtpAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNiatelemetryNxosVtpAllOf(val *NiatelemetryNxosVtpAllOf) *NullableNiatelemetryNxosVtpAllOf {
	return &NullableNiatelemetryNxosVtpAllOf{value: val, isSet: true}
}

func (v NullableNiatelemetryNxosVtpAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNiatelemetryNxosVtpAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
