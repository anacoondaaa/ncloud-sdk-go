package sdk

import (
	"encoding/xml"
	"errors"
	"fmt"
	"net/http"

	common "github.com/NaverCloudPlatform/ncloud-sdk-go/common"
	request "github.com/NaverCloudPlatform/ncloud-sdk-go/request"
)

func processDeleteBlockStorageInstancesParams(blockStorageInstanceNoList []string) (map[string]string, error) {
	params := make(map[string]string)

	if len(blockStorageInstanceNoList) == 0 {
		return params, errors.New("BlockStorageInstanceNoList is required field")
	}

	for k, v := range blockStorageInstanceNoList {
		params[fmt.Sprintf("blockStorageInstanceNoList.%d", k+1)] = v
	}

	return params, nil
}

// DeleteBlockStorageInstances delete block storage instances
func (s *Conn) DeleteBlockStorageInstances(blockStorageInstanceNoList []string) (*BlockStorageInstanceList, error) {
	params, err := processDeleteBlockStorageInstancesParams(blockStorageInstanceNoList)
	if err != nil {
		return nil, err
	}

	bytes, resp, err := request.NewRequest(s.accessKey, s.secretKey, "POST", s.apiURL, "/server/v2/deleteBlockStorageInstances", params)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		responseError, err := common.ParseErrorResponse(bytes)
		if err != nil {
			return nil, err
		}

		respError := BlockStorageInstanceList{}
		respError.ReturnCode = responseError.ReturnCode
		respError.ReturnMessage = responseError.ReturnMessage

		return &respError, fmt.Errorf("%s %s - error code: %d , error message: %s", resp.Status, string(bytes), responseError.ReturnCode, responseError.ReturnMessage)
	}

	var blockStorageInstanceList = BlockStorageInstanceList{}
	if err := xml.Unmarshal([]byte(bytes), &blockStorageInstanceList); err != nil {
		fmt.Println(err)
		return nil, err
	}

	return &blockStorageInstanceList, nil
}
