package intersight

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"reflect"
	"strings"
	"time"

	models "github.com/CiscoDevNet/terraform-provider-intersight/intersight_gosdk"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceCapabilitySwitchCapability() *schema.Resource {
	var subSchema = map[string]*schema.Schema{"account_moid": {
		Description: "The Account ID for this managed object.",
		Type:        schema.TypeString,
		Optional:    true,
	},
		"additional_properties": {
			Type:             schema.TypeString,
			Optional:         true,
			DiffSuppressFunc: SuppressDiffAdditionProps,
		},
		"ancestors": {
			Description: "An array of relationships to moBaseMo resources.",
			Type:        schema.TypeList,
			Optional:    true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"additional_properties": {
						Type:             schema.TypeString,
						Optional:         true,
						DiffSuppressFunc: SuppressDiffAdditionProps,
					},
					"class_id": {
						Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"moid": {
						Description: "The Moid of the referenced REST resource.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"object_type": {
						Description: "The fully-qualified name of the remote type referred by this relationship.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"selector": {
						Description: "An OData $filter expression which describes the REST resource to be referenced. This field may\nbe set instead of 'moid' by clients.\n1. If 'moid' is set this field is ignored.\n1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of the\nresource matching the filter expression and populates it in the MoRef that is part of the object\ninstance being inserted/updated to fulfill the REST request.\nAn error is returned if the filter matches zero or more than one REST resource.\nAn example filter string is: Serial eq '3AA8B7T11'.",
						Type:        schema.TypeString,
						Optional:    true,
					},
				},
			},
		},
		"class_id": {
			Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"create_time": {
			Description: "The time when this managed object was created.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"default_fcoe_vlan": {
			Description: "Default Fcoe VLAN associated with this switch.",
			Type:        schema.TypeInt,
			Optional:    true,
		},
		"domain_group_moid": {
			Description: "The DomainGroup ID for this managed object.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"dynamic_vifs_supported": {
			Description: "Dynamic VIFs support on this switch.",
			Type:        schema.TypeBool,
			Optional:    true,
		},
		"fan_modules_supported": {
			Description: "Fan Modules support on this switch.",
			Type:        schema.TypeBool,
			Optional:    true,
		},
		"fc_end_host_mode_reserved_vsans": {
			Type:     schema.TypeList,
			Optional: true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"additional_properties": {
						Type:             schema.TypeString,
						Optional:         true,
						DiffSuppressFunc: SuppressDiffAdditionProps,
					},
					"class_id": {
						Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"end_port_id": {
						Description: "Ending Port ID in this range of ports.",
						Type:        schema.TypeInt,
						Optional:    true,
					},
					"end_slot_id": {
						Description: "Ending Slot ID in this range of ports.",
						Type:        schema.TypeInt,
						Optional:    true,
					},
					"object_type": {
						Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"start_port_id": {
						Description: "Starting Port ID in this range of ports.",
						Type:        schema.TypeInt,
						Optional:    true,
					},
					"start_slot_id": {
						Description: "Starting Slot ID in this range of ports.",
						Type:        schema.TypeInt,
						Optional:    true,
					},
				},
			},
		},
		"fc_uplink_ports_auto_negotiation_supported": {
			Description: "Fc Uplink ports auto negotiation speed support on this switch.",
			Type:        schema.TypeBool,
			Optional:    true,
		},
		"locator_beacon_supported": {
			Description: "Locator Beacon LED support on this switch.",
			Type:        schema.TypeBool,
			Optional:    true,
		},
		"max_ports": {
			Description: "Maximum allowed physical ports on this switch.",
			Type:        schema.TypeInt,
			Optional:    true,
		},
		"max_slots": {
			Description: "Maximum allowed physical slots on this switch.",
			Type:        schema.TypeInt,
			Optional:    true,
		},
		"mod_time": {
			Description: "The time when this managed object was last modified.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"moid": {
			Description: "The unique identifier of this Managed Object instance.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"name": {
			Description: "An unique identifer for a capability descriptor.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"network_limits": {
			Description: "List of network limitations for this switch.",
			Type:        schema.TypeList,
			MaxItems:    1,
			Optional:    true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"additional_properties": {
						Type:             schema.TypeString,
						Optional:         true,
						DiffSuppressFunc: SuppressDiffAdditionProps,
					},
					"class_id": {
						Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"max_compressed_port_vlan_count": {
						Description: "Maximum Compressed configurable VLANs on Switch/Fabric-Interconnect.",
						Type:        schema.TypeInt,
						Optional:    true,
					},
					"max_uncompressed_port_vlan_count": {
						Description: "Maximum configurable VLANs on Switch/Fabric-Interconnect.",
						Type:        schema.TypeInt,
						Optional:    true,
					},
					"maximum_active_traffic_monitoring_sessions": {
						Description: "Maximum configured and enabled Traffic Monitoring sessions on this Switch/Fabric-Interconnect.",
						Type:        schema.TypeInt,
						Optional:    true,
					},
					"maximum_ethernet_port_channels": {
						Description: "Maximum configurable Ethernet port-channels on Switch/Fabric-Interconnect.",
						Type:        schema.TypeInt,
						Optional:    true,
					},
					"maximum_ethernet_uplink_ports": {
						Description: "Maximum configurable Ethernet Uplink ports on Switch/Fabric-Interconnect.",
						Type:        schema.TypeInt,
						Optional:    true,
					},
					"maximum_fc_port_channel_members": {
						Description: "Maximum configurable Fibre Channel port-channel member ports on Switch/Fabric-Interconnect.",
						Type:        schema.TypeInt,
						Optional:    true,
					},
					"maximum_fc_port_channels": {
						Description: "Maximum configurable Fibre Channel port-channels on Switch/Fabric-Interconnect.",
						Type:        schema.TypeInt,
						Optional:    true,
					},
					"maximum_igmp_groups": {
						Description: "Maximum configurable IGMP Groups on Switch/Fabric-Interconnect.",
						Type:        schema.TypeInt,
						Optional:    true,
					},
					"maximum_port_channel_members": {
						Description: "Maximum configurable ports per each port-channel on Switch/Fabric-Interconnect.",
						Type:        schema.TypeInt,
						Optional:    true,
					},
					"maximum_primary_vlan": {
						Description: "Maximum configurable Primary Private VLANs on Switch/Fabric-Interconnect.",
						Type:        schema.TypeInt,
						Optional:    true,
					},
					"maximum_secondary_vlan": {
						Description: "Maximum configurable Secondary Private VLANs on Switch/Fabric-Interconnect.",
						Type:        schema.TypeInt,
						Optional:    true,
					},
					"maximum_secondary_vlan_per_primary": {
						Description: "Maximum configurable Secondary VLANs per each Primary VLAN on Switch/Fabric-Interconnect.",
						Type:        schema.TypeInt,
						Optional:    true,
					},
					"maximum_vifs": {
						Description: "Maximum allowes VIFs on Switch/Fabric-Interconnect.",
						Type:        schema.TypeInt,
						Optional:    true,
					},
					"maximum_vlans": {
						Description: "Maximum configurable VLANs on Switch/Fabric-Interconnect.",
						Type:        schema.TypeInt,
						Optional:    true,
					},
					"minimum_active_fans": {
						Description: "Minimum required fans in 'active' state for this Switch/Fabric-Interconnect.",
						Type:        schema.TypeInt,
						Optional:    true,
					},
					"object_type": {
						Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
						Type:        schema.TypeString,
						Optional:    true,
					},
				},
			},
		},
		"object_type": {
			Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"owners": {
			Type:     schema.TypeList,
			Optional: true,
			Elem: &schema.Schema{
				Type: schema.TypeString}},
		"parent": {
			Description: "A reference to a moBaseMo resource.\nWhen the $expand query parameter is specified, the referenced resource is returned inline.",
			Type:        schema.TypeList,
			MaxItems:    1,
			Optional:    true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"additional_properties": {
						Type:             schema.TypeString,
						Optional:         true,
						DiffSuppressFunc: SuppressDiffAdditionProps,
					},
					"class_id": {
						Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"moid": {
						Description: "The Moid of the referenced REST resource.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"object_type": {
						Description: "The fully-qualified name of the remote type referred by this relationship.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"selector": {
						Description: "An OData $filter expression which describes the REST resource to be referenced. This field may\nbe set instead of 'moid' by clients.\n1. If 'moid' is set this field is ignored.\n1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of the\nresource matching the filter expression and populates it in the MoRef that is part of the object\ninstance being inserted/updated to fulfill the REST request.\nAn error is returned if the filter matches zero or more than one REST resource.\nAn example filter string is: Serial eq '3AA8B7T11'.",
						Type:        schema.TypeString,
						Optional:    true,
					},
				},
			},
		},
		"permission_resources": {
			Description: "An array of relationships to moBaseMo resources.",
			Type:        schema.TypeList,
			Optional:    true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"additional_properties": {
						Type:             schema.TypeString,
						Optional:         true,
						DiffSuppressFunc: SuppressDiffAdditionProps,
					},
					"class_id": {
						Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"moid": {
						Description: "The Moid of the referenced REST resource.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"object_type": {
						Description: "The fully-qualified name of the remote type referred by this relationship.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"selector": {
						Description: "An OData $filter expression which describes the REST resource to be referenced. This field may\nbe set instead of 'moid' by clients.\n1. If 'moid' is set this field is ignored.\n1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of the\nresource matching the filter expression and populates it in the MoRef that is part of the object\ninstance being inserted/updated to fulfill the REST request.\nAn error is returned if the filter matches zero or more than one REST resource.\nAn example filter string is: Serial eq '3AA8B7T11'.",
						Type:        schema.TypeString,
						Optional:    true,
					},
				},
			},
		},
		"pid": {
			Description: "Product Identifier for a Switch/Fabric-Interconnect.\n* `UCS-FI-6454` - The standard 4th generation UCS Fabric Interconnect with 54 ports.\n* `UCS-FI-64108` - The expanded 4th generation UCS Fabric Interconnect with 108 ports.\n* `unknown` - Unknown device type, usage is TBD.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"ports_supporting100g_speed": {
			Type:     schema.TypeList,
			Optional: true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"additional_properties": {
						Type:             schema.TypeString,
						Optional:         true,
						DiffSuppressFunc: SuppressDiffAdditionProps,
					},
					"class_id": {
						Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"end_port_id": {
						Description: "Ending Port ID in this range of ports.",
						Type:        schema.TypeInt,
						Optional:    true,
					},
					"end_slot_id": {
						Description: "Ending Slot ID in this range of ports.",
						Type:        schema.TypeInt,
						Optional:    true,
					},
					"object_type": {
						Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"start_port_id": {
						Description: "Starting Port ID in this range of ports.",
						Type:        schema.TypeInt,
						Optional:    true,
					},
					"start_slot_id": {
						Description: "Starting Slot ID in this range of ports.",
						Type:        schema.TypeInt,
						Optional:    true,
					},
				},
			},
		},
		"ports_supporting10g_speed": {
			Type:     schema.TypeList,
			Optional: true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"additional_properties": {
						Type:             schema.TypeString,
						Optional:         true,
						DiffSuppressFunc: SuppressDiffAdditionProps,
					},
					"class_id": {
						Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"end_port_id": {
						Description: "Ending Port ID in this range of ports.",
						Type:        schema.TypeInt,
						Optional:    true,
					},
					"end_slot_id": {
						Description: "Ending Slot ID in this range of ports.",
						Type:        schema.TypeInt,
						Optional:    true,
					},
					"object_type": {
						Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"start_port_id": {
						Description: "Starting Port ID in this range of ports.",
						Type:        schema.TypeInt,
						Optional:    true,
					},
					"start_slot_id": {
						Description: "Starting Slot ID in this range of ports.",
						Type:        schema.TypeInt,
						Optional:    true,
					},
				},
			},
		},
		"ports_supporting1g_speed": {
			Type:     schema.TypeList,
			Optional: true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"additional_properties": {
						Type:             schema.TypeString,
						Optional:         true,
						DiffSuppressFunc: SuppressDiffAdditionProps,
					},
					"class_id": {
						Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"end_port_id": {
						Description: "Ending Port ID in this range of ports.",
						Type:        schema.TypeInt,
						Optional:    true,
					},
					"end_slot_id": {
						Description: "Ending Slot ID in this range of ports.",
						Type:        schema.TypeInt,
						Optional:    true,
					},
					"object_type": {
						Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"start_port_id": {
						Description: "Starting Port ID in this range of ports.",
						Type:        schema.TypeInt,
						Optional:    true,
					},
					"start_slot_id": {
						Description: "Starting Slot ID in this range of ports.",
						Type:        schema.TypeInt,
						Optional:    true,
					},
				},
			},
		},
		"ports_supporting25g_speed": {
			Type:     schema.TypeList,
			Optional: true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"additional_properties": {
						Type:             schema.TypeString,
						Optional:         true,
						DiffSuppressFunc: SuppressDiffAdditionProps,
					},
					"class_id": {
						Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"end_port_id": {
						Description: "Ending Port ID in this range of ports.",
						Type:        schema.TypeInt,
						Optional:    true,
					},
					"end_slot_id": {
						Description: "Ending Slot ID in this range of ports.",
						Type:        schema.TypeInt,
						Optional:    true,
					},
					"object_type": {
						Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"start_port_id": {
						Description: "Starting Port ID in this range of ports.",
						Type:        schema.TypeInt,
						Optional:    true,
					},
					"start_slot_id": {
						Description: "Starting Slot ID in this range of ports.",
						Type:        schema.TypeInt,
						Optional:    true,
					},
				},
			},
		},
		"ports_supporting40g_speed": {
			Type:     schema.TypeList,
			Optional: true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"additional_properties": {
						Type:             schema.TypeString,
						Optional:         true,
						DiffSuppressFunc: SuppressDiffAdditionProps,
					},
					"class_id": {
						Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"end_port_id": {
						Description: "Ending Port ID in this range of ports.",
						Type:        schema.TypeInt,
						Optional:    true,
					},
					"end_slot_id": {
						Description: "Ending Slot ID in this range of ports.",
						Type:        schema.TypeInt,
						Optional:    true,
					},
					"object_type": {
						Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"start_port_id": {
						Description: "Starting Port ID in this range of ports.",
						Type:        schema.TypeInt,
						Optional:    true,
					},
					"start_slot_id": {
						Description: "Starting Slot ID in this range of ports.",
						Type:        schema.TypeInt,
						Optional:    true,
					},
				},
			},
		},
		"ports_supporting_breakout": {
			Type:     schema.TypeList,
			Optional: true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"additional_properties": {
						Type:             schema.TypeString,
						Optional:         true,
						DiffSuppressFunc: SuppressDiffAdditionProps,
					},
					"class_id": {
						Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"end_port_id": {
						Description: "Ending Port ID in this range of ports.",
						Type:        schema.TypeInt,
						Optional:    true,
					},
					"end_slot_id": {
						Description: "Ending Slot ID in this range of ports.",
						Type:        schema.TypeInt,
						Optional:    true,
					},
					"object_type": {
						Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"start_port_id": {
						Description: "Starting Port ID in this range of ports.",
						Type:        schema.TypeInt,
						Optional:    true,
					},
					"start_slot_id": {
						Description: "Starting Slot ID in this range of ports.",
						Type:        schema.TypeInt,
						Optional:    true,
					},
				},
			},
		},
		"ports_supporting_fcoe": {
			Type:     schema.TypeList,
			Optional: true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"additional_properties": {
						Type:             schema.TypeString,
						Optional:         true,
						DiffSuppressFunc: SuppressDiffAdditionProps,
					},
					"class_id": {
						Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"end_port_id": {
						Description: "Ending Port ID in this range of ports.",
						Type:        schema.TypeInt,
						Optional:    true,
					},
					"end_slot_id": {
						Description: "Ending Slot ID in this range of ports.",
						Type:        schema.TypeInt,
						Optional:    true,
					},
					"object_type": {
						Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"start_port_id": {
						Description: "Starting Port ID in this range of ports.",
						Type:        schema.TypeInt,
						Optional:    true,
					},
					"start_slot_id": {
						Description: "Starting Slot ID in this range of ports.",
						Type:        schema.TypeInt,
						Optional:    true,
					},
				},
			},
		},
		"ports_supporting_server_role": {
			Type:     schema.TypeList,
			Optional: true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"additional_properties": {
						Type:             schema.TypeString,
						Optional:         true,
						DiffSuppressFunc: SuppressDiffAdditionProps,
					},
					"class_id": {
						Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"end_port_id": {
						Description: "Ending Port ID in this range of ports.",
						Type:        schema.TypeInt,
						Optional:    true,
					},
					"end_slot_id": {
						Description: "Ending Slot ID in this range of ports.",
						Type:        schema.TypeInt,
						Optional:    true,
					},
					"object_type": {
						Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"start_port_id": {
						Description: "Starting Port ID in this range of ports.",
						Type:        schema.TypeInt,
						Optional:    true,
					},
					"start_slot_id": {
						Description: "Starting Slot ID in this range of ports.",
						Type:        schema.TypeInt,
						Optional:    true,
					},
				},
			},
		},
		"reserved_vsans": {
			Type:     schema.TypeList,
			Optional: true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"additional_properties": {
						Type:             schema.TypeString,
						Optional:         true,
						DiffSuppressFunc: SuppressDiffAdditionProps,
					},
					"class_id": {
						Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"end_port_id": {
						Description: "Ending Port ID in this range of ports.",
						Type:        schema.TypeInt,
						Optional:    true,
					},
					"end_slot_id": {
						Description: "Ending Slot ID in this range of ports.",
						Type:        schema.TypeInt,
						Optional:    true,
					},
					"object_type": {
						Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"start_port_id": {
						Description: "Starting Port ID in this range of ports.",
						Type:        schema.TypeInt,
						Optional:    true,
					},
					"start_slot_id": {
						Description: "Starting Slot ID in this range of ports.",
						Type:        schema.TypeInt,
						Optional:    true,
					},
				},
			},
		},
		"sereno_netflow_supported": {
			Description: "Sereno Adaptor with Netflow support on this switch.",
			Type:        schema.TypeBool,
			Optional:    true,
		},
		"shared_scope": {
			Description: "Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.\nObjects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"sku": {
			Description: "SKU information for Switch/Fabric-Interconnect.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"storage_limits": {
			Description: "List of storage limitations for this switch.",
			Type:        schema.TypeList,
			MaxItems:    1,
			Optional:    true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"additional_properties": {
						Type:             schema.TypeString,
						Optional:         true,
						DiffSuppressFunc: SuppressDiffAdditionProps,
					},
					"class_id": {
						Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"maximum_user_zone_count": {
						Description: "Maximum user zones per Switch/Fabric-Interconnect.",
						Type:        schema.TypeInt,
						Optional:    true,
					},
					"maximum_virtual_fc_interfaces": {
						Description: "Maximum configurable Virtual Fibre Channel interfaces on Switch/Fabric-Interconnect.",
						Type:        schema.TypeInt,
						Optional:    true,
					},
					"maximum_virtual_fc_interfaces_per_blade_server": {
						Description: "Maximum configurable Virtual Fibre Channel interfaces per blade.",
						Type:        schema.TypeInt,
						Optional:    true,
					},
					"maximum_vsans": {
						Description: "Maximum configurable VSANs on Switch/Fabric-Interconnect.",
						Type:        schema.TypeInt,
						Optional:    true,
					},
					"maximum_zone_count": {
						Description: "Zone limit per Switch/Fabric-Interconnect.",
						Type:        schema.TypeInt,
						Optional:    true,
					},
					"object_type": {
						Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
						Type:        schema.TypeString,
						Optional:    true,
					},
				},
			},
		},
		"switching_mode_capabilities": {
			Type:     schema.TypeList,
			Optional: true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"additional_properties": {
						Type:             schema.TypeString,
						Optional:         true,
						DiffSuppressFunc: SuppressDiffAdditionProps,
					},
					"class_id": {
						Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"object_type": {
						Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"switching_mode": {
						Description: "Switching mode type (endhost, switch) of the switch.\n* `end-host` - In end-host mode, the fabric interconnects appear to the upstream devices as end hosts with multiple links.In this mode, the switch does not run Spanning Tree Protocol and avoids loops by following a set of rules for traffic forwarding.In case of ethernet switching mode - Ethernet end-host mode is also known as Ethernet host virtualizer.\n* `switch` - In switch mode, the switch runs Spanning Tree Protocol to avoid loops, and broadcast and multicast packets are handled in the traditional way.This is the traditional switch mode.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"vp_compression_supported": {
						Description: "VP Compression support on this switch.",
						Type:        schema.TypeBool,
						Optional:    true,
					},
				},
			},
		},
		"system_limits": {
			Description: "List of system limitations for this switch.",
			Type:        schema.TypeList,
			MaxItems:    1,
			Optional:    true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"additional_properties": {
						Type:             schema.TypeString,
						Optional:         true,
						DiffSuppressFunc: SuppressDiffAdditionProps,
					},
					"class_id": {
						Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"maximum_chassis_count": {
						Description: "Maximum UCS chassis that can be connected to this Switch/Fabric-Interconnect.",
						Type:        schema.TypeInt,
						Optional:    true,
					},
					"maximum_fex_per_domain": {
						Description: "Maximum UCS Fabric-extenders (FEX) per Switch/Fabric-Interconnect.",
						Type:        schema.TypeInt,
						Optional:    true,
					},
					"maximum_servers_per_domain": {
						Description: "Maximum UCS servers per Switch/Fabric-Interconnect.",
						Type:        schema.TypeInt,
						Optional:    true,
					},
					"object_type": {
						Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
						Type:        schema.TypeString,
						Optional:    true,
					},
				},
			},
		},
		"tags": {
			Type:     schema.TypeList,
			Optional: true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"additional_properties": {
						Type:             schema.TypeString,
						Optional:         true,
						DiffSuppressFunc: SuppressDiffAdditionProps,
					},
					"key": {
						Description: "The string representation of a tag key.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"value": {
						Description: "The string representation of a tag value.",
						Type:        schema.TypeString,
						Optional:    true,
					},
				},
			},
		},
		"unified_ports": {
			Type:     schema.TypeList,
			Optional: true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"additional_properties": {
						Type:             schema.TypeString,
						Optional:         true,
						DiffSuppressFunc: SuppressDiffAdditionProps,
					},
					"class_id": {
						Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"end_port_id": {
						Description: "Ending Port ID in this range of ports.",
						Type:        schema.TypeInt,
						Optional:    true,
					},
					"end_slot_id": {
						Description: "Ending Slot ID in this range of ports.",
						Type:        schema.TypeInt,
						Optional:    true,
					},
					"object_type": {
						Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"start_port_id": {
						Description: "Starting Port ID in this range of ports.",
						Type:        schema.TypeInt,
						Optional:    true,
					},
					"start_slot_id": {
						Description: "Starting Slot ID in this range of ports.",
						Type:        schema.TypeInt,
						Optional:    true,
					},
				},
			},
		},
		"unified_rule": {
			Description: "The Slider rule for Unified ports on this switch.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"version_context": {
			Description: "The versioning info for this managed object.",
			Type:        schema.TypeList,
			MaxItems:    1,
			Optional:    true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"additional_properties": {
						Type:             schema.TypeString,
						Optional:         true,
						DiffSuppressFunc: SuppressDiffAdditionProps,
					},
					"class_id": {
						Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"interested_mos": {
						Type:     schema.TypeList,
						Optional: true,
						Elem: &schema.Resource{
							Schema: map[string]*schema.Schema{
								"additional_properties": {
									Type:             schema.TypeString,
									Optional:         true,
									DiffSuppressFunc: SuppressDiffAdditionProps,
								},
								"class_id": {
									Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
									Type:        schema.TypeString,
									Optional:    true,
								},
								"moid": {
									Description: "The Moid of the referenced REST resource.",
									Type:        schema.TypeString,
									Optional:    true,
								},
								"object_type": {
									Description: "The fully-qualified name of the remote type referred by this relationship.",
									Type:        schema.TypeString,
									Optional:    true,
								},
								"selector": {
									Description: "An OData $filter expression which describes the REST resource to be referenced. This field may\nbe set instead of 'moid' by clients.\n1. If 'moid' is set this field is ignored.\n1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of the\nresource matching the filter expression and populates it in the MoRef that is part of the object\ninstance being inserted/updated to fulfill the REST request.\nAn error is returned if the filter matches zero or more than one REST resource.\nAn example filter string is: Serial eq '3AA8B7T11'.",
									Type:        schema.TypeString,
									Optional:    true,
								},
							},
						},
					},
					"object_type": {
						Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"ref_mo": {
						Description: "A reference to the original Managed Object.",
						Type:        schema.TypeList,
						MaxItems:    1,
						Optional:    true,
						Elem: &schema.Resource{
							Schema: map[string]*schema.Schema{
								"additional_properties": {
									Type:             schema.TypeString,
									Optional:         true,
									DiffSuppressFunc: SuppressDiffAdditionProps,
								},
								"class_id": {
									Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
									Type:        schema.TypeString,
									Optional:    true,
								},
								"moid": {
									Description: "The Moid of the referenced REST resource.",
									Type:        schema.TypeString,
									Optional:    true,
								},
								"object_type": {
									Description: "The fully-qualified name of the remote type referred by this relationship.",
									Type:        schema.TypeString,
									Optional:    true,
								},
								"selector": {
									Description: "An OData $filter expression which describes the REST resource to be referenced. This field may\nbe set instead of 'moid' by clients.\n1. If 'moid' is set this field is ignored.\n1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of the\nresource matching the filter expression and populates it in the MoRef that is part of the object\ninstance being inserted/updated to fulfill the REST request.\nAn error is returned if the filter matches zero or more than one REST resource.\nAn example filter string is: Serial eq '3AA8B7T11'.",
									Type:        schema.TypeString,
									Optional:    true,
								},
							},
						},
					},
					"timestamp": {
						Description: "The time this versioned Managed Object was created.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"nr_version": {
						Description: "The version of the Managed Object, e.g. an incrementing number or a hash id.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"version_type": {
						Description: "Specifies type of version. Currently the only supported value is \"Configured\"\nthat is used to keep track of snapshots of policies and profiles that are intended\nto be configured to target endpoints.\n* `Modified` - Version created every time an object is modified.\n* `Configured` - Version created every time an object is configured to the service profile.\n* `Deployed` - Version created for objects related to a service profile when it is deployed.",
						Type:        schema.TypeString,
						Optional:    true,
					},
				},
			},
		},
		"vid": {
			Description: "VID information for Switch/Fabric-Interconnect.",
			Type:        schema.TypeString,
			Optional:    true,
		},
	}
	var model = map[string]*schema.Schema{"account_moid": {
		Description: "The Account ID for this managed object.",
		Type:        schema.TypeString,
		Optional:    true,
	},
		"additional_properties": {
			Type:             schema.TypeString,
			Optional:         true,
			DiffSuppressFunc: SuppressDiffAdditionProps,
		},
		"ancestors": {
			Description: "An array of relationships to moBaseMo resources.",
			Type:        schema.TypeList,
			Optional:    true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"additional_properties": {
						Type:             schema.TypeString,
						Optional:         true,
						DiffSuppressFunc: SuppressDiffAdditionProps,
					},
					"class_id": {
						Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"moid": {
						Description: "The Moid of the referenced REST resource.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"object_type": {
						Description: "The fully-qualified name of the remote type referred by this relationship.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"selector": {
						Description: "An OData $filter expression which describes the REST resource to be referenced. This field may\nbe set instead of 'moid' by clients.\n1. If 'moid' is set this field is ignored.\n1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of the\nresource matching the filter expression and populates it in the MoRef that is part of the object\ninstance being inserted/updated to fulfill the REST request.\nAn error is returned if the filter matches zero or more than one REST resource.\nAn example filter string is: Serial eq '3AA8B7T11'.",
						Type:        schema.TypeString,
						Optional:    true,
					},
				},
			},
		},
		"class_id": {
			Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"create_time": {
			Description: "The time when this managed object was created.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"default_fcoe_vlan": {
			Description: "Default Fcoe VLAN associated with this switch.",
			Type:        schema.TypeInt,
			Optional:    true,
		},
		"domain_group_moid": {
			Description: "The DomainGroup ID for this managed object.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"dynamic_vifs_supported": {
			Description: "Dynamic VIFs support on this switch.",
			Type:        schema.TypeBool,
			Optional:    true,
		},
		"fan_modules_supported": {
			Description: "Fan Modules support on this switch.",
			Type:        schema.TypeBool,
			Optional:    true,
		},
		"fc_end_host_mode_reserved_vsans": {
			Type:     schema.TypeList,
			Optional: true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"additional_properties": {
						Type:             schema.TypeString,
						Optional:         true,
						DiffSuppressFunc: SuppressDiffAdditionProps,
					},
					"class_id": {
						Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"end_port_id": {
						Description: "Ending Port ID in this range of ports.",
						Type:        schema.TypeInt,
						Optional:    true,
					},
					"end_slot_id": {
						Description: "Ending Slot ID in this range of ports.",
						Type:        schema.TypeInt,
						Optional:    true,
					},
					"object_type": {
						Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"start_port_id": {
						Description: "Starting Port ID in this range of ports.",
						Type:        schema.TypeInt,
						Optional:    true,
					},
					"start_slot_id": {
						Description: "Starting Slot ID in this range of ports.",
						Type:        schema.TypeInt,
						Optional:    true,
					},
				},
			},
		},
		"fc_uplink_ports_auto_negotiation_supported": {
			Description: "Fc Uplink ports auto negotiation speed support on this switch.",
			Type:        schema.TypeBool,
			Optional:    true,
		},
		"locator_beacon_supported": {
			Description: "Locator Beacon LED support on this switch.",
			Type:        schema.TypeBool,
			Optional:    true,
		},
		"max_ports": {
			Description: "Maximum allowed physical ports on this switch.",
			Type:        schema.TypeInt,
			Optional:    true,
		},
		"max_slots": {
			Description: "Maximum allowed physical slots on this switch.",
			Type:        schema.TypeInt,
			Optional:    true,
		},
		"mod_time": {
			Description: "The time when this managed object was last modified.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"moid": {
			Description: "The unique identifier of this Managed Object instance.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"name": {
			Description: "An unique identifer for a capability descriptor.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"network_limits": {
			Description: "List of network limitations for this switch.",
			Type:        schema.TypeList,
			MaxItems:    1,
			Optional:    true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"additional_properties": {
						Type:             schema.TypeString,
						Optional:         true,
						DiffSuppressFunc: SuppressDiffAdditionProps,
					},
					"class_id": {
						Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"max_compressed_port_vlan_count": {
						Description: "Maximum Compressed configurable VLANs on Switch/Fabric-Interconnect.",
						Type:        schema.TypeInt,
						Optional:    true,
					},
					"max_uncompressed_port_vlan_count": {
						Description: "Maximum configurable VLANs on Switch/Fabric-Interconnect.",
						Type:        schema.TypeInt,
						Optional:    true,
					},
					"maximum_active_traffic_monitoring_sessions": {
						Description: "Maximum configured and enabled Traffic Monitoring sessions on this Switch/Fabric-Interconnect.",
						Type:        schema.TypeInt,
						Optional:    true,
					},
					"maximum_ethernet_port_channels": {
						Description: "Maximum configurable Ethernet port-channels on Switch/Fabric-Interconnect.",
						Type:        schema.TypeInt,
						Optional:    true,
					},
					"maximum_ethernet_uplink_ports": {
						Description: "Maximum configurable Ethernet Uplink ports on Switch/Fabric-Interconnect.",
						Type:        schema.TypeInt,
						Optional:    true,
					},
					"maximum_fc_port_channel_members": {
						Description: "Maximum configurable Fibre Channel port-channel member ports on Switch/Fabric-Interconnect.",
						Type:        schema.TypeInt,
						Optional:    true,
					},
					"maximum_fc_port_channels": {
						Description: "Maximum configurable Fibre Channel port-channels on Switch/Fabric-Interconnect.",
						Type:        schema.TypeInt,
						Optional:    true,
					},
					"maximum_igmp_groups": {
						Description: "Maximum configurable IGMP Groups on Switch/Fabric-Interconnect.",
						Type:        schema.TypeInt,
						Optional:    true,
					},
					"maximum_port_channel_members": {
						Description: "Maximum configurable ports per each port-channel on Switch/Fabric-Interconnect.",
						Type:        schema.TypeInt,
						Optional:    true,
					},
					"maximum_primary_vlan": {
						Description: "Maximum configurable Primary Private VLANs on Switch/Fabric-Interconnect.",
						Type:        schema.TypeInt,
						Optional:    true,
					},
					"maximum_secondary_vlan": {
						Description: "Maximum configurable Secondary Private VLANs on Switch/Fabric-Interconnect.",
						Type:        schema.TypeInt,
						Optional:    true,
					},
					"maximum_secondary_vlan_per_primary": {
						Description: "Maximum configurable Secondary VLANs per each Primary VLAN on Switch/Fabric-Interconnect.",
						Type:        schema.TypeInt,
						Optional:    true,
					},
					"maximum_vifs": {
						Description: "Maximum allowes VIFs on Switch/Fabric-Interconnect.",
						Type:        schema.TypeInt,
						Optional:    true,
					},
					"maximum_vlans": {
						Description: "Maximum configurable VLANs on Switch/Fabric-Interconnect.",
						Type:        schema.TypeInt,
						Optional:    true,
					},
					"minimum_active_fans": {
						Description: "Minimum required fans in 'active' state for this Switch/Fabric-Interconnect.",
						Type:        schema.TypeInt,
						Optional:    true,
					},
					"object_type": {
						Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
						Type:        schema.TypeString,
						Optional:    true,
					},
				},
			},
		},
		"object_type": {
			Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"owners": {
			Type:     schema.TypeList,
			Optional: true,
			Elem: &schema.Schema{
				Type: schema.TypeString}},
		"parent": {
			Description: "A reference to a moBaseMo resource.\nWhen the $expand query parameter is specified, the referenced resource is returned inline.",
			Type:        schema.TypeList,
			MaxItems:    1,
			Optional:    true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"additional_properties": {
						Type:             schema.TypeString,
						Optional:         true,
						DiffSuppressFunc: SuppressDiffAdditionProps,
					},
					"class_id": {
						Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"moid": {
						Description: "The Moid of the referenced REST resource.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"object_type": {
						Description: "The fully-qualified name of the remote type referred by this relationship.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"selector": {
						Description: "An OData $filter expression which describes the REST resource to be referenced. This field may\nbe set instead of 'moid' by clients.\n1. If 'moid' is set this field is ignored.\n1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of the\nresource matching the filter expression and populates it in the MoRef that is part of the object\ninstance being inserted/updated to fulfill the REST request.\nAn error is returned if the filter matches zero or more than one REST resource.\nAn example filter string is: Serial eq '3AA8B7T11'.",
						Type:        schema.TypeString,
						Optional:    true,
					},
				},
			},
		},
		"permission_resources": {
			Description: "An array of relationships to moBaseMo resources.",
			Type:        schema.TypeList,
			Optional:    true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"additional_properties": {
						Type:             schema.TypeString,
						Optional:         true,
						DiffSuppressFunc: SuppressDiffAdditionProps,
					},
					"class_id": {
						Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"moid": {
						Description: "The Moid of the referenced REST resource.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"object_type": {
						Description: "The fully-qualified name of the remote type referred by this relationship.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"selector": {
						Description: "An OData $filter expression which describes the REST resource to be referenced. This field may\nbe set instead of 'moid' by clients.\n1. If 'moid' is set this field is ignored.\n1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of the\nresource matching the filter expression and populates it in the MoRef that is part of the object\ninstance being inserted/updated to fulfill the REST request.\nAn error is returned if the filter matches zero or more than one REST resource.\nAn example filter string is: Serial eq '3AA8B7T11'.",
						Type:        schema.TypeString,
						Optional:    true,
					},
				},
			},
		},
		"pid": {
			Description: "Product Identifier for a Switch/Fabric-Interconnect.\n* `UCS-FI-6454` - The standard 4th generation UCS Fabric Interconnect with 54 ports.\n* `UCS-FI-64108` - The expanded 4th generation UCS Fabric Interconnect with 108 ports.\n* `unknown` - Unknown device type, usage is TBD.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"ports_supporting100g_speed": {
			Type:     schema.TypeList,
			Optional: true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"additional_properties": {
						Type:             schema.TypeString,
						Optional:         true,
						DiffSuppressFunc: SuppressDiffAdditionProps,
					},
					"class_id": {
						Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"end_port_id": {
						Description: "Ending Port ID in this range of ports.",
						Type:        schema.TypeInt,
						Optional:    true,
					},
					"end_slot_id": {
						Description: "Ending Slot ID in this range of ports.",
						Type:        schema.TypeInt,
						Optional:    true,
					},
					"object_type": {
						Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"start_port_id": {
						Description: "Starting Port ID in this range of ports.",
						Type:        schema.TypeInt,
						Optional:    true,
					},
					"start_slot_id": {
						Description: "Starting Slot ID in this range of ports.",
						Type:        schema.TypeInt,
						Optional:    true,
					},
				},
			},
		},
		"ports_supporting10g_speed": {
			Type:     schema.TypeList,
			Optional: true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"additional_properties": {
						Type:             schema.TypeString,
						Optional:         true,
						DiffSuppressFunc: SuppressDiffAdditionProps,
					},
					"class_id": {
						Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"end_port_id": {
						Description: "Ending Port ID in this range of ports.",
						Type:        schema.TypeInt,
						Optional:    true,
					},
					"end_slot_id": {
						Description: "Ending Slot ID in this range of ports.",
						Type:        schema.TypeInt,
						Optional:    true,
					},
					"object_type": {
						Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"start_port_id": {
						Description: "Starting Port ID in this range of ports.",
						Type:        schema.TypeInt,
						Optional:    true,
					},
					"start_slot_id": {
						Description: "Starting Slot ID in this range of ports.",
						Type:        schema.TypeInt,
						Optional:    true,
					},
				},
			},
		},
		"ports_supporting1g_speed": {
			Type:     schema.TypeList,
			Optional: true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"additional_properties": {
						Type:             schema.TypeString,
						Optional:         true,
						DiffSuppressFunc: SuppressDiffAdditionProps,
					},
					"class_id": {
						Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"end_port_id": {
						Description: "Ending Port ID in this range of ports.",
						Type:        schema.TypeInt,
						Optional:    true,
					},
					"end_slot_id": {
						Description: "Ending Slot ID in this range of ports.",
						Type:        schema.TypeInt,
						Optional:    true,
					},
					"object_type": {
						Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"start_port_id": {
						Description: "Starting Port ID in this range of ports.",
						Type:        schema.TypeInt,
						Optional:    true,
					},
					"start_slot_id": {
						Description: "Starting Slot ID in this range of ports.",
						Type:        schema.TypeInt,
						Optional:    true,
					},
				},
			},
		},
		"ports_supporting25g_speed": {
			Type:     schema.TypeList,
			Optional: true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"additional_properties": {
						Type:             schema.TypeString,
						Optional:         true,
						DiffSuppressFunc: SuppressDiffAdditionProps,
					},
					"class_id": {
						Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"end_port_id": {
						Description: "Ending Port ID in this range of ports.",
						Type:        schema.TypeInt,
						Optional:    true,
					},
					"end_slot_id": {
						Description: "Ending Slot ID in this range of ports.",
						Type:        schema.TypeInt,
						Optional:    true,
					},
					"object_type": {
						Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"start_port_id": {
						Description: "Starting Port ID in this range of ports.",
						Type:        schema.TypeInt,
						Optional:    true,
					},
					"start_slot_id": {
						Description: "Starting Slot ID in this range of ports.",
						Type:        schema.TypeInt,
						Optional:    true,
					},
				},
			},
		},
		"ports_supporting40g_speed": {
			Type:     schema.TypeList,
			Optional: true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"additional_properties": {
						Type:             schema.TypeString,
						Optional:         true,
						DiffSuppressFunc: SuppressDiffAdditionProps,
					},
					"class_id": {
						Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"end_port_id": {
						Description: "Ending Port ID in this range of ports.",
						Type:        schema.TypeInt,
						Optional:    true,
					},
					"end_slot_id": {
						Description: "Ending Slot ID in this range of ports.",
						Type:        schema.TypeInt,
						Optional:    true,
					},
					"object_type": {
						Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"start_port_id": {
						Description: "Starting Port ID in this range of ports.",
						Type:        schema.TypeInt,
						Optional:    true,
					},
					"start_slot_id": {
						Description: "Starting Slot ID in this range of ports.",
						Type:        schema.TypeInt,
						Optional:    true,
					},
				},
			},
		},
		"ports_supporting_breakout": {
			Type:     schema.TypeList,
			Optional: true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"additional_properties": {
						Type:             schema.TypeString,
						Optional:         true,
						DiffSuppressFunc: SuppressDiffAdditionProps,
					},
					"class_id": {
						Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"end_port_id": {
						Description: "Ending Port ID in this range of ports.",
						Type:        schema.TypeInt,
						Optional:    true,
					},
					"end_slot_id": {
						Description: "Ending Slot ID in this range of ports.",
						Type:        schema.TypeInt,
						Optional:    true,
					},
					"object_type": {
						Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"start_port_id": {
						Description: "Starting Port ID in this range of ports.",
						Type:        schema.TypeInt,
						Optional:    true,
					},
					"start_slot_id": {
						Description: "Starting Slot ID in this range of ports.",
						Type:        schema.TypeInt,
						Optional:    true,
					},
				},
			},
		},
		"ports_supporting_fcoe": {
			Type:     schema.TypeList,
			Optional: true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"additional_properties": {
						Type:             schema.TypeString,
						Optional:         true,
						DiffSuppressFunc: SuppressDiffAdditionProps,
					},
					"class_id": {
						Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"end_port_id": {
						Description: "Ending Port ID in this range of ports.",
						Type:        schema.TypeInt,
						Optional:    true,
					},
					"end_slot_id": {
						Description: "Ending Slot ID in this range of ports.",
						Type:        schema.TypeInt,
						Optional:    true,
					},
					"object_type": {
						Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"start_port_id": {
						Description: "Starting Port ID in this range of ports.",
						Type:        schema.TypeInt,
						Optional:    true,
					},
					"start_slot_id": {
						Description: "Starting Slot ID in this range of ports.",
						Type:        schema.TypeInt,
						Optional:    true,
					},
				},
			},
		},
		"ports_supporting_server_role": {
			Type:     schema.TypeList,
			Optional: true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"additional_properties": {
						Type:             schema.TypeString,
						Optional:         true,
						DiffSuppressFunc: SuppressDiffAdditionProps,
					},
					"class_id": {
						Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"end_port_id": {
						Description: "Ending Port ID in this range of ports.",
						Type:        schema.TypeInt,
						Optional:    true,
					},
					"end_slot_id": {
						Description: "Ending Slot ID in this range of ports.",
						Type:        schema.TypeInt,
						Optional:    true,
					},
					"object_type": {
						Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"start_port_id": {
						Description: "Starting Port ID in this range of ports.",
						Type:        schema.TypeInt,
						Optional:    true,
					},
					"start_slot_id": {
						Description: "Starting Slot ID in this range of ports.",
						Type:        schema.TypeInt,
						Optional:    true,
					},
				},
			},
		},
		"reserved_vsans": {
			Type:     schema.TypeList,
			Optional: true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"additional_properties": {
						Type:             schema.TypeString,
						Optional:         true,
						DiffSuppressFunc: SuppressDiffAdditionProps,
					},
					"class_id": {
						Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"end_port_id": {
						Description: "Ending Port ID in this range of ports.",
						Type:        schema.TypeInt,
						Optional:    true,
					},
					"end_slot_id": {
						Description: "Ending Slot ID in this range of ports.",
						Type:        schema.TypeInt,
						Optional:    true,
					},
					"object_type": {
						Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"start_port_id": {
						Description: "Starting Port ID in this range of ports.",
						Type:        schema.TypeInt,
						Optional:    true,
					},
					"start_slot_id": {
						Description: "Starting Slot ID in this range of ports.",
						Type:        schema.TypeInt,
						Optional:    true,
					},
				},
			},
		},
		"sereno_netflow_supported": {
			Description: "Sereno Adaptor with Netflow support on this switch.",
			Type:        schema.TypeBool,
			Optional:    true,
		},
		"shared_scope": {
			Description: "Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.\nObjects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"sku": {
			Description: "SKU information for Switch/Fabric-Interconnect.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"storage_limits": {
			Description: "List of storage limitations for this switch.",
			Type:        schema.TypeList,
			MaxItems:    1,
			Optional:    true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"additional_properties": {
						Type:             schema.TypeString,
						Optional:         true,
						DiffSuppressFunc: SuppressDiffAdditionProps,
					},
					"class_id": {
						Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"maximum_user_zone_count": {
						Description: "Maximum user zones per Switch/Fabric-Interconnect.",
						Type:        schema.TypeInt,
						Optional:    true,
					},
					"maximum_virtual_fc_interfaces": {
						Description: "Maximum configurable Virtual Fibre Channel interfaces on Switch/Fabric-Interconnect.",
						Type:        schema.TypeInt,
						Optional:    true,
					},
					"maximum_virtual_fc_interfaces_per_blade_server": {
						Description: "Maximum configurable Virtual Fibre Channel interfaces per blade.",
						Type:        schema.TypeInt,
						Optional:    true,
					},
					"maximum_vsans": {
						Description: "Maximum configurable VSANs on Switch/Fabric-Interconnect.",
						Type:        schema.TypeInt,
						Optional:    true,
					},
					"maximum_zone_count": {
						Description: "Zone limit per Switch/Fabric-Interconnect.",
						Type:        schema.TypeInt,
						Optional:    true,
					},
					"object_type": {
						Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
						Type:        schema.TypeString,
						Optional:    true,
					},
				},
			},
		},
		"switching_mode_capabilities": {
			Type:     schema.TypeList,
			Optional: true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"additional_properties": {
						Type:             schema.TypeString,
						Optional:         true,
						DiffSuppressFunc: SuppressDiffAdditionProps,
					},
					"class_id": {
						Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"object_type": {
						Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"switching_mode": {
						Description: "Switching mode type (endhost, switch) of the switch.\n* `end-host` - In end-host mode, the fabric interconnects appear to the upstream devices as end hosts with multiple links.In this mode, the switch does not run Spanning Tree Protocol and avoids loops by following a set of rules for traffic forwarding.In case of ethernet switching mode - Ethernet end-host mode is also known as Ethernet host virtualizer.\n* `switch` - In switch mode, the switch runs Spanning Tree Protocol to avoid loops, and broadcast and multicast packets are handled in the traditional way.This is the traditional switch mode.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"vp_compression_supported": {
						Description: "VP Compression support on this switch.",
						Type:        schema.TypeBool,
						Optional:    true,
					},
				},
			},
		},
		"system_limits": {
			Description: "List of system limitations for this switch.",
			Type:        schema.TypeList,
			MaxItems:    1,
			Optional:    true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"additional_properties": {
						Type:             schema.TypeString,
						Optional:         true,
						DiffSuppressFunc: SuppressDiffAdditionProps,
					},
					"class_id": {
						Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"maximum_chassis_count": {
						Description: "Maximum UCS chassis that can be connected to this Switch/Fabric-Interconnect.",
						Type:        schema.TypeInt,
						Optional:    true,
					},
					"maximum_fex_per_domain": {
						Description: "Maximum UCS Fabric-extenders (FEX) per Switch/Fabric-Interconnect.",
						Type:        schema.TypeInt,
						Optional:    true,
					},
					"maximum_servers_per_domain": {
						Description: "Maximum UCS servers per Switch/Fabric-Interconnect.",
						Type:        schema.TypeInt,
						Optional:    true,
					},
					"object_type": {
						Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
						Type:        schema.TypeString,
						Optional:    true,
					},
				},
			},
		},
		"tags": {
			Type:     schema.TypeList,
			Optional: true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"additional_properties": {
						Type:             schema.TypeString,
						Optional:         true,
						DiffSuppressFunc: SuppressDiffAdditionProps,
					},
					"key": {
						Description: "The string representation of a tag key.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"value": {
						Description: "The string representation of a tag value.",
						Type:        schema.TypeString,
						Optional:    true,
					},
				},
			},
		},
		"unified_ports": {
			Type:     schema.TypeList,
			Optional: true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"additional_properties": {
						Type:             schema.TypeString,
						Optional:         true,
						DiffSuppressFunc: SuppressDiffAdditionProps,
					},
					"class_id": {
						Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"end_port_id": {
						Description: "Ending Port ID in this range of ports.",
						Type:        schema.TypeInt,
						Optional:    true,
					},
					"end_slot_id": {
						Description: "Ending Slot ID in this range of ports.",
						Type:        schema.TypeInt,
						Optional:    true,
					},
					"object_type": {
						Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"start_port_id": {
						Description: "Starting Port ID in this range of ports.",
						Type:        schema.TypeInt,
						Optional:    true,
					},
					"start_slot_id": {
						Description: "Starting Slot ID in this range of ports.",
						Type:        schema.TypeInt,
						Optional:    true,
					},
				},
			},
		},
		"unified_rule": {
			Description: "The Slider rule for Unified ports on this switch.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"version_context": {
			Description: "The versioning info for this managed object.",
			Type:        schema.TypeList,
			MaxItems:    1,
			Optional:    true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"additional_properties": {
						Type:             schema.TypeString,
						Optional:         true,
						DiffSuppressFunc: SuppressDiffAdditionProps,
					},
					"class_id": {
						Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"interested_mos": {
						Type:     schema.TypeList,
						Optional: true,
						Elem: &schema.Resource{
							Schema: map[string]*schema.Schema{
								"additional_properties": {
									Type:             schema.TypeString,
									Optional:         true,
									DiffSuppressFunc: SuppressDiffAdditionProps,
								},
								"class_id": {
									Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
									Type:        schema.TypeString,
									Optional:    true,
								},
								"moid": {
									Description: "The Moid of the referenced REST resource.",
									Type:        schema.TypeString,
									Optional:    true,
								},
								"object_type": {
									Description: "The fully-qualified name of the remote type referred by this relationship.",
									Type:        schema.TypeString,
									Optional:    true,
								},
								"selector": {
									Description: "An OData $filter expression which describes the REST resource to be referenced. This field may\nbe set instead of 'moid' by clients.\n1. If 'moid' is set this field is ignored.\n1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of the\nresource matching the filter expression and populates it in the MoRef that is part of the object\ninstance being inserted/updated to fulfill the REST request.\nAn error is returned if the filter matches zero or more than one REST resource.\nAn example filter string is: Serial eq '3AA8B7T11'.",
									Type:        schema.TypeString,
									Optional:    true,
								},
							},
						},
					},
					"object_type": {
						Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"ref_mo": {
						Description: "A reference to the original Managed Object.",
						Type:        schema.TypeList,
						MaxItems:    1,
						Optional:    true,
						Elem: &schema.Resource{
							Schema: map[string]*schema.Schema{
								"additional_properties": {
									Type:             schema.TypeString,
									Optional:         true,
									DiffSuppressFunc: SuppressDiffAdditionProps,
								},
								"class_id": {
									Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
									Type:        schema.TypeString,
									Optional:    true,
								},
								"moid": {
									Description: "The Moid of the referenced REST resource.",
									Type:        schema.TypeString,
									Optional:    true,
								},
								"object_type": {
									Description: "The fully-qualified name of the remote type referred by this relationship.",
									Type:        schema.TypeString,
									Optional:    true,
								},
								"selector": {
									Description: "An OData $filter expression which describes the REST resource to be referenced. This field may\nbe set instead of 'moid' by clients.\n1. If 'moid' is set this field is ignored.\n1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of the\nresource matching the filter expression and populates it in the MoRef that is part of the object\ninstance being inserted/updated to fulfill the REST request.\nAn error is returned if the filter matches zero or more than one REST resource.\nAn example filter string is: Serial eq '3AA8B7T11'.",
									Type:        schema.TypeString,
									Optional:    true,
								},
							},
						},
					},
					"timestamp": {
						Description: "The time this versioned Managed Object was created.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"nr_version": {
						Description: "The version of the Managed Object, e.g. an incrementing number or a hash id.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"version_type": {
						Description: "Specifies type of version. Currently the only supported value is \"Configured\"\nthat is used to keep track of snapshots of policies and profiles that are intended\nto be configured to target endpoints.\n* `Modified` - Version created every time an object is modified.\n* `Configured` - Version created every time an object is configured to the service profile.\n* `Deployed` - Version created for objects related to a service profile when it is deployed.",
						Type:        schema.TypeString,
						Optional:    true,
					},
				},
			},
		},
		"vid": {
			Description: "VID information for Switch/Fabric-Interconnect.",
			Type:        schema.TypeString,
			Optional:    true,
		},
	}
	model["results"] = &schema.Schema{
		Type:     schema.TypeList,
		Elem:     &schema.Resource{Schema: subSchema},
		Computed: true,
	}
	return &schema.Resource{
		ReadContext: dataSourceCapabilitySwitchCapabilityRead,
		Schema:      model}
}

