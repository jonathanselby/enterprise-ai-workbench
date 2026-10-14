package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/jonathanselby/enterprise-ai-workbench/internal/workbench"
)

const version = "0.1.0"

func main() { os.Exit(run(os.Args[1:])) }

func run(args []string) int {
	if len(args) == 0 {
		usage()
		return 2
	}
	switch args[0] {
	case "version", "-version", "--version":
		fmt.Println(version)
		return 0
	case "list":
		return listCmd(args[1:])
	case "validate":
		return validateCmd(args[1:])
	case "render":
		return renderCmd(args[1:])
	case "check":
		return checkCmd(args[1:])
	case "review":
		return reviewCmd(args[1:])
	default:
		fmt.Fprintf(os.Stderr, "unknown command %q\n", args[0])
		usage()
		return 2
	}
}

func listCmd(args []string) int {
	fs := flag.NewFlagSet("list", flag.ContinueOnError)
	fs.SetOutput(os.Stderr)
	root := fs.String("root", "workflows", "workflow root directory")
	if fs.Parse(args) != nil {
		return 2
	}
	entries, err := workbench.List(*root)
	if err != nil {
		return executionError(err)
	}
	return printJSON(entries)
}

func validateCmd(args []string) int {
	fs := flag.NewFlagSet("validate", flag.ContinueOnError)
	fs.SetOutput(os.Stderr)
	root := fs.String("root", "workflows", "workflow root directory")
	if fs.Parse(args) != nil {
		return 2
	}
	result := workbench.ValidateCatalog(*root)
	printJSON(result)
	if !result.Valid {
		return 1
	}
	return 0
}

func renderCmd(args []string) int {
	fs := flag.NewFlagSet("render", flag.ContinueOnError)
	fs.SetOutput(os.Stderr)
	root := fs.String("root", "workflows", "workflow root directory")
	workflowID := fs.String("workflow", "", "workflow id")
	inputPath := fs.String("input", "", "input Markdown/text file")
	outPath := fs.String("out", "", "optional output path; stdout when omitted")
	if fs.Parse(args) != nil {
		return 2
	}
	if *workflowID == "" || *inputPath == "" {
		fmt.Fprintln(os.Stderr, "-workflow and -input are required")
		return 2
	}
	wf, err := workbench.LoadWorkflow(*root, *workflowID)
	if err != nil {
		return executionError(err)
	}
	input, err := workbench.ReadText(*inputPath)
	if err != nil {
		return executionError(err)
	}
	bundle, err := workbench.Render(wf, input)
	if err != nil {
		return executionError(err)
	}
	if *outPath == "" {
		fmt.Print(bundle)
		return 0
	}
	if err := os.WriteFile(*outPath, []byte(bundle), 0644); err != nil {
		return executionError(err)
	}
	return 0
}

func checkCmd(args []string) int {
	fs := flag.NewFlagSet("check", flag.ContinueOnError)
	fs.SetOutput(os.Stderr)
	root := fs.String("root", "workflows", "workflow root directory")
	workflowID := fs.String("workflow", "", "workflow id")
	outputPath := fs.String("output", "", "model output Markdown file")
	if fs.Parse(args) != nil {
		return 2
	}
	if *workflowID == "" || *outputPath == "" {
		fmt.Fprintln(os.Stderr, "-workflow and -output are required")
		return 2
	}
	wf, err := workbench.LoadWorkflow(*root, *workflowID)
	if err != nil {
		return executionError(err)
	}
	output, err := workbench.ReadText(*outputPath)
	if err != nil {
		return executionError(err)
	}
	result := workbench.CheckOutput(wf.Manifest, output)
	printJSON(result)
	if !result.Valid {
		return 1
	}
	return 0
}

func reviewCmd(args []string) int {
	fs := flag.NewFlagSet("review", flag.ContinueOnError)
	fs.SetOutput(os.Stderr)
	root := fs.String("root", "workflows", "workflow root directory")
	workflowID := fs.String("workflow", "", "workflow id")
	inputPath := fs.String("input", "", "original input file")
	outputPath := fs.String("output", "", "model output file")
	decision := fs.String("decision", "", "accepted or rejected")
	reviewer := fs.String("reviewer", "", "human reviewer or deterministic control name")
	reason := fs.String("reason", "", "review reason")
	if fs.Parse(args) != nil {
		return 2
	}
	if *workflowID == "" || *inputPath == "" || *outputPath == "" || *decision == "" || *reviewer == "" || *reason == "" {
		fmt.Fprintln(os.Stderr, "-workflow, -input, -output, -decision, -reviewer and -reason are required")
		return 2
	}
	wf, err := workbench.LoadWorkflow(*root, *workflowID)
	if err != nil {
		return executionError(err)
	}
	input, err := workbench.ReadText(*inputPath)
	if err != nil {
		return executionError(err)
	}
	output, err := workbench.ReadText(*outputPath)
	if err != nil {
		return executionError(err)
	}
	record, err := workbench.Review(wf, input, output, *decision, *reviewer, *reason, time.Now())
	if err != nil {
		if strings.Contains(err.Error(), "cannot accept output") {
			fmt.Fprintln(os.Stderr, err)
			return 1
		}
		return executionError(err)
	}
	return printJSON(record)
}

func printJSON(v any) int {
	enc := json.NewEncoder(os.Stdout)
	enc.SetIndent("", "  ")
	if err := enc.Encode(v); err != nil {
		fmt.Fprintln(os.Stderr, err)
		return 3
	}
	return 0
}

func executionError(err error) int { fmt.Fprintln(os.Stderr, err); return 3 }

func usage() {
	fmt.Fprintln(os.Stderr, "usage: workbench <version|list|validate|render|check|review> [options]")
}
