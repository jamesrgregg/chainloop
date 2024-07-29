// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/chainloop-dev/chainloop/app/controlplane/pkg/data/ent/integrationattachment"
	"github.com/chainloop-dev/chainloop/app/controlplane/pkg/data/ent/organization"
	"github.com/chainloop-dev/chainloop/app/controlplane/pkg/data/ent/referrer"
	"github.com/chainloop-dev/chainloop/app/controlplane/pkg/data/ent/robotaccount"
	"github.com/chainloop-dev/chainloop/app/controlplane/pkg/data/ent/workflow"
	"github.com/chainloop-dev/chainloop/app/controlplane/pkg/data/ent/workflowcontract"
	"github.com/chainloop-dev/chainloop/app/controlplane/pkg/data/ent/workflowrun"
	"github.com/google/uuid"
)

// WorkflowCreate is the builder for creating a Workflow entity.
type WorkflowCreate struct {
	config
	mutation *WorkflowMutation
	hooks    []Hook
}

// SetName sets the "name" field.
func (wc *WorkflowCreate) SetName(s string) *WorkflowCreate {
	wc.mutation.SetName(s)
	return wc
}

// SetProject sets the "project" field.
func (wc *WorkflowCreate) SetProject(s string) *WorkflowCreate {
	wc.mutation.SetProject(s)
	return wc
}

// SetNillableProject sets the "project" field if the given value is not nil.
func (wc *WorkflowCreate) SetNillableProject(s *string) *WorkflowCreate {
	if s != nil {
		wc.SetProject(*s)
	}
	return wc
}

// SetTeam sets the "team" field.
func (wc *WorkflowCreate) SetTeam(s string) *WorkflowCreate {
	wc.mutation.SetTeam(s)
	return wc
}

// SetNillableTeam sets the "team" field if the given value is not nil.
func (wc *WorkflowCreate) SetNillableTeam(s *string) *WorkflowCreate {
	if s != nil {
		wc.SetTeam(*s)
	}
	return wc
}

// SetRunsCount sets the "runs_count" field.
func (wc *WorkflowCreate) SetRunsCount(i int) *WorkflowCreate {
	wc.mutation.SetRunsCount(i)
	return wc
}

// SetNillableRunsCount sets the "runs_count" field if the given value is not nil.
func (wc *WorkflowCreate) SetNillableRunsCount(i *int) *WorkflowCreate {
	if i != nil {
		wc.SetRunsCount(*i)
	}
	return wc
}

