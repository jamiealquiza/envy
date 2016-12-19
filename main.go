package envy

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

// Parse takes a string p that is used
// as the environment variable prefix
// for each flag configured.
func Parse(p string) {
	flag.VisitAll(func(f *flag.Flag) {
		// Create an env var name
		// based on the supplied prefix.
		envVar := fmt.Sprintf("%s_%s", p, strings.ToUpper(f.Name))
		// Replace hyphens with underscores.
		envVar = strings.Replace(envVar, "-", "_", -1)

		// Traverse the available flags
		// and look for env overrides.
		if val := os.Getenv(envVar); val != "" {
			flag.Set(f.Name, val)
		}
	})
}
