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

package resource

import (
	"github.com/aws-controllers-k8s/pkg/compare"
	"github.com/aws-controllers-k8s/pkg/path/fieldpath"
	"github.com/aws-controllers-k8s/pkg/types/resource/schema"
)

// Resource is a logical representation of an AWS service resource.
type Resource interface {
	// IsValid returns true if the desired state has a possibility of being in
	// sync with the latest observed state, false otherwise. If a user sets a
	// desired field to a value that caused an Update operation to fail,
	// IsValid() would return false for the resource and a call to Errors()
	// would return a non-empty set of errors that placed the resource into an
	// invalid state.
	IsValid() bool
	// IsReady returns true if the resource's state indicates that the resource
	// is "active", "available" or "ready".
	IsReady() bool
	// IsMutable returns true if the resource's state indicates that the
	// resource may be modified.
	IsMutable() bool
	// Errors returns zero or more errors that indicate why the resource may be
	// in an invalid state.
	Errors() []error
	// Identifiers returns an Identifiers which contain all the information
	// needed to identify the resource.
	Identifiers() Identifiers
	// Schema returns a Schema that describes the resource's fields and
	// identifiers
	Schema() schema.Schema
	// Delta returns a Delta object containing the difference between this
	// Resource and another.
	Delta(Resource) *compare.Delta
	// Values returns a map, keyed by stringified field path, of field
	// values.
	Values() map[string]interface{}
	// ValueAt returns the value stored in the Resource for a Field identified by
	// a supplied field path. Note that there is no way to retrieve a single
	// element in a list field or a single key from a map field. Instead, Value
	// returns the entire slice or map for a field identified by the supplied
	// field path.
	ValueAt(fieldpath.Path) interface{}
	// SetAt sets the value of a Resource field at the specified field path.
	SetAt(fieldpath.Path, interface{}) error
}
