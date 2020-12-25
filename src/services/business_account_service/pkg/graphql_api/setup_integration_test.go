package graphql_api_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/BlackspaceInc/BlackspacePlatform/src/services/business_account_service/pkg/api"
	"github.com/BlackspaceInc/BlackspacePlatform/src/services/business_account_service/pkg/database"
	"github.com/BlackspaceInc/BlackspacePlatform/src/services/business_account_service/pkg/graphql_api/proto"
)

var (
	db        *database.Db
	host = "localhost"
	port = 5433
	user = "postgres"
	password = "postgres"
	dbname = "postgres"
	testBusinessAccount = &proto.BusinessAccount{
		Id:                          0,
		CompanyName:                 "BlackspaceInc",
		CompanyAddress:              "340 clifton pl",
		PhoneNumber:                 &proto.PhoneNumber{
			Number: "424-410-6123",
			Type:   0,
		},
		Category:                    "small business",
		Media:                       &proto.Media{
			Id:        0,
			Website:   "blackspaceInc.com",
			Instagram: "blkspace",
			Facebook:  "blkspace_",
			LinkedIn:  "",
			Pinterest: "",
		},
		Password:                    "Granada123",
		Email:                       "BlackspaceInc@gmail.com",
		IsActive:                    false,
		TypeOfBusiness:              &proto.BusinessType{
			Category:    proto.BusinessCategory_Tech,
			SubCategory: proto.BusinessSubCategory_Technology,
		},
		BusinessGoals:              []string{"onboard as many customers as possible"},
		BusinessStage:               "small business",
		MerchantType:                proto.MerchantType_CasualUse,
		PaymentDetails:              &proto.PaymentProcessingMethods{
			PaymentOptions: []proto.PaymentOptions{proto.PaymentOptions_Online},
			Medium:         nil,
		},
		ServicesManagedByBlackspace: proto.ServicesManagedByBlackspace_FundingYourBusiness,
		FounderAddress:              &proto.Address{
			Address:       "340 Clifton Pl",
			ApartmentUnit: "3D",
			ZipCode:       "19101",
			City:          "Brooklyn",
			State:         "NY",
			Birthdate:     &proto.DateOfBirth{
				Month: "july",
				Day:   "12",
				Year:  "1996",
			},
		},
		SubscribedTopics:            &proto.Topics{Business: true},
		AuthnId:                     0,
	}
)

func InitializeDatabaseConnection() *database.Db {
	const serviceName string = "test"
	// initiate tracing engine
	tracerEngine, closer := api.InitializeTracingEngine(serviceName)
	defer closer.Close()
	ctx := context.Background()

	// initiate metrics engine
	serviceMetrics := api.InitializeMetricsEngine(serviceName)

	// initiate logging client
	logger := api.InitializeLoggingEngine(ctx)

	connectionString := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	return database.Setup(ctx, connectionString, tracerEngine, serviceMetrics, logger, "http://localhost:9898/v1/account")
}

func TestMain(m *testing.M) {
	db = InitializeDatabaseConnection()
	defer db.Engine.Close()

	// defer deleting all created entries
	cleanupHandler := database.DeleteCreatedEntities(db.Engine)
	defer cleanupHandler()

	_ = m.Run()
	return
}
