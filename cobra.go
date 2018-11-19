package envy

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

func ParseCobra(c *cobra.Command, p string, r bool) {
	// Append subcommand names to the prefix.
	if c.Root() != c {
		p = fmt.Sprintf("%s_%s", p, strings.ToUpper(c.Name()))
	}

	// Update the current command.
	updateCobra(p, c.Flags())

	// Recursively update child commands.
	if r {
		for _, child := range c.Commands() {
			if child.Name() == "help" {
				continue
			}

			ParseCobra(child, p, r)
		}
	}
}

func updateCobra(p string, fs *pflag.FlagSet) {
	// Build a map of explicitly set flags.
	set := map[string]interface{}{}
	fs.Visit(func(f *pflag.Flag) {
		set[f.Name] = nil
	})

	fs.VisitAll(func(f *pflag.Flag) {
		if f.Name == "help" {
			return
		}

		// Create an env var name
		// based on the supplied prefix.
		envVar := fmt.Sprintf("%s_%s", p, strings.ToUpper(f.Name))
		envVar = strings.Replace(envVar, "-", "_", -1)

		// Update the Flag.Value if the
		// env var is non "".
		if val := os.Getenv(envVar); val != "" {
			// Update the value if it hasn't
			// already been set.
			if _, defined := set[f.Name]; !defined {
				fs.Set(f.Name, val)
			}
		}

		// Append the env var to the
		// Flag.Usage field.
		f.Usage = fmt.Sprintf("%s [%s]", f.Usage, envVar)
	})
}
