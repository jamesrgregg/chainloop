// Code generated by ent, DO NOT EDIT.

package ent

import (
	"time"

	"github.com/chainloop-dev/chainloop/app/controlplane/internal/data/ent/casbackend"
	"github.com/chainloop-dev/chainloop/app/controlplane/internal/data/ent/integration"
	"github.com/chainloop-dev/chainloop/app/controlplane/internal/data/ent/integrationattachment"
	"github.com/chainloop-dev/chainloop/app/controlplane/internal/data/ent/membership"
	"github.com/chainloop-dev/chainloop/app/controlplane/internal/data/ent/organization"
	"github.com/chainloop-dev/chainloop/app/controlplane/internal/data/ent/robotaccount"
	"github.com/chainloop-dev/chainloop/app/controlplane/internal/data/ent/schema"
	"github.com/chainloop-dev/chainloop/app/controlplane/internal/data/ent/user"
	"github.com/chainloop-dev/chainloop/app/controlplane/internal/data/ent/workflow"
	"github.com/chainloop-dev/chainloop/app/controlplane/internal/data/ent/workflowcontract"
	"github.com/chainloop-dev/chainloop/app/controlplane/internal/data/ent/workflowcontractversion"
	"github.com/chainloop-dev/chainloop/app/controlplane/internal/data/ent/workflowrun"
	"github.com/google/uuid"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	casbackendFields := schema.CASBackend{}.Fields()
	_ = casbackendFields
	// casbackendDescCreatedAt is the schema descriptor for created_at field.
	casbackendDescCreatedAt := casbackendFields[5].Descriptor()
	// casbackend.DefaultCreatedAt holds the default value on creation for the created_at field.
	casbackend.DefaultCreatedAt = casbackendDescCreatedAt.Default.(func() time.Time)
	// casbackendDescValidatedAt is the schema descriptor for validated_at field.
	casbackendDescValidatedAt := casbackendFields[7].Descriptor()
	// casbackend.DefaultValidatedAt holds the default value on creation for the validated_at field.
	casbackend.DefaultValidatedAt = casbackendDescValidatedAt.Default.(func() time.Time)
	// casbackendDescDefault is the schema descriptor for default field.
	casbackendDescDefault := casbackendFields[8].Descriptor()
	// casbackend.DefaultDefault holds the default value on creation for the default field.
	casbackend.DefaultDefault = casbackendDescDefault.Default.(bool)
	// casbackendDescFallback is the schema descriptor for fallback field.
	casbackendDescFallback := casbackendFields[10].Descriptor()
	// casbackend.DefaultFallback holds the default value on creation for the fallback field.
	casbackend.DefaultFallback = casbackendDescFallback.Default.(bool)
	// casbackendDescID is the schema descriptor for id field.
	casbackendDescID := casbackendFields[0].Descriptor()
	// casbackend.DefaultID holds the default value on creation for the id field.
	casbackend.DefaultID = casbackendDescID.Default.(func() uuid.UUID)
	integrationFields := schema.Integration{}.Fields()
	_ = integrationFields
	// integrationDescCreatedAt is the schema descriptor for created_at field.
	integrationDescCreatedAt := integrationFields[4].Descriptor()
	// integration.DefaultCreatedAt holds the default value on creation for the created_at field.
	integration.DefaultCreatedAt = integrationDescCreatedAt.Default.(func() time.Time)
	// integrationDescID is the schema descriptor for id field.
	integrationDescID := integrationFields[0].Descriptor()
	// integration.DefaultID holds the default value on creation for the id field.
	integration.DefaultID = integrationDescID.Default.(func() uuid.UUID)
	integrationattachmentFields := schema.IntegrationAttachment{}.Fields()
	_ = integrationattachmentFields
	// integrationattachmentDescCreatedAt is the schema descriptor for created_at field.
	integrationattachmentDescCreatedAt := integrationattachmentFields[1].Descriptor()
	// integrationattachment.DefaultCreatedAt holds the default value on creation for the created_at field.
	integrationattachment.DefaultCreatedAt = integrationattachmentDescCreatedAt.Default.(func() time.Time)
	// integrationattachmentDescID is the schema descriptor for id field.
	integrationattachmentDescID := integrationattachmentFields[0].Descriptor()
	// integrationattachment.DefaultID holds the default value on creation for the id field.
	integrationattachment.DefaultID = integrationattachmentDescID.Default.(func() uuid.UUID)
	membershipFields := schema.Membership{}.Fields()
	_ = membershipFields
	// membershipDescCurrent is the schema descriptor for current field.
	membershipDescCurrent := membershipFields[1].Descriptor()
	// membership.DefaultCurrent holds the default value on creation for the current field.
	membership.DefaultCurrent = membershipDescCurrent.Default.(bool)
	// membershipDescCreatedAt is the schema descriptor for created_at field.
	membershipDescCreatedAt := membershipFields[2].Descriptor()
	// membership.DefaultCreatedAt holds the default value on creation for the created_at field.
	membership.DefaultCreatedAt = membershipDescCreatedAt.Default.(func() time.Time)
	// membershipDescUpdatedAt is the schema descriptor for updated_at field.
	membershipDescUpdatedAt := membershipFields[3].Descriptor()
	// membership.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	membership.DefaultUpdatedAt = membershipDescUpdatedAt.Default.(func() time.Time)
	// membershipDescID is the schema descriptor for id field.
	membershipDescID := membershipFields[0].Descriptor()
	// membership.DefaultID holds the default value on creation for the id field.
	membership.DefaultID = membershipDescID.Default.(func() uuid.UUID)
	organizationFields := schema.Organization{}.Fields()
	_ = organizationFields
	// organizationDescName is the schema descriptor for name field.
	organizationDescName := organizationFields[1].Descriptor()
	// organization.DefaultName holds the default value on creation for the name field.
	organization.DefaultName = organizationDescName.Default.(string)
	// organizationDescCreatedAt is the schema descriptor for created_at field.
	organizationDescCreatedAt := organizationFields[2].Descriptor()
	// organization.DefaultCreatedAt holds the default value on creation for the created_at field.
	organization.DefaultCreatedAt = organizationDescCreatedAt.Default.(func() time.Time)
	// organizationDescID is the schema descriptor for id field.
	organizationDescID := organizationFields[0].Descriptor()
	// organization.DefaultID holds the default value on creation for the id field.
	organization.DefaultID = organizationDescID.Default.(func() uuid.UUID)
	robotaccountFields := schema.RobotAccount{}.Fields()
	_ = robotaccountFields
	// robotaccountDescCreatedAt is the schema descriptor for created_at field.
	robotaccountDescCreatedAt := robotaccountFields[2].Descriptor()
	// robotaccount.DefaultCreatedAt holds the default value on creation for the created_at field.
	robotaccount.DefaultCreatedAt = robotaccountDescCreatedAt.Default.(func() time.Time)
	// robotaccountDescID is the schema descriptor for id field.
	robotaccountDescID := robotaccountFields[0].Descriptor()
	// robotaccount.DefaultID holds the default value on creation for the id field.
	robotaccount.DefaultID = robotaccountDescID.Default.(func() uuid.UUID)
	userFields := schema.User{}.Fields()
	_ = userFields
	// userDescEmail is the schema descriptor for email field.
	userDescEmail := userFields[0].Descriptor()
	// user.EmailValidator is a validator for the "email" field. It is called by the builders before save.
	user.EmailValidator = userDescEmail.Validators[0].(func(string) error)
	// userDescCreatedAt is the schema descriptor for created_at field.
	userDescCreatedAt := userFields[2].Descriptor()
	// user.DefaultCreatedAt holds the default value on creation for the created_at field.
	user.DefaultCreatedAt = userDescCreatedAt.Default.(func() time.Time)
	// userDescID is the schema descriptor for id field.
	userDescID := userFields[1].Descriptor()
	// user.DefaultID holds the default value on creation for the id field.
	user.DefaultID = userDescID.Default.(func() uuid.UUID)
	workflowFields := schema.Workflow{}.Fields()
	_ = workflowFields
	// workflowDescRunsCount is the schema descriptor for runs_count field.
	workflowDescRunsCount := workflowFields[3].Descriptor()
	// workflow.DefaultRunsCount holds the default value on creation for the runs_count field.
	workflow.DefaultRunsCount = workflowDescRunsCount.Default.(int)
	// workflowDescCreatedAt is the schema descriptor for created_at field.
	workflowDescCreatedAt := workflowFields[5].Descriptor()
	// workflow.DefaultCreatedAt holds the default value on creation for the created_at field.
	workflow.DefaultCreatedAt = workflowDescCreatedAt.Default.(func() time.Time)
	// workflowDescID is the schema descriptor for id field.
	workflowDescID := workflowFields[4].Descriptor()
	// workflow.DefaultID holds the default value on creation for the id field.
	workflow.DefaultID = workflowDescID.Default.(func() uuid.UUID)
	workflowcontractFields := schema.WorkflowContract{}.Fields()
	_ = workflowcontractFields
	// workflowcontractDescCreatedAt is the schema descriptor for created_at field.
	workflowcontractDescCreatedAt := workflowcontractFields[2].Descriptor()
	// workflowcontract.DefaultCreatedAt holds the default value on creation for the created_at field.
	workflowcontract.DefaultCreatedAt = workflowcontractDescCreatedAt.Default.(func() time.Time)
	// workflowcontractDescID is the schema descriptor for id field.
	workflowcontractDescID := workflowcontractFields[0].Descriptor()
	// workflowcontract.DefaultID holds the default value on creation for the id field.
	workflowcontract.DefaultID = workflowcontractDescID.Default.(func() uuid.UUID)
	workflowcontractversionFields := schema.WorkflowContractVersion{}.Fields()
	_ = workflowcontractversionFields
	// workflowcontractversionDescBody is the schema descriptor for body field.
	workflowcontractversionDescBody := workflowcontractversionFields[1].Descriptor()
	// workflowcontractversion.BodyValidator is a validator for the "body" field. It is called by the builders before save.
	workflowcontractversion.BodyValidator = workflowcontractversionDescBody.Validators[0].(func([]byte) error)
	// workflowcontractversionDescRevision is the schema descriptor for revision field.
	workflowcontractversionDescRevision := workflowcontractversionFields[2].Descriptor()
	// workflowcontractversion.DefaultRevision holds the default value on creation for the revision field.
	workflowcontractversion.DefaultRevision = workflowcontractversionDescRevision.Default.(int)
	// workflowcontractversionDescCreatedAt is the schema descriptor for created_at field.
	workflowcontractversionDescCreatedAt := workflowcontractversionFields[3].Descriptor()
	// workflowcontractversion.DefaultCreatedAt holds the default value on creation for the created_at field.
	workflowcontractversion.DefaultCreatedAt = workflowcontractversionDescCreatedAt.Default.(func() time.Time)
	// workflowcontractversionDescID is the schema descriptor for id field.
	workflowcontractversionDescID := workflowcontractversionFields[0].Descriptor()
	// workflowcontractversion.DefaultID holds the default value on creation for the id field.
	workflowcontractversion.DefaultID = workflowcontractversionDescID.Default.(func() uuid.UUID)
	workflowrunFields := schema.WorkflowRun{}.Fields()
	_ = workflowrunFields
	// workflowrunDescCreatedAt is the schema descriptor for created_at field.
	workflowrunDescCreatedAt := workflowrunFields[1].Descriptor()
	// workflowrun.DefaultCreatedAt holds the default value on creation for the created_at field.
	workflowrun.DefaultCreatedAt = workflowrunDescCreatedAt.Default.(func() time.Time)
	// workflowrunDescID is the schema descriptor for id field.
	workflowrunDescID := workflowrunFields[0].Descriptor()
	// workflowrun.DefaultID holds the default value on creation for the id field.
	workflowrun.DefaultID = workflowrunDescID.Default.(func() uuid.UUID)
}
