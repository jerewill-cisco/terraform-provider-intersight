/*
 * Cisco Intersight
 *
 * Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document. This document was created on 2021-05-12T14:10:48Z.
 *
 * API version: 1.0.9-4289
 * Contact: intersight@cisco.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package intersight

import (
	"encoding/json"
)

// StorageNetAppExportPolicyRuleAllOf Definition of the list of properties defined in 'storage.NetAppExportPolicyRule', excluding properties defined in parent classes.
type StorageNetAppExportPolicyRuleAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType  string   `json:"ObjectType"`
	ClientMatch []string `json:"ClientMatch,omitempty"`
	// Position of export rule in the list of rules.
	Index     *int64   `json:"Index,omitempty"`
	RoRule    []string `json:"RoRule,omitempty"`
	RwRule    []string `json:"RwRule,omitempty"`
	SuperUser []string `json:"SuperUser,omitempty"`
	// Export Policy rule that are mapped to this User ID.
	User                 *string `json:"User,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _StorageNetAppExportPolicyRuleAllOf StorageNetAppExportPolicyRuleAllOf

// NewStorageNetAppExportPolicyRuleAllOf instantiates a new StorageNetAppExportPolicyRuleAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStorageNetAppExportPolicyRuleAllOf(classId string, objectType string) *StorageNetAppExportPolicyRuleAllOf {
	this := StorageNetAppExportPolicyRuleAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewStorageNetAppExportPolicyRuleAllOfWithDefaults instantiates a new StorageNetAppExportPolicyRuleAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStorageNetAppExportPolicyRuleAllOfWithDefaults() *StorageNetAppExportPolicyRuleAllOf {
	this := StorageNetAppExportPolicyRuleAllOf{}
	var classId string = "storage.NetAppExportPolicyRule"
	this.ClassId = classId
	var objectType string = "storage.NetAppExportPolicyRule"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *StorageNetAppExportPolicyRuleAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *StorageNetAppExportPolicyRuleAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *StorageNetAppExportPolicyRuleAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *StorageNetAppExportPolicyRuleAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *StorageNetAppExportPolicyRuleAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *StorageNetAppExportPolicyRuleAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetClientMatch returns the ClientMatch field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StorageNetAppExportPolicyRuleAllOf) GetClientMatch() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.ClientMatch
}

// GetClientMatchOk returns a tuple with the ClientMatch field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StorageNetAppExportPolicyRuleAllOf) GetClientMatchOk() (*[]string, bool) {
	if o == nil || o.ClientMatch == nil {
		return nil, false
	}
	return &o.ClientMatch, true
}

// HasClientMatch returns a boolean if a field has been set.
func (o *StorageNetAppExportPolicyRuleAllOf) HasClientMatch() bool {
	if o != nil && o.ClientMatch != nil {
		return true
	}

	return false
}

// SetClientMatch gets a reference to the given []string and assigns it to the ClientMatch field.
func (o *StorageNetAppExportPolicyRuleAllOf) SetClientMatch(v []string) {
	o.ClientMatch = v
}

// GetIndex returns the Index field value if set, zero value otherwise.
func (o *StorageNetAppExportPolicyRuleAllOf) GetIndex() int64 {
	if o == nil || o.Index == nil {
		var ret int64
		return ret
	}
	return *o.Index
}

// GetIndexOk returns a tuple with the Index field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNetAppExportPolicyRuleAllOf) GetIndexOk() (*int64, bool) {
	if o == nil || o.Index == nil {
		return nil, false
	}
	return o.Index, true
}

// HasIndex returns a boolean if a field has been set.
func (o *StorageNetAppExportPolicyRuleAllOf) HasIndex() bool {
	if o != nil && o.Index != nil {
		return true
	}

	return false
}

// SetIndex gets a reference to the given int64 and assigns it to the Index field.
func (o *StorageNetAppExportPolicyRuleAllOf) SetIndex(v int64) {
	o.Index = &v
}

// GetRoRule returns the RoRule field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StorageNetAppExportPolicyRuleAllOf) GetRoRule() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.RoRule
}

// GetRoRuleOk returns a tuple with the RoRule field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StorageNetAppExportPolicyRuleAllOf) GetRoRuleOk() (*[]string, bool) {
	if o == nil || o.RoRule == nil {
		return nil, false
	}
	return &o.RoRule, true
}

// HasRoRule returns a boolean if a field has been set.
func (o *StorageNetAppExportPolicyRuleAllOf) HasRoRule() bool {
	if o != nil && o.RoRule != nil {
		return true
	}

	return false
}

// SetRoRule gets a reference to the given []string and assigns it to the RoRule field.
func (o *StorageNetAppExportPolicyRuleAllOf) SetRoRule(v []string) {
	o.RoRule = v
}

// GetRwRule returns the RwRule field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StorageNetAppExportPolicyRuleAllOf) GetRwRule() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.RwRule
}

// GetRwRuleOk returns a tuple with the RwRule field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StorageNetAppExportPolicyRuleAllOf) GetRwRuleOk() (*[]string, bool) {
	if o == nil || o.RwRule == nil {
		return nil, false
	}
	return &o.RwRule, true
}

// HasRwRule returns a boolean if a field has been set.
func (o *StorageNetAppExportPolicyRuleAllOf) HasRwRule() bool {
	if o != nil && o.RwRule != nil {
		return true
	}

	return false
}

// SetRwRule gets a reference to the given []string and assigns it to the RwRule field.
func (o *StorageNetAppExportPolicyRuleAllOf) SetRwRule(v []string) {
	o.RwRule = v
}

// GetSuperUser returns the SuperUser field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StorageNetAppExportPolicyRuleAllOf) GetSuperUser() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.SuperUser
}

// GetSuperUserOk returns a tuple with the SuperUser field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StorageNetAppExportPolicyRuleAllOf) GetSuperUserOk() (*[]string, bool) {
	if o == nil || o.SuperUser == nil {
		return nil, false
	}
	return &o.SuperUser, true
}

// HasSuperUser returns a boolean if a field has been set.
func (o *StorageNetAppExportPolicyRuleAllOf) HasSuperUser() bool {
	if o != nil && o.SuperUser != nil {
		return true
	}

	return false
}

// SetSuperUser gets a reference to the given []string and assigns it to the SuperUser field.
func (o *StorageNetAppExportPolicyRuleAllOf) SetSuperUser(v []string) {
	o.SuperUser = v
}

// GetUser returns the User field value if set, zero value otherwise.
func (o *StorageNetAppExportPolicyRuleAllOf) GetUser() string {
	if o == nil || o.User == nil {
		var ret string
		return ret
	}
	return *o.User
}

// GetUserOk returns a tuple with the User field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNetAppExportPolicyRuleAllOf) GetUserOk() (*string, bool) {
	if o == nil || o.User == nil {
		return nil, false
	}
	return o.User, true
}

// HasUser returns a boolean if a field has been set.
func (o *StorageNetAppExportPolicyRuleAllOf) HasUser() bool {
	if o != nil && o.User != nil {
		return true
	}

	return false
}

// SetUser gets a reference to the given string and assigns it to the User field.
func (o *StorageNetAppExportPolicyRuleAllOf) SetUser(v string) {
	o.User = &v
}

func (o StorageNetAppExportPolicyRuleAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.ClientMatch != nil {
		toSerialize["ClientMatch"] = o.ClientMatch
	}
	if o.Index != nil {
		toSerialize["Index"] = o.Index
	}
	if o.RoRule != nil {
		toSerialize["RoRule"] = o.RoRule
	}
	if o.RwRule != nil {
		toSerialize["RwRule"] = o.RwRule
	}
	if o.SuperUser != nil {
		toSerialize["SuperUser"] = o.SuperUser
	}
	if o.User != nil {
		toSerialize["User"] = o.User
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *StorageNetAppExportPolicyRuleAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varStorageNetAppExportPolicyRuleAllOf := _StorageNetAppExportPolicyRuleAllOf{}

	if err = json.Unmarshal(bytes, &varStorageNetAppExportPolicyRuleAllOf); err == nil {
		*o = StorageNetAppExportPolicyRuleAllOf(varStorageNetAppExportPolicyRuleAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "ClientMatch")
		delete(additionalProperties, "Index")
		delete(additionalProperties, "RoRule")
		delete(additionalProperties, "RwRule")
		delete(additionalProperties, "SuperUser")
		delete(additionalProperties, "User")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableStorageNetAppExportPolicyRuleAllOf struct {
	value *StorageNetAppExportPolicyRuleAllOf
	isSet bool
}

func (v NullableStorageNetAppExportPolicyRuleAllOf) Get() *StorageNetAppExportPolicyRuleAllOf {
	return v.value
}

func (v *NullableStorageNetAppExportPolicyRuleAllOf) Set(val *StorageNetAppExportPolicyRuleAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableStorageNetAppExportPolicyRuleAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableStorageNetAppExportPolicyRuleAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStorageNetAppExportPolicyRuleAllOf(val *StorageNetAppExportPolicyRuleAllOf) *NullableStorageNetAppExportPolicyRuleAllOf {
	return &NullableStorageNetAppExportPolicyRuleAllOf{value: val, isSet: true}
}

func (v NullableStorageNetAppExportPolicyRuleAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStorageNetAppExportPolicyRuleAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
