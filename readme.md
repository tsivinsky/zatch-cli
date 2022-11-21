# zatch-cli

Command-Line Interface for [zatch.ru api](https://zatch.ru).

## Install

### with Go cli

```bash
go install github.com/tsivinsky/zatch-cli@latest
```

### with Git

```bash
git clone https://github.com/tsivinsky/zatch-cli && cd zatch-cli && go build -o zatch && chmod +x zatch && sudo cp zatch /usr/bin/zatch
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
