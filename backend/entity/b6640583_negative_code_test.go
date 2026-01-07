package entity

import (
	"testing"
	"github.com/asaskevich/govalidator"
	. "github.com/onsi/gomega"
)

func TestCode(t *testing.T){
	g := NewGomegaWithT(t)
	t.Run(`Code must start with BK followed by 6 digits (0-9)` , func(t *testing.T){
		books := Books{
			Title :"SUT",
			Price :500,
			Code : "B123456",
		}
		ok, err := govalidator.ValidateStruct(books)
		g.Expect(ok).NotTo(BeTrue())
		g.Expect(err).NotTo(BeNil())
		g.Expect(err.Error()).To(Equal("Code must start with BK followed by 6 digits (0-9)"))
	})
}