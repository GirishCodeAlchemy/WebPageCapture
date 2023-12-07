# WebPageCapture

## Description

This is a simple tool written in Go that allows users to capture either a screenshot or a PDF of a web page using headless Chrome.

## Usage

### To Capture as a Screenshot:

```bash
go run main.go <URL> image [filename]

```

#### Example

```bash
go run main.go https://example.com image
```

or

```bash
go run main.go https://example.com image output.jpg
```

If `filename` is not provided, it will use the default filename `output.jpg`

### To Generate PDF

```bash
go run main.go <URL> pdf [filename]
```

#### Example

```bash
go run main.go https://example.com pdf
```

or

```bash
go run main.go https://example.com pdf output.pdf
```

If `filename` is not provided, it will use the default filename `output.pdf`.

## Building the Binary

To build the binary, use the following commands:

```bash
go build -o web-capture-tool main.go
```

This will generate an executable file named web-capture-tool. You can then run it as follows:

```bash
./web-capture-tool <URL> <pdf|image> [filename]
```
