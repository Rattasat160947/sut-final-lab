package entity

import (
	"testing"
	"github.com/asaskevich/govalidator"
	. "github.com/onsi/gomega"
)

func TestPrice(t *testing.T){
	g := NewGomegaWithT(t)
	t.Run(`Price must be between 50 and 5000` , func(t *testing.T){
		books := Books{
			Title: "SUT",
			Price: 10,
			Code:  "BK123456",
		}
		ok, err := govalidator.ValidateStruct(books)
		g.Expect(ok).NotTo(BeTrue())
		g.Expect(err).NotTo(BeNil())
		g.Expect(err.Error()).To(Equal("Price must be between 50 and 5000"))
	})
}