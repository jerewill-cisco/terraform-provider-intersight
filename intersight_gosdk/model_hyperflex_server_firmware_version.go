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

// HyperflexServerFirmwareVersion The server firmware version represents the UCS server firmware details.
type HyperflexServerFirmwareVersion struct {
	MoBaseMo
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string                           `json:"ObjectType"`
	AppCatalog *HyperflexAppCatalogRelationship `json:"AppCatalog,omitempty"`
	// An array of relationships to hyperflexServerFirmwareVersionEntry resources.
	ServerFirmwareVersionEntries []HyperflexServerFirmwareVersionEntryRelationship `json:"ServerFirmwareVersionEntries,omitempty"`
	AdditionalProperties         map[string]interface{}
}

type _HyperflexServerFirmwareVersion HyperflexServerFirmwareVersion

// NewHyperflexServerFirmwareVersion instantiates a new HyperflexServerFirmwareVersion object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHyperflexServerFirmwareVersion(classId string, objectType string) *HyperflexServerFirmwareVersion {
	this := HyperflexServerFirmwareVersion{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewHyperflexServerFirmwareVersionWithDefaults instantiates a new HyperflexServerFirmwareVersion object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHyperflexServerFirmwareVersionWithDefaults() *HyperflexServerFirmwareVersion {
	this := HyperflexServerFirmwareVersion{}
	var classId string = "hyperflex.ServerFirmwareVersion"
	this.ClassId = classId
	var objectType string = "hyperflex.ServerFirmwareVersion"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *HyperflexServerFirmwareVersion) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *HyperflexServerFirmwareVersion) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *HyperflexServerFirmwareVersion) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *HyperflexServerFirmwareVersion) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *HyperflexServerFirmwareVersion) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *HyperflexServerFirmwareVersion) SetObjectType(v string) {
	o.ObjectType = v
}

// GetAppCatalog returns the AppCatalog field value if set, zero value otherwise.
func (o *HyperflexServerFirmwareVersion) GetAppCatalog() HyperflexAppCatalogRelationship {
	if o == nil || o.AppCatalog == nil {
		var ret HyperflexAppCatalogRelationship
		return ret
	}
	return *o.AppCatalog
}

// GetAppCatalogOk returns a tuple with the AppCatalog field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexServerFirmwareVersion) GetAppCatalogOk() (*HyperflexAppCatalogRelationship, bool) {
	if o == nil || o.AppCatalog == nil {
		return nil, false
	}
	return o.AppCatalog, true
}

// HasAppCatalog returns a boolean if a field has been set.
func (o *HyperflexServerFirmwareVersion) HasAppCatalog() bool {
	if o != nil && o.AppCatalog != nil {
		return true
	}

	return false
}

// SetAppCatalog gets a reference to the given HyperflexAppCatalogRelationship and assigns it to the AppCatalog field.
func (o *HyperflexServerFirmwareVersion) SetAppCatalog(v HyperflexAppCatalogRelationship) {
	o.AppCatalog = &v
}

// GetServerFirmwareVersionEntries returns the ServerFirmwareVersionEntries field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HyperflexServerFirmwareVersion) GetServerFirmwareVersionEntries() []HyperflexServerFirmwareVersionEntryRelationship {
	if o == nil {
		var ret []HyperflexServerFirmwareVersionEntryRelationship
		return ret
	}
	return o.ServerFirmwareVersionEntries
}

// GetServerFirmwareVersionEntriesOk returns a tuple with the ServerFirmwareVersionEntries field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HyperflexServerFirmwareVersion) GetServerFirmwareVersionEntriesOk() (*[]HyperflexServerFirmwareVersionEntryRelationship, bool) {
	if o == nil || o.ServerFirmwareVersionEntries == nil {
		return nil, false
	}
	return &o.ServerFirmwareVersionEntries, true
}

