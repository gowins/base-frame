package env

import (
	"os"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestSetup(t *testing.T) {
	Convey("setup", t, func() {
		_ = os.Setenv(APP_ENV_KEY, "abcd")
		So(Setup, ShouldPanic)
		_ = os.Setenv(APP_ENV_KEY, "test")
		So(Setup, ShouldNotPanic)
		_ = os.Setenv(APP_ENV_KEY, "develop")
		So(Setup, ShouldNotPanic)
	})
}

func TestGetRunEnv(t *testing.T) {
	Setup()
	t.Log(GetRunEnv())
}

func TestIsDev(t *testing.T) {
	Convey("ISDev", t, func() {
		_ = os.Setenv(APP_ENV_KEY, "test")
		Setup()
		So(IsDev(), ShouldBeFalse)
		_ = os.Setenv(APP_ENV_KEY, "")
		Setup()
		So(IsDev(), ShouldBeTrue)
	})
}

func TestIsTest(t *testing.T) {
	Convey("ISTest", t, func() {
		_ = os.Setenv(APP_ENV_KEY, "test")
		Setup()
		So(IsTest(), ShouldBeTrue)
		_ = os.Setenv(APP_ENV_KEY, "")
		Setup()
		So(IsTest(), ShouldBeFalse)
	})
}

func TestIsProduct(t *testing.T) {
	Convey("IsProduct", t, func() {
		_ = os.Setenv(APP_ENV_KEY, "product")
		Setup()
		So(IsProduct(), ShouldBeTrue)
		_ = os.Setenv(APP_ENV_KEY, "")
		Setup()
		So(IsProduct(), ShouldBeFalse)
	})
}
