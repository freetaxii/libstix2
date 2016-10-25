/*
 FreeTAXII Database Structure
 Copyright 2016 Bret Jordan, All rights reserved.

 Use of this source code is governed by an Apache 2.0 license
 that can be found in the LICENSE file in the root of the source
 tree.

 Source Server         : FreeTAXII
 Source Server Type    : SQLite
 Source Server Version : 3007013
 Source Database       : main

 Target Server Type    : SQLite
 Target Server Version : 3007013
 File Encoding         : utf-8

 Date: 10/25/2016 17:19:53 PM

*/

PRAGMA foreign_keys = false;

-- ----------------------------
--  Table structure for "common_aliases"
-- ----------------------------
DROP TABLE IF EXISTS "common_aliases";
CREATE TABLE "common_aliases" (
	 "auto_id" integer PRIMARY KEY AUTOINCREMENT,
	 "parent_id" integer,
	 "type" text,
	 "id" text,
	 "version" integer,
	 "aliases" text
);

-- ----------------------------
--  Table structure for "common_external_references"
-- ----------------------------
DROP TABLE IF EXISTS "common_external_references";
CREATE TABLE "common_external_references" (
	 "auto_id" integer PRIMARY KEY AUTOINCREMENT,
	 "parent_id" integer,
	 "type" text,
	 "id" text,
	 "version" integer,
	 "source_name" text,
	 "description" text,
	 "url" text,
	 "external_id" text
);

-- ----------------------------
--  Table structure for "common_kill_chain_phases"
-- ----------------------------
DROP TABLE IF EXISTS "common_kill_chain_phases";
CREATE TABLE "common_kill_chain_phases" (
	 "auto_id" integer PRIMARY KEY AUTOINCREMENT,
	 "parent_id" integer,
	 "type" text,
	 "id" text,
	 "version" integer,
	 "kill_chain_name" text,
	 "phase_name" text
);

-- ----------------------------
--  Table structure for "common_labels"
-- ----------------------------
DROP TABLE IF EXISTS "common_labels";
CREATE TABLE "common_labels" (
	 "auto_id" integer PRIMARY KEY AUTOINCREMENT,
	 "parent_id" integer,
	 "type" text,
	 "id" text,
	 "version" integer,
	 "labels" text
);

-- ----------------------------
--  Table structure for "common_nationalities"
-- ----------------------------
DROP TABLE IF EXISTS "common_nationalities";
CREATE TABLE "common_nationalities" (
	 "auto_id" integer PRIMARY KEY AUTOINCREMENT,
	 "parent_id" integer,
	 "type" text,
	 "id" text,
	 "version" integer,
	 "nationalities" text
);

-- ----------------------------
--  Table structure for "common_object_marking_refs"
-- ----------------------------
DROP TABLE IF EXISTS "common_object_marking_refs";
CREATE TABLE "common_object_marking_refs" (
	 "auto_id" integer PRIMARY KEY AUTOINCREMENT,
	 "parent_id" integer,
	 "type" text,
	 "id" text,
	 "version" integer,
	 "object_marking_refs" text
);

-- ----------------------------
--  Table structure for "common_personal_motivations"
-- ----------------------------
DROP TABLE IF EXISTS "common_personal_motivations";
CREATE TABLE "common_personal_motivations" (
	 "auto_id" integer PRIMARY KEY AUTOINCREMENT,
	 "parent_id" integer,
	 "type" text,
	 "id" text,
	 "version" integer,
	 "personal_motivations" integer
);

-- ----------------------------
--  Table structure for "common_regions"
-- ----------------------------
DROP TABLE IF EXISTS "common_regions";
CREATE TABLE "common_regions" (
	 "auto_id" integer PRIMARY KEY AUTOINCREMENT,
	 "parent_id" integer,
	 "type" text,
	 "id" text,
	 "version" integer,
	 "regions" text
);

