package main

import (
	"flag"
	"log"

	"github.com/truewebber/goitunes"
)

func main() {
	appleID := flag.String("apple-id", "", "Твой AppleID.")
	authToken := flag.String("auth-token", "", "Токен для авторизации в itunes.")
	buyingСert := flag.String("buying-cert", "", "Какой то сертификат для покупок, для покупок обязателен.")
	dsid := flag.String("dsid", "", "Не знаю что такое.")
	geo := flag.String("geo", "", "Обязательно, но не понятно чем заполнять.")
	machineGUID := flag.String("machine_guid", "", "machine guid, обязательно для покупок.")
	userAgent := flag.String("user-agent", "", "User Agent.")
	machineName := flag.String("machine-name", "", "Название компьютера, скорее всего в сети.")

	itunes, err := goitunes.NewGOiTunes(
		*appleID,
		*authToken,
		*buyingСert,
		*dsid,
		*geo,
		*machineGUID,
		*userAgent,
		*machineName,
	)
	if err != nil {
		log.Fatalln(err)
	}

	authResponse, err := itunes.Login("password")
	if err != nil {
		log.Fatalln(err)
	}

	log.Println("CreditBalance:", authResponse.CreditBalance)
	log.Println("DSID:", authResponse.DSID)
	log.Println("FreeSongBalance:", authResponse.FreeSongBalance)
	log.Println("PasswordToken:", authResponse.PasswordToken)
}
