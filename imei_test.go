package imei_test

import (
	"testing"

	"github.com/onsi/ginkgo/v2"
	"github.com/onsi/gomega"

	"entrlcom.dev/imei"
)

func Test(t *testing.T) {
	t.Parallel()

	gomega.RegisterFailHandler(ginkgo.Fail)
	ginkgo.RunSpecs(t, "IMEI Test Suite")
}

var _ = ginkgo.Describe("NewIMEI", func() {
	ginkgo.It("IMEI", func() {
		v, err := imei.NewIMEI("35-209900-176148-1")
		gomega.Expect(err).To(gomega.Succeed())
		gomega.Expect(v.CD().String()).To(gomega.Equal("1"))
		gomega.Expect(v.IsIMEI()).To(gomega.BeTrue())
		gomega.Expect(v.IsIMEISV()).To(gomega.BeFalse())
		gomega.Expect(v.SNR().String()).To(gomega.Equal("176148"))
		gomega.Expect(v.String()).To(gomega.Equal("35 209900 176148 1"))
		gomega.Expect(v.SVN().IsZero()).To(gomega.BeTrue())
		gomega.Expect(v.TAC().RBI().String()).To(gomega.Equal("35"))
		gomega.Expect(v.TAC().ID()).To(gomega.Equal("209900"))
	})

	ginkgo.It("IMEISV", func() {
		v, err := imei.NewIMEI("35-209900-176148-23")
		gomega.Expect(err).To(gomega.Succeed())
		gomega.Expect(v.CD().IsZero()).To(gomega.BeTrue())
		gomega.Expect(v.IsIMEI()).To(gomega.BeFalse())
		gomega.Expect(v.IsIMEISV()).To(gomega.BeTrue())
		gomega.Expect(v.SNR().String()).To(gomega.Equal("176148"))
		gomega.Expect(v.String()).To(gomega.Equal("35 209900 176148 23"))
		gomega.Expect(v.SVN().String()).To(gomega.Equal("23"))
		gomega.Expect(v.TAC().RBI().String()).To(gomega.Equal("35"))
		gomega.Expect(v.TAC().ID()).To(gomega.Equal("209900"))
	})
})
