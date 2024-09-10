package types

type ChipEncryptedV1 struct {
	Data string `json:"data,omitempty"`
}

type ChipEncryptedV2 struct {
	Version int      `json:"version"`
	Keys    []string `json:"keys"`
	Data    string   `json:"data"`
}

type Chip struct {
	Version int      `json:"version"`
	Keys    []string `json:"keys"`
	Data    NfcData  `json:"data"`
}

type NfcData struct {
	MEDID       string       `json:"med-id"`
	DNR         bool         `json:"dnr"`
	BloodType   string       `json:"blood-type"`
	Name        string       `json:"name"`
	DateOfBirth string       `json:"date-of-birth"`
	ICE         string       `json:"ICE"`
	Allergies   []string     `json:"allergies"`
	Medication  []string     `json:"medications"`
	Implants    []string     `json:"implants"`
	Conditions  []Conditions `json:"conditions"`
}

type Conditions struct {
	Name        string `json:"name"`
	Severity    int    `json:"severity"`
	Description string `json:"description"`
}
