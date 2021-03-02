package project

import (
	"fmt"
	"os"
	"path"
	"text/template"

	"github.com/alandtsang/shine/internal/license"
	"github.com/alandtsang/shine/internal/tpl"
)

type Project struct {
	AbsolutePath string
	PkgName      string
	Author       string
	Copyright    string
	License      *license.License
	AppName      string
}

func InitProject(args []string) (string, error) {
	curDir, err := os.Getwd()
	if err != nil {
		return "", err
	}

	if len(args) > 0 {
		if args[0] != "." {
			curDir = fmt.Sprintf("%s/%s", curDir, args[0])
		}
	}

	pkgName := args[0]
	author := args[1]
	copyright := license.CreateCopyright(author)
	licenseText := license.New(license.Apache2License)
	appName := path.Base(pkgName)

	proj := New(curDir, pkgName, author, copyright, appName, licenseText)
	if err = proj.Create(); err != nil {
		return "", err
	}

	return proj.AbsolutePath, nil
}

func New(path, pkg, author, copyright, appName string, licenseText *license.License) *Project {
	return &Project{
		AbsolutePath: path,
		PkgName:      pkg,
		Author:       author,
		Copyright:    copyright,
		License:      licenseText,
		AppName:      appName,
	}
}

func (p *Project) Create() error {
	var err error

	// create root directory
	if err = p.createRootDir(); err != nil {
		return err
	}

	// create cmd directory
	if err = p.createCmdDir(); err != nil {
		return err
	}

	// create configs directory
	if err = p.createConfigsDir(); err != nil {
		return err
	}

	// create internal directory
	if err = p.createInternalDir(); err != nil {
		return err
	}

	// create go.mod and go.sum file
	if err = p.createGoModAndSum(); err != nil {
		return err
	}

	// create Makefile file
	if err = p.createMakefile(); err != nil {
		return err
	}

	// create license file
	return license.CreateLicenseFile(p.AbsolutePath, p.Author)
}

func (p *Project) createRootDir() error {
	if _, err := os.Stat(p.AbsolutePath); os.IsNotExist(err) {
		if err = os.Mkdir(p.AbsolutePath, 0755); err != nil {
			return err
		}
	}
	return nil
}

func (p *Project) createCmdDir() error {
	var err error

	cmdDirStr := fmt.Sprintf("%s/cmd", p.AbsolutePath)
	mainFileStr := fmt.Sprintf("%s/cmd/main.go", p.AbsolutePath)

	if _, err = os.Stat(cmdDirStr); os.IsNotExist(err) {
		if err = os.Mkdir(cmdDirStr, 0755); err != nil {
			return err
		}
	}

	mainFile, err := os.Create(mainFileStr)
	if err != nil {
		return err
	}

	mainTemplate := template.Must(template.New("main").Parse(string(tpl.MainTemplate())))
	if err = mainTemplate.Execute(mainFile, p); err != nil {
		return err
	}

	return license.CreateLicenseFile(p.AbsolutePath, p.Author)
}

func (p *Project) createInternalDir() error {
	var err error

	internalDirStr := fmt.Sprintf("%s/internal", p.AbsolutePath)
	if _, err = os.Stat(internalDirStr); os.IsNotExist(err) {
		if err = os.Mkdir(internalDirStr, 0755); err != nil {
			return err
		}
	}

	if err = p.createInternalBinderDir(); err != nil {
		return err
	}

	if err = p.createInternalValidatorDir(); err != nil {
		return err
	}

	if err = p.createInternalRouterDir(); err != nil {
		return err
	}

	return nil
}

func (p *Project) createConfigsDir() error {
	var err error

	configsDirStr := fmt.Sprintf("%s/configs", p.AbsolutePath)
	configsFileStr := fmt.Sprintf("%s/configs/configs.go", p.AbsolutePath)

	if _, err = os.Stat(configsDirStr); os.IsNotExist(err) {
		if err = os.Mkdir(configsDirStr, 0755); err != nil {
			return err
		}
	}

	configsFile, err := os.Create(configsFileStr)
	if err != nil {
		return err
	}

	configsTemplate := template.Must(template.New("configs").Parse(string(tpl.ConfigsTemplate())))
	return configsTemplate.Execute(configsFile, p)
}

