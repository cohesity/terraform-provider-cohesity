package services

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"
	"time"

	"golang.org/x/crypto/ssh"
)

func NewSSHClient(host, user, password string, timeout time.Duration) (*ssh.Client, error) {
	var authMethods []ssh.AuthMethod
	if password != "" {
		authMethods = append(authMethods, ssh.Password(password))
	}
	config := &ssh.ClientConfig{
		User:            user,
		Auth:            authMethods,
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		Timeout:         timeout,
	}

	conn, err := ssh.Dial("tcp", host+":22", config)
	if err != nil {
		return nil, err
	}

	return conn, nil
}

func ExecuteSSHCommand(client *ssh.Client, command string) (string, error) {
	session, err := client.NewSession()
	if err != nil {
		return "", err
	}
	defer session.Close()

	// Set up stdout and stderr streaming
	stdout, err := session.StdoutPipe()
	if err != nil {
		return "", err
	}
	stderr, err := session.StderrPipe()
	if err != nil {
		return "", err
	}

	// Start the command
	if err := session.Start(command); err != nil {
		return "", err
	}

	// Buffer to store output for later use
	var outputBuffer bytes.Buffer

	// Use a scanner to read stdout line by line
	scanner := bufio.NewScanner(stdout)
	go func() {
		for scanner.Scan() {
			line := scanner.Text()
			log.Printf("[INFO] STDOUT: %s", line)
			outputBuffer.WriteString(line + "\n")
		}
	}()

	// Use a scanner to read stderr line by line
	errScanner := bufio.NewScanner(stderr)
	go func() {
		for errScanner.Scan() {
			line := errScanner.Text()
			log.Printf("[INFO] STDERR: %s", line)
			outputBuffer.WriteString(line + "\n")
		}
	}()

	// Wait for the session to finish
	if err := session.Wait(); err != nil {
		return outputBuffer.String(), err
	}

	return outputBuffer.String(), nil

}

func TransferFile(client *ssh.Client, source string, destination string) error {
	session, err := client.NewSession()
	if err != nil {
		return err
	}
	defer session.Close()

	content, err := os.ReadFile(source)
	if err != nil {
		return err
	}

	stdin, err := session.StdinPipe()
	if err != nil {
		return err
	}

	go func() {
		defer stdin.Close()
		stdin.Write(content)
	}()

	if err := session.Run("cat > " + destination); err != nil {
		return err
	}

	return nil
}

// UploadScriptToVM uploads a script to the VM
func UploadScriptToVM(client *ssh.Client, path string, content string) error {
	// Init the ssh session.
	session, err := client.NewSession()
	if err != nil {
		return err
	}
	defer session.Close()

	// Write the script to a file on the VM
	scriptFile := fmt.Sprintf("echo '%s' > %s && chmod 777 %s", content, path, path)
	err = session.Run(scriptFile)
	if err != nil {
		return fmt.Errorf("failed to write script to VM: %s", err)
	}

	return nil
}
