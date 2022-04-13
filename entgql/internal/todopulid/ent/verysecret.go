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
//
// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"github.com/asfsadas/contrib/entgql/internal/todopulid/ent/schema/pulid"
	"github.com/asfsadas/contrib/entgql/internal/todopulid/ent/verysecret"
	"entgo.io/ent/dialect/sql"
)

// VerySecret is the model entity for the VerySecret schema.
type VerySecret struct {
	config `json:"-"`
	// ID of the ent.
	ID pulid.ID `json:"id,omitempty"`
	// Password holds the value of the "password" field.
	Password string `json:"password,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*VerySecret) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case verysecret.FieldID:
			values[i] = new(pulid.ID)
		case verysecret.FieldPassword:
			values[i] = new(sql.NullString)
		default:
			return nil, fmt.Errorf("unexpected column %q for type VerySecret", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the VerySecret fields.
func (vs *VerySecret) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case verysecret.FieldID:
			if value, ok := values[i].(*pulid.ID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				vs.ID = *value
			}
		case verysecret.FieldPassword:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field password", values[i])
			} else if value.Valid {
				vs.Password = value.String
			}
		}
	}
	return nil
}

// Update returns a builder for updating this VerySecret.
// Note that you need to call VerySecret.Unwrap() before calling this method if this VerySecret
// was returned from a transaction, and the transaction was committed or rolled back.
func (vs *VerySecret) Update() *VerySecretUpdateOne {
	return (&VerySecretClient{config: vs.config}).UpdateOne(vs)
}

// Unwrap unwraps the VerySecret entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (vs *VerySecret) Unwrap() *VerySecret {
	tx, ok := vs.config.driver.(*txDriver)
	if !ok {
		panic("ent: VerySecret is not a transactional entity")
	}
	vs.config.driver = tx.drv
	return vs
}

// String implements the fmt.Stringer.
func (vs *VerySecret) String() string {
	var builder strings.Builder
	builder.WriteString("VerySecret(")
	builder.WriteString(fmt.Sprintf("id=%v", vs.ID))
	builder.WriteString(", password=")
	builder.WriteString(vs.Password)
	builder.WriteByte(')')
	return builder.String()
}

// VerySecrets is a parsable slice of VerySecret.
type VerySecrets []*VerySecret

func (vs VerySecrets) config(cfg config) {
	for _i := range vs {
		vs[_i].config = cfg
	}
}
