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

package manager

import "github.com/aws-controllers-k8s/pkg/types/resource/schema"

// Factory returns a Manager that can be used to manage AWS resources for a
// particular Resource Kind, AWS account and AWS region.
type Factory interface {
	// Manager returns a Manager that manages AWS resources on behalf of a
	// particular AWS account and in a specific AWS region
	Manager(
		schema.Kind,
		string, // AWS Account ID
		string, // AWS Region
	) (Manager, error)
}
