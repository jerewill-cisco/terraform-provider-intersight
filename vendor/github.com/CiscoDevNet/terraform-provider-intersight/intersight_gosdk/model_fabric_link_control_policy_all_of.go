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

// FabricLinkControlPolicyAllOf Definition of the list of properties defined in 'fabric.LinkControlPolicy', excluding properties defined in parent classes.
type FabricLinkControlPolicyAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType           string                                `json:"ObjectType"`
	UdldSettings         NullableFabricUdldSettings            `json:"UdldSettings,omitempty"`
	Organization         *OrganizationOrganizationRelationship `json:"Organization,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _FabricLinkControlPolicyAllOf FabricLinkControlPolicyAllOf

// NewFabricLinkControlPolicyAllOf instantiates a new FabricLinkControlPolicyAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFabricLinkControlPolicyAllOf(classId string, objectType string) *FabricLinkControlPolicyAllOf {
	this := FabricLinkControlPolicyAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewFabricLinkControlPolicyAllOfWithDefaults instantiates a new FabricLinkControlPolicyAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFabricLinkControlPolicyAllOfWithDefaults() *FabricLinkControlPolicyAllOf {
	this := FabricLinkControlPolicyAllOf{}
	var classId string = "fabric.LinkControlPolicy"
	this.ClassId = classId
	var objectType string = "fabric.LinkControlPolicy"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *FabricLinkControlPolicyAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *FabricLinkControlPolicyAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *FabricLinkControlPolicyAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *FabricLinkControlPolicyAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *FabricLinkControlPolicyAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *FabricLinkControlPolicyAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetUdldSettings returns the UdldSettings field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FabricLinkControlPolicyAllOf) GetUdldSettings() FabricUdldSettings {
	if o == nil || o.UdldSettings.Get() == nil {
		var ret FabricUdldSettings
		return ret
	}
	return *o.UdldSettings.Get()
}

// GetUdldSettingsOk returns a tuple with the UdldSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FabricLinkControlPolicyAllOf) GetUdldSettingsOk() (*FabricUdldSettings, bool) {
	if o == nil {
		return nil, false
	}
	return o.UdldSettings.Get(), o.UdldSettings.IsSet()
}

// HasUdldSettings returns a boolean if a field has been set.
func (o *FabricLinkControlPolicyAllOf) HasUdldSettings() bool {
	if o != nil && o.UdldSettings.IsSet() {
		return true
	}

	return false
}

// SetUdldSettings gets a reference to the given NullableFabricUdldSettings and assigns it to the UdldSettings field.
func (o *FabricLinkControlPolicyAllOf) SetUdldSettings(v FabricUdldSettings) {
	o.UdldSettings.Set(&v)
}

// SetUdldSettingsNil sets the value for UdldSettings to be an explicit nil
func (o *FabricLinkControlPolicyAllOf) SetUdldSettingsNil() {
	o.UdldSettings.Set(nil)
}

// UnsetUdldSettings ensures that no value is present for UdldSettings, not even an explicit nil
func (o *FabricLinkControlPolicyAllOf) UnsetUdldSettings() {
	o.UdldSettings.Unset()
}

// GetOrganization returns the Organization field value if set, zero value otherwise.
func (o *FabricLinkControlPolicyAllOf) GetOrganization() OrganizationOrganizationRelationship {
	if o == nil || o.Organization == nil {
		var ret OrganizationOrganizationRelationship
		return ret
	}
	return *o.Organization
}

// GetOrganizationOk returns a tuple with the Organization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FabricLinkControlPolicyAllOf) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool) {
	if o == nil || o.Organization == nil {
		return nil, false
	}
	return o.Organization, true
}

// HasOrganization returns a boolean if a field has been set.
func (o *FabricLinkControlPolicyAllOf) HasOrganization() bool {
	if o != nil && o.Organization != nil {
		return true
	}

	return false
}

// SetOrganization gets a reference to the given OrganizationOrganizationRelationship and assigns it to the Organization field.
func (o *FabricLinkControlPolicyAllOf) SetOrganization(v OrganizationOrganizationRelationship) {
	o.Organization = &v
}

func (o FabricLinkControlPolicyAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.UdldSettings.IsSet() {
		toSerialize["UdldSettings"] = o.UdldSettings.Get()
	}
	if o.Organization != nil {
		toSerialize["Organization"] = o.Organization
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *FabricLinkControlPolicyAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varFabricLinkControlPolicyAllOf := _FabricLinkControlPolicyAllOf{}

	if err = json.Unmarshal(bytes, &varFabricLinkControlPolicyAllOf); err == nil {
		*o = FabricLinkControlPolicyAllOf(varFabricLinkControlPolicyAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "UdldSettings")
		delete(additionalProperties, "Organization")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableFabricLinkControlPolicyAllOf struct {
	value *FabricLinkControlPolicyAllOf
	isSet bool
}

func (v NullableFabricLinkControlPolicyAllOf) Get() *FabricLinkControlPolicyAllOf {
	return v.value
}

func (v *NullableFabricLinkControlPolicyAllOf) Set(val *FabricLinkControlPolicyAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableFabricLinkControlPolicyAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableFabricLinkControlPolicyAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFabricLinkControlPolicyAllOf(val *FabricLinkControlPolicyAllOf) *NullableFabricLinkControlPolicyAllOf {
	return &NullableFabricLinkControlPolicyAllOf{value: val, isSet: true}
}

func (v NullableFabricLinkControlPolicyAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFabricLinkControlPolicyAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
