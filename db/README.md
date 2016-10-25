# FreeTAXII/libstix2/db #

This is a simple SQLite relational database that can be used to store STIX 2.0
content.

## Table Structures ##

You will find that I am using the same STIX property names in the table structures 
even when that means the name is plural but the value is singular. I did this on 
purpose to make it easier to match up with STIX property names when marshalling 
and unmarshalling the content. 


### SDO Tables ###

#### sdo_attack_pattern ####
| Field                 | Type     | Notes
| --------------------- | -------- | -------
| auto_id               | integer  | Auto incrementing key field  
| date_added            | test     | The date in STIX timestamp format when this record was added to the database
| id                    | text     | See STIX
| created_by_ref        | text     | See STIX
| created               | text     | See STIX
| modified              | text     | See STIX
| version               | integer  | See STIX
| revoked               | integer  | See STIX (0=false, 1=true)
| name                  | text     | See STIX 
| description           | text     | See STIX 



---


### Common Property Tables ####

In each of these tables you will see a the following fields:

1. `parent_id` field. This is the auto_id number from the parent table. This is needed to make sure you match the right record in the case where a duplicate record may be found.
2. `type` field. This is the parent record type. This will allow you to easily query this table for all records tied to for example an Indicator. By recording it here, this prevents you from having to figure it out from the STIX ID.
3. `id` field. This is the STIX ID from the parent record. This is recorded here as well in case you need to go backwards in your graph search.
4. `version` field. This is the version of the STIX ID from the parent record that this record is tied to. You need both the STIX ID and its Version to make sure you match against the right record. 


#### common_external_references ####
| Field                 | Type     | Notes
| --------------------- | -------- | -------
| auto_id               | integer  | Auto incrementing key field  
| parent_id             | integer  | The auto_id number from the parent table
| type                  | text     | The parent record type
| id                    | text     | The STIX ID from the parent record
| version               | integer  | The version of the STIX ID from the parent record
| source_name           | text     | See STIX 
| description           | text     | See STIX 
| url                   | text     | See STIX 
| external_id           | text     | See STIX 



#### common_kill_chain_phases ####
| Field                 | Type     | Notes
| --------------------- | -------- | -------
| auto_id               | integer  | Auto incrementing key field  
| parent_id             | integer  | The auto_id number from the parent table
| type                  | text     | The parent record type
| id                    | text     | The STIX ID from the parent record
| version               | integer  | The version of the STIX ID from the parent record
| kill_chain_name       | text     | See STIX 
| phase_name            | text     | See STIX 



#### common_labels ####

| Field                 | Type     | Notes
| --------------------- | -------- | -------
| auto_id               | integer  | Auto incrementing key field  
| parent_id             | integer  | The auto_id number from the parent table
| type                  | text     | The parent record type
| id                    | text     | The STIX ID from the parent record
| version               | integer  | The version of the STIX ID from the parent record 
| labels                | text     | See STIX



#### common_object_marking_refs ####

| Field                 | Type     | Notes
| --------------------- | -------- | -------
| auto_id               | integer  | Auto incrementing key field  
| parent_id             | integer  | The auto_id number from the parent table
| type                  | text     | The parent record type
| id                    | text     | The STIX ID from the parent record
| version               | integer  | The version of the STIX ID from the parent record
| object_marking_refs   | text     | See STIX





## License ##

This is free software, licensed under the Apache License, Version 2.0.


## Copyright ##

Copyright 2016 Bret Jordan, All rights reserved.