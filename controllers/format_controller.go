/*
Copyright 2022.

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
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	cachev1alpha1 "github.com/pixiake/format-operator/api/v1alpha1"
	"io"
	kubeErr "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/util/yaml"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"
)

// FormatReconciler reconciles a Format object
type FormatReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

type Fruit struct {
	Name     string  `json:"name,omitempty"`
	Features Feature `json:"features,omitempty"`
}

type Feature struct {
	Size   string `json:"size,omitempty"`
	Origin string `json:"origin,omitempty"`
}

//+kubebuilder:rbac:groups=cache.example.com,resources=formats,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=cache.example.com,resources=formats/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=cache.example.com,resources=formats/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the Format object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.11.0/pkg/reconcile
func (r *FormatReconciler) Reconcile(ctx context.Context, req ctrl.Request) (_ ctrl.Result, reterr error) {
	log := log.FromContext(ctx)

	format := &cachev1alpha1.Format{}

	// Fetch the Format object
	if err := r.Get(ctx, req.NamespacedName, format); err != nil {
		if kubeErr.IsNotFound(err) {
			log.Info("Cluster resource not found. Ignoring since object must be deleted")
			return ctrl.Result{}, nil
		}
		log.Error(err, "Failed to get Cluster")
		return ctrl.Result{}, err
	}

	if len(format.Status.Results) >= len(format.Spec.Foo) {
		return ctrl.Result{}, nil
	}

	for _, f := range format.Spec.Foo {
		if len(format.Status.Results) == 0 {

			if f == "apple" {
				fruit := Fruit{
					Name: "apple",
					Features: Feature{
						Size:   "100",
						Origin: "China",
					},
				}

				fruitsJson, err := json.Marshal(fruit)
				if err != nil {
					return ctrl.Result{}, err
				}
				ext := runtime.RawExtension{}

				d := yaml.NewYAMLOrJSONDecoder(bytes.NewReader(fruitsJson), 4096)
				if err := d.Decode(&ext); err != nil {
					if err != io.EOF {
						return ctrl.Result{}, err
					}
				}
				if err != nil {
					return ctrl.Result{}, err
				}
				format.Status.Results = append(format.Status.Results, ext)
			}
		}
		if f == "orange" {
			fruit := Fruit{
				Name: "orange",
				Features: Feature{
					Size:   "50",
					Origin: "Japan",
				},
			}

			fruitsJson, err := json.Marshal(fruit)
			if err != nil {
				return ctrl.Result{}, err
			}
			ext := runtime.RawExtension{}
			err = ext.UnmarshalJSON(fruitsJson)
			if err != nil {
				return ctrl.Result{}, err
			}
			format.Status.Results = append(format.Status.Results, ext)
		}
	}

	//ext := runtime.RawExtension{}
	//d := yaml.NewYAMLOrJSONDecoder(strings.NewReader(string(fruitsJson)), 4096)
	//if err = d.Decode(&ext); err != nil {
	//	if err != io.EOF {
	//		return ctrl.Result{}, err
	//	}
	//}
	//fmt.Println(string(format.Spec.Test.Raw))
	//format.Status.Results = format.Spec.Test

	defer func() {
		if err := r.Status().Update(ctx, format); err != nil {
			fmt.Println("test2")
			reterr = err
		}
	}()

	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *FormatReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&cachev1alpha1.Format{}).
		Complete(r)
}
