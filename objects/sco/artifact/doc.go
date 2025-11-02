// Copyright 2015-2022 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

/*
Package artifact implements the STIX 2.1 Artifact SCO object.

The following information comes directly from the STIX 2.1 specification.

This implements STIX 2.1 specification section 6.1 - Artifact Object.

The Artifact object permits capturing an array of bytes (8-bits), as a
base64-encoded string, or linking to a file-like payload. One of payload_bin or
url MUST be provided. It is incumbent on object creators to ensure that the URL
is accessible for downstream consumers.
*/
package artifact
