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

// PortInterfaceBase Abstract Interface Base class for a virtual interface card/Fex Network Port.
type PortInterfaceBase struct {
	InventoryBase
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property. The enum values provides the list of concrete types that can be instantiated from this abstract type.
	ObjectType string `json:"ObjectType"`
	// Operational state of an Interface.
	OperState                 *string                            `json:"OperState,omitempty"`
	AcknowledgedPeerInterface *EtherPhysicalPortBaseRelationship `json:"AcknowledgedPeerInterface,omitempty"`
	PeerInterface             *EtherPhysicalPortBaseRelationship `json:"PeerInterface,omitempty"`
	AdditionalProperties      map[string]interface{}
}

type _PortInterfaceBase PortInterfaceBase

// NewPortInterfaceBase instantiates a new PortInterfaceBase object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPortInterfaceBase(classId string, objectType string) *PortInterfaceBase {
	this := PortInterfaceBase{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewPortInterfaceBaseWithDefaults instantiates a new PortInterfaceBase object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPortInterfaceBaseWithDefaults() *PortInterfaceBase {
	this := PortInterfaceBase{}
	return &this
}

// GetClassId returns the ClassId field value
func (o *PortInterfaceBase) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *PortInterfaceBase) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *PortInterfaceBase) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *PortInterfaceBase) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *PortInterfaceBase) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *PortInterfaceBase) SetObjectType(v string) {
	o.ObjectType = v
}

// GetOperState returns the OperState field value if set, zero value otherwise.
func (o *PortInterfaceBase) GetOperState() string {
	if o == nil || o.OperState == nil {
		var ret string
		return ret
	}
	return *o.OperState
}

// GetOperStateOk returns a tuple with the OperState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PortInterfaceBase) GetOperStateOk() (*string, bool) {
	if o == nil || o.OperState == nil {
		return nil, false
	}
	return o.OperState, true
}

// HasOperState returns a boolean if a field has been set.
func (o *PortInterfaceBase) HasOperState() bool {
	if o != nil && o.OperState != nil {
		return true
	}

	return false
}

// SetOperState gets a reference to the given string and assigns it to the OperState field.
func (o *PortInterfaceBase) SetOperState(v string) {
	o.OperState = &v
}

// GetAcknowledgedPeerInterface returns the AcknowledgedPeerInterface field value if set, zero value otherwise.
func (o *PortInterfaceBase) GetAcknowledgedPeerInterface() EtherPhysicalPortBaseRelationship {
	if o == nil || o.AcknowledgedPeerInterface == nil {
		var ret EtherPhysicalPortBaseRelationship
		return ret
	}
	return *o.AcknowledgedPeerInterface
}

// GetAcknowledgedPeerInterfaceOk returns a tuple with the AcknowledgedPeerInterface field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PortInterfaceBase) GetAcknowledgedPeerInterfaceOk() (*EtherPhysicalPortBaseRelationship, bool) {
	if o == nil || o.AcknowledgedPeerInterface == nil {
		return nil, false
	}
	return o.AcknowledgedPeerInterface, true
}

// HasAcknowledgedPeerInterface returns a boolean if a field has been set.
func (o *PortInterfaceBase) HasAcknowledgedPeerInterface() bool {
	if o != nil && o.AcknowledgedPeerInterface != nil {
		return true
	}

	return false
}

// SetAcknowledgedPeerInterface gets a reference to the given EtherPhysicalPortBaseRelationship and assigns it to the AcknowledgedPeerInterface field.
func (o *PortInterfaceBase) SetAcknowledgedPeerInterface(v EtherPhysicalPortBaseRelationship) {
	o.AcknowledgedPeerInterface = &v
}

// GetPeerInterface returns the PeerInterface field value if set, zero value otherwise.
func (o *PortInterfaceBase) GetPeerInterface() EtherPhysicalPortBaseRelationship {
	if o == nil || o.PeerInterface == nil {
		var ret EtherPhysicalPortBaseRelationship
		return ret
	}
	return *o.PeerInterface
}

// GetPeerInterfaceOk returns a tuple with the PeerInterface field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PortInterfaceBase) GetPeerInterfaceOk() (*EtherPhysicalPortBaseRelationship, bool) {
	if o == nil || o.PeerInterface == nil {
		return nil, false
	}
	return o.PeerInterface, true
}

