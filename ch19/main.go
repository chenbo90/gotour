package main

import "os"

/**
ch19 代码规范检查

localhost:gotour chenbo$ go get github.com/golangci/golangci-lint/cmd/golangci-lint@v1.32.2		//下载golangci-lint
go: downloading github.com/golangci/golangci-lint v1.32.2
go: downloading github.com/gofrs/flock v0.8.0
go: downloading github.com/spf13/viper v1.7.1
go: downloading golang.org/x/tools v0.0.0-20201013201025-64a9e34f3752
go: downloading github.com/go-critic/go-critic v0.5.2
go: downloading github.com/go-xmlfmt/xmlfmt v0.0.0-20191208150333-d5b6f63a941b
go: downloading github.com/golangci/revgrep v0.0.0-20180526074752-d9c87f5ffaf0
go: downloading 4d63.com/gochecknoglobals v0.0.0-20201008074935-acfc0b28355a
go: downloading github.com/Djarvur/go-err113 v0.0.0-20200511133814-5174e21577d5
go: downloading github.com/OpenPeeDeeP/depguard v1.0.1
go: downloading github.com/bombsimon/wsl/v3 v3.1.0
go: downloading github.com/daixiang0/gci v0.2.4
go: downloading github.com/denis-tingajkin/go-header v0.3.1
go: downloading github.com/golangci/check v0.0.0-20180506172741-cfe4005ccda2
go: downloading github.com/golangci/errcheck v0.0.0-20181223084120-ef45e06d44b6
go: downloading github.com/golangci/dupl v0.0.0-20180902072040-3e9179ac440a
go: downloading github.com/golangci/go-misc v0.0.0-20180628070357-927a3d87b613
go: downloading github.com/golangci/goconst v0.0.0-20180610141641-041c5f2b40f3
go: downloading github.com/golangci/gocyclo v0.0.0-20180528144436-0a533e8fa43d
go: downloading github.com/golangci/gofmt v0.0.0-20190930125516-244bba706f1a
go: downloading github.com/golangci/ineffassign v0.0.0-20190609212857-42439a7714cc
go: downloading github.com/golangci/lint-1 v0.0.0-20191013205115-297bf364a8e0
go: downloading github.com/golangci/maligned v0.0.0-20180506175553-b1d89398deca
go: downloading github.com/golangci/misspell v0.0.0-20180809174111-950f5d19e770
go: downloading github.com/golangci/prealloc v0.0.0-20180630174525-215b22d4de21
go: downloading github.com/golangci/unconvert v0.0.0-20180507085042-28b1c447d1f4
go: downloading github.com/jingyugao/rowserrcheck v0.0.0-20191204022205-72ab7603b68a
go: downloading github.com/jirfag/go-printf-func-name v0.0.0-20191110105641-45db9963cdd3
go: downloading github.com/kyoh86/exportloopref v0.1.7
go: downloading github.com/maratori/testpackage v1.0.1
go: downloading github.com/matoous/godox v0.0.0-20190911065817-5d6d842e92eb
go: downloading github.com/mbilski/exhaustivestruct v1.1.0
go: downloading github.com/moricho/tparallel v0.2.1
go: downloading github.com/nakabonne/nestif v0.3.0
go: downloading github.com/nishanths/exhaustive v0.1.0
go: downloading github.com/polyfloyd/go-errorlint v0.0.0-20201006195004-351e25ade6e3
go: downloading github.com/ryancurrah/gomodguard v1.1.0
go: downloading github.com/ryanrolds/sqlclosecheck v0.3.0
go: downloading github.com/securego/gosec/v2 v2.5.0
go: downloading github.com/shazow/go-diff v0.0.0-20160112020656-b6b7b6733b8c
go: downloading github.com/sonatard/noctx v0.0.1
go: downloading github.com/sourcegraph/go-diff v0.6.1
go: downloading github.com/ssgreg/nlreturn/v2 v2.1.0
go: downloading github.com/tdakkota/asciicheck v0.0.0-20200416190851-d7f85be797a2
go: downloading github.com/tetafro/godot v0.4.9
go: downloading github.com/timakin/bodyclose v0.0.0-20190930140734-f7f2e9bca95e
go: downloading github.com/tomarrell/wrapcheck v0.0.0-20200807122107-df9e8bcb914d
go: downloading github.com/tommy-muehle/go-mnd v1.3.1-0.20200224220436-e6f9a994e8fa
go: downloading github.com/ultraware/funlen v0.0.3
go: downloading github.com/ultraware/whitespace v0.0.4
go: downloading github.com/uudashr/gocognit v1.0.1
go: downloading honnef.co/go/tools v0.0.1-2020.1.6
go: downloading mvdan.cc/gofumpt v0.0.0-20200802201014-ab5a8192947d
go: downloading mvdan.cc/interfacer v0.0.0-20180901003855-c20040233aed
go: downloading mvdan.cc/unparam v0.0.0-20200501210554-b37ab49443f7
go: downloading github.com/stretchr/objx v0.1.1
go: downloading github.com/spf13/afero v1.1.2
go: downloading github.com/spf13/cast v1.3.0
go: downloading github.com/subosito/gotenv v1.2.0
go: downloading gopkg.in/ini.v1 v1.51.0
go: downloading github.com/go-toolsmith/astfmt v1.0.0
go: downloading github.com/go-toolsmith/astcast v1.0.0
go: downloading github.com/go-toolsmith/astcopy v1.0.0
go: downloading github.com/go-toolsmith/astequal v1.0.0
go: downloading github.com/go-toolsmith/astp v1.0.0
go: downloading github.com/go-toolsmith/strparse v1.0.0
go: downloading github.com/go-toolsmith/typep v1.0.2
go: downloading github.com/quasilyte/go-ruleguard v0.2.0
go: downloading github.com/quasilyte/regex/syntax v0.0.0-20200407221936-30656e2c4a95
go: downloading github.com/gostaticanalysis/analysisutil v0.1.0
go: downloading github.com/Masterminds/semver v1.5.0
go: downloading github.com/phayes/checkstyle v0.0.0-20170904204023-bfd46e6a821d
go: downloading github.com/nbutton23/zxcvbn-go v0.0.0-20180912185939-ae427f1e4c1d
go: downloading mvdan.cc/lint v0.0.0-20170908181259-adc824a0674b
go: downloading github.com/gostaticanalysis/comment v1.3.0
go get: added github.com/go-critic/checkers v0.0.0-20181031185637-879460b6c936
go get: added github.com/go-lintpack/lintpack v0.0.0-20181105152233-7ff0297828fc
go get: added github.com/golangci/go-tools v0.0.0-20180902103155-93eecd106a0b
go get: added github.com/golangci/golangci-lint v1.32.2
go get: added github.com/golangci/gosec v0.0.0-20180901114220-8afd9cbb6cfb
go get: added github.com/golangci/govet v0.0.0-20180818181408-44ddbe260190
go get: added github.com/golangci/interfacer v0.0.0-20180902080945-01958817a6ec
go get: added github.com/golangci/lint v0.0.0-20180902080404-c2187e7932b5
go get: added github.com/golangci/tools v0.0.0-20180902102414-2cefd77fef9b
go get: added github.com/golangci/unparam v0.0.0-20180902112548-7ad9dbcccc16
go get: added gopkg.in/airbrake/gobrake.v2 v2.0.9
go get: added gopkg.in/gemnasium/logrus-airbrake-hook.v2 v2.1.2
go get: added sourcegraph.com/sourcegraph/go-diff v0.0.0-20171119081133-3f415a150aec
go get: added sourcegraph.com/sqs/pbtypes v0.0.0-20160107090929-4d1b9dc7ffc3
localhost:gotour chenbo$ golangci-lint version
golangci-lint has version v1.32.2 built from (unknown, mod sum: "h1:CgIeFWTLJ3Nt1w/WU1RO351j/CjN6LIVjppbJfI9nMk=") on (unknown)
localhost:gotour chenbo$ golangci-lint run ch19/				//	执行代码检查
ch19/main.go:8:7: `name` is unused (deadcode)
const name  = "树下听雨"
      ^
ch19/main.go:10:10: Error return value of `os.Mkdir` is not checked (errcheck)
        os.Mkdir("tmp",0666)//创建一个临时目录
                ^
localhost:gotour chenbo$



*/
const name  = "树下听雨"
func main(){
	os.Mkdir("tmp",0666)//创建一个临时目录
	//os.Remove("tmp1")
}
