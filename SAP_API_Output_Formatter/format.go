package sap_api_output_formatter

import (
	"encoding/json"
	"sap-api-integrations-production-routing-creates/SAP_API_Caller/responses"

	"github.com/latonaio/golang-logging-library-for-sap/logger"
	"golang.org/x/xerrors"
)

func ConvertToHeader(raw []byte, l *logger.Logger) (*Header, error) {
	pm := &responses.Header{}
	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to Header. raw data is:\n%v\nunmarshal error: %w", string(raw), err)
	}
	data := pm.D
	header := &Header{
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

	return header, nil
}

func ConvertToMaterialAssignment(raw []byte, l *logger.Logger) (*MaterialAssignment, error) {
	pm := &responses.MaterialAssignment{}
	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to MaterialAssignment. raw data is:\n%v\nunmarshal error: %w", string(raw), err)
	}
	data := pm.D
	materialAssignment := &MaterialAssignment{
		Product:                        data.Product,
		Plant:                          data.Plant,
		ProductionRoutingGroup:         data.ProductionRoutingGroup,
		ProductionRouting:              data.ProductionRouting,
		ProductionRoutingMatlAssgmt:    data.ProductionRoutingMatlAssgmt,
		ProductionRtgMatlAssgmtIntVers: data.ProductionRtgMatlAssgmtIntVers,
		CreationDate:                   data.CreationDate,
		LastChangeDate:                 data.LastChangeDate,
		ValidityStartDate:              data.ValidityStartDate,
		ValidityEndDate:                data.ValidityEndDate,
		ChangeNumber:                   data.ChangeNumber,
	}

	return materialAssignment, nil
}

func ConvertToSequence(raw []byte, l *logger.Logger) (*Sequence, error) {
	pm := &responses.Sequence{}
	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to Sequence. raw data is:\n%v\nunmarshal error: %w", string(raw), err)
	}
	data := pm.D
	sequence := &Sequence{
		ProductionRoutingGroup:       data.ProductionRoutingGroup,
		ProductionRouting:            data.ProductionRouting,
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

	return sequence, nil
}

func ConvertToOperation(raw []byte, l *logger.Logger) (*Operation, error) {
	pm := &responses.Operation{}
	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to Operation. raw data is:\n%v\nunmarshal error: %w", string(raw), err)
	}
	data := pm.D
	operation := &Operation{
		ProductionRoutingGroup:        data.ProductionRoutingGroup,
		ProductionRouting:             data.ProductionRouting,
		ProductionRoutingSequence:     data.ProductionRoutingSequence,
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

	return operation, nil
}

func ConvertToComponentAllocation(raw []byte, l *logger.Logger) (*ComponentAllocation, error) {
	pm := &responses.ComponentAllocation{}
	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to ComponentAllocation. raw data is:\n%v\nunmarshal error: %w", string(raw), err)
	}
	data := pm.D
	componentAllocation := &ComponentAllocation{
		ProductionRoutingGroup:       data.ProductionRoutingGroup,
		ProductionRouting:            data.ProductionRouting,
		ProductionRoutingSequence:    data.ProductionRoutingSequence,
		ProductionRoutingOpIntID:     data.ProductionRoutingOpIntID,
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

	return componentAllocation, nil
}
