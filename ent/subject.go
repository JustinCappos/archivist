// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/testifysec/archivist/ent/statement"
	"github.com/testifysec/archivist/ent/subject"
)

// Subject is the model entity for the Subject schema.
type Subject struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the SubjectQuery when eager-loading is set.
	Edges              SubjectEdges `json:"edges"`
	statement_subjects *int
}

// SubjectEdges holds the relations/edges for other nodes in the graph.
type SubjectEdges struct {
	// SubjectDigests holds the value of the subject_digests edge.
	SubjectDigests []*SubjectDigest `json:"subject_digests,omitempty"`
	// Statement holds the value of the statement edge.
	Statement *Statement `json:"statement,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
	// totalCount holds the count of the edges above.
	totalCount [2]*int
}

// SubjectDigestsOrErr returns the SubjectDigests value or an error if the edge
// was not loaded in eager-loading.
func (e SubjectEdges) SubjectDigestsOrErr() ([]*SubjectDigest, error) {
	if e.loadedTypes[0] {
		return e.SubjectDigests, nil
	}
	return nil, &NotLoadedError{edge: "subject_digests"}
}

// StatementOrErr returns the Statement value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e SubjectEdges) StatementOrErr() (*Statement, error) {
	if e.loadedTypes[1] {
		if e.Statement == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: statement.Label}
		}
		return e.Statement, nil
	}
	return nil, &NotLoadedError{edge: "statement"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Subject) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case subject.FieldID:
			values[i] = new(sql.NullInt64)
		case subject.FieldName:
			values[i] = new(sql.NullString)
		case subject.ForeignKeys[0]: // statement_subjects
			values[i] = new(sql.NullInt64)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Subject", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Subject fields.
func (s *Subject) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case subject.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			s.ID = int(value.Int64)
		case subject.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				s.Name = value.String
			}
		case subject.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field statement_subjects", value)
			} else if value.Valid {
				s.statement_subjects = new(int)
				*s.statement_subjects = int(value.Int64)
			}
		}
	}
	return nil
}

// QuerySubjectDigests queries the "subject_digests" edge of the Subject entity.
func (s *Subject) QuerySubjectDigests() *SubjectDigestQuery {
	return (&SubjectClient{config: s.config}).QuerySubjectDigests(s)
}

// QueryStatement queries the "statement" edge of the Subject entity.
func (s *Subject) QueryStatement() *StatementQuery {
	return (&SubjectClient{config: s.config}).QueryStatement(s)
}

// Update returns a builder for updating this Subject.
// Note that you need to call Subject.Unwrap() before calling this method if this Subject
// was returned from a transaction, and the transaction was committed or rolled back.
func (s *Subject) Update() *SubjectUpdateOne {
	return (&SubjectClient{config: s.config}).UpdateOne(s)
}

// Unwrap unwraps the Subject entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (s *Subject) Unwrap() *Subject {
	_tx, ok := s.config.driver.(*txDriver)
	if !ok {
		panic("ent: Subject is not a transactional entity")
	}
	s.config.driver = _tx.drv
	return s
}

// String implements the fmt.Stringer.
func (s *Subject) String() string {
	var builder strings.Builder
	builder.WriteString("Subject(")
	builder.WriteString(fmt.Sprintf("id=%v, ", s.ID))
	builder.WriteString("name=")
	builder.WriteString(s.Name)
	builder.WriteByte(')')
	return builder.String()
}

// Subjects is a parsable slice of Subject.
type Subjects []*Subject

func (s Subjects) config(cfg config) {
	for _i := range s {
		s[_i].config = cfg
	}
}
