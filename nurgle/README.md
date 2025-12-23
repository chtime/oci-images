# Nurgle

A chaos engineering tool for testing container and environment limits.

## Commands

- `nurgle eat memory <amount>` — Allocates and holds the specified amount of memory (e.g., `50Mi`, `10G`)
- `nurgle eat cpu <n>` — Consumes CPU cycles across `n` parallel processes
- `nurgle poison <delay>` — Crashes the process after `<delay>` seconds

## Usage

```sh
nurgle eat memory 512Mi
nurgle eat cpu 4
nurgle poison 30
```

Use `--verbose` for debug output.