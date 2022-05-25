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

// Identifiers returns information about a resource's identifying fields and
// those fields' values.
type Identifiers interface {
	// Fields returns an ordered slice of slices of Fields that contain values
	// that can be used to uniquely identify the resource.
	//
	// The returned slice is sorted by efficiency of fetch operation. For
	// example, assume an S3 Bucket resource. It may be identified by ARN or by
	// Bucket Location in addition to the Bucket's Name field, which happens to
	// be globally-unique, so the returned value from Fields() would be:
	//
	// [][]Field{
	//   {Field("ARN")},
	//   {Field("Location")},
	//   {Field("Name")},
	// }
	//
	// For something like an RDS DBInstance, the DB instance's ARN is a
	// globally unique identifier, but the instance's DBInstanceIdentifier
	// field is only unique when coupled with the customer's AWS Account and
	// Region. So, the returned value from Fields() would be:
	//
	// [][]Field{
	//   {Field("ARN")},
	//   {Field("DBInstanceIdentifier"), Field("AWSAccountID"), Field("AWSRegion")},
	// }
	Fields() [][]Field
}
