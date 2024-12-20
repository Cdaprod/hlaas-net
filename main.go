package main

import (
    "encoding/csv"
    "fmt"
    "os"
    "strings"
)

type NetworkDevice struct {
    IPAddress     string
    Name          string
    IPv6Addresses string
    MAC           string
    Vendor        string
    Properties    string
    mDNSName      string
    LLMNRName     string
    NetBIOSName   string
    NetBIOSDomain string
    DNSName       string
}

func main() {
    devices := loadCSVData("net-analyzer/network_devices.csv")

    for {
        displayMenu()
        choice := getUserChoice()

        switch choice {
        case 1:
            displayAllDevices(devices)
        case 2:
            searchDevicesByIP(devices)
        case 3:
            searchDevicesByName(devices)
        case 4:
            displaySummary(devices)
        case 5:
            fmt.Println("Exiting...")
            return
        default:
            fmt.Println("Invalid choice. Please try again.")
        }

        fmt.Println()
    }
}

func loadCSVData(filePath string) []NetworkDevice {
    file, err := os.Open(filePath)
    if err != nil {
        fmt.Println("Error opening file:", err)
        os.Exit(1)
    }
    defer file.Close()

    reader := csv.NewReader(file)
    _, err = reader.Read() // Skip header row
    if err != nil {
        fmt.Println("Error reading header:", err)
        os.Exit(1)
    }

    records, err := reader.ReadAll()
    if err != nil {
        fmt.Println("Error reading records:", err)
        os.Exit(1)
    }

    var devices []NetworkDevice
    for _, record := range records {
        devices = append(devices, NetworkDevice{
            IPAddress:     record[0],
            Name:          record[1],
            IPv6Addresses: record[2],
            MAC:           record[3],
            Vendor:        record[4],
            Properties:    record[5],
            mDNSName:      record[6],
            LLMNRName:     record[7],
            NetBIOSName:   record[8],
            NetBIOSDomain: record[9],
            DNSName:       record[10],
        })
    }

    return devices
}

func displayMenu() {
    fmt.Println("Network Inventory Management System")
    fmt.Println("1. Display All Devices")
    fmt.Println("2. Search Devices by IP")
    fmt.Println("3. Search Devices by Name")
    fmt.Println("4. Display Summary")
    fmt.Println("5. Exit")
}

func getUserChoice() int {
    var choice int
    fmt.Print("Enter your choice: ")
    fmt.Scanln(&choice)
    return choice
}

func displayAllDevices(devices []NetworkDevice) {
    for _, device := range devices {
        fmt.Printf("IP: %s, Name: %s, Properties: %s\n", device.IPAddress, device.Name, device.Properties)
    }
}

func searchDevicesByIP(devices []NetworkDevice) {
    var searchIP string
    fmt.Print("Enter IP to search: ")
    fmt.Scanln(&searchIP)

    found := false
    for _, device := range devices {
        if device.IPAddress == searchIP {
            fmt.Printf("IP: %s, Name: %s, Properties: %s\n", device.IPAddress, device.Name, device.Properties)
            found = true
            break
        }
    }

    if !found {
        fmt.Printf("No device found with IP %s\n", searchIP)
    }
}

func searchDevicesByName(devices []NetworkDevice) {
    var searchName string
    fmt.Print("Enter name to search: ")
    fmt.Scanln(&searchName)

    found := false
    for _, device := range devices {
        if strings.Contains(strings.ToLower(device.Name), strings.ToLower(searchName)) {
            fmt.Printf("IP: %s, Name: %s, Properties: %s\n", device.IPAddress, device.Name, device.Properties)
            found = true
        }
    }

    if !found {
        fmt.Printf("No devices found containing name %s\n", searchName)
    }
}

func displaySummary(devices []NetworkDevice) {
    totalDevices := len(devices)
    pingableDevices := 0
    forwardedDevices := 0

    for _, device := range devices {
        if strings.Contains(device.Properties, "Pingable") {
            pingableDevices++
        }
        if device.IPv6Addresses != "" {
            forwardedDevices++
        }
    }

    fmt.Println("Summary:")
    fmt.Printf("Total Devices: %d\n", totalDevices)
    fmt.Printf("Pingable Devices: %d\n", pingableDevices)
    fmt.Printf("IPv6 Forwarded Devices: %d\n", forwardedDevices)
}