package pique

import (
	"fmt"
	"log"
	"reflect"

	"github.com/hashicorp/hcl/v2/hclsimple"
	"github.com/hashicorp/hcl2/gohcl"
	"github.com/hashicorp/hcl2/hclparse"
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
	Id                string `hcl:"id,label"`
	BlockName         string `hcl:"block_name`
	Pluginname        string `hcl:"plugin_name"`
	Functionname      string `hcl:"plugin_function_name"`
	Functionarguments string `hcl:"plugin_function_arguments"`
}

func parsehcl() {
	var config Program
	err := hclsimple.DecodeFile("config.hcl", nil, &config)
	if err != nil {
		log.Fatalf("Failed to load configuration: %s", err)
	}
	log.Printf("Configuration is %#v", config)

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
