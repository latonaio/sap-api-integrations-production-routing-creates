package responses

type Sequence struct {
	D struct {
		ProductionRoutingGroup       string `json:"ProductionRoutingGroup"`
		ProductionRouting            string `json:"ProductionRouting"`
		ProductionRoutingSequence    string `json:"ProductionRoutingSequence"`
		ProductionRoutingSqncIntVers string `json:"ProductionRoutingSqncIntVers"`
		ChangeNumber                 string `json:"ChangeNumber"`
		ValidityStartDate            string `json:"ValidityStartDate"`
		ValidityEndDate              string `json:"ValidityEndDate"`
		SequenceCategory             string `json:"SequenceCategory"`
		BillOfOperationsRefSequence  string `json:"BillOfOperationsRefSequence"`
		MinimumLotSizeQuantity       string `json:"MinimumLotSizeQuantity"`
		MaximumLotSizeQuantity       string `json:"MaximumLotSizeQuantity"`
		BillOfOperationsUnit         string `json:"BillOfOperationsUnit"`
		SequenceText                 string `json:"SequenceText"`
		CreationDate                 string `json:"CreationDate"`
		LastChangeDate               string `json:"LastChangeDate"`
	} `json:"d"`
}
