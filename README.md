This application provides a simple menu-based interface for interacting with the network device data. It includes the following features:

- Loading the CSV data from a file named "network_devices.csv" in the current directory.
- Displaying all devices with their IP address, name, and properties.
- Searching for a device by IP address.
- Searching for devices by name (case-insensitive partial match).
- Displaying a summary of the devices, including the total count, pingable devices, and IPv6 forwarded devices.
- The application uses a NetworkDevice struct to represent each device and stores the loaded devices in a slice. The loadCSVData function reads the CSV file and populates the slice of NetworkDevice structs.

The main loop of the application displays the menu, prompts the user for their choice, and calls the appropriate function based on the user's input.