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

// FieldType indicates the underlying Go builtin type for a Field.
type FieldType int

const (
	// FieldTypeNil is a placeholder for an nil type
	FieldTypeNil FieldType = iota
	// FieldTypeInt is for integer types
	FieldTypeInt
	// FieldTypeFloat is for float/double types
	FieldTypeFloat
	// FieldTypeTime is for time.Time types
	FieldTypeTime
	// FieldTypeString is for string types
	FieldTypeString
	// FieldTypeList is for list/slice types
	FieldTypeList
	// FieldTypeMap is for map[string]interface{} types
	FieldTypeMap
	// FieldTypeStruct is for (nested) struct types
	FieldTypeStruct
)

// Field has methods that return information about a field in a resource.
type Field interface {
	// Name returns the normalized, camel-cased name of the field
	Name() string
	// Type returns the underlying type of the field.
	Type() FieldType
	// ElementType returns the type of the list's elements.
	//
	// If Type is FieldTypeList, the ElementType() method is guaranteed to
	// return the type of the list element. If Type is not FieldTypeList,
	// ElementType is guaranteed to be FieldTypeNil.
	ElementType() FieldType
	// ValueType returns the type of the map's values.
	//
	// If Type is FieldTypeMap, the ValueType() method is guaranteed to return
	// the type of the map values. If Type is not FieldTypeMap, ValueType will
	// always return FieldTypeNil
	ValueType() FieldType
	// KeyType returns the type of the map's keys.
	//
	// If Type is FieldTypeMap, the KeyType() method is guaranteed to return
	// the type of the map keys. If Type is not FieldTypeMap, KeyType will
	// always return FieldTypeNil
	KeyType() FieldType
	// MemberFields returns a map, keyed by member field name, of nested Fields
	// when this Field has a Type of FieldTypeStruct. Returns nil when Type is
	// not FieldTypeStruct.
	MemberFields() map[string]Field
	// IsReadOnly returns true if the field is not settable by the user
	IsReadOnly() bool
	// IsImmutable returns true if the field cannot be changed once set
	IsImmutable() bool
	// IsLateInitialized returns true if the field is "late initialized"
	// with a service-side default value
	IsLateInitialized() bool
	// IsSecret returns true if the field contains secret information
	IsSecret() bool
	// References returns the Kind for a referred type if the field contains a
	// reference to another resource, or nil otherwise.
	//
	// For example, consider a Resource `rds.DBInstance` with a field
	// `Subnets`. This field contains EC2 VPC Subnet identifiers. The Type() of
	// this field would be FieldTypeList. The ElementType() of this field would
	// be FieldTypeString. The References() of this field would return a Kind
	// containing "ec2.Subnet".
	References() Kind
}
