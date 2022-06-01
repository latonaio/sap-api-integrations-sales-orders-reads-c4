package sap_api_output_formatter

type CustomerOrderCollection struct {
			ObjectID                                                 string    `json:"ObjectID"`
			InformationLifeCycleStatusCode                           string    `json:"InformationLifeCycleStatusCode"`
			InformationLifeCycleStatusCodeText                       string    `json:"InformationLifeCycleStatusCodeText"`
			ClassificationCode                                       string    `json:"ClassificationCode"`
			ClassificationCodeText                                   string    `json:"ClassificationCodeText"`
			BuyerID                                                  string    `json:"BuyerID"`
			FulfilmentBlockingReasonCode                             string    `json:"FulfilmentBlockingReasonCode"`
			FulfilmentBlockingReasonCodeText                         string    `json:"FulfilmentBlockingReasonCodeText"`
			ID                                                       string    `json:"ID"`
			InvoicingBlockingReasonCode                              string    `json:"InvoicingBlockingReasonCode"`
			InvoicingBlockingReasonCodeText                          string    `json:"InvoicingBlockingReasonCodeText"`
			LastChangeDate                                           string    `json:"LastChangeDate"`
			OrderExternalLifeCycleStatusCode                         string    `json:"OrderExternalLifeCycleStatusCode"`
			OrderExternalLifeCycleStatusCodeText                     string    `json:"OrderExternalLifeCycleStatusCodeText"`
			Name                                                     string    `json:"Name"`
			LanguageCode                                             string    `json:"LanguageCode"`
			LanguageCodeText                                         string    `json:"LanguageCodeText"`
			ProcessingTypeCode                                       string    `json:"ProcessingTypeCode"`
			ProcessingTypeCodeText                                   string    `json:"ProcessingTypeCodeText"`
			ItemListCancellationStatusCode                           string    `json:"ItemListCancellationStatusCode"`
			ItemListCancellationStatusCodeText                       string    `json:"ItemListCancellationStatusCodeText"`
			ItemListFulfilmentProcessingStatusCode                   string    `json:"ItemListFulfilmentProcessingStatusCode"`
			ItemListFulfilmentProcessingStatusCodeText               string    `json:"ItemListFulfilmentProcessingStatusCodeText"`
			ApprovalStatusCode                                       string    `json:"ApprovalStatusCode"`
			ApprovalStatusCodeText                                   string    `json:"ApprovalStatusCodeText"`
			ConsistencyStatusCode                                    string    `json:"ConsistencyStatusCode"`
			ConsistencyStatusCodeText                                string    `json:"ConsistencyStatusCodeText"`
			OverallBlockingStatusCode                                string    `json:"OverallBlockingStatusCode"`
			OverallBlockingStatusCodeText                            string    `json:"OverallBlockingStatusCodeText"`
			ReplicationProcessingStatusCode                          string    `json:"ReplicationProcessingStatusCode"`
			ReplicationProcessingStatusCodeText                      string    `json:"ReplicationProcessingStatusCodeText"`
			DistributionChannelCode                                  string    `json:"DistributionChannelCode"`
			DistributionChannelCodeText                              string    `json:"DistributionChannelCodeText"`
			DivisionCode                                             string    `json:"DivisionCode"`
			DivisionCodeText                                         string    `json:"DivisionCodeText"`
			SalesGroupID                                             string    `json:"SalesGroupID"`
			SalesOfficeID                                            string    `json:"SalesOfficeID"`
			SalesOrganisationID                                      string    `json:"SalesOrganisationID"`
			SalesTerritoryID                                         string    `json:"SalesTerritoryID"`
			BuyerPartyID                                             string    `json:"BuyerPartyID"`
			DeliveryPriorityCode                                     string    `json:"DeliveryPriorityCode"`
			DeliveryPriorityCodeText                                 string    `json:"DeliveryPriorityCodeText"`
			TransferLocationName                                     string    `json:"TransferLocationName"`
			EmployeeResponsiblePartyID                               string    `json:"EmployeeResponsiblePartyID"`
			BuyerPartyName                                           string    `json:"BuyerPartyName"`
			EmployeeResponsiblePartyName                             string    `json:"EmployeeResponsiblePartyName"`
			BuyerContactPartyID                                      string    `json:"BuyerContactPartyID"`
			BuyerContactPartyName                                    string    `json:"BuyerContactPartyName"`
			CurrencyCode                                             string    `json:"CurrencyCode"`
			CurrencyCodeText                                         string    `json:"CurrencyCodeText"`
			PriceDateTime                                            string    `json:"PriceDateTime"`
			ProductRecipientPartyID                                  string    `json:"ProductRecipientPartyID"`
			ProductRecipientPartyName                                string    `json:"ProductRecipientPartyName"`
			RequestedFulfillmentStartDateTime                        string    `json:"RequestedFulfillmentStartDateTime"`
			TimeZoneCode                                             string    `json:"timeZoneCode"`
			TimeZoneCodeText                                         string    `json:"timeZoneCodeText"`
			CancellationReasonCode                                   string    `json:"CancellationReasonCode"`
			CancellationReasonCodeText                               string    `json:"CancellationReasonCodeText"`
			SalesUnitPartyID                                         string    `json:"SalesUnitPartyID"`
			CalculationStatusCode                                    string    `json:"CalculationStatusCode"`
			CalculationStatusCodeText                                string    `json:"CalculationStatusCodeText"`
			ExternalPriceDocumentBaseBusinessTransactionDocumentUUID string    `json:"ExternalPriceDocumentBaseBusinessTransactionDocumentUUID"`
			PricingProcedureCode                                     string    `json:"PricingProcedureCode"`
			PricingProcedureCodeText                                 string    `json:"PricingProcedureCodeText"`
			DateTime                                                 string    `json:"DateTime"`
			OrderReasonCode                                          string    `json:"OrderReasonCode"`
			OrderReasonCodeText                                      string    `json:"OrderReasonCodeText"`
			MaintenanceModeInternalOnlyMainDiscount                  string    `json:"MaintenanceModeInternalOnlyMainDiscount"`
			NetAmount                                                string    `json:"NetAmount"`
			NetAmountCurrencyCode                                    string    `json:"NetAmountCurrencyCode"`
			NetAmountCurrencyCodeText                                string    `json:"NetAmountCurrencyCodeText"`
			GrossAmount                                              string    `json:"GrossAmount"`
			GrossAmountCurrencyCode                                  string    `json:"GrossAmountCurrencyCode"`
			GrossAmountCurrencyCodeText                              string    `json:"GrossAmountCurrencyCodeText"`
			TaxAmount                                                string    `json:"TaxAmount"`
			TaxAmountCurrencyCode                                    string    `json:"TaxAmountCurrencyCode"`
			TaxAmountCurrencyCodeText                                string    `json:"TaxAmountCurrencyCodeText"`
			InternalPricingProcedureCode                             string    `json:"InternalPricingProcedureCode"`
			InternalPricingProcedureCodeText                         string    `json:"InternalPricingProcedureCodeText"`
			InternalPricingCalculationStatusCode                     string    `json:"InternalPricingCalculationStatusCode"`
			InternalPricingCalculationStatusCodeText                 string    `json:"InternalPricingCalculationStatusCodeText"`
			NetWeightMeasure                                         string    `json:"NetWeightMeasure"`
			NetWeightUnitCode                                        string    `json:"NetWeightUnitCode"`
			NetWeightUnitCodeText                                    string    `json:"NetWeightUnitCodeText"`
			GrossWeightMeasure                                       string    `json:"GrossWeightMeasure"`
			GrossWeightUnitCode                                      string    `json:"GrossWeightUnitCode"`
			GrossWeightUnitCodeText                                  string    `json:"GrossWeightUnitCodeText"`
			VolumeMeasure                                            string    `json:"VolumeMeasure"`
			VolumeUnitCode                                           string    `json:"VolumeUnitCode"`
			VolumeUnitCodeText                                       string    `json:"VolumeUnitCodeText"`
			Simulate                                                 bool      `json:"Simulate"`
			SubmitForApproval                                        bool      `json:"SubmitForApproval"`
			Transfer                                                 bool      `json:"Transfer"`
			WithdrawFromApproval                                     bool      `json:"WithdrawFromApproval"`
			SetAsCompleted                                           bool      `json:"SetAsCompleted"`
			PlantPartyID                                             string    `json:"PlantPartyID"`
			PlantPartyName                                           string    `json:"PlantPartyName"`
			CreatedBy                                                string    `json:"CreatedBy"`
			LastChangedBy                                            string    `json:"LastChangedBy"`
			CreationIdentityUUID                                     string    `json:"CreationIdentityUUID"`
			LastChangeIdentityUUID                                   string    `json:"LastChangeIdentityUUID"`
			EntityLastChangedOn                                      string    `json:"EntityLastChangedOn"`
			ETag                                                     string    `json:"ETag"`
			DataloaderKUT                                            string    `json:"dataloader_KUT"`
			ToCustomerOrderCashDiscountTerms                         string    `json:"CustomerOrderCashDiscountTerms"`
			ToCustomerOrderItem                         		     string    `json:"CustomerOrderItem"`
			ToCustomerOrderParty                        		     string    `json:"CustomerOrderParty"`
}

