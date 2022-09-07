package sap_api_input_reader

import (
	"sap-api-integrations-production-routing-creates/SAP_API_Caller/requests"
)

func (sdc *SDC) ConvertToHeader() *requests.Header {
	data := sdc.ProductionRouting
	return &requests.Header{
		ProductionRoutingGroup:        data.ProductionRoutingGroup,
		ProductionRouting:             data.ProductionRouting,
		ProductionRoutingInternalVers: data.ProductionRoutingInternalVers,
		IsMarkedForDeletion:           data.IsMarkedForDeletion,
		BillOfOperationsDesc:          data.BillOfOperationsDesc,
		Plant:                         data.Plant,
		BillOfOperationsUsage:         data.BillOfOperationsUsage,
		BillOfOperationsStatus:        data.BillOfOperationsStatus,
		ResponsiblePlannerGroup:       data.ResponsiblePlannerGroup,
		MinimumLotSizeQuantity:        data.MinimumLotSizeQuantity,
		MaximumLotSizeQuantity:        data.MaximumLotSizeQuantity,
		BillOfOperationsUnit:          data.BillOfOperationsUnit,
		CreationDate:                  data.CreationDate,
		CreatedByUser:                 data.CreatedByUser,
		LastChangeDate:                data.LastChangeDate,
		ValidityStartDate:             data.ValidityStartDate,
		ValidityEndDate:               data.ValidityEndDate,
		ChangeNumber:                  data.ChangeNumber,
		PlainLongText:                 data.PlainLongText,
	}
}

func (sdc *SDC) ConvertToMaterialAssignment() *requests.MaterialAssignment {
	dataProductionRouting := sdc.ProductionRouting
	data := sdc.ProductionRouting.MaterialAssignment
	return &requests.MaterialAssignment{
		Product:                        data.Product,
		Plant:                          dataProductionRouting.Plant,
		ProductionRoutingMatlAssgmt:    data.ProductionRoutingMatlAssgmt,
		ProductionRtgMatlAssgmtIntVers: data.ProductionRtgMatlAssgmtIntVers,
		CreationDate:                   data.CreationDate,
		LastChangeDate:                 data.LastChangeDate,
		ValidityStartDate:              data.ValidityStartDate,
		ValidityEndDate:                data.ValidityEndDate,
		ChangeNumber:                   data.ChangeNumber,
	}
}

func (sdc *SDC) ConvertToSequence() *requests.Sequence {
	data := sdc.ProductionRouting.Sequence
	return &requests.Sequence{
		ProductionRoutingSequence:    data.ProductionRoutingSequence,
		ProductionRoutingSqncIntVers: data.ProductionRoutingSqncIntVers,
		ChangeNumber:                 data.ChangeNumber,
		ValidityStartDate:            data.ValidityStartDate,
		ValidityEndDate:              data.ValidityEndDate,
		SequenceCategory:             data.SequenceCategory,
		BillOfOperationsRefSequence:  data.BillOfOperationsRefSequence,
		MinimumLotSizeQuantity:       data.MinimumLotSizeQuantity,
		MaximumLotSizeQuantity:       data.MaximumLotSizeQuantity,
		BillOfOperationsUnit:         data.BillOfOperationsUnit,
		SequenceText:                 data.SequenceText,
		CreationDate:                 data.CreationDate,
		LastChangeDate:               data.LastChangeDate,
	}
}

