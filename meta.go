package main

import (
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"

	"cloud.google.com/go/storage"
)

const (
	poem = `a little poem for you to read
something simple, no great deed
these words mark the placeno intention to intercede.
`
	meta = "http://metadata.google.internal/computeMetadata/v1/project/"
)

func projectID() string {
	url := meta + "project-id"
	client := &http.Client{}
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("Metadata-Flavor", "Google")
	resp, err := client.Do(req)

	if err != nil {
		fmt.Printf("meta failure: " + err.Error())
		return ""
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	return string(body)
}

func create(bucket, name, text string) error {
	ctx := context.Background()
	client, err := storage.NewClient(ctx)
	if err != nil {
		return err
	}
	bkt := client.Bucket(bucket)
	obj := bkt.Object(name)
	w := obj.NewWriter(ctx)
	if _, err := fmt.Fprintf(w, text); err != nil {
		return err
	}
	// Close, just like writing a file.
	if err := w.Close(); err != nil {
		return err
	}
	return nil
}

func attributes(bucket string) error {
	ctx := context.Background()
	client, err := storage.NewClient(ctx)
	if err != nil {
		return err
	}
	bkt := client.Bucket(bucket)
	attrs, err := bkt.Attrs(ctx)
	if err != nil {
		return err
	}
	fmt.Println(attrs)
	return nil
}

func readOut(bucket, name string) error {
	ctx := context.Background()
	client, err := storage.NewClient(ctx)
	if err != nil {
		return err
	}
	bkt := client.Bucket(bucket)
	obj := bkt.Object(name)
	// Read it back.
	r, err := obj.NewReader(ctx)
	if err != nil {
		return err
	}
	defer r.Close()
	if _, err := io.Copy(os.Stdout, r); err != nil {
		return err
	}
	return nil
}

func creds(projID string) {
	ctx := context.Background()
	// Use Google Application Default Credentials to authorize and authenticate the client.
	// More information about Application Default Credentials and how to enable is at
	// https://developers.google.com/identity/protocols/application-default-credentials.
	//
	// This is the recommended way of authorizing and authenticating.
	//
	// Note: The example uses the datastore client, but the same steps apply to
	// the other client libraries underneath this package.
	client, err := storage.NewClient(ctx)
	if err != nil {
		// TODO: handle error.
	}
	// Use the client.
	_ = client
}

/*
func hello() {
	bucketName := _projID + ".appspot.com"
	readOut(bucketName, "hello.txt")
}

func sayWhat() {
	gLog("hey man")
	logReader(_projID, _logName, os.Stdout)
}
*/
