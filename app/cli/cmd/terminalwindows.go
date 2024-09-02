//
// Copyright 2024 The Chainloop Authors.
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

//go:build windows

package cmd

import (
	"fmt"
	"os"
	"syscall"

	"golang.org/x/term"
)

func readPasswordFromTerminal() ([]byte, error) {
	fd := int(syscall.Handle(os.Stdin.Fd()))

	token, err := term.ReadPassword(fd)
	if err != nil {
		return nil, fmt.Errorf("retrieving password from stdin: %w", err)
	}

	return token, nil
}
