// Copyright 2017 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

package sqlite3

import (
	"bytes"
	"database/sql"
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/freetaxii/libstix2/common/stixid"
	"github.com/freetaxii/libstix2/common/timestamp"
	"github.com/freetaxii/libstix2/defs"
	"github.com/freetaxii/libstix2/objects"
	"github.com/freetaxii/libstix2/resources"
)

// ----------------------------------------------------------------------
//
// Collection Data Table Private Functions
// Table property names and SQL statements
//
// ----------------------------------------------------------------------

/*
collectionDataProperties - This function will return the properties that make up
the collection content table

date_added    = The date that this object was added to the collection
collection_id = The collection ID that this object is tied to
stix_id       = The STIX ID for the object that is being mapped to a collection.
  We do not use the object_id here or the row_id as that would point to a
  specific version and we need to be able to find all versions of an object.
  and if we used row_id for example, it would require two queries, the first
  to get the SITX ID and then the second to get all objects with that STIX ID.
*/
func collectionDataProperties() string {
	return `
	"row_id" INTEGER PRIMARY KEY,
 	"date_added" TEXT NOT NULL,
 	"collection_id" INTEGER NOT NULL,
 	"stix_id" TEXT NOT NULL
 	`
}

// ----------------------------------------------------------------------
//
// Collection Data Table Private Functions and Methods
// getCollectionSize
//
// ----------------------------------------------------------------------

/*
sqlGetCollectionSize - This function will return an SQL statement that will
get the size of a collection from the t_collection_data table.
*/
func sqlGetCollectionSize() (string, error) {
	tblColData := DB_TABLE_TAXII_COLLECTION_DATA

	/*
		SELECT
			count(t_collection_data.collection_id)
		FROM
			t_collection_data
		WHERE
			t_collection_data.collection_id = ?
	*/

	var s bytes.Buffer
	s.WriteString("SELECT ")
	s.WriteString("count(")
	s.WriteString(tblColData)
	s.WriteString(".collection_id) ")
	s.WriteString("FROM ")
	s.WriteString(tblColData)
	s.WriteString(" WHERE ")
	s.WriteString(tblColData)
	s.WriteString(".collection_id = ?")

	return s.String(), nil
}

/*
getCollectionSize - This method will return the size of a given collection
*/
func (ds *Datastore) getCollectionSize(collectionID string) (int, error) {
	ds.Logger.Traceln("TRACE getCollectionSize(): Start")
	var index int
	sqlStmt, _ := sqlGetCollectionSize()

	collectionDatastoreID := ds.Cache.Collections[collectionID].DatastoreID
	err := ds.DB.QueryRow(sqlStmt, collectionDatastoreID).Scan(&index)
	if err != nil {
		if err == sql.ErrNoRows {
			return 0, errors.New("no collection data found")
		}
		return 0, fmt.Errorf("getCollectionSize database execution error: ", err)
	}

	ds.Logger.Debugln("DEBUG: Collection ID", collectionID, "has a size of", index)

	return index, nil
}

// ----------------------------------------------------------------------
//
// Collection Data Table Private Functions and Methods
// addObjectToCollection
//
// ----------------------------------------------------------------------

/*
sqlAddObjectToColleciton - This function will return an SQL statement that will
add an object to a collection by adding an entry in the taxii_collection_data
table. In this table we use the STIX ID not the Object ID because we need to
make sure we include all versions of an object. So we need to store just the
STIX ID
*/
func sqlAddObjectToCollection() (string, error) {
	tblColData := DB_TABLE_TAXII_COLLECTION_DATA

	/*
		INSERT INTO
			t_collection_data (
				"date_added",
				"collection_id",
				"stix_id"
			)
			values (?, ?, ?)
	*/
	var s bytes.Buffer
	s.WriteString("INSERT INTO ")
	s.WriteString(tblColData)
	s.WriteString(" (")
	s.WriteString("date_added, ")
	s.WriteString("collection_id, ")
	s.WriteString("stix_id) ")
	s.WriteString("values (?, ?, ?) ")

	return s.String(), nil
}

