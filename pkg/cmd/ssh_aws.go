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
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
	"time"

	"github.com/Masterminds/semver"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// AwsInstanceAttribute stores all the critical information for creating an instance on AWS.
type AwsInstanceAttribute struct {
	ShootName                string
	InstanceID               string
	SecurityGroupName        string
	SecurityGroupID          string
	ImageID                  string
	VpcID                    string
	KeyName                  string
	SubnetID                 string
	BastionSecurityGroupID   string
	BastionInstanceName      string
	BastionIP                string
	BastionInstanceID        string
	BastionSecurityGroupName string
	UserData                 []byte
	SSHPublicKey             []byte
}

// sshToAWSNode provides cmds to ssh to aws via a bastions host and clean it up afterwards
func sshToAWSNode(nodeIP, path, user string, sshPublicKey []byte) {
	a := &AwsInstanceAttribute{}
	a.SSHPublicKey = sshPublicKey
	fmt.Println("")

	fmt.Println("(1/4) Fetching data from target shoot cluster")
	a.fetchAwsAttributes(nodeIP, path)
	fmt.Println("Data fetched from target shoot cluster.")
	fmt.Println("")

	fmt.Println("(2/4) Setting up bastion host security group")
	a.createBastionHostSecurityGroup()
	fmt.Println("")

	fmt.Println("(3/4) Creating bastion host")
	a.createBastionHostInstance()

	bastionNode := user + "@" + a.BastionIP
	node := user + "@" + nodeIP
	fmt.Println("Waiting 45 seconds until ports are open.")
	time.Sleep(45 * time.Second)

	sshCmd := fmt.Sprintf("ssh -i key -o \"ProxyCommand ssh -W %%h:%%p -i key -o StrictHostKeyChecking=no " + bastionNode + "\" " + node + " -o StrictHostKeyChecking=no")
	cmd := exec.Command("bash", "-c", sshCmd)
	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	checkError(err)

	fmt.Println("(4/4) Cleanup")
	a.cleanupAwsBastionHost()
}

// fetchAwsAttributes gets all the needed attributes for creating bastion host and its security group with given <nodeIP>.
func (a *AwsInstanceAttribute) fetchAwsAttributes(nodeIP, path string) {
	a.ShootName = getShootClusterName()
	a.InstanceID = getAwsInstanceIDForIP(nodeIP)
	terraformVersion, err := ExecCmdReturnOutput("bash", "-c", "cat "+path+"  | jq -r .terraform_version")
	checkError(err)
	c, err := semver.NewConstraint(">= 0.12.0")
	if err != nil {
		fmt.Println("Handle version not being parsable.")
		os.Exit(2)
	}
	v, err := semver.NewVersion(terraformVersion)
	if err != nil {
		fmt.Println("Handle version not being parsable.")
		os.Exit(2)
	}
	if c.Check(v) {
		subnetID, err := ExecCmdReturnOutput("bash", "-c", "cat "+path+" | jq -r .outputs.subnet_public_utility_z0.value")
		checkError(err)
		a.SubnetID = subnetID
		vpcID, err := ExecCmdReturnOutput("bash", "-c", "cat "+path+" | jq -r .outputs.vpc_id.value")
		checkError(err)
		a.VpcID = vpcID
	} else {
		subnetID, err := ExecCmdReturnOutput("bash", "-c", "cat "+path+" | jq -r .modules[].outputs.subnet_public_utility_z0.value")
		checkError(err)
		a.SubnetID = subnetID
		vpcID, err := ExecCmdReturnOutput("bash", "-c", "cat "+path+" | jq -r .modules[].outputs.vpc_id.value")
		checkError(err)
		a.VpcID = vpcID
	}
	a.SecurityGroupName = a.ShootName + "-nodes"
	a.getSecurityGroupID()
	a.BastionInstanceName = a.ShootName + "-bastions"
	a.BastionSecurityGroupName = a.ShootName + "-bsg"
	a.ImageID, err = fetchAWSImageIDByInstancePrivateIP(nodeIP)
	checkError(err)
	a.KeyName = a.ShootName + "-ssh-publickey"
	a.UserData = getBastionUserData(a.SSHPublicKey)
}