-- ----------------------------
--  Table structure for "common_secondary_motivations"
-- ----------------------------
DROP TABLE IF EXISTS "common_secondary_motivations";
CREATE TABLE "common_secondary_motivations" (
	 "auto_id" integer PRIMARY KEY AUTOINCREMENT,
	 "parent_id" integer,
	 "type" text,
	 "id" text,
	 "version" integer,
	 "secondary_motivations" integer
);

-- ----------------------------
--  Table structure for "common_sectors"
-- ----------------------------
DROP TABLE IF EXISTS "common_sectors";
CREATE TABLE "common_sectors" (
	 "auto_id" integer PRIMARY KEY AUTOINCREMENT,
	 "parent_id" integer,
	 "type" text,
	 "id" text,
	 "version" integer,
	 "sectors" integer
);

-- ----------------------------
--  Table structure for "sdo_attack_pattern"
-- ----------------------------
DROP TABLE IF EXISTS "sdo_attack_pattern";
CREATE TABLE "sdo_attack_pattern" (
	 "auto_id" integer PRIMARY KEY AUTOINCREMENT,
	 "date_added" text,
	 "id" text,
	 "created_by_ref" text,
	 "created" text,
	 "modified" text,
	 "version" integer,
	 "revoked" integer(1,0),
	 "name" text,
	 "description" text
);

-- ----------------------------
--  Table structure for "sdo_campaigns"
-- ----------------------------
DROP TABLE IF EXISTS "sdo_campaigns";
CREATE TABLE "sdo_campaigns" (
	 "auto_id" integer PRIMARY KEY AUTOINCREMENT,
	 "date" text,
	 "id" text,
	 "created_by_ref" text,
	 "created" text,
	 "modified" text,
	 "version" integer,
	 "revoked" integer(1,0),
	 "name" text,
	 "description" text,
	 "first_seen" text,
	 "first_seen_precision" text,
	 "objective" text
);

-- ----------------------------
--  Table structure for "sdo_course_of_action"
-- ----------------------------
DROP TABLE IF EXISTS "sdo_course_of_action";
CREATE TABLE "sdo_course_of_action" (
	 "auto_id" integer PRIMARY KEY AUTOINCREMENT,
	 "date" text,
	 "id" text,
	 "created_by_ref" text,
	 "created" text,
	 "modified" text,
	 "version" integer,
	 "revoked" integer(1,0),
	 "name" text,
	 "description" text
);

-- ----------------------------
--  Table structure for "sdo_identity"
-- ----------------------------
DROP TABLE IF EXISTS "sdo_identity";
CREATE TABLE "sdo_identity" (
	 "auto_id" integer PRIMARY KEY AUTOINCREMENT,
	 "date" text,
	 "id" text,
	 "created_by_ref" text,
	 "created" text,
	 "modified" text,
	 "version" integer,
	 "revoked" integer(1,0),
	 "name" text,
	 "description" text,
	 "identity_class" integer,
	 "contact_information" text
);

-- ----------------------------
--  Table structure for "sdo_indicators"
-- ----------------------------
DROP TABLE IF EXISTS "sdo_indicators";
CREATE TABLE "sdo_indicators" (
	 "auto_id" integer PRIMARY KEY AUTOINCREMENT,
	 "date" text,
	 "id" text,
	 "created_by_ref" text,
	 "created" text,
	 "modified" text,
	 "version" integer,
	 "revoked" integer(1,0),
	 "name" text,
	 "description" text,
	 "pattern_lang" integer,
	 "pattern_lang_version" text,
	 "pattern" text,
	 "valid_from" text,
	 "valid_from_precision" text,
	 "valid_until" text,
	 "valid_until_precision" text
);

