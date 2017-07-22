package main

import (
	"bufio"
	"flag"
	"log"
	"os"

	"github.com/venkatramachandran/robot/commands"
	"github.com/venkatramachandran/robot/objects"
)

func main() {
	var fileName string
	flag.StringVar(&fileName, "input", "commands.txt", "Input Commands file")
	flag.Parse()

	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	robot := objects.NewRobot(5, 5)

	for scanner.Scan() {
		command, err := commands.From(scanner.Text())
		if err != nil {
			log.Printf("Unable to understand command : %s\n", err)
			continue
		}
		command.Execute(robot)
	}
}
