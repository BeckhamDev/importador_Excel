package repository

import (
	"database/sql"
	"fmt"
	"log"
	"strings"

	"github.com/google/uuid"
)

type Reconfile struct {
	db *sql.DB
}

func NewReconRepository(db *sql.DB) *Reconfile {
	return &Reconfile{db}
}

func (c *Reconfile) SaveData(data [][]string) {
	for _, row := range data {

		id1 := strings.ReplaceAll(row[0], " ", "")
		partnerId, err := uuid.Parse(id1)
		if err != nil {
			log.Fatalf("Erro ao fazer o parse do UUID: %v", err)
		}


		_, err = c.db.Exec("INSERT INTO reconfile (PartnerId, PartnerName, CustomerId, CustomerName, CustomerDomainName, CustomerCountry, MpnId, Tier2MpnId, InvoiceNumber, ProductId, SkuId, AvaliabilityId, SkuName, ProductName, PublisherName, PublisherId, SubscriptionDescription, SubscriptionId, ChargeStartDate, ChargeEndDate, UsageDate, MeterType, MeterCategory, MeterId, MeterSubCategory, MeterName, MeterRegion, Unit, ResourceLocation, ConsumedService, ResourceGroup, ResourceURI, ChargeType, UnitPrice, Quantity, UnitType, BillingPreTaxTotal, BillingCurrency, PricingPreTaxTotal, PricingCurrency, ServiceInfo1, ServiceInfo2, Tags, AdditionalInfo, EffectiveUnitPrice, PCToBCExchangeRate, PCToBCExchangeRateDate, EntitlementId, EntitlementDescription, PartnerEarnedCreditPercentage, CreditPercentage, CreditType, BenefitOrderId, BenefitId, BenefitType) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21, $22, $23, $24, $25, $26, $27, $28, $29, $30, $31, $32, $33, $34, $35, $36, $37, $38, $39, $40, $41, $42, $43, $44, $45, $46, $47, $48, $49, $50, $51, $52, $53, $54, $55)", partnerId, row[1], row[2] , row[3], row[4], row[5], row[6], row[7], row[8], row[9], row[10], row[11], row[12], row[13], row[14], 123, row[16], row[17], row[18], row[19], row[20], row[21], row[22], row[23], row[24], row[25], row[26], row[27], row[28], row[29], row[30], row[31], row[32], row[33], row[34], row[35], row[36], row[37], row[38], row[39], row[40], row[41], row[42], row[43], row[44], row[45], row[46], row[47], row[48], row[49], row[50], row[51], row[52], row[53], row[54])
		if err != nil {
			fmt.Println(err)
		}
	}
}