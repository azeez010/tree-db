package db
 
import (
	"fmt"
	"treedb/helper"
	"treedb/dberrors"
)

type Table interface {
	NewNode(val int)
	AddData(val int) bool
	Search(val int) (bool *Node)
	PrintAll()
	InitDB()
}

type Node struct {
	index int
	Number int
	fields []string
	indexfield string
	Data map[string]interface{}
	isEmpty bool
	left *Node
	right *Node
	Len int
	indexType interface{}
}

type TableMap struct {
	tables map[string]*Node
}
type DBMap struct {
	databases map[string]TableMap
} 

type User struct {
	dbs DBMap
	name string
	curdb string
	password string
}

type Users struct {
	users []User
} 

func NewUsers(name, password string) *Users {
	return &Users{users: []User{{name: name, password: password}}}
}

func (userStruct Users) GetUser(name, password string) (*User, error) {
	for _, user := range userStruct.users {
		if name == user.name && password == user.password {
			return &user, nil
		}
	}
	return nil, dberrors.ErrFailedToconnect
}

// func Storage(){
// 	Database := make(DBMap)
// 	Users := []User{}
// 	// re
// 	} 

// func GetStorage (){
// 	Storage()
// }
// func init(){
// }

// func Connect(name, password string) (*User, error) {
// 	for _, user := range Users {
// 		if name == user.name && password == user.password {
// 			return &user, nil
// 		}
// 	}
// 	return nil, dberrors.ErrFailedToconnect
// }

func (userStruct Users) CreateUser(name, password string) (*User) {
	for _, user := range userStruct.users {
		if name == user.name {
			return nil
		}
	}

	user := User{name: name, password: password}
	users := userStruct.users
	if users != nil {
		users = append(users, user) 
	}
	return &user
}

func (user *User) CreateDatabase(dbname string) (string) {
	curDblen := len(user.dbs.databases)
	_, ok := user.dbs.databases[dbname]
	
	if curDblen == 0 {
		database := DBMap{databases: map[string]TableMap{dbname: TableMap{}}}
		user.dbs = database
		return fmt.Sprintf("Database is created successfully!")
	} else {
		if !ok {
			user.dbs.databases[dbname] = TableMap{}
			return fmt.Sprintf("Database is created successfully!")
	
		} else {
			return fmt.Sprintf("Error, Database already exists!")
		}
	}
}

func (user *User) ChooseDb(dbname string) error {
	_, ok := user.dbs.databases[dbname]
	if ok {
		user.curdb = dbname 
		return nil
	}
	return dberrors.ErrDbName
}

func (user *User) CreateTable(tablename string, index string, fields []string, indexType string) (*Node, error) {
	dbname := user.curdb
	if len(dbname) > 0 {
		if stringPresent := helper.StringInSlice(index, fields); stringPresent {
			if indexTypeCorrect := helper.StringInSlice(indexType, []string{"string", "int"}); indexTypeCorrect {
				user.dbs.databases[dbname] = TableMap{tables: map[string]*Node{tablename: &Node{isEmpty: true, indexfield: index, indexType: indexType, fields: fields}}}
				return user.dbs.databases[dbname].tables[tablename],  nil
			} else {
				return nil, dberrors.ErrIndexType
			}	
		} else {
			return nil, dberrors.ErrIndexNotInFields
		}
	} else {
		return nil, dberrors.ErrDbName
	}
}

func (tree *Node ) NewTable(index string, fields []string) (*Node, error) {
	if stringPresent := helper.StringInSlice(index, fields); stringPresent {
		return &Node{isEmpty: true, fields: fields}, nil
	} else {
		return nil, dberrors.ErrIndexNotInFields
	}
}

func NewNode(val int, data map[string]interface{}) *Node {
	return &Node{Number: val, isEmpty: false, Len: 1, Data: data}
}

func (tree *Node ) AddData(data map[string]interface{}) error {
	correctFields := helper.MapKeysInSlice(data, tree.fields)
	if correctFields {
		fielddata := data[tree.indexfield]
		var val int;
		val = helper.GetIndex(fielddata)

		success := tree.Add(tree, val, data)
		
		if success {
			return nil
		} else {
			return dberrors.ErrDbInsert
		}
	} else {
		return dberrors.ErrDbInsert
	}
}

func (tree *Node ) Add(mainRoot *Node, val int, data map[string]interface{}) bool {	
	if tree.isEmpty {
		tree.Number = val
		tree.isEmpty = false
		tree.Data = data
		mainRoot.Len++
		return true
	} else {
		if val < tree.Number {
			if tree.left == nil {
				mainRoot.Len++
				tree.left = NewNode(val, data)
				tree.left.index = mainRoot.Len
			return true
		} else {
			tree.left.Add(mainRoot, val, data)
		}
	} else if val > tree.Number {
		if tree.right == nil {
			mainRoot.Len++
			tree.right = NewNode(val, data)
			tree.right.index = mainRoot.Len
			return true
		} else {
			tree.right.Add(mainRoot, val, data)
			}
		}
	}
	return false
}

func Search (tree *Node, searchdata interface{}) (bool, *Node) {
	var val int;
	val = helper.GetIndex(searchdata)
	
	searchSuccess := false
	nodeData := tree

	if tree.isEmpty {
		return false, nil
	} else  {
		if tree.Number == val {
			return true, tree
		} else if val < tree.Number {
			if tree.left == nil {
				return false, nil
			} else {
				searchSuccess, nodeData = Search(tree.left, val)
			}
		} else if val > tree.Number {
			if tree.right == nil {
				return false, nil
			} else {
				searchSuccess, nodeData = Search(tree.right, val)
			}
		}
	}
	return searchSuccess, nodeData
}


func (tree *Node ) PrintAll (){
	if tree.isEmpty {
		fmt.Printf("Nothing to print")
	} else {
			fmt.Println(tree.Number, tree.Data)

			if tree.right != nil {
				tree.right.PrintAll()
			} 
			
			if tree.left != nil {
				tree.left.PrintAll()
			}
		}
}