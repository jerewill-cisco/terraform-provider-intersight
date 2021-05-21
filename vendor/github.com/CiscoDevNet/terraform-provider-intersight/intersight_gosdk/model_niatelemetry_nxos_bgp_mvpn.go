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

// NiatelemetryNxosBgpMvpn Stores data about BGP MVPN.
type NiatelemetryNxosBgpMvpn struct {
	MoBaseComplexType
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Return count of BGP MVPN table capable peers.
	CapablePeers *int64 `json:"CapablePeers,omitempty"`
	// Return count of BGP MVPN configured peers.
	ConfiguredPeers *int64 `json:"ConfiguredPeers,omitempty"`
	// Return value of BGP MVPN memory used.
	MemoryUsed *int64 `json:"MemoryUsed,omitempty"`
	// Return value of BGP MVPN cluster list.
	NumberOfClusterLists *int64 `json:"NumberOfClusterLists,omitempty"`
	// Return count of BGP MVPN communities.
	NumberOfCommunities *int64 `json:"NumberOfCommunities,omitempty"`
	// Return value of BGP MVPN table version.
	TableVersion *int64 `json:"TableVersion,omitempty"`
	// Return count of BGP MVPN networks.
	TotalNetworks *int64 `json:"TotalNetworks,omitempty"`
	// Return count of BGP MVPN paths.
	TotalPaths           *int64 `json:"TotalPaths,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _NiatelemetryNxosBgpMvpn NiatelemetryNxosBgpMvpn

// NewNiatelemetryNxosBgpMvpn instantiates a new NiatelemetryNxosBgpMvpn object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNiatelemetryNxosBgpMvpn(classId string, objectType string) *NiatelemetryNxosBgpMvpn {
	this := NiatelemetryNxosBgpMvpn{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewNiatelemetryNxosBgpMvpnWithDefaults instantiates a new NiatelemetryNxosBgpMvpn object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNiatelemetryNxosBgpMvpnWithDefaults() *NiatelemetryNxosBgpMvpn {
	this := NiatelemetryNxosBgpMvpn{}
	var classId string = "niatelemetry.NxosBgpMvpn"
	this.ClassId = classId
	var objectType string = "niatelemetry.NxosBgpMvpn"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *NiatelemetryNxosBgpMvpn) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *NiatelemetryNxosBgpMvpn) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *NiatelemetryNxosBgpMvpn) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *NiatelemetryNxosBgpMvpn) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *NiatelemetryNxosBgpMvpn) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *NiatelemetryNxosBgpMvpn) SetObjectType(v string) {
	o.ObjectType = v
}

// GetCapablePeers returns the CapablePeers field value if set, zero value otherwise.
func (o *NiatelemetryNxosBgpMvpn) GetCapablePeers() int64 {
	if o == nil || o.CapablePeers == nil {
		var ret int64
		return ret
	}
	return *o.CapablePeers
}

// GetCapablePeersOk returns a tuple with the CapablePeers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryNxosBgpMvpn) GetCapablePeersOk() (*int64, bool) {
	if o == nil || o.CapablePeers == nil {
		return nil, false
	}
	return o.CapablePeers, true
}

// HasCapablePeers returns a boolean if a field has been set.
func (o *NiatelemetryNxosBgpMvpn) HasCapablePeers() bool {
	if o != nil && o.CapablePeers != nil {
		return true
	}

	return false
}

// SetCapablePeers gets a reference to the given int64 and assigns it to the CapablePeers field.
func (o *NiatelemetryNxosBgpMvpn) SetCapablePeers(v int64) {
	o.CapablePeers = &v
}

// GetConfiguredPeers returns the ConfiguredPeers field value if set, zero value otherwise.
func (o *NiatelemetryNxosBgpMvpn) GetConfiguredPeers() int64 {
	if o == nil || o.ConfiguredPeers == nil {
		var ret int64
		return ret
	}
	return *o.ConfiguredPeers
}

// GetConfiguredPeersOk returns a tuple with the ConfiguredPeers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryNxosBgpMvpn) GetConfiguredPeersOk() (*int64, bool) {
	if o == nil || o.ConfiguredPeers == nil {
		return nil, false
	}
	return o.ConfiguredPeers, true
}

// HasConfiguredPeers returns a boolean if a field has been set.
func (o *NiatelemetryNxosBgpMvpn) HasConfiguredPeers() bool {
	if o != nil && o.ConfiguredPeers != nil {
		return true
	}

	return false
}

// SetConfiguredPeers gets a reference to the given int64 and assigns it to the ConfiguredPeers field.
func (o *NiatelemetryNxosBgpMvpn) SetConfiguredPeers(v int64) {
	o.ConfiguredPeers = &v
}

// GetMemoryUsed returns the MemoryUsed field value if set, zero value otherwise.
func (o *NiatelemetryNxosBgpMvpn) GetMemoryUsed() int64 {
	if o == nil || o.MemoryUsed == nil {
		var ret int64
		return ret
	}
	return *o.MemoryUsed
}

// GetMemoryUsedOk returns a tuple with the MemoryUsed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryNxosBgpMvpn) GetMemoryUsedOk() (*int64, bool) {
	if o == nil || o.MemoryUsed == nil {
		return nil, false
	}
	return o.MemoryUsed, true
}

// HasMemoryUsed returns a boolean if a field has been set.
func (o *NiatelemetryNxosBgpMvpn) HasMemoryUsed() bool {
	if o != nil && o.MemoryUsed != nil {
		return true
	}

	return false
}

// SetMemoryUsed gets a reference to the given int64 and assigns it to the MemoryUsed field.
func (o *NiatelemetryNxosBgpMvpn) SetMemoryUsed(v int64) {
	o.MemoryUsed = &v
}

// GetNumberOfClusterLists returns the NumberOfClusterLists field value if set, zero value otherwise.
func (o *NiatelemetryNxosBgpMvpn) GetNumberOfClusterLists() int64 {
	if o == nil || o.NumberOfClusterLists == nil {
		var ret int64
		return ret
	}
	return *o.NumberOfClusterLists
}

// GetNumberOfClusterListsOk returns a tuple with the NumberOfClusterLists field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryNxosBgpMvpn) GetNumberOfClusterListsOk() (*int64, bool) {
	if o == nil || o.NumberOfClusterLists == nil {
		return nil, false
	}
	return o.NumberOfClusterLists, true
}

// HasNumberOfClusterLists returns a boolean if a field has been set.
func (o *NiatelemetryNxosBgpMvpn) HasNumberOfClusterLists() bool {
	if o != nil && o.NumberOfClusterLists != nil {
		return true
	}

	return false
}

// SetNumberOfClusterLists gets a reference to the given int64 and assigns it to the NumberOfClusterLists field.
func (o *NiatelemetryNxosBgpMvpn) SetNumberOfClusterLists(v int64) {
	o.NumberOfClusterLists = &v
}

// GetNumberOfCommunities returns the NumberOfCommunities field value if set, zero value otherwise.
func (o *NiatelemetryNxosBgpMvpn) GetNumberOfCommunities() int64 {
	if o == nil || o.NumberOfCommunities == nil {
		var ret int64
		return ret
	}
	return *o.NumberOfCommunities
}

// GetNumberOfCommunitiesOk returns a tuple with the NumberOfCommunities field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryNxosBgpMvpn) GetNumberOfCommunitiesOk() (*int64, bool) {
	if o == nil || o.NumberOfCommunities == nil {
		return nil, false
	}
	return o.NumberOfCommunities, true
}

// HasNumberOfCommunities returns a boolean if a field has been set.
func (o *NiatelemetryNxosBgpMvpn) HasNumberOfCommunities() bool {
	if o != nil && o.NumberOfCommunities != nil {
		return true
	}

	return false
}

// SetNumberOfCommunities gets a reference to the given int64 and assigns it to the NumberOfCommunities field.
func (o *NiatelemetryNxosBgpMvpn) SetNumberOfCommunities(v int64) {
	o.NumberOfCommunities = &v
}

// GetTableVersion returns the TableVersion field value if set, zero value otherwise.
func (o *NiatelemetryNxosBgpMvpn) GetTableVersion() int64 {
	if o == nil || o.TableVersion == nil {
		var ret int64
		return ret
	}
	return *o.TableVersion
}

// GetTableVersionOk returns a tuple with the TableVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryNxosBgpMvpn) GetTableVersionOk() (*int64, bool) {
	if o == nil || o.TableVersion == nil {
		return nil, false
	}
	return o.TableVersion, true
}

// HasTableVersion returns a boolean if a field has been set.
func (o *NiatelemetryNxosBgpMvpn) HasTableVersion() bool {
	if o != nil && o.TableVersion != nil {
		return true
	}

	return false
}

// SetTableVersion gets a reference to the given int64 and assigns it to the TableVersion field.
func (o *NiatelemetryNxosBgpMvpn) SetTableVersion(v int64) {
	o.TableVersion = &v
}

// GetTotalNetworks returns the TotalNetworks field value if set, zero value otherwise.
func (o *NiatelemetryNxosBgpMvpn) GetTotalNetworks() int64 {
	if o == nil || o.TotalNetworks == nil {
		var ret int64
		return ret
	}
	return *o.TotalNetworks
}

// GetTotalNetworksOk returns a tuple with the TotalNetworks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryNxosBgpMvpn) GetTotalNetworksOk() (*int64, bool) {
	if o == nil || o.TotalNetworks == nil {
		return nil, false
	}
	return o.TotalNetworks, true
}

// HasTotalNetworks returns a boolean if a field has been set.
func (o *NiatelemetryNxosBgpMvpn) HasTotalNetworks() bool {
	if o != nil && o.TotalNetworks != nil {
		return true
	}

	return false
}

// SetTotalNetworks gets a reference to the given int64 and assigns it to the TotalNetworks field.
func (o *NiatelemetryNxosBgpMvpn) SetTotalNetworks(v int64) {
	o.TotalNetworks = &v
}

// GetTotalPaths returns the TotalPaths field value if set, zero value otherwise.
func (o *NiatelemetryNxosBgpMvpn) GetTotalPaths() int64 {
	if o == nil || o.TotalPaths == nil {
		var ret int64
		return ret
	}
	return *o.TotalPaths
}

// GetTotalPathsOk returns a tuple with the TotalPaths field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryNxosBgpMvpn) GetTotalPathsOk() (*int64, bool) {
	if o == nil || o.TotalPaths == nil {
		return nil, false
	}
	return o.TotalPaths, true
}

// HasTotalPaths returns a boolean if a field has been set.
func (o *NiatelemetryNxosBgpMvpn) HasTotalPaths() bool {
	if o != nil && o.TotalPaths != nil {
		return true
	}

	return false
}

// SetTotalPaths gets a reference to the given int64 and assigns it to the TotalPaths field.
func (o *NiatelemetryNxosBgpMvpn) SetTotalPaths(v int64) {
	o.TotalPaths = &v
}

func (o NiatelemetryNxosBgpMvpn) MarshalJSON() ([]byte, error) {
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
	if o.CapablePeers != nil {
		toSerialize["CapablePeers"] = o.CapablePeers
	}
	if o.ConfiguredPeers != nil {
		toSerialize["ConfiguredPeers"] = o.ConfiguredPeers
	}
	if o.MemoryUsed != nil {
		toSerialize["MemoryUsed"] = o.MemoryUsed
	}
	if o.NumberOfClusterLists != nil {
		toSerialize["NumberOfClusterLists"] = o.NumberOfClusterLists
	}
	if o.NumberOfCommunities != nil {
		toSerialize["NumberOfCommunities"] = o.NumberOfCommunities
	}
	if o.TableVersion != nil {
		toSerialize["TableVersion"] = o.TableVersion
	}
	if o.TotalNetworks != nil {
		toSerialize["TotalNetworks"] = o.TotalNetworks
	}
	if o.TotalPaths != nil {
		toSerialize["TotalPaths"] = o.TotalPaths
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *NiatelemetryNxosBgpMvpn) UnmarshalJSON(bytes []byte) (err error) {
	type NiatelemetryNxosBgpMvpnWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Return count of BGP MVPN table capable peers.
		CapablePeers *int64 `json:"CapablePeers,omitempty"`
		// Return count of BGP MVPN configured peers.
		ConfiguredPeers *int64 `json:"ConfiguredPeers,omitempty"`
		// Return value of BGP MVPN memory used.
		MemoryUsed *int64 `json:"MemoryUsed,omitempty"`
		// Return value of BGP MVPN cluster list.
		NumberOfClusterLists *int64 `json:"NumberOfClusterLists,omitempty"`
		// Return count of BGP MVPN communities.
		NumberOfCommunities *int64 `json:"NumberOfCommunities,omitempty"`
		// Return value of BGP MVPN table version.
		TableVersion *int64 `json:"TableVersion,omitempty"`
		// Return count of BGP MVPN networks.
		TotalNetworks *int64 `json:"TotalNetworks,omitempty"`
		// Return count of BGP MVPN paths.
		TotalPaths *int64 `json:"TotalPaths,omitempty"`
	}

	varNiatelemetryNxosBgpMvpnWithoutEmbeddedStruct := NiatelemetryNxosBgpMvpnWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varNiatelemetryNxosBgpMvpnWithoutEmbeddedStruct)
	if err == nil {
		varNiatelemetryNxosBgpMvpn := _NiatelemetryNxosBgpMvpn{}
		varNiatelemetryNxosBgpMvpn.ClassId = varNiatelemetryNxosBgpMvpnWithoutEmbeddedStruct.ClassId
		varNiatelemetryNxosBgpMvpn.ObjectType = varNiatelemetryNxosBgpMvpnWithoutEmbeddedStruct.ObjectType
		varNiatelemetryNxosBgpMvpn.CapablePeers = varNiatelemetryNxosBgpMvpnWithoutEmbeddedStruct.CapablePeers
		varNiatelemetryNxosBgpMvpn.ConfiguredPeers = varNiatelemetryNxosBgpMvpnWithoutEmbeddedStruct.ConfiguredPeers
		varNiatelemetryNxosBgpMvpn.MemoryUsed = varNiatelemetryNxosBgpMvpnWithoutEmbeddedStruct.MemoryUsed
		varNiatelemetryNxosBgpMvpn.NumberOfClusterLists = varNiatelemetryNxosBgpMvpnWithoutEmbeddedStruct.NumberOfClusterLists
		varNiatelemetryNxosBgpMvpn.NumberOfCommunities = varNiatelemetryNxosBgpMvpnWithoutEmbeddedStruct.NumberOfCommunities
		varNiatelemetryNxosBgpMvpn.TableVersion = varNiatelemetryNxosBgpMvpnWithoutEmbeddedStruct.TableVersion
		varNiatelemetryNxosBgpMvpn.TotalNetworks = varNiatelemetryNxosBgpMvpnWithoutEmbeddedStruct.TotalNetworks
		varNiatelemetryNxosBgpMvpn.TotalPaths = varNiatelemetryNxosBgpMvpnWithoutEmbeddedStruct.TotalPaths
		*o = NiatelemetryNxosBgpMvpn(varNiatelemetryNxosBgpMvpn)
	} else {
		return err
	}

	varNiatelemetryNxosBgpMvpn := _NiatelemetryNxosBgpMvpn{}

	err = json.Unmarshal(bytes, &varNiatelemetryNxosBgpMvpn)
	if err == nil {
		o.MoBaseComplexType = varNiatelemetryNxosBgpMvpn.MoBaseComplexType
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "CapablePeers")
		delete(additionalProperties, "ConfiguredPeers")
		delete(additionalProperties, "MemoryUsed")
		delete(additionalProperties, "NumberOfClusterLists")
		delete(additionalProperties, "NumberOfCommunities")
		delete(additionalProperties, "TableVersion")
		delete(additionalProperties, "TotalNetworks")
		delete(additionalProperties, "TotalPaths")

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

type NullableNiatelemetryNxosBgpMvpn struct {
	value *NiatelemetryNxosBgpMvpn
	isSet bool
}

func (v NullableNiatelemetryNxosBgpMvpn) Get() *NiatelemetryNxosBgpMvpn {
	return v.value
}

func (v *NullableNiatelemetryNxosBgpMvpn) Set(val *NiatelemetryNxosBgpMvpn) {
	v.value = val
	v.isSet = true
}

func (v NullableNiatelemetryNxosBgpMvpn) IsSet() bool {
	return v.isSet
}

func (v *NullableNiatelemetryNxosBgpMvpn) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNiatelemetryNxosBgpMvpn(val *NiatelemetryNxosBgpMvpn) *NullableNiatelemetryNxosBgpMvpn {
	return &NullableNiatelemetryNxosBgpMvpn{value: val, isSet: true}
}

func (v NullableNiatelemetryNxosBgpMvpn) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNiatelemetryNxosBgpMvpn) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
