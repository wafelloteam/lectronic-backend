package auth

import (
	"fmt"
	"os"

	sendinblue "github.com/sendinblue/APIv3-go-library/v2/lib"
	"github.com/wafellofazztrack/lectronic-backend/database/orm/model"
	"github.com/wafellofazztrack/lectronic-backend/interfaces"
	"github.com/wafellofazztrack/lectronic-backend/lib"
)

type auth_service struct {
	repo interfaces.UserRepoIF
}

type tokenResponse struct {
	Token string `json:token`
}

func NewService(repo interfaces.UserRepoIF) *auth_service {
	return &auth_service{repo}
}

func (s *auth_service) Login(body *model.User) *lib.Response {

	user, err := s.repo.FindEmail(body.Email)
	if err != nil {
		return lib.NewRes("Email not registered", 401, true)
	}

	if lib.CheckPassword(user.Password, body.Password) {
		return lib.NewRes("Wrong password", 401, true)

	}

	jwt := lib.NewToken(user.ID, user.Role)
	token, err := jwt.CreateToken()
	if err != nil {
		return lib.NewRes(err, 501, true)
	}
	return lib.NewRes(tokenResponse{Token: token}, 200, false)

}

func (s *auth_service) Register(body *model.User) *lib.Response {
	_, err := s.repo.FindEmail(body.Email)
	if err == nil {
		return lib.NewRes("Email has been registered", 401, true)
	}
	hashPassword, err := lib.HashPassword(body.Password)
	if err != nil {
		return lib.NewRes(err.Error(), 400, true)
	}
	body.Password = hashPassword
	data, err := s.repo.Add(body)
	if err != nil {
		return lib.NewRes(err.Error(), 400, true)
	}
	return lib.NewRes(data, 200, false)

}

func (s *auth_service) ForgetPassword(body *model.UserPassword) *lib.Response {

	data, err := s.repo.FindEmail(body.Email)
	if err != nil {
		return lib.NewRes("Email not found", 401, true)
	}
	link := fmt.Sprintf("%s:%s/auth/update-password/%s", os.Getenv("BASE_URL"), os.Getenv("APP_PORT"), data.ID)
	message := sendinblue.SendSmtpEmail{
		Sender: &sendinblue.SendSmtpEmailSender{
			Name:  "Lectronic",
			Email: "lectronic@mail.com",
		},
		To: []sendinblue.SendSmtpEmailTo{
			{
				Name:  data.FullName,
				Email: data.Email,
			},
		},
		Subject: "Reset Your Password",

		HtmlContent: fmt.Sprintf(`
			<div style="font-family: Arial, sans-serif; padding: 20px;">
				<h1 style="font-size: 24px; font-weight: bold;">Reset Your Password</h1>
				<p style="font-size: 16px;">Hello %s,</p>
				<p style="font-size: 16px;">You recently requested to reset your password. Please click the button below to create a new password:</p>
				<a href="%s" style="width: 200px; height: 50px; font-size: 16px; font-weight: bold; text-align: center; color: #fff; background-color: #007bff; padding: 10px; border-radius: 5px; text-decoration: none; margin: 20px auto 0;">Reset Password</a>
				<p style="font-size: 16px;">If you did not make this request, please ignore this email.</p>
				<p style="font-size: 16px;">Thank you,</p>
				<p style="font-size: 16px; font-weight: bold;">The Lectronic Team</p>
			</div>
		`, data.FullName, link),
	}

	_, err = lib.EmailSend(message)

	if err != nil {
		return lib.NewRes(err.Error(), 400, true)
	}
	return lib.NewRes("Email sent successfully", 200, false)
}

func (s *auth_service) UpdatePassword(id string, body *model.UserUpdatePassword) *lib.Response {

	hashPassword, err := lib.HashPassword(body.Password)
	if err != nil {
		return lib.NewRes(err.Error(), 400, true)
	}
	body.Password = hashPassword

	data, err := s.repo.UpdatePassword(id, body)
	if err != nil {
		return lib.NewRes(err.Error(), 400, true)
	}
	return lib.NewRes(data, 200, false)

}
