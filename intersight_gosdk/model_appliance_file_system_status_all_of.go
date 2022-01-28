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

// ApplianceFileSystemStatusAllOf Definition of the list of properties defined in 'appliance.FileSystemStatus', excluding properties defined in parent classes.
type ApplianceFileSystemStatusAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Capacity of the file system in bytes.
	Capacity *int64 `json:"Capacity,omitempty"`
	// Mount point of this file system.
	Mountpoint *string `json:"Mountpoint,omitempty"`
	// Operational status of the file system. Operational status is based on the result of the status checks. If result of any check is Critical, then its value is Impaired. Otherwise, if result of any check is Warning, then its value is AttentionNeeded. If all checks are OK, then its value is Operational. * `Unknown` - Operational status of the Intersight Appliance entity is Unknown. * `Operational` - Operational status of the Intersight Appliance entity is Operational. * `Impaired` - Operational status of the Intersight Appliance entity is Impaired. * `AttentionNeeded` - Operational status of the Intersight Appliance entity is AttentionNeeded.
	OperationalStatus *string                `json:"OperationalStatus,omitempty"`
	StatusChecks      []ApplianceStatusCheck `json:"StatusChecks,omitempty"`
	// Percentage of the file system capacity currently in use.
	Usage                *float32                         `json:"Usage,omitempty"`
	NodeStatus           *ApplianceNodeStatusRelationship `json:"NodeStatus,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ApplianceFileSystemStatusAllOf ApplianceFileSystemStatusAllOf

// NewApplianceFileSystemStatusAllOf instantiates a new ApplianceFileSystemStatusAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApplianceFileSystemStatusAllOf(classId string, objectType string) *ApplianceFileSystemStatusAllOf {
	this := ApplianceFileSystemStatusAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewApplianceFileSystemStatusAllOfWithDefaults instantiates a new ApplianceFileSystemStatusAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApplianceFileSystemStatusAllOfWithDefaults() *ApplianceFileSystemStatusAllOf {
	this := ApplianceFileSystemStatusAllOf{}
	var classId string = "appliance.FileSystemStatus"
	this.ClassId = classId
	var objectType string = "appliance.FileSystemStatus"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *ApplianceFileSystemStatusAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *ApplianceFileSystemStatusAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *ApplianceFileSystemStatusAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *ApplianceFileSystemStatusAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *ApplianceFileSystemStatusAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *ApplianceFileSystemStatusAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetCapacity returns the Capacity field value if set, zero value otherwise.
func (o *ApplianceFileSystemStatusAllOf) GetCapacity() int64 {
	if o == nil || o.Capacity == nil {
		var ret int64
		return ret
	}
	return *o.Capacity
}

// GetCapacityOk returns a tuple with the Capacity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceFileSystemStatusAllOf) GetCapacityOk() (*int64, bool) {
	if o == nil || o.Capacity == nil {
		return nil, false
	}
	return o.Capacity, true
}

// HasCapacity returns a boolean if a field has been set.
func (o *ApplianceFileSystemStatusAllOf) HasCapacity() bool {
	if o != nil && o.Capacity != nil {
		return true
	}

	return false
}

// SetCapacity gets a reference to the given int64 and assigns it to the Capacity field.
func (o *ApplianceFileSystemStatusAllOf) SetCapacity(v int64) {
	o.Capacity = &v
}

// GetMountpoint returns the Mountpoint field value if set, zero value otherwise.
func (o *ApplianceFileSystemStatusAllOf) GetMountpoint() string {
	if o == nil || o.Mountpoint == nil {
		var ret string
		return ret
	}
	return *o.Mountpoint
}

// GetMountpointOk returns a tuple with the Mountpoint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceFileSystemStatusAllOf) GetMountpointOk() (*string, bool) {
	if o == nil || o.Mountpoint == nil {
		return nil, false
	}
	return o.Mountpoint, true
}

// HasMountpoint returns a boolean if a field has been set.
func (o *ApplianceFileSystemStatusAllOf) HasMountpoint() bool {
	if o != nil && o.Mountpoint != nil {
		return true
	}

	return false
}

// SetMountpoint gets a reference to the given string and assigns it to the Mountpoint field.
func (o *ApplianceFileSystemStatusAllOf) SetMountpoint(v string) {
	o.Mountpoint = &v
}

// GetOperationalStatus returns the OperationalStatus field value if set, zero value otherwise.
func (o *ApplianceFileSystemStatusAllOf) GetOperationalStatus() string {
	if o == nil || o.OperationalStatus == nil {
		var ret string
		return ret
	}
	return *o.OperationalStatus
}

// GetOperationalStatusOk returns a tuple with the OperationalStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceFileSystemStatusAllOf) GetOperationalStatusOk() (*string, bool) {
	if o == nil || o.OperationalStatus == nil {
		return nil, false
	}
	return o.OperationalStatus, true
}

// HasOperationalStatus returns a boolean if a field has been set.
func (o *ApplianceFileSystemStatusAllOf) HasOperationalStatus() bool {
	if o != nil && o.OperationalStatus != nil {
		return true
	}

	return false
}

// SetOperationalStatus gets a reference to the given string and assigns it to the OperationalStatus field.
func (o *ApplianceFileSystemStatusAllOf) SetOperationalStatus(v string) {
	o.OperationalStatus = &v
}

// GetStatusChecks returns the StatusChecks field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApplianceFileSystemStatusAllOf) GetStatusChecks() []ApplianceStatusCheck {
	if o == nil {
		var ret []ApplianceStatusCheck
		return ret
	}
	return o.StatusChecks
}

// GetStatusChecksOk returns a tuple with the StatusChecks field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApplianceFileSystemStatusAllOf) GetStatusChecksOk() (*[]ApplianceStatusCheck, bool) {
	if o == nil || o.StatusChecks == nil {
		return nil, false
	}
	return &o.StatusChecks, true
}

// HasStatusChecks returns a boolean if a field has been set.
func (o *ApplianceFileSystemStatusAllOf) HasStatusChecks() bool {
	if o != nil && o.StatusChecks != nil {
		return true
	}

	return false
}

// SetStatusChecks gets a reference to the given []ApplianceStatusCheck and assigns it to the StatusChecks field.
func (o *ApplianceFileSystemStatusAllOf) SetStatusChecks(v []ApplianceStatusCheck) {
	o.StatusChecks = v
}

// GetUsage returns the Usage field value if set, zero value otherwise.
func (o *ApplianceFileSystemStatusAllOf) GetUsage() float32 {
	if o == nil || o.Usage == nil {
		var ret float32
		return ret
	}
	return *o.Usage
}

// GetUsageOk returns a tuple with the Usage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceFileSystemStatusAllOf) GetUsageOk() (*float32, bool) {
	if o == nil || o.Usage == nil {
		return nil, false
	}
	return o.Usage, true
}

// HasUsage returns a boolean if a field has been set.
func (o *ApplianceFileSystemStatusAllOf) HasUsage() bool {
	if o != nil && o.Usage != nil {
		return true
	}

	return false
}

// SetUsage gets a reference to the given float32 and assigns it to the Usage field.
func (o *ApplianceFileSystemStatusAllOf) SetUsage(v float32) {
	o.Usage = &v
}

// GetNodeStatus returns the NodeStatus field value if set, zero value otherwise.
func (o *ApplianceFileSystemStatusAllOf) GetNodeStatus() ApplianceNodeStatusRelationship {
	if o == nil || o.NodeStatus == nil {
		var ret ApplianceNodeStatusRelationship
		return ret
	}
	return *o.NodeStatus
}

// GetNodeStatusOk returns a tuple with the NodeStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceFileSystemStatusAllOf) GetNodeStatusOk() (*ApplianceNodeStatusRelationship, bool) {
	if o == nil || o.NodeStatus == nil {
		return nil, false
	}
	return o.NodeStatus, true
}

// HasNodeStatus returns a boolean if a field has been set.
func (o *ApplianceFileSystemStatusAllOf) HasNodeStatus() bool {
	if o != nil && o.NodeStatus != nil {
		return true
	}

	return false
}

// SetNodeStatus gets a reference to the given ApplianceNodeStatusRelationship and assigns it to the NodeStatus field.
func (o *ApplianceFileSystemStatusAllOf) SetNodeStatus(v ApplianceNodeStatusRelationship) {
	o.NodeStatus = &v
}

func (o ApplianceFileSystemStatusAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Capacity != nil {
		toSerialize["Capacity"] = o.Capacity
	}
	if o.Mountpoint != nil {
		toSerialize["Mountpoint"] = o.Mountpoint
	}
	if o.OperationalStatus != nil {
		toSerialize["OperationalStatus"] = o.OperationalStatus
	}
	if o.StatusChecks != nil {
		toSerialize["StatusChecks"] = o.StatusChecks
	}
	if o.Usage != nil {
		toSerialize["Usage"] = o.Usage
	}
	if o.NodeStatus != nil {
		toSerialize["NodeStatus"] = o.NodeStatus
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ApplianceFileSystemStatusAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varApplianceFileSystemStatusAllOf := _ApplianceFileSystemStatusAllOf{}

	if err = json.Unmarshal(bytes, &varApplianceFileSystemStatusAllOf); err == nil {
		*o = ApplianceFileSystemStatusAllOf(varApplianceFileSystemStatusAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Capacity")
		delete(additionalProperties, "Mountpoint")
		delete(additionalProperties, "OperationalStatus")
		delete(additionalProperties, "StatusChecks")
		delete(additionalProperties, "Usage")
		delete(additionalProperties, "NodeStatus")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableApplianceFileSystemStatusAllOf struct {
	value *ApplianceFileSystemStatusAllOf
	isSet bool
}

func (v NullableApplianceFileSystemStatusAllOf) Get() *ApplianceFileSystemStatusAllOf {
	return v.value
}

func (v *NullableApplianceFileSystemStatusAllOf) Set(val *ApplianceFileSystemStatusAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableApplianceFileSystemStatusAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableApplianceFileSystemStatusAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApplianceFileSystemStatusAllOf(val *ApplianceFileSystemStatusAllOf) *NullableApplianceFileSystemStatusAllOf {
	return &NullableApplianceFileSystemStatusAllOf{value: val, isSet: true}
}

func (v NullableApplianceFileSystemStatusAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApplianceFileSystemStatusAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
