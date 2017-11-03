// Copyright 2017 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license
// that can be found in the LICENSE file in the root of the source
// tree.

package sqlite3

const (
	DB_TABLE_STIX_BASE_OBJECT           = "stix_base_object"
	DB_TABLE_STIX_ATTACK_PATTERN        = "sdo_attack_pattern"
	DB_TABLE_STIX_CAMPAIGN              = "sdo_campaign"
	DB_TABLE_STIX_COURSE_OF_ACTION      = "sdo_course_of_action"
	DB_TABLE_STIX_IDENTITY              = "sdo_identity"
	DB_TABLE_STIX_IDENTITY_SECTORS      = "identity_sectors"
	DB_TABLE_STIX_INDICATOR             = "sdo_indicator"
	DB_TABLE_STIX_INTRUSION_SET         = "sdo_intrusion_set"
	DB_TABLE_STIX_LOCATION              = "sdo_location"
	DB_TABLE_STIX_MALWARE               = "sdo_malware"
	DB_TABLE_STIX_NOTE                  = "sdo_note"
	DB_TABLE_STIX_OBSERVED_DATA         = "sdo_observed_data"
	DB_TABLE_STIX_OPINION               = "sdo_opinion"
	DB_TABLE_STIX_REPORT                = "sdo_report"
	DB_TABLE_STIX_THREAT_ACTOR          = "sdo_threat_actor"
	DB_TABLE_STIX_THREAT_ACTOR_ROLES    = "threat_actor_roles"
	DB_TABLE_STIX_TOOL                  = "sdo_tool"
	DB_TABLE_STIX_VULNERABILITY         = "sdo_vulnerability"
	DB_TABLE_STIX_ALIASES               = "aliases"
	DB_TABLE_STIX_AUTHORS               = "authors"
	DB_TABLE_STIX_EXTERNAL_REFERENCES   = "external_references"
	DB_TABLE_STIX_GOALS                 = "goals"
	DB_TABLE_STIX_HASHES                = "hashes"
	DB_TABLE_STIX_KILL_CHAIN_PHASES     = "kill_chain_phases"
	DB_TABLE_STIX_LABELS                = "labels"
	DB_TABLE_STIX_OBJECT_MARKING_REFS   = "object_marking_refs"
	DB_TABLE_STIX_OBJECT_REFS           = "object_refs"
	DB_TABLE_STIX_SECONDARY_MOTIVATIONS = "secondary_motivations"
	DB_TABLE_STIX_PERSONAL_MOTIVATIONS  = "personal_motivations"

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

	DB_TABLE_TAXII_COLLECTION            = "taxii_collection"
	DB_TABLE_TAXII_COLLECTION_MEDIA_TYPE = "taxii_collection_media_type"
	DB_TABLE_TAXII_COLLECTION_CONTENT    = "taxii_collection_content"
)
