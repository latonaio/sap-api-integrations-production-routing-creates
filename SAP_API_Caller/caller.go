package sap_api_caller

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"sap-api-integrations-production-routing-creates/SAP_API_Caller/requests"
	sap_api_output_formatter "sap-api-integrations-production-routing-creates/SAP_API_Output_Formatter"
	"strings"
	"sync"

	"github.com/latonaio/golang-logging-library-for-sap/logger"
	sap_api_request_client_header_setup "github.com/latonaio/sap-api-request-client-header-setup"
	"golang.org/x/xerrors"
)

type SAPAPICaller struct {
	baseURL         string
	sapClientNumber string
	requestClient   *sap_api_request_client_header_setup.SAPRequestClient
	log             *logger.Logger
}

func NewSAPAPICaller(baseUrl, sapClientNumber string, requestClient *sap_api_request_client_header_setup.SAPRequestClient, l *logger.Logger) *SAPAPICaller {
	return &SAPAPICaller{
		baseURL:         baseUrl,
		requestClient:   requestClient,
		sapClientNumber: sapClientNumber,
		log:             l,
	}
}

func (c *SAPAPICaller) AsyncPostProductionRouting(
	header *requests.Header,
	materialAssignment *requests.MaterialAssignment,
	sequence *requests.Sequence,
	operation *requests.Operation,
	componentAllocation *requests.ComponentAllocation,
	accepter []string) {
	wg := &sync.WaitGroup{}
	wg.Add(len(accepter))
	for _, fn := range accepter {
		switch fn {
		case "Header":
			func() {
				c.Header(header)
				wg.Done()
			}()
		case "MaterialAssignment":
			func() {
				c.MaterialAssignment(materialAssignment)
				wg.Done()
			}()
		case "Sequence":
			func() {
				c.Sequence(sequence)
				wg.Done()
			}()
		case "Operation":
			func() {
				c.Operation(operation)
				wg.Done()
			}()
		case "ComponentAllocation":
			func() {
				c.ComponentAllocation(componentAllocation)
				wg.Done()
			}()
		default:
			wg.Done()
		}
	}

	wg.Wait()
}

func (c *SAPAPICaller) Header(header *requests.Header) {
	outputDataHeader, err := c.callProductionRoutingSrvAPIRequirementHeader("ProductionRoutingHeader", header)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(outputDataHeader)
}