// SetCreatedAt sets the "created_at" field.
func (wc *WorkflowCreate) SetCreatedAt(t time.Time) *WorkflowCreate {
	wc.mutation.SetCreatedAt(t)
	return wc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (wc *WorkflowCreate) SetNillableCreatedAt(t *time.Time) *WorkflowCreate {
	if t != nil {
		wc.SetCreatedAt(*t)
	}
	return wc
}

// SetDeletedAt sets the "deleted_at" field.
func (wc *WorkflowCreate) SetDeletedAt(t time.Time) *WorkflowCreate {
	wc.mutation.SetDeletedAt(t)
	return wc
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (wc *WorkflowCreate) SetNillableDeletedAt(t *time.Time) *WorkflowCreate {
	if t != nil {
		wc.SetDeletedAt(*t)
	}
	return wc
}

// SetPublic sets the "public" field.
func (wc *WorkflowCreate) SetPublic(b bool) *WorkflowCreate {
	wc.mutation.SetPublic(b)
	return wc
}

// SetNillablePublic sets the "public" field if the given value is not nil.
func (wc *WorkflowCreate) SetNillablePublic(b *bool) *WorkflowCreate {
	if b != nil {
		wc.SetPublic(*b)
	}
	return wc
}

// SetOrganizationID sets the "organization_id" field.
func (wc *WorkflowCreate) SetOrganizationID(u uuid.UUID) *WorkflowCreate {
	wc.mutation.SetOrganizationID(u)
	return wc
}

// SetDescription sets the "description" field.
func (wc *WorkflowCreate) SetDescription(s string) *WorkflowCreate {
	wc.mutation.SetDescription(s)
	return wc
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (wc *WorkflowCreate) SetNillableDescription(s *string) *WorkflowCreate {
	if s != nil {
		wc.SetDescription(*s)
	}
	return wc
}

// SetID sets the "id" field.
func (wc *WorkflowCreate) SetID(u uuid.UUID) *WorkflowCreate {
	wc.mutation.SetID(u)
	return wc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (wc *WorkflowCreate) SetNillableID(u *uuid.UUID) *WorkflowCreate {
	if u != nil {
		wc.SetID(*u)
	}
	return wc
}

// AddRobotaccountIDs adds the "robotaccounts" edge to the RobotAccount entity by IDs.
func (wc *WorkflowCreate) AddRobotaccountIDs(ids ...uuid.UUID) *WorkflowCreate {
	wc.mutation.AddRobotaccountIDs(ids...)
	return wc
}

// AddRobotaccounts adds the "robotaccounts" edges to the RobotAccount entity.
func (wc *WorkflowCreate) AddRobotaccounts(r ...*RobotAccount) *WorkflowCreate {
	ids := make([]uuid.UUID, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return wc.AddRobotaccountIDs(ids...)
}

// AddWorkflowrunIDs adds the "workflowruns" edge to the WorkflowRun entity by IDs.
func (wc *WorkflowCreate) AddWorkflowrunIDs(ids ...uuid.UUID) *WorkflowCreate {
	wc.mutation.AddWorkflowrunIDs(ids...)
	return wc
}

// AddWorkflowruns adds the "workflowruns" edges to the WorkflowRun entity.
func (wc *WorkflowCreate) AddWorkflowruns(w ...*WorkflowRun) *WorkflowCreate {
	ids := make([]uuid.UUID, len(w))
	for i := range w {
		ids[i] = w[i].ID
	}
	return wc.AddWorkflowrunIDs(ids...)
}

// SetOrganization sets the "organization" edge to the Organization entity.
func (wc *WorkflowCreate) SetOrganization(o *Organization) *WorkflowCreate {
	return wc.SetOrganizationID(o.ID)
}

// SetContractID sets the "contract" edge to the WorkflowContract entity by ID.
func (wc *WorkflowCreate) SetContractID(id uuid.UUID) *WorkflowCreate {
	wc.mutation.SetContractID(id)
	return wc
}

// SetContract sets the "contract" edge to the WorkflowContract entity.
func (wc *WorkflowCreate) SetContract(w *WorkflowContract) *WorkflowCreate {
	return wc.SetContractID(w.ID)
}

// AddIntegrationAttachmentIDs adds the "integration_attachments" edge to the IntegrationAttachment entity by IDs.
func (wc *WorkflowCreate) AddIntegrationAttachmentIDs(ids ...uuid.UUID) *WorkflowCreate {
	wc.mutation.AddIntegrationAttachmentIDs(ids...)
	return wc
}

// AddIntegrationAttachments adds the "integration_attachments" edges to the IntegrationAttachment entity.
func (wc *WorkflowCreate) AddIntegrationAttachments(i ...*IntegrationAttachment) *WorkflowCreate {
	ids := make([]uuid.UUID, len(i))
	for j := range i {
		ids[j] = i[j].ID
	}
	return wc.AddIntegrationAttachmentIDs(ids...)
}

// AddReferrerIDs adds the "referrers" edge to the Referrer entity by IDs.
func (wc *WorkflowCreate) AddReferrerIDs(ids ...uuid.UUID) *WorkflowCreate {
	wc.mutation.AddReferrerIDs(ids...)
	return wc
}

// AddReferrers adds the "referrers" edges to the Referrer entity.
func (wc *WorkflowCreate) AddReferrers(r ...*Referrer) *WorkflowCreate {
	ids := make([]uuid.UUID, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return wc.AddReferrerIDs(ids...)
}

// Mutation returns the WorkflowMutation object of the builder.
func (wc *WorkflowCreate) Mutation() *WorkflowMutation {
	return wc.mutation
}

// Save creates the Workflow in the database.
func (wc *WorkflowCreate) Save(ctx context.Context) (*Workflow, error) {
	wc.defaults()
	return withHooks(ctx, wc.sqlSave, wc.mutation, wc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (wc *WorkflowCreate) SaveX(ctx context.Context) *Workflow {
	v, err := wc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (wc *WorkflowCreate) Exec(ctx context.Context) error {
	_, err := wc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (wc *WorkflowCreate) ExecX(ctx context.Context) {
	if err := wc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (wc *WorkflowCreate) defaults() {
	if _, ok := wc.mutation.RunsCount(); !ok {
		v := workflow.DefaultRunsCount
		wc.mutation.SetRunsCount(v)
	}
	if _, ok := wc.mutation.CreatedAt(); !ok {
		v := workflow.DefaultCreatedAt()
		wc.mutation.SetCreatedAt(v)
	}
	if _, ok := wc.mutation.Public(); !ok {
		v := workflow.DefaultPublic
		wc.mutation.SetPublic(v)
	}
	if _, ok := wc.mutation.ID(); !ok {
		v := workflow.DefaultID()
		wc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (wc *WorkflowCreate) check() error {
	if _, ok := wc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "Workflow.name"`)}
	}
	if _, ok := wc.mutation.RunsCount(); !ok {
		return &ValidationError{Name: "runs_count", err: errors.New(`ent: missing required field "Workflow.runs_count"`)}
	}
	if _, ok := wc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Workflow.created_at"`)}
	}
	if _, ok := wc.mutation.Public(); !ok {
		return &ValidationError{Name: "public", err: errors.New(`ent: missing required field "Workflow.public"`)}
	}
	if _, ok := wc.mutation.OrganizationID(); !ok {
		return &ValidationError{Name: "organization_id", err: errors.New(`ent: missing required field "Workflow.organization_id"`)}
	}
	if len(wc.mutation.OrganizationIDs()) == 0 {
		return &ValidationError{Name: "organization", err: errors.New(`ent: missing required edge "Workflow.organization"`)}
	}
	if len(wc.mutation.ContractIDs()) == 0 {
		return &ValidationError{Name: "contract", err: errors.New(`ent: missing required edge "Workflow.contract"`)}
	}
	return nil
}

func (wc *WorkflowCreate) sqlSave(ctx context.Context) (*Workflow, error) {
	if err := wc.check(); err != nil {
		return nil, err
	}
	_node, _spec := wc.createSpec()
	if err := sqlgraph.CreateNode(ctx, wc.driver, _spec); err != nil {
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
	wc.mutation.id = &_node.ID
	wc.mutation.done = true
	return _node, nil
}

func (wc *WorkflowCreate) createSpec() (*Workflow, *sqlgraph.CreateSpec) {
	var (
		_node = &Workflow{config: wc.config}
		_spec = sqlgraph.NewCreateSpec(workflow.Table, sqlgraph.NewFieldSpec(workflow.FieldID, field.TypeUUID))
	)
	if id, ok := wc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := wc.mutation.Name(); ok {
		_spec.SetField(workflow.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := wc.mutation.Project(); ok {
		_spec.SetField(workflow.FieldProject, field.TypeString, value)
		_node.Project = value
	}
	if value, ok := wc.mutation.Team(); ok {
		_spec.SetField(workflow.FieldTeam, field.TypeString, value)
		_node.Team = value
	}
	if value, ok := wc.mutation.RunsCount(); ok {
		_spec.SetField(workflow.FieldRunsCount, field.TypeInt, value)
		_node.RunsCount = value
	}
	if value, ok := wc.mutation.CreatedAt(); ok {
		_spec.SetField(workflow.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := wc.mutation.DeletedAt(); ok {
		_spec.SetField(workflow.FieldDeletedAt, field.TypeTime, value)
		_node.DeletedAt = value
	}
	if value, ok := wc.mutation.Public(); ok {
		_spec.SetField(workflow.FieldPublic, field.TypeBool, value)
		_node.Public = value
	}
	if value, ok := wc.mutation.Description(); ok {
		_spec.SetField(workflow.FieldDescription, field.TypeString, value)
		_node.Description = value
	}
	if nodes := wc.mutation.RobotaccountsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   workflow.RobotaccountsTable,
			Columns: []string{workflow.RobotaccountsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(robotaccount.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := wc.mutation.WorkflowrunsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   workflow.WorkflowrunsTable,
			Columns: []string{workflow.WorkflowrunsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(workflowrun.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := wc.mutation.OrganizationIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   workflow.OrganizationTable,
			Columns: []string{workflow.OrganizationColumn},
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
	if nodes := wc.mutation.ContractIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   workflow.ContractTable,
			Columns: []string{workflow.ContractColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(workflowcontract.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.workflow_contract = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := wc.mutation.IntegrationAttachmentsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   workflow.IntegrationAttachmentsTable,
			Columns: []string{workflow.IntegrationAttachmentsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(integrationattachment.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := wc.mutation.ReferrersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   workflow.ReferrersTable,
			Columns: workflow.ReferrersPrimaryKey,
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
	return _node, _spec
}

// WorkflowCreateBulk is the builder for creating many Workflow entities in bulk.
type WorkflowCreateBulk struct {
	config
	err      error
	builders []*WorkflowCreate
}

// Save creates the Workflow entities in the database.
func (wcb *WorkflowCreateBulk) Save(ctx context.Context) ([]*Workflow, error) {
	if wcb.err != nil {
		return nil, wcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(wcb.builders))
	nodes := make([]*Workflow, len(wcb.builders))
	mutators := make([]Mutator, len(wcb.builders))
	for i := range wcb.builders {
		func(i int, root context.Context) {
			builder := wcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*WorkflowMutation)
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
					_, err = mutators[i+1].Mutate(root, wcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, wcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, wcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (wcb *WorkflowCreateBulk) SaveX(ctx context.Context) []*Workflow {
	v, err := wcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (wcb *WorkflowCreateBulk) Exec(ctx context.Context) error {
	_, err := wcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (wcb *WorkflowCreateBulk) ExecX(ctx context.Context) {
	if err := wcb.Exec(ctx); err != nil {
		panic(err)
	}
}
