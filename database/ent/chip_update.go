// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"
	"viseh-api/database/ent/chip"
	"viseh-api/database/ent/predicate"
	"viseh-api/database/ent/user"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// ChipUpdate is the builder for updating Chip entities.
type ChipUpdate struct {
	config
	hooks    []Hook
	mutation *ChipMutation
}

// Where appends a list predicates to the ChipUpdate builder.
func (cu *ChipUpdate) Where(ps ...predicate.Chip) *ChipUpdate {
	cu.mutation.Where(ps...)
	return cu
}

// SetName sets the "name" field.
func (cu *ChipUpdate) SetName(s string) *ChipUpdate {
	cu.mutation.SetName(s)
	return cu
}

// SetNillableName sets the "name" field if the given value is not nil.
func (cu *ChipUpdate) SetNillableName(s *string) *ChipUpdate {
	if s != nil {
		cu.SetName(*s)
	}
	return cu
}

// SetChipType sets the "chip_type" field.
func (cu *ChipUpdate) SetChipType(s string) *ChipUpdate {
	cu.mutation.SetChipType(s)
	return cu
}

// SetNillableChipType sets the "chip_type" field if the given value is not nil.
func (cu *ChipUpdate) SetNillableChipType(s *string) *ChipUpdate {
	if s != nil {
		cu.SetChipType(*s)
	}
	return cu
}

// SetLastUpdated sets the "last_updated" field.
func (cu *ChipUpdate) SetLastUpdated(t time.Time) *ChipUpdate {
	cu.mutation.SetLastUpdated(t)
	return cu
}

// SetNillableLastUpdated sets the "last_updated" field if the given value is not nil.
func (cu *ChipUpdate) SetNillableLastUpdated(t *time.Time) *ChipUpdate {
	if t != nil {
		cu.SetLastUpdated(*t)
	}
	return cu
}

// SetUsersID sets the "users" edge to the User entity by ID.
func (cu *ChipUpdate) SetUsersID(id int) *ChipUpdate {
	cu.mutation.SetUsersID(id)
	return cu
}

// SetNillableUsersID sets the "users" edge to the User entity by ID if the given value is not nil.
func (cu *ChipUpdate) SetNillableUsersID(id *int) *ChipUpdate {
	if id != nil {
		cu = cu.SetUsersID(*id)
	}
	return cu
}

// SetUsers sets the "users" edge to the User entity.
func (cu *ChipUpdate) SetUsers(u *User) *ChipUpdate {
	return cu.SetUsersID(u.ID)
}

// Mutation returns the ChipMutation object of the builder.
func (cu *ChipUpdate) Mutation() *ChipMutation {
	return cu.mutation
}