/*
addObjectToColleciton - This method will add an object to a collection by adding
an entry in the taxii_collection_data table. In this table we use the STIX ID
not the Object ID because we need to make sure we include all versions of an
object. So we need to store just the STIX ID.
*/
func (ds *Datastore) addObjectToCollection(obj *resources.CollectionRecord) error {
	ds.Logger.Traceln("TRACE addObjectToCollection(): Start")
	dateAdded := time.Now().UTC().Format(defs.TIME_RFC_3339_MICRO)

	// We are storing the Collection DatastoreID which is an integer instead
	// of the long collection ID string (UUID). So lets get the DatastoreID from
	// the cache.
	collectionDatastoreID := ds.Cache.Collections[obj.CollectionID].DatastoreID

	// TODO before we add an object to the database, we need to make sure the
	// object is not already in the table. Another option would be to add them
	// to a secondary table and then have a second process go through and merge
	// them.  This way the end client would not be held up by the transaction.
	sqlStmt, _ := sqlAddObjectToCollection()
	ds.Logger.Debugln("DEBUG: Collection ID", collectionDatastoreID)
	ds.Logger.Debugln("DEBUG: Object ID", obj.STIXID)
	_, err := ds.DB.Exec(sqlStmt, dateAdded, collectionDatastoreID, obj.STIXID)

	if err != nil {
		ds.Logger.Levelln("Function", "FUNC: END addObjectToCollection() with error")
		return fmt.Errorf("database execution error inserting collection data: ", err)
	}

	// If the operation was successful, lets increment the collection cache size
	ds.Cache.Collections[obj.CollectionID].Size++
	ds.Logger.Debugln("DEBUG: Collection ID", obj.CollectionID, "now has a size of", ds.Cache.Collections[obj.CollectionID].Size)
	ds.Logger.Levelln("Function", "FUNC: END addObjectToCollection()")
	return nil
}

// ----------------------------------------------------------------------
//
// Collection Data Table Private Functions and Methods
// getBundle
//
// ----------------------------------------------------------------------

/*
getBundle - This method will return a STIX bundle based on the query provided.
*/
func (ds *Datastore) getBundle(query resources.CollectionQuery) (*resources.CollectionQueryResult, error) {
	ds.Logger.Traceln("TRACE getBundle(): Start")

	// Lets first make sure the collection exists in the cache
	if _, found := ds.Cache.Collections[query.CollectionID]; !found {
		return nil, fmt.Errorf("the following collection id was not found in the cache", query.CollectionID)
	}

	stixBundle := objects.NewBundle()

	// First get a list of all of the objects that are in the collection that
	// meet the query requirements. This is done with the manifest records.
	resultData, err := ds.getManifestData(query)
	if err != nil {
		return nil, err
	}

	// Loop through all of the STIX IDs in the list and get the actual object
	for _, v := range resultData.ManifestData.Objects {
		obj, err := ds.GetObject(v.ID, v.Version)

		if err != nil {
			return nil, err
		}
		stixBundle.AddObject(obj)
	}
	resultData.BundleData = *stixBundle
	return resultData, nil
}

// ----------------------------------------------------------------------
//
// Collection Data Table Private Functions and Methods
// getManifestData
//
// ----------------------------------------------------------------------