-- ----------------------------
--  Table structure for "sdo_intrusion_sets"
-- ----------------------------
DROP TABLE IF EXISTS "sdo_intrusion_sets";
CREATE TABLE "sdo_intrusion_sets" (
	 "auto_id" integer PRIMARY KEY AUTOINCREMENT,
	 "date" text,
	 "id" text,
	 "created_by_ref" text,
	 "created" text,
	 "modified" text,
	 "version" integer,
	 "revoked" integer(1,0),
	 "name" text,
	 "description" text,
	 "first_seen" text,
	 "first_seen_precision" text,
	 "resource_level" integer,
	 "primary_motivation" integer,
	 "region" text,
	 "country" text
);

-- ----------------------------
--  Table structure for "sdo_malware"
-- ----------------------------
DROP TABLE IF EXISTS "sdo_malware";
CREATE TABLE "sdo_malware" (
	 "auto_id" integer PRIMARY KEY AUTOINCREMENT,
	 "date" text,
	 "id" text,
	 "created_by_ref" text,
	 "created" text,
	 "modified" text,
	 "version" integer,
	 "revoked" integer(1,0),
	 "name" text,
	 "description" text
);

-- ----------------------------
--  Table structure for "sdo_observed_data"
-- ----------------------------
DROP TABLE IF EXISTS "sdo_observed_data";
CREATE TABLE "sdo_observed_data" (
	 "auto_id" integer PRIMARY KEY AUTOINCREMENT,
	 "date" text,
	 "id" text,
	 "created_by_ref" text,
	 "created" text,
	 "modified" text,
	 "version" integer,
	 "revoked" integer(1,0),
	 "first_observed" text,
	 "last_observed" text,
	 "number_observed" integer,
	 "objects" text
);

-- ----------------------------
--  Table structure for "sdo_reports"
-- ----------------------------
DROP TABLE IF EXISTS "sdo_reports";
CREATE TABLE "sdo_reports" (
	 "auto_id" integer PRIMARY KEY AUTOINCREMENT,
	 "date" text,
	 "id" text,
	 "created_by_ref" text,
	 "created" text,
	 "modified" text,
	 "version" integer,
	 "revoked" integer(1,0),
	 "name" text,
	 "description" text,
	 "published" text
);

-- ----------------------------
--  Table structure for "sdo_reports_object_refs"
-- ----------------------------
DROP TABLE IF EXISTS "sdo_reports_object_refs";
CREATE TABLE "sdo_reports_object_refs" (
	 "auto_id" integer PRIMARY KEY AUTOINCREMENT,
	 "parent_id" integer,
	 "date" text,
	 "id" text,
	 "version" integer,
	 "object_refs" text
);

-- ----------------------------
--  Table structure for "sdo_threat_actors"
-- ----------------------------
DROP TABLE IF EXISTS "sdo_threat_actors";
CREATE TABLE "sdo_threat_actors" (
	 "auto_id" integer PRIMARY KEY AUTOINCREMENT,
	 "date" text,
	 "id" text,
	 "created_by_ref" text,
	 "created" text,
	 "modified" text,
	 "version" integer,
	 "revoked" integer(1,0),
	 "name" text,
	 "description" text,
	 "sophistication" integer,
	 "resource_level" integer,
	 "primary_motivation" integer
);

-- ----------------------------
--  Table structure for "sdo_threat_actors_goals"
-- ----------------------------
DROP TABLE IF EXISTS "sdo_threat_actors_goals";
CREATE TABLE "sdo_threat_actors_goals" (
	 "auto_id" integer PRIMARY KEY AUTOINCREMENT,
	 "parent_id" integer,
	 "date" text,
	 "id" text,
	 "version" integer,
	 "goals" text
);

-- ----------------------------
--  Table structure for "sdo_tool"
-- ----------------------------
DROP TABLE IF EXISTS "sdo_tool";
CREATE TABLE "sdo_tool" (
	 "auto_id" integer PRIMARY KEY AUTOINCREMENT,
	 "date" text,
	 "id" text,
	 "created_by_ref" text,
	 "created" text,
	 "modified" text,
	 "version" integer,
	 "revoked" integer(1,0),
	 "name" text,
	 "description" text,
	 "tool_version" text
);

