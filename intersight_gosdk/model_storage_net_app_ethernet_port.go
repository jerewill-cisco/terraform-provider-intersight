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

// StorageNetAppEthernetPort Ethernet port is a port on a node in a storage array.
type StorageNetAppEthernetPort struct {
	MoBaseMo
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Status of Port to determine if its enabled or not.
	Enabled *string `json:"Enabled,omitempty"`
	// Macaddress  of the port available in storage array.
	MacAddress *string `json:"MacAddress,omitempty"`
	// Maximum transmission unit of the physical port available in storage array.
	Mtu *string `json:"Mtu,omitempty"`
	// Name of the port available in storage array.
	Name *string `json:"Name,omitempty"`
	// Operational speed of port measured.
	Speed *int64 `json:"Speed,omitempty"`
	// State of the port available in storage array. * `down` - An inactive port is listed as Down. * `up` - An active port is listed as Up. * `present` - An active port is listed as present.
	State *string `json:"State,omitempty"`
	// Type of the port available in storage array. * `LAG` - Storage port of type lag. * `physical` - LIFs can be configured directly on physical ports. * `VLAN` - A logical port that receives and sends VLAN-tagged (IEEE 802.1Q standard) traffic. VLAN port characteristics include the VLAN ID for the port.
	Type *string `json:"Type,omitempty"`
	// UUID of physical port.
	Uuid                 *string                        `json:"Uuid,omitempty"`
	ArrayController      *StorageNetAppNodeRelationship `json:"ArrayController,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _StorageNetAppEthernetPort StorageNetAppEthernetPort

// NewStorageNetAppEthernetPort instantiates a new StorageNetAppEthernetPort object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStorageNetAppEthernetPort(classId string, objectType string) *StorageNetAppEthernetPort {
	this := StorageNetAppEthernetPort{}
	this.ClassId = classId
	this.ObjectType = objectType
	var state string = "down"
	this.State = &state
	var type_ string = "LAG"
	this.Type = &type_
	return &this
}

// NewStorageNetAppEthernetPortWithDefaults instantiates a new StorageNetAppEthernetPort object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStorageNetAppEthernetPortWithDefaults() *StorageNetAppEthernetPort {
	this := StorageNetAppEthernetPort{}
	var classId string = "storage.NetAppEthernetPort"
	this.ClassId = classId
	var objectType string = "storage.NetAppEthernetPort"
	this.ObjectType = objectType
	var state string = "down"
	this.State = &state
	var type_ string = "LAG"
	this.Type = &type_
	return &this
}

// GetClassId returns the ClassId field value
func (o *StorageNetAppEthernetPort) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *StorageNetAppEthernetPort) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *StorageNetAppEthernetPort) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *StorageNetAppEthernetPort) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *StorageNetAppEthernetPort) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *StorageNetAppEthernetPort) SetObjectType(v string) {
	o.ObjectType = v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *StorageNetAppEthernetPort) GetEnabled() string {
	if o == nil || o.Enabled == nil {
		var ret string
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNetAppEthernetPort) GetEnabledOk() (*string, bool) {
	if o == nil || o.Enabled == nil {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *StorageNetAppEthernetPort) HasEnabled() bool {
	if o != nil && o.Enabled != nil {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given string and assigns it to the Enabled field.
func (o *StorageNetAppEthernetPort) SetEnabled(v string) {
	o.Enabled = &v
}

// GetMacAddress returns the MacAddress field value if set, zero value otherwise.
func (o *StorageNetAppEthernetPort) GetMacAddress() string {
	if o == nil || o.MacAddress == nil {
		var ret string
		return ret
	}
	return *o.MacAddress
}

// GetMacAddressOk returns a tuple with the MacAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNetAppEthernetPort) GetMacAddressOk() (*string, bool) {
	if o == nil || o.MacAddress == nil {
		return nil, false
	}
	return o.MacAddress, true
}

// HasMacAddress returns a boolean if a field has been set.
func (o *StorageNetAppEthernetPort) HasMacAddress() bool {
	if o != nil && o.MacAddress != nil {
		return true
	}

	return false
}

// SetMacAddress gets a reference to the given string and assigns it to the MacAddress field.
func (o *StorageNetAppEthernetPort) SetMacAddress(v string) {
	o.MacAddress = &v
}

// GetMtu returns the Mtu field value if set, zero value otherwise.
func (o *StorageNetAppEthernetPort) GetMtu() string {
	if o == nil || o.Mtu == nil {
		var ret string
		return ret
	}
	return *o.Mtu
}

// GetMtuOk returns a tuple with the Mtu field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNetAppEthernetPort) GetMtuOk() (*string, bool) {
	if o == nil || o.Mtu == nil {
		return nil, false
	}
	return o.Mtu, true
}

// HasMtu returns a boolean if a field has been set.
func (o *StorageNetAppEthernetPort) HasMtu() bool {
	if o != nil && o.Mtu != nil {
		return true
	}

	return false
}

// SetMtu gets a reference to the given string and assigns it to the Mtu field.
func (o *StorageNetAppEthernetPort) SetMtu(v string) {
	o.Mtu = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *StorageNetAppEthernetPort) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNetAppEthernetPort) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *StorageNetAppEthernetPort) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *StorageNetAppEthernetPort) SetName(v string) {
	o.Name = &v
}

// GetSpeed returns the Speed field value if set, zero value otherwise.
func (o *StorageNetAppEthernetPort) GetSpeed() int64 {
	if o == nil || o.Speed == nil {
		var ret int64
		return ret
	}
	return *o.Speed
}

// GetSpeedOk returns a tuple with the Speed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNetAppEthernetPort) GetSpeedOk() (*int64, bool) {
	if o == nil || o.Speed == nil {
		return nil, false
	}
	return o.Speed, true
}

// HasSpeed returns a boolean if a field has been set.
func (o *StorageNetAppEthernetPort) HasSpeed() bool {
	if o != nil && o.Speed != nil {
		return true
	}

	return false
}

// SetSpeed gets a reference to the given int64 and assigns it to the Speed field.
func (o *StorageNetAppEthernetPort) SetSpeed(v int64) {
	o.Speed = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *StorageNetAppEthernetPort) GetState() string {
	if o == nil || o.State == nil {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNetAppEthernetPort) GetStateOk() (*string, bool) {
	if o == nil || o.State == nil {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *StorageNetAppEthernetPort) HasState() bool {
	if o != nil && o.State != nil {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *StorageNetAppEthernetPort) SetState(v string) {
	o.State = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *StorageNetAppEthernetPort) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNetAppEthernetPort) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *StorageNetAppEthernetPort) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *StorageNetAppEthernetPort) SetType(v string) {
	o.Type = &v
}

// GetUuid returns the Uuid field value if set, zero value otherwise.
func (o *StorageNetAppEthernetPort) GetUuid() string {
	if o == nil || o.Uuid == nil {
		var ret string
		return ret
	}
	return *o.Uuid
}

// GetUuidOk returns a tuple with the Uuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNetAppEthernetPort) GetUuidOk() (*string, bool) {
	if o == nil || o.Uuid == nil {
		return nil, false
	}
	return o.Uuid, true
}

// HasUuid returns a boolean if a field has been set.
func (o *StorageNetAppEthernetPort) HasUuid() bool {
	if o != nil && o.Uuid != nil {
		return true
	}

	return false
}

// SetUuid gets a reference to the given string and assigns it to the Uuid field.
func (o *StorageNetAppEthernetPort) SetUuid(v string) {
	o.Uuid = &v
}

// GetArrayController returns the ArrayController field value if set, zero value otherwise.
func (o *StorageNetAppEthernetPort) GetArrayController() StorageNetAppNodeRelationship {
	if o == nil || o.ArrayController == nil {
		var ret StorageNetAppNodeRelationship
		return ret
	}
	return *o.ArrayController
}

// GetArrayControllerOk returns a tuple with the ArrayController field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNetAppEthernetPort) GetArrayControllerOk() (*StorageNetAppNodeRelationship, bool) {
	if o == nil || o.ArrayController == nil {
		return nil, false
	}
	return o.ArrayController, true
}

// HasArrayController returns a boolean if a field has been set.
func (o *StorageNetAppEthernetPort) HasArrayController() bool {
	if o != nil && o.ArrayController != nil {
		return true
	}

	return false
}

// SetArrayController gets a reference to the given StorageNetAppNodeRelationship and assigns it to the ArrayController field.
func (o *StorageNetAppEthernetPort) SetArrayController(v StorageNetAppNodeRelationship) {
	o.ArrayController = &v
}

func (o StorageNetAppEthernetPort) MarshalJSON() ([]byte, error) {
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
	if o.Enabled != nil {
		toSerialize["Enabled"] = o.Enabled
	}
	if o.MacAddress != nil {
		toSerialize["MacAddress"] = o.MacAddress
	}
	if o.Mtu != nil {
		toSerialize["Mtu"] = o.Mtu
	}
	if o.Name != nil {
		toSerialize["Name"] = o.Name
	}
	if o.Speed != nil {
		toSerialize["Speed"] = o.Speed
	}
	if o.State != nil {
		toSerialize["State"] = o.State
	}
	if o.Type != nil {
		toSerialize["Type"] = o.Type
	}
	if o.Uuid != nil {
		toSerialize["Uuid"] = o.Uuid
	}
	if o.ArrayController != nil {
		toSerialize["ArrayController"] = o.ArrayController
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *StorageNetAppEthernetPort) UnmarshalJSON(bytes []byte) (err error) {
	type StorageNetAppEthernetPortWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Status of Port to determine if its enabled or not.
		Enabled *string `json:"Enabled,omitempty"`
		// Macaddress  of the port available in storage array.
		MacAddress *string `json:"MacAddress,omitempty"`
		// Maximum transmission unit of the physical port available in storage array.
		Mtu *string `json:"Mtu,omitempty"`
		// Name of the port available in storage array.
		Name *string `json:"Name,omitempty"`
		// Operational speed of port measured.
		Speed *int64 `json:"Speed,omitempty"`
		// State of the port available in storage array. * `down` - An inactive port is listed as Down. * `up` - An active port is listed as Up. * `present` - An active port is listed as present.
		State *string `json:"State,omitempty"`
		// Type of the port available in storage array. * `LAG` - Storage port of type lag. * `physical` - LIFs can be configured directly on physical ports. * `VLAN` - A logical port that receives and sends VLAN-tagged (IEEE 802.1Q standard) traffic. VLAN port characteristics include the VLAN ID for the port.
		Type *string `json:"Type,omitempty"`
		// UUID of physical port.
		Uuid            *string                        `json:"Uuid,omitempty"`
		ArrayController *StorageNetAppNodeRelationship `json:"ArrayController,omitempty"`
	}

	varStorageNetAppEthernetPortWithoutEmbeddedStruct := StorageNetAppEthernetPortWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varStorageNetAppEthernetPortWithoutEmbeddedStruct)
	if err == nil {
		varStorageNetAppEthernetPort := _StorageNetAppEthernetPort{}
		varStorageNetAppEthernetPort.ClassId = varStorageNetAppEthernetPortWithoutEmbeddedStruct.ClassId
		varStorageNetAppEthernetPort.ObjectType = varStorageNetAppEthernetPortWithoutEmbeddedStruct.ObjectType
		varStorageNetAppEthernetPort.Enabled = varStorageNetAppEthernetPortWithoutEmbeddedStruct.Enabled
		varStorageNetAppEthernetPort.MacAddress = varStorageNetAppEthernetPortWithoutEmbeddedStruct.MacAddress
		varStorageNetAppEthernetPort.Mtu = varStorageNetAppEthernetPortWithoutEmbeddedStruct.Mtu
		varStorageNetAppEthernetPort.Name = varStorageNetAppEthernetPortWithoutEmbeddedStruct.Name
		varStorageNetAppEthernetPort.Speed = varStorageNetAppEthernetPortWithoutEmbeddedStruct.Speed
		varStorageNetAppEthernetPort.State = varStorageNetAppEthernetPortWithoutEmbeddedStruct.State
		varStorageNetAppEthernetPort.Type = varStorageNetAppEthernetPortWithoutEmbeddedStruct.Type
		varStorageNetAppEthernetPort.Uuid = varStorageNetAppEthernetPortWithoutEmbeddedStruct.Uuid
		varStorageNetAppEthernetPort.ArrayController = varStorageNetAppEthernetPortWithoutEmbeddedStruct.ArrayController
		*o = StorageNetAppEthernetPort(varStorageNetAppEthernetPort)
	} else {
		return err
	}

	varStorageNetAppEthernetPort := _StorageNetAppEthernetPort{}

	err = json.Unmarshal(bytes, &varStorageNetAppEthernetPort)
	if err == nil {
		o.MoBaseMo = varStorageNetAppEthernetPort.MoBaseMo
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Enabled")
		delete(additionalProperties, "MacAddress")
		delete(additionalProperties, "Mtu")
		delete(additionalProperties, "Name")
		delete(additionalProperties, "Speed")
		delete(additionalProperties, "State")
		delete(additionalProperties, "Type")
		delete(additionalProperties, "Uuid")
		delete(additionalProperties, "ArrayController")

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

type NullableStorageNetAppEthernetPort struct {
	value *StorageNetAppEthernetPort
	isSet bool
}

func (v NullableStorageNetAppEthernetPort) Get() *StorageNetAppEthernetPort {
	return v.value
}

func (v *NullableStorageNetAppEthernetPort) Set(val *StorageNetAppEthernetPort) {
	v.value = val
	v.isSet = true
}

func (v NullableStorageNetAppEthernetPort) IsSet() bool {
	return v.isSet
}

func (v *NullableStorageNetAppEthernetPort) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStorageNetAppEthernetPort(val *StorageNetAppEthernetPort) *NullableStorageNetAppEthernetPort {
	return &NullableStorageNetAppEthernetPort{value: val, isSet: true}
}

func (v NullableStorageNetAppEthernetPort) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStorageNetAppEthernetPort) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
