// Copyright 2015-2022 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

package bundle

import (
	"encoding/json"
	"github.com/freetaxii/libstix2/objects/malwareanalysis"
	"github.com/freetaxii/libstix2/objects/sco/autonomoussystem"
	"github.com/freetaxii/libstix2/objects/sco/domainname"
	"github.com/freetaxii/libstix2/objects/sco/emailaddr"
	"github.com/freetaxii/libstix2/objects/sco/emailmessage"
	"github.com/freetaxii/libstix2/objects/sco/file"
	"github.com/freetaxii/libstix2/objects/sco/ipv4addr"
	"github.com/freetaxii/libstix2/objects/sco/ipv6addr"
	"github.com/freetaxii/libstix2/objects/sco/networktraffic"
	"github.com/freetaxii/libstix2/objects/sco/software"
	"github.com/freetaxii/libstix2/objects/sco/urlobject"
	"github.com/freetaxii/libstix2/objects/sco/x509certificate"
	"io"

	"github.com/freetaxii/libstix2/objects"
	"github.com/freetaxii/libstix2/objects/attackpattern"
	"github.com/freetaxii/libstix2/objects/campaign"
	"github.com/freetaxii/libstix2/objects/courseofaction"
	"github.com/freetaxii/libstix2/objects/identity"
	"github.com/freetaxii/libstix2/objects/indicator"
	"github.com/freetaxii/libstix2/objects/infrastructure"
	"github.com/freetaxii/libstix2/objects/intrusionset"
	"github.com/freetaxii/libstix2/objects/location"
	"github.com/freetaxii/libstix2/objects/malware"
	"github.com/freetaxii/libstix2/objects/observeddata"
	"github.com/freetaxii/libstix2/objects/relationship"
	"github.com/freetaxii/libstix2/objects/report"
	"github.com/freetaxii/libstix2/objects/sighting"
	"github.com/freetaxii/libstix2/objects/threatactor"
	"github.com/freetaxii/libstix2/objects/tool"
	"github.com/freetaxii/libstix2/objects/vulnerability"
)

type DecodeFunc func([]byte) (objects.STIXObject, error)

// ----------------------------------------------------------------------
// Public Functions - JSON Decoder
// ----------------------------------------------------------------------

/*
DecodeWithCustomObjects - This function will decode a bundle and return the object as a pointer
along with any errors found.
*/
func DecodeWithCustomObjects(r io.Reader, customDecoders map[string]DecodeFunc) (*Bundle, []error) {
	allErrors := make([]error, 0)
	decoders := map[string]DecodeFunc{
		"attack-pattern": func(bytes []byte) (objects.STIXObject, error) {
			return attackpattern.Decode(bytes)
		},
		"campaign": func(bytes []byte) (objects.STIXObject, error) {
			return campaign.Decode(bytes)
		},
		"course-of-action": func(bytes []byte) (objects.STIXObject, error) {
			return courseofaction.Decode(bytes)
		},
		"identity": func(bytes []byte) (objects.STIXObject, error) {
			return identity.Decode(bytes)
		},
		"indicator": func(bytes []byte) (objects.STIXObject, error) {
			return indicator.Decode(bytes)
		},
		"infrastructure": func(bytes []byte) (objects.STIXObject, error) {
			return infrastructure.Decode(bytes)
		},
		"intrusion-set": func(bytes []byte) (objects.STIXObject, error) {
			return intrusionset.Decode(bytes)
		},
		"malware": func(bytes []byte) (objects.STIXObject, error) {
			return malware.Decode(bytes)
		},
		"malware-analysis": func(bytes []byte) (objects.STIXObject, error) {
			return malwareanalysis.Decode(bytes)
		},
		"observed-data": func(bytes []byte) (objects.STIXObject, error) {
			return observeddata.Decode(bytes)
		},
		"relationship": func(bytes []byte) (objects.STIXObject, error) {
			return relationship.Decode(bytes)
		},
		"sighting": func(bytes []byte) (objects.STIXObject, error) {
			return sighting.Decode(bytes)
		},
		"location": func(bytes []byte) (objects.STIXObject, error) {
			return location.Decode(bytes)
		},
		"threat-actor": func(bytes []byte) (objects.STIXObject, error) {
			return threatactor.Decode(bytes)
		},
		"tool": func(bytes []byte) (objects.STIXObject, error) {
			return tool.Decode(bytes)
		},
		"vulnerability": func(bytes []byte) (objects.STIXObject, error) {
			return vulnerability.Decode(bytes)
		},
		"report": func(bytes []byte) (objects.STIXObject, error) {
			return report.Decode(bytes)
		},
		"ipv4-addr": func(bytes []byte) (objects.STIXObject, error) {
			return ipv4addr.Decode(bytes)
		},
		"ipv6-addr": func(bytes []byte) (objects.STIXObject, error) {
			return ipv6addr.Decode(bytes)
		},
		"domain-name": func(bytes []byte) (objects.STIXObject, error) {
			return domainname.Decode(bytes)
		},
		"email-addr": func(bytes []byte) (objects.STIXObject, error) {
			return emailaddr.Decode(bytes)
		},
		"email-message": func(bytes []byte) (objects.STIXObject, error) {
			return emailmessage.Decode(bytes)
		},
		"url": func(bytes []byte) (objects.STIXObject, error) {
			return urlobject.Decode(bytes)
		},
		"x509-certificate": func(bytes []byte) (objects.STIXObject, error) {
			return x509certificate.Decode(bytes)
		},
		"autonomous-system": func(bytes []byte) (objects.STIXObject, error) {
			return autonomoussystem.Decode(bytes)
		},
		"network-traffic": func(bytes []byte) (objects.STIXObject, error) {
			return networktraffic.Decode(bytes)
		},
		"software": func(bytes []byte) (objects.STIXObject, error) {
			return software.Decode(bytes)
		},
		"file": func(bytes []byte) (objects.STIXObject, error) {
			return file.Decode(bytes)
		},
	}
	for k, v := range customDecoders {
		decoders[k] = v
	}

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

		if function, ok := decoders[stixtype]; ok {
			obj, err := function(v)
			if err != nil {
				allErrors = append(allErrors, err)
				continue
			}
			b.AddObject(obj)
		} else {
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

/*
Decode - This function will decode a bundle and return the object as a pointer
along with any errors found.
*/
func Decode(r io.Reader) (*Bundle, []error) {
	return DecodeWithCustomObjects(r, map[string]DecodeFunc{})
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
