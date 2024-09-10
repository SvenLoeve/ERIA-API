package nfc

import (
	"encoding/json"
	"viseh-api/services/cryptography/aes"
	"viseh-api/types"
)

var nfc types.Chip
var nfcEncrypted types.ChipEncryptedV2

func CreateEncryptedV2(nfc types.Chip) types.ChipEncryptedV2 {
	medData, _ := json.Marshal(nfc.Data)
	encryptedNfc, userEncryptionKeys := aes.EncryptNfc(string(medData))

	var encryptedMedKeys []string

	for _, key := range userEncryptionKeys {
		encryptedKey := aes.EncryptKeys(key)
		encryptedMedKeys = append(encryptedMedKeys, encryptedKey)
	}

	nfcEncrypted.Version = nfc.Version
	nfcEncrypted.Keys = encryptedMedKeys
	nfcEncrypted.Data = encryptedNfc

	return nfcEncrypted
}

func CreateDecryptedV2(nfcEncrypted types.ChipEncryptedV2) (types.Chip, error) {
	decryptedNfc := aes.Decrypt(nfcEncrypted.Data, nfcEncrypted.Keys[0])

	nfc.Version = nfcEncrypted.Version
	nfc.Keys = nfcEncrypted.Keys

	var medData types.NfcData

	err := json.Unmarshal([]byte(decryptedNfc), &medData)

	nfc.Data = medData

	return nfc, err
}
