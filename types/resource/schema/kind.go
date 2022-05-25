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

// Kind describes the type/kind of a Resource. It is similar to a GroupKind in
// Kubernetes.
type Kind interface {
	// Service returns the name of the AWS service this resource is associated
	// with. The string returned matches the service package name in
	// aws-sdk-go.
	Service() string
	// Name returns the camel-cased name of the resource (i.e. the Kind, in Kubernetes
	// speak).
	//
	// Note that the combination of Service and Name is a unique identifier for
	// this type of Resource.
	Name() string
	// PluralName returns camel-cased name of the pluralized resource.
	//
	// Note that the combination of Service and PluralName is a unique identifier for
	// this type of Resource.
	PluralName() string
}
