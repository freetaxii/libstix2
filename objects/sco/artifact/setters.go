// Copyright 2015-2022 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

package artifact

// ----------------------------------------------------------------------
// Public Methods - Artifact - Setters
// ----------------------------------------------------------------------

/*
SetMimeType - This method takes in a string value representing the MIME type and
updates the mime_type property.
*/
func (o *Artifact) SetMimeType(s string) error {
	o.MimeType = s
	return nil
}

/*
SetPayloadBin - This method takes in a string value representing the base64
encoded payload and updates the payload_bin property.
*/
func (o *Artifact) SetPayloadBin(s string) error {
	o.PayloadBin = s
	return nil
}

/*
SetURL - This method takes in a string value representing a URL and updates the
url property.
*/
func (o *Artifact) SetURL(s string) error {
	o.URL = s
	return nil
}

/*
SetEncryptionAlgorithm - This method takes in a string value representing the
encryption algorithm and updates the encryption_algorithm property.
*/
func (o *Artifact) SetEncryptionAlgorithm(s string) error {
	o.EncryptionAlgo = s
	return nil
}

/*
SetDecryptionKey - This method takes in a string value representing the
decryption key and updates the decryption_key property.
*/
func (o *Artifact) SetDecryptionKey(s string) error {
	o.DecryptionKey = s
	return nil
}

/*
AddHash - This method takes in two parameters and adds the hash to the map. The
first is a string value representing a hash type from the STIX hashing-algorithm-ov
vocabulary. The second is a string value representing the actual hash.
*/
func (o *Artifact) AddHash(k, v string) error {
	if o.Hashes == nil {
		o.Hashes = make(map[string]string, 0)
	}
	o.Hashes[k] = v
	return nil
}