// HasServerFirmwareVersionEntries returns a boolean if a field has been set.
func (o *HyperflexServerFirmwareVersion) HasServerFirmwareVersionEntries() bool {
	if o != nil && o.ServerFirmwareVersionEntries != nil {
		return true
	}

	return false
}

// SetServerFirmwareVersionEntries gets a reference to the given []HyperflexServerFirmwareVersionEntryRelationship and assigns it to the ServerFirmwareVersionEntries field.
func (o *HyperflexServerFirmwareVersion) SetServerFirmwareVersionEntries(v []HyperflexServerFirmwareVersionEntryRelationship) {
	o.ServerFirmwareVersionEntries = v
}

func (o HyperflexServerFirmwareVersion) MarshalJSON() ([]byte, error) {
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
	if o.AppCatalog != nil {
		toSerialize["AppCatalog"] = o.AppCatalog
	}
	if o.ServerFirmwareVersionEntries != nil {
		toSerialize["ServerFirmwareVersionEntries"] = o.ServerFirmwareVersionEntries
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *HyperflexServerFirmwareVersion) UnmarshalJSON(bytes []byte) (err error) {
	type HyperflexServerFirmwareVersionWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string                           `json:"ObjectType"`
		AppCatalog *HyperflexAppCatalogRelationship `json:"AppCatalog,omitempty"`
		// An array of relationships to hyperflexServerFirmwareVersionEntry resources.
		ServerFirmwareVersionEntries []HyperflexServerFirmwareVersionEntryRelationship `json:"ServerFirmwareVersionEntries,omitempty"`
	}

	varHyperflexServerFirmwareVersionWithoutEmbeddedStruct := HyperflexServerFirmwareVersionWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varHyperflexServerFirmwareVersionWithoutEmbeddedStruct)
	if err == nil {
		varHyperflexServerFirmwareVersion := _HyperflexServerFirmwareVersion{}
		varHyperflexServerFirmwareVersion.ClassId = varHyperflexServerFirmwareVersionWithoutEmbeddedStruct.ClassId
		varHyperflexServerFirmwareVersion.ObjectType = varHyperflexServerFirmwareVersionWithoutEmbeddedStruct.ObjectType
		varHyperflexServerFirmwareVersion.AppCatalog = varHyperflexServerFirmwareVersionWithoutEmbeddedStruct.AppCatalog
		varHyperflexServerFirmwareVersion.ServerFirmwareVersionEntries = varHyperflexServerFirmwareVersionWithoutEmbeddedStruct.ServerFirmwareVersionEntries
		*o = HyperflexServerFirmwareVersion(varHyperflexServerFirmwareVersion)
	} else {
		return err
	}

	varHyperflexServerFirmwareVersion := _HyperflexServerFirmwareVersion{}

	err = json.Unmarshal(bytes, &varHyperflexServerFirmwareVersion)
	if err == nil {
		o.MoBaseMo = varHyperflexServerFirmwareVersion.MoBaseMo
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "AppCatalog")
		delete(additionalProperties, "ServerFirmwareVersionEntries")

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

type NullableHyperflexServerFirmwareVersion struct {
	value *HyperflexServerFirmwareVersion
	isSet bool
}

func (v NullableHyperflexServerFirmwareVersion) Get() *HyperflexServerFirmwareVersion {
	return v.value
}

func (v *NullableHyperflexServerFirmwareVersion) Set(val *HyperflexServerFirmwareVersion) {
	v.value = val
	v.isSet = true
}

func (v NullableHyperflexServerFirmwareVersion) IsSet() bool {
	return v.isSet
}

func (v *NullableHyperflexServerFirmwareVersion) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHyperflexServerFirmwareVersion(val *HyperflexServerFirmwareVersion) *NullableHyperflexServerFirmwareVersion {
	return &NullableHyperflexServerFirmwareVersion{value: val, isSet: true}
}

func (v NullableHyperflexServerFirmwareVersion) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHyperflexServerFirmwareVersion) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
