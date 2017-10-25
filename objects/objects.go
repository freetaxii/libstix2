// Copyright 2017 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license
// that can be found in the LICENSE file in the root of the source
// tree.

package objects

import (
	"github.com/freetaxii/libstix2/objects/attackPattern"
	"github.com/freetaxii/libstix2/objects/bundle"
	"github.com/freetaxii/libstix2/objects/campaign"
	"github.com/freetaxii/libstix2/objects/courseOfAction"
	"github.com/freetaxii/libstix2/objects/identity"
	"github.com/freetaxii/libstix2/objects/indicator"
	"github.com/freetaxii/libstix2/objects/infrastructure"
	"github.com/freetaxii/libstix2/objects/intrusionSet"
	"github.com/freetaxii/libstix2/objects/malware"
	"github.com/freetaxii/libstix2/objects/observedData"
	"github.com/freetaxii/libstix2/objects/relationship"
	"github.com/freetaxii/libstix2/objects/report"
	"github.com/freetaxii/libstix2/objects/sighting"
	"github.com/freetaxii/libstix2/objects/threatActor"
	"github.com/freetaxii/libstix2/objects/tool"
	"github.com/freetaxii/libstix2/objects/vulnerability"
)

// ----------------------------------------------------------------------
// Public Create Functions
// ----------------------------------------------------------------------

// NewAttackPattern - Create and return a new Attack Pattern object.
func NewAttackPattern(ver string) attackPattern.AttackPatternType {
	return attackPattern.New(ver)
}

// NewBundle - Create and return a new Bundle object.
func NewBundle() bundle.BundleType {
	return bundle.New()
}

// NewCampaign - Create and return a new Campaign object.
func NewCampaign(ver string) campaign.CampaignType {
	return campaign.New(ver)
}

// NewCourseOfAction - Create and return a new Course of Action object.
func NewCourseOfAction(ver string) courseOfAction.CourseOfActionType {
	return courseOfAction.New(ver)
}

// NewIdentity - Create and return a new Identity object.
func NewIdentity(ver string) identity.IdentityType {
	return identity.New(ver)
}

// NewIndicator - Create and return a new Indicator object.
func NewIndicator(ver string) indicator.IndicatorType {
	return indicator.New(ver)
}

// NewInfrastructure - Create and return a new Infrastructure object.
func NewInfrastructure(ver string) infrastructure.InfrastructureType {
	return infrastructure.New(ver)
}

// NewIntrusionSet - Create and return a new Intrusion Set object.
func NewIntrusionSet(ver string) intrusionSet.IntrusionSetType {
	return intrusionSet.New(ver)
}

// NewMalware - Create and return a new Malware object.
func NewMalware(ver string) malware.MalwareType {
	return malware.New(ver)
}

// NewObservedData - Create and return a new Observed Data object.
func NewObservedData(ver string) observedData.ObservedDataType {
	return observedData.New(ver)
}

// NewRelationship - Create and return a new Relationship object.
func NewRelationship(ver string) relationship.RelationshipType {
	return relationship.New(ver)
}

// NewReport - Create and return a new Report object.
func NewReport(ver string) report.ReportType {
	return report.New(ver)
}

// NewSighting - Create and return a new Sighting object.
func NewSighting(ver string) sighting.SightingType {
	return sighting.New(ver)
}

// NewThreatActor - Create and return a new Threat Actor object.
func NewThreatActor(ver string) threatActor.ThreatActorType {
	return threatActor.New(ver)
}

// NewTool - Create and return a new Tool object.
func NewTool(ver string) tool.ToolType {
	return tool.New(ver)
}

// NewVulnerability - Create and return a new Vulnerability object.
func NewVulnerability(ver string) vulnerability.VulnerabilityType {
	return vulnerability.New(ver)
}
