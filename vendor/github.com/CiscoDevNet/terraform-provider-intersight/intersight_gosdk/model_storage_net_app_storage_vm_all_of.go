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

// StorageNetAppStorageVmAllOf Definition of the list of properties defined in 'storage.NetAppStorageVm', excluding properties defined in parent classes.
type StorageNetAppStorageVmAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Status for Common Internet File System protocol ( CIFS ) allowed to run on Vservers.
	CifsEnabled *bool `json:"CifsEnabled,omitempty"`
	// Status for Fibre Channel Protocol ( FCP ) allowed to run on Vservers.
	FcpEnabled *bool `json:"FcpEnabled,omitempty"`
	// Status for iSCSI protocol allowed to run on Vservers.
	IscsiEnabled *bool `json:"IscsiEnabled,omitempty"`
	// Status for Network File System Protocol ( NFS ) allowed to run on  Vservers.
	NfsEnabled *bool `json:"NfsEnabled,omitempty"`
	// Status for NVME protocol allowed to run on Vservers.
	NvmeEnabled *bool                             `json:"NvmeEnabled,omitempty"`
	Array       *StorageNetAppClusterRelationship `json:"Array,omitempty"`
	// An array of relationships to storageNetAppAggregate resources.
	DiskPool             []StorageNetAppAggregateRelationship `json:"DiskPool,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _StorageNetAppStorageVmAllOf StorageNetAppStorageVmAllOf

// NewStorageNetAppStorageVmAllOf instantiates a new StorageNetAppStorageVmAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStorageNetAppStorageVmAllOf(classId string, objectType string) *StorageNetAppStorageVmAllOf {
	this := StorageNetAppStorageVmAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewStorageNetAppStorageVmAllOfWithDefaults instantiates a new StorageNetAppStorageVmAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStorageNetAppStorageVmAllOfWithDefaults() *StorageNetAppStorageVmAllOf {
	this := StorageNetAppStorageVmAllOf{}
	var classId string = "storage.NetAppStorageVm"
	this.ClassId = classId
	var objectType string = "storage.NetAppStorageVm"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *StorageNetAppStorageVmAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *StorageNetAppStorageVmAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *StorageNetAppStorageVmAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *StorageNetAppStorageVmAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *StorageNetAppStorageVmAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *StorageNetAppStorageVmAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetCifsEnabled returns the CifsEnabled field value if set, zero value otherwise.
func (o *StorageNetAppStorageVmAllOf) GetCifsEnabled() bool {
	if o == nil || o.CifsEnabled == nil {
		var ret bool
		return ret
	}
	return *o.CifsEnabled
}

// GetCifsEnabledOk returns a tuple with the CifsEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNetAppStorageVmAllOf) GetCifsEnabledOk() (*bool, bool) {
	if o == nil || o.CifsEnabled == nil {
		return nil, false
	}
	return o.CifsEnabled, true
}

// HasCifsEnabled returns a boolean if a field has been set.
func (o *StorageNetAppStorageVmAllOf) HasCifsEnabled() bool {
	if o != nil && o.CifsEnabled != nil {
		return true
	}

	return false
}

// SetCifsEnabled gets a reference to the given bool and assigns it to the CifsEnabled field.
func (o *StorageNetAppStorageVmAllOf) SetCifsEnabled(v bool) {
	o.CifsEnabled = &v
}

// GetFcpEnabled returns the FcpEnabled field value if set, zero value otherwise.
func (o *StorageNetAppStorageVmAllOf) GetFcpEnabled() bool {
	if o == nil || o.FcpEnabled == nil {
		var ret bool
		return ret
	}
	return *o.FcpEnabled
}

// GetFcpEnabledOk returns a tuple with the FcpEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNetAppStorageVmAllOf) GetFcpEnabledOk() (*bool, bool) {
	if o == nil || o.FcpEnabled == nil {
		return nil, false
	}
	return o.FcpEnabled, true
}

// HasFcpEnabled returns a boolean if a field has been set.
func (o *StorageNetAppStorageVmAllOf) HasFcpEnabled() bool {
	if o != nil && o.FcpEnabled != nil {
		return true
	}

	return false
}

// SetFcpEnabled gets a reference to the given bool and assigns it to the FcpEnabled field.
func (o *StorageNetAppStorageVmAllOf) SetFcpEnabled(v bool) {
	o.FcpEnabled = &v
}

// GetIscsiEnabled returns the IscsiEnabled field value if set, zero value otherwise.
func (o *StorageNetAppStorageVmAllOf) GetIscsiEnabled() bool {
	if o == nil || o.IscsiEnabled == nil {
		var ret bool
		return ret
	}
	return *o.IscsiEnabled
}

// GetIscsiEnabledOk returns a tuple with the IscsiEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNetAppStorageVmAllOf) GetIscsiEnabledOk() (*bool, bool) {
	if o == nil || o.IscsiEnabled == nil {
		return nil, false
	}
	return o.IscsiEnabled, true
}

// HasIscsiEnabled returns a boolean if a field has been set.
func (o *StorageNetAppStorageVmAllOf) HasIscsiEnabled() bool {
	if o != nil && o.IscsiEnabled != nil {
		return true
	}

	return false
}

// SetIscsiEnabled gets a reference to the given bool and assigns it to the IscsiEnabled field.
func (o *StorageNetAppStorageVmAllOf) SetIscsiEnabled(v bool) {
	o.IscsiEnabled = &v
}

// GetNfsEnabled returns the NfsEnabled field value if set, zero value otherwise.
func (o *StorageNetAppStorageVmAllOf) GetNfsEnabled() bool {
	if o == nil || o.NfsEnabled == nil {
		var ret bool
		return ret
	}
	return *o.NfsEnabled
}

// GetNfsEnabledOk returns a tuple with the NfsEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNetAppStorageVmAllOf) GetNfsEnabledOk() (*bool, bool) {
	if o == nil || o.NfsEnabled == nil {
		return nil, false
	}
	return o.NfsEnabled, true
}

// HasNfsEnabled returns a boolean if a field has been set.
func (o *StorageNetAppStorageVmAllOf) HasNfsEnabled() bool {
	if o != nil && o.NfsEnabled != nil {
		return true
	}

	return false
}

// SetNfsEnabled gets a reference to the given bool and assigns it to the NfsEnabled field.
func (o *StorageNetAppStorageVmAllOf) SetNfsEnabled(v bool) {
	o.NfsEnabled = &v
}

// GetNvmeEnabled returns the NvmeEnabled field value if set, zero value otherwise.
func (o *StorageNetAppStorageVmAllOf) GetNvmeEnabled() bool {
	if o == nil || o.NvmeEnabled == nil {
		var ret bool
		return ret
	}
	return *o.NvmeEnabled
}

// GetNvmeEnabledOk returns a tuple with the NvmeEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNetAppStorageVmAllOf) GetNvmeEnabledOk() (*bool, bool) {
	if o == nil || o.NvmeEnabled == nil {
		return nil, false
	}
	return o.NvmeEnabled, true
}

// HasNvmeEnabled returns a boolean if a field has been set.
func (o *StorageNetAppStorageVmAllOf) HasNvmeEnabled() bool {
	if o != nil && o.NvmeEnabled != nil {
		return true
	}

	return false
}

// SetNvmeEnabled gets a reference to the given bool and assigns it to the NvmeEnabled field.
func (o *StorageNetAppStorageVmAllOf) SetNvmeEnabled(v bool) {
	o.NvmeEnabled = &v
}

// GetArray returns the Array field value if set, zero value otherwise.
func (o *StorageNetAppStorageVmAllOf) GetArray() StorageNetAppClusterRelationship {
	if o == nil || o.Array == nil {
		var ret StorageNetAppClusterRelationship
		return ret
	}
	return *o.Array
}

// GetArrayOk returns a tuple with the Array field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNetAppStorageVmAllOf) GetArrayOk() (*StorageNetAppClusterRelationship, bool) {
	if o == nil || o.Array == nil {
		return nil, false
	}
	return o.Array, true
}

// HasArray returns a boolean if a field has been set.
func (o *StorageNetAppStorageVmAllOf) HasArray() bool {
	if o != nil && o.Array != nil {
		return true
	}

	return false
}

// SetArray gets a reference to the given StorageNetAppClusterRelationship and assigns it to the Array field.
func (o *StorageNetAppStorageVmAllOf) SetArray(v StorageNetAppClusterRelationship) {
	o.Array = &v
}

// GetDiskPool returns the DiskPool field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StorageNetAppStorageVmAllOf) GetDiskPool() []StorageNetAppAggregateRelationship {
	if o == nil {
		var ret []StorageNetAppAggregateRelationship
		return ret
	}
	return o.DiskPool
}

// GetDiskPoolOk returns a tuple with the DiskPool field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StorageNetAppStorageVmAllOf) GetDiskPoolOk() (*[]StorageNetAppAggregateRelationship, bool) {
	if o == nil || o.DiskPool == nil {
		return nil, false
	}
	return &o.DiskPool, true
}

// HasDiskPool returns a boolean if a field has been set.
func (o *StorageNetAppStorageVmAllOf) HasDiskPool() bool {
	if o != nil && o.DiskPool != nil {
		return true
	}

	return false
}

// SetDiskPool gets a reference to the given []StorageNetAppAggregateRelationship and assigns it to the DiskPool field.
func (o *StorageNetAppStorageVmAllOf) SetDiskPool(v []StorageNetAppAggregateRelationship) {
	o.DiskPool = v
}

func (o StorageNetAppStorageVmAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.CifsEnabled != nil {
		toSerialize["CifsEnabled"] = o.CifsEnabled
	}
	if o.FcpEnabled != nil {
		toSerialize["FcpEnabled"] = o.FcpEnabled
	}
	if o.IscsiEnabled != nil {
		toSerialize["IscsiEnabled"] = o.IscsiEnabled
	}
	if o.NfsEnabled != nil {
		toSerialize["NfsEnabled"] = o.NfsEnabled
	}
	if o.NvmeEnabled != nil {
		toSerialize["NvmeEnabled"] = o.NvmeEnabled
	}
	if o.Array != nil {
		toSerialize["Array"] = o.Array
	}
	if o.DiskPool != nil {
		toSerialize["DiskPool"] = o.DiskPool
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *StorageNetAppStorageVmAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varStorageNetAppStorageVmAllOf := _StorageNetAppStorageVmAllOf{}

	if err = json.Unmarshal(bytes, &varStorageNetAppStorageVmAllOf); err == nil {
		*o = StorageNetAppStorageVmAllOf(varStorageNetAppStorageVmAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "CifsEnabled")
		delete(additionalProperties, "FcpEnabled")
		delete(additionalProperties, "IscsiEnabled")
		delete(additionalProperties, "NfsEnabled")
		delete(additionalProperties, "NvmeEnabled")
		delete(additionalProperties, "Array")
		delete(additionalProperties, "DiskPool")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableStorageNetAppStorageVmAllOf struct {
	value *StorageNetAppStorageVmAllOf
	isSet bool
}

func (v NullableStorageNetAppStorageVmAllOf) Get() *StorageNetAppStorageVmAllOf {
	return v.value
}

func (v *NullableStorageNetAppStorageVmAllOf) Set(val *StorageNetAppStorageVmAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableStorageNetAppStorageVmAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableStorageNetAppStorageVmAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStorageNetAppStorageVmAllOf(val *StorageNetAppStorageVmAllOf) *NullableStorageNetAppStorageVmAllOf {
	return &NullableStorageNetAppStorageVmAllOf{value: val, isSet: true}
}

func (v NullableStorageNetAppStorageVmAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStorageNetAppStorageVmAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
