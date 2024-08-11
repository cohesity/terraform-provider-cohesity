package cohesity

import (
	"golang.org/x/crypto/ssh"
	"io"
	"io/ioutil"
	"log"
	"os"
	"time"
)

func newSSHClient(host, user, password string, timeout time.Duration) (*ssh.Client, error) {
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

func executeSSHCommand(client *ssh.Client, command string) error {
	session, err := client.NewSession()
	if err != nil {
		return err
	}
	defer session.Close()

	stdout, err := session.StdoutPipe()
	if err != nil {
		return err
	}

	stderr, err := session.StderrPipe()
	if err != nil {
		return err
	}

	err = session.Start(command)
	if err != nil {
		return err
	}

	go func() {
		_, err := io.Copy(os.Stdout, stdout)
		if err != nil {
			log.Printf("Error copying stdout: %v", err)
		}
	}()

	go func() {
		_, err := io.Copy(os.Stderr, stderr)
		if err != nil {
			log.Printf("Error copying stderr: %v", err)
		}
	}()

	err = session.Wait()
	if err != nil {
		return err
	}

	return nil
}

func transferFile(client *ssh.Client, source string, destination string) error {
	session, err := client.NewSession()
	if err != nil {
		return err
	}
	defer session.Close()

	content, err := ioutil.ReadFile(source)
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
