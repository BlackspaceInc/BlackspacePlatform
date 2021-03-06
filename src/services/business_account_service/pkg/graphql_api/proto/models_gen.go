// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package proto

import (
	"github.com/BlackspaceInc/BlackspacePlatform/src/services/business_account_service/pkg/graphql_api/model"
)

type Address struct {
	Address       *string      `json:"Address"`
	ApartmentUnit *string      `json:"ApartmentUnit"`
	ZipCode       *string      `json:"ZipCode"`
	City          *string      `json:"City"`
	State         *string      `json:"State"`
	Birthdate     *DateOfBirth `json:"birthdate"`
}

type BusinessCategory struct {
	Tech                         *bool `json:"Tech"`
	CharitiesEducationMembership *bool `json:"CharitiesEducationMembership"`
	FoodAndDrink                 *bool `json:"FoodAndDrink"`
	HealthCareAndFitness         *bool `json:"HealthCareAndFitness"`
	HomeAndRepair                *bool `json:"HomeAndRepair"`
	LeisureAndEntertainment      *bool `json:"LeisureAndEntertainment"`
	ProfessionalServices         *bool `json:"ProfessionalServices"`
	Retail                       *bool `json:"Retail"`
	Transportation               *bool `json:"Transportation"`
	BeautyAndPersonalCare        *bool `json:"BeautyAndPersonalCare"`
}

type BusinessSubCategory struct {
	Marketing       *bool `json:"Marketing"`
	Travel          *bool `json:"Travel"`
	InteriorDesign  *bool `json:"Interior_Design"`
	Music           *bool `json:"Music"`
	Technology      *bool `json:"Technology"`
	Food            *bool `json:"Food"`
	Restaurants     *bool `json:"Restaurants"`
	Polictics       *bool `json:"Polictics"`
	HealthAndBeauty *bool `json:"Health_And_Beauty"`
	Design          *bool `json:"Design"`
	NonProfit       *bool `json:"Non_Profit"`
	Jewelry         *bool `json:"Jewelry"`
	Gaming          *bool `json:"Gaming"`
	Magazine        *bool `json:"Magazine"`
	Photography     *bool `json:"Photography"`
	Fitenss         *bool `json:"Fitenss"`
	Consulting      *bool `json:"Consulting"`
	Fashion         *bool `json:"Fashion"`
	Services        *bool `json:"Services"`
	Art             *bool `json:"Art"`
}

type BusinessType struct {
	Category    *BusinessCategory    `json:"category"`
	SubCategory *BusinessSubCategory `json:"subCategory"`
}

type CreateBusinessAccountRequest struct {
	BusinessAccount *model.BusinessAccount `json:"businessAccount"`
	AuthnID         *int                   `json:"authnId"`
}

type DateOfBirth struct {
	Month *string `json:"Month"`
	Day   *string `json:"Day"`
	Year  *string `json:"Year"`
}

type DeleteBusinessAccountRequest struct {
	ID *int `json:"id"`
}

type DeleteBusinessAccountResponse struct {
	Result *bool `json:"result"`
}

type DeleteBusinessAccountsRequest struct {
	ID []*int `json:"id"`
}

type GetBusinessAccountRequest struct {
	ID *int `json:"id"`
}

type GetBusinessAccountsRequest struct {
	Limit *int `json:"limit"`
}

type Media struct {
	ID        *int    `json:"id"`
	Website   *string `json:"website"`
	Instagram *string `json:"instagram"`
	Facebook  *string `json:"facebook"`
	LinkedIn  *string `json:"linkedIn"`
	Pinterest *string `json:"pinterest"`
}

type MerchantType struct {
	SoleProprietor        *bool `json:"SoleProprietor"`
	SideProject           *bool `json:"SideProject"`
	CasualUse             *bool `json:"CasualUse"`
	LLCCorporation        *bool `json:"LLCCorporation"`
	Partnership           *bool `json:"Partnership"`
	Charity               *bool `json:"Charity"`
	ReligiousOrganization *bool `json:"ReligiousOrganization"`
	OnePersonBusiness     *bool `json:"OnePersonBusiness"`
}

type PaymentMedium struct {
	MobilePhone *bool `json:"MobilePhone"`
	Tablet      *bool `json:"Tablet"`
	Computer    *bool `json:"Computer"`
}

type PaymentOptions struct {
	BrickAndMortar  *bool `json:"BrickAndMortar"`
	OnTheGo         *bool `json:"OnTheGo"`
	Online          *bool `json:"Online"`
	ThroughInvoices *bool `json:"ThroughInvoices"`
}

type PaymentProcessingMethods struct {
	PaymentOptions []*PaymentOptions `json:"paymentOptions"`
	Medium         []*PaymentMedium  `json:"medium"`
}

type PhoneNumber struct {
	Number *string    `json:"number"`
	Type   *PhoneType `json:"type"`
}

type PhoneType struct {
	Home   *bool `json:"home"`
	Work   *bool `json:"work"`
	Mobile *bool `json:"mobile"`
}

type ServicesManagedByBlackspace struct {
	ItemCatalog         *bool `json:"ItemCatalog"`
	GiftCards           *bool `json:"GiftCards"`
	Discounts           *bool `json:"Discounts"`
	MarketingCampaigns  *bool `json:"MarketingCampaigns"`
	LoyaltyProgram      *bool `json:"LoyaltyProgram"`
	FundingYourBusiness *bool `json:"FundingYourBusiness"`
	Analytics           *bool `json:"Analytics"`
}

type Topics struct {
	ID              *int  `json:"id"`
	Technology      *bool `json:"Technology"`
	Health          *bool `json:"Health"`
	Food            *bool `json:"Food"`
	Science         *bool `json:"Science"`
	Music           *bool `json:"Music"`
	Travel          *bool `json:"Travel"`
	Business        *bool `json:"Business"`
	Cooking         *bool `json:"Cooking"`
	FashionAndStyle *bool `json:"FashionAndStyle"`
	Design          *bool `json:"Design"`
	Art             *bool `json:"Art"`
}

type UpdateBusinessAccountRequest struct {
	ID              *int                   `json:"id"`
	BusinessAccount *model.BusinessAccount `json:"businessAccount"`
}
