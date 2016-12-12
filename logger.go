// Copyright 2016 Google Inc. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

// Sample simplelog writes some entries, lists them, then deletes the log.
package main

import (
	"fmt"
	"io"
	"log"
	"time"

	"cloud.google.com/go/logging"
	"cloud.google.com/go/logging/logadmin"
	"golang.org/x/net/context"
	"google.golang.org/api/iterator"
)

type gLogger struct {
	client *logging.Client
	admin  *logadmin.Client
}

var (
	common *gLogger
)

func newClient(projID string) (*gLogger, error) {
	ctx := context.Background()
	client, err := logging.NewClient(ctx, projID)
	if err != nil {
		log.Fatalf("Failed to create logging client: %v", err)
		return nil, err
	}

	admin, err := logadmin.NewClient(ctx, projID)
	if err != nil {
		log.Fatalf("Failed to create logadmin client: %v", err)
		return nil, err
	}

	client.OnError = func(err error) {
		// Print an error to the local log.
		// For example, if Flush() failed.
		log.Printf("client.OnError: %v", err)
	}
	return &gLogger{client, admin}, nil
}

func (l *gLogger) writeEntry(logName, msg string) {
	logger := l.client.Logger(logName)

	infolog := logger.StandardLogger(logging.Info)
	infolog.Printf(msg)
	logger.Flush() // Ensure the entry is written.
}

func (l *gLogger) structuredWrite(logName string, payload interface{}) {
	logger := l.client.Logger(logName)

	logger.Log(logging.Entry{
		Payload:  payload,
		Severity: logging.Debug,
	})
	logger.Flush()
}

func (l *gLogger) deleteLog(logName string) error {
	ctx := context.Background()
	return l.admin.DeleteLog(ctx, logName)
}

func (l *gLogger) getEntries(projID, logName string) ([]*logging.Entry, error) {
	ctx := context.Background()

	fmt.Println("ENTRIES PROJ:", projID, "LOG:", logName)
	var entries []*logging.Entry
	iter := l.admin.Entries(ctx,
		// Only get entries from the log-example log.
		logadmin.Filter(fmt.Sprintf(`logName = "projects/%s/logs/%s"`, projID, logName)),
		// Get most recent entries first.
		logadmin.NewestFirst(),
	)
	pi := iter.PageInfo()
	fmt.Println("REMAINING:", pi.Remaining())

	if iter == nil {
		return []*logging.Entry{}, fmt.Errorf("no entries")
	}
	// Fetch the most recent 20 entries.
	for len(entries) < 20 {
		entry, err := iter.Next()
		if err == iterator.Done {
			return entries, nil
		}
		if err != nil {
			return nil, err
		}
		entries = append(entries, entry)
	}
	return entries, nil
}

func logReader(projID, logName string, w io.Writer) error {
	glog, err := newClient(projID)
	if err != nil {
		log.Printf("Failed to create logging client: %v\n", err)
		return err
	}
	//log.Print("Fetching and printing log entries.")
	entries, err := glog.getEntries(projID, logName)
	if err != nil {
		log.Printf("Could not get entries: %v\n", err)
		return err
	}
	//log.Printf("Found %d entries.", len(entries))
	for _, entry := range entries {
		fmt.Fprintf(w, "Entry: %6s @%s: %v\n",
			entry.Severity,
			entry.Timestamp.Format(time.RFC3339),
			entry.Payload)
	}
	return nil
}

func defaultLogger(projID string) {
	var err error
	common, err = newClient(projID)
	if err != nil {
		panic(err)
	}

}
