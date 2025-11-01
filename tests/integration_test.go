// Copyright 2015-2022 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

package tests

import (
	"testing"
	
	"github.com/freetaxii/libstix2/objects"
	"github.com/freetaxii/libstix2/objects/sco/artifact"
	"github.com/freetaxii/libstix2/objects/sco/file"
	"github.com/freetaxii/libstix2/objects/sco/ipv6addr"
	"github.com/freetaxii/libstix2/objects/sco/macaddr"
	"github.com/freetaxii/libstix2/objects/sco/mutex"
	"github.com/freetaxii/libstix2/objects/languagecontent"
)

// TestNewSCOImplementations tests that all new SCO implementations can be created
func TestNewSCOImplementations(t *testing.T) {
	tests := []struct {
		name       string
		createFunc func() interface{}
		expectedType string
	}{
		{
			name:         "IPv6 Address",
			createFunc:   func() interface{} { return ipv6addr.New() },
			expectedType: "ipv6-addr",
		},
		{
			name:         "MAC Address",
			createFunc:   func() interface{} { return macaddr.New() },
			expectedType: "mac-addr",
		},
		{
			name:         "File",
			createFunc:   func() interface{} { return file.New() },
			expectedType: "file",
		},
		{
			name:         "Artifact",
			createFunc:   func() interface{} { return artifact.New() },
			expectedType: "artifact",
		},
		{
			name:         "Mutex",
			createFunc:   func() interface{} { return mutex.New() },
			expectedType: "mutex",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			obj := tt.createFunc()
			
			// Use type assertion to get CommonObjectProperties
			type hasCommonProps interface {
				GetCommonProperties() *objects.CommonObjectProperties
			}
			
			if cobj, ok := obj.(hasCommonProps); ok {
				props := cobj.GetCommonProperties()
				if props.ObjectType != tt.expectedType {
					t.Errorf("%s: Expected type %s, got %s", tt.name, tt.expectedType, props.ObjectType)
				}
				if props.SpecVersion == "" {
					t.Errorf("%s: SpecVersion not set", tt.name)
				}
				if props.ID == "" {
					t.Errorf("%s: ID not set", tt.name)
				}
			} else {
				t.Errorf("%s: Object does not implement GetCommonProperties", tt.name)
			}
		})
	}
}

// TestLanguageContentImplementation tests the language-content meta object
func TestLanguageContentImplementation(t *testing.T) {
	obj := languagecontent.New()
	
	if obj.ObjectType != "language-content" {
		t.Errorf("Expected type language-content, got %s", obj.ObjectType)
	}
	
	if obj.SpecVersion == "" {
		t.Error("SpecVersion not set")
	}
	
	if obj.ID == "" {
		t.Error("ID not set")
	}
	
	// Test setting properties
	obj.SetObjectRef("indicator--8e2e2d2b-17d4-4cbf-938f-98ee46b3cd3f")
	obj.AddContent("es", "name", "Indicador")
	obj.AddContent("fr", "name", "Indicateur")
	
	if obj.ObjectRef == "" {
		t.Error("ObjectRef not set")
	}
	
	if len(obj.Contents) != 2 {
		t.Errorf("Expected 2 languages, got %d", len(obj.Contents))
	}
	
	// Test validation
	valid, problems, _ := obj.Valid(false)
	if !valid {
		t.Errorf("Object should be valid, found %d problems", problems)
	}
}

// TestSCOPropertySetters tests that property setters work correctly
func TestSCOPropertySetters(t *testing.T) {
	t.Run("IPv6 Address Properties", func(t *testing.T) {
		obj := ipv6addr.New()
		obj.SetValue("2001:db8::1")
		obj.AddResolvesToRefs("mac-addr--test")
		obj.AddBelongsToRefs("autonomous-system--test")
		
		valid, _, _ := obj.Valid(false)
		if !valid {
			t.Error("IPv6 object with all properties should be valid")
		}
	})
	
	t.Run("File Properties", func(t *testing.T) {
		obj := file.New()
		obj.SetName("malware.exe")
		obj.AddHash("SHA-256", "abc123")
		obj.SetSize(1024)
		obj.SetMimeType("application/octet-stream")
		
		if obj.Name != "malware.exe" {
			t.Error("Name not set correctly")
		}
		if len(obj.Hashes) != 1 {
			t.Error("Hash not added correctly")
		}
		if obj.Size != 1024 {
			t.Error("Size not set correctly")
		}
		
		valid, _, _ := obj.Valid(false)
		if !valid {
			t.Error("File object with all properties should be valid")
		}
	})
	
	t.Run("Artifact Properties", func(t *testing.T) {
		obj := artifact.New()
		obj.SetPayloadBin("VGVzdA==")
		obj.SetMimeType("text/plain")
		obj.AddHash("MD5", "test-hash")
		
		if obj.PayloadBin != "VGVzdA==" {
			t.Error("PayloadBin not set correctly")
		}
		if len(obj.Hashes) != 1 {
			t.Error("Hash not added correctly")
		}
		
		valid, _, _ := obj.Valid(false)
		if !valid {
			t.Error("Artifact object should be valid")
		}
	})
}

// TestObjectTypeValidation tests that ValidObjectType includes all new types
func TestObjectTypeValidation(t *testing.T) {
	newTypes := []string{
		"ipv6-addr",
		"mac-addr",
		"file",
		"artifact",
		"mutex",
		"autonomous-system",
		"directory",
		"email-addr",
		"email-message",
		"network-traffic",
		"process",
		"software",
		"user-account",
		"windows-registry-key",
		"x509-certificate",
		"language-content",
	}
	
	for _, objType := range newTypes {
		if !objects.ValidObjectType(objType) {
			t.Errorf("Object type %s should be valid", objType)
		}
	}
}
