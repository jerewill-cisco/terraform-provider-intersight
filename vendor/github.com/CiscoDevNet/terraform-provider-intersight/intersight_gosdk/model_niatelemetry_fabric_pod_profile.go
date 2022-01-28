/*
Cisco Intersight

Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document.

API version: 1.0.9-5208
Contact: intersight@cisco.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package intersight

import (
	"encoding/json"
	"reflect"
	"strings"
)

// NiatelemetryFabricPodProfile Object to capture Fabric Pod Profile details.
type NiatelemetryFabricPodProfile struct {
	MoBaseMo
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Dn of the Fabric pod profile for APIC.
	Dn *string `json:"Dn,omitempty"`
	// Dn of the Children PodS for the above Pod Profile .
	FabricPodsList *string `json:"FabricPodsList,omitempty"`
	// Name of the Fabric pod profile for APIC.
	Name *string `json:"Name,omitempty"`
	// Type of record DCNM / APIC / SE. This determines the type of platform where inventory was collected.
	RecordType *string `json:"RecordType,omitempty"`
	// Version of record being pushed. This determines what was the API version for data available from the device.
	RecordVersion *string `json:"RecordVersion,omitempty"`
	// Name of the APIC site from which this data is being collected.
	SiteName             *string                              `json:"SiteName,omitempty"`
	RegisteredDevice     *AssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _NiatelemetryFabricPodProfile NiatelemetryFabricPodProfile

// NewNiatelemetryFabricPodProfile instantiates a new NiatelemetryFabricPodProfile object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNiatelemetryFabricPodProfile(classId string, objectType string) *NiatelemetryFabricPodProfile {
	this := NiatelemetryFabricPodProfile{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewNiatelemetryFabricPodProfileWithDefaults instantiates a new NiatelemetryFabricPodProfile object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNiatelemetryFabricPodProfileWithDefaults() *NiatelemetryFabricPodProfile {
	this := NiatelemetryFabricPodProfile{}
	var classId string = "niatelemetry.FabricPodProfile"
	this.ClassId = classId
	var objectType string = "niatelemetry.FabricPodProfile"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *NiatelemetryFabricPodProfile) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *NiatelemetryFabricPodProfile) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *NiatelemetryFabricPodProfile) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *NiatelemetryFabricPodProfile) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *NiatelemetryFabricPodProfile) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *NiatelemetryFabricPodProfile) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDn returns the Dn field value if set, zero value otherwise.
func (o *NiatelemetryFabricPodProfile) GetDn() string {
	if o == nil || o.Dn == nil {
		var ret string
		return ret
	}
	return *o.Dn
}

// GetDnOk returns a tuple with the Dn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryFabricPodProfile) GetDnOk() (*string, bool) {
	if o == nil || o.Dn == nil {
		return nil, false
	}
	return o.Dn, true
}

// HasDn returns a boolean if a field has been set.
func (o *NiatelemetryFabricPodProfile) HasDn() bool {
	if o != nil && o.Dn != nil {
		return true
	}

	return false
}

// SetDn gets a reference to the given string and assigns it to the Dn field.
func (o *NiatelemetryFabricPodProfile) SetDn(v string) {
	o.Dn = &v
}

// GetFabricPodsList returns the FabricPodsList field value if set, zero value otherwise.
func (o *NiatelemetryFabricPodProfile) GetFabricPodsList() string {
	if o == nil || o.FabricPodsList == nil {
		var ret string
		return ret
	}
	return *o.FabricPodsList
}

// GetFabricPodsListOk returns a tuple with the FabricPodsList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryFabricPodProfile) GetFabricPodsListOk() (*string, bool) {
	if o == nil || o.FabricPodsList == nil {
		return nil, false
	}
	return o.FabricPodsList, true
}

// HasFabricPodsList returns a boolean if a field has been set.
func (o *NiatelemetryFabricPodProfile) HasFabricPodsList() bool {
	if o != nil && o.FabricPodsList != nil {
		return true
	}

	return false
}

// SetFabricPodsList gets a reference to the given string and assigns it to the FabricPodsList field.
func (o *NiatelemetryFabricPodProfile) SetFabricPodsList(v string) {
	o.FabricPodsList = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *NiatelemetryFabricPodProfile) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryFabricPodProfile) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *NiatelemetryFabricPodProfile) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *NiatelemetryFabricPodProfile) SetName(v string) {
	o.Name = &v
}

// GetRecordType returns the RecordType field value if set, zero value otherwise.
func (o *NiatelemetryFabricPodProfile) GetRecordType() string {
	if o == nil || o.RecordType == nil {
		var ret string
		return ret
	}
	return *o.RecordType
}

// GetRecordTypeOk returns a tuple with the RecordType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryFabricPodProfile) GetRecordTypeOk() (*string, bool) {
	if o == nil || o.RecordType == nil {
		return nil, false
	}
	return o.RecordType, true
}

// HasRecordType returns a boolean if a field has been set.
func (o *NiatelemetryFabricPodProfile) HasRecordType() bool {
	if o != nil && o.RecordType != nil {
		return true
	}

	return false
}

// SetRecordType gets a reference to the given string and assigns it to the RecordType field.
func (o *NiatelemetryFabricPodProfile) SetRecordType(v string) {
	o.RecordType = &v
}

// GetRecordVersion returns the RecordVersion field value if set, zero value otherwise.
func (o *NiatelemetryFabricPodProfile) GetRecordVersion() string {
	if o == nil || o.RecordVersion == nil {
		var ret string
		return ret
	}
	return *o.RecordVersion
}

// GetRecordVersionOk returns a tuple with the RecordVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryFabricPodProfile) GetRecordVersionOk() (*string, bool) {
	if o == nil || o.RecordVersion == nil {
		return nil, false
	}
	return o.RecordVersion, true
}

// HasRecordVersion returns a boolean if a field has been set.
func (o *NiatelemetryFabricPodProfile) HasRecordVersion() bool {
	if o != nil && o.RecordVersion != nil {
		return true
	}

	return false
}

// SetRecordVersion gets a reference to the given string and assigns it to the RecordVersion field.
func (o *NiatelemetryFabricPodProfile) SetRecordVersion(v string) {
	o.RecordVersion = &v
}

// GetSiteName returns the SiteName field value if set, zero value otherwise.
func (o *NiatelemetryFabricPodProfile) GetSiteName() string {
	if o == nil || o.SiteName == nil {
		var ret string
		return ret
	}
	return *o.SiteName
}

// GetSiteNameOk returns a tuple with the SiteName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryFabricPodProfile) GetSiteNameOk() (*string, bool) {
	if o == nil || o.SiteName == nil {
		return nil, false
	}
	return o.SiteName, true
}

// HasSiteName returns a boolean if a field has been set.
func (o *NiatelemetryFabricPodProfile) HasSiteName() bool {
	if o != nil && o.SiteName != nil {
		return true
	}

	return false
}

// SetSiteName gets a reference to the given string and assigns it to the SiteName field.
func (o *NiatelemetryFabricPodProfile) SetSiteName(v string) {
	o.SiteName = &v
}

// GetRegisteredDevice returns the RegisteredDevice field value if set, zero value otherwise.
func (o *NiatelemetryFabricPodProfile) GetRegisteredDevice() AssetDeviceRegistrationRelationship {
	if o == nil || o.RegisteredDevice == nil {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.RegisteredDevice
}

// GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryFabricPodProfile) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil || o.RegisteredDevice == nil {
		return nil, false
	}
	return o.RegisteredDevice, true
}

// HasRegisteredDevice returns a boolean if a field has been set.
func (o *NiatelemetryFabricPodProfile) HasRegisteredDevice() bool {
	if o != nil && o.RegisteredDevice != nil {
		return true
	}

	return false
}

// SetRegisteredDevice gets a reference to the given AssetDeviceRegistrationRelationship and assigns it to the RegisteredDevice field.
func (o *NiatelemetryFabricPodProfile) SetRegisteredDevice(v AssetDeviceRegistrationRelationship) {
	o.RegisteredDevice = &v
}

func (o NiatelemetryFabricPodProfile) MarshalJSON() ([]byte, error) {
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
	if o.FabricPodsList != nil {
		toSerialize["FabricPodsList"] = o.FabricPodsList
	}
	if o.Name != nil {
		toSerialize["Name"] = o.Name
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
	if o.RegisteredDevice != nil {
		toSerialize["RegisteredDevice"] = o.RegisteredDevice
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *NiatelemetryFabricPodProfile) UnmarshalJSON(bytes []byte) (err error) {
	type NiatelemetryFabricPodProfileWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Dn of the Fabric pod profile for APIC.
		Dn *string `json:"Dn,omitempty"`
		// Dn of the Children PodS for the above Pod Profile .
		FabricPodsList *string `json:"FabricPodsList,omitempty"`
		// Name of the Fabric pod profile for APIC.
		Name *string `json:"Name,omitempty"`
		// Type of record DCNM / APIC / SE. This determines the type of platform where inventory was collected.
		RecordType *string `json:"RecordType,omitempty"`
		// Version of record being pushed. This determines what was the API version for data available from the device.
		RecordVersion *string `json:"RecordVersion,omitempty"`
		// Name of the APIC site from which this data is being collected.
		SiteName         *string                              `json:"SiteName,omitempty"`
		RegisteredDevice *AssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	}

	varNiatelemetryFabricPodProfileWithoutEmbeddedStruct := NiatelemetryFabricPodProfileWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varNiatelemetryFabricPodProfileWithoutEmbeddedStruct)
	if err == nil {
		varNiatelemetryFabricPodProfile := _NiatelemetryFabricPodProfile{}
		varNiatelemetryFabricPodProfile.ClassId = varNiatelemetryFabricPodProfileWithoutEmbeddedStruct.ClassId
		varNiatelemetryFabricPodProfile.ObjectType = varNiatelemetryFabricPodProfileWithoutEmbeddedStruct.ObjectType
		varNiatelemetryFabricPodProfile.Dn = varNiatelemetryFabricPodProfileWithoutEmbeddedStruct.Dn
		varNiatelemetryFabricPodProfile.FabricPodsList = varNiatelemetryFabricPodProfileWithoutEmbeddedStruct.FabricPodsList
		varNiatelemetryFabricPodProfile.Name = varNiatelemetryFabricPodProfileWithoutEmbeddedStruct.Name
		varNiatelemetryFabricPodProfile.RecordType = varNiatelemetryFabricPodProfileWithoutEmbeddedStruct.RecordType
		varNiatelemetryFabricPodProfile.RecordVersion = varNiatelemetryFabricPodProfileWithoutEmbeddedStruct.RecordVersion
		varNiatelemetryFabricPodProfile.SiteName = varNiatelemetryFabricPodProfileWithoutEmbeddedStruct.SiteName
		varNiatelemetryFabricPodProfile.RegisteredDevice = varNiatelemetryFabricPodProfileWithoutEmbeddedStruct.RegisteredDevice
		*o = NiatelemetryFabricPodProfile(varNiatelemetryFabricPodProfile)
	} else {
		return err
	}

	varNiatelemetryFabricPodProfile := _NiatelemetryFabricPodProfile{}

	err = json.Unmarshal(bytes, &varNiatelemetryFabricPodProfile)
	if err == nil {
		o.MoBaseMo = varNiatelemetryFabricPodProfile.MoBaseMo
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Dn")
		delete(additionalProperties, "FabricPodsList")
		delete(additionalProperties, "Name")
		delete(additionalProperties, "RecordType")
		delete(additionalProperties, "RecordVersion")
		delete(additionalProperties, "SiteName")
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

type NullableNiatelemetryFabricPodProfile struct {
	value *NiatelemetryFabricPodProfile
	isSet bool
}

func (v NullableNiatelemetryFabricPodProfile) Get() *NiatelemetryFabricPodProfile {
	return v.value
}

func (v *NullableNiatelemetryFabricPodProfile) Set(val *NiatelemetryFabricPodProfile) {
	v.value = val
	v.isSet = true
}

func (v NullableNiatelemetryFabricPodProfile) IsSet() bool {
	return v.isSet
}

func (v *NullableNiatelemetryFabricPodProfile) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNiatelemetryFabricPodProfile(val *NiatelemetryFabricPodProfile) *NullableNiatelemetryFabricPodProfile {
	return &NullableNiatelemetryFabricPodProfile{value: val, isSet: true}
}

func (v NullableNiatelemetryFabricPodProfile) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNiatelemetryFabricPodProfile) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
