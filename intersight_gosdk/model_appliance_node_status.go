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
	"reflect"
	"strings"
)

// ApplianceNodeStatus Status of an Intersight Appliance node.
type ApplianceNodeStatus struct {
	MoBaseMo
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Percentage of CPU currently in use.
	CpuUsage *float32 `json:"CpuUsage,omitempty"`
	// Percentage of memory currently in use.
	MemUsage *float32 `json:"MemUsage,omitempty"`
	// System assigned unique ID of the Intersight Appliance node. The system incrementally assigns identifiers to each node in the Intersight Appliance cluster starting with a value of 1.
	NodeId *int64 `json:"NodeId,omitempty"`
	// State of the node in terms of its readiness to host Kubernetes pods. * `Down` - The node is yet to come up and join as a member of theKubernetes cluster. * `Preparing` - The node has come up and joined the Kubernetes cluster,preparing to host Kubernetes pods. * `Ready` - The node is ready to host Kubernetes pods.
	NodeState *string `json:"NodeState,omitempty"`
	// Operational status of the Intersight Appliance node. Operational status is based on the result of the status checks. If result of any check is Critical, then its value is Impaired. Otherwise, if result of any check is Warning, then its value is AttentionNeeded. If all checks are OK, then its value is Operational. * `Unknown` - Operational status of the Intersight Appliance entity is Unknown. * `Operational` - Operational status of the Intersight Appliance entity is Operational. * `Impaired` - Operational status of the Intersight Appliance entity is Impaired. * `AttentionNeeded` - Operational status of the Intersight Appliance entity is AttentionNeeded.
	OperationalStatus *string                `json:"OperationalStatus,omitempty"`
	StatusChecks      []ApplianceStatusCheck `json:"StatusChecks,omitempty"`
	// An array of relationships to applianceFileSystemStatus resources.
	FileSystemStatuses   []ApplianceFileSystemStatusRelationship `json:"FileSystemStatuses,omitempty"`
	NodeInfo             *ApplianceNodeInfoRelationship          `json:"NodeInfo,omitempty"`
	SystemStatus         *ApplianceSystemStatusRelationship      `json:"SystemStatus,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ApplianceNodeStatus ApplianceNodeStatus

// NewApplianceNodeStatus instantiates a new ApplianceNodeStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApplianceNodeStatus(classId string, objectType string) *ApplianceNodeStatus {
	this := ApplianceNodeStatus{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewApplianceNodeStatusWithDefaults instantiates a new ApplianceNodeStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApplianceNodeStatusWithDefaults() *ApplianceNodeStatus {
	this := ApplianceNodeStatus{}
	var classId string = "appliance.NodeStatus"
	this.ClassId = classId
	var objectType string = "appliance.NodeStatus"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *ApplianceNodeStatus) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *ApplianceNodeStatus) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *ApplianceNodeStatus) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *ApplianceNodeStatus) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *ApplianceNodeStatus) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *ApplianceNodeStatus) SetObjectType(v string) {
	o.ObjectType = v
}

// GetCpuUsage returns the CpuUsage field value if set, zero value otherwise.
func (o *ApplianceNodeStatus) GetCpuUsage() float32 {
	if o == nil || o.CpuUsage == nil {
		var ret float32
		return ret
	}
	return *o.CpuUsage
}

// GetCpuUsageOk returns a tuple with the CpuUsage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceNodeStatus) GetCpuUsageOk() (*float32, bool) {
	if o == nil || o.CpuUsage == nil {
		return nil, false
	}
	return o.CpuUsage, true
}

// HasCpuUsage returns a boolean if a field has been set.
func (o *ApplianceNodeStatus) HasCpuUsage() bool {
	if o != nil && o.CpuUsage != nil {
		return true
	}

	return false
}

// SetCpuUsage gets a reference to the given float32 and assigns it to the CpuUsage field.
func (o *ApplianceNodeStatus) SetCpuUsage(v float32) {
	o.CpuUsage = &v
}

// GetMemUsage returns the MemUsage field value if set, zero value otherwise.
func (o *ApplianceNodeStatus) GetMemUsage() float32 {
	if o == nil || o.MemUsage == nil {
		var ret float32
		return ret
	}
	return *o.MemUsage
}

// GetMemUsageOk returns a tuple with the MemUsage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceNodeStatus) GetMemUsageOk() (*float32, bool) {
	if o == nil || o.MemUsage == nil {
		return nil, false
	}
	return o.MemUsage, true
}

// HasMemUsage returns a boolean if a field has been set.
func (o *ApplianceNodeStatus) HasMemUsage() bool {
	if o != nil && o.MemUsage != nil {
		return true
	}

	return false
}

// SetMemUsage gets a reference to the given float32 and assigns it to the MemUsage field.
func (o *ApplianceNodeStatus) SetMemUsage(v float32) {
	o.MemUsage = &v
}

// GetNodeId returns the NodeId field value if set, zero value otherwise.
func (o *ApplianceNodeStatus) GetNodeId() int64 {
	if o == nil || o.NodeId == nil {
		var ret int64
		return ret
	}
	return *o.NodeId
}

// GetNodeIdOk returns a tuple with the NodeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceNodeStatus) GetNodeIdOk() (*int64, bool) {
	if o == nil || o.NodeId == nil {
		return nil, false
	}
	return o.NodeId, true
}

// HasNodeId returns a boolean if a field has been set.
func (o *ApplianceNodeStatus) HasNodeId() bool {
	if o != nil && o.NodeId != nil {
		return true
	}

	return false
}

// SetNodeId gets a reference to the given int64 and assigns it to the NodeId field.
func (o *ApplianceNodeStatus) SetNodeId(v int64) {
	o.NodeId = &v
}

// GetNodeState returns the NodeState field value if set, zero value otherwise.
func (o *ApplianceNodeStatus) GetNodeState() string {
	if o == nil || o.NodeState == nil {
		var ret string
		return ret
	}
	return *o.NodeState
}

// GetNodeStateOk returns a tuple with the NodeState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceNodeStatus) GetNodeStateOk() (*string, bool) {
	if o == nil || o.NodeState == nil {
		return nil, false
	}
	return o.NodeState, true
}

// HasNodeState returns a boolean if a field has been set.
func (o *ApplianceNodeStatus) HasNodeState() bool {
	if o != nil && o.NodeState != nil {
		return true
	}

	return false
}

// SetNodeState gets a reference to the given string and assigns it to the NodeState field.
func (o *ApplianceNodeStatus) SetNodeState(v string) {
	o.NodeState = &v
}

// GetOperationalStatus returns the OperationalStatus field value if set, zero value otherwise.
func (o *ApplianceNodeStatus) GetOperationalStatus() string {
	if o == nil || o.OperationalStatus == nil {
		var ret string
		return ret
	}
	return *o.OperationalStatus
}

// GetOperationalStatusOk returns a tuple with the OperationalStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceNodeStatus) GetOperationalStatusOk() (*string, bool) {
	if o == nil || o.OperationalStatus == nil {
		return nil, false
	}
	return o.OperationalStatus, true
}

// HasOperationalStatus returns a boolean if a field has been set.
func (o *ApplianceNodeStatus) HasOperationalStatus() bool {
	if o != nil && o.OperationalStatus != nil {
		return true
	}

	return false
}

// SetOperationalStatus gets a reference to the given string and assigns it to the OperationalStatus field.
func (o *ApplianceNodeStatus) SetOperationalStatus(v string) {
	o.OperationalStatus = &v
}

// GetStatusChecks returns the StatusChecks field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApplianceNodeStatus) GetStatusChecks() []ApplianceStatusCheck {
	if o == nil {
		var ret []ApplianceStatusCheck
		return ret
	}
	return o.StatusChecks
}

// GetStatusChecksOk returns a tuple with the StatusChecks field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApplianceNodeStatus) GetStatusChecksOk() (*[]ApplianceStatusCheck, bool) {
	if o == nil || o.StatusChecks == nil {
		return nil, false
	}
	return &o.StatusChecks, true
}

// HasStatusChecks returns a boolean if a field has been set.
func (o *ApplianceNodeStatus) HasStatusChecks() bool {
	if o != nil && o.StatusChecks != nil {
		return true
	}

	return false
}

// SetStatusChecks gets a reference to the given []ApplianceStatusCheck and assigns it to the StatusChecks field.
func (o *ApplianceNodeStatus) SetStatusChecks(v []ApplianceStatusCheck) {
	o.StatusChecks = v
}

// GetFileSystemStatuses returns the FileSystemStatuses field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApplianceNodeStatus) GetFileSystemStatuses() []ApplianceFileSystemStatusRelationship {
	if o == nil {
		var ret []ApplianceFileSystemStatusRelationship
		return ret
	}
	return o.FileSystemStatuses
}

// GetFileSystemStatusesOk returns a tuple with the FileSystemStatuses field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApplianceNodeStatus) GetFileSystemStatusesOk() (*[]ApplianceFileSystemStatusRelationship, bool) {
	if o == nil || o.FileSystemStatuses == nil {
		return nil, false
	}
	return &o.FileSystemStatuses, true
}

// HasFileSystemStatuses returns a boolean if a field has been set.
func (o *ApplianceNodeStatus) HasFileSystemStatuses() bool {
	if o != nil && o.FileSystemStatuses != nil {
		return true
	}

	return false
}

// SetFileSystemStatuses gets a reference to the given []ApplianceFileSystemStatusRelationship and assigns it to the FileSystemStatuses field.
func (o *ApplianceNodeStatus) SetFileSystemStatuses(v []ApplianceFileSystemStatusRelationship) {
	o.FileSystemStatuses = v
}

// GetNodeInfo returns the NodeInfo field value if set, zero value otherwise.
func (o *ApplianceNodeStatus) GetNodeInfo() ApplianceNodeInfoRelationship {
	if o == nil || o.NodeInfo == nil {
		var ret ApplianceNodeInfoRelationship
		return ret
	}
	return *o.NodeInfo
}

// GetNodeInfoOk returns a tuple with the NodeInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceNodeStatus) GetNodeInfoOk() (*ApplianceNodeInfoRelationship, bool) {
	if o == nil || o.NodeInfo == nil {
		return nil, false
	}
	return o.NodeInfo, true
}

// HasNodeInfo returns a boolean if a field has been set.
func (o *ApplianceNodeStatus) HasNodeInfo() bool {
	if o != nil && o.NodeInfo != nil {
		return true
	}

	return false
}

// SetNodeInfo gets a reference to the given ApplianceNodeInfoRelationship and assigns it to the NodeInfo field.
func (o *ApplianceNodeStatus) SetNodeInfo(v ApplianceNodeInfoRelationship) {
	o.NodeInfo = &v
}

// GetSystemStatus returns the SystemStatus field value if set, zero value otherwise.
func (o *ApplianceNodeStatus) GetSystemStatus() ApplianceSystemStatusRelationship {
	if o == nil || o.SystemStatus == nil {
		var ret ApplianceSystemStatusRelationship
		return ret
	}
	return *o.SystemStatus
}

// GetSystemStatusOk returns a tuple with the SystemStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceNodeStatus) GetSystemStatusOk() (*ApplianceSystemStatusRelationship, bool) {
	if o == nil || o.SystemStatus == nil {
		return nil, false
	}
	return o.SystemStatus, true
}

// HasSystemStatus returns a boolean if a field has been set.
func (o *ApplianceNodeStatus) HasSystemStatus() bool {
	if o != nil && o.SystemStatus != nil {
		return true
	}

	return false
}

// SetSystemStatus gets a reference to the given ApplianceSystemStatusRelationship and assigns it to the SystemStatus field.
func (o *ApplianceNodeStatus) SetSystemStatus(v ApplianceSystemStatusRelationship) {
	o.SystemStatus = &v
}

func (o ApplianceNodeStatus) MarshalJSON() ([]byte, error) {
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
	if o.CpuUsage != nil {
		toSerialize["CpuUsage"] = o.CpuUsage
	}
	if o.MemUsage != nil {
		toSerialize["MemUsage"] = o.MemUsage
	}
	if o.NodeId != nil {
		toSerialize["NodeId"] = o.NodeId
	}
	if o.NodeState != nil {
		toSerialize["NodeState"] = o.NodeState
	}
	if o.OperationalStatus != nil {
		toSerialize["OperationalStatus"] = o.OperationalStatus
	}
	if o.StatusChecks != nil {
		toSerialize["StatusChecks"] = o.StatusChecks
	}
	if o.FileSystemStatuses != nil {
		toSerialize["FileSystemStatuses"] = o.FileSystemStatuses
	}
	if o.NodeInfo != nil {
		toSerialize["NodeInfo"] = o.NodeInfo
	}
	if o.SystemStatus != nil {
		toSerialize["SystemStatus"] = o.SystemStatus
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ApplianceNodeStatus) UnmarshalJSON(bytes []byte) (err error) {
	type ApplianceNodeStatusWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Percentage of CPU currently in use.
		CpuUsage *float32 `json:"CpuUsage,omitempty"`
		// Percentage of memory currently in use.
		MemUsage *float32 `json:"MemUsage,omitempty"`
		// System assigned unique ID of the Intersight Appliance node. The system incrementally assigns identifiers to each node in the Intersight Appliance cluster starting with a value of 1.
		NodeId *int64 `json:"NodeId,omitempty"`
		// State of the node in terms of its readiness to host Kubernetes pods. * `Down` - The node is yet to come up and join as a member of theKubernetes cluster. * `Preparing` - The node has come up and joined the Kubernetes cluster,preparing to host Kubernetes pods. * `Ready` - The node is ready to host Kubernetes pods.
		NodeState *string `json:"NodeState,omitempty"`
		// Operational status of the Intersight Appliance node. Operational status is based on the result of the status checks. If result of any check is Critical, then its value is Impaired. Otherwise, if result of any check is Warning, then its value is AttentionNeeded. If all checks are OK, then its value is Operational. * `Unknown` - Operational status of the Intersight Appliance entity is Unknown. * `Operational` - Operational status of the Intersight Appliance entity is Operational. * `Impaired` - Operational status of the Intersight Appliance entity is Impaired. * `AttentionNeeded` - Operational status of the Intersight Appliance entity is AttentionNeeded.
		OperationalStatus *string                `json:"OperationalStatus,omitempty"`
		StatusChecks      []ApplianceStatusCheck `json:"StatusChecks,omitempty"`
		// An array of relationships to applianceFileSystemStatus resources.
		FileSystemStatuses []ApplianceFileSystemStatusRelationship `json:"FileSystemStatuses,omitempty"`
		NodeInfo           *ApplianceNodeInfoRelationship          `json:"NodeInfo,omitempty"`
		SystemStatus       *ApplianceSystemStatusRelationship      `json:"SystemStatus,omitempty"`
	}

	varApplianceNodeStatusWithoutEmbeddedStruct := ApplianceNodeStatusWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varApplianceNodeStatusWithoutEmbeddedStruct)
	if err == nil {
		varApplianceNodeStatus := _ApplianceNodeStatus{}
		varApplianceNodeStatus.ClassId = varApplianceNodeStatusWithoutEmbeddedStruct.ClassId
		varApplianceNodeStatus.ObjectType = varApplianceNodeStatusWithoutEmbeddedStruct.ObjectType
		varApplianceNodeStatus.CpuUsage = varApplianceNodeStatusWithoutEmbeddedStruct.CpuUsage
		varApplianceNodeStatus.MemUsage = varApplianceNodeStatusWithoutEmbeddedStruct.MemUsage
		varApplianceNodeStatus.NodeId = varApplianceNodeStatusWithoutEmbeddedStruct.NodeId
		varApplianceNodeStatus.NodeState = varApplianceNodeStatusWithoutEmbeddedStruct.NodeState
		varApplianceNodeStatus.OperationalStatus = varApplianceNodeStatusWithoutEmbeddedStruct.OperationalStatus
		varApplianceNodeStatus.StatusChecks = varApplianceNodeStatusWithoutEmbeddedStruct.StatusChecks
		varApplianceNodeStatus.FileSystemStatuses = varApplianceNodeStatusWithoutEmbeddedStruct.FileSystemStatuses
		varApplianceNodeStatus.NodeInfo = varApplianceNodeStatusWithoutEmbeddedStruct.NodeInfo
		varApplianceNodeStatus.SystemStatus = varApplianceNodeStatusWithoutEmbeddedStruct.SystemStatus
		*o = ApplianceNodeStatus(varApplianceNodeStatus)
	} else {
		return err
	}

	varApplianceNodeStatus := _ApplianceNodeStatus{}

	err = json.Unmarshal(bytes, &varApplianceNodeStatus)
	if err == nil {
		o.MoBaseMo = varApplianceNodeStatus.MoBaseMo
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "CpuUsage")
		delete(additionalProperties, "MemUsage")
		delete(additionalProperties, "NodeId")
		delete(additionalProperties, "NodeState")
		delete(additionalProperties, "OperationalStatus")
		delete(additionalProperties, "StatusChecks")
		delete(additionalProperties, "FileSystemStatuses")
		delete(additionalProperties, "NodeInfo")
		delete(additionalProperties, "SystemStatus")

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

type NullableApplianceNodeStatus struct {
	value *ApplianceNodeStatus
	isSet bool
}

func (v NullableApplianceNodeStatus) Get() *ApplianceNodeStatus {
	return v.value
}

func (v *NullableApplianceNodeStatus) Set(val *ApplianceNodeStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableApplianceNodeStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableApplianceNodeStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApplianceNodeStatus(val *ApplianceNodeStatus) *NullableApplianceNodeStatus {
	return &NullableApplianceNodeStatus{value: val, isSet: true}
}

func (v NullableApplianceNodeStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApplianceNodeStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