// fetchAWSImageIDByInstancePrivateIP returns the image ID (AMI) for instance with the given <privateIP>.
func fetchAWSImageIDByInstancePrivateIP(privateIP string) (string, error) {
	command := fmt.Sprintf("aws ec2 describe-instances --filters Name=private-ip-address,Values=%s --query Reservations[0].Instances[0].ImageId", privateIP)
	captured := capture()
	operate("aws", command)
	output, err := captured()
	if err != nil || output == "None" {
		return "", fmt.Errorf("could not fetch image ID (AMI) of instance with private IP address %q", privateIP)
	}
	return output, nil
}

// createBastionHostSecurityGroup finds the or creates a security group for the bastion host.
func (a *AwsInstanceAttribute) createBastionHostSecurityGroup() {
	var err error
	// check if security group exists
	a.getBastionSecurityGroupID()
	if a.BastionSecurityGroupID != "" {
		fmt.Println("Security Group exists, skipping creation.")
		return
	}

	// create security group and ssh rule
	arguments := fmt.Sprintf("aws ec2 create-security-group --group-name %s --description ssh-access --vpc-id %s", a.BastionSecurityGroupName, a.VpcID)
	captured := capture()
	operate("aws", arguments)
	capturedOutput, err := captured()
	checkError(err)
	a.BastionSecurityGroupID = strings.Trim((capturedOutput), "\n")
	arguments = fmt.Sprintf("aws ec2 authorize-security-group-ingress --group-id %s --protocol tcp --port 22 --cidr 0.0.0.0/0", a.BastionSecurityGroupID)
	operate("aws", arguments)
	fmt.Println("Bastion host security group set up.")

	// add shh rule to ec2 instance
	arguments = fmt.Sprintf("aws ec2 authorize-security-group-ingress --group-id %s --protocol tcp --port 22 --cidr 0.0.0.0/0", a.SecurityGroupID)
	captured = capture()
	operate("aws", arguments)
	_, err = captured()
	checkError(err)
	fmt.Println("Opend SSH Port on Node.")
}

// getSecurityGroupID extracts security group id of ec2 instance
func (a *AwsInstanceAttribute) getSecurityGroupID() {
	var err error
	arguments := fmt.Sprintf("aws ec2 describe-security-groups --filters Name=vpc-id,Values=%s Name=group-name,Values=%s --query SecurityGroups[*].{ID:GroupId}", a.VpcID, a.SecurityGroupName)
	captured := capture()
	operate("aws", arguments)
	a.SecurityGroupID, err = captured()
	checkError(err)
}

// getBastionSecurityGroupID extracts security group id for bastion security group
func (a *AwsInstanceAttribute) getBastionSecurityGroupID() {
	var err error
	arguments := fmt.Sprintf("aws ec2 describe-security-groups --filters Name=vpc-id,Values=%s Name=group-name,Values=%s --query SecurityGroups[*].{ID:GroupId}", a.VpcID, a.BastionSecurityGroupName)
	captured := capture()
	operate("aws", arguments)
	a.BastionSecurityGroupID, err = captured()
	checkError(err)
}

// getAwsInstanceIDForIP returns the instance ID with the given <nodeIP>.
func getAwsInstanceIDForIP(ip string) string {
	typeName, err := getTargetType()
	checkError(err)
	Client, err = clientToTarget(typeName)
	checkError(err)
	nodes, err := Client.CoreV1().Nodes().List(metav1.ListOptions{})
	checkError(err)
	for _, node := range nodes.Items {
		if ip == node.Status.Addresses[0].Address {
			for _, v := range strings.Split(node.Spec.ProviderID, "/") {
				if strings.HasPrefix(v, "i-") {
					return v
				}
			}
		}
	}
	return ""
}

// getBastionHostInstance gets bastion host instance if it exists
func (a *AwsInstanceAttribute) getBastionHostInstance() {
	var err error
	arguments := fmt.Sprintf("aws ec2 describe-instances --filter Name=vpc-id,Values=%s Name=tag:Name,Values=%s Name=instance-state-name,Values=running --query Reservations[*].Instances[].{Instance:InstanceId} --output text", a.VpcID, a.BastionSecurityGroupName)
	captured := capture()
	operate("aws", arguments)
	a.BastionInstanceID, err = captured()
	checkError(err)
}

