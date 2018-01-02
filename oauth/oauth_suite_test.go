package oauth_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

const accessKey = "XYt3YgyYWSqQfhrotmca"
const secretKey = "7mlqMgs9QxeT1ybgUWZJI4z6VAVmK8goHHIABYdM"

func TestOauth(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Oauth Suite")
}
