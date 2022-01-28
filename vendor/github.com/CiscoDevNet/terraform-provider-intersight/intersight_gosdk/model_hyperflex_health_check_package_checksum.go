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
	"time"
)

// HyperflexHealthCheckPackageChecksum HyperFlex health check Debian Package SHA512 checksum.
type HyperflexHealthCheckPackageChecksum struct {
	MoBaseMo
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// SHA512 checksum of the health check package.
	Checksum *string `json:"Checksum,omitempty"`
	// Name of the health check Debian package.
	Name *string `json:"Name,omitempty"`
	// HyperFlex health check Debian package file name.
	PackageName *string `json:"PackageName,omitempty"`
	// Timestamp of last update of HyperFlex health check package checksum.
	Timestamp *time.Time `json:"Timestamp,omitempty"`
	// HyperFlex health check Debian Package Version.
	Version              *string `json:"Version,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _HyperflexHealthCheckPackageChecksum HyperflexHealthCheckPackageChecksum

// NewHyperflexHealthCheckPackageChecksum instantiates a new HyperflexHealthCheckPackageChecksum object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHyperflexHealthCheckPackageChecksum(classId string, objectType string) *HyperflexHealthCheckPackageChecksum {
	this := HyperflexHealthCheckPackageChecksum{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewHyperflexHealthCheckPackageChecksumWithDefaults instantiates a new HyperflexHealthCheckPackageChecksum object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHyperflexHealthCheckPackageChecksumWithDefaults() *HyperflexHealthCheckPackageChecksum {
	this := HyperflexHealthCheckPackageChecksum{}
	var classId string = "hyperflex.HealthCheckPackageChecksum"
	this.ClassId = classId
	var objectType string = "hyperflex.HealthCheckPackageChecksum"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *HyperflexHealthCheckPackageChecksum) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *HyperflexHealthCheckPackageChecksum) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *HyperflexHealthCheckPackageChecksum) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *HyperflexHealthCheckPackageChecksum) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *HyperflexHealthCheckPackageChecksum) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *HyperflexHealthCheckPackageChecksum) SetObjectType(v string) {
	o.ObjectType = v
}

// GetChecksum returns the Checksum field value if set, zero value otherwise.
func (o *HyperflexHealthCheckPackageChecksum) GetChecksum() string {
	if o == nil || o.Checksum == nil {
		var ret string
		return ret
	}
	return *o.Checksum
}

// GetChecksumOk returns a tuple with the Checksum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexHealthCheckPackageChecksum) GetChecksumOk() (*string, bool) {
	if o == nil || o.Checksum == nil {
		return nil, false
	}
	return o.Checksum, true
}

// HasChecksum returns a boolean if a field has been set.
func (o *HyperflexHealthCheckPackageChecksum) HasChecksum() bool {
	if o != nil && o.Checksum != nil {
		return true
	}

	return false
}

// SetChecksum gets a reference to the given string and assigns it to the Checksum field.
func (o *HyperflexHealthCheckPackageChecksum) SetChecksum(v string) {
	o.Checksum = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *HyperflexHealthCheckPackageChecksum) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexHealthCheckPackageChecksum) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *HyperflexHealthCheckPackageChecksum) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *HyperflexHealthCheckPackageChecksum) SetName(v string) {
	o.Name = &v
}

// GetPackageName returns the PackageName field value if set, zero value otherwise.
func (o *HyperflexHealthCheckPackageChecksum) GetPackageName() string {
	if o == nil || o.PackageName == nil {
		var ret string
		return ret
	}
	return *o.PackageName
}

// GetPackageNameOk returns a tuple with the PackageName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexHealthCheckPackageChecksum) GetPackageNameOk() (*string, bool) {
	if o == nil || o.PackageName == nil {
		return nil, false
	}
	return o.PackageName, true
}

// HasPackageName returns a boolean if a field has been set.
func (o *HyperflexHealthCheckPackageChecksum) HasPackageName() bool {
	if o != nil && o.PackageName != nil {
		return true
	}

	return false
}

// SetPackageName gets a reference to the given string and assigns it to the PackageName field.
func (o *HyperflexHealthCheckPackageChecksum) SetPackageName(v string) {
	o.PackageName = &v
}

// GetTimestamp returns the Timestamp field value if set, zero value otherwise.
func (o *HyperflexHealthCheckPackageChecksum) GetTimestamp() time.Time {
	if o == nil || o.Timestamp == nil {
		var ret time.Time
		return ret
	}
	return *o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexHealthCheckPackageChecksum) GetTimestampOk() (*time.Time, bool) {
	if o == nil || o.Timestamp == nil {
		return nil, false
	}
	return o.Timestamp, true
}

// HasTimestamp returns a boolean if a field has been set.
func (o *HyperflexHealthCheckPackageChecksum) HasTimestamp() bool {
	if o != nil && o.Timestamp != nil {
		return true
	}

	return false
}

// SetTimestamp gets a reference to the given time.Time and assigns it to the Timestamp field.
func (o *HyperflexHealthCheckPackageChecksum) SetTimestamp(v time.Time) {
	o.Timestamp = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *HyperflexHealthCheckPackageChecksum) GetVersion() string {
	if o == nil || o.Version == nil {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexHealthCheckPackageChecksum) GetVersionOk() (*string, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *HyperflexHealthCheckPackageChecksum) HasVersion() bool {
	if o != nil && o.Version != nil {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *HyperflexHealthCheckPackageChecksum) SetVersion(v string) {
	o.Version = &v
}

func (o HyperflexHealthCheckPackageChecksum) MarshalJSON() ([]byte, error) {
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
	if o.Checksum != nil {
		toSerialize["Checksum"] = o.Checksum
	}
	if o.Name != nil {
		toSerialize["Name"] = o.Name
	}
	if o.PackageName != nil {
		toSerialize["PackageName"] = o.PackageName
	}
	if o.Timestamp != nil {
		toSerialize["Timestamp"] = o.Timestamp
	}
	if o.Version != nil {
		toSerialize["Version"] = o.Version
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *HyperflexHealthCheckPackageChecksum) UnmarshalJSON(bytes []byte) (err error) {
	type HyperflexHealthCheckPackageChecksumWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// SHA512 checksum of the health check package.
		Checksum *string `json:"Checksum,omitempty"`
		// Name of the health check Debian package.
		Name *string `json:"Name,omitempty"`
		// HyperFlex health check Debian package file name.
		PackageName *string `json:"PackageName,omitempty"`
		// Timestamp of last update of HyperFlex health check package checksum.
		Timestamp *time.Time `json:"Timestamp,omitempty"`
		// HyperFlex health check Debian Package Version.
		Version *string `json:"Version,omitempty"`
	}

	varHyperflexHealthCheckPackageChecksumWithoutEmbeddedStruct := HyperflexHealthCheckPackageChecksumWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varHyperflexHealthCheckPackageChecksumWithoutEmbeddedStruct)
	if err == nil {
		varHyperflexHealthCheckPackageChecksum := _HyperflexHealthCheckPackageChecksum{}
		varHyperflexHealthCheckPackageChecksum.ClassId = varHyperflexHealthCheckPackageChecksumWithoutEmbeddedStruct.ClassId
		varHyperflexHealthCheckPackageChecksum.ObjectType = varHyperflexHealthCheckPackageChecksumWithoutEmbeddedStruct.ObjectType
		varHyperflexHealthCheckPackageChecksum.Checksum = varHyperflexHealthCheckPackageChecksumWithoutEmbeddedStruct.Checksum
		varHyperflexHealthCheckPackageChecksum.Name = varHyperflexHealthCheckPackageChecksumWithoutEmbeddedStruct.Name
		varHyperflexHealthCheckPackageChecksum.PackageName = varHyperflexHealthCheckPackageChecksumWithoutEmbeddedStruct.PackageName
		varHyperflexHealthCheckPackageChecksum.Timestamp = varHyperflexHealthCheckPackageChecksumWithoutEmbeddedStruct.Timestamp
		varHyperflexHealthCheckPackageChecksum.Version = varHyperflexHealthCheckPackageChecksumWithoutEmbeddedStruct.Version
		*o = HyperflexHealthCheckPackageChecksum(varHyperflexHealthCheckPackageChecksum)
	} else {
		return err
	}

	varHyperflexHealthCheckPackageChecksum := _HyperflexHealthCheckPackageChecksum{}

	err = json.Unmarshal(bytes, &varHyperflexHealthCheckPackageChecksum)
	if err == nil {
		o.MoBaseMo = varHyperflexHealthCheckPackageChecksum.MoBaseMo
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Checksum")
		delete(additionalProperties, "Name")
		delete(additionalProperties, "PackageName")
		delete(additionalProperties, "Timestamp")
		delete(additionalProperties, "Version")

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

type NullableHyperflexHealthCheckPackageChecksum struct {
	value *HyperflexHealthCheckPackageChecksum
	isSet bool
}

func (v NullableHyperflexHealthCheckPackageChecksum) Get() *HyperflexHealthCheckPackageChecksum {
	return v.value
}

func (v *NullableHyperflexHealthCheckPackageChecksum) Set(val *HyperflexHealthCheckPackageChecksum) {
	v.value = val
	v.isSet = true
}

func (v NullableHyperflexHealthCheckPackageChecksum) IsSet() bool {
	return v.isSet
}

func (v *NullableHyperflexHealthCheckPackageChecksum) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHyperflexHealthCheckPackageChecksum(val *HyperflexHealthCheckPackageChecksum) *NullableHyperflexHealthCheckPackageChecksum {
	return &NullableHyperflexHealthCheckPackageChecksum{value: val, isSet: true}
}

func (v NullableHyperflexHealthCheckPackageChecksum) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHyperflexHealthCheckPackageChecksum) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
