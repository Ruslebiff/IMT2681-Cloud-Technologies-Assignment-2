# IMT2681_assignment_2
This is a submission for Assignment 2 in IMT2681 Cloud Technologies at NTNU 2019


# Endpoints: 
"/repocheck/v1/commits/"   (?limit=<number>&auth=<access-token>)
"/repocheck/v1/languages/" (?limit=<number>&auth=<access-token>)
"/repocheck/v1/issues/"    (?type=users|labels&auth=<access-token>)
"/repocheck/v1/status/"
"/repocheck/v1/webhooks/"

# Webhook usage:
GET-request to "/repocheck/v1/webhooks/" to list all webhooks
GET-request to "/repocheck/v1/webhooks/<webhookid>" to get specific webhook
DELETE-request to "/repocheck/v1/webhooks/<webhookid>" to get specific webhook

POST-request to "/repocheck/v1/webhooks/" to register new webhook, need JSON body like:
{
    "event": "your-event-type",
	"url": "yout-webhook-url"
}