package main

import (
	"fmt"
	"reloaded/core"
	"reloaded/helper"
)

func FileManipulate() {
	var inputFile string
	var outputFile string
	fmt.Println("--------------------------------------------Go ReLoaded-------------------------------------------------")
	fmt.Println("|Hello they, welcome to GoReloaded, GoReloaded is a program that takes an Input file and an output file|")
	fmt.Println("|Format the text in the input file and writes it in the output file.-----------------------------------|")
	fmt.Println("|Please ensure both input file and output file exist---------------------------------------------------|")
	fmt.Println("|======================Try Not To Break GoReloaded,as you use may you be blessed=======================|")
	fmt.Println("--------------------------------------------Go ReLoaded-------------------------------------------------")
	fmt.Print("Enter Input File Name: ")
	fmt.Scanf("%s", &inputFile)
	fmt.Println()
	fmt.Print("Enter Output File Name: ")
	fmt.Scanf("%s", &outputFile)
	readWrite(inputFile, outputFile)
}
func readWrite(inputFiles, outputFiles string) {
	files := helper.FileStruct{
		InputFileName:  inputFiles,
		OutputFileName: outputFiles,
	}
	text, err := files.ReadFile()
	if err != nil {
		fmt.Println(err)
		return
	}
	transText := core.ParseAndTransform(text)
	files.WriteTo(transText)
}
