// Copyright 2015-2019 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

package baseobject

import (
	"github.com/freetaxii/libstix2/objects/properties"
)

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

/*
Compare - This function will compare the common properties from two objects
(object 1 and object 2) to make sure they are the same. This function will
return an integer that tracks the number of problems and a slice of strings that
contain the detailed results, whether good or bad.
*/
func Compare(obj1, obj2 *CommonObjectProperties) (bool, int, []string) {
	problemsFound := 0
	resultDetails := make([]string, 0)

	// Check Type Value
	_, pTypes, dTypes := properties.CompareTypeProperties(&obj1.TypeProperty, &obj2.TypeProperty)
	problemsFound += pTypes
	resultDetails = append(resultDetails, dTypes...)

	// Check Spec Version Value
	_, pSpecVersions, dSpecVersions := properties.CompareSpecVersionProperties(&obj1.SpecVersionProperty, &obj2.SpecVersionProperty)
	problemsFound += pSpecVersions
	resultDetails = append(resultDetails, dSpecVersions...)

	// Check ID Value
	_, pIDs, dIDs := properties.CompareIDProperties(&obj1.IDProperty, &obj2.IDProperty)
	problemsFound += pIDs
	resultDetails = append(resultDetails, dIDs...)

	// Check Created By Ref Value
	_, pCreatedByRefs, dCreatedByRefs := properties.CompareCreatedByRefProperties(&obj1.CreatedByRefProperty, &obj2.CreatedByRefProperty)
	problemsFound += pCreatedByRefs
	resultDetails = append(resultDetails, dCreatedByRefs...)

	// Check Created and Modified Values
	_, pCreatedModified, dCreatedModified := properties.CompareCreatedModifiedProperties(&obj1.CreatedModifiedProperty, &obj2.CreatedModifiedProperty)
	problemsFound += pCreatedModified
	resultDetails = append(resultDetails, dCreatedModified...)

	// Check Revoked Value
	_, pRevoked, dRevoked := properties.CompareRevokedProperties(&obj1.RevokedProperty, &obj2.RevokedProperty)
	problemsFound += pRevoked
	resultDetails = append(resultDetails, dRevoked...)

	// Check Labels Values
	_, pLabels, dLabels := properties.CompareLabelsProperties(&obj1.LabelsProperty, &obj2.LabelsProperty)
	problemsFound += pLabels
	resultDetails = append(resultDetails, dLabels...)

	// Check Confidence Value
	_, pConfidences, dConfidences := properties.CompareConfidenceProperties(&obj1.ConfidenceProperty, &obj2.ConfidenceProperty)
	problemsFound += pConfidences
	resultDetails = append(resultDetails, dConfidences...)

	// Check Lang Value
	_, pLangs, dLangs := properties.CompareLangProperties(&obj1.LangProperty, &obj2.LangProperty)
	problemsFound += pLangs
	resultDetails = append(resultDetails, dLangs...)

	// Check External References
	_, pExternalRefereces, dExternalRefereces := properties.CompareExternalReferencesProperties(&obj1.ExternalReferencesProperty, &obj2.ExternalReferencesProperty)
	problemsFound += pExternalRefereces
	resultDetails = append(resultDetails, dExternalRefereces...)

	// Check Object Marking Refs and Granular Markings
	_, pMarkings, dMarkings := properties.CompareMarkingProperties(&obj1.MarkingProperty, &obj2.MarkingProperty)
	problemsFound += pMarkings
	resultDetails = append(resultDetails, dMarkings...)

	if problemsFound > 0 {
		return false, problemsFound, resultDetails
	}

	return true, 0, resultDetails
}
