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

// NiatelemetryApicSnmpTrapDetails Object to capture Snmp trap details in APIC.
type NiatelemetryApicSnmpTrapDetails struct {
	MoBaseMo
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Dn of SNMP Trap attribute in APIC.
	Dn *string `json:"Dn,omitempty"`
	// Type of record DCNM / APIC / SE. This determines the type of platform where inventory was collected.
	RecordType *string `json:"RecordType,omitempty"`
	// Version of record being pushed. This determines what was the API version for data available from the device.
	RecordVersion *string `json:"RecordVersion,omitempty"`
	// Name of the APIC site from which this data is being collected.
	SiteName *string `json:"SiteName,omitempty"`
	// V3SecLevel of SNMP Trap in APIC.
	V3SecLevel *string `json:"V3SecLevel,omitempty"`
	// Version of SNMP Trap in APIC.
	Ver                  *string                              `json:"Ver,omitempty"`
	RegisteredDevice     *AssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _NiatelemetryApicSnmpTrapDetails NiatelemetryApicSnmpTrapDetails

// NewNiatelemetryApicSnmpTrapDetails instantiates a new NiatelemetryApicSnmpTrapDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNiatelemetryApicSnmpTrapDetails(classId string, objectType string) *NiatelemetryApicSnmpTrapDetails {
	this := NiatelemetryApicSnmpTrapDetails{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewNiatelemetryApicSnmpTrapDetailsWithDefaults instantiates a new NiatelemetryApicSnmpTrapDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNiatelemetryApicSnmpTrapDetailsWithDefaults() *NiatelemetryApicSnmpTrapDetails {
	this := NiatelemetryApicSnmpTrapDetails{}
	var classId string = "niatelemetry.ApicSnmpTrapDetails"
	this.ClassId = classId
	var objectType string = "niatelemetry.ApicSnmpTrapDetails"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *NiatelemetryApicSnmpTrapDetails) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *NiatelemetryApicSnmpTrapDetails) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *NiatelemetryApicSnmpTrapDetails) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *NiatelemetryApicSnmpTrapDetails) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *NiatelemetryApicSnmpTrapDetails) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *NiatelemetryApicSnmpTrapDetails) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDn returns the Dn field value if set, zero value otherwise.
func (o *NiatelemetryApicSnmpTrapDetails) GetDn() string {
	if o == nil || o.Dn == nil {
		var ret string
		return ret
	}
	return *o.Dn
}

// GetDnOk returns a tuple with the Dn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryApicSnmpTrapDetails) GetDnOk() (*string, bool) {
	if o == nil || o.Dn == nil {
		return nil, false
	}
	return o.Dn, true
}

// HasDn returns a boolean if a field has been set.
func (o *NiatelemetryApicSnmpTrapDetails) HasDn() bool {
	if o != nil && o.Dn != nil {
		return true
	}

	return false
}

// SetDn gets a reference to the given string and assigns it to the Dn field.
func (o *NiatelemetryApicSnmpTrapDetails) SetDn(v string) {
	o.Dn = &v
}

// GetRecordType returns the RecordType field value if set, zero value otherwise.
func (o *NiatelemetryApicSnmpTrapDetails) GetRecordType() string {
	if o == nil || o.RecordType == nil {
		var ret string
		return ret
	}
	return *o.RecordType
}

// GetRecordTypeOk returns a tuple with the RecordType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryApicSnmpTrapDetails) GetRecordTypeOk() (*string, bool) {
	if o == nil || o.RecordType == nil {
		return nil, false
	}
	return o.RecordType, true
}

// HasRecordType returns a boolean if a field has been set.
func (o *NiatelemetryApicSnmpTrapDetails) HasRecordType() bool {
	if o != nil && o.RecordType != nil {
		return true
	}

	return false
}

// SetRecordType gets a reference to the given string and assigns it to the RecordType field.
func (o *NiatelemetryApicSnmpTrapDetails) SetRecordType(v string) {
	o.RecordType = &v
}

// GetRecordVersion returns the RecordVersion field value if set, zero value otherwise.
func (o *NiatelemetryApicSnmpTrapDetails) GetRecordVersion() string {
	if o == nil || o.RecordVersion == nil {
		var ret string
		return ret
	}
	return *o.RecordVersion
}

// GetRecordVersionOk returns a tuple with the RecordVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryApicSnmpTrapDetails) GetRecordVersionOk() (*string, bool) {
	if o == nil || o.RecordVersion == nil {
		return nil, false
	}
	return o.RecordVersion, true
}

// HasRecordVersion returns a boolean if a field has been set.
func (o *NiatelemetryApicSnmpTrapDetails) HasRecordVersion() bool {
	if o != nil && o.RecordVersion != nil {
		return true
	}

	return false
}

