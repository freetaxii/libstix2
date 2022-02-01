// Copyright 2015-2022 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

package bundle

import (
	"encoding/json"
	"io"

	"github.com/freetaxii/libstix2/objects"
	"github.com/freetaxii/libstix2/objects/attackpattern"
	"github.com/freetaxii/libstix2/objects/campaign"
	"github.com/freetaxii/libstix2/objects/courseofaction"
	"github.com/freetaxii/libstix2/objects/identity"
	"github.com/freetaxii/libstix2/objects/indicator"
	"github.com/freetaxii/libstix2/objects/infrastructure"
	"github.com/freetaxii/libstix2/objects/intrusionset"
	"github.com/freetaxii/libstix2/objects/malware"
	"github.com/freetaxii/libstix2/objects/observeddata"
	"github.com/freetaxii/libstix2/objects/relationship"
	"github.com/freetaxii/libstix2/objects/report"
	"github.com/freetaxii/libstix2/objects/sighting"
	"github.com/freetaxii/libstix2/objects/threatactor"
	"github.com/freetaxii/libstix2/objects/tool"
	"github.com/freetaxii/libstix2/objects/vulnerability"
)

// ----------------------------------------------------------------------
// Public Functions - JSON Decoder
// ----------------------------------------------------------------------

/*
Decode - This function will decode a bundle and return the object as a pointer
along with any errors found.
*/
func Decode(r io.Reader) (*Bundle, []error) {
	allErrors := make([]error, 0)

	var b Bundle
	var rawBundle bundleRawDecode

	// This will decode the outer layer of the bundle and leave all of the
	// objects as a slice of json.rawMessage bytes.
	err := json.NewDecoder(r).Decode(&rawBundle)
	if err != nil {
		// If we can not decode the outer Bundle, we can not do anything so return
		allErrors = append(allErrors, err)
		return nil, allErrors
	}

	// Populate the ID just in case a client needs or wants it
	b.SetID(rawBundle.GetID())

	// Loop through all of the raw objects and decode them
	for _, v := range rawBundle.Objects {

		// Make a first pass to decode just the object type value. Once we have this
		// value we can easily make a second pass and decode the rest of the object.
		stixtype, err := objects.DecodeType(v)
		if err != nil {
			allErrors = append(allErrors, err)
			return nil, allErrors
		}

		switch stixtype {
		case "attack-pattern":
			obj, err := attackpattern.Decode(v)
			if err != nil {
				allErrors = append(allErrors, err)
				continue
			}
			b.AddObject(obj)
		case "campaign":
			obj, err := campaign.Decode(v)
			if err != nil {
				allErrors = append(allErrors, err)
				continue
			}
			b.AddObject(obj)
		case "course-of-action":
			obj, err := courseofaction.Decode(v)
			if err != nil {
				allErrors = append(allErrors, err)
				continue
			}
			b.AddObject(obj)
		case "identity":
			obj, err := identity.Decode(v)
			if err != nil {
				allErrors = append(allErrors, err)
				continue
			}
			b.AddObject(obj)
		case "indicator":
			obj, err := indicator.Decode(v)
			if err != nil {
				allErrors = append(allErrors, err)
				continue
			}
			b.AddObject(obj)
		case "infrastructure":
			obj, err := infrastructure.Decode(v)
			if err != nil {
				allErrors = append(allErrors, err)
				continue
			}
			b.AddObject(obj)
		case "intrusion-set":
			obj, err := intrusionset.Decode(v)
			if err != nil {
				allErrors = append(allErrors, err)
				continue
			}
			b.AddObject(obj)
		case "malware":
			obj, err := malware.Decode(v)
			if err != nil {
				allErrors = append(allErrors, err)
				continue
			}
			b.AddObject(obj)
		case "observed-data":
			obj, err := observeddata.Decode(v)
			if err != nil {
				allErrors = append(allErrors, err)
				continue
			}
			b.AddObject(obj)
		case "relationship":
			obj, err := relationship.Decode(v)
			if err != nil {
				allErrors = append(allErrors, err)
				continue
			}
			b.AddObject(obj)
		case "report":
			obj, err := report.Decode(v)
			if err != nil {
				allErrors = append(allErrors, err)
				continue
			}
			b.AddObject(obj)
		case "sighting":
			obj, err := sighting.Decode(v)
			if err != nil {
				allErrors = append(allErrors, err)
				continue
			}
			b.AddObject(obj)
		case "threat-actor":
			obj, err := threatactor.Decode(v)
			if err != nil {
				allErrors = append(allErrors, err)
				continue
			}
			b.AddObject(obj)
		case "tool":
			obj, err := tool.Decode(v)
			if err != nil {
				allErrors = append(allErrors, err)
				continue
			}
			b.AddObject(obj)
		case "vulnerability":
			obj, err := vulnerability.Decode(v)
			if err != nil {
				allErrors = append(allErrors, err)
				continue
			}
			b.AddObject(obj)
		default:
			obj, err := objects.Decode(v)
			if err != nil {
				allErrors = append(allErrors, err)
				continue
			}
			b.AddObject(obj)
		}
	}

	return &b, allErrors
}

// ----------------------------------------------------------------------
// Public Methods JSON Encoders
// The encoding is done here at the individual object level instead of at
// the STIX Object level so that individual pre/post processing rules can
// be applied. Since some of the STIX Objects do not follow a universal
// model, we need to cleanup some things that were inherited but not valid
// for the object.
// ----------------------------------------------------------------------

/*
Encode - This method is a simple wrapper for encoding an object into JSON
*/
func (o *Bundle) Encode() ([]byte, error) {
	data, err := json.MarshalIndent(o, "", "    ")
	if err != nil {
		return nil, err
	}

	// Any needed preprocessing would be done here
	return data, nil
}

/*
EncodeToString - This method is a simple wrapper for encoding an object into
JSON
*/
func (o *Bundle) EncodeToString() (string, error) {
	data, err := o.Encode()
	if err != nil {
		return "", err
	}
	return string(data), nil
}
