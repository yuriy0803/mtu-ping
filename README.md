# MTU and Ping Finder

This Go program determines the **optimal MTU (Maximum Transmission Unit)** and the **best ping time** for a given host.  
It supports both **Windows** and **Linux**.

## Features
- Detects the best **MTU** using a **binary search method**.
- Measures the **best ping time** based on multiple responses.
- Compatible with **Windows** and **Linux**.

## Prerequisites
- Install **Go**: [golang.org](https://go.dev/dl/)
- The program relies on the built-in `ping` command for each operating system:
  - **Windows**: Uses `ping -f -l`
  - **Linux**: Uses `ping -M do -s`

## Installation
Clone the repository:

```sh
git clone https://github.com/yuriy0803/mtu-ping.git
cd mtu-ping
```

## Usage
Run the program:

```sh
go run main.go
```

Or build and execute it:

```sh
go build -o mtu-ping
./mtu-ping
```

## Example Output
```yaml
Best MTU for 8.8.8.8: 1444
Best Ping for 8.8.8.8: 30.00 ms
```

## OS Compatibility
| Feature        | Windows | Linux |
|---------------|---------|-------|
| MTU Detection | ✅ Yes  | ✅ Yes |
| Ping Test     | ✅ Yes  | ✅ Yes |

## MTU Change on Windows (CMD)
To change the MTU in Windows, use the following command in **CMD (as administrator)**:

```sh
netsh interface ipv4 set subinterface "<Interface Name>" mtu=<MTU Value> store=persistent
```

Example (for an Ethernet connection with an MTU of **1444**):

```sh
netsh interface ipv4 set subinterface "Ethernet" mtu=1444 store=persistent
```

To check your network interfaces and their MTU values:

```sh
netsh interface ipv4 show interfaces
```

## MTU Change on Linux
To change the MTU on Linux, use:

```sh
sudo ip link set <interface> mtu <MTU Value>
```

Example (for an Ethernet connection `eth0` with an MTU of **1444**):

```sh
sudo ip link set eth0 mtu 1444
```

To check your current MTU settings:

```sh
ip link show
```

## Notes
- On **Linux**, you may need **root privileges** to run `ping` with the required flags.
- If you receive errors, ensure your **firewall allows ICMP** requests.

## License
This project is open-source under the **MIT License**.

---

Feel free to contribute and improve the project! 🚀

## Adding to GitHub
To add this README to your repository, follow these steps:

```sh
cd mtu-ping
git add README.md
git commit -m "Add README with Windows & Linux MTU change instructions"
git push origin main
```