/*
sqlGetManifestData - This function will return an SQL statement that will
return a list of objects from a given collection and all of the information
needed to create the manifest resource. It will use the query struct to
determine the requirements and parameters for the where clause of the SQL
statement. A byte array is used instead of sting concatenation as it is the most
efficient way to do string concatenation in Go.
*/
func sqlGetManifestData(query resources.CollectionQuery) (string, error) {
	tblColData := DB_TABLE_TAXII_COLLECTION_DATA
	tblBaseObj := DB_TABLE_STIX_BASE_OBJECT

	// If an error is found, that means a query parameter was passed incorrectly
	// and we should return an error versus just skipping the option.
	whereQuery, err := sqlCollectionDataQueryOptions(query)
	if err != nil {
		return "", err
	}

	// The only types of errors here are if the client value is invalid or is
	// bigger than the server max value. Either way, we should just log. In both
	// cases the server max value is returned.
	limitQuery, err1 := sqlQueryLimit(query)
	if err1 != nil {
		//ds.Logger.Debugln("DEBUG: SQL limit error", err1)
	}

	/*
		SELECT
			t_collection_data.stix_id,
			s.base_object.date_added,
			s_base_object.modified,
			s_base_object.spec_version
		FROM
			t_collection_data
		JOIN
			s_base_object ON
			t_collection_data.stix_id = s_base_object.id
		WHERE
			t_collection_data.collection_id = ?
		LIMIT 5
	*/
	var s bytes.Buffer
	s.WriteString("SELECT ")
	s.WriteString(tblColData)
	s.WriteString(".stix_id, ")
	s.WriteString(tblBaseObj)
	s.WriteString(".date_added, ")
	s.WriteString(tblBaseObj)
	s.WriteString(".modified, ")
	s.WriteString(tblBaseObj)
	s.WriteString(".spec_version ")

	s.WriteString("FROM ")
	s.WriteString(tblColData)

	s.WriteString(" JOIN ")
	s.WriteString(tblBaseObj)
	s.WriteString(" ON ")
	s.WriteString(tblColData)
	s.WriteString(".stix_id = ")
	s.WriteString(tblBaseObj)
	s.WriteString(".id ")

	s.WriteString("WHERE ")
	s.WriteString(whereQuery)

	if limitQuery != 0 {
		s.WriteString(" LIMIT ")
		i := strconv.Itoa(limitQuery)
		s.WriteString(i)
	}

	return s.String(), nil
}

/*
getManifestData - This method will return manifest data based on the query provided.
*/
func (ds *Datastore) getManifestData(query resources.CollectionQuery) (*resources.CollectionQueryResult, error) {
	ds.Logger.Traceln("TRACE getManifestData(): Start")

	// Lets first make sure the collection exists in the cache
	if _, found := ds.Cache.Collections[query.CollectionID]; !found {
		return nil, fmt.Errorf("the following collection id was not found in the cache", query.CollectionID)
	}
	query.CollectionDatastoreID = ds.Cache.Collections[query.CollectionID].DatastoreID

	var resultData resources.CollectionQueryResult
	manifest := resources.NewManifest()

	// If an error is found, that means a query parameter was passed incorrectly
	// and we should return an error versus just skipping the option.
	sqlStmt, err := sqlGetManifestData(query)
	if err != nil {
		return nil, err
	}

	// Query database for all the collection entries
	rows, err := ds.DB.Query(sqlStmt)

	if err != nil {
		return nil, fmt.Errorf("database execution error getting collection data: ", err)
	}
	defer rows.Close()

	for rows.Next() {
		var stixid, dateAdded, modified, specVersion string
		if err := rows.Scan(&stixid, &dateAdded, &modified, &specVersion); err != nil {
			rows.Close()
			return nil, fmt.Errorf("database scan error getting collection data: ", err)
		}

		switch specVersion {
		case "2.0":
			specVersion = defs.MEDIA_TYPE_STIX20
		case "2.1":
			specVersion = defs.MEDIA_TYPE_STIX21
		case "2.2":
			specVersion = defs.MEDIA_TYPE_STIX22
		case "2.3":
			specVersion = defs.MEDIA_TYPE_STIX23
		case "2.4":
			specVersion = defs.MEDIA_TYPE_STIX24
		default:
			specVersion = defs.MEDIA_TYPE_STIX
		}
		manifest.CreateManifestEntry(stixid, dateAdded, modified, specVersion)
	}

	// Errors can cause the rows.Next() to exit prematurely, if this happens lets
	// check for the error and handle it.
	if err := rows.Err(); err != nil {
		rows.Close()
		return nil, fmt.Errorf("database rows error getting manifest data: ", err)
	}

	if len(manifest.Objects) == 0 {
		return nil, fmt.Errorf("no records returned getting manifest data")
	}

	resultData.Size = ds.Cache.Collections[query.CollectionID].Size

	ds.Logger.Traceln("TRACE getManifestData(): Query Collection ID", query.CollectionID)
	ds.Logger.Traceln("TRACE getManifestData(): Cache ID", ds.Cache.Collections[query.CollectionID].ID, "Cache Datastore ID", ds.Cache.Collections[query.CollectionID].DatastoreID, "Size in Cache", ds.Cache.Collections[query.CollectionID].Size)

	resultData.ManifestData.Objects = manifest.Objects

	// ----------------------------------------------------------------------
	// This is the old pagination code
	// first, last, errRange = ds.processRangeValues(query.RangeBegin, query.RangeEnd, query.ServerRecordLimit, resultData.Size)

	// if errRange != nil {
	// 	return nil, errRange
	// }

	// Get a new slice based on the range of records
	// resultData.ManifestData.Objects = manifest.Objects[first:last]
	// ----------------------------------------------------------------------

	resultData.DateAddedFirst = resultData.ManifestData.Objects[0].DateAdded
	resultData.DateAddedLast = resultData.ManifestData.Objects[len(resultData.ManifestData.Objects)-1].DateAdded

	return &resultData, nil
}

