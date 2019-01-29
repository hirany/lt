package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"

	"github.com/codegangsta/cli"
)

const (
	tabSpace  = "    "
	tabColumn = "|   "
)

func main() {

	app := cli.NewApp()
	app.Name = "lt"
	app.Usage = "show directory"
	app.Action = func(c *cli.Context) {
		err := printCurrentDir("./")
		checkError(err)
		err = scanDir("./", 1, 1)
		checkError(err)
		fmt.Println()
	}

	app.Run(os.Args)

}

func scanDir(currentDir string, deepLevel, columnBit int) error {

	list, err := ioutil.ReadDir(currentDir)
	checkError(err)
	dirNum := len(list)

	for i := 0; i < dirNum; i++ {
		if list[i].IsDir() {
			if i+1 == dirNum {
				printTab(deepLevel, columnBit)
				fmt.Println()
				printTab(deepLevel-1, columnBit)
				fmt.Printf("└-%s\n", list[i].Name())
				scanDir(path.Join(currentDir, list[i].Name()), deepLevel+1, columnBit+1<<uint(deepLevel)-1<<uint(deepLevel-1))
			} else {
				printTab(deepLevel, columnBit)
				fmt.Println()
				printTab(deepLevel-1, columnBit)
				fmt.Printf("├-%s\n", list[i].Name())
				scanDir(path.Join(currentDir, list[i].Name()), deepLevel+1, columnBit+1<<uint(deepLevel))
			}
		} else {
			if i+1 == dirNum {
				printTab(deepLevel, columnBit)
				fmt.Printf("%s\n", list[i].Name())
				printTab(deepLevel-1, columnBit)
				fmt.Println()
			} else {
				printTab(deepLevel, columnBit)
				fmt.Printf("%s\n", list[i].Name())
			}
		}
	}

	return nil

}

func printTab(deepLevel, columnBit int) {
	for i := 0; i < deepLevel; i++ {
		if columnBit&1 == 1 {
			fmt.Print(tabColumn)
		} else {
			fmt.Print(tabSpace)
		}
		columnBit >>= 1
	}
}

func printCurrentDir(dir string) error {

	file, err := os.Open(dir)
	defer file.Close()
	if err != nil {
		return err
	}
	fmt.Printf("%s\n", file.Name())

	return nil

}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
