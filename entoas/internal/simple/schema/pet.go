// Copyright 2019-present Facebook
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
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/asfsadas/contrib/entoas"
)

// Pet holds the schema definition for the Pet entity.
type Pet struct {
	ent.Schema
}

// Fields of the Pet.
func (Pet) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id").
			Annotations(entoas.Groups("test:view")),
		field.String("name").
			Annotations(
				entoas.Groups("pet"),
				entoas.Example("Kuro"),
			),
		field.JSON("nicknames", []string{}).
			Optional().
			Annotations(entoas.Groups("pet")),
		field.Int("age").
			Optional().
			Annotations(
				entoas.Groups("pet"),
				entoas.Example(1),
			),
	}
}

// Edges of the Pet.
func (Pet) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("categories", Category.Type).
			Ref("pets"),
		edge.From("owner", User.Type).
			Ref("pets").
			Unique().
			Annotations(
				entoas.Groups("pet:list", "pet:read", "test:edge", "test:view"),
			),
		edge.To("friends", Pet.Type),
	}
}

// Annotations of the Pet.
func (Pet) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entoas.ReadOperation(
			entoas.OperationGroups("pet", "pet:read"),
		),
	}
}
