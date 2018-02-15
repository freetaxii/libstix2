// Copyright 2017 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

package datastore

import (
	"github.com/freetaxii/libstix2/objects"
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
	GetBundle(query CollectionQueryType) (*CollectionQueryResultType, error)
	GetManifestData(query CollectionQueryType) (*CollectionQueryResultType, error)
}

/*
CacheType - This struct will hold a cache of various database elements
that will be loaded during initialization and updated along the way.
*/
type CacheType struct {
	BaseObjectIDIndex int
	Collections       map[string]*resources.CollectionType
}

/*
CollectionQueryType - This struct will hold all of the variables that a user can
query a collection on.
*/
type CollectionQueryType struct {
	CollectionID          string
	CollectionDatastoreID int
	STIXID                []string // Passed in from the URL
	STIXType              []string // Passed in from the URL
	STIXVersion           []string // Passed in from the URL
	AddedAfter            []string // Passed in from the URL
	AddedBefore           []string // Passed in from the URL
	Limit                 []string // Passed in from the URL
	RangeBegin            int      // Passed in from Range Headers
	RangeEnd              int      // Passed in from Range Headers
	ClientRecordLimit     int      // Passed in from Range Headers
	ServerRecordLimit     int      // Server defined value in the configuration file
}

/*
CollectionQueryResultType - This struct contains the various bits of meta data
that are returned from a query against a collection on a TAXII server. This is
done so that the method signatures do not need to change as time goes on and we
add more meta data that needs to be returned. It is important to note that a
collection may have more entries than the server or client wants to transmit. So
it is important to keep track of which records are actually being delivered to
the client.

Size           = The total size of the dataset returned from the database query.
RangeBegin     = The range value of the first record being sent to the client.
RangeEnd       = The range value of the last record being sent to the client.
DateAddedFirst = The added date of the first record being sent to the client.
DateAddedLast  = The added date of the last record being sent to the client.
BundleData     = The STIX bundle that contains the requested data from the collection.
ManifestData   = The TAXII manifest resource that contains the requested data from the collection.
*/
type CollectionQueryResultType struct {
	Size           int
	RangeBegin     int
	RangeEnd       int
	DateAddedFirst string
	DateAddedLast  string
	BundleData     objects.BundleType
	ManifestData   resources.ManifestType
}
