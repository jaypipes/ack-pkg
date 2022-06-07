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

	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/aws-controllers-k8s/pkg/types/manager"
)

// Mapper provides methods that marry the Kubernetes client semantics and
// Kubernetes Resource Model with the AWS service-level CRUD interface.
type Mapper interface {
	manager.Manager
	// EmptyObject returns the empty Kubernetes controller-runtime.Object
	// representation for the Kind of Resource mapped with this Mapper
	EmptyObject() client.Object
	// ToObject returns the Kubernetes controller-runtime.Object representation
	// of the supplied Resource
	ToObject(Resource) client.Object
	// FromObject returns the Resource from the supplied Kubernetes
	// controller-runtime.Object
	FromObject(client.Object) Resource
	// Resource returns the Resource matching the supplied Kubernetes Namespace
	// and Name
	Resource(
		context.Context,
		client.Reader,
		types.NamespacedName,
	) (Resource, error)
	// ResolveReferences finds if there are any Reference field(s) present
	// inside the Resource and attempts to resolve those reference field(s)
	// into target field(s).
	//
	// It returns an error if the Resource's reference field(s) cannot be
	// resolved.
	//
	// NOTE: This method also adds/updates the ConditionTypeReferencesResolved
	// for the Resource.
	ResolveReferences(
		context.Context,
		client.Reader,
		Resource, // resource to resolve references for
	) error
}
