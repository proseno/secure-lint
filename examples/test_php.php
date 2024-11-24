<?php

// Hardcoded credentials vulnerability
function hardcodedCredentials() {
    $username = "admin"; // Hardcoded username
    $password = "supersecret"; // Hardcoded password
    echo "Connecting with username: $username and password: $password\n";
}

// SQL Injection vulnerability
function sqlInjection($userInput) {
    $conn = new mysqli("localhost", "root", "password", "test");
    if ($conn->connect_error) {
        die("Connection failed: " . $conn->connect_error);
    }

    $query = "SELECT * FROM users WHERE username = '$userInput'"; // Vulnerable to SQL Injection
    $result = $conn->query($query);

    if ($result) {
        while ($row = $result->fetch_assoc()) {
            echo "User: " . $row['username'] . "\n";
        }
    }

    $conn->close();
}

// Weak cryptographic hash function
function weakHash($data) {
    $hash = md5($data); // Weak MD5 hash
    echo "MD5 hash of $data: $hash\n";
}

// Command injection vulnerability
function commandInjection($userInput) {
    $command = "echo " . $userInput; // Vulnerable to command injection
    system($command); // Using system() with unvalidated input is dangerous!
}

// Insecure file inclusion
function fileInclusion($filePath) {
    include($filePath); // Vulnerable to arbitrary file inclusion
}

// Unrestricted file upload
function unrestrictedFileUpload($uploadedFile) {
    $targetDir = "/var/www/uploads/";
    $targetFile = $targetDir . basename($uploadedFile['name']);
    move_uploaded_file($uploadedFile['tmp_name'], $targetFile); // No validation on file type or name
    echo "File uploaded to: $targetFile\n";
}

// Insecure direct object reference (IDOR)
function insecureIDOR($userId) {
    $filePath = "/var/www/user_data/" . $userId . ".txt"; // Direct file access without validation
    if (file_exists($filePath)) {
        echo "User Data: " . file_get_contents($filePath);
    } else {
        echo "File not found.\n";
    }
}

// Main function for testing
function main() {
    echo "Testing vulnerable functions...\n";

    // Call vulnerable functions
    hardcodedCredentials();
    sqlInjection("'; DROP TABLE users; --");
    weakHash("sensitive_data");
    commandInjection("rm -rf /");
    fileInclusion("../../etc/passwd");
    insecureIDOR("12345");

    // Simulate file upload
    $uploadedFile = [
        'name' => 'malicious.php',
        'tmp_name' => '/tmp/phpYzdqkD'
    ];
    unrestrictedFileUpload($uploadedFile);
}

main();
?>
