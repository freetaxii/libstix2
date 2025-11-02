// Copyright 2015-2020 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

package markingdefinition

import (
	"encoding/json"

	"github.com/freetaxii/libstix2/defs"
)

func Decode(data []byte) (*MarkingDefinition, error) {
	var o MarkingDefinition

	if err := json.Unmarshal(data, &o); err != nil {
		return nil, err
	}

	return &o, nil
}

/*
UnmarshalJSON - This method will over write the default UnmarshalJSON method
to enable custom properties that this library does not know about. It will store
them as map where the value of each key is a byte arrays. This way a tool that
does know how to deal with them can then further process them after this is
done. This will also allow the storage of the raw JSON data.
*/
func (o *MarkingDefinition) UnmarshalJSON(b []byte) error {

	type alias MarkingDefinition
	temp := &struct {
		*alias
	}{
		alias: (*alias)(o),
	}
	if err := json.Unmarshal(b, &temp); err != nil {
		return err
	}

	// This will create a map of all of the custom properties and store them in a
	// property called o.Custom
	if err := o.FindCustomProperties(b, o.GetPropertyList()); err != nil {
		return err
	}

	// This will store a complete copy of the original JSON in a byte array called
	// o.Raw. This could be useful if you need to digitally sign the JSON or do
	// verification on what was actually received.
	if defs.KEEP_RAW_DATA == true {
		o.SetRawData(b)
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

/*
Encode - This method is a simple wrapper for encoding an object into JSON
*/
func (o *MarkingDefinition) Encode() ([]byte, error) {
	data, err := json.MarshalIndent(o, "", "    ")
	if err != nil {
		return nil, err
	}

	// Any needed preprocessing would be done here
	return data, nil
}

/*
MarshalJSON - This method creates a custom JSON marshaling for MarkingDefinition
to handle extensions properly. When extensions are present, we don't include
definition and definition_type fields in the output.
*/
func (o *MarkingDefinition) MarshalJSON() ([]byte, error) {
	// If extensions are present, create a version without definition and definition_type
	if len(o.Extensions) > 0 {
		// Create a custom struct that includes only the fields we want
		temp := struct {
			Type               string                 `json:"type"`
			SpecVersion        string                 `json:"spec_version,omitempty"`
			ID                 string                 `json:"id,omitempty"`
			Created            string                 `json:"created,omitempty"`
			Modified           string                 `json:"modified,omitempty"`
			CreatedByRef       string                 `json:"created_by_ref,omitempty"`
			Revoked            bool                   `json:"revoked,omitempty"`
			Labels             []string               `json:"labels,omitempty"`
			Confidence         int                    `json:"confidence,omitempty"`
			Lang               string                 `json:"lang,omitempty"`
			ExternalReferences []interface{}          `json:"external_references,omitempty"`
			ObjectMarkingRefs  []string               `json:"object_marking_refs,omitempty"`
			GranularMarkings   []interface{}          `json:"granular_markings,omitempty"`
			Extensions         map[string]interface{} `json:"extensions,omitempty"`
		}{
			Type:         o.ObjectType,
			SpecVersion:  o.SpecVersion,
			ID:           o.ID,
			Created:      o.Created,
			Modified:     o.Modified,
			CreatedByRef: o.CreatedByRef,
			Revoked:      o.Revoked,
			Labels:       o.Labels,
			Confidence:   o.Confidence,
			Lang:         o.Lang,
			Extensions:   o.Extensions,
		}

		// Handle external references
		if len(o.ExternalReferences) > 0 {
			temp.ExternalReferences = make([]interface{}, len(o.ExternalReferences))
			for i, er := range o.ExternalReferences {
				temp.ExternalReferences[i] = er
			}
		}

		// Handle granular markings
		if len(o.GranularMarkings) > 0 {
			temp.GranularMarkings = make([]interface{}, len(o.GranularMarkings))
			for i, gm := range o.GranularMarkings {
				temp.GranularMarkings[i] = gm
			}
		}

		return json.MarshalIndent(&temp, "", "    ")
	}

	// For traditional marking definitions, include definition and definition_type
	type Alias MarkingDefinition
	return json.MarshalIndent((*Alias)(o), "", "    ")
}

/*
EncodeToString - This method is a simple wrapper for encoding an object into
JSON
*/
func (o *MarkingDefinition) EncodeToString() (string, error) {
	data, err := o.Encode()
	if err != nil {
		return "", err
	}
	return string(data), nil
}
