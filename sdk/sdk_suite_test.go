package sdk_test

import (
	"math/rand"
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

var accessKey, secretKey string

const charset = "abcdefghijklmnopqrstuvwxyz" +
	"ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

var seededRand *rand.Rand = rand.New(
	rand.NewSource(time.Now().UnixNano()))

func StringWithCharset(length int, charset string) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

func String(length int) string {
	return StringWithCharset(length, charset)
}

func TestSDK(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "SDK Suite")
}

var _ = BeforeSuite(func() {
	accessKey = "YZ8WMQ0Mo2D7Sy3HgcFj"
	secretKey = "pYN0Lld36TipU6FbJ05TQO76nVqaDXPnvm5Q4ZnE"
})
