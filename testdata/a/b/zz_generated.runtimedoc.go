/*
Package b GENERATED BY gengo:runtimedoc
DON'T EDIT THIS FILE
*/
package b

type canRuntimeDoc interface {
	RuntimeDoc(names ...string) ([]string, bool)
}

func runtimeDoc(v any, names ...string) ([]string, bool) {
	if c, ok := v.(canRuntimeDoc); ok {
		return c.RuntimeDoc(names...)
	}
	return nil, false
}

func (B) RuntimeDoc(names ...string) ([]string, bool) {
	return []string{
		"B is a type for testing",
	}, true
}
func (v Obj) RuntimeDoc(names ...string) ([]string, bool) {
	if len(names) > 0 {
		switch names[0] {
		case "Name":
			return []string{
				"name",
				"姓名",
			}, true
		case "SubObj":
			return []string{}, true

		}
		if doc, ok := runtimeDoc(v.SubObj, names...); ok {
			return doc, ok
		}

		return nil, false
	}
	return []string{}, true
}

func (v SubObj) RuntimeDoc(names ...string) ([]string, bool) {
	if len(names) > 0 {
		switch names[0] {
		case "Age":
			return []string{
				"Age",
			}, true

		}

		return nil, false
	}
	return []string{}, true
}

func (v Third) RuntimeDoc(names ...string) ([]string, bool) {
	if len(names) > 0 {
		switch names[0] {
		case "VCS":
			return []string{}, true
		case "Repo":
			return []string{
				"Repo is the repository URL, including scheme.",
			}, true
		case "Root":
			return []string{
				"Root is the import path corresponding to the root of the",
				"repository.",
			}, true

		}

		return nil, false
	}
	return []string{}, true
}
