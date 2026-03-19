package main

import (
	"flag"
	"fmt"

	"github.com/zhangga/luban-go/core"
	"github.com/zhangga/luban-go/core/pipeline"
)

func main() {
	var target string
	var codeTarget string
	var inputDataDir string
	var outputCodeDir string
	var outputDataDir string
	var templateDir string

	// CLI参数解析
	flag.StringVar(&target, "target", "all", "the target name defined in xml")
	flag.StringVar(&codeTarget, "code_target", "go", "target language for code generation (e.g. go, cs)")
	flag.StringVar(&inputDataDir, "input_data_dir", ".", "input data directory")
	flag.StringVar(&outputCodeDir, "output_code_dir", "./gen_code", "output code directory")
	flag.StringVar(&outputDataDir, "output_data_dir", "./gen_data", "output data directory")
	flag.StringVar(&templateDir, "template_dir", "./templates", "template directory for code generation")

	flag.Parse()

	args := &pipeline.PipelineArguments{
		Target:        target,
		CodeTargets:   []string{codeTarget},
		InputDataDir:  inputDataDir,
		OutputCodeDir: outputCodeDir,
		OutputDataDir: outputDataDir,
		TemplateDir:   templateDir,
	}

	launcher := core.NewSimpleLauncher()

	fmt.Println("Luban Go started")
	launcher.Start(args)
	fmt.Println("Luban Go finished")
}
