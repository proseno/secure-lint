import os
import sqlite3
import hashlib
import subprocess
import requests

# Hardcoded credentials vulnerability
def hardcoded_credentials():
    username = "admin"  # Hardcoded username
    password = "supersecret"  # Hardcoded password
    print(f"Connecting with username: {username} and password: {password}")

# SQL Injection vulnerability
def sql_injection(user_input):
    conn = sqlite3.connect("example.db")
    cursor = conn.cursor()
    query = f"SELECT * FROM users WHERE username = '{user_input}'"  # Vulnerable to SQL Injection
    cursor.execute(query)
    print(cursor.fetchall())
    conn.close()

# Weak cryptographic hash function
def weak_hash(data):
    hash_value = hashlib.md5(data.encode()).hexdigest()  # Weak MD5 hash
    print(f"MD5 hash of {data}: {hash_value}")

# Command injection vulnerability
def command_injection(user_input):
    command = f"echo {user_input}"  # Vulnerable to command injection
    subprocess.run(command, shell=True)  # Using shell=True is dangerous!

# Insecure HTTP connection
def insecure_http_request(url):
    response = requests.get(url)  # Unsecured HTTP request
    print(f"Response from {url}: {response.text}")

# File inclusion vulnerability
def file_inclusion(file_path):
    with open(file_path, "r") as file:  # Vulnerable to arbitrary file inclusion
        print(file.read())

# Main function for testing
def main():
    print("Testing vulnerable functions...")

    # Call vulnerable functions
    hardcoded_credentials()
    sql_injection("'; DROP TABLE users; --")
    weak_hash("sensitive_data")
    command_injection("rm -rf /")
    insecure_http_request("http://example.com")
    file_inclusion("../../etc/passwd")

if __name__ == "__main__":
    main()
