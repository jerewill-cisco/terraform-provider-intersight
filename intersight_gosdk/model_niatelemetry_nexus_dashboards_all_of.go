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
)

// NiatelemetryNexusDashboardsAllOf Definition of the list of properties defined in 'niatelemetry.NexusDashboards', excluding properties defined in parent classes.
type NiatelemetryNexusDashboardsAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Nexus Dashboard can onboard multiple APIC clusters/sites.
	ClusterName *string `json:"ClusterName,omitempty"`
	// Dn of the objects present for Nexus Dashboard devices.
	Dn *string `json:"Dn,omitempty"`
	// Health of Nexus Dashboard cluster.
	IsClusterHealthy *string `json:"IsClusterHealthy,omitempty"`
	// Number of nodes in Nexus Dashboard cluster.
	NdClusterSize *int64 `json:"NdClusterSize,omitempty"`
	// Node type in Nexus Dashboard cluster.
	NdType *string `json:"NdType,omitempty"`
	// Version running on Nexus Dashboard.
	NdVersion *string `json:"NdVersion,omitempty"`
	// Number of applications installed in the Nexus Dashboard.
	NumberOfApps *int64 `json:"NumberOfApps,omitempty"`
	// Number of total schemas in Multi-Site Orchestrator.
	NumberOfSchemasInMso *int64 `json:"NumberOfSchemasInMso,omitempty"`
	// Number of sites in Multi-Site Orchestrator.
	NumberOfSitesInMso *int64 `json:"NumberOfSitesInMso,omitempty"`
	// Number of sites serviced by ND.
	NumberOfSitesServiced *int64 `json:"NumberOfSitesServiced,omitempty"`
	// Number of total tenants in Multi-Site Orchestrator.
	NumberOfTenantsInMso *int64 `json:"NumberOfTenantsInMso,omitempty"`
	// Type of record DCNM / APIC / SE. This determines the type of platform where inventory was collected.
	RecordType *string `json:"RecordType,omitempty"`
	// Type of site added to Multi-Site Orchestrator.
	TypeOfSiteInMso      *string                              `json:"TypeOfSiteInMso,omitempty"`
	RegisteredDevice     *AssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _NiatelemetryNexusDashboardsAllOf NiatelemetryNexusDashboardsAllOf

// NewNiatelemetryNexusDashboardsAllOf instantiates a new NiatelemetryNexusDashboardsAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNiatelemetryNexusDashboardsAllOf(classId string, objectType string) *NiatelemetryNexusDashboardsAllOf {
	this := NiatelemetryNexusDashboardsAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewNiatelemetryNexusDashboardsAllOfWithDefaults instantiates a new NiatelemetryNexusDashboardsAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNiatelemetryNexusDashboardsAllOfWithDefaults() *NiatelemetryNexusDashboardsAllOf {
	this := NiatelemetryNexusDashboardsAllOf{}
	var classId string = "niatelemetry.NexusDashboards"
	this.ClassId = classId
	var objectType string = "niatelemetry.NexusDashboards"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *NiatelemetryNexusDashboardsAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *NiatelemetryNexusDashboardsAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *NiatelemetryNexusDashboardsAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *NiatelemetryNexusDashboardsAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *NiatelemetryNexusDashboardsAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *NiatelemetryNexusDashboardsAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetClusterName returns the ClusterName field value if set, zero value otherwise.
func (o *NiatelemetryNexusDashboardsAllOf) GetClusterName() string {
	if o == nil || o.ClusterName == nil {
		var ret string
		return ret
	}
	return *o.ClusterName
}

// GetClusterNameOk returns a tuple with the ClusterName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryNexusDashboardsAllOf) GetClusterNameOk() (*string, bool) {
	if o == nil || o.ClusterName == nil {
		return nil, false
	}
	return o.ClusterName, true
}

// HasClusterName returns a boolean if a field has been set.
func (o *NiatelemetryNexusDashboardsAllOf) HasClusterName() bool {
	if o != nil && o.ClusterName != nil {
		return true
	}

	return false
}

// SetClusterName gets a reference to the given string and assigns it to the ClusterName field.
func (o *NiatelemetryNexusDashboardsAllOf) SetClusterName(v string) {
	o.ClusterName = &v
}

// GetDn returns the Dn field value if set, zero value otherwise.
func (o *NiatelemetryNexusDashboardsAllOf) GetDn() string {
	if o == nil || o.Dn == nil {
		var ret string
		return ret
	}
	return *o.Dn
}

