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

package controller

import (
	"context"

	v1 "github.com/forselli-stratio/grafana-operator/api/v1"
	grafanaclient "github.com/forselli-stratio/grafana-operator/internal/controller/grafana"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"
)

// GrafanaOrganizationReconciler reconciles a GrafanaOrganization object
type GrafanaOrganizationReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

//+kubebuilder:rbac:groups=grafana.stratio.com,resources=grafanaorganizations,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=grafana.stratio.com,resources=grafanaorganizations/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=grafana.stratio.com,resources=grafanaorganizations/finalizers,verbs=update

func (r *GrafanaOrganizationReconciler) syncOrganizations(ctx context.Context) (ctrl.Result, error) {
	log := log.FromContext(ctx)

	// get all grafana organizations crds
	allGrafanaOrganizationsCrds := &v1.GrafanaOrganizationList{}
	var opts []client.ListOption
	err := r.Client.List(ctx, allGrafanaOrganizationsCrds, opts...)
	if err != nil {
		return ctrl.Result{
			Requeue: true,
		}, err
	}

	// get all grafana organizations
	g, err := grafanaclient.NewGrafanaClient("http://localhost:3000")
	if err != nil {
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}
	allGrafanaOrganizations, err := g.Orgs()

	// sync organizations, delete organizations from grafana that do no longer have a cr
	for _, organization := range allGrafanaOrganizations {
		if allGrafanaOrganizationsCrds.Find(organization.Name) == nil {
			g, err := grafanaclient.NewGrafanaClient("http://localhost:3000")
			if err != nil {
				return ctrl.Result{}, client.IgnoreNotFound(err)
			}
			err = g.DeleteOrg(organization.ID)
			if err != nil {
				log.Error(err, "Unable to delete Grafana organization")
			}
		}
	}


	return ctrl.Result{Requeue: false}, nil
}

func (r *GrafanaOrganizationReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	log := log.FromContext(ctx).WithName("GrafanaOrganizationReconciler")

	// Fetch the GrafanaOrganization
	var grafanaOrganization v1.GrafanaOrganization
    if err := r.Get(ctx, req.NamespacedName, &grafanaOrganization); err != nil {
        log.Error(err, "unable to fetch GrafanaOrganization CR")
        // we'll ignore not-found errors, since they can't be fixed by an immediate
        // requeue (we'll need to wait for a new notification), and we can get them
        // on deleted requests.
        return ctrl.Result{}, client.IgnoreNotFound(err)
    }

	log.Info("Reconciling", "spec", grafanaOrganization.Spec)

    // Create Grafana client
//	g, err := grafanaclient.NewGrafanaClient("http://localhost:3000")
//	if err != nil {
//		return ctrl.Result{}, client.IgnoreNotFound(err)
//	}

	// Check if organization exists in Grafana
//	orgExists, err := r.Exists(g, grafanaOrganization.Spec.Name)
//	if err != nil {
//		log.Error(err, "Unable to fetch Organization from Grafana")
//		return ctrl.Result{}, client.IgnoreNotFound(err)
//	}

	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *GrafanaOrganizationReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&v1.GrafanaOrganization{}).
		Complete(r)
}
