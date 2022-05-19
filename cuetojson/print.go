package cuetojson

import (
	"encoding/json"
	"fmt"
)

func PrintAsJSON(infos []CueInfos) {
	packed, err := json.MarshalIndent(infos, "", "  ")

	if err != nil {
		fmt.Printf("An error occured: %s\n", err.Error())
		return
	}
	fmt.Printf("%s\n", string(packed))
}
