package restserver

import (
	"fmt"
	"net/http"
	"strconv"

	e "go-rest-api/extra"
	"strings"

	"github.com/gorilla/mux"
)

var ResultAssets []e.Asset

func GetAssets(w http.ResponseWriter, r *http.Request) {
	var allPoints []float64
	params := mux.Vars(r)
	if params["points"] != "" {
		fmt.Println(params["points"])
		pointsParam := strings.Split(params["points"], ",")
		for _, curPoint := range pointsParam {
			appendPoint, err := strconv.ParseFloat(curPoint, 64)
			if err != nil {
				continue
			}
			allPoints = append(allPoints, appendPoint)
		}
		e.DBGetAllPostgres(allPoints)
	}

	//json.NewEncoder(w).Encode(&extra.Assets{})
}
