package cmd

import (
	"embed"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

//go:embed tpl
var f embed.FS

var dumpCmd = &cobra.Command{
	Use:   "dump",
	Short: "Output docker-compose.yml template",
}

func buildSetupCmd(s string) *cobra.Command {
	return &cobra.Command{
		Use:   s,
		Short: "",
		Run: func(cmd *cobra.Command, args []string) {
			d, err := f.ReadFile("tpl/" + s + ".yml")
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}

			fmt.Print(string(d))
		},
	}
}

func init() {
	rootCmd.AddCommand(dumpCmd)
	dumpCmd.AddCommand(buildSetupCmd("clojure"))
	dumpCmd.AddCommand(buildSetupCmd("django"))
	dumpCmd.AddCommand(buildSetupCmd("dotnet"))
	dumpCmd.AddCommand(buildSetupCmd("golang"))
	dumpCmd.AddCommand(buildSetupCmd("mongo-express"))
	dumpCmd.AddCommand(buildSetupCmd("node"))
	dumpCmd.AddCommand(buildSetupCmd("postgres"))
	dumpCmd.AddCommand(buildSetupCmd("python"))
	dumpCmd.AddCommand(buildSetupCmd("rails-webpacker"))
	dumpCmd.AddCommand(buildSetupCmd("rails"))
	dumpCmd.AddCommand(buildSetupCmd("ruby"))
}
