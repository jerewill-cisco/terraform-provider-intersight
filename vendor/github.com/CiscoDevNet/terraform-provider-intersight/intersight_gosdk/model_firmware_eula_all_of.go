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

// FirmwareEulaAllOf Definition of the list of properties defined in 'firmware.Eula', excluding properties defined in parent classes.
type FirmwareEulaAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// EULA acceptance status for the account.
	Accepted *bool `json:"Accepted,omitempty"`
	// EULA acceptance form content provided by cisco.com.
	Content              *string                 `json:"Content,omitempty"`
	Account              *IamAccountRelationship `json:"Account,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _FirmwareEulaAllOf FirmwareEulaAllOf

// NewFirmwareEulaAllOf instantiates a new FirmwareEulaAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFirmwareEulaAllOf(classId string, objectType string) *FirmwareEulaAllOf {
	this := FirmwareEulaAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewFirmwareEulaAllOfWithDefaults instantiates a new FirmwareEulaAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFirmwareEulaAllOfWithDefaults() *FirmwareEulaAllOf {
	this := FirmwareEulaAllOf{}
	var classId string = "firmware.Eula"
	this.ClassId = classId
	var objectType string = "firmware.Eula"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *FirmwareEulaAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *FirmwareEulaAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *FirmwareEulaAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *FirmwareEulaAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *FirmwareEulaAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *FirmwareEulaAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetAccepted returns the Accepted field value if set, zero value otherwise.
func (o *FirmwareEulaAllOf) GetAccepted() bool {
	if o == nil || o.Accepted == nil {
		var ret bool
		return ret
	}
	return *o.Accepted
}

// GetAcceptedOk returns a tuple with the Accepted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FirmwareEulaAllOf) GetAcceptedOk() (*bool, bool) {
	if o == nil || o.Accepted == nil {
		return nil, false
	}
	return o.Accepted, true
}

// HasAccepted returns a boolean if a field has been set.
func (o *FirmwareEulaAllOf) HasAccepted() bool {
	if o != nil && o.Accepted != nil {
		return true
	}

	return false
}

// SetAccepted gets a reference to the given bool and assigns it to the Accepted field.
func (o *FirmwareEulaAllOf) SetAccepted(v bool) {
	o.Accepted = &v
}

// GetContent returns the Content field value if set, zero value otherwise.
func (o *FirmwareEulaAllOf) GetContent() string {
	if o == nil || o.Content == nil {
		var ret string
		return ret
	}
	return *o.Content
}

// GetContentOk returns a tuple with the Content field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FirmwareEulaAllOf) GetContentOk() (*string, bool) {
	if o == nil || o.Content == nil {
		return nil, false
	}
	return o.Content, true
}

// HasContent returns a boolean if a field has been set.
func (o *FirmwareEulaAllOf) HasContent() bool {
	if o != nil && o.Content != nil {
		return true
	}

	return false
}

// SetContent gets a reference to the given string and assigns it to the Content field.
func (o *FirmwareEulaAllOf) SetContent(v string) {
	o.Content = &v
}

// GetAccount returns the Account field value if set, zero value otherwise.
func (o *FirmwareEulaAllOf) GetAccount() IamAccountRelationship {
	if o == nil || o.Account == nil {
		var ret IamAccountRelationship
		return ret
	}
	return *o.Account
}

// GetAccountOk returns a tuple with the Account field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FirmwareEulaAllOf) GetAccountOk() (*IamAccountRelationship, bool) {
	if o == nil || o.Account == nil {
		return nil, false
	}
	return o.Account, true
}

// HasAccount returns a boolean if a field has been set.
func (o *FirmwareEulaAllOf) HasAccount() bool {
	if o != nil && o.Account != nil {
		return true
	}

	return false
}

// SetAccount gets a reference to the given IamAccountRelationship and assigns it to the Account field.
func (o *FirmwareEulaAllOf) SetAccount(v IamAccountRelationship) {
	o.Account = &v
}

func (o FirmwareEulaAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Accepted != nil {
		toSerialize["Accepted"] = o.Accepted
	}
	if o.Content != nil {
		toSerialize["Content"] = o.Content
	}
	if o.Account != nil {
		toSerialize["Account"] = o.Account
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *FirmwareEulaAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varFirmwareEulaAllOf := _FirmwareEulaAllOf{}

	if err = json.Unmarshal(bytes, &varFirmwareEulaAllOf); err == nil {
		*o = FirmwareEulaAllOf(varFirmwareEulaAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Accepted")
		delete(additionalProperties, "Content")
		delete(additionalProperties, "Account")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableFirmwareEulaAllOf struct {
	value *FirmwareEulaAllOf
	isSet bool
}

func (v NullableFirmwareEulaAllOf) Get() *FirmwareEulaAllOf {
	return v.value
}

func (v *NullableFirmwareEulaAllOf) Set(val *FirmwareEulaAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableFirmwareEulaAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableFirmwareEulaAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFirmwareEulaAllOf(val *FirmwareEulaAllOf) *NullableFirmwareEulaAllOf {
	return &NullableFirmwareEulaAllOf{value: val, isSet: true}
}

func (v NullableFirmwareEulaAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFirmwareEulaAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
