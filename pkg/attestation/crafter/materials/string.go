//
// Copyright 2023 The Chainloop Authors.
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

package materials

import (
	"context"
	"fmt"
	"strings"

	schemaapi "github.com/chainloop-dev/chainloop/app/controlplane/api/workflowcontract/v1"
	api "github.com/chainloop-dev/chainloop/pkg/attestation/crafter/api/attestation/v1"
	cr_v1 "github.com/google/go-containerregistry/pkg/v1"
)

type StringCrafter struct {
	*crafterCommon
}

func NewStringCrafter(materialSchema *schemaapi.CraftingSchema_Material) (*StringCrafter, error) {
	if materialSchema.Type != schemaapi.CraftingSchema_Material_STRING {
		return nil, fmt.Errorf("material type is not string")
	}

	return &StringCrafter{
		&crafterCommon{input: materialSchema},
	}, nil
}

func (i *StringCrafter) Craft(_ context.Context, value string) (*api.Attestation_Material, error) {
	hash, _, err := cr_v1.SHA256(strings.NewReader(value))
	if err != nil {
		return nil, fmt.Errorf("generating digest: %w", err)
	}
	return &api.Attestation_Material{
		MaterialType: i.input.Type,
		M: &api.Attestation_Material_String_{
			String_: &api.Attestation_Material_KeyVal{
				// TODO: remove once we know servers are not running server-side validation
				Id:    i.input.Name,
				Value: value, Digest: hash.String(),
			},
		},
	}, nil
}