// HasPeerInterface returns a boolean if a field has been set.
func (o *PortInterfaceBase) HasPeerInterface() bool {
	if o != nil && o.PeerInterface != nil {
		return true
	}

	return false
}

// SetPeerInterface gets a reference to the given EtherPhysicalPortBaseRelationship and assigns it to the PeerInterface field.
func (o *PortInterfaceBase) SetPeerInterface(v EtherPhysicalPortBaseRelationship) {
	o.PeerInterface = &v
}

func (o PortInterfaceBase) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedInventoryBase, errInventoryBase := json.Marshal(o.InventoryBase)
	if errInventoryBase != nil {
		return []byte{}, errInventoryBase
	}
	errInventoryBase = json.Unmarshal([]byte(serializedInventoryBase), &toSerialize)
	if errInventoryBase != nil {
		return []byte{}, errInventoryBase
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.OperState != nil {
		toSerialize["OperState"] = o.OperState
	}
	if o.AcknowledgedPeerInterface != nil {
		toSerialize["AcknowledgedPeerInterface"] = o.AcknowledgedPeerInterface
	}
	if o.PeerInterface != nil {
		toSerialize["PeerInterface"] = o.PeerInterface
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *PortInterfaceBase) UnmarshalJSON(bytes []byte) (err error) {
	type PortInterfaceBaseWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property. The enum values provides the list of concrete types that can be instantiated from this abstract type.
		ObjectType string `json:"ObjectType"`
		// Operational state of an Interface.
		OperState                 *string                            `json:"OperState,omitempty"`
		AcknowledgedPeerInterface *EtherPhysicalPortBaseRelationship `json:"AcknowledgedPeerInterface,omitempty"`
		PeerInterface             *EtherPhysicalPortBaseRelationship `json:"PeerInterface,omitempty"`
	}

	varPortInterfaceBaseWithoutEmbeddedStruct := PortInterfaceBaseWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varPortInterfaceBaseWithoutEmbeddedStruct)
	if err == nil {
		varPortInterfaceBase := _PortInterfaceBase{}
		varPortInterfaceBase.ClassId = varPortInterfaceBaseWithoutEmbeddedStruct.ClassId
		varPortInterfaceBase.ObjectType = varPortInterfaceBaseWithoutEmbeddedStruct.ObjectType
		varPortInterfaceBase.OperState = varPortInterfaceBaseWithoutEmbeddedStruct.OperState
		varPortInterfaceBase.AcknowledgedPeerInterface = varPortInterfaceBaseWithoutEmbeddedStruct.AcknowledgedPeerInterface
		varPortInterfaceBase.PeerInterface = varPortInterfaceBaseWithoutEmbeddedStruct.PeerInterface
		*o = PortInterfaceBase(varPortInterfaceBase)
	} else {
		return err
	}

	varPortInterfaceBase := _PortInterfaceBase{}

	err = json.Unmarshal(bytes, &varPortInterfaceBase)
	if err == nil {
		o.InventoryBase = varPortInterfaceBase.InventoryBase
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "OperState")
		delete(additionalProperties, "AcknowledgedPeerInterface")
		delete(additionalProperties, "PeerInterface")

		// remove fields from embedded structs
		reflectInventoryBase := reflect.ValueOf(o.InventoryBase)
		for i := 0; i < reflectInventoryBase.Type().NumField(); i++ {
			t := reflectInventoryBase.Type().Field(i)

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

type NullablePortInterfaceBase struct {
	value *PortInterfaceBase
	isSet bool
}

func (v NullablePortInterfaceBase) Get() *PortInterfaceBase {
	return v.value
}

func (v *NullablePortInterfaceBase) Set(val *PortInterfaceBase) {
	v.value = val
	v.isSet = true
}

func (v NullablePortInterfaceBase) IsSet() bool {
	return v.isSet
}

func (v *NullablePortInterfaceBase) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePortInterfaceBase(val *PortInterfaceBase) *NullablePortInterfaceBase {
	return &NullablePortInterfaceBase{value: val, isSet: true}
}

func (v NullablePortInterfaceBase) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePortInterfaceBase) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
