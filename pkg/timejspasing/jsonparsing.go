package timejspasing

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

type Timezones struct {
	Timezones []Timezo `json:"timezones"`
}

type Timezo struct {
	Value  string  `json:"value"`
	Abbr   string  `json:"abbr"`
	Offset float32 `json:"offset"`
}

func GetOffset (s string ) float32 {

	fileContent, err := os.Open("./assets/timezone.json")

	if err != nil {
		log.Fatal(err)
		return 0
	}
	defer fileContent.Close()

	byteResult, _ := ioutil.ReadAll(fileContent)

	var tz Timezones
	json.Unmarshal(byteResult, &tz)

	var off float32
	for i := 0; i < len(tz.Timezones); i++ {
		if tz.Timezones[i].Abbr == s {
			off = tz.Timezones[i].Offset
		}
	}
	return off 
	
}