#!/bin/sh
# Copyright © 2020 Yashvardhan <yash.kukreja.98@gmail.com>

# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at

#     http://www.apache.org/licenses/LICENSE-2.0

# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

# script to run whatsupstream in the background
CONFIG_PATH="$1"
if [ -z $CONFIG_PATH ]; then
    echo "running whatsupstream with default config"
    whatsupstream notify &
else
    whatsupstream notify --config $CONFIG_PATH &
fi