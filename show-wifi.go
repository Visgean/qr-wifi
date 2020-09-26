package main

import (
	"fmt"
	"github.com/fumiyas/go-tty"
	qrc "github.com/fumiyas/qrc/lib"
	"github.com/mattn/go-colorable"
	"github.com/qpliu/qrencode-go/qrencode"
	"gopkg.in/ini.v1"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func parse_wifi_config(name string) (string, string) {
	cfg, err := ini.Load(name)
	if err != nil {
		log.Fatal("Fail to read file: %v", err)
	}

	ssid := cfg.Section("connection").Key("id").String()
	security_type := cfg.Section("wifi-security").Key("key-mgmt").String()
	if security_type == "wpa-psk" {
		security_type = "WPA"
	}
	password := cfg.Section("wifi-security").Key("psk").String()

	return ssid, fmt.Sprintf("WIFI:S:%s;T:%s;P:%s;;", ssid, security_type, password)

}

func print_qr(text string) {
	grid, err := qrencode.Encode(text, qrencode.ECLevelL)
	if err != nil {
		log.Fatal("encode failed: %v\n", err)

	}
	da1, err := tty.GetDeviceAttributes1(os.Stdout)
	if err == nil && da1[tty.DA1_SIXEL] {
		qrc.PrintSixel(os.Stdout, grid, false)
	} else {
		stdout := colorable.NewColorableStdout()
		qrc.PrintAA(stdout, grid, false)
	}
}

func match_with_args(filename string) bool {
	argsWithoutProg := os.Args[1:]
	if len(argsWithoutProg) <= 0 {
		return true
	}

	for _, arg := range argsWithoutProg {
		if strings.Contains(strings.ToLower(filename), strings.ToLower(arg)) {
			return true
		}
	}
	return false
}

func main() {
	files, err := filepath.Glob("/etc/NetworkManager/system-connections/*")
	if err != nil {
		log.Fatal(err)
	}
	for _, file := range files {
		if !match_with_args(file) {
			continue
		}

		ssid, passwd_string := parse_wifi_config(file)
		fmt.Println(ssid, "\n")
		//fmt.Println(passwd_string)
		print_qr(passwd_string)
		println("\n\n")
	}
}
