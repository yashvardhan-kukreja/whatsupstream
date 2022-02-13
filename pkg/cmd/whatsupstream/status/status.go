/*
Copyright © 2020 Yashvardhan Kukreja <yash.kukreja.98@gmail.com>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package status

import (
	"fmt"
	"os/exec"
	"strings"

	"github.com/spf13/cobra"
)

func NewCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "status",
		Short: "Shows number of whatsupstream's processes active.",
		Long:  "Shows information about the active whatsupstream's processes along with their config files.",
		RunE: func(cmd *cobra.Command, args []string) error {
			return runE()
		},
	}
	return cmd
}

func runE() error {
	statusCmd := `ps aux | grep "whatsupstream notify" | grep -v "grep" | awk '{ print $14; }'`
	execCmd := exec.Command("bash", "-c", statusCmd)
	out, err := execCmd.Output()
	if err != nil {
		return fmt.Errorf("%w", err)
	}
	configs := strings.Split(string(out), "\n")
	fmt.Printf("No. of Active Processes: %d \n\n", len(configs)-1)
	if len(configs) > 1 {
		fmt.Println("Targetted Repositories & Issue Labels:")
		for _, config := range configs {
			if config == "" {
				continue
			}
			fmt.Println(config)
			issues, err := getReposFromConfigFile(config)
			if err != nil {
				return fmt.Errorf("%w", err)
			}
			for _, issue := range issues {
				fmt.Println("- ", issue)
			}
		}
	}
	return nil
}
