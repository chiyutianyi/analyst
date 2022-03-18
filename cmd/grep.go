package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strconv"

	"github.com/spf13/cobra"
)

type grepCmd struct {
}

func (cmd *grepCmd) Run(_ *cobra.Command, args []string) {
	if len(args) < 4 {
		fmt.Println("Usage: analyst grep <jsonfile> <key> <cmp> <value>")
		os.Exit(1)
	}
	file := args[0]
	key := args[1]
	cmp := args[2]
	value, err := strconv.ParseFloat(args[3], 64)
	if err != nil {
		fmt.Fprintf(os.Stderr, "invalid value: %v", err)
		return
	}
	f, err := os.Open(file)
	if err != nil {
		fmt.Fprintf(os.Stderr, "can not open %v: %v\n", file, err)
		return
	}
	buf := bufio.NewReader(f)
	var data = map[string]interface{}{}
	for {
		line, _, err := buf.ReadLine()
		if err != nil {
			if err == io.EOF {
				return
			}
			fmt.Fprintf(os.Stderr, "read line error: %v\n", err)
			return
		}
		err = json.Unmarshal(line, &data)
		if err != nil {
			fmt.Fprintf(os.Stderr, "unmarshal line err: %v\nline: %v\n", err, line)
			continue
		}
		if data[key] == nil {
			continue
		}
		switch cmp {
		case "=":
			if data[key].(string) == args[3] {
				fmt.Println(string(line))
			}
		case "!=":
			if data[key].(string) != args[3] {
				fmt.Println(string(line))
			}
		case ">":
			if data[key].(float64) > value {
				fmt.Println(string(line))
			}
		case ">=":
			if data[key].(float64) >= value {
				fmt.Println(string(line))
			}
		case "<":
			if data[key].(float64) < value {
				fmt.Println(string(line))
			}
		case "<=":
			if data[key].(float64) <= value {
				fmt.Println(string(line))
			}
		}
	}
}

func init() {
	readPack := &grepCmd{}

	cmd := &cobra.Command{
		Use:   "grep",
		Short: "grep a given json file",
		Run:   readPack.Run,
	}

	Cmd.AddCommand(cmd)
}
