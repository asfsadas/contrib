// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/asfsadas/contrib/schemast/internal/mutatetest/ent/user"
	"github.com/asfsadas/contrib/schemast/internal/mutatetest/ent/withmodifiedfield"
)

// WithModifiedField is the model entity for the WithModifiedField schema.
type WithModifiedField struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the WithModifiedFieldQuery when eager-loading is set.
	Edges                     WithModifiedFieldEdges `json:"edges"`
	with_modified_field_owner *int
}

// WithModifiedFieldEdges holds the relations/edges for other nodes in the graph.
type WithModifiedFieldEdges struct {
	// Owner holds the value of the owner edge.
	Owner *User `json:"owner,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// OwnerOrErr returns the Owner value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e WithModifiedFieldEdges) OwnerOrErr() (*User, error) {
	if e.loadedTypes[0] {
		if e.Owner == nil {
			// The edge owner was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: user.Label}
		}
		return e.Owner, nil
	}
	return nil, &NotLoadedError{edge: "owner"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*WithModifiedField) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case withmodifiedfield.FieldID:
			values[i] = new(sql.NullInt64)
		case withmodifiedfield.FieldName:
			values[i] = new(sql.NullString)
		case withmodifiedfield.ForeignKeys[0]: // with_modified_field_owner
			values[i] = new(sql.NullInt64)
		default:
			return nil, fmt.Errorf("unexpected column %q for type WithModifiedField", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the WithModifiedField fields.
func (wmf *WithModifiedField) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case withmodifiedfield.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			wmf.ID = int(value.Int64)
		case withmodifiedfield.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				wmf.Name = value.String
			}
		case withmodifiedfield.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field with_modified_field_owner", value)
			} else if value.Valid {
				wmf.with_modified_field_owner = new(int)
				*wmf.with_modified_field_owner = int(value.Int64)
			}
		}
	}
	return nil
}

// QueryOwner queries the "owner" edge of the WithModifiedField entity.
func (wmf *WithModifiedField) QueryOwner() *UserQuery {
	return (&WithModifiedFieldClient{config: wmf.config}).QueryOwner(wmf)
}

// Update returns a builder for updating this WithModifiedField.
// Note that you need to call WithModifiedField.Unwrap() before calling this method if this WithModifiedField
// was returned from a transaction, and the transaction was committed or rolled back.
func (wmf *WithModifiedField) Update() *WithModifiedFieldUpdateOne {
	return (&WithModifiedFieldClient{config: wmf.config}).UpdateOne(wmf)
}

// Unwrap unwraps the WithModifiedField entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (wmf *WithModifiedField) Unwrap() *WithModifiedField {
	tx, ok := wmf.config.driver.(*txDriver)
	if !ok {
		panic("ent: WithModifiedField is not a transactional entity")
	}
	wmf.config.driver = tx.drv
	return wmf
}

// String implements the fmt.Stringer.
func (wmf *WithModifiedField) String() string {
	var builder strings.Builder
	builder.WriteString("WithModifiedField(")
	builder.WriteString(fmt.Sprintf("id=%v", wmf.ID))
	builder.WriteString(", name=")
	builder.WriteString(wmf.Name)
	builder.WriteByte(')')
	return builder.String()
}

// WithModifiedFields is a parsable slice of WithModifiedField.
type WithModifiedFields []*WithModifiedField

func (wmf WithModifiedFields) config(cfg config) {
	for _i := range wmf {
		wmf[_i].config = cfg
	}
}