// GetDnOk returns a tuple with the Dn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryNexusDashboardsAllOf) GetDnOk() (*string, bool) {
	if o == nil || o.Dn == nil {
		return nil, false
	}
	return o.Dn, true
}

// HasDn returns a boolean if a field has been set.
func (o *NiatelemetryNexusDashboardsAllOf) HasDn() bool {
	if o != nil && o.Dn != nil {
		return true
	}

	return false
}

// SetDn gets a reference to the given string and assigns it to the Dn field.
func (o *NiatelemetryNexusDashboardsAllOf) SetDn(v string) {
	o.Dn = &v
}

// GetIsClusterHealthy returns the IsClusterHealthy field value if set, zero value otherwise.
func (o *NiatelemetryNexusDashboardsAllOf) GetIsClusterHealthy() string {
	if o == nil || o.IsClusterHealthy == nil {
		var ret string
		return ret
	}
	return *o.IsClusterHealthy
}

// GetIsClusterHealthyOk returns a tuple with the IsClusterHealthy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryNexusDashboardsAllOf) GetIsClusterHealthyOk() (*string, bool) {
	if o == nil || o.IsClusterHealthy == nil {
		return nil, false
	}
	return o.IsClusterHealthy, true
}

// HasIsClusterHealthy returns a boolean if a field has been set.
func (o *NiatelemetryNexusDashboardsAllOf) HasIsClusterHealthy() bool {
	if o != nil && o.IsClusterHealthy != nil {
		return true
	}

	return false
}

// SetIsClusterHealthy gets a reference to the given string and assigns it to the IsClusterHealthy field.
func (o *NiatelemetryNexusDashboardsAllOf) SetIsClusterHealthy(v string) {
	o.IsClusterHealthy = &v
}

// GetNdClusterSize returns the NdClusterSize field value if set, zero value otherwise.
func (o *NiatelemetryNexusDashboardsAllOf) GetNdClusterSize() int64 {
	if o == nil || o.NdClusterSize == nil {
		var ret int64
		return ret
	}
	return *o.NdClusterSize
}

// GetNdClusterSizeOk returns a tuple with the NdClusterSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryNexusDashboardsAllOf) GetNdClusterSizeOk() (*int64, bool) {
	if o == nil || o.NdClusterSize == nil {
		return nil, false
	}
	return o.NdClusterSize, true
}

// HasNdClusterSize returns a boolean if a field has been set.
func (o *NiatelemetryNexusDashboardsAllOf) HasNdClusterSize() bool {
	if o != nil && o.NdClusterSize != nil {
		return true
	}

	return false
}

// SetNdClusterSize gets a reference to the given int64 and assigns it to the NdClusterSize field.
func (o *NiatelemetryNexusDashboardsAllOf) SetNdClusterSize(v int64) {
	o.NdClusterSize = &v
}

// GetNdType returns the NdType field value if set, zero value otherwise.
func (o *NiatelemetryNexusDashboardsAllOf) GetNdType() string {
	if o == nil || o.NdType == nil {
		var ret string
		return ret
	}
	return *o.NdType
}

// GetNdTypeOk returns a tuple with the NdType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryNexusDashboardsAllOf) GetNdTypeOk() (*string, bool) {
	if o == nil || o.NdType == nil {
		return nil, false
	}
	return o.NdType, true
}

// HasNdType returns a boolean if a field has been set.
func (o *NiatelemetryNexusDashboardsAllOf) HasNdType() bool {
	if o != nil && o.NdType != nil {
		return true
	}

	return false
}

// SetNdType gets a reference to the given string and assigns it to the NdType field.
func (o *NiatelemetryNexusDashboardsAllOf) SetNdType(v string) {
	o.NdType = &v
}

// GetNdVersion returns the NdVersion field value if set, zero value otherwise.
func (o *NiatelemetryNexusDashboardsAllOf) GetNdVersion() string {
	if o == nil || o.NdVersion == nil {
		var ret string
		return ret
	}
	return *o.NdVersion
}

// GetNdVersionOk returns a tuple with the NdVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryNexusDashboardsAllOf) GetNdVersionOk() (*string, bool) {
	if o == nil || o.NdVersion == nil {
		return nil, false
	}
	return o.NdVersion, true
}

// HasNdVersion returns a boolean if a field has been set.
func (o *NiatelemetryNexusDashboardsAllOf) HasNdVersion() bool {
	if o != nil && o.NdVersion != nil {
		return true
	}

	return false
}

