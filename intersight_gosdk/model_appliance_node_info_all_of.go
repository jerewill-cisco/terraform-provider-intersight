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

// ApplianceNodeInfoAllOf Definition of the list of properties defined in 'appliance.NodeInfo', excluding properties defined in parent classes.
type ApplianceNodeInfoAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Cluster node's FQDN or IP address.
	Hostname *string `json:"Hostname,omitempty"`
	// System assigned unique ID of the Intersight Appliance node. The system incrementally assigns identifiers to each node in the Intersight Appliance cluster starting with a value of 1.
	NodeId         *int64                    `json:"NodeId,omitempty"`
	NodeIpV4Config NullableCommIpV4Interface `json:"NodeIpV4Config,omitempty"`
	NodeIpV6Config NullableCommIpV6Interface `json:"NodeIpV6Config,omitempty"`
	// Operational status of the Intersight Appliance node. * `Unknown` - Operational status of the Intersight Appliance entity is Unknown. * `Operational` - Operational status of the Intersight Appliance entity is Operational. * `Impaired` - Operational status of the Intersight Appliance entity is Impaired. * `AttentionNeeded` - Operational status of the Intersight Appliance entity is AttentionNeeded.
	OperationalStatus    *string `json:"OperationalStatus,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ApplianceNodeInfoAllOf ApplianceNodeInfoAllOf

// NewApplianceNodeInfoAllOf instantiates a new ApplianceNodeInfoAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApplianceNodeInfoAllOf(classId string, objectType string) *ApplianceNodeInfoAllOf {
	this := ApplianceNodeInfoAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewApplianceNodeInfoAllOfWithDefaults instantiates a new ApplianceNodeInfoAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApplianceNodeInfoAllOfWithDefaults() *ApplianceNodeInfoAllOf {
	this := ApplianceNodeInfoAllOf{}
	var classId string = "appliance.NodeInfo"
	this.ClassId = classId
	var objectType string = "appliance.NodeInfo"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *ApplianceNodeInfoAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *ApplianceNodeInfoAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *ApplianceNodeInfoAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *ApplianceNodeInfoAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *ApplianceNodeInfoAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *ApplianceNodeInfoAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetHostname returns the Hostname field value if set, zero value otherwise.
func (o *ApplianceNodeInfoAllOf) GetHostname() string {
	if o == nil || o.Hostname == nil {
		var ret string
		return ret
	}
	return *o.Hostname
}

// GetHostnameOk returns a tuple with the Hostname field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceNodeInfoAllOf) GetHostnameOk() (*string, bool) {
	if o == nil || o.Hostname == nil {
		return nil, false
	}
	return o.Hostname, true
}

// HasHostname returns a boolean if a field has been set.
func (o *ApplianceNodeInfoAllOf) HasHostname() bool {
	if o != nil && o.Hostname != nil {
		return true
	}

	return false
}

// SetHostname gets a reference to the given string and assigns it to the Hostname field.
func (o *ApplianceNodeInfoAllOf) SetHostname(v string) {
	o.Hostname = &v
}

// GetNodeId returns the NodeId field value if set, zero value otherwise.
func (o *ApplianceNodeInfoAllOf) GetNodeId() int64 {
	if o == nil || o.NodeId == nil {
		var ret int64
		return ret
	}
	return *o.NodeId
}

// GetNodeIdOk returns a tuple with the NodeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceNodeInfoAllOf) GetNodeIdOk() (*int64, bool) {
	if o == nil || o.NodeId == nil {
		return nil, false
	}
	return o.NodeId, true
}

// HasNodeId returns a boolean if a field has been set.
func (o *ApplianceNodeInfoAllOf) HasNodeId() bool {
	if o != nil && o.NodeId != nil {
		return true
	}

	return false
}

// SetNodeId gets a reference to the given int64 and assigns it to the NodeId field.
func (o *ApplianceNodeInfoAllOf) SetNodeId(v int64) {
	o.NodeId = &v
}

// GetNodeIpV4Config returns the NodeIpV4Config field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApplianceNodeInfoAllOf) GetNodeIpV4Config() CommIpV4Interface {
	if o == nil || o.NodeIpV4Config.Get() == nil {
		var ret CommIpV4Interface
		return ret
	}
	return *o.NodeIpV4Config.Get()
}

// GetNodeIpV4ConfigOk returns a tuple with the NodeIpV4Config field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApplianceNodeInfoAllOf) GetNodeIpV4ConfigOk() (*CommIpV4Interface, bool) {
	if o == nil {
		return nil, false
	}
	return o.NodeIpV4Config.Get(), o.NodeIpV4Config.IsSet()
}

// HasNodeIpV4Config returns a boolean if a field has been set.
func (o *ApplianceNodeInfoAllOf) HasNodeIpV4Config() bool {
	if o != nil && o.NodeIpV4Config.IsSet() {
		return true
	}

	return false
}

// SetNodeIpV4Config gets a reference to the given NullableCommIpV4Interface and assigns it to the NodeIpV4Config field.
func (o *ApplianceNodeInfoAllOf) SetNodeIpV4Config(v CommIpV4Interface) {
	o.NodeIpV4Config.Set(&v)
}

// SetNodeIpV4ConfigNil sets the value for NodeIpV4Config to be an explicit nil
func (o *ApplianceNodeInfoAllOf) SetNodeIpV4ConfigNil() {
	o.NodeIpV4Config.Set(nil)
}

// UnsetNodeIpV4Config ensures that no value is present for NodeIpV4Config, not even an explicit nil
func (o *ApplianceNodeInfoAllOf) UnsetNodeIpV4Config() {
	o.NodeIpV4Config.Unset()
}

// GetNodeIpV6Config returns the NodeIpV6Config field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApplianceNodeInfoAllOf) GetNodeIpV6Config() CommIpV6Interface {
	if o == nil || o.NodeIpV6Config.Get() == nil {
		var ret CommIpV6Interface
		return ret
	}
	return *o.NodeIpV6Config.Get()
}

// GetNodeIpV6ConfigOk returns a tuple with the NodeIpV6Config field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApplianceNodeInfoAllOf) GetNodeIpV6ConfigOk() (*CommIpV6Interface, bool) {
	if o == nil {
		return nil, false
	}
	return o.NodeIpV6Config.Get(), o.NodeIpV6Config.IsSet()
}

// HasNodeIpV6Config returns a boolean if a field has been set.
func (o *ApplianceNodeInfoAllOf) HasNodeIpV6Config() bool {
	if o != nil && o.NodeIpV6Config.IsSet() {
		return true
	}

	return false
}

// SetNodeIpV6Config gets a reference to the given NullableCommIpV6Interface and assigns it to the NodeIpV6Config field.
func (o *ApplianceNodeInfoAllOf) SetNodeIpV6Config(v CommIpV6Interface) {
	o.NodeIpV6Config.Set(&v)
}

// SetNodeIpV6ConfigNil sets the value for NodeIpV6Config to be an explicit nil
func (o *ApplianceNodeInfoAllOf) SetNodeIpV6ConfigNil() {
	o.NodeIpV6Config.Set(nil)
}

// UnsetNodeIpV6Config ensures that no value is present for NodeIpV6Config, not even an explicit nil
func (o *ApplianceNodeInfoAllOf) UnsetNodeIpV6Config() {
	o.NodeIpV6Config.Unset()
}

// GetOperationalStatus returns the OperationalStatus field value if set, zero value otherwise.
func (o *ApplianceNodeInfoAllOf) GetOperationalStatus() string {
	if o == nil || o.OperationalStatus == nil {
		var ret string
		return ret
	}
	return *o.OperationalStatus
}

// GetOperationalStatusOk returns a tuple with the OperationalStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceNodeInfoAllOf) GetOperationalStatusOk() (*string, bool) {
	if o == nil || o.OperationalStatus == nil {
		return nil, false
	}
	return o.OperationalStatus, true
}

// HasOperationalStatus returns a boolean if a field has been set.
func (o *ApplianceNodeInfoAllOf) HasOperationalStatus() bool {
	if o != nil && o.OperationalStatus != nil {
		return true
	}

	return false
}

// SetOperationalStatus gets a reference to the given string and assigns it to the OperationalStatus field.
func (o *ApplianceNodeInfoAllOf) SetOperationalStatus(v string) {
	o.OperationalStatus = &v
}

func (o ApplianceNodeInfoAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Hostname != nil {
		toSerialize["Hostname"] = o.Hostname
	}
	if o.NodeId != nil {
		toSerialize["NodeId"] = o.NodeId
	}
	if o.NodeIpV4Config.IsSet() {
		toSerialize["NodeIpV4Config"] = o.NodeIpV4Config.Get()
	}
	if o.NodeIpV6Config.IsSet() {
		toSerialize["NodeIpV6Config"] = o.NodeIpV6Config.Get()
	}
	if o.OperationalStatus != nil {
		toSerialize["OperationalStatus"] = o.OperationalStatus
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ApplianceNodeInfoAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varApplianceNodeInfoAllOf := _ApplianceNodeInfoAllOf{}

	if err = json.Unmarshal(bytes, &varApplianceNodeInfoAllOf); err == nil {
		*o = ApplianceNodeInfoAllOf(varApplianceNodeInfoAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Hostname")
		delete(additionalProperties, "NodeId")
		delete(additionalProperties, "NodeIpV4Config")
		delete(additionalProperties, "NodeIpV6Config")
		delete(additionalProperties, "OperationalStatus")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableApplianceNodeInfoAllOf struct {
	value *ApplianceNodeInfoAllOf
	isSet bool
}

func (v NullableApplianceNodeInfoAllOf) Get() *ApplianceNodeInfoAllOf {
	return v.value
}

func (v *NullableApplianceNodeInfoAllOf) Set(val *ApplianceNodeInfoAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableApplianceNodeInfoAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableApplianceNodeInfoAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApplianceNodeInfoAllOf(val *ApplianceNodeInfoAllOf) *NullableApplianceNodeInfoAllOf {
	return &NullableApplianceNodeInfoAllOf{value: val, isSet: true}
}

func (v NullableApplianceNodeInfoAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApplianceNodeInfoAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
