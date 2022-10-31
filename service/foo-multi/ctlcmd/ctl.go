package main

import (
	"fmt"

	"github.com/gowins/dionysus"
	"github.com/gowins/dionysus/cmd"
	"github.com/spf13/cobra"
	"gorm.io/driver/mysql"
	"gorm.io/gen"
	"gorm.io/gorm"
)

type gormGen struct{}

var (
	gDSN         string
	outPath      string
	modelPkgPath string
)

func (g *gormGen) GetCmd() *cobra.Command {
	c := &cobra.Command{
		Use:   "gengorm",
		Short: "generate gorm code",
	}
	c.Flags().StringVar(&gDSN, "dsn", "", "connect to mysql dsn, user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local")
	c.Flags().StringVar(&outPath, "out", "", "generate repository outpath")
	c.Flags().StringVar(&modelPkgPath, "pkg", "", "generate model outpath")
	c.RunE = func(*cobra.Command, []string) error {
		if outPath == "" {
			outPath = "./common/repository"
		}
		if modelPkgPath == "" {
			modelPkgPath = "./common/model"
		}
		if gDSN == "" {
			return fmt.Errorf("dsn is empty")
		}
		genGormCode()
		return nil
	}
	return c
}
func (cc *gormGen) GetShutdownFunc() cmd.StopFunc {
	return func() {}
}
func (cc *gormGen) RegShutdownFunc(stopSteps ...cmd.StopStep) {
}

// go run service/foo-multi/ctlcmd/ctl.go gengorm
// go run ctl.go ctl
func main() {
	ctl := cmd.NewCtlCommand()
	ctl.RegRunFunc(func() error {
		fmt.Println("ctl test")
		return nil
	})
	cc := &gormGen{}
	dio := dionysus.NewDio()
	dio.DioStart("ctl_test", ctl, cc)
}

func genGormCode() {
	db, _ := gorm.Open(mysql.Open(gDSN), &gorm.Config{})
	// specify the output directory (default: "./query")
	// ### if you want to query without context constrain, set mode gen.WithoutContext ###
	g := gen.NewGenerator(gen.Config{
		OutPath:      outPath,
		ModelPkgPath: modelPkgPath,
		Mode:         gen.WithDefaultQuery,
		/* Mode: gen.WithoutContext|gen.WithDefaultQuery|gen.WithQueryInterface*/
		//if you want the nullable field generation property to be pointer type, set FieldNullable true
		/* FieldNullable: true,*/
		//if you want to assign field which has default value in `Create` API, set FieldCoverable true, reference: https://gorm.io/docs/create.html#Default-Values
		/* FieldCoverable: true,*/
		// if you want generate field with unsigned integer type, set FieldSignable true
		FieldSignable: true,
		//if you want to generate index tags from database, set FieldWithIndexTag true
		/* FieldWithIndexTag: true,*/
		//if you want to generate type tags from database, set FieldWithTypeTag true
		/* FieldWithTypeTag: true,*/
		//if you need unit tests for query code, set WithUnitTest true
		// WithUnitTest: true,
	})

	// reuse the database connection in Project or create a connection here
	// if you want to use GenerateModel/GenerateModelAs, UseDB is necessary or it will panic
	// db, _ := gorm.Open(mysql.Open("root:@(127.0.0.1:3306)/demo?charset=utf8mb4&parseTime=True&loc=Local"))
	g.UseDB(db)

	var dataMap = map[string]func(detailType string) (dataType string){
		//"int":     func(detailType string) (dataType string) { return "uint" },
		"tinyint": func(detailType string) (dataType string) { return "int8" },
	}
	g.WithDataTypeMap(dataMap)
	// g.WithJSONTagNameStrategy(func(c string) string { return "-" })
	g.FieldWithIndexTag = true
	g.WithNewTagNameStrategy(func(c string) string { return `mapstructure:"` + c + `"` })

	// apply basic crud api on structs or table models which is specified by table name with function
	// GenerateModel/GenerateModelAs. And generator will generate table models' code when calling Excute.
	// 想对已有的model生成crud等基础方法可以直接指定model struct ，例如model.User{}
	// 如果是想直接生成表的model和crud方法，则可以指定表的名称，例如g.GenerateModel("company")
	// 想自定义某个表生成特性，比如struct的名称/字段类型/tag等，可以指定opt，例如g.GenerateModel("company",gen.FieldIgnore("address")), g.GenerateModelAs("people", "Person", gen.FieldIgnore("address"))
	// g.ApplyBasic(model.User{}, g.GenerateModel("company"), g.GenerateModelAs("people", "Person", gen.FieldIgnore("address")))

	// apply diy interfaces on structs or table models
	// 如果想给某些表或者model生成自定义方法，可以用ApplyInterface，第一个参数是方法接口，可以参考DIY部分文档定义
	// g.ApplyInterface(func(method model.Method) {}, model.User{}, g.GenerateModel("company"))

	g.ApplyBasic(g.GenerateAllTable()...)
	// execute the action of code generation
	g.Execute()
}
