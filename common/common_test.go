package common_test

import (
	. "github.com/anacoondaaa/ncloud-sdk-go/common"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Common", func() {
	It("should parse error response", func() {
		response, err := ParseErrorResponse([]byte(`<responseError>
			<returnCode>800</returnCode>
			<returnMessage>Expired url.</returnMessage>
			</responseError>`))

		Expect(err).To(BeNil())
		Expect(response.ReturnCode).To(Equal(800))
		Expect(response.ReturnMessage).To(Equal("Expired url."))
	})

})
