// Copyright 2015-2022 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

package objects

import "fmt"

// ----------------------------------------------------------------------
// Aliases Property
// ----------------------------------------------------------------------

// AliasesProperty - A property used by one or more STIX objects.
type AliasesProperty struct {
	Aliases []string `json:"aliases,omitempty" bson:"aliases,omitempty"`
}

// AddAliases - This method takes in a string value, a comma separated list of
// string values, or a slice of string values that represents an alias and adds it
// to the aliases property.
func (o *AliasesProperty) AddAliases(values interface{}) error {
	return AddValuesToList(&o.Aliases, values)
}

// Compare - This method will compare two properties to make sure they are the
// same and will return a boolean, an integer that tracks the number of problems
// found, and a slice of strings that contain the detailed results, whether good or
// bad.
func (o *AliasesProperty) Compare(obj2 *AliasesProperty, debug bool) (bool, int, []string) {
	var r *results = new(results)
	r.debug = debug

	if len(o.Aliases) != len(obj2.Aliases) {
		str := fmt.Sprintf("-- The number of entries in aliases do not match: %d | %d", len(o.Aliases), len(obj2.Aliases))
		logProblem(r, str)
	} else {
		str := fmt.Sprintf("++ The number of entries in aliases match: %d | %d", len(o.Aliases), len(obj2.Aliases))
		logValid(r, str)

		// If lengths are the same, then check each value
		for index := range o.Aliases {
			if o.Aliases[index] != obj2.Aliases[index] {
				str := fmt.Sprintf("-- The alias values do not match: %s | %s", o.Aliases[index], obj2.Aliases[index])
				logProblem(r, str)
			} else {
				str := fmt.Sprintf("++ The alias values match: %s | %s", o.Aliases[index], obj2.Aliases[index])
				logValid(r, str)
			}
		}
	}

	// Return real values not pointers
	if r.problemsFound > 0 {
		return false, r.problemsFound, r.resultDetails
	}
	return true, 0, r.resultDetails
}

// ----------------------------------------------------------------------
// Authors Property
// ----------------------------------------------------------------------

// AuthorsProperty - A property used by one or more STIX objects.
type AuthorsProperty struct {
	Authors []string `json:"authors,omitempty" bson:"authors,omitempty"`
}

// AddAuthors - This method takes in a string value, a comma separated list of
// string values, or a slice of string values that represents a alias and adds it
// to the authors property.
func (o *AuthorsProperty) AddAuthors(values interface{}) error {
	return AddValuesToList(&o.Authors, values)
}

// Compare - This method will compare two properties to make sure they are the
// same and will return a boolean, an integer that tracks the number of problems
// found, and a slice of strings that contain the detailed results, whether good or
// bad.
func (o *AuthorsProperty) Compare(obj2 *AuthorsProperty, debug bool) (bool, int, []string) {
	var r *results = new(results)
	r.debug = debug

	if len(o.Authors) != len(obj2.Authors) {
		str := fmt.Sprintf("-- The number of entries in authors do not match: %d | %d", len(o.Authors), len(obj2.Authors))
		logProblem(r, str)
	} else {
		str := fmt.Sprintf("++ The number of entries in authors match: %d | %d", len(o.Authors), len(obj2.Authors))
		logValid(r, str)

		// If lengths are the same, then check each value
		for index := range o.Authors {
			if o.Authors[index] != obj2.Authors[index] {
				str := fmt.Sprintf("-- The author values do not match: %s | %s", o.Authors[index], obj2.Authors[index])
				logProblem(r, str)
			} else {
				str := fmt.Sprintf("++ The author values match: %s | %s", o.Authors[index], obj2.Authors[index])
				logValid(r, str)
			}
		}
	}

	// Return real values not pointers
	if r.problemsFound > 0 {
		return false, r.problemsFound, r.resultDetails
	}
	return true, 0, r.resultDetails
}

// ----------------------------------------------------------------------
// Belongs to Refs Property
// ----------------------------------------------------------------------

// BelongsToRefsProperty -
type BelongsToRefsProperty struct {
	BelongsToRefs []string `json:"belongs_to_refs,omitempty" bson:"belongs_to_refs,omitempty"`
}

