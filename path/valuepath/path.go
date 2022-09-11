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

// Package valuepath allows calling code to specify a path to a singular value
// within a structure's fields.
//
// A `valuepath.Path` is a JSONPath-like structure that allows specifying a
// single value within a structure's fields at an infinite depth.
//
// For the following examples, assume a struct with a number of fields of
// various scalar and non-scalar types:
//
// type Address struct {
//   Street string
//   City string
//   Postcode string
//   Country string
// }
//
// type Publisher struct {
//   Name string
//   Addresses []Address
// }
//
// type Book struct {
//   Title string
//   ChapterTitles []string
// }
//
// type Author struct {
//   FirstName string
//   LastName string
//   Publisher Publisher
//   Books map[string]Book // Books keyed by book title
// }
//
// To specify the value of a scalar top-level field such as the author's last
// name, simply use the name of the top-level field:
//
// p := valuepath.FromString("LastName")
//
// To specify the value of a scalar field such as an author's publisher's name,
// use dotted-notation, like so:
//
// p := valuepath.FromString("Publisher.Name")
//
// To specify the value of a list field (where the value is the entire contents
// of the list), similarly use dotted-notation. For instance, if we wanted to
// create a value path that pointed to the entire list of an author's
// publisher's addresses:
//
// p := valuepath.FromString("Publisher.Addresses")
//
// Likewise, to specify the value of a map field (where the value is the entire
// contents of the map), also use dotted-notation. For instance, if we wanted
// to create a value path that pointed to the entire map of an author's books:
//
// p := valuepath.FromString("Publisher.Books")
//
// To specify the value of a list field at a numeric index, such as the city of
// the first address of the author's publisher, specify the numeric index
// within square brackets:
//
// p := valuepath.FromString("Publisher.Addresses[0].City")
//
// To specify the value of a map field at a specific key, such as an author's
// book with the title "Gone with the Wind", use a single-quoteed escaped
// string between square brackets:
//
// p := valuepath.FromString("Books['Gone With the Wind']")
//
// NOTE: key-matching is case-sensitive for map field value keys
package valuepath

import (
	"encoding/json"
	"errors"
	"strconv"
	"strings"
	"unicode"
)

var (
	// ErrInvalidPath indicates a supplied string path was invalid
	ErrInvalidPath = errors.New("invalid path")
)

// partType indicates the type of the PathPart
type partType int

const (
	partTypeField   partType = iota
	partTypeElement          // An element in a list field
	partTypeKey              // A key in a map field
)

// pathPart provides a "route" to a particular value within a part of a Path.
// For parts of a Path that represent a non-scalar field, the PathPart has the
// index or key (for list and map types, respectively) to locate the value.
type pathPart struct {
	partType partType
	// index is the element's index within a list field
	index int
	// fieldOrKey is the name of the field or the key in a map field
	fieldOrKey string
}

// Path provides a JSONPath-like struct and "route" to a particular field's
// value within a resource. Path implements json.Marshaler interface.
type Path struct {
	parts []pathPart
}

// String returns the dotted-notation representation of the Path
func (p *Path) String() string {
	return strings.Join(p.parts, ".")
}

// MarshalJSON returns the JSON encoding of a Path object.
func (p *Path) MarshalJSON() ([]byte, error) {
	// Since json.Marshal doesn't encode unexported struct fields we have to
	// copy the Path instance into a new struct object with exported fields.
	// See https://github.com/aws-controllers-k8s/community/issues/772
	return json.Marshal(
		struct {
			Parts []string
		}{
			p.parts,
		},
	)
}

// Pop removes the last part from the Path and returns it.
func (p *Path) Pop() (part string) {
	if len(p.parts) > 0 {
		part = p.parts[len(p.parts)-1]
		p.parts = p.parts[:len(p.parts)-1]
	}
	return part
}

// At returns the part of the Path at the supplied index, or empty string if
// index exceeds boundary.
func (p *Path) At(index int) string {
	if index < 0 || len(p.parts) == 0 || index > len(p.parts)-1 {
		return ""
	}
	return p.parts[index]
}

// Front returns the first part of the Path or empty string if the Path has no
// parts.
func (p *Path) Front() string {
	if len(p.parts) == 0 {
		return ""
	}
	return p.parts[0]
}

// PopFront removes the first part of the Path and returns it.
func (p *Path) PopFront() (part string) {
	if len(p.parts) > 0 {
		part = p.parts[0]
		p.parts = p.parts[1:]
	}
	return part
}

// Back returns the last part of the Path or empty string if the Path has no
// parts.
func (p *Path) Back() string {
	if len(p.parts) == 0 {
		return ""
	}
	return p.parts[len(p.parts)-1]
}

// PushBack adds a new part to the end of the Path.
func (p *Path) PushBack(part string) {
	p.parts = append(p.parts, part)
}

// Copy returns a new Path that is a copy of this Path
func (p *Path) Copy() *Path {
	return &Path{p.parts}
}

// CopyAt returns a new Path that is a copy of this Path up to the supplied
// index.
//
// e.g. given Path $A containing "X.Y", $A.CopyAt(0) would return a new Path
// containing just "X". $A.CopyAt(1) would return a new Path containing "X.Y".
func (p *Path) CopyAt(index int) *Path {
	if index < 0 || len(p.parts) == 0 || index > len(p.parts)-1 {
		return nil
	}
	return &Path{p.parts[0 : index+1]}
}

