// Code is generated. DO NOT EDIT.

package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/codestarnotifications"
)

func ListCodestarnotificationsNotificationRule(client *Client) ([]Resource, error) {
	req := client.codestarnotificationsconn.ListNotificationRulesRequest(&codestarnotifications.ListNotificationRulesInput{})

	var result []Resource

	p := codestarnotifications.NewListNotificationRulesPaginator(req)
	for p.Next(context.Background()) {
		page := p.CurrentPage()

		for _, r := range page.NotificationRules {

			result = append(result, Resource{
				Type:   "aws_codestarnotifications_notification_rule",
				ID:     *r.Arn,
				Region: client.codestarnotificationsconn.Config.Region,
			})
		}
	}

	if err := p.Err(); err != nil {
		return nil, err
	}

	return result, nil
}
