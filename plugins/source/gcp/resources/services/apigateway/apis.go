package apigateway

import (
	"context"

	"google.golang.org/api/iterator"

	pb "cloud.google.com/go/apigateway/apiv1/apigatewaypb"
	"github.com/cloudquery/plugin-sdk/v2/schema"
	"github.com/cloudquery/plugin-sdk/v2/transformers"
	"github.com/cloudquery/plugins/source/gcp/client"

	apigateway "cloud.google.com/go/apigateway/apiv1"
)

func Apis() *schema.Table {
	return &schema.Table{
		Name:        "gcp_apigateway_apis",
		Description: `https://cloud.google.com/api-gateway/docs/reference/rest/v1/projects.locations.apis#Api`,
		Resolver:    fetchApis,
		Multiplex:   client.ProjectMultiplexEnabledServices("apigateway.googleapis.com"),
		Transform:   client.TransformWithStruct(&pb.Api{}, transformers.WithPrimaryKeys("Name")),
		Columns: []schema.Column{
			{
				Name:     "project_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveProject,
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
		},
	}
}

func fetchApis(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)
	req := &pb.ListApisRequest{
		Parent: "projects/" + c.ProjectId + "/locations/-",
	}
	gcpClient, err := apigateway.NewClient(ctx, c.ClientOptions...)
	if err != nil {
		return err
	}
	it := gcpClient.ListApis(ctx, req, c.CallOptions...)
	for {
		resp, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return err
		}

		res <- resp
	}
	return nil
}
