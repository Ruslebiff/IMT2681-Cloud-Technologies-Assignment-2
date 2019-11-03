package assignment2

import (
	"context"
	"fmt"
	"log"

	"github.com/pkg/errors"

	firebase "firebase.google.com/go"

	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
)

var fb = Firebase{} // firebase db instance

// DBInit initialises a Firebase DB
func DBInit() error {
	fb.Ctx = context.Background()

	// opt = service account: load credentials file that you downloaded from project's settings. This is the access token to the database
	opt := option.WithCredentialsFile(Firebasecredentials)
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		return fmt.Errorf("error initializing app: %v", err)
	}

	fb.Client, err = app.Firestore(fb.Ctx)
	if err != nil {
		log.Fatalln(err)
	}

	return nil

}

// DBClose closes database connection
func DBClose() {
	fb.Client.Close()
}

// DBSave saves webhook to DB
func DBSave(webhook *Webhookreg) (string, error) {
	ref := fb.Client.Collection("Webhooks").NewDoc()
	webhook.ID = ref.ID
	_, err := ref.Set(fb.Ctx, webhook)
	if err != nil {
		fmt.Println("Error saving webhook to db")
		return "", errors.Wrap(err, "Error in firebase save")
	}
	return ref.ID, nil
}

// DBDelete deletes a webhook from DB
func DBDelete(id string) error {
	ref := fb.Client.Collection("Webhooks").Doc(id)
	_, err := ref.Delete(fb.Ctx)
	if err != nil {
		fmt.Println("Error deleting webhook from db: " + id)
		return errors.Wrap(err, "Error in firebase delete")
	}
	return nil

}

// DBReadall reads all webhook registrations from database
func DBReadall() ([]Webhookreg, error) {
	var tempwebhooks []Webhookreg
	webhook := Webhookreg{}
	iter := fb.Client.Collection("Webhooks").Documents(fb.Ctx)
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatalf("Failed to iterate: %v", err)
		}
		err = doc.DataTo(&webhook) // put data into temp struct
		if err != nil {
			fmt.Println("Error when converting retrieved document to webhook struct: ", err)
		}

		tempwebhooks = append(tempwebhooks, webhook) // add to temp array of webhookregs

	}
	return tempwebhooks, nil

}

// DBReadid reads a webhook from database with specified id
func DBReadid(id string) (Webhookreg, error) {
	res := Webhookreg{}
	ref, err := fb.Client.Collection("Webhooks").Doc(id).Get(fb.Ctx) // reference to document in firebase collection
	if err != nil {
		return res, err
	}
	err = ref.DataTo(&res) // data to temp struct
	if err != nil {
		return res, err
	}
	return res, nil // all good, return temp struct and no error
}
