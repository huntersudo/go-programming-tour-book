package util

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

func main() {
	//dbEngine := getRealDbEngine()

}


func getRealDbEngine() *gorm.DB {
	client, err := DialWithPasswd("10.253.26.43:17536", "deployer", "eDEy!oJ40mpOL#")
	if err != nil {
		panic(err)
	}
	//out, err := client.Cmd("ls -l / ").Output()
	_, err = client.Cmd("ls -l / ").Output()
	if err != nil {
		panic(err)
	}
	//fmt.Println(string(out))
	// Now we register the ViaSSHDialer with the ssh connection as a parameter
	mysql.RegisterDialContext("mysql+tcp", (&ViaSSHDialer{client.client, nil}).Dial)

	// above is dial  ssh tunnel

	dbEngine, err := gorm.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=%t&loc=Local",
		"e_kcs",
		"NdSKzd#g_1B2lo",
		"10.253.41.2",
		"3306",
		"kcs-cluster-db",
		"utf8",
		true,
	))

	if err != nil {
		log.Fatalf("Get mysql gorm engine error %s.", err)
		return nil
	}
	//if global.ServerSetting.RunMode == "debug" {
	dbEngine.LogMode(true)

	dbEngine.DB().SetMaxIdleConns(1000)
	dbEngine.DB().SetMaxOpenConns(2000)

	//if rows, err := dbEngine.DB().Query("select clusterName,clusterid from cluster where uid='CIDC-U-00d0315528a045d28588d8083ed4ab8f';"); err == nil {
	//	for rows.Next() {
	//		var clusterName string
	//		var clusterid string
	//
	//		rows.Scan(&clusterName,&clusterid)
	//		fmt.Printf("clusterName: %s  clusterId: %s\n", clusterName, clusterid)
	//		// +-------------+--------------------------------------+
	//		//| clusterName | clusterid                            |
	//		//+-------------+--------------------------------------+
	//		//| cluster-ycx | fd8556fc-97a6-4f8a-8629-5edd78b838bf |
	//		//| test1       | 37a43ace-994d-4102-a750-f8ecb292037c |
	//		//| test2       | c800ee0c-bfc1-4ed0-aec3-c350e24aa6d5 |
	//		//| test3       | 8a7f26e6-d34b-4b2e-9283-a3c8f5542515 |
	//		//| test4       | e7f82f99-0578-489f-ad92-44b03a6dcf98 |
	//		//+-------------+--------------------------------------+
	//	}
	//	rows.Close()
	//} else {
	//	fmt.Printf("Failure: %s", err.Error())
	//}

	return dbEngine
}

func GetSQLManager() *sql.DB {

	client, err := DialWithPasswd("10.253.26.43:17536", "deployer", "eDEy!oJ40mpOL#")
	if err != nil {
		panic(err)
	}
	_, err = client.Cmd("ls -l / ").Output()
	if err != nil {
		panic(err)
	}
	//fmt.Println(string(out))
	mysql.RegisterDialContext("mysql+tcp", (&ViaSSHDialer{client.client, nil}).Dial)


	if db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@mysql+tcp(%s)/%s", "e_kcs", "NdSKzd#g_1B2lo", "10.253.41.2:3306", "kcs-cluster-db"));
		err == nil {
		//fmt.Printf("Successfully connected to the db\n")
		//if rows, err := db.Query("select clusterName,clusterid from cluster where uid='CIDC-U-00d0315528a045d28588d8083ed4ab8f';"); err == nil {
		//	for rows.Next() {
		//		var clusterName string
		//		var clusterId string
		//		rows.Scan(&clusterName,&clusterId)
		//		fmt.Printf("clusterName: %s  clusterId: %s\n", clusterName, clusterId)
		//	}
		//	rows.Close()

		return db
	} else {
		fmt.Printf("Failure: %s", err.Error())
	}
	return nil

}