// AddBelongsToRefs - This method takes in a string value, a comma separated
// list of string values, or a slice of string values that represents an id that
// this object belongs to and adds it to the belongs to refs property.
func (o *BelongsToRefsProperty) AddBelongsToRefs(values interface{}) error {
	return AddValuesToList(&o.BelongsToRefs, values)
}

// ----------------------------------------------------------------------
// Description Property
// ----------------------------------------------------------------------

// DescriptionProperty - A property used by one or more STIX objects that
// captures the description for the object as a string.
type DescriptionProperty struct {
	Description string `json:"description,omitempty" bson:"description,omitempty"`
}

// SetDescription - This method takes in a string value representing a text
// description and updates the description property.
func (o *DescriptionProperty) SetDescription(s string) error {
	o.Description = s
	return nil
}

// GetDescription - This method returns the description for an object as a
// string.
func (o *DescriptionProperty) GetDescription() string {
	return o.Description
}

// Compare - This method will compare two properties to make sure they are the
// same and will return a boolean, an integer that tracks the number of problems
// found, and a slice of strings that contain the detailed results, whether good or
// bad.
func (o *DescriptionProperty) Compare(obj2 *DescriptionProperty, debug bool) (bool, int, []string) {
	var r *results = new(results)
	r.debug = debug

	// Check Description Value
	if o.Description != obj2.Description {
		str := fmt.Sprintf("-- The description values do not match: %s | %s", o.Description, obj2.Description)
		logProblem(r, str)
	} else {
		str := fmt.Sprintf("++ The description values match: %s | %s", o.Description, obj2.Description)
		logValid(r, str)
	}

	// Return real values not pointers
	if r.problemsFound > 0 {
		return false, r.problemsFound, r.resultDetails
	}
	return true, 0, r.resultDetails
}

// ----------------------------------------------------------------------
// Goals Property
// ----------------------------------------------------------------------

// GoalsProperty - A property used by one or more STIX objects that captures a
// list of goals that are part of the STIX object.
type GoalsProperty struct {
	Goals []string `json:"goals,omitempty" bson:"goals,omitempty"`
}

// AddGoals - This method takes in a string value, a comma separated list of
// string values, or a slice of string values that represents a goal and adds it to
// the goals property.
func (o *GoalsProperty) AddGoals(values interface{}) error {
	return AddValuesToList(&o.Goals, values)
}

// ----------------------------------------------------------------------
// Kill Chain Property
// ----------------------------------------------------------------------

// KillChainPhasesProperty - A property used by one or more STIX objects that
// captures a list of kll chain phases as defined by STIX.
type KillChainPhasesProperty struct {
	KillChainPhases []KillChainPhase `json:"kill_chain_phases,omitempty" bson:"kill_chain_phases,omitempty"`
}

// KillChainPhase - This type defines all of the properties associated with the
// STIX Kill Chain Phase type.
type KillChainPhase struct {
	KillChainName string `json:"kill_chain_name,omitempty" bson:"kill_chain_name,omitempty"`
	PhaseName     string `json:"phase_name,omitempty" bson:"phase_name,omitempty"`
}

// CreateKillChainPhase - This method takes in two parameters and creates and
// adds a new kill chain phase to the list. The first value is a string value
// representing the name of the kill chain being used. The second value is a string
// value representing the phase name from that kill chain.
func (o *KillChainPhasesProperty) CreateKillChainPhase(name, phase string) error {
	k, _ := o.newKillChainPhase()
	k.SetName(name)
	k.SetPhase(phase)
	return nil
}

// newKillChainPhase - This method returns a reference to a slice location. This
// will enable the code to update an object located at that slice location.
func (o *KillChainPhasesProperty) newKillChainPhase() (*KillChainPhase, error) {
	var s KillChainPhase

	// if o.KillChainPhases == nil {
	// 	a := make([]KillChainPhase, 0)
	// 	o.KillChainPhases = a
	// }

	positionThatAppendWillUse := len(o.KillChainPhases)
	o.KillChainPhases = append(o.KillChainPhases, s)
	return &o.KillChainPhases[positionThatAppendWillUse], nil
}

// SetName - This method takes in a string value representing the name of a kill
// chain and updates the kill chain name property.
func (o *KillChainPhase) SetName(s string) error {
	o.KillChainName = s
	return nil
}

// SetPhase - This method takes in a string value representing the phase of a
// kill chain and updates the phase name property.
func (o *KillChainPhase) SetPhase(s string) error {
	o.PhaseName = s
	return nil
}

