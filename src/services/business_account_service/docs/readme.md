## Business Account Service

#### Table Of Contents
- [Business Account Service](#business-account-service)

#### Overview
The business account service manages all interactions and features specific to our merchant/business profiles. Some interactions that this service
 handles include, authentication and authorization of business accounts, CRUD operations against business accounts, as well as business acount
  onboarding amongst many other features.

#### Dependencies
This service witholds strict dependencies on the authentication handler service as well as the api-gateway. It leverages the authentication service
 to perform distributed account operations some of which include account locking, unlocking, archiving, ...etc

#### Service Level Interactions (SLAs)
##### Account Sign Up Flow

##### Account Login Flow

##### Account Deletion Flow

##### Account Archive Flow

#### Testing

##### Unit Testing

##### E2E Testing

##### Stress Testing
