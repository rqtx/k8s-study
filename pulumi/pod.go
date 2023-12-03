package main

import "github.com/pulumi/pulumi/sdk/v3/go/pulumi"

type K8sPodArgs struct {
	Name                     pulumi.StringInput `pulumi:"Name"     copier:"-"`
	Image                    pulumi.StringInput `pulumi:"Image"    copier:"-"`
	Limits                   pulumi.StringMap   `pulumi:"Limits"   copier:"-"`
	Requests                 pulumi.StringMap   `pulumi:"Requests" copier:"-"`
	Command                  pulumi.StringArray `pulumi:"Command"  copier:"-"`
	Args                     pulumi.StringArray `pulumi:"Args"     copier:"-"`
	AllowPrivilegeEscalation pulumi.Bool        `pulumi:"AllowPrivilegeEscalation" copier:"must"`
	Privileged               pulumi.Bool        `pulumi:"Privileged"               copier:"must"`
	ReadOnlyRootFilesystem   pulumi.Bool        `pulumi:"ReadOnlyRootFilesystem"   copier:"must"`
	RunAsNonRoot             pulumi.Bool        `pulumi:"RunAsNonRoot"             copier:"must"`
	RunAsUser                pulumi.Int         `pulumi:"RunAsUser"                copier:"must"`
	RunAsGroup               pulumi.Int         `pulumi:"RunAsGroup"               copier:"must"`
}
