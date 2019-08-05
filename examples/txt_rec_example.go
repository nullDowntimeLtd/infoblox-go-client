// TXT record example
package main

import (
	"log"

	ibclient "github.com/nullDowntimeLTD/infoblox-go-client"
)

func main() {
	hostConfig := ibclient.HostConfig{
		Host:     "<your (w)api host>",
		Version:  "<your wapi version e.g.: 2.6>",
		Port:     "<your port eg.: 443>",
		Username: "<your username>",
		Password: "<your pass>",
	}

	transportConfig := ibclient.NewTransportConfig("false", 20, 10)
	requestBuilder := &ibclient.WapiRequestBuilder{}
	requestor := &ibclient.WapiHttpRequestor{}
	conn, err := ibclient.NewConnector(hostConfig, transportConfig, requestBuilder, requestor)
	if err != nil {
		log.Println(err)
	}

	defer conn.Logout()

	objMgr := ibclient.NewObjectManager(conn, "myclient", "")

	zone := "<your zone>"
	dnsView := "<your view>"
	text := "exampletext"
	name := "exampletextrecord" + "." + zone

	recordTXT, err := objMgr.CreateTXTRecord(text, name, dnsView)
	if err != nil {
		log.Fatalf("Creating TXT Record in dns view(%s) failed: %s", dnsView, err)
	}
	log.Printf("Created TXT record: %+v", recordTXT)

	txtRef := recordTXT.Ref
	txtRec, err := objMgr.GetTXTRecordByRef(txtRef)
	if err != nil {
		log.Fatalf("Getting TXT RECORD ref: %s from dns view(%s) failed: %s", txtRef, dnsView, err)
	}
	log.Printf("Retrieved TXT record: %+v", txtRec)

	delRec, err := objMgr.DeleteTXTRecord(txtRef)
	if err != nil {
		log.Fatalf("Deleting TXT record ref: %s from dns view(%s) failed: %s", txtRef, dnsView, err)
	}
	log.Printf("Deleted TXT record: %s", delRec)
}
