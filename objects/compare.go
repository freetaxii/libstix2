// Copyright 2015-2020 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

package objects

/*
Compare - This method will compare the common properties from two objects to
make sure they are the same. The common properties receiver is object 1 and the
common properties passed in is object 2. This method will return an integer that
tracks the number of problems and a slice of strings that contain the detailed
results, whether good or bad.
*/
func (o *CommonObjectProperties) Compare(obj2 *CommonObjectProperties) (bool, int, []string) {
	return Compare(o, obj2)
}

/* Compare - This function will compare two objects to make sure they are the
same and will return a boolean, an integer that tracks the number of problems
found, and a slice of strings that contain the detailed results, whether good or
bad. */
func Compare(obj1, obj2 *CommonObjectProperties) (bool, int, []string) {
	problemsFound := 0
	resultDetails := make([]string, 0)

	// Check Type Value
	_, pTypes, dTypes := obj1.TypeProperty.Compare(&obj2.TypeProperty)
	problemsFound += pTypes
	resultDetails = append(resultDetails, dTypes...)

	// Check Spec Version Value
	_, pSpecVersions, dSpecVersions := obj1.SpecVersionProperty.Compare(&obj2.SpecVersionProperty)
	problemsFound += pSpecVersions
	resultDetails = append(resultDetails, dSpecVersions...)

	// Check ID Value
	_, pIDs, dIDs := obj1.IDProperty.Compare(&obj2.IDProperty)
	problemsFound += pIDs
	resultDetails = append(resultDetails, dIDs...)

	// Check Created By Ref Value
	_, pCreatedByRefs, dCreatedByRefs := obj1.CreatedByRefProperty.Compare(&obj2.CreatedByRefProperty)
	problemsFound += pCreatedByRefs
	resultDetails = append(resultDetails, dCreatedByRefs...)

	// Check Created Values
	_, pCreated, dCreated := obj1.CreatedProperty.Compare(&obj2.CreatedProperty)
	problemsFound += pCreated
	resultDetails = append(resultDetails, dCreated...)

	// Check Modified Values
	_, pModified, dModified := obj1.ModifiedProperty.Compare(&obj2.ModifiedProperty)
	problemsFound += pModified
	resultDetails = append(resultDetails, dModified...)

	// Check Revoked Value
	_, pRevoked, dRevoked := obj1.RevokedProperty.Compare(&obj2.RevokedProperty)
	problemsFound += pRevoked
	resultDetails = append(resultDetails, dRevoked...)

	// Check Labels Values
	_, pLabels, dLabels := obj1.LabelsProperty.Compare(&obj2.LabelsProperty)
	problemsFound += pLabels
	resultDetails = append(resultDetails, dLabels...)

	// Check Confidence Value
	_, pConfidences, dConfidences := obj1.ConfidenceProperty.Compare(&obj2.ConfidenceProperty)
	problemsFound += pConfidences
	resultDetails = append(resultDetails, dConfidences...)

	// Check Lang Value
	_, pLangs, dLangs := obj1.LangProperty.Compare(&obj2.LangProperty)
	problemsFound += pLangs
	resultDetails = append(resultDetails, dLangs...)

	// Check External References
	_, pExternalRefereces, dExternalRefereces := obj1.ExternalReferencesProperty.Compare(&obj2.ExternalReferencesProperty)
	problemsFound += pExternalRefereces
	resultDetails = append(resultDetails, dExternalRefereces...)

	// Check Object Marking Refs and Granular Markings
	_, pMarkings, dMarkings := obj1.MarkingProperties.Compare(&obj2.MarkingProperties)
	problemsFound += pMarkings
	resultDetails = append(resultDetails, dMarkings...)

	if problemsFound > 0 {
		return false, problemsFound, resultDetails
	}

	return true, 0, resultDetails
}
