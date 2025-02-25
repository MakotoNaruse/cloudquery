package security

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/security/armsecurity"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/v2/schema"
	"github.com/cloudquery/plugin-sdk/v2/transformers"
)

func WorkspaceSettings() *schema.Table {
	return &schema.Table{
		Name:        "azure_security_workspace_settings",
		Resolver:    fetchWorkspaceSettings,
		Description: "https://learn.microsoft.com/en-us/rest/api/defenderforcloud/workspace-settings/list?tabs=HTTP#workspacesetting",
		Multiplex:   client.SubscriptionMultiplexRegisteredNamespace("azure_security_workspace_settings", client.Namespacemicrosoft_security),
		Transform:   transformers.TransformWithStruct(&armsecurity.WorkspaceSetting{}, transformers.WithPrimaryKeys("ID")),
		Columns:     schema.ColumnList{client.SubscriptionID},
	}
}

func fetchWorkspaceSettings(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc, err := armsecurity.NewWorkspaceSettingsClient(cl.SubscriptionId, cl.Creds, cl.Options)
	if err != nil {
		return err
	}
	pager := svc.NewListPager(nil)
	for pager.More() {
		p, err := pager.NextPage(ctx)
		if err != nil {
			return err
		}
		res <- p.Value
	}
	return nil
}
