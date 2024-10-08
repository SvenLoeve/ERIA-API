// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"
	"viseh-api/database/ent/chip"
	"viseh-api/database/ent/user"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// ChipCreate is the builder for creating a Chip entity.
type ChipCreate struct {
	config
	mutation *ChipMutation
	hooks    []Hook
}

// SetName sets the "name" field.
func (cc *ChipCreate) SetName(s string) *ChipCreate {
	cc.mutation.SetName(s)
	return cc
}

// SetNillableName sets the "name" field if the given value is not nil.
func (cc *ChipCreate) SetNillableName(s *string) *ChipCreate {
	if s != nil {
		cc.SetName(*s)
	}
	return cc
}

// SetChipType sets the "chip_type" field.
func (cc *ChipCreate) SetChipType(s string) *ChipCreate {
	cc.mutation.SetChipType(s)
	return cc
}

// SetLastUpdated sets the "last_updated" field.
func (cc *ChipCreate) SetLastUpdated(t time.Time) *ChipCreate {
	cc.mutation.SetLastUpdated(t)
	return cc
}

// SetUsersID sets the "users" edge to the User entity by ID.
func (cc *ChipCreate) SetUsersID(id int) *ChipCreate {
	cc.mutation.SetUsersID(id)
	return cc
}

// SetNillableUsersID sets the "users" edge to the User entity by ID if the given value is not nil.
func (cc *ChipCreate) SetNillableUsersID(id *int) *ChipCreate {
	if id != nil {
		cc = cc.SetUsersID(*id)
	}
	return cc
}

// SetUsers sets the "users" edge to the User entity.
func (cc *ChipCreate) SetUsers(u *User) *ChipCreate {
	return cc.SetUsersID(u.ID)
}

// Mutation returns the ChipMutation object of the builder.
func (cc *ChipCreate) Mutation() *ChipMutation {
	return cc.mutation
}

// Save creates the Chip in the database.
func (cc *ChipCreate) Save(ctx context.Context) (*Chip, error) {
	cc.defaults()
	return withHooks(ctx, cc.sqlSave, cc.mutation, cc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (cc *ChipCreate) SaveX(ctx context.Context) *Chip {
	v, err := cc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (cc *ChipCreate) Exec(ctx context.Context) error {
	_, err := cc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cc *ChipCreate) ExecX(ctx context.Context) {
	if err := cc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (cc *ChipCreate) defaults() {
	if _, ok := cc.mutation.Name(); !ok {
		v := chip.DefaultName
		cc.mutation.SetName(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cc *ChipCreate) check() error {
	if _, ok := cc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "Chip.name"`)}
	}
	if v, ok := cc.mutation.Name(); ok {
		if err := chip.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Chip.name": %w`, err)}
		}
	}
	if _, ok := cc.mutation.ChipType(); !ok {
		return &ValidationError{Name: "chip_type", err: errors.New(`ent: missing required field "Chip.chip_type"`)}
	}
	if v, ok := cc.mutation.ChipType(); ok {
		if err := chip.ChipTypeValidator(v); err != nil {
			return &ValidationError{Name: "chip_type", err: fmt.Errorf(`ent: validator failed for field "Chip.chip_type": %w`, err)}
		}
	}
	if _, ok := cc.mutation.LastUpdated(); !ok {
		return &ValidationError{Name: "last_updated", err: errors.New(`ent: missing required field "Chip.last_updated"`)}
	}
	return nil
}

func (cc *ChipCreate) sqlSave(ctx context.Context) (*Chip, error) {
	if err := cc.check(); err != nil {
		return nil, err
	}
	_node, _spec := cc.createSpec()
	if err := sqlgraph.CreateNode(ctx, cc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	cc.mutation.id = &_node.ID
	cc.mutation.done = true
	return _node, nil
}

func (cc *ChipCreate) createSpec() (*Chip, *sqlgraph.CreateSpec) {
	var (
		_node = &Chip{config: cc.config}
		_spec = sqlgraph.NewCreateSpec(chip.Table, sqlgraph.NewFieldSpec(chip.FieldID, field.TypeInt))
	)
	if value, ok := cc.mutation.Name(); ok {
		_spec.SetField(chip.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := cc.mutation.ChipType(); ok {
		_spec.SetField(chip.FieldChipType, field.TypeString, value)
		_node.ChipType = value
	}
	if value, ok := cc.mutation.LastUpdated(); ok {
		_spec.SetField(chip.FieldLastUpdated, field.TypeTime, value)
		_node.LastUpdated = value
	}
	if nodes := cc.mutation.UsersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   chip.UsersTable,
			Columns: []string{chip.UsersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.user_chips = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// ChipCreateBulk is the builder for creating many Chip entities in bulk.
type ChipCreateBulk struct {
	config
	err      error
	builders []*ChipCreate
}

// Save creates the Chip entities in the database.
func (ccb *ChipCreateBulk) Save(ctx context.Context) ([]*Chip, error) {
	if ccb.err != nil {
		return nil, ccb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(ccb.builders))
	nodes := make([]*Chip, len(ccb.builders))
	mutators := make([]Mutator, len(ccb.builders))
	for i := range ccb.builders {
		func(i int, root context.Context) {
			builder := ccb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ChipMutation)
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
					_, err = mutators[i+1].Mutate(root, ccb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ccb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, ccb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ccb *ChipCreateBulk) SaveX(ctx context.Context) []*Chip {
	v, err := ccb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ccb *ChipCreateBulk) Exec(ctx context.Context) error {
	_, err := ccb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ccb *ChipCreateBulk) ExecX(ctx context.Context) {
	if err := ccb.Exec(ctx); err != nil {
		panic(err)
	}
}