func dataSourceCapabilitySwitchCapabilityRead(c context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)
	var de diag.Diagnostics
	var o = &models.CapabilitySwitchCapability{}
	if _, ok := d.GetOk("account_moid"); ok {
		v := d.Get("account_moid")
		x := (v.(string))
		o.SetAccountMoid(x)
	}

	if _, ok := d.GetOk("additional_properties"); ok {
		v := d.Get("additional_properties")
		x := []byte(v.(string))
		var x1 interface{}
		err := json.Unmarshal(x, &x1)
		if err == nil && x1 != nil {
			o.AdditionalProperties = x1.(map[string]interface{})
		}
	}

	if _, ok := d.GetOk("ancestors"); ok {
		v := d.Get("ancestors")
		x := make([]models.MoBaseMoRelationship, 0)
		s := v.([]interface{})
		for i := 0; i < len(s); i++ {
			o := &models.MoMoRef{}
			l := s[i].(map[string]interface{})
			if v, ok := l["additional_properties"]; ok {
				{
					x := []byte(v.(string))
					var x1 interface{}
					err := json.Unmarshal(x, &x1)
					if err == nil && x1 != nil {
						o.AdditionalProperties = x1.(map[string]interface{})
					}
				}
			}
			o.SetClassId("mo.MoRef")
			if v, ok := l["moid"]; ok {
				{
					x := (v.(string))
					o.SetMoid(x)
				}
			}
			if v, ok := l["object_type"]; ok {
				{
					x := (v.(string))
					o.SetObjectType(x)
				}
			}
			if v, ok := l["selector"]; ok {
				{
					x := (v.(string))
					o.SetSelector(x)
				}
			}
			x = append(x, models.MoMoRefAsMoBaseMoRelationship(o))
		}
		o.SetAncestors(x)
	}

	if _, ok := d.GetOk("class_id"); ok {
		v := d.Get("class_id")
		x := (v.(string))
		o.SetClassId(x)
	}

	if _, ok := d.GetOk("create_time"); ok {
		v := d.Get("create_time")
		x, _ := time.Parse(v.(string), time.RFC1123)
		o.SetCreateTime(x)
	}

	if _, ok := d.GetOk("default_fcoe_vlan"); ok {
		v := d.Get("default_fcoe_vlan")
		x := int64(v.(int))
		o.SetDefaultFcoeVlan(x)
	}

	if _, ok := d.GetOk("domain_group_moid"); ok {
		v := d.Get("domain_group_moid")
		x := (v.(string))
		o.SetDomainGroupMoid(x)
	}

	o.SetDynamicVifsSupported(d.Get("dynamic_vifs_supported").(bool))

	o.SetFanModulesSupported(d.Get("fan_modules_supported").(bool))

	if _, ok := d.GetOk("fc_end_host_mode_reserved_vsans"); ok {
		v := d.Get("fc_end_host_mode_reserved_vsans")
		x := make([]models.CapabilityPortRange, 0)
		s := v.([]interface{})
		for i := 0; i < len(s); i++ {
			o := &models.CapabilityPortRange{}
			l := s[i].(map[string]interface{})
			if v, ok := l["additional_properties"]; ok {
				{
					x := []byte(v.(string))
					var x1 interface{}
					err := json.Unmarshal(x, &x1)
					if err == nil && x1 != nil {
						o.AdditionalProperties = x1.(map[string]interface{})
					}
				}
			}
			o.SetClassId("capability.PortRange")
			if v, ok := l["end_port_id"]; ok {
				{
					x := int64(v.(int))
					o.SetEndPortId(x)
				}
			}
			if v, ok := l["end_slot_id"]; ok {
				{
					x := int64(v.(int))
					o.SetEndSlotId(x)
				}
			}
			if v, ok := l["object_type"]; ok {
				{
					x := (v.(string))
					o.SetObjectType(x)
				}
			}
			if v, ok := l["start_port_id"]; ok {
				{
					x := int64(v.(int))
					o.SetStartPortId(x)
				}
			}
			if v, ok := l["start_slot_id"]; ok {
				{
					x := int64(v.(int))
					o.SetStartSlotId(x)
				}
			}
			x = append(x, *o)
		}
		o.SetFcEndHostModeReservedVsans(x)
	}

	o.SetFcUplinkPortsAutoNegotiationSupported(d.Get("fc_uplink_ports_auto_negotiation_supported").(bool))

	o.SetLocatorBeaconSupported(d.Get("locator_beacon_supported").(bool))

	if _, ok := d.GetOk("max_ports"); ok {
		v := d.Get("max_ports")
		x := int64(v.(int))
		o.SetMaxPorts(x)
	}

	if _, ok := d.GetOk("max_slots"); ok {
		v := d.Get("max_slots")
		x := int64(v.(int))
		o.SetMaxSlots(x)
	}

	if _, ok := d.GetOk("mod_time"); ok {
		v := d.Get("mod_time")
		x, _ := time.Parse(v.(string), time.RFC1123)
		o.SetModTime(x)
	}

	if _, ok := d.GetOk("moid"); ok {
		v := d.Get("moid")
		x := (v.(string))
		o.SetMoid(x)
	}

	if _, ok := d.GetOk("name"); ok {
		v := d.Get("name")
		x := (v.(string))
		o.SetName(x)
	}

	if _, ok := d.GetOk("network_limits"); ok {
		v := d.Get("network_limits")
		p := make([]models.CapabilitySwitchNetworkLimits, 0, 1)
		s := v.([]interface{})
		for i := 0; i < len(s); i++ {
			l := s[i].(map[string]interface{})
			o := &models.CapabilitySwitchNetworkLimits{}
			if v, ok := l["additional_properties"]; ok {
				{
					x := []byte(v.(string))
					var x1 interface{}
					err := json.Unmarshal(x, &x1)
					if err == nil && x1 != nil {
						o.AdditionalProperties = x1.(map[string]interface{})
					}
				}
			}
			o.SetClassId("capability.SwitchNetworkLimits")
			if v, ok := l["max_compressed_port_vlan_count"]; ok {
				{
					x := int64(v.(int))
					o.SetMaxCompressedPortVlanCount(x)
				}
			}
			if v, ok := l["max_uncompressed_port_vlan_count"]; ok {
				{
					x := int64(v.(int))
					o.SetMaxUncompressedPortVlanCount(x)
				}
			}
			if v, ok := l["maximum_active_traffic_monitoring_sessions"]; ok {
				{
					x := int64(v.(int))
					o.SetMaximumActiveTrafficMonitoringSessions(x)
				}
			}
			if v, ok := l["maximum_ethernet_port_channels"]; ok {
				{
					x := int64(v.(int))
					o.SetMaximumEthernetPortChannels(x)
				}
			}
			if v, ok := l["maximum_ethernet_uplink_ports"]; ok {
				{
					x := int64(v.(int))
					o.SetMaximumEthernetUplinkPorts(x)
				}
			}
			if v, ok := l["maximum_fc_port_channel_members"]; ok {
				{
					x := int64(v.(int))
					o.SetMaximumFcPortChannelMembers(x)
				}
			}
			if v, ok := l["maximum_fc_port_channels"]; ok {
				{
					x := int64(v.(int))
					o.SetMaximumFcPortChannels(x)
				}
			}
			if v, ok := l["maximum_igmp_groups"]; ok {
				{
					x := int64(v.(int))
					o.SetMaximumIgmpGroups(x)
				}
			}
			if v, ok := l["maximum_port_channel_members"]; ok {
				{
					x := int64(v.(int))
					o.SetMaximumPortChannelMembers(x)
				}
			}
			if v, ok := l["maximum_primary_vlan"]; ok {
				{
					x := int64(v.(int))
					o.SetMaximumPrimaryVlan(x)
				}
			}
			if v, ok := l["maximum_secondary_vlan"]; ok {
				{
					x := int64(v.(int))
					o.SetMaximumSecondaryVlan(x)
				}
			}
			if v, ok := l["maximum_secondary_vlan_per_primary"]; ok {
				{
					x := int64(v.(int))
					o.SetMaximumSecondaryVlanPerPrimary(x)
				}
			}
			if v, ok := l["maximum_vifs"]; ok {
				{
					x := int64(v.(int))
					o.SetMaximumVifs(x)
				}
			}
			if v, ok := l["maximum_vlans"]; ok {
				{
					x := int64(v.(int))
					o.SetMaximumVlans(x)
				}
			}
			if v, ok := l["minimum_active_fans"]; ok {
				{
					x := int64(v.(int))
					o.SetMinimumActiveFans(x)
				}
			}
			if v, ok := l["object_type"]; ok {
				{
					x := (v.(string))
					o.SetObjectType(x)
				}
			}
			p = append(p, *o)
		}
		if len(p) > 0 {
			x := p[0]
			o.SetNetworkLimits(x)
		}
	}

	if _, ok := d.GetOk("object_type"); ok {
		v := d.Get("object_type")
		x := (v.(string))
		o.SetObjectType(x)
	}

	if _, ok := d.GetOk("owners"); ok {
		v := d.Get("owners")
		x := make([]string, 0)
		y := reflect.ValueOf(v)
		for i := 0; i < y.Len(); i++ {
			x = append(x, y.Index(i).Interface().(string))
		}
		o.SetOwners(x)
	}

	if _, ok := d.GetOk("parent"); ok {
		v := d.Get("parent")
		p := make([]models.MoBaseMoRelationship, 0, 1)
		s := v.([]interface{})
		for i := 0; i < len(s); i++ {
			l := s[i].(map[string]interface{})
			o := &models.MoMoRef{}
			if v, ok := l["additional_properties"]; ok {
				{
					x := []byte(v.(string))
					var x1 interface{}
					err := json.Unmarshal(x, &x1)
					if err == nil && x1 != nil {
						o.AdditionalProperties = x1.(map[string]interface{})
					}
				}
			}
			o.SetClassId("mo.MoRef")
			if v, ok := l["moid"]; ok {
				{
					x := (v.(string))
					o.SetMoid(x)
				}
			}
			if v, ok := l["object_type"]; ok {
				{
					x := (v.(string))
					o.SetObjectType(x)
				}
			}
			if v, ok := l["selector"]; ok {
				{
					x := (v.(string))
					o.SetSelector(x)
				}
			}
			p = append(p, models.MoMoRefAsMoBaseMoRelationship(o))
		}
		if len(p) > 0 {
			x := p[0]
			o.SetParent(x)
		}
	}

	if _, ok := d.GetOk("permission_resources"); ok {
		v := d.Get("permission_resources")
		x := make([]models.MoBaseMoRelationship, 0)
		s := v.([]interface{})
		for i := 0; i < len(s); i++ {
			o := &models.MoMoRef{}
			l := s[i].(map[string]interface{})
			if v, ok := l["additional_properties"]; ok {
				{
					x := []byte(v.(string))
					var x1 interface{}
					err := json.Unmarshal(x, &x1)
					if err == nil && x1 != nil {
						o.AdditionalProperties = x1.(map[string]interface{})
					}
				}
			}
			o.SetClassId("mo.MoRef")
			if v, ok := l["moid"]; ok {
				{
					x := (v.(string))
					o.SetMoid(x)
				}
			}
			if v, ok := l["object_type"]; ok {
				{
					x := (v.(string))
					o.SetObjectType(x)
				}
			}
			if v, ok := l["selector"]; ok {
				{
					x := (v.(string))
					o.SetSelector(x)
				}
			}
			x = append(x, models.MoMoRefAsMoBaseMoRelationship(o))
		}
		o.SetPermissionResources(x)
	}

	if _, ok := d.GetOk("pid"); ok {
		v := d.Get("pid")
		x := (v.(string))
		o.SetPid(x)
	}

	if _, ok := d.GetOk("ports_supporting100g_speed"); ok {
		v := d.Get("ports_supporting100g_speed")
		x := make([]models.CapabilityPortRange, 0)
		s := v.([]interface{})
		for i := 0; i < len(s); i++ {
			o := &models.CapabilityPortRange{}
			l := s[i].(map[string]interface{})
			if v, ok := l["additional_properties"]; ok {
				{
					x := []byte(v.(string))
					var x1 interface{}
					err := json.Unmarshal(x, &x1)
					if err == nil && x1 != nil {
						o.AdditionalProperties = x1.(map[string]interface{})
					}
				}
			}
			o.SetClassId("capability.PortRange")
			if v, ok := l["end_port_id"]; ok {
				{
					x := int64(v.(int))
					o.SetEndPortId(x)
				}
			}
			if v, ok := l["end_slot_id"]; ok {
				{
					x := int64(v.(int))
					o.SetEndSlotId(x)
				}
			}
			if v, ok := l["object_type"]; ok {
				{
					x := (v.(string))
					o.SetObjectType(x)
				}
			}
			if v, ok := l["start_port_id"]; ok {
				{
					x := int64(v.(int))
					o.SetStartPortId(x)
				}
			}
			if v, ok := l["start_slot_id"]; ok {
				{
					x := int64(v.(int))
					o.SetStartSlotId(x)
				}
			}
			x = append(x, *o)
		}
		o.SetPortsSupporting100gSpeed(x)
	}

	if _, ok := d.GetOk("ports_supporting10g_speed"); ok {
		v := d.Get("ports_supporting10g_speed")
		x := make([]models.CapabilityPortRange, 0)
		s := v.([]interface{})
		for i := 0; i < len(s); i++ {
			o := &models.CapabilityPortRange{}
			l := s[i].(map[string]interface{})
			if v, ok := l["additional_properties"]; ok {
				{
					x := []byte(v.(string))
					var x1 interface{}
					err := json.Unmarshal(x, &x1)
					if err == nil && x1 != nil {
						o.AdditionalProperties = x1.(map[string]interface{})
					}
				}
			}
			o.SetClassId("capability.PortRange")
			if v, ok := l["end_port_id"]; ok {
				{
					x := int64(v.(int))
					o.SetEndPortId(x)
				}
			}
			if v, ok := l["end_slot_id"]; ok {
				{
					x := int64(v.(int))
					o.SetEndSlotId(x)
				}
			}
			if v, ok := l["object_type"]; ok {
				{
					x := (v.(string))
					o.SetObjectType(x)
				}
			}
			if v, ok := l["start_port_id"]; ok {
				{
					x := int64(v.(int))
					o.SetStartPortId(x)
				}
			}
			if v, ok := l["start_slot_id"]; ok {
				{
					x := int64(v.(int))
					o.SetStartSlotId(x)
				}
			}
			x = append(x, *o)
		}
		o.SetPortsSupporting10gSpeed(x)
	}

	if _, ok := d.GetOk("ports_supporting1g_speed"); ok {
		v := d.Get("ports_supporting1g_speed")
		x := make([]models.CapabilityPortRange, 0)
		s := v.([]interface{})
		for i := 0; i < len(s); i++ {
			o := &models.CapabilityPortRange{}
			l := s[i].(map[string]interface{})
			if v, ok := l["additional_properties"]; ok {
				{
					x := []byte(v.(string))
					var x1 interface{}
					err := json.Unmarshal(x, &x1)
					if err == nil && x1 != nil {
						o.AdditionalProperties = x1.(map[string]interface{})
					}
				}
			}
			o.SetClassId("capability.PortRange")
			if v, ok := l["end_port_id"]; ok {
				{
					x := int64(v.(int))
					o.SetEndPortId(x)
				}
			}
			if v, ok := l["end_slot_id"]; ok {
				{
					x := int64(v.(int))
					o.SetEndSlotId(x)
				}
			}
			if v, ok := l["object_type"]; ok {
				{
					x := (v.(string))
					o.SetObjectType(x)
				}
			}
			if v, ok := l["start_port_id"]; ok {
				{
					x := int64(v.(int))
					o.SetStartPortId(x)
				}
			}
			if v, ok := l["start_slot_id"]; ok {
				{
					x := int64(v.(int))
					o.SetStartSlotId(x)
				}
			}
			x = append(x, *o)
		}
		o.SetPortsSupporting1gSpeed(x)
	}

	if _, ok := d.GetOk("ports_supporting25g_speed"); ok {
		v := d.Get("ports_supporting25g_speed")
		x := make([]models.CapabilityPortRange, 0)
		s := v.([]interface{})
		for i := 0; i < len(s); i++ {
			o := &models.CapabilityPortRange{}
			l := s[i].(map[string]interface{})
			if v, ok := l["additional_properties"]; ok {
				{
					x := []byte(v.(string))
					var x1 interface{}
					err := json.Unmarshal(x, &x1)
					if err == nil && x1 != nil {
						o.AdditionalProperties = x1.(map[string]interface{})
					}
				}
			}
			o.SetClassId("capability.PortRange")
			if v, ok := l["end_port_id"]; ok {
				{
					x := int64(v.(int))
					o.SetEndPortId(x)
				}
			}
			if v, ok := l["end_slot_id"]; ok {
				{
					x := int64(v.(int))
					o.SetEndSlotId(x)
				}
			}
			if v, ok := l["object_type"]; ok {
				{
					x := (v.(string))
					o.SetObjectType(x)
				}
			}
			if v, ok := l["start_port_id"]; ok {
				{
					x := int64(v.(int))
					o.SetStartPortId(x)
				}
			}
			if v, ok := l["start_slot_id"]; ok {
				{
					x := int64(v.(int))
					o.SetStartSlotId(x)
				}
			}
			x = append(x, *o)
		}
		o.SetPortsSupporting25gSpeed(x)
	}

	if _, ok := d.GetOk("ports_supporting40g_speed"); ok {
		v := d.Get("ports_supporting40g_speed")
		x := make([]models.CapabilityPortRange, 0)
		s := v.([]interface{})
		for i := 0; i < len(s); i++ {
			o := &models.CapabilityPortRange{}
			l := s[i].(map[string]interface{})
			if v, ok := l["additional_properties"]; ok {
				{
					x := []byte(v.(string))
					var x1 interface{}
					err := json.Unmarshal(x, &x1)
					if err == nil && x1 != nil {
						o.AdditionalProperties = x1.(map[string]interface{})
					}
				}
			}
			o.SetClassId("capability.PortRange")
			if v, ok := l["end_port_id"]; ok {
				{
					x := int64(v.(int))
					o.SetEndPortId(x)
				}
			}
			if v, ok := l["end_slot_id"]; ok {
				{
					x := int64(v.(int))
					o.SetEndSlotId(x)
				}
			}
			if v, ok := l["object_type"]; ok {
				{
					x := (v.(string))
					o.SetObjectType(x)
				}
			}
			if v, ok := l["start_port_id"]; ok {
				{
					x := int64(v.(int))
					o.SetStartPortId(x)
				}
			}
			if v, ok := l["start_slot_id"]; ok {
				{
					x := int64(v.(int))
					o.SetStartSlotId(x)
				}
			}
			x = append(x, *o)
		}
		o.SetPortsSupporting40gSpeed(x)
	}

	if _, ok := d.GetOk("ports_supporting_breakout"); ok {
		v := d.Get("ports_supporting_breakout")
		x := make([]models.CapabilityPortRange, 0)
		s := v.([]interface{})
		for i := 0; i < len(s); i++ {
			o := &models.CapabilityPortRange{}
			l := s[i].(map[string]interface{})
			if v, ok := l["additional_properties"]; ok {
				{
					x := []byte(v.(string))
					var x1 interface{}
					err := json.Unmarshal(x, &x1)
					if err == nil && x1 != nil {
						o.AdditionalProperties = x1.(map[string]interface{})
					}
				}
			}
			o.SetClassId("capability.PortRange")
			if v, ok := l["end_port_id"]; ok {
				{
					x := int64(v.(int))
					o.SetEndPortId(x)
				}
			}
			if v, ok := l["end_slot_id"]; ok {
				{
					x := int64(v.(int))
					o.SetEndSlotId(x)
				}
			}
			if v, ok := l["object_type"]; ok {
				{
					x := (v.(string))
					o.SetObjectType(x)
				}
			}
			if v, ok := l["start_port_id"]; ok {
				{
					x := int64(v.(int))
					o.SetStartPortId(x)
				}
			}
			if v, ok := l["start_slot_id"]; ok {
				{
					x := int64(v.(int))
					o.SetStartSlotId(x)
				}
			}
			x = append(x, *o)
		}
		o.SetPortsSupportingBreakout(x)
	}

	if _, ok := d.GetOk("ports_supporting_fcoe"); ok {
		v := d.Get("ports_supporting_fcoe")
		x := make([]models.CapabilityPortRange, 0)
		s := v.([]interface{})
		for i := 0; i < len(s); i++ {
			o := &models.CapabilityPortRange{}
			l := s[i].(map[string]interface{})
			if v, ok := l["additional_properties"]; ok {
				{
					x := []byte(v.(string))
					var x1 interface{}
					err := json.Unmarshal(x, &x1)
					if err == nil && x1 != nil {
						o.AdditionalProperties = x1.(map[string]interface{})
					}
				}
			}
			o.SetClassId("capability.PortRange")
			if v, ok := l["end_port_id"]; ok {
				{
					x := int64(v.(int))
					o.SetEndPortId(x)
				}
			}
			if v, ok := l["end_slot_id"]; ok {
				{
					x := int64(v.(int))
					o.SetEndSlotId(x)
				}
			}
			if v, ok := l["object_type"]; ok {
				{
					x := (v.(string))
					o.SetObjectType(x)
				}
			}
			if v, ok := l["start_port_id"]; ok {
				{
					x := int64(v.(int))
					o.SetStartPortId(x)
				}
			}
			if v, ok := l["start_slot_id"]; ok {
				{
					x := int64(v.(int))
					o.SetStartSlotId(x)
				}
			}
			x = append(x, *o)
		}
		o.SetPortsSupportingFcoe(x)
	}

	if _, ok := d.GetOk("ports_supporting_server_role"); ok {
		v := d.Get("ports_supporting_server_role")
		x := make([]models.CapabilityPortRange, 0)
		s := v.([]interface{})
		for i := 0; i < len(s); i++ {
			o := &models.CapabilityPortRange{}
			l := s[i].(map[string]interface{})
			if v, ok := l["additional_properties"]; ok {
				{
					x := []byte(v.(string))
					var x1 interface{}
					err := json.Unmarshal(x, &x1)
					if err == nil && x1 != nil {
						o.AdditionalProperties = x1.(map[string]interface{})
					}
				}
			}
			o.SetClassId("capability.PortRange")
			if v, ok := l["end_port_id"]; ok {
				{
					x := int64(v.(int))
					o.SetEndPortId(x)
				}
			}
			if v, ok := l["end_slot_id"]; ok {
				{
					x := int64(v.(int))
					o.SetEndSlotId(x)
				}
			}
			if v, ok := l["object_type"]; ok {
				{
					x := (v.(string))
					o.SetObjectType(x)
				}
			}
			if v, ok := l["start_port_id"]; ok {
				{
					x := int64(v.(int))
					o.SetStartPortId(x)
				}
			}
			if v, ok := l["start_slot_id"]; ok {
				{
					x := int64(v.(int))
					o.SetStartSlotId(x)
				}
			}
			x = append(x, *o)
		}
		o.SetPortsSupportingServerRole(x)
	}

	if _, ok := d.GetOk("reserved_vsans"); ok {
		v := d.Get("reserved_vsans")
		x := make([]models.CapabilityPortRange, 0)
		s := v.([]interface{})
		for i := 0; i < len(s); i++ {
			o := &models.CapabilityPortRange{}
			l := s[i].(map[string]interface{})
			if v, ok := l["additional_properties"]; ok {
				{
					x := []byte(v.(string))
					var x1 interface{}
					err := json.Unmarshal(x, &x1)
					if err == nil && x1 != nil {
						o.AdditionalProperties = x1.(map[string]interface{})
					}
				}
			}
			o.SetClassId("capability.PortRange")
			if v, ok := l["end_port_id"]; ok {
				{
					x := int64(v.(int))
					o.SetEndPortId(x)
				}
			}
			if v, ok := l["end_slot_id"]; ok {
				{
					x := int64(v.(int))
					o.SetEndSlotId(x)
				}
			}
			if v, ok := l["object_type"]; ok {
				{
					x := (v.(string))
					o.SetObjectType(x)
				}
			}
			if v, ok := l["start_port_id"]; ok {
				{
					x := int64(v.(int))
					o.SetStartPortId(x)
				}
			}
			if v, ok := l["start_slot_id"]; ok {
				{
					x := int64(v.(int))
					o.SetStartSlotId(x)
				}
			}
			x = append(x, *o)
		}
		o.SetReservedVsans(x)
	}

	o.SetSerenoNetflowSupported(d.Get("sereno_netflow_supported").(bool))

	if _, ok := d.GetOk("shared_scope"); ok {
		v := d.Get("shared_scope")
		x := (v.(string))
		o.SetSharedScope(x)
	}

	if _, ok := d.GetOk("sku"); ok {
		v := d.Get("sku")
		x := (v.(string))
		o.SetSku(x)
	}

	if _, ok := d.GetOk("storage_limits"); ok {
		v := d.Get("storage_limits")
		p := make([]models.CapabilitySwitchStorageLimits, 0, 1)
		s := v.([]interface{})
		for i := 0; i < len(s); i++ {
			l := s[i].(map[string]interface{})
			o := &models.CapabilitySwitchStorageLimits{}
			if v, ok := l["additional_properties"]; ok {
				{
					x := []byte(v.(string))
					var x1 interface{}
					err := json.Unmarshal(x, &x1)
					if err == nil && x1 != nil {
						o.AdditionalProperties = x1.(map[string]interface{})
					}
				}
			}
			o.SetClassId("capability.SwitchStorageLimits")
			if v, ok := l["maximum_user_zone_count"]; ok {
				{
					x := int64(v.(int))
					o.SetMaximumUserZoneCount(x)
				}
			}
			if v, ok := l["maximum_virtual_fc_interfaces"]; ok {
				{
					x := int64(v.(int))
					o.SetMaximumVirtualFcInterfaces(x)
				}
			}
			if v, ok := l["maximum_virtual_fc_interfaces_per_blade_server"]; ok {
				{
					x := int64(v.(int))
					o.SetMaximumVirtualFcInterfacesPerBladeServer(x)
				}
			}
			if v, ok := l["maximum_vsans"]; ok {
				{
					x := int64(v.(int))
					o.SetMaximumVsans(x)
				}
			}
			if v, ok := l["maximum_zone_count"]; ok {
				{
					x := int64(v.(int))
					o.SetMaximumZoneCount(x)
				}
			}
			if v, ok := l["object_type"]; ok {
				{
					x := (v.(string))
					o.SetObjectType(x)
				}
			}
			p = append(p, *o)
		}
		if len(p) > 0 {
			x := p[0]
			o.SetStorageLimits(x)
		}
	}

	if _, ok := d.GetOk("switching_mode_capabilities"); ok {
		v := d.Get("switching_mode_capabilities")
		x := make([]models.CapabilitySwitchingModeCapability, 0)
		s := v.([]interface{})
		for i := 0; i < len(s); i++ {
			o := &models.CapabilitySwitchingModeCapability{}
			l := s[i].(map[string]interface{})
			if v, ok := l["additional_properties"]; ok {
				{
					x := []byte(v.(string))
					var x1 interface{}
					err := json.Unmarshal(x, &x1)
					if err == nil && x1 != nil {
						o.AdditionalProperties = x1.(map[string]interface{})
					}
				}
			}
			o.SetClassId("capability.SwitchingModeCapability")
			if v, ok := l["object_type"]; ok {
				{
					x := (v.(string))
					o.SetObjectType(x)
				}
			}
			if v, ok := l["switching_mode"]; ok {
				{
					x := (v.(string))
					o.SetSwitchingMode(x)
				}
			}
			if v, ok := l["vp_compression_supported"]; ok {
				{
					x := (v.(bool))
					o.SetVpCompressionSupported(x)
				}
			}
			x = append(x, *o)
		}
		o.SetSwitchingModeCapabilities(x)
	}

	if _, ok := d.GetOk("system_limits"); ok {
		v := d.Get("system_limits")
		p := make([]models.CapabilitySwitchSystemLimits, 0, 1)
		s := v.([]interface{})
		for i := 0; i < len(s); i++ {
			l := s[i].(map[string]interface{})
			o := &models.CapabilitySwitchSystemLimits{}
			if v, ok := l["additional_properties"]; ok {
				{
					x := []byte(v.(string))
					var x1 interface{}
					err := json.Unmarshal(x, &x1)
					if err == nil && x1 != nil {
						o.AdditionalProperties = x1.(map[string]interface{})
					}
				}
			}
			o.SetClassId("capability.SwitchSystemLimits")
			if v, ok := l["maximum_chassis_count"]; ok {
				{
					x := int64(v.(int))
					o.SetMaximumChassisCount(x)
				}
			}
			if v, ok := l["maximum_fex_per_domain"]; ok {
				{
					x := int64(v.(int))
					o.SetMaximumFexPerDomain(x)
				}
			}
			if v, ok := l["maximum_servers_per_domain"]; ok {
				{
					x := int64(v.(int))
					o.SetMaximumServersPerDomain(x)
				}
			}
			if v, ok := l["object_type"]; ok {
				{
					x := (v.(string))
					o.SetObjectType(x)
				}
			}
			p = append(p, *o)
		}
		if len(p) > 0 {
			x := p[0]
			o.SetSystemLimits(x)
		}
	}

	if _, ok := d.GetOk("tags"); ok {
		v := d.Get("tags")
		x := make([]models.MoTag, 0)
		s := v.([]interface{})
		for i := 0; i < len(s); i++ {
			o := &models.MoTag{}
			l := s[i].(map[string]interface{})
			if v, ok := l["additional_properties"]; ok {
				{
					x := []byte(v.(string))
					var x1 interface{}
					err := json.Unmarshal(x, &x1)
					if err == nil && x1 != nil {
						o.AdditionalProperties = x1.(map[string]interface{})
					}
				}
			}
			if v, ok := l["key"]; ok {
				{
					x := (v.(string))
					o.SetKey(x)
				}
			}
			if v, ok := l["value"]; ok {
				{
					x := (v.(string))
					o.SetValue(x)
				}
			}
			x = append(x, *o)
		}
		o.SetTags(x)
	}

	if _, ok := d.GetOk("unified_ports"); ok {
		v := d.Get("unified_ports")
		x := make([]models.CapabilityPortRange, 0)
		s := v.([]interface{})
		for i := 0; i < len(s); i++ {
			o := &models.CapabilityPortRange{}
			l := s[i].(map[string]interface{})
			if v, ok := l["additional_properties"]; ok {
				{
					x := []byte(v.(string))
					var x1 interface{}
					err := json.Unmarshal(x, &x1)
					if err == nil && x1 != nil {
						o.AdditionalProperties = x1.(map[string]interface{})
					}
				}
			}
			o.SetClassId("capability.PortRange")
			if v, ok := l["end_port_id"]; ok {
				{
					x := int64(v.(int))
					o.SetEndPortId(x)
				}
			}
			if v, ok := l["end_slot_id"]; ok {
				{
					x := int64(v.(int))
					o.SetEndSlotId(x)
				}
			}
			if v, ok := l["object_type"]; ok {
				{
					x := (v.(string))
					o.SetObjectType(x)
				}
			}
			if v, ok := l["start_port_id"]; ok {
				{
					x := int64(v.(int))
					o.SetStartPortId(x)
				}
			}
			if v, ok := l["start_slot_id"]; ok {
				{
					x := int64(v.(int))
					o.SetStartSlotId(x)
				}
			}
			x = append(x, *o)
		}
		o.SetUnifiedPorts(x)
	}

	if _, ok := d.GetOk("unified_rule"); ok {
		v := d.Get("unified_rule")
		x := (v.(string))
		o.SetUnifiedRule(x)
	}

	if _, ok := d.GetOk("version_context"); ok {
		v := d.Get("version_context")
		p := make([]models.MoVersionContext, 0, 1)
		s := v.([]interface{})
		for i := 0; i < len(s); i++ {
			l := s[i].(map[string]interface{})
			o := &models.MoVersionContext{}
			if v, ok := l["additional_properties"]; ok {
				{
					x := []byte(v.(string))
					var x1 interface{}
					err := json.Unmarshal(x, &x1)
					if err == nil && x1 != nil {
						o.AdditionalProperties = x1.(map[string]interface{})
					}
				}
			}
			o.SetClassId("mo.VersionContext")
			if v, ok := l["interested_mos"]; ok {
				{
					x := make([]models.MoMoRef, 0)
					s := v.([]interface{})
					for i := 0; i < len(s); i++ {
						o := models.NewMoMoRefWithDefaults()
						l := s[i].(map[string]interface{})
						if v, ok := l["additional_properties"]; ok {
							{
								x := []byte(v.(string))
								var x1 interface{}
								err := json.Unmarshal(x, &x1)
								if err == nil && x1 != nil {
									o.AdditionalProperties = x1.(map[string]interface{})
								}
							}
						}
						o.SetClassId("mo.MoRef")
						if v, ok := l["moid"]; ok {
							{
								x := (v.(string))
								o.SetMoid(x)
							}
						}
						if v, ok := l["object_type"]; ok {
							{
								x := (v.(string))
								o.SetObjectType(x)
							}
						}
						if v, ok := l["selector"]; ok {
							{
								x := (v.(string))
								o.SetSelector(x)
							}
						}
						x = append(x, *o)
					}
					if len(x) > 0 {
						o.SetInterestedMos(x)
					}
				}
			}
			if v, ok := l["object_type"]; ok {
				{
					x := (v.(string))
					o.SetObjectType(x)
				}
			}
			p = append(p, *o)
		}
		if len(p) > 0 {
			x := p[0]
			o.SetVersionContext(x)
		}
	}

	if _, ok := d.GetOk("vid"); ok {
		v := d.Get("vid")
		x := (v.(string))
		o.SetVid(x)
	}

	if v, ok := d.GetOk("account_moid"); ok {
		x := (v.(string))
		o.SetAccountMoid(x)
	}
	if v, ok := d.GetOk("class_id"); ok {
		x := (v.(string))
		o.SetClassId(x)
	}
	if v, ok := d.GetOk("create_time"); ok {
		x, _ := time.Parse(v.(string), time.RFC1123)
		o.SetCreateTime(x)
	}
	if v, ok := d.GetOk("default_fcoe_vlan"); ok {
		x := int64(v.(int))
		o.SetDefaultFcoeVlan(x)
	}
	if v, ok := d.GetOk("domain_group_moid"); ok {
		x := (v.(string))
		o.SetDomainGroupMoid(x)
	}
	if v, ok := d.GetOk("dynamic_vifs_supported"); ok {
		x := (v.(bool))
		o.SetDynamicVifsSupported(x)
	}
	if v, ok := d.GetOk("fan_modules_supported"); ok {
		x := (v.(bool))
		o.SetFanModulesSupported(x)
	}
	if v, ok := d.GetOk("fc_uplink_ports_auto_negotiation_supported"); ok {
		x := (v.(bool))
		o.SetFcUplinkPortsAutoNegotiationSupported(x)
	}
	if v, ok := d.GetOk("locator_beacon_supported"); ok {
		x := (v.(bool))
		o.SetLocatorBeaconSupported(x)
	}
	if v, ok := d.GetOk("max_ports"); ok {
		x := int64(v.(int))
		o.SetMaxPorts(x)
	}
	if v, ok := d.GetOk("max_slots"); ok {
		x := int64(v.(int))
		o.SetMaxSlots(x)
	}
	if v, ok := d.GetOk("mod_time"); ok {
		x, _ := time.Parse(v.(string), time.RFC1123)
		o.SetModTime(x)
	}
	if v, ok := d.GetOk("moid"); ok {
		x := (v.(string))
		o.SetMoid(x)
	}
	if v, ok := d.GetOk("name"); ok {
		x := (v.(string))
		o.SetName(x)
	}
	if v, ok := d.GetOk("object_type"); ok {
		x := (v.(string))
		o.SetObjectType(x)
	}
	if v, ok := d.GetOk("pid"); ok {
		x := (v.(string))
		o.SetPid(x)
	}
	if v, ok := d.GetOk("sereno_netflow_supported"); ok {
		x := (v.(bool))
		o.SetSerenoNetflowSupported(x)
	}
	if v, ok := d.GetOk("shared_scope"); ok {
		x := (v.(string))
		o.SetSharedScope(x)
	}
	if v, ok := d.GetOk("sku"); ok {
		x := (v.(string))
		o.SetSku(x)
	}
	if v, ok := d.GetOk("unified_rule"); ok {
		x := (v.(string))
		o.SetUnifiedRule(x)
	}
	if v, ok := d.GetOk("vid"); ok {
		x := (v.(string))
		o.SetVid(x)
	}

	data, err := o.MarshalJSON()
	if err != nil {
		return diag.Errorf("json marshal of CapabilitySwitchCapability object failed with error : %s", err.Error())
	}
	countResponse, _, responseErr := conn.ApiClient.CapabilityApi.GetCapabilitySwitchCapabilityList(conn.ctx).Filter(getRequestParams(data)).Inlinecount("allpages").Execute()
	if responseErr != nil {
		errorType := fmt.Sprintf("%T", responseErr)
		if strings.Contains(errorType, "GenericOpenAPIError") {
			responseErr := responseErr.(models.GenericOpenAPIError)
			return diag.Errorf("error occurred while fetching count of CapabilitySwitchCapability: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
		}
		return diag.Errorf("error occurred while fetching count of CapabilitySwitchCapability: %s", responseErr.Error())
	}
	count := countResponse.CapabilitySwitchCapabilityList.GetCount()
	if count == 0 {
		return diag.Errorf("your query for CapabilitySwitchCapability data source did not return any results. Please change your search criteria and try again")
	}
	var i int32
	var capabilitySwitchCapabilityResults = make([]map[string]interface{}, count, count)
	var j = 0
	for i = 0; i < count; i += 100 {
		resMo, _, responseErr := conn.ApiClient.CapabilityApi.GetCapabilitySwitchCapabilityList(conn.ctx).Filter(getRequestParams(data)).Top(100).Skip(i).Execute()
		if responseErr != nil {
			errorType := fmt.Sprintf("%T", responseErr)
			if strings.Contains(errorType, "GenericOpenAPIError") {
				responseErr := responseErr.(models.GenericOpenAPIError)
				return diag.Errorf("error occurred while fetching CapabilitySwitchCapability: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
			}
			return diag.Errorf("error occurred while fetching CapabilitySwitchCapability: %s", responseErr.Error())
		}
		results := resMo.CapabilitySwitchCapabilityList.GetResults()
		switch reflect.TypeOf(results).Kind() {
		case reflect.Slice:
			for i := 0; i < len(results); i++ {
				var s = results[i]
				var temp = make(map[string]interface{})
				temp["account_moid"] = (s.GetAccountMoid())
				temp["additional_properties"] = flattenAdditionalProperties(s.AdditionalProperties)

				temp["ancestors"] = flattenListMoBaseMoRelationship(s.GetAncestors(), d)
				temp["class_id"] = (s.GetClassId())

				temp["create_time"] = (s.GetCreateTime()).String()
				temp["default_fcoe_vlan"] = (s.GetDefaultFcoeVlan())
				temp["domain_group_moid"] = (s.GetDomainGroupMoid())
				temp["dynamic_vifs_supported"] = (s.GetDynamicVifsSupported())
				temp["fan_modules_supported"] = (s.GetFanModulesSupported())

				temp["fc_end_host_mode_reserved_vsans"] = flattenListCapabilityPortRange(s.GetFcEndHostModeReservedVsans(), d)
				temp["fc_uplink_ports_auto_negotiation_supported"] = (s.GetFcUplinkPortsAutoNegotiationSupported())
				temp["locator_beacon_supported"] = (s.GetLocatorBeaconSupported())
				temp["max_ports"] = (s.GetMaxPorts())
				temp["max_slots"] = (s.GetMaxSlots())

				temp["mod_time"] = (s.GetModTime()).String()
				temp["moid"] = (s.GetMoid())
				temp["name"] = (s.GetName())

				temp["network_limits"] = flattenMapCapabilitySwitchNetworkLimits(s.GetNetworkLimits(), d)
				temp["object_type"] = (s.GetObjectType())
				temp["owners"] = (s.GetOwners())

				temp["parent"] = flattenMapMoBaseMoRelationship(s.GetParent(), d)

				temp["permission_resources"] = flattenListMoBaseMoRelationship(s.GetPermissionResources(), d)
				temp["pid"] = (s.GetPid())

				temp["ports_supporting100g_speed"] = flattenListCapabilityPortRange(s.GetPortsSupporting100gSpeed(), d)

				temp["ports_supporting10g_speed"] = flattenListCapabilityPortRange(s.GetPortsSupporting10gSpeed(), d)

				temp["ports_supporting1g_speed"] = flattenListCapabilityPortRange(s.GetPortsSupporting1gSpeed(), d)

				temp["ports_supporting25g_speed"] = flattenListCapabilityPortRange(s.GetPortsSupporting25gSpeed(), d)

				temp["ports_supporting40g_speed"] = flattenListCapabilityPortRange(s.GetPortsSupporting40gSpeed(), d)

				temp["ports_supporting_breakout"] = flattenListCapabilityPortRange(s.GetPortsSupportingBreakout(), d)

				temp["ports_supporting_fcoe"] = flattenListCapabilityPortRange(s.GetPortsSupportingFcoe(), d)

				temp["ports_supporting_server_role"] = flattenListCapabilityPortRange(s.GetPortsSupportingServerRole(), d)

				temp["reserved_vsans"] = flattenListCapabilityPortRange(s.GetReservedVsans(), d)
				temp["sereno_netflow_supported"] = (s.GetSerenoNetflowSupported())
				temp["shared_scope"] = (s.GetSharedScope())
				temp["sku"] = (s.GetSku())

				temp["storage_limits"] = flattenMapCapabilitySwitchStorageLimits(s.GetStorageLimits(), d)

				temp["switching_mode_capabilities"] = flattenListCapabilitySwitchingModeCapability(s.GetSwitchingModeCapabilities(), d)

				temp["system_limits"] = flattenMapCapabilitySwitchSystemLimits(s.GetSystemLimits(), d)

				temp["tags"] = flattenListMoTag(s.GetTags(), d)

				temp["unified_ports"] = flattenListCapabilityPortRange(s.GetUnifiedPorts(), d)
				temp["unified_rule"] = (s.GetUnifiedRule())

				temp["version_context"] = flattenMapMoVersionContext(s.GetVersionContext(), d)
				temp["vid"] = (s.GetVid())
				capabilitySwitchCapabilityResults[j] = temp
				j += 1
			}
		}
	}
	log.Println("length of results: ", len(capabilitySwitchCapabilityResults))
	if err := d.Set("results", capabilitySwitchCapabilityResults); err != nil {
		return diag.Errorf("error occurred while setting results: %s", err.Error())
	}
	d.SetId(capabilitySwitchCapabilityResults[0]["moid"].(string))
	return de
}
