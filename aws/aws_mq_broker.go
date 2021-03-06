// Code is generated. DO NOT EDIT.

package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/mq"
)

func ListMqBroker(client *Client) ([]Resource, error) {
	req := client.mqconn.ListBrokersRequest(&mq.ListBrokersInput{})

	var result []Resource

	resp, err := req.Send(context.Background())
	if err != nil {
		return nil, err
	}

	if len(resp.BrokerSummaries) > 0 {
		for _, r := range resp.BrokerSummaries {

			result = append(result, Resource{
				Type:   "aws_mq_broker",
				ID:     *r.BrokerId,
				Region: client.mqconn.Config.Region,
			})
		}
	}

	return result, nil
}
