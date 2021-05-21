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

// IppoolIpV6Config Network interface configuration data for IPv6 interfaces. Netmask, Gateway and DNS settings.
type IppoolIpV6Config struct {
	MoBaseComplexType
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// IP address of the default IPv6 gateway.
	Gateway *string `json:"Gateway,omitempty"`
	// A prefix length which masks the  IP address and divides the IP address into network address and host address.
	Prefix *int64 `json:"Prefix,omitempty"`
	// IP Address of the primary Domain Name System (DNS) server.
	PrimaryDns *string `json:"PrimaryDns,omitempty"`
	// IP Address of the secondary Domain Name System (DNS) server.
	SecondaryDns         *string `json:"SecondaryDns,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _IppoolIpV6Config IppoolIpV6Config

// NewIppoolIpV6Config instantiates a new IppoolIpV6Config object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIppoolIpV6Config(classId string, objectType string) *IppoolIpV6Config {
	this := IppoolIpV6Config{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewIppoolIpV6ConfigWithDefaults instantiates a new IppoolIpV6Config object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIppoolIpV6ConfigWithDefaults() *IppoolIpV6Config {
	this := IppoolIpV6Config{}
	var classId string = "ippool.IpV6Config"
	this.ClassId = classId
	var objectType string = "ippool.IpV6Config"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *IppoolIpV6Config) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *IppoolIpV6Config) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *IppoolIpV6Config) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *IppoolIpV6Config) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *IppoolIpV6Config) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *IppoolIpV6Config) SetObjectType(v string) {
	o.ObjectType = v
}

// GetGateway returns the Gateway field value if set, zero value otherwise.
func (o *IppoolIpV6Config) GetGateway() string {
	if o == nil || o.Gateway == nil {
		var ret string
		return ret
	}
	return *o.Gateway
}

// GetGatewayOk returns a tuple with the Gateway field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IppoolIpV6Config) GetGatewayOk() (*string, bool) {
	if o == nil || o.Gateway == nil {
		return nil, false
	}
	return o.Gateway, true
}

// HasGateway returns a boolean if a field has been set.
func (o *IppoolIpV6Config) HasGateway() bool {
	if o != nil && o.Gateway != nil {
		return true
	}

	return false
}

// SetGateway gets a reference to the given string and assigns it to the Gateway field.
func (o *IppoolIpV6Config) SetGateway(v string) {
	o.Gateway = &v
}

// GetPrefix returns the Prefix field value if set, zero value otherwise.
func (o *IppoolIpV6Config) GetPrefix() int64 {
	if o == nil || o.Prefix == nil {
		var ret int64
		return ret
	}
	return *o.Prefix
}

// GetPrefixOk returns a tuple with the Prefix field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IppoolIpV6Config) GetPrefixOk() (*int64, bool) {
	if o == nil || o.Prefix == nil {
		return nil, false
	}
	return o.Prefix, true
}

// HasPrefix returns a boolean if a field has been set.
func (o *IppoolIpV6Config) HasPrefix() bool {
	if o != nil && o.Prefix != nil {
		return true
	}

	return false
}

// SetPrefix gets a reference to the given int64 and assigns it to the Prefix field.
func (o *IppoolIpV6Config) SetPrefix(v int64) {
	o.Prefix = &v
}

// GetPrimaryDns returns the PrimaryDns field value if set, zero value otherwise.
func (o *IppoolIpV6Config) GetPrimaryDns() string {
	if o == nil || o.PrimaryDns == nil {
		var ret string
		return ret
	}
	return *o.PrimaryDns
}

// GetPrimaryDnsOk returns a tuple with the PrimaryDns field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IppoolIpV6Config) GetPrimaryDnsOk() (*string, bool) {
	if o == nil || o.PrimaryDns == nil {
		return nil, false
	}
	return o.PrimaryDns, true
}

// HasPrimaryDns returns a boolean if a field has been set.
func (o *IppoolIpV6Config) HasPrimaryDns() bool {
	if o != nil && o.PrimaryDns != nil {
		return true
	}

	return false
}

// SetPrimaryDns gets a reference to the given string and assigns it to the PrimaryDns field.
func (o *IppoolIpV6Config) SetPrimaryDns(v string) {
	o.PrimaryDns = &v
}

// GetSecondaryDns returns the SecondaryDns field value if set, zero value otherwise.
func (o *IppoolIpV6Config) GetSecondaryDns() string {
	if o == nil || o.SecondaryDns == nil {
		var ret string
		return ret
	}
	return *o.SecondaryDns
}

// GetSecondaryDnsOk returns a tuple with the SecondaryDns field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IppoolIpV6Config) GetSecondaryDnsOk() (*string, bool) {
	if o == nil || o.SecondaryDns == nil {
		return nil, false
	}
	return o.SecondaryDns, true
}

// HasSecondaryDns returns a boolean if a field has been set.
func (o *IppoolIpV6Config) HasSecondaryDns() bool {
	if o != nil && o.SecondaryDns != nil {
		return true
	}

	return false
}

// SetSecondaryDns gets a reference to the given string and assigns it to the SecondaryDns field.
func (o *IppoolIpV6Config) SetSecondaryDns(v string) {
	o.SecondaryDns = &v
}

func (o IppoolIpV6Config) MarshalJSON() ([]byte, error) {
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
	if o.Gateway != nil {
		toSerialize["Gateway"] = o.Gateway
	}
	if o.Prefix != nil {
		toSerialize["Prefix"] = o.Prefix
	}
	if o.PrimaryDns != nil {
		toSerialize["PrimaryDns"] = o.PrimaryDns
	}
	if o.SecondaryDns != nil {
		toSerialize["SecondaryDns"] = o.SecondaryDns
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *IppoolIpV6Config) UnmarshalJSON(bytes []byte) (err error) {
	type IppoolIpV6ConfigWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// IP address of the default IPv6 gateway.
		Gateway *string `json:"Gateway,omitempty"`
		// A prefix length which masks the  IP address and divides the IP address into network address and host address.
		Prefix *int64 `json:"Prefix,omitempty"`
		// IP Address of the primary Domain Name System (DNS) server.
		PrimaryDns *string `json:"PrimaryDns,omitempty"`
		// IP Address of the secondary Domain Name System (DNS) server.
		SecondaryDns *string `json:"SecondaryDns,omitempty"`
	}

	varIppoolIpV6ConfigWithoutEmbeddedStruct := IppoolIpV6ConfigWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varIppoolIpV6ConfigWithoutEmbeddedStruct)
	if err == nil {
		varIppoolIpV6Config := _IppoolIpV6Config{}
		varIppoolIpV6Config.ClassId = varIppoolIpV6ConfigWithoutEmbeddedStruct.ClassId
		varIppoolIpV6Config.ObjectType = varIppoolIpV6ConfigWithoutEmbeddedStruct.ObjectType
		varIppoolIpV6Config.Gateway = varIppoolIpV6ConfigWithoutEmbeddedStruct.Gateway
		varIppoolIpV6Config.Prefix = varIppoolIpV6ConfigWithoutEmbeddedStruct.Prefix
		varIppoolIpV6Config.PrimaryDns = varIppoolIpV6ConfigWithoutEmbeddedStruct.PrimaryDns
		varIppoolIpV6Config.SecondaryDns = varIppoolIpV6ConfigWithoutEmbeddedStruct.SecondaryDns
		*o = IppoolIpV6Config(varIppoolIpV6Config)
	} else {
		return err
	}

	varIppoolIpV6Config := _IppoolIpV6Config{}

	err = json.Unmarshal(bytes, &varIppoolIpV6Config)
	if err == nil {
		o.MoBaseComplexType = varIppoolIpV6Config.MoBaseComplexType
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Gateway")
		delete(additionalProperties, "Prefix")
		delete(additionalProperties, "PrimaryDns")
		delete(additionalProperties, "SecondaryDns")

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

type NullableIppoolIpV6Config struct {
	value *IppoolIpV6Config
	isSet bool
}

func (v NullableIppoolIpV6Config) Get() *IppoolIpV6Config {
	return v.value
}

func (v *NullableIppoolIpV6Config) Set(val *IppoolIpV6Config) {
	v.value = val
	v.isSet = true
}

func (v NullableIppoolIpV6Config) IsSet() bool {
	return v.isSet
}

func (v *NullableIppoolIpV6Config) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIppoolIpV6Config(val *IppoolIpV6Config) *NullableIppoolIpV6Config {
	return &NullableIppoolIpV6Config{value: val, isSet: true}
}

func (v NullableIppoolIpV6Config) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIppoolIpV6Config) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
