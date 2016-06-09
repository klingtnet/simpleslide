package cmd

import (
	"bufio"
	"bytes"
	"io/ioutil"
	"os"
	"strings"
	"text/template"

	"github.com/russross/blackfriday"
	"github.com/spf13/cobra"

	"github.com/klingtnet/simpleslide/lib"
)

const SEP = "<hr />"

var verbose bool

var buildCmd = &cobra.Command{
	Use:   "build [input] [output]",
	Short: "Builds the slideshow",
	Run: func(cmd *cobra.Command, args []string) {
		utils.ExpectExactArgs(cmd, 2, args)
		data, err := ioutil.ReadFile(args[0])
		utils.ExitIfErr(err)
		rendered := blackfriday.MarkdownCommon(data)
		slides := strings.Split(string(rendered), SEP)
		utils.PrintIfVerbose(verbose, "Found %d slides\n", len(slides))
		raw, err := renderTemplate(slides)
		utils.ExitIfErr(err)
		utils.ExitIfErr(ioutil.WriteFile(args[1], raw, os.FileMode(0644)))
	},
}

func init() {
	buildCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "print what is being done")
}

type Content struct {
	Style  string
	Script string
	Slides []string
}

func renderTemplate(slides []string) ([]byte, error) {
	htmlTmpl, err := utils.ReadFileAsString("assets/template.html")
	if err != nil {
		return nil, err
	}
	style, err := utils.ReadFileAsString("assets/style.css")
	if err != nil {
		return nil, err
	}
	script, err := utils.ReadFileAsString("assets/slides.js")
	if err != nil {
		return nil, err
	}
	content := Content{
		Style:  style,
		Script: script,
		Slides: slides,
	}

	tmpl, err := template.New("slideshow").Parse(htmlTmpl)
	if err != nil {
		return nil, err
	}
	var renderBuf bytes.Buffer
	byteWriter := bufio.NewWriter(&renderBuf)
	err = tmpl.Execute(byteWriter, content)
	return renderBuf.Bytes(), err
}