// SetRecordVersion gets a reference to the given string and assigns it to the RecordVersion field.
func (o *NiatelemetryApicSnmpTrapDetails) SetRecordVersion(v string) {
	o.RecordVersion = &v
}

// GetSiteName returns the SiteName field value if set, zero value otherwise.
func (o *NiatelemetryApicSnmpTrapDetails) GetSiteName() string {
	if o == nil || o.SiteName == nil {
		var ret string
		return ret
	}
	return *o.SiteName
}

// GetSiteNameOk returns a tuple with the SiteName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryApicSnmpTrapDetails) GetSiteNameOk() (*string, bool) {
	if o == nil || o.SiteName == nil {
		return nil, false
	}
	return o.SiteName, true
}

// HasSiteName returns a boolean if a field has been set.
func (o *NiatelemetryApicSnmpTrapDetails) HasSiteName() bool {
	if o != nil && o.SiteName != nil {
		return true
	}

	return false
}

// SetSiteName gets a reference to the given string and assigns it to the SiteName field.
func (o *NiatelemetryApicSnmpTrapDetails) SetSiteName(v string) {
	o.SiteName = &v
}

// GetV3SecLevel returns the V3SecLevel field value if set, zero value otherwise.
func (o *NiatelemetryApicSnmpTrapDetails) GetV3SecLevel() string {
	if o == nil || o.V3SecLevel == nil {
		var ret string
		return ret
	}
	return *o.V3SecLevel
}

// GetV3SecLevelOk returns a tuple with the V3SecLevel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryApicSnmpTrapDetails) GetV3SecLevelOk() (*string, bool) {
	if o == nil || o.V3SecLevel == nil {
		return nil, false
	}
	return o.V3SecLevel, true
}

// HasV3SecLevel returns a boolean if a field has been set.
func (o *NiatelemetryApicSnmpTrapDetails) HasV3SecLevel() bool {
	if o != nil && o.V3SecLevel != nil {
		return true
	}

	return false
}

// SetV3SecLevel gets a reference to the given string and assigns it to the V3SecLevel field.
func (o *NiatelemetryApicSnmpTrapDetails) SetV3SecLevel(v string) {
	o.V3SecLevel = &v
}

// GetVer returns the Ver field value if set, zero value otherwise.
func (o *NiatelemetryApicSnmpTrapDetails) GetVer() string {
	if o == nil || o.Ver == nil {
		var ret string
		return ret
	}
	return *o.Ver
}

// GetVerOk returns a tuple with the Ver field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryApicSnmpTrapDetails) GetVerOk() (*string, bool) {
	if o == nil || o.Ver == nil {
		return nil, false
	}
	return o.Ver, true
}

// HasVer returns a boolean if a field has been set.
func (o *NiatelemetryApicSnmpTrapDetails) HasVer() bool {
	if o != nil && o.Ver != nil {
		return true
	}

	return false
}

// SetVer gets a reference to the given string and assigns it to the Ver field.
func (o *NiatelemetryApicSnmpTrapDetails) SetVer(v string) {
	o.Ver = &v
}

// GetRegisteredDevice returns the RegisteredDevice field value if set, zero value otherwise.
func (o *NiatelemetryApicSnmpTrapDetails) GetRegisteredDevice() AssetDeviceRegistrationRelationship {
	if o == nil || o.RegisteredDevice == nil {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.RegisteredDevice
}

// GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryApicSnmpTrapDetails) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil || o.RegisteredDevice == nil {
		return nil, false
	}
	return o.RegisteredDevice, true
}

// HasRegisteredDevice returns a boolean if a field has been set.
func (o *NiatelemetryApicSnmpTrapDetails) HasRegisteredDevice() bool {
	if o != nil && o.RegisteredDevice != nil {
		return true
	}

	return false
}

// SetRegisteredDevice gets a reference to the given AssetDeviceRegistrationRelationship and assigns it to the RegisteredDevice field.
func (o *NiatelemetryApicSnmpTrapDetails) SetRegisteredDevice(v AssetDeviceRegistrationRelationship) {
	o.RegisteredDevice = &v
}

