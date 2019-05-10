/*
* Copyright (c) 2019, WSO2 Inc. (http://www.wso2.org) All Rights Reserved.
*
* WSO2 Inc. licenses this file to you under the Apache License,
* Version 2.0 (the "License"); you may not use this file except
* in compliance with the License.
* You may obtain a copy of the License at
*
*    http://www.apache.org/licenses/LICENSE-2.0
*
* Unless required by applicable law or agreed to in writing,
* software distributed under the License is distributed on an
* "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
* KIND, either express or implied. See the License for the
* specific language governing permissions and limitations
* under the License.
 */

package cmd

import (
    "os"
    "fmt"
    "strconv"
    "github.com/spf13/cobra"
    "github.com/wso2/micro-integrator/cmd/utils"
)

// Init command related usage info
const initCmdLiteral = "init"
const initCmdShortDesc = "Set Management API configuration"

var initCmdLongDesc = "Set the Hostname and the Port of the Management API\n"

// initCmd represents the init command
var initCmd = &cobra.Command{
    Use:   initCmdLiteral,
    Short: initCmdShortDesc,
    Long:  initCmdLongDesc,
    Run: func(cmd *cobra.Command, args []string) {
        utils.Logln(utils.LogPrefixInfo + "Init called")
        promptUserForConfig()
    },
}

func init() {
    rootCmd.AddCommand(initCmd)
}

func promptUserForConfig(){
    fmt.Println("Follow the instructions below to configure the CLI")
    fmt.Print("Enter Host name: ")
    var host string
    fmt.Scanln(&host)
    if len(host) == 0 {
        fmt.Println("Host name is not specified, default value is used")
        host = utils.DefaultHost
    }
    fmt.Print("Enter Port number: ")
    var strPort string
    fmt.Scanln(&strPort)
    if len(strPort) == 0 {
        fmt.Println("Port number is not specified, default value is used")
        strPort = utils.DefaultPort
    } else {
        port, err := strconv.Atoi(strPort)
        if err == nil {
            if (port < 1000) || (port > 10000) {
                fmt.Println("Port number is out of range. Please specify a port number between 1000 and 10000")
                os.Exit(1)
            }
            strPort = strconv.Itoa(int(port))
        } else {
            fmt.Println("Port number is invalid. Please specify a port number between 1000 and 10000")
            os.Exit(1)
        }
    }
    writeConfig( utils.HTTPProtocol + host, strPort)
}

func writeConfig(url, port string){
    serverConfig := utils.ServerConfig{Url:url, Port:port}
    utils.WriteServerConfigFile(serverConfig, utils.ServerConfigFilePath)
    fmt.Println("CLI configuration is successful")
}