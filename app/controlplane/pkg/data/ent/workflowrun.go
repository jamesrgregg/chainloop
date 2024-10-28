// Code generated by ent, DO NOT EDIT.

package ent

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/chainloop-dev/chainloop/app/controlplane/pkg/biz"
	"github.com/chainloop-dev/chainloop/app/controlplane/pkg/data/ent/projectversion"
	"github.com/chainloop-dev/chainloop/app/controlplane/pkg/data/ent/workflow"
	"github.com/chainloop-dev/chainloop/app/controlplane/pkg/data/ent/workflowcontractversion"
	"github.com/chainloop-dev/chainloop/app/controlplane/pkg/data/ent/workflowrun"
	"github.com/google/uuid"
	"github.com/secure-systems-lab/go-securesystemslib/dsse"
)

// WorkflowRun is the model entity for the WorkflowRun schema.
type WorkflowRun struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// FinishedAt holds the value of the "finished_at" field.
	FinishedAt time.Time `json:"finished_at,omitempty"`
	// State holds the value of the "state" field.
	State biz.WorkflowRunStatus `json:"state,omitempty"`
	// Reason holds the value of the "reason" field.
	Reason string `json:"reason,omitempty"`
	// RunURL holds the value of the "run_url" field.
	RunURL string `json:"run_url,omitempty"`
	// RunnerType holds the value of the "runner_type" field.
	RunnerType string `json:"runner_type,omitempty"`
	// Attestation holds the value of the "attestation" field.
	Attestation *dsse.Envelope `json:"attestation,omitempty"`
	// AttestationDigest holds the value of the "attestation_digest" field.
	AttestationDigest string `json:"attestation_digest,omitempty"`
	// AttestationState holds the value of the "attestation_state" field.
	AttestationState []byte `json:"attestation_state,omitempty"`
	// ContractRevisionUsed holds the value of the "contract_revision_used" field.
	ContractRevisionUsed int `json:"contract_revision_used,omitempty"`
	// ContractRevisionLatest holds the value of the "contract_revision_latest" field.
	ContractRevisionLatest int `json:"contract_revision_latest,omitempty"`
	// VersionID holds the value of the "version_id" field.
	VersionID uuid.UUID `json:"version_id,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the WorkflowRunQuery when eager-loading is set.
	Edges                         WorkflowRunEdges `json:"edges"`
	workflow_workflowruns         *uuid.UUID
	workflow_run_contract_version *uuid.UUID
	selectValues                  sql.SelectValues
}

// WorkflowRunEdges holds the relations/edges for other nodes in the graph.
type WorkflowRunEdges struct {
	// Workflow holds the value of the workflow edge.
	Workflow *Workflow `json:"workflow,omitempty"`
	// ContractVersion holds the value of the contract_version edge.
	ContractVersion *WorkflowContractVersion `json:"contract_version,omitempty"`
	// CasBackends holds the value of the cas_backends edge.
	CasBackends []*CASBackend `json:"cas_backends,omitempty"`
	// Version holds the value of the version edge.
	Version *ProjectVersion `json:"version,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [4]bool
}

// WorkflowOrErr returns the Workflow value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e WorkflowRunEdges) WorkflowOrErr() (*Workflow, error) {
	if e.Workflow != nil {
		return e.Workflow, nil
	} else if e.loadedTypes[0] {
		return nil, &NotFoundError{label: workflow.Label}
	}
	return nil, &NotLoadedError{edge: "workflow"}
}

// ContractVersionOrErr returns the ContractVersion value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e WorkflowRunEdges) ContractVersionOrErr() (*WorkflowContractVersion, error) {
	if e.ContractVersion != nil {
		return e.ContractVersion, nil
	} else if e.loadedTypes[1] {
		return nil, &NotFoundError{label: workflowcontractversion.Label}
	}
	return nil, &NotLoadedError{edge: "contract_version"}
}

