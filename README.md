# Hetzner Certbot DNS Validation Hooks

This project provides [Pre- and Post-Validation Hooks](https://eff-certbot.readthedocs.io/en/stable/using.html#pre-and-post-validation-hooks) for [Certbot](https://certbot.eff.org) for [Hetzner DNS](https://dns.hetzner.com/) users to use to fulfil DNS challenges.

## Compilation

Compilation requires at least [Go](https://go.dev/dl/) 1.19.

```bash
git clone https://github.com/hitecherik/hetzner-certbot-dns.git
cd hetzner-certbot-dns
make
```

## Usage

This project provides two binaries:

| Binary          | Purpose              |
| --------------- | -------------------- |
| `authenticator` | Pre-validation hook  |
| `cleanup`       | Post-validation hook |

The Hetzner DNS API key is passed to the binaries using the `HETZNER_API_KEY` environment variable.

For details on specifying these binaries as the pre- and post-validation hooks, see [Certbot's documentation](https://eff-certbot.readthedocs.io/en/stable/using.html#pre-and-post-validation-hooks).

## License

Copyright &copy; 2022 Alexander Nielsen. Licensed under the [MIT License](LICENSE.md).