package sdk

import (
	"encoding/xml"
	"errors"
	"fmt"
	"net/http"

	common "github.com/NaverCloudPlatform/ncloud-sdk-go/common"
	request "github.com/NaverCloudPlatform/ncloud-sdk-go/request"
)

func processCreatePublicIPInstanceParams(reqParams *RequestCreatePublicIPInstance) (map[string]string, error) {
	params := make(map[string]string)

	if reqParams.ServerInstanceNo != "" {
		params["serverInstanceNo"] = reqParams.ServerInstanceNo
	}

	if reqParams.PublicIPDescription != "" {
		if len := len(reqParams.PublicIPDescription); len > 1000 {
			return params, errors.New("Length of publicIpDescription should be max 1000")
		}
		params["publicIpDescription"] = reqParams.PublicIPDescription
	}

	if reqParams.InternetLineTypeCode != "" {
		if reqParams.InternetLineTypeCode != "PUBLC" && reqParams.InternetLineTypeCode != "GLBL" {
			return params, errors.New("InternetLineTypeCode should be PUBLC or GLBL")
		}
		params["internetLineTypeCode"] = reqParams.InternetLineTypeCode
	}

	if reqParams.RegionNo != "" {
		params["regionNo"] = reqParams.RegionNo
	}

	if reqParams.ZoneNo != "" {
		params["zoneNo"] = reqParams.ZoneNo
	}

	return params, nil
}

// CreatePublicIPInstance create public ip instance and allocate it to server instance
func (s *Conn) CreatePublicIPInstance(reqParams *RequestCreatePublicIPInstance) (*PublicIPInstanceList, error) {
	params, err := processCreatePublicIPInstanceParams(reqParams)
	if err != nil {
		return nil, err
	}

	bytes, resp, err := request.NewRequest(s.accessKey, s.secretKey, "POST", s.apiURL, "/server/v2/createPublicIpInstance", params)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		responseError, err := common.ParseErrorResponse(bytes)
		if err != nil {
			return nil, err
		}

		respError := PublicIPInstanceList{}
		respError.ReturnCode = responseError.ReturnCode
		respError.ReturnMessage = responseError.ReturnMessage

		return &respError, fmt.Errorf("%s %s - error code: %d , error message: %s", resp.Status, string(bytes), responseError.ReturnCode, responseError.ReturnMessage)
	}

	var responseCreatePublicIPInstances = PublicIPInstanceList{}
	if err := xml.Unmarshal([]byte(bytes), &responseCreatePublicIPInstances); err != nil {
		fmt.Println(err)
		return nil, err
	}

	return &responseCreatePublicIPInstances, nil
}
