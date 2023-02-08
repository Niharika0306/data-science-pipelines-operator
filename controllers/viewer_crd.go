/*
Copyright 2023.

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

package controllers

import (
	dspipelinesiov1alpha1 "github.com/opendatahub-io/ds-pipelines-controller/api/v1alpha1"
)

var viewerCRDTemplates = []string{
	"config/internal/viewer-crd/deployment.yaml.tmpl",
	"config/internal/viewer-crd/role.yaml.tmpl",
	"config/internal/viewer-crd/rolebinding.yaml.tmpl",
	"config/internal/viewer-crd/sa.yaml.tmpl",
}

func (r *DSPipelineReconciler) ReconcileViewerCRD(dsp *dspipelinesiov1alpha1.DSPipeline,
	params *DSPipelineParams) error {

	r.Log.Info("Applying ViewerCRD Resources")

	for _, template := range viewerCRDTemplates {
		err := r.Apply(dsp, params, template)
		if err != nil {
			return err
		}
	}

	r.Log.Info("Finished applying ViewerCRD Resources")
	return nil
}