// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/chainloop-dev/chainloop/app/controlplane/pkg/authz"
	"github.com/chainloop-dev/chainloop/app/controlplane/pkg/biz"
	"github.com/chainloop-dev/chainloop/app/controlplane/pkg/data/ent/organization"
	"github.com/chainloop-dev/chainloop/app/controlplane/pkg/data/ent/orginvitation"
	"github.com/chainloop-dev/chainloop/app/controlplane/pkg/data/ent/user"
	"github.com/google/uuid"
)

// OrgInvitationCreate is the builder for creating a OrgInvitation entity.
type OrgInvitationCreate struct {
	config
	mutation *OrgInvitationMutation
	hooks    []Hook
}

// SetReceiverEmail sets the "receiver_email" field.
func (oic *OrgInvitationCreate) SetReceiverEmail(s string) *OrgInvitationCreate {
	oic.mutation.SetReceiverEmail(s)
	return oic
}

// SetStatus sets the "status" field.
func (oic *OrgInvitationCreate) SetStatus(bis biz.OrgInvitationStatus) *OrgInvitationCreate {
	oic.mutation.SetStatus(bis)
	return oic
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (oic *OrgInvitationCreate) SetNillableStatus(bis *biz.OrgInvitationStatus) *OrgInvitationCreate {
	if bis != nil {
		oic.SetStatus(*bis)
	}
	return oic
}

// SetCreatedAt sets the "created_at" field.
func (oic *OrgInvitationCreate) SetCreatedAt(t time.Time) *OrgInvitationCreate {
	oic.mutation.SetCreatedAt(t)
	return oic
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (oic *OrgInvitationCreate) SetNillableCreatedAt(t *time.Time) *OrgInvitationCreate {
	if t != nil {
		oic.SetCreatedAt(*t)
	}
	return oic
}

// SetDeletedAt sets the "deleted_at" field.
func (oic *OrgInvitationCreate) SetDeletedAt(t time.Time) *OrgInvitationCreate {
	oic.mutation.SetDeletedAt(t)
	return oic
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (oic *OrgInvitationCreate) SetNillableDeletedAt(t *time.Time) *OrgInvitationCreate {
	if t != nil {
		oic.SetDeletedAt(*t)
	}
	return oic
}

// SetOrganizationID sets the "organization_id" field.
func (oic *OrgInvitationCreate) SetOrganizationID(u uuid.UUID) *OrgInvitationCreate {
	oic.mutation.SetOrganizationID(u)
	return oic
}

// SetSenderID sets the "sender_id" field.
func (oic *OrgInvitationCreate) SetSenderID(u uuid.UUID) *OrgInvitationCreate {
	oic.mutation.SetSenderID(u)
	return oic
}

// SetRole sets the "role" field.
func (oic *OrgInvitationCreate) SetRole(a authz.Role) *OrgInvitationCreate {
	oic.mutation.SetRole(a)
	return oic
}

// SetNillableRole sets the "role" field if the given value is not nil.
func (oic *OrgInvitationCreate) SetNillableRole(a *authz.Role) *OrgInvitationCreate {
	if a != nil {
		oic.SetRole(*a)
	}
	return oic
}

// SetID sets the "id" field.
func (oic *OrgInvitationCreate) SetID(u uuid.UUID) *OrgInvitationCreate {
	oic.mutation.SetID(u)
	return oic
}

// SetNillableID sets the "id" field if the given value is not nil.
func (oic *OrgInvitationCreate) SetNillableID(u *uuid.UUID) *OrgInvitationCreate {
	if u != nil {
		oic.SetID(*u)
	}
	return oic
}

// SetOrganization sets the "organization" edge to the Organization entity.
func (oic *OrgInvitationCreate) SetOrganization(o *Organization) *OrgInvitationCreate {
	return oic.SetOrganizationID(o.ID)
}

// SetSender sets the "sender" edge to the User entity.
func (oic *OrgInvitationCreate) SetSender(u *User) *OrgInvitationCreate {
	return oic.SetSenderID(u.ID)
}

// Mutation returns the OrgInvitationMutation object of the builder.
func (oic *OrgInvitationCreate) Mutation() *OrgInvitationMutation {
	return oic.mutation
}

// Save creates the OrgInvitation in the database.
func (oic *OrgInvitationCreate) Save(ctx context.Context) (*OrgInvitation, error) {
	oic.defaults()
	return withHooks(ctx, oic.sqlSave, oic.mutation, oic.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (oic *OrgInvitationCreate) SaveX(ctx context.Context) *OrgInvitation {
	v, err := oic.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (oic *OrgInvitationCreate) Exec(ctx context.Context) error {
	_, err := oic.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (oic *OrgInvitationCreate) ExecX(ctx context.Context) {
	if err := oic.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (oic *OrgInvitationCreate) defaults() {
	if _, ok := oic.mutation.Status(); !ok {
		v := orginvitation.DefaultStatus
		oic.mutation.SetStatus(v)
	}
	if _, ok := oic.mutation.CreatedAt(); !ok {
		v := orginvitation.DefaultCreatedAt()
		oic.mutation.SetCreatedAt(v)
	}
	if _, ok := oic.mutation.ID(); !ok {
		v := orginvitation.DefaultID()
		oic.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (oic *OrgInvitationCreate) check() error {
	if _, ok := oic.mutation.ReceiverEmail(); !ok {
		return &ValidationError{Name: "receiver_email", err: errors.New(`ent: missing required field "OrgInvitation.receiver_email"`)}
	}
	if _, ok := oic.mutation.Status(); !ok {
		return &ValidationError{Name: "status", err: errors.New(`ent: missing required field "OrgInvitation.status"`)}
	}
	if v, ok := oic.mutation.Status(); ok {
		if err := orginvitation.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`ent: validator failed for field "OrgInvitation.status": %w`, err)}
		}
	}
	if _, ok := oic.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "OrgInvitation.created_at"`)}
	}
	if _, ok := oic.mutation.OrganizationID(); !ok {
		return &ValidationError{Name: "organization_id", err: errors.New(`ent: missing required field "OrgInvitation.organization_id"`)}
	}
	if _, ok := oic.mutation.SenderID(); !ok {
		return &ValidationError{Name: "sender_id", err: errors.New(`ent: missing required field "OrgInvitation.sender_id"`)}
	}
	if v, ok := oic.mutation.Role(); ok {
		if err := orginvitation.RoleValidator(v); err != nil {
			return &ValidationError{Name: "role", err: fmt.Errorf(`ent: validator failed for field "OrgInvitation.role": %w`, err)}
		}
	}
	if len(oic.mutation.OrganizationIDs()) == 0 {
		return &ValidationError{Name: "organization", err: errors.New(`ent: missing required edge "OrgInvitation.organization"`)}
	}
	if len(oic.mutation.SenderIDs()) == 0 {
		return &ValidationError{Name: "sender", err: errors.New(`ent: missing required edge "OrgInvitation.sender"`)}
	}
	return nil
}

func (oic *OrgInvitationCreate) sqlSave(ctx context.Context) (*OrgInvitation, error) {
	if err := oic.check(); err != nil {
		return nil, err
	}
	_node, _spec := oic.createSpec()
	if err := sqlgraph.CreateNode(ctx, oic.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(*uuid.UUID); ok {
			_node.ID = *id
		} else if err := _node.ID.Scan(_spec.ID.Value); err != nil {
			return nil, err
		}
	}
	oic.mutation.id = &_node.ID
	oic.mutation.done = true
	return _node, nil
}

func (oic *OrgInvitationCreate) createSpec() (*OrgInvitation, *sqlgraph.CreateSpec) {
	var (
		_node = &OrgInvitation{config: oic.config}
		_spec = sqlgraph.NewCreateSpec(orginvitation.Table, sqlgraph.NewFieldSpec(orginvitation.FieldID, field.TypeUUID))
	)
	if id, ok := oic.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := oic.mutation.ReceiverEmail(); ok {
		_spec.SetField(orginvitation.FieldReceiverEmail, field.TypeString, value)
		_node.ReceiverEmail = value
	}
	if value, ok := oic.mutation.Status(); ok {
		_spec.SetField(orginvitation.FieldStatus, field.TypeEnum, value)
		_node.Status = value
	}
	if value, ok := oic.mutation.CreatedAt(); ok {
		_spec.SetField(orginvitation.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := oic.mutation.DeletedAt(); ok {
		_spec.SetField(orginvitation.FieldDeletedAt, field.TypeTime, value)
		_node.DeletedAt = value
	}
	if value, ok := oic.mutation.Role(); ok {
		_spec.SetField(orginvitation.FieldRole, field.TypeEnum, value)
		_node.Role = value
	}
	if nodes := oic.mutation.OrganizationIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   orginvitation.OrganizationTable,
			Columns: []string{orginvitation.OrganizationColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(organization.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.OrganizationID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := oic.mutation.SenderIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   orginvitation.SenderTable,
			Columns: []string{orginvitation.SenderColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.SenderID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// OrgInvitationCreateBulk is the builder for creating many OrgInvitation entities in bulk.
type OrgInvitationCreateBulk struct {
	config
	err      error
	builders []*OrgInvitationCreate
}

// Save creates the OrgInvitation entities in the database.
func (oicb *OrgInvitationCreateBulk) Save(ctx context.Context) ([]*OrgInvitation, error) {
	if oicb.err != nil {
		return nil, oicb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(oicb.builders))
	nodes := make([]*OrgInvitation, len(oicb.builders))
	mutators := make([]Mutator, len(oicb.builders))
	for i := range oicb.builders {
		func(i int, root context.Context) {
			builder := oicb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*OrgInvitationMutation)
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
					_, err = mutators[i+1].Mutate(root, oicb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, oicb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
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
		if _, err := mutators[0].Mutate(ctx, oicb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (oicb *OrgInvitationCreateBulk) SaveX(ctx context.Context) []*OrgInvitation {
	v, err := oicb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (oicb *OrgInvitationCreateBulk) Exec(ctx context.Context) error {
	_, err := oicb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (oicb *OrgInvitationCreateBulk) ExecX(ctx context.Context) {
	if err := oicb.Exec(ctx); err != nil {
		panic(err)
	}
}
