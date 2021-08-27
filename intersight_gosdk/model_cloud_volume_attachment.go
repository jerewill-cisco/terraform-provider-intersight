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
	"reflect"
	"strings"
	"time"
)

// CloudVolumeAttachment Volume attachment details for the virtual machine.
type CloudVolumeAttachment struct {
	MoBaseComplexType
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// The time stamp when the attachment of volume to virtual machine is initiated.
	AttachedTime *time.Time `json:"AttachedTime,omitempty"`
	// If true, volume is deleted on virtual machine termination.
	AutoDelete *bool `json:"AutoDelete,omitempty"`
	// The time stamp when the detached of volume to virtual machine is initiated.
	DetachedTime *time.Time `json:"DetachedTime,omitempty"`
	// The device name.For example, /dev/sdh or xvdh in case of AWS cloud.
	DeviceName *string `json:"DeviceName,omitempty"`
	// The internally generated identity of this volume.
	Identity *string `json:"Identity,omitempty"`
	// The position of the volume attachment in the virtual machine.
	Index *int64 `json:"Index,omitempty"`
	// If set to true, then it is the root volume.
	IsRoot               *bool `json:"IsRoot,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CloudVolumeAttachment CloudVolumeAttachment

// NewCloudVolumeAttachment instantiates a new CloudVolumeAttachment object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCloudVolumeAttachment(classId string, objectType string) *CloudVolumeAttachment {
	this := CloudVolumeAttachment{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewCloudVolumeAttachmentWithDefaults instantiates a new CloudVolumeAttachment object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCloudVolumeAttachmentWithDefaults() *CloudVolumeAttachment {
	this := CloudVolumeAttachment{}
	var classId string = "cloud.VolumeAttachment"
	this.ClassId = classId
	var objectType string = "cloud.VolumeAttachment"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *CloudVolumeAttachment) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *CloudVolumeAttachment) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *CloudVolumeAttachment) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *CloudVolumeAttachment) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *CloudVolumeAttachment) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *CloudVolumeAttachment) SetObjectType(v string) {
	o.ObjectType = v
}

// GetAttachedTime returns the AttachedTime field value if set, zero value otherwise.
func (o *CloudVolumeAttachment) GetAttachedTime() time.Time {
	if o == nil || o.AttachedTime == nil {
		var ret time.Time
		return ret
	}
	return *o.AttachedTime
}

// GetAttachedTimeOk returns a tuple with the AttachedTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudVolumeAttachment) GetAttachedTimeOk() (*time.Time, bool) {
	if o == nil || o.AttachedTime == nil {
		return nil, false
	}
	return o.AttachedTime, true
}

// HasAttachedTime returns a boolean if a field has been set.
func (o *CloudVolumeAttachment) HasAttachedTime() bool {
	if o != nil && o.AttachedTime != nil {
		return true
	}

	return false
}

// SetAttachedTime gets a reference to the given time.Time and assigns it to the AttachedTime field.
func (o *CloudVolumeAttachment) SetAttachedTime(v time.Time) {
	o.AttachedTime = &v
}

// GetAutoDelete returns the AutoDelete field value if set, zero value otherwise.
func (o *CloudVolumeAttachment) GetAutoDelete() bool {
	if o == nil || o.AutoDelete == nil {
		var ret bool
		return ret
	}
	return *o.AutoDelete
}

// GetAutoDeleteOk returns a tuple with the AutoDelete field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudVolumeAttachment) GetAutoDeleteOk() (*bool, bool) {
	if o == nil || o.AutoDelete == nil {
		return nil, false
	}
	return o.AutoDelete, true
}

// HasAutoDelete returns a boolean if a field has been set.
func (o *CloudVolumeAttachment) HasAutoDelete() bool {
	if o != nil && o.AutoDelete != nil {
		return true
	}

	return false
}

// SetAutoDelete gets a reference to the given bool and assigns it to the AutoDelete field.
func (o *CloudVolumeAttachment) SetAutoDelete(v bool) {
	o.AutoDelete = &v
}

// GetDetachedTime returns the DetachedTime field value if set, zero value otherwise.
func (o *CloudVolumeAttachment) GetDetachedTime() time.Time {
	if o == nil || o.DetachedTime == nil {
		var ret time.Time
		return ret
	}
	return *o.DetachedTime
}

// GetDetachedTimeOk returns a tuple with the DetachedTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudVolumeAttachment) GetDetachedTimeOk() (*time.Time, bool) {
	if o == nil || o.DetachedTime == nil {
		return nil, false
	}
	return o.DetachedTime, true
}

// HasDetachedTime returns a boolean if a field has been set.
func (o *CloudVolumeAttachment) HasDetachedTime() bool {
	if o != nil && o.DetachedTime != nil {
		return true
	}

	return false
}

// SetDetachedTime gets a reference to the given time.Time and assigns it to the DetachedTime field.
func (o *CloudVolumeAttachment) SetDetachedTime(v time.Time) {
	o.DetachedTime = &v
}

// GetDeviceName returns the DeviceName field value if set, zero value otherwise.
func (o *CloudVolumeAttachment) GetDeviceName() string {
	if o == nil || o.DeviceName == nil {
		var ret string
		return ret
	}
	return *o.DeviceName
}

// GetDeviceNameOk returns a tuple with the DeviceName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudVolumeAttachment) GetDeviceNameOk() (*string, bool) {
	if o == nil || o.DeviceName == nil {
		return nil, false
	}
	return o.DeviceName, true
}

// HasDeviceName returns a boolean if a field has been set.
func (o *CloudVolumeAttachment) HasDeviceName() bool {
	if o != nil && o.DeviceName != nil {
		return true
	}

	return false
}

// SetDeviceName gets a reference to the given string and assigns it to the DeviceName field.
func (o *CloudVolumeAttachment) SetDeviceName(v string) {
	o.DeviceName = &v
}

// GetIdentity returns the Identity field value if set, zero value otherwise.
func (o *CloudVolumeAttachment) GetIdentity() string {
	if o == nil || o.Identity == nil {
		var ret string
		return ret
	}
	return *o.Identity
}

// GetIdentityOk returns a tuple with the Identity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudVolumeAttachment) GetIdentityOk() (*string, bool) {
	if o == nil || o.Identity == nil {
		return nil, false
	}
	return o.Identity, true
}

// HasIdentity returns a boolean if a field has been set.
func (o *CloudVolumeAttachment) HasIdentity() bool {
	if o != nil && o.Identity != nil {
		return true
	}

	return false
}

// SetIdentity gets a reference to the given string and assigns it to the Identity field.
func (o *CloudVolumeAttachment) SetIdentity(v string) {
	o.Identity = &v
}

// GetIndex returns the Index field value if set, zero value otherwise.
func (o *CloudVolumeAttachment) GetIndex() int64 {
	if o == nil || o.Index == nil {
		var ret int64
		return ret
	}
	return *o.Index
}

// GetIndexOk returns a tuple with the Index field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudVolumeAttachment) GetIndexOk() (*int64, bool) {
	if o == nil || o.Index == nil {
		return nil, false
	}
	return o.Index, true
}

// HasIndex returns a boolean if a field has been set.
func (o *CloudVolumeAttachment) HasIndex() bool {
	if o != nil && o.Index != nil {
		return true
	}

	return false
}

// SetIndex gets a reference to the given int64 and assigns it to the Index field.
func (o *CloudVolumeAttachment) SetIndex(v int64) {
	o.Index = &v
}

// GetIsRoot returns the IsRoot field value if set, zero value otherwise.
func (o *CloudVolumeAttachment) GetIsRoot() bool {
	if o == nil || o.IsRoot == nil {
		var ret bool
		return ret
	}
	return *o.IsRoot
}

// GetIsRootOk returns a tuple with the IsRoot field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudVolumeAttachment) GetIsRootOk() (*bool, bool) {
	if o == nil || o.IsRoot == nil {
		return nil, false
	}
	return o.IsRoot, true
}

// HasIsRoot returns a boolean if a field has been set.
func (o *CloudVolumeAttachment) HasIsRoot() bool {
	if o != nil && o.IsRoot != nil {
		return true
	}

	return false
}

// SetIsRoot gets a reference to the given bool and assigns it to the IsRoot field.
func (o *CloudVolumeAttachment) SetIsRoot(v bool) {
	o.IsRoot = &v
}

func (o CloudVolumeAttachment) MarshalJSON() ([]byte, error) {
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
	if o.AttachedTime != nil {
		toSerialize["AttachedTime"] = o.AttachedTime
	}
	if o.AutoDelete != nil {
		toSerialize["AutoDelete"] = o.AutoDelete
	}
	if o.DetachedTime != nil {
		toSerialize["DetachedTime"] = o.DetachedTime
	}
	if o.DeviceName != nil {
		toSerialize["DeviceName"] = o.DeviceName
	}
	if o.Identity != nil {
		toSerialize["Identity"] = o.Identity
	}
	if o.Index != nil {
		toSerialize["Index"] = o.Index
	}
	if o.IsRoot != nil {
		toSerialize["IsRoot"] = o.IsRoot
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *CloudVolumeAttachment) UnmarshalJSON(bytes []byte) (err error) {
	type CloudVolumeAttachmentWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// The time stamp when the attachment of volume to virtual machine is initiated.
		AttachedTime *time.Time `json:"AttachedTime,omitempty"`
		// If true, volume is deleted on virtual machine termination.
		AutoDelete *bool `json:"AutoDelete,omitempty"`
		// The time stamp when the detached of volume to virtual machine is initiated.
		DetachedTime *time.Time `json:"DetachedTime,omitempty"`
		// The device name.For example, /dev/sdh or xvdh in case of AWS cloud.
		DeviceName *string `json:"DeviceName,omitempty"`
		// The internally generated identity of this volume.
		Identity *string `json:"Identity,omitempty"`
		// The position of the volume attachment in the virtual machine.
		Index *int64 `json:"Index,omitempty"`
		// If set to true, then it is the root volume.
		IsRoot *bool `json:"IsRoot,omitempty"`
	}

	varCloudVolumeAttachmentWithoutEmbeddedStruct := CloudVolumeAttachmentWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varCloudVolumeAttachmentWithoutEmbeddedStruct)
	if err == nil {
		varCloudVolumeAttachment := _CloudVolumeAttachment{}
		varCloudVolumeAttachment.ClassId = varCloudVolumeAttachmentWithoutEmbeddedStruct.ClassId
		varCloudVolumeAttachment.ObjectType = varCloudVolumeAttachmentWithoutEmbeddedStruct.ObjectType
		varCloudVolumeAttachment.AttachedTime = varCloudVolumeAttachmentWithoutEmbeddedStruct.AttachedTime
		varCloudVolumeAttachment.AutoDelete = varCloudVolumeAttachmentWithoutEmbeddedStruct.AutoDelete
		varCloudVolumeAttachment.DetachedTime = varCloudVolumeAttachmentWithoutEmbeddedStruct.DetachedTime
		varCloudVolumeAttachment.DeviceName = varCloudVolumeAttachmentWithoutEmbeddedStruct.DeviceName
		varCloudVolumeAttachment.Identity = varCloudVolumeAttachmentWithoutEmbeddedStruct.Identity
		varCloudVolumeAttachment.Index = varCloudVolumeAttachmentWithoutEmbeddedStruct.Index
		varCloudVolumeAttachment.IsRoot = varCloudVolumeAttachmentWithoutEmbeddedStruct.IsRoot
		*o = CloudVolumeAttachment(varCloudVolumeAttachment)
	} else {
		return err
	}

	varCloudVolumeAttachment := _CloudVolumeAttachment{}

	err = json.Unmarshal(bytes, &varCloudVolumeAttachment)
	if err == nil {
		o.MoBaseComplexType = varCloudVolumeAttachment.MoBaseComplexType
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "AttachedTime")
		delete(additionalProperties, "AutoDelete")
		delete(additionalProperties, "DetachedTime")
		delete(additionalProperties, "DeviceName")
		delete(additionalProperties, "Identity")
		delete(additionalProperties, "Index")
		delete(additionalProperties, "IsRoot")

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

type NullableCloudVolumeAttachment struct {
	value *CloudVolumeAttachment
	isSet bool
}

func (v NullableCloudVolumeAttachment) Get() *CloudVolumeAttachment {
	return v.value
}

func (v *NullableCloudVolumeAttachment) Set(val *CloudVolumeAttachment) {
	v.value = val
	v.isSet = true
}

func (v NullableCloudVolumeAttachment) IsSet() bool {
	return v.isSet
}

func (v *NullableCloudVolumeAttachment) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCloudVolumeAttachment(val *CloudVolumeAttachment) *NullableCloudVolumeAttachment {
	return &NullableCloudVolumeAttachment{value: val, isSet: true}
}

func (v NullableCloudVolumeAttachment) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCloudVolumeAttachment) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
