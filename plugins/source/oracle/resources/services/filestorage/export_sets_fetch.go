package filestorage

import (
	"context"

	"github.com/cloudquery/cloudquery/plugins/source/oracle/client"
	"github.com/cloudquery/plugin-sdk/v2/schema"
	"github.com/oracle/oci-go-sdk/v65/common"
	"github.com/oracle/oci-go-sdk/v65/filestorage"
)

func fetchExportSets(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cqClient := meta.(*client.Client)

	var page *string
	for {
		request := filestorage.ListExportSetsRequest{
			CompartmentId:      common.String(cqClient.CompartmentOcid),
			AvailabilityDomain: common.String(cqClient.AvailibilityDomain),
			Page:               page,
		}

		response, err := cqClient.OracleClients[cqClient.Region].FilestorageFilestorageClient.ListExportSets(ctx, request)

		if err != nil {
			return err
		}

		res <- response.Items

		if response.OpcNextPage == nil {
			break
		}

		page = response.OpcNextPage
	}

	return nil
}
