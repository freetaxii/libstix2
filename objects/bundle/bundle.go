// Copyright 2018 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

package bundle

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"

	"github.com/freetaxii/libstix2/objects/baseobject"
	"github.com/freetaxii/libstix2/objects/campaign"
	"github.com/freetaxii/libstix2/objects/indicator"
	"github.com/freetaxii/libstix2/objects/infrastructure"
	"github.com/freetaxii/libstix2/objects/malware"
	"github.com/freetaxii/libstix2/objects/observeddata"
	"github.com/freetaxii/libstix2/objects/relationship"
	"github.com/freetaxii/libstix2/objects/sighting"
)

// ----------------------------------------------------------------------
//
// Define Object Type
//
// ----------------------------------------------------------------------

/*
Bundle - This type implements the STIX 2 Bundle and defines
all of the properties methods needed to create and work with the STIX Bundle.
All of the methods not defined local to this type are inherited from
the individual properties.

The following information comes directly from the STIX 2 specification documents.

A Bundle is a collection of arbitrary STIX Objects and Marking Definitions
grouped together in a single container. A Bundle does not have any semantic
meaning and Objects are not considered related by virtue of being in the same
Bundle.

Bundle is not STIX Object, so it does not have any of the Common Properties
other than the type and id properties. Bundle is transient and implementations
should not assume that other implementations will treat it as a persistent
object.
*/
type Bundle struct {
	baseobject.CommonBundleProperties
	Objects []interface{} `json:"objects,omitempty"`
}

type BundleDecode struct {
	baseobject.CommonBundleProperties
	Objects []json.RawMessage `json:"objects,omitempty"`
}

// ----------------------------------------------------------------------
//
// Initialization Functions
//
// ----------------------------------------------------------------------

/*
New - This function will create a new STIX Bundle object and return it as
a pointer. This function can not use the InitNewObject() function as a Bundle
does not have all of the fields that are common to a standard object.
*/
func New() *Bundle {
	var obj Bundle
	obj.SetObjectType("bundle")
	obj.SetNewID("bundle")
	return &obj
}

/*
Decode - This function will decode the outer later of a bundle and stop
processing when it gets to the objects. It will leave the objects as a slice of
json.RawMessage objects. This way, later on, we can decode each one individually
*/
func Decode(r io.Reader) (*BundleDecode, error) {
	var b BundleDecode
	err := json.NewDecoder(r).Decode(&b)
	if err != nil {
		return nil, err
	}

	// // Check to make sure the object type is valid.
	// if err := b.CommonBundleProperties.TypeProperty.Verify(); err != nil {
	// 	return nil, err
	// }

	return &b, nil
}

/*
DecodeObjectType - This function will take in a slice of bytes representing a
random STIX object encoded as JSON and return the STIX object type as a string.
*/
func DecodeObjectType(data []byte) (string, error) {
	var o baseobject.CommonObjectProperties
	err := json.Unmarshal(data, &o)
	if err != nil {
		return "", err
	}

	if o.ObjectType == "" {
		return "", errors.New("Invalid STIX object: object type is missing")
	}

	return o.ObjectType, nil
}

/*
DecodeObject - This function will take in a slice of bytes representing a
random STIX object encoded as JSON, decode it to the appropriate STIX object
struct, and return object itself, its STIX ID, and any possible errors.
*/
func DecodeObject(data []byte) (interface{}, string, error) {
	var err error

	// Make a first pass to decode just the object type value. Once we have this
	// value we can easily make a second pass and decode the rest of the object.
	stixtype, err := DecodeObjectType(data)
	if err != nil {
		return nil, "", err
	}

	switch stixtype {
	case "campaign":
		o, err := campaign.Decode(data)
		if err != nil {
			return nil, "", err
		}
		return o, o.ID, nil
	case "indicator":
		o, err := indicator.Decode(data)
		if err != nil {
			return nil, "", err
		}
		return o, o.ID, nil
	case "infrastructure":
		var o infrastructure.Infrastructure
		err = json.Unmarshal(data, &o)
		fmt.Println(o.ID)
		return o, "", nil
	case "malware":
		var o malware.Malware
		err = json.Unmarshal(data, &o)
		fmt.Println(o.ID)
		return o, "", nil
	case "observed-data":
		var o observeddata.ObservedData
		err = json.Unmarshal(data, &o)
		fmt.Println(o.ID)
		return o, "", nil
	case "relationship":
		var o relationship.Relationship
		err = json.Unmarshal(data, &o)
		fmt.Println(o.ID)
		return o, "", nil
	case "sighting":
		var o sighting.Sighting
		err = json.Unmarshal(data, &o)
		fmt.Println(o.ID)
		return o, "", nil
	}
	//TODO add a default that just stores the custom object
	return nil, "", nil
}

// ----------------------------------------------------------------------
//
// Public Methods - Bundle
//
// ----------------------------------------------------------------------

/*
Encode - This method is a simple wrapper for encoding an object in to JSON
*/
func (o *BundleDecode) Encode() ([]byte, error) {
	data, err := json.MarshalIndent(o, "", "    ")
	if err != nil {
		return nil, err
	}
	return data, nil
}

/*
EncodeToString - This method is a simple wrapper for encoding an object in to JSON
*/
func (o *BundleDecode) EncodeToString() (string, error) {
	data, err := json.MarshalIndent(o, "", "    ")
	if err != nil {
		return "", err
	}
	return string(data), nil
}

/*
AddObject - This method will take in an object as an interface and add it to
the list of objects in the bundle.
*/
func (o *Bundle) AddObject(i interface{}) error {
	o.Objects = append(o.Objects, i)
	return nil
}
