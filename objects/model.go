// Copyright 2015-2020 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

package objects

import (
	"fmt"

	"github.com/freetaxii/libstix2/defs"
	"github.com/freetaxii/libstix2/objects/properties"
)

// ----------------------------------------------------------------------
// Define Object Model
// ----------------------------------------------------------------------

/* STIXObject - This interface defines what methods an object must have to be
considered a STIX Object. So any new object that is created that inherits the
CommonObjectProperties is considered a STIX Object by this code. This interface
is currently used by the Bundle object to add objects to the Bundle. */
type STIXObject interface {
	GetCommonProperties() *CommonObjectProperties
}

/* CommonObjectProperties - This type defines the properties that are common to
most STIX objects. If an object does not use all of these properties, then the
Encode() function for that object will clean up and remove the properties that
might get populated by mistake. Also, there will be Init() functions for each
type of STIX object to help with populating the right properties for that type
of object. This was done so that we would only need one type that could be used
by all objects, to simplify the code. */
type CommonObjectProperties struct {
	properties.DatastoreIDProperty
	properties.TypeProperty
	properties.SpecVersionProperty
	properties.IDProperty
	properties.CreatedByRefProperty
	properties.CreatedProperty
	properties.ModifiedProperty
	properties.RevokedProperty
	properties.LabelsProperty
	properties.ConfidenceProperty
	properties.LangProperty
	properties.ExternalReferencesProperty
	properties.MarkingProperties
	properties.RawProperty
}

// ----------------------------------------------------------------------
// Initialization Functions
// ----------------------------------------------------------------------

/* InitSDO - This method will initialize a STIX Domain Object by setting all
of the basic properties and is called by the New() function from each object. */
func (o *CommonObjectProperties) InitSDO(objectType string) error {
	if defs.STRICT_TYPES {
		if valid := ValidObjectType(objectType); valid != true {
			return fmt.Errorf("invalid object type for InitSDO with strict checks enabled")
		}
	}

	o.SetSpecVersion(defs.STIX_VERSION)
	o.SetObjectType(objectType)
	o.SetNewSTIXID(objectType)
	o.SetCreatedToCurrentTime()
	o.SetModified(o.GetCreated())
	return nil
}

/* InitSRO - This method will initialize a STIX Relationship Object by setting
all of the basic properties and is called by the New() function from each
object. */
func (o *CommonObjectProperties) InitSRO(objectType string) error {
	if defs.STRICT_TYPES {
		if valid := ValidObjectType(objectType); valid != true {
			return fmt.Errorf("invalid object type for InitSRO with strict checks enabled")
		}
	}

	o.SetSpecVersion(defs.STIX_VERSION)
	o.SetObjectType(objectType)
	o.SetNewSTIXID(objectType)
	o.SetCreatedToCurrentTime()
	o.SetModified(o.GetCreated())
	return nil
}

/* InitSCO - This method will initialize a STIX Cyber Observable Object by
setting all of the basic properties and is called by the New() function from
each object. */
func (o *CommonObjectProperties) InitSCO(objectType string) error {
	if defs.STRICT_TYPES {
		if valid := ValidObjectType(objectType); valid != true {
			return fmt.Errorf("invalid object type for InitSCO with strict checks enabled")
		}
	}

	o.SetSpecVersion(defs.STIX_VERSION)
	o.SetObjectType(objectType)
	o.SetNewSTIXID(objectType)
	return nil
}

/* InitBundle - This method will initialize a STIX Bundle by setting all of the
basic properties and is called by the New() function from that object. */
func (o *CommonObjectProperties) InitBundle() error {
	o.SetObjectType("bundle")
	o.SetNewSTIXID("bundle")
	return nil
}
