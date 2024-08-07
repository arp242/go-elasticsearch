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
// https://github.com/elastic/elasticsearch-specification/tree/8e91c0692c0235474a0c21bb7e9716a8430e8533

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strconv"
)

// ShingleTokenFilter type.
//
// https://github.com/elastic/elasticsearch-specification/blob/8e91c0692c0235474a0c21bb7e9716a8430e8533/specification/_types/analysis/token_filters.ts#L87-L95
type ShingleTokenFilter struct {
	FillerToken                *string `json:"filler_token,omitempty"`
	MaxShingleSize             string  `json:"max_shingle_size,omitempty"`
	MinShingleSize             string  `json:"min_shingle_size,omitempty"`
	OutputUnigrams             *bool   `json:"output_unigrams,omitempty"`
	OutputUnigramsIfNoShingles *bool   `json:"output_unigrams_if_no_shingles,omitempty"`
	TokenSeparator             *string `json:"token_separator,omitempty"`
	Type                       string  `json:"type,omitempty"`
	Version                    *string `json:"version,omitempty"`
}

func (s *ShingleTokenFilter) UnmarshalJSON(data []byte) error {

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

		case "filler_token":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "FillerToken", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.FillerToken = &o

		case "max_shingle_size":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "MaxShingleSize", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.MaxShingleSize = o

		case "min_shingle_size":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "MinShingleSize", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.MinShingleSize = o

		case "output_unigrams":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "OutputUnigrams", err)
				}
				s.OutputUnigrams = &value
			case bool:
				s.OutputUnigrams = &v
			}

		case "output_unigrams_if_no_shingles":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "OutputUnigramsIfNoShingles", err)
				}
				s.OutputUnigramsIfNoShingles = &value
			case bool:
				s.OutputUnigramsIfNoShingles = &v
			}

		case "token_separator":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "TokenSeparator", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.TokenSeparator = &o

		case "type":
			if err := dec.Decode(&s.Type); err != nil {
				return fmt.Errorf("%s | %w", "Type", err)
			}

		case "version":
			if err := dec.Decode(&s.Version); err != nil {
				return fmt.Errorf("%s | %w", "Version", err)
			}

		}
	}
	return nil
}

// MarshalJSON override marshalling to include literal value
func (s ShingleTokenFilter) MarshalJSON() ([]byte, error) {
	type innerShingleTokenFilter ShingleTokenFilter
	tmp := innerShingleTokenFilter{
		FillerToken:                s.FillerToken,
		MaxShingleSize:             s.MaxShingleSize,
		MinShingleSize:             s.MinShingleSize,
		OutputUnigrams:             s.OutputUnigrams,
		OutputUnigramsIfNoShingles: s.OutputUnigramsIfNoShingles,
		TokenSeparator:             s.TokenSeparator,
		Type:                       s.Type,
		Version:                    s.Version,
	}

	tmp.Type = "shingle"

	return json.Marshal(tmp)
}

// NewShingleTokenFilter returns a ShingleTokenFilter.
func NewShingleTokenFilter() *ShingleTokenFilter {
	r := &ShingleTokenFilter{}

	return r
}
