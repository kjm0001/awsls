// Code is generated. DO NOT EDIT.

package aws

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/service/ec2"
)

func ListEc2ClientVpnEndpoint(client *Client) error {
	req := client.ec2conn.DescribeClientVpnEndpointsRequest(&ec2.DescribeClientVpnEndpointsInput{})

	p := ec2.NewDescribeClientVpnEndpointsPaginator(req)
	for p.Next(context.Background()) {
		page := p.CurrentPage()

		for _, r := range page.ClientVpnEndpoints {
			fmt.Println(*r.ClientVpnEndpointId)
			for _, t := range r.Tags {
				fmt.Printf("\t%s: %s\n", *t.Key, *t.Value)
			}
		}
	}

	if err := p.Err(); err != nil {
		return err
	}

	return nil
}
