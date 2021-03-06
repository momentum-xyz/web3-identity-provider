// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/OdysseyMomentumExperience/web3-identity-provider/ent/web3user"
	"github.com/google/uuid"
)

// Web3User is the model entity for the Web3User schema.
type Web3User struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// UUID holds the value of the "uuid" field.
	UUID uuid.UUID `json:"uuid,omitempty"`
	// Address holds the value of the "address" field.
	Address string `json:"address,omitempty"`
	// WalletType holds the value of the "wallet_type" field.
	WalletType web3user.WalletType `json:"wallet_type,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the Web3UserQuery when eager-loading is set.
	Edges Web3UserEdges `json:"edges"`
}

// Web3UserEdges holds the relations/edges for other nodes in the graph.
type Web3UserEdges struct {
	// Web3Challenges holds the value of the web3_challenges edge.
	Web3Challenges []*Web3Challenge `json:"web3_challenges,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// Web3ChallengesOrErr returns the Web3Challenges value or an error if the edge
// was not loaded in eager-loading.
func (e Web3UserEdges) Web3ChallengesOrErr() ([]*Web3Challenge, error) {
	if e.loadedTypes[0] {
		return e.Web3Challenges, nil
	}
	return nil, &NotLoadedError{edge: "web3_challenges"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Web3User) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case web3user.FieldID:
			values[i] = new(sql.NullInt64)
		case web3user.FieldAddress, web3user.FieldWalletType:
			values[i] = new(sql.NullString)
		case web3user.FieldUUID:
			values[i] = new(uuid.UUID)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Web3User", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Web3User fields.
func (w *Web3User) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case web3user.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			w.ID = int(value.Int64)
		case web3user.FieldUUID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field uuid", values[i])
			} else if value != nil {
				w.UUID = *value
			}
		case web3user.FieldAddress:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field address", values[i])
			} else if value.Valid {
				w.Address = value.String
			}
		case web3user.FieldWalletType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field wallet_type", values[i])
			} else if value.Valid {
				w.WalletType = web3user.WalletType(value.String)
			}
		}
	}
	return nil
}

// QueryWeb3Challenges queries the "web3_challenges" edge of the Web3User entity.
func (w *Web3User) QueryWeb3Challenges() *Web3ChallengeQuery {
	return (&Web3UserClient{config: w.config}).QueryWeb3Challenges(w)
}

// Update returns a builder for updating this Web3User.
// Note that you need to call Web3User.Unwrap() before calling this method if this Web3User
// was returned from a transaction, and the transaction was committed or rolled back.
func (w *Web3User) Update() *Web3UserUpdateOne {
	return (&Web3UserClient{config: w.config}).UpdateOne(w)
}

// Unwrap unwraps the Web3User entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (w *Web3User) Unwrap() *Web3User {
	tx, ok := w.config.driver.(*txDriver)
	if !ok {
		panic("ent: Web3User is not a transactional entity")
	}
	w.config.driver = tx.drv
	return w
}

// String implements the fmt.Stringer.
func (w *Web3User) String() string {
	var builder strings.Builder
	builder.WriteString("Web3User(")
	builder.WriteString(fmt.Sprintf("id=%v", w.ID))
	builder.WriteString(", uuid=")
	builder.WriteString(fmt.Sprintf("%v", w.UUID))
	builder.WriteString(", address=")
	builder.WriteString(w.Address)
	builder.WriteString(", wallet_type=")
	builder.WriteString(fmt.Sprintf("%v", w.WalletType))
	builder.WriteByte(')')
	return builder.String()
}

// Web3Users is a parsable slice of Web3User.
type Web3Users []*Web3User

func (w Web3Users) config(cfg config) {
	for _i := range w {
		w[_i].config = cfg
	}
}
