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

// HyperflexTrackedDisk Tracked disk for a snapshot.
type HyperflexTrackedDisk struct {
	MoBaseComplexType
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string                 `json:"ObjectType"`
	DiskFiles  []HyperflexTrackedFile `json:"DiskFiles,omitempty"`
	// Disk type for a vm virtual disk. * `NONE` - The disk type for this VM is None. * `NATIVE` - The disk type for this VM is Native. * `NONNATIVE` - The disk type for this VM is Non-Native.
	DiskType             *string `json:"DiskType,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _HyperflexTrackedDisk HyperflexTrackedDisk

// NewHyperflexTrackedDisk instantiates a new HyperflexTrackedDisk object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHyperflexTrackedDisk(classId string, objectType string) *HyperflexTrackedDisk {
	this := HyperflexTrackedDisk{}
	this.ClassId = classId
	this.ObjectType = objectType
	var diskType string = "NONE"
	this.DiskType = &diskType
	return &this
}

// NewHyperflexTrackedDiskWithDefaults instantiates a new HyperflexTrackedDisk object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHyperflexTrackedDiskWithDefaults() *HyperflexTrackedDisk {
	this := HyperflexTrackedDisk{}
	var classId string = "hyperflex.TrackedDisk"
	this.ClassId = classId
	var objectType string = "hyperflex.TrackedDisk"
	this.ObjectType = objectType
	var diskType string = "NONE"
	this.DiskType = &diskType
	return &this
}

// GetClassId returns the ClassId field value
func (o *HyperflexTrackedDisk) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *HyperflexTrackedDisk) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *HyperflexTrackedDisk) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *HyperflexTrackedDisk) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *HyperflexTrackedDisk) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *HyperflexTrackedDisk) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDiskFiles returns the DiskFiles field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HyperflexTrackedDisk) GetDiskFiles() []HyperflexTrackedFile {
	if o == nil {
		var ret []HyperflexTrackedFile
		return ret
	}
	return o.DiskFiles
}

// GetDiskFilesOk returns a tuple with the DiskFiles field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HyperflexTrackedDisk) GetDiskFilesOk() (*[]HyperflexTrackedFile, bool) {
	if o == nil || o.DiskFiles == nil {
		return nil, false
	}
	return &o.DiskFiles, true
}

// HasDiskFiles returns a boolean if a field has been set.
func (o *HyperflexTrackedDisk) HasDiskFiles() bool {
	if o != nil && o.DiskFiles != nil {
		return true
	}

	return false
}

// SetDiskFiles gets a reference to the given []HyperflexTrackedFile and assigns it to the DiskFiles field.
func (o *HyperflexTrackedDisk) SetDiskFiles(v []HyperflexTrackedFile) {
	o.DiskFiles = v
}

// GetDiskType returns the DiskType field value if set, zero value otherwise.
func (o *HyperflexTrackedDisk) GetDiskType() string {
	if o == nil || o.DiskType == nil {
		var ret string
		return ret
	}
	return *o.DiskType
}

// GetDiskTypeOk returns a tuple with the DiskType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexTrackedDisk) GetDiskTypeOk() (*string, bool) {
	if o == nil || o.DiskType == nil {
		return nil, false
	}
	return o.DiskType, true
}

// HasDiskType returns a boolean if a field has been set.
func (o *HyperflexTrackedDisk) HasDiskType() bool {
	if o != nil && o.DiskType != nil {
		return true
	}

	return false
}

// SetDiskType gets a reference to the given string and assigns it to the DiskType field.
func (o *HyperflexTrackedDisk) SetDiskType(v string) {
	o.DiskType = &v
}

func (o HyperflexTrackedDisk) MarshalJSON() ([]byte, error) {
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
	if o.DiskFiles != nil {
		toSerialize["DiskFiles"] = o.DiskFiles
	}
	if o.DiskType != nil {
		toSerialize["DiskType"] = o.DiskType
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *HyperflexTrackedDisk) UnmarshalJSON(bytes []byte) (err error) {
	type HyperflexTrackedDiskWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string                 `json:"ObjectType"`
		DiskFiles  []HyperflexTrackedFile `json:"DiskFiles,omitempty"`
		// Disk type for a vm virtual disk. * `NONE` - The disk type for this VM is None. * `NATIVE` - The disk type for this VM is Native. * `NONNATIVE` - The disk type for this VM is Non-Native.
		DiskType *string `json:"DiskType,omitempty"`
	}

	varHyperflexTrackedDiskWithoutEmbeddedStruct := HyperflexTrackedDiskWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varHyperflexTrackedDiskWithoutEmbeddedStruct)
	if err == nil {
		varHyperflexTrackedDisk := _HyperflexTrackedDisk{}
		varHyperflexTrackedDisk.ClassId = varHyperflexTrackedDiskWithoutEmbeddedStruct.ClassId
		varHyperflexTrackedDisk.ObjectType = varHyperflexTrackedDiskWithoutEmbeddedStruct.ObjectType
		varHyperflexTrackedDisk.DiskFiles = varHyperflexTrackedDiskWithoutEmbeddedStruct.DiskFiles
		varHyperflexTrackedDisk.DiskType = varHyperflexTrackedDiskWithoutEmbeddedStruct.DiskType
		*o = HyperflexTrackedDisk(varHyperflexTrackedDisk)
	} else {
		return err
	}

	varHyperflexTrackedDisk := _HyperflexTrackedDisk{}

	err = json.Unmarshal(bytes, &varHyperflexTrackedDisk)
	if err == nil {
		o.MoBaseComplexType = varHyperflexTrackedDisk.MoBaseComplexType
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "DiskFiles")
		delete(additionalProperties, "DiskType")

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

type NullableHyperflexTrackedDisk struct {
	value *HyperflexTrackedDisk
	isSet bool
}

func (v NullableHyperflexTrackedDisk) Get() *HyperflexTrackedDisk {
	return v.value
}

func (v *NullableHyperflexTrackedDisk) Set(val *HyperflexTrackedDisk) {
	v.value = val
	v.isSet = true
}

func (v NullableHyperflexTrackedDisk) IsSet() bool {
	return v.isSet
}

func (v *NullableHyperflexTrackedDisk) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHyperflexTrackedDisk(val *HyperflexTrackedDisk) *NullableHyperflexTrackedDisk {
	return &NullableHyperflexTrackedDisk{value: val, isSet: true}
}

func (v NullableHyperflexTrackedDisk) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHyperflexTrackedDisk) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
