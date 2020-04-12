// Code is generated. DO NOT EDIT.

package aws

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/service/ec2"
)

func ListNetworkAcl(client *Client) error {
	req := client.ec2conn.DescribeNetworkAclsRequest(&ec2.DescribeNetworkAclsInput{})

	p := ec2.NewDescribeNetworkAclsPaginator(req)
	for p.Next(context.Background()) {
		page := p.CurrentPage()

		for _, r := range page.NetworkAcls {
			fmt.Println(*r.NetworkAclId)
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
