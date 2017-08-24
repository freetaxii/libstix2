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

func NewAttackPattern() attack_pattern.AttackPatternType {
	return attack_pattern.New()
}

func NewBundle() bundle.BundleType {
	return bundle.New()
}

func NewCampaign() campaign.CampaignType {
	return campaign.New()
}

func NewCourseOfAction() course_of_action.CourseOfActionType {
	return course_of_action.New()
}

func NewIdentity() identity.IdentityType {
	return identity.New()
}

func NewIndicator() indicator.IndicatorType {
	return indicator.New()
}

func NewInfrastructure() infrastructure.InfrastructureType {
	return infrastructure.New()
}

func NewIntrusionSet() intrusion_set.IntrusionSetType {
	return intrusion_set.New()
}

func NewMalware() malware.MalwareType {
	return malware.New()
}

func NewObservedData() observed_data.ObservedDataType {
	return observed_data.New()
}

func NewRelationship() relationship.RelationshipType {
	return relationship.New()
}

func NewReport() report.ReportType {
	return report.New()
}

func NewSighting() sighting.SightingType {
	return sighting.New()
}

func NewThreatActor() threat_actor.ThreatActorType {
	return threat_actor.New()
}

func NewTool() tool.ToolType {
	return tool.New()
}

func NewVulnerability() vulnerability.VulnerabilityType {
	return vulnerability.New()
}
