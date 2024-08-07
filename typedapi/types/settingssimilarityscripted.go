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
	"encoding/json"
)

// SettingsSimilarityScripted type.
//
// https://github.com/elastic/elasticsearch-specification/blob/8e91c0692c0235474a0c21bb7e9716a8430e8533/specification/indices/_types/IndexSettings.ts#L224-L228
type SettingsSimilarityScripted struct {
	Script       Script  `json:"script"`
	Type         string  `json:"type,omitempty"`
	WeightScript *Script `json:"weight_script,omitempty"`
}

// MarshalJSON override marshalling to include literal value
func (s SettingsSimilarityScripted) MarshalJSON() ([]byte, error) {
	type innerSettingsSimilarityScripted SettingsSimilarityScripted
	tmp := innerSettingsSimilarityScripted{
		Script:       s.Script,
		Type:         s.Type,
		WeightScript: s.WeightScript,
	}

	tmp.Type = "scripted"

	return json.Marshal(tmp)
}

// NewSettingsSimilarityScripted returns a SettingsSimilarityScripted.
func NewSettingsSimilarityScripted() *SettingsSimilarityScripted {
	r := &SettingsSimilarityScripted{}

	return r
}