// Empty returns true if there are no parts to the Path
func (p *Path) Empty() bool {
	return len(p.parts) == 0
}

// Size returns the Path number of parts
func (p *Path) Size() int {
	return len(p.parts)
}

// HasPrefix returns true if the supplied string, delimited on ".", matches
// p.parts up to the length of the supplied string.
// e.g. if the Path p represents "A.B":
//  subject "A" -> true
//  subject "A.B" -> true
//  subject "A.B.C" -> false
//  subject "B" -> false
//  subject "A.C" -> false
func (p *Path) HasPrefix(subject string) bool {
	subjectSplit := strings.Split(subject, ".")

	if len(subjectSplit) > len(p.parts) {
		return false
	}

	for i, s := range subjectSplit {
		if p.parts[i] != s {
			return false
		}
	}

	return true
}

// HasPrefixFold is the same as HasPrefix but uses case-insensitive comparisons
func (p *Path) HasPrefixFold(subject string) bool {
	subjectSplit := strings.Split(subject, ".")

	if len(subjectSplit) > len(p.parts) {
		return false
	}

	for i, s := range subjectSplit {
		if !strings.EqualFold(p.parts[i], s) {
			return false
		}
	}

	return true
}

// FromString returns a new Path from a valuepath dotted-notation-like string,
// e.g.  "Author.Publishers.Addresses[0].City".
//
// If the supplied path string is invalid, returns `ErrInvalidPath`
func FromString(subject string) (*Path, error) {
	// For the typical case of simple dotted notation like
	// "Author.Publisher.Name", we can just check if there are any square
	// brackets in the string and if not, create path parts indicating whole
	// field values.
	parts := []pathPart{}
	if !strings.IndexAny(subject, "[]") {
		fields := strings.Split(subject, ".")
		for _, field := range fields {
			parts = append(parts, pathPart{
				pathType:   pathTypeField,
				fieldOrKey: field,
			})
		}
		return &Path{parts}, nil
	}
	subjectLen := len(subject)
	var cursor int
	var idxSpecial int
	var leftBracketPos int = -1
	for {
		idxSpecial = strings.IndexAny(subject[cursor:], "[.")
		if idxSpecial < 0 {
			// no more special chars, so if there is anything left after the
			// cursor it's a field name part
			if cursor < subjectLen {
				parts = append(parts, pathPart{
					pathType:   pathTypeField,
					fieldOrKey: subject[cursor:],
				})
			}
			break
		}
		special := subject[idxSpecial]
		switch special {
		case '[':
			if leftBracketPos > 0 {
				// we got two left brackets before a right bracket
				return nil, ErrInvalidPath
			}
			cursor = idxSpecial + 1
			// we expect either a number or a single quote directly after the
			// left bracket. If not, it's not a valid path
			if cursor >= subjectLen {
				// left bracket is last character in path, which is invalid
				return nil, ErrInvalidPath
			}
			nextChar := subject[cursor]
			if nextChar == '\'' {
				// we need to find the corresponding right bracket, check for
				// the corresponding single quote and then use the characters
				// in between as our map field key
				if leftBracketPos < 0 || leftBracketPos > idxSpecial {
					return nil, ErrInvalidPath
				}
				mapKey := subject[leftBracketPos:idxSpecial]
				parts = append(parts, pathPart{
					partType:   partTypeKey,
					fieldOrKey: mapKey,
				})
				leftBracketPos = -1
				cursor = idxSpecial + 1
			} else if unicode.IsDigit(nextChar) {
				idx, offset, err := findElementIndexWithClosingOffset(
					subject[cursor:],
				)
				if err != nil {
					return nil, ErrInvalidPath
				}
				parts = append(parts, pathPart{
					partType: partTypeElement,
					index:    idx,
				})
				cursor += offset
			} else {
				return nil, ErrInvalidPath
			}
		case '.':
			if leftBracketPos > 0 {
				//
			}
		}
	}
	return &Path{parts}, nil
}

// findElementIndexWithClosingOffset returns an integer that is the index of a
// referred element value in a partial string path along with the index of the
// closing right bracket in the string.
//
// For example, given the string "2089].Name", the function will return the
// integer 2089 and 6, since the right enclosing bracket is at the 6th position
// in the 0-based string array.
func findElementIndexWithClosingOffset(subject string) (int, int, error) {
	var cursor int
	subjectLen := len(subject)
	nextChar := subject[cursor]
	// consume all digits up to right square bracket, erroring out
	// if a non-digit char is found.
	idxStr := string(nextChar)
	cursor++
	for {
		if cursor >= subjectLen {
			// no right bracket after a left bracket with following
			// digits...
			return nil, ErrInvalidPath
		}
		nextChar = subject[cursor]
		if unicode.IsDigit(nextChar) {
			idxStr += nextChar
		} else if nextChar == ']' {
			idx, err := strconv.Atoi(idxStr)
			if err != nil {
				return -1. - 1, ErrInvalidPath
			}
			return idx, cursor, nil
		} else {
			return -1, -1, ErrInvalidPath
		}
		cursor++
	}
	return -1, -1, ErrInvalidPath
}
