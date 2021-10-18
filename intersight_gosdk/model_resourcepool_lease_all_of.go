/*
Cisco Intersight

Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document.

API version: 1.0.9-4663
Contact: intersight@cisco.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package intersight

import (
	"encoding/json"
)

// ResourcepoolLeaseAllOf Definition of the list of properties defined in 'resourcepool.Lease', excluding properties defined in parent classes.
type ResourcepoolLeaseAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string             `json:"ObjectType"`
	Condition  []ResourceSelector `json:"Condition,omitempty"`
	// Lease opertion applied for the feature.
	Feature         *string                             `json:"Feature,omitempty"`
	LeaseParameters NullableResourcepoolLeaseParameters `json:"LeaseParameters,omitempty"`
	Resource        *MoBaseMo                           `json:"Resource,omitempty"`
	// The type of the resource present in the pool, example 'server' its combination of RackUnit and Blade. * `None` - The resource cannot consider for Resource Pool. * `Server` - Resource Pool holds the server kind of resources, example - RackServer, Blade.
	ResourceType         *string                                `json:"ResourceType,omitempty"`
	AssignedToEntity     *MoBaseMoRelationship                  `json:"AssignedToEntity,omitempty"`
	LeasedResource       *ResourcepoolLeaseResourceRelationship `json:"LeasedResource,omitempty"`
	Pool                 *ResourcepoolPoolRelationship          `json:"Pool,omitempty"`
	PoolMember           *ResourcepoolPoolMemberRelationship    `json:"PoolMember,omitempty"`
	Universe             *ResourcepoolUniverseRelationship      `json:"Universe,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ResourcepoolLeaseAllOf ResourcepoolLeaseAllOf

// NewResourcepoolLeaseAllOf instantiates a new ResourcepoolLeaseAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResourcepoolLeaseAllOf(classId string, objectType string) *ResourcepoolLeaseAllOf {
	this := ResourcepoolLeaseAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	var resourceType string = "None"
	this.ResourceType = &resourceType
	return &this
}

// NewResourcepoolLeaseAllOfWithDefaults instantiates a new ResourcepoolLeaseAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResourcepoolLeaseAllOfWithDefaults() *ResourcepoolLeaseAllOf {
	this := ResourcepoolLeaseAllOf{}
	var classId string = "resourcepool.Lease"
	this.ClassId = classId
	var objectType string = "resourcepool.Lease"
	this.ObjectType = objectType
	var resourceType string = "None"
	this.ResourceType = &resourceType
	return &this
}

// GetClassId returns the ClassId field value
func (o *ResourcepoolLeaseAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *ResourcepoolLeaseAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *ResourcepoolLeaseAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *ResourcepoolLeaseAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *ResourcepoolLeaseAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *ResourcepoolLeaseAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetCondition returns the Condition field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ResourcepoolLeaseAllOf) GetCondition() []ResourceSelector {
	if o == nil {
		var ret []ResourceSelector
		return ret
	}
	return o.Condition
}

// GetConditionOk returns a tuple with the Condition field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ResourcepoolLeaseAllOf) GetConditionOk() (*[]ResourceSelector, bool) {
	if o == nil || o.Condition == nil {
		return nil, false
	}
	return &o.Condition, true
}

// HasCondition returns a boolean if a field has been set.
func (o *ResourcepoolLeaseAllOf) HasCondition() bool {
	if o != nil && o.Condition != nil {
		return true
	}

	return false
}

// SetCondition gets a reference to the given []ResourceSelector and assigns it to the Condition field.
func (o *ResourcepoolLeaseAllOf) SetCondition(v []ResourceSelector) {
	o.Condition = v
}

// GetFeature returns the Feature field value if set, zero value otherwise.
func (o *ResourcepoolLeaseAllOf) GetFeature() string {
	if o == nil || o.Feature == nil {
		var ret string
		return ret
	}
	return *o.Feature
}

// GetFeatureOk returns a tuple with the Feature field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourcepoolLeaseAllOf) GetFeatureOk() (*string, bool) {
	if o == nil || o.Feature == nil {
		return nil, false
	}
	return o.Feature, true
}

// HasFeature returns a boolean if a field has been set.
func (o *ResourcepoolLeaseAllOf) HasFeature() bool {
	if o != nil && o.Feature != nil {
		return true
	}

	return false
}

// SetFeature gets a reference to the given string and assigns it to the Feature field.
func (o *ResourcepoolLeaseAllOf) SetFeature(v string) {
	o.Feature = &v
}

// GetLeaseParameters returns the LeaseParameters field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ResourcepoolLeaseAllOf) GetLeaseParameters() ResourcepoolLeaseParameters {
	if o == nil || o.LeaseParameters.Get() == nil {
		var ret ResourcepoolLeaseParameters
		return ret
	}
	return *o.LeaseParameters.Get()
}

// GetLeaseParametersOk returns a tuple with the LeaseParameters field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ResourcepoolLeaseAllOf) GetLeaseParametersOk() (*ResourcepoolLeaseParameters, bool) {
	if o == nil {
		return nil, false
	}
	return o.LeaseParameters.Get(), o.LeaseParameters.IsSet()
}

// HasLeaseParameters returns a boolean if a field has been set.
func (o *ResourcepoolLeaseAllOf) HasLeaseParameters() bool {
	if o != nil && o.LeaseParameters.IsSet() {
		return true
	}

	return false
}

// SetLeaseParameters gets a reference to the given NullableResourcepoolLeaseParameters and assigns it to the LeaseParameters field.
func (o *ResourcepoolLeaseAllOf) SetLeaseParameters(v ResourcepoolLeaseParameters) {
	o.LeaseParameters.Set(&v)
}

// SetLeaseParametersNil sets the value for LeaseParameters to be an explicit nil
func (o *ResourcepoolLeaseAllOf) SetLeaseParametersNil() {
	o.LeaseParameters.Set(nil)
}

// UnsetLeaseParameters ensures that no value is present for LeaseParameters, not even an explicit nil
func (o *ResourcepoolLeaseAllOf) UnsetLeaseParameters() {
	o.LeaseParameters.Unset()
}

// GetResource returns the Resource field value if set, zero value otherwise.
func (o *ResourcepoolLeaseAllOf) GetResource() MoBaseMo {
	if o == nil || o.Resource == nil {
		var ret MoBaseMo
		return ret
	}
	return *o.Resource
}

// GetResourceOk returns a tuple with the Resource field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourcepoolLeaseAllOf) GetResourceOk() (*MoBaseMo, bool) {
	if o == nil || o.Resource == nil {
		return nil, false
	}
	return o.Resource, true
}

// HasResource returns a boolean if a field has been set.
func (o *ResourcepoolLeaseAllOf) HasResource() bool {
	if o != nil && o.Resource != nil {
		return true
	}

	return false
}

// SetResource gets a reference to the given MoBaseMo and assigns it to the Resource field.
func (o *ResourcepoolLeaseAllOf) SetResource(v MoBaseMo) {
	o.Resource = &v
}

// GetResourceType returns the ResourceType field value if set, zero value otherwise.
func (o *ResourcepoolLeaseAllOf) GetResourceType() string {
	if o == nil || o.ResourceType == nil {
		var ret string
		return ret
	}
	return *o.ResourceType
}

// GetResourceTypeOk returns a tuple with the ResourceType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourcepoolLeaseAllOf) GetResourceTypeOk() (*string, bool) {
	if o == nil || o.ResourceType == nil {
		return nil, false
	}
	return o.ResourceType, true
}

// HasResourceType returns a boolean if a field has been set.
func (o *ResourcepoolLeaseAllOf) HasResourceType() bool {
	if o != nil && o.ResourceType != nil {
		return true
	}

	return false
}

// SetResourceType gets a reference to the given string and assigns it to the ResourceType field.
func (o *ResourcepoolLeaseAllOf) SetResourceType(v string) {
	o.ResourceType = &v
}

// GetAssignedToEntity returns the AssignedToEntity field value if set, zero value otherwise.
func (o *ResourcepoolLeaseAllOf) GetAssignedToEntity() MoBaseMoRelationship {
	if o == nil || o.AssignedToEntity == nil {
		var ret MoBaseMoRelationship
		return ret
	}
	return *o.AssignedToEntity
}

// GetAssignedToEntityOk returns a tuple with the AssignedToEntity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourcepoolLeaseAllOf) GetAssignedToEntityOk() (*MoBaseMoRelationship, bool) {
	if o == nil || o.AssignedToEntity == nil {
		return nil, false
	}
	return o.AssignedToEntity, true
}

// HasAssignedToEntity returns a boolean if a field has been set.
func (o *ResourcepoolLeaseAllOf) HasAssignedToEntity() bool {
	if o != nil && o.AssignedToEntity != nil {
		return true
	}

	return false
}

// SetAssignedToEntity gets a reference to the given MoBaseMoRelationship and assigns it to the AssignedToEntity field.
func (o *ResourcepoolLeaseAllOf) SetAssignedToEntity(v MoBaseMoRelationship) {
	o.AssignedToEntity = &v
}

// GetLeasedResource returns the LeasedResource field value if set, zero value otherwise.
func (o *ResourcepoolLeaseAllOf) GetLeasedResource() ResourcepoolLeaseResourceRelationship {
	if o == nil || o.LeasedResource == nil {
		var ret ResourcepoolLeaseResourceRelationship
		return ret
	}
	return *o.LeasedResource
}

// GetLeasedResourceOk returns a tuple with the LeasedResource field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourcepoolLeaseAllOf) GetLeasedResourceOk() (*ResourcepoolLeaseResourceRelationship, bool) {
	if o == nil || o.LeasedResource == nil {
		return nil, false
	}
	return o.LeasedResource, true
}

// HasLeasedResource returns a boolean if a field has been set.
func (o *ResourcepoolLeaseAllOf) HasLeasedResource() bool {
	if o != nil && o.LeasedResource != nil {
		return true
	}

	return false
}

// SetLeasedResource gets a reference to the given ResourcepoolLeaseResourceRelationship and assigns it to the LeasedResource field.
func (o *ResourcepoolLeaseAllOf) SetLeasedResource(v ResourcepoolLeaseResourceRelationship) {
	o.LeasedResource = &v
}

// GetPool returns the Pool field value if set, zero value otherwise.
func (o *ResourcepoolLeaseAllOf) GetPool() ResourcepoolPoolRelationship {
	if o == nil || o.Pool == nil {
		var ret ResourcepoolPoolRelationship
		return ret
	}
	return *o.Pool
}

// GetPoolOk returns a tuple with the Pool field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourcepoolLeaseAllOf) GetPoolOk() (*ResourcepoolPoolRelationship, bool) {
	if o == nil || o.Pool == nil {
		return nil, false
	}
	return o.Pool, true
}

// HasPool returns a boolean if a field has been set.
func (o *ResourcepoolLeaseAllOf) HasPool() bool {
	if o != nil && o.Pool != nil {
		return true
	}

	return false
}

// SetPool gets a reference to the given ResourcepoolPoolRelationship and assigns it to the Pool field.
func (o *ResourcepoolLeaseAllOf) SetPool(v ResourcepoolPoolRelationship) {
	o.Pool = &v
}

// GetPoolMember returns the PoolMember field value if set, zero value otherwise.
func (o *ResourcepoolLeaseAllOf) GetPoolMember() ResourcepoolPoolMemberRelationship {
	if o == nil || o.PoolMember == nil {
		var ret ResourcepoolPoolMemberRelationship
		return ret
	}
	return *o.PoolMember
}

// GetPoolMemberOk returns a tuple with the PoolMember field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourcepoolLeaseAllOf) GetPoolMemberOk() (*ResourcepoolPoolMemberRelationship, bool) {
	if o == nil || o.PoolMember == nil {
		return nil, false
	}
	return o.PoolMember, true
}

// HasPoolMember returns a boolean if a field has been set.
func (o *ResourcepoolLeaseAllOf) HasPoolMember() bool {
	if o != nil && o.PoolMember != nil {
		return true
	}

	return false
}

// SetPoolMember gets a reference to the given ResourcepoolPoolMemberRelationship and assigns it to the PoolMember field.
func (o *ResourcepoolLeaseAllOf) SetPoolMember(v ResourcepoolPoolMemberRelationship) {
	o.PoolMember = &v
}

// GetUniverse returns the Universe field value if set, zero value otherwise.
func (o *ResourcepoolLeaseAllOf) GetUniverse() ResourcepoolUniverseRelationship {
	if o == nil || o.Universe == nil {
		var ret ResourcepoolUniverseRelationship
		return ret
	}
	return *o.Universe
}

// GetUniverseOk returns a tuple with the Universe field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourcepoolLeaseAllOf) GetUniverseOk() (*ResourcepoolUniverseRelationship, bool) {
	if o == nil || o.Universe == nil {
		return nil, false
	}
	return o.Universe, true
}

// HasUniverse returns a boolean if a field has been set.
func (o *ResourcepoolLeaseAllOf) HasUniverse() bool {
	if o != nil && o.Universe != nil {
		return true
	}

	return false
}

// SetUniverse gets a reference to the given ResourcepoolUniverseRelationship and assigns it to the Universe field.
func (o *ResourcepoolLeaseAllOf) SetUniverse(v ResourcepoolUniverseRelationship) {
	o.Universe = &v
}

func (o ResourcepoolLeaseAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Condition != nil {
		toSerialize["Condition"] = o.Condition
	}
	if o.Feature != nil {
		toSerialize["Feature"] = o.Feature
	}
	if o.LeaseParameters.IsSet() {
		toSerialize["LeaseParameters"] = o.LeaseParameters.Get()
	}
	if o.Resource != nil {
		toSerialize["Resource"] = o.Resource
	}
	if o.ResourceType != nil {
		toSerialize["ResourceType"] = o.ResourceType
	}
	if o.AssignedToEntity != nil {
		toSerialize["AssignedToEntity"] = o.AssignedToEntity
	}
	if o.LeasedResource != nil {
		toSerialize["LeasedResource"] = o.LeasedResource
	}
	if o.Pool != nil {
		toSerialize["Pool"] = o.Pool
	}
	if o.PoolMember != nil {
		toSerialize["PoolMember"] = o.PoolMember
	}
	if o.Universe != nil {
		toSerialize["Universe"] = o.Universe
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ResourcepoolLeaseAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varResourcepoolLeaseAllOf := _ResourcepoolLeaseAllOf{}

	if err = json.Unmarshal(bytes, &varResourcepoolLeaseAllOf); err == nil {
		*o = ResourcepoolLeaseAllOf(varResourcepoolLeaseAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Condition")
		delete(additionalProperties, "Feature")
		delete(additionalProperties, "LeaseParameters")
		delete(additionalProperties, "Resource")
		delete(additionalProperties, "ResourceType")
		delete(additionalProperties, "AssignedToEntity")
		delete(additionalProperties, "LeasedResource")
		delete(additionalProperties, "Pool")
		delete(additionalProperties, "PoolMember")
		delete(additionalProperties, "Universe")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableResourcepoolLeaseAllOf struct {
	value *ResourcepoolLeaseAllOf
	isSet bool
}

func (v NullableResourcepoolLeaseAllOf) Get() *ResourcepoolLeaseAllOf {
	return v.value
}

func (v *NullableResourcepoolLeaseAllOf) Set(val *ResourcepoolLeaseAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableResourcepoolLeaseAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableResourcepoolLeaseAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResourcepoolLeaseAllOf(val *ResourcepoolLeaseAllOf) *NullableResourcepoolLeaseAllOf {
	return &NullableResourcepoolLeaseAllOf{value: val, isSet: true}
}

func (v NullableResourcepoolLeaseAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResourcepoolLeaseAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