-- ----------------------------
--  Table structure for "sdo_vulnerability"
-- ----------------------------
DROP TABLE IF EXISTS "sdo_vulnerability";
CREATE TABLE "sdo_vulnerability" (
	 "auto_id" integer PRIMARY KEY AUTOINCREMENT,
	 "date" text,
	 "id" text,
	 "created_by_ref" text,
	 "created" text,
	 "modified" text,
	 "version" integer,
	 "revoked" integer(1,0),
	 "name" text,
	 "description" text
);

-- ----------------------------
--  Table structure for "vocab_attack_motivation"
-- ----------------------------
DROP TABLE IF EXISTS "vocab_attack_motivation";
CREATE TABLE "vocab_attack_motivation" (
	 "auto_id" integer PRIMARY KEY AUTOINCREMENT,
	 "value" text,
	 "active" integer(1,0),
	 "major_version" integer,
	 "minor_version" integer
);

-- ----------------------------
--  Records of "vocab_attack_motivation"
-- ----------------------------
BEGIN;
INSERT INTO "vocab_attack_motivation" VALUES (1, 'accidental', 1, 2, 0);
INSERT INTO "vocab_attack_motivation" VALUES (2, 'coercion', 1, 2, 0);
INSERT INTO "vocab_attack_motivation" VALUES (3, 'dominance', 1, 2, 0);
INSERT INTO "vocab_attack_motivation" VALUES (4, 'ideology', 1, 2, 0);
INSERT INTO "vocab_attack_motivation" VALUES (5, 'notoriety', 1, 2, 0);
INSERT INTO "vocab_attack_motivation" VALUES (6, 'organizational-gain', 1, 2, 0);
INSERT INTO "vocab_attack_motivation" VALUES (7, 'personal-gain', 1, 2, 0);
INSERT INTO "vocab_attack_motivation" VALUES (8, 'personal-satisfaction', 1, 2, 0);
INSERT INTO "vocab_attack_motivation" VALUES (9, 'revenge', 1, 2, 0);
INSERT INTO "vocab_attack_motivation" VALUES (10, 'unpredictable', 1, 2, 0);
COMMIT;

-- ----------------------------
--  Table structure for "vocab_attack_resource_level"
-- ----------------------------
DROP TABLE IF EXISTS "vocab_attack_resource_level";
CREATE TABLE "vocab_attack_resource_level" (
	 "auto_id" integer PRIMARY KEY AUTOINCREMENT,
	 "value" text,
	 "active" integer(1,0),
	 "major_version" integer,
	 "minor_version" integer
);

-- ----------------------------
--  Records of "vocab_attack_resource_level"
-- ----------------------------
BEGIN;
INSERT INTO "vocab_attack_resource_level" VALUES (1, 'individual', 1, 2, 0);
INSERT INTO "vocab_attack_resource_level" VALUES (2, 'club', 1, 2, 0);
INSERT INTO "vocab_attack_resource_level" VALUES (3, 'content', 1, 2, 0);
INSERT INTO "vocab_attack_resource_level" VALUES (4, 'team', 1, 2, 0);
INSERT INTO "vocab_attack_resource_level" VALUES (5, 'organization', 1, 2, 0);
INSERT INTO "vocab_attack_resource_level" VALUES (6, 'government', 1, 2, 0);
COMMIT;

-- ----------------------------
--  Table structure for "vocab_identity_class"
-- ----------------------------
DROP TABLE IF EXISTS "vocab_identity_class";
CREATE TABLE "vocab_identity_class" (
	 "auto_id" integer PRIMARY KEY AUTOINCREMENT,
	 "value" text,
	 "active" integer(1,0),
	 "major_version" integer,
	 "minor_version" integer
);

