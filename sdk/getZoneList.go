package sdk

import (
	"encoding/xml"
	"fmt"
	"net/http"

	common "github.com/NaverCloudPlatform/ncloud-sdk-go/common"
	request "github.com/NaverCloudPlatform/ncloud-sdk-go/request"
)

// GetZoneList get zone list
func (s *Conn) GetZoneList(regionNo string) (*ZoneList, error) {
	params := make(map[string]string)

	if regionNo != "" {
		params["regionNo"] = regionNo
	}

	bytes, resp, err := request.NewRequest(s.accessKey, s.secretKey, "GET", s.apiURL, "/server/v2/getZoneList", params)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		responseError, err := common.ParseErrorResponse(bytes)
		if err != nil {
			return nil, err
		}

		respError := ZoneList{}
		respError.ReturnCode = responseError.ReturnCode
		respError.ReturnMessage = responseError.ReturnMessage

		return &respError, fmt.Errorf("%s %s - error code: %d , error message: %s", resp.Status, string(bytes), responseError.ReturnCode, responseError.ReturnMessage)
	}

	ZoneList := ZoneList{}
	if err := xml.Unmarshal([]byte(bytes), &ZoneList); err != nil {
		return nil, err
	}

	return &ZoneList, nil
}
