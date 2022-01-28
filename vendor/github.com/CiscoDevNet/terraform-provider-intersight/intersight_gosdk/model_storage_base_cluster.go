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

// StorageBaseCluster Common attributes of a storage cluster.
type StorageBaseCluster struct {
	InfraBaseCluster
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property. The enum values provides the list of concrete types that can be instantiated from this abstract type.
	ObjectType string `json:"ObjectType"`
	// The storage capacity in this cluster.
	StorageCapacity      *int64                          `json:"StorageCapacity,omitempty"`
	ParentCluster        *ComputeBaseClusterRelationship `json:"ParentCluster,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _StorageBaseCluster StorageBaseCluster

// NewStorageBaseCluster instantiates a new StorageBaseCluster object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStorageBaseCluster(classId string, objectType string) *StorageBaseCluster {
	this := StorageBaseCluster{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewStorageBaseClusterWithDefaults instantiates a new StorageBaseCluster object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStorageBaseClusterWithDefaults() *StorageBaseCluster {
	this := StorageBaseCluster{}
	var classId string = "hyperflex.Cluster"
	this.ClassId = classId
	var objectType string = "hyperflex.Cluster"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *StorageBaseCluster) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *StorageBaseCluster) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *StorageBaseCluster) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *StorageBaseCluster) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *StorageBaseCluster) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *StorageBaseCluster) SetObjectType(v string) {
	o.ObjectType = v
}

// GetStorageCapacity returns the StorageCapacity field value if set, zero value otherwise.
func (o *StorageBaseCluster) GetStorageCapacity() int64 {
	if o == nil || o.StorageCapacity == nil {
		var ret int64
		return ret
	}
	return *o.StorageCapacity
}

// GetStorageCapacityOk returns a tuple with the StorageCapacity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageBaseCluster) GetStorageCapacityOk() (*int64, bool) {
	if o == nil || o.StorageCapacity == nil {
		return nil, false
	}
	return o.StorageCapacity, true
}

// HasStorageCapacity returns a boolean if a field has been set.
func (o *StorageBaseCluster) HasStorageCapacity() bool {
	if o != nil && o.StorageCapacity != nil {
		return true
	}

	return false
}

// SetStorageCapacity gets a reference to the given int64 and assigns it to the StorageCapacity field.
func (o *StorageBaseCluster) SetStorageCapacity(v int64) {
	o.StorageCapacity = &v
}

// GetParentCluster returns the ParentCluster field value if set, zero value otherwise.
func (o *StorageBaseCluster) GetParentCluster() ComputeBaseClusterRelationship {
	if o == nil || o.ParentCluster == nil {
		var ret ComputeBaseClusterRelationship
		return ret
	}
	return *o.ParentCluster
}

// GetParentClusterOk returns a tuple with the ParentCluster field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageBaseCluster) GetParentClusterOk() (*ComputeBaseClusterRelationship, bool) {
	if o == nil || o.ParentCluster == nil {
		return nil, false
	}
	return o.ParentCluster, true
}

// HasParentCluster returns a boolean if a field has been set.
func (o *StorageBaseCluster) HasParentCluster() bool {
	if o != nil && o.ParentCluster != nil {
		return true
	}

	return false
}

// SetParentCluster gets a reference to the given ComputeBaseClusterRelationship and assigns it to the ParentCluster field.
func (o *StorageBaseCluster) SetParentCluster(v ComputeBaseClusterRelationship) {
	o.ParentCluster = &v
}

func (o StorageBaseCluster) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedInfraBaseCluster, errInfraBaseCluster := json.Marshal(o.InfraBaseCluster)
	if errInfraBaseCluster != nil {
		return []byte{}, errInfraBaseCluster
	}
	errInfraBaseCluster = json.Unmarshal([]byte(serializedInfraBaseCluster), &toSerialize)
	if errInfraBaseCluster != nil {
		return []byte{}, errInfraBaseCluster
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.StorageCapacity != nil {
		toSerialize["StorageCapacity"] = o.StorageCapacity
	}
	if o.ParentCluster != nil {
		toSerialize["ParentCluster"] = o.ParentCluster
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *StorageBaseCluster) UnmarshalJSON(bytes []byte) (err error) {
	type StorageBaseClusterWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property. The enum values provides the list of concrete types that can be instantiated from this abstract type.
		ObjectType string `json:"ObjectType"`
		// The storage capacity in this cluster.
		StorageCapacity *int64                          `json:"StorageCapacity,omitempty"`
		ParentCluster   *ComputeBaseClusterRelationship `json:"ParentCluster,omitempty"`
	}

	varStorageBaseClusterWithoutEmbeddedStruct := StorageBaseClusterWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varStorageBaseClusterWithoutEmbeddedStruct)
	if err == nil {
		varStorageBaseCluster := _StorageBaseCluster{}
		varStorageBaseCluster.ClassId = varStorageBaseClusterWithoutEmbeddedStruct.ClassId
		varStorageBaseCluster.ObjectType = varStorageBaseClusterWithoutEmbeddedStruct.ObjectType
		varStorageBaseCluster.StorageCapacity = varStorageBaseClusterWithoutEmbeddedStruct.StorageCapacity
		varStorageBaseCluster.ParentCluster = varStorageBaseClusterWithoutEmbeddedStruct.ParentCluster
		*o = StorageBaseCluster(varStorageBaseCluster)
	} else {
		return err
	}

	varStorageBaseCluster := _StorageBaseCluster{}

	err = json.Unmarshal(bytes, &varStorageBaseCluster)
	if err == nil {
		o.InfraBaseCluster = varStorageBaseCluster.InfraBaseCluster
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "StorageCapacity")
		delete(additionalProperties, "ParentCluster")

		// remove fields from embedded structs
		reflectInfraBaseCluster := reflect.ValueOf(o.InfraBaseCluster)
		for i := 0; i < reflectInfraBaseCluster.Type().NumField(); i++ {
			t := reflectInfraBaseCluster.Type().Field(i)

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

type NullableStorageBaseCluster struct {
	value *StorageBaseCluster
	isSet bool
}

func (v NullableStorageBaseCluster) Get() *StorageBaseCluster {
	return v.value
}

func (v *NullableStorageBaseCluster) Set(val *StorageBaseCluster) {
	v.value = val
	v.isSet = true
}

func (v NullableStorageBaseCluster) IsSet() bool {
	return v.isSet
}

func (v *NullableStorageBaseCluster) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStorageBaseCluster(val *StorageBaseCluster) *NullableStorageBaseCluster {
	return &NullableStorageBaseCluster{value: val, isSet: true}
}

func (v NullableStorageBaseCluster) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStorageBaseCluster) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
