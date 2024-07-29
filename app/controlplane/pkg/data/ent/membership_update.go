// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/chainloop-dev/chainloop/app/controlplane/pkg/authz"
	"github.com/chainloop-dev/chainloop/app/controlplane/pkg/data/ent/membership"
	"github.com/chainloop-dev/chainloop/app/controlplane/pkg/data/ent/organization"
	"github.com/chainloop-dev/chainloop/app/controlplane/pkg/data/ent/predicate"
	"github.com/chainloop-dev/chainloop/app/controlplane/pkg/data/ent/user"
	"github.com/google/uuid"
)

// MembershipUpdate is the builder for updating Membership entities.
type MembershipUpdate struct {
	config
	hooks     []Hook
	mutation  *MembershipMutation
	modifiers []func(*sql.UpdateBuilder)
}

// Where appends a list predicates to the MembershipUpdate builder.
func (mu *MembershipUpdate) Where(ps ...predicate.Membership) *MembershipUpdate {
	mu.mutation.Where(ps...)
	return mu
}

// SetCurrent sets the "current" field.
func (mu *MembershipUpdate) SetCurrent(b bool) *MembershipUpdate {
	mu.mutation.SetCurrent(b)
	return mu
}

// SetNillableCurrent sets the "current" field if the given value is not nil.
func (mu *MembershipUpdate) SetNillableCurrent(b *bool) *MembershipUpdate {
	if b != nil {
		mu.SetCurrent(*b)
	}
	return mu
}

