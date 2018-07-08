package component

import (
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/factorapp/factor/files"
	"github.com/gobuffalo/envy"
)

// ProcessAll processes components starting at base
func ProcessAll(base string) error {
	packagePath := envy.CurrentPackage()
	err := filepath.Walk(base, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && files.IsHTML(info) {
			f, err := os.Open(path)
			if err != nil {
				return err
			}
			//c, _ := component.Parse(f, componentName(path))

			comp := files.ComponentName(path)
			gfn := filepath.Join(base, strings.ToLower(comp)+".go")
			_, err = os.Stat(gfn)
			var makeStruct bool
			if os.IsNotExist(err) {
				makeStruct = true
			}
			/*gofile, err := os.Create(goFileName(base, componentName(path)))
			if err != nil {
				return err
			}
			defer gofile.Close()

			c.Transform(gofile)
			*/
			transpiler, err := NewTranspiler(f, makeStruct, comp, "components", packagePath)
			if err != nil {
				log.Println("ERROR", err)
				return err
			}

			gofile, err := os.Create(files.GeneratedGoFileName(base, comp))
			if err != nil {
				log.Println("ERROR", err)
				return err
			}
			defer gofile.Close()
			_, err = io.WriteString(gofile, transpiler.Code())
			if err != nil {
				log.Println("ERROR", err)
				return err
			}
		}
		return nil
	})

	if err != nil {
		log.Printf("error walking the path %q: %v\n", base, err)
	}
	return err
}
