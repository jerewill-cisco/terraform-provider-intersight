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

// AssetDeviceConfigurationAllOf Definition of the list of properties defined in 'asset.DeviceConfiguration', excluding properties defined in parent classes.
type AssetDeviceConfigurationAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Specifies whether configuration through the platforms local management interface has been disabled, with only configuration through the Intersight service enabled.
	LocalConfigurationLocked *bool `json:"LocalConfigurationLocked,omitempty"`
	// The log level of the device connector service.
	LogLevel             *string                              `json:"LogLevel,omitempty"`
	Device               *AssetDeviceRegistrationRelationship `json:"Device,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AssetDeviceConfigurationAllOf AssetDeviceConfigurationAllOf

// NewAssetDeviceConfigurationAllOf instantiates a new AssetDeviceConfigurationAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAssetDeviceConfigurationAllOf(classId string, objectType string) *AssetDeviceConfigurationAllOf {
	this := AssetDeviceConfigurationAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewAssetDeviceConfigurationAllOfWithDefaults instantiates a new AssetDeviceConfigurationAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAssetDeviceConfigurationAllOfWithDefaults() *AssetDeviceConfigurationAllOf {
	this := AssetDeviceConfigurationAllOf{}
	var classId string = "asset.DeviceConfiguration"
	this.ClassId = classId
	var objectType string = "asset.DeviceConfiguration"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *AssetDeviceConfigurationAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *AssetDeviceConfigurationAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *AssetDeviceConfigurationAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *AssetDeviceConfigurationAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *AssetDeviceConfigurationAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *AssetDeviceConfigurationAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetLocalConfigurationLocked returns the LocalConfigurationLocked field value if set, zero value otherwise.
func (o *AssetDeviceConfigurationAllOf) GetLocalConfigurationLocked() bool {
	if o == nil || o.LocalConfigurationLocked == nil {
		var ret bool
		return ret
	}
	return *o.LocalConfigurationLocked
}

// GetLocalConfigurationLockedOk returns a tuple with the LocalConfigurationLocked field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetDeviceConfigurationAllOf) GetLocalConfigurationLockedOk() (*bool, bool) {
	if o == nil || o.LocalConfigurationLocked == nil {
		return nil, false
	}
	return o.LocalConfigurationLocked, true
}

// HasLocalConfigurationLocked returns a boolean if a field has been set.
func (o *AssetDeviceConfigurationAllOf) HasLocalConfigurationLocked() bool {
	if o != nil && o.LocalConfigurationLocked != nil {
		return true
	}

	return false
}

// SetLocalConfigurationLocked gets a reference to the given bool and assigns it to the LocalConfigurationLocked field.
func (o *AssetDeviceConfigurationAllOf) SetLocalConfigurationLocked(v bool) {
	o.LocalConfigurationLocked = &v
}

// GetLogLevel returns the LogLevel field value if set, zero value otherwise.
func (o *AssetDeviceConfigurationAllOf) GetLogLevel() string {
	if o == nil || o.LogLevel == nil {
		var ret string
		return ret
	}
	return *o.LogLevel
}

// GetLogLevelOk returns a tuple with the LogLevel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetDeviceConfigurationAllOf) GetLogLevelOk() (*string, bool) {
	if o == nil || o.LogLevel == nil {
		return nil, false
	}
	return o.LogLevel, true
}

// HasLogLevel returns a boolean if a field has been set.
func (o *AssetDeviceConfigurationAllOf) HasLogLevel() bool {
	if o != nil && o.LogLevel != nil {
		return true
	}

	return false
}

// SetLogLevel gets a reference to the given string and assigns it to the LogLevel field.
func (o *AssetDeviceConfigurationAllOf) SetLogLevel(v string) {
	o.LogLevel = &v
}

// GetDevice returns the Device field value if set, zero value otherwise.
func (o *AssetDeviceConfigurationAllOf) GetDevice() AssetDeviceRegistrationRelationship {
	if o == nil || o.Device == nil {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.Device
}

// GetDeviceOk returns a tuple with the Device field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetDeviceConfigurationAllOf) GetDeviceOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil || o.Device == nil {
		return nil, false
	}
	return o.Device, true
}

// HasDevice returns a boolean if a field has been set.
func (o *AssetDeviceConfigurationAllOf) HasDevice() bool {
	if o != nil && o.Device != nil {
		return true
	}

	return false
}

// SetDevice gets a reference to the given AssetDeviceRegistrationRelationship and assigns it to the Device field.
func (o *AssetDeviceConfigurationAllOf) SetDevice(v AssetDeviceRegistrationRelationship) {
	o.Device = &v
}

func (o AssetDeviceConfigurationAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.LocalConfigurationLocked != nil {
		toSerialize["LocalConfigurationLocked"] = o.LocalConfigurationLocked
	}
	if o.LogLevel != nil {
		toSerialize["LogLevel"] = o.LogLevel
	}
	if o.Device != nil {
		toSerialize["Device"] = o.Device
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *AssetDeviceConfigurationAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varAssetDeviceConfigurationAllOf := _AssetDeviceConfigurationAllOf{}

	if err = json.Unmarshal(bytes, &varAssetDeviceConfigurationAllOf); err == nil {
		*o = AssetDeviceConfigurationAllOf(varAssetDeviceConfigurationAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "LocalConfigurationLocked")
		delete(additionalProperties, "LogLevel")
		delete(additionalProperties, "Device")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAssetDeviceConfigurationAllOf struct {
	value *AssetDeviceConfigurationAllOf
	isSet bool
}

func (v NullableAssetDeviceConfigurationAllOf) Get() *AssetDeviceConfigurationAllOf {
	return v.value
}

func (v *NullableAssetDeviceConfigurationAllOf) Set(val *AssetDeviceConfigurationAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableAssetDeviceConfigurationAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableAssetDeviceConfigurationAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAssetDeviceConfigurationAllOf(val *AssetDeviceConfigurationAllOf) *NullableAssetDeviceConfigurationAllOf {
	return &NullableAssetDeviceConfigurationAllOf{value: val, isSet: true}
}

func (v NullableAssetDeviceConfigurationAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAssetDeviceConfigurationAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