func (sdc *SDC) ConvertToOperation() *requests.Operation {
	dataProductionRoutingSequence := sdc.ProductionRouting.Sequence
	data := sdc.ProductionRouting.Sequence.Operation
	return &requests.Operation{
		ProductionRoutingSequence:     dataProductionRoutingSequence.ProductionRoutingSequence,
		ProductionRoutingOpIntID:      data.ProductionRoutingOpIntID,
		ProductionRoutingOpIntVersion: data.ProductionRoutingOpIntVersion,
		Operation:                     data.Operation,
		CreationDate:                  data.CreationDate,
		LastChangeDate:                data.LastChangeDate,
		ChangeNumber:                  data.ChangeNumber,
		ValidityStartDate:             data.ValidityStartDate,
		ValidityEndDate:               data.ValidityEndDate,
		OperationText:                 data.OperationText,
		LongTextLanguageCode:          data.LongTextLanguageCode,
		Plant:                         data.Plant,
		OperationControlProfile:       data.OperationControlProfile,
		OperationStandardTextCode:     data.OperationStandardTextCode,
		WorkCenterTypeCode:            data.WorkCenterTypeCode,
		WorkCenterInternalID:          data.WorkCenterInternalID,
		CapacityCategoryCode:          data.CapacityCategoryCode,
		OperationCostingRelevancyType: data.OperationCostingRelevancyType,
		NumberOfTimeTickets:           data.NumberOfTimeTickets,
		NumberOfConfirmationSlips:     data.NumberOfConfirmationSlips,
		OperationSetupType:            data.OperationSetupType,
		OperationSetupGroupCategory:   data.OperationSetupGroupCategory,
		OperationSetupGroup:           data.OperationSetupGroup,
		OperationReferenceQuantity:    data.OperationReferenceQuantity,
		OperationUnit:                 data.OperationUnit,
		OpQtyToBaseQtyNmrtr:           data.OpQtyToBaseQtyNmrtr,
		OpQtyToBaseQtyDnmntr:          data.OpQtyToBaseQtyDnmntr,
		MaximumWaitDuration:           data.MaximumWaitDuration,
		MaximumWaitDurationUnit:       data.MaximumWaitDurationUnit,
		MinimumWaitDuration:           data.MinimumWaitDuration,
		MinimumWaitDurationUnit:       data.MinimumWaitDurationUnit,
		StandardQueueDuration:         data.StandardQueueDuration,
		StandardQueueDurationUnit:     data.StandardQueueDurationUnit,
		MinimumQueueDuration:          data.MinimumQueueDuration,
		MinimumQueueDurationUnit:      data.MinimumQueueDurationUnit,
		StandardMoveDuration:          data.StandardMoveDuration,
		StandardMoveDurationUnit:      data.StandardMoveDurationUnit,
		MinimumMoveDuration:           data.MinimumMoveDuration,
		MinimumMoveDurationUnit:       data.MinimumMoveDurationUnit,
		OpIsExtlyProcdWithSubcontrg:   data.OpIsExtlyProcdWithSubcontrg,
		PurchasingInfoRecord:          data.PurchasingInfoRecord,
		PurchasingOrganization:        data.PurchasingOrganization,
		PlannedDeliveryDuration:       data.PlannedDeliveryDuration,
		MaterialGroup:                 data.MaterialGroup,
		PurchasingGroup:               data.PurchasingGroup,
		Supplier:                      data.Supplier,
		NumberOfOperationPriceUnits:   data.NumberOfOperationPriceUnits,
		CostElement:                   data.CostElement,
		OpExternalProcessingPrice:     data.OpExternalProcessingPrice,
		OpExternalProcessingCurrency:  data.OpExternalProcessingCurrency,
		OperationScrapPercent:         data.OperationScrapPercent,
		ChangedDateTime:               data.ChangedDateTime,
		PlainLongText:                 data.PlainLongText,
	}
}

func (sdc *SDC) ConvertToComponentAllocation() *requests.ComponentAllocation {
	dataProductionRoutingSequence := sdc.ProductionRouting.Sequence
	dataProductionRoutingSequenceOperation := sdc.ProductionRouting.Sequence.Operation
	data := sdc.ProductionRouting.Sequence.Operation.ComponentAllocation
	return &requests.ComponentAllocation{
		ProductionRoutingSequence:    dataProductionRoutingSequence.ProductionRoutingSequence,
		ProductionRoutingOpIntID:     dataProductionRoutingSequenceOperation.ProductionRoutingOpIntID,
		ProdnRtgOpBOMItemInternalID:  data.ProdnRtgOpBOMItemInternalID,
		ProdnRtgOpBOMItemIntVersion:  data.ProdnRtgOpBOMItemIntVersion,
		BillOfMaterialCategory:       data.BillOfMaterialCategory,
		BillOfMaterial:               data.BillOfMaterial,
		BillOfMaterialVariant:        data.BillOfMaterialVariant,
		BillOfMaterialItemNodeNumber: data.BillOfMaterialItemNodeNumber,
		MatlCompIsMarkedForBackflush: data.MatlCompIsMarkedForBackflush,
		CreationDate:                 data.CreationDate,
		LastChangeDate:               data.LastChangeDate,
		ValidityStartDate:            data.ValidityStartDate,
		ValidityEndDate:              data.ValidityEndDate,
		ChangeNumber:                 data.ChangeNumber,
		ChangedDateTime:              data.ChangedDateTime,
	}
}
