# Secure Code Analyzer

**Secure Code Analyzer** is a modular tool for static code analysis, designed to identify security vulnerabilities in Go and Python projects. The system supports automatic vulnerability detection and generates comprehensive reports for further evaluation.

---

## **Features**
- **Multi-language Support**: Analyze code written in Go and Python.
- **Customizable Reports**: Generates detailed security reports in text format.
- **Modular Architecture**: Easily extendable for additional languages and analysis techniques.

---

## **Installation**
### **Requirements**
- Go (version 1.16 or later)
- Python 3.x
- Installed tools:
    - [gosec](https://github.com/securego/gosec) for Go analysis
    - [bandit](https://github.com/PyCQA/bandit) for Python analysis

### **Setup**
1. Clone the repository:
   ```bash
   git clone https://github.com/your-username/secure-code-analyzer.git
   cd secure-code-analyzer
   ```

2. Install dependencies:
    - Install `gosec`:
      ```bash
      go install github.com/securego/gosec/v2/cmd/gosec@latest
      ```
    - Install `bandit`:
      ```bash
      pip install bandit
      ```

---

## **Usage**
### **Run the Analyzer**
To analyze Go and Python projects, use the following commands:

**Go Analysis**:
   ```bash
   go run main.go ./examples
   ```

---

## **Sample Reports**
### **Go Analysis Report**
```
File: examples/go/vulnerable.go
Line: 12
Issue: Hardcoded credentials detected.
Severity: Medium
```

### **Python Analysis Report**
```
File: examples/python/vulnerable.py
Line: 23
Issue: Use of eval() is insecure and should be avoided.
Severity: High
```

---

## **Development**
The tool uses a modular architecture that makes it easy to extend:
- **Analyzer Module**: Contains logic to run analysis for specific languages.
- **CLI Interface**: Provides a command-line interface for interaction.

### **Adding Support for New Languages**
1. Implement a new analyzer in the `Analyzer` struct.
2. Define how the analysis tool for the language should be executed.
3. Add the language to the `RunAnalysis` function.

---

## **Contributing**
Contributions are welcome! If you want to report a bug, suggest a feature, or contribute code, please open an issue or a pull request.

---

## **License**
This project is licensed under the MIT License. See the [LICENSE](LICENSE.txt) file for details.
