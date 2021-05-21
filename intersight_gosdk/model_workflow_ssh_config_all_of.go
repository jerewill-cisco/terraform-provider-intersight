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

// WorkflowSshConfigAllOf Definition of the list of properties defined in 'workflow.SshConfig', excluding properties defined in parent classes.
type WorkflowSshConfigAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Password to use in the SSH connection credentials (If empty then private key will be used).
	Password *string `json:"Password,omitempty"`
	// The remote server to connect to. IPv4 address represented in dot decimal notation.
	Target *string `json:"Target,omitempty"`
	// Username for the remote SSH connection.
	User                 *string `json:"User,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _WorkflowSshConfigAllOf WorkflowSshConfigAllOf

// NewWorkflowSshConfigAllOf instantiates a new WorkflowSshConfigAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkflowSshConfigAllOf(classId string, objectType string) *WorkflowSshConfigAllOf {
	this := WorkflowSshConfigAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewWorkflowSshConfigAllOfWithDefaults instantiates a new WorkflowSshConfigAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkflowSshConfigAllOfWithDefaults() *WorkflowSshConfigAllOf {
	this := WorkflowSshConfigAllOf{}
	var classId string = "workflow.SshConfig"
	this.ClassId = classId
	var objectType string = "workflow.SshConfig"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *WorkflowSshConfigAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *WorkflowSshConfigAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *WorkflowSshConfigAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *WorkflowSshConfigAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *WorkflowSshConfigAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *WorkflowSshConfigAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetPassword returns the Password field value if set, zero value otherwise.
func (o *WorkflowSshConfigAllOf) GetPassword() string {
	if o == nil || o.Password == nil {
		var ret string
		return ret
	}
	return *o.Password
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowSshConfigAllOf) GetPasswordOk() (*string, bool) {
	if o == nil || o.Password == nil {
		return nil, false
	}
	return o.Password, true
}

// HasPassword returns a boolean if a field has been set.
func (o *WorkflowSshConfigAllOf) HasPassword() bool {
	if o != nil && o.Password != nil {
		return true
	}

	return false
}

// SetPassword gets a reference to the given string and assigns it to the Password field.
func (o *WorkflowSshConfigAllOf) SetPassword(v string) {
	o.Password = &v
}

// GetTarget returns the Target field value if set, zero value otherwise.
func (o *WorkflowSshConfigAllOf) GetTarget() string {
	if o == nil || o.Target == nil {
		var ret string
		return ret
	}
	return *o.Target
}

// GetTargetOk returns a tuple with the Target field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowSshConfigAllOf) GetTargetOk() (*string, bool) {
	if o == nil || o.Target == nil {
		return nil, false
	}
	return o.Target, true
}

// HasTarget returns a boolean if a field has been set.
func (o *WorkflowSshConfigAllOf) HasTarget() bool {
	if o != nil && o.Target != nil {
		return true
	}

	return false
}

// SetTarget gets a reference to the given string and assigns it to the Target field.
func (o *WorkflowSshConfigAllOf) SetTarget(v string) {
	o.Target = &v
}

// GetUser returns the User field value if set, zero value otherwise.
func (o *WorkflowSshConfigAllOf) GetUser() string {
	if o == nil || o.User == nil {
		var ret string
		return ret
	}
	return *o.User
}

// GetUserOk returns a tuple with the User field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowSshConfigAllOf) GetUserOk() (*string, bool) {
	if o == nil || o.User == nil {
		return nil, false
	}
	return o.User, true
}

// HasUser returns a boolean if a field has been set.
func (o *WorkflowSshConfigAllOf) HasUser() bool {
	if o != nil && o.User != nil {
		return true
	}

	return false
}

// SetUser gets a reference to the given string and assigns it to the User field.
func (o *WorkflowSshConfigAllOf) SetUser(v string) {
	o.User = &v
}

func (o WorkflowSshConfigAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Password != nil {
		toSerialize["Password"] = o.Password
	}
	if o.Target != nil {
		toSerialize["Target"] = o.Target
	}
	if o.User != nil {
		toSerialize["User"] = o.User
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *WorkflowSshConfigAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varWorkflowSshConfigAllOf := _WorkflowSshConfigAllOf{}

	if err = json.Unmarshal(bytes, &varWorkflowSshConfigAllOf); err == nil {
		*o = WorkflowSshConfigAllOf(varWorkflowSshConfigAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Password")
		delete(additionalProperties, "Target")
		delete(additionalProperties, "User")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableWorkflowSshConfigAllOf struct {
	value *WorkflowSshConfigAllOf
	isSet bool
}

func (v NullableWorkflowSshConfigAllOf) Get() *WorkflowSshConfigAllOf {
	return v.value
}

func (v *NullableWorkflowSshConfigAllOf) Set(val *WorkflowSshConfigAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkflowSshConfigAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkflowSshConfigAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkflowSshConfigAllOf(val *WorkflowSshConfigAllOf) *NullableWorkflowSshConfigAllOf {
	return &NullableWorkflowSshConfigAllOf{value: val, isSet: true}
}

func (v NullableWorkflowSshConfigAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkflowSshConfigAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
