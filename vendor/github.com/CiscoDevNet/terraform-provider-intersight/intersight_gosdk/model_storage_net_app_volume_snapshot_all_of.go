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
)

// StorageNetAppVolumeSnapshotAllOf Definition of the list of properties defined in 'storage.NetAppVolumeSnapshot', excluding properties defined in parent classes.
type StorageNetAppVolumeSnapshotAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// UUID of the volume snapshot.
	Uuid                 *string                           `json:"Uuid,omitempty"`
	Array                *StorageNetAppClusterRelationship `json:"Array,omitempty"`
	StorageContainer     *StorageNetAppVolumeRelationship  `json:"StorageContainer,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _StorageNetAppVolumeSnapshotAllOf StorageNetAppVolumeSnapshotAllOf

// NewStorageNetAppVolumeSnapshotAllOf instantiates a new StorageNetAppVolumeSnapshotAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStorageNetAppVolumeSnapshotAllOf(classId string, objectType string) *StorageNetAppVolumeSnapshotAllOf {
	this := StorageNetAppVolumeSnapshotAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewStorageNetAppVolumeSnapshotAllOfWithDefaults instantiates a new StorageNetAppVolumeSnapshotAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStorageNetAppVolumeSnapshotAllOfWithDefaults() *StorageNetAppVolumeSnapshotAllOf {
	this := StorageNetAppVolumeSnapshotAllOf{}
	var classId string = "storage.NetAppVolumeSnapshot"
	this.ClassId = classId
	var objectType string = "storage.NetAppVolumeSnapshot"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *StorageNetAppVolumeSnapshotAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *StorageNetAppVolumeSnapshotAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *StorageNetAppVolumeSnapshotAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *StorageNetAppVolumeSnapshotAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *StorageNetAppVolumeSnapshotAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *StorageNetAppVolumeSnapshotAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetUuid returns the Uuid field value if set, zero value otherwise.
func (o *StorageNetAppVolumeSnapshotAllOf) GetUuid() string {
	if o == nil || o.Uuid == nil {
		var ret string
		return ret
	}
	return *o.Uuid
}

// GetUuidOk returns a tuple with the Uuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNetAppVolumeSnapshotAllOf) GetUuidOk() (*string, bool) {
	if o == nil || o.Uuid == nil {
		return nil, false
	}
	return o.Uuid, true
}

// HasUuid returns a boolean if a field has been set.
func (o *StorageNetAppVolumeSnapshotAllOf) HasUuid() bool {
	if o != nil && o.Uuid != nil {
		return true
	}

	return false
}

// SetUuid gets a reference to the given string and assigns it to the Uuid field.
func (o *StorageNetAppVolumeSnapshotAllOf) SetUuid(v string) {
	o.Uuid = &v
}

// GetArray returns the Array field value if set, zero value otherwise.
func (o *StorageNetAppVolumeSnapshotAllOf) GetArray() StorageNetAppClusterRelationship {
	if o == nil || o.Array == nil {
		var ret StorageNetAppClusterRelationship
		return ret
	}
	return *o.Array
}

// GetArrayOk returns a tuple with the Array field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNetAppVolumeSnapshotAllOf) GetArrayOk() (*StorageNetAppClusterRelationship, bool) {
	if o == nil || o.Array == nil {
		return nil, false
	}
	return o.Array, true
}

// HasArray returns a boolean if a field has been set.
func (o *StorageNetAppVolumeSnapshotAllOf) HasArray() bool {
	if o != nil && o.Array != nil {
		return true
	}

	return false
}

// SetArray gets a reference to the given StorageNetAppClusterRelationship and assigns it to the Array field.
func (o *StorageNetAppVolumeSnapshotAllOf) SetArray(v StorageNetAppClusterRelationship) {
	o.Array = &v
}

// GetStorageContainer returns the StorageContainer field value if set, zero value otherwise.
func (o *StorageNetAppVolumeSnapshotAllOf) GetStorageContainer() StorageNetAppVolumeRelationship {
	if o == nil || o.StorageContainer == nil {
		var ret StorageNetAppVolumeRelationship
		return ret
	}
	return *o.StorageContainer
}

// GetStorageContainerOk returns a tuple with the StorageContainer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNetAppVolumeSnapshotAllOf) GetStorageContainerOk() (*StorageNetAppVolumeRelationship, bool) {
	if o == nil || o.StorageContainer == nil {
		return nil, false
	}
	return o.StorageContainer, true
}

// HasStorageContainer returns a boolean if a field has been set.
func (o *StorageNetAppVolumeSnapshotAllOf) HasStorageContainer() bool {
	if o != nil && o.StorageContainer != nil {
		return true
	}

	return false
}

// SetStorageContainer gets a reference to the given StorageNetAppVolumeRelationship and assigns it to the StorageContainer field.
func (o *StorageNetAppVolumeSnapshotAllOf) SetStorageContainer(v StorageNetAppVolumeRelationship) {
	o.StorageContainer = &v
}

func (o StorageNetAppVolumeSnapshotAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Uuid != nil {
		toSerialize["Uuid"] = o.Uuid
	}
	if o.Array != nil {
		toSerialize["Array"] = o.Array
	}
	if o.StorageContainer != nil {
		toSerialize["StorageContainer"] = o.StorageContainer
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *StorageNetAppVolumeSnapshotAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varStorageNetAppVolumeSnapshotAllOf := _StorageNetAppVolumeSnapshotAllOf{}

	if err = json.Unmarshal(bytes, &varStorageNetAppVolumeSnapshotAllOf); err == nil {
		*o = StorageNetAppVolumeSnapshotAllOf(varStorageNetAppVolumeSnapshotAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Uuid")
		delete(additionalProperties, "Array")
		delete(additionalProperties, "StorageContainer")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableStorageNetAppVolumeSnapshotAllOf struct {
	value *StorageNetAppVolumeSnapshotAllOf
	isSet bool
}

func (v NullableStorageNetAppVolumeSnapshotAllOf) Get() *StorageNetAppVolumeSnapshotAllOf {
	return v.value
}

func (v *NullableStorageNetAppVolumeSnapshotAllOf) Set(val *StorageNetAppVolumeSnapshotAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableStorageNetAppVolumeSnapshotAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableStorageNetAppVolumeSnapshotAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStorageNetAppVolumeSnapshotAllOf(val *StorageNetAppVolumeSnapshotAllOf) *NullableStorageNetAppVolumeSnapshotAllOf {
	return &NullableStorageNetAppVolumeSnapshotAllOf{value: val, isSet: true}
}

func (v NullableStorageNetAppVolumeSnapshotAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStorageNetAppVolumeSnapshotAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
