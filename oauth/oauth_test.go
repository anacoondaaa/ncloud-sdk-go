package oauth_test

import (
	"net/http"

	. "github.com/NaverCloudPlatform/ncloud-sdk-go/oauth"
	gock "gopkg.in/h2non/gock.v1"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Oauth", func() {
	Describe("Get method", func() {
		BeforeEach(func() {
			gock.New("https://api.ncloud.com").
				Get("/server").
				Reply(http.StatusOK)
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

			reqUrl, _, err := c.GetRequest()
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

	Describe("Post method", func() {
		BeforeEach(func() {
			gock.New("https://api.ncloud.com").
				Post("/server").
				Reply(http.StatusOK)
		})
		AfterEach(func() {
			gock.Off()
		})
		It("should create server instance list", func() {
			const requestMethod = "POST"
			const requestURL = "https://api.ncloud.com/server/"

			c := NewConsumer(accessKey, secretKey, requestMethod, requestURL)

			c.AdditionalParams["action"] = "createLoginKey"
			c.AdditionalParams["responseFormatType"] = "xml"
			c.AdditionalParams["keyName"] = "keyName"

			reqUrl, body, err := c.GetRequest()
			Expect(err).To(BeNil())

			req, err := http.NewRequest("POST", reqUrl, body)
			Expect(err).To(BeNil())

			client := &http.Client{}
			resp, err := client.Do(req)
			Expect(err).To(BeNil())

			defer resp.Body.Close()

			Expect(resp.StatusCode).To(Equal(200))
		})
	})
})