func (p *Project) createInternalBinderDir() error {
	var err error

	binderDirStr := fmt.Sprintf("%s/internal/binder", p.AbsolutePath)
	binderFileStr := fmt.Sprintf("%s/internal/binder/binder.go", p.AbsolutePath)

	if _, err = os.Stat(binderDirStr); os.IsNotExist(err) {
		if err = os.Mkdir(binderDirStr, 0755); err != nil {
			return err
		}
	}

	binderFile, err := os.Create(binderFileStr)
	if err != nil {
		return err
	}

	binderTemplate := template.Must(template.New("binder").Parse(string(tpl.BinderTemplate())))
	return binderTemplate.Execute(binderFile, p)
}

func (p *Project) createInternalValidatorDir() error {
	var err error

	validatorDirStr := fmt.Sprintf("%s/internal/validator", p.AbsolutePath)
	validatorFileStr := fmt.Sprintf("%s/internal/validator/validator.go", p.AbsolutePath)

	if _, err = os.Stat(validatorDirStr); os.IsNotExist(err) {
		if err = os.Mkdir(validatorDirStr, 0755); err != nil {
			return err
		}
	}

	validatorFile, err := os.Create(validatorFileStr)
	if err != nil {
		return err
	}

	binderTemplate := template.Must(template.New("validator").Parse(string(tpl.ValidatorTemplate())))
	return binderTemplate.Execute(validatorFile, p)
}

func (p *Project) createInternalRouterDir() error {
	var err error

	routerDirStr := fmt.Sprintf("%s/internal/router", p.AbsolutePath)
	routerFileStr := fmt.Sprintf("%s/internal/router/router.go", p.AbsolutePath)

	if _, err = os.Stat(routerDirStr); os.IsNotExist(err) {
		if err = os.Mkdir(routerDirStr, 0755); err != nil {
			return err
		}
	}

	if err = p.createInternalRoutesDir(); err != nil {
		return err
	}

	routerFile, err := os.Create(routerFileStr)
	if err != nil {
		return err
	}

	routerTemplate := template.Must(template.New("router").Parse(string(tpl.RouterTemplate())))
	return routerTemplate.Execute(routerFile, p)
}

func (p *Project) createInternalRoutesDir() error {
	var err error

	routesDirStr := fmt.Sprintf("%s/internal/router/routes", p.AbsolutePath)
	routesFileStr := fmt.Sprintf("%s/internal/router/routes/routes.go", p.AbsolutePath)

	if _, err = os.Stat(routesDirStr); os.IsNotExist(err) {
		if err = os.Mkdir(routesDirStr, 0755); err != nil {
			return err
		}
	}

	routersFile, err := os.Create(routesFileStr)
	if err != nil {
		return err
	}

	routesTemplate := template.Must(template.New("routes").Parse(string(tpl.RoutesTemplate())))
	return routesTemplate.Execute(routersFile, p)
}

func (p *Project) createGoModAndSum() error {
	if err := p.createGoMod(); err != nil {
		return err
	}
	return p.createGoSum()
}

func (p *Project) createGoMod() error {
	var err error

	goModFileStr := fmt.Sprintf("%s/go.mod", p.AbsolutePath)
	goModFile, err := os.Create(goModFileStr)
	if err != nil {
		return err
	}

	goModTemplate := template.Must(template.New("gomod").Parse(string(tpl.GoModTemplate())))
	return goModTemplate.Execute(goModFile, p)
}

func (p *Project) createGoSum() error {
	var err error

	goSumFileStr := fmt.Sprintf("%s/go.sum", p.AbsolutePath)

	goSumFile, err := os.Create(goSumFileStr)
	if err != nil {
		return err
	}

	goSumTemplate := template.Must(template.New("gosum").Parse(string(tpl.GoSumTemplate())))
	return goSumTemplate.Execute(goSumFile, p)
}

func (p *Project) createMakefile() error {
	var err error

	makefileStr := fmt.Sprintf("%s/Makefile", p.AbsolutePath)

	makefile, err := os.Create(makefileStr)
	if err != nil {
		return err
	}

	makefileTemplate := template.Must(template.New("makefile").Parse(string(tpl.MakefileTemplate())))
	return makefileTemplate.Execute(makefile, p)
}
