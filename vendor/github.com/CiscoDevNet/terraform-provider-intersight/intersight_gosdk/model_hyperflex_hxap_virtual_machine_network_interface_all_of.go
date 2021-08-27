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
)

// HyperflexHxapVirtualMachineNetworkInterfaceAllOf Definition of the list of properties defined in 'hyperflex.HxapVirtualMachineNetworkInterface', excluding properties defined in parent classes.
type HyperflexHxapVirtualMachineNetworkInterfaceAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Operating system assigned name for network interface.
	InterfaceName *string  `json:"InterfaceName,omitempty"`
	IpAddress     []string `json:"IpAddress,omitempty"`
	// Primary IP address of the network interface.
	PrimaryIpAddress *string `json:"PrimaryIpAddress,omitempty"`
	// Current status of virtual network interface status. * `Up` - Virtual network interface is up and running. * `Down` - Virtual network interface is down and not running.
	Status *string `json:"Status,omitempty"`
	// A reference to the virtual machine where this network object is attached to.
	VirtualMachineName   *string                                  `json:"VirtualMachineName,omitempty"`
	Cluster              *HyperflexHxapClusterRelationship        `json:"Cluster,omitempty"`
	Network              *HyperflexHxapNetworkRelationship        `json:"Network,omitempty"`
	VirtualMachine       *HyperflexHxapVirtualMachineRelationship `json:"VirtualMachine,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _HyperflexHxapVirtualMachineNetworkInterfaceAllOf HyperflexHxapVirtualMachineNetworkInterfaceAllOf

// NewHyperflexHxapVirtualMachineNetworkInterfaceAllOf instantiates a new HyperflexHxapVirtualMachineNetworkInterfaceAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHyperflexHxapVirtualMachineNetworkInterfaceAllOf(classId string, objectType string) *HyperflexHxapVirtualMachineNetworkInterfaceAllOf {
	this := HyperflexHxapVirtualMachineNetworkInterfaceAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	var status string = "Up"
	this.Status = &status
	return &this
}

// NewHyperflexHxapVirtualMachineNetworkInterfaceAllOfWithDefaults instantiates a new HyperflexHxapVirtualMachineNetworkInterfaceAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHyperflexHxapVirtualMachineNetworkInterfaceAllOfWithDefaults() *HyperflexHxapVirtualMachineNetworkInterfaceAllOf {
	this := HyperflexHxapVirtualMachineNetworkInterfaceAllOf{}
	var classId string = "hyperflex.HxapVirtualMachineNetworkInterface"
	this.ClassId = classId
	var objectType string = "hyperflex.HxapVirtualMachineNetworkInterface"
	this.ObjectType = objectType
	var status string = "Up"
	this.Status = &status
	return &this
}

// GetClassId returns the ClassId field value
func (o *HyperflexHxapVirtualMachineNetworkInterfaceAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *HyperflexHxapVirtualMachineNetworkInterfaceAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *HyperflexHxapVirtualMachineNetworkInterfaceAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *HyperflexHxapVirtualMachineNetworkInterfaceAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *HyperflexHxapVirtualMachineNetworkInterfaceAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *HyperflexHxapVirtualMachineNetworkInterfaceAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetInterfaceName returns the InterfaceName field value if set, zero value otherwise.
func (o *HyperflexHxapVirtualMachineNetworkInterfaceAllOf) GetInterfaceName() string {
	if o == nil || o.InterfaceName == nil {
		var ret string
		return ret
	}
	return *o.InterfaceName
}

// GetInterfaceNameOk returns a tuple with the InterfaceName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexHxapVirtualMachineNetworkInterfaceAllOf) GetInterfaceNameOk() (*string, bool) {
	if o == nil || o.InterfaceName == nil {
		return nil, false
	}
	return o.InterfaceName, true
}

// HasInterfaceName returns a boolean if a field has been set.
func (o *HyperflexHxapVirtualMachineNetworkInterfaceAllOf) HasInterfaceName() bool {
	if o != nil && o.InterfaceName != nil {
		return true
	}

	return false
}

// SetInterfaceName gets a reference to the given string and assigns it to the InterfaceName field.
func (o *HyperflexHxapVirtualMachineNetworkInterfaceAllOf) SetInterfaceName(v string) {
	o.InterfaceName = &v
}

// GetIpAddress returns the IpAddress field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HyperflexHxapVirtualMachineNetworkInterfaceAllOf) GetIpAddress() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.IpAddress
}

// GetIpAddressOk returns a tuple with the IpAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HyperflexHxapVirtualMachineNetworkInterfaceAllOf) GetIpAddressOk() (*[]string, bool) {
	if o == nil || o.IpAddress == nil {
		return nil, false
	}
	return &o.IpAddress, true
}

// HasIpAddress returns a boolean if a field has been set.
func (o *HyperflexHxapVirtualMachineNetworkInterfaceAllOf) HasIpAddress() bool {
	if o != nil && o.IpAddress != nil {
		return true
	}

	return false
}

// SetIpAddress gets a reference to the given []string and assigns it to the IpAddress field.
func (o *HyperflexHxapVirtualMachineNetworkInterfaceAllOf) SetIpAddress(v []string) {
	o.IpAddress = v
}

// GetPrimaryIpAddress returns the PrimaryIpAddress field value if set, zero value otherwise.
func (o *HyperflexHxapVirtualMachineNetworkInterfaceAllOf) GetPrimaryIpAddress() string {
	if o == nil || o.PrimaryIpAddress == nil {
		var ret string
		return ret
	}
	return *o.PrimaryIpAddress
}

// GetPrimaryIpAddressOk returns a tuple with the PrimaryIpAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexHxapVirtualMachineNetworkInterfaceAllOf) GetPrimaryIpAddressOk() (*string, bool) {
	if o == nil || o.PrimaryIpAddress == nil {
		return nil, false
	}
	return o.PrimaryIpAddress, true
}

// HasPrimaryIpAddress returns a boolean if a field has been set.
func (o *HyperflexHxapVirtualMachineNetworkInterfaceAllOf) HasPrimaryIpAddress() bool {
	if o != nil && o.PrimaryIpAddress != nil {
		return true
	}

	return false
}

// SetPrimaryIpAddress gets a reference to the given string and assigns it to the PrimaryIpAddress field.
func (o *HyperflexHxapVirtualMachineNetworkInterfaceAllOf) SetPrimaryIpAddress(v string) {
	o.PrimaryIpAddress = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *HyperflexHxapVirtualMachineNetworkInterfaceAllOf) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexHxapVirtualMachineNetworkInterfaceAllOf) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *HyperflexHxapVirtualMachineNetworkInterfaceAllOf) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *HyperflexHxapVirtualMachineNetworkInterfaceAllOf) SetStatus(v string) {
	o.Status = &v
}

// GetVirtualMachineName returns the VirtualMachineName field value if set, zero value otherwise.
func (o *HyperflexHxapVirtualMachineNetworkInterfaceAllOf) GetVirtualMachineName() string {
	if o == nil || o.VirtualMachineName == nil {
		var ret string
		return ret
	}
	return *o.VirtualMachineName
}

// GetVirtualMachineNameOk returns a tuple with the VirtualMachineName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexHxapVirtualMachineNetworkInterfaceAllOf) GetVirtualMachineNameOk() (*string, bool) {
	if o == nil || o.VirtualMachineName == nil {
		return nil, false
	}
	return o.VirtualMachineName, true
}

// HasVirtualMachineName returns a boolean if a field has been set.
func (o *HyperflexHxapVirtualMachineNetworkInterfaceAllOf) HasVirtualMachineName() bool {
	if o != nil && o.VirtualMachineName != nil {
		return true
	}

	return false
}

// SetVirtualMachineName gets a reference to the given string and assigns it to the VirtualMachineName field.
func (o *HyperflexHxapVirtualMachineNetworkInterfaceAllOf) SetVirtualMachineName(v string) {
	o.VirtualMachineName = &v
}

// GetCluster returns the Cluster field value if set, zero value otherwise.
func (o *HyperflexHxapVirtualMachineNetworkInterfaceAllOf) GetCluster() HyperflexHxapClusterRelationship {
	if o == nil || o.Cluster == nil {
		var ret HyperflexHxapClusterRelationship
		return ret
	}
	return *o.Cluster
}

// GetClusterOk returns a tuple with the Cluster field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexHxapVirtualMachineNetworkInterfaceAllOf) GetClusterOk() (*HyperflexHxapClusterRelationship, bool) {
	if o == nil || o.Cluster == nil {
		return nil, false
	}
	return o.Cluster, true
}

// HasCluster returns a boolean if a field has been set.
func (o *HyperflexHxapVirtualMachineNetworkInterfaceAllOf) HasCluster() bool {
	if o != nil && o.Cluster != nil {
		return true
	}

	return false
}

// SetCluster gets a reference to the given HyperflexHxapClusterRelationship and assigns it to the Cluster field.
func (o *HyperflexHxapVirtualMachineNetworkInterfaceAllOf) SetCluster(v HyperflexHxapClusterRelationship) {
	o.Cluster = &v
}

// GetNetwork returns the Network field value if set, zero value otherwise.
func (o *HyperflexHxapVirtualMachineNetworkInterfaceAllOf) GetNetwork() HyperflexHxapNetworkRelationship {
	if o == nil || o.Network == nil {
		var ret HyperflexHxapNetworkRelationship
		return ret
	}
	return *o.Network
}

// GetNetworkOk returns a tuple with the Network field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexHxapVirtualMachineNetworkInterfaceAllOf) GetNetworkOk() (*HyperflexHxapNetworkRelationship, bool) {
	if o == nil || o.Network == nil {
		return nil, false
	}
	return o.Network, true
}

// HasNetwork returns a boolean if a field has been set.
func (o *HyperflexHxapVirtualMachineNetworkInterfaceAllOf) HasNetwork() bool {
	if o != nil && o.Network != nil {
		return true
	}

	return false
}

// SetNetwork gets a reference to the given HyperflexHxapNetworkRelationship and assigns it to the Network field.
func (o *HyperflexHxapVirtualMachineNetworkInterfaceAllOf) SetNetwork(v HyperflexHxapNetworkRelationship) {
	o.Network = &v
}

// GetVirtualMachine returns the VirtualMachine field value if set, zero value otherwise.
func (o *HyperflexHxapVirtualMachineNetworkInterfaceAllOf) GetVirtualMachine() HyperflexHxapVirtualMachineRelationship {
	if o == nil || o.VirtualMachine == nil {
		var ret HyperflexHxapVirtualMachineRelationship
		return ret
	}
	return *o.VirtualMachine
}

// GetVirtualMachineOk returns a tuple with the VirtualMachine field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexHxapVirtualMachineNetworkInterfaceAllOf) GetVirtualMachineOk() (*HyperflexHxapVirtualMachineRelationship, bool) {
	if o == nil || o.VirtualMachine == nil {
		return nil, false
	}
	return o.VirtualMachine, true
}

// HasVirtualMachine returns a boolean if a field has been set.
func (o *HyperflexHxapVirtualMachineNetworkInterfaceAllOf) HasVirtualMachine() bool {
	if o != nil && o.VirtualMachine != nil {
		return true
	}

	return false
}

// SetVirtualMachine gets a reference to the given HyperflexHxapVirtualMachineRelationship and assigns it to the VirtualMachine field.
func (o *HyperflexHxapVirtualMachineNetworkInterfaceAllOf) SetVirtualMachine(v HyperflexHxapVirtualMachineRelationship) {
	o.VirtualMachine = &v
}

func (o HyperflexHxapVirtualMachineNetworkInterfaceAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.InterfaceName != nil {
		toSerialize["InterfaceName"] = o.InterfaceName
	}
	if o.IpAddress != nil {
		toSerialize["IpAddress"] = o.IpAddress
	}
	if o.PrimaryIpAddress != nil {
		toSerialize["PrimaryIpAddress"] = o.PrimaryIpAddress
	}
	if o.Status != nil {
		toSerialize["Status"] = o.Status
	}
	if o.VirtualMachineName != nil {
		toSerialize["VirtualMachineName"] = o.VirtualMachineName
	}
	if o.Cluster != nil {
		toSerialize["Cluster"] = o.Cluster
	}
	if o.Network != nil {
		toSerialize["Network"] = o.Network
	}
	if o.VirtualMachine != nil {
		toSerialize["VirtualMachine"] = o.VirtualMachine
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *HyperflexHxapVirtualMachineNetworkInterfaceAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varHyperflexHxapVirtualMachineNetworkInterfaceAllOf := _HyperflexHxapVirtualMachineNetworkInterfaceAllOf{}

	if err = json.Unmarshal(bytes, &varHyperflexHxapVirtualMachineNetworkInterfaceAllOf); err == nil {
		*o = HyperflexHxapVirtualMachineNetworkInterfaceAllOf(varHyperflexHxapVirtualMachineNetworkInterfaceAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "InterfaceName")
		delete(additionalProperties, "IpAddress")
		delete(additionalProperties, "PrimaryIpAddress")
		delete(additionalProperties, "Status")
		delete(additionalProperties, "VirtualMachineName")
		delete(additionalProperties, "Cluster")
		delete(additionalProperties, "Network")
		delete(additionalProperties, "VirtualMachine")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableHyperflexHxapVirtualMachineNetworkInterfaceAllOf struct {
	value *HyperflexHxapVirtualMachineNetworkInterfaceAllOf
	isSet bool
}

func (v NullableHyperflexHxapVirtualMachineNetworkInterfaceAllOf) Get() *HyperflexHxapVirtualMachineNetworkInterfaceAllOf {
	return v.value
}

func (v *NullableHyperflexHxapVirtualMachineNetworkInterfaceAllOf) Set(val *HyperflexHxapVirtualMachineNetworkInterfaceAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableHyperflexHxapVirtualMachineNetworkInterfaceAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableHyperflexHxapVirtualMachineNetworkInterfaceAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHyperflexHxapVirtualMachineNetworkInterfaceAllOf(val *HyperflexHxapVirtualMachineNetworkInterfaceAllOf) *NullableHyperflexHxapVirtualMachineNetworkInterfaceAllOf {
	return &NullableHyperflexHxapVirtualMachineNetworkInterfaceAllOf{value: val, isSet: true}
}

func (v NullableHyperflexHxapVirtualMachineNetworkInterfaceAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHyperflexHxapVirtualMachineNetworkInterfaceAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
