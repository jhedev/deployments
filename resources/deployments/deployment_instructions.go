// Copyright 2016 Mender Software AS
//
//    Licensed under the Apache License, Version 2.0 (the "License");
//    you may not use this file except in compliance with the License.
//    You may obtain a copy of the License at
//
//        http://www.apache.org/licenses/LICENSE-2.0
//
//    Unless required by applicable law or agreed to in writing, software
//    distributed under the License is distributed on an "AS IS" BASIS,
//    WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//    See the License for the specific language governing permissions and
//    limitations under the License.

package deployments

import "github.com/mendersoftware/deployments/resources/images"

type ArtifactDeploymentInstructions struct {
	ArtifactName          string      `json:"artifact_name"`
	Source                images.Link `json:"source"`
	DeviceTypesCompatible []string    `json:"device_types_compatible"`
}

type DeploymentInstructions struct {
	ID       string                         `json:"id"`
	Artifact ArtifactDeploymentInstructions `json:"artifact"`
}
