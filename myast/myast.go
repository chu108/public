package myast

import (
	"fmt"
	"go/ast"
	"strings"
)

// GetObjFunMp find all exported func of sturct objName
// GetObjFunMp 类中的所有导出函数
func GetObjFunMp(astPkg *ast.Package, objName string) map[string]*ast.FuncDecl {
	funMp := make(map[string]*ast.FuncDecl)
	// find all exported func of sturct objName
	for _, fl := range astPkg.Files {
		for _, d := range fl.Decls {
			switch specDecl := d.(type) {
			case *ast.FuncDecl:
				if specDecl.Recv != nil {
					if exp, ok := specDecl.Recv.List[0].Type.(*ast.StarExpr); ok { // Check that the type is correct first beforing throwing to parser
						if strings.Compare(fmt.Sprint(exp.X), objName) == 0 { // is the same struct
							funMp[specDecl.Name.String()] = specDecl // catch
						}
					}
				}
			}
		}
	}

	return funMp
}

// AnalysisImport 分析整合import相关信息
func AnalysisImport(astPkgs *ast.Package) map[string]string {
	imports := make(map[string]string)
	for _, f := range astPkgs.Files {
		for _, p := range f.Imports {
			k := ""
			if p.Name != nil {
				k = p.Name.Name
			}
			v := strings.Trim(p.Path.Value, `"`)
			if len(k) == 0 {
				n := strings.LastIndex(v, "/")
				if n > 0 {
					k = v[n+1:]
				} else {
					k = v
				}
			}
			imports[k] = v
		}
	}

	return imports
}

// GetImportString 获取包名 import
func GetImportString(imports map[string]*ast.ImportSpec, pkgName string) string {
	for _, v := range imports {
		fmt.Println(v)
	}
	return ""
}
