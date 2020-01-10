// Copyright 2015-2020 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

package attackpattern

import (
	"encoding/json"
)

// ----------------------------------------------------------------------
// Public Functions - JSON Decoder
// ----------------------------------------------------------------------

/* Decode - This function is a simple wrapper for decoding JSON data. It will
decode a slice of bytes into an actual struct and return a pointer to that
object along with any errors. */
func Decode(data []byte) (*AttackPattern, error) {
	var o AttackPattern

	err := json.Unmarshal(data, &o)
	if err != nil {
		return nil, err
	}

	o.SetRawData(data)

	return &o, nil
}

/* UnmarshalJSON - This method will over right the default UnmarshalJSON method
to enable custom properties that this library does not know about. It will store
them as map of byte arrays. This way a tool that does know how to deal with them
can then further process them after this is done. */
func (o *AttackPattern) UnmarshalJSON(b []byte) error {
	// First thing is to capture all of the properties in a map so we can remove
	// what we know about. This will leave us with just the custom properties.
	var customProperties map[string]*json.RawMessage
	if err := json.Unmarshal(b, &customProperties); err != nil {
		return err
	}

	// Now delete the properties we know about
	delete(customProperties, "type")
	delete(customProperties, "spec_version")
	delete(customProperties, "id")
	delete(customProperties, "created_by_ref")
	delete(customProperties, "created")
	delete(customProperties, "modified")
	delete(customProperties, "revoked")
	delete(customProperties, "labels")
	delete(customProperties, "confidence")
	delete(customProperties, "lang")
	delete(customProperties, "external_references")
	delete(customProperties, "object_marking_refs")
	delete(customProperties, "granular_markings")

	delete(customProperties, "name")
	delete(customProperties, "description")
	delete(customProperties, "aliases")
	delete(customProperties, "kill_chain_phases")

	// Unmarshal the properties that we understand. We need to alias the object
	// so that we do not recursively call Unmarshal on this object.
	type alias AttackPattern
	temp := &struct {
		*alias
	}{
		alias: (*alias)(o),
	}
	if err := json.Unmarshal(b, &temp); err != nil {
		return err
	}

	// If there are any custom properties left store them in the custom property
	if len(customProperties) > 0 {
		o.Custom = make(map[string][]byte)
		for k, v := range customProperties {
			o.Custom[k] = *v
		}
	}
	return nil
}

// ----------------------------------------------------------------------
// Public Methods JSON Encoders
// The encoding is done here at the individual object level instead of at
// the STIX Object level so that individual pre/post processing rules can
// be applied. Since some of the STIX Objects do not follow a universal
// model, we need to cleanup some things that were inherited but not valid
// for the object.
// ----------------------------------------------------------------------

/* Encode - This method is a simple wrapper for encoding an object into JSON */
func (o *AttackPattern) Encode() ([]byte, error) {
	data, err := json.MarshalIndent(o, "", "    ")
	if err != nil {
		return nil, err
	}

	// Any needed preprocessing would be done here
	return data, nil
}

/* EncodeToString - This method is a simple wrapper for encoding an object into
JSON */
func (o *AttackPattern) EncodeToString() (string, error) {
	data, err := o.Encode()
	if err != nil {
		return "", err
	}
	return string(data), nil
}
