const fs = require("fs");
const { exec } = require("child_process");
const http = require("http");
const crypto = require("crypto");
const mysql = require("mysql");

// Hardcoded credentials vulnerability
function hardcodedCredentials() {
    const username = "admin"; // Hardcoded username
    const password = "supersecret"; // Hardcoded password
    console.log(`Logging in with username: ${username} and password: ${password}`);
}

// SQL Injection vulnerability
function sqlInjection(userInput) {
    const connection = mysql.createConnection({
        host: "localhost",
        user: "root",
        password: "password",
        database: "test",
    });

    connection.connect();

    // Vulnerable to SQL Injection
    const query = `SELECT * FROM users WHERE username = '${userInput}'`;
    connection.query(query, (error, results) => {
        if (error) throw error;
        console.log("User data:", results);
    });

    connection.end();
}

// Weak cryptographic hash function
function weakHash(data) {
    const hash = crypto.createHash("md5").update(data).digest("hex"); // Weak MD5 hash
    console.log(`MD5 hash of ${data}: ${hash}`);
}

// Command injection vulnerability
function commandInjection(userInput) {
    const command = `echo ${userInput}`; // Vulnerable to command injection
    exec(command, (error, stdout, stderr) => {
        if (error) {
            console.error(`Error: ${error.message}`);
            return;
        }
        if (stderr) {
            console.error(`Stderr: ${stderr}`);
            return;
        }
        console.log(`Output: ${stdout}`);
    });
}

// Insecure HTTP server
function insecureHTTPServer() {
    http
        .createServer((req, res) => {
            if (req.url === "/data") {
                res.writeHead(200, { "Content-Type": "text/plain" });
                res.end("Sensitive data exposed over HTTP"); // Exposed sensitive data
            }
        })
        .listen(8080, () => {
            console.log("HTTP server running on http://localhost:8080");
        });
}

// File inclusion vulnerability
function fileInclusion(filePath) {
    try {
        const content = fs.readFileSync(filePath, "utf8"); // Vulnerable to arbitrary file inclusion
        console.log("File content:", content);
    } catch (error) {
        console.error("Error reading file:", error.message);
    }
}

// Unrestricted file upload
function unrestrictedFileUpload(req, res) {
    const filePath = `./uploads/${req.headers["file-name"]}`; // File name from headers without validation
    const fileStream = fs.createWriteStream(filePath);

    req.pipe(fileStream);
    req.on("end", () => {
        res.writeHead(200, { "Content-Type": "text/plain" });
        res.end("File uploaded successfully.");
        console.log(`File uploaded to ${filePath}`);
    });
}

// Main function to test vulnerabilities
function main() {
    console.log("Testing vulnerable functions...");

    hardcodedCredentials();
    sqlInjection("admin' OR 1=1 --");
    weakHash("sensitive_data");
    commandInjection("rm -rf /");
    insecureHTTPServer();
    fileInclusion("../../etc/passwd");

    // Simulate HTTP request for file upload (example usage in a real server)
    console.log(
        "Unrestricted file upload can only be tested in a server environment."
    );
}

main();
