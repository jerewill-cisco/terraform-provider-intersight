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
	"reflect"
	"strings"
)

// KubernetesAddonPolicy A policy that defines which AddonDefinitions to use.
type KubernetesAddonPolicy struct {
	PolicyAbstractPolicy
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType           string                                 `json:"ObjectType"`
	AddonConfiguration   NullableKubernetesAddonConfiguration   `json:"AddonConfiguration,omitempty"`
	AddonDefinition      *KubernetesAddonDefinitionRelationship `json:"AddonDefinition,omitempty"`
	Organization         *OrganizationOrganizationRelationship  `json:"Organization,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _KubernetesAddonPolicy KubernetesAddonPolicy

// NewKubernetesAddonPolicy instantiates a new KubernetesAddonPolicy object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKubernetesAddonPolicy(classId string, objectType string) *KubernetesAddonPolicy {
	this := KubernetesAddonPolicy{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewKubernetesAddonPolicyWithDefaults instantiates a new KubernetesAddonPolicy object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKubernetesAddonPolicyWithDefaults() *KubernetesAddonPolicy {
	this := KubernetesAddonPolicy{}
	var classId string = "kubernetes.AddonPolicy"
	this.ClassId = classId
	var objectType string = "kubernetes.AddonPolicy"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *KubernetesAddonPolicy) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *KubernetesAddonPolicy) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *KubernetesAddonPolicy) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *KubernetesAddonPolicy) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *KubernetesAddonPolicy) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *KubernetesAddonPolicy) SetObjectType(v string) {
	o.ObjectType = v
}

// GetAddonConfiguration returns the AddonConfiguration field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *KubernetesAddonPolicy) GetAddonConfiguration() KubernetesAddonConfiguration {
	if o == nil || o.AddonConfiguration.Get() == nil {
		var ret KubernetesAddonConfiguration
		return ret
	}
	return *o.AddonConfiguration.Get()
}

// GetAddonConfigurationOk returns a tuple with the AddonConfiguration field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KubernetesAddonPolicy) GetAddonConfigurationOk() (*KubernetesAddonConfiguration, bool) {
	if o == nil {
		return nil, false
	}
	return o.AddonConfiguration.Get(), o.AddonConfiguration.IsSet()
}

// HasAddonConfiguration returns a boolean if a field has been set.
func (o *KubernetesAddonPolicy) HasAddonConfiguration() bool {
	if o != nil && o.AddonConfiguration.IsSet() {
		return true
	}

	return false
}

// SetAddonConfiguration gets a reference to the given NullableKubernetesAddonConfiguration and assigns it to the AddonConfiguration field.
func (o *KubernetesAddonPolicy) SetAddonConfiguration(v KubernetesAddonConfiguration) {
	o.AddonConfiguration.Set(&v)
}

// SetAddonConfigurationNil sets the value for AddonConfiguration to be an explicit nil
func (o *KubernetesAddonPolicy) SetAddonConfigurationNil() {
	o.AddonConfiguration.Set(nil)
}

// UnsetAddonConfiguration ensures that no value is present for AddonConfiguration, not even an explicit nil
func (o *KubernetesAddonPolicy) UnsetAddonConfiguration() {
	o.AddonConfiguration.Unset()
}

// GetAddonDefinition returns the AddonDefinition field value if set, zero value otherwise.
func (o *KubernetesAddonPolicy) GetAddonDefinition() KubernetesAddonDefinitionRelationship {
	if o == nil || o.AddonDefinition == nil {
		var ret KubernetesAddonDefinitionRelationship
		return ret
	}
	return *o.AddonDefinition
}

// GetAddonDefinitionOk returns a tuple with the AddonDefinition field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesAddonPolicy) GetAddonDefinitionOk() (*KubernetesAddonDefinitionRelationship, bool) {
	if o == nil || o.AddonDefinition == nil {
		return nil, false
	}
	return o.AddonDefinition, true
}

// HasAddonDefinition returns a boolean if a field has been set.
func (o *KubernetesAddonPolicy) HasAddonDefinition() bool {
	if o != nil && o.AddonDefinition != nil {
		return true
	}

	return false
}

// SetAddonDefinition gets a reference to the given KubernetesAddonDefinitionRelationship and assigns it to the AddonDefinition field.
func (o *KubernetesAddonPolicy) SetAddonDefinition(v KubernetesAddonDefinitionRelationship) {
	o.AddonDefinition = &v
}

// GetOrganization returns the Organization field value if set, zero value otherwise.
func (o *KubernetesAddonPolicy) GetOrganization() OrganizationOrganizationRelationship {
	if o == nil || o.Organization == nil {
		var ret OrganizationOrganizationRelationship
		return ret
	}
	return *o.Organization
}

// GetOrganizationOk returns a tuple with the Organization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesAddonPolicy) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool) {
	if o == nil || o.Organization == nil {
		return nil, false
	}
	return o.Organization, true
}

// HasOrganization returns a boolean if a field has been set.
func (o *KubernetesAddonPolicy) HasOrganization() bool {
	if o != nil && o.Organization != nil {
		return true
	}

	return false
}

// SetOrganization gets a reference to the given OrganizationOrganizationRelationship and assigns it to the Organization field.
func (o *KubernetesAddonPolicy) SetOrganization(v OrganizationOrganizationRelationship) {
	o.Organization = &v
}

func (o KubernetesAddonPolicy) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedPolicyAbstractPolicy, errPolicyAbstractPolicy := json.Marshal(o.PolicyAbstractPolicy)
	if errPolicyAbstractPolicy != nil {
		return []byte{}, errPolicyAbstractPolicy
	}
	errPolicyAbstractPolicy = json.Unmarshal([]byte(serializedPolicyAbstractPolicy), &toSerialize)
	if errPolicyAbstractPolicy != nil {
		return []byte{}, errPolicyAbstractPolicy
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.AddonConfiguration.IsSet() {
		toSerialize["AddonConfiguration"] = o.AddonConfiguration.Get()
	}
	if o.AddonDefinition != nil {
		toSerialize["AddonDefinition"] = o.AddonDefinition
	}
	if o.Organization != nil {
		toSerialize["Organization"] = o.Organization
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *KubernetesAddonPolicy) UnmarshalJSON(bytes []byte) (err error) {
	type KubernetesAddonPolicyWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType         string                                 `json:"ObjectType"`
		AddonConfiguration NullableKubernetesAddonConfiguration   `json:"AddonConfiguration,omitempty"`
		AddonDefinition    *KubernetesAddonDefinitionRelationship `json:"AddonDefinition,omitempty"`
		Organization       *OrganizationOrganizationRelationship  `json:"Organization,omitempty"`
	}

	varKubernetesAddonPolicyWithoutEmbeddedStruct := KubernetesAddonPolicyWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varKubernetesAddonPolicyWithoutEmbeddedStruct)
	if err == nil {
		varKubernetesAddonPolicy := _KubernetesAddonPolicy{}
		varKubernetesAddonPolicy.ClassId = varKubernetesAddonPolicyWithoutEmbeddedStruct.ClassId
		varKubernetesAddonPolicy.ObjectType = varKubernetesAddonPolicyWithoutEmbeddedStruct.ObjectType
		varKubernetesAddonPolicy.AddonConfiguration = varKubernetesAddonPolicyWithoutEmbeddedStruct.AddonConfiguration
		varKubernetesAddonPolicy.AddonDefinition = varKubernetesAddonPolicyWithoutEmbeddedStruct.AddonDefinition
		varKubernetesAddonPolicy.Organization = varKubernetesAddonPolicyWithoutEmbeddedStruct.Organization
		*o = KubernetesAddonPolicy(varKubernetesAddonPolicy)
	} else {
		return err
	}

	varKubernetesAddonPolicy := _KubernetesAddonPolicy{}

	err = json.Unmarshal(bytes, &varKubernetesAddonPolicy)
	if err == nil {
		o.PolicyAbstractPolicy = varKubernetesAddonPolicy.PolicyAbstractPolicy
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "AddonConfiguration")
		delete(additionalProperties, "AddonDefinition")
		delete(additionalProperties, "Organization")

		// remove fields from embedded structs
		reflectPolicyAbstractPolicy := reflect.ValueOf(o.PolicyAbstractPolicy)
		for i := 0; i < reflectPolicyAbstractPolicy.Type().NumField(); i++ {
			t := reflectPolicyAbstractPolicy.Type().Field(i)

			if jsonTag := t.Tag.Get("json"); jsonTag != "" {
				fieldName := ""
				if commaIdx := strings.Index(jsonTag, ","); commaIdx > 0 {
					fieldName = jsonTag[:commaIdx]
				} else {
					fieldName = jsonTag
				}
				if fieldName != "AdditionalProperties" {
					delete(additionalProperties, fieldName)
				}
			}
		}

		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableKubernetesAddonPolicy struct {
	value *KubernetesAddonPolicy
	isSet bool
}

func (v NullableKubernetesAddonPolicy) Get() *KubernetesAddonPolicy {
	return v.value
}

func (v *NullableKubernetesAddonPolicy) Set(val *KubernetesAddonPolicy) {
	v.value = val
	v.isSet = true
}

func (v NullableKubernetesAddonPolicy) IsSet() bool {
	return v.isSet
}

func (v *NullableKubernetesAddonPolicy) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKubernetesAddonPolicy(val *KubernetesAddonPolicy) *NullableKubernetesAddonPolicy {
	return &NullableKubernetesAddonPolicy{value: val, isSet: true}
}

func (v NullableKubernetesAddonPolicy) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKubernetesAddonPolicy) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
