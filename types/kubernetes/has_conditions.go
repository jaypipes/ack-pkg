// Copyright Amazon.com Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"). You may
// not use this file except in compliance with the License. A copy of the
// License is located at
//
//     http://aws.amazon.com/apache2.0/
//
// or in the "license" file accompanying this file. This file is distributed
// on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
// express or implied. See the License for the specific language governing
// permissions and limitations under the License.

package kubernetes

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// HasConditions describes a thing that can set and retrieve Condition
// objects.
type HasConditions interface {
	// Conditions returns a collection of Conditionss
	Conditions() []*metav1.Condition
	// ReplaceConditions replaces the resource's set of Condition structs with
	// the supplied slice of Conditions.
	ReplaceConditions([]*metav1.Condition)
}
