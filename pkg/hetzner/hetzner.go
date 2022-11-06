package hetzner

import (
	"fmt"
	"strings"
)

type Hetzner struct {
	api     *hetznerApi
	zoneIds map[string]string
}

func New(apiKey string) *Hetzner {
	return &Hetzner{
		api: newHetznerApi(apiKey),
	}
}

func (h *Hetzner) FetchZones() error {
	zones, err := h.api.getZones()
	if err != nil {
		return err
	}

	h.zoneIds = make(map[string]string)
	for _, zone := range zones {
		h.zoneIds[zone.Name] = zone.Id
	}

	return nil
}

func (h *Hetzner) CreateRecord(name string, recordType string, ttl int, value string) (string, error) {
	zoneId, nameInZone, err := h.getZoneId(name)
	if err != nil {
		return "", err
	}

	record := &hetznerRecord{
		ZoneId: zoneId,
		Name:   nameInZone,
		Type:   recordType,
		Ttl:    ttl,
		Value:  value,
	}
	return h.api.createRecord(record)
}

func (h *Hetzner) DeleteRecords(name string, recordType string) error {
	zoneId, nameInZone, err := h.getZoneId(name)
	if err != nil {
		return err
	}

	records, err := h.api.getRecords(zoneId)
	if err != nil {
		return err
	}

	for _, record := range records {
		if record.Name == nameInZone && record.Type == recordType {
			if err := h.api.deleteRecord(record.Id); err != nil {
				return err
			}
		}
	}

	return nil
}

func (h *Hetzner) DeleteRecord(recordId string) error {
	return h.api.deleteRecord(recordId)
}

func (h *Hetzner) getZoneId(name string) (string, string, error) {
	candidates := splitName(name)

	for i, candidate := range candidates {
		if id, ok := h.zoneIds[candidate]; ok {
			if i == 0 {
				return id, "@", nil
			}
			return id, name[:len(name)-len(candidate)-1], nil
		}
	}

	return "", "", fmt.Errorf("no zone ID found for %v", name)
}

func splitName(name string) []string {
	split := strings.Split(name, ".")
	result := make([]string, len(split))

	for i := range split {
		result[i] = strings.Join(split[i:], ".")
	}

	return result
}
