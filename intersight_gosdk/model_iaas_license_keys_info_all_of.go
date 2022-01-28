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

// IaasLicenseKeysInfoAllOf Definition of the list of properties defined in 'iaas.LicenseKeysInfo', excluding properties defined in parent classes.
type IaasLicenseKeysInfoAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Number of licenses available for the UCSD PID (Product ID).
	Count *int64 `json:"Count,omitempty"`
	// Expiration date for the license.
	ExpirationDate *string `json:"ExpirationDate,omitempty"`
	// UCS Director Unique license ID.
	LicenseId *string `json:"LicenseId,omitempty"`
	// PID (Product ID) for UCSD License.
	Pid                  *string `json:"Pid,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _IaasLicenseKeysInfoAllOf IaasLicenseKeysInfoAllOf

// NewIaasLicenseKeysInfoAllOf instantiates a new IaasLicenseKeysInfoAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIaasLicenseKeysInfoAllOf(classId string, objectType string) *IaasLicenseKeysInfoAllOf {
	this := IaasLicenseKeysInfoAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewIaasLicenseKeysInfoAllOfWithDefaults instantiates a new IaasLicenseKeysInfoAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIaasLicenseKeysInfoAllOfWithDefaults() *IaasLicenseKeysInfoAllOf {
	this := IaasLicenseKeysInfoAllOf{}
	var classId string = "iaas.LicenseKeysInfo"
	this.ClassId = classId
	var objectType string = "iaas.LicenseKeysInfo"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *IaasLicenseKeysInfoAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *IaasLicenseKeysInfoAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *IaasLicenseKeysInfoAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *IaasLicenseKeysInfoAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *IaasLicenseKeysInfoAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *IaasLicenseKeysInfoAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *IaasLicenseKeysInfoAllOf) GetCount() int64 {
	if o == nil || o.Count == nil {
		var ret int64
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IaasLicenseKeysInfoAllOf) GetCountOk() (*int64, bool) {
	if o == nil || o.Count == nil {
		return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *IaasLicenseKeysInfoAllOf) HasCount() bool {
	if o != nil && o.Count != nil {
		return true
	}

	return false
}

// SetCount gets a reference to the given int64 and assigns it to the Count field.
func (o *IaasLicenseKeysInfoAllOf) SetCount(v int64) {
	o.Count = &v
}

// GetExpirationDate returns the ExpirationDate field value if set, zero value otherwise.
func (o *IaasLicenseKeysInfoAllOf) GetExpirationDate() string {
	if o == nil || o.ExpirationDate == nil {
		var ret string
		return ret
	}
	return *o.ExpirationDate
}

// GetExpirationDateOk returns a tuple with the ExpirationDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IaasLicenseKeysInfoAllOf) GetExpirationDateOk() (*string, bool) {
	if o == nil || o.ExpirationDate == nil {
		return nil, false
	}
	return o.ExpirationDate, true
}

// HasExpirationDate returns a boolean if a field has been set.
func (o *IaasLicenseKeysInfoAllOf) HasExpirationDate() bool {
	if o != nil && o.ExpirationDate != nil {
		return true
	}

	return false
}

// SetExpirationDate gets a reference to the given string and assigns it to the ExpirationDate field.
func (o *IaasLicenseKeysInfoAllOf) SetExpirationDate(v string) {
	o.ExpirationDate = &v
}

// GetLicenseId returns the LicenseId field value if set, zero value otherwise.
func (o *IaasLicenseKeysInfoAllOf) GetLicenseId() string {
	if o == nil || o.LicenseId == nil {
		var ret string
		return ret
	}
	return *o.LicenseId
}

// GetLicenseIdOk returns a tuple with the LicenseId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IaasLicenseKeysInfoAllOf) GetLicenseIdOk() (*string, bool) {
	if o == nil || o.LicenseId == nil {
		return nil, false
	}
	return o.LicenseId, true
}

// HasLicenseId returns a boolean if a field has been set.
func (o *IaasLicenseKeysInfoAllOf) HasLicenseId() bool {
	if o != nil && o.LicenseId != nil {
		return true
	}

	return false
}

// SetLicenseId gets a reference to the given string and assigns it to the LicenseId field.
func (o *IaasLicenseKeysInfoAllOf) SetLicenseId(v string) {
	o.LicenseId = &v
}

// GetPid returns the Pid field value if set, zero value otherwise.
func (o *IaasLicenseKeysInfoAllOf) GetPid() string {
	if o == nil || o.Pid == nil {
		var ret string
		return ret
	}
	return *o.Pid
}

// GetPidOk returns a tuple with the Pid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IaasLicenseKeysInfoAllOf) GetPidOk() (*string, bool) {
	if o == nil || o.Pid == nil {
		return nil, false
	}
	return o.Pid, true
}

// HasPid returns a boolean if a field has been set.
func (o *IaasLicenseKeysInfoAllOf) HasPid() bool {
	if o != nil && o.Pid != nil {
		return true
	}

	return false
}

// SetPid gets a reference to the given string and assigns it to the Pid field.
func (o *IaasLicenseKeysInfoAllOf) SetPid(v string) {
	o.Pid = &v
}

func (o IaasLicenseKeysInfoAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Count != nil {
		toSerialize["Count"] = o.Count
	}
	if o.ExpirationDate != nil {
		toSerialize["ExpirationDate"] = o.ExpirationDate
	}
	if o.LicenseId != nil {
		toSerialize["LicenseId"] = o.LicenseId
	}
	if o.Pid != nil {
		toSerialize["Pid"] = o.Pid
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *IaasLicenseKeysInfoAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varIaasLicenseKeysInfoAllOf := _IaasLicenseKeysInfoAllOf{}

	if err = json.Unmarshal(bytes, &varIaasLicenseKeysInfoAllOf); err == nil {
		*o = IaasLicenseKeysInfoAllOf(varIaasLicenseKeysInfoAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Count")
		delete(additionalProperties, "ExpirationDate")
		delete(additionalProperties, "LicenseId")
		delete(additionalProperties, "Pid")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableIaasLicenseKeysInfoAllOf struct {
	value *IaasLicenseKeysInfoAllOf
	isSet bool
}

func (v NullableIaasLicenseKeysInfoAllOf) Get() *IaasLicenseKeysInfoAllOf {
	return v.value
}

func (v *NullableIaasLicenseKeysInfoAllOf) Set(val *IaasLicenseKeysInfoAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableIaasLicenseKeysInfoAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableIaasLicenseKeysInfoAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIaasLicenseKeysInfoAllOf(val *IaasLicenseKeysInfoAllOf) *NullableIaasLicenseKeysInfoAllOf {
	return &NullableIaasLicenseKeysInfoAllOf{value: val, isSet: true}
}

func (v NullableIaasLicenseKeysInfoAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIaasLicenseKeysInfoAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
