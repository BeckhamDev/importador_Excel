package repository

import (
	"database/sql"
	"importador_Excel/internal/models"
	"strings"

	"github.com/google/uuid"
)

type Reconfile struct {
	db *sql.DB
}

func NewReconRepository(db *sql.DB) *Reconfile {
	return &Reconfile{db}
}

func (c *Reconfile) SaveData(data [][]string) error {
	for _, row := range data {	

		//Passando partnerId de string para uuid
		id := strings.ReplaceAll(row[0], " ", "")
		partnerId, err := uuid.Parse(id)
		if err != nil {
			return err
		}

		_, err = c.db.Exec("INSERT INTO reconfile (PartnerId, PartnerName, CustomerId, CustomerName, CustomerDomainName, CustomerCountry, MpnId, Tier2MpnId, InvoiceNumber, ProductId, SkuId, AvaliabilityId, SkuName, ProductName, PublisherName, PublisherId, SubscriptionDescription, SubscriptionId, ChargeStartDate, ChargeEndDate, UsageDate, MeterType, MeterCategory, MeterId, MeterSubCategory, MeterName, MeterRegion, Unit, ResourceLocation, ConsumedService, ResourceGroup, ResourceURI, ChargeType, UnitPrice, Quantity, UnitType, BillingPreTaxTotal, BillingCurrency, PricingPreTaxTotal, PricingCurrency, ServiceInfo1, ServiceInfo2, Tags, AdditionalInfo, EffectiveUnitPrice, PCToBCExchangeRate, PCToBCExchangeRateDate, EntitlementId, EntitlementDescription, PartnerEarnedCreditPercentage, CreditPercentage, CreditType, BenefitOrderId, BenefitId, BenefitType) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21, $22, $23, $24, $25, $26, $27, $28, $29, $30, $31, $32, $33, $34, $35, $36, $37, $38, $39, $40, $41, $42, $43, $44, $45, $46, $47, $48, $49, $50, $51, $52, $53, $54, $55)", partnerId, row[1], row[2] , row[3], row[4], row[5], row[6], row[7], row[8], row[9], row[10], row[11], row[12], row[13], row[14], row[15], row[16], row[17], row[18], row[19], row[20], row[21], row[22], row[23], row[24], row[25], row[26], row[27], row[28], row[29], row[30], row[31], row[32], row[33], row[34], row[35], row[36], row[37], row[38], row[39], row[40], row[41], row[42], row[43], row[44], row[45], row[46], row[47], row[48], row[49], row[50], row[51], row[52], row[53], row[54])
		if err != nil {
			return err
		}
	}

	return nil
}

func (c *Reconfile) GetByCustomerName(customerName string) ([]models.FileRow, error){
	sql := "SELECT * FROM reconfile WHERE customername = $1"
	stmt, err := c.db.Query(sql, customerName)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	var rows []models.FileRow

	for stmt.Next() {
		var row models.FileRow
		if err = stmt.Scan(
			&row.PartnerId,                  
			&row.PartnerName,      
			&row.CustomerId,   
			&row.CustomerName,
			&row.CustomerDomainName,
			&row.CustomerCountry,
			&row.MpnId,
			&row.Tier2MpnId,
			&row.InvoiceNumber,
			&row.ProductId,
			&row.SkuId,
			&row.AvaliabilityId,
			&row.SkuName,
			&row.ProductName,
			&row.PublisherName,
			&row.PublisherId,
			&row.SubscriptionDescription,
			&row.SubscriptionId,
			&row.ChargeStartDate,
			&row.ChargeEndDate,
			&row.UsageDate,
			&row.MeterType,
			&row.MeterCategory,
			&row.MeterId,
			&row.MeterSubCategory,
			&row.MeterName,
			&row.MeterRegion,
			&row.Unit,
			&row.ResourceLocation,
			&row.ConsumedService,
			&row.ResourceGroup,
			&row.ResourceURI,
			&row.ChargeType,
			&row.UnitPrice,
			&row.Quantity,
			&row.UnitType,
			&row.BillingPreTaxTotal,
			&row.BillingCurrency,
			&row.PricingPreTaxTotal,
			&row.PricingCurrency,
			&row.ServiceInfo1,
			&row.ServiceInfo2,
			&row.Tags,
			&row.AdditionalInfo,
			&row.EffectiveUnitPrice,
			&row.PCToBCExchangeRate,
			&row.PCToBCExchangeRateDate,
			&row.EntitlementId,
			&row.EntitlementDescription,
			&row.PartnerEarnedCreditPercentage,
			&row.CreditPercentage,
			&row.CreditType,
			&row.BenefitOrderId,
			&row.BenefitId,
			&row.BenefitType,
			
		); err != nil {
			return nil, err
		}

		rows = append(rows, row)
	}

	return rows, nil
}

