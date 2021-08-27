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

// UcsdUcsdRestoreParametersAllOf Definition of the list of properties defined in 'ucsd.UcsdRestoreParameters', excluding properties defined in parent classes.
type UcsdUcsdRestoreParametersAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Indicates whether the value of the 'password' property has been set.
	IsPasswordSet *bool `json:"IsPasswordSet,omitempty"`
	// The complete location of the path on the server. The location should be specified in the following format- hostname-or-ipv4address<:port>/absolute-file-path.
	Location *string `json:"Location,omitempty"`
	// The password of the target backup server. Only required if the target server is accessed using SFTP or SCP protocol.
	Password *string `json:"Password,omitempty"`
	// The protocol used to backup the UCS Director.
	Protocol *string `json:"Protocol,omitempty"`
	// Decides whether UCS Director property files should also be restored.
	RestoreConfigurationFiles *bool `json:"RestoreConfigurationFiles,omitempty"`
	// Decides whether license should also be restored.
	RestoreLicense *bool `json:"RestoreLicense,omitempty"`
	// The username of the target backup server. Only required if the target server is accessed using SFTP or SCP protocol.
	Username             *string `json:"Username,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _UcsdUcsdRestoreParametersAllOf UcsdUcsdRestoreParametersAllOf

// NewUcsdUcsdRestoreParametersAllOf instantiates a new UcsdUcsdRestoreParametersAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUcsdUcsdRestoreParametersAllOf(classId string, objectType string) *UcsdUcsdRestoreParametersAllOf {
	this := UcsdUcsdRestoreParametersAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewUcsdUcsdRestoreParametersAllOfWithDefaults instantiates a new UcsdUcsdRestoreParametersAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUcsdUcsdRestoreParametersAllOfWithDefaults() *UcsdUcsdRestoreParametersAllOf {
	this := UcsdUcsdRestoreParametersAllOf{}
	var classId string = "ucsd.UcsdRestoreParameters"
	this.ClassId = classId
	var objectType string = "ucsd.UcsdRestoreParameters"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *UcsdUcsdRestoreParametersAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *UcsdUcsdRestoreParametersAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *UcsdUcsdRestoreParametersAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *UcsdUcsdRestoreParametersAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *UcsdUcsdRestoreParametersAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *UcsdUcsdRestoreParametersAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetIsPasswordSet returns the IsPasswordSet field value if set, zero value otherwise.
func (o *UcsdUcsdRestoreParametersAllOf) GetIsPasswordSet() bool {
	if o == nil || o.IsPasswordSet == nil {
		var ret bool
		return ret
	}
	return *o.IsPasswordSet
}

// GetIsPasswordSetOk returns a tuple with the IsPasswordSet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UcsdUcsdRestoreParametersAllOf) GetIsPasswordSetOk() (*bool, bool) {
	if o == nil || o.IsPasswordSet == nil {
		return nil, false
	}
	return o.IsPasswordSet, true
}

// HasIsPasswordSet returns a boolean if a field has been set.
func (o *UcsdUcsdRestoreParametersAllOf) HasIsPasswordSet() bool {
	if o != nil && o.IsPasswordSet != nil {
		return true
	}

	return false
}

// SetIsPasswordSet gets a reference to the given bool and assigns it to the IsPasswordSet field.
func (o *UcsdUcsdRestoreParametersAllOf) SetIsPasswordSet(v bool) {
	o.IsPasswordSet = &v
}

// GetLocation returns the Location field value if set, zero value otherwise.
func (o *UcsdUcsdRestoreParametersAllOf) GetLocation() string {
	if o == nil || o.Location == nil {
		var ret string
		return ret
	}
	return *o.Location
}

// GetLocationOk returns a tuple with the Location field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UcsdUcsdRestoreParametersAllOf) GetLocationOk() (*string, bool) {
	if o == nil || o.Location == nil {
		return nil, false
	}
	return o.Location, true
}

// HasLocation returns a boolean if a field has been set.
func (o *UcsdUcsdRestoreParametersAllOf) HasLocation() bool {
	if o != nil && o.Location != nil {
		return true
	}

	return false
}

// SetLocation gets a reference to the given string and assigns it to the Location field.
func (o *UcsdUcsdRestoreParametersAllOf) SetLocation(v string) {
	o.Location = &v
}

// GetPassword returns the Password field value if set, zero value otherwise.
func (o *UcsdUcsdRestoreParametersAllOf) GetPassword() string {
	if o == nil || o.Password == nil {
		var ret string
		return ret
	}
	return *o.Password
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UcsdUcsdRestoreParametersAllOf) GetPasswordOk() (*string, bool) {
	if o == nil || o.Password == nil {
		return nil, false
	}
	return o.Password, true
}

// HasPassword returns a boolean if a field has been set.
func (o *UcsdUcsdRestoreParametersAllOf) HasPassword() bool {
	if o != nil && o.Password != nil {
		return true
	}

	return false
}

// SetPassword gets a reference to the given string and assigns it to the Password field.
func (o *UcsdUcsdRestoreParametersAllOf) SetPassword(v string) {
	o.Password = &v
}

// GetProtocol returns the Protocol field value if set, zero value otherwise.
func (o *UcsdUcsdRestoreParametersAllOf) GetProtocol() string {
	if o == nil || o.Protocol == nil {
		var ret string
		return ret
	}
	return *o.Protocol
}

// GetProtocolOk returns a tuple with the Protocol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UcsdUcsdRestoreParametersAllOf) GetProtocolOk() (*string, bool) {
	if o == nil || o.Protocol == nil {
		return nil, false
	}
	return o.Protocol, true
}

// HasProtocol returns a boolean if a field has been set.
func (o *UcsdUcsdRestoreParametersAllOf) HasProtocol() bool {
	if o != nil && o.Protocol != nil {
		return true
	}

	return false
}

// SetProtocol gets a reference to the given string and assigns it to the Protocol field.
func (o *UcsdUcsdRestoreParametersAllOf) SetProtocol(v string) {
	o.Protocol = &v
}

// GetRestoreConfigurationFiles returns the RestoreConfigurationFiles field value if set, zero value otherwise.
func (o *UcsdUcsdRestoreParametersAllOf) GetRestoreConfigurationFiles() bool {
	if o == nil || o.RestoreConfigurationFiles == nil {
		var ret bool
		return ret
	}
	return *o.RestoreConfigurationFiles
}

// GetRestoreConfigurationFilesOk returns a tuple with the RestoreConfigurationFiles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UcsdUcsdRestoreParametersAllOf) GetRestoreConfigurationFilesOk() (*bool, bool) {
	if o == nil || o.RestoreConfigurationFiles == nil {
		return nil, false
	}
	return o.RestoreConfigurationFiles, true
}

// HasRestoreConfigurationFiles returns a boolean if a field has been set.
func (o *UcsdUcsdRestoreParametersAllOf) HasRestoreConfigurationFiles() bool {
	if o != nil && o.RestoreConfigurationFiles != nil {
		return true
	}

	return false
}

// SetRestoreConfigurationFiles gets a reference to the given bool and assigns it to the RestoreConfigurationFiles field.
func (o *UcsdUcsdRestoreParametersAllOf) SetRestoreConfigurationFiles(v bool) {
	o.RestoreConfigurationFiles = &v
}

// GetRestoreLicense returns the RestoreLicense field value if set, zero value otherwise.
func (o *UcsdUcsdRestoreParametersAllOf) GetRestoreLicense() bool {
	if o == nil || o.RestoreLicense == nil {
		var ret bool
		return ret
	}
	return *o.RestoreLicense
}

// GetRestoreLicenseOk returns a tuple with the RestoreLicense field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UcsdUcsdRestoreParametersAllOf) GetRestoreLicenseOk() (*bool, bool) {
	if o == nil || o.RestoreLicense == nil {
		return nil, false
	}
	return o.RestoreLicense, true
}

// HasRestoreLicense returns a boolean if a field has been set.
func (o *UcsdUcsdRestoreParametersAllOf) HasRestoreLicense() bool {
	if o != nil && o.RestoreLicense != nil {
		return true
	}

	return false
}

// SetRestoreLicense gets a reference to the given bool and assigns it to the RestoreLicense field.
func (o *UcsdUcsdRestoreParametersAllOf) SetRestoreLicense(v bool) {
	o.RestoreLicense = &v
}

// GetUsername returns the Username field value if set, zero value otherwise.
func (o *UcsdUcsdRestoreParametersAllOf) GetUsername() string {
	if o == nil || o.Username == nil {
		var ret string
		return ret
	}
	return *o.Username
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UcsdUcsdRestoreParametersAllOf) GetUsernameOk() (*string, bool) {
	if o == nil || o.Username == nil {
		return nil, false
	}
	return o.Username, true
}

// HasUsername returns a boolean if a field has been set.
func (o *UcsdUcsdRestoreParametersAllOf) HasUsername() bool {
	if o != nil && o.Username != nil {
		return true
	}

	return false
}

// SetUsername gets a reference to the given string and assigns it to the Username field.
func (o *UcsdUcsdRestoreParametersAllOf) SetUsername(v string) {
	o.Username = &v
}

func (o UcsdUcsdRestoreParametersAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.IsPasswordSet != nil {
		toSerialize["IsPasswordSet"] = o.IsPasswordSet
	}
	if o.Location != nil {
		toSerialize["Location"] = o.Location
	}
	if o.Password != nil {
		toSerialize["Password"] = o.Password
	}
	if o.Protocol != nil {
		toSerialize["Protocol"] = o.Protocol
	}
	if o.RestoreConfigurationFiles != nil {
		toSerialize["RestoreConfigurationFiles"] = o.RestoreConfigurationFiles
	}
	if o.RestoreLicense != nil {
		toSerialize["RestoreLicense"] = o.RestoreLicense
	}
	if o.Username != nil {
		toSerialize["Username"] = o.Username
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *UcsdUcsdRestoreParametersAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varUcsdUcsdRestoreParametersAllOf := _UcsdUcsdRestoreParametersAllOf{}

	if err = json.Unmarshal(bytes, &varUcsdUcsdRestoreParametersAllOf); err == nil {
		*o = UcsdUcsdRestoreParametersAllOf(varUcsdUcsdRestoreParametersAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "IsPasswordSet")
		delete(additionalProperties, "Location")
		delete(additionalProperties, "Password")
		delete(additionalProperties, "Protocol")
		delete(additionalProperties, "RestoreConfigurationFiles")
		delete(additionalProperties, "RestoreLicense")
		delete(additionalProperties, "Username")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableUcsdUcsdRestoreParametersAllOf struct {
	value *UcsdUcsdRestoreParametersAllOf
	isSet bool
}

func (v NullableUcsdUcsdRestoreParametersAllOf) Get() *UcsdUcsdRestoreParametersAllOf {
	return v.value
}

func (v *NullableUcsdUcsdRestoreParametersAllOf) Set(val *UcsdUcsdRestoreParametersAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableUcsdUcsdRestoreParametersAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableUcsdUcsdRestoreParametersAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUcsdUcsdRestoreParametersAllOf(val *UcsdUcsdRestoreParametersAllOf) *NullableUcsdUcsdRestoreParametersAllOf {
	return &NullableUcsdUcsdRestoreParametersAllOf{value: val, isSet: true}
}

func (v NullableUcsdUcsdRestoreParametersAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUcsdUcsdRestoreParametersAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
