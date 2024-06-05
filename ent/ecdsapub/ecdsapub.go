// Code generated by ent, DO NOT EDIT.

package ecdsapub

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the ecdsapub type in the database.
	Label = "ecdsa_pub"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCurve holds the string denoting the curve field in the database.
	FieldCurve = "curve"
	// FieldCoords holds the string denoting the coords field in the database.
	FieldCoords = "coords"
	// EdgeVault holds the string denoting the vault edge name in mutations.
	EdgeVault = "vault"
	// Table holds the table name of the ecdsapub in the database.
	Table = "ecdsa_pubs"
	// VaultTable is the table that holds the vault relation/edge.
	VaultTable = "ecdsa_pubs"
	// VaultInverseTable is the table name for the Vault entity.
	// It exists in this package in order to avoid circular dependency with the "vault" package.
	VaultInverseTable = "vaults"
	// VaultColumn is the table column denoting the vault relation/edge.
	VaultColumn = "ecdsa_pub_vault"
)

// Columns holds all SQL columns for ecdsapub fields.
var Columns = []string{
	FieldID,
	FieldCurve,
	FieldCoords,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "ecdsa_pubs"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"ecdsa_pub_vault",
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
	// CurveValidator is a validator for the "curve" field. It is called by the builders before save.
	CurveValidator func(string) error
)

// OrderOption defines the ordering options for the ECDSAPub queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByCurve orders the results by the curve field.
func ByCurve(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCurve, opts...).ToFunc()
}

// ByVaultField orders the results by vault field.
func ByVaultField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newVaultStep(), sql.OrderByField(field, opts...))
	}
}
func newVaultStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(VaultInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, false, VaultTable, VaultColumn),
	)
}