-- ----------------------------
--  Records of "vocab_identity_class"
-- ----------------------------
BEGIN;
INSERT INTO "vocab_identity_class" VALUES (1, 'individual', 1, 2, 0);
INSERT INTO "vocab_identity_class" VALUES (2, 'group', 1, 2, 0);
INSERT INTO "vocab_identity_class" VALUES (3, 'organization', 1, 2, 0);
INSERT INTO "vocab_identity_class" VALUES (4, 'class', 1, 2, 0);
INSERT INTO "vocab_identity_class" VALUES (5, 'unknown', 1, 2, 0);
COMMIT;

-- ----------------------------
--  Table structure for "vocab_indicator_label"
-- ----------------------------
DROP TABLE IF EXISTS "vocab_indicator_label";
CREATE TABLE "vocab_indicator_label" (
	 "auto_id" integer PRIMARY KEY AUTOINCREMENT,
	 "value" text,
	 "active" integer(1,0),
	 "major_version" integer,
	 "minor_version" integer
);

-- ----------------------------
--  Records of "vocab_indicator_label"
-- ----------------------------
BEGIN;
INSERT INTO "vocab_indicator_label" VALUES (1, 'anomalous-activity', 1, 2, 0);
INSERT INTO "vocab_indicator_label" VALUES (2, 'anonymization', 1, 2, 0);
INSERT INTO "vocab_indicator_label" VALUES (3, 'benign', 1, 2, 0);
INSERT INTO "vocab_indicator_label" VALUES (4, 'compromised', 1, 2, 0);
INSERT INTO "vocab_indicator_label" VALUES (5, 'malicious-activity', 1, 2, 0);
INSERT INTO "vocab_indicator_label" VALUES (6, 'attribution', 1, 2, 0);
COMMIT;

-- ----------------------------
--  Table structure for "vocab_industry_sector"
-- ----------------------------
DROP TABLE IF EXISTS "vocab_industry_sector";
CREATE TABLE "vocab_industry_sector" (
	 "auto_id" integer PRIMARY KEY AUTOINCREMENT,
	 "sector" text,
	 "active" integer(1,0),
	 "major_version" integer,
	 "minor_version" integer
);

-- ----------------------------
--  Records of "vocab_industry_sector"
-- ----------------------------
BEGIN;
INSERT INTO "vocab_industry_sector" VALUES (1, 'agriculture', 1, 2, 0);
INSERT INTO "vocab_industry_sector" VALUES (2, 'aerospace', 1, 2, 0);
INSERT INTO "vocab_industry_sector" VALUES (3, 'automotive', 1, 2, 0);
INSERT INTO "vocab_industry_sector" VALUES (4, 'communications', 1, 2, 0);
INSERT INTO "vocab_industry_sector" VALUES (5, 'construction', 1, 2, 0);
INSERT INTO "vocab_industry_sector" VALUES (6, 'defense', 1, 2, 0);
INSERT INTO "vocab_industry_sector" VALUES (7, 'education', 1, 2, 0);
INSERT INTO "vocab_industry_sector" VALUES (8, 'energy', 1, 2, 0);
INSERT INTO "vocab_industry_sector" VALUES (9, 'entertainment', 1, 2, 0);
INSERT INTO "vocab_industry_sector" VALUES (10, 'financial-services', 1, 2, 0);
INSERT INTO "vocab_industry_sector" VALUES (11, 'government-national', 1, 2, 0);
INSERT INTO "vocab_industry_sector" VALUES (12, 'government-regional', 1, 2, 0);
INSERT INTO "vocab_industry_sector" VALUES (13, 'government-local', 1, 2, 0);
INSERT INTO "vocab_industry_sector" VALUES (14, 'government-public-services', 1, 2, 0);
INSERT INTO "vocab_industry_sector" VALUES (15, 'healthcare', 1, 2, 0);
INSERT INTO "vocab_industry_sector" VALUES (16, 'hospitality-leisure', 1, 2, 0);
INSERT INTO "vocab_industry_sector" VALUES (17, 'infrastructure', 1, 2, 0);
INSERT INTO "vocab_industry_sector" VALUES (18, 'insurance', 1, 2, 0);
INSERT INTO "vocab_industry_sector" VALUES (19, 'manufacturing', 1, 2, 0);
INSERT INTO "vocab_industry_sector" VALUES (20, 'mining', 1, 2, 0);
INSERT INTO "vocab_industry_sector" VALUES (21, 'non-profit', 1, 2, 0);
INSERT INTO "vocab_industry_sector" VALUES (22, 'pharmaceuticals', 1, 2, 0);
INSERT INTO "vocab_industry_sector" VALUES (23, 'retail', 1, 2, 0);
INSERT INTO "vocab_industry_sector" VALUES (24, 'technology', 1, 2, 0);
INSERT INTO "vocab_industry_sector" VALUES (25, 'telecommunications', 1, 2, 0);
INSERT INTO "vocab_industry_sector" VALUES (26, 'transportation', 1, 2, 0);
INSERT INTO "vocab_industry_sector" VALUES (27, 'utilities', 1, 2, 0);
COMMIT;

