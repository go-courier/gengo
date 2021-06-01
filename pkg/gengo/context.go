package gengo

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"os"
	"path"
	"sort"
	"strings"

	gengotypes "github.com/go-courier/gengo/pkg/types"
	"github.com/pkg/errors"
)

func NewContext(args *GeneratorArgs) (*Context, error) {
	u, err := gengotypes.Load(args.Inputs)
	if err != nil {
		return nil, err
	}
	c := &Context{
		Args:     args,
		Universe: u,
	}
	return c, nil
}

type Context struct {
	Args     *GeneratorArgs
	Universe gengotypes.Universe

	Package gengotypes.Package
	Tags    map[string][]string
}

type CreateGenerator = func() Generator

func (c *Context) Execute(ctx context.Context, createGenerators ...CreateGenerator) error {
	for _, pkg := range c.Args.Inputs {
		if err := c.execPkg(ctx, pkg, createGenerators...); err != nil {
			return err
		}
	}
	return nil
}

func (c *Context) execPkg(ctx context.Context, pkg string, createGenerators ...CreateGenerator) error {
	p, ok := c.Universe[pkg]
	if !ok {
		return errors.Errorf("invalid pkg `%s`", pkg)
	}

	ctxWithPkg := &Context{
		Args:     c.Args,
		Universe: c.Universe,
		Package:  p,
		Tags:     map[string][]string{},
	}

	for _, f := range p.Files() {
		if f.Doc != nil && len(f.Doc.List) > 0 {
			tags := gengotypes.ExtractCommentTags("+", strings.Split(f.Doc.Text(), "\n"))

			for k := range tags {
				ctxWithPkg.Tags[k] = tags[k]
			}
		}
	}

	for i := range createGenerators {
		createGenerator := createGenerators[i]

		g := createGenerator()

		generatorEnabled := "gengo:" + g.Name()

		if _, ok := ctxWithPkg.Tags[generatorEnabled]; ok {
			if err := ctxWithPkg.doGenerating(ctx, g); err != nil {
				return errors.Wrapf(err, "`%s` generate failed for %s", g.Name(), ctxWithPkg.Package.Pkg().Path())
			}
		}
	}

	return nil
}

func (c *Context) doGenerating(ctx context.Context, g Generator) error {
	if c.Package == nil {
		return nil
	}

	filename := path.Join(c.Package.SourceDir(), fmt.Sprintf("%s.%s.go", c.Args.OutputFileBaseName, g.Name()))

	f, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		if os.IsNotExist(err) {
			f, err = os.Create(filename)
			if err != nil {
				return err
			}
		}
		return err
	}
	defer f.Close()

	if err := g.Init(c, f); err != nil {
		return err
	}

	types := c.Package.Types()

	body := bytes.NewBuffer(nil)

	for i := range types {
		if err := g.GenerateType(c, types[i].Type(), body); err != nil {
			return err
		}
	}

	if _, err := fmt.Fprintf(f, `// GENERATED BY gengo:%s DON'T EDIT
package %s
`, g.Name(), c.Package.Pkg().Name()); err != nil {
		return err
	}

	writeImports(g.Imports(c), f)

	if _, err := io.Copy(f, body); err != nil {
		return err
	}

	return nil
}

func writeImports(pathToName map[string]string, w io.Writer) {

	importPaths := make([]string, 0)
	for path := range pathToName {
		importPaths = append(importPaths, path)
	}
	sort.Sort(sort.StringSlice(importPaths))

	if len(importPaths) > 0 {
		_, _ = fmt.Fprintf(w, `import (`)

		for _, path := range importPaths {
			_, _ = fmt.Fprintf(w, `
	%s "%s"
`, pathToName[path], path)
		}

		_, _ = fmt.Fprintf(w, `)`)
	}
}
