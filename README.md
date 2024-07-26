# subnet-calc

## Usage

A simple, intuitive command-line tool for calculating IP subnets and related information.

### Installation

To install the tool, make sure you have Go installed on your machine. Then, run the following command:

```shell
go get github.com/barturba/subnet-calc
```

### Usage

Once the tool is installed, you can use it by running the following command:

```shell
subnet-calc <ip-address> <subnet-mask>
```

Replace `<ip-address>` with the IP address you want to calculate the subnet for, and `<subnet-mask>` with the subnet mask in CIDR notation.

### Example

Here's an example of how to use the tool:

```shell
subnet-calc 192.168.0.1 24
```

This will calculate the subnet for the IP address `192.168.0.1` with a subnet mask of `24`.

### Contributing

If you find any issues or have suggestions for improvement, feel free to open an issue or submit a pull request on [GitHub](https://github.com/barturba/subnet-calc).
