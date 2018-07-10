// Copyright 2017 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

package datastore

import (
	"github.com/freetaxii/libstix2/resources"
)

/*
Datastorer - This interface enables access to the STIX/TAXII datastore.

Close           - This will close the connection to the datastore
GetSTIXObject   - This takes in a STIX ID and modified timestamp and returns a specific STIX object
AddSTIXObject   - This takes in a STIX object and writes it to the datastore
AddTAXIIObject  - This takes in a TAXII object and writes it to the datastore
GetBundle       - This will return a STIX Bundle object based on the query parameters provided
GetManifestData - This will return a TAXII Manifest resource based on the query parameters provide
*/
type Datastorer interface {
	Close() error
	GetSTIXObject(stixid, version string) (interface{}, error)
	AddSTIXObject(obj interface{}) error
	AddTAXIIObject(obj interface{}) error
	GetBundle(query resources.CollectionQuery) (*resources.CollectionQueryResult, error)
	GetManifestData(query resources.CollectionQuery) (*resources.CollectionQueryResult, error)
}

/*
MemCache - This struct will hold a cache of various datastore elements
that will be loaded during initialization and updated along the way.
*/
type MemCache struct {
	BaseObjectIDIndex int
	Collections       map[string]*resources.Collection
}

/*
The following constants define database tables names for a relational database.
All of the SQL statements and other code uses these constants, so it should be
pretty safe, if needed, to change the actual table names without problems.
*/
const (
	DB_TABLE_STIX_BASE_OBJECT           = "s_base_object"
	DB_TABLE_STIX_ATTACK_PATTERN        = "s_attack_pattern"
	DB_TABLE_STIX_CAMPAIGN              = "s_campaign"
	DB_TABLE_STIX_COURSE_OF_ACTION      = "s_course_of_action"
	DB_TABLE_STIX_IDENTITY              = "s_identity"
	DB_TABLE_STIX_IDENTITY_SECTORS      = "s_identity_sectors"
	DB_TABLE_STIX_INDICATOR             = "s_indicator"
	DB_TABLE_STIX_INTRUSION_SET         = "s_intrusion_set"
	DB_TABLE_STIX_LOCATION              = "s_location"
	DB_TABLE_STIX_MALWARE               = "s_malware"
	DB_TABLE_STIX_NOTE                  = "s_note"
	DB_TABLE_STIX_OBSERVED_DATA         = "s_observed_data"
	DB_TABLE_STIX_OPINION               = "s_opinion"
	DB_TABLE_STIX_REPORT                = "s_report"
	DB_TABLE_STIX_THREAT_ACTOR          = "s_threat_actor"
	DB_TABLE_STIX_THREAT_ACTOR_ROLES    = "s_threat_actor_roles"
	DB_TABLE_STIX_TOOL                  = "s_tool"
	DB_TABLE_STIX_VULNERABILITY         = "s_vulnerability"
	DB_TABLE_STIX_ALIASES               = "s_aliases"
	DB_TABLE_STIX_AUTHORS               = "s_authors"
	DB_TABLE_STIX_EXTERNAL_REFERENCES   = "s_external_references"
	DB_TABLE_STIX_GOALS                 = "s_goals"
	DB_TABLE_STIX_HASHES                = "s_hashes"
	DB_TABLE_STIX_KILL_CHAIN_PHASES     = "s_kill_chain_phases"
	DB_TABLE_STIX_LABELS                = "s_labels"
	DB_TABLE_STIX_OBJECT_MARKING_REFS   = "s_object_marking_refs"
	DB_TABLE_STIX_OBJECT_REFS           = "s_object_refs"
	DB_TABLE_STIX_SECONDARY_MOTIVATIONS = "s_secondary_motivations"
	DB_TABLE_STIX_PERSONAL_MOTIVATIONS  = "s_personal_motivations"

	DB_TABLE_VOCAB_ATTACK_MOTIVATIONS          = "v_attack_motivation"
	DB_TABLE_VOCAB_ATTACK_RESOURCE_LEVEL       = "v_attack_resource_level"
	DB_TABLE_VOCAB_IDENTITY_CLASS              = "v_identity_class"
	DB_TABLE_VOCAB_INDICATOR_LABEL             = "v_indicator_label"
	DB_TABLE_VOCAB_INDUSTRY_SECTOR             = "v_industry_sector"
	DB_TABLE_VOCAB_MALWARE_LABEL               = "v_malware_label"
	DB_TABLE_VOCAB_REPORT_LABEL                = "v_report_label"
	DB_TABLE_VOCAB_THREAT_ACTOR_LABEL          = "v_threat_actor_label"
	DB_TABLE_VOCAB_THREAT_ACTOR_ROLE           = "v_threat_actor_role"
	DB_TABLE_VOCAB_THREAT_ACTOR_SOPHISTICATION = "v_threat_actor_sophistication"
	DB_TABLE_VOCAB_TOOL_LABEL                  = "v_tool_label"

	DB_TABLE_TAXII_COLLECTIONS           = "t_collections"
	DB_TABLE_TAXII_COLLECTION_MEDIA_TYPE = "t_collection_media_type"
	DB_TABLE_TAXII_COLLECTION_DATA       = "t_collection_data"
	DB_TABLE_TAXII_MEDIA_TYPES           = "t_media_types"
)
