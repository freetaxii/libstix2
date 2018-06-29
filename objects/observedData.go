// Copyright 2017 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

package objects

import (
	"github.com/freetaxii/libstix2/common/timestamp"
	"github.com/freetaxii/libstix2/objects/properties"
)

// ----------------------------------------------------------------------
//
// Define Object Type
//
// ----------------------------------------------------------------------

/*
ObservedData - This type implements the STIX 2 Observed Data SDO and defines
all of the properties methods needed to create and work with the STIX Observed Data
SDO. All of the methods not defined local to this type are inherited from
the individual properties.

The following information comes directly from the STIX 2 specification documents.

Observed Data conveys information that was observed on systems and networks
using the Cyber Observable specification defined in parts 3 and 4 of this
specification. For example, Observed Data can capture the observation of an IP
address, a network connection, a file, or a registry key. Observed Data is not
an intelligence assertion, it is simply information: this file was seen, without
any context for what it means.

Observed Data captures both a single observation of a single entity (file,
network connection) as well as the aggregation of multiple observations of an
entity. When the number_observed property is 1 the Observed Data is of a single
entity. When the number_observed property is greater than 1, the observed data
consists of several instances of an entity collected over the time window
specified by the first_observed and last_observed properties. When used to
collect aggregate data, it is likely that some fields in the Cyber Observable
Object (e.g., timestamp fields) will be omitted because they would differ for
each of the individual observations.

Observed Data may be used by itself (without relationships) to convey raw data
collected from network and host-based detection tools. A firewall could emit a
single Observed Data instance containing a single Network Traffic object for
each connection it sees. The firewall could also aggregate data and instead send
out an Observed Data instance every ten minutes with an IP address and an
appropriate number_observed value to indicate the number of times that IP
address was observed in that window.

Observed Data may also be related to other SDOs to represent raw data that is
relevant to those objects. The Sighting object, which captures the sighting of
an Indicator, Malware, or other SDO, uses Observed Data to represent the raw
information that led to the creation of the Sighting (e.g., what was actually
seen that suggested that a particular instance of malware was active).
*/
type ObservedData struct {
	properties.CommonObjectProperties
	FirstObserved  string `json:"first_observed,omitempty"`
	LastObserved   string `json:"last_observed,omitempty"`
	NumberObserved int    `json:"number_observed,omitempty"`
	Objects        string `json:"objects,omitempty"`
}

// ----------------------------------------------------------------------
//
// Initialization Functions
//
// ----------------------------------------------------------------------

/*
NewObservedData - This function will create a new STIX Observed Data object
and return it as a pointer.
*/
func NewObservedData(ver string) *ObservedData {
	var obj ObservedData
	obj.InitObjectProperties("observed-data", ver)
	return &obj
}

// ----------------------------------------------------------------------
//
// Public Methods - ObservedData
//
// ----------------------------------------------------------------------

/*
SetFirstObservedToCurrentTime - This methods sets the first observed time to the
current time
*/
func (o *ObservedData) SetFirstObservedToCurrentTime() error {
	o.FirstObserved = timestamp.GetCurrentTime("micro")
	return nil
}

/*
SetFirstObserved - This method takes in a timestamp in either time.Time or
string format and updates the first observed property.
*/
func (o *ObservedData) SetFirstObserved(t interface{}) error {
	ts, _ := timestamp.ToString(t, "micro")
	o.FirstObserved = ts
	return nil
}

/*
SetLastObservedToCurrentTime - This methods sets the last observed time to the
current time
*/
func (o *ObservedData) SetLastObservedToCurrentTime() error {
	o.LastObserved = timestamp.GetCurrentTime("micro")
	return nil
}

/*
SetLastObserved - This method takes in a timestamp in either time.Time or
string format and updates the last observed property.
*/
func (o *ObservedData) SetLastObserved(t interface{}) error {
	ts, _ := timestamp.ToString(t, "micro")
	o.LastObserved = ts
	return nil
}

/*
SetNumberObserved - This method takes in an integer that represents the
number of objects that were observed and updates the number observed property.
*/
func (o *ObservedData) SetNumberObserved(i int) error {
	o.NumberObserved = i
	return nil
}

/*
SetObjects - This takes in a string value that represents represents a cyber
observable JSON object and updates the objects property.
*/
func (o *ObservedData) SetObjects(s string) error {
	o.Objects = s
	return nil
}
