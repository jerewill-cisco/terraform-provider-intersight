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

// HyperflexHxapHostVswitch A HyperFlex Application Platform vSwitch entity that is part of a cluster wide dvSwitch.
type HyperflexHxapHostVswitch struct {
	VirtualizationBaseVswitch
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

type _HyperflexHxapHostVswitch HyperflexHxapHostVswitch

// NewHyperflexHxapHostVswitch instantiates a new HyperflexHxapHostVswitch object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHyperflexHxapHostVswitch(classId string, objectType string) *HyperflexHxapHostVswitch {
	this := HyperflexHxapHostVswitch{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewHyperflexHxapHostVswitchWithDefaults instantiates a new HyperflexHxapHostVswitch object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHyperflexHxapHostVswitchWithDefaults() *HyperflexHxapHostVswitch {
	this := HyperflexHxapHostVswitch{}
	var classId string = "hyperflex.HxapHostVswitch"
	this.ClassId = classId
	var objectType string = "hyperflex.HxapHostVswitch"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *HyperflexHxapHostVswitch) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *HyperflexHxapHostVswitch) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *HyperflexHxapHostVswitch) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *HyperflexHxapHostVswitch) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *HyperflexHxapHostVswitch) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *HyperflexHxapHostVswitch) SetObjectType(v string) {
	o.ObjectType = v
}

// GetHostName returns the HostName field value if set, zero value otherwise.
func (o *HyperflexHxapHostVswitch) GetHostName() string {
	if o == nil || o.HostName == nil {
		var ret string
		return ret
	}
	return *o.HostName
}

// GetHostNameOk returns a tuple with the HostName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexHxapHostVswitch) GetHostNameOk() (*string, bool) {
	if o == nil || o.HostName == nil {
		return nil, false
	}
	return o.HostName, true
}

// HasHostName returns a boolean if a field has been set.
func (o *HyperflexHxapHostVswitch) HasHostName() bool {
	if o != nil && o.HostName != nil {
		return true
	}

	return false
}

// SetHostName gets a reference to the given string and assigns it to the HostName field.
func (o *HyperflexHxapHostVswitch) SetHostName(v string) {
	o.HostName = &v
}

// GetPorts returns the Ports field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HyperflexHxapHostVswitch) GetPorts() []HyperflexNetworkPort {
	if o == nil {
		var ret []HyperflexNetworkPort
		return ret
	}
	return o.Ports
}

// GetPortsOk returns a tuple with the Ports field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HyperflexHxapHostVswitch) GetPortsOk() (*[]HyperflexNetworkPort, bool) {
	if o == nil || o.Ports == nil {
		return nil, false
	}
	return &o.Ports, true
}

// HasPorts returns a boolean if a field has been set.
func (o *HyperflexHxapHostVswitch) HasPorts() bool {
	if o != nil && o.Ports != nil {
		return true
	}

	return false
}

// SetPorts gets a reference to the given []HyperflexNetworkPort and assigns it to the Ports field.
func (o *HyperflexHxapHostVswitch) SetPorts(v []HyperflexNetworkPort) {
	o.Ports = v
}

// GetCluster returns the Cluster field value if set, zero value otherwise.
func (o *HyperflexHxapHostVswitch) GetCluster() HyperflexHxapClusterRelationship {
	if o == nil || o.Cluster == nil {
		var ret HyperflexHxapClusterRelationship
		return ret
	}
	return *o.Cluster
}

// GetClusterOk returns a tuple with the Cluster field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexHxapHostVswitch) GetClusterOk() (*HyperflexHxapClusterRelationship, bool) {
	if o == nil || o.Cluster == nil {
		return nil, false
	}
	return o.Cluster, true
}

// HasCluster returns a boolean if a field has been set.
func (o *HyperflexHxapHostVswitch) HasCluster() bool {
	if o != nil && o.Cluster != nil {
		return true
	}

	return false
}

// SetCluster gets a reference to the given HyperflexHxapClusterRelationship and assigns it to the Cluster field.
func (o *HyperflexHxapHostVswitch) SetCluster(v HyperflexHxapClusterRelationship) {
	o.Cluster = &v
}

// GetDvSwitch returns the DvSwitch field value if set, zero value otherwise.
func (o *HyperflexHxapHostVswitch) GetDvSwitch() HyperflexHxapDvswitchRelationship {
	if o == nil || o.DvSwitch == nil {
		var ret HyperflexHxapDvswitchRelationship
		return ret
	}
	return *o.DvSwitch
}

// GetDvSwitchOk returns a tuple with the DvSwitch field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexHxapHostVswitch) GetDvSwitchOk() (*HyperflexHxapDvswitchRelationship, bool) {
	if o == nil || o.DvSwitch == nil {
		return nil, false
	}
	return o.DvSwitch, true
}

// HasDvSwitch returns a boolean if a field has been set.
func (o *HyperflexHxapHostVswitch) HasDvSwitch() bool {
	if o != nil && o.DvSwitch != nil {
		return true
	}

	return false
}

// SetDvSwitch gets a reference to the given HyperflexHxapDvswitchRelationship and assigns it to the DvSwitch field.
func (o *HyperflexHxapHostVswitch) SetDvSwitch(v HyperflexHxapDvswitchRelationship) {
	o.DvSwitch = &v
}

