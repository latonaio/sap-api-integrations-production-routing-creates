package responses

type Header struct {
	D struct {
		ProductionRoutingGroup        string `json:"ProductionRoutingGroup"`
		ProductionRouting             string `json:"ProductionRouting"`
		ProductionRoutingInternalVers string `json:"ProductionRoutingInternalVers"`
		IsMarkedForDeletion           bool   `json:"IsMarkedForDeletion"`
		BillOfOperationsDesc          string `json:"BillOfOperationsDesc"`
		Plant                         string `json:"Plant"`
		BillOfOperationsUsage         string `json:"BillOfOperationsUsage"`
		BillOfOperationsStatus        string `json:"BillOfOperationsStatus"`
		ResponsiblePlannerGroup       string `json:"ResponsiblePlannerGroup"`
		MinimumLotSizeQuantity        string `json:"MinimumLotSizeQuantity"`
		MaximumLotSizeQuantity        string `json:"MaximumLotSizeQuantity"`
		BillOfOperationsUnit          string `json:"BillOfOperationsUnit"`
		CreationDate                  string `json:"CreationDate"`
		CreatedByUser                 string `json:"CreatedByUser"`
		LastChangeDate                string `json:"LastChangeDate"`
		ValidityStartDate             string `json:"ValidityStartDate"`
		ValidityEndDate               string `json:"ValidityEndDate"`
		ChangeNumber                  string `json:"ChangeNumber"`
		PlainLongText                 string `json:"PlainLongText"`
	} `json:"d"`
}
