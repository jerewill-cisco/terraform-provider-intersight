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
	"reflect"
	"strings"
)

// EquipmentIoCard I/O module on a chassis which multiplexes traffic from blade servers.
type EquipmentIoCard struct {
	EquipmentIoCardBase
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Switch Id to which the IOM is connected to. The value can be A or B.
	ConnectionPath *string `json:"ConnectionPath,omitempty"`
	// IOM device connector support.
	DcSupported       *bool              `json:"DcSupported,omitempty"`
	InbandIpAddresses []ComputeIpAddress `json:"InbandIpAddresses,omitempty"`
	// Location of IOM within a chassis. The value can be left or right.
	Side                       *string                              `json:"Side,omitempty"`
	EquipmentChassis           *EquipmentChassisRelationship        `json:"EquipmentChassis,omitempty"`
	EquipmentFex               *EquipmentFexRelationship            `json:"EquipmentFex,omitempty"`
	InventoryDeviceInfo        *InventoryDeviceInfoRelationship     `json:"InventoryDeviceInfo,omitempty"`
	PhysicalDeviceRegistration *AssetDeviceRegistrationRelationship `json:"PhysicalDeviceRegistration,omitempty"`
	RegisteredDevice           *AssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	AdditionalProperties       map[string]interface{}
}

type _EquipmentIoCard EquipmentIoCard

// NewEquipmentIoCard instantiates a new EquipmentIoCard object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEquipmentIoCard(classId string, objectType string) *EquipmentIoCard {
	this := EquipmentIoCard{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewEquipmentIoCardWithDefaults instantiates a new EquipmentIoCard object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEquipmentIoCardWithDefaults() *EquipmentIoCard {
	this := EquipmentIoCard{}
	var classId string = "equipment.IoCard"
	this.ClassId = classId
	var objectType string = "equipment.IoCard"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *EquipmentIoCard) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *EquipmentIoCard) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *EquipmentIoCard) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *EquipmentIoCard) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *EquipmentIoCard) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *EquipmentIoCard) SetObjectType(v string) {
	o.ObjectType = v
}

// GetConnectionPath returns the ConnectionPath field value if set, zero value otherwise.
func (o *EquipmentIoCard) GetConnectionPath() string {
	if o == nil || o.ConnectionPath == nil {
		var ret string
		return ret
	}
	return *o.ConnectionPath
}

// GetConnectionPathOk returns a tuple with the ConnectionPath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentIoCard) GetConnectionPathOk() (*string, bool) {
	if o == nil || o.ConnectionPath == nil {
		return nil, false
	}
	return o.ConnectionPath, true
}

// HasConnectionPath returns a boolean if a field has been set.
func (o *EquipmentIoCard) HasConnectionPath() bool {
	if o != nil && o.ConnectionPath != nil {
		return true
	}

	return false
}

// SetConnectionPath gets a reference to the given string and assigns it to the ConnectionPath field.
func (o *EquipmentIoCard) SetConnectionPath(v string) {
	o.ConnectionPath = &v
}

// GetDcSupported returns the DcSupported field value if set, zero value otherwise.
func (o *EquipmentIoCard) GetDcSupported() bool {
	if o == nil || o.DcSupported == nil {
		var ret bool
		return ret
	}
	return *o.DcSupported
}

// GetDcSupportedOk returns a tuple with the DcSupported field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentIoCard) GetDcSupportedOk() (*bool, bool) {
	if o == nil || o.DcSupported == nil {
		return nil, false
	}
	return o.DcSupported, true
}

// HasDcSupported returns a boolean if a field has been set.
func (o *EquipmentIoCard) HasDcSupported() bool {
	if o != nil && o.DcSupported != nil {
		return true
	}

	return false
}

// SetDcSupported gets a reference to the given bool and assigns it to the DcSupported field.
func (o *EquipmentIoCard) SetDcSupported(v bool) {
	o.DcSupported = &v
}

// GetInbandIpAddresses returns the InbandIpAddresses field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EquipmentIoCard) GetInbandIpAddresses() []ComputeIpAddress {
	if o == nil {
		var ret []ComputeIpAddress
		return ret
	}
	return o.InbandIpAddresses
}

// GetInbandIpAddressesOk returns a tuple with the InbandIpAddresses field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EquipmentIoCard) GetInbandIpAddressesOk() (*[]ComputeIpAddress, bool) {
	if o == nil || o.InbandIpAddresses == nil {
		return nil, false
	}
	return &o.InbandIpAddresses, true
}