type CustomerOrderCashDiscountTerms struct {
	        ObjectID                                                 string `json:"ObjectID"`
	        ParentObjectID                                           string `json:"ParentObjectID"`
	        Code                                                     string `json:"Code"`
	        CodeText                                                 string `json:"CodeText"`
	        SalesOrderID                                             string `json:"SalesOrderID"`
	        ETag                                                     string `json:"ETag"`
}

type CustomerOrderItem struct {
			ObjectID                                                 string `json:"ObjectID"`
			ParentObjectID                                           string `json:"ParentObjectID"`
			Description                                              string `json:"Description"`
			LanguageCode                                             string `json:"languageCode"`
			LanguageCodeText                                         string `json:"languageCodeText"`
			CreationDate                                             string `json:"CreationDate"`
			ParentItemID                                             string `json:"ParentItemID"`
			ID                                                       string `json:"ID"`
			ProcessingTypeCode                                       string `json:"ProcessingTypeCode"`
			ProcessingTypeCodeText                                   string `json:"ProcessingTypeCodeText"`
			ProductID                                                string `json:"ProductID"`
			ProductStandardID                                        string `json:"ProductStandardID"`
			ProductInternalID                                        string `json:"ProductInternalID"`
			Quantity                                                 string `json:"Quantity"`
			QuantityMeasureUnitCode                                  string `json:"QuantityMeasureUnitCode"`
			QuantityMeasureUnitCodeText                              string `json:"QuantityMeasureUnitCodeText"`
			StartDateTime                                            string `json:"StartDateTime"`
			TimeZoneCode                                             string `json:"TimeZoneCode"`
			TimeZoneCodeText                                         string `json:"TimeZoneCodeText"`
			SalesTermsCancellationReasonCode                         string `json:"SalesTermsCancellationReasonCode"`
			SalesTermsCancellationReasonCodeText                     string `json:"SalesTermsCancellationReasonCodeText"`
			MaintenanceModeInternalOnlyItemMainDiscount              string `json:"MaintenanceModeInternalOnlyItemMainDiscount"`
			ItemProductRecipientBusinessPartnerInternalID            string `json:"ItemProductRecipientBusinessPartnerInternalID"`
			BuyerID                                                  string `json:"BuyerID"`
			ProductBuyerID                                           string `json:"ProductBuyerID"`
			SalesOrderID                                             string `json:"SalesOrderID"`
			NetAmount                                                string `json:"NetAmount"`
			NetAmountCurrencyCode                                    string `json:"NetAmountCurrencyCode"`
			NetAmountCurrencyCodeText                                string `json:"NetAmountCurrencyCodeText"`
			NetPriceAmount                                           string `json:"NetPriceAmount"`
			NetPriceAmountCurrencyCode                               string `json:"NetPriceAmountCurrencyCode"`
			NetPriceAmountCurrencyCodeText                           string `json:"NetPriceAmountCurrencyCodeText"`
			TaxAmount                                                string `json:"TaxAmount"`
			TaxAmountCurrencyCode                                    string `json:"TaxAmountCurrencyCode"`
			TaxAmountCurrencyCodeText                                string `json:"TaxAmountCurrencyCodeText"`
			NetWeightMeasure                                         string `json:"NetWeightMeasure"`
			NetWeightUnitCode                                        string `json:"NetWeightUnitCode"`
			NetWeightUnitCodeText                                    string `json:"NetWeightUnitCodeText"`
			GrossWeightMeasure                                       string `json:"GrossWeightMeasure"`
			GrossWeightUnitCode                                      string `json:"GrossWeightUnitCode"`
			GrossWeightUnitCodeText                                  string `json:"GrossWeightUnitCodeText"`
			VolumeMeasure                                            string `json:"VolumeMeasure"`
			VolumeUnitCode                                           string `json:"VolumeUnitCode"`
			VolumeUnitCodeText                                       string `json:"VolumeUnitCodeText"`
			NetPriceBaseQuantity                                     string `json:"NetPriceBaseQuantity"`
			NetPriceBaseQuantityunitCode                             string `json:"NetPriceBaseQuantityunitCode"`
			NetPriceBaseQuantityunitCodeText                         string `json:"NetPriceBaseQuantityunitCodeText"`
			PlantPartyID                                             string `json:"PlantPartyID"`
			PlantName                                                string `json:"PlantName"`
			LifeCycleStatusCode                                      string `json:"LifeCycleStatusCode"`
			LifeCycleStatusCodeText                                  string `json:"LifeCycleStatusCodeText"`
			CancellationStatusCode                                   string `json:"CancellationStatusCode"`
			CancellationStatusCodeText                               string `json:"CancellationStatusCodeText"`
			FulfilmentProcessingStatusCode                           string `json:"FulfilmentProcessingStatusCode"`
			FulfilmentProcessingStatusCodeText                       string `json:"FulfilmentProcessingStatusCodeText"`
			InvoiceProcessingStatusCode                              string `json:"InvoiceProcessingStatusCode"`
			InvoiceProcessingStatusCodeText                          string `json:"InvoiceProcessingStatusCodeText"`
			EntityLastChangedOn                                      string `json:"EntityLastChangedOn"`
			ETag                                                     string `json:"ETag"`
}