-- ----------------------------
--  Table structure for "vocab_malware_label"
-- ----------------------------
DROP TABLE IF EXISTS "vocab_malware_label";
CREATE TABLE "vocab_malware_label" (
	 "auto_id" integer PRIMARY KEY AUTOINCREMENT,
	 "value" text,
	 "active" integer(1,0),
	 "major_version" integer,
	 "minor_version" integer
);

-- ----------------------------
--  Records of "vocab_malware_label"
-- ----------------------------
BEGIN;
INSERT INTO "vocab_malware_label" VALUES (1, 'adware', 1, 2, 0);
INSERT INTO "vocab_malware_label" VALUES (2, 'backdoor', 1, 2, 0);
INSERT INTO "vocab_malware_label" VALUES (3, 'bot', 1, 2, 0);
INSERT INTO "vocab_malware_label" VALUES (4, 'ddos', 1, 2, 0);
INSERT INTO "vocab_malware_label" VALUES (5, 'dropper', 1, 2, 0);
INSERT INTO "vocab_malware_label" VALUES (6, 'exploit-kit', 1, 2, 0);
INSERT INTO "vocab_malware_label" VALUES (7, 'keylogger', 1, 2, 0);
INSERT INTO "vocab_malware_label" VALUES (8, 'ransomware', 1, 2, 0);
INSERT INTO "vocab_malware_label" VALUES (9, 'remote-access-trojan', 1, 2, 0);
INSERT INTO "vocab_malware_label" VALUES (10, 'resource-exploitation', 1, 2, 0);
INSERT INTO "vocab_malware_label" VALUES (11, 'rogue-security-software', 1, 2, 0);
INSERT INTO "vocab_malware_label" VALUES (12, 'rootkit', 1, 2, 0);
INSERT INTO "vocab_malware_label" VALUES (13, 'screen-capture', 1, 2, 0);
INSERT INTO "vocab_malware_label" VALUES (14, 'spyware', 1, 2, 0);
INSERT INTO "vocab_malware_label" VALUES (15, 'trojan', 1, 2, 0);
INSERT INTO "vocab_malware_label" VALUES (16, 'virus', 1, 2, 0);
INSERT INTO "vocab_malware_label" VALUES (17, 'worm', 1, 2, 0);
COMMIT;

-- ----------------------------
--  Table structure for "vocab_pattern_lang"
-- ----------------------------
DROP TABLE IF EXISTS "vocab_pattern_lang";
CREATE TABLE "vocab_pattern_lang" (
	 "auto_id" integer PRIMARY KEY AUTOINCREMENT,
	 "value" text,
	 "active" integer(1,0),
	 "major_version" integer,
	 "minor_version" integer
);