// HasInbandIpAddresses returns a boolean if a field has been set.
func (o *EquipmentIoCard) HasInbandIpAddresses() bool {
	if o != nil && o.InbandIpAddresses != nil {
		return true
	}

	return false
}

// SetInbandIpAddresses gets a reference to the given []ComputeIpAddress and assigns it to the InbandIpAddresses field.
func (o *EquipmentIoCard) SetInbandIpAddresses(v []ComputeIpAddress) {
	o.InbandIpAddresses = v
}

// GetSide returns the Side field value if set, zero value otherwise.
func (o *EquipmentIoCard) GetSide() string {
	if o == nil || o.Side == nil {
		var ret string
		return ret
	}
	return *o.Side
}

// GetSideOk returns a tuple with the Side field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentIoCard) GetSideOk() (*string, bool) {
	if o == nil || o.Side == nil {
		return nil, false
	}
	return o.Side, true
}

// HasSide returns a boolean if a field has been set.
func (o *EquipmentIoCard) HasSide() bool {
	if o != nil && o.Side != nil {
		return true
	}

	return false
}

// SetSide gets a reference to the given string and assigns it to the Side field.
func (o *EquipmentIoCard) SetSide(v string) {
	o.Side = &v
}

// GetEquipmentChassis returns the EquipmentChassis field value if set, zero value otherwise.
func (o *EquipmentIoCard) GetEquipmentChassis() EquipmentChassisRelationship {
	if o == nil || o.EquipmentChassis == nil {
		var ret EquipmentChassisRelationship
		return ret
	}
	return *o.EquipmentChassis
}

// GetEquipmentChassisOk returns a tuple with the EquipmentChassis field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentIoCard) GetEquipmentChassisOk() (*EquipmentChassisRelationship, bool) {
	if o == nil || o.EquipmentChassis == nil {
		return nil, false
	}
	return o.EquipmentChassis, true
}

// HasEquipmentChassis returns a boolean if a field has been set.
func (o *EquipmentIoCard) HasEquipmentChassis() bool {
	if o != nil && o.EquipmentChassis != nil {
		return true
	}

	return false
}

// SetEquipmentChassis gets a reference to the given EquipmentChassisRelationship and assigns it to the EquipmentChassis field.
func (o *EquipmentIoCard) SetEquipmentChassis(v EquipmentChassisRelationship) {
	o.EquipmentChassis = &v
}

// GetEquipmentFex returns the EquipmentFex field value if set, zero value otherwise.
func (o *EquipmentIoCard) GetEquipmentFex() EquipmentFexRelationship {
	if o == nil || o.EquipmentFex == nil {
		var ret EquipmentFexRelationship
		return ret
	}
	return *o.EquipmentFex
}

// GetEquipmentFexOk returns a tuple with the EquipmentFex field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentIoCard) GetEquipmentFexOk() (*EquipmentFexRelationship, bool) {
	if o == nil || o.EquipmentFex == nil {
		return nil, false
	}
	return o.EquipmentFex, true
}

// HasEquipmentFex returns a boolean if a field has been set.
func (o *EquipmentIoCard) HasEquipmentFex() bool {
	if o != nil && o.EquipmentFex != nil {
		return true
	}

	return false
}

// SetEquipmentFex gets a reference to the given EquipmentFexRelationship and assigns it to the EquipmentFex field.
func (o *EquipmentIoCard) SetEquipmentFex(v EquipmentFexRelationship) {
	o.EquipmentFex = &v
}

// GetInventoryDeviceInfo returns the InventoryDeviceInfo field value if set, zero value otherwise.
func (o *EquipmentIoCard) GetInventoryDeviceInfo() InventoryDeviceInfoRelationship {
	if o == nil || o.InventoryDeviceInfo == nil {
		var ret InventoryDeviceInfoRelationship
		return ret
	}
	return *o.InventoryDeviceInfo
}

// GetInventoryDeviceInfoOk returns a tuple with the InventoryDeviceInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentIoCard) GetInventoryDeviceInfoOk() (*InventoryDeviceInfoRelationship, bool) {
	if o == nil || o.InventoryDeviceInfo == nil {
		return nil, false
	}
	return o.InventoryDeviceInfo, true
}

