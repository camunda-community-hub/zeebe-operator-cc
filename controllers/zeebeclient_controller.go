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

	if zeebeClient.Spec.ClientId == "" && zeebeClient.Spec.ClusterId != "" && zeebeClient.Status == (zeebev1.ZeebeClientStatus{}) {
		log.Info("Creating Zeebe Client for Cluster: " + zeebeClient.Spec.ClusterId + " and " + zeebeClient.Status.Status)

		createdClientResponse, err := ccClient.CreateZeebeClient(zeebeClient.Spec.ClusterId, zeebeClient.Spec.ClientName)

		if err != nil {
			log.Error(err, "failed to create zeebe client for cluster: "+zeebeClient.Spec.ClusterId)
			return reconcile.Result{}, err
		}
		log.Info("Updating Zeebe Client with: ", "ClientId", zeebeClient.Spec.ClusterId)
		zeebeClient.Spec.ClientId = createdClientResponse.ClientID
		zeebeClient.Spec.SecretName = zeebeClient.Spec.ClientName + "-secret"
		zeebeClient.Spec.ConfigMapName = zeebeClient.Spec.ClientName + "-configmap"

		objectMetaSecret := metav1.ObjectMeta{
			Name:      zeebeClient.Spec.SecretName,
			Namespace: "default",
		}
		var secret = v1.Secret{
			ObjectMeta: objectMetaSecret,
			Data:       map[string][]byte{},
		}
		secret.Data["ZEEBE_CLIENT_SECRET"] = []byte(createdClientResponse.ClientSecret)
		if err := r.Create(context.Background(), &secret); err != nil {
			log.Error(err, "Secret: "+zeebeClient.Spec.SecretName+" already exist and it was not updated")
		}

		details, err := ccClient.GetZeebeClientDetails(zeebeClient.Spec.ClusterId, createdClientResponse.ClientID)

		objectMetaConfigMap := metav1.ObjectMeta{
			Name:      zeebeClient.Spec.ConfigMapName,
			Namespace: "default",
		}

		var configMap = v1.ConfigMap{
			ObjectMeta: objectMetaConfigMap,
			Data: map[string]string{
				"ZEEBE_ADDRESS":                  details.ZEEBEADDRESS,
				"ZEEBE_AUTHORIZATION_SERVER_URL": details.ZEEBEAUTHORIZATIONSERVERURL,
				"ZEEBE_CLIENT_ID":                details.ZEEBECLIENTID},
		}

		if err := r.Create(context.Background(), &configMap); err != nil {
			//return reconcile.Result{}, err
			log.Error(err, "ConfigMap: "+zeebeClient.Spec.ConfigMapName+" already exist and it was not updated")
		}

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
