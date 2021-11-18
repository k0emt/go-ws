package time

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

type worldTime struct {
	TimeZoneAbbreviation string `json:"abbreviation"` // "abbreviation":"EDT"
	ClientIp             string `json:"client_ip"`    // "client_ip":"2601:801:4100:302:ff04:b11a:aa92:1094"
	DateTime             string `json:"datetime"`     // "datetime":"2021-10-27T23:26:45.709286-04:00"
	DayOfWeek            int    `json:"day_of_week"`  // "day_of_week":3
	DayOfYear            int    `json:"day_of_year"`  // "day_of_year":300
	CurrentlyDst         bool   `json:"dst"`          // "dst":true
	DstFrom              string `json:"dst_from"`     // "dst_from":"2021-03-14T07:00:00+00:00"
	DstOffset            int    `json:"dst_offset"`   // "dst_offset":3600
	DstUntil             string `json:"dst_until"`    // "dst_until":"2021-11-07T06:00:00+00:00"
	RawOffset            int    `json:"raw_offset"`   // "raw_offset":-18000
	TimeZone             string `json:"timezone"`     // "timezone":"America/Indiana/Indianapolis"
	UnixTime             int64  `json:"unixtime"`     // "unixtime":1635391605
	UtcDateTime          string `json:"utc_datetime"` // "utc_datetime":"2021-10-28T03:26:45.709286+00:00"
	UtcOffset            string `json:"utc_offset"`   // "utc_offset":"-04:00"
	WeekNumber           int    `json:"week_number"`  // "week_number":43
}

// TODO: refactor for testability
func GetUtcTime(c *gin.Context) {

	response, err := http.Get("http://worldtimeapi.org/api/ip")

	if err != nil {
		errorMessage := fmt.Sprintf("The HTTP request failed with error %s\n", err)
		c.JSON(http.StatusFailedDependency, errorMessage)
	} else {
		data, _ := ioutil.ReadAll(response.Body)
		fmt.Println(string(data))

		var wt worldTime
		json.Unmarshal(data, &wt)
		fmt.Printf("--------\n%#v\n", wt)

		jsonMessage := json.RawMessage(data)
		c.JSON(http.StatusOK, jsonMessage)
	}

}

// examples include mapping results into structs for more concise response
// https://appdividend.com/2020/03/04/golang-how-to-serialize-json-string-in-go-example/
// https://www.thepolyglotdeveloper.com/2017/07/consume-restful-api-endpoints-golang-application/
