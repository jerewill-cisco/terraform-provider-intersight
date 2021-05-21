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

// FirmwareSwitchUpgrade Firmware upgrade operation for Fabric Interconnects that downloads the image located at Cisco/appliance/user provided HTTP repository or uses the image from a network share and upgrade. Direct download is used for upgrade that uses the image from a Cisco repository or an appliance repository. Network share is used for upgrade that use the image from a network share from your data center.
type FirmwareSwitchUpgrade struct {
	FirmwareUpgradeBase
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// The flag to enable or disable fabric evacuation during the switch firmware upgrade. In case of IMM, it is mandatory to have the Fabric Interconnects associated with domain profile for fabric evacuation to happen.
	EnableFabricEvacuation *bool                                `json:"EnableFabricEvacuation,omitempty"`
	Device                 *AssetDeviceRegistrationRelationship `json:"Device,omitempty"`
	// An array of relationships to networkElement resources.
	NetworkElements      []NetworkElementRelationship `json:"NetworkElements,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _FirmwareSwitchUpgrade FirmwareSwitchUpgrade

// NewFirmwareSwitchUpgrade instantiates a new FirmwareSwitchUpgrade object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFirmwareSwitchUpgrade(classId string, objectType string) *FirmwareSwitchUpgrade {
	this := FirmwareSwitchUpgrade{}
	this.ClassId = classId
	this.ObjectType = objectType
	var enableFabricEvacuation bool = true
	this.EnableFabricEvacuation = &enableFabricEvacuation
	return &this
}

// NewFirmwareSwitchUpgradeWithDefaults instantiates a new FirmwareSwitchUpgrade object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFirmwareSwitchUpgradeWithDefaults() *FirmwareSwitchUpgrade {
	this := FirmwareSwitchUpgrade{}
	var classId string = "firmware.SwitchUpgrade"
	this.ClassId = classId
	var objectType string = "firmware.SwitchUpgrade"
	this.ObjectType = objectType
	var enableFabricEvacuation bool = true
	this.EnableFabricEvacuation = &enableFabricEvacuation
	return &this
}

// GetClassId returns the ClassId field value
func (o *FirmwareSwitchUpgrade) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *FirmwareSwitchUpgrade) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *FirmwareSwitchUpgrade) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *FirmwareSwitchUpgrade) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *FirmwareSwitchUpgrade) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *FirmwareSwitchUpgrade) SetObjectType(v string) {
	o.ObjectType = v
}

// GetEnableFabricEvacuation returns the EnableFabricEvacuation field value if set, zero value otherwise.
func (o *FirmwareSwitchUpgrade) GetEnableFabricEvacuation() bool {
	if o == nil || o.EnableFabricEvacuation == nil {
		var ret bool
		return ret
	}
	return *o.EnableFabricEvacuation
}

// GetEnableFabricEvacuationOk returns a tuple with the EnableFabricEvacuation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FirmwareSwitchUpgrade) GetEnableFabricEvacuationOk() (*bool, bool) {
	if o == nil || o.EnableFabricEvacuation == nil {
		return nil, false
	}
	return o.EnableFabricEvacuation, true
}

// HasEnableFabricEvacuation returns a boolean if a field has been set.
func (o *FirmwareSwitchUpgrade) HasEnableFabricEvacuation() bool {
	if o != nil && o.EnableFabricEvacuation != nil {
		return true
	}

	return false
}

// SetEnableFabricEvacuation gets a reference to the given bool and assigns it to the EnableFabricEvacuation field.
func (o *FirmwareSwitchUpgrade) SetEnableFabricEvacuation(v bool) {
	o.EnableFabricEvacuation = &v
}

// GetDevice returns the Device field value if set, zero value otherwise.
func (o *FirmwareSwitchUpgrade) GetDevice() AssetDeviceRegistrationRelationship {
	if o == nil || o.Device == nil {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.Device
}

// GetDeviceOk returns a tuple with the Device field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FirmwareSwitchUpgrade) GetDeviceOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil || o.Device == nil {
		return nil, false
	}
	return o.Device, true
}

// HasDevice returns a boolean if a field has been set.
func (o *FirmwareSwitchUpgrade) HasDevice() bool {
	if o != nil && o.Device != nil {
		return true
	}

	return false
}

// SetDevice gets a reference to the given AssetDeviceRegistrationRelationship and assigns it to the Device field.
func (o *FirmwareSwitchUpgrade) SetDevice(v AssetDeviceRegistrationRelationship) {
	o.Device = &v
}

// GetNetworkElements returns the NetworkElements field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FirmwareSwitchUpgrade) GetNetworkElements() []NetworkElementRelationship {
	if o == nil {
		var ret []NetworkElementRelationship
		return ret
	}
	return o.NetworkElements
}

// GetNetworkElementsOk returns a tuple with the NetworkElements field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FirmwareSwitchUpgrade) GetNetworkElementsOk() (*[]NetworkElementRelationship, bool) {
	if o == nil || o.NetworkElements == nil {
		return nil, false
	}
	return &o.NetworkElements, true
}

// HasNetworkElements returns a boolean if a field has been set.
func (o *FirmwareSwitchUpgrade) HasNetworkElements() bool {
	if o != nil && o.NetworkElements != nil {
		return true
	}

	return false
}

// SetNetworkElements gets a reference to the given []NetworkElementRelationship and assigns it to the NetworkElements field.
func (o *FirmwareSwitchUpgrade) SetNetworkElements(v []NetworkElementRelationship) {
	o.NetworkElements = v
}

func (o FirmwareSwitchUpgrade) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedFirmwareUpgradeBase, errFirmwareUpgradeBase := json.Marshal(o.FirmwareUpgradeBase)
	if errFirmwareUpgradeBase != nil {
		return []byte{}, errFirmwareUpgradeBase
	}
	errFirmwareUpgradeBase = json.Unmarshal([]byte(serializedFirmwareUpgradeBase), &toSerialize)
	if errFirmwareUpgradeBase != nil {
		return []byte{}, errFirmwareUpgradeBase
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.EnableFabricEvacuation != nil {
		toSerialize["EnableFabricEvacuation"] = o.EnableFabricEvacuation
	}
	if o.Device != nil {
		toSerialize["Device"] = o.Device
	}
	if o.NetworkElements != nil {
		toSerialize["NetworkElements"] = o.NetworkElements
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *FirmwareSwitchUpgrade) UnmarshalJSON(bytes []byte) (err error) {
	type FirmwareSwitchUpgradeWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// The flag to enable or disable fabric evacuation during the switch firmware upgrade. In case of IMM, it is mandatory to have the Fabric Interconnects associated with domain profile for fabric evacuation to happen.
		EnableFabricEvacuation *bool                                `json:"EnableFabricEvacuation,omitempty"`
		Device                 *AssetDeviceRegistrationRelationship `json:"Device,omitempty"`
		// An array of relationships to networkElement resources.
		NetworkElements []NetworkElementRelationship `json:"NetworkElements,omitempty"`
	}

	varFirmwareSwitchUpgradeWithoutEmbeddedStruct := FirmwareSwitchUpgradeWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varFirmwareSwitchUpgradeWithoutEmbeddedStruct)
	if err == nil {
		varFirmwareSwitchUpgrade := _FirmwareSwitchUpgrade{}
		varFirmwareSwitchUpgrade.ClassId = varFirmwareSwitchUpgradeWithoutEmbeddedStruct.ClassId
		varFirmwareSwitchUpgrade.ObjectType = varFirmwareSwitchUpgradeWithoutEmbeddedStruct.ObjectType
		varFirmwareSwitchUpgrade.EnableFabricEvacuation = varFirmwareSwitchUpgradeWithoutEmbeddedStruct.EnableFabricEvacuation
		varFirmwareSwitchUpgrade.Device = varFirmwareSwitchUpgradeWithoutEmbeddedStruct.Device
		varFirmwareSwitchUpgrade.NetworkElements = varFirmwareSwitchUpgradeWithoutEmbeddedStruct.NetworkElements
		*o = FirmwareSwitchUpgrade(varFirmwareSwitchUpgrade)
	} else {
		return err
	}

	varFirmwareSwitchUpgrade := _FirmwareSwitchUpgrade{}

	err = json.Unmarshal(bytes, &varFirmwareSwitchUpgrade)
	if err == nil {
		o.FirmwareUpgradeBase = varFirmwareSwitchUpgrade.FirmwareUpgradeBase
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "EnableFabricEvacuation")
		delete(additionalProperties, "Device")
		delete(additionalProperties, "NetworkElements")

		// remove fields from embedded structs
		reflectFirmwareUpgradeBase := reflect.ValueOf(o.FirmwareUpgradeBase)
		for i := 0; i < reflectFirmwareUpgradeBase.Type().NumField(); i++ {
			t := reflectFirmwareUpgradeBase.Type().Field(i)

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

type NullableFirmwareSwitchUpgrade struct {
	value *FirmwareSwitchUpgrade
	isSet bool
}

func (v NullableFirmwareSwitchUpgrade) Get() *FirmwareSwitchUpgrade {
	return v.value
}

func (v *NullableFirmwareSwitchUpgrade) Set(val *FirmwareSwitchUpgrade) {
	v.value = val
	v.isSet = true
}

func (v NullableFirmwareSwitchUpgrade) IsSet() bool {
	return v.isSet
}

func (v *NullableFirmwareSwitchUpgrade) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFirmwareSwitchUpgrade(val *FirmwareSwitchUpgrade) *NullableFirmwareSwitchUpgrade {
	return &NullableFirmwareSwitchUpgrade{value: val, isSet: true}
}

func (v NullableFirmwareSwitchUpgrade) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFirmwareSwitchUpgrade) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
