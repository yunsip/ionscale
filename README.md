# ionscale

> **Note**:
> ionscale is currently beta quality, actively being developed and so subject to changes

**What is Tailscale?**

[Tailscale](https://tailscale.com) is a VPN service that makes the devices and applications you own accessible anywhere in the world, securely and effortlessly. 
It enables encrypted point-to-point connections using the open source [WireGuard](https://www.wireguard.com/) protocol, which means only devices on your private network can communicate with each other.

**What is ionscale?**

While the Tailscale software running on each node is open source, their centralized "coordination server" which act as a shared drop box for public keys is not.

_ionscale_ aims to implement such lightweight, open source alternative Tailscale control server.

⭐ If you find _ionscale_ useful, please consider giving it a star on GitHub, or [drop a note](https://github.com/jsiebens/ionscale/discussions/new?category=show-and-tell) on how you are using _ionscale_.

## Features

- multi [tailnet](https://tailscale.com/kb/1136/tailnet/) support
- multi user support
- OIDC integration (not required, although recommended)
- [Auth keys](https://tailscale.com/kb/1085/auth-keys/)
- [Access control list](https://tailscale.com/kb/1018/acls/) (with support for `autogroups`, `tagOwners`, `autoApprovers`, `nodeAttrs`, `grants` ...)  
- [DNS](https://tailscale.com/kb/1054/dns/)
  - nameservers
  - Split DNS
  - MagicDNS
- [Subnet routers](https://tailscale.com/kb/1019/subnets) and [Exit Nodes](https://tailscale.com/kb/1103/exit-nodes)
- [HTTPS Certs](https://tailscale.com/kb/1153/enabling-https/)
- [Tailscale SSH](https://tailscale.com/kb/1193/tailscale-ssh/) 
- [Tailscale Serve](https://tailscale.com/kb/1312/serve)
- [Service collection](https://tailscale.com/kb/1100/services/)
- [Taildrop](https://tailscale.com/kb/1106/taildrop/)

## Documentation

Some documentation can be found [here](https://jsiebens.github.io/ionscale)

## Disclaimer

This is not an official Tailscale or Tailscale Inc. project.