// ----------------------------------------------------------------------
//
// HTTP Range Values for Collection Data Queries
//
// ----------------------------------------------------------------------

/*
processRangeValues - This method will take in the various range parameters and size
of the dataset and will return the correct first and last index values to be used.
*/
func (ds *Datastore) processRangeValues(first, last, max, size int) (int, int, error) {
	ds.Logger.Traceln("TRACE processRangeValues(): Start")

	if first < 0 {
		return 0, 0, errors.New("the starting value can not be negative")
	}

	if first > last {
		return 0, 0, errors.New("the starting range value is larger than the ending range value")
	}

	if first >= size {
		return 0, 0, errors.New("the starting range value is out of scope")
	}

	// If no range is requested and the server is not forcing it, do nothing.
	if last == 0 && first == 0 && max != 0 {
		last = first + max
	} else {
		// We need to be inclusive of the last value that was provided
		last++
	}

	// If the last record requested is bigger than the total size of the data
	// set the last size to be the size of the data
	if last > size {
		last = size
	}

	// If the request is for more records than the max size will allow, then
	// compute where the new last record should be, but only if the server is
	// forcing a max size.
	if max != 0 && (last-first) > max {
		last = first + max
	}

	return first, last, nil
}

// ----------------------------------------------------------------------
//
// LIMIT statements for Collection Data Queries
//
// ----------------------------------------------------------------------

/*
sqlQueryLimit - This function will take in a query struct and
build an SQL LIMIT statement based on the values provided in the query object.
*/
func sqlQueryLimit(query resources.CollectionQuery) (int, error) {
	srv := 0
	client := 0
	var err error

	if query.ServerRecordLimit > 0 {
		srv = query.ServerRecordLimit
	}

	if query.Limit != nil {
		client, err = strconv.Atoi(query.Limit[0])
	}
	if err != nil {
		return srv, fmt.Errorf("client limit value is not valid: ", err)
	}

	if client > srv {
		return srv, fmt.Errorf("client limit value is greater than the server limit, using server limit")
	} else if client == 0 {
		return srv, nil
	} else if client == srv {
		return srv, nil
	} else if client < srv {
		return client, nil
	}

	return 0, nil
}

// ----------------------------------------------------------------------
//
// WHERE statements for Collection Data Queries
//
// ----------------------------------------------------------------------