// Compare - This method will compare two properties to make sure they are the
// same and will return a boolean, an integer that tracks the number of problems
// found, and a slice of strings that contain the detailed results, whether good or
// bad.
func (o *KillChainPhasesProperty) Compare(obj2 *KillChainPhasesProperty, debug bool) (bool, int, []string) {
	var r *results = new(results)
	r.debug = debug

	// Check Kill Chain Phases Property Length
	if len(o.KillChainPhases) != len(obj2.KillChainPhases) {
		str := fmt.Sprintf("-- The number of entries in kill chain phases do not match: %d | %d", len(o.KillChainPhases), len(obj2.KillChainPhases))
		logProblem(r, str)
	} else {
		str := fmt.Sprintf("++ The number of entries in kill chain phases match: %d | %d", len(o.KillChainPhases), len(obj2.KillChainPhases))
		logValid(r, str)

		for index := range o.KillChainPhases {
			// Check Kill Chain Phases values
			if o.KillChainPhases[index].KillChainName != obj2.KillChainPhases[index].KillChainName {
				str := fmt.Sprintf("-- The kill chain name values do not match: %s | %s", o.KillChainPhases[index].KillChainName, obj2.KillChainPhases[index].KillChainName)
				logProblem(r, str)
			} else {
				str := fmt.Sprintf("++ The kill chain name values match: %s | %s", o.KillChainPhases[index].KillChainName, obj2.KillChainPhases[index].KillChainName)
				logValid(r, str)
			}

			// Check Kill Chain Phases values
			if o.KillChainPhases[index].PhaseName != obj2.KillChainPhases[index].PhaseName {
				str := fmt.Sprintf("-- The kill chain phase values do not match: %s | %s", o.KillChainPhases[index].PhaseName, obj2.KillChainPhases[index].PhaseName)
				logProblem(r, str)
			} else {
				str := fmt.Sprintf("++ The kill chain phase values match: %s | %s", o.KillChainPhases[index].PhaseName, obj2.KillChainPhases[index].PhaseName)
				logValid(r, str)
			}
		}
	}

	// Return real values not pointers
	if r.problemsFound > 0 {
		return false, r.problemsFound, r.resultDetails
	}
	return true, 0, r.resultDetails
}

// ----------------------------------------------------------------------
// Motivation Properties
// ----------------------------------------------------------------------

// MotivationProperties - Properties used by one or more STIX objects that
// capture the primary and secondary motivations.
type MotivationProperties struct {
	PrimaryMotivation    string   `json:"primary_motivation,omitempty" bson:"primary_motivation,omitempty"`
	SecondaryMotivations []string `json:"secondary_motivations,omitempty" bson:"secondary_motivations,omitempty"`
}

// ----------------------------------------------------------------------
// Public Methods - MotivationProperties - Setters
// ----------------------------------------------------------------------

// SetPrimaryMotivation - This methods takes in a string value representing a
// motivation from the attack-motivation-ov vocab and updates the primary
// motivation property.
func (o *MotivationProperties) SetPrimaryMotivation(s string) error {
	o.PrimaryMotivation = s
	return nil
}

// GetPrimaryMotivation - This method returns the primary motivation.
func (o *MotivationProperties) GetPrimaryMotivation() string {
	return o.PrimaryMotivation
}

// AddSecondaryMotivations - This method takes in a string value, a comma
// separated list of string values, or a slice of string values that represents a
// secondary motivation and adds it to the secondary motivations property.
func (o *MotivationProperties) AddSecondaryMotivations(values interface{}) error {
	return AddValuesToList(&o.SecondaryMotivations, values)
}

// ----------------------------------------------------------------------
// Name Property
// ----------------------------------------------------------------------

// NameProperty - A property used by one or more STIX objects that captures a
// vanity name for the STIX object in string format.
type NameProperty struct {
	Name string `json:"name,omitempty" bson:"name,omitempty"`
}

// SetName - This method takes in a string value representing a name of the
// object and updates the name property.
func (o *NameProperty) SetName(s string) error {
	o.Name = s
	return nil
}

// GetName - This method returns the current name of the object.
func (o *NameProperty) GetName() string {
	return o.Name
}

