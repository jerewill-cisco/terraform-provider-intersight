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

// IamOAuthTokenAllOf Definition of the list of properties defined in 'iam.OAuthToken', excluding properties defined in parent classes.
type IamOAuthTokenAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Expiration time for the JWT token to which it can be used for api calls.
	AccessExpirationTime *time.Time `json:"AccessExpirationTime,omitempty"`
	// The identifier of the registered application to which the token belongs.
	ClientId *string `json:"ClientId,omitempty"`
	// The user agent IP address from which the auth token is launched.
	ClientIpAddress *string `json:"ClientIpAddress,omitempty"`
	// The name of the registered application to which the token belongs.
	ClientName *string `json:"ClientName,omitempty"`
	// Expiration time for the JWT token to which it can be refreshed.
	ExpirationTime *time.Time `json:"ExpirationTime,omitempty"`
	// The client address from which last login is initiated.
	LastLoginClient *string `json:"LastLoginClient,omitempty"`
	// The last login time for user.
	LastLoginTime *time.Time `json:"LastLoginTime,omitempty"`
	// Token identifier. Not the Access Token itself.
	TokenId              *string                         `json:"TokenId,omitempty"`
	UserMeta             NullableIamClientMeta           `json:"UserMeta,omitempty"`
	AppRegistration      *IamAppRegistrationRelationship `json:"AppRegistration,omitempty"`
	Permission           *IamPermissionRelationship      `json:"Permission,omitempty"`
	User                 *IamUserRelationship            `json:"User,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _IamOAuthTokenAllOf IamOAuthTokenAllOf

// NewIamOAuthTokenAllOf instantiates a new IamOAuthTokenAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIamOAuthTokenAllOf(classId string, objectType string) *IamOAuthTokenAllOf {
	this := IamOAuthTokenAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewIamOAuthTokenAllOfWithDefaults instantiates a new IamOAuthTokenAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIamOAuthTokenAllOfWithDefaults() *IamOAuthTokenAllOf {
	this := IamOAuthTokenAllOf{}
	var classId string = "iam.OAuthToken"
	this.ClassId = classId
	var objectType string = "iam.OAuthToken"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *IamOAuthTokenAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *IamOAuthTokenAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *IamOAuthTokenAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *IamOAuthTokenAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *IamOAuthTokenAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *IamOAuthTokenAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetAccessExpirationTime returns the AccessExpirationTime field value if set, zero value otherwise.
func (o *IamOAuthTokenAllOf) GetAccessExpirationTime() time.Time {
	if o == nil || o.AccessExpirationTime == nil {
		var ret time.Time
		return ret
	}
	return *o.AccessExpirationTime
}

// GetAccessExpirationTimeOk returns a tuple with the AccessExpirationTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamOAuthTokenAllOf) GetAccessExpirationTimeOk() (*time.Time, bool) {
	if o == nil || o.AccessExpirationTime == nil {
		return nil, false
	}
	return o.AccessExpirationTime, true
}

// HasAccessExpirationTime returns a boolean if a field has been set.
func (o *IamOAuthTokenAllOf) HasAccessExpirationTime() bool {
	if o != nil && o.AccessExpirationTime != nil {
		return true
	}

	return false
}

// SetAccessExpirationTime gets a reference to the given time.Time and assigns it to the AccessExpirationTime field.
func (o *IamOAuthTokenAllOf) SetAccessExpirationTime(v time.Time) {
	o.AccessExpirationTime = &v
}

// GetClientId returns the ClientId field value if set, zero value otherwise.
func (o *IamOAuthTokenAllOf) GetClientId() string {
	if o == nil || o.ClientId == nil {
		var ret string
		return ret
	}
	return *o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamOAuthTokenAllOf) GetClientIdOk() (*string, bool) {
	if o == nil || o.ClientId == nil {
		return nil, false
	}
	return o.ClientId, true
}

// HasClientId returns a boolean if a field has been set.
func (o *IamOAuthTokenAllOf) HasClientId() bool {
	if o != nil && o.ClientId != nil {
		return true
	}

	return false
}

// SetClientId gets a reference to the given string and assigns it to the ClientId field.
func (o *IamOAuthTokenAllOf) SetClientId(v string) {
	o.ClientId = &v
}

// GetClientIpAddress returns the ClientIpAddress field value if set, zero value otherwise.
func (o *IamOAuthTokenAllOf) GetClientIpAddress() string {
	if o == nil || o.ClientIpAddress == nil {
		var ret string
		return ret
	}
	return *o.ClientIpAddress
}

// GetClientIpAddressOk returns a tuple with the ClientIpAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamOAuthTokenAllOf) GetClientIpAddressOk() (*string, bool) {
	if o == nil || o.ClientIpAddress == nil {
		return nil, false
	}
	return o.ClientIpAddress, true
}

// HasClientIpAddress returns a boolean if a field has been set.
func (o *IamOAuthTokenAllOf) HasClientIpAddress() bool {
	if o != nil && o.ClientIpAddress != nil {
		return true
	}

	return false
}

// SetClientIpAddress gets a reference to the given string and assigns it to the ClientIpAddress field.
func (o *IamOAuthTokenAllOf) SetClientIpAddress(v string) {
	o.ClientIpAddress = &v
}

// GetClientName returns the ClientName field value if set, zero value otherwise.
func (o *IamOAuthTokenAllOf) GetClientName() string {
	if o == nil || o.ClientName == nil {
		var ret string
		return ret
	}
	return *o.ClientName
}

// GetClientNameOk returns a tuple with the ClientName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamOAuthTokenAllOf) GetClientNameOk() (*string, bool) {
	if o == nil || o.ClientName == nil {
		return nil, false
	}
	return o.ClientName, true
}

// HasClientName returns a boolean if a field has been set.
func (o *IamOAuthTokenAllOf) HasClientName() bool {
	if o != nil && o.ClientName != nil {
		return true
	}

	return false
}

// SetClientName gets a reference to the given string and assigns it to the ClientName field.
func (o *IamOAuthTokenAllOf) SetClientName(v string) {
	o.ClientName = &v
}

// GetExpirationTime returns the ExpirationTime field value if set, zero value otherwise.
func (o *IamOAuthTokenAllOf) GetExpirationTime() time.Time {
	if o == nil || o.ExpirationTime == nil {
		var ret time.Time
		return ret
	}
	return *o.ExpirationTime
}

// GetExpirationTimeOk returns a tuple with the ExpirationTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamOAuthTokenAllOf) GetExpirationTimeOk() (*time.Time, bool) {
	if o == nil || o.ExpirationTime == nil {
		return nil, false
	}
	return o.ExpirationTime, true
}

// HasExpirationTime returns a boolean if a field has been set.
func (o *IamOAuthTokenAllOf) HasExpirationTime() bool {
	if o != nil && o.ExpirationTime != nil {
		return true
	}

	return false
}

// SetExpirationTime gets a reference to the given time.Time and assigns it to the ExpirationTime field.
func (o *IamOAuthTokenAllOf) SetExpirationTime(v time.Time) {
	o.ExpirationTime = &v
}

// GetLastLoginClient returns the LastLoginClient field value if set, zero value otherwise.
func (o *IamOAuthTokenAllOf) GetLastLoginClient() string {
	if o == nil || o.LastLoginClient == nil {
		var ret string
		return ret
	}
	return *o.LastLoginClient
}

// GetLastLoginClientOk returns a tuple with the LastLoginClient field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamOAuthTokenAllOf) GetLastLoginClientOk() (*string, bool) {
	if o == nil || o.LastLoginClient == nil {
		return nil, false
	}
	return o.LastLoginClient, true
}

// HasLastLoginClient returns a boolean if a field has been set.
func (o *IamOAuthTokenAllOf) HasLastLoginClient() bool {
	if o != nil && o.LastLoginClient != nil {
		return true
	}

	return false
}

// SetLastLoginClient gets a reference to the given string and assigns it to the LastLoginClient field.
func (o *IamOAuthTokenAllOf) SetLastLoginClient(v string) {
	o.LastLoginClient = &v
}

// GetLastLoginTime returns the LastLoginTime field value if set, zero value otherwise.
func (o *IamOAuthTokenAllOf) GetLastLoginTime() time.Time {
	if o == nil || o.LastLoginTime == nil {
		var ret time.Time
		return ret
	}
	return *o.LastLoginTime
}

// GetLastLoginTimeOk returns a tuple with the LastLoginTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamOAuthTokenAllOf) GetLastLoginTimeOk() (*time.Time, bool) {
	if o == nil || o.LastLoginTime == nil {
		return nil, false
	}
	return o.LastLoginTime, true
}

// HasLastLoginTime returns a boolean if a field has been set.
func (o *IamOAuthTokenAllOf) HasLastLoginTime() bool {
	if o != nil && o.LastLoginTime != nil {
		return true
	}

	return false
}

// SetLastLoginTime gets a reference to the given time.Time and assigns it to the LastLoginTime field.
func (o *IamOAuthTokenAllOf) SetLastLoginTime(v time.Time) {
	o.LastLoginTime = &v
}

// GetTokenId returns the TokenId field value if set, zero value otherwise.
func (o *IamOAuthTokenAllOf) GetTokenId() string {
	if o == nil || o.TokenId == nil {
		var ret string
		return ret
	}
	return *o.TokenId
}

// GetTokenIdOk returns a tuple with the TokenId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamOAuthTokenAllOf) GetTokenIdOk() (*string, bool) {
	if o == nil || o.TokenId == nil {
		return nil, false
	}
	return o.TokenId, true
}

// HasTokenId returns a boolean if a field has been set.
func (o *IamOAuthTokenAllOf) HasTokenId() bool {
	if o != nil && o.TokenId != nil {
		return true
	}

	return false
}

// SetTokenId gets a reference to the given string and assigns it to the TokenId field.
func (o *IamOAuthTokenAllOf) SetTokenId(v string) {
	o.TokenId = &v
}

// GetUserMeta returns the UserMeta field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IamOAuthTokenAllOf) GetUserMeta() IamClientMeta {
	if o == nil || o.UserMeta.Get() == nil {
		var ret IamClientMeta
		return ret
	}
	return *o.UserMeta.Get()
}

// GetUserMetaOk returns a tuple with the UserMeta field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IamOAuthTokenAllOf) GetUserMetaOk() (*IamClientMeta, bool) {
	if o == nil {
		return nil, false
	}
	return o.UserMeta.Get(), o.UserMeta.IsSet()
}

// HasUserMeta returns a boolean if a field has been set.
func (o *IamOAuthTokenAllOf) HasUserMeta() bool {
	if o != nil && o.UserMeta.IsSet() {
		return true
	}

	return false
}

// SetUserMeta gets a reference to the given NullableIamClientMeta and assigns it to the UserMeta field.
func (o *IamOAuthTokenAllOf) SetUserMeta(v IamClientMeta) {
	o.UserMeta.Set(&v)
}

// SetUserMetaNil sets the value for UserMeta to be an explicit nil
func (o *IamOAuthTokenAllOf) SetUserMetaNil() {
	o.UserMeta.Set(nil)
}

// UnsetUserMeta ensures that no value is present for UserMeta, not even an explicit nil
func (o *IamOAuthTokenAllOf) UnsetUserMeta() {
	o.UserMeta.Unset()
}

// GetAppRegistration returns the AppRegistration field value if set, zero value otherwise.
func (o *IamOAuthTokenAllOf) GetAppRegistration() IamAppRegistrationRelationship {
	if o == nil || o.AppRegistration == nil {
		var ret IamAppRegistrationRelationship
		return ret
	}
	return *o.AppRegistration
}

// GetAppRegistrationOk returns a tuple with the AppRegistration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamOAuthTokenAllOf) GetAppRegistrationOk() (*IamAppRegistrationRelationship, bool) {
	if o == nil || o.AppRegistration == nil {
		return nil, false
	}
	return o.AppRegistration, true
}

// HasAppRegistration returns a boolean if a field has been set.
func (o *IamOAuthTokenAllOf) HasAppRegistration() bool {
	if o != nil && o.AppRegistration != nil {
		return true
	}

	return false
}

// SetAppRegistration gets a reference to the given IamAppRegistrationRelationship and assigns it to the AppRegistration field.
func (o *IamOAuthTokenAllOf) SetAppRegistration(v IamAppRegistrationRelationship) {
	o.AppRegistration = &v
}

// GetPermission returns the Permission field value if set, zero value otherwise.
func (o *IamOAuthTokenAllOf) GetPermission() IamPermissionRelationship {
	if o == nil || o.Permission == nil {
		var ret IamPermissionRelationship
		return ret
	}
	return *o.Permission
}

// GetPermissionOk returns a tuple with the Permission field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamOAuthTokenAllOf) GetPermissionOk() (*IamPermissionRelationship, bool) {
	if o == nil || o.Permission == nil {
		return nil, false
	}
	return o.Permission, true
}

// HasPermission returns a boolean if a field has been set.
func (o *IamOAuthTokenAllOf) HasPermission() bool {
	if o != nil && o.Permission != nil {
		return true
	}

	return false
}

// SetPermission gets a reference to the given IamPermissionRelationship and assigns it to the Permission field.
func (o *IamOAuthTokenAllOf) SetPermission(v IamPermissionRelationship) {
	o.Permission = &v
}

// GetUser returns the User field value if set, zero value otherwise.
func (o *IamOAuthTokenAllOf) GetUser() IamUserRelationship {
	if o == nil || o.User == nil {
		var ret IamUserRelationship
		return ret
	}
	return *o.User
}

// GetUserOk returns a tuple with the User field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamOAuthTokenAllOf) GetUserOk() (*IamUserRelationship, bool) {
	if o == nil || o.User == nil {
		return nil, false
	}
	return o.User, true
}

// HasUser returns a boolean if a field has been set.
func (o *IamOAuthTokenAllOf) HasUser() bool {
	if o != nil && o.User != nil {
		return true
	}

	return false
}

// SetUser gets a reference to the given IamUserRelationship and assigns it to the User field.
func (o *IamOAuthTokenAllOf) SetUser(v IamUserRelationship) {
	o.User = &v
}

func (o IamOAuthTokenAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.AccessExpirationTime != nil {
		toSerialize["AccessExpirationTime"] = o.AccessExpirationTime
	}
	if o.ClientId != nil {
		toSerialize["ClientId"] = o.ClientId
	}
	if o.ClientIpAddress != nil {
		toSerialize["ClientIpAddress"] = o.ClientIpAddress
	}
	if o.ClientName != nil {
		toSerialize["ClientName"] = o.ClientName
	}
	if o.ExpirationTime != nil {
		toSerialize["ExpirationTime"] = o.ExpirationTime
	}
	if o.LastLoginClient != nil {
		toSerialize["LastLoginClient"] = o.LastLoginClient
	}
	if o.LastLoginTime != nil {
		toSerialize["LastLoginTime"] = o.LastLoginTime
	}
	if o.TokenId != nil {
		toSerialize["TokenId"] = o.TokenId
	}
	if o.UserMeta.IsSet() {
		toSerialize["UserMeta"] = o.UserMeta.Get()
	}
	if o.AppRegistration != nil {
		toSerialize["AppRegistration"] = o.AppRegistration
	}
	if o.Permission != nil {
		toSerialize["Permission"] = o.Permission
	}
	if o.User != nil {
		toSerialize["User"] = o.User
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *IamOAuthTokenAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varIamOAuthTokenAllOf := _IamOAuthTokenAllOf{}

	if err = json.Unmarshal(bytes, &varIamOAuthTokenAllOf); err == nil {
		*o = IamOAuthTokenAllOf(varIamOAuthTokenAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "AccessExpirationTime")
		delete(additionalProperties, "ClientId")
		delete(additionalProperties, "ClientIpAddress")
		delete(additionalProperties, "ClientName")
		delete(additionalProperties, "ExpirationTime")
		delete(additionalProperties, "LastLoginClient")
		delete(additionalProperties, "LastLoginTime")
		delete(additionalProperties, "TokenId")
		delete(additionalProperties, "UserMeta")
		delete(additionalProperties, "AppRegistration")
		delete(additionalProperties, "Permission")
		delete(additionalProperties, "User")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableIamOAuthTokenAllOf struct {
	value *IamOAuthTokenAllOf
	isSet bool
}

func (v NullableIamOAuthTokenAllOf) Get() *IamOAuthTokenAllOf {
	return v.value
}

func (v *NullableIamOAuthTokenAllOf) Set(val *IamOAuthTokenAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableIamOAuthTokenAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableIamOAuthTokenAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIamOAuthTokenAllOf(val *IamOAuthTokenAllOf) *NullableIamOAuthTokenAllOf {
	return &NullableIamOAuthTokenAllOf{value: val, isSet: true}
}

func (v NullableIamOAuthTokenAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIamOAuthTokenAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
