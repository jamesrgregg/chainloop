// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/chainloop-dev/chainloop/app/controlplane/extensions"
	"github.com/chainloop-dev/chainloop/app/controlplane/internal/biz"
	"github.com/chainloop-dev/chainloop/app/controlplane/internal/conf"
	"github.com/chainloop-dev/chainloop/app/controlplane/internal/data"
	"github.com/chainloop-dev/chainloop/app/controlplane/internal/dispatcher"
	"github.com/chainloop-dev/chainloop/app/controlplane/internal/server"
	"github.com/chainloop-dev/chainloop/app/controlplane/internal/service"
	"github.com/chainloop-dev/chainloop/internal/blobmanager/oci"
	"github.com/chainloop-dev/chainloop/internal/credentials"
	"github.com/go-kratos/kratos/v2/log"
)

// Injectors from wire.go:

func wireApp(bootstrap *conf.Bootstrap, readerWriter credentials.ReaderWriter, logger log.Logger) (*app, func(), error) {
	confData := bootstrap.Data
	dataData, cleanup, err := data.NewData(confData, logger)
	if err != nil {
		return nil, nil, err
	}
	userRepo := data.NewUserRepo(dataData, logger)
	membershipRepo := data.NewMembershipRepo(dataData, logger)
	membershipUseCase := biz.NewMembershipUseCase(membershipRepo, logger)
	organizationRepo := data.NewOrganizationRepo(dataData, logger)
	ociRepositoryRepo := data.NewOCIRepositoryRepo(dataData, logger)
	backendProvider := oci.NewBackendProvider(readerWriter)
	ociRepositoryUseCase := biz.NewOCIRepositoryUseCase(ociRepositoryRepo, readerWriter, backendProvider, logger)
	integrationRepo := data.NewIntegrationRepo(dataData, logger)
	integrationAttachmentRepo := data.NewIntegrationAttachmentRepo(dataData, logger)
	workflowRepo := data.NewWorkflowRepo(dataData, logger)
	newIntegrationUseCaseOpts := &biz.NewIntegrationUseCaseOpts{
		IRepo:   integrationRepo,
		IaRepo:  integrationAttachmentRepo,
		WfRepo:  workflowRepo,
		CredsRW: readerWriter,
		Logger:  logger,
	}
	integrationUseCase := biz.NewIntegrationUseCase(newIntegrationUseCaseOpts)
	organizationUseCase := biz.NewOrganizationUseCase(organizationRepo, ociRepositoryUseCase, integrationUseCase, logger)
	newUserUseCaseParams := &biz.NewUserUseCaseParams{
		UserRepo:            userRepo,
		MembershipUseCase:   membershipUseCase,
		OrganizationUseCase: organizationUseCase,
		Logger:              logger,
	}
	userUseCase := biz.NewUserUseCase(newUserUseCaseParams)
	robotAccountRepo := data.NewRobotAccountRepo(dataData, logger)
	auth := bootstrap.Auth
	robotAccountUseCase := biz.NewRootAccountUseCase(robotAccountRepo, workflowRepo, auth, logger)
	casCredentialsUseCase, err := biz.NewCASCredentialsUseCase(auth)
	if err != nil {
		cleanup()
		return nil, nil, err
	}
	bootstrap_CASServer := bootstrap.CasServer
	v := _wireValue
	casClientUseCase := biz.NewCASClientUseCase(casCredentialsUseCase, bootstrap_CASServer, logger, v...)
	workflowContractRepo := data.NewWorkflowContractRepo(dataData, logger)
	workflowContractUseCase := biz.NewWorkflowContractUseCase(workflowContractRepo, logger)
	workflowUseCase := biz.NewWorkflowUsecase(workflowRepo, workflowContractUseCase, logger)
	v2 := serviceOpts(logger)
	workflowService := service.NewWorkflowService(workflowUseCase, v2...)
	confServer := bootstrap.Server
	authService, err := service.NewAuthService(userUseCase, organizationUseCase, membershipUseCase, auth, confServer, v2...)
	if err != nil {
		cleanup()
		return nil, nil, err
	}
	robotAccountService := service.NewRobotAccountService(robotAccountUseCase, v2...)
	workflowRunRepo := data.NewWorkflowRunRepo(dataData, logger)
	workflowRunUseCase, err := biz.NewWorkflowRunUseCase(workflowRunRepo, workflowRepo, logger)
	if err != nil {
		cleanup()
		return nil, nil, err
	}
	attestationUseCase := biz.NewAttestationUseCase(casClientUseCase, logger)
	newWorkflowRunServiceOpts := &service.NewWorkflowRunServiceOpts{
		WorkflowRunUC:      workflowRunUseCase,
		WorkflowUC:         workflowUseCase,
		AttestationUC:      attestationUseCase,
		WorkflowContractUC: workflowContractUseCase,
		CredsReader:        readerWriter,
		Opts:               v2,
	}
	workflowRunService := service.NewWorkflowRunService(newWorkflowRunServiceOpts)
	availableExtensions, err := extensions.Load(logger)
	if err != nil {
		cleanup()
		return nil, nil, err
	}
	fanOutDispatcher := dispatcher.New(integrationUseCase, readerWriter, casClientUseCase, availableExtensions, logger)
	newAttestationServiceOpts := &service.NewAttestationServiceOpts{
		WorkflowRunUC:      workflowRunUseCase,
		WorkflowUC:         workflowUseCase,
		WorkflowContractUC: workflowContractUseCase,
		OCIUC:              ociRepositoryUseCase,
		AttestationUC:      attestationUseCase,
		CredsReader:        readerWriter,
		IntegrationUseCase: integrationUseCase,
		CasCredsUseCase:    casCredentialsUseCase,
		FanoutDispatcher:   fanOutDispatcher,
		Opts:               v2,
	}
	attestationService := service.NewAttestationService(newAttestationServiceOpts)
	workflowContractService := service.NewWorkflowSchemaService(workflowContractUseCase, v2...)
	contextService := service.NewContextService(ociRepositoryUseCase, v2...)
	casCredentialsService := service.NewCASCredentialsService(casCredentialsUseCase, ociRepositoryUseCase, v2...)
	orgMetricsRepo := data.NewOrgMetricsRepo(dataData, logger)
	orgMetricsUseCase, err := biz.NewOrgMetricsUseCase(orgMetricsRepo, logger)
	if err != nil {
		cleanup()
		return nil, nil, err
	}
	orgMetricsService := service.NewOrgMetricsService(orgMetricsUseCase, v2...)
	ociRepositoryService := service.NewOCIRepositoryService(ociRepositoryUseCase, v2...)
	integrationsService := service.NewIntegrationsService(integrationUseCase, workflowUseCase, availableExtensions, v2...)
	organizationService := service.NewOrganizationService(membershipUseCase, v2...)
	opts := &server.Opts{
		UserUseCase:          userUseCase,
		RobotAccountUseCase:  robotAccountUseCase,
		OCIRepositoryUseCase: ociRepositoryUseCase,
		CASClientUseCase:     casClientUseCase,
		IntegrationUseCase:   integrationUseCase,
		WorkflowSvc:          workflowService,
		AuthSvc:              authService,
		RobotAccountSvc:      robotAccountService,
		WorkflowRunSvc:       workflowRunService,
		AttesstationSvc:      attestationService,
		WorkflowContractSvc:  workflowContractService,
		ContextSvc:           contextService,
		CASCredsSvc:          casCredentialsService,
		OrgMetricsSvc:        orgMetricsService,
		OCIRepositorySvc:     ociRepositoryService,
		IntegrationsSvc:      integrationsService,
		OrganizationSvc:      organizationService,
		Logger:               logger,
		ServerConfig:         confServer,
		AuthConfig:           auth,
		Credentials:          readerWriter,
	}
	grpcServer := server.NewGRPCServer(opts)
	httpServer, err := server.NewHTTPServer(opts, grpcServer)
	if err != nil {
		cleanup()
		return nil, nil, err
	}
	httpMetricsServer, err := server.NewHTTPMetricsServer(opts)
	if err != nil {
		cleanup()
		return nil, nil, err
	}
	workflowRunExpirerUseCase := biz.NewWorkflowRunExpirerUseCase(workflowRunRepo, logger)
	mainApp := newApp(logger, grpcServer, httpServer, httpMetricsServer, workflowRunExpirerUseCase)
	return mainApp, func() {
		cleanup()
	}, nil
}

var (
	_wireValue = []biz.CASClientOpts{}
)

// wire.go:

func serviceOpts(l log.Logger) []service.NewOpt {
	return []service.NewOpt{service.WithLogger(l)}
}
