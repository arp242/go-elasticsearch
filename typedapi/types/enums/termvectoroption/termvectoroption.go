// Licensed to Elasticsearch B.V. under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Elasticsearch B.V. licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

// Code generated from the elasticsearch-specification DO NOT EDIT.
// https://github.com/elastic/elasticsearch-specification/tree/1b56d7e58f5c59f05d1641c6d6a8117c5e01d741

// Package termvectoroption
// https://github.com/elastic/elasticsearch-specification/blob/1b56d7e58f5c59f05d1641c6d6a8117c5e01d741/specification/_types/mapping/TermVectorOption.ts#L20-L28
package termvectoroption

import "strings"

type TermVectorOption struct {
	name string
}

var (
	No = TermVectorOption{"no"}

	Yes = TermVectorOption{"yes"}

	Withoffsets = TermVectorOption{"with_offsets"}

	Withpositions = TermVectorOption{"with_positions"}

	Withpositionsoffsets = TermVectorOption{"with_positions_offsets"}

	Withpositionsoffsetspayloads = TermVectorOption{"with_positions_offsets_payloads"}

	Withpositionspayloads = TermVectorOption{"with_positions_payloads"}
)

func (t TermVectorOption) MarshalText() (text []byte, err error) {
	return []byte(t.String()), nil
}

func (t *TermVectorOption) UnmarshalText(text []byte) error {
	switch strings.ToLower(string(text)) {

	case "no":
		*t = No
	case "yes":
		*t = Yes
	case "with_offsets":
		*t = Withoffsets
	case "with_positions":
		*t = Withpositions
	case "with_positions_offsets":
		*t = Withpositionsoffsets
	case "with_positions_offsets_payloads":
		*t = Withpositionsoffsetspayloads
	case "with_positions_payloads":
		*t = Withpositionspayloads
	default:
		*t = TermVectorOption{string(text)}
	}

	return nil
}

func (t TermVectorOption) String() string {
	return t.name
}