func (c *SAPAPICaller) callProductionRoutingSrvAPIRequirementHeader(api string, header *requests.Header) (*sap_api_output_formatter.Header, error) {
	body, err := json.Marshal(header)
	if err != nil {
		return nil, xerrors.Errorf("API request error: %w", err)
	}
	url := strings.Join([]string{c.baseURL, "API_PRODUCTION_ROUTING", api}, "/")
	params := c.addQuerySAPClient(map[string]string{})
	resp, err := c.requestClient.Request("POST", url, params, string(body))
	if err != nil {
		return nil, xerrors.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()
	byteArray, _ := ioutil.ReadAll(resp.Body)
	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
		return nil, xerrors.Errorf("bad response:%s", string(byteArray))
	}

	data, err := sap_api_output_formatter.ConvertToHeader(byteArray, c.log)
	if err != nil {
		return nil, xerrors.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) MaterialAssignment(materialAssignment *requests.MaterialAssignment) {
	outputDataMaterialAssignment, err := c.callProductionRoutingSrvAPIRequirementMaterialAssignment("ProductionRoutingMatlAssgmt", materialAssignment)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(outputDataMaterialAssignment)
}

func (c *SAPAPICaller) callProductionRoutingSrvAPIRequirementMaterialAssignment(api string, materialAssignment *requests.MaterialAssignment) (*sap_api_output_formatter.MaterialAssignment, error) {
	body, err := json.Marshal(materialAssignment)
	if err != nil {
		return nil, xerrors.Errorf("API request error: %w", err)
	}
	url := strings.Join([]string{c.baseURL, "API_PRODUCTION_ROUTING", api}, "/")
	params := c.addQuerySAPClient(map[string]string{})
	resp, err := c.requestClient.Request("POST", url, params, string(body))
	if err != nil {
		return nil, xerrors.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()
	byteArray, _ := ioutil.ReadAll(resp.Body)
	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
		return nil, xerrors.Errorf("bad response:%s", string(byteArray))
	}

	data, err := sap_api_output_formatter.ConvertToMaterialAssignment(byteArray, c.log)
	if err != nil {
		return nil, xerrors.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) Sequence(sequence *requests.Sequence) {
	outputDataSequence, err := c.callProductionRoutingSrvAPIRequirementSequence("ProductionRoutingSequence", sequence)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(outputDataSequence)
}

func (c *SAPAPICaller) callProductionRoutingSrvAPIRequirementSequence(api string, sequence *requests.Sequence) (*sap_api_output_formatter.Sequence, error) {
	body, err := json.Marshal(sequence)
	if err != nil {
		return nil, xerrors.Errorf("API request error: %w", err)
	}
	url := strings.Join([]string{c.baseURL, "API_PRODUCTION_ROUTING", api}, "/")
	params := c.addQuerySAPClient(map[string]string{})
	resp, err := c.requestClient.Request("POST", url, params, string(body))
	if err != nil {
		return nil, xerrors.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()
	byteArray, _ := ioutil.ReadAll(resp.Body)
	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
		return nil, xerrors.Errorf("bad response:%s", string(byteArray))
	}

	data, err := sap_api_output_formatter.ConvertToSequence(byteArray, c.log)
	if err != nil {
		return nil, xerrors.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) Operation(operation *requests.Operation) {
	outputDataOperation, err := c.callProductionRoutingSrvAPIRequirementOperation("ProductionRoutingOperation", operation)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(outputDataOperation)
}

func (c *SAPAPICaller) callProductionRoutingSrvAPIRequirementOperation(api string, operation *requests.Operation) (*sap_api_output_formatter.Operation, error) {
	body, err := json.Marshal(operation)
	if err != nil {
		return nil, xerrors.Errorf("API request error: %w", err)
	}
	url := strings.Join([]string{c.baseURL, "API_PRODUCTION_ROUTING", api}, "/")
	params := c.addQuerySAPClient(map[string]string{})
	resp, err := c.requestClient.Request("POST", url, params, string(body))
	if err != nil {
		return nil, xerrors.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()
	byteArray, _ := ioutil.ReadAll(resp.Body)
	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
		return nil, xerrors.Errorf("bad response:%s", string(byteArray))
	}

	data, err := sap_api_output_formatter.ConvertToOperation(byteArray, c.log)
	if err != nil {
		return nil, xerrors.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) ComponentAllocation(componentAllocation *requests.ComponentAllocation) {
	outputDataComponentAllocation, err := c.callProductionRoutingSrvAPIRequirementComponentAllocation("ProductionRoutingOpCompAlloc", componentAllocation)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(outputDataComponentAllocation)
}

func (c *SAPAPICaller) callProductionRoutingSrvAPIRequirementComponentAllocation(api string, componentAllocation *requests.ComponentAllocation) (*sap_api_output_formatter.ComponentAllocation, error) {
	body, err := json.Marshal(componentAllocation)
	if err != nil {
		return nil, xerrors.Errorf("API request error: %w", err)
	}
	url := strings.Join([]string{c.baseURL, "API_PRODUCTION_ROUTING", api}, "/")
	params := c.addQuerySAPClient(map[string]string{})
	resp, err := c.requestClient.Request("POST", url, params, string(body))
	if err != nil {
		return nil, xerrors.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()
	byteArray, _ := ioutil.ReadAll(resp.Body)
	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
		return nil, xerrors.Errorf("bad response:%s", string(byteArray))
	}

	data, err := sap_api_output_formatter.ConvertToComponentAllocation(byteArray, c.log)
	if err != nil {
		return nil, xerrors.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) addQuerySAPClient(params map[string]string) map[string]string {
	if len(params) == 0 {
		params = make(map[string]string, 1)
	}
	params["sap-client"] = c.sapClientNumber
	return params
}
