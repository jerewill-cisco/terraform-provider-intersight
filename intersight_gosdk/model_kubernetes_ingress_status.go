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

// KubernetesIngressStatus The current state of the Ingress.
type KubernetesIngressStatus struct {
	MoBaseComplexType
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType           string                         `json:"ObjectType"`
	LoadBalancer         NullableKubernetesLoadBalancer `json:"LoadBalancer,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _KubernetesIngressStatus KubernetesIngressStatus

// NewKubernetesIngressStatus instantiates a new KubernetesIngressStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKubernetesIngressStatus(classId string, objectType string) *KubernetesIngressStatus {
	this := KubernetesIngressStatus{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewKubernetesIngressStatusWithDefaults instantiates a new KubernetesIngressStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKubernetesIngressStatusWithDefaults() *KubernetesIngressStatus {
	this := KubernetesIngressStatus{}
	var classId string = "kubernetes.IngressStatus"
	this.ClassId = classId
	var objectType string = "kubernetes.IngressStatus"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *KubernetesIngressStatus) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *KubernetesIngressStatus) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *KubernetesIngressStatus) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *KubernetesIngressStatus) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *KubernetesIngressStatus) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *KubernetesIngressStatus) SetObjectType(v string) {
	o.ObjectType = v
}

// GetLoadBalancer returns the LoadBalancer field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *KubernetesIngressStatus) GetLoadBalancer() KubernetesLoadBalancer {
	if o == nil || o.LoadBalancer.Get() == nil {
		var ret KubernetesLoadBalancer
		return ret
	}
	return *o.LoadBalancer.Get()
}

// GetLoadBalancerOk returns a tuple with the LoadBalancer field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KubernetesIngressStatus) GetLoadBalancerOk() (*KubernetesLoadBalancer, bool) {
	if o == nil {
		return nil, false
	}
	return o.LoadBalancer.Get(), o.LoadBalancer.IsSet()
}

// HasLoadBalancer returns a boolean if a field has been set.
func (o *KubernetesIngressStatus) HasLoadBalancer() bool {
	if o != nil && o.LoadBalancer.IsSet() {
		return true
	}

	return false
}

// SetLoadBalancer gets a reference to the given NullableKubernetesLoadBalancer and assigns it to the LoadBalancer field.
func (o *KubernetesIngressStatus) SetLoadBalancer(v KubernetesLoadBalancer) {
	o.LoadBalancer.Set(&v)
}

// SetLoadBalancerNil sets the value for LoadBalancer to be an explicit nil
func (o *KubernetesIngressStatus) SetLoadBalancerNil() {
	o.LoadBalancer.Set(nil)
}

// UnsetLoadBalancer ensures that no value is present for LoadBalancer, not even an explicit nil
func (o *KubernetesIngressStatus) UnsetLoadBalancer() {
	o.LoadBalancer.Unset()
}

func (o KubernetesIngressStatus) MarshalJSON() ([]byte, error) {
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
	if o.LoadBalancer.IsSet() {
		toSerialize["LoadBalancer"] = o.LoadBalancer.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *KubernetesIngressStatus) UnmarshalJSON(bytes []byte) (err error) {
	type KubernetesIngressStatusWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType   string                         `json:"ObjectType"`
		LoadBalancer NullableKubernetesLoadBalancer `json:"LoadBalancer,omitempty"`
	}

	varKubernetesIngressStatusWithoutEmbeddedStruct := KubernetesIngressStatusWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varKubernetesIngressStatusWithoutEmbeddedStruct)
	if err == nil {
		varKubernetesIngressStatus := _KubernetesIngressStatus{}
		varKubernetesIngressStatus.ClassId = varKubernetesIngressStatusWithoutEmbeddedStruct.ClassId
		varKubernetesIngressStatus.ObjectType = varKubernetesIngressStatusWithoutEmbeddedStruct.ObjectType
		varKubernetesIngressStatus.LoadBalancer = varKubernetesIngressStatusWithoutEmbeddedStruct.LoadBalancer
		*o = KubernetesIngressStatus(varKubernetesIngressStatus)
	} else {
		return err
	}

	varKubernetesIngressStatus := _KubernetesIngressStatus{}

	err = json.Unmarshal(bytes, &varKubernetesIngressStatus)
	if err == nil {
		o.MoBaseComplexType = varKubernetesIngressStatus.MoBaseComplexType
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "LoadBalancer")

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

type NullableKubernetesIngressStatus struct {
	value *KubernetesIngressStatus
	isSet bool
}

func (v NullableKubernetesIngressStatus) Get() *KubernetesIngressStatus {
	return v.value
}

func (v *NullableKubernetesIngressStatus) Set(val *KubernetesIngressStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableKubernetesIngressStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableKubernetesIngressStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKubernetesIngressStatus(val *KubernetesIngressStatus) *NullableKubernetesIngressStatus {
	return &NullableKubernetesIngressStatus{value: val, isSet: true}
}

func (v NullableKubernetesIngressStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKubernetesIngressStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
