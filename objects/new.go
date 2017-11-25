// Copyright 2017 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license
// that can be found in the LICENSE file in the root of the source
// tree.

package objects

// ----------------------------------------------------------------------
// Public Create Functions
// ----------------------------------------------------------------------

// NewAttackPattern - Create and return a new Attack Pattern object.
func NewAttackPattern(ver string) AttackPatternType {
	var obj AttackPatternType
	obj.InitNewObject("attack-pattern", ver)
	return obj
}

// NewBundle - Create and return a new Bundle object.
// This function can not use the InitNewObject() function as a Bundle does not
// have all of the fields that are common to a standard object.
func NewBundle() BundleType {
	var obj BundleType
	obj.SetObjectType("bundle")
	obj.NewID("bundle")
	obj.SetSpecVersion20()
	return obj
}

// NewCampaign - Create and return a new Campaign object.
func NewCampaign(ver string) CampaignType {
	var obj CampaignType
	obj.InitNewObject("campaign", ver)
	return obj
}

// NewCourseOfAction - Create and return a new Course of Action object.
func NewCourseOfAction(ver string) CourseOfActionType {
	var obj CourseOfActionType
	obj.InitNewObject("course-of-action", ver)
	return obj
}

// NewIdentity - Create and return a new Identity object.
func NewIdentity(ver string) IdentityType {
	var obj IdentityType
	obj.InitNewObject("identity", ver)
	return obj
}

// NewIndicator - Create and return a new Indicator object.
func NewIndicator(ver string) IndicatorType {
	var obj IndicatorType
	obj.InitNewObject("indicator", ver)
	return obj
}

// NewInfrastructure - Create and return a new Infrastructure object.
func NewInfrastructure(ver string) InfrastructureType {
	var obj InfrastructureType
	obj.InitNewObject("infrastructure", ver)
	return obj
}

// NewIntrusionSet - Create and return a new Intrusion Set object.
func NewIntrusionSet(ver string) IntrusionSetType {
	var obj IntrusionSetType
	obj.InitNewObject("intrusion-set", ver)
	return obj
}

// NewMalware - Create and return a new Malware object.
func NewMalware(ver string) MalwareType {
	var obj MalwareType
	obj.InitNewObject("malware", ver)
	return obj
}

// NewObservedData - Create and return a new Observed Data object.
func NewObservedData(ver string) ObservedDataType {
	var obj ObservedDataType
	obj.InitNewObject("observed-data", ver)
	return obj
}

// NewRelationship - Create and return a new Relationship object.
func NewRelationship(ver string) RelationshipType {
	var obj RelationshipType
	obj.InitNewObject("relationship", ver)
	return obj
}

// NewReport - Create and return a new Report object.
func NewReport(ver string) ReportType {
	var obj ReportType
	obj.InitNewObject("report", ver)
	return obj
}

// NewSighting - Create and return a new Sighting object.
func NewSighting(ver string) SightingType {
	var obj SightingType
	obj.InitNewObject("sighting", ver)
	return obj
}

// NewThreatActor - Create and return a new Threat Actor object.
func NewThreatActor(ver string) ThreatActorType {
	var obj ThreatActorType
	obj.InitNewObject("threat-actor", ver)
	return obj
}

// NewTool - Create and return a new Tool object.
func NewTool(ver string) ToolType {
	var obj ToolType
	obj.InitNewObject("tool", ver)
	return obj
}

// NewVulnerability - Create and return a new Vulnerability object.
func NewVulnerability(ver string) VulnerabilityType {
	var obj VulnerabilityType
	obj.InitNewObject("vulnerability", ver)
	return obj
}
