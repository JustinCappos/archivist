// Code generated by ent, DO NOT EDIT.

package attestation

const (
	// Label holds the string label denoting the attestation type in the database.
	Label = "attestation"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldType holds the string denoting the type field in the database.
	FieldType = "type"
	// EdgeAttestationCollection holds the string denoting the attestation_collection edge name in mutations.
	EdgeAttestationCollection = "attestation_collection"
	// Table holds the table name of the attestation in the database.
	Table = "attestations"
	// AttestationCollectionTable is the table that holds the attestation_collection relation/edge.
	AttestationCollectionTable = "attestations"
	// AttestationCollectionInverseTable is the table name for the AttestationCollection entity.
	// It exists in this package in order to avoid circular dependency with the "attestationcollection" package.
	AttestationCollectionInverseTable = "attestation_collections"
	// AttestationCollectionColumn is the table column denoting the attestation_collection relation/edge.
	AttestationCollectionColumn = "attestation_collection_attestations"
)

// Columns holds all SQL columns for attestation fields.
var Columns = []string{
	FieldID,
	FieldType,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "attestations"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"attestation_collection_attestations",
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
	// TypeValidator is a validator for the "type" field. It is called by the builders before save.
	TypeValidator func(string) error
)
