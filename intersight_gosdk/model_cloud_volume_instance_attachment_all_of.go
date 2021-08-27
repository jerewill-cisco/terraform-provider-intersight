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
	"time"
)

// CloudVolumeInstanceAttachmentAllOf Definition of the list of properties defined in 'cloud.VolumeInstanceAttachment', excluding properties defined in parent classes.
type CloudVolumeInstanceAttachmentAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Time stamp when the attachment initiated.
	AttachTime *time.Time `json:"AttachTime,omitempty"`
	// If true, volume is deleted on virtual machine termination.
	AutoDelete *bool `json:"AutoDelete,omitempty"`
	// Device name assigned to the volume.
	DeviceName *string `json:"DeviceName,omitempty"`
	// ID of the virtual machine, the volume is attached to.
	InstanceId *string `json:"InstanceId,omitempty"`
	// The attachment state of the volume. * `UnRecognized` - Volume is in unrecognized state. * `Attached` - Volume is attached to the virtual machine. * `Attaching` - Volume is being attached to the virtual machine. * `Detaching` - Volume is being detached from the virtual machine.
	State                *string `json:"State,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CloudVolumeInstanceAttachmentAllOf CloudVolumeInstanceAttachmentAllOf

// NewCloudVolumeInstanceAttachmentAllOf instantiates a new CloudVolumeInstanceAttachmentAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCloudVolumeInstanceAttachmentAllOf(classId string, objectType string) *CloudVolumeInstanceAttachmentAllOf {
	this := CloudVolumeInstanceAttachmentAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewCloudVolumeInstanceAttachmentAllOfWithDefaults instantiates a new CloudVolumeInstanceAttachmentAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCloudVolumeInstanceAttachmentAllOfWithDefaults() *CloudVolumeInstanceAttachmentAllOf {
	this := CloudVolumeInstanceAttachmentAllOf{}
	var classId string = "cloud.VolumeInstanceAttachment"
	this.ClassId = classId
	var objectType string = "cloud.VolumeInstanceAttachment"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *CloudVolumeInstanceAttachmentAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *CloudVolumeInstanceAttachmentAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *CloudVolumeInstanceAttachmentAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *CloudVolumeInstanceAttachmentAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *CloudVolumeInstanceAttachmentAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *CloudVolumeInstanceAttachmentAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetAttachTime returns the AttachTime field value if set, zero value otherwise.
func (o *CloudVolumeInstanceAttachmentAllOf) GetAttachTime() time.Time {
	if o == nil || o.AttachTime == nil {
		var ret time.Time
		return ret
	}
	return *o.AttachTime
}

// GetAttachTimeOk returns a tuple with the AttachTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudVolumeInstanceAttachmentAllOf) GetAttachTimeOk() (*time.Time, bool) {
	if o == nil || o.AttachTime == nil {
		return nil, false
	}
	return o.AttachTime, true
}

// HasAttachTime returns a boolean if a field has been set.
func (o *CloudVolumeInstanceAttachmentAllOf) HasAttachTime() bool {
	if o != nil && o.AttachTime != nil {
		return true
	}

	return false
}

// SetAttachTime gets a reference to the given time.Time and assigns it to the AttachTime field.
func (o *CloudVolumeInstanceAttachmentAllOf) SetAttachTime(v time.Time) {
	o.AttachTime = &v
}

// GetAutoDelete returns the AutoDelete field value if set, zero value otherwise.
func (o *CloudVolumeInstanceAttachmentAllOf) GetAutoDelete() bool {
	if o == nil || o.AutoDelete == nil {
		var ret bool
		return ret
	}
	return *o.AutoDelete
}

// GetAutoDeleteOk returns a tuple with the AutoDelete field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudVolumeInstanceAttachmentAllOf) GetAutoDeleteOk() (*bool, bool) {
	if o == nil || o.AutoDelete == nil {
		return nil, false
	}
	return o.AutoDelete, true
}

// HasAutoDelete returns a boolean if a field has been set.
func (o *CloudVolumeInstanceAttachmentAllOf) HasAutoDelete() bool {
	if o != nil && o.AutoDelete != nil {
		return true
	}

	return false
}

// SetAutoDelete gets a reference to the given bool and assigns it to the AutoDelete field.
func (o *CloudVolumeInstanceAttachmentAllOf) SetAutoDelete(v bool) {
	o.AutoDelete = &v
}

// GetDeviceName returns the DeviceName field value if set, zero value otherwise.
func (o *CloudVolumeInstanceAttachmentAllOf) GetDeviceName() string {
	if o == nil || o.DeviceName == nil {
		var ret string
		return ret
	}
	return *o.DeviceName
}

// GetDeviceNameOk returns a tuple with the DeviceName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudVolumeInstanceAttachmentAllOf) GetDeviceNameOk() (*string, bool) {
	if o == nil || o.DeviceName == nil {
		return nil, false
	}
	return o.DeviceName, true
}

// HasDeviceName returns a boolean if a field has been set.
func (o *CloudVolumeInstanceAttachmentAllOf) HasDeviceName() bool {
	if o != nil && o.DeviceName != nil {
		return true
	}

	return false
}

// SetDeviceName gets a reference to the given string and assigns it to the DeviceName field.
func (o *CloudVolumeInstanceAttachmentAllOf) SetDeviceName(v string) {
	o.DeviceName = &v
}

// GetInstanceId returns the InstanceId field value if set, zero value otherwise.
func (o *CloudVolumeInstanceAttachmentAllOf) GetInstanceId() string {
	if o == nil || o.InstanceId == nil {
		var ret string
		return ret
	}
	return *o.InstanceId
}

// GetInstanceIdOk returns a tuple with the InstanceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudVolumeInstanceAttachmentAllOf) GetInstanceIdOk() (*string, bool) {
	if o == nil || o.InstanceId == nil {
		return nil, false
	}
	return o.InstanceId, true
}

// HasInstanceId returns a boolean if a field has been set.
func (o *CloudVolumeInstanceAttachmentAllOf) HasInstanceId() bool {
	if o != nil && o.InstanceId != nil {
		return true
	}

	return false
}

// SetInstanceId gets a reference to the given string and assigns it to the InstanceId field.
func (o *CloudVolumeInstanceAttachmentAllOf) SetInstanceId(v string) {
	o.InstanceId = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *CloudVolumeInstanceAttachmentAllOf) GetState() string {
	if o == nil || o.State == nil {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudVolumeInstanceAttachmentAllOf) GetStateOk() (*string, bool) {
	if o == nil || o.State == nil {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *CloudVolumeInstanceAttachmentAllOf) HasState() bool {
	if o != nil && o.State != nil {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *CloudVolumeInstanceAttachmentAllOf) SetState(v string) {
	o.State = &v
}

func (o CloudVolumeInstanceAttachmentAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.AttachTime != nil {
		toSerialize["AttachTime"] = o.AttachTime
	}
	if o.AutoDelete != nil {
		toSerialize["AutoDelete"] = o.AutoDelete
	}
	if o.DeviceName != nil {
		toSerialize["DeviceName"] = o.DeviceName
	}
	if o.InstanceId != nil {
		toSerialize["InstanceId"] = o.InstanceId
	}
	if o.State != nil {
		toSerialize["State"] = o.State
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *CloudVolumeInstanceAttachmentAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varCloudVolumeInstanceAttachmentAllOf := _CloudVolumeInstanceAttachmentAllOf{}

	if err = json.Unmarshal(bytes, &varCloudVolumeInstanceAttachmentAllOf); err == nil {
		*o = CloudVolumeInstanceAttachmentAllOf(varCloudVolumeInstanceAttachmentAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "AttachTime")
		delete(additionalProperties, "AutoDelete")
		delete(additionalProperties, "DeviceName")
		delete(additionalProperties, "InstanceId")
		delete(additionalProperties, "State")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCloudVolumeInstanceAttachmentAllOf struct {
	value *CloudVolumeInstanceAttachmentAllOf
	isSet bool
}

func (v NullableCloudVolumeInstanceAttachmentAllOf) Get() *CloudVolumeInstanceAttachmentAllOf {
	return v.value
}

func (v *NullableCloudVolumeInstanceAttachmentAllOf) Set(val *CloudVolumeInstanceAttachmentAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableCloudVolumeInstanceAttachmentAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableCloudVolumeInstanceAttachmentAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCloudVolumeInstanceAttachmentAllOf(val *CloudVolumeInstanceAttachmentAllOf) *NullableCloudVolumeInstanceAttachmentAllOf {
	return &NullableCloudVolumeInstanceAttachmentAllOf{value: val, isSet: true}
}

func (v NullableCloudVolumeInstanceAttachmentAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCloudVolumeInstanceAttachmentAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
