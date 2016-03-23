// Copyright 2015 go-swagger maintainers
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package generator provides the code generation library for go-swagger
//
// The general idea is that you should rarely see interface{} in the generated code.
// You get a complete representation of a swagger document in somewhat idiomatic go.
//
// To do so there is set of mapping patterns that are applied to a spec to go types:
//
//    definition of primitive => type alias/name
//    definition of array => type alias/name
//    definition of map => type alias/name
//    definition of object with properties => struct
//    definition of ref => type alias/name
//    object with only additional properties => map[string]T
//    object with additional properties and properties => custom serializer
//    schema with schema array in items => tuple (struct with properties, custom serializer)
//    schema with all of => struct
//      * all of schema with ref => embedded value
//      * all of schema with properties => properties are included in struct
//      * adding an all of schema with just "x-isnullable": true or "x-nullable": true turns
//      the schema into a pointer when there are only other extension properties provided
//
// JSONSchema and by extension swagger allow for items that have a fixed size array
// with schema's describing the items at each index. This can be combined with additional items
// to form some kind of tuple with varargs.
// To map this to go it creates a struct that has fixed names and a custom json serializer.
//
// The code that is generated also gets the doc comments that are used by the scanner
// to generate a spec from go code. So that after generation you should be able to reverse
// generate a spec from the code that was generated by your spec.
//
// It should be equivalent to the original spec but might miss some default values and examples.
package generator
