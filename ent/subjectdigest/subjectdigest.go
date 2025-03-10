// Code generated by ent, DO NOT EDIT.

package subjectdigest

const (
	// Label holds the string label denoting the subjectdigest type in the database.
	Label = "subject_digest"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldAlgorithm holds the string denoting the algorithm field in the database.
	FieldAlgorithm = "algorithm"
	// FieldValue holds the string denoting the value field in the database.
	FieldValue = "value"
	// EdgeSubject holds the string denoting the subject edge name in mutations.
	EdgeSubject = "subject"
	// Table holds the table name of the subjectdigest in the database.
	Table = "subject_digests"
	// SubjectTable is the table that holds the subject relation/edge.
	SubjectTable = "subject_digests"
	// SubjectInverseTable is the table name for the Subject entity.
	// It exists in this package in order to avoid circular dependency with the "subject" package.
	SubjectInverseTable = "subjects"
	// SubjectColumn is the table column denoting the subject relation/edge.
	SubjectColumn = "subject_subject_digests"
)

// Columns holds all SQL columns for subjectdigest fields.
var Columns = []string{
	FieldID,
	FieldAlgorithm,
	FieldValue,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "subject_digests"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"subject_subject_digests",
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}

var (
	// AlgorithmValidator is a validator for the "algorithm" field. It is called by the builders before save.
	AlgorithmValidator func(string) error
	// ValueValidator is a validator for the "value" field. It is called by the builders before save.
	ValueValidator func(string) error
)
