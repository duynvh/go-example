package main

import "food-delivery-service/cmd"

//func main() {
//	dsn := os.Getenv("DB_CONN_STR")
//	secretKey := os.Getenv("SECRET_KEY")
//
//	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
//
//	if err != nil {
//		log.Fatalln("Cannot connect to MySQL:", err)
//	}
//
//	log.Println("Connected:", db)
//
//	router := gin.Default()
//	router.Use(middleware.Recover())
//
//	appCtx := component.NewAppContext(db, secretKey)
//
//	mainRoute(router, appCtx)
//
//	router.Run(":3000")
//}

func main() {
	cmd.Execute()
}