/*
sqlCollectionDataQueryOptions - This function will take in a query struct and
build an SQL where statement based on all of the provided query parameters.
*/
func sqlCollectionDataQueryOptions(query resources.CollectionQuery) (string, error) {
	var wherestmt bytes.Buffer
	var err error

	// ----------------------------------------------------------------------
	// Lets first add the collection ID to the where clause.
	// ----------------------------------------------------------------------
	if err = sqlCollectionDataWhereCollectionID(query.CollectionDatastoreID, &wherestmt); err != nil {
		return "", err
	}

	// ----------------------------------------------------------------------
	// Check to see if an added after query was supplied. There can only be one
	// added after option, it does not make sense to have multiple.
	// ----------------------------------------------------------------------
	if err = sqlCollectionDataWhereAddedAfter(query.AddedAfter, &wherestmt); err != nil {
		return "", err
	}

	// ----------------------------------------------------------------------
	// Check to see if one or more STIX IDs was supplied.
	// If there is more than one option given we need to enclose the options in
	// parentheses as the comma represents an OR operator.
	// ----------------------------------------------------------------------
	if err = sqlCollectionDataWhereSTIXID(query.STIXID, &wherestmt); err != nil {
		return "", err
	}

	// ----------------------------------------------------------------------
	// Check to see if one or more STIX types, to query on, was supplied.
	// If there is more than one option given we need to enclose the options in
	// parentheses as the comma represents an OR operator.
	// ----------------------------------------------------------------------
	if err = sqlCollectionDataWhereSTIXType(query.STIXType, &wherestmt); err != nil {
		return "", err
	}

	// ----------------------------------------------------------------------
	// Check to see if one or more STIX versions to query on was supplied.
	// If there is more than one option given, we need to enclose the options in
	// parentheses as the comma represents an OR operator.
	// ----------------------------------------------------------------------
	if err = sqlCollectionDataWhereSTIXVersion(query.STIXVersion, &wherestmt); err != nil {
		return "", err
	}

	return wherestmt.String(), nil
}

/*
sqlCollectionDataWhereCollectionID - This function will build the correct WHERE
statement for a provided collection ID value and is called from
func sqlCollectionDataQueryOptions(query resources.CollectionQueryType) (string, error)
*/
func sqlCollectionDataWhereCollectionID(id int, b *bytes.Buffer) error {
	tblColData := DB_TABLE_TAXII_COLLECTION_DATA

	/*
		This sql where statement should look like:
		t_collection_data.collection_id = "some collection id"
	*/
	if id != 0 {
		b.WriteString(tblColData)
		b.WriteString(`.collection_id = `)
		b.WriteString(strconv.Itoa(id))
	} else {
		return errors.New("no collection ID was provided")
	}
	return nil
}

/*
sqlCollectionDataWhereAddedAfter - This function will build the correct WHERE
statement for a provided added after value and is called from
func sqlCollectionDataQueryOptions(query resources.CollectionQueryType) (string, error)

This method only supports a single added after value, since more than one does
not make sense.
*/
func sqlCollectionDataWhereAddedAfter(date []string, b *bytes.Buffer) error {
	tblBaseObj := DB_TABLE_STIX_BASE_OBJECT

	/*
		This sql where statement should look like:
		t_collection_data.collection_id = "aa" AND
		t_collection_data.date_added > "2017"
	*/
	if date != nil {
		// We are only allowing a single added after value, since having more does
		// not make sense.
		if timestamp.Valid(date[0]) {
			b.WriteString(" AND ")
			b.WriteString(tblBaseObj)
			b.WriteString(`.date_added > "`)
			b.WriteString(date[0])
			b.WriteString(`"`)
		} else {
			return errors.New("the provided timestamp for added after is invalid")
		}
	}
	return nil
}