func (o NiatelemetryApicSnmpTrapDetails) MarshalJSON() ([]byte, error) {
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
	if o.Dn != nil {
		toSerialize["Dn"] = o.Dn
	}
	if o.RecordType != nil {
		toSerialize["RecordType"] = o.RecordType
	}
	if o.RecordVersion != nil {
		toSerialize["RecordVersion"] = o.RecordVersion
	}
	if o.SiteName != nil {
		toSerialize["SiteName"] = o.SiteName
	}
	if o.V3SecLevel != nil {
		toSerialize["V3SecLevel"] = o.V3SecLevel
	}
	if o.Ver != nil {
		toSerialize["Ver"] = o.Ver
	}
	if o.RegisteredDevice != nil {
		toSerialize["RegisteredDevice"] = o.RegisteredDevice
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *NiatelemetryApicSnmpTrapDetails) UnmarshalJSON(bytes []byte) (err error) {
	type NiatelemetryApicSnmpTrapDetailsWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Dn of SNMP Trap attribute in APIC.
		Dn *string `json:"Dn,omitempty"`
		// Type of record DCNM / APIC / SE. This determines the type of platform where inventory was collected.
		RecordType *string `json:"RecordType,omitempty"`
		// Version of record being pushed. This determines what was the API version for data available from the device.
		RecordVersion *string `json:"RecordVersion,omitempty"`
		// Name of the APIC site from which this data is being collected.
		SiteName *string `json:"SiteName,omitempty"`
		// V3SecLevel of SNMP Trap in APIC.
		V3SecLevel *string `json:"V3SecLevel,omitempty"`
		// Version of SNMP Trap in APIC.
		Ver              *string                              `json:"Ver,omitempty"`
		RegisteredDevice *AssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	}

	varNiatelemetryApicSnmpTrapDetailsWithoutEmbeddedStruct := NiatelemetryApicSnmpTrapDetailsWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varNiatelemetryApicSnmpTrapDetailsWithoutEmbeddedStruct)
	if err == nil {
		varNiatelemetryApicSnmpTrapDetails := _NiatelemetryApicSnmpTrapDetails{}
		varNiatelemetryApicSnmpTrapDetails.ClassId = varNiatelemetryApicSnmpTrapDetailsWithoutEmbeddedStruct.ClassId
		varNiatelemetryApicSnmpTrapDetails.ObjectType = varNiatelemetryApicSnmpTrapDetailsWithoutEmbeddedStruct.ObjectType
		varNiatelemetryApicSnmpTrapDetails.Dn = varNiatelemetryApicSnmpTrapDetailsWithoutEmbeddedStruct.Dn
		varNiatelemetryApicSnmpTrapDetails.RecordType = varNiatelemetryApicSnmpTrapDetailsWithoutEmbeddedStruct.RecordType
		varNiatelemetryApicSnmpTrapDetails.RecordVersion = varNiatelemetryApicSnmpTrapDetailsWithoutEmbeddedStruct.RecordVersion
		varNiatelemetryApicSnmpTrapDetails.SiteName = varNiatelemetryApicSnmpTrapDetailsWithoutEmbeddedStruct.SiteName
		varNiatelemetryApicSnmpTrapDetails.V3SecLevel = varNiatelemetryApicSnmpTrapDetailsWithoutEmbeddedStruct.V3SecLevel
		varNiatelemetryApicSnmpTrapDetails.Ver = varNiatelemetryApicSnmpTrapDetailsWithoutEmbeddedStruct.Ver
		varNiatelemetryApicSnmpTrapDetails.RegisteredDevice = varNiatelemetryApicSnmpTrapDetailsWithoutEmbeddedStruct.RegisteredDevice
		*o = NiatelemetryApicSnmpTrapDetails(varNiatelemetryApicSnmpTrapDetails)
	} else {
		return err
	}

	varNiatelemetryApicSnmpTrapDetails := _NiatelemetryApicSnmpTrapDetails{}

	err = json.Unmarshal(bytes, &varNiatelemetryApicSnmpTrapDetails)
	if err == nil {
		o.MoBaseMo = varNiatelemetryApicSnmpTrapDetails.MoBaseMo
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Dn")
		delete(additionalProperties, "RecordType")
		delete(additionalProperties, "RecordVersion")
		delete(additionalProperties, "SiteName")
		delete(additionalProperties, "V3SecLevel")
		delete(additionalProperties, "Ver")
		delete(additionalProperties, "RegisteredDevice")

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

type NullableNiatelemetryApicSnmpTrapDetails struct {
	value *NiatelemetryApicSnmpTrapDetails
	isSet bool
}

func (v NullableNiatelemetryApicSnmpTrapDetails) Get() *NiatelemetryApicSnmpTrapDetails {
	return v.value
}

func (v *NullableNiatelemetryApicSnmpTrapDetails) Set(val *NiatelemetryApicSnmpTrapDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableNiatelemetryApicSnmpTrapDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableNiatelemetryApicSnmpTrapDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNiatelemetryApicSnmpTrapDetails(val *NiatelemetryApicSnmpTrapDetails) *NullableNiatelemetryApicSnmpTrapDetails {
	return &NullableNiatelemetryApicSnmpTrapDetails{value: val, isSet: true}
}

func (v NullableNiatelemetryApicSnmpTrapDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNiatelemetryApicSnmpTrapDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
