// Copyright 2015-2020 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

/*
Package opinion implements the STIX 2.1 Relationship object.

The following information comes directly from the STIX 2.1 specification.

An Opinion is an assessment of the correctness of the information in a STIX
Object produced by a different entity. The primary property is the opinion
property, which captures the level of agreement or disagreement using a fixed
scale. That fixed scale also supports a numeric mapping to allow for consistent
statistical operations across opinions.

For example, an analyst from a consuming organization might say that they
"strongly disagree" with a Campaign object and provide an explanation about why.
In a more automated workflow, a SOC operator might give an Indicator "one star"
in their TIP (expressing "strongly disagree") because it is considered to be a
false positive within their environment. Opinions are subjective, and the
specification does not address how best to interpret them. Sharing communities
are encouraged to provide clear guidelines to their constituents regarding best
practice for the use of Opinion objects within the community.

Because Opinions are typically (though not always) created by human analysts and
are comprised of human-oriented text, they contain an additional property to
capture the analyst(s) that created the Opinion. This is distinct from the
created_by_ref property, which is meant to capture the organization that created
the object.
*/
package opinion
