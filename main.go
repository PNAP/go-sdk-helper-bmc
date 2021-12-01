package main

import (
	"fmt"
	"github.com/PNAP/go-sdk-helper-bmc/command/bmcapi"
	"github.com/PNAP/go-sdk-helper-bmc/receiver"
	"github.com/PNAP/go-sdk-helper-bmc/dto"
	bmcapiclient "github.com/phoenixnap/go-sdk-bmc/bmcapi"
)




func main(){

	configuration := dto.Configuration{}
	configuration.UserAgent = "sdk-helper"
	sdk, err := receiver.NewBMCSDKWithDefaultConfig(configuration)

	if err!=nil{
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

	srv := bmcapiclient.ServerReserve{}
	//boolVar := true
	srv.PricingModel = "bla bla"
	cmd3 := bmcapi.NewReserveServerCommand(sdk,"61a6a8d968496e7abc7e8497", srv)

	s, errr := cmd3.Execute()
	

	if errr!= nil{
		fmt.Println("Error is", errr)
	}else{
		fmt.Println("Server is is", s.Hostname)
	}


}