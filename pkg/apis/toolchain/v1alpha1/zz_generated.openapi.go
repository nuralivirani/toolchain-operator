// +build !ignore_autogenerated

// This file was autogenerated by openapi-gen. Do not edit it manually!

package v1alpha1

import (
	spec "github.com/go-openapi/spec"
	common "k8s.io/kube-openapi/pkg/common"
)

func GetOpenAPIDefinitions(ref common.ReferenceCallback) map[string]common.OpenAPIDefinition {
	return map[string]common.OpenAPIDefinition{
		"github.com/codeready-toolchain/toolchain-operator/pkg/apis/toolchain/v1alpha1.CheInstallation":          schema_pkg_apis_toolchain_v1alpha1_CheInstallation(ref),
		"github.com/codeready-toolchain/toolchain-operator/pkg/apis/toolchain/v1alpha1.CheInstallationSpec":      schema_pkg_apis_toolchain_v1alpha1_CheInstallationSpec(ref),
		"github.com/codeready-toolchain/toolchain-operator/pkg/apis/toolchain/v1alpha1.CheInstallationStatus":    schema_pkg_apis_toolchain_v1alpha1_CheInstallationStatus(ref),
		"github.com/codeready-toolchain/toolchain-operator/pkg/apis/toolchain/v1alpha1.TektonInstallation":       schema_pkg_apis_toolchain_v1alpha1_TektonInstallation(ref),
		"github.com/codeready-toolchain/toolchain-operator/pkg/apis/toolchain/v1alpha1.TektonInstallationSpec":   schema_pkg_apis_toolchain_v1alpha1_TektonInstallationSpec(ref),
		"github.com/codeready-toolchain/toolchain-operator/pkg/apis/toolchain/v1alpha1.TektonInstallationStatus": schema_pkg_apis_toolchain_v1alpha1_TektonInstallationStatus(ref),
	}
}

func schema_pkg_apis_toolchain_v1alpha1_CheInstallation(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "CheInstallation is the Schema for the cheinstallations API",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"kind": {
						SchemaProps: spec.SchemaProps{
							Description: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"apiVersion": {
						SchemaProps: spec.SchemaProps{
							Description: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"metadata": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"),
						},
					},
					"spec": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/codeready-toolchain/toolchain-operator/pkg/apis/toolchain/v1alpha1.CheInstallationSpec"),
						},
					},
					"status": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/codeready-toolchain/toolchain-operator/pkg/apis/toolchain/v1alpha1.CheInstallationStatus"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"github.com/codeready-toolchain/toolchain-operator/pkg/apis/toolchain/v1alpha1.CheInstallationSpec", "github.com/codeready-toolchain/toolchain-operator/pkg/apis/toolchain/v1alpha1.CheInstallationStatus", "k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"},
	}
}

func schema_pkg_apis_toolchain_v1alpha1_CheInstallationSpec(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "CheInstallationSpec defines the desired state of CheInstallation",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"cheOperatorSpec": {
						SchemaProps: spec.SchemaProps{
							Description: "The configuration required for che operator",
							Ref:         ref("github.com/codeready-toolchain/toolchain-operator/pkg/apis/toolchain/v1alpha1.CheOperator"),
						},
					},
				},
				Required: []string{"cheOperatorSpec"},
			},
		},
		Dependencies: []string{
			"github.com/codeready-toolchain/toolchain-operator/pkg/apis/toolchain/v1alpha1.CheOperator"},
	}
}

func schema_pkg_apis_toolchain_v1alpha1_CheInstallationStatus(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "CheInstallationStatus defines the observed state of CheInstallation",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"conditions": {
						VendorExtensible: spec.VendorExtensible{
							Extensions: spec.Extensions{
								"x-kubernetes-patch-merge-key": "type",
								"x-kubernetes-patch-strategy":  "merge",
							},
						},
						SchemaProps: spec.SchemaProps{
							Description: "Conditions is an array of current CheInstallation conditions Supported condition types: CheInstalled, FailedToInstallChe",
							Type:        []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Ref: ref("github.com/codeready-toolchain/api/pkg/apis/toolchain/v1alpha1.Condition"),
									},
								},
							},
						},
					},
				},
			},
		},
		Dependencies: []string{
			"github.com/codeready-toolchain/api/pkg/apis/toolchain/v1alpha1.Condition"},
	}
}

func schema_pkg_apis_toolchain_v1alpha1_TektonInstallation(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "TektonInstallation is the Schema for the tektoninstallations API",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"kind": {
						SchemaProps: spec.SchemaProps{
							Description: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"apiVersion": {
						SchemaProps: spec.SchemaProps{
							Description: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"metadata": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"),
						},
					},
					"spec": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/codeready-toolchain/toolchain-operator/pkg/apis/toolchain/v1alpha1.TektonInstallationSpec"),
						},
					},
					"status": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/codeready-toolchain/toolchain-operator/pkg/apis/toolchain/v1alpha1.TektonInstallationStatus"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"github.com/codeready-toolchain/toolchain-operator/pkg/apis/toolchain/v1alpha1.TektonInstallationSpec", "github.com/codeready-toolchain/toolchain-operator/pkg/apis/toolchain/v1alpha1.TektonInstallationStatus", "k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"},
	}
}

func schema_pkg_apis_toolchain_v1alpha1_TektonInstallationSpec(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "TektonInstallationSpec defines the desired state of TektonInstallation",
				Type:        []string{"object"},
			},
		},
	}
}

func schema_pkg_apis_toolchain_v1alpha1_TektonInstallationStatus(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "TektonInstallationStatus defines the observed state of TektonInstallation",
				Type:        []string{"object"},
			},
		},
	}
}
