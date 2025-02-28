package config

import "os"

var DbName string = os.Getenv("DB_NAME")

var MongoURI string = os.Getenv("MONGO_URI")
