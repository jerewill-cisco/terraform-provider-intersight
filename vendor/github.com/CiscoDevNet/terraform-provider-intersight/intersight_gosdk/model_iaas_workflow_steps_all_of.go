/*
 * Cisco Intersight
 *
 * Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document. This document was created on 2021-06-30T12:14:04Z.
 *
 * API version: 1.0.9-4375
 * Contact: intersight@cisco.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package intersight

import (
	"encoding/json"
)

// IaasWorkflowStepsAllOf Definition of the list of properties defined in 'iaas.WorkflowSteps', excluding properties defined in parent classes.
type IaasWorkflowStepsAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Completed time of the workflow step.
	CompletedTime *string `json:"CompletedTime,omitempty"`
	// Name of the workflow step.
	Name *string `json:"Name,omitempty"`
	// Status of the workflow step.
	Status *string `json:"Status,omitempty"`
	// Status message of the workflow step.
	StatusMessage        *string `json:"StatusMessage,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _IaasWorkflowStepsAllOf IaasWorkflowStepsAllOf

// NewIaasWorkflowStepsAllOf instantiates a new IaasWorkflowStepsAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIaasWorkflowStepsAllOf(classId string, objectType string) *IaasWorkflowStepsAllOf {
	this := IaasWorkflowStepsAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewIaasWorkflowStepsAllOfWithDefaults instantiates a new IaasWorkflowStepsAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIaasWorkflowStepsAllOfWithDefaults() *IaasWorkflowStepsAllOf {
	this := IaasWorkflowStepsAllOf{}
	var classId string = "iaas.WorkflowSteps"
	this.ClassId = classId
	var objectType string = "iaas.WorkflowSteps"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *IaasWorkflowStepsAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *IaasWorkflowStepsAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *IaasWorkflowStepsAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *IaasWorkflowStepsAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *IaasWorkflowStepsAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *IaasWorkflowStepsAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetCompletedTime returns the CompletedTime field value if set, zero value otherwise.
func (o *IaasWorkflowStepsAllOf) GetCompletedTime() string {
	if o == nil || o.CompletedTime == nil {
		var ret string
		return ret
	}
	return *o.CompletedTime
}

// GetCompletedTimeOk returns a tuple with the CompletedTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IaasWorkflowStepsAllOf) GetCompletedTimeOk() (*string, bool) {
	if o == nil || o.CompletedTime == nil {
		return nil, false
	}
	return o.CompletedTime, true
}

// HasCompletedTime returns a boolean if a field has been set.
func (o *IaasWorkflowStepsAllOf) HasCompletedTime() bool {
	if o != nil && o.CompletedTime != nil {
		return true
	}

	return false
}

// SetCompletedTime gets a reference to the given string and assigns it to the CompletedTime field.
func (o *IaasWorkflowStepsAllOf) SetCompletedTime(v string) {
	o.CompletedTime = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *IaasWorkflowStepsAllOf) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IaasWorkflowStepsAllOf) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *IaasWorkflowStepsAllOf) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *IaasWorkflowStepsAllOf) SetName(v string) {
	o.Name = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *IaasWorkflowStepsAllOf) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IaasWorkflowStepsAllOf) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *IaasWorkflowStepsAllOf) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *IaasWorkflowStepsAllOf) SetStatus(v string) {
	o.Status = &v
}

// GetStatusMessage returns the StatusMessage field value if set, zero value otherwise.
func (o *IaasWorkflowStepsAllOf) GetStatusMessage() string {
	if o == nil || o.StatusMessage == nil {
		var ret string
		return ret
	}
	return *o.StatusMessage
}

// GetStatusMessageOk returns a tuple with the StatusMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IaasWorkflowStepsAllOf) GetStatusMessageOk() (*string, bool) {
	if o == nil || o.StatusMessage == nil {
		return nil, false
	}
	return o.StatusMessage, true
}

// HasStatusMessage returns a boolean if a field has been set.
func (o *IaasWorkflowStepsAllOf) HasStatusMessage() bool {
	if o != nil && o.StatusMessage != nil {
		return true
	}

	return false
}

// SetStatusMessage gets a reference to the given string and assigns it to the StatusMessage field.
func (o *IaasWorkflowStepsAllOf) SetStatusMessage(v string) {
	o.StatusMessage = &v
}

func (o IaasWorkflowStepsAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.CompletedTime != nil {
		toSerialize["CompletedTime"] = o.CompletedTime
	}
	if o.Name != nil {
		toSerialize["Name"] = o.Name
	}
	if o.Status != nil {
		toSerialize["Status"] = o.Status
	}
	if o.StatusMessage != nil {
		toSerialize["StatusMessage"] = o.StatusMessage
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *IaasWorkflowStepsAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varIaasWorkflowStepsAllOf := _IaasWorkflowStepsAllOf{}

	if err = json.Unmarshal(bytes, &varIaasWorkflowStepsAllOf); err == nil {
		*o = IaasWorkflowStepsAllOf(varIaasWorkflowStepsAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "CompletedTime")
		delete(additionalProperties, "Name")
		delete(additionalProperties, "Status")
		delete(additionalProperties, "StatusMessage")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableIaasWorkflowStepsAllOf struct {
	value *IaasWorkflowStepsAllOf
	isSet bool
}

func (v NullableIaasWorkflowStepsAllOf) Get() *IaasWorkflowStepsAllOf {
	return v.value
}

func (v *NullableIaasWorkflowStepsAllOf) Set(val *IaasWorkflowStepsAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableIaasWorkflowStepsAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableIaasWorkflowStepsAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIaasWorkflowStepsAllOf(val *IaasWorkflowStepsAllOf) *NullableIaasWorkflowStepsAllOf {
	return &NullableIaasWorkflowStepsAllOf{value: val, isSet: true}
}

func (v NullableIaasWorkflowStepsAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIaasWorkflowStepsAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
