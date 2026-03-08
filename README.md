# jots

**jots** is a minimal CLI journaling tool

## Installation

### with Nix

### with Nix

Simply add the repo to your flake inputs...

```nix
inputs = {
  spyglass.url = "github:stikypiston/jots";
};
```

...and pass it into your `environment.systemPackages`...

```nix
environment.systemPackages = [
  inputs.jots.packages.${pkgs.stdenv.hostPlatform.system}.jots
];
```

### with Go

To install, simply run:

```shell
go install github.com/stikypiston/jots@latest
```

## Usage

## Inspirations

**jots** was mainly inspired by [jrnl](https://jrnl.sh)
