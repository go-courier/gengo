package gengo

import (
	"bytes"
	testingx "github.com/octohelm/x/testing"
	"reflect"
	"testing"

	"github.com/octohelm/gengo/pkg/namer"
)

type Item struct {
	Name string `json:"name"`
}

type List[T any] struct {
	Items []T `json:"items,omitempty"`
}

func TestDumper_TypeLit(t *testing.T) {
	d := NewDumper(namer.NewRawNamer("", namer.NewDefaultImportTracker()))

	t.Run("TypeLit", func(t *testing.T) {
		testingx.Expect(t, "*bytes.Buffer", testingx.Be(d.ReflectTypeLit(reflect.TypeOf(&bytes.Buffer{}))))
		testingx.Expect(t, "[]string", testingx.Be(d.ReflectTypeLit(reflect.TypeOf([]string{}))))
		testingx.Expect(t, "map[string]string", testingx.Be(d.ReflectTypeLit(reflect.TypeOf(map[string]string{}))))
		testingx.Expect(t, "*struct {V string `json:\"v\" validate:\"@int[0,10]\"`\n}", testingx.Be(d.ReflectTypeLit(reflect.TypeOf(&struct {
			V string `json:"v" validate:"@int[0,10]"`
		}{}))))
	})

	t.Run("TypeListWithGenerics", func(t *testing.T) {
		testingx.Expect(t,
			"*github_com_octohelm_gengo_pkg_gengo.List[github_com_octohelm_gengo_pkg_gengo.Item]",
			testingx.Equal(d.ReflectTypeLit(reflect.TypeOf(&List[Item]{}))))

		testingx.Expect(t,
			"*github_com_octohelm_gengo_pkg_gengo.List[github_com_octohelm_gengo_pkg_gengo.List[github_com_octohelm_gengo_pkg_gengo.Item]]",
			testingx.Equal(d.ReflectTypeLit(reflect.TypeOf(&List[List[Item]]{}))))
	})

	t.Run("ValueLit", func(t *testing.T) {
		testingx.Expect(t, "&(bytes.Buffer{})", testingx.Be(d.ValueLit(reflect.ValueOf(&(bytes.Buffer{})))))
		testingx.Expect(t, `[]string{
"1",
"2",
}`, testingx.Be(d.ValueLit(reflect.ValueOf([]string{"1", "2"}))))
	})
}