type CustomerOrderParty struct {
			ObjectID                                                 string `json:"ObjectID"`
			ParentObjectID                                           string `json:"ParentObjectID"`
			TimeZoneCode                                             string `json:"TimeZoneCode"`
			TimeZoneCodeText                                         string `json:"TimeZoneCodeText"`
			PartyName                                                string `json:"PartyName"`
			PartyID                                                  string `json:"PartyID"`
			RoleCode                                                 string `json:"RoleCode"`
			RoleCodeText                                             string `json:"RoleCodeText"`
			SalesOrderID                                             string `json:"SalesOrderID"`
			PreferredCommunicationMediumTypeCode                     string `json:"PreferredCommunicationMediumTypeCode"`
			PreferredCommunicationMediumTypeCodeText                 string `json:"PreferredCommunicationMediumTypeCodeText"`
			MainIndicator                                            bool   `json:"MainIndicator"`
			ConventionalPhone                                        string `json:"ConventionalPhone"`
			Email                                                    string `json:"Email"`
			Web                                                      string `json:"Web"`
			Fax                                                      string `json:"Fax"`
			MobilePhone                                              string `json:"MobilePhone"`
			FirstLineName                                            string `json:"FirstLineName"`
			SecondLineName                                           string `json:"SecondLineName"`
			ThirdLineName                                            string `json:"ThirdLineName"`
			FourthLineName                                           string `json:"FourthLineName"`
			CareOfName                                               string `json:"CareOfName"`
			CityName                                                 string `json:"CityName"`
			CountryCode                                              string `json:"CountryCode"`
			CountryCodeText                                          string `json:"CountryCodeText"`
			CountyName                                               string `json:"CountyName"`
			DistrictName                                             string `json:"DistrictName"`
			StreetPrefixName                                         string `json:"StreetPrefixName"`
			AdditionalStreetPrefixName                               string `json:"AdditionalStreetPrefixName"`
			HouseID                                                  string `json:"HouseID"`
			StreetName                                               string `json:"StreetName"`
			StreetSuffixName                                         string `json:"StreetSuffixName"`
			AdditionalStreetSuffixName                               string `json:"AdditionalStreetSuffixName"`
			RegionCode                                               string `json:"RegionCode"`
			RegionCodeText                                           string `json:"RegionCodeText"`
			StreetPostalCode                                         string `json:"StreetPostalCode"`
			POBoxIndicator                                           bool   `json:"POBoxIndicator"`
			POBoxID                                                  string `json:"POBoxID"`
			POBoxPostalCode                                          string `json:"POBoxPostalCode"`
			POBoxDeviatingCityName                                   string `json:"POBoxDeviatingCityName"`
			POBoxDeviatingCountryCode                                string `json:"POBoxDeviatingCountryCode"`
			POBoxDeviatingCountryCodeText                            string `json:"POBoxDeviatingCountryCodeText"`
			POBoxDeviatingRegionCode                                 string `json:"POBoxDeviatingRegionCode"`
			POBoxDeviatingRegionCodeText                             string `json:"POBoxDeviatingRegionCodeText"`
			ETag                                                     string `json:"ETag"`
}