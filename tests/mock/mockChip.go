package mock

import (
	"encoding/json"
	"fmt"
	"github.com/brianvoe/gofakeit/v7"
	"time"
	"viseh-api/types"
)

func Chip() {

	var ChipData types.Chip

	ChipData = types.Chip{
		Version: 3,
		Keys:    []string{"ab27748e0d21b268f835dd2db0734407841383627938c9199a6abf58f8b5dcadad156d32898638a04ed1e7cff07b1c08a78eae9bb5ac1b09301244f4d2a27d1aeffc7fcc7f866188fa5ef92190c38468ee0e550bd72f7766d590507d"},
		Data: types.NfcData{
			Name:        gofakeit.Name(),
			DNR:         gofakeit.Bool(),
			BloodType:   gofakeit.RandomString([]string{"A", "B", "AB"}),
			DateOfBirth: gofakeit.PastDate().Format(time.DateOnly),
			ICE:         gofakeit.Phone(),
			Allergies:   []string{},
			Medication:  []string{},
			Implants:    []string{},
			Conditions:  []types.Conditions{},
		},
	}
	chipString, _ := json.Marshal(ChipData)
	fmt.Println(string(chipString))
}
