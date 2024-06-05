// Code generated by ent, DO NOT EDIT.

package vault

import (
	"vultisigner/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Vault {
	return predicate.Vault(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Vault {
	return predicate.Vault(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Vault {
	return predicate.Vault(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Vault {
	return predicate.Vault(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Vault {
	return predicate.Vault(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Vault {
	return predicate.Vault(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Vault {
	return predicate.Vault(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Vault {
	return predicate.Vault(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Vault {
	return predicate.Vault(sql.FieldLTE(FieldID, id))
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.Vault {
	return predicate.Vault(sql.FieldEQ(FieldName, v))
}

// PubKey applies equality check predicate on the "pub_key" field. It's identical to PubKeyEQ.
func PubKey(v string) predicate.Vault {
	return predicate.Vault(sql.FieldEQ(FieldPubKey, v))
}

// LocalPartyKey applies equality check predicate on the "local_party_key" field. It's identical to LocalPartyKeyEQ.
func LocalPartyKey(v string) predicate.Vault {
	return predicate.Vault(sql.FieldEQ(FieldLocalPartyKey, v))
}

// ChainCodeHex applies equality check predicate on the "chain_code_hex" field. It's identical to ChainCodeHexEQ.
func ChainCodeHex(v string) predicate.Vault {
	return predicate.Vault(sql.FieldEQ(FieldChainCodeHex, v))
}

// ResharePrefix applies equality check predicate on the "reshare_prefix" field. It's identical to ResharePrefixEQ.
func ResharePrefix(v string) predicate.Vault {
	return predicate.Vault(sql.FieldEQ(FieldResharePrefix, v))
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.Vault {
	return predicate.Vault(sql.FieldEQ(FieldName, v))
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.Vault {
	return predicate.Vault(sql.FieldNEQ(FieldName, v))
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.Vault {
	return predicate.Vault(sql.FieldIn(FieldName, vs...))
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.Vault {
	return predicate.Vault(sql.FieldNotIn(FieldName, vs...))
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.Vault {
	return predicate.Vault(sql.FieldGT(FieldName, v))
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.Vault {
	return predicate.Vault(sql.FieldGTE(FieldName, v))
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.Vault {
	return predicate.Vault(sql.FieldLT(FieldName, v))
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.Vault {
	return predicate.Vault(sql.FieldLTE(FieldName, v))
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.Vault {
	return predicate.Vault(sql.FieldContains(FieldName, v))
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.Vault {
	return predicate.Vault(sql.FieldHasPrefix(FieldName, v))
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.Vault {
	return predicate.Vault(sql.FieldHasSuffix(FieldName, v))
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.Vault {
	return predicate.Vault(sql.FieldEqualFold(FieldName, v))
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.Vault {
	return predicate.Vault(sql.FieldContainsFold(FieldName, v))
}

// PubKeyEQ applies the EQ predicate on the "pub_key" field.
func PubKeyEQ(v string) predicate.Vault {
	return predicate.Vault(sql.FieldEQ(FieldPubKey, v))
}

// PubKeyNEQ applies the NEQ predicate on the "pub_key" field.
func PubKeyNEQ(v string) predicate.Vault {
	return predicate.Vault(sql.FieldNEQ(FieldPubKey, v))
}

// PubKeyIn applies the In predicate on the "pub_key" field.
func PubKeyIn(vs ...string) predicate.Vault {
	return predicate.Vault(sql.FieldIn(FieldPubKey, vs...))
}

// PubKeyNotIn applies the NotIn predicate on the "pub_key" field.
func PubKeyNotIn(vs ...string) predicate.Vault {
	return predicate.Vault(sql.FieldNotIn(FieldPubKey, vs...))
}

// PubKeyGT applies the GT predicate on the "pub_key" field.
func PubKeyGT(v string) predicate.Vault {
	return predicate.Vault(sql.FieldGT(FieldPubKey, v))
}

// PubKeyGTE applies the GTE predicate on the "pub_key" field.
func PubKeyGTE(v string) predicate.Vault {
	return predicate.Vault(sql.FieldGTE(FieldPubKey, v))
}

// PubKeyLT applies the LT predicate on the "pub_key" field.
func PubKeyLT(v string) predicate.Vault {
	return predicate.Vault(sql.FieldLT(FieldPubKey, v))
}

// PubKeyLTE applies the LTE predicate on the "pub_key" field.
func PubKeyLTE(v string) predicate.Vault {
	return predicate.Vault(sql.FieldLTE(FieldPubKey, v))
}

// PubKeyContains applies the Contains predicate on the "pub_key" field.
func PubKeyContains(v string) predicate.Vault {
	return predicate.Vault(sql.FieldContains(FieldPubKey, v))
}

// PubKeyHasPrefix applies the HasPrefix predicate on the "pub_key" field.
func PubKeyHasPrefix(v string) predicate.Vault {
	return predicate.Vault(sql.FieldHasPrefix(FieldPubKey, v))
}

// PubKeyHasSuffix applies the HasSuffix predicate on the "pub_key" field.
func PubKeyHasSuffix(v string) predicate.Vault {
	return predicate.Vault(sql.FieldHasSuffix(FieldPubKey, v))
}

// PubKeyIsNil applies the IsNil predicate on the "pub_key" field.
func PubKeyIsNil() predicate.Vault {
	return predicate.Vault(sql.FieldIsNull(FieldPubKey))
}

// PubKeyNotNil applies the NotNil predicate on the "pub_key" field.
func PubKeyNotNil() predicate.Vault {
	return predicate.Vault(sql.FieldNotNull(FieldPubKey))
}

// PubKeyEqualFold applies the EqualFold predicate on the "pub_key" field.
func PubKeyEqualFold(v string) predicate.Vault {
	return predicate.Vault(sql.FieldEqualFold(FieldPubKey, v))
}

// PubKeyContainsFold applies the ContainsFold predicate on the "pub_key" field.
func PubKeyContainsFold(v string) predicate.Vault {
	return predicate.Vault(sql.FieldContainsFold(FieldPubKey, v))
}

// KeygenCommitteeKeysIsNil applies the IsNil predicate on the "keygen_committee_keys" field.
func KeygenCommitteeKeysIsNil() predicate.Vault {
	return predicate.Vault(sql.FieldIsNull(FieldKeygenCommitteeKeys))
}

// KeygenCommitteeKeysNotNil applies the NotNil predicate on the "keygen_committee_keys" field.
func KeygenCommitteeKeysNotNil() predicate.Vault {
	return predicate.Vault(sql.FieldNotNull(FieldKeygenCommitteeKeys))
}

// LocalPartyKeyEQ applies the EQ predicate on the "local_party_key" field.
func LocalPartyKeyEQ(v string) predicate.Vault {
	return predicate.Vault(sql.FieldEQ(FieldLocalPartyKey, v))
}

// LocalPartyKeyNEQ applies the NEQ predicate on the "local_party_key" field.
func LocalPartyKeyNEQ(v string) predicate.Vault {
	return predicate.Vault(sql.FieldNEQ(FieldLocalPartyKey, v))
}

// LocalPartyKeyIn applies the In predicate on the "local_party_key" field.
func LocalPartyKeyIn(vs ...string) predicate.Vault {
	return predicate.Vault(sql.FieldIn(FieldLocalPartyKey, vs...))
}

// LocalPartyKeyNotIn applies the NotIn predicate on the "local_party_key" field.
func LocalPartyKeyNotIn(vs ...string) predicate.Vault {
	return predicate.Vault(sql.FieldNotIn(FieldLocalPartyKey, vs...))
}

// LocalPartyKeyGT applies the GT predicate on the "local_party_key" field.
func LocalPartyKeyGT(v string) predicate.Vault {
	return predicate.Vault(sql.FieldGT(FieldLocalPartyKey, v))
}

// LocalPartyKeyGTE applies the GTE predicate on the "local_party_key" field.
func LocalPartyKeyGTE(v string) predicate.Vault {
	return predicate.Vault(sql.FieldGTE(FieldLocalPartyKey, v))
}

// LocalPartyKeyLT applies the LT predicate on the "local_party_key" field.
func LocalPartyKeyLT(v string) predicate.Vault {
	return predicate.Vault(sql.FieldLT(FieldLocalPartyKey, v))
}

// LocalPartyKeyLTE applies the LTE predicate on the "local_party_key" field.
func LocalPartyKeyLTE(v string) predicate.Vault {
	return predicate.Vault(sql.FieldLTE(FieldLocalPartyKey, v))
}

// LocalPartyKeyContains applies the Contains predicate on the "local_party_key" field.
func LocalPartyKeyContains(v string) predicate.Vault {
	return predicate.Vault(sql.FieldContains(FieldLocalPartyKey, v))
}

// LocalPartyKeyHasPrefix applies the HasPrefix predicate on the "local_party_key" field.
func LocalPartyKeyHasPrefix(v string) predicate.Vault {
	return predicate.Vault(sql.FieldHasPrefix(FieldLocalPartyKey, v))
}

// LocalPartyKeyHasSuffix applies the HasSuffix predicate on the "local_party_key" field.
func LocalPartyKeyHasSuffix(v string) predicate.Vault {
	return predicate.Vault(sql.FieldHasSuffix(FieldLocalPartyKey, v))
}

// LocalPartyKeyEqualFold applies the EqualFold predicate on the "local_party_key" field.
func LocalPartyKeyEqualFold(v string) predicate.Vault {
	return predicate.Vault(sql.FieldEqualFold(FieldLocalPartyKey, v))
}

// LocalPartyKeyContainsFold applies the ContainsFold predicate on the "local_party_key" field.
func LocalPartyKeyContainsFold(v string) predicate.Vault {
	return predicate.Vault(sql.FieldContainsFold(FieldLocalPartyKey, v))
}

// ChainCodeHexEQ applies the EQ predicate on the "chain_code_hex" field.
func ChainCodeHexEQ(v string) predicate.Vault {
	return predicate.Vault(sql.FieldEQ(FieldChainCodeHex, v))
}

// ChainCodeHexNEQ applies the NEQ predicate on the "chain_code_hex" field.
func ChainCodeHexNEQ(v string) predicate.Vault {
	return predicate.Vault(sql.FieldNEQ(FieldChainCodeHex, v))
}

// ChainCodeHexIn applies the In predicate on the "chain_code_hex" field.
func ChainCodeHexIn(vs ...string) predicate.Vault {
	return predicate.Vault(sql.FieldIn(FieldChainCodeHex, vs...))
}

// ChainCodeHexNotIn applies the NotIn predicate on the "chain_code_hex" field.
func ChainCodeHexNotIn(vs ...string) predicate.Vault {
	return predicate.Vault(sql.FieldNotIn(FieldChainCodeHex, vs...))
}

// ChainCodeHexGT applies the GT predicate on the "chain_code_hex" field.
func ChainCodeHexGT(v string) predicate.Vault {
	return predicate.Vault(sql.FieldGT(FieldChainCodeHex, v))
}

// ChainCodeHexGTE applies the GTE predicate on the "chain_code_hex" field.
func ChainCodeHexGTE(v string) predicate.Vault {
	return predicate.Vault(sql.FieldGTE(FieldChainCodeHex, v))
}

// ChainCodeHexLT applies the LT predicate on the "chain_code_hex" field.
func ChainCodeHexLT(v string) predicate.Vault {
	return predicate.Vault(sql.FieldLT(FieldChainCodeHex, v))
}

// ChainCodeHexLTE applies the LTE predicate on the "chain_code_hex" field.
func ChainCodeHexLTE(v string) predicate.Vault {
	return predicate.Vault(sql.FieldLTE(FieldChainCodeHex, v))
}

// ChainCodeHexContains applies the Contains predicate on the "chain_code_hex" field.
func ChainCodeHexContains(v string) predicate.Vault {
	return predicate.Vault(sql.FieldContains(FieldChainCodeHex, v))
}

// ChainCodeHexHasPrefix applies the HasPrefix predicate on the "chain_code_hex" field.
func ChainCodeHexHasPrefix(v string) predicate.Vault {
	return predicate.Vault(sql.FieldHasPrefix(FieldChainCodeHex, v))
}

// ChainCodeHexHasSuffix applies the HasSuffix predicate on the "chain_code_hex" field.
func ChainCodeHexHasSuffix(v string) predicate.Vault {
	return predicate.Vault(sql.FieldHasSuffix(FieldChainCodeHex, v))
}

// ChainCodeHexEqualFold applies the EqualFold predicate on the "chain_code_hex" field.
func ChainCodeHexEqualFold(v string) predicate.Vault {
	return predicate.Vault(sql.FieldEqualFold(FieldChainCodeHex, v))
}

// ChainCodeHexContainsFold applies the ContainsFold predicate on the "chain_code_hex" field.
func ChainCodeHexContainsFold(v string) predicate.Vault {
	return predicate.Vault(sql.FieldContainsFold(FieldChainCodeHex, v))
}

// ResharePrefixEQ applies the EQ predicate on the "reshare_prefix" field.
func ResharePrefixEQ(v string) predicate.Vault {
	return predicate.Vault(sql.FieldEQ(FieldResharePrefix, v))
}

// ResharePrefixNEQ applies the NEQ predicate on the "reshare_prefix" field.
func ResharePrefixNEQ(v string) predicate.Vault {
	return predicate.Vault(sql.FieldNEQ(FieldResharePrefix, v))
}

// ResharePrefixIn applies the In predicate on the "reshare_prefix" field.
func ResharePrefixIn(vs ...string) predicate.Vault {
	return predicate.Vault(sql.FieldIn(FieldResharePrefix, vs...))
}

// ResharePrefixNotIn applies the NotIn predicate on the "reshare_prefix" field.
func ResharePrefixNotIn(vs ...string) predicate.Vault {
	return predicate.Vault(sql.FieldNotIn(FieldResharePrefix, vs...))
}

// ResharePrefixGT applies the GT predicate on the "reshare_prefix" field.
func ResharePrefixGT(v string) predicate.Vault {
	return predicate.Vault(sql.FieldGT(FieldResharePrefix, v))
}

// ResharePrefixGTE applies the GTE predicate on the "reshare_prefix" field.
func ResharePrefixGTE(v string) predicate.Vault {
	return predicate.Vault(sql.FieldGTE(FieldResharePrefix, v))
}

// ResharePrefixLT applies the LT predicate on the "reshare_prefix" field.
func ResharePrefixLT(v string) predicate.Vault {
	return predicate.Vault(sql.FieldLT(FieldResharePrefix, v))
}

// ResharePrefixLTE applies the LTE predicate on the "reshare_prefix" field.
func ResharePrefixLTE(v string) predicate.Vault {
	return predicate.Vault(sql.FieldLTE(FieldResharePrefix, v))
}

// ResharePrefixContains applies the Contains predicate on the "reshare_prefix" field.
func ResharePrefixContains(v string) predicate.Vault {
	return predicate.Vault(sql.FieldContains(FieldResharePrefix, v))
}

// ResharePrefixHasPrefix applies the HasPrefix predicate on the "reshare_prefix" field.
func ResharePrefixHasPrefix(v string) predicate.Vault {
	return predicate.Vault(sql.FieldHasPrefix(FieldResharePrefix, v))
}

// ResharePrefixHasSuffix applies the HasSuffix predicate on the "reshare_prefix" field.
func ResharePrefixHasSuffix(v string) predicate.Vault {
	return predicate.Vault(sql.FieldHasSuffix(FieldResharePrefix, v))
}

// ResharePrefixIsNil applies the IsNil predicate on the "reshare_prefix" field.
func ResharePrefixIsNil() predicate.Vault {
	return predicate.Vault(sql.FieldIsNull(FieldResharePrefix))
}

// ResharePrefixNotNil applies the NotNil predicate on the "reshare_prefix" field.
func ResharePrefixNotNil() predicate.Vault {
	return predicate.Vault(sql.FieldNotNull(FieldResharePrefix))
}

// ResharePrefixEqualFold applies the EqualFold predicate on the "reshare_prefix" field.
func ResharePrefixEqualFold(v string) predicate.Vault {
	return predicate.Vault(sql.FieldEqualFold(FieldResharePrefix, v))
}

// ResharePrefixContainsFold applies the ContainsFold predicate on the "reshare_prefix" field.
func ResharePrefixContainsFold(v string) predicate.Vault {
	return predicate.Vault(sql.FieldContainsFold(FieldResharePrefix, v))
}

// HasPaillierSk applies the HasEdge predicate on the "paillier_sk" edge.
func HasPaillierSk() predicate.Vault {
	return predicate.Vault(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, PaillierSkTable, PaillierSkColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasPaillierSkWith applies the HasEdge predicate on the "paillier_sk" edge with a given conditions (other predicates).
func HasPaillierSkWith(preds ...predicate.PaillierSK) predicate.Vault {
	return predicate.Vault(func(s *sql.Selector) {
		step := newPaillierSkStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasEcdsaLocalData applies the HasEdge predicate on the "ecdsa_local_data" edge.
func HasEcdsaLocalData() predicate.Vault {
	return predicate.Vault(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, EcdsaLocalDataTable, EcdsaLocalDataColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasEcdsaLocalDataWith applies the HasEdge predicate on the "ecdsa_local_data" edge with a given conditions (other predicates).
func HasEcdsaLocalDataWith(preds ...predicate.EcdsaLocalData) predicate.Vault {
	return predicate.Vault(func(s *sql.Selector) {
		step := newEcdsaLocalDataStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasEddsaLocalData applies the HasEdge predicate on the "eddsa_local_data" edge.
func HasEddsaLocalData() predicate.Vault {
	return predicate.Vault(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, EddsaLocalDataTable, EddsaLocalDataColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasEddsaLocalDataWith applies the HasEdge predicate on the "eddsa_local_data" edge with a given conditions (other predicates).
func HasEddsaLocalDataWith(preds ...predicate.EddsaLocalData) predicate.Vault {
	return predicate.Vault(func(s *sql.Selector) {
		step := newEddsaLocalDataStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Vault) predicate.Vault {
	return predicate.Vault(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Vault) predicate.Vault {
	return predicate.Vault(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Vault) predicate.Vault {
	return predicate.Vault(sql.NotPredicates(p))
}
