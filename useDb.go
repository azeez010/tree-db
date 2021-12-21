package main

import (
	"fmt"
	"treedb/db"	
)

func main(){
	users := db.NewUsers("Azeez", "Olabode")
	user, ok := users.GetUser("Azeez", "Olabode")
	if ok != nil {
		panic("Connection Error")
	}

	user.CreateDatabase("Genesis")
	user.CreateDatabase("Genesiss")
	user.CreateDatabase("20Genesiss")
	user.ChooseDb("Genesis")
	
	table, err := user.CreateTable("School", "Name", []string{"Age", "Name", "Class"}, "int")
	
	data := map[string]interface{}{
		"Name": "Mangakan", 
		"Age": 1, 
		"Class": 400,
	}
	data2 := map[string]interface{}{
		"Name": "Mangakans", 
		"Age": 12, 
		"Class": 400,
	}
	

	err = table.AddData(data)
	err = table.AddData(data2)
	
	if err != nil {
		fmt.Println(err)
	}

	found, datafound := db.Search(table, "Mangakan")
	if found {
		fmt.Println(datafound.Data["Name"])
	}

	fmt.Println("Table Len: ", table.Len)	
	table.PrintAll()

	user.ChooseDb("Genesiss")
	table, err = user.CreateTable("School", "Name", []string{"Age", "Name", "Class"}, "int")
	
	data = map[string]interface{}{
		"Name": "Mangakan", 
		"Age": 1, 
		"Class": 400,
	}
	data2 = map[string]interface{}{
		"Name": "Mangakans", 
		"Age": 12, 
		"Class": 400,
	}

	
	err = table.AddData(data)
	found, datafound = db.Search(table, "Mangakan")
	if found {
		fmt.Println(datafound.Data["Name"])
	}

	fmt.Println("Table Len: ", table.Len)	
	table.PrintAll()
}
