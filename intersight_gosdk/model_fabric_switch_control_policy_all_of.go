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

// FabricSwitchControlPolicyAllOf Definition of the list of properties defined in 'fabric.SwitchControlPolicy', excluding properties defined in parent classes.
type FabricSwitchControlPolicyAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType       string                           `json:"ObjectType"`
	MacAgingSettings NullableFabricMacAgingSettings   `json:"MacAgingSettings,omitempty"`
	UdldSettings     NullableFabricUdldGlobalSettings `json:"UdldSettings,omitempty"`
	// To enable or disable the VLAN port count optimization.
	VlanPortOptimizationEnabled *bool                                 `json:"VlanPortOptimizationEnabled,omitempty"`
	Organization                *OrganizationOrganizationRelationship `json:"Organization,omitempty"`
	// An array of relationships to fabricSwitchProfile resources.
	Profiles             []FabricSwitchProfileRelationship `json:"Profiles,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _FabricSwitchControlPolicyAllOf FabricSwitchControlPolicyAllOf

// NewFabricSwitchControlPolicyAllOf instantiates a new FabricSwitchControlPolicyAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFabricSwitchControlPolicyAllOf(classId string, objectType string) *FabricSwitchControlPolicyAllOf {
	this := FabricSwitchControlPolicyAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	var vlanPortOptimizationEnabled bool = false
	this.VlanPortOptimizationEnabled = &vlanPortOptimizationEnabled
	return &this
}

// NewFabricSwitchControlPolicyAllOfWithDefaults instantiates a new FabricSwitchControlPolicyAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFabricSwitchControlPolicyAllOfWithDefaults() *FabricSwitchControlPolicyAllOf {
	this := FabricSwitchControlPolicyAllOf{}
	var classId string = "fabric.SwitchControlPolicy"
	this.ClassId = classId
	var objectType string = "fabric.SwitchControlPolicy"
	this.ObjectType = objectType
	var vlanPortOptimizationEnabled bool = false
	this.VlanPortOptimizationEnabled = &vlanPortOptimizationEnabled
	return &this
}

// GetClassId returns the ClassId field value
func (o *FabricSwitchControlPolicyAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *FabricSwitchControlPolicyAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *FabricSwitchControlPolicyAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *FabricSwitchControlPolicyAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *FabricSwitchControlPolicyAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *FabricSwitchControlPolicyAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetMacAgingSettings returns the MacAgingSettings field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FabricSwitchControlPolicyAllOf) GetMacAgingSettings() FabricMacAgingSettings {
	if o == nil || o.MacAgingSettings.Get() == nil {
		var ret FabricMacAgingSettings
		return ret
	}
	return *o.MacAgingSettings.Get()
}

// GetMacAgingSettingsOk returns a tuple with the MacAgingSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FabricSwitchControlPolicyAllOf) GetMacAgingSettingsOk() (*FabricMacAgingSettings, bool) {
	if o == nil {
		return nil, false
	}
	return o.MacAgingSettings.Get(), o.MacAgingSettings.IsSet()
}

// HasMacAgingSettings returns a boolean if a field has been set.
func (o *FabricSwitchControlPolicyAllOf) HasMacAgingSettings() bool {
	if o != nil && o.MacAgingSettings.IsSet() {
		return true
	}

	return false
}

// SetMacAgingSettings gets a reference to the given NullableFabricMacAgingSettings and assigns it to the MacAgingSettings field.
func (o *FabricSwitchControlPolicyAllOf) SetMacAgingSettings(v FabricMacAgingSettings) {
	o.MacAgingSettings.Set(&v)
}

// SetMacAgingSettingsNil sets the value for MacAgingSettings to be an explicit nil
func (o *FabricSwitchControlPolicyAllOf) SetMacAgingSettingsNil() {
	o.MacAgingSettings.Set(nil)
}

// UnsetMacAgingSettings ensures that no value is present for MacAgingSettings, not even an explicit nil
func (o *FabricSwitchControlPolicyAllOf) UnsetMacAgingSettings() {
	o.MacAgingSettings.Unset()
}

// GetUdldSettings returns the UdldSettings field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FabricSwitchControlPolicyAllOf) GetUdldSettings() FabricUdldGlobalSettings {
	if o == nil || o.UdldSettings.Get() == nil {
		var ret FabricUdldGlobalSettings
		return ret
	}
	return *o.UdldSettings.Get()
}

// GetUdldSettingsOk returns a tuple with the UdldSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FabricSwitchControlPolicyAllOf) GetUdldSettingsOk() (*FabricUdldGlobalSettings, bool) {
	if o == nil {
		return nil, false
	}
	return o.UdldSettings.Get(), o.UdldSettings.IsSet()
}

// HasUdldSettings returns a boolean if a field has been set.
func (o *FabricSwitchControlPolicyAllOf) HasUdldSettings() bool {
	if o != nil && o.UdldSettings.IsSet() {
		return true
	}

	return false
}

// SetUdldSettings gets a reference to the given NullableFabricUdldGlobalSettings and assigns it to the UdldSettings field.
func (o *FabricSwitchControlPolicyAllOf) SetUdldSettings(v FabricUdldGlobalSettings) {
	o.UdldSettings.Set(&v)
}

// SetUdldSettingsNil sets the value for UdldSettings to be an explicit nil
func (o *FabricSwitchControlPolicyAllOf) SetUdldSettingsNil() {
	o.UdldSettings.Set(nil)
}

// UnsetUdldSettings ensures that no value is present for UdldSettings, not even an explicit nil
func (o *FabricSwitchControlPolicyAllOf) UnsetUdldSettings() {
	o.UdldSettings.Unset()
}

// GetVlanPortOptimizationEnabled returns the VlanPortOptimizationEnabled field value if set, zero value otherwise.
func (o *FabricSwitchControlPolicyAllOf) GetVlanPortOptimizationEnabled() bool {
	if o == nil || o.VlanPortOptimizationEnabled == nil {
		var ret bool
		return ret
	}
	return *o.VlanPortOptimizationEnabled
}

// GetVlanPortOptimizationEnabledOk returns a tuple with the VlanPortOptimizationEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FabricSwitchControlPolicyAllOf) GetVlanPortOptimizationEnabledOk() (*bool, bool) {
	if o == nil || o.VlanPortOptimizationEnabled == nil {
		return nil, false
	}
	return o.VlanPortOptimizationEnabled, true
}

// HasVlanPortOptimizationEnabled returns a boolean if a field has been set.
func (o *FabricSwitchControlPolicyAllOf) HasVlanPortOptimizationEnabled() bool {
	if o != nil && o.VlanPortOptimizationEnabled != nil {
		return true
	}

	return false
}

// SetVlanPortOptimizationEnabled gets a reference to the given bool and assigns it to the VlanPortOptimizationEnabled field.
func (o *FabricSwitchControlPolicyAllOf) SetVlanPortOptimizationEnabled(v bool) {
	o.VlanPortOptimizationEnabled = &v
}

// GetOrganization returns the Organization field value if set, zero value otherwise.
func (o *FabricSwitchControlPolicyAllOf) GetOrganization() OrganizationOrganizationRelationship {
	if o == nil || o.Organization == nil {
		var ret OrganizationOrganizationRelationship
		return ret
	}
	return *o.Organization
}

// GetOrganizationOk returns a tuple with the Organization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FabricSwitchControlPolicyAllOf) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool) {
	if o == nil || o.Organization == nil {
		return nil, false
	}
	return o.Organization, true
}

// HasOrganization returns a boolean if a field has been set.
func (o *FabricSwitchControlPolicyAllOf) HasOrganization() bool {
	if o != nil && o.Organization != nil {
		return true
	}

	return false
}

// SetOrganization gets a reference to the given OrganizationOrganizationRelationship and assigns it to the Organization field.
func (o *FabricSwitchControlPolicyAllOf) SetOrganization(v OrganizationOrganizationRelationship) {
	o.Organization = &v
}

// GetProfiles returns the Profiles field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FabricSwitchControlPolicyAllOf) GetProfiles() []FabricSwitchProfileRelationship {
	if o == nil {
		var ret []FabricSwitchProfileRelationship
		return ret
	}
	return o.Profiles
}

// GetProfilesOk returns a tuple with the Profiles field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FabricSwitchControlPolicyAllOf) GetProfilesOk() (*[]FabricSwitchProfileRelationship, bool) {
	if o == nil || o.Profiles == nil {
		return nil, false
	}
	return &o.Profiles, true
}

// HasProfiles returns a boolean if a field has been set.
func (o *FabricSwitchControlPolicyAllOf) HasProfiles() bool {
	if o != nil && o.Profiles != nil {
		return true
	}

	return false
}

// SetProfiles gets a reference to the given []FabricSwitchProfileRelationship and assigns it to the Profiles field.
func (o *FabricSwitchControlPolicyAllOf) SetProfiles(v []FabricSwitchProfileRelationship) {
	o.Profiles = v
}

func (o FabricSwitchControlPolicyAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.MacAgingSettings.IsSet() {
		toSerialize["MacAgingSettings"] = o.MacAgingSettings.Get()
	}
	if o.UdldSettings.IsSet() {
		toSerialize["UdldSettings"] = o.UdldSettings.Get()
	}
	if o.VlanPortOptimizationEnabled != nil {
		toSerialize["VlanPortOptimizationEnabled"] = o.VlanPortOptimizationEnabled
	}
	if o.Organization != nil {
		toSerialize["Organization"] = o.Organization
	}
	if o.Profiles != nil {
		toSerialize["Profiles"] = o.Profiles
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *FabricSwitchControlPolicyAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varFabricSwitchControlPolicyAllOf := _FabricSwitchControlPolicyAllOf{}

	if err = json.Unmarshal(bytes, &varFabricSwitchControlPolicyAllOf); err == nil {
		*o = FabricSwitchControlPolicyAllOf(varFabricSwitchControlPolicyAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "MacAgingSettings")
		delete(additionalProperties, "UdldSettings")
		delete(additionalProperties, "VlanPortOptimizationEnabled")
		delete(additionalProperties, "Organization")
		delete(additionalProperties, "Profiles")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableFabricSwitchControlPolicyAllOf struct {
	value *FabricSwitchControlPolicyAllOf
	isSet bool
}

func (v NullableFabricSwitchControlPolicyAllOf) Get() *FabricSwitchControlPolicyAllOf {
	return v.value
}

func (v *NullableFabricSwitchControlPolicyAllOf) Set(val *FabricSwitchControlPolicyAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableFabricSwitchControlPolicyAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableFabricSwitchControlPolicyAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFabricSwitchControlPolicyAllOf(val *FabricSwitchControlPolicyAllOf) *NullableFabricSwitchControlPolicyAllOf {
	return &NullableFabricSwitchControlPolicyAllOf{value: val, isSet: true}
}

func (v NullableFabricSwitchControlPolicyAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFabricSwitchControlPolicyAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
