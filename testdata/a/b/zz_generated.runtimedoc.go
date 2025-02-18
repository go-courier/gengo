/*
Package b GENERATED BY gengo:runtimedoc
DON'T EDIT THIS FILE
*/
package b

import (
	embed "embed"
)

func (*B) RuntimeDoc(names ...string) ([]string, bool) {
	return []string{
		"is a type for testing",
	}, true
}

func (v *List[T]) RuntimeDoc(names ...string) ([]string, bool) {
	if len(names) > 0 {
		switch names[0] {
		case "Name":
			return []string{}, true
		}

		return nil, false
	}
	return []string{}, true
}

//go:embed doc/b.md
var embedDocOfObj1 string
var _ embed.FS

func (v *Obj) RuntimeDoc(names ...string) ([]string, bool) {
	if len(names) > 0 {
		switch names[0] {
		case "Name":
			return []string{
				"name",
				"姓名",
			}, true
		}
		if doc, ok := runtimeDoc(&v.SubObj, "", names...); ok {
			return doc, ok
		}

		return nil, false
	}
	return []string{
		"some object",
		embedDocOfObj1,
	}, true
}

func (v *SubObj) RuntimeDoc(names ...string) ([]string, bool) {
	if len(names) > 0 {
		switch names[0] {
		case "Age":
			return []string{}, true
		}

		return nil, false
	}
	return []string{}, true
}

func (v *Third) RuntimeDoc(names ...string) ([]string, bool) {
	if len(names) > 0 {
		switch names[0] {
		case "Path":
			return []string{
				"is a module path, like \"golang.org/x/text\" or \"rsc.io/quote/v2\".",
			}, true
		case "Version":
			return []string{
				"is usually a semantic version in canonical form.",
				"There are three exceptions to this general rule.",
				"First, the top-level target of a build has no specific version",
				"and uses Version = \"\".",
				"Second, during MVS calculations the version \"none\" is used",
				"to represent the decision to take no version of a given module.",
				"Third, filesystem paths found in \"replace\" directives are",
				"represented by a path with an empty version.",
			}, true

		}

		return nil, false
	}
	return []string{}, true
}

// nolint:deadcode,unused
func runtimeDoc(v any, prefix string, names ...string) ([]string, bool) {
	if c, ok := v.(interface {
		RuntimeDoc(names ...string) ([]string, bool)
	}); ok {
		doc, ok := c.RuntimeDoc(names...)
		if ok {
			if prefix != "" && len(doc) > 0 {
				doc[0] = prefix + doc[0]
				return doc, true
			}

			return doc, true
		}
	}
	return nil, false
}
