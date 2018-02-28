// Copyright 2018 The Gardener Authors.
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

package terraformer

import (
	"errors"
	"fmt"
	"regexp"
	"sort"
	"strings"

	gardenv1beta1 "github.com/gardener/gardener/pkg/apis/garden/v1beta1"
)

// retrieveTerraformErrors gets a map <logList> whose keys are pod names and whose values are the corresponding logs,
// and it parses the logs for Terraform errors. If none are found, it will return nil, and otherwhise the list of
// found errors as string slice.
func retrieveTerraformErrors(logList map[string]string) []string {
	var foundErrors = map[string]string{}
	var errorList = []string{}

	for podName, output := range logList {
		errorMessage := findTerraformErrors(output)
		_, ok := foundErrors[errorMessage]

		// Add the errorMessage to the list of found errors (only if it does not already exist).
		if errorMessage != "" && !ok {
			foundErrors[errorMessage] = podName
		}
	}

	for errorMessage, podName := range foundErrors {
		errorList = append(errorList, fmt.Sprintf("-> Pod '%s' reported:\n%s", podName, errorMessage))
	}

	if len(errorList) == 0 {
		return nil
	}
	return errorList
}

// determineErrorCode determines the Garden error code for the given error message.
func determineErrorCode(message string) error {
	var code gardenv1beta1.ErrorCode

	switch {
	case strings.Contains(message, "UnauthorizedOperation") || strings.Contains(message, "SignatureDoesNotMatch"):
		code = gardenv1beta1.ErrorInfraUnauthorized
	case strings.Contains(message, "LimitExceeded"):
		code = gardenv1beta1.ErrorInfraQuotaExceeded
	case strings.Contains(message, "AccessDenied"):
		code = gardenv1beta1.ErrorInfraInsufficientPrivileges
	case strings.Contains(message, "DependencyViolation"):
		code = gardenv1beta1.ErrorInfraDependencies
	}

	if len(code) != 0 {
		message = fmt.Sprintf("CODE:%s %s", code, message)
	}

	return errors.New(message)
}

// findTerraformErrors gets the <output> of a Terraform run and parses it to find the occurred
// errors (which will be returned). If no errors occurred, an empty string will be returned.
func findTerraformErrors(output string) string {
	var (
		regexTerraformError = regexp.MustCompile(`(?:Error [^:]*|Errors): *([\s\S]*)`)
		regexUUID           = regexp.MustCompile(`(?i)[0-9a-f]{8}(?:-[0-9a-f]{4}){3}-[0-9a-f]{12}`)
		regexMultiNewline   = regexp.MustCompile(`\n{2,}`)

		errorMessage = output
		valid        = []string{}
	)

	// Strip optional explaination how Terraform behaves in case of errors.
	suffixIndex := strings.Index(errorMessage, "\n\nTerraform does not automatically rollback")
	if suffixIndex != -1 {
		errorMessage = errorMessage[:suffixIndex]
	}

	// Search for errors in Terraform output.
	terraformErrorMatch := regexTerraformError.FindStringSubmatch(errorMessage)
	if len(terraformErrorMatch) > 1 {
		// Remove leading and tailing spaces and newlines.
		errorMessage = strings.TrimSpace(terraformErrorMatch[1])

		// Omit (request) uuid's to allow easy determination of duplicates.
		errorMessage = regexUUID.ReplaceAllString(errorMessage, "<omitted>")

		// Sort the occurred errors alphabetically
		lines := strings.Split(errorMessage, "*")
		sort.Strings(lines)

		// Only keep the lines beginning with ' ' (actual errors)
		for _, line := range lines {
			if strings.HasPrefix(line, " ") {
				valid = append(valid, line)
			}
		}
		errorMessage = "*" + strings.Join(valid, "\n*")

		// Strip multiple newlines to one newline
		errorMessage = regexMultiNewline.ReplaceAllString(errorMessage, "\n")

		// Remove leading and tailing spaces and newlines.
		errorMessage = strings.TrimSpace(errorMessage)

		return errorMessage
	}
	return ""
}