func (c *Reconfile) GetByMeterCategory(meterCategory string) ([]models.FileRow, error){
	sql := "SELECT * FROM reconfile WHERE metercategory = $1"
	stmt, err := c.db.Query(sql, meterCategory)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	var rows []models.FileRow

	for stmt.Next() {
		var row models.FileRow
		if err = stmt.Scan(
			&row.PartnerId,                  
			&row.PartnerName,      
			&row.CustomerId,   
			&row.CustomerName,
			&row.CustomerDomainName,
			&row.CustomerCountry,
			&row.MpnId,
			&row.Tier2MpnId,
			&row.InvoiceNumber,
			&row.ProductId,
			&row.SkuId,
			&row.AvaliabilityId,
			&row.SkuName,
			&row.ProductName,
			&row.PublisherName,
			&row.PublisherId,
			&row.SubscriptionDescription,
			&row.SubscriptionId,
			&row.ChargeStartDate,
			&row.ChargeEndDate,
			&row.UsageDate,
			&row.MeterType,
			&row.MeterCategory,
			&row.MeterId,
			&row.MeterSubCategory,
			&row.MeterName,
			&row.MeterRegion,
			&row.Unit,
			&row.ResourceLocation,
			&row.ConsumedService,
			&row.ResourceGroup,
			&row.ResourceURI,
			&row.ChargeType,
			&row.UnitPrice,
			&row.Quantity,
			&row.UnitType,
			&row.BillingPreTaxTotal,
			&row.BillingCurrency,
			&row.PricingPreTaxTotal,
			&row.PricingCurrency,
			&row.ServiceInfo1,
			&row.ServiceInfo2,
			&row.Tags,
			&row.AdditionalInfo,
			&row.EffectiveUnitPrice,
			&row.PCToBCExchangeRate,
			&row.PCToBCExchangeRateDate,
			&row.EntitlementId,
			&row.EntitlementDescription,
			&row.PartnerEarnedCreditPercentage,
			&row.CreditPercentage,
			&row.CreditType,
			&row.BenefitOrderId,
			&row.BenefitId,
			&row.BenefitType,
			
		); err != nil {
			return nil, err
		}

		rows = append(rows, row)
	}

	return rows, nil
}

func (c *Reconfile) GetByMeterRegion(meterRegion string) ([]models.FileRow, error){
	sql := "SELECT * FROM reconfile WHERE meterregion = $1"
	stmt, err := c.db.Query(sql, meterRegion)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	var rows []models.FileRow

	for stmt.Next() {
		var row models.FileRow
		if err = stmt.Scan(
			&row.PartnerId,                  
			&row.PartnerName,      
			&row.CustomerId,   
			&row.CustomerName,
			&row.CustomerDomainName,
			&row.CustomerCountry,
			&row.MpnId,
			&row.Tier2MpnId,
			&row.InvoiceNumber,
			&row.ProductId,
			&row.SkuId,
			&row.AvaliabilityId,
			&row.SkuName,
			&row.ProductName,
			&row.PublisherName,
			&row.PublisherId,
			&row.SubscriptionDescription,
			&row.SubscriptionId,
			&row.ChargeStartDate,
			&row.ChargeEndDate,
			&row.UsageDate,
			&row.MeterType,
			&row.MeterCategory,
			&row.MeterId,
			&row.MeterSubCategory,
			&row.MeterName,
			&row.MeterRegion,
			&row.Unit,
			&row.ResourceLocation,
			&row.ConsumedService,
			&row.ResourceGroup,
			&row.ResourceURI,
			&row.ChargeType,
			&row.UnitPrice,
			&row.Quantity,
			&row.UnitType,
			&row.BillingPreTaxTotal,
			&row.BillingCurrency,
			&row.PricingPreTaxTotal,
			&row.PricingCurrency,
			&row.ServiceInfo1,
			&row.ServiceInfo2,
			&row.Tags,
			&row.AdditionalInfo,
			&row.EffectiveUnitPrice,
			&row.PCToBCExchangeRate,
			&row.PCToBCExchangeRateDate,
			&row.EntitlementId,
			&row.EntitlementDescription,
			&row.PartnerEarnedCreditPercentage,
			&row.CreditPercentage,
			&row.CreditType,
			&row.BenefitOrderId,
			&row.BenefitId,
			&row.BenefitType,
			
		); err != nil {
			return nil, err
		}

		rows = append(rows, row)
	}

	return rows, nil
}

func (c *Reconfile) GetByResourceGroup(resourceGroup string) ([]models.FileRow, error){
	sql := "SELECT * FROM reconfile WHERE resourcegroup = $1"
	stmt, err := c.db.Query(sql, resourceGroup)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	var rows []models.FileRow

	for stmt.Next() {
		var row models.FileRow
		if err = stmt.Scan(
			&row.PartnerId,                  
			&row.PartnerName,      
			&row.CustomerId,   
			&row.CustomerName,
			&row.CustomerDomainName,
			&row.CustomerCountry,
			&row.MpnId,
			&row.Tier2MpnId,
			&row.InvoiceNumber,
			&row.ProductId,
			&row.SkuId,
			&row.AvaliabilityId,
			&row.SkuName,
			&row.ProductName,
			&row.PublisherName,
			&row.PublisherId,
			&row.SubscriptionDescription,
			&row.SubscriptionId,
			&row.ChargeStartDate,
			&row.ChargeEndDate,
			&row.UsageDate,
			&row.MeterType,
			&row.MeterCategory,
			&row.MeterId,
			&row.MeterSubCategory,
			&row.MeterName,
			&row.MeterRegion,
			&row.Unit,
			&row.ResourceLocation,
			&row.ConsumedService,
			&row.ResourceGroup,
			&row.ResourceURI,
			&row.ChargeType,
			&row.UnitPrice,
			&row.Quantity,
			&row.UnitType,
			&row.BillingPreTaxTotal,
			&row.BillingCurrency,
			&row.PricingPreTaxTotal,
			&row.PricingCurrency,
			&row.ServiceInfo1,
			&row.ServiceInfo2,
			&row.Tags,
			&row.AdditionalInfo,
			&row.EffectiveUnitPrice,
			&row.PCToBCExchangeRate,
			&row.PCToBCExchangeRateDate,
			&row.EntitlementId,
			&row.EntitlementDescription,
			&row.PartnerEarnedCreditPercentage,
			&row.CreditPercentage,
			&row.CreditType,
			&row.BenefitOrderId,
			&row.BenefitId,
			&row.BenefitType,
			
		); err != nil {
			return nil, err
		}

		rows = append(rows, row)
	}

	return rows, nil
}