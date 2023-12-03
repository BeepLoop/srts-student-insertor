package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	filePtr := flag.String("i", "", "file to read from")
	failsPrt := flag.String("o", "fails.txt", "file to write failed inserts into")
	dbPtr := flag.String("db", "", "database Data Source Name (DSN): example root:password@tcp/dbName")
	limitPtr := flag.Bool("limit", false, "limit the number of lines to read")
	maxLimitPtr := flag.Int("maxLimit", 0, "number of lines to read")
	programPtr := flag.String("p", "", "program to insert students into")

	programId := -1

	flag.Parse()

	err := CheckFlags(limitPtr, maxLimitPtr, filePtr, dbPtr)
	if err != nil {
		fmt.Println(err.Error())
		flag.PrintDefaults()
		os.Exit(1)
	}

	err = InitStore(*dbPtr)
	if err != nil {
		panic(err)
	}

	fmt.Println("Reading from file: ", *filePtr)
	if programPtr != nil {
		fmt.Println("Inserting students into program: ", *programPtr)
	}

	if *programPtr != "" {
		err := DB_Conn.Get(&programId, "SELECT id FROM Program WHERE program = ?", *programPtr)
		if err != nil {
			panic(err)
		}
	}

	file, err := os.Open(*filePtr)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	reader := bufio.NewReader(file)

	logFile, err := InitLogFile(*failsPrt)
	if err != nil {
		panic(err)
	}
	defer logFile.Close()

	inserted := 0
	var line string
	maxLines := *maxLimitPtr

	for {
		line, err = reader.ReadString('\n')
		if err != nil {
			break
		}

		r := csv.NewReader(strings.NewReader(string(line)))
		record, err := r.ReadAll()
		if err != nil {
			panic(err)
		}

		for _, value := range record {
			fmt.Println("inserting: ", value)

			if programId == -1 {
				err := InsertToDB(value)
				if err != nil {
					fmt.Println("Failed to insert: ", value)

					err := WriteToFile(logFile, value, err)
					if err != nil {
						fmt.Println("Failed to write to file: ", err)
					}
				} else {
					inserted++
				}
			} else {
				err := InsertToDBWithProgram(value, programId)
				if err != nil {
					err := WriteToFile(logFile, value, err)
					if err != nil {
						fmt.Println("Failed to write to file: ", err)
					}
				} else {
					inserted++
				}
			}
		}

		if *limitPtr {
			if maxLines <= 0 {
				fmt.Println("Reached maxLimitPtr of lines to read")
				break
			}
			maxLines--
		}
	}

	fmt.Println("Done!")
	fmt.Println("Students inserted: ", inserted)
}
