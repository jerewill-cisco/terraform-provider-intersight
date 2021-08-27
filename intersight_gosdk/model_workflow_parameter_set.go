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

// WorkflowParameterSet The parameter set is a rule with a condition to match the controlling parameter against a value and make a set of parameters applicable.
type WorkflowParameterSet struct {
	MoBaseComplexType
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// The condition to be evaluated. * `eq` - Checks if the values of the two parameters are equal. * `ne` - Checks if the values of the two parameters are not equal. * `contains` - Checks if the second parameter string value is a substring of the first parameter string value. * `matchesPattern` - Checks if a string matches a regular expression.
	Condition *string `json:"Condition,omitempty"`
	// Name of the controlling entity, whose value will be used for evaluating the parameter set.
	ControlParameter *string  `json:"ControlParameter,omitempty"`
	EnableParameters []string `json:"EnableParameters,omitempty"`
	// Name for the parameter set.  Limited to 64 alphanumeric characters (upper and lower case), and special characters '-' and '_'.
	Name *string `json:"Name,omitempty"`
	// The controlling parameter will be evaluated against this 'value'.
	Value                *string `json:"Value,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _WorkflowParameterSet WorkflowParameterSet

// NewWorkflowParameterSet instantiates a new WorkflowParameterSet object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkflowParameterSet(classId string, objectType string) *WorkflowParameterSet {
	this := WorkflowParameterSet{}
	this.ClassId = classId
	this.ObjectType = objectType
	var condition string = "eq"
	this.Condition = &condition
	return &this
}

// NewWorkflowParameterSetWithDefaults instantiates a new WorkflowParameterSet object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkflowParameterSetWithDefaults() *WorkflowParameterSet {
	this := WorkflowParameterSet{}
	var classId string = "workflow.ParameterSet"
	this.ClassId = classId
	var objectType string = "workflow.ParameterSet"
	this.ObjectType = objectType
	var condition string = "eq"
	this.Condition = &condition
	return &this
}

// GetClassId returns the ClassId field value
func (o *WorkflowParameterSet) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *WorkflowParameterSet) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *WorkflowParameterSet) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *WorkflowParameterSet) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *WorkflowParameterSet) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *WorkflowParameterSet) SetObjectType(v string) {
	o.ObjectType = v
}

// GetCondition returns the Condition field value if set, zero value otherwise.
func (o *WorkflowParameterSet) GetCondition() string {
	if o == nil || o.Condition == nil {
		var ret string
		return ret
	}
	return *o.Condition
}

// GetConditionOk returns a tuple with the Condition field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowParameterSet) GetConditionOk() (*string, bool) {
	if o == nil || o.Condition == nil {
		return nil, false
	}
	return o.Condition, true
}

// HasCondition returns a boolean if a field has been set.
func (o *WorkflowParameterSet) HasCondition() bool {
	if o != nil && o.Condition != nil {
		return true
	}

	return false
}

// SetCondition gets a reference to the given string and assigns it to the Condition field.
func (o *WorkflowParameterSet) SetCondition(v string) {
	o.Condition = &v
}

// GetControlParameter returns the ControlParameter field value if set, zero value otherwise.
func (o *WorkflowParameterSet) GetControlParameter() string {
	if o == nil || o.ControlParameter == nil {
		var ret string
		return ret
	}
	return *o.ControlParameter
}

// GetControlParameterOk returns a tuple with the ControlParameter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowParameterSet) GetControlParameterOk() (*string, bool) {
	if o == nil || o.ControlParameter == nil {
		return nil, false
	}
	return o.ControlParameter, true
}

// HasControlParameter returns a boolean if a field has been set.
func (o *WorkflowParameterSet) HasControlParameter() bool {
	if o != nil && o.ControlParameter != nil {
		return true
	}

	return false
}

// SetControlParameter gets a reference to the given string and assigns it to the ControlParameter field.
func (o *WorkflowParameterSet) SetControlParameter(v string) {
	o.ControlParameter = &v
}

// GetEnableParameters returns the EnableParameters field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WorkflowParameterSet) GetEnableParameters() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.EnableParameters
}

// GetEnableParametersOk returns a tuple with the EnableParameters field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WorkflowParameterSet) GetEnableParametersOk() (*[]string, bool) {
	if o == nil || o.EnableParameters == nil {
		return nil, false
	}
	return &o.EnableParameters, true
}

// HasEnableParameters returns a boolean if a field has been set.
func (o *WorkflowParameterSet) HasEnableParameters() bool {
	if o != nil && o.EnableParameters != nil {
		return true
	}

	return false
}

// SetEnableParameters gets a reference to the given []string and assigns it to the EnableParameters field.
func (o *WorkflowParameterSet) SetEnableParameters(v []string) {
	o.EnableParameters = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *WorkflowParameterSet) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowParameterSet) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *WorkflowParameterSet) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *WorkflowParameterSet) SetName(v string) {
	o.Name = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *WorkflowParameterSet) GetValue() string {
	if o == nil || o.Value == nil {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowParameterSet) GetValueOk() (*string, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *WorkflowParameterSet) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *WorkflowParameterSet) SetValue(v string) {
	o.Value = &v
}

func (o WorkflowParameterSet) MarshalJSON() ([]byte, error) {
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
	if o.Condition != nil {
		toSerialize["Condition"] = o.Condition
	}
	if o.ControlParameter != nil {
		toSerialize["ControlParameter"] = o.ControlParameter
	}
	if o.EnableParameters != nil {
		toSerialize["EnableParameters"] = o.EnableParameters
	}
	if o.Name != nil {
		toSerialize["Name"] = o.Name
	}
	if o.Value != nil {
		toSerialize["Value"] = o.Value
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *WorkflowParameterSet) UnmarshalJSON(bytes []byte) (err error) {
	type WorkflowParameterSetWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// The condition to be evaluated. * `eq` - Checks if the values of the two parameters are equal. * `ne` - Checks if the values of the two parameters are not equal. * `contains` - Checks if the second parameter string value is a substring of the first parameter string value. * `matchesPattern` - Checks if a string matches a regular expression.
		Condition *string `json:"Condition,omitempty"`
		// Name of the controlling entity, whose value will be used for evaluating the parameter set.
		ControlParameter *string  `json:"ControlParameter,omitempty"`
		EnableParameters []string `json:"EnableParameters,omitempty"`
		// Name for the parameter set.  Limited to 64 alphanumeric characters (upper and lower case), and special characters '-' and '_'.
		Name *string `json:"Name,omitempty"`
		// The controlling parameter will be evaluated against this 'value'.
		Value *string `json:"Value,omitempty"`
	}

	varWorkflowParameterSetWithoutEmbeddedStruct := WorkflowParameterSetWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varWorkflowParameterSetWithoutEmbeddedStruct)
	if err == nil {
		varWorkflowParameterSet := _WorkflowParameterSet{}
		varWorkflowParameterSet.ClassId = varWorkflowParameterSetWithoutEmbeddedStruct.ClassId
		varWorkflowParameterSet.ObjectType = varWorkflowParameterSetWithoutEmbeddedStruct.ObjectType
		varWorkflowParameterSet.Condition = varWorkflowParameterSetWithoutEmbeddedStruct.Condition
		varWorkflowParameterSet.ControlParameter = varWorkflowParameterSetWithoutEmbeddedStruct.ControlParameter
		varWorkflowParameterSet.EnableParameters = varWorkflowParameterSetWithoutEmbeddedStruct.EnableParameters
		varWorkflowParameterSet.Name = varWorkflowParameterSetWithoutEmbeddedStruct.Name
		varWorkflowParameterSet.Value = varWorkflowParameterSetWithoutEmbeddedStruct.Value
		*o = WorkflowParameterSet(varWorkflowParameterSet)
	} else {
		return err
	}

	varWorkflowParameterSet := _WorkflowParameterSet{}

	err = json.Unmarshal(bytes, &varWorkflowParameterSet)
	if err == nil {
		o.MoBaseComplexType = varWorkflowParameterSet.MoBaseComplexType
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Condition")
		delete(additionalProperties, "ControlParameter")
		delete(additionalProperties, "EnableParameters")
		delete(additionalProperties, "Name")
		delete(additionalProperties, "Value")

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

type NullableWorkflowParameterSet struct {
	value *WorkflowParameterSet
	isSet bool
}

func (v NullableWorkflowParameterSet) Get() *WorkflowParameterSet {
	return v.value
}

func (v *NullableWorkflowParameterSet) Set(val *WorkflowParameterSet) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkflowParameterSet) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkflowParameterSet) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkflowParameterSet(val *WorkflowParameterSet) *NullableWorkflowParameterSet {
	return &NullableWorkflowParameterSet{value: val, isSet: true}
}

func (v NullableWorkflowParameterSet) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkflowParameterSet) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
