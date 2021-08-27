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

// CapabilitySwitchStorageLimits Object combines and lists the storage-config limits for each of the Fabric/Switch platforms.
type CapabilitySwitchStorageLimits struct {
	MoBaseComplexType
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Maximum user zones per Switch/Fabric-Interconnect.
	MaximumUserZoneCount *int64 `json:"MaximumUserZoneCount,omitempty"`
	// Maximum configurable Virtual Fibre Channel interfaces on Switch/Fabric-Interconnect.
	MaximumVirtualFcInterfaces *int64 `json:"MaximumVirtualFcInterfaces,omitempty"`
	// Maximum configurable Virtual Fibre Channel interfaces per blade.
	MaximumVirtualFcInterfacesPerBladeServer *int64 `json:"MaximumVirtualFcInterfacesPerBladeServer,omitempty"`
	// Maximum configurable VSANs on Switch/Fabric-Interconnect.
	MaximumVsans *int64 `json:"MaximumVsans,omitempty"`
	// Zone limit per Switch/Fabric-Interconnect.
	MaximumZoneCount     *int64 `json:"MaximumZoneCount,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CapabilitySwitchStorageLimits CapabilitySwitchStorageLimits

// NewCapabilitySwitchStorageLimits instantiates a new CapabilitySwitchStorageLimits object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCapabilitySwitchStorageLimits(classId string, objectType string) *CapabilitySwitchStorageLimits {
	this := CapabilitySwitchStorageLimits{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewCapabilitySwitchStorageLimitsWithDefaults instantiates a new CapabilitySwitchStorageLimits object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCapabilitySwitchStorageLimitsWithDefaults() *CapabilitySwitchStorageLimits {
	this := CapabilitySwitchStorageLimits{}
	var classId string = "capability.SwitchStorageLimits"
	this.ClassId = classId
	var objectType string = "capability.SwitchStorageLimits"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *CapabilitySwitchStorageLimits) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *CapabilitySwitchStorageLimits) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *CapabilitySwitchStorageLimits) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *CapabilitySwitchStorageLimits) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *CapabilitySwitchStorageLimits) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *CapabilitySwitchStorageLimits) SetObjectType(v string) {
	o.ObjectType = v
}

// GetMaximumUserZoneCount returns the MaximumUserZoneCount field value if set, zero value otherwise.
func (o *CapabilitySwitchStorageLimits) GetMaximumUserZoneCount() int64 {
	if o == nil || o.MaximumUserZoneCount == nil {
		var ret int64
		return ret
	}
	return *o.MaximumUserZoneCount
}

// GetMaximumUserZoneCountOk returns a tuple with the MaximumUserZoneCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CapabilitySwitchStorageLimits) GetMaximumUserZoneCountOk() (*int64, bool) {
	if o == nil || o.MaximumUserZoneCount == nil {
		return nil, false
	}
	return o.MaximumUserZoneCount, true
}

// HasMaximumUserZoneCount returns a boolean if a field has been set.
func (o *CapabilitySwitchStorageLimits) HasMaximumUserZoneCount() bool {
	if o != nil && o.MaximumUserZoneCount != nil {
		return true
	}

	return false
}

// SetMaximumUserZoneCount gets a reference to the given int64 and assigns it to the MaximumUserZoneCount field.
func (o *CapabilitySwitchStorageLimits) SetMaximumUserZoneCount(v int64) {
	o.MaximumUserZoneCount = &v
}

// GetMaximumVirtualFcInterfaces returns the MaximumVirtualFcInterfaces field value if set, zero value otherwise.
func (o *CapabilitySwitchStorageLimits) GetMaximumVirtualFcInterfaces() int64 {
	if o == nil || o.MaximumVirtualFcInterfaces == nil {
		var ret int64
		return ret
	}
	return *o.MaximumVirtualFcInterfaces
}

// GetMaximumVirtualFcInterfacesOk returns a tuple with the MaximumVirtualFcInterfaces field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CapabilitySwitchStorageLimits) GetMaximumVirtualFcInterfacesOk() (*int64, bool) {
	if o == nil || o.MaximumVirtualFcInterfaces == nil {
		return nil, false
	}
	return o.MaximumVirtualFcInterfaces, true
}

// HasMaximumVirtualFcInterfaces returns a boolean if a field has been set.
func (o *CapabilitySwitchStorageLimits) HasMaximumVirtualFcInterfaces() bool {
	if o != nil && o.MaximumVirtualFcInterfaces != nil {
		return true
	}

	return false
}

// SetMaximumVirtualFcInterfaces gets a reference to the given int64 and assigns it to the MaximumVirtualFcInterfaces field.
func (o *CapabilitySwitchStorageLimits) SetMaximumVirtualFcInterfaces(v int64) {
	o.MaximumVirtualFcInterfaces = &v
}

// GetMaximumVirtualFcInterfacesPerBladeServer returns the MaximumVirtualFcInterfacesPerBladeServer field value if set, zero value otherwise.
func (o *CapabilitySwitchStorageLimits) GetMaximumVirtualFcInterfacesPerBladeServer() int64 {
	if o == nil || o.MaximumVirtualFcInterfacesPerBladeServer == nil {
		var ret int64
		return ret
	}
	return *o.MaximumVirtualFcInterfacesPerBladeServer
}

// GetMaximumVirtualFcInterfacesPerBladeServerOk returns a tuple with the MaximumVirtualFcInterfacesPerBladeServer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CapabilitySwitchStorageLimits) GetMaximumVirtualFcInterfacesPerBladeServerOk() (*int64, bool) {
	if o == nil || o.MaximumVirtualFcInterfacesPerBladeServer == nil {
		return nil, false
	}
	return o.MaximumVirtualFcInterfacesPerBladeServer, true
}

// HasMaximumVirtualFcInterfacesPerBladeServer returns a boolean if a field has been set.
func (o *CapabilitySwitchStorageLimits) HasMaximumVirtualFcInterfacesPerBladeServer() bool {
	if o != nil && o.MaximumVirtualFcInterfacesPerBladeServer != nil {
		return true
	}

	return false
}

// SetMaximumVirtualFcInterfacesPerBladeServer gets a reference to the given int64 and assigns it to the MaximumVirtualFcInterfacesPerBladeServer field.
func (o *CapabilitySwitchStorageLimits) SetMaximumVirtualFcInterfacesPerBladeServer(v int64) {
	o.MaximumVirtualFcInterfacesPerBladeServer = &v
}

// GetMaximumVsans returns the MaximumVsans field value if set, zero value otherwise.
func (o *CapabilitySwitchStorageLimits) GetMaximumVsans() int64 {
	if o == nil || o.MaximumVsans == nil {
		var ret int64
		return ret
	}
	return *o.MaximumVsans
}

// GetMaximumVsansOk returns a tuple with the MaximumVsans field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CapabilitySwitchStorageLimits) GetMaximumVsansOk() (*int64, bool) {
	if o == nil || o.MaximumVsans == nil {
		return nil, false
	}
	return o.MaximumVsans, true
}

// HasMaximumVsans returns a boolean if a field has been set.
func (o *CapabilitySwitchStorageLimits) HasMaximumVsans() bool {
	if o != nil && o.MaximumVsans != nil {
		return true
	}

	return false
}

// SetMaximumVsans gets a reference to the given int64 and assigns it to the MaximumVsans field.
func (o *CapabilitySwitchStorageLimits) SetMaximumVsans(v int64) {
	o.MaximumVsans = &v
}

// GetMaximumZoneCount returns the MaximumZoneCount field value if set, zero value otherwise.
func (o *CapabilitySwitchStorageLimits) GetMaximumZoneCount() int64 {
	if o == nil || o.MaximumZoneCount == nil {
		var ret int64
		return ret
	}
	return *o.MaximumZoneCount
}

// GetMaximumZoneCountOk returns a tuple with the MaximumZoneCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CapabilitySwitchStorageLimits) GetMaximumZoneCountOk() (*int64, bool) {
	if o == nil || o.MaximumZoneCount == nil {
		return nil, false
	}
	return o.MaximumZoneCount, true
}

// HasMaximumZoneCount returns a boolean if a field has been set.
func (o *CapabilitySwitchStorageLimits) HasMaximumZoneCount() bool {
	if o != nil && o.MaximumZoneCount != nil {
		return true
	}

	return false
}

// SetMaximumZoneCount gets a reference to the given int64 and assigns it to the MaximumZoneCount field.
func (o *CapabilitySwitchStorageLimits) SetMaximumZoneCount(v int64) {
	o.MaximumZoneCount = &v
}

func (o CapabilitySwitchStorageLimits) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedMoBaseComplexType, errMoBaseComplexType := json.Marshal(o.MoBaseComplexType)
	if errMoBaseComplexType != nil {
		return []byte{}, errMoBaseComplexType
	}
	errMoBaseComplexType = json.Unmarshal([]byte(serializedMoBaseComplexType), &toSerialize)
	if errMoBaseComplexType != nil {
		return []byte{}, errMoBaseComplexType
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.MaximumUserZoneCount != nil {
		toSerialize["MaximumUserZoneCount"] = o.MaximumUserZoneCount
	}
	if o.MaximumVirtualFcInterfaces != nil {
		toSerialize["MaximumVirtualFcInterfaces"] = o.MaximumVirtualFcInterfaces
	}
	if o.MaximumVirtualFcInterfacesPerBladeServer != nil {
		toSerialize["MaximumVirtualFcInterfacesPerBladeServer"] = o.MaximumVirtualFcInterfacesPerBladeServer
	}
	if o.MaximumVsans != nil {
		toSerialize["MaximumVsans"] = o.MaximumVsans
	}
	if o.MaximumZoneCount != nil {
		toSerialize["MaximumZoneCount"] = o.MaximumZoneCount
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *CapabilitySwitchStorageLimits) UnmarshalJSON(bytes []byte) (err error) {
	type CapabilitySwitchStorageLimitsWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Maximum user zones per Switch/Fabric-Interconnect.
		MaximumUserZoneCount *int64 `json:"MaximumUserZoneCount,omitempty"`
		// Maximum configurable Virtual Fibre Channel interfaces on Switch/Fabric-Interconnect.
		MaximumVirtualFcInterfaces *int64 `json:"MaximumVirtualFcInterfaces,omitempty"`
		// Maximum configurable Virtual Fibre Channel interfaces per blade.
		MaximumVirtualFcInterfacesPerBladeServer *int64 `json:"MaximumVirtualFcInterfacesPerBladeServer,omitempty"`
		// Maximum configurable VSANs on Switch/Fabric-Interconnect.
		MaximumVsans *int64 `json:"MaximumVsans,omitempty"`
		// Zone limit per Switch/Fabric-Interconnect.
		MaximumZoneCount *int64 `json:"MaximumZoneCount,omitempty"`
	}

	varCapabilitySwitchStorageLimitsWithoutEmbeddedStruct := CapabilitySwitchStorageLimitsWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varCapabilitySwitchStorageLimitsWithoutEmbeddedStruct)
	if err == nil {
		varCapabilitySwitchStorageLimits := _CapabilitySwitchStorageLimits{}
		varCapabilitySwitchStorageLimits.ClassId = varCapabilitySwitchStorageLimitsWithoutEmbeddedStruct.ClassId
		varCapabilitySwitchStorageLimits.ObjectType = varCapabilitySwitchStorageLimitsWithoutEmbeddedStruct.ObjectType
		varCapabilitySwitchStorageLimits.MaximumUserZoneCount = varCapabilitySwitchStorageLimitsWithoutEmbeddedStruct.MaximumUserZoneCount
		varCapabilitySwitchStorageLimits.MaximumVirtualFcInterfaces = varCapabilitySwitchStorageLimitsWithoutEmbeddedStruct.MaximumVirtualFcInterfaces
		varCapabilitySwitchStorageLimits.MaximumVirtualFcInterfacesPerBladeServer = varCapabilitySwitchStorageLimitsWithoutEmbeddedStruct.MaximumVirtualFcInterfacesPerBladeServer
		varCapabilitySwitchStorageLimits.MaximumVsans = varCapabilitySwitchStorageLimitsWithoutEmbeddedStruct.MaximumVsans
		varCapabilitySwitchStorageLimits.MaximumZoneCount = varCapabilitySwitchStorageLimitsWithoutEmbeddedStruct.MaximumZoneCount
		*o = CapabilitySwitchStorageLimits(varCapabilitySwitchStorageLimits)
	} else {
		return err
	}

	varCapabilitySwitchStorageLimits := _CapabilitySwitchStorageLimits{}

	err = json.Unmarshal(bytes, &varCapabilitySwitchStorageLimits)
	if err == nil {
		o.MoBaseComplexType = varCapabilitySwitchStorageLimits.MoBaseComplexType
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "MaximumUserZoneCount")
		delete(additionalProperties, "MaximumVirtualFcInterfaces")
		delete(additionalProperties, "MaximumVirtualFcInterfacesPerBladeServer")
		delete(additionalProperties, "MaximumVsans")
		delete(additionalProperties, "MaximumZoneCount")

		// remove fields from embedded structs
		reflectMoBaseComplexType := reflect.ValueOf(o.MoBaseComplexType)
		for i := 0; i < reflectMoBaseComplexType.Type().NumField(); i++ {
			t := reflectMoBaseComplexType.Type().Field(i)

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

type NullableCapabilitySwitchStorageLimits struct {
	value *CapabilitySwitchStorageLimits
	isSet bool
}

func (v NullableCapabilitySwitchStorageLimits) Get() *CapabilitySwitchStorageLimits {
	return v.value
}

func (v *NullableCapabilitySwitchStorageLimits) Set(val *CapabilitySwitchStorageLimits) {
	v.value = val
	v.isSet = true
}

func (v NullableCapabilitySwitchStorageLimits) IsSet() bool {
	return v.isSet
}

func (v *NullableCapabilitySwitchStorageLimits) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCapabilitySwitchStorageLimits(val *CapabilitySwitchStorageLimits) *NullableCapabilitySwitchStorageLimits {
	return &NullableCapabilitySwitchStorageLimits{value: val, isSet: true}
}

func (v NullableCapabilitySwitchStorageLimits) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCapabilitySwitchStorageLimits) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
