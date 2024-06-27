//
// Copyright 2024 The Chainloop Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package data

import (
	"bytes"
	"context"
	"fmt"
	"time"

	schemav1 "github.com/chainloop-dev/chainloop/app/controlplane/api/workflowcontract/v1"
	"github.com/chainloop-dev/chainloop/app/controlplane/pkg/biz"
	"github.com/chainloop-dev/chainloop/app/controlplane/pkg/data/ent"
	"github.com/chainloop-dev/chainloop/app/controlplane/pkg/data/ent/organization"
	"github.com/chainloop-dev/chainloop/app/controlplane/pkg/data/ent/workflow"
	"github.com/chainloop-dev/chainloop/app/controlplane/pkg/data/ent/workflowcontract"
	"github.com/chainloop-dev/chainloop/app/controlplane/pkg/data/ent/workflowcontractversion"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/uuid"
	"google.golang.org/protobuf/proto"
)

type WorkflowContractRepo struct {
	data *Data
	log  *log.Helper
}

func NewWorkflowContractRepo(data *Data, logger log.Logger) biz.WorkflowContractRepo {
	return &WorkflowContractRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *WorkflowContractRepo) List(ctx context.Context, orgID uuid.UUID) ([]*biz.WorkflowContract, error) {
	contracts, err := orgScopedQuery(r.data.DB, orgID).
		QueryWorkflowContracts().
		Where(workflowcontract.DeletedAtIsNil()).
		WithWorkflows(func(q *ent.WorkflowQuery) {
			q.Where(workflow.DeletedAtIsNil())
		}).
		Order(ent.Desc(workflow.FieldCreatedAt)).
		All(ctx)
	if err != nil {
		return nil, err
	}

	result := make([]*biz.WorkflowContract, 0, len(contracts))
	for _, s := range contracts {
		latestV, err := latestVersion(ctx, s)
		if err != nil {
			return nil, err
		}

		workflowIDs, err := getWorkflowIDs(ctx, s)
		if err != nil {
			return nil, err
		}
		res := entContractToBizContract(s, latestV, workflowIDs)

		result = append(result, res)
	}

	return result, nil
}

func (r *WorkflowContractRepo) Create(ctx context.Context, opts *biz.ContractCreateOpts) (*biz.WorkflowContract, error) {
	tx, err := r.data.DB.Tx(ctx)
	if err != nil {
		return nil, err
	}

	contract, err := tx.WorkflowContract.Create().
		SetName(opts.Name).SetOrganizationID(opts.OrgID).
		SetNillableDescription(opts.Description).
		Save(ctx)
	if err != nil {
		return nil, rollback(tx, err)
	}

	// Add version
	version, err := tx.WorkflowContractVersion.Create().SetBody(opts.ContractBody).SetContract(contract).Save(ctx)
	if err != nil {
		return nil, rollback(tx, err)
	}

	if err := tx.Commit(); err != nil {
		return nil, err
	}

	res := entContractToBizContract(contract, version, nil)
	return res, nil
}

func (r *WorkflowContractRepo) FindVersionByID(ctx context.Context, versionID uuid.UUID) (*biz.WorkflowContractVersion, error) {
	version, err := r.data.DB.WorkflowContractVersion.Get(ctx, versionID)
	if err != nil && !ent.IsNotFound(err) {
		return nil, err
	} else if version == nil {
		return nil, nil
	}

	return entContractVersionToBizContractVersion(version)
}

