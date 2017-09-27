// Copyright 2017 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license
// that can be found in the LICENSE file in the root of the source
// tree.

package objects

import (
	"github.com/freetaxii/libstix2/objects/attack_pattern"
	"github.com/freetaxii/libstix2/objects/bundle"
	"github.com/freetaxii/libstix2/objects/campaign"
	"github.com/freetaxii/libstix2/objects/course_of_action"
	"github.com/freetaxii/libstix2/objects/identity"
	"github.com/freetaxii/libstix2/objects/indicator"
	"github.com/freetaxii/libstix2/objects/infrastructure"
	"github.com/freetaxii/libstix2/objects/intrusion_set"
	"github.com/freetaxii/libstix2/objects/malware"
	"github.com/freetaxii/libstix2/objects/observed_data"
	"github.com/freetaxii/libstix2/objects/relationship"
	"github.com/freetaxii/libstix2/objects/report"
	"github.com/freetaxii/libstix2/objects/sighting"
	"github.com/freetaxii/libstix2/objects/threat_actor"
	"github.com/freetaxii/libstix2/objects/tool"
	"github.com/freetaxii/libstix2/objects/vulnerability"
)

// ----------------------------------------------------------------------
// Public Create Functions
// ----------------------------------------------------------------------

// NewAttackPattern - Create and return a new Attack Pattern object.
func NewAttackPattern() attack_pattern.AttackPatternType {
	return attack_pattern.New()
}

// NewBundle - Create and return a new Bundle object.
func NewBundle() bundle.BundleType {
	return bundle.New()
}

// NewCampaign - Create and return a new Campaign object.
func NewCampaign() campaign.CampaignType {
	return campaign.New()
}

// NewCourseOfAction - Create and return a new Course of Action object.
func NewCourseOfAction() course_of_action.CourseOfActionType {
	return course_of_action.New()
}

// NewIdentity - Create and return a new Identity object.
func NewIdentity() identity.IdentityType {
	return identity.New()
}

// NewIndicator - Create and return a new Indicator object.
func NewIndicator() indicator.IndicatorType {
	return indicator.New()
}

// NewInfrastructure - Create and return a new Infrastructure object.
func NewInfrastructure() infrastructure.InfrastructureType {
	return infrastructure.New()
}

// NewIntrusionSet - Create and return a new Intrusion Set object.
func NewIntrusionSet() intrusion_set.IntrusionSetType {
	return intrusion_set.New()
}

// NewMalware - Create and return a new Malware object.
func NewMalware() malware.MalwareType {
	return malware.New()
}

// NewObservedData - Create and return a new Observed Data object.
func NewObservedData() observed_data.ObservedDataType {
	return observed_data.New()
}

// NewRelationship - Create and return a new Relationship object.
func NewRelationship() relationship.RelationshipType {
	return relationship.New()
}

// NewReport - Create and return a new Report object.
func NewReport() report.ReportType {
	return report.New()
}

// NewSighting - Create and return a new Sighting object.
func NewSighting() sighting.SightingType {
	return sighting.New()
}

// NewThreatActor - Create and return a new Threat Actor object.
func NewThreatActor() threat_actor.ThreatActorType {
	return threat_actor.New()
}

// NewTool - Create and return a new Tool object.
func NewTool() tool.ToolType {
	return tool.New()
}

// NewVulnerability - Create and return a new Vulnerability object.
func NewVulnerability() vulnerability.VulnerabilityType {
	return vulnerability.New()
}
