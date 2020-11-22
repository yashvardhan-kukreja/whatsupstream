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
package notify

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	homedir "github.com/mitchellh/go-homedir"
)

type flagpole struct {
	Config string
}

func NewCommand() *cobra.Command {
	flags := &flagpole{}
	cmd := &cobra.Command{
		Use:   "notify",
		Short: "Enables notifications",
		Long: `Enables desktop notifications for the current user as per the provided configuration.
The user would receive notifications around new issues depending on the filters/conditions provided in the config.
The notifications would be in the form of desktop notifications.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			return runE(flags)
		},
	}
	home, err := homedir.Dir()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	cmd.Flags().StringVar(
		&flags.Config, "config",
		fmt.Sprintf("%s/.whatsupstream/config.yaml", home),
		"Path to the config containing preferences associated with the notifications to receive.",
	)
	return cmd
}