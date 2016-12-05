// Copyright 2016 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license
// that can be found in the LICENSE file in the root of the source
// tree.

package common

import (
	"code.google.com/p/go-uuid/uuid"
	"fmt"
	"github.com/freetaxii/libstix2/objects/defs"
	"strings"
	"time"
)

// ----------------------------------------------------------------------
// Types
// ----------------------------------------------------------------------

type CommonPropertiesType struct {
	MessageType         string                 `json:"type,omitempty"`
	Id                  string                 `json:"id,omitempty"`
	Created_by_ref      string                 `json:"created_by_ref,omitempty"`
	Created             string                 `json:"created,omitempty"`
	Modified            string                 `json:"modified,omitempty"`
	Version             int                    `json:"version,omitempty"`
	Revoked             bool                   `json:"revoked,omitempty"`
	Labels              []string               `json:"labels,omitempty"`
	External_references []ExteralReferenceType `json:"external_references,omitempty"`
	Object_marking_refs []string               `json:"object_marking_refs,omitempty"`
}

// ----------------------------------------------------------------------
// Public Methods
// ----------------------------------------------------------------------

// VerifyTimestamp will accept a timestamp in one of the two formats
// time.Time
// string
func VerifyTimestamp(t interface{}) (string, error) {
	switch ts := t.(type) {
	case time.Time:
		return ts.UTC().Format(defs.TIME_RFC_3339), nil
	case string:
		//TODO verify format of timestamp when in string format
		return ts, nil
	default:
		return "", fmt.Errorf("timestamp format of \"%s\" is not a valid format", ts)
	}
}

// VerifyPrecision will verify the supplied precision string to make sure it
// is valid per the STIX specification.
func VerifyPrecision(s string) (string, error) {

	if s == "" {
		return "", nil
	}

	s = strings.ToLower(s)
	switch s {
	case "year":
		return s, nil
	case "month":
		return s, nil
	case "day":
		return s, nil
	case "hour":
		return s, nil
	case "minute":
		return s, nil
	case "full":
		return s, nil
	default:
		return "", fmt.Errorf("invalid precision \"%s\", setting requested precision to \"\"", s)
	}
}

func NewId(s string) string {
	// TODO Add check to validate input value
	id := s + "--" + uuid.New()
	return id
}

// ----------------------------------------------------------------------
// Public Methods - CommonPropertiesType
// ----------------------------------------------------------------------

func (this *CommonPropertiesType) NewId(s string) string {
	// TODO Add check to validate input value
	id := s + "--" + uuid.New()
	return id
}

func (this *CommonPropertiesType) GetId() string {
	return this.Id
}

func (this *CommonPropertiesType) SetCreatedBy(s string) {
	this.Created_by_ref = s
}

// VerifyTimestamp will accept a timestamp in one of the two formats
// time.Time
// string
func (this *CommonPropertiesType) VerifyTimestamp(t interface{}) (string, error) {
	return VerifyTimestamp(t)
}

// VerifyPrecision will verify the supplied precision string to make sure it
// is valid per the STIX specification.
func (this *CommonPropertiesType) VerifyPrecision(s string) (string, error) {
	return VerifyPrecision(s)
}

func (this *CommonPropertiesType) GetCurrentTime() string {
	return time.Now().UTC().Format(defs.TIME_RFC_3339)
}

func (this *CommonPropertiesType) SetCreated(t interface{}) error {
	ts, err := this.VerifyTimestamp(t)
	if err != nil {
		return err
	}
	this.Created = ts
	return nil
}

func (this *CommonPropertiesType) SetModified(t interface{}) error {
	ts, err := this.VerifyTimestamp(t)
	if err != nil {
		return err
	}
	this.Modified = ts
	return nil
}

func (this *CommonPropertiesType) VerifyVersion(newVersion, currentVersion int) error {
	if newVersion < defs.MIN_VERSION_SIZE {
		return fmt.Errorf("no change made, new version is smaller than min size")
	}

	if newVersion > defs.MAX_VERSION_SIZE {
		return fmt.Errorf("no change made, new version is larger than max size")
	}

	if newVersion <= currentVersion {
		return fmt.Errorf("no change made, new version is not larger than original")
	}
	return nil
}

func (this *CommonPropertiesType) SetVersion(i int) error {
	err := this.VerifyVersion(i, this.Version)
	if err != nil {
		return err
	}
	this.Version = i
	return nil
}

func (this *CommonPropertiesType) SetRevoked() {
	this.Revoked = true
}

func (this *CommonPropertiesType) AddLabel(value string) {
	if this.Labels == nil {
		a := make([]string, 0)
		this.Labels = a
	}
	this.Labels = append(this.Labels, value)
}