// Compare - This method will compare two properties to make sure they are the
// same and will return a boolean, an integer that tracks the number of problems
// found, and a slice of strings that contain the detailed results, whether good or
// bad.
func (o *NameProperty) Compare(obj2 *NameProperty, debug bool) (bool, int, []string) {
	var r *results = new(results)
	r.debug = debug

	// Check Name Value
	if o.Name != obj2.Name {
		str := fmt.Sprintf("-- The names do not match: %s | %s", o.Name, obj2.Name)
		logProblem(r, str)
	} else {
		str := fmt.Sprintf("++ The names match: %s | %s", o.Name, obj2.Name)
		logValid(r, str)
	}

	// Return real values not pointers
	if r.problemsFound > 0 {
		return false, r.problemsFound, r.resultDetails
	}
	return true, 0, r.resultDetails
}

// ----------------------------------------------------------------------
// Object refs property
// ----------------------------------------------------------------------

// ObjectRefsProperty - A property used by one or more STIX objects.
type ObjectRefsProperty struct {
	ObjectRefs []string `json:"object_refs,omitempty" bson:"object_refs,omitempty"`
}

// AddObjectRefs - This method takes in a string value, a comma separated list of
// string values, or a slice of string values that represents a alias and adds it
// to the object refs property.
func (o *ObjectRefsProperty) AddObjectRefs(values interface{}) error {
	return AddValuesToList(&o.ObjectRefs, values)
}

// Compare - This method will compare two properties to make sure they are the
// same and will return a boolean, an integer that tracks the number of problems
// found, and a slice of strings that contain the detailed results, whether good or
// bad.
func (o *ObjectRefsProperty) Compare(obj2 *ObjectRefsProperty, debug bool) (bool, int, []string) {
	var r *results = new(results)
	r.debug = debug

	if len(o.ObjectRefs) != len(obj2.ObjectRefs) {
		str := fmt.Sprintf("-- The number of entries in object refs do not match: %d | %d", len(o.ObjectRefs), len(obj2.ObjectRefs))
		logProblem(r, str)
	} else {
		str := fmt.Sprintf("++ The number of entries in object refs match: %d | %d", len(o.ObjectRefs), len(obj2.ObjectRefs))
		logValid(r, str)

		// If lengths are the same, then check each value
		for index := range o.ObjectRefs {
			if o.ObjectRefs[index] != obj2.ObjectRefs[index] {
				str := fmt.Sprintf("-- The object ref values do not match: %s | %s", o.ObjectRefs[index], obj2.ObjectRefs[index])
				logProblem(r, str)
			} else {
				str := fmt.Sprintf("++ The object ref values match: %s | %s", o.ObjectRefs[index], obj2.ObjectRefs[index])
				logValid(r, str)
			}
		}
	}

	// Return real values not pointers
	if r.problemsFound > 0 {
		return false, r.problemsFound, r.resultDetails
	}
	return true, 0, r.resultDetails
}

// ----------------------------------------------------------------------
// Resource Level Property
// ----------------------------------------------------------------------

// ResourceLevelProperty - A property used by one or more STIX objects that
// captures the resource level.
type ResourceLevelProperty struct {
	ResourceLevel string `json:"resource_level,omitempty" bson:"resource_level,omitempty"`
}

// SetResourceLevel - This method takes in a string value representing a
// resource level from the attack-resrouce-level-ov vocab and updates the resource
// level property.
func (o *ResourceLevelProperty) SetResourceLevel(s string) error {
	o.ResourceLevel = s
	return nil
}

// GetResourceLevel - This method returns the resource level.
func (o *ResourceLevelProperty) GetResourceLevel() string {
	return o.ResourceLevel
}

// ----------------------------------------------------------------------
// Roles Property
// ----------------------------------------------------------------------

// RolesProperty - A property used by one or more STIX objects that captures a
// list of roles that are part of the STIX object.
type RolesProperty struct {
	Roles []string `json:"roles,omitempty" bson:"roles,omitempty"`
}

// AddRoles - This method takes in a string value, a comma separated list of
// string values, or a slice of string values that represents a role and adds it to
// the roles property.
func (o *RolesProperty) AddRoles(values interface{}) error {
	return AddValuesToList(&o.Roles, values)
}

// ----------------------------------------------------------------------
// Resolves To Ref Property
// ----------------------------------------------------------------------

