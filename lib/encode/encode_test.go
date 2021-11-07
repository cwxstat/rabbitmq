package encode_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/cwxstat/rabbitmq/lib/encode"
)

var _ = Describe("Encode", func() {

	var (
		file      string
		msgString string
	)

	BeforeEach(func() {
		file = "readFile"
		msgString = "This is message"
	})

	Describe("Check Encoding", func() {

		Context("Should encode", func() {
			It("should send messages", func() {

				encoded := encode.Encode([]byte(msgString))
				err := encode.WriteDecode(file, encoded)

				Expect(err).To(BeNil())

				result, err := encode.ReadEncode(file)
				Expect(err).To(BeNil())
				decoded, err := encode.Decode(result)
				Expect(string(decoded)).To(Equal(msgString))
				Expect(err).To(BeNil())

			})
		})

	})

})
