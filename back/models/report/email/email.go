package email

import (
	"crypto/tls"
	"fmt"
	"net/smtp"
	"strings"
)

func SendEmail(email string, file []byte) error {
	var (
		serverAddr = "smtp.gmail.com"
		password   = "shady_password"
		emailAddr  = "ShadyElectricCompany@gmail.com"
		portNumber = 465
		tos        = []string{
			email,
		}
		filename  = "recibo.pdf"
		delimeter = "**=myohmy689407924327"
	)

	tlsConfig := tls.Config{
		ServerName:         serverAddr,
		InsecureSkipVerify: true,
	}

	conn, err := tls.Dial("tcp", fmt.Sprintf("%s:%d", serverAddr, portNumber), &tlsConfig)
	if err != nil {
		return err
	}
	defer conn.Close()

	client, err := smtp.NewClient(conn, serverAddr)
	if err != nil {
		return err
	}
	defer client.Close()

	auth := smtp.PlainAuth("", emailAddr, password, serverAddr)

	err = client.Auth(auth)
	if err != nil {
		return err
	}

	err = client.Mail(emailAddr)
	if err != nil {
		return err
	}

	for _, to := range tos {
		err = client.Rcpt(to)
		if err != nil {
			return err
		}
	}

	writer, err := client.Data()
	if err != nil {
		return err
	}

	// basic email headers
	sampleMsg := fmt.Sprintf("From: %s\r\n", emailAddr)
	sampleMsg += fmt.Sprintf("To: %s\r\n", strings.Join(tos, ";"))
	sampleMsg += "Subject: Factura digital\r\n"

	// log.Println("Mark content to accept multiple contents")
	sampleMsg += "MIME-Version: 1.0\r\n"
	sampleMsg += fmt.Sprintf("Content-Type: multipart/mixed; boundary=\"%s\"\r\n", delimeter)

	// place HTML message
	// log.Println("Put HTML message")
	sampleMsg += fmt.Sprintf("\r\n--%s\r\n", delimeter)
	sampleMsg += "Content-Type: text/html; charset=\"utf-8\"\r\n"
	sampleMsg += "Content-Transfer-Encoding: 7bit\r\n"
	sampleMsg += fmt.Sprintf("\r\n%s", "<html><body><h1>Hola</h1>"+
		"<p>Recuerda que si no realizas el pago de tus facturas oportunamente se generan cobros de reconexión y la Empresa te reportará a los operadores de Bancos de Datos y/o Centrales de Riesgo.</p></body></html>\r\n")

	// place file
	// log.Println("Put file attachment")
	sampleMsg += fmt.Sprintf("\r\n--%s\r\n", delimeter)
	sampleMsg += "Content-Type: application/pdf\r\n"
	sampleMsg += "Content-Transfer-Encoding: base64\r\n"
	sampleMsg += "Content-Disposition: attachment;filename=\"" + filename + "\"\r\n"

	// write into email client stream writter
	// log.Println("Write content into client writter I/O")
	sampleMsg += "\r\n" + string(file)

	_, err = writer.Write([]byte(sampleMsg))
	if err != nil {
		return err
	}

	err = writer.Close()
	if err != nil {
		return err
	}

	err = client.Quit()
	if err != nil {
		return err
	}

	return nil
}
