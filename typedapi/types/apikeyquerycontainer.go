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
// https://github.com/elastic/elasticsearch-specification/tree/48e2d9de9de2911b8cb1cf715e4bc0a2b1f4b827

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
)

// ApiKeyQueryContainer type.
//
// https://github.com/elastic/elasticsearch-specification/blob/48e2d9de9de2911b8cb1cf715e4bc0a2b1f4b827/specification/security/query_api_keys/types.ts#L141-L205
type ApiKeyQueryContainer struct {
	// Bool matches documents matching boolean combinations of other queries.
	Bool *BoolQuery `json:"bool,omitempty"`
	// Exists Returns documents that contain an indexed value for a field.
	Exists *ExistsQuery `json:"exists,omitempty"`
	// Ids Returns documents based on their IDs.
	// This query uses document IDs stored in the `_id` field.
	Ids *IdsQuery `json:"ids,omitempty"`
	// Match Returns documents that match a provided text, number, date or boolean value.
	// The provided text is analyzed before matching.
	Match map[string]MatchQuery `json:"match,omitempty"`
	// MatchAll Matches all documents, giving them all a `_score` of 1.0.
	MatchAll *MatchAllQuery `json:"match_all,omitempty"`
	// Prefix Returns documents that contain a specific prefix in a provided field.
	Prefix map[string]PrefixQuery `json:"prefix,omitempty"`
	// Range Returns documents that contain terms within a provided range.
	Range map[string]RangeQuery `json:"range,omitempty"`
	// SimpleQueryString Returns documents based on a provided query string, using a parser with a
	// limited but fault-tolerant syntax.
	SimpleQueryString *SimpleQueryStringQuery `json:"simple_query_string,omitempty"`
	// Term Returns documents that contain an exact term in a provided field.
	// To return a document, the query term must exactly match the queried field's
	// value, including whitespace and capitalization.
	Term map[string]TermQuery `json:"term,omitempty"`
	// Terms Returns documents that contain one or more exact terms in a provided field.
	// To return a document, one or more terms must exactly match a field value,
	// including whitespace and capitalization.
	Terms *TermsQuery `json:"terms,omitempty"`
	// Wildcard Returns documents that contain terms matching a wildcard pattern.
	Wildcard map[string]WildcardQuery `json:"wildcard,omitempty"`
}

func (s *ApiKeyQueryContainer) UnmarshalJSON(data []byte) error {

	dec := json.NewDecoder(bytes.NewReader(data))

	for {
		t, err := dec.Token()
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			return err
		}

		switch t {

		case "bool":
			if err := dec.Decode(&s.Bool); err != nil {
				return fmt.Errorf("%s | %w", "Bool", err)
			}

		case "exists":
			if err := dec.Decode(&s.Exists); err != nil {
				return fmt.Errorf("%s | %w", "Exists", err)
			}

		case "ids":
			if err := dec.Decode(&s.Ids); err != nil {
				return fmt.Errorf("%s | %w", "Ids", err)
			}

		case "match":
			if s.Match == nil {
				s.Match = make(map[string]MatchQuery, 0)
			}
			if err := dec.Decode(&s.Match); err != nil {
				return fmt.Errorf("%s | %w", "Match", err)
			}

		case "match_all":
			if err := dec.Decode(&s.MatchAll); err != nil {
				return fmt.Errorf("%s | %w", "MatchAll", err)
			}

		case "prefix":
			if s.Prefix == nil {
				s.Prefix = make(map[string]PrefixQuery, 0)
			}
			if err := dec.Decode(&s.Prefix); err != nil {
				return fmt.Errorf("%s | %w", "Prefix", err)
			}

		case "range":
			if s.Range == nil {
				s.Range = make(map[string]RangeQuery, 0)
			}
			messages := make(map[string]json.RawMessage)
			err := dec.Decode(&messages)
			if err != nil {
				return fmt.Errorf("%s | %w", "Range", err)
			}

			untyped := NewUntypedRangeQuery()
			for key, message := range messages {
				err := json.Unmarshal(message, &untyped)
				if err != nil {
					return fmt.Errorf("%s | %w", "Range", err)
				}
				s.Range[key] = untyped
			}

		case "simple_query_string":
			if err := dec.Decode(&s.SimpleQueryString); err != nil {
				return fmt.Errorf("%s | %w", "SimpleQueryString", err)
			}

		case "term":
			if s.Term == nil {
				s.Term = make(map[string]TermQuery, 0)
			}
			if err := dec.Decode(&s.Term); err != nil {
				return fmt.Errorf("%s | %w", "Term", err)
			}

		case "terms":
			if err := dec.Decode(&s.Terms); err != nil {
				return fmt.Errorf("%s | %w", "Terms", err)
			}

		case "wildcard":
			if s.Wildcard == nil {
				s.Wildcard = make(map[string]WildcardQuery, 0)
			}
			if err := dec.Decode(&s.Wildcard); err != nil {
				return fmt.Errorf("%s | %w", "Wildcard", err)
			}

		}
	}
	return nil
}

// NewApiKeyQueryContainer returns a ApiKeyQueryContainer.
func NewApiKeyQueryContainer() *ApiKeyQueryContainer {
	r := &ApiKeyQueryContainer{
		Match:    make(map[string]MatchQuery, 0),
		Prefix:   make(map[string]PrefixQuery, 0),
		Range:    make(map[string]RangeQuery, 0),
		Term:     make(map[string]TermQuery, 0),
		Wildcard: make(map[string]WildcardQuery, 0),
	}

	return r
}