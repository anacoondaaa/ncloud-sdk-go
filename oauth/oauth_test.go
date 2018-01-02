package oauth_test

import (
	"net/http"

	gock "gopkg.in/h2non/gock.v1"
	. "github.com/NaverCloudPlatform/ncloud-sdk-go/oauth"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Oauth", func() {
	BeforeEach(func() {
		gock.New("https://api.ncloud.com").
			Get("/server").
			Reply(200)
	})
	AfterEach(func() {
		gock.Off()
	})
	It("should get server instance list", func() {
		const requestMethod = "GET"
		const requestURL = "https://api.ncloud.com/server/"

		c := NewConsumer(accessKey, secretKey, requestMethod, requestURL)

		c.AdditionalParams["action"] = "createServerInstances"
		c.AdditionalParams["responseFormatType"] = "xml"
		c.AdditionalParams["serverImageProductCode"] = "SPSW0LINUX000046"
		c.AdditionalParams["serverProductCode"] = "SPSVRSSD00000003"
		c.AdditionalParams["serverCreateCount"] = "2"
		c.AdditionalParams["serverCreateStartNo"] = "001"

		reqUrl, err := c.GetRequestUrl()
		Expect(err).To(BeNil())

		req, err := http.NewRequest("GET", reqUrl, nil)
		Expect(err).To(BeNil())

		client := &http.Client{}
		resp, err := client.Do(req)
		Expect(err).To(BeNil())

		defer resp.Body.Close()

		Expect(resp.StatusCode).To(Equal(200))
	})
})