func (r *WorkflowContractRepo) Describe(ctx context.Context, orgID, contractID uuid.UUID, revision int) (*biz.WorkflowContractWithVersion, error) {
	contract, err := contractInOrg(ctx, r.data.DB, orgID, &contractID, nil)
	if err != nil && !ent.IsNotFound(err) {
		return nil, err
	} else if contract == nil {
		return nil, nil
	}

	latestV, err := latestVersion(ctx, contract)
	if err != nil {
		return nil, err
	}

	// revision 0 means latest
	version := latestV
	if revision != 0 {
		version, err = contract.QueryVersions().Where(workflowcontractversion.RevisionEQ(revision)).Only(ctx)
		if err != nil && !ent.IsNotFound(err) {
			return nil, err
		} else if version == nil {
			return nil, nil
		}
	}

	v, err := entContractVersionToBizContractVersion(version)
	if err != nil {
		return nil, err
	}

	workflowIDs, err := getWorkflowIDs(ctx, contract)
	if err != nil {
		return nil, err
	}
	s := entContractToBizContract(contract, latestV, workflowIDs)

	return &biz.WorkflowContractWithVersion{
		Contract: s,
		Version:  v,
	}, nil
}

// Update will add a new version of the contract.
// NOTE: ContractVersions are immutable
func (r *WorkflowContractRepo) Update(ctx context.Context, orgID uuid.UUID, name string, opts *biz.ContractUpdateOpts) (*biz.WorkflowContractWithVersion, error) {
	tx, err := r.data.DB.Debug().Tx(ctx)
	if err != nil {
		return nil, err
	}

	contract, err := contractInOrgTx(ctx, tx, orgID, nil, &name)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, nil
		}

		return nil, rollback(tx, err)
	}

	if contract == nil {
		return nil, nil
	}

	contract, err = contract.Update().SetNillableDescription(opts.Description).Save(ctx)
	if err != nil {
		return nil, rollback(tx, err)
	}

	lv, err := latestVersion(ctx, contract)
	if err != nil {
		return nil, err
	}

	// Create a revision only if we are providing a new contract and it has changed
	if opts.ContractBody != nil && !bytes.Equal(lv.Body, opts.ContractBody) {
		// TODO: Add pessimist locking to make sure we are incrementing the latest revision
		lv, err = tx.WorkflowContractVersion.Create().
			SetBody(opts.ContractBody).
			SetContract(contract).
			SetRevision(lv.Revision + 1).
			Save(ctx)
		if err != nil {
			return nil, rollback(tx, err)
		}
	}

	workflowIDs, err := getWorkflowIDs(ctx, contract)
	if err != nil {
		return nil, err
	}

	if err := tx.Commit(); err != nil {
		return nil, err
	}

	// The transaction is committed, we can now return the result
	contract, err = contractInOrg(ctx, r.data.DB, orgID, nil, &name)
	if err != nil {
		return nil, err
	}

	v, err := entContractVersionToBizContractVersion(lv)
	if err != nil {
		return nil, err
	}

	return &biz.WorkflowContractWithVersion{
		Contract: entContractToBizContract(contract, lv, workflowIDs),
		Version:  v,
	}, nil
}

func (r *WorkflowContractRepo) FindByIDInOrg(ctx context.Context, orgID, contractID uuid.UUID) (*biz.WorkflowContract, error) {
	contract, err := contractInOrg(ctx, r.data.DB, orgID, &contractID, nil)
	if err != nil && !ent.IsNotFound(err) {
		return nil, err
	} else if contract == nil {
		return nil, nil
	}

	workflowIDs, err := getWorkflowIDs(ctx, contract)
	if err != nil {
		return nil, err
	}

	latestV, err := latestVersion(ctx, contract)
	if err != nil {
		return nil, err
	}

	return entContractToBizContract(contract, latestV, workflowIDs), nil
}

func (r *WorkflowContractRepo) FindByNameInOrg(ctx context.Context, orgID uuid.UUID, name string) (*biz.WorkflowContract, error) {
	contract, err := contractInOrg(ctx, r.data.DB, orgID, nil, &name)
	if err != nil && !ent.IsNotFound(err) {
		return nil, fmt.Errorf("failed to find contract: %w", err)
	} else if contract == nil {
		return nil, biz.NewErrNotFound("contract")
	}

	workflowIDs, err := getWorkflowIDs(ctx, contract)
	if err != nil {
		return nil, fmt.Errorf("failed to get workflows: %w", err)
	}

	latestV, err := latestVersion(ctx, contract)
	if err != nil {
		return nil, fmt.Errorf("failed to get latest version: %w", err)
	}

	return entContractToBizContract(contract, latestV, workflowIDs), nil
}

