package goromadvanced

dsn := "host=SG-storm-bath-6462-5504-pgsql-master.servers.mongodirector.com user=ecom2 password=Ecom2#2024 dbname=ecom sslmode=disable"

if err := db.Migrator().DropTable(&User{}); err != nil {
	fmt.Println("Error dropping table:", err)
} else {
	fmt.Println("User table dropped successfully")
}
if err := db.Migrator().DropTable(&Order{}); err != nil {
	fmt.Println("Error dropping table:", err)
} else {
	fmt.Println("Order table dropped successfully")
}
if err := db.Migrator().DropTable(&Group{}); err != nil {
	fmt.Println("Error dropping table:", err)
} else {
	fmt.Println("Group table dropped successfully")
}
if err := db.Migrator().DropTable(&Profile{}); err != nil {
	fmt.Println("Error dropping table:", err)
} else {
	fmt.Println("Group table dropped successfully")
}