package examples

import (
	"context"
	"fmt"
	"github.com/libdns/libdns"
	"github.com/libdns/nicrudns"
	"github.com/pkg/errors"
	"time"
)

var (
	clientID         string
	secretID         string
	username         string
	password         string
	zoneName         string
	nicruServiceName string
	cachePath        string
)

func ExampleLibdnsProvider() error {
	provider := nicrudns.Provider{
		OAuth2ClientID: clientID,
		OAuth2SecretID: secretID,
		Username:       username,
		Password:       password,
		DnsServiceName: nicruServiceName,
		CachePath:      cachePath,
	}
	ctx := context.TODO()
	var records = []libdns.Record{
		{
			Type:  `A`,
			Name:  `www`,
			Value: `1.2.3.4`,
			TTL:   time.Hour,
		},
	}
	if records, err := provider.AppendRecords(ctx, zoneName, records); err != nil {
		return errors.Wrap(err, `append records error`)
	} else {
		for _, record := range records {
			fmt.Println(record.Name, record.TTL, record.TTL, record.Value)
		}
		return nil
	}
}

func ExampleNicruClient() error {
	provider := nicrudns.Provider{
		OAuth2ClientID: clientID,
		OAuth2SecretID: secretID,
		Username:       username,
		Password:       password,
		DnsServiceName: nicruServiceName,
		CachePath:      cachePath,
	}
	client := nicrudns.NewClient(&provider)
	var names = []string{`www`}
	if response, err := client.AddA(zoneName, names, `1.2.3.4`, `3600`); err != nil {
		return errors.Wrap(err, `add records error`)
	} else {
		for _, rr := range response.Data.Zone[0].Rr {
			fmt.Println(rr.Name, rr.Type, rr.Ttl, rr.A.String())
		}
		return nil
	}
}
