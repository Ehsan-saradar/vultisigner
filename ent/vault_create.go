// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"vultisigner/ent/ecdsalocaldata"
	"vultisigner/ent/eddsalocaldata"
	"vultisigner/ent/pailliersk"
	"vultisigner/ent/vault"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// VaultCreate is the builder for creating a Vault entity.
type VaultCreate struct {
	config
	mutation *VaultMutation
	hooks    []Hook
}

// SetName sets the "name" field.
func (vc *VaultCreate) SetName(s string) *VaultCreate {
	vc.mutation.SetName(s)
	return vc
}

// SetPubKey sets the "pub_key" field.
func (vc *VaultCreate) SetPubKey(s string) *VaultCreate {
	vc.mutation.SetPubKey(s)
	return vc
}

// SetNillablePubKey sets the "pub_key" field if the given value is not nil.
func (vc *VaultCreate) SetNillablePubKey(s *string) *VaultCreate {
	if s != nil {
		vc.SetPubKey(*s)
	}
	return vc
}

// SetKeygenCommitteeKeys sets the "keygen_committee_keys" field.
func (vc *VaultCreate) SetKeygenCommitteeKeys(s []string) *VaultCreate {
	vc.mutation.SetKeygenCommitteeKeys(s)
	return vc
}

// SetLocalPartyKey sets the "local_party_key" field.
func (vc *VaultCreate) SetLocalPartyKey(s string) *VaultCreate {
	vc.mutation.SetLocalPartyKey(s)
	return vc
}

// SetChainCodeHex sets the "chain_code_hex" field.
func (vc *VaultCreate) SetChainCodeHex(s string) *VaultCreate {
	vc.mutation.SetChainCodeHex(s)
	return vc
}

// SetResharePrefix sets the "reshare_prefix" field.
func (vc *VaultCreate) SetResharePrefix(s string) *VaultCreate {
	vc.mutation.SetResharePrefix(s)
	return vc
}

// SetNillableResharePrefix sets the "reshare_prefix" field if the given value is not nil.
func (vc *VaultCreate) SetNillableResharePrefix(s *string) *VaultCreate {
	if s != nil {
		vc.SetResharePrefix(*s)
	}
	return vc
}

// SetPaillierSkID sets the "paillier_sk" edge to the PaillierSK entity by ID.
func (vc *VaultCreate) SetPaillierSkID(id int) *VaultCreate {
	vc.mutation.SetPaillierSkID(id)
	return vc
}

// SetNillablePaillierSkID sets the "paillier_sk" edge to the PaillierSK entity by ID if the given value is not nil.
func (vc *VaultCreate) SetNillablePaillierSkID(id *int) *VaultCreate {
	if id != nil {
		vc = vc.SetPaillierSkID(*id)
	}
	return vc
}

// SetPaillierSk sets the "paillier_sk" edge to the PaillierSK entity.
func (vc *VaultCreate) SetPaillierSk(p *PaillierSK) *VaultCreate {
	return vc.SetPaillierSkID(p.ID)
}

// SetEcdsaLocalDataID sets the "ecdsa_local_data" edge to the EcdsaLocalData entity by ID.
func (vc *VaultCreate) SetEcdsaLocalDataID(id int) *VaultCreate {
	vc.mutation.SetEcdsaLocalDataID(id)
	return vc
}

// SetNillableEcdsaLocalDataID sets the "ecdsa_local_data" edge to the EcdsaLocalData entity by ID if the given value is not nil.
func (vc *VaultCreate) SetNillableEcdsaLocalDataID(id *int) *VaultCreate {
	if id != nil {
		vc = vc.SetEcdsaLocalDataID(*id)
	}
	return vc
}

// SetEcdsaLocalData sets the "ecdsa_local_data" edge to the EcdsaLocalData entity.
func (vc *VaultCreate) SetEcdsaLocalData(e *EcdsaLocalData) *VaultCreate {
	return vc.SetEcdsaLocalDataID(e.ID)
}

// SetEddsaLocalDataID sets the "eddsa_local_data" edge to the EddsaLocalData entity by ID.
func (vc *VaultCreate) SetEddsaLocalDataID(id int) *VaultCreate {
	vc.mutation.SetEddsaLocalDataID(id)
	return vc
}

// SetNillableEddsaLocalDataID sets the "eddsa_local_data" edge to the EddsaLocalData entity by ID if the given value is not nil.
func (vc *VaultCreate) SetNillableEddsaLocalDataID(id *int) *VaultCreate {
	if id != nil {
		vc = vc.SetEddsaLocalDataID(*id)
	}
	return vc
}

// SetEddsaLocalData sets the "eddsa_local_data" edge to the EddsaLocalData entity.
func (vc *VaultCreate) SetEddsaLocalData(e *EddsaLocalData) *VaultCreate {
	return vc.SetEddsaLocalDataID(e.ID)
}

// Mutation returns the VaultMutation object of the builder.
func (vc *VaultCreate) Mutation() *VaultMutation {
	return vc.mutation
}

