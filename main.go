package main

import (
	"fmt"

	"github.com/PNAP/go-sdk-helper-bmc/command/bmcapi/server"
	"github.com/PNAP/go-sdk-helper-bmc/command/networkapi/privatenetwork"
	"github.com/PNAP/go-sdk-helper-bmc/dto"
	"github.com/PNAP/go-sdk-helper-bmc/receiver"
)

func main() {

	configuration := dto.Configuration{}
	configuration.UserAgent = "sdk-helper"
	sdk, err := receiver.NewBMCSDKWithDefaultConfig(configuration)

	if err != nil {
		fmt.Println("Error is", err)
	}
	//cmd := bmcapi.NewGetServerCommand(sdk,"sdasdad")
	/* srv := bmcapiclient.ServerCreate{}
	srv.Hostname = "test"
	boolVar := true
	srv.InstallDefaultSshKeys = &boolVar
	srv.Location = "PHX"
	srv.Type =  "s1.c1.small"
	srv.Os = "ubuntu/bionic"
	cmd1 := bmcapi.NewCreateServerCommand(sdk, srv) */

	//cmd2 := bmcapi.NewGetServersCommand(sdk)

	// srv := bmcapiclient.ServerReserve{}
	//boolVar := true
	// srv.PricingModel = "bla bla"
	cmd3 := server.NewDeleteServerCommand(sdk, "61a6a8d968496e7abc7e8497")

	s, errr := cmd3.Execute()

	if errr != nil {
		fmt.Println("Error is", errr)
	} else {
		fmt.Println("Server is is", s.Result)
	}

	cmd4 := privatenetwork.NewGetPrivateNetworkCommand(sdk, "5fe997d274432c34c12adf8a")

	s4, errr4 := cmd4.Execute()

	if errr4 != nil {
		fmt.Println("Error is", errr4)
	} else {
		fmt.Println("Server is is", s4.Id)
	}

}
