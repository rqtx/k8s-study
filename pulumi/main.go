package main

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		d := NewK8sDeployment("boilerplate", "busybox", 1,
			WithCommandArgs([]string{"/bin/sh"}, []string{"-c", "while :; do echo date; sleep 900; done"}),
			WithRequestLimits(
				map[string]string{
					"cpu":    "1m",
					"memory": "32M",
				},
				map[string]string{
					"cpu":    "1m",
					"memory": "32M",
				},
			),
			WithPodSecurityStandard(Restricted),
		)
		return d.up(ctx)
	})
}

/*
func NewK8sDeployment(ctx *pulumi.Context, name string, args *K8sDeploymentArgs, opts ...pulumi.ResourceOption) (*provider.ConstructResult, error) {

	if args == nil {
		args = &K8sDeploymentArgs{}
	}

	// Initialize resource
	k8sDeployment := &K8sDeployment{}
	err := ctx.RegisterComponentResource("custom:k8s:Deployment", name, k8sDeployment, opts...)
	if err != nil {
		return nil, err
	}

	appLabels := pulumi.StringMap{
		"app": args.Name,
	}

	deployment, err := appsv1.NewDeployment(ctx, name, &appsv1.DeploymentArgs{
		Spec: appsv1.DeploymentSpecArgs{
			Selector: &metav1.LabelSelectorArgs{
				MatchLabels: appLabels,
			},
			Replicas: pulumi.Int(1),
			Template: &corev1.PodTemplateSpecArgs{
				Metadata: &metav1.ObjectMetaArgs{
					Labels: appLabels,
				},
				Spec: &corev1.PodSpecArgs{
					AutomountServiceAccountToken: pulumi.BoolPtr(true),
					Containers: corev1.ContainerArray{
						corev1.ContainerArgs{
							Name:    args.Name,
							Image:   args.Image,
							Command: pulumi.ToStringArray([]string{"/bin/sh"}),
							Args:    pulumi.ToStringArray([]string{"-c", "while :; do echo date; sleep 900; done"}),
							Resources: corev1.ResourceRequirementsArgs{
								Limits:   args.Limits,
								Requests: args.Requests,
							},
							SecurityContext: corev1.SecurityContextArgs{
								AllowPrivilegeEscalation: pulumi.BoolPtr(false),
								Privileged:               pulumi.BoolPtr(false),
								ReadOnlyRootFilesystem:   pulumi.BoolPtr(true),
								RunAsNonRoot:             pulumi.BoolPtr(true),
								SeccompProfile: corev1.SeccompProfileArgs{
									Type: pulumi.String("RuntimeDefault"),
								},
								RunAsUser:  pulumi.IntPtr(10001),
								RunAsGroup: pulumi.IntPtr(10001),
								Capabilities: corev1.CapabilitiesArgs{
									Add:  pulumi.ToStringArray([]string{}),
									Drop: pulumi.ToStringArray([]string{"All"}),
								},
							},
						},
					},
				},
			},
		},
	}, pulumi.Parent(k8sDeployment))
	if err != nil {
		return nil, errors.Wrap(err, "")
	}

	return provider.NewConstructResult(deployment)
}
*/
