package time

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

/*
curl "http://worldtimeapi.org/api/ip"
{
	"abbreviation":"EDT","client_ip":"2601:801:4100:302:ff04:b11a:aa92:1094",
	"datetime":"2021-10-27T23:26:45.709286-04:00","day_of_week":3,"day_of_year":300,
	"dst":true,"dst_from":"2021-03-14T07:00:00+00:00","dst_offset":3600,"dst_until":"2021-11-07T06:00:00+00:00","raw_offset":-18000,
	"timezone":"America/Indiana/Indianapolis",
	"unixtime":1635391605,
	"utc_datetime":"2021-10-28T03:26:45.709286+00:00",
	"utc_offset":"-04:00","week_number":43
}
*/
func GetUtcTime(c *gin.Context) {

	response, err := http.Get("http://worldtimeapi.org/api/ip")

	if err != nil {
		errorMessage := fmt.Sprintf("The HTTP request failed with error %s\n", err)
		c.JSON(http.StatusFailedDependency, errorMessage)
	} else {
		data, _ := ioutil.ReadAll(response.Body)
		fmt.Println(string(data))
		jsonMessage := json.RawMessage(data)
		c.JSON(http.StatusOK, jsonMessage)
	}

}

// examples include mapping results into structs for more concise response
// https://appdividend.com/2020/03/04/golang-how-to-serialize-json-string-in-go-example/
// https://www.thepolyglotdeveloper.com/2017/07/consume-restful-api-endpoints-golang-application/
