package pique

import (
	"fmt"
	"log"
	"reflect"

	"github.com/hashicorp/hcl/v2/hclsimple"
	"github.com/hashicorp/hcl2/gohcl"
	"github.com/hashicorp/hcl2/hclparse"
	job "github.com/veeruns/pique/internal/job"
)

//Program struct is the Program that will get interpretted to protcol buffer and saved in a file
type Program struct {
	LogLevel       string `hcl:"log_level"`
	Name           string `hcl:"name"`
	CheckFrequency int32  `hcl:"check_frequency"`

	Nodes []Node `hcl:"node,block"`
}

//Node Struct is each node tin the Program
type Node struct {
	ID                string `hcl:"id,label"`
	Blockname         string `hcl:"blockname"`
	Pluginname        string `hcl:"plugin_name"`
	Functionname      string `hcl:"plugin_function_name"`
	Functionarguments string `hcl:"plugin_function_arguments"`
}

//Parsehcl parses hcl if you need more than simple parsing. Not required now.
func Parsehcl() {
	var config Program
	err := hclsimple.DecodeFile("config.hcl", nil, &config)
	if err != nil {
		log.Fatalf("Failed to load configuration: %s", err)
	}
	//log.Printf("Configuration is %#v", config)

	parser := hclparse.NewParser()

	parsedhcl, parseDiags := parser.ParseHCLFile("config.hcl")
	if parseDiags.HasErrors() {
		fmt.Println(parseDiags.Error())
	}

	schema, _ := gohcl.ImpliedBodySchema(&Program{})
	content, _ := parsedhcl.Body.Content(schema)
	for k, v := range content.Blocks.ByType() {
		fmt.Printf("Nodes %#v : %#v\n", k, v)

		fmt.Printf("Type is %s\n", reflect.TypeOf(v))
		for _, w := range v {
			u := w.Body
			fmt.Printf("%#v\n", u)
			var c Program
			fmt.Printf("%#v\n", parsedhcl)
			decodeDiags := gohcl.DecodeBody(u, nil, &c)
			if decodeDiags.HasErrors() {
				fmt.Println(decodeDiags.Error())
			}
			fmt.Printf("%#v\n", c)
		}

	}
	//fmt.Printf("%#v", c)
}

//ParseAndLoad parses the config and loads it to the type
func ParseAndLoad() {
	var config Program
	localjob := new(job.Job)

	err := hclsimple.DecodeFile("../examples/program_examples/config.hcl", nil, &config)

	for key, value := range config.Nodes {
		log.Printf("%d and %#v\n", key, value)
	}
	//localjob.GlobalOptions.LogLevel
	goption := &job.GlobalOptions{
		LogLevel:       config.LogLevel,
		Name:           config.Name,
		Checkfrequency: config.CheckFrequency,
	}
	localjob.GlobalOptions = goption

	if err != nil {
		log.Fatalf("Failed to load configuration: %s", err)
	}
	log.Printf("Configuration is %s\n %s\n", config.LogLevel, localjob.GlobalOptions.GetLogLevel())
	startserver()
}