// SetNdVersion gets a reference to the given string and assigns it to the NdVersion field.
func (o *NiatelemetryNexusDashboardsAllOf) SetNdVersion(v string) {
	o.NdVersion = &v
}

// GetNumberOfApps returns the NumberOfApps field value if set, zero value otherwise.
func (o *NiatelemetryNexusDashboardsAllOf) GetNumberOfApps() int64 {
	if o == nil || o.NumberOfApps == nil {
		var ret int64
		return ret
	}
	return *o.NumberOfApps
}

// GetNumberOfAppsOk returns a tuple with the NumberOfApps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryNexusDashboardsAllOf) GetNumberOfAppsOk() (*int64, bool) {
	if o == nil || o.NumberOfApps == nil {
		return nil, false
	}
	return o.NumberOfApps, true
}

// HasNumberOfApps returns a boolean if a field has been set.
func (o *NiatelemetryNexusDashboardsAllOf) HasNumberOfApps() bool {
	if o != nil && o.NumberOfApps != nil {
		return true
	}

	return false
}

// SetNumberOfApps gets a reference to the given int64 and assigns it to the NumberOfApps field.
func (o *NiatelemetryNexusDashboardsAllOf) SetNumberOfApps(v int64) {
	o.NumberOfApps = &v
}

// GetNumberOfSchemasInMso returns the NumberOfSchemasInMso field value if set, zero value otherwise.
func (o *NiatelemetryNexusDashboardsAllOf) GetNumberOfSchemasInMso() int64 {
	if o == nil || o.NumberOfSchemasInMso == nil {
		var ret int64
		return ret
	}
	return *o.NumberOfSchemasInMso
}

// GetNumberOfSchemasInMsoOk returns a tuple with the NumberOfSchemasInMso field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryNexusDashboardsAllOf) GetNumberOfSchemasInMsoOk() (*int64, bool) {
	if o == nil || o.NumberOfSchemasInMso == nil {
		return nil, false
	}
	return o.NumberOfSchemasInMso, true
}

// HasNumberOfSchemasInMso returns a boolean if a field has been set.
func (o *NiatelemetryNexusDashboardsAllOf) HasNumberOfSchemasInMso() bool {
	if o != nil && o.NumberOfSchemasInMso != nil {
		return true
	}

	return false
}

// SetNumberOfSchemasInMso gets a reference to the given int64 and assigns it to the NumberOfSchemasInMso field.
func (o *NiatelemetryNexusDashboardsAllOf) SetNumberOfSchemasInMso(v int64) {
	o.NumberOfSchemasInMso = &v
}

// GetNumberOfSitesInMso returns the NumberOfSitesInMso field value if set, zero value otherwise.
func (o *NiatelemetryNexusDashboardsAllOf) GetNumberOfSitesInMso() int64 {
	if o == nil || o.NumberOfSitesInMso == nil {
		var ret int64
		return ret
	}
	return *o.NumberOfSitesInMso
}

// GetNumberOfSitesInMsoOk returns a tuple with the NumberOfSitesInMso field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryNexusDashboardsAllOf) GetNumberOfSitesInMsoOk() (*int64, bool) {
	if o == nil || o.NumberOfSitesInMso == nil {
		return nil, false
	}
	return o.NumberOfSitesInMso, true
}

// HasNumberOfSitesInMso returns a boolean if a field has been set.
func (o *NiatelemetryNexusDashboardsAllOf) HasNumberOfSitesInMso() bool {
	if o != nil && o.NumberOfSitesInMso != nil {
		return true
	}

	return false
}

// SetNumberOfSitesInMso gets a reference to the given int64 and assigns it to the NumberOfSitesInMso field.
func (o *NiatelemetryNexusDashboardsAllOf) SetNumberOfSitesInMso(v int64) {
	o.NumberOfSitesInMso = &v
}

// GetNumberOfSitesServiced returns the NumberOfSitesServiced field value if set, zero value otherwise.
func (o *NiatelemetryNexusDashboardsAllOf) GetNumberOfSitesServiced() int64 {
	if o == nil || o.NumberOfSitesServiced == nil {
		var ret int64
		return ret
	}
	return *o.NumberOfSitesServiced
}

// GetNumberOfSitesServicedOk returns a tuple with the NumberOfSitesServiced field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryNexusDashboardsAllOf) GetNumberOfSitesServicedOk() (*int64, bool) {
	if o == nil || o.NumberOfSitesServiced == nil {
		return nil, false
	}
	return o.NumberOfSitesServiced, true
}

