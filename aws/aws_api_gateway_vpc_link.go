// Code is generated. DO NOT EDIT.

package aws

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/service/apigateway"
)

func ListApiGatewayVpcLink(client *Client) error {
	req := client.apigatewayconn.GetVpcLinksRequest(&apigateway.GetVpcLinksInput{})

	p := apigateway.NewGetVpcLinksPaginator(req)
	for p.Next(context.Background()) {
		page := p.CurrentPage()

		for _, r := range page.Items {
			fmt.Println(*r.Id)
			for k, v := range r.Tags {
				fmt.Printf("\t%s: %s\n", k, v)
			}
		}
	}

	if err := p.Err(); err != nil {
		return err
	}

	return nil
}
