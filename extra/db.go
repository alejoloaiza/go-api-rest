package extra

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

var dbpostgre *sql.DB
var err error

func DBConnectPostgres() {
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", Localconfig.DBHost[0], Localconfig.DBPort[0], Localconfig.DBUser[0], Localconfig.DBPass[0], Localconfig.DBName[0])
	//fmt.Println(psqlInfo)
	dbpostgre, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	//fmt.Println(">>>>>>>>>>>>>>>>> Successfully connected to Database <<<<<<<<<<<<<<<<<")
}

func DBGetAllPostgres(points []float64) []Asset {
	var resultAsset []Asset
	if len(points)%2 == 0 && len(points) > 1 {
		polygon := `'POLYGON((`
		for i := 0; i < len(points); i = i + 2 {
			polygon = polygon + fmt.Sprintf(`%.6f %.6f`, points[i], points[i+1])
			if i+2 < len(points) {
				polygon = polygon + ","
			}
		}
		polygon = polygon + `))' `
		fmt.Println(polygon)
		sqlStatement := `
			SELECT "Agency","Business","Code","Type","Location","City","Area","Price","Numrooms","Numbaths","Status","Link"
			FROM parallel.webscrapingresults 
			WHERE ST_Contains(ST_GeomFromText( ` + polygon + ` ) ,"Coordinate" ); `
		results, err := dbpostgre.Query(sqlStatement)
		fmt.Println(sqlStatement)
		if err != nil {
			fmt.Println(err)
		}
		for results.Next() {
			var dbAsset Asset
			//	var dbAgency string
			err := results.Scan(&dbAsset.Agency, &dbAsset.Business, &dbAsset.Code, &dbAsset.Type, &dbAsset.Location, &dbAsset.City, &dbAsset.Area, &dbAsset.Price, &dbAsset.Numrooms, &dbAsset.Numbaths, &dbAsset.Status, &dbAsset.Link)
			//fmt.Printf("%s %s %s %s %s %s %s %s %s %t %s \n", dbAsset.Agency, dbAsset.Business, dbAsset.Code, dbAsset.Type, dbAsset.Location, dbAsset.City, dbAsset.Area, dbAsset.Numrooms, dbAsset.Numbaths, dbAsset.Status, dbAsset.Link)
			resultAsset = append(resultAsset, dbAsset)
			if err != nil {
				panic(err)
			}
			//fmt.Println("uid | username | department | created ")
			//fmt.Printf("%3v | %8v | %6v | %6v\n", uid, username, department, created)
		}
		return resultAsset
	}
	return nil
}
