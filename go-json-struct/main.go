package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
)

const (
	TUN_CLIENT_FNAME = "./client.json"
)

var CliTunID map[string]string

type ClientConfig struct {
	Version				string `json:"version"`
	ClientId			string `json:"client_id"`
	CliApp				string `json:"cli"`
	ClientTunIds		[] string `json:"client_tun_ids"`
}

func printConfig(cc *ClientConfig) {
	//fmt.Println(cc)
	fmt.Printf("%s\n", cc.Version)
	fmt.Printf("%s\n", cc.ClientId)
	fmt.Printf("%s\n", cc.CliApp)
	fmt.Printf("%v\n", cc.ClientTunIds)
}

func readConfig(fname string) (error, *ClientConfig) {

	cc := &ClientConfig{}

    if fname == "" {
        fmt.Printf("Error in file cannot be null: %s\n", fname)
        return errors.New("filename is null"), nil
    }

    fd, err := os.Open(fname)
    if err != nil {
        fmt.Printf("Error in opening file: %s\n", fname)
        return err, nil
    }
    defer fd.Close()

    data, err := ioutil.ReadAll(fd)
    if err != nil {
        return err, nil
    }

    json.Unmarshal(data, &cc)

    return nil, cc
}

func main() {
	fmt.Println("vim-go")
	CliTunID = make(map[string]string)
	err, conf := readConfig(TUN_CLIENT_FNAME)
	if err != nil {
		fmt.Printf("Error reading file: %s err: %v\n",
			TUN_CLIENT_FNAME, err)
		return;
	}
	for _, v := range conf.ClientTunIds {
		fmt.Printf("%s\n", v)
		CliTunID[v] = v
	}
	printConfig(conf)
	fmt.Println(CliTunID)

	if v, ok := CliTunID["1002"]; ok {
		fmt.Println(v)
	}
	return
}