// createBastionHostInstance find or creates a bastion host instance.
func (a *AwsInstanceAttribute) createBastionHostInstance() {

	// check if bastion host exists
	a.getBastionHostInstance()
	if a.BastionInstanceID != "" {
		fmt.Println("Bastion Host exists, skipping creation.")
		return
	}

	tmpfile, err := ioutil.TempFile(os.TempDir(), "gardener-user.sh")
	checkError(err)
	defer os.Remove(tmpfile.Name())
	_, err = tmpfile.Write(a.UserData)
	checkError(err)

	// create bastion host
	arguments := "aws " + fmt.Sprintf("ec2 run-instances --iam-instance-profile Name=%s --image-id %s --count 1 --instance-type t2.nano --key-name %s --security-group-ids %s --subnet-id %s --associate-public-ip-address --user-data file://%s --tag-specifications ResourceType=instance,Tags=[{Key=Name,Value=%s}]", a.BastionInstanceName, a.ImageID, a.KeyName, a.BastionSecurityGroupID, a.SubnetID, tmpfile.Name(), a.BastionInstanceName)
	captured := capture()
	operate("aws", arguments)
	capturedOutput, err := captured()
	checkError(err)
	words := strings.Fields(capturedOutput)
	for _, value := range words {
		if strings.HasPrefix(value, "i-") {
			a.BastionInstanceID = value
		}
	}
	fmt.Println("Bastion host instance created.")
	fmt.Println("")

	// check if bastion host is up and running, timeout after 3 minutes
	attemptCnt := 0
	for attemptCnt < 60 {
		arguments = "aws ec2 describe-instances --instance-id=" + a.BastionInstanceID + " --query Reservations[*].Instances[].[State.Name] --output text"
		captured = capture()
		operate("aws", arguments)
		capturedOutput, err = captured()
		checkError(err)
		fmt.Println("Instance State: " + capturedOutput)
		if strings.Trim(capturedOutput, "\n") == "running" {
			arguments := "aws ec2 describe-instances --instance-id " + a.BastionInstanceID
			captured := capture()
			operate("aws", arguments)
			capturedOutput, err := captured()
			words := strings.Fields(capturedOutput)
			checkError(err)
			ip := ""
			for _, value := range words {
				if isIP(value) && !strings.HasPrefix(value, "10.") {
					ip = value
					break
				}
			}
			a.BastionIP = ip
			return
		}
		time.Sleep(time.Second * 2)
		attemptCnt++
	}
	if attemptCnt == 90 {
		fmt.Println("Bastion server instance timeout. Please try again.")
		os.Exit(2)
	}
}

// cleanupAwsBastionHost cleans up the bastion host for the targeted cluster.
func (a *AwsInstanceAttribute) cleanupAwsBastionHost() {
	fmt.Println("Cleaning up bastion host configurations...")
	fmt.Println("")
	fmt.Println("Starting cleanup")
	fmt.Println("")

	// clean up bastion host instance
	fmt.Println("  (1/3) Cleaning up bastion host instance")
	arguments := fmt.Sprintf("aws ec2 terminate-instances --instance-ids %s", a.BastionInstanceID)
	captured := capture()
	operate("aws", arguments)
	capturedOutput, err := captured()
	checkError(err)
	fmt.Println(capturedOutput)

	// remove shh rule from ec2 instance
	fmt.Println("  (2/3) Close SSH Port on Node.")
	arguments = fmt.Sprintf("aws ec2 revoke-security-group-ingress --group-id %s --protocol tcp --port 22 --cidr 0.0.0.0/0", a.SecurityGroupID)
	captured = capture()
	operate("aws", arguments)
	capturedOutput, err = captured()
	checkError(err)
	fmt.Println("  Closed SSH Port on Node.")
	fmt.Println(capturedOutput)

	// clean up bastion security group
	fmt.Println("  (3/3) Clean up bastion host security group")
	fmt.Println("")
	fmt.Println("  Waiting 45 seconds until instance is deleted to remove all dependencies.")
	time.Sleep(time.Second * 45)
	arguments = fmt.Sprintf("aws ec2 delete-security-group --group-id %s", a.BastionSecurityGroupID)
	captured = capture()
	operate("aws", arguments)
	_, err = captured()
	checkError(err)
	fmt.Println("")
	fmt.Println("Bastion host configurations successfully cleaned up.")
}
