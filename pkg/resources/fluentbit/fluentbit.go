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

package fluentbit

import (
	"github.com/banzaicloud/logging-operator/api/v1beta1"
	"github.com/banzaicloud/logging-operator/pkg/k8sutil"
	"github.com/banzaicloud/logging-operator/pkg/resources"
	"github.com/go-logr/logr"
	"github.com/goph/emperror"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
)

const (
	serviceAccountName        = "logging"
	clusterRoleBindingName    = "logging"
	clusterRoleName           = "logging"
	fluentBitSecretConfigName = "fluentbit"
	fluentbitDaemonSetName    = "fluentbit"
	fluentbitServiceName      = "fluentbit"
)

var labelSelector = map[string]string{
	"app": "fluent-bit",
}

// Reconciler holds info what resource to reconcile
type Reconciler struct {
	Logging *v1beta1.Logging
	*k8sutil.GenericResourceReconciler
}

// NewReconciler creates a new Fluentbit reconciler
func New(client client.Client, logger logr.Logger, logging *v1beta1.Logging) *Reconciler {
	return &Reconciler{
		Logging:                   logging,
		GenericResourceReconciler: k8sutil.NewReconciler(client, logger),
	}
}

// Reconcile reconciles the fluentBit resource
func (r *Reconciler) Reconcile() (*reconcile.Result, error) {
	for _, res := range []resources.Resource{
		r.serviceAccount,
		r.clusterRole,
		r.clusterRoleBinding,
		r.configSecret,
		r.daemonSet,
		r.monitorService,
	} {
		o, state := res()
		err := r.ReconcileResource(o, state)
		if err != nil {
			return nil, emperror.WrapWith(err, "failed to reconcile resource", "resource", o.GetObjectKind().GroupVersionKind())
		}
	}
	return nil, nil
}
