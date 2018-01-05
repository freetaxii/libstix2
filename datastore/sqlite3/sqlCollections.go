// Copyright 2017 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license
// that can be found in the LICENSE file in the root of the source
// tree.

package sqlite3

import (
	"bytes"
	"github.com/freetaxii/libstix2/datastore"
	//"log"
)

/*
sqlEnabledCollections - This method will take in a query struct and return
an SQL select statement that matches the requirements and parameters given in the
query struct. We are using the byte array as it is the most efficient way to do
string concatenation in Go.
*/
func (ds *Sqlite3DatastoreType) sqlEnabledCollections() (string, error) {
	tblCol := datastore.DB_TABLE_TAXII_COLLECTIONS
	tblColMedia := datastore.DB_TABLE_TAXII_COLLECTION_MEDIA_TYPE
	tblMediaTypes := datastore.DB_TABLE_TAXII_MEDIA_TYPES

	/*
		SELECT
			t_collections.id,
			t_collections.title,
			t_collections.description,
			t_collections.can_read,
			t_collections.can_write
			group_concat(t_media_types.media_type)
		FROM
			t_collections
		JOIN
			t_collection_media_type ON
			t_collections.id = t_collection_media_type.collection_id
		JOIN
			t_media_types ON
			t_collection_media_type.media_type_id = t_media_types.row_id
		WHERE
			t_collections.enabled == '1'
		GROUP BY
			t_collections.id
	*/
	var s bytes.Buffer
	s.WriteString("SELECT \n\t")
	s.WriteString(tblCol)
	s.WriteString(".id, \n\t")
	s.WriteString(tblCol)
	s.WriteString(".title, \n\t")
	s.WriteString(tblCol)
	s.WriteString(".description, \n\t")
	s.WriteString(tblCol)
	s.WriteString(".can_read, \n\t")
	s.WriteString(tblCol)
	s.WriteString(".can_write, \n\t")
	s.WriteString("group_concat(")
	s.WriteString(tblMediaTypes)
	s.WriteString(".media_type) \n")

	s.WriteString("FROM \n\t")
	s.WriteString(tblCol)
	s.WriteString("\n")

	s.WriteString("JOIN \n\t")
	s.WriteString(tblColMedia)
	s.WriteString(" ON \n\t")
	s.WriteString(tblCol)
	s.WriteString(".id = ")
	s.WriteString(tblColMedia)
	s.WriteString(".collection_id \n")

	s.WriteString("JOIN \n\t")
	s.WriteString(tblMediaTypes)
	s.WriteString(" ON \n\t")
	s.WriteString(tblColMedia)
	s.WriteString(".media_type_id = ")
	s.WriteString(tblMediaTypes)
	s.WriteString(".row_id \n")

	s.WriteString("WHERE \n\t")
	s.WriteString(tblCol)
	s.WriteString(".enabled == '1' \n")

	s.WriteString("GROUP BY \n\t")
	s.WriteString(tblCol)
	s.WriteString(".id \n")

	//log.Println("DEBUG: \n", s.String())
	return s.String(), nil
}
