/*
 * Cisco Intersight
 *
 * Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document. This document was created on 2021-06-30T12:14:04Z.
 *
 * API version: 1.0.9-4375
 * Contact: intersight@cisco.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package intersight

import (
	"encoding/json"
	"fmt"
)

// ForecastInstanceRelationship - A relationship to the 'forecast.Instance' resource, or the expanded 'forecast.Instance' resource, or the 'null' value.
type ForecastInstanceRelationship struct {
	ForecastInstance *ForecastInstance
	MoMoRef          *MoMoRef
}

// ForecastInstanceAsForecastInstanceRelationship is a convenience function that returns ForecastInstance wrapped in ForecastInstanceRelationship
func ForecastInstanceAsForecastInstanceRelationship(v *ForecastInstance) ForecastInstanceRelationship {
	return ForecastInstanceRelationship{ForecastInstance: v}
}

// MoMoRefAsForecastInstanceRelationship is a convenience function that returns MoMoRef wrapped in ForecastInstanceRelationship
func MoMoRefAsForecastInstanceRelationship(v *MoMoRef) ForecastInstanceRelationship {
	return ForecastInstanceRelationship{MoMoRef: v}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *ForecastInstanceRelationship) UnmarshalJSON(data []byte) error {
	var err error
	// this object is nullable so check if the payload is null or empty string
	if string(data) == "" || string(data) == "{}" {
		return nil
	}

	// use discriminator value to speed up the lookup
	var jsonDict map[string]interface{}
	err = json.Unmarshal(data, &jsonDict)
	if err != nil {
		return fmt.Errorf("Failed to unmarshal JSON into map for the discrimintor lookup.")
	}

	// check if the discriminator value is 'forecast.Instance'
	if jsonDict["ClassId"] == "forecast.Instance" {
		// try to unmarshal JSON data into ForecastInstance
		err = json.Unmarshal(data, &dst.ForecastInstance)
		if err == nil {
			return nil // data stored in dst.ForecastInstance, return on the first match
		} else {
			dst.ForecastInstance = nil
			return fmt.Errorf("Failed to unmarshal ForecastInstanceRelationship as ForecastInstance: %s", err.Error())
		}
	}

	// check if the discriminator value is 'mo.MoRef'
	if jsonDict["ClassId"] == "mo.MoRef" {
		// try to unmarshal JSON data into MoMoRef
		err = json.Unmarshal(data, &dst.MoMoRef)
		if err == nil {
			return nil // data stored in dst.MoMoRef, return on the first match
		} else {
			dst.MoMoRef = nil
			return fmt.Errorf("Failed to unmarshal ForecastInstanceRelationship as MoMoRef: %s", err.Error())
		}
	}

	return nil
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src ForecastInstanceRelationship) MarshalJSON() ([]byte, error) {
	if src.ForecastInstance != nil {
		return json.Marshal(&src.ForecastInstance)
	}

	if src.MoMoRef != nil {
		return json.Marshal(&src.MoMoRef)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *ForecastInstanceRelationship) GetActualInstance() interface{} {
	if obj.ForecastInstance != nil {
		return obj.ForecastInstance
	}

	if obj.MoMoRef != nil {
		return obj.MoMoRef
	}

	// all schemas are nil
	return nil
}

type NullableForecastInstanceRelationship struct {
	value *ForecastInstanceRelationship
	isSet bool
}

func (v NullableForecastInstanceRelationship) Get() *ForecastInstanceRelationship {
	return v.value
}

func (v *NullableForecastInstanceRelationship) Set(val *ForecastInstanceRelationship) {
	v.value = val
	v.isSet = true
}

func (v NullableForecastInstanceRelationship) IsSet() bool {
	return v.isSet
}

func (v *NullableForecastInstanceRelationship) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableForecastInstanceRelationship(val *ForecastInstanceRelationship) *NullableForecastInstanceRelationship {
	return &NullableForecastInstanceRelationship{value: val, isSet: true}
}

func (v NullableForecastInstanceRelationship) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableForecastInstanceRelationship) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