// CasBackendsOrErr returns the CasBackends value or an error if the edge
// was not loaded in eager-loading.
func (e WorkflowRunEdges) CasBackendsOrErr() ([]*CASBackend, error) {
	if e.loadedTypes[2] {
		return e.CasBackends, nil
	}
	return nil, &NotLoadedError{edge: "cas_backends"}
}

// VersionOrErr returns the Version value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e WorkflowRunEdges) VersionOrErr() (*ProjectVersion, error) {
	if e.Version != nil {
		return e.Version, nil
	} else if e.loadedTypes[3] {
		return nil, &NotFoundError{label: projectversion.Label}
	}
	return nil, &NotLoadedError{edge: "version"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*WorkflowRun) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case workflowrun.FieldAttestation, workflowrun.FieldAttestationState:
			values[i] = new([]byte)
		case workflowrun.FieldContractRevisionUsed, workflowrun.FieldContractRevisionLatest:
			values[i] = new(sql.NullInt64)
		case workflowrun.FieldState, workflowrun.FieldReason, workflowrun.FieldRunURL, workflowrun.FieldRunnerType, workflowrun.FieldAttestationDigest:
			values[i] = new(sql.NullString)
		case workflowrun.FieldCreatedAt, workflowrun.FieldFinishedAt:
			values[i] = new(sql.NullTime)
		case workflowrun.FieldID, workflowrun.FieldVersionID:
			values[i] = new(uuid.UUID)
		case workflowrun.ForeignKeys[0]: // workflow_workflowruns
			values[i] = &sql.NullScanner{S: new(uuid.UUID)}
		case workflowrun.ForeignKeys[1]: // workflow_run_contract_version
			values[i] = &sql.NullScanner{S: new(uuid.UUID)}
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the WorkflowRun fields.
func (wr *WorkflowRun) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case workflowrun.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				wr.ID = *value
			}
		case workflowrun.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				wr.CreatedAt = value.Time
			}
		case workflowrun.FieldFinishedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field finished_at", values[i])
			} else if value.Valid {
				wr.FinishedAt = value.Time
			}
		case workflowrun.FieldState:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field state", values[i])
			} else if value.Valid {
				wr.State = biz.WorkflowRunStatus(value.String)
			}
		case workflowrun.FieldReason:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field reason", values[i])
			} else if value.Valid {
				wr.Reason = value.String
			}
		case workflowrun.FieldRunURL:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field run_url", values[i])
			} else if value.Valid {
				wr.RunURL = value.String
			}
		case workflowrun.FieldRunnerType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field runner_type", values[i])
			} else if value.Valid {
				wr.RunnerType = value.String
			}
		case workflowrun.FieldAttestation:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field attestation", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &wr.Attestation); err != nil {
					return fmt.Errorf("unmarshal field attestation: %w", err)
				}
			}
		case workflowrun.FieldAttestationDigest:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field attestation_digest", values[i])
			} else if value.Valid {
				wr.AttestationDigest = value.String
			}
		case workflowrun.FieldAttestationState:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field attestation_state", values[i])
			} else if value != nil {
				wr.AttestationState = *value
			}
		case workflowrun.FieldContractRevisionUsed:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field contract_revision_used", values[i])
			} else if value.Valid {
				wr.ContractRevisionUsed = int(value.Int64)
			}
		case workflowrun.FieldContractRevisionLatest:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field contract_revision_latest", values[i])
			} else if value.Valid {
				wr.ContractRevisionLatest = int(value.Int64)
			}
		case workflowrun.FieldVersionID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field version_id", values[i])
			} else if value != nil {
				wr.VersionID = *value
			}
		case workflowrun.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullScanner); !ok {
				return fmt.Errorf("unexpected type %T for field workflow_workflowruns", values[i])
			} else if value.Valid {
				wr.workflow_workflowruns = new(uuid.UUID)
				*wr.workflow_workflowruns = *value.S.(*uuid.UUID)
			}
		case workflowrun.ForeignKeys[1]:
			if value, ok := values[i].(*sql.NullScanner); !ok {
				return fmt.Errorf("unexpected type %T for field workflow_run_contract_version", values[i])
			} else if value.Valid {
				wr.workflow_run_contract_version = new(uuid.UUID)
				*wr.workflow_run_contract_version = *value.S.(*uuid.UUID)
			}
		default:
			wr.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the WorkflowRun.
