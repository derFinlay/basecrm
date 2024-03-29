// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/derfinlay/basecrm/ent/user"
	"github.com/derfinlay/basecrm/ent/usersession"
)

// UserSession is the model entity for the UserSession schema.
type UserSession struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Token holds the value of the "token" field.
	Token string `json:"token,omitempty"`
	// Active holds the value of the "active" field.
	Active bool `json:"active,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the UserSessionQuery when eager-loading is set.
	Edges             UserSessionEdges `json:"edges"`
	user_session_user *int
	selectValues      sql.SelectValues
}

// UserSessionEdges holds the relations/edges for other nodes in the graph.
type UserSessionEdges struct {
	// User holds the value of the user edge.
	User *User `json:"user,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// UserOrErr returns the User value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e UserSessionEdges) UserOrErr() (*User, error) {
	if e.loadedTypes[0] {
		if e.User == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: user.Label}
		}
		return e.User, nil
	}
	return nil, &NotLoadedError{edge: "user"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*UserSession) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case usersession.FieldActive:
			values[i] = new(sql.NullBool)
		case usersession.FieldID:
			values[i] = new(sql.NullInt64)
		case usersession.FieldToken:
			values[i] = new(sql.NullString)
		case usersession.FieldCreatedAt, usersession.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		case usersession.ForeignKeys[0]: // user_session_user
			values[i] = new(sql.NullInt64)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the UserSession fields.
func (us *UserSession) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case usersession.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			us.ID = int(value.Int64)
		case usersession.FieldToken:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field token", values[i])
			} else if value.Valid {
				us.Token = value.String
			}
		case usersession.FieldActive:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field active", values[i])
			} else if value.Valid {
				us.Active = value.Bool
			}
		case usersession.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				us.CreatedAt = value.Time
			}
		case usersession.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				us.UpdatedAt = value.Time
			}
		case usersession.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field user_session_user", value)
			} else if value.Valid {
				us.user_session_user = new(int)
				*us.user_session_user = int(value.Int64)
			}
		default:
			us.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the UserSession.
// This includes values selected through modifiers, order, etc.
func (us *UserSession) Value(name string) (ent.Value, error) {
	return us.selectValues.Get(name)
}

// QueryUser queries the "user" edge of the UserSession entity.
func (us *UserSession) QueryUser() *UserQuery {
	return NewUserSessionClient(us.config).QueryUser(us)
}

// Update returns a builder for updating this UserSession.
// Note that you need to call UserSession.Unwrap() before calling this method if this UserSession
// was returned from a transaction, and the transaction was committed or rolled back.
func (us *UserSession) Update() *UserSessionUpdateOne {
	return NewUserSessionClient(us.config).UpdateOne(us)
}

// Unwrap unwraps the UserSession entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (us *UserSession) Unwrap() *UserSession {
	_tx, ok := us.config.driver.(*txDriver)
	if !ok {
		panic("ent: UserSession is not a transactional entity")
	}
	us.config.driver = _tx.drv
	return us
}

// String implements the fmt.Stringer.
func (us *UserSession) String() string {
	var builder strings.Builder
	builder.WriteString("UserSession(")
	builder.WriteString(fmt.Sprintf("id=%v, ", us.ID))
	builder.WriteString("token=")
	builder.WriteString(us.Token)
	builder.WriteString(", ")
	builder.WriteString("active=")
	builder.WriteString(fmt.Sprintf("%v", us.Active))
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(us.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(us.UpdatedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// UserSessions is a parsable slice of UserSession.
type UserSessions []*UserSession