// HasNumberOfSitesServiced returns a boolean if a field has been set.
func (o *NiatelemetryNexusDashboardsAllOf) HasNumberOfSitesServiced() bool {
	if o != nil && o.NumberOfSitesServiced != nil {
		return true
	}

	return false
}

// SetNumberOfSitesServiced gets a reference to the given int64 and assigns it to the NumberOfSitesServiced field.
func (o *NiatelemetryNexusDashboardsAllOf) SetNumberOfSitesServiced(v int64) {
	o.NumberOfSitesServiced = &v
}

// GetNumberOfTenantsInMso returns the NumberOfTenantsInMso field value if set, zero value otherwise.
func (o *NiatelemetryNexusDashboardsAllOf) GetNumberOfTenantsInMso() int64 {
	if o == nil || o.NumberOfTenantsInMso == nil {
		var ret int64
		return ret
	}
	return *o.NumberOfTenantsInMso
}

// GetNumberOfTenantsInMsoOk returns a tuple with the NumberOfTenantsInMso field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryNexusDashboardsAllOf) GetNumberOfTenantsInMsoOk() (*int64, bool) {
	if o == nil || o.NumberOfTenantsInMso == nil {
		return nil, false
	}
	return o.NumberOfTenantsInMso, true
}

// HasNumberOfTenantsInMso returns a boolean if a field has been set.
func (o *NiatelemetryNexusDashboardsAllOf) HasNumberOfTenantsInMso() bool {
	if o != nil && o.NumberOfTenantsInMso != nil {
		return true
	}

	return false
}

// SetNumberOfTenantsInMso gets a reference to the given int64 and assigns it to the NumberOfTenantsInMso field.
func (o *NiatelemetryNexusDashboardsAllOf) SetNumberOfTenantsInMso(v int64) {
	o.NumberOfTenantsInMso = &v
}

// GetRecordType returns the RecordType field value if set, zero value otherwise.
func (o *NiatelemetryNexusDashboardsAllOf) GetRecordType() string {
	if o == nil || o.RecordType == nil {
		var ret string
		return ret
	}
	return *o.RecordType
}

// GetRecordTypeOk returns a tuple with the RecordType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryNexusDashboardsAllOf) GetRecordTypeOk() (*string, bool) {
	if o == nil || o.RecordType == nil {
		return nil, false
	}
	return o.RecordType, true
}

// HasRecordType returns a boolean if a field has been set.
func (o *NiatelemetryNexusDashboardsAllOf) HasRecordType() bool {
	if o != nil && o.RecordType != nil {
		return true
	}

	return false
}

// SetRecordType gets a reference to the given string and assigns it to the RecordType field.
func (o *NiatelemetryNexusDashboardsAllOf) SetRecordType(v string) {
	o.RecordType = &v
}

// GetTypeOfSiteInMso returns the TypeOfSiteInMso field value if set, zero value otherwise.
func (o *NiatelemetryNexusDashboardsAllOf) GetTypeOfSiteInMso() string {
	if o == nil || o.TypeOfSiteInMso == nil {
		var ret string
		return ret
	}
	return *o.TypeOfSiteInMso
}

// GetTypeOfSiteInMsoOk returns a tuple with the TypeOfSiteInMso field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryNexusDashboardsAllOf) GetTypeOfSiteInMsoOk() (*string, bool) {
	if o == nil || o.TypeOfSiteInMso == nil {
		return nil, false
	}
	return o.TypeOfSiteInMso, true
}

// HasTypeOfSiteInMso returns a boolean if a field has been set.
func (o *NiatelemetryNexusDashboardsAllOf) HasTypeOfSiteInMso() bool {
	if o != nil && o.TypeOfSiteInMso != nil {
		return true
	}

	return false
}

// SetTypeOfSiteInMso gets a reference to the given string and assigns it to the TypeOfSiteInMso field.
func (o *NiatelemetryNexusDashboardsAllOf) SetTypeOfSiteInMso(v string) {
	o.TypeOfSiteInMso = &v
}

// GetRegisteredDevice returns the RegisteredDevice field value if set, zero value otherwise.
func (o *NiatelemetryNexusDashboardsAllOf) GetRegisteredDevice() AssetDeviceRegistrationRelationship {
	if o == nil || o.RegisteredDevice == nil {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.RegisteredDevice
}

// GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryNexusDashboardsAllOf) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil || o.RegisteredDevice == nil {
		return nil, false
	}
	return o.RegisteredDevice, true
}

