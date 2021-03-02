/*

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
	"context"
	cc "github.com/camunda-community-hub/camunda-cloud-go-client/pkg/cc/client"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"

	"github.com/go-logr/logr"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"

	zeebev1 "zeebe.io/m/v2/api/v1"
)

// ZeebeClientReconciler reconciles a ZeebeClient object
type ZeebeClientReconciler struct {
	client.Client
	Log logr.Logger
}

// +kubebuilder:rbac:groups=zeebe.io.zeebe,resources=zeebeclients,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=zeebe.io.zeebe,resources=zeebeclients/status,verbs=get;update;patch

func (r *ZeebeClientReconciler) Reconcile(req ctrl.Request) (ctrl.Result, error) {
	ctx := context.Background()
	log := r.Log.WithValues("zeebeclient", req.NamespacedName)

	var zeebeClient zeebev1.ZeebeClient

	err := r.Get(ctx, req.NamespacedName, &zeebeClient)
	if err != nil {
		// handle error
		// The cluster doesn't exist.. NOOP
		//log.Error(err, "Failed to get Zeebe Cluster")
	}

	if zeebeClient.Spec.ClusterId != "" && zeebeClient.Status == (zeebev1.ZeebeClientStatus{}) {
		log.Info("Creating Zeebe Client for Cluster: " + zeebeClient.Spec.ClusterId)

		createdClientResponse, err := cc.CreateZeebeClient(zeebeClient.Spec.ClusterId, zeebeClient.Spec.ClientName)

		if err != nil {
			log.Error(err, "failed to create zeebe client for cluster: "+zeebeClient.Spec.ClusterId)
			return reconcile.Result{}, err
		}
		log.Info("Updating Zeebe Client with: ", "ClientId", zeebeClient.Spec.ClusterId)
		zeebeClient.Spec.ClientId = createdClientResponse.ClientID
		zeebeClient.Spec.SecretName = zeebeClient.Spec.ClientName + "-secret"
		objectMeta := metav1.ObjectMeta{
			Name:      zeebeClient.Spec.SecretName,
			Namespace: "default"}
		var secret = v1.Secret{

			ObjectMeta: objectMeta,
			Data:       map[string][]byte{},
		}
		secret.Data["secret"] = []byte(createdClientResponse.ClientSecret)
		if err := r.Create(context.Background(), &secret); err != nil {
			return reconcile.Result{}, err
		}

		details, err := cc.GetZeebeClientDetails(zeebeClient.Spec.ClusterId, createdClientResponse.ClientID)
		zeebeClient.Spec.ZeebeClientDetails = details

		zeebeClient.Status = zeebev1.ZeebeClientStatus{
			Status: "Created",
		}

		if err := r.Update(context.Background(), &zeebeClient); err != nil {
			return reconcile.Result{}, err
		}
	}

	return ctrl.Result{}, nil
}

func (r *ZeebeClientReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&zeebev1.ZeebeClient{}).
		Complete(r)
}