// ClearUsers clears the "users" edge to the User entity.
func (cu *ChipUpdate) ClearUsers() *ChipUpdate {
	cu.mutation.ClearUsers()
	return cu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (cu *ChipUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, cu.sqlSave, cu.mutation, cu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (cu *ChipUpdate) SaveX(ctx context.Context) int {
	affected, err := cu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (cu *ChipUpdate) Exec(ctx context.Context) error {
	_, err := cu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cu *ChipUpdate) ExecX(ctx context.Context) {
	if err := cu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cu *ChipUpdate) check() error {
	if v, ok := cu.mutation.Name(); ok {
		if err := chip.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Chip.name": %w`, err)}
		}
	}
	if v, ok := cu.mutation.ChipType(); ok {
		if err := chip.ChipTypeValidator(v); err != nil {
			return &ValidationError{Name: "chip_type", err: fmt.Errorf(`ent: validator failed for field "Chip.chip_type": %w`, err)}
		}
	}
	return nil
}

func (cu *ChipUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := cu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(chip.Table, chip.Columns, sqlgraph.NewFieldSpec(chip.FieldID, field.TypeInt))
	if ps := cu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cu.mutation.Name(); ok {
		_spec.SetField(chip.FieldName, field.TypeString, value)
	}
	if value, ok := cu.mutation.ChipType(); ok {
		_spec.SetField(chip.FieldChipType, field.TypeString, value)
	}
	if value, ok := cu.mutation.LastUpdated(); ok {
		_spec.SetField(chip.FieldLastUpdated, field.TypeTime, value)
	}
	if cu.mutation.UsersCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.UsersIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, cu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{chip.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	cu.mutation.done = true
	return n, nil
}

// ChipUpdateOne is the builder for updating a single Chip entity.
type ChipUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *ChipMutation
}

// SetName sets the "name" field.
func (cuo *ChipUpdateOne) SetName(s string) *ChipUpdateOne {
	cuo.mutation.SetName(s)
	return cuo
}

// SetNillableName sets the "name" field if the given value is not nil.
func (cuo *ChipUpdateOne) SetNillableName(s *string) *ChipUpdateOne {
	if s != nil {
		cuo.SetName(*s)
	}
	return cuo
}

// SetChipType sets the "chip_type" field.
func (cuo *ChipUpdateOne) SetChipType(s string) *ChipUpdateOne {
	cuo.mutation.SetChipType(s)
	return cuo
}

// SetNillableChipType sets the "chip_type" field if the given value is not nil.
func (cuo *ChipUpdateOne) SetNillableChipType(s *string) *ChipUpdateOne {
	if s != nil {
		cuo.SetChipType(*s)
	}
	return cuo
}

// SetLastUpdated sets the "last_updated" field.
func (cuo *ChipUpdateOne) SetLastUpdated(t time.Time) *ChipUpdateOne {
	cuo.mutation.SetLastUpdated(t)
	return cuo
}

// SetNillableLastUpdated sets the "last_updated" field if the given value is not nil.
func (cuo *ChipUpdateOne) SetNillableLastUpdated(t *time.Time) *ChipUpdateOne {
	if t != nil {
		cuo.SetLastUpdated(*t)
	}
	return cuo
}

// SetUsersID sets the "users" edge to the User entity by ID.
func (cuo *ChipUpdateOne) SetUsersID(id int) *ChipUpdateOne {
	cuo.mutation.SetUsersID(id)
	return cuo
}

// SetNillableUsersID sets the "users" edge to the User entity by ID if the given value is not nil.
func (cuo *ChipUpdateOne) SetNillableUsersID(id *int) *ChipUpdateOne {
	if id != nil {
		cuo = cuo.SetUsersID(*id)
	}
	return cuo
}

// SetUsers sets the "users" edge to the User entity.
func (cuo *ChipUpdateOne) SetUsers(u *User) *ChipUpdateOne {
	return cuo.SetUsersID(u.ID)
}

// Mutation returns the ChipMutation object of the builder.
func (cuo *ChipUpdateOne) Mutation() *ChipMutation {
	return cuo.mutation
}

// ClearUsers clears the "users" edge to the User entity.
func (cuo *ChipUpdateOne) ClearUsers() *ChipUpdateOne {
	cuo.mutation.ClearUsers()
	return cuo
}

// Where appends a list predicates to the ChipUpdate builder.
func (cuo *ChipUpdateOne) Where(ps ...predicate.Chip) *ChipUpdateOne {
	cuo.mutation.Where(ps...)
	return cuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (cuo *ChipUpdateOne) Select(field string, fields ...string) *ChipUpdateOne {
	cuo.fields = append([]string{field}, fields...)
	return cuo
}

// Save executes the query and returns the updated Chip entity.
func (cuo *ChipUpdateOne) Save(ctx context.Context) (*Chip, error) {
	return withHooks(ctx, cuo.sqlSave, cuo.mutation, cuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (cuo *ChipUpdateOne) SaveX(ctx context.Context) *Chip {
	node, err := cuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (cuo *ChipUpdateOne) Exec(ctx context.Context) error {
	_, err := cuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cuo *ChipUpdateOne) ExecX(ctx context.Context) {
	if err := cuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cuo *ChipUpdateOne) check() error {
	if v, ok := cuo.mutation.Name(); ok {
		if err := chip.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Chip.name": %w`, err)}
		}
	}
	if v, ok := cuo.mutation.ChipType(); ok {
		if err := chip.ChipTypeValidator(v); err != nil {
			return &ValidationError{Name: "chip_type", err: fmt.Errorf(`ent: validator failed for field "Chip.chip_type": %w`, err)}
		}
	}
	return nil
}

func (cuo *ChipUpdateOne) sqlSave(ctx context.Context) (_node *Chip, err error) {
	if err := cuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(chip.Table, chip.Columns, sqlgraph.NewFieldSpec(chip.FieldID, field.TypeInt))
	id, ok := cuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Chip.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := cuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, chip.FieldID)
		for _, f := range fields {
			if !chip.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != chip.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := cuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cuo.mutation.Name(); ok {
		_spec.SetField(chip.FieldName, field.TypeString, value)
	}
	if value, ok := cuo.mutation.ChipType(); ok {
		_spec.SetField(chip.FieldChipType, field.TypeString, value)
	}
	if value, ok := cuo.mutation.LastUpdated(); ok {
		_spec.SetField(chip.FieldLastUpdated, field.TypeTime, value)
	}
	if cuo.mutation.UsersCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.UsersIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Chip{config: cuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, cuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{chip.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	cuo.mutation.done = true
	return _node, nil
}
