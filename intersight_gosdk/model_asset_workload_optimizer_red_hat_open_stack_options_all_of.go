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

// AssetWorkloadOptimizerRedHatOpenStackOptionsAllOf Definition of the list of properties defined in 'asset.WorkloadOptimizerRedHatOpenStackOptions', excluding properties defined in parent classes.
type AssetWorkloadOptimizerRedHatOpenStackOptionsAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// RedHat OpenStack Identity Service (keystone) domain name. Domain is an additional namespaces you can create in keystone to partition users, groups, and projects. Default domain name value is \"Default\".
	Domain *string `json:"Domain,omitempty"`
	// The name of tenant which has assigned administrator role this RedHat OpenStack target user is in. A tenant or project is referred to as a group of users of RedHat OpenStack. Each tenant can be assigned a role which gives all its member users their rights and privileges to perform a specific set of operations.
	Tenant               *string `json:"Tenant,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AssetWorkloadOptimizerRedHatOpenStackOptionsAllOf AssetWorkloadOptimizerRedHatOpenStackOptionsAllOf

// NewAssetWorkloadOptimizerRedHatOpenStackOptionsAllOf instantiates a new AssetWorkloadOptimizerRedHatOpenStackOptionsAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAssetWorkloadOptimizerRedHatOpenStackOptionsAllOf(classId string, objectType string) *AssetWorkloadOptimizerRedHatOpenStackOptionsAllOf {
	this := AssetWorkloadOptimizerRedHatOpenStackOptionsAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewAssetWorkloadOptimizerRedHatOpenStackOptionsAllOfWithDefaults instantiates a new AssetWorkloadOptimizerRedHatOpenStackOptionsAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAssetWorkloadOptimizerRedHatOpenStackOptionsAllOfWithDefaults() *AssetWorkloadOptimizerRedHatOpenStackOptionsAllOf {
	this := AssetWorkloadOptimizerRedHatOpenStackOptionsAllOf{}
	var classId string = "asset.WorkloadOptimizerRedHatOpenStackOptions"
	this.ClassId = classId
	var objectType string = "asset.WorkloadOptimizerRedHatOpenStackOptions"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *AssetWorkloadOptimizerRedHatOpenStackOptionsAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *AssetWorkloadOptimizerRedHatOpenStackOptionsAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *AssetWorkloadOptimizerRedHatOpenStackOptionsAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *AssetWorkloadOptimizerRedHatOpenStackOptionsAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *AssetWorkloadOptimizerRedHatOpenStackOptionsAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *AssetWorkloadOptimizerRedHatOpenStackOptionsAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDomain returns the Domain field value if set, zero value otherwise.
func (o *AssetWorkloadOptimizerRedHatOpenStackOptionsAllOf) GetDomain() string {
	if o == nil || o.Domain == nil {
		var ret string
		return ret
	}
	return *o.Domain
}

// GetDomainOk returns a tuple with the Domain field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetWorkloadOptimizerRedHatOpenStackOptionsAllOf) GetDomainOk() (*string, bool) {
	if o == nil || o.Domain == nil {
		return nil, false
	}
	return o.Domain, true
}

// HasDomain returns a boolean if a field has been set.
func (o *AssetWorkloadOptimizerRedHatOpenStackOptionsAllOf) HasDomain() bool {
	if o != nil && o.Domain != nil {
		return true
	}

	return false
}

// SetDomain gets a reference to the given string and assigns it to the Domain field.
func (o *AssetWorkloadOptimizerRedHatOpenStackOptionsAllOf) SetDomain(v string) {
	o.Domain = &v
}

// GetTenant returns the Tenant field value if set, zero value otherwise.
func (o *AssetWorkloadOptimizerRedHatOpenStackOptionsAllOf) GetTenant() string {
	if o == nil || o.Tenant == nil {
		var ret string
		return ret
	}
	return *o.Tenant
}

// GetTenantOk returns a tuple with the Tenant field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetWorkloadOptimizerRedHatOpenStackOptionsAllOf) GetTenantOk() (*string, bool) {
	if o == nil || o.Tenant == nil {
		return nil, false
	}
	return o.Tenant, true
}

// HasTenant returns a boolean if a field has been set.
func (o *AssetWorkloadOptimizerRedHatOpenStackOptionsAllOf) HasTenant() bool {
	if o != nil && o.Tenant != nil {
		return true
	}

	return false
}

// SetTenant gets a reference to the given string and assigns it to the Tenant field.
func (o *AssetWorkloadOptimizerRedHatOpenStackOptionsAllOf) SetTenant(v string) {
	o.Tenant = &v
}

func (o AssetWorkloadOptimizerRedHatOpenStackOptionsAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Domain != nil {
		toSerialize["Domain"] = o.Domain
	}
	if o.Tenant != nil {
		toSerialize["Tenant"] = o.Tenant
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *AssetWorkloadOptimizerRedHatOpenStackOptionsAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varAssetWorkloadOptimizerRedHatOpenStackOptionsAllOf := _AssetWorkloadOptimizerRedHatOpenStackOptionsAllOf{}

	if err = json.Unmarshal(bytes, &varAssetWorkloadOptimizerRedHatOpenStackOptionsAllOf); err == nil {
		*o = AssetWorkloadOptimizerRedHatOpenStackOptionsAllOf(varAssetWorkloadOptimizerRedHatOpenStackOptionsAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Domain")
		delete(additionalProperties, "Tenant")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAssetWorkloadOptimizerRedHatOpenStackOptionsAllOf struct {
	value *AssetWorkloadOptimizerRedHatOpenStackOptionsAllOf
	isSet bool
}

func (v NullableAssetWorkloadOptimizerRedHatOpenStackOptionsAllOf) Get() *AssetWorkloadOptimizerRedHatOpenStackOptionsAllOf {
	return v.value
}

func (v *NullableAssetWorkloadOptimizerRedHatOpenStackOptionsAllOf) Set(val *AssetWorkloadOptimizerRedHatOpenStackOptionsAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableAssetWorkloadOptimizerRedHatOpenStackOptionsAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableAssetWorkloadOptimizerRedHatOpenStackOptionsAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAssetWorkloadOptimizerRedHatOpenStackOptionsAllOf(val *AssetWorkloadOptimizerRedHatOpenStackOptionsAllOf) *NullableAssetWorkloadOptimizerRedHatOpenStackOptionsAllOf {
	return &NullableAssetWorkloadOptimizerRedHatOpenStackOptionsAllOf{value: val, isSet: true}
}

func (v NullableAssetWorkloadOptimizerRedHatOpenStackOptionsAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAssetWorkloadOptimizerRedHatOpenStackOptionsAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
