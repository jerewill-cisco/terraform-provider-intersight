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

// WorkflowCustomDataProperty Capture all the properties for Custom data type.
type WorkflowCustomDataProperty struct {
	MoBaseComplexType
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Specify the catalog moid that this custom data type belongs.
	CatalogMoid *string `json:"CatalogMoid,omitempty"`
	// The resolved custom data type definition managed object.
	CustomDataTypeId *string `json:"CustomDataTypeId,omitempty"`
	// Name of the custom data type for this input.
	CustomDataTypeName   *string `json:"CustomDataTypeName,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _WorkflowCustomDataProperty WorkflowCustomDataProperty

// NewWorkflowCustomDataProperty instantiates a new WorkflowCustomDataProperty object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkflowCustomDataProperty(classId string, objectType string) *WorkflowCustomDataProperty {
	this := WorkflowCustomDataProperty{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewWorkflowCustomDataPropertyWithDefaults instantiates a new WorkflowCustomDataProperty object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkflowCustomDataPropertyWithDefaults() *WorkflowCustomDataProperty {
	this := WorkflowCustomDataProperty{}
	var classId string = "workflow.CustomDataProperty"
	this.ClassId = classId
	var objectType string = "workflow.CustomDataProperty"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *WorkflowCustomDataProperty) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *WorkflowCustomDataProperty) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *WorkflowCustomDataProperty) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *WorkflowCustomDataProperty) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *WorkflowCustomDataProperty) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *WorkflowCustomDataProperty) SetObjectType(v string) {
	o.ObjectType = v
}

// GetCatalogMoid returns the CatalogMoid field value if set, zero value otherwise.
func (o *WorkflowCustomDataProperty) GetCatalogMoid() string {
	if o == nil || o.CatalogMoid == nil {
		var ret string
		return ret
	}
	return *o.CatalogMoid
}

// GetCatalogMoidOk returns a tuple with the CatalogMoid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowCustomDataProperty) GetCatalogMoidOk() (*string, bool) {
	if o == nil || o.CatalogMoid == nil {
		return nil, false
	}
	return o.CatalogMoid, true
}

// HasCatalogMoid returns a boolean if a field has been set.
func (o *WorkflowCustomDataProperty) HasCatalogMoid() bool {
	if o != nil && o.CatalogMoid != nil {
		return true
	}

	return false
}

// SetCatalogMoid gets a reference to the given string and assigns it to the CatalogMoid field.
func (o *WorkflowCustomDataProperty) SetCatalogMoid(v string) {
	o.CatalogMoid = &v
}

// GetCustomDataTypeId returns the CustomDataTypeId field value if set, zero value otherwise.
func (o *WorkflowCustomDataProperty) GetCustomDataTypeId() string {
	if o == nil || o.CustomDataTypeId == nil {
		var ret string
		return ret
	}
	return *o.CustomDataTypeId
}

// GetCustomDataTypeIdOk returns a tuple with the CustomDataTypeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowCustomDataProperty) GetCustomDataTypeIdOk() (*string, bool) {
	if o == nil || o.CustomDataTypeId == nil {
		return nil, false
	}
	return o.CustomDataTypeId, true
}

// HasCustomDataTypeId returns a boolean if a field has been set.
func (o *WorkflowCustomDataProperty) HasCustomDataTypeId() bool {
	if o != nil && o.CustomDataTypeId != nil {
		return true
	}

	return false
}

// SetCustomDataTypeId gets a reference to the given string and assigns it to the CustomDataTypeId field.
func (o *WorkflowCustomDataProperty) SetCustomDataTypeId(v string) {
	o.CustomDataTypeId = &v
}

// GetCustomDataTypeName returns the CustomDataTypeName field value if set, zero value otherwise.
func (o *WorkflowCustomDataProperty) GetCustomDataTypeName() string {
	if o == nil || o.CustomDataTypeName == nil {
		var ret string
		return ret
	}
	return *o.CustomDataTypeName
}

// GetCustomDataTypeNameOk returns a tuple with the CustomDataTypeName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowCustomDataProperty) GetCustomDataTypeNameOk() (*string, bool) {
	if o == nil || o.CustomDataTypeName == nil {
		return nil, false
	}
	return o.CustomDataTypeName, true
}

// HasCustomDataTypeName returns a boolean if a field has been set.
func (o *WorkflowCustomDataProperty) HasCustomDataTypeName() bool {
	if o != nil && o.CustomDataTypeName != nil {
		return true
	}

	return false
}

// SetCustomDataTypeName gets a reference to the given string and assigns it to the CustomDataTypeName field.
func (o *WorkflowCustomDataProperty) SetCustomDataTypeName(v string) {
	o.CustomDataTypeName = &v
}

func (o WorkflowCustomDataProperty) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedMoBaseComplexType, errMoBaseComplexType := json.Marshal(o.MoBaseComplexType)
	if errMoBaseComplexType != nil {
		return []byte{}, errMoBaseComplexType
	}
	errMoBaseComplexType = json.Unmarshal([]byte(serializedMoBaseComplexType), &toSerialize)
	if errMoBaseComplexType != nil {
		return []byte{}, errMoBaseComplexType
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.CatalogMoid != nil {
		toSerialize["CatalogMoid"] = o.CatalogMoid
	}
	if o.CustomDataTypeId != nil {
		toSerialize["CustomDataTypeId"] = o.CustomDataTypeId
	}
	if o.CustomDataTypeName != nil {
		toSerialize["CustomDataTypeName"] = o.CustomDataTypeName
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *WorkflowCustomDataProperty) UnmarshalJSON(bytes []byte) (err error) {
	type WorkflowCustomDataPropertyWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Specify the catalog moid that this custom data type belongs.
		CatalogMoid *string `json:"CatalogMoid,omitempty"`
		// The resolved custom data type definition managed object.
		CustomDataTypeId *string `json:"CustomDataTypeId,omitempty"`
		// Name of the custom data type for this input.
		CustomDataTypeName *string `json:"CustomDataTypeName,omitempty"`
	}

	varWorkflowCustomDataPropertyWithoutEmbeddedStruct := WorkflowCustomDataPropertyWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varWorkflowCustomDataPropertyWithoutEmbeddedStruct)
	if err == nil {
		varWorkflowCustomDataProperty := _WorkflowCustomDataProperty{}
		varWorkflowCustomDataProperty.ClassId = varWorkflowCustomDataPropertyWithoutEmbeddedStruct.ClassId
		varWorkflowCustomDataProperty.ObjectType = varWorkflowCustomDataPropertyWithoutEmbeddedStruct.ObjectType
		varWorkflowCustomDataProperty.CatalogMoid = varWorkflowCustomDataPropertyWithoutEmbeddedStruct.CatalogMoid
		varWorkflowCustomDataProperty.CustomDataTypeId = varWorkflowCustomDataPropertyWithoutEmbeddedStruct.CustomDataTypeId
		varWorkflowCustomDataProperty.CustomDataTypeName = varWorkflowCustomDataPropertyWithoutEmbeddedStruct.CustomDataTypeName
		*o = WorkflowCustomDataProperty(varWorkflowCustomDataProperty)
	} else {
		return err
	}

	varWorkflowCustomDataProperty := _WorkflowCustomDataProperty{}

	err = json.Unmarshal(bytes, &varWorkflowCustomDataProperty)
	if err == nil {
		o.MoBaseComplexType = varWorkflowCustomDataProperty.MoBaseComplexType
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "CatalogMoid")
		delete(additionalProperties, "CustomDataTypeId")
		delete(additionalProperties, "CustomDataTypeName")

		// remove fields from embedded structs
		reflectMoBaseComplexType := reflect.ValueOf(o.MoBaseComplexType)
		for i := 0; i < reflectMoBaseComplexType.Type().NumField(); i++ {
			t := reflectMoBaseComplexType.Type().Field(i)

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

type NullableWorkflowCustomDataProperty struct {
	value *WorkflowCustomDataProperty
	isSet bool
}

func (v NullableWorkflowCustomDataProperty) Get() *WorkflowCustomDataProperty {
	return v.value
}

func (v *NullableWorkflowCustomDataProperty) Set(val *WorkflowCustomDataProperty) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkflowCustomDataProperty) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkflowCustomDataProperty) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkflowCustomDataProperty(val *WorkflowCustomDataProperty) *NullableWorkflowCustomDataProperty {
	return &NullableWorkflowCustomDataProperty{value: val, isSet: true}
}

func (v NullableWorkflowCustomDataProperty) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkflowCustomDataProperty) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