/*
sqlCollectionDataWhereSTIXID - This function will build the correct WHERE
statement when one or more STIX IDs is provided and is called from
func sqlCollectionDataQueryOptions(query resources.CollectionQueryType) (string, error)
*/
func sqlCollectionDataWhereSTIXID(id []string, b *bytes.Buffer) error {
	tblColData := DB_TABLE_TAXII_COLLECTION_DATA

	/*
		This sql where statement should look like one of these two:
		t_collection_data.collection_id = "aa" AND
		t_collection_data.stix_id = "indicator--37abef16-7616-439c-86be-23712030c4b7"

		t_collection_data.collection_id = "aa" AND
		(t_collection_data.stix_id = "indicator--37abef16-7616-439c-86be-23712030c4b7" OR
		t_collection_data.stix_id = "attack-pattern--c7c8a099-70a9-487b-a95f-2498d2941104" OR
		t_collection_data.stix_id = "campaign--6f938db5-6648-4ec1-81cb-5b65138c3c66")
	*/
	if id != nil {
		if len(id) == 1 {
			if stixid.ValidSTIXID(id[0]) {
				b.WriteString(" AND ")
				b.WriteString(tblColData)
				b.WriteString(`.stix_id = "`)
				b.WriteString(id[0])
				b.WriteString(`"`)
			} else {
				return errors.New("invalid SQL where statement, the provided object id is invalid")
			}
		} else if len(id) > 1 {
			b.WriteString(" AND (")
			addOR := false
			for _, v := range id {

				// Lets only add the OR after the first object id and not after the last object id
				if addOR == true {
					b.WriteString(" OR ")
					addOR = false
				}
				// Lets make sure the value that was passed in is actually a valid id

				if stixid.ValidSTIXID(v) {
					b.WriteString(tblColData)
					b.WriteString(`.stix_id = "`)
					b.WriteString(v)
					b.WriteString(`"`)
					addOR = true
				} else {
					return errors.New("invalid SQL where statement, the provided object id is invalid")
				}
			}
			b.WriteString(")")
		}
	}
	return nil
}

/*
sqlCollectionDataWhereSTIXType - This function will build the correct WHERE
statement when one or more STIX types is provided and is called from
func sqlCollectionDataQueryOptions(query resources.CollectionQueryType) (string, error)
*/
func sqlCollectionDataWhereSTIXType(t []string, b *bytes.Buffer) error {
	tblColData := DB_TABLE_TAXII_COLLECTION_DATA

	/*
		This sql where statement should look like one of these two:
		t_collection_data.collection_id = "aa" AND
		t_collection_data.stix_id LIKE "indicator%"

		t_collection_data.collection_id = "aa" AND
		(t_collection_data.stix_id LIKE "indicator%" OR
		t_collection_data.stix_id LIKE "attack-pattern%" OR
		t_collection_data.stix_id LIKE "campaign%")
	*/
	if t != nil {
		if len(t) == 1 {
			if stixid.ValidSTIXObjectType(t[0]) {
				b.WriteString(" AND ")
				b.WriteString(tblColData)
				b.WriteString(`.stix_id LIKE "`)
				b.WriteString(t[0])
				b.WriteString(`%"`)
			} else {
				return errors.New("the provided object type is invalid")
			}
		} else if len(t) > 1 {
			b.WriteString(" AND (")
			addOR := false
			for _, v := range t {

				// Lets only add the OR after the first object and not after the last object
				if addOR == true {
					b.WriteString(" OR ")
					addOR = false
				}
				// Lets make sure the value that was passed in is actually a valid object
				if stixid.ValidSTIXObjectType(v) {
					b.WriteString(tblColData)
					b.WriteString(`.stix_id LIKE "`)
					b.WriteString(v)
					b.WriteString(`%"`)
					addOR = true
				} else {
					return errors.New("the provided object type is invalid")
				}
			}
			b.WriteString(`)`)
		}
	}
	return nil
}

