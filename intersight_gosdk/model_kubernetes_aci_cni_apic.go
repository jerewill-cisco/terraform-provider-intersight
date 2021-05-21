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

// KubernetesAciCniApic Internally generated object of claimed APIC device known to Razor.
type KubernetesAciCniApic struct {
	MoBaseMo
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// The Moid of the apic device under the asset service.
	AssetApicMoid *string `json:"AssetApicMoid,omitempty"`
	// The number of ACI CNI profiles configured for this APIC.
	NumAciCniProfiles *int64 `json:"NumAciCniProfiles,omitempty"`
	// An array of relationships to kubernetesAciCniProfile resources.
	AciCniProfiles       []KubernetesAciCniProfileRelationship `json:"AciCniProfiles,omitempty"`
	Organization         *OrganizationOrganizationRelationship `json:"Organization,omitempty"`
	RegisteredDevice     *AssetDeviceRegistrationRelationship  `json:"RegisteredDevice,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _KubernetesAciCniApic KubernetesAciCniApic

// NewKubernetesAciCniApic instantiates a new KubernetesAciCniApic object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKubernetesAciCniApic(classId string, objectType string) *KubernetesAciCniApic {
	this := KubernetesAciCniApic{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewKubernetesAciCniApicWithDefaults instantiates a new KubernetesAciCniApic object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKubernetesAciCniApicWithDefaults() *KubernetesAciCniApic {
	this := KubernetesAciCniApic{}
	var classId string = "kubernetes.AciCniApic"
	this.ClassId = classId
	var objectType string = "kubernetes.AciCniApic"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *KubernetesAciCniApic) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *KubernetesAciCniApic) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *KubernetesAciCniApic) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *KubernetesAciCniApic) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *KubernetesAciCniApic) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *KubernetesAciCniApic) SetObjectType(v string) {
	o.ObjectType = v
}

// GetAssetApicMoid returns the AssetApicMoid field value if set, zero value otherwise.
func (o *KubernetesAciCniApic) GetAssetApicMoid() string {
	if o == nil || o.AssetApicMoid == nil {
		var ret string
		return ret
	}
	return *o.AssetApicMoid
}

// GetAssetApicMoidOk returns a tuple with the AssetApicMoid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesAciCniApic) GetAssetApicMoidOk() (*string, bool) {
	if o == nil || o.AssetApicMoid == nil {
		return nil, false
	}
	return o.AssetApicMoid, true
}

// HasAssetApicMoid returns a boolean if a field has been set.
func (o *KubernetesAciCniApic) HasAssetApicMoid() bool {
	if o != nil && o.AssetApicMoid != nil {
		return true
	}

	return false
}

// SetAssetApicMoid gets a reference to the given string and assigns it to the AssetApicMoid field.
func (o *KubernetesAciCniApic) SetAssetApicMoid(v string) {
	o.AssetApicMoid = &v
}

// GetNumAciCniProfiles returns the NumAciCniProfiles field value if set, zero value otherwise.
func (o *KubernetesAciCniApic) GetNumAciCniProfiles() int64 {
	if o == nil || o.NumAciCniProfiles == nil {
		var ret int64
		return ret
	}
	return *o.NumAciCniProfiles
}

// GetNumAciCniProfilesOk returns a tuple with the NumAciCniProfiles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesAciCniApic) GetNumAciCniProfilesOk() (*int64, bool) {
	if o == nil || o.NumAciCniProfiles == nil {
		return nil, false
	}
	return o.NumAciCniProfiles, true
}

// HasNumAciCniProfiles returns a boolean if a field has been set.
func (o *KubernetesAciCniApic) HasNumAciCniProfiles() bool {
	if o != nil && o.NumAciCniProfiles != nil {
		return true
	}

	return false
}

// SetNumAciCniProfiles gets a reference to the given int64 and assigns it to the NumAciCniProfiles field.
func (o *KubernetesAciCniApic) SetNumAciCniProfiles(v int64) {
	o.NumAciCniProfiles = &v
}

// GetAciCniProfiles returns the AciCniProfiles field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *KubernetesAciCniApic) GetAciCniProfiles() []KubernetesAciCniProfileRelationship {
	if o == nil {
		var ret []KubernetesAciCniProfileRelationship
		return ret
	}
	return o.AciCniProfiles
}

// GetAciCniProfilesOk returns a tuple with the AciCniProfiles field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KubernetesAciCniApic) GetAciCniProfilesOk() (*[]KubernetesAciCniProfileRelationship, bool) {
	if o == nil || o.AciCniProfiles == nil {
		return nil, false
	}
	return &o.AciCniProfiles, true
}

// HasAciCniProfiles returns a boolean if a field has been set.
func (o *KubernetesAciCniApic) HasAciCniProfiles() bool {
	if o != nil && o.AciCniProfiles != nil {
		return true
	}

	return false
}

// SetAciCniProfiles gets a reference to the given []KubernetesAciCniProfileRelationship and assigns it to the AciCniProfiles field.
func (o *KubernetesAciCniApic) SetAciCniProfiles(v []KubernetesAciCniProfileRelationship) {
	o.AciCniProfiles = v
}

// GetOrganization returns the Organization field value if set, zero value otherwise.
func (o *KubernetesAciCniApic) GetOrganization() OrganizationOrganizationRelationship {
	if o == nil || o.Organization == nil {
		var ret OrganizationOrganizationRelationship
		return ret
	}
	return *o.Organization
}

// GetOrganizationOk returns a tuple with the Organization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesAciCniApic) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool) {
	if o == nil || o.Organization == nil {
		return nil, false
	}
	return o.Organization, true
}

// HasOrganization returns a boolean if a field has been set.
func (o *KubernetesAciCniApic) HasOrganization() bool {
	if o != nil && o.Organization != nil {
		return true
	}

	return false
}

// SetOrganization gets a reference to the given OrganizationOrganizationRelationship and assigns it to the Organization field.
func (o *KubernetesAciCniApic) SetOrganization(v OrganizationOrganizationRelationship) {
	o.Organization = &v
}

// GetRegisteredDevice returns the RegisteredDevice field value if set, zero value otherwise.
func (o *KubernetesAciCniApic) GetRegisteredDevice() AssetDeviceRegistrationRelationship {
	if o == nil || o.RegisteredDevice == nil {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.RegisteredDevice
}

// GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesAciCniApic) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil || o.RegisteredDevice == nil {
		return nil, false
	}
	return o.RegisteredDevice, true
}

// HasRegisteredDevice returns a boolean if a field has been set.
func (o *KubernetesAciCniApic) HasRegisteredDevice() bool {
	if o != nil && o.RegisteredDevice != nil {
		return true
	}

	return false
}

// SetRegisteredDevice gets a reference to the given AssetDeviceRegistrationRelationship and assigns it to the RegisteredDevice field.
func (o *KubernetesAciCniApic) SetRegisteredDevice(v AssetDeviceRegistrationRelationship) {
	o.RegisteredDevice = &v
}

func (o KubernetesAciCniApic) MarshalJSON() ([]byte, error) {
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
	if o.AssetApicMoid != nil {
		toSerialize["AssetApicMoid"] = o.AssetApicMoid
	}
	if o.NumAciCniProfiles != nil {
		toSerialize["NumAciCniProfiles"] = o.NumAciCniProfiles
	}
	if o.AciCniProfiles != nil {
		toSerialize["AciCniProfiles"] = o.AciCniProfiles
	}
	if o.Organization != nil {
		toSerialize["Organization"] = o.Organization
	}
	if o.RegisteredDevice != nil {
		toSerialize["RegisteredDevice"] = o.RegisteredDevice
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *KubernetesAciCniApic) UnmarshalJSON(bytes []byte) (err error) {
	type KubernetesAciCniApicWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// The Moid of the apic device under the asset service.
		AssetApicMoid *string `json:"AssetApicMoid,omitempty"`
		// The number of ACI CNI profiles configured for this APIC.
		NumAciCniProfiles *int64 `json:"NumAciCniProfiles,omitempty"`
		// An array of relationships to kubernetesAciCniProfile resources.
		AciCniProfiles   []KubernetesAciCniProfileRelationship `json:"AciCniProfiles,omitempty"`
		Organization     *OrganizationOrganizationRelationship `json:"Organization,omitempty"`
		RegisteredDevice *AssetDeviceRegistrationRelationship  `json:"RegisteredDevice,omitempty"`
	}

	varKubernetesAciCniApicWithoutEmbeddedStruct := KubernetesAciCniApicWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varKubernetesAciCniApicWithoutEmbeddedStruct)
	if err == nil {
		varKubernetesAciCniApic := _KubernetesAciCniApic{}
		varKubernetesAciCniApic.ClassId = varKubernetesAciCniApicWithoutEmbeddedStruct.ClassId
		varKubernetesAciCniApic.ObjectType = varKubernetesAciCniApicWithoutEmbeddedStruct.ObjectType
		varKubernetesAciCniApic.AssetApicMoid = varKubernetesAciCniApicWithoutEmbeddedStruct.AssetApicMoid
		varKubernetesAciCniApic.NumAciCniProfiles = varKubernetesAciCniApicWithoutEmbeddedStruct.NumAciCniProfiles
		varKubernetesAciCniApic.AciCniProfiles = varKubernetesAciCniApicWithoutEmbeddedStruct.AciCniProfiles
		varKubernetesAciCniApic.Organization = varKubernetesAciCniApicWithoutEmbeddedStruct.Organization
		varKubernetesAciCniApic.RegisteredDevice = varKubernetesAciCniApicWithoutEmbeddedStruct.RegisteredDevice
		*o = KubernetesAciCniApic(varKubernetesAciCniApic)
	} else {
		return err
	}

	varKubernetesAciCniApic := _KubernetesAciCniApic{}

	err = json.Unmarshal(bytes, &varKubernetesAciCniApic)
	if err == nil {
		o.MoBaseMo = varKubernetesAciCniApic.MoBaseMo
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "AssetApicMoid")
		delete(additionalProperties, "NumAciCniProfiles")
		delete(additionalProperties, "AciCniProfiles")
		delete(additionalProperties, "Organization")
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

type NullableKubernetesAciCniApic struct {
	value *KubernetesAciCniApic
	isSet bool
}

func (v NullableKubernetesAciCniApic) Get() *KubernetesAciCniApic {
	return v.value
}

func (v *NullableKubernetesAciCniApic) Set(val *KubernetesAciCniApic) {
	v.value = val
	v.isSet = true
}

func (v NullableKubernetesAciCniApic) IsSet() bool {
	return v.isSet
}

func (v *NullableKubernetesAciCniApic) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKubernetesAciCniApic(val *KubernetesAciCniApic) *NullableKubernetesAciCniApic {
	return &NullableKubernetesAciCniApic{value: val, isSet: true}
}

func (v NullableKubernetesAciCniApic) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKubernetesAciCniApic) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
