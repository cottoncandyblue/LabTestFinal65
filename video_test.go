package Entity

import(
	"testing"
	"github.com/asaskevich/govalidator"
	. "github.com/onsi/gomega"
	"gorm.io/gorm"

)

type Video struct {
	gorm.Model
	Name string `valid:"required~Name cannot be blank"`
	Url  string `gorm:"uniqueIndex" valid:"url"`
}

func TestValidation(t *testing.T) {

	g := NewGomegaWithT(t)

	t.Run("check data valid", func(t *testing.T) {

		u := Video{

			Name: "KKK",
			Url: "http://www.youtube.com/",
		}

		ok, err := govalidator.ValidateStruct(u)

		g.Expect(ok).To(BeTrue())
		g.Expect(err).To(BeNil())
	})
}

func TestName(t *testing.T) {
	g:= NewGomegaWithT(t)

	t.Run("check name be blank", func(t *testing.T) {

		u := Video{
			Name: "",
			Url: "http://www.youtube.com/",
		}

		ok, err := govalidator.ValidateStruct(u)
		g.Expect(ok).NotTo(BeTrue())
		g.Expect(err).ToNot(BeNil())
		g.Expect(err.Error()).To(Equal("Name cannot be blank"))
	})

}

func TestUrl(t *testing.T) {

	g := NewGomegaWithT(t)

	t.Run("check Url is vaild", func(t *testing.T) {

		u := Video{
			Name: "kkkkk",
			Url: "://www.youtubess.com/",
		}

		ok, err := govalidator.ValidateStruct(u)

		g.Expect(ok).NotTo(BeTrue())
		g.Expect(err).ToNot(BeNil())
		g.Expect(err.Error()).To(Equal("Url: ://www.youtubess.com/ does not validate as url"))

	})

}
