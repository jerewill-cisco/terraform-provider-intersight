/*
 * Cisco Intersight
 *
 * Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document. This document was created on 2021-06-30T12:14:04Z.
 *
 * API version: 1.0.9-4375
 * Contact: intersight@cisco.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package intersight

import (
	"encoding/json"
)

// HyperflexHxapHostVswitchAllOf Definition of the list of properties defined in 'hyperflex.HxapHostVswitch', excluding properties defined in parent classes.
type HyperflexHxapHostVswitchAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// The name of the host to which this vSwitch belongs to.
	HostName             *string                            `json:"HostName,omitempty"`
	Ports                []HyperflexNetworkPort             `json:"Ports,omitempty"`
	Cluster              *HyperflexHxapClusterRelationship  `json:"Cluster,omitempty"`
	DvSwitch             *HyperflexHxapDvswitchRelationship `json:"DvSwitch,omitempty"`
	Host                 *HyperflexHxapHostRelationship     `json:"Host,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _HyperflexHxapHostVswitchAllOf HyperflexHxapHostVswitchAllOf

// NewHyperflexHxapHostVswitchAllOf instantiates a new HyperflexHxapHostVswitchAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHyperflexHxapHostVswitchAllOf(classId string, objectType string) *HyperflexHxapHostVswitchAllOf {
	this := HyperflexHxapHostVswitchAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewHyperflexHxapHostVswitchAllOfWithDefaults instantiates a new HyperflexHxapHostVswitchAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHyperflexHxapHostVswitchAllOfWithDefaults() *HyperflexHxapHostVswitchAllOf {
	this := HyperflexHxapHostVswitchAllOf{}
	var classId string = "hyperflex.HxapHostVswitch"
	this.ClassId = classId
	var objectType string = "hyperflex.HxapHostVswitch"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *HyperflexHxapHostVswitchAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *HyperflexHxapHostVswitchAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *HyperflexHxapHostVswitchAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *HyperflexHxapHostVswitchAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *HyperflexHxapHostVswitchAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *HyperflexHxapHostVswitchAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetHostName returns the HostName field value if set, zero value otherwise.
func (o *HyperflexHxapHostVswitchAllOf) GetHostName() string {
	if o == nil || o.HostName == nil {
		var ret string
		return ret
	}
	return *o.HostName
}

// GetHostNameOk returns a tuple with the HostName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexHxapHostVswitchAllOf) GetHostNameOk() (*string, bool) {
	if o == nil || o.HostName == nil {
		return nil, false
	}
	return o.HostName, true
}

// HasHostName returns a boolean if a field has been set.
func (o *HyperflexHxapHostVswitchAllOf) HasHostName() bool {
	if o != nil && o.HostName != nil {
		return true
	}

	return false
}

// SetHostName gets a reference to the given string and assigns it to the HostName field.
func (o *HyperflexHxapHostVswitchAllOf) SetHostName(v string) {
	o.HostName = &v
}

// GetPorts returns the Ports field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HyperflexHxapHostVswitchAllOf) GetPorts() []HyperflexNetworkPort {
	if o == nil {
		var ret []HyperflexNetworkPort
		return ret
	}
	return o.Ports
}

// GetPortsOk returns a tuple with the Ports field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HyperflexHxapHostVswitchAllOf) GetPortsOk() (*[]HyperflexNetworkPort, bool) {
	if o == nil || o.Ports == nil {
		return nil, false
	}
	return &o.Ports, true
}

// HasPorts returns a boolean if a field has been set.
func (o *HyperflexHxapHostVswitchAllOf) HasPorts() bool {
	if o != nil && o.Ports != nil {
		return true
	}

	return false
}

// SetPorts gets a reference to the given []HyperflexNetworkPort and assigns it to the Ports field.
func (o *HyperflexHxapHostVswitchAllOf) SetPorts(v []HyperflexNetworkPort) {
	o.Ports = v
}

// GetCluster returns the Cluster field value if set, zero value otherwise.
func (o *HyperflexHxapHostVswitchAllOf) GetCluster() HyperflexHxapClusterRelationship {
	if o == nil || o.Cluster == nil {
		var ret HyperflexHxapClusterRelationship
		return ret
	}
	return *o.Cluster
}

// GetClusterOk returns a tuple with the Cluster field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexHxapHostVswitchAllOf) GetClusterOk() (*HyperflexHxapClusterRelationship, bool) {
	if o == nil || o.Cluster == nil {
		return nil, false
	}
	return o.Cluster, true
}

// HasCluster returns a boolean if a field has been set.
func (o *HyperflexHxapHostVswitchAllOf) HasCluster() bool {
	if o != nil && o.Cluster != nil {
		return true
	}

	return false
}

// SetCluster gets a reference to the given HyperflexHxapClusterRelationship and assigns it to the Cluster field.
func (o *HyperflexHxapHostVswitchAllOf) SetCluster(v HyperflexHxapClusterRelationship) {
	o.Cluster = &v
}

// GetDvSwitch returns the DvSwitch field value if set, zero value otherwise.
func (o *HyperflexHxapHostVswitchAllOf) GetDvSwitch() HyperflexHxapDvswitchRelationship {
	if o == nil || o.DvSwitch == nil {
		var ret HyperflexHxapDvswitchRelationship
		return ret
	}
	return *o.DvSwitch
}

// GetDvSwitchOk returns a tuple with the DvSwitch field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexHxapHostVswitchAllOf) GetDvSwitchOk() (*HyperflexHxapDvswitchRelationship, bool) {
	if o == nil || o.DvSwitch == nil {
		return nil, false
	}
	return o.DvSwitch, true
}

// HasDvSwitch returns a boolean if a field has been set.
func (o *HyperflexHxapHostVswitchAllOf) HasDvSwitch() bool {
	if o != nil && o.DvSwitch != nil {
		return true
	}

	return false
}

// SetDvSwitch gets a reference to the given HyperflexHxapDvswitchRelationship and assigns it to the DvSwitch field.
func (o *HyperflexHxapHostVswitchAllOf) SetDvSwitch(v HyperflexHxapDvswitchRelationship) {
	o.DvSwitch = &v
}

// GetHost returns the Host field value if set, zero value otherwise.
func (o *HyperflexHxapHostVswitchAllOf) GetHost() HyperflexHxapHostRelationship {
	if o == nil || o.Host == nil {
		var ret HyperflexHxapHostRelationship
		return ret
	}
	return *o.Host
}

// GetHostOk returns a tuple with the Host field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexHxapHostVswitchAllOf) GetHostOk() (*HyperflexHxapHostRelationship, bool) {
	if o == nil || o.Host == nil {
		return nil, false
	}
	return o.Host, true
}

// HasHost returns a boolean if a field has been set.
func (o *HyperflexHxapHostVswitchAllOf) HasHost() bool {
	if o != nil && o.Host != nil {
		return true
	}

	return false
}

// SetHost gets a reference to the given HyperflexHxapHostRelationship and assigns it to the Host field.
func (o *HyperflexHxapHostVswitchAllOf) SetHost(v HyperflexHxapHostRelationship) {
	o.Host = &v
}

func (o HyperflexHxapHostVswitchAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.HostName != nil {
		toSerialize["HostName"] = o.HostName
	}
	if o.Ports != nil {
		toSerialize["Ports"] = o.Ports
	}
	if o.Cluster != nil {
		toSerialize["Cluster"] = o.Cluster
	}
	if o.DvSwitch != nil {
		toSerialize["DvSwitch"] = o.DvSwitch
	}
	if o.Host != nil {
		toSerialize["Host"] = o.Host
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *HyperflexHxapHostVswitchAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varHyperflexHxapHostVswitchAllOf := _HyperflexHxapHostVswitchAllOf{}

	if err = json.Unmarshal(bytes, &varHyperflexHxapHostVswitchAllOf); err == nil {
		*o = HyperflexHxapHostVswitchAllOf(varHyperflexHxapHostVswitchAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "HostName")
		delete(additionalProperties, "Ports")
		delete(additionalProperties, "Cluster")
		delete(additionalProperties, "DvSwitch")
		delete(additionalProperties, "Host")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableHyperflexHxapHostVswitchAllOf struct {
	value *HyperflexHxapHostVswitchAllOf
	isSet bool
}

func (v NullableHyperflexHxapHostVswitchAllOf) Get() *HyperflexHxapHostVswitchAllOf {
	return v.value
}

func (v *NullableHyperflexHxapHostVswitchAllOf) Set(val *HyperflexHxapHostVswitchAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableHyperflexHxapHostVswitchAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableHyperflexHxapHostVswitchAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHyperflexHxapHostVswitchAllOf(val *HyperflexHxapHostVswitchAllOf) *NullableHyperflexHxapHostVswitchAllOf {
	return &NullableHyperflexHxapHostVswitchAllOf{value: val, isSet: true}
}

func (v NullableHyperflexHxapHostVswitchAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHyperflexHxapHostVswitchAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
