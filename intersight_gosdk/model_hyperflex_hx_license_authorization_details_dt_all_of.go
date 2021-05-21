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

// HyperflexHxLicenseAuthorizationDetailsDtAllOf Definition of the list of properties defined in 'hyperflex.HxLicenseAuthorizationDetailsDt', excluding properties defined in parent classes.
type HyperflexHxLicenseAuthorizationDetailsDtAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Deadline date for next communication
	CommunicationDeadlineDate *string `json:"CommunicationDeadlineDate,omitempty"`
	// Evaluation period expired date
	EvaluationPeriodExpiredAt *string `json:"EvaluationPeriodExpiredAt,omitempty"`
	// Remaining evaluation period
	EvaluationPeriodRemaining *string `json:"EvaluationPeriodRemaining,omitempty"`
	// Last Communication Attempt Date
	LastCommunicationAttemptDate *string `json:"LastCommunicationAttemptDate,omitempty"`
	// Timestamp of Next Communication Attempt
	NextCommunicationAttemptDate *string `json:"NextCommunicationAttemptDate,omitempty"`
	// License Authorization Status
	Status               *string `json:"Status,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _HyperflexHxLicenseAuthorizationDetailsDtAllOf HyperflexHxLicenseAuthorizationDetailsDtAllOf

// NewHyperflexHxLicenseAuthorizationDetailsDtAllOf instantiates a new HyperflexHxLicenseAuthorizationDetailsDtAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHyperflexHxLicenseAuthorizationDetailsDtAllOf(classId string, objectType string) *HyperflexHxLicenseAuthorizationDetailsDtAllOf {
	this := HyperflexHxLicenseAuthorizationDetailsDtAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewHyperflexHxLicenseAuthorizationDetailsDtAllOfWithDefaults instantiates a new HyperflexHxLicenseAuthorizationDetailsDtAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHyperflexHxLicenseAuthorizationDetailsDtAllOfWithDefaults() *HyperflexHxLicenseAuthorizationDetailsDtAllOf {
	this := HyperflexHxLicenseAuthorizationDetailsDtAllOf{}
	var classId string = "hyperflex.HxLicenseAuthorizationDetailsDt"
	this.ClassId = classId
	var objectType string = "hyperflex.HxLicenseAuthorizationDetailsDt"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *HyperflexHxLicenseAuthorizationDetailsDtAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *HyperflexHxLicenseAuthorizationDetailsDtAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *HyperflexHxLicenseAuthorizationDetailsDtAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *HyperflexHxLicenseAuthorizationDetailsDtAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *HyperflexHxLicenseAuthorizationDetailsDtAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *HyperflexHxLicenseAuthorizationDetailsDtAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetCommunicationDeadlineDate returns the CommunicationDeadlineDate field value if set, zero value otherwise.
func (o *HyperflexHxLicenseAuthorizationDetailsDtAllOf) GetCommunicationDeadlineDate() string {
	if o == nil || o.CommunicationDeadlineDate == nil {
		var ret string
		return ret
	}
	return *o.CommunicationDeadlineDate
}

// GetCommunicationDeadlineDateOk returns a tuple with the CommunicationDeadlineDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexHxLicenseAuthorizationDetailsDtAllOf) GetCommunicationDeadlineDateOk() (*string, bool) {
	if o == nil || o.CommunicationDeadlineDate == nil {
		return nil, false
	}
	return o.CommunicationDeadlineDate, true
}

// HasCommunicationDeadlineDate returns a boolean if a field has been set.
func (o *HyperflexHxLicenseAuthorizationDetailsDtAllOf) HasCommunicationDeadlineDate() bool {
	if o != nil && o.CommunicationDeadlineDate != nil {
		return true
	}

	return false
}

// SetCommunicationDeadlineDate gets a reference to the given string and assigns it to the CommunicationDeadlineDate field.
func (o *HyperflexHxLicenseAuthorizationDetailsDtAllOf) SetCommunicationDeadlineDate(v string) {
	o.CommunicationDeadlineDate = &v
}

// GetEvaluationPeriodExpiredAt returns the EvaluationPeriodExpiredAt field value if set, zero value otherwise.
func (o *HyperflexHxLicenseAuthorizationDetailsDtAllOf) GetEvaluationPeriodExpiredAt() string {
	if o == nil || o.EvaluationPeriodExpiredAt == nil {
		var ret string
		return ret
	}
	return *o.EvaluationPeriodExpiredAt
}

// GetEvaluationPeriodExpiredAtOk returns a tuple with the EvaluationPeriodExpiredAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexHxLicenseAuthorizationDetailsDtAllOf) GetEvaluationPeriodExpiredAtOk() (*string, bool) {
	if o == nil || o.EvaluationPeriodExpiredAt == nil {
		return nil, false
	}
	return o.EvaluationPeriodExpiredAt, true
}

// HasEvaluationPeriodExpiredAt returns a boolean if a field has been set.
func (o *HyperflexHxLicenseAuthorizationDetailsDtAllOf) HasEvaluationPeriodExpiredAt() bool {
	if o != nil && o.EvaluationPeriodExpiredAt != nil {
		return true
	}

	return false
}

// SetEvaluationPeriodExpiredAt gets a reference to the given string and assigns it to the EvaluationPeriodExpiredAt field.
func (o *HyperflexHxLicenseAuthorizationDetailsDtAllOf) SetEvaluationPeriodExpiredAt(v string) {
	o.EvaluationPeriodExpiredAt = &v
}

// GetEvaluationPeriodRemaining returns the EvaluationPeriodRemaining field value if set, zero value otherwise.
func (o *HyperflexHxLicenseAuthorizationDetailsDtAllOf) GetEvaluationPeriodRemaining() string {
	if o == nil || o.EvaluationPeriodRemaining == nil {
		var ret string
		return ret
	}
	return *o.EvaluationPeriodRemaining
}

// GetEvaluationPeriodRemainingOk returns a tuple with the EvaluationPeriodRemaining field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexHxLicenseAuthorizationDetailsDtAllOf) GetEvaluationPeriodRemainingOk() (*string, bool) {
	if o == nil || o.EvaluationPeriodRemaining == nil {
		return nil, false
	}
	return o.EvaluationPeriodRemaining, true
}

// HasEvaluationPeriodRemaining returns a boolean if a field has been set.
func (o *HyperflexHxLicenseAuthorizationDetailsDtAllOf) HasEvaluationPeriodRemaining() bool {
	if o != nil && o.EvaluationPeriodRemaining != nil {
		return true
	}

	return false
}

// SetEvaluationPeriodRemaining gets a reference to the given string and assigns it to the EvaluationPeriodRemaining field.
func (o *HyperflexHxLicenseAuthorizationDetailsDtAllOf) SetEvaluationPeriodRemaining(v string) {
	o.EvaluationPeriodRemaining = &v
}

// GetLastCommunicationAttemptDate returns the LastCommunicationAttemptDate field value if set, zero value otherwise.
func (o *HyperflexHxLicenseAuthorizationDetailsDtAllOf) GetLastCommunicationAttemptDate() string {
	if o == nil || o.LastCommunicationAttemptDate == nil {
		var ret string
		return ret
	}
	return *o.LastCommunicationAttemptDate
}

// GetLastCommunicationAttemptDateOk returns a tuple with the LastCommunicationAttemptDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexHxLicenseAuthorizationDetailsDtAllOf) GetLastCommunicationAttemptDateOk() (*string, bool) {
	if o == nil || o.LastCommunicationAttemptDate == nil {
		return nil, false
	}
	return o.LastCommunicationAttemptDate, true
}

// HasLastCommunicationAttemptDate returns a boolean if a field has been set.
func (o *HyperflexHxLicenseAuthorizationDetailsDtAllOf) HasLastCommunicationAttemptDate() bool {
	if o != nil && o.LastCommunicationAttemptDate != nil {
		return true
	}

	return false
}

// SetLastCommunicationAttemptDate gets a reference to the given string and assigns it to the LastCommunicationAttemptDate field.
func (o *HyperflexHxLicenseAuthorizationDetailsDtAllOf) SetLastCommunicationAttemptDate(v string) {
	o.LastCommunicationAttemptDate = &v
}

// GetNextCommunicationAttemptDate returns the NextCommunicationAttemptDate field value if set, zero value otherwise.
func (o *HyperflexHxLicenseAuthorizationDetailsDtAllOf) GetNextCommunicationAttemptDate() string {
	if o == nil || o.NextCommunicationAttemptDate == nil {
		var ret string
		return ret
	}
	return *o.NextCommunicationAttemptDate
}

// GetNextCommunicationAttemptDateOk returns a tuple with the NextCommunicationAttemptDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexHxLicenseAuthorizationDetailsDtAllOf) GetNextCommunicationAttemptDateOk() (*string, bool) {
	if o == nil || o.NextCommunicationAttemptDate == nil {
		return nil, false
	}
	return o.NextCommunicationAttemptDate, true
}

// HasNextCommunicationAttemptDate returns a boolean if a field has been set.
func (o *HyperflexHxLicenseAuthorizationDetailsDtAllOf) HasNextCommunicationAttemptDate() bool {
	if o != nil && o.NextCommunicationAttemptDate != nil {
		return true
	}

	return false
}

// SetNextCommunicationAttemptDate gets a reference to the given string and assigns it to the NextCommunicationAttemptDate field.
func (o *HyperflexHxLicenseAuthorizationDetailsDtAllOf) SetNextCommunicationAttemptDate(v string) {
	o.NextCommunicationAttemptDate = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *HyperflexHxLicenseAuthorizationDetailsDtAllOf) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexHxLicenseAuthorizationDetailsDtAllOf) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *HyperflexHxLicenseAuthorizationDetailsDtAllOf) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *HyperflexHxLicenseAuthorizationDetailsDtAllOf) SetStatus(v string) {
	o.Status = &v
}

func (o HyperflexHxLicenseAuthorizationDetailsDtAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.CommunicationDeadlineDate != nil {
		toSerialize["CommunicationDeadlineDate"] = o.CommunicationDeadlineDate
	}
	if o.EvaluationPeriodExpiredAt != nil {
		toSerialize["EvaluationPeriodExpiredAt"] = o.EvaluationPeriodExpiredAt
	}
	if o.EvaluationPeriodRemaining != nil {
		toSerialize["EvaluationPeriodRemaining"] = o.EvaluationPeriodRemaining
	}
	if o.LastCommunicationAttemptDate != nil {
		toSerialize["LastCommunicationAttemptDate"] = o.LastCommunicationAttemptDate
	}
	if o.NextCommunicationAttemptDate != nil {
		toSerialize["NextCommunicationAttemptDate"] = o.NextCommunicationAttemptDate
	}
	if o.Status != nil {
		toSerialize["Status"] = o.Status
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *HyperflexHxLicenseAuthorizationDetailsDtAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varHyperflexHxLicenseAuthorizationDetailsDtAllOf := _HyperflexHxLicenseAuthorizationDetailsDtAllOf{}

	if err = json.Unmarshal(bytes, &varHyperflexHxLicenseAuthorizationDetailsDtAllOf); err == nil {
		*o = HyperflexHxLicenseAuthorizationDetailsDtAllOf(varHyperflexHxLicenseAuthorizationDetailsDtAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "CommunicationDeadlineDate")
		delete(additionalProperties, "EvaluationPeriodExpiredAt")
		delete(additionalProperties, "EvaluationPeriodRemaining")
		delete(additionalProperties, "LastCommunicationAttemptDate")
		delete(additionalProperties, "NextCommunicationAttemptDate")
		delete(additionalProperties, "Status")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableHyperflexHxLicenseAuthorizationDetailsDtAllOf struct {
	value *HyperflexHxLicenseAuthorizationDetailsDtAllOf
	isSet bool
}

func (v NullableHyperflexHxLicenseAuthorizationDetailsDtAllOf) Get() *HyperflexHxLicenseAuthorizationDetailsDtAllOf {
	return v.value
}

func (v *NullableHyperflexHxLicenseAuthorizationDetailsDtAllOf) Set(val *HyperflexHxLicenseAuthorizationDetailsDtAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableHyperflexHxLicenseAuthorizationDetailsDtAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableHyperflexHxLicenseAuthorizationDetailsDtAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHyperflexHxLicenseAuthorizationDetailsDtAllOf(val *HyperflexHxLicenseAuthorizationDetailsDtAllOf) *NullableHyperflexHxLicenseAuthorizationDetailsDtAllOf {
	return &NullableHyperflexHxLicenseAuthorizationDetailsDtAllOf{value: val, isSet: true}
}

func (v NullableHyperflexHxLicenseAuthorizationDetailsDtAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHyperflexHxLicenseAuthorizationDetailsDtAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
