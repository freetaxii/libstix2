// Copyright 2015-2019 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

package sqlite3

import (
	"bytes"
	"database/sql"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/freetaxii/libstix2/defs"
	"github.com/freetaxii/libstix2/objects"
	"github.com/freetaxii/libstix2/objects/taxii/collections"
	"github.com/freetaxii/libstix2/objects/taxii/envelope"
	"github.com/freetaxii/libstix2/objects/taxii/versions"
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

	We do not use the datastore_id here or the row_id as that would point to a
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
getCollectionSize - This method takes in a collection datastore ID and will
return the size of a given collection.
*/
func (ds *Store) getCollectionSize(collectionDatastoreID int) (int, error) {
	ds.Logger.Levelln("Function", "FUNC: getCollectionSize start")
	ds.Logger.Debugln("DEBUG: Getting collection size for collection ID", collectionDatastoreID)
	var size int

	// Create SQL Statement
	/*
		SELECT
			count(collection_id)
		FROM
			t_collection_data
		WHERE
			collection_id = ?
	*/
	tblColData := DB_TABLE_TAXII_COLLECTION_DATA
	var sqlstmt bytes.Buffer
	sqlstmt.WriteString("SELECT count(collection_id) FROM ")
	sqlstmt.WriteString(tblColData)
	sqlstmt.WriteString(" WHERE collection_id = ?")
	stmt := sqlstmt.String()

	// Make SQL Call
	err := ds.DB.QueryRow(stmt, collectionDatastoreID).Scan(&size)
	if err != nil {
		if err == sql.ErrNoRows {
			ds.Logger.Levelln("Function", "FUNC: getCollectionSize exited with an error,", err)
			return 0, errors.New("no collection data found")
		}
		ds.Logger.Levelln("Function", "FUNC: getCollectionSize exited with an error,", err)
		return 0, fmt.Errorf("getCollectionSize database execution error: ", err)
	}

	ds.Logger.Debugln("DEBUG: Collection ID", collectionDatastoreID, "has a size of", size)

	ds.Logger.Levelln("Function", "FUNC: getCollectionSize end")
	return size, nil
}

// ----------------------------------------------------------------------
//
// Collection Data Table Private Functions and Methods
// addToCollection
//
// ----------------------------------------------------------------------

/*
addToCollection - This method will add an entry to a collection by adding
an entry in the taxii_collection_data table. In this table we use the STIX ID
not the Object ID because we need to make sure we include all versions of an
object. So we need to store just the STIX ID.
*/
func (ds *Store) addToCollection(collectionUUID, stixid string) error {
	ds.Logger.Levelln("Function", "FUNC: addToCollection start")

	// Lets first make sure the collection exists in the cache
	if found := ds.doesCollectionExistInTheCache(collectionUUID); !found {
		ds.Logger.Levelln("Function", "FUNC: addToCollection exited with an error")
		return fmt.Errorf("the following collection id was not found in the cache", collectionUUID)
	}

	// We are storing the Collection DatastoreID which is an integer instead
	// of the long collection ID string (UUID). So lets get the DatastoreID from
	// the cache.
	collectionDatastoreID := ds.Cache.Collections[collectionUUID].DatastoreID
	ds.Logger.Debugln("DEBUG: Collection Datastore ID", collectionDatastoreID)
	ds.Logger.Debugln("DEBUG: Object ID", stixid)

	// Create SQL Statement
	/*
		INSERT INTO
			t_collection_data (
				"date_added",
				"collection_id",
				"stix_id"
			)
			values (?, ?, ?)
	*/
	tblColData := DB_TABLE_TAXII_COLLECTION_DATA
	var sqlstmt bytes.Buffer
	sqlstmt.WriteString("INSERT INTO ")
	sqlstmt.WriteString(tblColData)
	sqlstmt.WriteString(" (date_added, collection_id, stix_id) values (?, ?, ?)")
	stmt := sqlstmt.String()

	// TODO before we add an object to the database, we need to make sure the
	// object is not already in the table. Another option would be to add them
	// to a secondary table and then have a second process go through and merge
	// them. This way the end client would not be held up by the transaction.

	dateAdded := time.Now().UTC().Format(defs.TimeRFC3339Micro)

	// Make SQL Call
	_, err := ds.DB.Exec(stmt, dateAdded, collectionDatastoreID, stixid)
	if err != nil {
		ds.Logger.Levelln("Function", "FUNC: addToCollection exited with an error,", err)
		return fmt.Errorf("database execution error inserting collection data: ", err)
	}

	// If the operation was successful, lets increment the collection cache size
	ds.Cache.Collections[collectionUUID].Size++
	ds.Logger.Debugln("DEBUG: Collection ID", collectionUUID, "now has a size of", ds.Cache.Collections[collectionUUID].Size)
	ds.Logger.Levelln("Function", "FUNC: addToCollection end")
	return nil
}

// ----------------------------------------------------------------------
//
// Collection Data Table Private Functions and Methods
// getObjects
// getVersions
//
// ----------------------------------------------------------------------

/*
getObjects - This method will return a TAXII Envelope resource based on the query provided.
*/
func (ds *Store) getObjects(query collections.CollectionQuery) (*collections.CollectionQueryResult, error) {
	ds.Logger.Levelln("Function", "FUNC: getObjects start")

	// Lets first make sure the collection exists in the cache
	if found := ds.doesCollectionExistInTheCache(query.CollectionUUID); !found {
		ds.Logger.Levelln("Function", "FUNC: getObjects exited with an error")
		return nil, fmt.Errorf("the following collection id was not found in the cache", query.CollectionUUID)
	}

	taxiiEnvelope := envelope.New()

	// First get a list of all of the objects that are in the collection that
	// meet the query requirements. This is done with the manifest records.
	resultData, err := ds.getManifestData(query)
	if err != nil {
		ds.Logger.Levelln("Function", "FUNC: getObjects exited with an error,", err)
		return nil, err
	}

	if resultData.ManifestData.More == true {
		taxiiEnvelope.SetMore()
	}

	// Loop through all of the STIX IDs in the list and get the actual object
	for _, v := range resultData.ManifestData.Objects {
		var obj interface{}
		var err error

		// ------------------------------------------------------------
		// Test STIX ID to see if it is valid
		// ------------------------------------------------------------
		idparts := strings.Split(v.ID, "--")

		// Is the UUIDv4 portion of the ID valid?
		if ds.Strict.IDs == true {
			if !objects.IsUUIDValid(idparts[1]) {
				ds.Logger.Debugln("DEBUG: Get STIX object error, invalid STIX UUID", idparts[1])
				continue
			}
		}

		// TODO: FIX
		// // Is the STIX type part of the ID valid?
		// if ds.Strict.Types == true {
		// 	if !stixid.ValidSTIXObjectType(idparts[0]) {
		// 		ds.Logger.Debugln("DEBUG: Get STIX object error, invalid STIX type", idparts[0])
		// 		continue
		// 	}
		// }

		// ------------------------------------------------------------
		// Get object by type
		// ------------------------------------------------------------

		switch idparts[0] {
		case "indicator":
			obj, err = ds.getIndicator(v.ID, v.Version)
			if err != nil {
				ds.Logger.Debugln("DEBUG: Get object error,", err)
				continue
			}
		default:
			ds.Logger.Debugln("DEBUG: Get object error, the following STIX type is not currently supported: ", idparts[0])
			continue
		}

		taxiiEnvelope.AddObject(obj)
	}

	resultData.ObjectData = *taxiiEnvelope
	ds.Logger.Levelln("Function", "FUNC: getObjects end")
	return resultData, nil
}

/*
getVersions - This method will return a TAXII Versions resource based on the query provided.
*/
func (ds *Store) getVersions(query collections.CollectionQuery) (*collections.CollectionQueryResult, error) {
	ds.Logger.Levelln("Function", "FUNC: getVersions start")

	// Lets first make sure the collection exists in the cache
	if found := ds.doesCollectionExistInTheCache(query.CollectionUUID); !found {
		ds.Logger.Levelln("Function", "FUNC: getVersions exited with an error")
		return nil, fmt.Errorf("the following collection id was not found in the cache", query.CollectionUUID)
	}

	taxiiVersions := versions.New()

	// First get a list of all of the objects that are in the collection that
	// meet the query requirements. This is done with the manifest records.
	resultData, err := ds.getManifestData(query)
	if err != nil {
		ds.Logger.Levelln("Function", "FUNC: getVersions exited with an error,", err)
		return nil, err
	}

	if resultData.ManifestData.More == true {
		taxiiVersions.SetMore()
	}

	// Loop through all of the STIX IDs in the list and get the actual object
	for _, v := range resultData.ManifestData.Objects {
		taxiiVersions.AddVersion(v.Version)
	}

	resultData.VersionsData = *taxiiVersions
	ds.Logger.Levelln("Function", "FUNC: getVersions end")
	return resultData, nil
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
func (ds *Store) processRangeValues(first, last, max, size int) (int, int, error) {
	ds.Logger.Levelln("Function", "FUNC: processRangeValues start")

	if first < 0 {
		ds.Logger.Levelln("Function", "FUNC: processRangeValues exited with an error")
		return 0, 0, errors.New("the starting value can not be negative")
	}

	if first > last {
		ds.Logger.Levelln("Function", "FUNC: processRangeValues exited with an error")
		return 0, 0, errors.New("the starting range value is larger than the ending range value")
	}

	if first >= size {
		ds.Logger.Levelln("Function", "FUNC: processRangeValues exited with an error")
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

	ds.Logger.Levelln("Function", "FUNC: processRangeValues end")
	return first, last, nil
}
