package main

import (
	"flag"
	"fmt"
	"log"
	"os/exec"
	"runtime"
	"strings"
)

func main() {

	applicationPath := flag.String("applicationPath", "", "applicationPath")
	project := flag.String("project", "", "project")
	data := flag.String("data", "", "data")
	openingMode := flag.String("openingMode", "", "openingMode")
	createData := flag.Bool("createData", false, "createData")
	userParam := flag.String("userParam", "", "userParam")
	headless := flag.Bool("headless", false, "headless")
	dataless := flag.Bool("dataless", false, "dataless")
	webadminSettingsFile := flag.String("webadminSettingsFile", "", "webadminSettingsFile")
	webadminAccessKey := flag.String("webadminAccessKey", "", "webadminAccessKey")
	webadminAutoStart := flag.Bool("webadminAutoStart", false, "webadminAutoStart")
	webadminStoreSettings := flag.Bool("webadminStoreSettings", false, "webadminStoreSettings")
	utility := flag.Bool("utility", false, "utility")
	skipOnstartup := flag.Bool("skipOnstartup", false, "skipOnstartup")
	startupMethod := flag.String("startupMethod", "", "startupMethod")
	/*
	   https://developer.4d.com/docs/Admin/cli/
	*/

	flag.Parse()

	var args []string
	if *project != "" {
		args = append(args, "--project="+(*project))
	}
	if *data != "" {
		args = append(args, "--data="+(*data))
	}
	if *headless {
		args = append(args, "--headless")
	}
	if *dataless {
		args = append(args, "--headless")
	}
	if *skipOnstartup {
		args = append(args, "--skip-onstartup")
	}
	if *startupMethod != "" {
		args = append(args, "--startup-method="+(*startupMethod))
	}
	if *utility {
		args = append(args, "--utility")
	}
	if *createData {
		args = append(args, "--create-data")
	}
	if *userParam != "" {
		args = append(args, "--user-param="+(*userParam))
	}
	if *openingMode != "" {
		args = append(args, "--opening-mode="+(*openingMode))
	}
	if *webadminAutoStart {
		args = append(args, "--webadmin-auto-start")
	}
	if *webadminStoreSettings {
		args = append(args, "--webadmin-store-settings")
	}
	if *webadminSettingsFile != "" {
		args = append(args, "--webadmin-settings-file="+(*webadminSettingsFile))
	}
	if *webadminAccessKey != "" {
		args = append(args, "--webadmin-access-key="+(*webadminAccessKey))
	}

	var cmd *exec.Cmd

	switch runtime.GOOS {
	case "darwin":
		cmd = exec.Command("open", "-a", *applicationPath, "--args", strings.Join(args, " "))
	case "windows":
		cmd = exec.Command("cmd.exe", "/c", "start", *applicationPath, strings.Join(args, " "))
	}

	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", out)

}
