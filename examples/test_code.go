package examples

import (
	"crypto/md5"
	"crypto/tls"
	"database/sql"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
)

func HardcodedCredentials() {
	username := "admin"
	password := "supersecret" // Hardcoded password
	fmt.Printf("Logging in as %s with password %s\n", username, password)
}

func SqlInjections() {
	db, err := sql.Open("postgres",
		"user=admin password=supersecret dbname=test sslmode=disable")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	userInput := "'; DROP TABLE users; --"
	query := fmt.Sprintf(
		"SELECT * FROM users WHERE username='%s'",
		userInput) // SQL Injection
	rows, err := db.Query(query)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
}

func WeakCryptography() {
	data := []byte("sensitive data")
	hash := md5.Sum(data) // Weak cryptographic hash
	fmt.Printf("MD5 hash: %x\n", hash)
}

func InsecureTlsConfiguration() {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true}, // Insecure TLS configuration
	}
	client := &http.Client{Transport: tr}
	resp, err := client.Get("https://example.com")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
}

func FileInclusionVulnerability() {
	userInput := "../../etc/passwd" // Arbitrary file inclusion
	data, err := ioutil.ReadFile(userInput)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	fmt.Println("File contents:", string(data))
}

func CommandInjection() {
	userInput := "ls; rm -rf /" // Command injection
	cmd := exec.Command("sh", "-c", userInput)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		fmt.Println("Command failed:", err)
	}
}
