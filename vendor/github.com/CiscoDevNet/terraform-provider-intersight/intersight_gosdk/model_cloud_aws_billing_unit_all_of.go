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

// CloudAwsBillingUnitAllOf Definition of the list of properties defined in 'cloud.AwsBillingUnit', excluding properties defined in parent classes.
type CloudAwsBillingUnitAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Email address of the account holder.
	Email                *string                              `json:"Email,omitempty"`
	RegisteredDevice     *AssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CloudAwsBillingUnitAllOf CloudAwsBillingUnitAllOf

// NewCloudAwsBillingUnitAllOf instantiates a new CloudAwsBillingUnitAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCloudAwsBillingUnitAllOf(classId string, objectType string) *CloudAwsBillingUnitAllOf {
	this := CloudAwsBillingUnitAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewCloudAwsBillingUnitAllOfWithDefaults instantiates a new CloudAwsBillingUnitAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCloudAwsBillingUnitAllOfWithDefaults() *CloudAwsBillingUnitAllOf {
	this := CloudAwsBillingUnitAllOf{}
	var classId string = "cloud.AwsBillingUnit"
	this.ClassId = classId
	var objectType string = "cloud.AwsBillingUnit"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *CloudAwsBillingUnitAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *CloudAwsBillingUnitAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *CloudAwsBillingUnitAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *CloudAwsBillingUnitAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *CloudAwsBillingUnitAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *CloudAwsBillingUnitAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *CloudAwsBillingUnitAllOf) GetEmail() string {
	if o == nil || o.Email == nil {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudAwsBillingUnitAllOf) GetEmailOk() (*string, bool) {
	if o == nil || o.Email == nil {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *CloudAwsBillingUnitAllOf) HasEmail() bool {
	if o != nil && o.Email != nil {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *CloudAwsBillingUnitAllOf) SetEmail(v string) {
	o.Email = &v
}

// GetRegisteredDevice returns the RegisteredDevice field value if set, zero value otherwise.
func (o *CloudAwsBillingUnitAllOf) GetRegisteredDevice() AssetDeviceRegistrationRelationship {
	if o == nil || o.RegisteredDevice == nil {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.RegisteredDevice
}

// GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudAwsBillingUnitAllOf) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil || o.RegisteredDevice == nil {
		return nil, false
	}
	return o.RegisteredDevice, true
}

// HasRegisteredDevice returns a boolean if a field has been set.
func (o *CloudAwsBillingUnitAllOf) HasRegisteredDevice() bool {
	if o != nil && o.RegisteredDevice != nil {
		return true
	}

	return false
}

// SetRegisteredDevice gets a reference to the given AssetDeviceRegistrationRelationship and assigns it to the RegisteredDevice field.
func (o *CloudAwsBillingUnitAllOf) SetRegisteredDevice(v AssetDeviceRegistrationRelationship) {
	o.RegisteredDevice = &v
}

func (o CloudAwsBillingUnitAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Email != nil {
		toSerialize["Email"] = o.Email
	}
	if o.RegisteredDevice != nil {
		toSerialize["RegisteredDevice"] = o.RegisteredDevice
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *CloudAwsBillingUnitAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varCloudAwsBillingUnitAllOf := _CloudAwsBillingUnitAllOf{}

	if err = json.Unmarshal(bytes, &varCloudAwsBillingUnitAllOf); err == nil {
		*o = CloudAwsBillingUnitAllOf(varCloudAwsBillingUnitAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Email")
		delete(additionalProperties, "RegisteredDevice")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCloudAwsBillingUnitAllOf struct {
	value *CloudAwsBillingUnitAllOf
	isSet bool
}

func (v NullableCloudAwsBillingUnitAllOf) Get() *CloudAwsBillingUnitAllOf {
	return v.value
}

func (v *NullableCloudAwsBillingUnitAllOf) Set(val *CloudAwsBillingUnitAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableCloudAwsBillingUnitAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableCloudAwsBillingUnitAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCloudAwsBillingUnitAllOf(val *CloudAwsBillingUnitAllOf) *NullableCloudAwsBillingUnitAllOf {
	return &NullableCloudAwsBillingUnitAllOf{value: val, isSet: true}
}

func (v NullableCloudAwsBillingUnitAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCloudAwsBillingUnitAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
