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

// VirtualizationEsxiVmConfiguration Specify ESXi virtual machine configuration data.
type VirtualizationEsxiVmConfiguration struct {
	VirtualizationBaseVmConfiguration
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Specify annotation (optional) for the virtual machine.
	Annotation *string                                          `json:"Annotation,omitempty"`
	Compute    NullableVirtualizationEsxiVmComputeConfiguration `json:"Compute,omitempty"`
	Customspec NullableVirtualizationBaseCustomSpec             `json:"Customspec,omitempty"`
	// Datacenter where virtual machine is deployed.
	Datacenter *string `json:"Datacenter,omitempty"`
	// Folder where virtual machine is deployed.
	Folder *string `json:"Folder,omitempty"`
	// Image path of OVA (The image can be from any location).
	Image   *string                                          `json:"Image,omitempty"`
	Network NullableVirtualizationEsxiVmNetworkConfiguration `json:"Network,omitempty"`
	Storage NullableVirtualizationEsxiVmStorageConfiguration `json:"Storage,omitempty"`
	// Template to be used for clone.
	Template             *string `json:"Template,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _VirtualizationEsxiVmConfiguration VirtualizationEsxiVmConfiguration

// NewVirtualizationEsxiVmConfiguration instantiates a new VirtualizationEsxiVmConfiguration object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVirtualizationEsxiVmConfiguration(classId string, objectType string) *VirtualizationEsxiVmConfiguration {
	this := VirtualizationEsxiVmConfiguration{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewVirtualizationEsxiVmConfigurationWithDefaults instantiates a new VirtualizationEsxiVmConfiguration object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVirtualizationEsxiVmConfigurationWithDefaults() *VirtualizationEsxiVmConfiguration {
	this := VirtualizationEsxiVmConfiguration{}
	var classId string = "virtualization.EsxiVmConfiguration"
	this.ClassId = classId
	var objectType string = "virtualization.EsxiVmConfiguration"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *VirtualizationEsxiVmConfiguration) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *VirtualizationEsxiVmConfiguration) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *VirtualizationEsxiVmConfiguration) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *VirtualizationEsxiVmConfiguration) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *VirtualizationEsxiVmConfiguration) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *VirtualizationEsxiVmConfiguration) SetObjectType(v string) {
	o.ObjectType = v
}

// GetAnnotation returns the Annotation field value if set, zero value otherwise.
func (o *VirtualizationEsxiVmConfiguration) GetAnnotation() string {
	if o == nil || o.Annotation == nil {
		var ret string
		return ret
	}
	return *o.Annotation
}

// GetAnnotationOk returns a tuple with the Annotation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationEsxiVmConfiguration) GetAnnotationOk() (*string, bool) {
	if o == nil || o.Annotation == nil {
		return nil, false
	}
	return o.Annotation, true
}

// HasAnnotation returns a boolean if a field has been set.
func (o *VirtualizationEsxiVmConfiguration) HasAnnotation() bool {
	if o != nil && o.Annotation != nil {
		return true
	}

	return false
}

// SetAnnotation gets a reference to the given string and assigns it to the Annotation field.
func (o *VirtualizationEsxiVmConfiguration) SetAnnotation(v string) {
	o.Annotation = &v
}

// GetCompute returns the Compute field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VirtualizationEsxiVmConfiguration) GetCompute() VirtualizationEsxiVmComputeConfiguration {
	if o == nil || o.Compute.Get() == nil {
		var ret VirtualizationEsxiVmComputeConfiguration
		return ret
	}
	return *o.Compute.Get()
}

// GetComputeOk returns a tuple with the Compute field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VirtualizationEsxiVmConfiguration) GetComputeOk() (*VirtualizationEsxiVmComputeConfiguration, bool) {
	if o == nil {
		return nil, false
	}
	return o.Compute.Get(), o.Compute.IsSet()
}

// HasCompute returns a boolean if a field has been set.
func (o *VirtualizationEsxiVmConfiguration) HasCompute() bool {
	if o != nil && o.Compute.IsSet() {
		return true
	}

	return false
}

// SetCompute gets a reference to the given NullableVirtualizationEsxiVmComputeConfiguration and assigns it to the Compute field.
func (o *VirtualizationEsxiVmConfiguration) SetCompute(v VirtualizationEsxiVmComputeConfiguration) {
	o.Compute.Set(&v)
}

// SetComputeNil sets the value for Compute to be an explicit nil
func (o *VirtualizationEsxiVmConfiguration) SetComputeNil() {
	o.Compute.Set(nil)
}

// UnsetCompute ensures that no value is present for Compute, not even an explicit nil
func (o *VirtualizationEsxiVmConfiguration) UnsetCompute() {
	o.Compute.Unset()
}

// GetCustomspec returns the Customspec field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VirtualizationEsxiVmConfiguration) GetCustomspec() VirtualizationBaseCustomSpec {
	if o == nil || o.Customspec.Get() == nil {
		var ret VirtualizationBaseCustomSpec
		return ret
	}
	return *o.Customspec.Get()
}

// GetCustomspecOk returns a tuple with the Customspec field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VirtualizationEsxiVmConfiguration) GetCustomspecOk() (*VirtualizationBaseCustomSpec, bool) {
	if o == nil {
		return nil, false
	}
	return o.Customspec.Get(), o.Customspec.IsSet()
}

// HasCustomspec returns a boolean if a field has been set.
func (o *VirtualizationEsxiVmConfiguration) HasCustomspec() bool {
	if o != nil && o.Customspec.IsSet() {
		return true
	}

	return false
}

// SetCustomspec gets a reference to the given NullableVirtualizationBaseCustomSpec and assigns it to the Customspec field.
func (o *VirtualizationEsxiVmConfiguration) SetCustomspec(v VirtualizationBaseCustomSpec) {
	o.Customspec.Set(&v)
}

// SetCustomspecNil sets the value for Customspec to be an explicit nil
func (o *VirtualizationEsxiVmConfiguration) SetCustomspecNil() {
	o.Customspec.Set(nil)
}

// UnsetCustomspec ensures that no value is present for Customspec, not even an explicit nil
func (o *VirtualizationEsxiVmConfiguration) UnsetCustomspec() {
	o.Customspec.Unset()
}

// GetDatacenter returns the Datacenter field value if set, zero value otherwise.
func (o *VirtualizationEsxiVmConfiguration) GetDatacenter() string {
	if o == nil || o.Datacenter == nil {
		var ret string
		return ret
	}
	return *o.Datacenter
}

// GetDatacenterOk returns a tuple with the Datacenter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationEsxiVmConfiguration) GetDatacenterOk() (*string, bool) {
	if o == nil || o.Datacenter == nil {
		return nil, false
	}
	return o.Datacenter, true
}

// HasDatacenter returns a boolean if a field has been set.
func (o *VirtualizationEsxiVmConfiguration) HasDatacenter() bool {
	if o != nil && o.Datacenter != nil {
		return true
	}

	return false
}

// SetDatacenter gets a reference to the given string and assigns it to the Datacenter field.
func (o *VirtualizationEsxiVmConfiguration) SetDatacenter(v string) {
	o.Datacenter = &v
}

// GetFolder returns the Folder field value if set, zero value otherwise.
func (o *VirtualizationEsxiVmConfiguration) GetFolder() string {
	if o == nil || o.Folder == nil {
		var ret string
		return ret
	}
	return *o.Folder
}

// GetFolderOk returns a tuple with the Folder field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationEsxiVmConfiguration) GetFolderOk() (*string, bool) {
	if o == nil || o.Folder == nil {
		return nil, false
	}
	return o.Folder, true
}

// HasFolder returns a boolean if a field has been set.
func (o *VirtualizationEsxiVmConfiguration) HasFolder() bool {
	if o != nil && o.Folder != nil {
		return true
	}

	return false
}

// SetFolder gets a reference to the given string and assigns it to the Folder field.
func (o *VirtualizationEsxiVmConfiguration) SetFolder(v string) {
	o.Folder = &v
}

// GetImage returns the Image field value if set, zero value otherwise.
func (o *VirtualizationEsxiVmConfiguration) GetImage() string {
	if o == nil || o.Image == nil {
		var ret string
		return ret
	}
	return *o.Image
}

// GetImageOk returns a tuple with the Image field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationEsxiVmConfiguration) GetImageOk() (*string, bool) {
	if o == nil || o.Image == nil {
		return nil, false
	}
	return o.Image, true
}

// HasImage returns a boolean if a field has been set.
func (o *VirtualizationEsxiVmConfiguration) HasImage() bool {
	if o != nil && o.Image != nil {
		return true
	}

	return false
}

// SetImage gets a reference to the given string and assigns it to the Image field.
func (o *VirtualizationEsxiVmConfiguration) SetImage(v string) {
	o.Image = &v
}

// GetNetwork returns the Network field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VirtualizationEsxiVmConfiguration) GetNetwork() VirtualizationEsxiVmNetworkConfiguration {
	if o == nil || o.Network.Get() == nil {
		var ret VirtualizationEsxiVmNetworkConfiguration
		return ret
	}
	return *o.Network.Get()
}

// GetNetworkOk returns a tuple with the Network field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VirtualizationEsxiVmConfiguration) GetNetworkOk() (*VirtualizationEsxiVmNetworkConfiguration, bool) {
	if o == nil {
		return nil, false
	}
	return o.Network.Get(), o.Network.IsSet()
}

// HasNetwork returns a boolean if a field has been set.
func (o *VirtualizationEsxiVmConfiguration) HasNetwork() bool {
	if o != nil && o.Network.IsSet() {
		return true
	}

	return false
}

// SetNetwork gets a reference to the given NullableVirtualizationEsxiVmNetworkConfiguration and assigns it to the Network field.
func (o *VirtualizationEsxiVmConfiguration) SetNetwork(v VirtualizationEsxiVmNetworkConfiguration) {
	o.Network.Set(&v)
}

// SetNetworkNil sets the value for Network to be an explicit nil
func (o *VirtualizationEsxiVmConfiguration) SetNetworkNil() {
	o.Network.Set(nil)
}

// UnsetNetwork ensures that no value is present for Network, not even an explicit nil
func (o *VirtualizationEsxiVmConfiguration) UnsetNetwork() {
	o.Network.Unset()
}

// GetStorage returns the Storage field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VirtualizationEsxiVmConfiguration) GetStorage() VirtualizationEsxiVmStorageConfiguration {
	if o == nil || o.Storage.Get() == nil {
		var ret VirtualizationEsxiVmStorageConfiguration
		return ret
	}
	return *o.Storage.Get()
}

// GetStorageOk returns a tuple with the Storage field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VirtualizationEsxiVmConfiguration) GetStorageOk() (*VirtualizationEsxiVmStorageConfiguration, bool) {
	if o == nil {
		return nil, false
	}
	return o.Storage.Get(), o.Storage.IsSet()
}

// HasStorage returns a boolean if a field has been set.
func (o *VirtualizationEsxiVmConfiguration) HasStorage() bool {
	if o != nil && o.Storage.IsSet() {
		return true
	}

	return false
}

// SetStorage gets a reference to the given NullableVirtualizationEsxiVmStorageConfiguration and assigns it to the Storage field.
func (o *VirtualizationEsxiVmConfiguration) SetStorage(v VirtualizationEsxiVmStorageConfiguration) {
	o.Storage.Set(&v)
}

// SetStorageNil sets the value for Storage to be an explicit nil
func (o *VirtualizationEsxiVmConfiguration) SetStorageNil() {
	o.Storage.Set(nil)
}

// UnsetStorage ensures that no value is present for Storage, not even an explicit nil
func (o *VirtualizationEsxiVmConfiguration) UnsetStorage() {
	o.Storage.Unset()
}

// GetTemplate returns the Template field value if set, zero value otherwise.
func (o *VirtualizationEsxiVmConfiguration) GetTemplate() string {
	if o == nil || o.Template == nil {
		var ret string
		return ret
	}
	return *o.Template
}

// GetTemplateOk returns a tuple with the Template field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationEsxiVmConfiguration) GetTemplateOk() (*string, bool) {
	if o == nil || o.Template == nil {
		return nil, false
	}
	return o.Template, true
}

// HasTemplate returns a boolean if a field has been set.
func (o *VirtualizationEsxiVmConfiguration) HasTemplate() bool {
	if o != nil && o.Template != nil {
		return true
	}

	return false
}

// SetTemplate gets a reference to the given string and assigns it to the Template field.
func (o *VirtualizationEsxiVmConfiguration) SetTemplate(v string) {
	o.Template = &v
}

func (o VirtualizationEsxiVmConfiguration) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedVirtualizationBaseVmConfiguration, errVirtualizationBaseVmConfiguration := json.Marshal(o.VirtualizationBaseVmConfiguration)
	if errVirtualizationBaseVmConfiguration != nil {
		return []byte{}, errVirtualizationBaseVmConfiguration
	}
	errVirtualizationBaseVmConfiguration = json.Unmarshal([]byte(serializedVirtualizationBaseVmConfiguration), &toSerialize)
	if errVirtualizationBaseVmConfiguration != nil {
		return []byte{}, errVirtualizationBaseVmConfiguration
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Annotation != nil {
		toSerialize["Annotation"] = o.Annotation
	}
	if o.Compute.IsSet() {
		toSerialize["Compute"] = o.Compute.Get()
	}
	if o.Customspec.IsSet() {
		toSerialize["Customspec"] = o.Customspec.Get()
	}
	if o.Datacenter != nil {
		toSerialize["Datacenter"] = o.Datacenter
	}
	if o.Folder != nil {
		toSerialize["Folder"] = o.Folder
	}
	if o.Image != nil {
		toSerialize["Image"] = o.Image
	}
	if o.Network.IsSet() {
		toSerialize["Network"] = o.Network.Get()
	}
	if o.Storage.IsSet() {
		toSerialize["Storage"] = o.Storage.Get()
	}
	if o.Template != nil {
		toSerialize["Template"] = o.Template
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *VirtualizationEsxiVmConfiguration) UnmarshalJSON(bytes []byte) (err error) {
	type VirtualizationEsxiVmConfigurationWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Specify annotation (optional) for the virtual machine.
		Annotation *string                                          `json:"Annotation,omitempty"`
		Compute    NullableVirtualizationEsxiVmComputeConfiguration `json:"Compute,omitempty"`
		Customspec NullableVirtualizationBaseCustomSpec             `json:"Customspec,omitempty"`
		// Datacenter where virtual machine is deployed.
		Datacenter *string `json:"Datacenter,omitempty"`
		// Folder where virtual machine is deployed.
		Folder *string `json:"Folder,omitempty"`
		// Image path of OVA (The image can be from any location).
		Image   *string                                          `json:"Image,omitempty"`
		Network NullableVirtualizationEsxiVmNetworkConfiguration `json:"Network,omitempty"`
		Storage NullableVirtualizationEsxiVmStorageConfiguration `json:"Storage,omitempty"`
		// Template to be used for clone.
		Template *string `json:"Template,omitempty"`
	}

	varVirtualizationEsxiVmConfigurationWithoutEmbeddedStruct := VirtualizationEsxiVmConfigurationWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varVirtualizationEsxiVmConfigurationWithoutEmbeddedStruct)
	if err == nil {
		varVirtualizationEsxiVmConfiguration := _VirtualizationEsxiVmConfiguration{}
		varVirtualizationEsxiVmConfiguration.ClassId = varVirtualizationEsxiVmConfigurationWithoutEmbeddedStruct.ClassId
		varVirtualizationEsxiVmConfiguration.ObjectType = varVirtualizationEsxiVmConfigurationWithoutEmbeddedStruct.ObjectType
		varVirtualizationEsxiVmConfiguration.Annotation = varVirtualizationEsxiVmConfigurationWithoutEmbeddedStruct.Annotation
		varVirtualizationEsxiVmConfiguration.Compute = varVirtualizationEsxiVmConfigurationWithoutEmbeddedStruct.Compute
		varVirtualizationEsxiVmConfiguration.Customspec = varVirtualizationEsxiVmConfigurationWithoutEmbeddedStruct.Customspec
		varVirtualizationEsxiVmConfiguration.Datacenter = varVirtualizationEsxiVmConfigurationWithoutEmbeddedStruct.Datacenter
		varVirtualizationEsxiVmConfiguration.Folder = varVirtualizationEsxiVmConfigurationWithoutEmbeddedStruct.Folder
		varVirtualizationEsxiVmConfiguration.Image = varVirtualizationEsxiVmConfigurationWithoutEmbeddedStruct.Image
		varVirtualizationEsxiVmConfiguration.Network = varVirtualizationEsxiVmConfigurationWithoutEmbeddedStruct.Network
		varVirtualizationEsxiVmConfiguration.Storage = varVirtualizationEsxiVmConfigurationWithoutEmbeddedStruct.Storage
		varVirtualizationEsxiVmConfiguration.Template = varVirtualizationEsxiVmConfigurationWithoutEmbeddedStruct.Template
		*o = VirtualizationEsxiVmConfiguration(varVirtualizationEsxiVmConfiguration)
	} else {
		return err
	}

	varVirtualizationEsxiVmConfiguration := _VirtualizationEsxiVmConfiguration{}

	err = json.Unmarshal(bytes, &varVirtualizationEsxiVmConfiguration)
	if err == nil {
		o.VirtualizationBaseVmConfiguration = varVirtualizationEsxiVmConfiguration.VirtualizationBaseVmConfiguration
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Annotation")
		delete(additionalProperties, "Compute")
		delete(additionalProperties, "Customspec")
		delete(additionalProperties, "Datacenter")
		delete(additionalProperties, "Folder")
		delete(additionalProperties, "Image")
		delete(additionalProperties, "Network")
		delete(additionalProperties, "Storage")
		delete(additionalProperties, "Template")

		// remove fields from embedded structs
		reflectVirtualizationBaseVmConfiguration := reflect.ValueOf(o.VirtualizationBaseVmConfiguration)
		for i := 0; i < reflectVirtualizationBaseVmConfiguration.Type().NumField(); i++ {
			t := reflectVirtualizationBaseVmConfiguration.Type().Field(i)

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

type NullableVirtualizationEsxiVmConfiguration struct {
	value *VirtualizationEsxiVmConfiguration
	isSet bool
}

func (v NullableVirtualizationEsxiVmConfiguration) Get() *VirtualizationEsxiVmConfiguration {
	return v.value
}

func (v *NullableVirtualizationEsxiVmConfiguration) Set(val *VirtualizationEsxiVmConfiguration) {
	v.value = val
	v.isSet = true
}

func (v NullableVirtualizationEsxiVmConfiguration) IsSet() bool {
	return v.isSet
}

func (v *NullableVirtualizationEsxiVmConfiguration) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVirtualizationEsxiVmConfiguration(val *VirtualizationEsxiVmConfiguration) *NullableVirtualizationEsxiVmConfiguration {
	return &NullableVirtualizationEsxiVmConfiguration{value: val, isSet: true}
}

func (v NullableVirtualizationEsxiVmConfiguration) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVirtualizationEsxiVmConfiguration) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
