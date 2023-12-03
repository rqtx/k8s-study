package main

import (
	"github.com/pkg/errors"
	appsv1 "github.com/pulumi/pulumi-kubernetes/sdk/v3/go/kubernetes/apps/v1"
	corev1 "github.com/pulumi/pulumi-kubernetes/sdk/v3/go/kubernetes/core/v1"
	metav1 "github.com/pulumi/pulumi-kubernetes/sdk/v3/go/kubernetes/meta/v1"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/provider"
)

type PodSecurityStandard uint8

const (
	Restricted PodSecurityStandard = iota
	Baseline
	Privileged
	Root
)

type K8sDeploymentInput struct {
	Name                string
	Image               string
	Replicas            int
	Command             []string
	Args                []string
	Limits              map[string]string
	Requests            map[string]string
	PodSecurityStandard PodSecurityStandard
}

type K8sDeploymentArgs struct {
	Name     string `pulumi:"Name"`
	Replicas int
	Pod      K8sPodArgs
}

type K8sDeployment struct {
	pulumi.ResourceState

	Args     *K8sDeploymentArgs
	Resource *appsv1.Deployment
}

type K8sPodOption func(*K8sPodArgs)

func NewK8sDeployment(name string, image string, replicas int, opts ...K8sPodOption) *K8sDeployment {

	// Initialize resource
	pod := &K8sPodArgs{
		Name:                     pulumi.String(name),
		Image:                    pulumi.String(image),
		Command:                  pulumi.ToStringArray([]string{}),
		Args:                     pulumi.ToStringArray([]string{}),
		Limits:                   pulumi.ToStringMap(map[string]string{}),
		Requests:                 pulumi.ToStringMap(map[string]string{}),
		AllowPrivilegeEscalation: pulumi.Bool(false),
		Privileged:               pulumi.Bool(false),
		ReadOnlyRootFilesystem:   pulumi.Bool(true),
		RunAsNonRoot:             pulumi.Bool(true),
		RunAsUser:                pulumi.Int(10001),
		RunAsGroup:               pulumi.Int(10001),
	}

	// Apply all the functional options to configure the client.
	for _, opt := range opts {
		opt(pod)
	}

	deployment := &K8sDeployment{Args: &K8sDeploymentArgs{Name: name, Replicas: replicas, Pod: *pod}}
	return deployment
}

func (k *K8sDeployment) up(ctx *pulumi.Context, opts ...pulumi.ResourceOption) error {
	err := ctx.RegisterComponentResource("custom:k8s:Deployment", k.Args.Name, k, opts...)
	if err != nil {
		return err
	}

	appLabels := pulumi.StringMap{
		"app": k.Args.Pod.Name,
	}

	deployment, err := appsv1.NewDeployment(ctx, k.Args.Name, &appsv1.DeploymentArgs{
		Spec: appsv1.DeploymentSpecArgs{
			Selector: &metav1.LabelSelectorArgs{
				MatchLabels: appLabels,
			},
			Replicas: pulumi.Int(k.Args.Replicas),
			Template: &corev1.PodTemplateSpecArgs{
				Metadata: &metav1.ObjectMetaArgs{
					Labels: appLabels,
				},
				Spec: &corev1.PodSpecArgs{
					AutomountServiceAccountToken: pulumi.BoolPtr(false),
					Containers: corev1.ContainerArray{
						corev1.ContainerArgs{
							Name:    k.Args.Pod.Name,
							Image:   k.Args.Pod.Image,
							Command: k.Args.Pod.Command,
							Args:    k.Args.Pod.Args,
							Resources: corev1.ResourceRequirementsArgs{
								Limits:   k.Args.Pod.Limits,
								Requests: k.Args.Pod.Requests,
							},
							SecurityContext: corev1.SecurityContextArgs{
								AllowPrivilegeEscalation: k.Args.Pod.AllowPrivilegeEscalation,
								Privileged:               k.Args.Pod.Privileged,
								ReadOnlyRootFilesystem:   k.Args.Pod.ReadOnlyRootFilesystem,
								RunAsNonRoot:             k.Args.Pod.RunAsNonRoot,
								SeccompProfile: corev1.SeccompProfileArgs{
									Type: pulumi.String("RuntimeDefault"),
								},
								RunAsUser:  k.Args.Pod.RunAsUser,
								RunAsGroup: k.Args.Pod.RunAsGroup,
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
	}, pulumi.Parent(k))
	if err != nil {
		return errors.Wrap(err, "")
	}

	k.Resource = deployment
	return nil
}

func (k *K8sDeployment) Result(ctx *pulumi.Context, opts ...pulumi.ResourceOption) (*provider.ConstructResult, error) {
	return provider.NewConstructResult(k.Resource)
}

// WithTimeout is a functional option to set Command and Args
func WithCommandArgs(cmd []string, args []string) K8sPodOption {
	return func(pod *K8sPodArgs) {
		pod.Command = pulumi.ToStringArray(cmd)
		pod.Args = pulumi.ToStringArray(args)
	}
}

// WithTimeout is a functional option to set Request and Limits
func WithRequestLimits(requests map[string]string, limits map[string]string) K8sPodOption {
	return func(pod *K8sPodArgs) {
		pod.Requests = pulumi.ToStringMap(requests)
		pod.Limits = pulumi.ToStringMap(limits)
	}
}

// WithTimeout is a functional option to set Request and Limits
func WithPodSecurityStandard(pss PodSecurityStandard) K8sPodOption {
	return func(pod *K8sPodArgs) {
		switch pss {
		case Root:
			pod.AllowPrivilegeEscalation = true
			pod.Privileged = true
			pod.ReadOnlyRootFilesystem = false
			pod.RunAsNonRoot = false
			pod.RunAsUser = 0
			pod.RunAsGroup = 0
		case Privileged:
			pod.AllowPrivilegeEscalation = true
			pod.Privileged = true
			pod.ReadOnlyRootFilesystem = false
			pod.RunAsNonRoot = true
			pod.RunAsUser = 10001
			pod.RunAsGroup = 10001
		case Baseline:
			pod.AllowPrivilegeEscalation = false
			pod.Privileged = false
			pod.ReadOnlyRootFilesystem = true
			pod.RunAsNonRoot = true
			pod.RunAsUser = 10001
			pod.RunAsGroup = 10001
		case Restricted:
			pod.AllowPrivilegeEscalation = false
			pod.Privileged = false
			pod.ReadOnlyRootFilesystem = true
			pod.RunAsNonRoot = true
			pod.RunAsUser = 10001
			pod.RunAsGroup = 10001
		}
	}
}