func (r *WorkflowContractRepo) SoftDelete(ctx context.Context, id uuid.UUID) error {
	return r.data.DB.WorkflowContract.UpdateOneID(id).SetDeletedAt(time.Now()).Exec(ctx)
}

func entContractVersionToBizContractVersion(w *ent.WorkflowContractVersion) (*biz.WorkflowContractVersion, error) {
	contract := &schemav1.CraftingSchema{}
	err := proto.Unmarshal(w.Body, contract)
	if err != nil {
		return nil, err
	}

	return &biz.WorkflowContractVersion{
		ID: w.ID, CreatedAt: toTimePtr(w.CreatedAt), Revision: w.Revision, BodyV1: contract,
	}, nil
}

// rollback calls to tx.Rollback and wraps the given error
// with the rollback error if occurred.
func rollback(tx *ent.Tx, err error) error {
	if rerr := tx.Rollback(); rerr != nil {
		err = fmt.Errorf("%w: %w", err, rerr)
	}

	// If the error is a constraint error, we return a more specific error to indicate the client that's a duplicate
	if ent.IsConstraintError(err) {
		return biz.ErrAlreadyExists
	}

	return err
}

func latestVersion(ctx context.Context, contract *ent.WorkflowContract) (*ent.WorkflowContractVersion, error) {
	return contract.QueryVersions().Order(ent.Desc(workflowcontractversion.FieldRevision)).First(ctx)
}

func contractInOrg(ctx context.Context, client *ent.Client, orgID uuid.UUID, contractID *uuid.UUID, name *string) (*ent.WorkflowContract, error) {
	return contractInOrgQuery(ctx, client.Organization.Query(), orgID, contractID, name)
}

func contractInOrgTx(ctx context.Context, tx *ent.Tx, orgID uuid.UUID, contractID *uuid.UUID, name *string) (*ent.WorkflowContract, error) {
	return contractInOrgQuery(ctx, tx.Organization.Query(), orgID, contractID, name)
}

// It can be loaded via by ID or name
func contractInOrgQuery(ctx context.Context, q *ent.OrganizationQuery, orgID uuid.UUID, contractID *uuid.UUID, name *string) (*ent.WorkflowContract, error) {
	if contractID == nil && name == nil {
		return nil, fmt.Errorf("either contractID or name must be provided")
	}

	query := q.
		Where(organization.ID(orgID)).
		QueryWorkflowContracts().
		Where(workflowcontract.DeletedAtIsNil()).
		WithWorkflows(func(q *ent.WorkflowQuery) {
			q.Where(workflow.DeletedAtIsNil())
		})

	if contractID != nil {
		query = query.Where(workflowcontract.ID(*contractID))
	}

	if name != nil {
		query = query.Where(workflowcontract.NameEQ(*name))
	}

	return query.Only(ctx)
}

func entContractToBizContract(w *ent.WorkflowContract, version *ent.WorkflowContractVersion, workflowIDs []string) *biz.WorkflowContract {
	c := &biz.WorkflowContract{
		Name: w.Name, ID: w.ID, CreatedAt: toTimePtr(w.CreatedAt), WorkflowIDs: workflowIDs, Description: w.Description,
	}

	c.LatestRevision = version.Revision
	return c
}

// get the list of workflows associated with a given contract
func getWorkflowIDs(ctx context.Context, schema *ent.WorkflowContract) ([]string, error) {
	// Either get it from preloaded entity or query it
	workflows := schema.Edges.Workflows
	if workflows == nil {
		var err error
		workflows, err = schema.QueryWorkflows().
			Where(workflow.DeletedAtIsNil()).
			Select(workflowcontract.FieldID).All(ctx)
		if err != nil {
			return nil, err
		}
	}

	wfIDs := make([]string, 0, len(workflows))
	for _, wf := range workflows {
		wfIDs = append(wfIDs, wf.ID.String())
	}

	return wfIDs, nil
}
