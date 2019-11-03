# IMT2681_assignment_2
This is a submission for Assignment 2 in IMT2681 Cloud Technologies at NTNU 2019


# Endpoints: 
"/repocheck/v1/commits/"   (?limit=<number>&auth=<access-token>)\
"/repocheck/v1/languages/" (?limit=<number>&auth=<access-token>)\
"/repocheck/v1/issues/"    (?type=users|labels&auth=<access-token>)\
"/repocheck/v1/status/"\
"/repocheck/v1/webhooks/"

# Webhook usage:
GET-request to "/repocheck/v1/webhooks/" to list all webhooks\
GET-request to "/repocheck/v1/webhooks/<webhookid>" to get specific webhook\
DELETE-request to "/repocheck/v1/webhooks/<webhookid>" to get specific webhook\
\
POST-request to "/repocheck/v1/webhooks/" to register new webhook, need JSON body like:\
{\
    "event": "your-event-type",\
	"url": "yout-webhook-url"\
}

# File descriptions: 
main.go         - main function to run the application\
globals.go      - Global variables and consts are stored here, including API URL.\
structs.go      - All struct types\
handler.go      - All HTTP handlers\
functions.go    - Global functions\
firebase.go     - Database functions

# Database usage: 
If you copy this repo, you can make your own Firebase database for this program to store webhooks. \
To connect to database, you need an access token file. Download this from the Firebase project settings (service accounts) and save it in the root folder for this application. Do NOT share this file with anyone. It contains your own private key for accessing the database. \
Example name for this file is assignment2-2c6b0-firebase-adminsdk-9dvth-77d8aa990f.json\
It is required to have a json file like this for the program to work. 