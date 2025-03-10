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
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Dsse represents some metadata about an archived DSSE envelope
type Dsse struct {
	ent.Schema
}

// Fields of the Statement.
func (Dsse) Fields() []ent.Field {
	return []ent.Field{
		field.String("gitoid_sha256").NotEmpty().Unique(),
		field.String("payload_type").NotEmpty(),
	}
}

// Edges of the Statement.
func (Dsse) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("statement", Statement.Type).Unique(),
		edge.To("signatures", Signature.Type),
		edge.To("payload_digests", PayloadDigest.Type),
	}
}

func (Dsse) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.RelayConnection(),
		entgql.QueryField(),
	}
}
