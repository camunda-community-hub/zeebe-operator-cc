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

package main

import (
	"flag"
	cc "github.com/salaboy/camunda-cloud-go-client/pkg/cc/client"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"os"
	"time"

	"context"
	"k8s.io/apimachinery/pkg/runtime"
	clientgoscheme "k8s.io/client-go/kubernetes/scheme"
	_ "k8s.io/client-go/plugin/pkg/client/auth/gcp"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/log/zap"
	zeebev1 "zeebe.io/m/v2/api/v1"
	"zeebe.io/m/v2/controllers"
	// +kubebuilder:scaffold:imports
)

var (
	scheme   = runtime.NewScheme()
	setupLog = ctrl.Log.WithName("setup")
)

func init() {
	_ = clientgoscheme.AddToScheme(scheme)

	_ = zeebev1.AddToScheme(scheme)
	// +kubebuilder:scaffold:scheme
}

var clientId = os.Getenv("CC_CLIENT_ID")
var clientSecret = os.Getenv("CC_CLIENT_SECRET")

func main() {
	var metricsAddr string
	var enableLeaderElection bool
	flag.StringVar(&metricsAddr, "metrics-addr", ":8080", "The address the metric endpoint binds to.")
	flag.BoolVar(&enableLeaderElection, "enable-leader-election", false,
		"Enable leader election for controller manager. Enabling this will ensure there is only one active controller manager.")
	flag.Parse()

	ctrl.SetLogger(zap.Logger(true))

	mgr, err := ctrl.NewManager(ctrl.GetConfigOrDie(), ctrl.Options{
		Scheme:             scheme,
		MetricsBindAddress: metricsAddr,
		LeaderElection:     enableLeaderElection,
		Port:               9443,
	})
	if err != nil {
		setupLog.Error(err, "unable to start manager")
		os.Exit(1)
	}

	if err = (&controllers.ZeebeClusterReconciler{
		Client: mgr.GetClient(),
		Log:    ctrl.Log.WithName("controllers").WithName("ZeebeCluster"),
	}).SetupWithManager(mgr); err != nil {
		setupLog.Error(err, "unable to create controller", "controller", "ZeebeCluster")
		os.Exit(1)
	}
	// +kubebuilder:scaffold:builder
	setupLog.Info("Attempting to login to Camunda Cloud ...")
	var loginOk, errCC = cc.Login(clientId, clientSecret)
	if errCC != nil {
		setupLog.Error(errCC, "Cannot connect to Camunda Cloud. Check CC_CLIENT_ID and CC_CLIENT_SECRET")
		os.Exit(1)
	}
	if loginOk {
		setupLog.Info("Logged in!")
		//Getting available plans on startup.. then working with those for all requests
		setupLog.Info("Retrieving Cluster Plans ...")
		cc.GetClusterParams()
		setupLog.Info("Retrieving Cluster Plans Done! ")

		go workerPollCCClusters(mgr)

	} else {
		setupLog.Info("Error login in!")
	}

	setupLog.Info("starting manager")
	if err := mgr.Start(ctrl.SetupSignalHandler()); err != nil {
		setupLog.Error(err, "problem running manager")
		os.Exit(1)
	}

}

func workerPollCCClusters(mgr ctrl.Manager) {
	ticker := time.NewTicker(10000 * time.Millisecond)
	// todo check if the cluster already exist
	for {
		select {
		case <-ticker.C:

			clusters, err := cc.GetClusters()
			if err != nil {
				setupLog.Error(err, "failed to get clusters")
			}
			setupLog.Info("reported clusters", "Clusters: ", clusters)
			for _, c := range clusters {
				ctx := context.Background()
				var zeebeCluster = zeebev1.ZeebeCluster{
					ObjectMeta: metav1.ObjectMeta{
						Name:      c.Name,
						Namespace: "default",
					},
					Spec: zeebev1.ZeebeClusterSpec{
						Owner:     "CC",
						Track:     true,
						ClusterId: c.ID,
					},
				}
				setupLog.Info("Creating Cluster: ", "ZeebeCluster: ", zeebeCluster)
				err := mgr.GetClient().Create(ctx, &zeebeCluster)
				if err != nil {
					setupLog.Error(err, "Error Creating ZeebeCluster: "+c.Name)
				}
			}
		}
	}

}
