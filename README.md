# scc

SuperCollider CLI

## Install

```
go get -u github.com/scgolang/scc
```

## Usage

### scsynth

Install and start a SuperCollider server.

```
scsynth -u 57120
```

Note that `scc` uses 57120 as the default port to connect to scsynth. This can be changed with the `-scsynth` flag.

For example, `-scsynth 127.0.0.1:57110`.

### synthdefs

scc ships with some synthdefs out of the box.

To use them you must send them to `scsynth` first.

```
scc senddefs
```

### sine

Create a 440kHz sine tone.

```
scc synth -def sine -id 1000
```

Turn it off!

```
scc nfree 1000
```
