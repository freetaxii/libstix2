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
Datastorer - This type enables access to the TAXII datastore.
GetSTIXObject returns a STIX object
*/
type Datastorer interface {
	Close() error
	GetEnabledCollections() resources.CollectionsType
	GetSTIXObject(stixid, version string) (interface{}, error)
	GetObjectsFromCollection(query CollectionQueryType) (*CollectionQueryResultType, error)
	GetListOfObjectsFromCollection(query CollectionQueryType) (*CollectionQueryResultType, error)
	GetManifestData(query CollectionQueryType) (*CollectionQueryResultType, error)
}

/*
DatabaseCacheType - This struct will hold a cache of various database elements
that will be loaded during initialization and updated along the way.
*/
type DatabaseCacheType struct {
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
	RangeBegin            int      // Passed in from Range Headers
	RangeEnd              int      // Passed in from Range Headers
	ServerRecordLimit     int      // Server defined value in the configuration file
	ClientRecordLimit     int      // Passed in from Range Headers
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
