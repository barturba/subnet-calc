# subnet-calc

A command-line IP subnet calculator that provides network information given an IP address and subnet mask.

## Installation

```bash
go install github.com/barturba/subnet-calc@latest
```

## Usage

```bash
subnet-calc <ip-address> <cidr>
```

### Example

```bash
$ subnet-calc 192.168.0.1 /24
Network address: 192.168.0.0
Usable host IP range: 192.168.0.1 - 192.168.0.254
Broadcast address: 192.168.0.255
Total number of hosts: 256
Number of usable hosts: 254
Subnet mask: 255.255.255.0
Wildcard mask: 0.0.0.255
```

## Development

### Build

```bash
make build
```

### Run tests

```bash
make test
```

### Run tests with coverage

```bash
make coverage
```

## License

MIT
