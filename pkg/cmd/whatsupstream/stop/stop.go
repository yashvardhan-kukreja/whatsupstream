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
package stop

import (
	"fmt"
	"github.com/spf13/cobra"
	"os/exec"
)

func NewCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "stop",
		Short: "Stops whatsupstream's notifications",
		Long:  "Stops all instances of whatsupstream's notification service",
		RunE: func(cmd *cobra.Command, args []string) error {
			return runE()
		},
	}
	return cmd
}

func runE() error {
	stopCmd := `ps aux | grep "whatsupstream notify" | grep -v "grep" | awk '{ print $2; }' | xargs kill `
	err := exec.Command("bash", "-c", stopCmd).Run()
	if err != nil {
		return fmt.Errorf("error occurred while stopping the instances of whatsupstream: %w", err)
	}
	return nil
}
