// Copyright (c) 2018 SAP SE or an SAP affiliate company. All rights reserved. This file is licensed under the Apache Software License, v. 2 except as noted otherwise in the LICENSE file
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
	"time"

	gardencorev1beta1 "github.com/gardener/gardener/pkg/apis/core/v1beta1"
	gardenerlogger "github.com/gardener/gardener/pkg/logger"
	kutil "github.com/gardener/gardener/pkg/utils/kubernetes"
	yaml "gopkg.in/yaml.v2"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	restclient "k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	clientcmdapi "k8s.io/client-go/tools/clientcmd/api"
)

// checkError checks if an error during execution occurred
func checkError(err error) {
	if err != nil {
		if exiterr, ok := err.(*exec.ExitError); ok {
			log.Println(string(exiterr.Stderr))
		}

		if debugSwitch {
			_, fn, line, _ := runtime.Caller(1)
			log.Fatalf("[error] %s:%d \n %v", fn, line, err)
		} else {
			log.Fatalf(err.Error())
		}
	}
}

// HomeDir returns homedirectory of user
func HomeDir() string {
	if h := os.Getenv("HOME"); h != "" {
		return h
	}
	return os.Getenv("USERPROFILE")
}

// FileExists check if the directory exists
func FileExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return true, err
}

// CreateDir creates a directory if it does no exist
func CreateDir(dirname string, permission os.FileMode) {
	exists, err := FileExists(dirname)
	checkError(err)
	if !exists {
		err := os.MkdirAll(dirname, permission)
		checkError(err)
	}
}

// CreateFileIfNotExists creates an empty file if it does not exist
func CreateFileIfNotExists(filename string, permission os.FileMode) {
	exists, err := FileExists(filename)
	checkError(err)
	if !exists {
		err = ioutil.WriteFile(filename, []byte{}, permission)
		checkError(err)
	}
}

// ExecCmd executes a command within set environment
func ExecCmd(input []byte, cmd string, suppressedOutput bool, environment ...string) (err error) {
	var command *exec.Cmd
	parts := strings.Fields(cmd)
	head := parts[0]
	if len(parts) > 1 {
		parts = parts[1:]
	} else {
		parts = nil
	}
	command = exec.Command(head, parts...)
	for index, env := range environment {
		if index == 0 {
			command.Env = append(os.Environ(),
				env,
			)
		} else {
			command.Env = append(command.Env,
				env,
			)
		}
	}
	var stdin = os.Stdin
	if input != nil {
		r, w, err := os.Pipe()
		if err != nil {
			return err
		}
		defer r.Close()
		go func() {
			_, err = w.Write([]byte(input))
			checkError(err)
			w.Close()
		}()
		stdin = r
	}
	if suppressedOutput {
		err = command.Run()
		if err != nil {
			os.Exit(2)
		}
	} else {
		command.Stdout = os.Stdout
		command.Stderr = os.Stderr
		command.Stdin = stdin
		err = command.Run()
		if err != nil {
			os.Exit(2)
		}
	}
	return nil
}

//ExecCmdSaveOutputFile save command output to file
func ExecCmdSaveOutputFile(input []byte, cmd string, fileName string, environment ...string) (err error) {
	var command *exec.Cmd
	parts := strings.Fields(cmd)
	head := parts[0]
	if len(parts) > 1 {
		parts = parts[1:]
	} else {
		parts = nil
	}
	command = exec.Command(head, parts...)
	for index, env := range environment {
		if index == 0 {
			command.Env = append(os.Environ(),
				env,
			)
		} else {
			command.Env = append(command.Env,
				env,
			)
		}
	}
	var stdin = os.Stdin
	if input != nil {
		r, w, err := os.Pipe()
		if err != nil {
			return err
		}
		defer r.Close()
		go func() {
			_, err = w.Write([]byte(input))
			checkError(err)
			w.Close()
		}()
		stdin = r
	}
	if _, err := os.Stat(fileName); err == nil {
		e := os.Remove(fileName)
		checkError(e)
	}
	outfile, err := os.Create(fileName)
	checkError(err)
	defer outfile.Close()
	command.Stdout = outfile
	command.Stderr = os.Stderr
	command.Stdin = stdin
	err = command.Run()
	if err != nil {
		os.Exit(2)
	}
	return nil
}

// ExecCmdReturnOutput execute cmd and return output
func ExecCmdReturnOutput(cmd string, args ...string) (output string, err error) {
	out, err := exec.Command(cmd, args...).Output()
	return strings.TrimSpace(string(out[:])), err
}

// ReadTarget file into Target
// DEPRECATED: Use `TargetReader` instead.
func ReadTarget(pathTarget string, target *Target) {
	targetFile, err := ioutil.ReadFile(pathTarget)
	checkError(err)
	err = yaml.Unmarshal(targetFile, target)
	checkError(err)
}