-- ----------------------------
--  Records of "vocab_pattern_lang"
-- ----------------------------
BEGIN;
INSERT INTO "vocab_pattern_lang" VALUES (1, 'stix', 1, 2, 0);
INSERT INTO "vocab_pattern_lang" VALUES (2, 'openioc', 1, 2, 0);
INSERT INTO "vocab_pattern_lang" VALUES (3, 'snort', 1, 2, 0);
INSERT INTO "vocab_pattern_lang" VALUES (4, 'yara', 1, 2, 0);
COMMIT;

-- ----------------------------
--  Table structure for "vocab_report_label"
-- ----------------------------
DROP TABLE IF EXISTS "vocab_report_label";
CREATE TABLE "vocab_report_label" (
	 "auto_id" integer PRIMARY KEY AUTOINCREMENT,
	 "value" text,
	 "active" integer(1,0),
	 "major_version" integer,
	 "minor_version" integer
);

-- ----------------------------
--  Records of "vocab_report_label"
-- ----------------------------
BEGIN;
INSERT INTO "vocab_report_label" VALUES (1, 'threat-report', 1, 2, 0);
INSERT INTO "vocab_report_label" VALUES (2, 'attack-pattern', 1, 2, 0);
INSERT INTO "vocab_report_label" VALUES (3, 'campaign', 1, 2, 0);
INSERT INTO "vocab_report_label" VALUES (4, 'indicator', 1, 2, 0);
INSERT INTO "vocab_report_label" VALUES (5, 'malware', 1, 2, 0);
INSERT INTO "vocab_report_label" VALUES (6, 'observed-data', 1, 2, 0);
INSERT INTO "vocab_report_label" VALUES (7, 'threat-actor', 1, 2, 0);
INSERT INTO "vocab_report_label" VALUES (8, 'tool', 1, 2, 0);
INSERT INTO "vocab_report_label" VALUES (9, 'victim-target', 1, 2, 0);
INSERT INTO "vocab_report_label" VALUES (10, 'vulnerability', 1, 2, 0);
COMMIT;

-- ----------------------------
--  Table structure for "vocab_threat_actor_label"
-- ----------------------------
DROP TABLE IF EXISTS "vocab_threat_actor_label";
CREATE TABLE "vocab_threat_actor_label" (
	 "auto_id" integer PRIMARY KEY AUTOINCREMENT,
	 "value" text,
	 "active" integer(1,0),
	 "major_version" integer,
	 "minor_version" integer
);

-- ----------------------------
--  Records of "vocab_threat_actor_label"
-- ----------------------------
BEGIN;
INSERT INTO "vocab_threat_actor_label" VALUES (1, 'activist', 1, 2, 0);
INSERT INTO "vocab_threat_actor_label" VALUES (2, 'competitor', 1, 2, 0);
INSERT INTO "vocab_threat_actor_label" VALUES (3, 'crime-syndicate', 1, 2, 0);
INSERT INTO "vocab_threat_actor_label" VALUES (4, 'criminal', 1, 2, 0);
INSERT INTO "vocab_threat_actor_label" VALUES (5, 'hacker', 1, 2, 0);
INSERT INTO "vocab_threat_actor_label" VALUES (6, 'insider-accidental', 1, 2, 0);
INSERT INTO "vocab_threat_actor_label" VALUES (7, 'insider-disgruntled', 1, 2, 0);
INSERT INTO "vocab_threat_actor_label" VALUES (8, 'nation-state', 1, 2, 0);
INSERT INTO "vocab_threat_actor_label" VALUES (9, 'sensationalist', 1, 2, 0);
INSERT INTO "vocab_threat_actor_label" VALUES (10, 'spy', 1, 2, 0);
INSERT INTO "vocab_threat_actor_label" VALUES (11, 'terrorist', 1, 2, 0);
COMMIT;

