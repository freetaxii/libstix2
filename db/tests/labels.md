# FreeTAXII/libstix2/db/tests/labels.md #

This documents the various unit tests that should be performed on your implementation 
of STIX labels 

## Unit Tests ##

### Test 1 ###

This test will verify that you are not accidentally duplicating records and thus 
you get back the right labels based on various conditions.

1. Start with a new SDO, version 0 with NO lables
2. Revision the SDO to version 1 and add 1 label "a"
3. Revision the SDO to version 2 and add 2 new labels "b", "c" so you have a total of 3
4. Revision the SDO to version 3 and add no new labels
5. Revision the SDO to version 4 and change label "c" to label "d"

Now you need to query the database for certain versions:

1. If you query for version 0 of the SDO you should get an object with NO labels
2. If you query for version 1 of the SDO you should get an object with a single label "a"
3. If you query for version 2 of the SDO you should get an object with 3 labels "a", "b", and "c"
4. If you query for version 3 of the SDO you should get an object with 3 labels "a", "b", and "c"
5. If you query for version 4 of the SDO you should get an object with 3 labels "a", "b", and "d"



## License ##

This is free software, licensed under the Apache License, Version 2.0.


## Copyright ##

Copyright 2016 Bret Jordan, All rights reserved.