// ResolvesToRefsProperty -
type ResolvesToRefsProperty struct {
	ResolvesToRefs []string `json:"resolves_to_refs,omitempty" bson:"resolves_to_refs,omitempty"`
}

// AddResolvesToRefs - This method takes in a string value, a comma separated
// list of string values, or a slice of string values that represents an id of an
// object that this resolves to and adds it to the resolves to refs property.
func (o *ResolvesToRefsProperty) AddResolvesToRefs(values interface{}) error {
	return AddValuesToList(&o.ResolvesToRefs, values)
}

// ----------------------------------------------------------------------
// Seen Properties
// ----------------------------------------------------------------------

// SeenProperties - Properties used by one or more STIX objects that
// captures the time that this object was first and last seen in STIX timestamp
// format.
type SeenProperties struct {
	FirstSeen string `json:"first_seen,omitempty" bson:"first_seen,omitempty"`
	LastSeen  string `json:"last_seen,omitempty" bson:"last_seen,omitempty"`
}

// SetFirstSeenToCurrentTime - This methods sets the first seen time to the
// current time
func (o *SeenProperties) SetFirstSeenToCurrentTime() error {
	o.FirstSeen = GetCurrentTime("micro")
	return nil
}

// SetFirstSeen -  This method takes in a timestamp in either time.Time or
// string format and updates the first seen property with it. The value is stored
// as a string, so if the value is in time.Time format, it will be converted to the
// correct STIX timestamp format.
func (o *SeenProperties) SetFirstSeen(t interface{}) error {
	ts, _ := TimeToString(t, "micro")
	o.FirstSeen = ts
	return nil
}

// SetLastSeenToCurrentTime - This methods sets the first seen time to the
// current time
func (o *SeenProperties) SetLastSeenToCurrentTime() error {
	o.LastSeen = GetCurrentTime("micro")
	return nil
}

// SetLastSeen -  This method takes in a time stamp in either time.Time or
// string format and updates the last seen property with it. The value is stored as
// a string, so if the value is in time.Time format, it will be converted to the
// correct STIX time stamp format.
func (o *SeenProperties) SetLastSeen(t interface{}) error {
	ts, _ := TimeToString(t, "micro")
	o.LastSeen = ts
	return nil
}

// ----------------------------------------------------------------------
// ID Property
// ----------------------------------------------------------------------

// IDProperty - A property used by one or more STIX objects that captures the
// unique identifier for the object.
type IDProperty struct {
	ID string `json:"id,omitempty" bson:"id,omitempty"`
}

// SetID - This method takes in a string value representing a unique identifier
// and updates the id property.
func (o *IDProperty) SetID(s string) error {
	o.ID = s
	return nil
}

// GetID - This method returns the id for an object as a string.
func (o *IDProperty) GetID() string {
	return o.ID
}

// ----------------------------------------------------------------------
// Title Properties
// ----------------------------------------------------------------------

// TitleProperty - A property used by one or more TAXII resources.
type TitleProperty struct {
	Title string `json:"title" bson:"title"`
}

// SetTitle - This method takes in a string value representing a title or name
// and updates the title property.
func (o *TitleProperty) SetTitle(s string) error {
	o.Title = s
	return nil
}

// GetTitle - This method returns the title.
func (o *TitleProperty) GetTitle() string {
	return o.Title
}

// ----------------------------------------------------------------------
// Value Property
// ----------------------------------------------------------------------

// ValueProperty -
type ValueProperty struct {
	Value string `json:"value,omitempty" bson:"value,omitempty"`
}

// SetValue -
func (o *ValueProperty) SetValue(val string) error {
	o.Value = val
	return nil
}

/*
	VerifyExists - This method will verify that the value property on an object

is present. It will return a boolean, an integer that tracks the number of
problems found, and a slice of strings that contain the detailed results,
whether good or bad.
*/
func (o *ValueProperty) VerifyExists() (bool, int, []string) {
	problemsFound := 0
	resultDetails := make([]string, 1)

	if o.Value == "" {
		problemsFound++
		resultDetails[0] = fmt.Sprintf("-- The value property is required but missing")
		return false, problemsFound, resultDetails
	}

	resultDetails[0] = fmt.Sprintf("++ The value property is required and is present")
	return true, problemsFound, resultDetails
}