-- ----------------------------
--  Table structure for "vocab_threat_actor_role"
-- ----------------------------
DROP TABLE IF EXISTS "vocab_threat_actor_role";
CREATE TABLE "vocab_threat_actor_role" (
	 "auto_id" integer PRIMARY KEY AUTOINCREMENT,
	 "value" text,
	 "active" integer(1,0),
	 "major_version" integer,
	 "minor_version" integer
);

-- ----------------------------
--  Records of "vocab_threat_actor_role"
-- ----------------------------
BEGIN;
INSERT INTO "vocab_threat_actor_role" VALUES (1, 'agent', 1, 2, 0);
INSERT INTO "vocab_threat_actor_role" VALUES (2, 'director', 1, 2, 0);
INSERT INTO "vocab_threat_actor_role" VALUES (3, 'independent', 1, 2, 0);
INSERT INTO "vocab_threat_actor_role" VALUES (4, 'infrastructure-architect', 1, 2, 0);
INSERT INTO "vocab_threat_actor_role" VALUES (5, 'infrastructure-operator', 1, 2, 0);
INSERT INTO "vocab_threat_actor_role" VALUES (6, 'malware-author', 1, 2, 0);
INSERT INTO "vocab_threat_actor_role" VALUES (7, 'sponsor', 1, 2, 0);
COMMIT;

-- ----------------------------
--  Table structure for "vocab_threat_actor_sophistication"
-- ----------------------------
DROP TABLE IF EXISTS "vocab_threat_actor_sophistication";
CREATE TABLE "vocab_threat_actor_sophistication" (
	 "auto_id" integer PRIMARY KEY AUTOINCREMENT,
	 "value" text,
	 "active" integer(1,0),
	 "major_version" integer,
	 "minor_version" integer
);

-- ----------------------------
--  Records of "vocab_threat_actor_sophistication"
-- ----------------------------
BEGIN;
INSERT INTO "vocab_threat_actor_sophistication" VALUES (1, 'none', 1, 2, 0);
INSERT INTO "vocab_threat_actor_sophistication" VALUES (2, 'minimal', 1, 2, 0);
INSERT INTO "vocab_threat_actor_sophistication" VALUES (3, 'intermediate', 1, 2, 0);
INSERT INTO "vocab_threat_actor_sophistication" VALUES (4, 'advanced', 1, 2, 0);
INSERT INTO "vocab_threat_actor_sophistication" VALUES (5, 'expert', 1, 2, 0);
INSERT INTO "vocab_threat_actor_sophistication" VALUES (6, 'innovator', 1, 2, 0);
INSERT INTO "vocab_threat_actor_sophistication" VALUES (7, 'strategic', 1, 2, 0);
COMMIT;

-- ----------------------------
--  Table structure for "vocab_tool_label"
-- ----------------------------
DROP TABLE IF EXISTS "vocab_tool_label";
CREATE TABLE "vocab_tool_label" (
	 "auto_id" integer PRIMARY KEY AUTOINCREMENT,
	 "value" text,
	 "active" integer(1,0),
	 "major_version" integer,
	 "minor_version" integer
);

-- ----------------------------
--  Records of "vocab_tool_label"
-- ----------------------------
BEGIN;
INSERT INTO "vocab_tool_label" VALUES (1, 'denial-of-service', 1, 2, 0);
INSERT INTO "vocab_tool_label" VALUES (2, 'exploitation', 1, 2, 0);
INSERT INTO "vocab_tool_label" VALUES (3, 'information-gathering', 1, 2, 0);
INSERT INTO "vocab_tool_label" VALUES (4, 'network-capture', 1, 2, 0);
INSERT INTO "vocab_tool_label" VALUES (5, 'credential-exploitation', 1, 2, 0);
INSERT INTO "vocab_tool_label" VALUES (6, 'remote-access', 1, 2, 0);
INSERT INTO "vocab_tool_label" VALUES (7, 'vulnerability-scanning', 1, 2, 0);
COMMIT;

PRAGMA foreign_keys = true;
