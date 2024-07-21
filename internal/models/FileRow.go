package models

import (
	"time"

	"github.com/google/uuid"
)

type FileRow struct {
	PartnerId                     uuid.UUID
	PartnerName                   string
	CustomerId                    string
	CustomerName                  string
	CustomerDomainName            string
	CustomerCountry               string
	MpnId                         int
	Tier2MpnId                    int
	InvoiceNumber                 string
	ProductId                     string
	SkuId                         string
	AvaliabilityId                string
	SkuName                       string
	ProductName                   string
	PublisherName                 string
	PublisherId                   string
	SubscriptionDescription       string
	SubscriptionId                string
	ChargeStartDate               time.Time
	ChargeEndDate                 time.Time
	UsageDate                     time.Time
	MeterType                     string
	MeterCategory                 string
	MeterId                       string
	MeterSubCategory              string
	MeterName                     string
	MeterRegion                   string
	Unit                          string
	ResourceLocation              string
	ConsumedService               string
	ResourceGroup                 string
	ResourceURI                   string
	ChargeType                    string
	UnitPrice                     float64
	Quantity                      float64
	UnitType                      string
	BillingPreTaxTotal            float64
	BillingCurrency               string
	PricingPreTaxTotal            float64
	PricingCurrency               string
	ServiceInfo1                  string
	ServiceInfo2                  string
	Tags                          string
	AdditionalInfo                string
	EffectiveUnitPrice            float64
	PCToBCExchangeRate            int
	PCToBCExchangeRateDate        time.Time
	EntitlementId                 string
	EntitlementDescription        string
	PartnerEarnedCreditPercentage int
	CreditPercentage              int
	CreditType                    string
	BenefitOrderId                string
	BenefitId                     string
	BenefitType                   string
}