// Save creates the Vault in the database.
func (vc *VaultCreate) Save(ctx context.Context) (*Vault, error) {
	return withHooks(ctx, vc.sqlSave, vc.mutation, vc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (vc *VaultCreate) SaveX(ctx context.Context) *Vault {
	v, err := vc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (vc *VaultCreate) Exec(ctx context.Context) error {
	_, err := vc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (vc *VaultCreate) ExecX(ctx context.Context) {
	if err := vc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (vc *VaultCreate) check() error {
	if _, ok := vc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "Vault.name"`)}
	}
	if v, ok := vc.mutation.Name(); ok {
		if err := vault.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Vault.name": %w`, err)}
		}
	}
	if v, ok := vc.mutation.PubKey(); ok {
		if err := vault.PubKeyValidator(v); err != nil {
			return &ValidationError{Name: "pub_key", err: fmt.Errorf(`ent: validator failed for field "Vault.pub_key": %w`, err)}
		}
	}
	if _, ok := vc.mutation.LocalPartyKey(); !ok {
		return &ValidationError{Name: "local_party_key", err: errors.New(`ent: missing required field "Vault.local_party_key"`)}
	}
	if v, ok := vc.mutation.LocalPartyKey(); ok {
		if err := vault.LocalPartyKeyValidator(v); err != nil {
			return &ValidationError{Name: "local_party_key", err: fmt.Errorf(`ent: validator failed for field "Vault.local_party_key": %w`, err)}
		}
	}
	if _, ok := vc.mutation.ChainCodeHex(); !ok {
		return &ValidationError{Name: "chain_code_hex", err: errors.New(`ent: missing required field "Vault.chain_code_hex"`)}
	}
	if v, ok := vc.mutation.ChainCodeHex(); ok {
		if err := vault.ChainCodeHexValidator(v); err != nil {
			return &ValidationError{Name: "chain_code_hex", err: fmt.Errorf(`ent: validator failed for field "Vault.chain_code_hex": %w`, err)}
		}
	}
	return nil
}

func (vc *VaultCreate) sqlSave(ctx context.Context) (*Vault, error) {
	if err := vc.check(); err != nil {
		return nil, err
	}
	_node, _spec := vc.createSpec()
	if err := sqlgraph.CreateNode(ctx, vc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	vc.mutation.id = &_node.ID
	vc.mutation.done = true
	return _node, nil
}

func (vc *VaultCreate) createSpec() (*Vault, *sqlgraph.CreateSpec) {
	var (
		_node = &Vault{config: vc.config}
		_spec = sqlgraph.NewCreateSpec(vault.Table, sqlgraph.NewFieldSpec(vault.FieldID, field.TypeInt))
	)
	if value, ok := vc.mutation.Name(); ok {
		_spec.SetField(vault.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := vc.mutation.PubKey(); ok {
		_spec.SetField(vault.FieldPubKey, field.TypeString, value)
		_node.PubKey = value
	}
	if value, ok := vc.mutation.KeygenCommitteeKeys(); ok {
		_spec.SetField(vault.FieldKeygenCommitteeKeys, field.TypeJSON, value)
		_node.KeygenCommitteeKeys = value
	}
	if value, ok := vc.mutation.LocalPartyKey(); ok {
		_spec.SetField(vault.FieldLocalPartyKey, field.TypeString, value)
		_node.LocalPartyKey = value
	}
	if value, ok := vc.mutation.ChainCodeHex(); ok {
		_spec.SetField(vault.FieldChainCodeHex, field.TypeString, value)
		_node.ChainCodeHex = value
	}
	if value, ok := vc.mutation.ResharePrefix(); ok {
		_spec.SetField(vault.FieldResharePrefix, field.TypeString, value)
		_node.ResharePrefix = value
	}
	if nodes := vc.mutation.PaillierSkIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   vault.PaillierSkTable,
			Columns: []string{vault.PaillierSkColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(pailliersk.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.paillier_sk_vault = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := vc.mutation.EcdsaLocalDataIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   vault.EcdsaLocalDataTable,
			Columns: []string{vault.EcdsaLocalDataColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(ecdsalocaldata.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.ecdsa_local_data_vault = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := vc.mutation.EddsaLocalDataIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   vault.EddsaLocalDataTable,
			Columns: []string{vault.EddsaLocalDataColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(eddsalocaldata.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.eddsa_local_data_vault = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// VaultCreateBulk is the builder for creating many Vault entities in bulk.
type VaultCreateBulk struct {
	config
	err      error
	builders []*VaultCreate
}

// Save creates the Vault entities in the database.
func (vcb *VaultCreateBulk) Save(ctx context.Context) ([]*Vault, error) {
	if vcb.err != nil {
		return nil, vcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(vcb.builders))
	nodes := make([]*Vault, len(vcb.builders))
	mutators := make([]Mutator, len(vcb.builders))
	for i := range vcb.builders {
		func(i int, root context.Context) {
			builder := vcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*VaultMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, vcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, vcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, vcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (vcb *VaultCreateBulk) SaveX(ctx context.Context) []*Vault {
	v, err := vcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (vcb *VaultCreateBulk) Exec(ctx context.Context) error {
	_, err := vcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (vcb *VaultCreateBulk) ExecX(ctx context.Context) {
	if err := vcb.Exec(ctx); err != nil {
		panic(err)
	}
}
