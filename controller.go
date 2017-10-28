package main

import (
	//"k8s.io/apimachinery/pkg/api/errors"
	"encoding/json"
	"fmt"
	"k8s.io/api/extensions/v1beta1"
	"k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"net/http"
)

var HOOK_TPR string = "hookjob.example.com"

type HookController struct {
	Client *kubernetes.Clientset
	TPR    *v1beta1.ThirdPartyResource
}

func NewHookController(config rest.Config) HookController {
	// creates the clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	// Get/Create the HookJob TPR
	tpr, err := clientset.ExtensionsV1beta1().ThirdPartyResources().Get(HOOK_TPR, v1.GetOptions{})
	if err != nil {
		hookjob := v1beta1.ThirdPartyResource{
			//Kind:        "ThirdPartyResource",
			//APIVersion:  "extensions/v1beta1",
			ObjectMeta: v1.ObjectMeta{
				Name: HOOK_TPR,
			},
			Description: "Webhook triggered Job",
			Versions:    []v1beta1.APIVersion{v1beta1.APIVersion{Name: "v1"}},
		}
		tpr, err = clientset.ExtensionsV1beta1().ThirdPartyResources().Create(&hookjob)
		if err != nil {
			panic(err.Error())
		}
	}

	return HookController{clientset, tpr}
}

func (c HookController) Run() {
}

func (c HookController) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	pretty, err := json.MarshalIndent(c.TPR.ObjectMeta, "", "  ")
	if err != nil {
		fmt.Fprintln(w, "error:", err)
	} else {
		fmt.Fprintf(w, "HookJob ThirdPartyResource: %s\n", pretty)
	}
}
