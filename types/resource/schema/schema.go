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

package schema

import "github.com/aws-controllers-k8s/pkg/path/fieldpath"

// Schema contains methods that returns information about a resource's schema.
type Schema interface {
	Kind
	// Field returns a Field at a given field path, or nil if there is no Field
	// at that path.
	Field(fieldpath.Path) Field
	// Fields returns a map, keyed by field path string, of Fields that
	// describe the resource's member fields.
	Fields() map[string]Field
	// Identifiers returns information about a resource's identifying fields
	// and those fields' values.
	Identifiers() Identifiers
}
