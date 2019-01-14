package authentication

import (
	"database/sql"
	// mysql needed
	_ "github.com/go-sql-driver/mysql"

	"fmt"
	"log"
)

// User type
type User struct {
	ID       int
	Username string
	Fullname string
	Hash     string
	Salt     string
	Disabled bool
	Invalid  bool
}

// InsertUser insert the indicated user
func InsertUser(user User) (int, error) {

	id := -1

	db, err := sql.Open("mysql",
		"edwinbrown:userpassword@tcp(127.0.0.1:3306)/dblearning")
	if err != nil {
		log.Fatal(err)
	}

	// 	INSERT INTO `dblearning`.`User`
	// (`ID`,
	// `Username`,
	// `Fullname`,
	// `Hash`,
	// `Salt`,
	// `disabled`)
	// VALUES
	// (<{ID: }>,
	// <{Username: }>,
	// <{Fullname: }>,
	// <{Hash: }>,
	// <{Salt: }>,
	// <{disabled: }>);

	fmt.Println("Seeing if database is available")
	err = db.Ping()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Let's see if the username '%s', exists\n", user.Username)

	var myuser User
	myuser, err = GetUserByName(user.Username)

	fmt.Printf("--- User returned: %+v. err returned %+v\n", myuser, err)

	fmt.Printf("--- user.Invalid %v. Negation: %v, myuser %+v\n", myuser.Invalid, !myuser.Invalid, myuser)

	if !myuser.Invalid {
		fmt.Printf("+++ user %s already exists\n", myuser.Username)
		return -3, err
	}

	fmt.Printf("OK. Now that invalid flag: %t\n", myuser.Invalid)

	rows, err := db.Query("select max(ID) as HIGH from User")

	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var high int

	for rows.Next() {
		err := rows.Scan(&high)
		if err != nil {
			log.Fatal(err)
		}

	}

	log.Println("Maximum id is: ", high)

	high++

	result, err := db.Exec("insert into dblearning.User (ID, Username, Fullname, Hash, Salt, disabled) VALUES (?,?,?,?,?,?)", high, user.Username, user.Fullname, user.Hash, user.Salt, user.Disabled)

	if err != nil {
		log.Fatal(err)
	}

	row, err := result.RowsAffected()

	log.Println("Number of rows affected: ", row)

	defer db.Close()

	return id, nil
}

// GetUserByName return the user based on the name
func GetUserByName(userName string) (User, error) {
	// spring.datasource.url: jdbc:mysql://localhost:3306/dblearning
	// spring.datasource.username: edwinbrown
	// spring.datasource.password: userpassword

	db, err := sql.Open("mysql",
		"edwinbrown:userpassword@tcp(127.0.0.1:3306)/dblearning")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Seeing if database is available")
	err = db.Ping()

	if err != nil {
		log.Fatal(err)
	}

	var dbUser User

	rows, err := db.Query("select ID, Username,Fullname,Hash,Salt,disabled from User where Username = ?", userName)

	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	dbUser.Invalid = true

	for rows.Next() {
		err := rows.Scan(&dbUser.ID, &dbUser.Username, &dbUser.Fullname, &dbUser.Hash, &dbUser.Salt, &dbUser.Disabled)
		if err != nil {
			// log.Fatal(err)
			fmt.Println(err)
			var emptyUser User
			return emptyUser, err
		}
		dbUser.Invalid = false
	}

	fmt.Printf("\t*** OK, after scan, dbUser is %+v. err %+v\n", dbUser, err)

	err = rows.Err()
	if dbUser.Invalid || err != nil {
		var emptyUser User
		emptyUser.Invalid = true
		return emptyUser, fmt.Errorf("User %s doesn't exist or db error %+v", userName, err)
	}

	if dbUser.Invalid {
		fmt.Println("User doesn't exist")
		var emptyUser User
		emptyUser.Invalid = true
		return emptyUser, fmt.Errorf("Username '%s' doesn't exist", userName)
	}

	fmt.Println("User exits **** ")

	defer db.Close()

	return dbUser, nil
}

// GetUserByID return the user based on the name
func GetUserByID(ID int) User {
	// spring.datasource.url: jdbc:mysql://localhost:3306/dblearning
	// spring.datasource.username: edwinbrown
	// spring.datasource.password: userpassword

	db, err := sql.Open("mysql",
		"edwinbrown:userpassword@tcp(127.0.0.1:3306)/dblearning")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Seeing if database is available")
	err = db.Ping()

	if err != nil {
		log.Fatal(err)
	}

	var dbUser User

	// rows, err := db.Query("select start_date, weather,temperature,precipitation,events from brooklyn_bridge")

	rows, err := db.Query("select ID, Username,Fullname,Hash,Salt,disabled from User where ID = ?", ID)

	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&dbUser.ID, &dbUser.Username, &dbUser.Fullname, &dbUser.Hash, &dbUser.Salt, &dbUser.Disabled)
		if err != nil {
			log.Fatal(err)
		}

	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	return dbUser
}
