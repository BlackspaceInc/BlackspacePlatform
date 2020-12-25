package graphql_api_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/99designs/gqlgen/client"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/stretchr/testify/assert"

	"github.com/BlackspaceInc/BlackspacePlatform/src/services/business_account_service/pkg/graphql_api"
	"github.com/BlackspaceInc/BlackspacePlatform/src/services/business_account_service/pkg/graphql_api/generated"
)

func TestE2ECreateAccount(t *testing.T) {
	t.Run("TestName:E2E_CreateAccount", CreateAccount)
	// case where account exists but is active
	// case where account exists but is inactive
	t.Run("TestName:E2E_CreateExistentInactiveAccount", CreateExistentInactiveAccount)
}

func CreateAccount(t *testing.T){
	randStr := graphql_api.GenerateRandomString(40)
	fakeEmail := randStr + "@gmail.com"
	fakeCompanyName := randStr
	fakePassword := randStr + "pwd"

	token, _ := graphql_api.CreateAccountInAuthServiceAndGetAuthToken(t, fakeEmail, fakePassword)

	resolvers := graphql_api.Resolver{Db: db}
	c := client.New(handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &resolvers})),
		client.AddHeader("Authorization", fmt.Sprintf("Bearer %s", token)))

	query := `
		mutation {
			CreateBusinessAccount(input: {
			    authnId: 10,
			    businessAccount: {
				  companyName: "%s"
			      companyAddress: "340 Clifton Pl"
			      category: "small business",
			      password: "Granada123"
			      email: "%s"
			      isActive: false
			      businessGoals: ["make money", "meet potential clients"]
			      businessStage: "early stage business"
			    }
			  }){
			    id
			  }
		}
	`

	// TODO - figure out how to pass bearer token through header
	resp, err := c.RawPost(fmt.Sprintf(query, fakeCompanyName, fakeEmail))
	ExpectedNoErrorToOccur(t, err, resp)
}

func CreateExistentInactiveAccount(t *testing.T){
	account := testBusinessAccount
	RandomizeAccount(account)
	ctx := context.TODO()

	token, authnId := graphql_api.CreateAccountInAuthServiceAndGetAuthToken(t, account.Email, account.Password)
	account.AuthnId = authnId

	// save the account
	createdAccount, err := db.CreateBusinessAccount(ctx, account, authnId)
	assert.NoError(t, err)
	assert.NotNil(t, createdAccount)

	// set the account as inactive
	err = db.ArchiveBusinessAccount(ctx, createdAccount.Id)
	assert.NoError(t, err)

	// lock the account from the context of the authentication service
	err = graphql_api.LockAccountInAuthService(t, createdAccount.AuthnId, token)
	assert.NoError(t, err)

	resolvers := graphql_api.Resolver{Db: db}
	c := client.New(handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &resolvers})),
		client.AddHeader("Authorization", fmt.Sprintf("Bearer %s", token)))

	// generate query attempting to create the same business account
	query := `
		mutation {
			CreateBusinessAccount(input: {
			    authnId: %d,
			    businessAccount: {
				  companyName: "%s"
			      companyAddress: "340 Clifton Pl"
			      category: "small business",
			      password: "%s"
			      email: "%s"
			      isActive: false
			      businessGoals: ["make money", "meet potential clients"]
			      businessStage: "early stage business"
			    }
			  }){
			    id
			  }
		}
	`

	q := fmt.Sprintf(query, account.AuthnId, account.CompanyName, account.Password, account.Email)

	resp, err := c.RawPost(q)
	assert.NoError(t, err)
	assert.NotEmpty(t, resp)
}
