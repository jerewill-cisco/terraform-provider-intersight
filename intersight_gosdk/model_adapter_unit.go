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
)

// AdapterUnit The physical adapter present on a server.
type AdapterUnit struct {
	EquipmentBase
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Unique Identifier of an adapter Unit within a Rack Interface.
	AdapterId *string `json:"AdapterId,omitempty"`
	// Original Base Mac address of an adapter unit.
	BaseMacAddress *string `json:"BaseMacAddress,omitempty"`
	// Connectivity Status of adapter - A or B or AB.
	ConnectionStatus *string `json:"ConnectionStatus,omitempty"`
	// Cisco Integrated adapter or other type.
	Integrated *string `json:"Integrated,omitempty"`
	// Operational state of an adapter unit.
	OperState *string `json:"OperState,omitempty"`
	// Operability state of an adapter unit.
	Operability *string `json:"Operability,omitempty"`
	// Part number of an adapter unit.
	PartNumber *string `json:"PartNumber,omitempty"`
	// PCIe slot of the adapter in the server.
	PciSlot *string `json:"PciSlot,omitempty"`
	// Power state of an adapter unit.
	Power *string `json:"Power,omitempty"`
	// Thermal state of an adapter unit.
	Thermal *string `json:"Thermal,omitempty"`
	// Virtual Id of the adapter in the server.
	Vid                 *string                           `json:"Vid,omitempty"`
	AdapterUnitExpander *AdapterUnitExpanderRelationship  `json:"AdapterUnitExpander,omitempty"`
	ComputeBlade        *ComputeBladeRelationship         `json:"ComputeBlade,omitempty"`
	ComputeRackUnit     *ComputeRackUnitRelationship      `json:"ComputeRackUnit,omitempty"`
	Controller          *ManagementControllerRelationship `json:"Controller,omitempty"`
	// An array of relationships to adapterExtEthInterface resources.
	ExtEthIfs []AdapterExtEthInterfaceRelationship `json:"ExtEthIfs,omitempty"`
	// An array of relationships to adapterHostEthInterface resources.
	HostEthIfs []AdapterHostEthInterfaceRelationship `json:"HostEthIfs,omitempty"`
	// An array of relationships to adapterHostFcInterface resources.
	HostFcIfs []AdapterHostFcInterfaceRelationship `json:"HostFcIfs,omitempty"`
	// An array of relationships to adapterHostIscsiInterface resources.
	HostIscsiIfs         []AdapterHostIscsiInterfaceRelationship `json:"HostIscsiIfs,omitempty"`
	InventoryDeviceInfo  *InventoryDeviceInfoRelationship        `json:"InventoryDeviceInfo,omitempty"`
	RegisteredDevice     *AssetDeviceRegistrationRelationship    `json:"RegisteredDevice,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AdapterUnit AdapterUnit

// NewAdapterUnit instantiates a new AdapterUnit object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAdapterUnit(classId string, objectType string) *AdapterUnit {
	this := AdapterUnit{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewAdapterUnitWithDefaults instantiates a new AdapterUnit object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAdapterUnitWithDefaults() *AdapterUnit {
	this := AdapterUnit{}
	var classId string = "adapter.Unit"
	this.ClassId = classId
	var objectType string = "adapter.Unit"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *AdapterUnit) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *AdapterUnit) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *AdapterUnit) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *AdapterUnit) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *AdapterUnit) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *AdapterUnit) SetObjectType(v string) {
	o.ObjectType = v
}

// GetAdapterId returns the AdapterId field value if set, zero value otherwise.
func (o *AdapterUnit) GetAdapterId() string {
	if o == nil || o.AdapterId == nil {
		var ret string
		return ret
	}
	return *o.AdapterId
}

// GetAdapterIdOk returns a tuple with the AdapterId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdapterUnit) GetAdapterIdOk() (*string, bool) {
	if o == nil || o.AdapterId == nil {
		return nil, false
	}
	return o.AdapterId, true
}

// HasAdapterId returns a boolean if a field has been set.
func (o *AdapterUnit) HasAdapterId() bool {
	if o != nil && o.AdapterId != nil {
		return true
	}

	return false
}

// SetAdapterId gets a reference to the given string and assigns it to the AdapterId field.
func (o *AdapterUnit) SetAdapterId(v string) {
	o.AdapterId = &v
}

// GetBaseMacAddress returns the BaseMacAddress field value if set, zero value otherwise.
func (o *AdapterUnit) GetBaseMacAddress() string {
	if o == nil || o.BaseMacAddress == nil {
		var ret string
		return ret
	}
	return *o.BaseMacAddress
}

// GetBaseMacAddressOk returns a tuple with the BaseMacAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdapterUnit) GetBaseMacAddressOk() (*string, bool) {
	if o == nil || o.BaseMacAddress == nil {
		return nil, false
	}
	return o.BaseMacAddress, true
}

// HasBaseMacAddress returns a boolean if a field has been set.
func (o *AdapterUnit) HasBaseMacAddress() bool {
	if o != nil && o.BaseMacAddress != nil {
		return true
	}

	return false
}

// SetBaseMacAddress gets a reference to the given string and assigns it to the BaseMacAddress field.
func (o *AdapterUnit) SetBaseMacAddress(v string) {
	o.BaseMacAddress = &v
}

// GetConnectionStatus returns the ConnectionStatus field value if set, zero value otherwise.
func (o *AdapterUnit) GetConnectionStatus() string {
	if o == nil || o.ConnectionStatus == nil {
		var ret string
		return ret
	}
	return *o.ConnectionStatus
}

// GetConnectionStatusOk returns a tuple with the ConnectionStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdapterUnit) GetConnectionStatusOk() (*string, bool) {
	if o == nil || o.ConnectionStatus == nil {
		return nil, false
	}
	return o.ConnectionStatus, true
}

// HasConnectionStatus returns a boolean if a field has been set.
func (o *AdapterUnit) HasConnectionStatus() bool {
	if o != nil && o.ConnectionStatus != nil {
		return true
	}

	return false
}

// SetConnectionStatus gets a reference to the given string and assigns it to the ConnectionStatus field.
func (o *AdapterUnit) SetConnectionStatus(v string) {
	o.ConnectionStatus = &v
}

// GetIntegrated returns the Integrated field value if set, zero value otherwise.
func (o *AdapterUnit) GetIntegrated() string {
	if o == nil || o.Integrated == nil {
		var ret string
		return ret
	}
	return *o.Integrated
}

// GetIntegratedOk returns a tuple with the Integrated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdapterUnit) GetIntegratedOk() (*string, bool) {
	if o == nil || o.Integrated == nil {
		return nil, false
	}
	return o.Integrated, true
}

// HasIntegrated returns a boolean if a field has been set.
func (o *AdapterUnit) HasIntegrated() bool {
	if o != nil && o.Integrated != nil {
		return true
	}

	return false
}

// SetIntegrated gets a reference to the given string and assigns it to the Integrated field.
func (o *AdapterUnit) SetIntegrated(v string) {
	o.Integrated = &v
}

// GetOperState returns the OperState field value if set, zero value otherwise.
func (o *AdapterUnit) GetOperState() string {
	if o == nil || o.OperState == nil {
		var ret string
		return ret
	}
	return *o.OperState
}

// GetOperStateOk returns a tuple with the OperState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdapterUnit) GetOperStateOk() (*string, bool) {
	if o == nil || o.OperState == nil {
		return nil, false
	}
	return o.OperState, true
}

// HasOperState returns a boolean if a field has been set.
func (o *AdapterUnit) HasOperState() bool {
	if o != nil && o.OperState != nil {
		return true
	}

	return false
}

// SetOperState gets a reference to the given string and assigns it to the OperState field.
func (o *AdapterUnit) SetOperState(v string) {
	o.OperState = &v
}

// GetOperability returns the Operability field value if set, zero value otherwise.
func (o *AdapterUnit) GetOperability() string {
	if o == nil || o.Operability == nil {
		var ret string
		return ret
	}
	return *o.Operability
}

// GetOperabilityOk returns a tuple with the Operability field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdapterUnit) GetOperabilityOk() (*string, bool) {
	if o == nil || o.Operability == nil {
		return nil, false
	}
	return o.Operability, true
}

// HasOperability returns a boolean if a field has been set.
func (o *AdapterUnit) HasOperability() bool {
	if o != nil && o.Operability != nil {
		return true
	}

	return false
}

// SetOperability gets a reference to the given string and assigns it to the Operability field.
func (o *AdapterUnit) SetOperability(v string) {
	o.Operability = &v
}

// GetPartNumber returns the PartNumber field value if set, zero value otherwise.
func (o *AdapterUnit) GetPartNumber() string {
	if o == nil || o.PartNumber == nil {
		var ret string
		return ret
	}
	return *o.PartNumber
}

// GetPartNumberOk returns a tuple with the PartNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdapterUnit) GetPartNumberOk() (*string, bool) {
	if o == nil || o.PartNumber == nil {
		return nil, false
	}
	return o.PartNumber, true
}

// HasPartNumber returns a boolean if a field has been set.
func (o *AdapterUnit) HasPartNumber() bool {
	if o != nil && o.PartNumber != nil {
		return true
	}

	return false
}

// SetPartNumber gets a reference to the given string and assigns it to the PartNumber field.
func (o *AdapterUnit) SetPartNumber(v string) {
	o.PartNumber = &v
}

// GetPciSlot returns the PciSlot field value if set, zero value otherwise.
func (o *AdapterUnit) GetPciSlot() string {
	if o == nil || o.PciSlot == nil {
		var ret string
		return ret
	}
	return *o.PciSlot
}

// GetPciSlotOk returns a tuple with the PciSlot field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdapterUnit) GetPciSlotOk() (*string, bool) {
	if o == nil || o.PciSlot == nil {
		return nil, false
	}
	return o.PciSlot, true
}

// HasPciSlot returns a boolean if a field has been set.
func (o *AdapterUnit) HasPciSlot() bool {
	if o != nil && o.PciSlot != nil {
		return true
	}

	return false
}

// SetPciSlot gets a reference to the given string and assigns it to the PciSlot field.
func (o *AdapterUnit) SetPciSlot(v string) {
	o.PciSlot = &v
}

// GetPower returns the Power field value if set, zero value otherwise.
func (o *AdapterUnit) GetPower() string {
	if o == nil || o.Power == nil {
		var ret string
		return ret
	}
	return *o.Power
}

// GetPowerOk returns a tuple with the Power field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdapterUnit) GetPowerOk() (*string, bool) {
	if o == nil || o.Power == nil {
		return nil, false
	}
	return o.Power, true
}

// HasPower returns a boolean if a field has been set.
func (o *AdapterUnit) HasPower() bool {
	if o != nil && o.Power != nil {
		return true
	}

	return false
}

// SetPower gets a reference to the given string and assigns it to the Power field.
func (o *AdapterUnit) SetPower(v string) {
	o.Power = &v
}

// GetThermal returns the Thermal field value if set, zero value otherwise.
func (o *AdapterUnit) GetThermal() string {
	if o == nil || o.Thermal == nil {
		var ret string
		return ret
	}
	return *o.Thermal
}

// GetThermalOk returns a tuple with the Thermal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdapterUnit) GetThermalOk() (*string, bool) {
	if o == nil || o.Thermal == nil {
		return nil, false
	}
	return o.Thermal, true
}

// HasThermal returns a boolean if a field has been set.
func (o *AdapterUnit) HasThermal() bool {
	if o != nil && o.Thermal != nil {
		return true
	}

	return false
}

// SetThermal gets a reference to the given string and assigns it to the Thermal field.
func (o *AdapterUnit) SetThermal(v string) {
	o.Thermal = &v
}

// GetVid returns the Vid field value if set, zero value otherwise.
func (o *AdapterUnit) GetVid() string {
	if o == nil || o.Vid == nil {
		var ret string
		return ret
	}
	return *o.Vid
}

// GetVidOk returns a tuple with the Vid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdapterUnit) GetVidOk() (*string, bool) {
	if o == nil || o.Vid == nil {
		return nil, false
	}
	return o.Vid, true
}

// HasVid returns a boolean if a field has been set.
func (o *AdapterUnit) HasVid() bool {
	if o != nil && o.Vid != nil {
		return true
	}

	return false
}

// SetVid gets a reference to the given string and assigns it to the Vid field.
func (o *AdapterUnit) SetVid(v string) {
	o.Vid = &v
}

// GetAdapterUnitExpander returns the AdapterUnitExpander field value if set, zero value otherwise.
func (o *AdapterUnit) GetAdapterUnitExpander() AdapterUnitExpanderRelationship {
	if o == nil || o.AdapterUnitExpander == nil {
		var ret AdapterUnitExpanderRelationship
		return ret
	}
	return *o.AdapterUnitExpander
}

// GetAdapterUnitExpanderOk returns a tuple with the AdapterUnitExpander field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdapterUnit) GetAdapterUnitExpanderOk() (*AdapterUnitExpanderRelationship, bool) {
	if o == nil || o.AdapterUnitExpander == nil {
		return nil, false
	}
	return o.AdapterUnitExpander, true
}

// HasAdapterUnitExpander returns a boolean if a field has been set.
func (o *AdapterUnit) HasAdapterUnitExpander() bool {
	if o != nil && o.AdapterUnitExpander != nil {
		return true
	}

	return false
}

// SetAdapterUnitExpander gets a reference to the given AdapterUnitExpanderRelationship and assigns it to the AdapterUnitExpander field.
func (o *AdapterUnit) SetAdapterUnitExpander(v AdapterUnitExpanderRelationship) {
	o.AdapterUnitExpander = &v
}

// GetComputeBlade returns the ComputeBlade field value if set, zero value otherwise.
func (o *AdapterUnit) GetComputeBlade() ComputeBladeRelationship {
	if o == nil || o.ComputeBlade == nil {
		var ret ComputeBladeRelationship
		return ret
	}
	return *o.ComputeBlade
}

// GetComputeBladeOk returns a tuple with the ComputeBlade field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdapterUnit) GetComputeBladeOk() (*ComputeBladeRelationship, bool) {
	if o == nil || o.ComputeBlade == nil {
		return nil, false
	}
	return o.ComputeBlade, true
}

// HasComputeBlade returns a boolean if a field has been set.
func (o *AdapterUnit) HasComputeBlade() bool {
	if o != nil && o.ComputeBlade != nil {
		return true
	}

	return false
}

// SetComputeBlade gets a reference to the given ComputeBladeRelationship and assigns it to the ComputeBlade field.
func (o *AdapterUnit) SetComputeBlade(v ComputeBladeRelationship) {
	o.ComputeBlade = &v
}

// GetComputeRackUnit returns the ComputeRackUnit field value if set, zero value otherwise.
func (o *AdapterUnit) GetComputeRackUnit() ComputeRackUnitRelationship {
	if o == nil || o.ComputeRackUnit == nil {
		var ret ComputeRackUnitRelationship
		return ret
	}
	return *o.ComputeRackUnit
}

// GetComputeRackUnitOk returns a tuple with the ComputeRackUnit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdapterUnit) GetComputeRackUnitOk() (*ComputeRackUnitRelationship, bool) {
	if o == nil || o.ComputeRackUnit == nil {
		return nil, false
	}
	return o.ComputeRackUnit, true
}

// HasComputeRackUnit returns a boolean if a field has been set.
func (o *AdapterUnit) HasComputeRackUnit() bool {
	if o != nil && o.ComputeRackUnit != nil {
		return true
	}

	return false
}

// SetComputeRackUnit gets a reference to the given ComputeRackUnitRelationship and assigns it to the ComputeRackUnit field.
func (o *AdapterUnit) SetComputeRackUnit(v ComputeRackUnitRelationship) {
	o.ComputeRackUnit = &v
}

// GetController returns the Controller field value if set, zero value otherwise.
func (o *AdapterUnit) GetController() ManagementControllerRelationship {
	if o == nil || o.Controller == nil {
		var ret ManagementControllerRelationship
		return ret
	}
	return *o.Controller
}

// GetControllerOk returns a tuple with the Controller field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdapterUnit) GetControllerOk() (*ManagementControllerRelationship, bool) {
	if o == nil || o.Controller == nil {
		return nil, false
	}
	return o.Controller, true
}

// HasController returns a boolean if a field has been set.
func (o *AdapterUnit) HasController() bool {
	if o != nil && o.Controller != nil {
		return true
	}

	return false
}

// SetController gets a reference to the given ManagementControllerRelationship and assigns it to the Controller field.
func (o *AdapterUnit) SetController(v ManagementControllerRelationship) {
	o.Controller = &v
}

// GetExtEthIfs returns the ExtEthIfs field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AdapterUnit) GetExtEthIfs() []AdapterExtEthInterfaceRelationship {
	if o == nil {
		var ret []AdapterExtEthInterfaceRelationship
		return ret
	}
	return o.ExtEthIfs
}

// GetExtEthIfsOk returns a tuple with the ExtEthIfs field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AdapterUnit) GetExtEthIfsOk() (*[]AdapterExtEthInterfaceRelationship, bool) {
	if o == nil || o.ExtEthIfs == nil {
		return nil, false
	}
	return &o.ExtEthIfs, true
}

// HasExtEthIfs returns a boolean if a field has been set.
func (o *AdapterUnit) HasExtEthIfs() bool {
	if o != nil && o.ExtEthIfs != nil {
		return true
	}

	return false
}

// SetExtEthIfs gets a reference to the given []AdapterExtEthInterfaceRelationship and assigns it to the ExtEthIfs field.
func (o *AdapterUnit) SetExtEthIfs(v []AdapterExtEthInterfaceRelationship) {
	o.ExtEthIfs = v
}

// GetHostEthIfs returns the HostEthIfs field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AdapterUnit) GetHostEthIfs() []AdapterHostEthInterfaceRelationship {
	if o == nil {
		var ret []AdapterHostEthInterfaceRelationship
		return ret
	}
	return o.HostEthIfs
}

// GetHostEthIfsOk returns a tuple with the HostEthIfs field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AdapterUnit) GetHostEthIfsOk() (*[]AdapterHostEthInterfaceRelationship, bool) {
	if o == nil || o.HostEthIfs == nil {
		return nil, false
	}
	return &o.HostEthIfs, true
}

// HasHostEthIfs returns a boolean if a field has been set.
func (o *AdapterUnit) HasHostEthIfs() bool {
	if o != nil && o.HostEthIfs != nil {
		return true
	}

	return false
}

// SetHostEthIfs gets a reference to the given []AdapterHostEthInterfaceRelationship and assigns it to the HostEthIfs field.
func (o *AdapterUnit) SetHostEthIfs(v []AdapterHostEthInterfaceRelationship) {
	o.HostEthIfs = v
}

// GetHostFcIfs returns the HostFcIfs field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AdapterUnit) GetHostFcIfs() []AdapterHostFcInterfaceRelationship {
	if o == nil {
		var ret []AdapterHostFcInterfaceRelationship
		return ret
	}
	return o.HostFcIfs
}

// GetHostFcIfsOk returns a tuple with the HostFcIfs field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AdapterUnit) GetHostFcIfsOk() (*[]AdapterHostFcInterfaceRelationship, bool) {
	if o == nil || o.HostFcIfs == nil {
		return nil, false
	}
	return &o.HostFcIfs, true
}

// HasHostFcIfs returns a boolean if a field has been set.
func (o *AdapterUnit) HasHostFcIfs() bool {
	if o != nil && o.HostFcIfs != nil {
		return true
	}

	return false
}

// SetHostFcIfs gets a reference to the given []AdapterHostFcInterfaceRelationship and assigns it to the HostFcIfs field.
func (o *AdapterUnit) SetHostFcIfs(v []AdapterHostFcInterfaceRelationship) {
	o.HostFcIfs = v
}

// GetHostIscsiIfs returns the HostIscsiIfs field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AdapterUnit) GetHostIscsiIfs() []AdapterHostIscsiInterfaceRelationship {
	if o == nil {
		var ret []AdapterHostIscsiInterfaceRelationship
		return ret
	}
	return o.HostIscsiIfs
}

// GetHostIscsiIfsOk returns a tuple with the HostIscsiIfs field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AdapterUnit) GetHostIscsiIfsOk() (*[]AdapterHostIscsiInterfaceRelationship, bool) {
	if o == nil || o.HostIscsiIfs == nil {
		return nil, false
	}
	return &o.HostIscsiIfs, true
}

// HasHostIscsiIfs returns a boolean if a field has been set.
func (o *AdapterUnit) HasHostIscsiIfs() bool {
	if o != nil && o.HostIscsiIfs != nil {
		return true
	}

	return false
}

// SetHostIscsiIfs gets a reference to the given []AdapterHostIscsiInterfaceRelationship and assigns it to the HostIscsiIfs field.
func (o *AdapterUnit) SetHostIscsiIfs(v []AdapterHostIscsiInterfaceRelationship) {
	o.HostIscsiIfs = v
}

// GetInventoryDeviceInfo returns the InventoryDeviceInfo field value if set, zero value otherwise.
func (o *AdapterUnit) GetInventoryDeviceInfo() InventoryDeviceInfoRelationship {
	if o == nil || o.InventoryDeviceInfo == nil {
		var ret InventoryDeviceInfoRelationship
		return ret
	}
	return *o.InventoryDeviceInfo
}

// GetInventoryDeviceInfoOk returns a tuple with the InventoryDeviceInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdapterUnit) GetInventoryDeviceInfoOk() (*InventoryDeviceInfoRelationship, bool) {
	if o == nil || o.InventoryDeviceInfo == nil {
		return nil, false
	}
	return o.InventoryDeviceInfo, true
}

// HasInventoryDeviceInfo returns a boolean if a field has been set.
func (o *AdapterUnit) HasInventoryDeviceInfo() bool {
	if o != nil && o.InventoryDeviceInfo != nil {
		return true
	}

	return false
}

// SetInventoryDeviceInfo gets a reference to the given InventoryDeviceInfoRelationship and assigns it to the InventoryDeviceInfo field.
func (o *AdapterUnit) SetInventoryDeviceInfo(v InventoryDeviceInfoRelationship) {
	o.InventoryDeviceInfo = &v
}

// GetRegisteredDevice returns the RegisteredDevice field value if set, zero value otherwise.
func (o *AdapterUnit) GetRegisteredDevice() AssetDeviceRegistrationRelationship {
	if o == nil || o.RegisteredDevice == nil {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.RegisteredDevice
}

// GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdapterUnit) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil || o.RegisteredDevice == nil {
		return nil, false
	}
	return o.RegisteredDevice, true
}

// HasRegisteredDevice returns a boolean if a field has been set.
func (o *AdapterUnit) HasRegisteredDevice() bool {
	if o != nil && o.RegisteredDevice != nil {
		return true
	}

	return false
}

// SetRegisteredDevice gets a reference to the given AssetDeviceRegistrationRelationship and assigns it to the RegisteredDevice field.
func (o *AdapterUnit) SetRegisteredDevice(v AssetDeviceRegistrationRelationship) {
	o.RegisteredDevice = &v
}

func (o AdapterUnit) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedEquipmentBase, errEquipmentBase := json.Marshal(o.EquipmentBase)
	if errEquipmentBase != nil {
		return []byte{}, errEquipmentBase
	}
	errEquipmentBase = json.Unmarshal([]byte(serializedEquipmentBase), &toSerialize)
	if errEquipmentBase != nil {
		return []byte{}, errEquipmentBase
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.AdapterId != nil {
		toSerialize["AdapterId"] = o.AdapterId
	}
	if o.BaseMacAddress != nil {
		toSerialize["BaseMacAddress"] = o.BaseMacAddress
	}
	if o.ConnectionStatus != nil {
		toSerialize["ConnectionStatus"] = o.ConnectionStatus
	}
	if o.Integrated != nil {
		toSerialize["Integrated"] = o.Integrated
	}
	if o.OperState != nil {
		toSerialize["OperState"] = o.OperState
	}
	if o.Operability != nil {
		toSerialize["Operability"] = o.Operability
	}
	if o.PartNumber != nil {
		toSerialize["PartNumber"] = o.PartNumber
	}
	if o.PciSlot != nil {
		toSerialize["PciSlot"] = o.PciSlot
	}
	if o.Power != nil {
		toSerialize["Power"] = o.Power
	}
	if o.Thermal != nil {
		toSerialize["Thermal"] = o.Thermal
	}
	if o.Vid != nil {
		toSerialize["Vid"] = o.Vid
	}
	if o.AdapterUnitExpander != nil {
		toSerialize["AdapterUnitExpander"] = o.AdapterUnitExpander
	}
	if o.ComputeBlade != nil {
		toSerialize["ComputeBlade"] = o.ComputeBlade
	}
	if o.ComputeRackUnit != nil {
		toSerialize["ComputeRackUnit"] = o.ComputeRackUnit
	}
	if o.Controller != nil {
		toSerialize["Controller"] = o.Controller
	}
	if o.ExtEthIfs != nil {
		toSerialize["ExtEthIfs"] = o.ExtEthIfs
	}
	if o.HostEthIfs != nil {
		toSerialize["HostEthIfs"] = o.HostEthIfs
	}
	if o.HostFcIfs != nil {
		toSerialize["HostFcIfs"] = o.HostFcIfs
	}
	if o.HostIscsiIfs != nil {
		toSerialize["HostIscsiIfs"] = o.HostIscsiIfs
	}
	if o.InventoryDeviceInfo != nil {
		toSerialize["InventoryDeviceInfo"] = o.InventoryDeviceInfo
	}
	if o.RegisteredDevice != nil {
		toSerialize["RegisteredDevice"] = o.RegisteredDevice
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *AdapterUnit) UnmarshalJSON(bytes []byte) (err error) {
	type AdapterUnitWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Unique Identifier of an adapter Unit within a Rack Interface.
		AdapterId *string `json:"AdapterId,omitempty"`
		// Original Base Mac address of an adapter unit.
		BaseMacAddress *string `json:"BaseMacAddress,omitempty"`
		// Connectivity Status of adapter - A or B or AB.
		ConnectionStatus *string `json:"ConnectionStatus,omitempty"`
		// Cisco Integrated adapter or other type.
		Integrated *string `json:"Integrated,omitempty"`
		// Operational state of an adapter unit.
		OperState *string `json:"OperState,omitempty"`
		// Operability state of an adapter unit.
		Operability *string `json:"Operability,omitempty"`
		// Part number of an adapter unit.
		PartNumber *string `json:"PartNumber,omitempty"`
		// PCIe slot of the adapter in the server.
		PciSlot *string `json:"PciSlot,omitempty"`
		// Power state of an adapter unit.
		Power *string `json:"Power,omitempty"`
		// Thermal state of an adapter unit.
		Thermal *string `json:"Thermal,omitempty"`
		// Virtual Id of the adapter in the server.
		Vid                 *string                           `json:"Vid,omitempty"`
		AdapterUnitExpander *AdapterUnitExpanderRelationship  `json:"AdapterUnitExpander,omitempty"`
		ComputeBlade        *ComputeBladeRelationship         `json:"ComputeBlade,omitempty"`
		ComputeRackUnit     *ComputeRackUnitRelationship      `json:"ComputeRackUnit,omitempty"`
		Controller          *ManagementControllerRelationship `json:"Controller,omitempty"`
		// An array of relationships to adapterExtEthInterface resources.
		ExtEthIfs []AdapterExtEthInterfaceRelationship `json:"ExtEthIfs,omitempty"`
		// An array of relationships to adapterHostEthInterface resources.
		HostEthIfs []AdapterHostEthInterfaceRelationship `json:"HostEthIfs,omitempty"`
		// An array of relationships to adapterHostFcInterface resources.
		HostFcIfs []AdapterHostFcInterfaceRelationship `json:"HostFcIfs,omitempty"`
		// An array of relationships to adapterHostIscsiInterface resources.
		HostIscsiIfs        []AdapterHostIscsiInterfaceRelationship `json:"HostIscsiIfs,omitempty"`
		InventoryDeviceInfo *InventoryDeviceInfoRelationship        `json:"InventoryDeviceInfo,omitempty"`
		RegisteredDevice    *AssetDeviceRegistrationRelationship    `json:"RegisteredDevice,omitempty"`
	}

	varAdapterUnitWithoutEmbeddedStruct := AdapterUnitWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varAdapterUnitWithoutEmbeddedStruct)
	if err == nil {
		varAdapterUnit := _AdapterUnit{}
		varAdapterUnit.ClassId = varAdapterUnitWithoutEmbeddedStruct.ClassId
		varAdapterUnit.ObjectType = varAdapterUnitWithoutEmbeddedStruct.ObjectType
		varAdapterUnit.AdapterId = varAdapterUnitWithoutEmbeddedStruct.AdapterId
		varAdapterUnit.BaseMacAddress = varAdapterUnitWithoutEmbeddedStruct.BaseMacAddress
		varAdapterUnit.ConnectionStatus = varAdapterUnitWithoutEmbeddedStruct.ConnectionStatus
		varAdapterUnit.Integrated = varAdapterUnitWithoutEmbeddedStruct.Integrated
		varAdapterUnit.OperState = varAdapterUnitWithoutEmbeddedStruct.OperState
		varAdapterUnit.Operability = varAdapterUnitWithoutEmbeddedStruct.Operability
		varAdapterUnit.PartNumber = varAdapterUnitWithoutEmbeddedStruct.PartNumber
		varAdapterUnit.PciSlot = varAdapterUnitWithoutEmbeddedStruct.PciSlot
		varAdapterUnit.Power = varAdapterUnitWithoutEmbeddedStruct.Power
		varAdapterUnit.Thermal = varAdapterUnitWithoutEmbeddedStruct.Thermal
		varAdapterUnit.Vid = varAdapterUnitWithoutEmbeddedStruct.Vid
		varAdapterUnit.AdapterUnitExpander = varAdapterUnitWithoutEmbeddedStruct.AdapterUnitExpander
		varAdapterUnit.ComputeBlade = varAdapterUnitWithoutEmbeddedStruct.ComputeBlade
		varAdapterUnit.ComputeRackUnit = varAdapterUnitWithoutEmbeddedStruct.ComputeRackUnit
		varAdapterUnit.Controller = varAdapterUnitWithoutEmbeddedStruct.Controller
		varAdapterUnit.ExtEthIfs = varAdapterUnitWithoutEmbeddedStruct.ExtEthIfs
		varAdapterUnit.HostEthIfs = varAdapterUnitWithoutEmbeddedStruct.HostEthIfs
		varAdapterUnit.HostFcIfs = varAdapterUnitWithoutEmbeddedStruct.HostFcIfs
		varAdapterUnit.HostIscsiIfs = varAdapterUnitWithoutEmbeddedStruct.HostIscsiIfs
		varAdapterUnit.InventoryDeviceInfo = varAdapterUnitWithoutEmbeddedStruct.InventoryDeviceInfo
		varAdapterUnit.RegisteredDevice = varAdapterUnitWithoutEmbeddedStruct.RegisteredDevice
		*o = AdapterUnit(varAdapterUnit)
	} else {
		return err
	}

	varAdapterUnit := _AdapterUnit{}

	err = json.Unmarshal(bytes, &varAdapterUnit)
	if err == nil {
		o.EquipmentBase = varAdapterUnit.EquipmentBase
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "AdapterId")
		delete(additionalProperties, "BaseMacAddress")
		delete(additionalProperties, "ConnectionStatus")
		delete(additionalProperties, "Integrated")
		delete(additionalProperties, "OperState")
		delete(additionalProperties, "Operability")
		delete(additionalProperties, "PartNumber")
		delete(additionalProperties, "PciSlot")
		delete(additionalProperties, "Power")
		delete(additionalProperties, "Thermal")
		delete(additionalProperties, "Vid")
		delete(additionalProperties, "AdapterUnitExpander")
		delete(additionalProperties, "ComputeBlade")
		delete(additionalProperties, "ComputeRackUnit")
		delete(additionalProperties, "Controller")
		delete(additionalProperties, "ExtEthIfs")
		delete(additionalProperties, "HostEthIfs")
		delete(additionalProperties, "HostFcIfs")
		delete(additionalProperties, "HostIscsiIfs")
		delete(additionalProperties, "InventoryDeviceInfo")
		delete(additionalProperties, "RegisteredDevice")

		// remove fields from embedded structs
		reflectEquipmentBase := reflect.ValueOf(o.EquipmentBase)
		for i := 0; i < reflectEquipmentBase.Type().NumField(); i++ {
			t := reflectEquipmentBase.Type().Field(i)

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

type NullableAdapterUnit struct {
	value *AdapterUnit
	isSet bool
}

func (v NullableAdapterUnit) Get() *AdapterUnit {
	return v.value
}

func (v *NullableAdapterUnit) Set(val *AdapterUnit) {
	v.value = val
	v.isSet = true
}

func (v NullableAdapterUnit) IsSet() bool {
	return v.isSet
}

func (v *NullableAdapterUnit) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAdapterUnit(val *AdapterUnit) *NullableAdapterUnit {
	return &NullableAdapterUnit{value: val, isSet: true}
}

func (v NullableAdapterUnit) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAdapterUnit) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