// HasRegisteredDevice returns a boolean if a field has been set.
func (o *NiatelemetryNexusDashboardsAllOf) HasRegisteredDevice() bool {
	if o != nil && o.RegisteredDevice != nil {
		return true
	}

	return false
}

// SetRegisteredDevice gets a reference to the given AssetDeviceRegistrationRelationship and assigns it to the RegisteredDevice field.
func (o *NiatelemetryNexusDashboardsAllOf) SetRegisteredDevice(v AssetDeviceRegistrationRelationship) {
	o.RegisteredDevice = &v
}

func (o NiatelemetryNexusDashboardsAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.ClusterName != nil {
		toSerialize["ClusterName"] = o.ClusterName
	}
	if o.Dn != nil {
		toSerialize["Dn"] = o.Dn
	}
	if o.IsClusterHealthy != nil {
		toSerialize["IsClusterHealthy"] = o.IsClusterHealthy
	}
	if o.NdClusterSize != nil {
		toSerialize["NdClusterSize"] = o.NdClusterSize
	}
	if o.NdType != nil {
		toSerialize["NdType"] = o.NdType
	}
	if o.NdVersion != nil {
		toSerialize["NdVersion"] = o.NdVersion
	}
	if o.NumberOfApps != nil {
		toSerialize["NumberOfApps"] = o.NumberOfApps
	}
	if o.NumberOfSchemasInMso != nil {
		toSerialize["NumberOfSchemasInMso"] = o.NumberOfSchemasInMso
	}
	if o.NumberOfSitesInMso != nil {
		toSerialize["NumberOfSitesInMso"] = o.NumberOfSitesInMso
	}
	if o.NumberOfSitesServiced != nil {
		toSerialize["NumberOfSitesServiced"] = o.NumberOfSitesServiced
	}
	if o.NumberOfTenantsInMso != nil {
		toSerialize["NumberOfTenantsInMso"] = o.NumberOfTenantsInMso
	}
	if o.RecordType != nil {
		toSerialize["RecordType"] = o.RecordType
	}
	if o.TypeOfSiteInMso != nil {
		toSerialize["TypeOfSiteInMso"] = o.TypeOfSiteInMso
	}
	if o.RegisteredDevice != nil {
		toSerialize["RegisteredDevice"] = o.RegisteredDevice
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *NiatelemetryNexusDashboardsAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varNiatelemetryNexusDashboardsAllOf := _NiatelemetryNexusDashboardsAllOf{}

	if err = json.Unmarshal(bytes, &varNiatelemetryNexusDashboardsAllOf); err == nil {
		*o = NiatelemetryNexusDashboardsAllOf(varNiatelemetryNexusDashboardsAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "ClusterName")
		delete(additionalProperties, "Dn")
		delete(additionalProperties, "IsClusterHealthy")
		delete(additionalProperties, "NdClusterSize")
		delete(additionalProperties, "NdType")
		delete(additionalProperties, "NdVersion")
		delete(additionalProperties, "NumberOfApps")
		delete(additionalProperties, "NumberOfSchemasInMso")
		delete(additionalProperties, "NumberOfSitesInMso")
		delete(additionalProperties, "NumberOfSitesServiced")
		delete(additionalProperties, "NumberOfTenantsInMso")
		delete(additionalProperties, "RecordType")
		delete(additionalProperties, "TypeOfSiteInMso")
		delete(additionalProperties, "RegisteredDevice")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableNiatelemetryNexusDashboardsAllOf struct {
	value *NiatelemetryNexusDashboardsAllOf
	isSet bool
}

func (v NullableNiatelemetryNexusDashboardsAllOf) Get() *NiatelemetryNexusDashboardsAllOf {
	return v.value
}

func (v *NullableNiatelemetryNexusDashboardsAllOf) Set(val *NiatelemetryNexusDashboardsAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableNiatelemetryNexusDashboardsAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableNiatelemetryNexusDashboardsAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNiatelemetryNexusDashboardsAllOf(val *NiatelemetryNexusDashboardsAllOf) *NullableNiatelemetryNexusDashboardsAllOf {
	return &NullableNiatelemetryNexusDashboardsAllOf{value: val, isSet: true}
}

func (v NullableNiatelemetryNexusDashboardsAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNiatelemetryNexusDashboardsAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
