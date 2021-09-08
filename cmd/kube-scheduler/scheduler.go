/*
Copyright 2014 The Kubernetes Authors.

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


/*
OVERALL:
Two important line:
1. command := app.NewSchedulerCommand()
2. command.Execute()
 */


package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/spf13/pflag"

	cliflag "k8s.io/component-base/cli/flag"
	"k8s.io/component-base/logs"
	_ "k8s.io/component-base/logs/json/register" // for JSON log format registration
	_ "k8s.io/component-base/metrics/prometheus/clientgo"
	_ "k8s.io/component-base/metrics/prometheus/version" // for version metric registration
	"k8s.io/kubernetes/cmd/kube-scheduler/app"
)

func main() {
	// runSchedulerCmd is the entry
	if err := runSchedulerCmd(); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
}

func runSchedulerCmd() error {
	// TODO WHY?
	rand.Seed(time.Now().UnixNano())

	/*
	pflag works on command line, SetNormalizeFunc do the normalize on command : [getURL to geturl] so that the command can be recognized,
	WordSepNormalizeFunc changes all flags that contain "_" separators
	*/
	pflag.CommandLine.SetNormalizeFunc(cliflag.WordSepNormalizeFunc)

	/*
	NewSchedulerCommand creates a *cobra.Command object with default parameters and registryOptions
	what is cobra.Command? == used to write command line tools. ps: https://segmentfault.com/a/1190000023382214
    ENTRY1: important
	 */
	command := app.NewSchedulerCommand()

	logs.InitLogs()

	// defer will execute before the func return
	defer logs.FlushLogs()

	if err := command.Execute(); err != nil {
		return err
	}

	return nil
}
