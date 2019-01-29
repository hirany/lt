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
		if len(c.Args()) < 1 {
			checkError(lt("./"))
		} else {
			checkError(lt(string(c.Args()[0])))
		}
	}

	app.Run(os.Args)

}

func lt(dirPath string) error {

	err := printCurrentDir(dirPath)
	if err != nil {
		return err
	}
	err = scanDir(dirPath, 1, 1)
	if err != nil {
		return err
	}
	fmt.Println()
	return nil

}

func scanDir(currentDir string, deepLevel, columnBit int) error {

	list, err := ioutil.ReadDir(currentDir)
	if err != nil {
		return err
	}
	dirNum := len(list)

	for i := 0; i < dirNum; i++ {
		if list[i].Name()[0] == '.' {
			continue
		}
		if !list[i].IsDir() {
			printTab(deepLevel-1, columnBit)
			fmt.Printf("| %s\n", list[i].Name())
		}
		if list[i].IsDir() {
			printTab(deepLevel, columnBit)
			fmt.Println()
			printTab(deepLevel-1, columnBit)
			if i+1 == dirNum {
				fmt.Printf("└-%s\n", list[i].Name())
				scanDir(path.Join(currentDir, list[i].Name()), deepLevel+1, columnBit+1<<uint(deepLevel)-1<<uint(deepLevel-1))
			} else {
				fmt.Printf("├-%s\n", list[i].Name())
				scanDir(path.Join(currentDir, list[i].Name()), deepLevel+1, columnBit+1<<uint(deepLevel))
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
