// Copyright 2022 The Archivist Contributors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// PayloadDigest represents the digest of the payload of a DSSE envelope
type PayloadDigest struct {
	ent.Schema
}

// Fields of the Digest.
func (PayloadDigest) Fields() []ent.Field {
	return []ent.Field{
		field.String("algorithm").NotEmpty(),
		field.String("value").NotEmpty(),
	}
}

// Edges of the Digest.
func (PayloadDigest) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("dsse", Dsse.Type).
			Ref("payload_digests").
			Unique(),
	}
}

func (PayloadDigest) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("value"),
	}
}
