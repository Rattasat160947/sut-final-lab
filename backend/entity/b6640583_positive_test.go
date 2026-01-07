package entity

import (
	"testing"

	"github.com/asaskevich/govalidator"
	. "github.com/onsi/gomega"
)

func TestValid(t *testing.T) {
	g := NewGomegaWithT(t)
	t.Run(`book is valid`, func(t *testing.T) {
		books := Books{
			Title: "SUT",
			Price: 500,
			Code:  "BK123456",
		}
		ok, err := govalidator.ValidateStruct(books)
		g.Expect(ok).To(BeTrue())
		g.Expect(err).To(BeNil())
	})
}