// GetHost returns the Host field value if set, zero value otherwise.
func (o *HyperflexHxapHostVswitch) GetHost() HyperflexHxapHostRelationship {
	if o == nil || o.Host == nil {
		var ret HyperflexHxapHostRelationship
		return ret
	}
	return *o.Host
}

// GetHostOk returns a tuple with the Host field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexHxapHostVswitch) GetHostOk() (*HyperflexHxapHostRelationship, bool) {
	if o == nil || o.Host == nil {
		return nil, false
	}
	return o.Host, true
}

// HasHost returns a boolean if a field has been set.
func (o *HyperflexHxapHostVswitch) HasHost() bool {
	if o != nil && o.Host != nil {
		return true
	}

	return false
}

// SetHost gets a reference to the given HyperflexHxapHostRelationship and assigns it to the Host field.
func (o *HyperflexHxapHostVswitch) SetHost(v HyperflexHxapHostRelationship) {
	o.Host = &v
}

func (o HyperflexHxapHostVswitch) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedVirtualizationBaseVswitch, errVirtualizationBaseVswitch := json.Marshal(o.VirtualizationBaseVswitch)
	if errVirtualizationBaseVswitch != nil {
		return []byte{}, errVirtualizationBaseVswitch
	}
	errVirtualizationBaseVswitch = json.Unmarshal([]byte(serializedVirtualizationBaseVswitch), &toSerialize)
	if errVirtualizationBaseVswitch != nil {
		return []byte{}, errVirtualizationBaseVswitch
	}
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

func (o *HyperflexHxapHostVswitch) UnmarshalJSON(bytes []byte) (err error) {
	type HyperflexHxapHostVswitchWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// The name of the host to which this vSwitch belongs to.
		HostName *string                            `json:"HostName,omitempty"`
		Ports    []HyperflexNetworkPort             `json:"Ports,omitempty"`
		Cluster  *HyperflexHxapClusterRelationship  `json:"Cluster,omitempty"`
		DvSwitch *HyperflexHxapDvswitchRelationship `json:"DvSwitch,omitempty"`
		Host     *HyperflexHxapHostRelationship     `json:"Host,omitempty"`
	}

	varHyperflexHxapHostVswitchWithoutEmbeddedStruct := HyperflexHxapHostVswitchWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varHyperflexHxapHostVswitchWithoutEmbeddedStruct)
	if err == nil {
		varHyperflexHxapHostVswitch := _HyperflexHxapHostVswitch{}
		varHyperflexHxapHostVswitch.ClassId = varHyperflexHxapHostVswitchWithoutEmbeddedStruct.ClassId
		varHyperflexHxapHostVswitch.ObjectType = varHyperflexHxapHostVswitchWithoutEmbeddedStruct.ObjectType
		varHyperflexHxapHostVswitch.HostName = varHyperflexHxapHostVswitchWithoutEmbeddedStruct.HostName
		varHyperflexHxapHostVswitch.Ports = varHyperflexHxapHostVswitchWithoutEmbeddedStruct.Ports
		varHyperflexHxapHostVswitch.Cluster = varHyperflexHxapHostVswitchWithoutEmbeddedStruct.Cluster
		varHyperflexHxapHostVswitch.DvSwitch = varHyperflexHxapHostVswitchWithoutEmbeddedStruct.DvSwitch
		varHyperflexHxapHostVswitch.Host = varHyperflexHxapHostVswitchWithoutEmbeddedStruct.Host
		*o = HyperflexHxapHostVswitch(varHyperflexHxapHostVswitch)
	} else {
		return err
	}

	varHyperflexHxapHostVswitch := _HyperflexHxapHostVswitch{}

	err = json.Unmarshal(bytes, &varHyperflexHxapHostVswitch)
	if err == nil {
		o.VirtualizationBaseVswitch = varHyperflexHxapHostVswitch.VirtualizationBaseVswitch
	} else {
		return err
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

		// remove fields from embedded structs
		reflectVirtualizationBaseVswitch := reflect.ValueOf(o.VirtualizationBaseVswitch)
		for i := 0; i < reflectVirtualizationBaseVswitch.Type().NumField(); i++ {
			t := reflectVirtualizationBaseVswitch.Type().Field(i)

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

type NullableHyperflexHxapHostVswitch struct {
	value *HyperflexHxapHostVswitch
	isSet bool
}

func (v NullableHyperflexHxapHostVswitch) Get() *HyperflexHxapHostVswitch {
	return v.value
}

func (v *NullableHyperflexHxapHostVswitch) Set(val *HyperflexHxapHostVswitch) {
	v.value = val
	v.isSet = true
}

func (v NullableHyperflexHxapHostVswitch) IsSet() bool {
	return v.isSet
}

func (v *NullableHyperflexHxapHostVswitch) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHyperflexHxapHostVswitch(val *HyperflexHxapHostVswitch) *NullableHyperflexHxapHostVswitch {
	return &NullableHyperflexHxapHostVswitch{value: val, isSet: true}
}

func (v NullableHyperflexHxapHostVswitch) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHyperflexHxapHostVswitch) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