// SetUpdatedAt sets the "updated_at" field.
func (mu *MembershipUpdate) SetUpdatedAt(t time.Time) *MembershipUpdate {
	mu.mutation.SetUpdatedAt(t)
	return mu
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (mu *MembershipUpdate) SetNillableUpdatedAt(t *time.Time) *MembershipUpdate {
	if t != nil {
		mu.SetUpdatedAt(*t)
	}
	return mu
}

// SetRole sets the "role" field.
func (mu *MembershipUpdate) SetRole(a authz.Role) *MembershipUpdate {
	mu.mutation.SetRole(a)
	return mu
}

// SetNillableRole sets the "role" field if the given value is not nil.
func (mu *MembershipUpdate) SetNillableRole(a *authz.Role) *MembershipUpdate {
	if a != nil {
		mu.SetRole(*a)
	}
	return mu
}

// SetOrganizationID sets the "organization" edge to the Organization entity by ID.
func (mu *MembershipUpdate) SetOrganizationID(id uuid.UUID) *MembershipUpdate {
	mu.mutation.SetOrganizationID(id)
	return mu
}

// SetOrganization sets the "organization" edge to the Organization entity.
func (mu *MembershipUpdate) SetOrganization(o *Organization) *MembershipUpdate {
	return mu.SetOrganizationID(o.ID)
}

// SetUserID sets the "user" edge to the User entity by ID.
func (mu *MembershipUpdate) SetUserID(id uuid.UUID) *MembershipUpdate {
	mu.mutation.SetUserID(id)
	return mu
}

// SetUser sets the "user" edge to the User entity.
func (mu *MembershipUpdate) SetUser(u *User) *MembershipUpdate {
	return mu.SetUserID(u.ID)
}

// Mutation returns the MembershipMutation object of the builder.
func (mu *MembershipUpdate) Mutation() *MembershipMutation {
	return mu.mutation
}

// ClearOrganization clears the "organization" edge to the Organization entity.
func (mu *MembershipUpdate) ClearOrganization() *MembershipUpdate {
	mu.mutation.ClearOrganization()
	return mu
}

// ClearUser clears the "user" edge to the User entity.
func (mu *MembershipUpdate) ClearUser() *MembershipUpdate {
	mu.mutation.ClearUser()
	return mu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (mu *MembershipUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, mu.sqlSave, mu.mutation, mu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (mu *MembershipUpdate) SaveX(ctx context.Context) int {
	affected, err := mu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (mu *MembershipUpdate) Exec(ctx context.Context) error {
	_, err := mu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mu *MembershipUpdate) ExecX(ctx context.Context) {
	if err := mu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (mu *MembershipUpdate) check() error {
	if v, ok := mu.mutation.Role(); ok {
		if err := membership.RoleValidator(v); err != nil {
			return &ValidationError{Name: "role", err: fmt.Errorf(`ent: validator failed for field "Membership.role": %w`, err)}
		}
	}
	if mu.mutation.OrganizationCleared() && len(mu.mutation.OrganizationIDs()) > 0 {
		return errors.New(`ent: clearing a required unique edge "Membership.organization"`)
	}
	if mu.mutation.UserCleared() && len(mu.mutation.UserIDs()) > 0 {
		return errors.New(`ent: clearing a required unique edge "Membership.user"`)
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (mu *MembershipUpdate) Modify(modifiers ...func(u *sql.UpdateBuilder)) *MembershipUpdate {
	mu.modifiers = append(mu.modifiers, modifiers...)
	return mu
}

func (mu *MembershipUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := mu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(membership.Table, membership.Columns, sqlgraph.NewFieldSpec(membership.FieldID, field.TypeUUID))
	if ps := mu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := mu.mutation.Current(); ok {
		_spec.SetField(membership.FieldCurrent, field.TypeBool, value)
	}
	if value, ok := mu.mutation.UpdatedAt(); ok {
		_spec.SetField(membership.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := mu.mutation.Role(); ok {
		_spec.SetField(membership.FieldRole, field.TypeEnum, value)
	}
	if mu.mutation.OrganizationCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   membership.OrganizationTable,
			Columns: []string{membership.OrganizationColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(organization.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := mu.mutation.OrganizationIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   membership.OrganizationTable,
			Columns: []string{membership.OrganizationColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(organization.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if mu.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   membership.UserTable,
			Columns: []string{membership.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := mu.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   membership.UserTable,
			Columns: []string{membership.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_spec.AddModifiers(mu.modifiers...)
	if n, err = sqlgraph.UpdateNodes(ctx, mu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{membership.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	mu.mutation.done = true
	return n, nil
}

// MembershipUpdateOne is the builder for updating a single Membership entity.
type MembershipUpdateOne struct {
	config
	fields    []string
	hooks     []Hook
	mutation  *MembershipMutation
	modifiers []func(*sql.UpdateBuilder)
}

// SetCurrent sets the "current" field.
func (muo *MembershipUpdateOne) SetCurrent(b bool) *MembershipUpdateOne {
	muo.mutation.SetCurrent(b)
	return muo
}

// SetNillableCurrent sets the "current" field if the given value is not nil.
func (muo *MembershipUpdateOne) SetNillableCurrent(b *bool) *MembershipUpdateOne {
	if b != nil {
		muo.SetCurrent(*b)
	}
	return muo
}

// SetUpdatedAt sets the "updated_at" field.
func (muo *MembershipUpdateOne) SetUpdatedAt(t time.Time) *MembershipUpdateOne {
	muo.mutation.SetUpdatedAt(t)
	return muo
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (muo *MembershipUpdateOne) SetNillableUpdatedAt(t *time.Time) *MembershipUpdateOne {
	if t != nil {
		muo.SetUpdatedAt(*t)
	}
	return muo
}

// SetRole sets the "role" field.
func (muo *MembershipUpdateOne) SetRole(a authz.Role) *MembershipUpdateOne {
	muo.mutation.SetRole(a)
	return muo
}

// SetNillableRole sets the "role" field if the given value is not nil.
func (muo *MembershipUpdateOne) SetNillableRole(a *authz.Role) *MembershipUpdateOne {
	if a != nil {
		muo.SetRole(*a)
	}
	return muo
}

// SetOrganizationID sets the "organization" edge to the Organization entity by ID.
func (muo *MembershipUpdateOne) SetOrganizationID(id uuid.UUID) *MembershipUpdateOne {
	muo.mutation.SetOrganizationID(id)
	return muo
}

// SetOrganization sets the "organization" edge to the Organization entity.
func (muo *MembershipUpdateOne) SetOrganization(o *Organization) *MembershipUpdateOne {
	return muo.SetOrganizationID(o.ID)
}

// SetUserID sets the "user" edge to the User entity by ID.
func (muo *MembershipUpdateOne) SetUserID(id uuid.UUID) *MembershipUpdateOne {
	muo.mutation.SetUserID(id)
	return muo
}

// SetUser sets the "user" edge to the User entity.
func (muo *MembershipUpdateOne) SetUser(u *User) *MembershipUpdateOne {
	return muo.SetUserID(u.ID)
}

// Mutation returns the MembershipMutation object of the builder.
func (muo *MembershipUpdateOne) Mutation() *MembershipMutation {
	return muo.mutation
}

// ClearOrganization clears the "organization" edge to the Organization entity.
func (muo *MembershipUpdateOne) ClearOrganization() *MembershipUpdateOne {
	muo.mutation.ClearOrganization()
	return muo
}

// ClearUser clears the "user" edge to the User entity.
func (muo *MembershipUpdateOne) ClearUser() *MembershipUpdateOne {
	muo.mutation.ClearUser()
	return muo
}

// Where appends a list predicates to the MembershipUpdate builder.
func (muo *MembershipUpdateOne) Where(ps ...predicate.Membership) *MembershipUpdateOne {
	muo.mutation.Where(ps...)
	return muo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (muo *MembershipUpdateOne) Select(field string, fields ...string) *MembershipUpdateOne {
	muo.fields = append([]string{field}, fields...)
	return muo
}

// Save executes the query and returns the updated Membership entity.
func (muo *MembershipUpdateOne) Save(ctx context.Context) (*Membership, error) {
	return withHooks(ctx, muo.sqlSave, muo.mutation, muo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (muo *MembershipUpdateOne) SaveX(ctx context.Context) *Membership {
	node, err := muo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (muo *MembershipUpdateOne) Exec(ctx context.Context) error {
	_, err := muo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (muo *MembershipUpdateOne) ExecX(ctx context.Context) {
	if err := muo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (muo *MembershipUpdateOne) check() error {
	if v, ok := muo.mutation.Role(); ok {
		if err := membership.RoleValidator(v); err != nil {
			return &ValidationError{Name: "role", err: fmt.Errorf(`ent: validator failed for field "Membership.role": %w`, err)}
		}
	}
	if muo.mutation.OrganizationCleared() && len(muo.mutation.OrganizationIDs()) > 0 {
		return errors.New(`ent: clearing a required unique edge "Membership.organization"`)
	}
	if muo.mutation.UserCleared() && len(muo.mutation.UserIDs()) > 0 {
		return errors.New(`ent: clearing a required unique edge "Membership.user"`)
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (muo *MembershipUpdateOne) Modify(modifiers ...func(u *sql.UpdateBuilder)) *MembershipUpdateOne {
	muo.modifiers = append(muo.modifiers, modifiers...)
	return muo
}

func (muo *MembershipUpdateOne) sqlSave(ctx context.Context) (_node *Membership, err error) {
	if err := muo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(membership.Table, membership.Columns, sqlgraph.NewFieldSpec(membership.FieldID, field.TypeUUID))
	id, ok := muo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Membership.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := muo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, membership.FieldID)
		for _, f := range fields {
			if !membership.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != membership.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := muo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := muo.mutation.Current(); ok {
		_spec.SetField(membership.FieldCurrent, field.TypeBool, value)
	}
	if value, ok := muo.mutation.UpdatedAt(); ok {
		_spec.SetField(membership.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := muo.mutation.Role(); ok {
		_spec.SetField(membership.FieldRole, field.TypeEnum, value)
	}
	if muo.mutation.OrganizationCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   membership.OrganizationTable,
			Columns: []string{membership.OrganizationColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(organization.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := muo.mutation.OrganizationIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   membership.OrganizationTable,
			Columns: []string{membership.OrganizationColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(organization.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if muo.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   membership.UserTable,
			Columns: []string{membership.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := muo.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   membership.UserTable,
			Columns: []string{membership.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_spec.AddModifiers(muo.modifiers...)
	_node = &Membership{config: muo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, muo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{membership.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	muo.mutation.done = true
	return _node, nil
}
