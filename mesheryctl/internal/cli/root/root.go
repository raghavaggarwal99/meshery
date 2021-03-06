// Copyright 2020 Layer5, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package root

import (
	"errors"
	"fmt"
	"os"

	"github.com/layer5io/meshery/models"
	"gopkg.in/yaml.v2"

	"github.com/layer5io/meshery/mesheryctl/internal/cli/root/experimental"
	"github.com/layer5io/meshery/mesheryctl/internal/cli/root/mesh"
	"github.com/layer5io/meshery/mesheryctl/internal/cli/root/perf"
	"github.com/layer5io/meshery/mesheryctl/internal/cli/root/system"
	"github.com/layer5io/meshery/mesheryctl/pkg/utils"
	log "github.com/sirupsen/logrus"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

//TerminalFormatter is exported
type TerminalFormatter struct{}

var (
	cfgFile     string
	mctlCfgFile string
	version     = "Not Set"
	commitsha   = "Not Set"
)

//Format is exported
func (f *TerminalFormatter) Format(entry *log.Entry) ([]byte, error) {
	return append([]byte(entry.Message), '\n'), nil
}

var (
	availableSubcommands = []*cobra.Command{}
)

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "mesheryctl",
	Short: "Meshery Command Line tool",
	Long:  `Meshery is the service mesh management plane, providing lifecycle, performance, and configuration management of service meshes and their workloads.`,
	Args:  cobra.MinimumNArgs(1),
	// Uncomment the following line if your bare application
	// has an action associated with it:
	PreRunE: func(cmd *cobra.Command, args []string) error {
		if ok := utils.IsValidSubcommand(availableSubcommands, args[0]); !ok {
			return errors.New(utils.RootError(fmt.Sprintf("invalid command: \"%s\"", args[0])))
		}
		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		b, err := cmd.Flags().GetBool("version")
		if err != nil {
			return err
		}

		if b {
			versionCmd.Run(nil, nil)
			return nil
		}

		return nil
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the RootCmd.
func Execute() {
	log.SetLevel(log.InfoLevel)

	if debug, err := RootCmd.Flags().GetBool("debug"); err == nil && debug {
		log.SetLevel(log.DebugLevel)
	}

	//log formatter for improved UX
	log.SetFormatter(new(TerminalFormatter))
	_ = RootCmd.Execute()
}

func init() {
	err := utils.SetFileLocation()
	if err != nil {
		log.Fatal(err)
	}
	cobra.OnInitialize(initConfig)

	RootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", fmt.Sprintf("config file (default location is: %s/%s)", utils.MesheryFolder, "config.yaml"))

	// Preparing for an "edge" channel
	// RootCmd.PersistentFlags().StringVar(&cfgFile, "edge", "", "flag to run Meshery as edge (one-time)")

	RootCmd.Flags().BoolP("version", "v", false, "Version flag")
	RootCmd.Flags().BoolP("debug", "d", false, "Debug flag")

	availableSubcommands = []*cobra.Command{
		versionCmd,
		system.SystemCmd,
		perf.PerfCmd,
		mesh.MeshCmd,
		experimental.ExpCmd,
	}

	RootCmd.AddCommand(availableSubcommands...)
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	// Allow user to override config file with use of --config global flag
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
		// Otherwise, use the default `config.yaml` config file
	} else {
		if _, err := os.Stat(fmt.Sprintf("%s/%s", utils.MesheryFolder, "config.yaml")); os.IsNotExist(err) {
			log.Printf("Missing Meshery config file.")
			userResponse := utils.AskForConfirmation("Create default config now?")

			// Check for Meshery existence and permission of application folder
			if userResponse {
				if _, err := os.Stat(utils.MesheryFolder); err != nil {
					if os.IsNotExist(err) {
						err = os.MkdirAll(utils.MesheryFolder, 0775)
						if err != nil {
							log.Fatal(err)
						}
					}
				}

				// Create Meshery config file using standard template
				if _, err := os.Stat(fmt.Sprintf("%s/%s", utils.MesheryFolder, "config.yaml")); os.IsNotExist(err) {
					localContext := models.Context{
						Endpoint: "localhost:9081",
						Token: models.Token{
							Name:     "Default",
							Location: fmt.Sprintf("%s/%s", utils.MesheryFolder, "auth.json"),
						},
						Platform: "docker",
						Adapters: []string{"meshery-istio", "meshery-linkerd", "meshery-consul", "meshery-octarine", "meshery-nsm", "meshery-kuma", "meshery-cpx", "meshery-osm", "meshery-nginx-sm"},
					}
					ContextMap := map[string]models.Context{}
					ContextMap["local"] = localContext
					basicMap := models.MesheryCtlConfig{
						Contexts:       ContextMap,
						CurrentContext: "local",
						Tokens:         nil,
					}
					filePointer, err := os.Create(fmt.Sprintf("%s/%s", utils.MesheryFolder, "config.yaml"))
					if err != nil {
						log.Fatalln(err)
					}
					content, err := yaml.Marshal(basicMap)
					if err != nil {
						log.Fatalln(err)
					}
					_, err = filePointer.Write(content)
					if err != nil {
						log.Fatalln(err)
					}
				}

				log.Println(
					fmt.Sprintf("Default config file created at %s",
						fmt.Sprintf("%s/%s", utils.MesheryFolder, "config.yaml"),
					))
			} else {
				// User choose not to have a config file created. User must provide location to config file or create one.
				log.Printf("Provide config file location using `--config <config-file>` or" +
					" run `mesheryctl system context create <name>` to " +
					"generate a config file.")

				os.Exit(1)
			}
		}

		viper.SetConfigFile(fmt.Sprintf("%s/%s", utils.MesheryFolder, "config.yaml"))
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		log.Debug("Using config file:", viper.ConfigFileUsed())
	}
}
