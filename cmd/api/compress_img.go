package main

import (
	"os/exec"
	"strings"

	"github.com/jmoiron/sqlx"
	_ "github.com/joho/godotenv/autoload"
)

func changeExtension(inputName string) string {
	splits := strings.Split(inputName, ".")
	splits = splits[:len(splits)-1]
	outName := strings.Join(splits, ".")
	outName += ".jpg"
	if outName == ".jpg" {
		return ""
	}
	return outName
}

func compressImgCli(srcImg string) error {
	outImg := changeExtension(srcImg)
	cmd := exec.Command("magick", "-quality", "90", srcImg, outImg)
	err := cmd.Run()
	if err != nil {
		return err
	}
	return nil
}

func batchCompress(dburl string) error {
	db, err := sqlx.Connect("postgres", dburl)
	if err != nil {
		return err
	}
	type Product struct {
		id    int64
		image string
	}
	products := make([]Product, 0)
	err = db.Select(products, `SELECT id, image FROM "Product" WHERE image LIKE %.png`)
	return err
}
