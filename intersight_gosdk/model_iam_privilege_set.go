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
	"reflect"
	"strings"
)

// IamPrivilegeSet A collection of privileges.
type IamPrivilegeSet struct {
	MoBaseMo
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Description of the privilege set.
	Description *string `json:"Description,omitempty"`
	// Name of the privilege set.
	Name           *string                 `json:"Name,omitempty"`
	PrivilegeNames []string                `json:"PrivilegeNames,omitempty"`
	Account        *IamAccountRelationship `json:"Account,omitempty"`
	// An array of relationships to iamPrivilegeSet resources.
	AssociatedPrivilegeSets []IamPrivilegeSetRelationship `json:"AssociatedPrivilegeSets,omitempty"`
	// An array of relationships to iamPrivilege resources.
	Privileges           []IamPrivilegeRelationship `json:"Privileges,omitempty"`
	System               *IamSystemRelationship     `json:"System,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _IamPrivilegeSet IamPrivilegeSet

// NewIamPrivilegeSet instantiates a new IamPrivilegeSet object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIamPrivilegeSet(classId string, objectType string) *IamPrivilegeSet {
	this := IamPrivilegeSet{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewIamPrivilegeSetWithDefaults instantiates a new IamPrivilegeSet object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIamPrivilegeSetWithDefaults() *IamPrivilegeSet {
	this := IamPrivilegeSet{}
	var classId string = "iam.PrivilegeSet"
	this.ClassId = classId
	var objectType string = "iam.PrivilegeSet"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *IamPrivilegeSet) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *IamPrivilegeSet) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *IamPrivilegeSet) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *IamPrivilegeSet) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *IamPrivilegeSet) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *IamPrivilegeSet) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *IamPrivilegeSet) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamPrivilegeSet) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *IamPrivilegeSet) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *IamPrivilegeSet) SetDescription(v string) {
	o.Description = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *IamPrivilegeSet) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamPrivilegeSet) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *IamPrivilegeSet) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *IamPrivilegeSet) SetName(v string) {
	o.Name = &v
}

// GetPrivilegeNames returns the PrivilegeNames field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IamPrivilegeSet) GetPrivilegeNames() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.PrivilegeNames
}

// GetPrivilegeNamesOk returns a tuple with the PrivilegeNames field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IamPrivilegeSet) GetPrivilegeNamesOk() (*[]string, bool) {
	if o == nil || o.PrivilegeNames == nil {
		return nil, false
	}
	return &o.PrivilegeNames, true
}

// HasPrivilegeNames returns a boolean if a field has been set.
func (o *IamPrivilegeSet) HasPrivilegeNames() bool {
	if o != nil && o.PrivilegeNames != nil {
		return true
	}

	return false
}

// SetPrivilegeNames gets a reference to the given []string and assigns it to the PrivilegeNames field.
func (o *IamPrivilegeSet) SetPrivilegeNames(v []string) {
	o.PrivilegeNames = v
}

// GetAccount returns the Account field value if set, zero value otherwise.
func (o *IamPrivilegeSet) GetAccount() IamAccountRelationship {
	if o == nil || o.Account == nil {
		var ret IamAccountRelationship
		return ret
	}
	return *o.Account
}

// GetAccountOk returns a tuple with the Account field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamPrivilegeSet) GetAccountOk() (*IamAccountRelationship, bool) {
	if o == nil || o.Account == nil {
		return nil, false
	}
	return o.Account, true
}

// HasAccount returns a boolean if a field has been set.
func (o *IamPrivilegeSet) HasAccount() bool {
	if o != nil && o.Account != nil {
		return true
	}

	return false
}

// SetAccount gets a reference to the given IamAccountRelationship and assigns it to the Account field.
func (o *IamPrivilegeSet) SetAccount(v IamAccountRelationship) {
	o.Account = &v
}

// GetAssociatedPrivilegeSets returns the AssociatedPrivilegeSets field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IamPrivilegeSet) GetAssociatedPrivilegeSets() []IamPrivilegeSetRelationship {
	if o == nil {
		var ret []IamPrivilegeSetRelationship
		return ret
	}
	return o.AssociatedPrivilegeSets
}

// GetAssociatedPrivilegeSetsOk returns a tuple with the AssociatedPrivilegeSets field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IamPrivilegeSet) GetAssociatedPrivilegeSetsOk() (*[]IamPrivilegeSetRelationship, bool) {
	if o == nil || o.AssociatedPrivilegeSets == nil {
		return nil, false
	}
	return &o.AssociatedPrivilegeSets, true
}

// HasAssociatedPrivilegeSets returns a boolean if a field has been set.
func (o *IamPrivilegeSet) HasAssociatedPrivilegeSets() bool {
	if o != nil && o.AssociatedPrivilegeSets != nil {
		return true
	}

	return false
}

// SetAssociatedPrivilegeSets gets a reference to the given []IamPrivilegeSetRelationship and assigns it to the AssociatedPrivilegeSets field.
func (o *IamPrivilegeSet) SetAssociatedPrivilegeSets(v []IamPrivilegeSetRelationship) {
	o.AssociatedPrivilegeSets = v
}

// GetPrivileges returns the Privileges field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IamPrivilegeSet) GetPrivileges() []IamPrivilegeRelationship {
	if o == nil {
		var ret []IamPrivilegeRelationship
		return ret
	}
	return o.Privileges
}

// GetPrivilegesOk returns a tuple with the Privileges field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IamPrivilegeSet) GetPrivilegesOk() (*[]IamPrivilegeRelationship, bool) {
	if o == nil || o.Privileges == nil {
		return nil, false
	}
	return &o.Privileges, true
}

// HasPrivileges returns a boolean if a field has been set.
func (o *IamPrivilegeSet) HasPrivileges() bool {
	if o != nil && o.Privileges != nil {
		return true
	}

	return false
}

// SetPrivileges gets a reference to the given []IamPrivilegeRelationship and assigns it to the Privileges field.
func (o *IamPrivilegeSet) SetPrivileges(v []IamPrivilegeRelationship) {
	o.Privileges = v
}

// GetSystem returns the System field value if set, zero value otherwise.
func (o *IamPrivilegeSet) GetSystem() IamSystemRelationship {
	if o == nil || o.System == nil {
		var ret IamSystemRelationship
		return ret
	}
	return *o.System
}

// GetSystemOk returns a tuple with the System field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamPrivilegeSet) GetSystemOk() (*IamSystemRelationship, bool) {
	if o == nil || o.System == nil {
		return nil, false
	}
	return o.System, true
}

// HasSystem returns a boolean if a field has been set.
func (o *IamPrivilegeSet) HasSystem() bool {
	if o != nil && o.System != nil {
		return true
	}

	return false
}

// SetSystem gets a reference to the given IamSystemRelationship and assigns it to the System field.
func (o *IamPrivilegeSet) SetSystem(v IamSystemRelationship) {
	o.System = &v
}

func (o IamPrivilegeSet) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedMoBaseMo, errMoBaseMo := json.Marshal(o.MoBaseMo)
	if errMoBaseMo != nil {
		return []byte{}, errMoBaseMo
	}
	errMoBaseMo = json.Unmarshal([]byte(serializedMoBaseMo), &toSerialize)
	if errMoBaseMo != nil {
		return []byte{}, errMoBaseMo
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Description != nil {
		toSerialize["Description"] = o.Description
	}
	if o.Name != nil {
		toSerialize["Name"] = o.Name
	}
	if o.PrivilegeNames != nil {
		toSerialize["PrivilegeNames"] = o.PrivilegeNames
	}
	if o.Account != nil {
		toSerialize["Account"] = o.Account
	}
	if o.AssociatedPrivilegeSets != nil {
		toSerialize["AssociatedPrivilegeSets"] = o.AssociatedPrivilegeSets
	}
	if o.Privileges != nil {
		toSerialize["Privileges"] = o.Privileges
	}
	if o.System != nil {
		toSerialize["System"] = o.System
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *IamPrivilegeSet) UnmarshalJSON(bytes []byte) (err error) {
	type IamPrivilegeSetWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Description of the privilege set.
		Description *string `json:"Description,omitempty"`
		// Name of the privilege set.
		Name           *string                 `json:"Name,omitempty"`
		PrivilegeNames []string                `json:"PrivilegeNames,omitempty"`
		Account        *IamAccountRelationship `json:"Account,omitempty"`
		// An array of relationships to iamPrivilegeSet resources.
		AssociatedPrivilegeSets []IamPrivilegeSetRelationship `json:"AssociatedPrivilegeSets,omitempty"`
		// An array of relationships to iamPrivilege resources.
		Privileges []IamPrivilegeRelationship `json:"Privileges,omitempty"`
		System     *IamSystemRelationship     `json:"System,omitempty"`
	}

	varIamPrivilegeSetWithoutEmbeddedStruct := IamPrivilegeSetWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varIamPrivilegeSetWithoutEmbeddedStruct)
	if err == nil {
		varIamPrivilegeSet := _IamPrivilegeSet{}
		varIamPrivilegeSet.ClassId = varIamPrivilegeSetWithoutEmbeddedStruct.ClassId
		varIamPrivilegeSet.ObjectType = varIamPrivilegeSetWithoutEmbeddedStruct.ObjectType
		varIamPrivilegeSet.Description = varIamPrivilegeSetWithoutEmbeddedStruct.Description
		varIamPrivilegeSet.Name = varIamPrivilegeSetWithoutEmbeddedStruct.Name
		varIamPrivilegeSet.PrivilegeNames = varIamPrivilegeSetWithoutEmbeddedStruct.PrivilegeNames
		varIamPrivilegeSet.Account = varIamPrivilegeSetWithoutEmbeddedStruct.Account
		varIamPrivilegeSet.AssociatedPrivilegeSets = varIamPrivilegeSetWithoutEmbeddedStruct.AssociatedPrivilegeSets
		varIamPrivilegeSet.Privileges = varIamPrivilegeSetWithoutEmbeddedStruct.Privileges
		varIamPrivilegeSet.System = varIamPrivilegeSetWithoutEmbeddedStruct.System
		*o = IamPrivilegeSet(varIamPrivilegeSet)
	} else {
		return err
	}

	varIamPrivilegeSet := _IamPrivilegeSet{}

	err = json.Unmarshal(bytes, &varIamPrivilegeSet)
	if err == nil {
		o.MoBaseMo = varIamPrivilegeSet.MoBaseMo
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Description")
		delete(additionalProperties, "Name")
		delete(additionalProperties, "PrivilegeNames")
		delete(additionalProperties, "Account")
		delete(additionalProperties, "AssociatedPrivilegeSets")
		delete(additionalProperties, "Privileges")
		delete(additionalProperties, "System")

		// remove fields from embedded structs
		reflectMoBaseMo := reflect.ValueOf(o.MoBaseMo)
		for i := 0; i < reflectMoBaseMo.Type().NumField(); i++ {
			t := reflectMoBaseMo.Type().Field(i)

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

type NullableIamPrivilegeSet struct {
	value *IamPrivilegeSet
	isSet bool
}

func (v NullableIamPrivilegeSet) Get() *IamPrivilegeSet {
	return v.value
}

func (v *NullableIamPrivilegeSet) Set(val *IamPrivilegeSet) {
	v.value = val
	v.isSet = true
}

func (v NullableIamPrivilegeSet) IsSet() bool {
	return v.isSet
}

func (v *NullableIamPrivilegeSet) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIamPrivilegeSet(val *IamPrivilegeSet) *NullableIamPrivilegeSet {
	return &NullableIamPrivilegeSet{value: val, isSet: true}
}

func (v NullableIamPrivilegeSet) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIamPrivilegeSet) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
