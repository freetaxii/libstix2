// Copyright 2015-2022 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

/*
Package note implements the STIX 2.1 Attack Pattern object.

The following information comes directly from the STIX 2.1 specification.

This implements STIX 2.1 specification section 4.12.
Reference: https://docs.oasis-open.org/cti/stix/v2.1/csprd01/stix-v2.1-csprd01.html#_Toc16070644

A Note is intended to convey informative text to provide further context and/or
to provide additional analysis not contained in the STIX Objects, Marking
Definition objects, or Language Content objects which the Note relates to. Notes
can be created by anyone (not just the original object creator).

For example, an analyst may add a Note to a Campaign object created by another
organization indicating that they've seen posts related to that Campaign on a
hacker forum.

Because Notes are typically (though not always) created by human analysts and
are comprised of human-oriented text, they contain an additional property to
capture the analyst(s) that created the Note. This is distinct from the
created_by_ref property, which is meant to capture the organization that created
the object.
*/
package note
