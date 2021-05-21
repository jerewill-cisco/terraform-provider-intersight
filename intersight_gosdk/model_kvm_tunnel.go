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

// KvmTunnel Tunneled Virtual KVM access to the vKVM console of a server. This must be specified while creating the vKVM session to gain tunneled access.
type KvmTunnel struct {
	TunnelingTunnel
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType           string                               `json:"ObjectType"`
	Var0Session          *KvmSessionRelationship              `json:"_0_Session,omitempty"`
	Device               *AssetDeviceRegistrationRelationship `json:"Device,omitempty"`
	Server               *ComputePhysicalRelationship         `json:"Server,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _KvmTunnel KvmTunnel

// NewKvmTunnel instantiates a new KvmTunnel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKvmTunnel(classId string, objectType string) *KvmTunnel {
	this := KvmTunnel{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewKvmTunnelWithDefaults instantiates a new KvmTunnel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKvmTunnelWithDefaults() *KvmTunnel {
	this := KvmTunnel{}
	var classId string = "kvm.Tunnel"
	this.ClassId = classId
	var objectType string = "kvm.Tunnel"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *KvmTunnel) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *KvmTunnel) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *KvmTunnel) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *KvmTunnel) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *KvmTunnel) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *KvmTunnel) SetObjectType(v string) {
	o.ObjectType = v
}

// GetVar0Session returns the Var0Session field value if set, zero value otherwise.
func (o *KvmTunnel) GetVar0Session() KvmSessionRelationship {
	if o == nil || o.Var0Session == nil {
		var ret KvmSessionRelationship
		return ret
	}
	return *o.Var0Session
}

// GetVar0SessionOk returns a tuple with the Var0Session field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KvmTunnel) GetVar0SessionOk() (*KvmSessionRelationship, bool) {
	if o == nil || o.Var0Session == nil {
		return nil, false
	}
	return o.Var0Session, true
}

// HasVar0Session returns a boolean if a field has been set.
func (o *KvmTunnel) HasVar0Session() bool {
	if o != nil && o.Var0Session != nil {
		return true
	}

	return false
}

// SetVar0Session gets a reference to the given KvmSessionRelationship and assigns it to the Var0Session field.
func (o *KvmTunnel) SetVar0Session(v KvmSessionRelationship) {
	o.Var0Session = &v
}

// GetDevice returns the Device field value if set, zero value otherwise.
func (o *KvmTunnel) GetDevice() AssetDeviceRegistrationRelationship {
	if o == nil || o.Device == nil {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.Device
}

// GetDeviceOk returns a tuple with the Device field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KvmTunnel) GetDeviceOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil || o.Device == nil {
		return nil, false
	}
	return o.Device, true
}

// HasDevice returns a boolean if a field has been set.
func (o *KvmTunnel) HasDevice() bool {
	if o != nil && o.Device != nil {
		return true
	}

	return false
}

// SetDevice gets a reference to the given AssetDeviceRegistrationRelationship and assigns it to the Device field.
func (o *KvmTunnel) SetDevice(v AssetDeviceRegistrationRelationship) {
	o.Device = &v
}

// GetServer returns the Server field value if set, zero value otherwise.
func (o *KvmTunnel) GetServer() ComputePhysicalRelationship {
	if o == nil || o.Server == nil {
		var ret ComputePhysicalRelationship
		return ret
	}
	return *o.Server
}

// GetServerOk returns a tuple with the Server field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KvmTunnel) GetServerOk() (*ComputePhysicalRelationship, bool) {
	if o == nil || o.Server == nil {
		return nil, false
	}
	return o.Server, true
}

// HasServer returns a boolean if a field has been set.
func (o *KvmTunnel) HasServer() bool {
	if o != nil && o.Server != nil {
		return true
	}

	return false
}

// SetServer gets a reference to the given ComputePhysicalRelationship and assigns it to the Server field.
func (o *KvmTunnel) SetServer(v ComputePhysicalRelationship) {
	o.Server = &v
}

func (o KvmTunnel) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedTunnelingTunnel, errTunnelingTunnel := json.Marshal(o.TunnelingTunnel)
	if errTunnelingTunnel != nil {
		return []byte{}, errTunnelingTunnel
	}
	errTunnelingTunnel = json.Unmarshal([]byte(serializedTunnelingTunnel), &toSerialize)
	if errTunnelingTunnel != nil {
		return []byte{}, errTunnelingTunnel
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Var0Session != nil {
		toSerialize["_0_Session"] = o.Var0Session
	}
	if o.Device != nil {
		toSerialize["Device"] = o.Device
	}
	if o.Server != nil {
		toSerialize["Server"] = o.Server
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *KvmTunnel) UnmarshalJSON(bytes []byte) (err error) {
	type KvmTunnelWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType  string                               `json:"ObjectType"`
		Var0Session *KvmSessionRelationship              `json:"_0_Session,omitempty"`
		Device      *AssetDeviceRegistrationRelationship `json:"Device,omitempty"`
		Server      *ComputePhysicalRelationship         `json:"Server,omitempty"`
	}

	varKvmTunnelWithoutEmbeddedStruct := KvmTunnelWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varKvmTunnelWithoutEmbeddedStruct)
	if err == nil {
		varKvmTunnel := _KvmTunnel{}
		varKvmTunnel.ClassId = varKvmTunnelWithoutEmbeddedStruct.ClassId
		varKvmTunnel.ObjectType = varKvmTunnelWithoutEmbeddedStruct.ObjectType
		varKvmTunnel.Var0Session = varKvmTunnelWithoutEmbeddedStruct.Var0Session
		varKvmTunnel.Device = varKvmTunnelWithoutEmbeddedStruct.Device
		varKvmTunnel.Server = varKvmTunnelWithoutEmbeddedStruct.Server
		*o = KvmTunnel(varKvmTunnel)
	} else {
		return err
	}

	varKvmTunnel := _KvmTunnel{}

	err = json.Unmarshal(bytes, &varKvmTunnel)
	if err == nil {
		o.TunnelingTunnel = varKvmTunnel.TunnelingTunnel
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "_0_Session")
		delete(additionalProperties, "Device")
		delete(additionalProperties, "Server")

		// remove fields from embedded structs
		reflectTunnelingTunnel := reflect.ValueOf(o.TunnelingTunnel)
		for i := 0; i < reflectTunnelingTunnel.Type().NumField(); i++ {
			t := reflectTunnelingTunnel.Type().Field(i)

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

type NullableKvmTunnel struct {
	value *KvmTunnel
	isSet bool
}

func (v NullableKvmTunnel) Get() *KvmTunnel {
	return v.value
}

func (v *NullableKvmTunnel) Set(val *KvmTunnel) {
	v.value = val
	v.isSet = true
}

func (v NullableKvmTunnel) IsSet() bool {
	return v.isSet
}

func (v *NullableKvmTunnel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKvmTunnel(val *KvmTunnel) *NullableKvmTunnel {
	return &NullableKvmTunnel{value: val, isSet: true}
}

func (v NullableKvmTunnel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKvmTunnel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
