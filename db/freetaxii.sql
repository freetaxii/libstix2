/*
 FreeTAXII Database Structure
 Copyright 2016 Bret Jordan, All rights reserved.

 Use of this source code is governed by an Apache 2.0 license
 that can be found in the LICENSE file in the root of the source
 tree.
*/

PRAGMA foreign_keys = false;

-- ----------------------------
--  Table structure for "common_external_references"
-- ----------------------------
DROP TABLE IF EXISTS "common_external_references";
CREATE TABLE "common_external_references" (
	 "auto_id" integer PRIMARY KEY AUTOINCREMENT,
	 "parent_id" integer,
	 "date" text,
	 "id" text,
	 "version" integer,
	 "source_name" text,
	 "description" text,
	 "url" text,
	 "external_id" text
);

-- ----------------------------
--  Table structure for "common_labels"
-- ----------------------------
DROP TABLE IF EXISTS "common_labels";
CREATE TABLE "common_labels" (
	 "auto_id" integer PRIMARY KEY AUTOINCREMENT,
	 "parent_id" integer,
	 "date" text,
	 "id" text,
	 "version" integer,
	 "labels" text
);

-- ----------------------------
--  Table structure for "common_object_marking_refs"
-- ----------------------------
DROP TABLE IF EXISTS "common_object_marking_refs";
CREATE TABLE "common_object_marking_refs" (
	 "auto_id" integer PRIMARY KEY AUTOINCREMENT,
	 "parent_id" integer,
	 "date" text,
	 "id" text,
	 "version" integer,
	 "object_marking_refs" text
);

-- ----------------------------
--  Table structure for "reports"
-- ----------------------------
DROP TABLE IF EXISTS "reports";
CREATE TABLE "reports" (
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
--  Table structure for "reports_object_refs"
-- ----------------------------
DROP TABLE IF EXISTS "reports_object_refs";
CREATE TABLE "reports_object_refs" (
	 "auto_id" integer PRIMARY KEY AUTOINCREMENT,
	 "parent_id" integer,
	 "date" text,
	 "id" text,
	 "version" integer,
	 "object_refs" text
);

PRAGMA foreign_keys = true;
