package applicationautoscaling

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/applicationautoscaling"
	"github.com/aws/aws-sdk-go-v2/service/applicationautoscaling/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/mocks"
	"github.com/cloudquery/plugin-sdk/v2/faker"
	"github.com/golang/mock/gomock"
)

func buildScalingActivities(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockApplicationautoscalingClient(ctrl)
	services := client.Services{
		Applicationautoscaling: m,
	}
	c := types.ScalingActivity{}
	if err := faker.FakeObject(&c); err != nil {
		t.Fatal(err)
	}
	output := &applicationautoscaling.DescribeScalingActivitiesOutput{
		ScalingActivities: []types.ScalingActivity{c},
	}
	m.EXPECT().DescribeScalingActivities(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		output,
		nil,
	)

	return services
}

func TestScalingActivities(t *testing.T) {
	client.AllNamespaces = []string{"test-namespace"} // Just one

	client.AwsMockTestHelper(t, ScalingActivities(), buildScalingActivities, client.TestOptions{})
}