/*
sqlCollectionDataWhereSTIXVersion - This function will build the correct WHERE
statement when one or more STIX versions is provided and is called from
func sqlCollectionDataQueryOptions(query resources.CollectionQueryType) (string, error).

It will return an error if multiple "all", "first", or "last" values is provided.
*/
func sqlCollectionDataWhereSTIXVersion(vers []string, b *bytes.Buffer) error {
	tblColData := DB_TABLE_TAXII_COLLECTION_DATA
	tblBaseObj := DB_TABLE_STIX_BASE_OBJECT

	// If no version parameter was supplied, then set "last" as the default
	if vers == nil {
		vers = append(vers, "last")
	}

	// Lets check the multiple version use case and see if the options are valid
	if len(vers) > 1 {
		first := 0
		last := 0
		all := 0

		for _, v := range vers {
			if v == "last" {
				last++
			} else if v == "first" {
				first++
			} else if v == "all" {
				all++
			}
		}

		if last > 1 {
			return errors.New("can not use the 'last' key word multiple time in the version selector")
		}
		if first > 1 {
			return errors.New("can not use the 'first' key word multiple time in the version selector")
		}
		if all > 0 {
			return errors.New("can not use the 'all' key word with a multiple version selector")
		}
	}

	/*
		This sql where statement should look like one of the following:
		t_collection_data.collection_id = "aa" AND
		s_base_object.modified = "2017-12-05T02:43:19.783Z"

		t_collection_data.collection_id = "aa" AND
		s_base_object.modified = (select max(modified) from s_base_object where t_collection_data.stix_id = s_base_object.id)

		t_collection_data.collection_id = "aa" AND
		s_base_object.modified = (select min(modified) from s_base_object where t_collection_data.stix_id = s_base_object.id)

		t_collection_data.collection_id = "aa" AND
		(s_base_object.modified = (select min(modified) from s_base_object where t_collection_data.stix_id = s_base_object.id)  OR
		s_base_object.modified = (select max(modified) from s_base_object where t_collection_data.stix_id = s_base_object.id))

		t_collection_data.collection_id = "aa" AND
		(s_base_object.modified = "2017-12-05T02:43:19.783Z" OR
		s_base_object.modified = "2017-12-05T02:43:23.822Z" OR
		s_base_object.modified = "2017-12-05T02:43:24.835Z")
	*/

	if len(vers) == 1 {
		if vers[0] == "last" {
			// s_base_object.modified = (select max(modified) from s_base_object where t_collection_data.stix_id = s_base_object.id)
			b.WriteString(" AND ")
			b.WriteString(tblBaseObj)
			b.WriteString(`.modified = (select max(modified) from `)
			b.WriteString(tblBaseObj)
			b.WriteString(` where `)
			b.WriteString(tblColData)
			b.WriteString(`.stix_id = `)
			b.WriteString(tblBaseObj)
			b.WriteString(`.id)`)

		} else if vers[0] == "first" {
			// s_base_object.modified = (select min(modified) from s_base_object where t_collection_data.stix_id = s_base_object.id)
			b.WriteString(" AND ")
			b.WriteString(tblBaseObj)
			b.WriteString(`.modified = (select min(modified) from `)
			b.WriteString(tblBaseObj)
			b.WriteString(` where `)
			b.WriteString(tblColData)
			b.WriteString(`.stix_id = `)
			b.WriteString(tblBaseObj)
			b.WriteString(`.id)`)

		} else if vers[0] == "all" {
			// Do nothing, since the default is to return all versions.
		} else {
			if timestamp.Valid(vers[0]) {
				b.WriteString(" AND ")
				b.WriteString(tblBaseObj)
				b.WriteString(`.modified = "`)
				b.WriteString(vers[0])
				b.WriteString(`"`)

			} else {
				return errors.New("the provided timestamp for the version is invalid")
			}
		}
	} else if len(vers) > 1 {
		b.WriteString(" AND (")
		for i, v := range vers {
			// Lets only add he OR after the first object and not after the
			// last object. Since skipOr starts as true, this takes care of
			// the first run case where i == 0

			if i > 0 {
				b.WriteString(" OR ")
			}

			if v == "last" {
				b.WriteString(tblBaseObj)
				b.WriteString(`.modified = (select max(modified) from `)
				b.WriteString(tblBaseObj)
				b.WriteString(` where `)
				b.WriteString(tblColData)
				b.WriteString(`.stix_id = `)
				b.WriteString(tblBaseObj)
				b.WriteString(`.id)`)

			} else if v == "first" {
				b.WriteString(tblBaseObj)
				b.WriteString(`.modified = (select min(modified) from `)
				b.WriteString(tblBaseObj)
				b.WriteString(` where `)
				b.WriteString(tblColData)
				b.WriteString(`.stix_id = `)
				b.WriteString(tblBaseObj)
				b.WriteString(`.id)`)

			} else {
				if timestamp.Valid(v) {
					b.WriteString(tblBaseObj)
					b.WriteString(`.modified = "`)
					b.WriteString(v)
					b.WriteString(`"`)
				} else {
					return errors.New("the provided timestamp for the version is invalid")
				}
			}
		}
		b.WriteString(`)`)
	}

	return nil
}
