package lib

import (
	"context"
	"fmt"
	"net/http"
	"os"

	sendinblue "github.com/sendinblue/APIv3-go-library/v2/lib"
)

func EmailSend(message sendinblue.SendSmtpEmail) (*http.Response, error) {
	apiKey := os.Getenv("SIB_KEY")
	var ctx context.Context
	cfg := sendinblue.NewConfiguration()
	//Configure API key authorization: api-key
	cfg.AddDefaultHeader("api-key", apiKey)
	//Configure API key authorization: partner-key
	cfg.AddDefaultHeader("partner-key", apiKey)

	sib := sendinblue.NewAPIClient(cfg)

	// Send the email
	result, resp, err := sib.TransactionalEmailsApi.SendTransacEmail(ctx, message)
	if err != nil {
		fmt.Println("Error when sending email: ", err.Error())
		return nil, err
	}
	fmt.Println("SendTransacEmail result:", result)
	fmt.Println(" SendTransacEmail response: ", resp)
	return resp, nil
}

// func SampleEmailSend() {
// 	// Create the email message
// 	message := sendinblue.SendSmtpEmail{
// 		Sender: &sendinblue.SendSmtpEmailSender{
// 			Name:  "Your Name",
// 			Email: "your@email.com",
// 		},
// 		To: []sendinblue.SendSmtpEmailTo{
// 			{
// 				Name:  "Rizaldi Fauzi",
// 				Email: "rizaldifz@gmail.com",
// 			},
// 		},
// 		Subject:     "Test email",
// 		TextContent: "This is a test email sent using Sendinblue.",
// 		HtmlContent: "<p>This is a test email sent using Sendinblue.</p>",
// 	}

// 	EmailSend(message)
// }
