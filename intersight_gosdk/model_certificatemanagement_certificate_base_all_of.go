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

// CertificatemanagementCertificateBaseAllOf Definition of the list of properties defined in 'certificatemanagement.CertificateBase', excluding properties defined in parent classes.
type CertificatemanagementCertificateBaseAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property. The enum values provides the list of concrete types that can be instantiated from this abstract type.
	ObjectType  string                  `json:"ObjectType"`
	Certificate NullableX509Certificate `json:"Certificate,omitempty"`
	// Enable/Disable the certificate in Certificate Management policy.
	Enabled *bool `json:"Enabled,omitempty"`
	// Indicates whether the value of the 'privatekey' property has been set.
	IsPrivatekeySet *bool `json:"IsPrivatekeySet,omitempty"`
	// Private Key which is used to validate the certificate.
	Privatekey           *string `json:"Privatekey,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CertificatemanagementCertificateBaseAllOf CertificatemanagementCertificateBaseAllOf

// NewCertificatemanagementCertificateBaseAllOf instantiates a new CertificatemanagementCertificateBaseAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCertificatemanagementCertificateBaseAllOf(classId string, objectType string) *CertificatemanagementCertificateBaseAllOf {
	this := CertificatemanagementCertificateBaseAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	var enabled bool = true
	this.Enabled = &enabled
	var isPrivatekeySet bool = false
	this.IsPrivatekeySet = &isPrivatekeySet
	return &this
}

// NewCertificatemanagementCertificateBaseAllOfWithDefaults instantiates a new CertificatemanagementCertificateBaseAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCertificatemanagementCertificateBaseAllOfWithDefaults() *CertificatemanagementCertificateBaseAllOf {
	this := CertificatemanagementCertificateBaseAllOf{}
	var classId string = "certificatemanagement.Imc"
	this.ClassId = classId
	var objectType string = "certificatemanagement.Imc"
	this.ObjectType = objectType
	var enabled bool = true
	this.Enabled = &enabled
	var isPrivatekeySet bool = false
	this.IsPrivatekeySet = &isPrivatekeySet
	return &this
}

// GetClassId returns the ClassId field value
func (o *CertificatemanagementCertificateBaseAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *CertificatemanagementCertificateBaseAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *CertificatemanagementCertificateBaseAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *CertificatemanagementCertificateBaseAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *CertificatemanagementCertificateBaseAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *CertificatemanagementCertificateBaseAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetCertificate returns the Certificate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CertificatemanagementCertificateBaseAllOf) GetCertificate() X509Certificate {
	if o == nil || o.Certificate.Get() == nil {
		var ret X509Certificate
		return ret
	}
	return *o.Certificate.Get()
}

// GetCertificateOk returns a tuple with the Certificate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CertificatemanagementCertificateBaseAllOf) GetCertificateOk() (*X509Certificate, bool) {
	if o == nil {
		return nil, false
	}
	return o.Certificate.Get(), o.Certificate.IsSet()
}

// HasCertificate returns a boolean if a field has been set.
func (o *CertificatemanagementCertificateBaseAllOf) HasCertificate() bool {
	if o != nil && o.Certificate.IsSet() {
		return true
	}

	return false
}

// SetCertificate gets a reference to the given NullableX509Certificate and assigns it to the Certificate field.
func (o *CertificatemanagementCertificateBaseAllOf) SetCertificate(v X509Certificate) {
	o.Certificate.Set(&v)
}

// SetCertificateNil sets the value for Certificate to be an explicit nil
func (o *CertificatemanagementCertificateBaseAllOf) SetCertificateNil() {
	o.Certificate.Set(nil)
}

// UnsetCertificate ensures that no value is present for Certificate, not even an explicit nil
func (o *CertificatemanagementCertificateBaseAllOf) UnsetCertificate() {
	o.Certificate.Unset()
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *CertificatemanagementCertificateBaseAllOf) GetEnabled() bool {
	if o == nil || o.Enabled == nil {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CertificatemanagementCertificateBaseAllOf) GetEnabledOk() (*bool, bool) {
	if o == nil || o.Enabled == nil {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *CertificatemanagementCertificateBaseAllOf) HasEnabled() bool {
	if o != nil && o.Enabled != nil {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *CertificatemanagementCertificateBaseAllOf) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetIsPrivatekeySet returns the IsPrivatekeySet field value if set, zero value otherwise.
func (o *CertificatemanagementCertificateBaseAllOf) GetIsPrivatekeySet() bool {
	if o == nil || o.IsPrivatekeySet == nil {
		var ret bool
		return ret
	}
	return *o.IsPrivatekeySet
}

// GetIsPrivatekeySetOk returns a tuple with the IsPrivatekeySet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CertificatemanagementCertificateBaseAllOf) GetIsPrivatekeySetOk() (*bool, bool) {
	if o == nil || o.IsPrivatekeySet == nil {
		return nil, false
	}
	return o.IsPrivatekeySet, true
}

// HasIsPrivatekeySet returns a boolean if a field has been set.
func (o *CertificatemanagementCertificateBaseAllOf) HasIsPrivatekeySet() bool {
	if o != nil && o.IsPrivatekeySet != nil {
		return true
	}

	return false
}

// SetIsPrivatekeySet gets a reference to the given bool and assigns it to the IsPrivatekeySet field.
func (o *CertificatemanagementCertificateBaseAllOf) SetIsPrivatekeySet(v bool) {
	o.IsPrivatekeySet = &v
}

// GetPrivatekey returns the Privatekey field value if set, zero value otherwise.
func (o *CertificatemanagementCertificateBaseAllOf) GetPrivatekey() string {
	if o == nil || o.Privatekey == nil {
		var ret string
		return ret
	}
	return *o.Privatekey
}

// GetPrivatekeyOk returns a tuple with the Privatekey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CertificatemanagementCertificateBaseAllOf) GetPrivatekeyOk() (*string, bool) {
	if o == nil || o.Privatekey == nil {
		return nil, false
	}
	return o.Privatekey, true
}

// HasPrivatekey returns a boolean if a field has been set.
func (o *CertificatemanagementCertificateBaseAllOf) HasPrivatekey() bool {
	if o != nil && o.Privatekey != nil {
		return true
	}

	return false
}

// SetPrivatekey gets a reference to the given string and assigns it to the Privatekey field.
func (o *CertificatemanagementCertificateBaseAllOf) SetPrivatekey(v string) {
	o.Privatekey = &v
}

func (o CertificatemanagementCertificateBaseAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Certificate.IsSet() {
		toSerialize["Certificate"] = o.Certificate.Get()
	}
	if o.Enabled != nil {
		toSerialize["Enabled"] = o.Enabled
	}
	if o.IsPrivatekeySet != nil {
		toSerialize["IsPrivatekeySet"] = o.IsPrivatekeySet
	}
	if o.Privatekey != nil {
		toSerialize["Privatekey"] = o.Privatekey
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *CertificatemanagementCertificateBaseAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varCertificatemanagementCertificateBaseAllOf := _CertificatemanagementCertificateBaseAllOf{}

	if err = json.Unmarshal(bytes, &varCertificatemanagementCertificateBaseAllOf); err == nil {
		*o = CertificatemanagementCertificateBaseAllOf(varCertificatemanagementCertificateBaseAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Certificate")
		delete(additionalProperties, "Enabled")
		delete(additionalProperties, "IsPrivatekeySet")
		delete(additionalProperties, "Privatekey")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCertificatemanagementCertificateBaseAllOf struct {
	value *CertificatemanagementCertificateBaseAllOf
	isSet bool
}

func (v NullableCertificatemanagementCertificateBaseAllOf) Get() *CertificatemanagementCertificateBaseAllOf {
	return v.value
}

func (v *NullableCertificatemanagementCertificateBaseAllOf) Set(val *CertificatemanagementCertificateBaseAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableCertificatemanagementCertificateBaseAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableCertificatemanagementCertificateBaseAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCertificatemanagementCertificateBaseAllOf(val *CertificatemanagementCertificateBaseAllOf) *NullableCertificatemanagementCertificateBaseAllOf {
	return &NullableCertificatemanagementCertificateBaseAllOf{value: val, isSet: true}
}

func (v NullableCertificatemanagementCertificateBaseAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCertificatemanagementCertificateBaseAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
