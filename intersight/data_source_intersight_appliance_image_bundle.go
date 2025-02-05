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

func dataSourceApplianceImageBundle() *schema.Resource {
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
		"ansible_packages": {
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
					"file_path": {
						Description: "Optional file path of the image package.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"file_sha": {
						Description: "Image file's fingerprint. Fingerprint is calculated using SHA256 algorithm.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"file_size": {
						Description: "Image file size in bytes.",
						Type:        schema.TypeInt,
						Optional:    true,
					},
					"file_time": {
						Description: "Image file's last modified date and time.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"filename": {
						Description: "Filename of the image package.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"name": {
						Description: "Name of the software image package.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"object_type": {
						Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"package_type": {
						Description: "Image package type (e.g. service, system etc.).",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"nr_version": {
						Description: "Image package version string.",
						Type:        schema.TypeString,
						Optional:    true,
					},
				},
			},
		},
		"auto_upgrade": {
			Description: "Indicates that the software upgrade was automatically initiated by the Intersight Appliance.",
			Type:        schema.TypeBool,
			Optional:    true,
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
		"dc_packages": {
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
					"file_path": {
						Description: "Optional file path of the image package.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"file_sha": {
						Description: "Image file's fingerprint. Fingerprint is calculated using SHA256 algorithm.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"file_size": {
						Description: "Image file size in bytes.",
						Type:        schema.TypeInt,
						Optional:    true,
					},
					"file_time": {
						Description: "Image file's last modified date and time.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"filename": {
						Description: "Filename of the image package.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"name": {
						Description: "Name of the software image package.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"object_type": {
						Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"package_type": {
						Description: "Image package type (e.g. service, system etc.).",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"nr_version": {
						Description: "Image package version string.",
						Type:        schema.TypeString,
						Optional:    true,
					},
				},
			},
		},
		"debug_packages": {
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
					"file_path": {
						Description: "Optional file path of the image package.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"file_sha": {
						Description: "Image file's fingerprint. Fingerprint is calculated using SHA256 algorithm.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"file_size": {
						Description: "Image file size in bytes.",
						Type:        schema.TypeInt,
						Optional:    true,
					},
					"file_time": {
						Description: "Image file's last modified date and time.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"filename": {
						Description: "Filename of the image package.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"name": {
						Description: "Name of the software image package.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"object_type": {
						Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"package_type": {
						Description: "Image package type (e.g. service, system etc.).",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"nr_version": {
						Description: "Image package version string.",
						Type:        schema.TypeString,
						Optional:    true,
					},
				},
			},
		},
		"description": {
			Description: "Short description of the software upgrade bundle.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"domain_group_moid": {
			Description: "The DomainGroup ID for this managed object.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"endpoint_packages": {
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
					"file_path": {
						Description: "Optional file path of the image package.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"file_sha": {
						Description: "Image file's fingerprint. Fingerprint is calculated using SHA256 algorithm.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"file_size": {
						Description: "Image file size in bytes.",
						Type:        schema.TypeInt,
						Optional:    true,
					},
					"file_time": {
						Description: "Image file's last modified date and time.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"filename": {
						Description: "Filename of the image package.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"name": {
						Description: "Name of the software image package.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"object_type": {
						Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"package_type": {
						Description: "Image package type (e.g. service, system etc.).",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"nr_version": {
						Description: "Image package version string.",
						Type:        schema.TypeString,
						Optional:    true,
					},
				},
			},
		},
		"fingerprint": {
			Description: "Fingerprint of the software manifest from which this bundle is created. Fingerprint is calculated using the SHA256 algorithm.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"has_error": {
			Description: "Indicates that the ImageBundle has errors. The upgrade service sets this field when it encounters errors during the manifest processing.",
			Type:        schema.TypeBool,
			Optional:    true,
		},
		"infra_packages": {
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
					"file_path": {
						Description: "Optional file path of the image package.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"file_sha": {
						Description: "Image file's fingerprint. Fingerprint is calculated using SHA256 algorithm.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"file_size": {
						Description: "Image file size in bytes.",
						Type:        schema.TypeInt,
						Optional:    true,
					},
					"file_time": {
						Description: "Image file's last modified date and time.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"filename": {
						Description: "Filename of the image package.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"name": {
						Description: "Name of the software image package.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"object_type": {
						Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"package_type": {
						Description: "Image package type (e.g. service, system etc.).",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"nr_version": {
						Description: "Image package version string.",
						Type:        schema.TypeString,
						Optional:    true,
					},
				},
			},
		},
		"init_packages": {
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
					"file_path": {
						Description: "Optional file path of the image package.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"file_sha": {
						Description: "Image file's fingerprint. Fingerprint is calculated using SHA256 algorithm.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"file_size": {
						Description: "Image file size in bytes.",
						Type:        schema.TypeInt,
						Optional:    true,
					},
					"file_time": {
						Description: "Image file's last modified date and time.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"filename": {
						Description: "Filename of the image package.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"name": {
						Description: "Name of the software image package.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"object_type": {
						Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"package_type": {
						Description: "Image package type (e.g. service, system etc.).",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"nr_version": {
						Description: "Image package version string.",
						Type:        schema.TypeString,
						Optional:    true,
					},
				},
			},
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
			Description: "Name of the software upgrade bundle.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"notes": {
			Description: "Detailed description of the software upgrade bundle.",
			Type:        schema.TypeString,
			Optional:    true,
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
		"priority": {
			Description: "Software upgrade manifest's upgrade priority. The upgrade service supports two priorities, Normal and Critical. Normal priority is used for regular software upgrades, and the upgrade service uses the Upgrade Policy to compute upgrade start time. Critical priority is used for the critical software security patches, and the upgrade service ignores the Upgrade Policy when it computes the upgrade start time.\n* `Normal` - Normal upgrade priority is used for all the software upgrades except for the critical security updates. The upgrade service of Intersight Appliance uses the Software Upgrade Policy settings to start the upgrade process.\n* `Critical` - Critical upgrade priority is used for critical updates such as security patches. The upgrade service of the Intersight Appliance starts the upgrade as specified by the upgrade properties in the software manifest file. The upgrade service will not use the settings specified in the Software Upgrade Policy.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"release_time": {
			Description: "Software upgrade manifest's release date and time.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"service_packages": {
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
					"file_path": {
						Description: "Optional file path of the image package.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"file_sha": {
						Description: "Image file's fingerprint. Fingerprint is calculated using SHA256 algorithm.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"file_size": {
						Description: "Image file size in bytes.",
						Type:        schema.TypeInt,
						Optional:    true,
					},
					"file_time": {
						Description: "Image file's last modified date and time.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"filename": {
						Description: "Filename of the image package.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"name": {
						Description: "Name of the software image package.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"object_type": {
						Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"package_type": {
						Description: "Image package type (e.g. service, system etc.).",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"nr_version": {
						Description: "Image package version string.",
						Type:        schema.TypeString,
						Optional:    true,
					},
				},
			},
		},
		"shared_scope": {
			Description: "Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.\nObjects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"status_message": {
			Description: "Status message set during the manifest processing.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"system_packages": {
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
					"file_path": {
						Description: "Optional file path of the image package.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"file_sha": {
						Description: "Image file's fingerprint. Fingerprint is calculated using SHA256 algorithm.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"file_size": {
						Description: "Image file size in bytes.",
						Type:        schema.TypeInt,
						Optional:    true,
					},
					"file_time": {
						Description: "Image file's last modified date and time.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"filename": {
						Description: "Filename of the image package.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"name": {
						Description: "Name of the software image package.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"object_type": {
						Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"package_type": {
						Description: "Image package type (e.g. service, system etc.).",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"nr_version": {
						Description: "Image package version string.",
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
		"ui_packages": {
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
					"file_path": {
						Description: "Optional file path of the image package.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"file_sha": {
						Description: "Image file's fingerprint. Fingerprint is calculated using SHA256 algorithm.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"file_size": {
						Description: "Image file size in bytes.",
						Type:        schema.TypeInt,
						Optional:    true,
					},
					"file_time": {
						Description: "Image file's last modified date and time.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"filename": {
						Description: "Filename of the image package.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"name": {
						Description: "Name of the software image package.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"object_type": {
						Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"package_type": {
						Description: "Image package type (e.g. service, system etc.).",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"nr_version": {
						Description: "Image package version string.",
						Type:        schema.TypeString,
						Optional:    true,
					},
				},
			},
		},
		"upgrade_end_time": {
			Description: "End date of the software upgrade process.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"upgrade_grace_period": {
			Description: "Grace period in seconds before the automatic upgrade is initiated. The upgrade service uses the grace period to compute the upgrade start time when it receives an upgrade notfication from the Intersight. If there is an Upgrade Policy configured for the Intersight Appliance, then the upgrade service uses the policy to compute the upgrade start time. However, the upgrade start time cannot not exceed the limit enforced by the grace period.",
			Type:        schema.TypeInt,
			Optional:    true,
		},
		"upgrade_impact_duration": {
			Description: "Duration (in minutes) for which services will be disrupted.",
			Type:        schema.TypeInt,
			Optional:    true,
		},
		"upgrade_impact_enum": {
			Description: "UpgradeImpactEnum is used to indicate the kind of impact the upgrade has on currently running services on the appliance.\n* `None` - The upgrade has no effect on the system.\n* `Disruptive` - The services will not be functional during the upgrade.\n* `Disruptive-reboot` - The upgrade needs a reboot.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"upgrade_start_time": {
			Description: "Start date of the software upgrade process.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"nr_version": {
			Description: "Software upgrade manifest's version.",
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
		"ansible_packages": {
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
					"file_path": {
						Description: "Optional file path of the image package.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"file_sha": {
						Description: "Image file's fingerprint. Fingerprint is calculated using SHA256 algorithm.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"file_size": {
						Description: "Image file size in bytes.",
						Type:        schema.TypeInt,
						Optional:    true,
					},
					"file_time": {
						Description: "Image file's last modified date and time.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"filename": {
						Description: "Filename of the image package.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"name": {
						Description: "Name of the software image package.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"object_type": {
						Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"package_type": {
						Description: "Image package type (e.g. service, system etc.).",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"nr_version": {
						Description: "Image package version string.",
						Type:        schema.TypeString,
						Optional:    true,
					},
				},
			},
		},
		"auto_upgrade": {
			Description: "Indicates that the software upgrade was automatically initiated by the Intersight Appliance.",
			Type:        schema.TypeBool,
			Optional:    true,
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
		"dc_packages": {
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
					"file_path": {
						Description: "Optional file path of the image package.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"file_sha": {
						Description: "Image file's fingerprint. Fingerprint is calculated using SHA256 algorithm.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"file_size": {
						Description: "Image file size in bytes.",
						Type:        schema.TypeInt,
						Optional:    true,
					},
					"file_time": {
						Description: "Image file's last modified date and time.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"filename": {
						Description: "Filename of the image package.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"name": {
						Description: "Name of the software image package.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"object_type": {
						Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"package_type": {
						Description: "Image package type (e.g. service, system etc.).",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"nr_version": {
						Description: "Image package version string.",
						Type:        schema.TypeString,
						Optional:    true,
					},
				},
			},
		},
		"debug_packages": {
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
					"file_path": {
						Description: "Optional file path of the image package.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"file_sha": {
						Description: "Image file's fingerprint. Fingerprint is calculated using SHA256 algorithm.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"file_size": {
						Description: "Image file size in bytes.",
						Type:        schema.TypeInt,
						Optional:    true,
					},
					"file_time": {
						Description: "Image file's last modified date and time.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"filename": {
						Description: "Filename of the image package.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"name": {
						Description: "Name of the software image package.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"object_type": {
						Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"package_type": {
						Description: "Image package type (e.g. service, system etc.).",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"nr_version": {
						Description: "Image package version string.",
						Type:        schema.TypeString,
						Optional:    true,
					},
				},
			},
		},
		"description": {
			Description: "Short description of the software upgrade bundle.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"domain_group_moid": {
			Description: "The DomainGroup ID for this managed object.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"endpoint_packages": {
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
					"file_path": {
						Description: "Optional file path of the image package.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"file_sha": {
						Description: "Image file's fingerprint. Fingerprint is calculated using SHA256 algorithm.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"file_size": {
						Description: "Image file size in bytes.",
						Type:        schema.TypeInt,
						Optional:    true,
					},
					"file_time": {
						Description: "Image file's last modified date and time.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"filename": {
						Description: "Filename of the image package.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"name": {
						Description: "Name of the software image package.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"object_type": {
						Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"package_type": {
						Description: "Image package type (e.g. service, system etc.).",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"nr_version": {
						Description: "Image package version string.",
						Type:        schema.TypeString,
						Optional:    true,
					},
				},
			},
		},
		"fingerprint": {
			Description: "Fingerprint of the software manifest from which this bundle is created. Fingerprint is calculated using the SHA256 algorithm.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"has_error": {
			Description: "Indicates that the ImageBundle has errors. The upgrade service sets this field when it encounters errors during the manifest processing.",
			Type:        schema.TypeBool,
			Optional:    true,
		},
		"infra_packages": {
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
					"file_path": {
						Description: "Optional file path of the image package.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"file_sha": {
						Description: "Image file's fingerprint. Fingerprint is calculated using SHA256 algorithm.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"file_size": {
						Description: "Image file size in bytes.",
						Type:        schema.TypeInt,
						Optional:    true,
					},
					"file_time": {
						Description: "Image file's last modified date and time.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"filename": {
						Description: "Filename of the image package.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"name": {
						Description: "Name of the software image package.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"object_type": {
						Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"package_type": {
						Description: "Image package type (e.g. service, system etc.).",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"nr_version": {
						Description: "Image package version string.",
						Type:        schema.TypeString,
						Optional:    true,
					},
				},
			},
		},
		"init_packages": {
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
					"file_path": {
						Description: "Optional file path of the image package.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"file_sha": {
						Description: "Image file's fingerprint. Fingerprint is calculated using SHA256 algorithm.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"file_size": {
						Description: "Image file size in bytes.",
						Type:        schema.TypeInt,
						Optional:    true,
					},
					"file_time": {
						Description: "Image file's last modified date and time.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"filename": {
						Description: "Filename of the image package.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"name": {
						Description: "Name of the software image package.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"object_type": {
						Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"package_type": {
						Description: "Image package type (e.g. service, system etc.).",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"nr_version": {
						Description: "Image package version string.",
						Type:        schema.TypeString,
						Optional:    true,
					},
				},
			},
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
			Description: "Name of the software upgrade bundle.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"notes": {
			Description: "Detailed description of the software upgrade bundle.",
			Type:        schema.TypeString,
			Optional:    true,
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
		"priority": {
			Description: "Software upgrade manifest's upgrade priority. The upgrade service supports two priorities, Normal and Critical. Normal priority is used for regular software upgrades, and the upgrade service uses the Upgrade Policy to compute upgrade start time. Critical priority is used for the critical software security patches, and the upgrade service ignores the Upgrade Policy when it computes the upgrade start time.\n* `Normal` - Normal upgrade priority is used for all the software upgrades except for the critical security updates. The upgrade service of Intersight Appliance uses the Software Upgrade Policy settings to start the upgrade process.\n* `Critical` - Critical upgrade priority is used for critical updates such as security patches. The upgrade service of the Intersight Appliance starts the upgrade as specified by the upgrade properties in the software manifest file. The upgrade service will not use the settings specified in the Software Upgrade Policy.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"release_time": {
			Description: "Software upgrade manifest's release date and time.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"service_packages": {
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
					"file_path": {
						Description: "Optional file path of the image package.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"file_sha": {
						Description: "Image file's fingerprint. Fingerprint is calculated using SHA256 algorithm.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"file_size": {
						Description: "Image file size in bytes.",
						Type:        schema.TypeInt,
						Optional:    true,
					},
					"file_time": {
						Description: "Image file's last modified date and time.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"filename": {
						Description: "Filename of the image package.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"name": {
						Description: "Name of the software image package.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"object_type": {
						Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"package_type": {
						Description: "Image package type (e.g. service, system etc.).",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"nr_version": {
						Description: "Image package version string.",
						Type:        schema.TypeString,
						Optional:    true,
					},
				},
			},
		},
		"shared_scope": {
			Description: "Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.\nObjects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"status_message": {
			Description: "Status message set during the manifest processing.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"system_packages": {
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
					"file_path": {
						Description: "Optional file path of the image package.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"file_sha": {
						Description: "Image file's fingerprint. Fingerprint is calculated using SHA256 algorithm.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"file_size": {
						Description: "Image file size in bytes.",
						Type:        schema.TypeInt,
						Optional:    true,
					},
					"file_time": {
						Description: "Image file's last modified date and time.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"filename": {
						Description: "Filename of the image package.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"name": {
						Description: "Name of the software image package.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"object_type": {
						Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"package_type": {
						Description: "Image package type (e.g. service, system etc.).",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"nr_version": {
						Description: "Image package version string.",
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
		"ui_packages": {
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
					"file_path": {
						Description: "Optional file path of the image package.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"file_sha": {
						Description: "Image file's fingerprint. Fingerprint is calculated using SHA256 algorithm.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"file_size": {
						Description: "Image file size in bytes.",
						Type:        schema.TypeInt,
						Optional:    true,
					},
					"file_time": {
						Description: "Image file's last modified date and time.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"filename": {
						Description: "Filename of the image package.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"name": {
						Description: "Name of the software image package.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"object_type": {
						Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"package_type": {
						Description: "Image package type (e.g. service, system etc.).",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"nr_version": {
						Description: "Image package version string.",
						Type:        schema.TypeString,
						Optional:    true,
					},
				},
			},
		},
		"upgrade_end_time": {
			Description: "End date of the software upgrade process.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"upgrade_grace_period": {
			Description: "Grace period in seconds before the automatic upgrade is initiated. The upgrade service uses the grace period to compute the upgrade start time when it receives an upgrade notfication from the Intersight. If there is an Upgrade Policy configured for the Intersight Appliance, then the upgrade service uses the policy to compute the upgrade start time. However, the upgrade start time cannot not exceed the limit enforced by the grace period.",
			Type:        schema.TypeInt,
			Optional:    true,
		},
		"upgrade_impact_duration": {
			Description: "Duration (in minutes) for which services will be disrupted.",
			Type:        schema.TypeInt,
			Optional:    true,
		},
		"upgrade_impact_enum": {
			Description: "UpgradeImpactEnum is used to indicate the kind of impact the upgrade has on currently running services on the appliance.\n* `None` - The upgrade has no effect on the system.\n* `Disruptive` - The services will not be functional during the upgrade.\n* `Disruptive-reboot` - The upgrade needs a reboot.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"upgrade_start_time": {
			Description: "Start date of the software upgrade process.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"nr_version": {
			Description: "Software upgrade manifest's version.",
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
	}
	model["results"] = &schema.Schema{
		Type:     schema.TypeList,
		Elem:     &schema.Resource{Schema: subSchema},
		Computed: true,
	}
	return &schema.Resource{
		ReadContext: dataSourceApplianceImageBundleRead,
		Schema:      model}
}

func dataSourceApplianceImageBundleRead(c context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)
	var de diag.Diagnostics
	var o = &models.ApplianceImageBundle{}
	if v, ok := d.GetOk("account_moid"); ok {
		x := (v.(string))
		o.SetAccountMoid(x)
	}

	if v, ok := d.GetOk("additional_properties"); ok {
		x := []byte(v.(string))
		var x1 interface{}
		err := json.Unmarshal(x, &x1)
		if err == nil && x1 != nil {
			o.AdditionalProperties = x1.(map[string]interface{})
		}
	}

	if v, ok := d.GetOk("ancestors"); ok {
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

	if v, ok := d.GetOk("ansible_packages"); ok {
		x := make([]models.OnpremImagePackage, 0)
		s := v.([]interface{})
		for i := 0; i < len(s); i++ {
			o := &models.OnpremImagePackage{}
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
			o.SetClassId("onprem.ImagePackage")
			if v, ok := l["object_type"]; ok {
				{
					x := (v.(string))
					o.SetObjectType(x)
				}
			}
			x = append(x, *o)
		}
		o.SetAnsiblePackages(x)
	}

	if v, ok := d.GetOkExists("auto_upgrade"); ok {
		x := (v.(bool))
		o.SetAutoUpgrade(x)
	}

	if v, ok := d.GetOk("class_id"); ok {
		x := (v.(string))
		o.SetClassId(x)
	}

	if v, ok := d.GetOk("create_time"); ok {
		x, _ := time.Parse(time.RFC1123, v.(string))
		o.SetCreateTime(x)
	}

	if v, ok := d.GetOk("dc_packages"); ok {
		x := make([]models.OnpremImagePackage, 0)
		s := v.([]interface{})
		for i := 0; i < len(s); i++ {
			o := &models.OnpremImagePackage{}
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
			o.SetClassId("onprem.ImagePackage")
			if v, ok := l["object_type"]; ok {
				{
					x := (v.(string))
					o.SetObjectType(x)
				}
			}
			x = append(x, *o)
		}
		o.SetDcPackages(x)
	}

	if v, ok := d.GetOk("debug_packages"); ok {
		x := make([]models.OnpremImagePackage, 0)
		s := v.([]interface{})
		for i := 0; i < len(s); i++ {
			o := &models.OnpremImagePackage{}
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
			o.SetClassId("onprem.ImagePackage")
			if v, ok := l["object_type"]; ok {
				{
					x := (v.(string))
					o.SetObjectType(x)
				}
			}
			x = append(x, *o)
		}
		o.SetDebugPackages(x)
	}

	if v, ok := d.GetOk("description"); ok {
		x := (v.(string))
		o.SetDescription(x)
	}

	if v, ok := d.GetOk("domain_group_moid"); ok {
		x := (v.(string))
		o.SetDomainGroupMoid(x)
	}

	if v, ok := d.GetOk("endpoint_packages"); ok {
		x := make([]models.OnpremImagePackage, 0)
		s := v.([]interface{})
		for i := 0; i < len(s); i++ {
			o := &models.OnpremImagePackage{}
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
			o.SetClassId("onprem.ImagePackage")
			if v, ok := l["object_type"]; ok {
				{
					x := (v.(string))
					o.SetObjectType(x)
				}
			}
			x = append(x, *o)
		}
		o.SetEndpointPackages(x)
	}

	if v, ok := d.GetOk("fingerprint"); ok {
		x := (v.(string))
		o.SetFingerprint(x)
	}

	if v, ok := d.GetOkExists("has_error"); ok {
		x := (v.(bool))
		o.SetHasError(x)
	}

	if v, ok := d.GetOk("infra_packages"); ok {
		x := make([]models.OnpremImagePackage, 0)
		s := v.([]interface{})
		for i := 0; i < len(s); i++ {
			o := &models.OnpremImagePackage{}
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
			o.SetClassId("onprem.ImagePackage")
			if v, ok := l["object_type"]; ok {
				{
					x := (v.(string))
					o.SetObjectType(x)
				}
			}
			x = append(x, *o)
		}
		o.SetInfraPackages(x)
	}

	if v, ok := d.GetOk("init_packages"); ok {
		x := make([]models.OnpremImagePackage, 0)
		s := v.([]interface{})
		for i := 0; i < len(s); i++ {
			o := &models.OnpremImagePackage{}
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
			o.SetClassId("onprem.ImagePackage")
			if v, ok := l["object_type"]; ok {
				{
					x := (v.(string))
					o.SetObjectType(x)
				}
			}
			x = append(x, *o)
		}
		o.SetInitPackages(x)
	}

	if v, ok := d.GetOk("mod_time"); ok {
		x, _ := time.Parse(time.RFC1123, v.(string))
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

	if v, ok := d.GetOk("notes"); ok {
		x := (v.(string))
		o.SetNotes(x)
	}

	if v, ok := d.GetOk("object_type"); ok {
		x := (v.(string))
		o.SetObjectType(x)
	}

	if v, ok := d.GetOk("owners"); ok {
		x := make([]string, 0)
		y := reflect.ValueOf(v)
		for i := 0; i < y.Len(); i++ {
			if y.Index(i).Interface() != nil {
				x = append(x, y.Index(i).Interface().(string))
			}
		}
		o.SetOwners(x)
	}

	if v, ok := d.GetOk("parent"); ok {
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
			o.SetClassId("")
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

	if v, ok := d.GetOk("permission_resources"); ok {
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

	if v, ok := d.GetOk("priority"); ok {
		x := (v.(string))
		o.SetPriority(x)
	}

	if v, ok := d.GetOk("release_time"); ok {
		x, _ := time.Parse(time.RFC1123, v.(string))
		o.SetReleaseTime(x)
	}

	if v, ok := d.GetOk("service_packages"); ok {
		x := make([]models.OnpremImagePackage, 0)
		s := v.([]interface{})
		for i := 0; i < len(s); i++ {
			o := &models.OnpremImagePackage{}
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
			o.SetClassId("onprem.ImagePackage")
			if v, ok := l["object_type"]; ok {
				{
					x := (v.(string))
					o.SetObjectType(x)
				}
			}
			x = append(x, *o)
		}
		o.SetServicePackages(x)
	}

	if v, ok := d.GetOk("shared_scope"); ok {
		x := (v.(string))
		o.SetSharedScope(x)
	}

	if v, ok := d.GetOk("status_message"); ok {
		x := (v.(string))
		o.SetStatusMessage(x)
	}

	if v, ok := d.GetOk("system_packages"); ok {
		x := make([]models.OnpremImagePackage, 0)
		s := v.([]interface{})
		for i := 0; i < len(s); i++ {
			o := &models.OnpremImagePackage{}
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
			o.SetClassId("onprem.ImagePackage")
			if v, ok := l["object_type"]; ok {
				{
					x := (v.(string))
					o.SetObjectType(x)
				}
			}
			x = append(x, *o)
		}
		o.SetSystemPackages(x)
	}

	if v, ok := d.GetOk("tags"); ok {
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

	if v, ok := d.GetOk("ui_packages"); ok {
		x := make([]models.OnpremImagePackage, 0)
		s := v.([]interface{})
		for i := 0; i < len(s); i++ {
			o := &models.OnpremImagePackage{}
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
			o.SetClassId("onprem.ImagePackage")
			if v, ok := l["object_type"]; ok {
				{
					x := (v.(string))
					o.SetObjectType(x)
				}
			}
			x = append(x, *o)
		}
		o.SetUiPackages(x)
	}

	if v, ok := d.GetOk("upgrade_end_time"); ok {
		x, _ := time.Parse(time.RFC1123, v.(string))
		o.SetUpgradeEndTime(x)
	}

	if v, ok := d.GetOkExists("upgrade_grace_period"); ok {
		x := int64(v.(int))
		o.SetUpgradeGracePeriod(x)
	}

	if v, ok := d.GetOkExists("upgrade_impact_duration"); ok {
		x := int64(v.(int))
		o.SetUpgradeImpactDuration(x)
	}

	if v, ok := d.GetOk("upgrade_impact_enum"); ok {
		x := (v.(string))
		o.SetUpgradeImpactEnum(x)
	}

	if v, ok := d.GetOk("upgrade_start_time"); ok {
		x, _ := time.Parse(time.RFC1123, v.(string))
		o.SetUpgradeStartTime(x)
	}

	if v, ok := d.GetOk("nr_version"); ok {
		x := (v.(string))
		o.SetVersion(x)
	}

	if v, ok := d.GetOk("version_context"); ok {
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
			o.SetClassId("")
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

	data, err := o.MarshalJSON()
	if err != nil {
		return diag.Errorf("json marshal of ApplianceImageBundle object failed with error : %s", err.Error())
	}
	countResponse, _, responseErr := conn.ApiClient.ApplianceApi.GetApplianceImageBundleList(conn.ctx).Filter(getRequestParams(data)).Inlinecount("allpages").Execute()
	if responseErr != nil {
		errorType := fmt.Sprintf("%T", responseErr)
		if strings.Contains(errorType, "GenericOpenAPIError") {
			responseErr := responseErr.(models.GenericOpenAPIError)
			return diag.Errorf("error occurred while fetching count of ApplianceImageBundle: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
		}
		return diag.Errorf("error occurred while fetching count of ApplianceImageBundle: %s", responseErr.Error())
	}
	count := countResponse.ApplianceImageBundleList.GetCount()
	if count == 0 {
		return diag.Errorf("your query for ApplianceImageBundle data source did not return any results. Please change your search criteria and try again")
	}
	var i int32
	var applianceImageBundleResults = make([]map[string]interface{}, count, count)
	var j = 0
	for i = 0; i < count; i += 100 {
		resMo, _, responseErr := conn.ApiClient.ApplianceApi.GetApplianceImageBundleList(conn.ctx).Filter(getRequestParams(data)).Top(100).Skip(i).Execute()
		if responseErr != nil {
			errorType := fmt.Sprintf("%T", responseErr)
			if strings.Contains(errorType, "GenericOpenAPIError") {
				responseErr := responseErr.(models.GenericOpenAPIError)
				return diag.Errorf("error occurred while fetching ApplianceImageBundle: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
			}
			return diag.Errorf("error occurred while fetching ApplianceImageBundle: %s", responseErr.Error())
		}
		results := resMo.ApplianceImageBundleList.GetResults()
		switch reflect.TypeOf(results).Kind() {
		case reflect.Slice:
			for i := 0; i < len(results); i++ {
				var s = results[i]
				var temp = make(map[string]interface{})
				temp["account_moid"] = (s.GetAccountMoid())
				temp["additional_properties"] = flattenAdditionalProperties(s.AdditionalProperties)

				temp["ancestors"] = flattenListMoBaseMoRelationship(s.GetAncestors(), d)

				temp["ansible_packages"] = flattenListOnpremImagePackage(s.GetAnsiblePackages(), d)
				temp["auto_upgrade"] = (s.GetAutoUpgrade())
				temp["class_id"] = (s.GetClassId())

				temp["create_time"] = (s.GetCreateTime()).String()

				temp["dc_packages"] = flattenListOnpremImagePackage(s.GetDcPackages(), d)

				temp["debug_packages"] = flattenListOnpremImagePackage(s.GetDebugPackages(), d)
				temp["description"] = (s.GetDescription())
				temp["domain_group_moid"] = (s.GetDomainGroupMoid())

				temp["endpoint_packages"] = flattenListOnpremImagePackage(s.GetEndpointPackages(), d)
				temp["fingerprint"] = (s.GetFingerprint())
				temp["has_error"] = (s.GetHasError())

				temp["infra_packages"] = flattenListOnpremImagePackage(s.GetInfraPackages(), d)

				temp["init_packages"] = flattenListOnpremImagePackage(s.GetInitPackages(), d)

				temp["mod_time"] = (s.GetModTime()).String()
				temp["moid"] = (s.GetMoid())
				temp["name"] = (s.GetName())
				temp["notes"] = (s.GetNotes())
				temp["object_type"] = (s.GetObjectType())
				temp["owners"] = (s.GetOwners())

				temp["parent"] = flattenMapMoBaseMoRelationship(s.GetParent(), d)

				temp["permission_resources"] = flattenListMoBaseMoRelationship(s.GetPermissionResources(), d)
				temp["priority"] = (s.GetPriority())

				temp["release_time"] = (s.GetReleaseTime()).String()

				temp["service_packages"] = flattenListOnpremImagePackage(s.GetServicePackages(), d)
				temp["shared_scope"] = (s.GetSharedScope())
				temp["status_message"] = (s.GetStatusMessage())

				temp["system_packages"] = flattenListOnpremImagePackage(s.GetSystemPackages(), d)

				temp["tags"] = flattenListMoTag(s.GetTags(), d)

				temp["ui_packages"] = flattenListOnpremImagePackage(s.GetUiPackages(), d)

				temp["upgrade_end_time"] = (s.GetUpgradeEndTime()).String()
				temp["upgrade_grace_period"] = (s.GetUpgradeGracePeriod())
				temp["upgrade_impact_duration"] = (s.GetUpgradeImpactDuration())
				temp["upgrade_impact_enum"] = (s.GetUpgradeImpactEnum())

				temp["upgrade_start_time"] = (s.GetUpgradeStartTime()).String()
				temp["nr_version"] = (s.GetVersion())

				temp["version_context"] = flattenMapMoVersionContext(s.GetVersionContext(), d)
				applianceImageBundleResults[j] = temp
				j += 1
			}
		}
	}
	log.Println("length of results: ", len(applianceImageBundleResults))
	if err := d.Set("results", applianceImageBundleResults); err != nil {
		return diag.Errorf("error occurred while setting results: %s", err.Error())
	}
	d.SetId(applianceImageBundleResults[0]["moid"].(string))
	return de
}
