package sdk

import (
	"encoding/xml"
	"errors"
	"fmt"
	"net/http"

	"github.com/NaverCloudPlatform/ncloud-sdk-go/common"
	"github.com/NaverCloudPlatform/ncloud-sdk-go/request"
)

func processDeletePortForwardingRules(reqParams *RequestDeletePortForwardingRules) (map[string]string, error) {
	params := make(map[string]string)

	if reqParams == nil {
		return nil, fmt.Errorf("PortForwardingConfigurationNo is required field")
	}
	if err := validateRequiredField("PortForwardingConfigurationNo", reqParams.PortForwardingConfigurationNo); err != nil {
		return nil, err
	}
	params["portForwardingConfigurationNo"] = reqParams.PortForwardingConfigurationNo

	if len(reqParams.PortForwardingRuleList) == 0 {
		return nil, errors.New("PortForwardingRuleList is required")
	}
	for k, v := range reqParams.PortForwardingRuleList {
		params[fmt.Sprintf("portForwardingRuleList.%d.serverInstanceNo", k+1)] = v.ServerInstanceNo
		params[fmt.Sprintf("portForwardingRuleList.%d.portForwardingInternalPort", k+1)] = v.PortForwardingInternalPort
		params[fmt.Sprintf("portForwardingRuleList.%d.portForwardingExternalPort", k+1)] = v.PortForwardingExternalPort
	}

	return params, nil
}

func (s *Conn) DeletePortForwardingRules(reqParams *RequestDeletePortForwardingRules) (*PortForwardingRuleList, error) {
	params, err := processDeletePortForwardingRules(reqParams)

	if err != nil {
		return nil, err
	}

	bytes, resp, err := request.NewRequest(s.accessKey, s.secretKey, "POST", s.apiURL, "/server/v2/deletePortForwardingRules", params)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		responseError, err := common.ParseErrorResponse(bytes)
		if err != nil {
			return nil, err
		}

		respError := PortForwardingRuleList{}
		respError.ReturnCode = responseError.ReturnCode
		respError.ReturnMessage = responseError.ReturnMessage

		return &respError, fmt.Errorf("%s %s - error code: %d , error message: %s", resp.Status, string(bytes), responseError.ReturnCode, responseError.ReturnMessage)
	}

	var portForwardingRuleList = PortForwardingRuleList{}
	if err := xml.Unmarshal([]byte(bytes), &portForwardingRuleList); err != nil {
		fmt.Println(err)
		return nil, err
	}

	return &portForwardingRuleList, nil
}
