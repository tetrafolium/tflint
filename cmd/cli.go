package cmd

import (
	"fmt"
	"io"
	"os"

	flags "github.com/jessevdk/go-flags"

	"github.com/wata727/tflint/issue"
	"github.com/wata727/tflint/printer"
	"github.com/wata727/tflint/rules"
	"github.com/wata727/tflint/tflint"
)

// Exit codes are int values that represent an exit code for a particular error.
const (
	ExitCodeOK    int = 0
	ExitCodeError int = 1 + iota
	ExitCodeIssuesFound
)

// CLI is the command line object
type CLI struct {
	// outStream and errStream are the stdout and stderr
	// to write message from the CLI.
	outStream, errStream io.Writer
	loader               tflint.AbstractLoader
	testMode             bool
}

// NewCLI returns new CLI initialized by input streams
func NewCLI(outStream io.Writer, errStream io.Writer) *CLI {
	return &CLI{
		outStream: outStream,
		errStream: errStream,
	}
}

// Run invokes the CLI with the given arguments.
func (cli *CLI) Run(args []string) int {
	var opts Options
	parser := flags.NewParser(&opts, flags.HelpFlag)
	parser.Usage = "[OPTIONS]"
	parser.UnknownOptionHandler = func(option string, arg flags.SplitArgument, args []string) ([]string, error) {
		return []string{}, fmt.Errorf("ERROR: `%s` is unknown option. Please run `tflint --help`", option)
	}
	// Parse commandline flag
	args, err := parser.ParseArgs(args)
	if err != nil {
		if flagsErr, ok := err.(*flags.Error); ok && flagsErr.Type == flags.ErrHelp {
			fmt.Fprintln(cli.outStream, err)
			return ExitCodeOK
		}
		fmt.Fprintln(cli.errStream, err)
		return ExitCodeError
	}
	argFiles := args[1:]

	// Show version
	if opts.Version {
		fmt.Fprintf(cli.outStream, "TFLint version %s\n", Version)
		return ExitCodeOK
	}

	// Setup config
	cfg, err := tflint.LoadConfig(opts.Config)
	if err != nil {
		fmt.Fprintln(cli.errStream, err)
		return ExitCodeError
	}
	cfg = cfg.Merge(opts.toConfig())

	// Load Terraform's configurations
	if !cli.testMode {
		cli.loader, err = tflint.NewLoader()
		if err != nil {
			fmt.Fprintln(cli.errStream, err)
			return ExitCodeError
		}
	}
	for _, file := range argFiles {
		if fileInfo, err := os.Stat(file); os.IsNotExist(err) {
			fmt.Fprintf(cli.errStream, "%s: configuration file is not found\n", file)
			return ExitCodeError
		} else if fileInfo.IsDir() {
			fmt.Fprintf(cli.errStream, "%s: TFLint doesn't accept directories as arguments\n", file)
			return ExitCodeError
		}

		if !cli.loader.IsConfigFile(file) {
			fmt.Fprintf(cli.errStream, "%s: This file is not a configuration file\n", file)
			return ExitCodeError
		}
	}
	configs, err := cli.loader.LoadConfig()
	if err != nil {
		fmt.Fprintln(cli.errStream, err)
		return ExitCodeError
	}
	valuesFiles, err := cli.loader.LoadValuesFiles(cfg.Varfile...)
	if err != nil {
		fmt.Fprintln(cli.errStream, err)
		return ExitCodeError
	}

	// Check configurations via Runner
	runner := tflint.NewRunner(cfg, configs, valuesFiles...)
	runners, err := tflint.NewModuleRunners(runner)
	if err != nil {
		fmt.Fprintln(cli.errStream, err)
		return ExitCodeError
	}
	runners = append(runners, runner)

	for _, rule := range rules.NewRules(cfg) {
		for _, runner := range runners {
			err := rule.Check(runner)
			if err != nil {
				fmt.Fprintln(cli.errStream, err)
				return ExitCodeError
			}
		}
	}

	issues := []*issue.Issue{}
	for _, runner := range runners {
		issues = append(issues, runner.LookupIssues(argFiles...)...)
	}

	// Print issues
	printer.NewPrinter(cli.outStream, cli.errStream).Print(issues, opts.Format, opts.Quiet)

	if opts.ErrorWithIssues && len(issues) > 0 {
		return ExitCodeIssuesFound
	}

	return ExitCodeOK
}