// This includes values selected through modifiers, order, etc.
func (wr *WorkflowRun) Value(name string) (ent.Value, error) {
	return wr.selectValues.Get(name)
}

// QueryWorkflow queries the "workflow" edge of the WorkflowRun entity.
func (wr *WorkflowRun) QueryWorkflow() *WorkflowQuery {
	return NewWorkflowRunClient(wr.config).QueryWorkflow(wr)
}

// QueryContractVersion queries the "contract_version" edge of the WorkflowRun entity.
func (wr *WorkflowRun) QueryContractVersion() *WorkflowContractVersionQuery {
	return NewWorkflowRunClient(wr.config).QueryContractVersion(wr)
}

// QueryCasBackends queries the "cas_backends" edge of the WorkflowRun entity.
func (wr *WorkflowRun) QueryCasBackends() *CASBackendQuery {
	return NewWorkflowRunClient(wr.config).QueryCasBackends(wr)
}

// QueryVersion queries the "version" edge of the WorkflowRun entity.
func (wr *WorkflowRun) QueryVersion() *ProjectVersionQuery {
	return NewWorkflowRunClient(wr.config).QueryVersion(wr)
}

// Update returns a builder for updating this WorkflowRun.
// Note that you need to call WorkflowRun.Unwrap() before calling this method if this WorkflowRun
// was returned from a transaction, and the transaction was committed or rolled back.
func (wr *WorkflowRun) Update() *WorkflowRunUpdateOne {
	return NewWorkflowRunClient(wr.config).UpdateOne(wr)
}

// Unwrap unwraps the WorkflowRun entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (wr *WorkflowRun) Unwrap() *WorkflowRun {
	_tx, ok := wr.config.driver.(*txDriver)
	if !ok {
		panic("ent: WorkflowRun is not a transactional entity")
	}
	wr.config.driver = _tx.drv
	return wr
}

// String implements the fmt.Stringer.
func (wr *WorkflowRun) String() string {
	var builder strings.Builder
	builder.WriteString("WorkflowRun(")
	builder.WriteString(fmt.Sprintf("id=%v, ", wr.ID))
	builder.WriteString("created_at=")
	builder.WriteString(wr.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("finished_at=")
	builder.WriteString(wr.FinishedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("state=")
	builder.WriteString(fmt.Sprintf("%v", wr.State))
	builder.WriteString(", ")
	builder.WriteString("reason=")
	builder.WriteString(wr.Reason)
	builder.WriteString(", ")
	builder.WriteString("run_url=")
	builder.WriteString(wr.RunURL)
	builder.WriteString(", ")
	builder.WriteString("runner_type=")
	builder.WriteString(wr.RunnerType)
	builder.WriteString(", ")
	builder.WriteString("attestation=")
	builder.WriteString(fmt.Sprintf("%v", wr.Attestation))
	builder.WriteString(", ")
	builder.WriteString("attestation_digest=")
	builder.WriteString(wr.AttestationDigest)
	builder.WriteString(", ")
	builder.WriteString("attestation_state=")
	builder.WriteString(fmt.Sprintf("%v", wr.AttestationState))
	builder.WriteString(", ")
	builder.WriteString("contract_revision_used=")
	builder.WriteString(fmt.Sprintf("%v", wr.ContractRevisionUsed))
	builder.WriteString(", ")
	builder.WriteString("contract_revision_latest=")
	builder.WriteString(fmt.Sprintf("%v", wr.ContractRevisionLatest))
	builder.WriteString(", ")
	builder.WriteString("version_id=")
	builder.WriteString(fmt.Sprintf("%v", wr.VersionID))
	builder.WriteByte(')')
	return builder.String()
}

// WorkflowRuns is a parsable slice of WorkflowRun.
type WorkflowRuns []*WorkflowRun