// HasInventoryDeviceInfo returns a boolean if a field has been set.
func (o *EquipmentIoCard) HasInventoryDeviceInfo() bool {
	if o != nil && o.InventoryDeviceInfo != nil {
		return true
	}

	return false
}

// SetInventoryDeviceInfo gets a reference to the given InventoryDeviceInfoRelationship and assigns it to the InventoryDeviceInfo field.
func (o *EquipmentIoCard) SetInventoryDeviceInfo(v InventoryDeviceInfoRelationship) {
	o.InventoryDeviceInfo = &v
}

// GetPhysicalDeviceRegistration returns the PhysicalDeviceRegistration field value if set, zero value otherwise.
func (o *EquipmentIoCard) GetPhysicalDeviceRegistration() AssetDeviceRegistrationRelationship {
	if o == nil || o.PhysicalDeviceRegistration == nil {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.PhysicalDeviceRegistration
}

// GetPhysicalDeviceRegistrationOk returns a tuple with the PhysicalDeviceRegistration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentIoCard) GetPhysicalDeviceRegistrationOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil || o.PhysicalDeviceRegistration == nil {
		return nil, false
	}
	return o.PhysicalDeviceRegistration, true
}

// HasPhysicalDeviceRegistration returns a boolean if a field has been set.
func (o *EquipmentIoCard) HasPhysicalDeviceRegistration() bool {
	if o != nil && o.PhysicalDeviceRegistration != nil {
		return true
	}

	return false
}

// SetPhysicalDeviceRegistration gets a reference to the given AssetDeviceRegistrationRelationship and assigns it to the PhysicalDeviceRegistration field.
func (o *EquipmentIoCard) SetPhysicalDeviceRegistration(v AssetDeviceRegistrationRelationship) {
	o.PhysicalDeviceRegistration = &v
}

// GetRegisteredDevice returns the RegisteredDevice field value if set, zero value otherwise.
func (o *EquipmentIoCard) GetRegisteredDevice() AssetDeviceRegistrationRelationship {
	if o == nil || o.RegisteredDevice == nil {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.RegisteredDevice
}

// GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentIoCard) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil || o.RegisteredDevice == nil {
		return nil, false
	}
	return o.RegisteredDevice, true
}

// HasRegisteredDevice returns a boolean if a field has been set.
func (o *EquipmentIoCard) HasRegisteredDevice() bool {
	if o != nil && o.RegisteredDevice != nil {
		return true
	}

	return false
}

// SetRegisteredDevice gets a reference to the given AssetDeviceRegistrationRelationship and assigns it to the RegisteredDevice field.
func (o *EquipmentIoCard) SetRegisteredDevice(v AssetDeviceRegistrationRelationship) {
	o.RegisteredDevice = &v
}

