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
	"reflect"
	"strings"
	"time"
)

// ApplianceSetupInfo SetupInfo will have only one managed object. SetupInfo managed object is to keep track of the Intersight Appliance's setup information and guide the UI through the initial configuration of the Intersight Appliance. The SetupInfo managed object is created during the Intersight Appliance setup. The Intersight UI uses this object to store the initial configuration states that the user has completed. If the user closes the Intersight UI without finishing all the initial configuration, then the Intersight UI will use this managed object to display the next configuration that the user needs to complete when the user uses the Intersight Appliance next time.
type ApplianceSetupInfo struct {
	MoBaseMo
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// The version of Intersight Appliance backup which can restore to.
	BackupVersion *string `json:"BackupVersion,omitempty"`
	// Build type of the Intersight Appliance setup (e.g. release or debug).
	BuildType    *string                 `json:"BuildType,omitempty"`
	Capabilities []ApplianceKeyValuePair `json:"Capabilities,omitempty"`
	// URL of the Intersight to which this Intersight Appliance is connected to.
	CloudUrl *string `json:"CloudUrl,omitempty"`
	// Indicates where Intersight Appliance is installed in air-gapped or connected mode. In connected mode, Intersight Appliance is claimed by Intesight SaaS. In air-gapped mode, Intersight Appliance does not connect to any Cisco services. * `Connected` - In connected mode, Intersight Appliance connects to Intersight SaaS and other cisco.com services. * `Private` - In private mode, Intersight Appliance does not connect to Intersight SaaS or any cisco.com services.
	DeploymentMode *string `json:"DeploymentMode,omitempty"`
	// End date of the Intersight Appliance's initial setup.
	EndTime *time.Time `json:"EndTime,omitempty"`
	// The most recent version which Intersight Appliance can upgrade to.
	LatestVersion *string  `json:"LatestVersion,omitempty"`
	SetupStates   []string `json:"SetupStates,omitempty"`
	// Start date of the Intersight Appliance's initial setup.
	StartTime            *time.Time              `json:"StartTime,omitempty"`
	Account              *IamAccountRelationship `json:"Account,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ApplianceSetupInfo ApplianceSetupInfo

// NewApplianceSetupInfo instantiates a new ApplianceSetupInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApplianceSetupInfo(classId string, objectType string) *ApplianceSetupInfo {
	this := ApplianceSetupInfo{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewApplianceSetupInfoWithDefaults instantiates a new ApplianceSetupInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApplianceSetupInfoWithDefaults() *ApplianceSetupInfo {
	this := ApplianceSetupInfo{}
	var classId string = "appliance.SetupInfo"
	this.ClassId = classId
	var objectType string = "appliance.SetupInfo"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *ApplianceSetupInfo) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *ApplianceSetupInfo) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *ApplianceSetupInfo) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *ApplianceSetupInfo) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *ApplianceSetupInfo) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *ApplianceSetupInfo) SetObjectType(v string) {
	o.ObjectType = v
}

// GetBackupVersion returns the BackupVersion field value if set, zero value otherwise.
func (o *ApplianceSetupInfo) GetBackupVersion() string {
	if o == nil || o.BackupVersion == nil {
		var ret string
		return ret
	}
	return *o.BackupVersion
}

// GetBackupVersionOk returns a tuple with the BackupVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceSetupInfo) GetBackupVersionOk() (*string, bool) {
	if o == nil || o.BackupVersion == nil {
		return nil, false
	}
	return o.BackupVersion, true
}

// HasBackupVersion returns a boolean if a field has been set.
func (o *ApplianceSetupInfo) HasBackupVersion() bool {
	if o != nil && o.BackupVersion != nil {
		return true
	}

	return false
}

// SetBackupVersion gets a reference to the given string and assigns it to the BackupVersion field.
func (o *ApplianceSetupInfo) SetBackupVersion(v string) {
	o.BackupVersion = &v
}

// GetBuildType returns the BuildType field value if set, zero value otherwise.
func (o *ApplianceSetupInfo) GetBuildType() string {
	if o == nil || o.BuildType == nil {
		var ret string
		return ret
	}
	return *o.BuildType
}

// GetBuildTypeOk returns a tuple with the BuildType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceSetupInfo) GetBuildTypeOk() (*string, bool) {
	if o == nil || o.BuildType == nil {
		return nil, false
	}
	return o.BuildType, true
}

// HasBuildType returns a boolean if a field has been set.
func (o *ApplianceSetupInfo) HasBuildType() bool {
	if o != nil && o.BuildType != nil {
		return true
	}

	return false
}

// SetBuildType gets a reference to the given string and assigns it to the BuildType field.
func (o *ApplianceSetupInfo) SetBuildType(v string) {
	o.BuildType = &v
}

// GetCapabilities returns the Capabilities field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApplianceSetupInfo) GetCapabilities() []ApplianceKeyValuePair {
	if o == nil {
		var ret []ApplianceKeyValuePair
		return ret
	}
	return o.Capabilities
}

// GetCapabilitiesOk returns a tuple with the Capabilities field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApplianceSetupInfo) GetCapabilitiesOk() (*[]ApplianceKeyValuePair, bool) {
	if o == nil || o.Capabilities == nil {
		return nil, false
	}
	return &o.Capabilities, true
}

// HasCapabilities returns a boolean if a field has been set.
func (o *ApplianceSetupInfo) HasCapabilities() bool {
	if o != nil && o.Capabilities != nil {
		return true
	}

	return false
}

// SetCapabilities gets a reference to the given []ApplianceKeyValuePair and assigns it to the Capabilities field.
func (o *ApplianceSetupInfo) SetCapabilities(v []ApplianceKeyValuePair) {
	o.Capabilities = v
}

// GetCloudUrl returns the CloudUrl field value if set, zero value otherwise.
func (o *ApplianceSetupInfo) GetCloudUrl() string {
	if o == nil || o.CloudUrl == nil {
		var ret string
		return ret
	}
	return *o.CloudUrl
}

// GetCloudUrlOk returns a tuple with the CloudUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceSetupInfo) GetCloudUrlOk() (*string, bool) {
	if o == nil || o.CloudUrl == nil {
		return nil, false
	}
	return o.CloudUrl, true
}

// HasCloudUrl returns a boolean if a field has been set.
func (o *ApplianceSetupInfo) HasCloudUrl() bool {
	if o != nil && o.CloudUrl != nil {
		return true
	}

	return false
}

// SetCloudUrl gets a reference to the given string and assigns it to the CloudUrl field.
func (o *ApplianceSetupInfo) SetCloudUrl(v string) {
	o.CloudUrl = &v
}

// GetDeploymentMode returns the DeploymentMode field value if set, zero value otherwise.
func (o *ApplianceSetupInfo) GetDeploymentMode() string {
	if o == nil || o.DeploymentMode == nil {
		var ret string
		return ret
	}
	return *o.DeploymentMode
}

// GetDeploymentModeOk returns a tuple with the DeploymentMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceSetupInfo) GetDeploymentModeOk() (*string, bool) {
	if o == nil || o.DeploymentMode == nil {
		return nil, false
	}
	return o.DeploymentMode, true
}

// HasDeploymentMode returns a boolean if a field has been set.
func (o *ApplianceSetupInfo) HasDeploymentMode() bool {
	if o != nil && o.DeploymentMode != nil {
		return true
	}

	return false
}

// SetDeploymentMode gets a reference to the given string and assigns it to the DeploymentMode field.
func (o *ApplianceSetupInfo) SetDeploymentMode(v string) {
	o.DeploymentMode = &v
}

// GetEndTime returns the EndTime field value if set, zero value otherwise.
func (o *ApplianceSetupInfo) GetEndTime() time.Time {
	if o == nil || o.EndTime == nil {
		var ret time.Time
		return ret
	}
	return *o.EndTime
}

// GetEndTimeOk returns a tuple with the EndTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceSetupInfo) GetEndTimeOk() (*time.Time, bool) {
	if o == nil || o.EndTime == nil {
		return nil, false
	}
	return o.EndTime, true
}

// HasEndTime returns a boolean if a field has been set.
func (o *ApplianceSetupInfo) HasEndTime() bool {
	if o != nil && o.EndTime != nil {
		return true
	}

	return false
}

// SetEndTime gets a reference to the given time.Time and assigns it to the EndTime field.
func (o *ApplianceSetupInfo) SetEndTime(v time.Time) {
	o.EndTime = &v
}

// GetLatestVersion returns the LatestVersion field value if set, zero value otherwise.
func (o *ApplianceSetupInfo) GetLatestVersion() string {
	if o == nil || o.LatestVersion == nil {
		var ret string
		return ret
	}
	return *o.LatestVersion
}

// GetLatestVersionOk returns a tuple with the LatestVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceSetupInfo) GetLatestVersionOk() (*string, bool) {
	if o == nil || o.LatestVersion == nil {
		return nil, false
	}
	return o.LatestVersion, true
}

// HasLatestVersion returns a boolean if a field has been set.
func (o *ApplianceSetupInfo) HasLatestVersion() bool {
	if o != nil && o.LatestVersion != nil {
		return true
	}

	return false
}

// SetLatestVersion gets a reference to the given string and assigns it to the LatestVersion field.
func (o *ApplianceSetupInfo) SetLatestVersion(v string) {
	o.LatestVersion = &v
}

// GetSetupStates returns the SetupStates field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApplianceSetupInfo) GetSetupStates() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.SetupStates
}

// GetSetupStatesOk returns a tuple with the SetupStates field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApplianceSetupInfo) GetSetupStatesOk() (*[]string, bool) {
	if o == nil || o.SetupStates == nil {
		return nil, false
	}
	return &o.SetupStates, true
}

// HasSetupStates returns a boolean if a field has been set.
func (o *ApplianceSetupInfo) HasSetupStates() bool {
	if o != nil && o.SetupStates != nil {
		return true
	}

	return false
}

// SetSetupStates gets a reference to the given []string and assigns it to the SetupStates field.
func (o *ApplianceSetupInfo) SetSetupStates(v []string) {
	o.SetupStates = v
}

// GetStartTime returns the StartTime field value if set, zero value otherwise.
func (o *ApplianceSetupInfo) GetStartTime() time.Time {
	if o == nil || o.StartTime == nil {
		var ret time.Time
		return ret
	}
	return *o.StartTime
}

// GetStartTimeOk returns a tuple with the StartTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceSetupInfo) GetStartTimeOk() (*time.Time, bool) {
	if o == nil || o.StartTime == nil {
		return nil, false
	}
	return o.StartTime, true
}

// HasStartTime returns a boolean if a field has been set.
func (o *ApplianceSetupInfo) HasStartTime() bool {
	if o != nil && o.StartTime != nil {
		return true
	}

	return false
}

// SetStartTime gets a reference to the given time.Time and assigns it to the StartTime field.
func (o *ApplianceSetupInfo) SetStartTime(v time.Time) {
	o.StartTime = &v
}

// GetAccount returns the Account field value if set, zero value otherwise.
func (o *ApplianceSetupInfo) GetAccount() IamAccountRelationship {
	if o == nil || o.Account == nil {
		var ret IamAccountRelationship
		return ret
	}
	return *o.Account
}

// GetAccountOk returns a tuple with the Account field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceSetupInfo) GetAccountOk() (*IamAccountRelationship, bool) {
	if o == nil || o.Account == nil {
		return nil, false
	}
	return o.Account, true
}

// HasAccount returns a boolean if a field has been set.
func (o *ApplianceSetupInfo) HasAccount() bool {
	if o != nil && o.Account != nil {
		return true
	}

	return false
}

// SetAccount gets a reference to the given IamAccountRelationship and assigns it to the Account field.
func (o *ApplianceSetupInfo) SetAccount(v IamAccountRelationship) {
	o.Account = &v
}

func (o ApplianceSetupInfo) MarshalJSON() ([]byte, error) {
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
	if o.BackupVersion != nil {
		toSerialize["BackupVersion"] = o.BackupVersion
	}
	if o.BuildType != nil {
		toSerialize["BuildType"] = o.BuildType
	}
	if o.Capabilities != nil {
		toSerialize["Capabilities"] = o.Capabilities
	}
	if o.CloudUrl != nil {
		toSerialize["CloudUrl"] = o.CloudUrl
	}
	if o.DeploymentMode != nil {
		toSerialize["DeploymentMode"] = o.DeploymentMode
	}
	if o.EndTime != nil {
		toSerialize["EndTime"] = o.EndTime
	}
	if o.LatestVersion != nil {
		toSerialize["LatestVersion"] = o.LatestVersion
	}
	if o.SetupStates != nil {
		toSerialize["SetupStates"] = o.SetupStates
	}
	if o.StartTime != nil {
		toSerialize["StartTime"] = o.StartTime
	}
	if o.Account != nil {
		toSerialize["Account"] = o.Account
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ApplianceSetupInfo) UnmarshalJSON(bytes []byte) (err error) {
	type ApplianceSetupInfoWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// The version of Intersight Appliance backup which can restore to.
		BackupVersion *string `json:"BackupVersion,omitempty"`
		// Build type of the Intersight Appliance setup (e.g. release or debug).
		BuildType    *string                 `json:"BuildType,omitempty"`
		Capabilities []ApplianceKeyValuePair `json:"Capabilities,omitempty"`
		// URL of the Intersight to which this Intersight Appliance is connected to.
		CloudUrl *string `json:"CloudUrl,omitempty"`
		// Indicates where Intersight Appliance is installed in air-gapped or connected mode. In connected mode, Intersight Appliance is claimed by Intesight SaaS. In air-gapped mode, Intersight Appliance does not connect to any Cisco services. * `Connected` - In connected mode, Intersight Appliance connects to Intersight SaaS and other cisco.com services. * `Private` - In private mode, Intersight Appliance does not connect to Intersight SaaS or any cisco.com services.
		DeploymentMode *string `json:"DeploymentMode,omitempty"`
		// End date of the Intersight Appliance's initial setup.
		EndTime *time.Time `json:"EndTime,omitempty"`
		// The most recent version which Intersight Appliance can upgrade to.
		LatestVersion *string  `json:"LatestVersion,omitempty"`
		SetupStates   []string `json:"SetupStates,omitempty"`
		// Start date of the Intersight Appliance's initial setup.
		StartTime *time.Time              `json:"StartTime,omitempty"`
		Account   *IamAccountRelationship `json:"Account,omitempty"`
	}

	varApplianceSetupInfoWithoutEmbeddedStruct := ApplianceSetupInfoWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varApplianceSetupInfoWithoutEmbeddedStruct)
	if err == nil {
		varApplianceSetupInfo := _ApplianceSetupInfo{}
		varApplianceSetupInfo.ClassId = varApplianceSetupInfoWithoutEmbeddedStruct.ClassId
		varApplianceSetupInfo.ObjectType = varApplianceSetupInfoWithoutEmbeddedStruct.ObjectType
		varApplianceSetupInfo.BackupVersion = varApplianceSetupInfoWithoutEmbeddedStruct.BackupVersion
		varApplianceSetupInfo.BuildType = varApplianceSetupInfoWithoutEmbeddedStruct.BuildType
		varApplianceSetupInfo.Capabilities = varApplianceSetupInfoWithoutEmbeddedStruct.Capabilities
		varApplianceSetupInfo.CloudUrl = varApplianceSetupInfoWithoutEmbeddedStruct.CloudUrl
		varApplianceSetupInfo.DeploymentMode = varApplianceSetupInfoWithoutEmbeddedStruct.DeploymentMode
		varApplianceSetupInfo.EndTime = varApplianceSetupInfoWithoutEmbeddedStruct.EndTime
		varApplianceSetupInfo.LatestVersion = varApplianceSetupInfoWithoutEmbeddedStruct.LatestVersion
		varApplianceSetupInfo.SetupStates = varApplianceSetupInfoWithoutEmbeddedStruct.SetupStates
		varApplianceSetupInfo.StartTime = varApplianceSetupInfoWithoutEmbeddedStruct.StartTime
		varApplianceSetupInfo.Account = varApplianceSetupInfoWithoutEmbeddedStruct.Account
		*o = ApplianceSetupInfo(varApplianceSetupInfo)
	} else {
		return err
	}

	varApplianceSetupInfo := _ApplianceSetupInfo{}

	err = json.Unmarshal(bytes, &varApplianceSetupInfo)
	if err == nil {
		o.MoBaseMo = varApplianceSetupInfo.MoBaseMo
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "BackupVersion")
		delete(additionalProperties, "BuildType")
		delete(additionalProperties, "Capabilities")
		delete(additionalProperties, "CloudUrl")
		delete(additionalProperties, "DeploymentMode")
		delete(additionalProperties, "EndTime")
		delete(additionalProperties, "LatestVersion")
		delete(additionalProperties, "SetupStates")
		delete(additionalProperties, "StartTime")
		delete(additionalProperties, "Account")

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

type NullableApplianceSetupInfo struct {
	value *ApplianceSetupInfo
	isSet bool
}

func (v NullableApplianceSetupInfo) Get() *ApplianceSetupInfo {
	return v.value
}

func (v *NullableApplianceSetupInfo) Set(val *ApplianceSetupInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableApplianceSetupInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableApplianceSetupInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApplianceSetupInfo(val *ApplianceSetupInfo) *NullableApplianceSetupInfo {
	return &NullableApplianceSetupInfo{value: val, isSet: true}
}

func (v NullableApplianceSetupInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApplianceSetupInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
