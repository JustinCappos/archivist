// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// AttestationsColumns holds the columns for the "attestations" table.
	AttestationsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "type", Type: field.TypeString},
		{Name: "attestation_collection_attestations", Type: field.TypeInt},
	}
	// AttestationsTable holds the schema information for the "attestations" table.
	AttestationsTable = &schema.Table{
		Name:       "attestations",
		Columns:    AttestationsColumns,
		PrimaryKey: []*schema.Column{AttestationsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "attestations_attestation_collections_attestations",
				Columns:    []*schema.Column{AttestationsColumns[2]},
				RefColumns: []*schema.Column{AttestationCollectionsColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "attestation_type",
				Unique:  false,
				Columns: []*schema.Column{AttestationsColumns[1]},
			},
		},
	}
	// AttestationCollectionsColumns holds the columns for the "attestation_collections" table.
	AttestationCollectionsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString},
		{Name: "statement_attestation_collections", Type: field.TypeInt, Unique: true},
	}
	// AttestationCollectionsTable holds the schema information for the "attestation_collections" table.
	AttestationCollectionsTable = &schema.Table{
		Name:       "attestation_collections",
		Columns:    AttestationCollectionsColumns,
		PrimaryKey: []*schema.Column{AttestationCollectionsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "attestation_collections_statements_attestation_collections",
				Columns:    []*schema.Column{AttestationCollectionsColumns[2]},
				RefColumns: []*schema.Column{StatementsColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "attestationcollection_name",
				Unique:  false,
				Columns: []*schema.Column{AttestationCollectionsColumns[1]},
			},
		},
	}
	// DssesColumns holds the columns for the "dsses" table.
	DssesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "gitoid_sha256", Type: field.TypeString, Unique: true},
		{Name: "payload_type", Type: field.TypeString},
		{Name: "dsse_statement", Type: field.TypeInt, Nullable: true},
	}
	// DssesTable holds the schema information for the "dsses" table.
	DssesTable = &schema.Table{
		Name:       "dsses",
		Columns:    DssesColumns,
		PrimaryKey: []*schema.Column{DssesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "dsses_statements_statement",
				Columns:    []*schema.Column{DssesColumns[3]},
				RefColumns: []*schema.Column{StatementsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// PayloadDigestsColumns holds the columns for the "payload_digests" table.
	PayloadDigestsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "algorithm", Type: field.TypeString},
		{Name: "value", Type: field.TypeString},
		{Name: "dsse_payload_digests", Type: field.TypeInt, Nullable: true},
	}
	// PayloadDigestsTable holds the schema information for the "payload_digests" table.
	PayloadDigestsTable = &schema.Table{
		Name:       "payload_digests",
		Columns:    PayloadDigestsColumns,
		PrimaryKey: []*schema.Column{PayloadDigestsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "payload_digests_dsses_payload_digests",
				Columns:    []*schema.Column{PayloadDigestsColumns[3]},
				RefColumns: []*schema.Column{DssesColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "payloaddigest_value",
				Unique:  false,
				Columns: []*schema.Column{PayloadDigestsColumns[2]},
			},
		},
	}
	// SignaturesColumns holds the columns for the "signatures" table.
	SignaturesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "key_id", Type: field.TypeString},
		{Name: "signature", Type: field.TypeString, SchemaType: map[string]string{"mysql": "text"}},
		{Name: "dsse_signatures", Type: field.TypeInt, Nullable: true},
	}
	// SignaturesTable holds the schema information for the "signatures" table.
	SignaturesTable = &schema.Table{
		Name:       "signatures",
		Columns:    SignaturesColumns,
		PrimaryKey: []*schema.Column{SignaturesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "signatures_dsses_signatures",
				Columns:    []*schema.Column{SignaturesColumns[3]},
				RefColumns: []*schema.Column{DssesColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "signature_key_id",
				Unique:  false,
				Columns: []*schema.Column{SignaturesColumns[1]},
			},
		},
	}
	// StatementsColumns holds the columns for the "statements" table.
	StatementsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "predicate", Type: field.TypeString},
	}
	// StatementsTable holds the schema information for the "statements" table.
	StatementsTable = &schema.Table{
		Name:       "statements",
		Columns:    StatementsColumns,
		PrimaryKey: []*schema.Column{StatementsColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "statement_predicate",
				Unique:  false,
				Columns: []*schema.Column{StatementsColumns[1]},
			},
		},
	}
	// SubjectsColumns holds the columns for the "subjects" table.
	SubjectsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString},
		{Name: "statement_subjects", Type: field.TypeInt, Nullable: true},
	}
	// SubjectsTable holds the schema information for the "subjects" table.
	SubjectsTable = &schema.Table{
		Name:       "subjects",
		Columns:    SubjectsColumns,
		PrimaryKey: []*schema.Column{SubjectsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "subjects_statements_subjects",
				Columns:    []*schema.Column{SubjectsColumns[2]},
				RefColumns: []*schema.Column{StatementsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "subject_name",
				Unique:  false,
				Columns: []*schema.Column{SubjectsColumns[1]},
			},
		},
	}
	// SubjectDigestsColumns holds the columns for the "subject_digests" table.
	SubjectDigestsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "algorithm", Type: field.TypeString},
		{Name: "value", Type: field.TypeString},
		{Name: "subject_subject_digests", Type: field.TypeInt, Nullable: true},
	}
	// SubjectDigestsTable holds the schema information for the "subject_digests" table.
	SubjectDigestsTable = &schema.Table{
		Name:       "subject_digests",
		Columns:    SubjectDigestsColumns,
		PrimaryKey: []*schema.Column{SubjectDigestsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "subject_digests_subjects_subject_digests",
				Columns:    []*schema.Column{SubjectDigestsColumns[3]},
				RefColumns: []*schema.Column{SubjectsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "subjectdigest_value",
				Unique:  false,
				Columns: []*schema.Column{SubjectDigestsColumns[2]},
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		AttestationsTable,
		AttestationCollectionsTable,
		DssesTable,
		PayloadDigestsTable,
		SignaturesTable,
		StatementsTable,
		SubjectsTable,
		SubjectDigestsTable,
	}
)

func init() {
	AttestationsTable.ForeignKeys[0].RefTable = AttestationCollectionsTable
	AttestationCollectionsTable.ForeignKeys[0].RefTable = StatementsTable
	DssesTable.ForeignKeys[0].RefTable = StatementsTable
	PayloadDigestsTable.ForeignKeys[0].RefTable = DssesTable
	SignaturesTable.ForeignKeys[0].RefTable = DssesTable
	SubjectsTable.ForeignKeys[0].RefTable = StatementsTable
	SubjectDigestsTable.ForeignKeys[0].RefTable = SubjectsTable
}