// NewConfigFromBytes returns a client from the given kubeconfig path
func NewConfigFromBytes(kubeconfig string) *restclient.Config {
	kubecf, err := ioutil.ReadFile(kubeconfig)
	checkError(err)
	configObj, err := clientcmd.Load(kubecf)
	if err != nil {
		fmt.Println("Could not load config")
		os.Exit(2)
	}
	clientConfig := clientcmd.NewDefaultClientConfig(*configObj, &clientcmd.ConfigOverrides{})
	client, err := clientConfig.ClientConfig()
	checkError(err)
	return client
}

// ValidateClientConfig validates that the auth info of a given kubeconfig doesn't have unsupported fields.
func ValidateClientConfig(config clientcmdapi.Config) error {
	validFields := []string{"client-certificate-data", "client-key-data", "token", "username", "password"}
	pathOfKubeconfig := getKubeConfigOfCurrentTarget()
	for user, authInfo := range config.AuthInfos {
		switch {
		case authInfo.ClientCertificate != "":
			return fmt.Errorf("client certificate files are not supported (user %q), these are the valid fields: %+v", user, validFields)
		case authInfo.ClientKey != "":
			return fmt.Errorf("client key files are not supported (user %q), these are the valid fields: %+v", user, validFields)
		case authInfo.TokenFile != "":
			return fmt.Errorf("token files are not supported (user %q), these are the valid fields: %+v", user, validFields)
		case authInfo.Impersonate != "" || len(authInfo.ImpersonateGroups) > 0:
			return fmt.Errorf("impersonation is not supported, these are the valid fields: %+v", validFields)
		case authInfo.AuthProvider != nil && len(authInfo.AuthProvider.Config) > 0:
			fmt.Printf("Kubeconfig under path %s contains auth provider configurations that could contain malicious code. Please only continue if you have verified it to be uncritical\n", pathOfKubeconfig)
			return nil
			// 	return fmt.Errorf("auth provider configurations are not supported (user %q), these are the valid fields: %+v", user, validFields)
		case authInfo.Exec != nil:
			fmt.Printf("Kubeconfig under path %s contains exec configurations that could contain malicious code. Please only continue if you have verified it to be uncritical\n", pathOfKubeconfig)
			return nil
			// 	return fmt.Errorf("exec configurations are not supported (user %q), these are the valid fields: %+v", user, validFields)
		}
	}

	return nil
}

// FetchShootFromTarget fetches shoot object from given target
func FetchShootFromTarget(target TargetInterface) (*gardencorev1beta1.Shoot, error) {
	gardenClientset, err := target.GardenerClient()
	if err != nil {
		return nil, err
	}

	var shoot *gardencorev1beta1.Shoot
	if target.Stack()[1].Kind == TargetKindProject {
		project, err := gardenClientset.CoreV1beta1().Projects().Get(target.Stack()[1].Name, metav1.GetOptions{})
		if err != nil {
			return nil, err
		}

		shoot, err = gardenClientset.CoreV1beta1().Shoots(*project.Spec.Namespace).Get(target.Stack()[2].Name, metav1.GetOptions{})
		if err != nil {
			return nil, err
		}
	} else {
		shootList, err := gardenClientset.CoreV1beta1().Shoots(metav1.NamespaceAll).List(metav1.ListOptions{})
		if err != nil {
			return nil, err
		}

		for index, s := range shootList.Items {
			if s.Name == target.Stack()[2].Name && *s.Spec.SeedName == target.Stack()[1].Name {
				shoot = &shootList.Items[index]
				break
			}
		}
	}

	return shoot, nil
}

//TidyKubeconfigWithHomeDir check if kubeconfig path contains ~, replace ~ with user home dir
func TidyKubeconfigWithHomeDir(pathToKubeconfig string) string {
	if strings.Contains(pathToKubeconfig, "~") {
		pathToKubeconfig = filepath.Clean(filepath.Join(HomeDir(), strings.Replace(pathToKubeconfig, "~", "", 1)))
	}
	return pathToKubeconfig
}

//CheckShootIsTargeted check if current target has shoot targeted
func CheckShootIsTargeted(target TargetInterface) bool {
	if (len(target.Stack()) < 3) || (target.Stack()[len(target.Stack())-1].Kind == "namespace" && target.Stack()[len(target.Stack())-2].Kind != "shoot") {
		return false
	}
	return true
}

//GardenctlDebugLog only outputs debug msg when gardencl -d or gardenctl --debug is specified
func GardenctlDebugLog(logMsg string) {
	if debugSwitch {
		var logger = gardenerlogger.NewLogger("debug")
		logger.Debugf(logMsg)
	}
}

