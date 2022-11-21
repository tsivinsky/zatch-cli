# zatch-cli

Command-Line Interface for [zatch.ru api](https://zatch.ru).

## Install

### Linux

```bash
curl -L https://github.com/tsivinsky/zatch-cli/releases/latest/download/zatch_linux -o zatch && chmod +x zatch
```

### Windows

```bash
curl -L https://github.com/tsivinsky/zatch-cli/releases/latest/download/zatch_windows.exe -o zatch && chmod +x zatch
```

### MacOS (intel)

```bash
curl -L https://github.com/tsivinsky/zatch-cli/releases/latest/download/zatch_macos -o zatch && chmod +x zatch
```

### MacOS (arm)

```bash
curl -L https://github.com/tsivinsky/zatch-cli/releases/latest/download/zatch_macos_arm -o zatch && chmod +x zatch
```

## Usage

### Simple shorten

```bash
zatch https://tsivinsky.com
```

### Shorten with custom name

```bash
zatch --name tsivinsky https://tsivinsky.com
```

### Shorten with auto_delete option

```bash
zatch --auto-delete 30 https://tsivinsky.com
```
