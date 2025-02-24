package namer

import (
	"go/types"
	"strings"

	gengotypes "github.com/octohelm/gengo/pkg/types"
)

type Namer interface {
	Name(gengotypes.TypeName) string
}

type NameSystems map[string]Namer

type Names map[gengotypes.TypeName]string

func NewRawNamer(pkgPath string, tracker ImportTracker) Namer {
	return &rawNamer{pkgPath: pkgPath, tracker: tracker}
}

type rawNamer struct {
	pkgPath string
	tracker ImportTracker
	Names
}

func (n *rawNamer) Name(typeName gengotypes.TypeName) string {
	if n.Names == nil {
		n.Names = Names{}
	}

	if name, ok := n.Names[typeName]; ok {
		return name
	}

	pkgPath := typeName.Pkg().Path()
	tn := &strings.Builder{}

	tn.WriteString(n.processName(typeName.Name()))

	if x, ok := typeName.(*types.TypeName); ok {
		if named, ok := x.Type().(*types.Named); ok {
			if p := named.TypeParams(); p != nil {
				tn.WriteString("[")

				for i := 0; i < p.Len(); i++ {
					if i > 0 {
						tn.WriteString(",")
					}
					tn.WriteString(p.At(i).String())
				}

				tn.WriteString("]")
			}
		}
	}

	if pkgPath == n.pkgPath {
		if tn.Len() != 0 {
			return tn.String()
		}
		return typeName.String()
	} else {
		n.tracker.AddType(typeName)
		return n.tracker.LocalNameOf(pkgPath) + "." + tn.String()
	}
}

func (n *rawNamer) processName(name string) string {
	t, err := gengotypes.ParseTypeRef(name)
	if err != nil {
		panic(err)
	}

	if len(t.TypeList) == 0 {
		return t.Name
	}

	for x := range t.Walk {
		if x.PkgPath == "" {
			continue
		}

		if x.PkgPath == n.pkgPath {
			x.PkgPath = ""
		} else {
			n.tracker.AddType(gengotypes.Ref(x.PkgPath, x.Name))
			x.PkgPath = n.tracker.LocalNameOf(x.PkgPath)
		}
	}

	return t.String()
}
