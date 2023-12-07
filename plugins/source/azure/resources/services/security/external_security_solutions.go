package security

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/security/armsecurity"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func ExternalSecuritySolutions() *schema.Table {
	return &schema.Table{
		Name:                 "azure_security_external_security_solutions",
		Resolver:             fetchExternalSecuritySolutions,
		PostResourceResolver: client.LowercaseIDResolver,
		Description:          "https://learn.microsoft.com/en-us/rest/api/defenderforcloud/external-security-solutions/list?tabs=HTTP#externalsecuritysolutionlist",
		Multiplex:            client.SubscriptionMultiplexRegisteredNamespace("azure_security_external_security_solutions", client.Namespacemicrosoft_security),
		Transform:            transformers.TransformWithStruct(&armsecurity.ExternalSecuritySolution{}, transformers.WithPrimaryKeys("ID")),
		Columns:              schema.ColumnList{client.SubscriptionID},
	}
}

func fetchExternalSecuritySolutions(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc, err := armsecurity.NewExternalSecuritySolutionsClient(cl.SubscriptionId, cl.Creds, cl.Options)
	if err != nil {
		return err
	}
	pager := svc.NewListPager(nil)
	for pager.More() {
		p, err := pager.NextPage(ctx)
		if err != nil {
			return err
		}
		results := make([]armsecurity.ExternalSecuritySolution, len(p.Value))
		for i, r := range p.Value {
			results[i] = *r.GetExternalSecuritySolution()
		}
		res <- results
	}
	return nil
}
