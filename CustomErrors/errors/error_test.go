package errors_test

import (
	. "github.com/b-entangled/GoExamples/CustomErrors/errors"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Error", func() {
	var (
		customError1 Error
		customError2 Error
	)
	BeforeEach(func() {
		customError1 = CError("CustomError1", 1)
		customError2 = CError("CustomError2", 2)
	})
	Describe("CustomError1", func() {
		Context("Success", func() {
			It("Message", func() {
				Expect(customError1.Error()).To(Equal("Error: CustomError1, Code: 1"))
			})
			It("Code", func() {
				Expect(customError1.Code()).To(Equal(1))
			})
		})
	})
	Describe("CustomError2", func() {
		Context("Success", func() {
			It("Message", func() {
				Expect(customError2.Error()).To(Equal("Error: CustomError2, Code: 2"))
			})
			It("Code", func() {
				Expect(customError2.Code()).To(Equal(2))
			})
		})
	})
})
