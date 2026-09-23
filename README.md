# sicpctl

A [Cobra](https://cobra.dev)-based command-line client for the **Philips
SICP** digital-signage protocol. Speaks SICP over TCP (default port 5000);
built on top of [`github.com/go-sicp/sicp`](https://github.com/go-sicp/sicp).

## Install

```sh
go install github.com/go-sicp/sicpctl@latest
```

## Usage

```sh
sicpctl --host 192.168.1.50 power on
sicpctl --host 192.168.1.50 power
sicpctl --host 192.168.1.50 source hdmi
sicpctl --host 192.168.1.50 inputs                     # list sources the model supports (V2.05+)
sicpctl --host 192.168.1.50 volume 30
sicpctl --host 192.168.1.50 mute on                    # V2.00+
sicpctl --host 192.168.1.50 backlight off              # V2.02+, blank screen, system stays up
sicpctl --host 192.168.1.50 video                      # report all 7 params
sicpctl --host 192.168.1.50 video 60 50 70 5 50 50 3   # set them
sicpctl --host 192.168.1.50 info
sicpctl --host 192.168.1.50 restart android            # V2.02+
sicpctl --host 192.168.1.50 raw 0x19                   # arbitrary command bytes
```

`sicpctl --help` and `sicpctl <cmd> --help` print full usage. Cobra also
provides a `completion` subcommand for shell autocompletion.

### Persistent flags

| Flag | Default | Description |
|---|---|---|
| `--host HOST[:PORT]` | (required) | Display address; port defaults to 5000 |
| `--monitor N` | 1 | Monitor ID 1..255. Use 0 for broadcast (no reply read) |
| `--group N` | 0 | Group ID 0..254 (0 = address by monitor ID) |
| `--timeout DURATION` | 3s | TCP I/O timeout |

## Layout

Each subcommand lives in its own file under [`cmd/`](cmd):

```
cmd/
├── root.go       # root command, persistent flags, newClient() helper
├── util.go       # parseByte (decimal or 0xHH)
├── power.go      # sicpctl power
├── source.go     # sicpctl source
├── inputs.go     # sicpctl inputs       (V2.05+)
├── volume.go     # sicpctl volume
├── mute.go       # sicpctl mute         (V2.00+)
├── backlight.go  # sicpctl backlight    (V2.02+)
├── video.go      # sicpctl video
├── info.go       # sicpctl info
├── restart.go    # sicpctl restart      (V2.02+)
└── raw.go        # sicpctl raw
```

`main.go` is a one-liner that calls `cmd.Execute()`. To add a new subcommand,
drop a new file in `cmd/`, declare a `cobra.Command`, and `rootCmd.AddCommand`
it from an `init()` function.

## Build from source

```sh
git clone https://github.com/go-sicp/sicpctl
cd sicpctl
go build .
```
