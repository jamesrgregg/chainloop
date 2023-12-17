// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/chainloop-dev/chainloop/app/controlplane/internal/data/ent/referrer"
	"github.com/chainloop-dev/chainloop/app/controlplane/internal/data/ent/workflow"
	"github.com/google/uuid"
)

// ReferrerCreate is the builder for creating a Referrer entity.
type ReferrerCreate struct {
	config
	mutation *ReferrerMutation
	hooks    []Hook
}

// SetDigest sets the "digest" field.
func (rc *ReferrerCreate) SetDigest(s string) *ReferrerCreate {
	rc.mutation.SetDigest(s)
	return rc
}

// SetKind sets the "kind" field.
func (rc *ReferrerCreate) SetKind(s string) *ReferrerCreate {
	rc.mutation.SetKind(s)
	return rc
}

// SetDownloadable sets the "downloadable" field.
func (rc *ReferrerCreate) SetDownloadable(b bool) *ReferrerCreate {
	rc.mutation.SetDownloadable(b)
	return rc
}

// SetCreatedAt sets the "created_at" field.
func (rc *ReferrerCreate) SetCreatedAt(t time.Time) *ReferrerCreate {
	rc.mutation.SetCreatedAt(t)
	return rc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (rc *ReferrerCreate) SetNillableCreatedAt(t *time.Time) *ReferrerCreate {
	if t != nil {
		rc.SetCreatedAt(*t)
	}
	return rc
}

// SetMetadata sets the "metadata" field.
func (rc *ReferrerCreate) SetMetadata(m map[string]string) *ReferrerCreate {
	rc.mutation.SetMetadata(m)
	return rc
}

// SetAnnotations sets the "annotations" field.
func (rc *ReferrerCreate) SetAnnotations(m map[string]string) *ReferrerCreate {
	rc.mutation.SetAnnotations(m)
	return rc
}

// SetID sets the "id" field.
func (rc *ReferrerCreate) SetID(u uuid.UUID) *ReferrerCreate {
	rc.mutation.SetID(u)
	return rc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (rc *ReferrerCreate) SetNillableID(u *uuid.UUID) *ReferrerCreate {
	if u != nil {
		rc.SetID(*u)
	}
	return rc
}

// AddReferredByIDs adds the "referred_by" edge to the Referrer entity by IDs.
func (rc *ReferrerCreate) AddReferredByIDs(ids ...uuid.UUID) *ReferrerCreate {
	rc.mutation.AddReferredByIDs(ids...)
	return rc
}

// AddReferredBy adds the "referred_by" edges to the Referrer entity.
func (rc *ReferrerCreate) AddReferredBy(r ...*Referrer) *ReferrerCreate {
	ids := make([]uuid.UUID, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return rc.AddReferredByIDs(ids...)
}

// AddReferenceIDs adds the "references" edge to the Referrer entity by IDs.
func (rc *ReferrerCreate) AddReferenceIDs(ids ...uuid.UUID) *ReferrerCreate {
	rc.mutation.AddReferenceIDs(ids...)
	return rc
}

// AddReferences adds the "references" edges to the Referrer entity.
func (rc *ReferrerCreate) AddReferences(r ...*Referrer) *ReferrerCreate {
	ids := make([]uuid.UUID, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return rc.AddReferenceIDs(ids...)
}

// AddWorkflowIDs adds the "workflows" edge to the Workflow entity by IDs.
func (rc *ReferrerCreate) AddWorkflowIDs(ids ...uuid.UUID) *ReferrerCreate {
	rc.mutation.AddWorkflowIDs(ids...)
	return rc
}

// AddWorkflows adds the "workflows" edges to the Workflow entity.
func (rc *ReferrerCreate) AddWorkflows(w ...*Workflow) *ReferrerCreate {
	ids := make([]uuid.UUID, len(w))
	for i := range w {
		ids[i] = w[i].ID
	}
	return rc.AddWorkflowIDs(ids...)
}

// Mutation returns the ReferrerMutation object of the builder.
func (rc *ReferrerCreate) Mutation() *ReferrerMutation {
	return rc.mutation
}

// Save creates the Referrer in the database.
func (rc *ReferrerCreate) Save(ctx context.Context) (*Referrer, error) {
	rc.defaults()
	return withHooks(ctx, rc.sqlSave, rc.mutation, rc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (rc *ReferrerCreate) SaveX(ctx context.Context) *Referrer {
	v, err := rc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (rc *ReferrerCreate) Exec(ctx context.Context) error {
	_, err := rc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (rc *ReferrerCreate) ExecX(ctx context.Context) {
	if err := rc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (rc *ReferrerCreate) defaults() {
	if _, ok := rc.mutation.CreatedAt(); !ok {
		v := referrer.DefaultCreatedAt()
		rc.mutation.SetCreatedAt(v)
	}
	if _, ok := rc.mutation.ID(); !ok {
		v := referrer.DefaultID()
		rc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (rc *ReferrerCreate) check() error {
	if _, ok := rc.mutation.Digest(); !ok {
		return &ValidationError{Name: "digest", err: errors.New(`ent: missing required field "Referrer.digest"`)}
	}
	if _, ok := rc.mutation.Kind(); !ok {
		return &ValidationError{Name: "kind", err: errors.New(`ent: missing required field "Referrer.kind"`)}
	}
	if _, ok := rc.mutation.Downloadable(); !ok {
		return &ValidationError{Name: "downloadable", err: errors.New(`ent: missing required field "Referrer.downloadable"`)}
	}
	if _, ok := rc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Referrer.created_at"`)}
	}
	return nil
}

func (rc *ReferrerCreate) sqlSave(ctx context.Context) (*Referrer, error) {
	if err := rc.check(); err != nil {
		return nil, err
	}
	_node, _spec := rc.createSpec()
	if err := sqlgraph.CreateNode(ctx, rc.driver, _spec); err != nil {
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
	rc.mutation.id = &_node.ID
	rc.mutation.done = true
	return _node, nil
}

func (rc *ReferrerCreate) createSpec() (*Referrer, *sqlgraph.CreateSpec) {
	var (
		_node = &Referrer{config: rc.config}
		_spec = sqlgraph.NewCreateSpec(referrer.Table, sqlgraph.NewFieldSpec(referrer.FieldID, field.TypeUUID))
	)
	if id, ok := rc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := rc.mutation.Digest(); ok {
		_spec.SetField(referrer.FieldDigest, field.TypeString, value)
		_node.Digest = value
	}
	if value, ok := rc.mutation.Kind(); ok {
		_spec.SetField(referrer.FieldKind, field.TypeString, value)
		_node.Kind = value
	}
	if value, ok := rc.mutation.Downloadable(); ok {
		_spec.SetField(referrer.FieldDownloadable, field.TypeBool, value)
		_node.Downloadable = value
	}
	if value, ok := rc.mutation.CreatedAt(); ok {
		_spec.SetField(referrer.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := rc.mutation.Metadata(); ok {
		_spec.SetField(referrer.FieldMetadata, field.TypeJSON, value)
		_node.Metadata = value
	}
	if value, ok := rc.mutation.Annotations(); ok {
		_spec.SetField(referrer.FieldAnnotations, field.TypeJSON, value)
		_node.Annotations = value
	}
	if nodes := rc.mutation.ReferredByIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   referrer.ReferredByTable,
			Columns: referrer.ReferredByPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(referrer.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := rc.mutation.ReferencesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   referrer.ReferencesTable,
			Columns: referrer.ReferencesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(referrer.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := rc.mutation.WorkflowsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   referrer.WorkflowsTable,
			Columns: referrer.WorkflowsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(workflow.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// ReferrerCreateBulk is the builder for creating many Referrer entities in bulk.
type ReferrerCreateBulk struct {
	config
	builders []*ReferrerCreate
}

// Save creates the Referrer entities in the database.
func (rcb *ReferrerCreateBulk) Save(ctx context.Context) ([]*Referrer, error) {
	specs := make([]*sqlgraph.CreateSpec, len(rcb.builders))
	nodes := make([]*Referrer, len(rcb.builders))
	mutators := make([]Mutator, len(rcb.builders))
	for i := range rcb.builders {
		func(i int, root context.Context) {
			builder := rcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ReferrerMutation)
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
					_, err = mutators[i+1].Mutate(root, rcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, rcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, rcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (rcb *ReferrerCreateBulk) SaveX(ctx context.Context) []*Referrer {
	v, err := rcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (rcb *ReferrerCreateBulk) Exec(ctx context.Context) error {
	_, err := rcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (rcb *ReferrerCreateBulk) ExecX(ctx context.Context) {
	if err := rcb.Exec(ctx); err != nil {
		panic(err)
	}
}
