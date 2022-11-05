package hetzner

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const apiBase string = "https://dns.hetzner.com/api/v1"

type hetznerApi struct {
	apiKey string
	client *http.Client
}

type hetznerZone struct {
	Id   string
	Name string
}

type hetznerRecord struct {
	ZoneId string `json:"zone_id"`
	Id     string `json:"id,omitempty"`
	Name   string `json:"name"`
	Type   string `json:"type"`
	Ttl    int    `json:"ttl"`
	Value  string `json:"value"`
}

func newHetznerApi(apiKey string) *hetznerApi {
	return &hetznerApi{
		apiKey: apiKey,
		client: &http.Client{},
	}
}

func (a *hetznerApi) getZones() ([]hetznerZone, error) {
	raw, err := a.makeRequest(http.MethodGet, "zones", nil)
	if err != nil {
		return nil, err
	}

	var zones struct {
		Zones []hetznerZone
	}
	if err := json.Unmarshal(raw, &zones); err != nil {
		return nil, err
	}

	return zones.Zones, nil
}

func (a *hetznerApi) getRecords(zoneId string) ([]hetznerRecord, error) {
	raw, err := a.makeRequest(http.MethodGet, fmt.Sprintf("records?zone_id=%v", zoneId), nil)
	if err != nil {
		return nil, err
	}

	var records struct {
		Records []hetznerRecord
	}
	if err := json.Unmarshal(raw, &records); err != nil {
		return nil, err
	}

	return records.Records, nil
}

func (a *hetznerApi) createRecord(newRecord *hetznerRecord) (string, error) {
	raw, err := json.Marshal(newRecord)
	if err != nil {
		return "", err
	}

	body, err := a.makeRequest(http.MethodPost, "records", bytes.NewReader(raw))
	if err != nil {
		return "", err
	}

	var record struct {
		Record hetznerRecord
	}
	if err := json.Unmarshal(body, &record); err != nil {
		return "", err
	}

	return record.Record.Id, nil
}

func (a *hetznerApi) deleteRecord(recordId string) error {
	_, err := a.makeRequest(http.MethodDelete, fmt.Sprintf("records/%v", recordId), nil)

	return err
}

func (a *hetznerApi) makeRequest(verb string, path string, body io.Reader) ([]byte, error) {
	req, err := http.NewRequest(verb, fmt.Sprintf("%v/%v", apiBase, path), body)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Auth-API-Token", a.apiKey)

	resp, err := a.client.Do(req)
	if err != nil {
		return nil, err
	}

	return io.ReadAll(resp.Body)
}
