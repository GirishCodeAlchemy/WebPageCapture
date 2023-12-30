# WebPageCapture

Go-Based Solution to capture screenshots or generate PDFs of web pages using headless Chrome.

## Description

In the realm of web development and automation, capturing screenshots or generating PDFs of web pages is a common need. Whether for documentation, testing, or monitoring, having a tool that simplifies this process can be invaluable. In this blog post, we introduce a lightweight solution built with Go that leverages headless Chrome for capturing screenshots and generating PDFs.

## Features

- **Capture Screenshots**: Take high-quality screenshots of web pages.

- **Generate PDFs**: Create PDF documents for the web pages.

- **Help and Usage Information**: Access help and usage information using the `-h` or `--help` flag.

- **Custom Output Path and Filename**: Specify the output path and filename for captured screenshots or generated PDFs.

- **URL Validation**: Ensure valid URLs are provided for capturing.

- **Output Type Validation**: Validate and support output types, such as PDF or image.

- **File Extension Validation**: Validate output filenames for proper file extensions.

## **Getting Started**

To get started with the Web Page Capture Tool, you can download the pre-built binary from the [GitHub repository](https://github.com/GirishCodeAlchemy/WebPageCapture.git)

1. Download the latest binary from the [Releases](https://github.com/GirishCodeAlchemy/WebPageCapture/releases/download/v1.0.0/web-capture-tool)

   ```bash
   wget https://github.com/GirishCodeAlchemy/WebPageCapture/releases/download/v1.0.0/web-capture-tool
   ```

2. Once downloaded, make the binary executable:

   ```bash
   chmod +x web-capture-tool
   ```

3. Move the binary to a directory in your system's PATH, so you can run it from any location:

   ```bash
   sudo mv web-capture-tool /usr/local/bin/
   ```

## **Usage**

Certainly! Below is an example of how you can include the help command usage in the README markdown:

markdown
Copy code

### **Help and Usage:**

For detailed information about the available flags and options, you can use the following command:

```bash
web-capture-tool -h
```

### **Capture Screenshot:**

To capture a screenshot of a web page, run the following command:

```bash
./web-capture-tool <URL> image [filename]
```

Example:

```bash
./web-capture-tool https://example.com image output.jpg
```

### **Generate PDF:**

To generate a PDF of a web page, use the following command:

```bash
./web-capture-tool <URL> pdf [filename]
```

Example:

```bash
./web-capture-tool https://example.com pdf output.pdf
```

## **Customization and Extension**

The tool is intentionally kept minimal to serve as a foundation that can be extended based on specific requirements. Users familiar with Go can explore the code and customize the tool to suit their needs. Additionally, integrating it into existing automation pipelines or scripts is straightforward.

## **Conclusion**

The Web Page Capture Tool provides a quick and efficient way to capture web pages in both image and PDF formats. Its simplicity makes it suitable for various use cases, and the underlying Go and Chrome integration offers flexibility and reliability. Whether you need to document web pages, perform visual testing, job hunting, and content documentation, or automate PDF generation, this tool can be a valuable addition to your toolkit.
