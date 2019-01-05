# FreeTAXII/libstix2/datastore #

## Design Problem 1 ##

When designing a relational database for STIX content, one might decide to make
a different table for each SDO. This would make sense, since each object in STIX
has different properties. However, doing this will require sub-tables for things
like "labels" to be created for every single SDO table that is created. 

If you tried to do this without creating a separate "labels" tables for each SDO
there would be no easy way to link them back to the right row in the right table.

Let me explain...

One could not simply use the ROWID, as the ROWID would not be unique across SDO
tables. You could also not just use the STIX ID, as different versions of the 
object will have different rows in the database. You would need some sort of 
deterministic ID or hash that is based on the STIX ID and the STIX modified 
timestamp (and maybe something else). You would also need to record the table 
name that the record was located in, otherwise you would need to search every 
table for the record in question (or keep a master index table some how). 

### My Solution ###

The way I have decided to solve this problem is to create a base table that 
contains just the common properties that are part of every STIX object (SDO and
SRO). This will then allow me to link all of the SDO Tables and extra sub-tables
directly to this base table. Yes, this means that every addition to the database
will result in multiple transactions, since you will need to get the right ID 
value for the base table record insertion to add that as a foreign ID in the 
corresponding SDO table or sub-table. 

## License ##

This is free software, licensed under the Apache License, Version 2.0.


## Copyright ##

Copyright 2015-2019 Bret Jordan, All rights reserved.