//GardenctlInfoLog outputs information msg at all time
func GardenctlInfoLog(logMsg string) {
	var logger = gardenerlogger.NewLogger("info")
	logger.Infof(logMsg)
}

//CheckToolInstalled checks whether cliName is installed on local machine
func CheckToolInstalled(cliName string) bool {
	_, err := exec.LookPath(cliName)
	if err != nil {
		fmt.Println(cliName + " is not installed on your system")
		return false
	}
	return true
}

//PrintoutObject print object in yaml or json format. Pass os.Stdout if desired
func PrintoutObject(objectToPrint interface{}, writer io.Writer, outputFormat string) error {
	if outputFormat == "yaml" {
		yaml, err := yaml.Marshal(objectToPrint)
		if err != nil {
			return err
		}
		fmt.Fprint(writer, string(yaml))
	} else if outputFormat == "json" {
		json, err := json.MarshalIndent(objectToPrint, "", "  ")
		if err != nil {
			return err
		}
		fmt.Fprint(writer, string(json))
	} else {
		return errors.New("output format not supported: '" + outputFormat + "'")
	}
	return nil
}

//CheckIPPortReachable check whether IP with port is reachable within certain period of time
func CheckIPPortReachable(ip string, port string) error {
	attemptCount := 0
	for attemptCount < 12 {
		timeout := time.Second * 10
		conn, _ := net.DialTimeout("tcp", net.JoinHostPort(ip, port), timeout)
		if conn != nil {
			defer conn.Close()
			fmt.Printf("IP %s port %s is reachable\n", ip, port)
			return nil
		}
		fmt.Println("waiting for 10 seconds to retry")
		time.Sleep(time.Second * 10)
		attemptCount++
	}
	return fmt.Errorf("IP %s port %s is not reachable", ip, port)
}

//setHibernation hibernate a shoot or wake up a shoot
func setHibernation(shoot *gardencorev1beta1.Shoot, hibernated bool) {
	if shoot.Spec.Hibernation != nil {
		shoot.Spec.Hibernation.Enabled = &hibernated
	}
	shoot.Spec.Hibernation = &gardencorev1beta1.Hibernation{
		Enabled: &hibernated,
	}
}

//waitShootReconciled wait for the shoot to be reconciled successfully
func waitShootReconciled(shoot *gardencorev1beta1.Shoot, targetReader TargetReader) error {
	attemptCnt := 0
	fmt.Println("Shoot is being reconciled, the progress will be updated 1 min interval")
	for attemptCnt < 20 {
		newShoot := &gardencorev1beta1.Shoot{}
		target := targetReader.ReadTarget(pathTarget)
		gardenClientset, err := target.GardenerClient()
		shootList, err := gardenClientset.CoreV1beta1().Shoots("").List(metav1.ListOptions{})
		checkError(err)
		for index, s := range shootList.Items {
			if s.Name == shoot.Name && *s.Spec.SeedName == *shoot.Spec.SeedName {
				newShoot = &shootList.Items[index]
				break
			}
		}
		checkError(err)
		if newShoot.Status.LastOperation.Description == "Shoot cluster state has been successfully reconciled." &&
			newShoot.Status.LastOperation.State == "Succeeded" &&
			newShoot.Status.LastOperation.Type == "Reconcile" &&
			newShoot.Status.LastOperation.Progress == 100 {
			fmt.Println("Shoot has been successfully reconciled")
			return nil
		}
		fmt.Println("Last operation descirption is " + newShoot.Status.LastOperation.Description)
		fmt.Println("Last operation state is " + newShoot.Status.LastOperation.State)
		fmt.Println("Last operation type is " + newShoot.Status.LastOperation.Type)
		fmt.Println("Last operation progress is " + strconv.Itoa(int(newShoot.Status.LastOperation.Progress)) + "%")
		time.Sleep(time.Second * 60)
		attemptCnt++
	}
	return fmt.Errorf("Shoot not successfully reconciled within 20 mins")

}

//MergePatchShoot Merge old shoot with new shoot
func MergePatchShoot(oldShoot, newShoot *gardencorev1beta1.Shoot, targetReader TargetReader) (*gardencorev1beta1.Shoot, error) {
	patchBytes, err := kutil.CreateTwoWayMergePatch(oldShoot, newShoot)
	if err != nil {
		return nil, fmt.Errorf("failed to patch bytes")
	}

	target := targetReader.ReadTarget(pathTarget)
	gardenClientset, err := target.GardenerClient()
	patchedShoot, err := gardenClientset.CoreV1beta1().Shoots(oldShoot.GetNamespace()).Patch(oldShoot.GetName(), types.StrategicMergePatchType, patchBytes)
	if err == nil {
		*oldShoot = *patchedShoot
	}
	return patchedShoot, err
}
