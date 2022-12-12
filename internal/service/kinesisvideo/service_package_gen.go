// Code generated by internal/generate/servicepackage/main.go; DO NOT EDIT.

package kinesisvideo

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-provider-aws/internal/experimental/intf"
)

type servicePackage struct {
	frameworkDataSourceFactories []func(context.Context) (datasource.DataSourceWithConfigure, error)
	frameworkResourceFactories   []func(context.Context) (resource.ResourceWithConfigure, error)
	sdkDataSourceFactories       []struct {
		TypeName string
		Factory  func() *schema.Resource
	}
	sdkResourceFactories []struct {
		TypeName string
		Factory  func() *schema.Resource
	}
}

func (p *servicePackage) Configure(ctx context.Context, meta any) error {
	return nil
}

func (p *servicePackage) FrameworkDataSources(ctx context.Context) []func(context.Context) (datasource.DataSourceWithConfigure, error) {
	return p.frameworkDataSourceFactories
}

func (p *servicePackage) FrameworkResources(ctx context.Context) []func(context.Context) (resource.ResourceWithConfigure, error) {
	return p.frameworkResourceFactories
}

func (p *servicePackage) SDKDataSources(ctx context.Context) []struct {
	TypeName string
	Factory  func() *schema.Resource
} {
	return p.sdkDataSourceFactories
}

func (p *servicePackage) SDKResources(ctx context.Context) []struct {
	TypeName string
	Factory  func() *schema.Resource
} {
	return p.sdkResourceFactories
}

func (p *servicePackage) ServicePackageName() string {
	return "kinesisvideo"
}

func (p *servicePackage) registerFrameworkDataSourceFactory(factory func(context.Context) (datasource.DataSourceWithConfigure, error)) {
	p.frameworkDataSourceFactories = append(p.frameworkDataSourceFactories, factory)
}

func (p *servicePackage) registerFrameworkResourceFactory(factory func(context.Context) (resource.ResourceWithConfigure, error)) {
	p.frameworkResourceFactories = append(p.frameworkResourceFactories, factory)
}

func (p *servicePackage) registerSDKDataSourceFactory(typeName string, factory func() *schema.Resource) {
	p.sdkDataSourceFactories = append(p.sdkDataSourceFactories, struct {
		TypeName string
		Factory  func() *schema.Resource
	}{TypeName: typeName, Factory: factory})
}

func (p *servicePackage) registerSDKResourceFactory(typeName string, factory func() *schema.Resource) {
	p.sdkResourceFactories = append(p.sdkResourceFactories, struct {
		TypeName string
		Factory  func() *schema.Resource
	}{TypeName: typeName, Factory: factory})
}

var (
	sp_                                = &servicePackage{}
	ServicePackage intf.ServicePackage = sp_
)
