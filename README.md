# CookieJar 🍪🔒

**CookieJar** is an open-source Golang tool designed to analyze web cookies for security vulnerabilities. It examines cookie attributes to ensure they adhere to best security practices, helping developers and security professionals secure their web applications against common threats like Cross-Site Scripting (XSS) and Cross-Site Request Forgery (CSRF) attacks.

---

## Table of Contents

- [CookieJar 🍪🔒](#cookiejar-)
  - [Table of Contents](#table-of-contents)
  - [Features](#features)
  - [Installation](#installation)
    - [**1. Clone the Repository**](#1-clone-the-repository)
    - [**2. Build the Environment**](#2-build-the-environment)
  - [Usage](#usage)
  - [Specify Output Format](#specify-output-format)
  - [Use a Configuration File](#use-a-configuration-file)
  - [Configuration](#configuration)
  - [Configuration Options](#configuration-options)
  - [Examples](#examples)
  - [Contributing](#contributing)
      - [1. **Fork the repository.**](#1-fork-the-repository)
      - [2. Clone your forked repository.](#2-clone-your-forked-repository)
      - [3. Create a new branch.](#3-create-a-new-branch)
      - [4. Make your changes and commit them.](#4-make-your-changes-and-commit-them)
      - [5. Push to your branch.](#5-push-to-your-branch)
      - [6. Create a Pull Request.](#6-create-a-pull-request)
  - [License](#license)
  - [Disclaimer](#disclaimer)
  - [Contact](#contact)
  - [Acknowledgments](#acknowledgments)
  - [Development Roadmap](#development-roadmap)
  - [Running Tests](#running-tests)
  - [Additional Resources](#additional-resources)
  - [Frequently Asked Questions (FAQ)](#frequently-asked-questions-faq)

---

## Features

- **Cookie Attribute Analysis**: Checks for essential security attributes like `HttpOnly`, `Secure`, and `SameSite`.
- **Domain and Path Scope Verification**: Ensures cookies are scoped appropriately to minimize exposure.
- **Expiration Analysis**: Identifies session and persistent cookies, evaluating their expiration policies.
- **Third-Party Cookie Detection**: Detects cookies set by third-party domains to assess privacy risks.
- **Size Evaluation**: Warns if cookies exceed recommended size limits, which can impact performance.
- **Weak Encryption Detection**: Flags cookies storing sensitive data in plaintext or using weak encryption.
- **Detailed Reporting**: Generates comprehensive reports with actionable recommendations.
- **Multiple Input Methods**: Supports URL scanning and HTTP request files.
- **Flexible Output Formats**: Outputs results in text, JSON, or CSV formats for easy integration.
- **Command-Line Interface**: User-friendly CLI with various flags and options.
- **Concurrency Support**: Scans multiple cookies concurrently for faster analysis.
- **Integration Ready**: Can be integrated into CI/CD pipelines for continuous security assessments.

---

## Installation

To install **CookieJar**, you need to have [Go](https://golang.org/dl/) installed (version 1.16 or higher is recommended).

### **1. Clone the Repository**

```bash
git clone https://github.com/BBennett92/cookiejar.git
cd cookiejar
```
### **2. Build the Environment**

```bash
go build -o bin/cookiejar ./cmd/cookiejar/main.go
```

This will compile the application and place the executable in the bin/ directory.

## Usage

**Basic Scan**

Scan a website's cookies using the default settings.

```bash
./bin/cookiejar -url https://example.com
```

## Specify Output Format

**Choose the output format: text, json, or csv.**

```bash
./bin/cookiejar -url https://example.com -output json
```

## Use a Configuration File

**Specify a custom configuration file.**

```bash
./bin/cookiejar -url https://example.com -config ./configs/default_config.yaml
```
## Configuration

CookieJar uses a configuration file in YAML format to customize scanning options. If no configuration file is specified, it uses the default settings from configs/default_config.yaml.

**Default Configuration (configs/default_config.yaml)**

```yaml
# Default configuration for CookieJar
scan:
  follow_redirects: true
  timeout: "10s"
output:
  verbose: false
  format: text
```

## Configuration Options

- **scan.follow_redirects:** Whether to follow HTTP redirects (**true** or **false**).
- **scan.timeout:** Timeout duration for HTTP requests (e.g., "10s", "30s").
- **output.verbose:** Enable verbose logging (**true** or **false**).
- **output.format:** Default output format (**text**, **json**, **csv**).

## Examples

**Example 1: Scan with Detailed Text Output**

```bash
./bin/cookiejar -url https://example.com -output text
```

**Sample Output:**

```yaml
Cookie Name: session_id
- HttpOnly: false
- Secure: true
- SameSite: Lax
- Expiration: Wed, 31 Dec 2024 23:59:59 UTC
Issues:
  - HttpOnly flag is not set.
Recommendations:
  - Add HttpOnly flag to prevent client-side scripts from accessing the cookie.
```

**Example 2: Scan and Output in CSV Format**

```bash
./bin/cookiejar -url https://example.com -output csv > results.csv
```

This command scans the website and saves the results in CSV format to **results.csv**.

## Contributing

Contributions are welcome! Please follow these steps:

#### 1. **Fork the repository.**

Click the "Fork" button at the top right of the repository page.

#### 2. Clone your forked repository.

```bash
git clone https://github.com/yourusername/cookiejar.git
```

#### 3. Create a new branch.

```bash
git checkout -b feature/new-feature
```

#### 4. Make your changes and commit them.

```bash
git commit -am 'Add a new feature'
```

#### 5. Push to your branch.

```bash
git push origin feature/new-feature
```

#### 6. Create a Pull Request.

Go to the original repository and click on "Pull Requests", then "New Pull Request".

Please ensure that your code adheres to the existing style and that all tests pass.

## License

This project is licensed under the MIT License.

## Disclaimer

CookieJar is intended for educational and ethical testing purposes only. Always obtain proper authorization before scanning any website or application. The developers are not responsible for any misuse of this tool.

## Contact

For any questions or inquiries, please contact:

- GitHub: [Brandon Bennett](https://github.com/BBennett92)
- LinkedIn: [Brandon Bennett](https://linkedin.com/in/brandon-bennett~)
- Email: infosec.brandon@protonmail.com

## Acknowledgments

- Inspired by the need for better cookie security analysis tools.
- Built with love using Golang.

## Development Roadmap

Planned features for future releases:

- **Size Analysis:** Check for cookies exceeding recommended size limits.
- **Encryption Detection:** Identify cookies storing sensitive data without proper encryption.
- **Batch URL Scanning:** Support scanning multiple URLs from a file.
- **Integration with CI/CD Pipelines:** Provide scripts and documentation for integration.
- **Enhanced Reporting:** Generate HTML reports with visual representations.

## Running Tests

To run the unit tests, execute:

```bash
go test ./...
```

Ensure that all tests pass before committing your changes.

## Additional Resources

- [Go Modules Reference](https://golang.org/ref/mod)
- [YAML Format Specification](https://yaml.org/spec/1.2/spec.html)

## Frequently Asked Questions (FAQ)

1. **Does CookieJar support HTTPS websites?**

- Yes, CookieJar supports scanning both HTTP and HTTPS websites.

2. **Can I integrate CookieJar into my existing security tools?**

- CookieJar provides output in multiple formats (JSON, CSV), making it easy to integrate with other tools.

3. **How do I add new features or customize CookieJar?**

- You can modify the source code to add new features. Please consider contributing back to the project via Pull Requests.