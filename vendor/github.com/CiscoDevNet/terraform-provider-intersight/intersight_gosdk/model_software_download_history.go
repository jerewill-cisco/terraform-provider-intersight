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
	"time"
)

// SoftwareDownloadHistory An object to keep track of software downloads from the Private Appliance portal in SaaS.
type SoftwareDownloadHistory struct {
	MoBaseMo
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// The name of software which was downloaded.
	Name *string `json:"Name,omitempty"`
	// The product type of the downloaded software.
	Product *string `json:"Product,omitempty"`
	// The download time of the software image.
	Timestamp *time.Time `json:"Timestamp,omitempty"`
	// The version of software which was downloaded.
	Version              *string                                `json:"Version,omitempty"`
	Account              *IamAccountRelationship                `json:"Account,omitempty"`
	Image                *FirmwareBaseDistributableRelationship `json:"Image,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SoftwareDownloadHistory SoftwareDownloadHistory

// NewSoftwareDownloadHistory instantiates a new SoftwareDownloadHistory object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSoftwareDownloadHistory(classId string, objectType string) *SoftwareDownloadHistory {
	this := SoftwareDownloadHistory{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewSoftwareDownloadHistoryWithDefaults instantiates a new SoftwareDownloadHistory object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSoftwareDownloadHistoryWithDefaults() *SoftwareDownloadHistory {
	this := SoftwareDownloadHistory{}
	var classId string = "software.DownloadHistory"
	this.ClassId = classId
	var objectType string = "software.DownloadHistory"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *SoftwareDownloadHistory) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *SoftwareDownloadHistory) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *SoftwareDownloadHistory) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *SoftwareDownloadHistory) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *SoftwareDownloadHistory) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *SoftwareDownloadHistory) SetObjectType(v string) {
	o.ObjectType = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *SoftwareDownloadHistory) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SoftwareDownloadHistory) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *SoftwareDownloadHistory) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *SoftwareDownloadHistory) SetName(v string) {
	o.Name = &v
}

// GetProduct returns the Product field value if set, zero value otherwise.
func (o *SoftwareDownloadHistory) GetProduct() string {
	if o == nil || o.Product == nil {
		var ret string
		return ret
	}
	return *o.Product
}

// GetProductOk returns a tuple with the Product field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SoftwareDownloadHistory) GetProductOk() (*string, bool) {
	if o == nil || o.Product == nil {
		return nil, false
	}
	return o.Product, true
}

// HasProduct returns a boolean if a field has been set.
func (o *SoftwareDownloadHistory) HasProduct() bool {
	if o != nil && o.Product != nil {
		return true
	}

	return false
}

// SetProduct gets a reference to the given string and assigns it to the Product field.
func (o *SoftwareDownloadHistory) SetProduct(v string) {
	o.Product = &v
}

// GetTimestamp returns the Timestamp field value if set, zero value otherwise.
func (o *SoftwareDownloadHistory) GetTimestamp() time.Time {
	if o == nil || o.Timestamp == nil {
		var ret time.Time
		return ret
	}
	return *o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SoftwareDownloadHistory) GetTimestampOk() (*time.Time, bool) {
	if o == nil || o.Timestamp == nil {
		return nil, false
	}
	return o.Timestamp, true
}

// HasTimestamp returns a boolean if a field has been set.
func (o *SoftwareDownloadHistory) HasTimestamp() bool {
	if o != nil && o.Timestamp != nil {
		return true
	}

	return false
}

// SetTimestamp gets a reference to the given time.Time and assigns it to the Timestamp field.
func (o *SoftwareDownloadHistory) SetTimestamp(v time.Time) {
	o.Timestamp = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *SoftwareDownloadHistory) GetVersion() string {
	if o == nil || o.Version == nil {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SoftwareDownloadHistory) GetVersionOk() (*string, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *SoftwareDownloadHistory) HasVersion() bool {
	if o != nil && o.Version != nil {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *SoftwareDownloadHistory) SetVersion(v string) {
	o.Version = &v
}

// GetAccount returns the Account field value if set, zero value otherwise.
func (o *SoftwareDownloadHistory) GetAccount() IamAccountRelationship {
	if o == nil || o.Account == nil {
		var ret IamAccountRelationship
		return ret
	}
	return *o.Account
}

// GetAccountOk returns a tuple with the Account field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SoftwareDownloadHistory) GetAccountOk() (*IamAccountRelationship, bool) {
	if o == nil || o.Account == nil {
		return nil, false
	}
	return o.Account, true
}

// HasAccount returns a boolean if a field has been set.
func (o *SoftwareDownloadHistory) HasAccount() bool {
	if o != nil && o.Account != nil {
		return true
	}

	return false
}

// SetAccount gets a reference to the given IamAccountRelationship and assigns it to the Account field.
func (o *SoftwareDownloadHistory) SetAccount(v IamAccountRelationship) {
	o.Account = &v
}

// GetImage returns the Image field value if set, zero value otherwise.
func (o *SoftwareDownloadHistory) GetImage() FirmwareBaseDistributableRelationship {
	if o == nil || o.Image == nil {
		var ret FirmwareBaseDistributableRelationship
		return ret
	}
	return *o.Image
}

// GetImageOk returns a tuple with the Image field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SoftwareDownloadHistory) GetImageOk() (*FirmwareBaseDistributableRelationship, bool) {
	if o == nil || o.Image == nil {
		return nil, false
	}
	return o.Image, true
}

// HasImage returns a boolean if a field has been set.
func (o *SoftwareDownloadHistory) HasImage() bool {
	if o != nil && o.Image != nil {
		return true
	}

	return false
}

// SetImage gets a reference to the given FirmwareBaseDistributableRelationship and assigns it to the Image field.
func (o *SoftwareDownloadHistory) SetImage(v FirmwareBaseDistributableRelationship) {
	o.Image = &v
}

func (o SoftwareDownloadHistory) MarshalJSON() ([]byte, error) {
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
	if o.Name != nil {
		toSerialize["Name"] = o.Name
	}
	if o.Product != nil {
		toSerialize["Product"] = o.Product
	}
	if o.Timestamp != nil {
		toSerialize["Timestamp"] = o.Timestamp
	}
	if o.Version != nil {
		toSerialize["Version"] = o.Version
	}
	if o.Account != nil {
		toSerialize["Account"] = o.Account
	}
	if o.Image != nil {
		toSerialize["Image"] = o.Image
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *SoftwareDownloadHistory) UnmarshalJSON(bytes []byte) (err error) {
	type SoftwareDownloadHistoryWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// The name of software which was downloaded.
		Name *string `json:"Name,omitempty"`
		// The product type of the downloaded software.
		Product *string `json:"Product,omitempty"`
		// The download time of the software image.
		Timestamp *time.Time `json:"Timestamp,omitempty"`
		// The version of software which was downloaded.
		Version *string                                `json:"Version,omitempty"`
		Account *IamAccountRelationship                `json:"Account,omitempty"`
		Image   *FirmwareBaseDistributableRelationship `json:"Image,omitempty"`
	}

	varSoftwareDownloadHistoryWithoutEmbeddedStruct := SoftwareDownloadHistoryWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varSoftwareDownloadHistoryWithoutEmbeddedStruct)
	if err == nil {
		varSoftwareDownloadHistory := _SoftwareDownloadHistory{}
		varSoftwareDownloadHistory.ClassId = varSoftwareDownloadHistoryWithoutEmbeddedStruct.ClassId
		varSoftwareDownloadHistory.ObjectType = varSoftwareDownloadHistoryWithoutEmbeddedStruct.ObjectType
		varSoftwareDownloadHistory.Name = varSoftwareDownloadHistoryWithoutEmbeddedStruct.Name
		varSoftwareDownloadHistory.Product = varSoftwareDownloadHistoryWithoutEmbeddedStruct.Product
		varSoftwareDownloadHistory.Timestamp = varSoftwareDownloadHistoryWithoutEmbeddedStruct.Timestamp
		varSoftwareDownloadHistory.Version = varSoftwareDownloadHistoryWithoutEmbeddedStruct.Version
		varSoftwareDownloadHistory.Account = varSoftwareDownloadHistoryWithoutEmbeddedStruct.Account
		varSoftwareDownloadHistory.Image = varSoftwareDownloadHistoryWithoutEmbeddedStruct.Image
		*o = SoftwareDownloadHistory(varSoftwareDownloadHistory)
	} else {
		return err
	}

	varSoftwareDownloadHistory := _SoftwareDownloadHistory{}

	err = json.Unmarshal(bytes, &varSoftwareDownloadHistory)
	if err == nil {
		o.MoBaseMo = varSoftwareDownloadHistory.MoBaseMo
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Name")
		delete(additionalProperties, "Product")
		delete(additionalProperties, "Timestamp")
		delete(additionalProperties, "Version")
		delete(additionalProperties, "Account")
		delete(additionalProperties, "Image")

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

type NullableSoftwareDownloadHistory struct {
	value *SoftwareDownloadHistory
	isSet bool
}

func (v NullableSoftwareDownloadHistory) Get() *SoftwareDownloadHistory {
	return v.value
}

func (v *NullableSoftwareDownloadHistory) Set(val *SoftwareDownloadHistory) {
	v.value = val
	v.isSet = true
}

func (v NullableSoftwareDownloadHistory) IsSet() bool {
	return v.isSet
}

func (v *NullableSoftwareDownloadHistory) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSoftwareDownloadHistory(val *SoftwareDownloadHistory) *NullableSoftwareDownloadHistory {
	return &NullableSoftwareDownloadHistory{value: val, isSet: true}
}

func (v NullableSoftwareDownloadHistory) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSoftwareDownloadHistory) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
