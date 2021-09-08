package main

import (
	"fmt"
	"os"

	"github.com/beego/beego/v2/client/orm"
	_ "github.com/lib/pq"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)



func init() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stdout})
	err := orm.RegisterDriver("postgres", orm.DRPostgres)
	if err != nil {
		log.Fatal().Err(err).Msg("register database driver")
	}
	err = orm.RegisterDataBase("default", "postgres", "user=gorm dbname=gorm password=gorm sslmode=disable")
	if err != nil {
		log.Fatal().Err(err).Msg("register database")
	}

}


func main() {

	o := orm.NewOrm()

	err := orm.RunSyncdb("default", true, true)
	if err != nil {
		log.Fatal().Err(err).Msg("sync database")
	}

	cat := PartCategory{Name: "#"}
	unit := PartMeasurementUnit{
		Name:      "Quantity",
		ShortName: "qty",
	}
	locCat := StorageLocationCategory{Name: "#"}
	loc := StorageLocation{
		Name:       "root",
		Category:   &locCat,
	}
	locCat2 := StorageLocationCategory{Name: "test", Parent: &locCat}
	part := Part{Name: "test", Category: &cat, Unit: &unit, StorageLocation: &loc}


	_, err = o.Insert(&locCat)
	if err != nil {
		log.Fatal().Err(err).Msg("insert locCat")
	}

	_, err = o.Insert(&locCat2)
	if err != nil {
		log.Fatal().Err(err).Msg("insert locCat2")
	}

	_, err = o.Insert(&loc)
	if err != nil {
		log.Fatal().Err(err).Msg("insert loc")
	}
	_, err = o.Insert(&cat)
	if err != nil {
		log.Fatal().Err(err).Msg("insert cat")
	}
	_, err = o.Insert(&unit)
	if err != nil {
		log.Fatal().Err(err).Msg("insert unit")
	}
	_, err = o.Insert(&part)
	if err != nil {
		log.Fatal().Err(err).Msg("insert part")
	}

	fmt.Println(locCat.Children)

}
