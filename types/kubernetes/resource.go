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
	"context"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/aws-controllers-k8s/pkg/types/resource"
)

// Resource represents a custom resource object in the Kubernetes API
// that corresponds to a resource in an AWS service API.
type Resource interface {
	resource.Resource
	HasConditions
	// IsBeingDeleted returns true if the Kubernetes resource has a non-zero
	// deletion timestamp
	IsBeingDeleted() bool
	// SetObjectMeta sets the ObjectMeta field for the resource
	SetObjectMeta(meta metav1.ObjectMeta)
	// DeepCopy will return a copy of the resource
	DeepCopy() Resource
	// LateInitialize merges any late-initialized fields from the supplied
	// Resource into this Resource.
	//
	// This method will initialize any optional fields which were not provided
	// by the k8s user but were defaulted by the AWS service.  If there are no
	// such fields to be initialized, no changes are made to the Resource.
	//
	// NOTE: This method also adds/updates the ConditionTypeLateInitialized for
	// the Resource.
	LateInitialize(
		context.Context,
		Resource,
	) error
}
