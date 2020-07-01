package cmd

import (
	"fmt"
	"log"

	tfplugin "github.com/tetrafolium/tflint/plugin"
	"github.com/tetrafolium/tflint/tflint"
)

func (cli *CLI) printVersion(opts Options) int {
	fmt.Fprintf(cli.outStream, "TFLint version %s\n", tflint.Version)

	cfg, err := tflint.LoadConfig(opts.Config)
	if err != nil {
		log.Printf("[ERROR] Failed to load TFLint config: %s", err)
		return ExitCodeOK
	}
	cfg = cfg.Merge(opts.toConfig())

	plugin, err := tfplugin.Discovery(cfg)
	if err != nil {
		log.Printf("[ERROR] Failed to initialize plugins: %s", err)
		return ExitCodeOK
	}
	defer plugin.Clean()

	for _, ruleset := range plugin.RuleSets {
		name, err := ruleset.RuleSetName()
		if err != nil {
			log.Printf("[ERROR] Failed to get ruleset name: %s", err)
			continue
		}
		version, err := ruleset.RuleSetVersion()
		if err != nil {
			log.Printf("[ERROR] Failed to get ruleset version: %s", err)
			continue
		}

		fmt.Fprintf(cli.outStream, "+ ruleset.%s (%s)\n", name, version)
	}

	return ExitCodeOK
}
