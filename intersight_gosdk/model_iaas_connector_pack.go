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

// IaasConnectorPack Describes about all the connector pack versions running currently in UCSD.
type IaasConnectorPack struct {
	MoBaseMo
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Complete version of the connector pack including build number.
	CompleteVersion *string  `json:"CompleteVersion,omitempty"`
	DependencyNames []string `json:"DependencyNames,omitempty"`
	// Version of the connector pack that is last downloaded successfully to UCSD.
	DownloadedVersion *string `json:"DownloadedVersion,omitempty"`
	// Name of the connector pack running on the UCSD.
	Name *string `json:"Name,omitempty"`
	// State of the connector pack whether it is enabled or disabled.
	State *string `json:"State,omitempty"`
	// Version of the connector pack.
	Version              *string                   `json:"Version,omitempty"`
	Guid                 *IaasUcsdInfoRelationship `json:"Guid,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _IaasConnectorPack IaasConnectorPack

// NewIaasConnectorPack instantiates a new IaasConnectorPack object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIaasConnectorPack(classId string, objectType string) *IaasConnectorPack {
	this := IaasConnectorPack{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewIaasConnectorPackWithDefaults instantiates a new IaasConnectorPack object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIaasConnectorPackWithDefaults() *IaasConnectorPack {
	this := IaasConnectorPack{}
	var classId string = "iaas.ConnectorPack"
	this.ClassId = classId
	var objectType string = "iaas.ConnectorPack"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *IaasConnectorPack) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *IaasConnectorPack) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *IaasConnectorPack) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *IaasConnectorPack) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *IaasConnectorPack) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *IaasConnectorPack) SetObjectType(v string) {
	o.ObjectType = v
}

// GetCompleteVersion returns the CompleteVersion field value if set, zero value otherwise.
func (o *IaasConnectorPack) GetCompleteVersion() string {
	if o == nil || o.CompleteVersion == nil {
		var ret string
		return ret
	}
	return *o.CompleteVersion
}

// GetCompleteVersionOk returns a tuple with the CompleteVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IaasConnectorPack) GetCompleteVersionOk() (*string, bool) {
	if o == nil || o.CompleteVersion == nil {
		return nil, false
	}
	return o.CompleteVersion, true
}

// HasCompleteVersion returns a boolean if a field has been set.
func (o *IaasConnectorPack) HasCompleteVersion() bool {
	if o != nil && o.CompleteVersion != nil {
		return true
	}

	return false
}

// SetCompleteVersion gets a reference to the given string and assigns it to the CompleteVersion field.
func (o *IaasConnectorPack) SetCompleteVersion(v string) {
	o.CompleteVersion = &v
}

// GetDependencyNames returns the DependencyNames field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IaasConnectorPack) GetDependencyNames() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.DependencyNames
}

// GetDependencyNamesOk returns a tuple with the DependencyNames field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IaasConnectorPack) GetDependencyNamesOk() (*[]string, bool) {
	if o == nil || o.DependencyNames == nil {
		return nil, false
	}
	return &o.DependencyNames, true
}

// HasDependencyNames returns a boolean if a field has been set.
func (o *IaasConnectorPack) HasDependencyNames() bool {
	if o != nil && o.DependencyNames != nil {
		return true
	}

	return false
}

// SetDependencyNames gets a reference to the given []string and assigns it to the DependencyNames field.
func (o *IaasConnectorPack) SetDependencyNames(v []string) {
	o.DependencyNames = v
}

// GetDownloadedVersion returns the DownloadedVersion field value if set, zero value otherwise.
func (o *IaasConnectorPack) GetDownloadedVersion() string {
	if o == nil || o.DownloadedVersion == nil {
		var ret string
		return ret
	}
	return *o.DownloadedVersion
}

// GetDownloadedVersionOk returns a tuple with the DownloadedVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IaasConnectorPack) GetDownloadedVersionOk() (*string, bool) {
	if o == nil || o.DownloadedVersion == nil {
		return nil, false
	}
	return o.DownloadedVersion, true
}

// HasDownloadedVersion returns a boolean if a field has been set.
func (o *IaasConnectorPack) HasDownloadedVersion() bool {
	if o != nil && o.DownloadedVersion != nil {
		return true
	}

	return false
}

// SetDownloadedVersion gets a reference to the given string and assigns it to the DownloadedVersion field.
func (o *IaasConnectorPack) SetDownloadedVersion(v string) {
	o.DownloadedVersion = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *IaasConnectorPack) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IaasConnectorPack) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *IaasConnectorPack) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *IaasConnectorPack) SetName(v string) {
	o.Name = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *IaasConnectorPack) GetState() string {
	if o == nil || o.State == nil {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IaasConnectorPack) GetStateOk() (*string, bool) {
	if o == nil || o.State == nil {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *IaasConnectorPack) HasState() bool {
	if o != nil && o.State != nil {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *IaasConnectorPack) SetState(v string) {
	o.State = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *IaasConnectorPack) GetVersion() string {
	if o == nil || o.Version == nil {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IaasConnectorPack) GetVersionOk() (*string, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *IaasConnectorPack) HasVersion() bool {
	if o != nil && o.Version != nil {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *IaasConnectorPack) SetVersion(v string) {
	o.Version = &v
}

// GetGuid returns the Guid field value if set, zero value otherwise.
func (o *IaasConnectorPack) GetGuid() IaasUcsdInfoRelationship {
	if o == nil || o.Guid == nil {
		var ret IaasUcsdInfoRelationship
		return ret
	}
	return *o.Guid
}

// GetGuidOk returns a tuple with the Guid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IaasConnectorPack) GetGuidOk() (*IaasUcsdInfoRelationship, bool) {
	if o == nil || o.Guid == nil {
		return nil, false
	}
	return o.Guid, true
}

// HasGuid returns a boolean if a field has been set.
func (o *IaasConnectorPack) HasGuid() bool {
	if o != nil && o.Guid != nil {
		return true
	}

	return false
}

// SetGuid gets a reference to the given IaasUcsdInfoRelationship and assigns it to the Guid field.
func (o *IaasConnectorPack) SetGuid(v IaasUcsdInfoRelationship) {
	o.Guid = &v
}

func (o IaasConnectorPack) MarshalJSON() ([]byte, error) {
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
	if o.CompleteVersion != nil {
		toSerialize["CompleteVersion"] = o.CompleteVersion
	}
	if o.DependencyNames != nil {
		toSerialize["DependencyNames"] = o.DependencyNames
	}
	if o.DownloadedVersion != nil {
		toSerialize["DownloadedVersion"] = o.DownloadedVersion
	}
	if o.Name != nil {
		toSerialize["Name"] = o.Name
	}
	if o.State != nil {
		toSerialize["State"] = o.State
	}
	if o.Version != nil {
		toSerialize["Version"] = o.Version
	}
	if o.Guid != nil {
		toSerialize["Guid"] = o.Guid
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *IaasConnectorPack) UnmarshalJSON(bytes []byte) (err error) {
	type IaasConnectorPackWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Complete version of the connector pack including build number.
		CompleteVersion *string  `json:"CompleteVersion,omitempty"`
		DependencyNames []string `json:"DependencyNames,omitempty"`
		// Version of the connector pack that is last downloaded successfully to UCSD.
		DownloadedVersion *string `json:"DownloadedVersion,omitempty"`
		// Name of the connector pack running on the UCSD.
		Name *string `json:"Name,omitempty"`
		// State of the connector pack whether it is enabled or disabled.
		State *string `json:"State,omitempty"`
		// Version of the connector pack.
		Version *string                   `json:"Version,omitempty"`
		Guid    *IaasUcsdInfoRelationship `json:"Guid,omitempty"`
	}

	varIaasConnectorPackWithoutEmbeddedStruct := IaasConnectorPackWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varIaasConnectorPackWithoutEmbeddedStruct)
	if err == nil {
		varIaasConnectorPack := _IaasConnectorPack{}
		varIaasConnectorPack.ClassId = varIaasConnectorPackWithoutEmbeddedStruct.ClassId
		varIaasConnectorPack.ObjectType = varIaasConnectorPackWithoutEmbeddedStruct.ObjectType
		varIaasConnectorPack.CompleteVersion = varIaasConnectorPackWithoutEmbeddedStruct.CompleteVersion
		varIaasConnectorPack.DependencyNames = varIaasConnectorPackWithoutEmbeddedStruct.DependencyNames
		varIaasConnectorPack.DownloadedVersion = varIaasConnectorPackWithoutEmbeddedStruct.DownloadedVersion
		varIaasConnectorPack.Name = varIaasConnectorPackWithoutEmbeddedStruct.Name
		varIaasConnectorPack.State = varIaasConnectorPackWithoutEmbeddedStruct.State
		varIaasConnectorPack.Version = varIaasConnectorPackWithoutEmbeddedStruct.Version
		varIaasConnectorPack.Guid = varIaasConnectorPackWithoutEmbeddedStruct.Guid
		*o = IaasConnectorPack(varIaasConnectorPack)
	} else {
		return err
	}

	varIaasConnectorPack := _IaasConnectorPack{}

	err = json.Unmarshal(bytes, &varIaasConnectorPack)
	if err == nil {
		o.MoBaseMo = varIaasConnectorPack.MoBaseMo
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "CompleteVersion")
		delete(additionalProperties, "DependencyNames")
		delete(additionalProperties, "DownloadedVersion")
		delete(additionalProperties, "Name")
		delete(additionalProperties, "State")
		delete(additionalProperties, "Version")
		delete(additionalProperties, "Guid")

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

type NullableIaasConnectorPack struct {
	value *IaasConnectorPack
	isSet bool
}

func (v NullableIaasConnectorPack) Get() *IaasConnectorPack {
	return v.value
}

func (v *NullableIaasConnectorPack) Set(val *IaasConnectorPack) {
	v.value = val
	v.isSet = true
}

func (v NullableIaasConnectorPack) IsSet() bool {
	return v.isSet
}

func (v *NullableIaasConnectorPack) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIaasConnectorPack(val *IaasConnectorPack) *NullableIaasConnectorPack {
	return &NullableIaasConnectorPack{value: val, isSet: true}
}

func (v NullableIaasConnectorPack) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIaasConnectorPack) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
