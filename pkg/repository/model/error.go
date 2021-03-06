// Copyright 2020 PingCAP, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// See the License for the specific language governing permissions and
// limitations under the License.

package model

import (
	"errors"
)

var (
	// ErrorConflict indicates manifest conflict
	ErrorConflict = errors.New("manifest conflict")
	// ErrorMissingKey indicates that the private key is missing
	ErrorMissingKey = errors.New("the private key is missing")
	// ErrorMissingOwner indicates that the owner is not found
	ErrorMissingOwner = errors.New("owner not found")
	// ErrorWrongSignature indicates that the signature is not correct
	ErrorWrongSignature = errors.New("the signature is not correct")
)
