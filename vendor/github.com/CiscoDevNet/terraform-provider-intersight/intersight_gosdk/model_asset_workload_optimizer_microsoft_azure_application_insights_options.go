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

// AssetWorkloadOptimizerMicrosoftAzureApplicationInsightsOptions Captures configuration specific to the Microsoft Azure Application Intersight target for the Workload Optimizer service.
type AssetWorkloadOptimizerMicrosoftAzureApplicationInsightsOptions struct {
	AssetServiceOptions
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Enrollment number of this Azure account.
	EnrollmentNumber *string `json:"EnrollmentNumber,omitempty"`
	// The offer ID of this account. Default value is \"MS-AZR-0003P\", a pay-as-you-go subscription lets you pay for the services and resources that you use on a monthly basis.
	OfferId *string `json:"OfferId,omitempty"`
	// The Azure Subscription ID.
	SubscriptionId *string `json:"SubscriptionId,omitempty"`
	// Tenant ID associated with Azure Account.
	TenantId             *string `json:"TenantId,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AssetWorkloadOptimizerMicrosoftAzureApplicationInsightsOptions AssetWorkloadOptimizerMicrosoftAzureApplicationInsightsOptions

// NewAssetWorkloadOptimizerMicrosoftAzureApplicationInsightsOptions instantiates a new AssetWorkloadOptimizerMicrosoftAzureApplicationInsightsOptions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAssetWorkloadOptimizerMicrosoftAzureApplicationInsightsOptions(classId string, objectType string) *AssetWorkloadOptimizerMicrosoftAzureApplicationInsightsOptions {
	this := AssetWorkloadOptimizerMicrosoftAzureApplicationInsightsOptions{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewAssetWorkloadOptimizerMicrosoftAzureApplicationInsightsOptionsWithDefaults instantiates a new AssetWorkloadOptimizerMicrosoftAzureApplicationInsightsOptions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAssetWorkloadOptimizerMicrosoftAzureApplicationInsightsOptionsWithDefaults() *AssetWorkloadOptimizerMicrosoftAzureApplicationInsightsOptions {
	this := AssetWorkloadOptimizerMicrosoftAzureApplicationInsightsOptions{}
	var classId string = "asset.WorkloadOptimizerMicrosoftAzureApplicationInsightsOptions"
	this.ClassId = classId
	var objectType string = "asset.WorkloadOptimizerMicrosoftAzureApplicationInsightsOptions"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *AssetWorkloadOptimizerMicrosoftAzureApplicationInsightsOptions) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *AssetWorkloadOptimizerMicrosoftAzureApplicationInsightsOptions) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *AssetWorkloadOptimizerMicrosoftAzureApplicationInsightsOptions) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *AssetWorkloadOptimizerMicrosoftAzureApplicationInsightsOptions) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *AssetWorkloadOptimizerMicrosoftAzureApplicationInsightsOptions) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *AssetWorkloadOptimizerMicrosoftAzureApplicationInsightsOptions) SetObjectType(v string) {
	o.ObjectType = v
}

// GetEnrollmentNumber returns the EnrollmentNumber field value if set, zero value otherwise.
func (o *AssetWorkloadOptimizerMicrosoftAzureApplicationInsightsOptions) GetEnrollmentNumber() string {
	if o == nil || o.EnrollmentNumber == nil {
		var ret string
		return ret
	}
	return *o.EnrollmentNumber
}

// GetEnrollmentNumberOk returns a tuple with the EnrollmentNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetWorkloadOptimizerMicrosoftAzureApplicationInsightsOptions) GetEnrollmentNumberOk() (*string, bool) {
	if o == nil || o.EnrollmentNumber == nil {
		return nil, false
	}
	return o.EnrollmentNumber, true
}

// HasEnrollmentNumber returns a boolean if a field has been set.
func (o *AssetWorkloadOptimizerMicrosoftAzureApplicationInsightsOptions) HasEnrollmentNumber() bool {
	if o != nil && o.EnrollmentNumber != nil {
		return true
	}

	return false
}

// SetEnrollmentNumber gets a reference to the given string and assigns it to the EnrollmentNumber field.
func (o *AssetWorkloadOptimizerMicrosoftAzureApplicationInsightsOptions) SetEnrollmentNumber(v string) {
	o.EnrollmentNumber = &v
}

// GetOfferId returns the OfferId field value if set, zero value otherwise.
func (o *AssetWorkloadOptimizerMicrosoftAzureApplicationInsightsOptions) GetOfferId() string {
	if o == nil || o.OfferId == nil {
		var ret string
		return ret
	}
	return *o.OfferId
}

// GetOfferIdOk returns a tuple with the OfferId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetWorkloadOptimizerMicrosoftAzureApplicationInsightsOptions) GetOfferIdOk() (*string, bool) {
	if o == nil || o.OfferId == nil {
		return nil, false
	}
	return o.OfferId, true
}

// HasOfferId returns a boolean if a field has been set.
func (o *AssetWorkloadOptimizerMicrosoftAzureApplicationInsightsOptions) HasOfferId() bool {
	if o != nil && o.OfferId != nil {
		return true
	}

	return false
}

// SetOfferId gets a reference to the given string and assigns it to the OfferId field.
func (o *AssetWorkloadOptimizerMicrosoftAzureApplicationInsightsOptions) SetOfferId(v string) {
	o.OfferId = &v
}

// GetSubscriptionId returns the SubscriptionId field value if set, zero value otherwise.
func (o *AssetWorkloadOptimizerMicrosoftAzureApplicationInsightsOptions) GetSubscriptionId() string {
	if o == nil || o.SubscriptionId == nil {
		var ret string
		return ret
	}
	return *o.SubscriptionId
}

// GetSubscriptionIdOk returns a tuple with the SubscriptionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetWorkloadOptimizerMicrosoftAzureApplicationInsightsOptions) GetSubscriptionIdOk() (*string, bool) {
	if o == nil || o.SubscriptionId == nil {
		return nil, false
	}
	return o.SubscriptionId, true
}

// HasSubscriptionId returns a boolean if a field has been set.
func (o *AssetWorkloadOptimizerMicrosoftAzureApplicationInsightsOptions) HasSubscriptionId() bool {
	if o != nil && o.SubscriptionId != nil {
		return true
	}

	return false
}

// SetSubscriptionId gets a reference to the given string and assigns it to the SubscriptionId field.
func (o *AssetWorkloadOptimizerMicrosoftAzureApplicationInsightsOptions) SetSubscriptionId(v string) {
	o.SubscriptionId = &v
}

// GetTenantId returns the TenantId field value if set, zero value otherwise.
func (o *AssetWorkloadOptimizerMicrosoftAzureApplicationInsightsOptions) GetTenantId() string {
	if o == nil || o.TenantId == nil {
		var ret string
		return ret
	}
	return *o.TenantId
}

// GetTenantIdOk returns a tuple with the TenantId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetWorkloadOptimizerMicrosoftAzureApplicationInsightsOptions) GetTenantIdOk() (*string, bool) {
	if o == nil || o.TenantId == nil {
		return nil, false
	}
	return o.TenantId, true
}

// HasTenantId returns a boolean if a field has been set.
func (o *AssetWorkloadOptimizerMicrosoftAzureApplicationInsightsOptions) HasTenantId() bool {
	if o != nil && o.TenantId != nil {
		return true
	}

	return false
}

// SetTenantId gets a reference to the given string and assigns it to the TenantId field.
func (o *AssetWorkloadOptimizerMicrosoftAzureApplicationInsightsOptions) SetTenantId(v string) {
	o.TenantId = &v
}

func (o AssetWorkloadOptimizerMicrosoftAzureApplicationInsightsOptions) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedAssetServiceOptions, errAssetServiceOptions := json.Marshal(o.AssetServiceOptions)
	if errAssetServiceOptions != nil {
		return []byte{}, errAssetServiceOptions
	}
	errAssetServiceOptions = json.Unmarshal([]byte(serializedAssetServiceOptions), &toSerialize)
	if errAssetServiceOptions != nil {
		return []byte{}, errAssetServiceOptions
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.EnrollmentNumber != nil {
		toSerialize["EnrollmentNumber"] = o.EnrollmentNumber
	}
	if o.OfferId != nil {
		toSerialize["OfferId"] = o.OfferId
	}
	if o.SubscriptionId != nil {
		toSerialize["SubscriptionId"] = o.SubscriptionId
	}
	if o.TenantId != nil {
		toSerialize["TenantId"] = o.TenantId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *AssetWorkloadOptimizerMicrosoftAzureApplicationInsightsOptions) UnmarshalJSON(bytes []byte) (err error) {
	type AssetWorkloadOptimizerMicrosoftAzureApplicationInsightsOptionsWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Enrollment number of this Azure account.
		EnrollmentNumber *string `json:"EnrollmentNumber,omitempty"`
		// The offer ID of this account. Default value is \"MS-AZR-0003P\", a pay-as-you-go subscription lets you pay for the services and resources that you use on a monthly basis.
		OfferId *string `json:"OfferId,omitempty"`
		// The Azure Subscription ID.
		SubscriptionId *string `json:"SubscriptionId,omitempty"`
		// Tenant ID associated with Azure Account.
		TenantId *string `json:"TenantId,omitempty"`
	}

	varAssetWorkloadOptimizerMicrosoftAzureApplicationInsightsOptionsWithoutEmbeddedStruct := AssetWorkloadOptimizerMicrosoftAzureApplicationInsightsOptionsWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varAssetWorkloadOptimizerMicrosoftAzureApplicationInsightsOptionsWithoutEmbeddedStruct)
	if err == nil {
		varAssetWorkloadOptimizerMicrosoftAzureApplicationInsightsOptions := _AssetWorkloadOptimizerMicrosoftAzureApplicationInsightsOptions{}
		varAssetWorkloadOptimizerMicrosoftAzureApplicationInsightsOptions.ClassId = varAssetWorkloadOptimizerMicrosoftAzureApplicationInsightsOptionsWithoutEmbeddedStruct.ClassId
		varAssetWorkloadOptimizerMicrosoftAzureApplicationInsightsOptions.ObjectType = varAssetWorkloadOptimizerMicrosoftAzureApplicationInsightsOptionsWithoutEmbeddedStruct.ObjectType
		varAssetWorkloadOptimizerMicrosoftAzureApplicationInsightsOptions.EnrollmentNumber = varAssetWorkloadOptimizerMicrosoftAzureApplicationInsightsOptionsWithoutEmbeddedStruct.EnrollmentNumber
		varAssetWorkloadOptimizerMicrosoftAzureApplicationInsightsOptions.OfferId = varAssetWorkloadOptimizerMicrosoftAzureApplicationInsightsOptionsWithoutEmbeddedStruct.OfferId
		varAssetWorkloadOptimizerMicrosoftAzureApplicationInsightsOptions.SubscriptionId = varAssetWorkloadOptimizerMicrosoftAzureApplicationInsightsOptionsWithoutEmbeddedStruct.SubscriptionId
		varAssetWorkloadOptimizerMicrosoftAzureApplicationInsightsOptions.TenantId = varAssetWorkloadOptimizerMicrosoftAzureApplicationInsightsOptionsWithoutEmbeddedStruct.TenantId
		*o = AssetWorkloadOptimizerMicrosoftAzureApplicationInsightsOptions(varAssetWorkloadOptimizerMicrosoftAzureApplicationInsightsOptions)
	} else {
		return err
	}

	varAssetWorkloadOptimizerMicrosoftAzureApplicationInsightsOptions := _AssetWorkloadOptimizerMicrosoftAzureApplicationInsightsOptions{}

	err = json.Unmarshal(bytes, &varAssetWorkloadOptimizerMicrosoftAzureApplicationInsightsOptions)
	if err == nil {
		o.AssetServiceOptions = varAssetWorkloadOptimizerMicrosoftAzureApplicationInsightsOptions.AssetServiceOptions
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "EnrollmentNumber")
		delete(additionalProperties, "OfferId")
		delete(additionalProperties, "SubscriptionId")
		delete(additionalProperties, "TenantId")

		// remove fields from embedded structs
		reflectAssetServiceOptions := reflect.ValueOf(o.AssetServiceOptions)
		for i := 0; i < reflectAssetServiceOptions.Type().NumField(); i++ {
			t := reflectAssetServiceOptions.Type().Field(i)

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

type NullableAssetWorkloadOptimizerMicrosoftAzureApplicationInsightsOptions struct {
	value *AssetWorkloadOptimizerMicrosoftAzureApplicationInsightsOptions
	isSet bool
}

func (v NullableAssetWorkloadOptimizerMicrosoftAzureApplicationInsightsOptions) Get() *AssetWorkloadOptimizerMicrosoftAzureApplicationInsightsOptions {
	return v.value
}

func (v *NullableAssetWorkloadOptimizerMicrosoftAzureApplicationInsightsOptions) Set(val *AssetWorkloadOptimizerMicrosoftAzureApplicationInsightsOptions) {
	v.value = val
	v.isSet = true
}

func (v NullableAssetWorkloadOptimizerMicrosoftAzureApplicationInsightsOptions) IsSet() bool {
	return v.isSet
}

func (v *NullableAssetWorkloadOptimizerMicrosoftAzureApplicationInsightsOptions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAssetWorkloadOptimizerMicrosoftAzureApplicationInsightsOptions(val *AssetWorkloadOptimizerMicrosoftAzureApplicationInsightsOptions) *NullableAssetWorkloadOptimizerMicrosoftAzureApplicationInsightsOptions {
	return &NullableAssetWorkloadOptimizerMicrosoftAzureApplicationInsightsOptions{value: val, isSet: true}
}

func (v NullableAssetWorkloadOptimizerMicrosoftAzureApplicationInsightsOptions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAssetWorkloadOptimizerMicrosoftAzureApplicationInsightsOptions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
