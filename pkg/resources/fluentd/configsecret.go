// Copyright © 2019 Banzai Cloud
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package fluentd

import (
	"bytes"
	"fmt"
	"html/template"

	"github.com/banzaicloud/logging-operator/pkg/k8sutil"
	"github.com/banzaicloud/logging-operator/pkg/resources/templates"
	"github.com/banzaicloud/logging-operator/pkg/util"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

type fluentdConfig struct {
	Monitor struct {
		Enabled bool
		Port    int32
		Path    string
	}
}

func generateConfig(input fluentdConfig) string {
	output := new(bytes.Buffer)
	tmpl, err := template.New("test").Parse(fluentdInputTemplate)
	if err != nil {
		return ""
	}
	err = tmpl.Execute(output, input)
	if err != nil {
		return ""
	}
	outputString := fmt.Sprint(output.String())
	return outputString
}

func (r *Reconciler) secretConfig() (runtime.Object, k8sutil.DesiredState) {
	input := fluentdConfig{Monitor: struct {
		Enabled bool
		Port    int32
		Path    string
	}{},
	}

	if r.Logging.Spec.FluentdSpec.Metrics != nil {
		input.Monitor.Enabled = true
		input.Monitor.Port = r.Logging.Spec.FluentdSpec.Metrics.Port
		input.Monitor.Path = r.Logging.Spec.FluentdSpec.Metrics.Path
	}

	return &corev1.Secret{
		ObjectMeta: templates.FluentdObjectMeta(
			r.Logging.QualifiedName(SecretConfigName), util.MergeLabels(r.Logging.Labels, labelSelector), r.Logging),
		Data: map[string][]byte{
			"fluent.conf":  []byte(fluentdDefaultTemplate),
			"input.conf":   []byte(generateConfig(input)),
			"devnull.conf": []byte(fluentdOutputTemplate),
		},
	}, k8sutil.StatePresent
}
