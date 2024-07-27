package main

import (
	"log"
	"os/exec"
	"strings"
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
	cmd := exec.Command("magick", srcImg, outImg)
	err := cmd.Run()
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}