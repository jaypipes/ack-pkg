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

import "github.com/aws-controllers-k8s/pkg/types/resource/schema"

// Identifiers contains all the information callers would need to identify a
// particular resource
type Identifiers interface {
	// ValuesIter returns a slice, ordered by efficiency of fetch operation,
	// of maps, keyed by identifying field, of identifying field values.
	ValuesIter() []map[schema.Field]string
	// ValuesBy returns a slice of strings representing the values of
	// supplied identifying Fields
	ValuesBy(...schema.Field) []string
	// ARN returns the resource's Amazon Resource Name (ARN) as a string
	// pointer or nil if the resource has no ARN
	ARN() *string
	// AWSAccountID returns the resource's AWS Account ID. This is guaranteed
	// to be a non-nil value.
	AWSAccountID() string
	// AWSRegion returns the AWS Region that the resource is in (or will be in
	// if the resource has yet to be created)
	AWSRegion() string
}