func (o EquipmentIoCard) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedEquipmentIoCardBase, errEquipmentIoCardBase := json.Marshal(o.EquipmentIoCardBase)
	if errEquipmentIoCardBase != nil {
		return []byte{}, errEquipmentIoCardBase
	}
	errEquipmentIoCardBase = json.Unmarshal([]byte(serializedEquipmentIoCardBase), &toSerialize)
	if errEquipmentIoCardBase != nil {
		return []byte{}, errEquipmentIoCardBase
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.ConnectionPath != nil {
		toSerialize["ConnectionPath"] = o.ConnectionPath
	}
	if o.DcSupported != nil {
		toSerialize["DcSupported"] = o.DcSupported
	}
	if o.InbandIpAddresses != nil {
		toSerialize["InbandIpAddresses"] = o.InbandIpAddresses
	}
	if o.Side != nil {
		toSerialize["Side"] = o.Side
	}
	if o.EquipmentChassis != nil {
		toSerialize["EquipmentChassis"] = o.EquipmentChassis
	}
	if o.EquipmentFex != nil {
		toSerialize["EquipmentFex"] = o.EquipmentFex
	}
	if o.InventoryDeviceInfo != nil {
		toSerialize["InventoryDeviceInfo"] = o.InventoryDeviceInfo
	}
	if o.PhysicalDeviceRegistration != nil {
		toSerialize["PhysicalDeviceRegistration"] = o.PhysicalDeviceRegistration
	}
	if o.RegisteredDevice != nil {
		toSerialize["RegisteredDevice"] = o.RegisteredDevice
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *EquipmentIoCard) UnmarshalJSON(bytes []byte) (err error) {
	type EquipmentIoCardWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Switch Id to which the IOM is connected to. The value can be A or B.
		ConnectionPath *string `json:"ConnectionPath,omitempty"`
		// IOM device connector support.
		DcSupported       *bool              `json:"DcSupported,omitempty"`
		InbandIpAddresses []ComputeIpAddress `json:"InbandIpAddresses,omitempty"`
		// Location of IOM within a chassis. The value can be left or right.
		Side                       *string                              `json:"Side,omitempty"`
		EquipmentChassis           *EquipmentChassisRelationship        `json:"EquipmentChassis,omitempty"`
		EquipmentFex               *EquipmentFexRelationship            `json:"EquipmentFex,omitempty"`
		InventoryDeviceInfo        *InventoryDeviceInfoRelationship     `json:"InventoryDeviceInfo,omitempty"`
		PhysicalDeviceRegistration *AssetDeviceRegistrationRelationship `json:"PhysicalDeviceRegistration,omitempty"`
		RegisteredDevice           *AssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	}

	varEquipmentIoCardWithoutEmbeddedStruct := EquipmentIoCardWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varEquipmentIoCardWithoutEmbeddedStruct)
	if err == nil {
		varEquipmentIoCard := _EquipmentIoCard{}
		varEquipmentIoCard.ClassId = varEquipmentIoCardWithoutEmbeddedStruct.ClassId
		varEquipmentIoCard.ObjectType = varEquipmentIoCardWithoutEmbeddedStruct.ObjectType
		varEquipmentIoCard.ConnectionPath = varEquipmentIoCardWithoutEmbeddedStruct.ConnectionPath
		varEquipmentIoCard.DcSupported = varEquipmentIoCardWithoutEmbeddedStruct.DcSupported
		varEquipmentIoCard.InbandIpAddresses = varEquipmentIoCardWithoutEmbeddedStruct.InbandIpAddresses
		varEquipmentIoCard.Side = varEquipmentIoCardWithoutEmbeddedStruct.Side
		varEquipmentIoCard.EquipmentChassis = varEquipmentIoCardWithoutEmbeddedStruct.EquipmentChassis
		varEquipmentIoCard.EquipmentFex = varEquipmentIoCardWithoutEmbeddedStruct.EquipmentFex
		varEquipmentIoCard.InventoryDeviceInfo = varEquipmentIoCardWithoutEmbeddedStruct.InventoryDeviceInfo
		varEquipmentIoCard.PhysicalDeviceRegistration = varEquipmentIoCardWithoutEmbeddedStruct.PhysicalDeviceRegistration
		varEquipmentIoCard.RegisteredDevice = varEquipmentIoCardWithoutEmbeddedStruct.RegisteredDevice
		*o = EquipmentIoCard(varEquipmentIoCard)
	} else {
		return err
	}

	varEquipmentIoCard := _EquipmentIoCard{}

	err = json.Unmarshal(bytes, &varEquipmentIoCard)
	if err == nil {
		o.EquipmentIoCardBase = varEquipmentIoCard.EquipmentIoCardBase
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "ConnectionPath")
		delete(additionalProperties, "DcSupported")
		delete(additionalProperties, "InbandIpAddresses")
		delete(additionalProperties, "Side")
		delete(additionalProperties, "EquipmentChassis")
		delete(additionalProperties, "EquipmentFex")
		delete(additionalProperties, "InventoryDeviceInfo")
		delete(additionalProperties, "PhysicalDeviceRegistration")
		delete(additionalProperties, "RegisteredDevice")

		// remove fields from embedded structs
		reflectEquipmentIoCardBase := reflect.ValueOf(o.EquipmentIoCardBase)
		for i := 0; i < reflectEquipmentIoCardBase.Type().NumField(); i++ {
			t := reflectEquipmentIoCardBase.Type().Field(i)

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

type NullableEquipmentIoCard struct {
	value *EquipmentIoCard
	isSet bool
}

func (v NullableEquipmentIoCard) Get() *EquipmentIoCard {
	return v.value
}

func (v *NullableEquipmentIoCard) Set(val *EquipmentIoCard) {
	v.value = val
	v.isSet = true
}

func (v NullableEquipmentIoCard) IsSet() bool {
	return v.isSet
}

func (v *NullableEquipmentIoCard) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEquipmentIoCard(val *EquipmentIoCard) *NullableEquipmentIoCard {
	return &NullableEquipmentIoCard{value: val, isSet: true}
}

func (v NullableEquipmentIoCard) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEquipmentIoCard) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
