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
	"entgo.io/ent/schema/field"
	"github.com/asfsadas/contrib/entgql/internal/todo/ent/schema"

	"github.com/google/uuid"
)

// Group defines the group type schema.
type Group struct {
	ent.Schema
}

// Mixin returns group mixed-in schema.
func (Group) Mixin() []ent.Mixin {
	return []ent.Mixin{
		schema.Group{},
	}
}

// Fields returns group fields.
func (Group) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New),
	}
}
