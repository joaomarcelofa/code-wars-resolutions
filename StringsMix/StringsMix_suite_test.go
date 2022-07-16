package stringsmix

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestStringsMix(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "StringsMix Suite")
}

func dotest(a1 string, a2 string, exp string) {
	var ans = Mix(a1, a2)
	Expect(ans).To(Equal(exp))
}

var _ = Describe("Tests Mix", func() {

	It("should handle basic cases 1", func() {
		dotest("looping is fun but dangerous", "less dangerous than coding", "1:ooo/1:uuu/2:sss/=:nnn/1:ii/2:aa/2:dd/2:ee/=:gg")
		dotest("Are they here", "yes, they are here", "2:eeeee/2:yy/=:hh/=:rr")
		dotest("uuuuuu", "uuuuuu", "=:uuuuuu")
	})
})
