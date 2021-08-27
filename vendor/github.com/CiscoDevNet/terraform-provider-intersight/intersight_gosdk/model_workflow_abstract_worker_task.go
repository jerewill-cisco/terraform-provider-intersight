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

// WorkflowAbstractWorkerTask An AbstractWorkerTask is used to model a task that does some end-user work. This can be another task or it can be another workflow.
type WorkflowAbstractWorkerTask struct {
	WorkflowWorkflowTask
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property. The enum values provides the list of concrete types that can be instantiated from this abstract type.
	ObjectType string `json:"ObjectType"`
	// JSON formatted map that defines the input given to the task. JSONPath is used for chaining output from previous tasks as inputs into the current task. The format to specify the mapping is '${Source.input/output.JsonPath}'. 'Source' can be either workflow or the name of the task within the workflow. You can map the task input to either a workflow input or a task output. Following this is JSON path expression to extract JSON fragment from source's input/output.
	InputParameters interface{} `json:"InputParameters,omitempty"`
	// This specifies the name of the next task to run if Task fails.  This is the unique name given to the task instance within the workflow. In a graph model, denotes an edge to another Task Node.
	OnFailure *string `json:"OnFailure,omitempty"`
	// This specifies the name of the next task to run if Task succeeds.  This is the unique name given to the task instance within the workflow. In a graph model, denotes an edge to another Task Node.
	OnSuccess *string `json:"OnSuccess,omitempty"`
	// The task is disabled/enabled for rollback operation in this workflow if the task has rollback support.
	RollbackDisabled *bool `json:"RollbackDisabled,omitempty"`
	// UseDefault when set to true, means the default version of the task or workflow will be used at the time of execution. When this property is set then version for task or subworkflow cannot be set. When workflow is created or updated the default version of task or subworkflow will be used for validation, but when the workflow is executed the default version that that time will be used for validation and subsequent execution.
	UseDefault           *bool `json:"UseDefault,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _WorkflowAbstractWorkerTask WorkflowAbstractWorkerTask

// NewWorkflowAbstractWorkerTask instantiates a new WorkflowAbstractWorkerTask object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkflowAbstractWorkerTask(classId string, objectType string) *WorkflowAbstractWorkerTask {
	this := WorkflowAbstractWorkerTask{}
	this.ClassId = classId
	this.ObjectType = objectType
	var rollbackDisabled bool = false
	this.RollbackDisabled = &rollbackDisabled
	var useDefault bool = false
	this.UseDefault = &useDefault
	return &this
}

// NewWorkflowAbstractWorkerTaskWithDefaults instantiates a new WorkflowAbstractWorkerTask object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkflowAbstractWorkerTaskWithDefaults() *WorkflowAbstractWorkerTask {
	this := WorkflowAbstractWorkerTask{}
	var rollbackDisabled bool = false
	this.RollbackDisabled = &rollbackDisabled
	var useDefault bool = false
	this.UseDefault = &useDefault
	return &this
}

// GetClassId returns the ClassId field value
func (o *WorkflowAbstractWorkerTask) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *WorkflowAbstractWorkerTask) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *WorkflowAbstractWorkerTask) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *WorkflowAbstractWorkerTask) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *WorkflowAbstractWorkerTask) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *WorkflowAbstractWorkerTask) SetObjectType(v string) {
	o.ObjectType = v
}

// GetInputParameters returns the InputParameters field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WorkflowAbstractWorkerTask) GetInputParameters() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.InputParameters
}

// GetInputParametersOk returns a tuple with the InputParameters field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WorkflowAbstractWorkerTask) GetInputParametersOk() (*interface{}, bool) {
	if o == nil || o.InputParameters == nil {
		return nil, false
	}
	return &o.InputParameters, true
}

// HasInputParameters returns a boolean if a field has been set.
func (o *WorkflowAbstractWorkerTask) HasInputParameters() bool {
	if o != nil && o.InputParameters != nil {
		return true
	}

	return false
}

// SetInputParameters gets a reference to the given interface{} and assigns it to the InputParameters field.
func (o *WorkflowAbstractWorkerTask) SetInputParameters(v interface{}) {
	o.InputParameters = v
}

// GetOnFailure returns the OnFailure field value if set, zero value otherwise.
func (o *WorkflowAbstractWorkerTask) GetOnFailure() string {
	if o == nil || o.OnFailure == nil {
		var ret string
		return ret
	}
	return *o.OnFailure
}

// GetOnFailureOk returns a tuple with the OnFailure field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowAbstractWorkerTask) GetOnFailureOk() (*string, bool) {
	if o == nil || o.OnFailure == nil {
		return nil, false
	}
	return o.OnFailure, true
}

// HasOnFailure returns a boolean if a field has been set.
func (o *WorkflowAbstractWorkerTask) HasOnFailure() bool {
	if o != nil && o.OnFailure != nil {
		return true
	}

	return false
}

// SetOnFailure gets a reference to the given string and assigns it to the OnFailure field.
func (o *WorkflowAbstractWorkerTask) SetOnFailure(v string) {
	o.OnFailure = &v
}

// GetOnSuccess returns the OnSuccess field value if set, zero value otherwise.
func (o *WorkflowAbstractWorkerTask) GetOnSuccess() string {
	if o == nil || o.OnSuccess == nil {
		var ret string
		return ret
	}
	return *o.OnSuccess
}

// GetOnSuccessOk returns a tuple with the OnSuccess field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowAbstractWorkerTask) GetOnSuccessOk() (*string, bool) {
	if o == nil || o.OnSuccess == nil {
		return nil, false
	}
	return o.OnSuccess, true
}

// HasOnSuccess returns a boolean if a field has been set.
func (o *WorkflowAbstractWorkerTask) HasOnSuccess() bool {
	if o != nil && o.OnSuccess != nil {
		return true
	}

	return false
}

// SetOnSuccess gets a reference to the given string and assigns it to the OnSuccess field.
func (o *WorkflowAbstractWorkerTask) SetOnSuccess(v string) {
	o.OnSuccess = &v
}

// GetRollbackDisabled returns the RollbackDisabled field value if set, zero value otherwise.
func (o *WorkflowAbstractWorkerTask) GetRollbackDisabled() bool {
	if o == nil || o.RollbackDisabled == nil {
		var ret bool
		return ret
	}
	return *o.RollbackDisabled
}

// GetRollbackDisabledOk returns a tuple with the RollbackDisabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowAbstractWorkerTask) GetRollbackDisabledOk() (*bool, bool) {
	if o == nil || o.RollbackDisabled == nil {
		return nil, false
	}
	return o.RollbackDisabled, true
}

// HasRollbackDisabled returns a boolean if a field has been set.
func (o *WorkflowAbstractWorkerTask) HasRollbackDisabled() bool {
	if o != nil && o.RollbackDisabled != nil {
		return true
	}

	return false
}

// SetRollbackDisabled gets a reference to the given bool and assigns it to the RollbackDisabled field.
func (o *WorkflowAbstractWorkerTask) SetRollbackDisabled(v bool) {
	o.RollbackDisabled = &v
}

// GetUseDefault returns the UseDefault field value if set, zero value otherwise.
func (o *WorkflowAbstractWorkerTask) GetUseDefault() bool {
	if o == nil || o.UseDefault == nil {
		var ret bool
		return ret
	}
	return *o.UseDefault
}

// GetUseDefaultOk returns a tuple with the UseDefault field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowAbstractWorkerTask) GetUseDefaultOk() (*bool, bool) {
	if o == nil || o.UseDefault == nil {
		return nil, false
	}
	return o.UseDefault, true
}

// HasUseDefault returns a boolean if a field has been set.
func (o *WorkflowAbstractWorkerTask) HasUseDefault() bool {
	if o != nil && o.UseDefault != nil {
		return true
	}

	return false
}

// SetUseDefault gets a reference to the given bool and assigns it to the UseDefault field.
func (o *WorkflowAbstractWorkerTask) SetUseDefault(v bool) {
	o.UseDefault = &v
}

func (o WorkflowAbstractWorkerTask) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedWorkflowWorkflowTask, errWorkflowWorkflowTask := json.Marshal(o.WorkflowWorkflowTask)
	if errWorkflowWorkflowTask != nil {
		return []byte{}, errWorkflowWorkflowTask
	}
	errWorkflowWorkflowTask = json.Unmarshal([]byte(serializedWorkflowWorkflowTask), &toSerialize)
	if errWorkflowWorkflowTask != nil {
		return []byte{}, errWorkflowWorkflowTask
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.InputParameters != nil {
		toSerialize["InputParameters"] = o.InputParameters
	}
	if o.OnFailure != nil {
		toSerialize["OnFailure"] = o.OnFailure
	}
	if o.OnSuccess != nil {
		toSerialize["OnSuccess"] = o.OnSuccess
	}
	if o.RollbackDisabled != nil {
		toSerialize["RollbackDisabled"] = o.RollbackDisabled
	}
	if o.UseDefault != nil {
		toSerialize["UseDefault"] = o.UseDefault
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *WorkflowAbstractWorkerTask) UnmarshalJSON(bytes []byte) (err error) {
	type WorkflowAbstractWorkerTaskWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property. The enum values provides the list of concrete types that can be instantiated from this abstract type.
		ObjectType string `json:"ObjectType"`
		// JSON formatted map that defines the input given to the task. JSONPath is used for chaining output from previous tasks as inputs into the current task. The format to specify the mapping is '${Source.input/output.JsonPath}'. 'Source' can be either workflow or the name of the task within the workflow. You can map the task input to either a workflow input or a task output. Following this is JSON path expression to extract JSON fragment from source's input/output.
		InputParameters interface{} `json:"InputParameters,omitempty"`
		// This specifies the name of the next task to run if Task fails.  This is the unique name given to the task instance within the workflow. In a graph model, denotes an edge to another Task Node.
		OnFailure *string `json:"OnFailure,omitempty"`
		// This specifies the name of the next task to run if Task succeeds.  This is the unique name given to the task instance within the workflow. In a graph model, denotes an edge to another Task Node.
		OnSuccess *string `json:"OnSuccess,omitempty"`
		// The task is disabled/enabled for rollback operation in this workflow if the task has rollback support.
		RollbackDisabled *bool `json:"RollbackDisabled,omitempty"`
		// UseDefault when set to true, means the default version of the task or workflow will be used at the time of execution. When this property is set then version for task or subworkflow cannot be set. When workflow is created or updated the default version of task or subworkflow will be used for validation, but when the workflow is executed the default version that that time will be used for validation and subsequent execution.
		UseDefault *bool `json:"UseDefault,omitempty"`
	}

	varWorkflowAbstractWorkerTaskWithoutEmbeddedStruct := WorkflowAbstractWorkerTaskWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varWorkflowAbstractWorkerTaskWithoutEmbeddedStruct)
	if err == nil {
		varWorkflowAbstractWorkerTask := _WorkflowAbstractWorkerTask{}
		varWorkflowAbstractWorkerTask.ClassId = varWorkflowAbstractWorkerTaskWithoutEmbeddedStruct.ClassId
		varWorkflowAbstractWorkerTask.ObjectType = varWorkflowAbstractWorkerTaskWithoutEmbeddedStruct.ObjectType
		varWorkflowAbstractWorkerTask.InputParameters = varWorkflowAbstractWorkerTaskWithoutEmbeddedStruct.InputParameters
		varWorkflowAbstractWorkerTask.OnFailure = varWorkflowAbstractWorkerTaskWithoutEmbeddedStruct.OnFailure
		varWorkflowAbstractWorkerTask.OnSuccess = varWorkflowAbstractWorkerTaskWithoutEmbeddedStruct.OnSuccess
		varWorkflowAbstractWorkerTask.RollbackDisabled = varWorkflowAbstractWorkerTaskWithoutEmbeddedStruct.RollbackDisabled
		varWorkflowAbstractWorkerTask.UseDefault = varWorkflowAbstractWorkerTaskWithoutEmbeddedStruct.UseDefault
		*o = WorkflowAbstractWorkerTask(varWorkflowAbstractWorkerTask)
	} else {
		return err
	}

	varWorkflowAbstractWorkerTask := _WorkflowAbstractWorkerTask{}

	err = json.Unmarshal(bytes, &varWorkflowAbstractWorkerTask)
	if err == nil {
		o.WorkflowWorkflowTask = varWorkflowAbstractWorkerTask.WorkflowWorkflowTask
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "InputParameters")
		delete(additionalProperties, "OnFailure")
		delete(additionalProperties, "OnSuccess")
		delete(additionalProperties, "RollbackDisabled")
		delete(additionalProperties, "UseDefault")

		// remove fields from embedded structs
		reflectWorkflowWorkflowTask := reflect.ValueOf(o.WorkflowWorkflowTask)
		for i := 0; i < reflectWorkflowWorkflowTask.Type().NumField(); i++ {
			t := reflectWorkflowWorkflowTask.Type().Field(i)

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

type NullableWorkflowAbstractWorkerTask struct {
	value *WorkflowAbstractWorkerTask
	isSet bool
}

func (v NullableWorkflowAbstractWorkerTask) Get() *WorkflowAbstractWorkerTask {
	return v.value
}

func (v *NullableWorkflowAbstractWorkerTask) Set(val *WorkflowAbstractWorkerTask) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkflowAbstractWorkerTask) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkflowAbstractWorkerTask) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkflowAbstractWorkerTask(val *WorkflowAbstractWorkerTask) *NullableWorkflowAbstractWorkerTask {
	return &NullableWorkflowAbstractWorkerTask{value: val, isSet: true}
}

func (v NullableWorkflowAbstractWorkerTask) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkflowAbstractWorkerTask) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
