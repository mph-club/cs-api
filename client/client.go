package main

import (
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cognitoidentityprovider"
)

func main() {
	initiateAuthorization()
}

func initiateAuthorization() *string {
	sess := session.Must(session.NewSession(&aws.Config{
		Region: aws.String("us-east-1"),
	}))

	svc := cognitoidentityprovider.New(sess)

	initAuth, err := svc.InitiateAuth(&cognitoidentityprovider.InitiateAuthInput{
		ClientId: aws.String("7lvbeb74tee3ovdi2c1b42a0pr"),
		AuthFlow: aws.String("USER_PASSWORD_AUTH"),
		AuthParameters: map[string]*string{
			"USERNAME": aws.String("oscar@mphclub.com"),
			"PASSWORD": aws.String("Hunter2!!"),
		},
	})

	if err != nil {
		log.Println("fail!")
		log.Println(err)
	} else {
		log.Println("success!")
		log.Println(initAuth)

		token := initAuth.AuthenticationResult.AccessToken

		return token
	}

	return aws.String("attempt to auth failed")
}

func createUserThroughAWS() {
	sess := session.Must(session.NewSession(&aws.Config{
		MaxRetries: aws.Int(3),
		Region:     aws.String("us-east-1"),
	}))

	svc := cognitoidentityprovider.New(sess)

	signUpOutput, err := svc.SignUp(&cognitoidentityprovider.SignUpInput{
		ClientId: aws.String("7lvbeb74tee3ovdi2c1b42a0pr"),
		Username: aws.String("oscar@mphclub.com"),
		Password: aws.String("Hunter2!!"),
	})

	if err != nil {
		log.Println("fail!")

		if awsErr, ok := err.(awserr.Error); ok {
			log.Println(awsErr.Code())
			log.Println(awsErr.Message())
		}
	} else {
		log.Println("success!")
		log.Println(signUpOutput)
	}
}

func confirmSignup(confirmCode string) {
	sess := session.Must(session.NewSession(&aws.Config{
		MaxRetries: aws.Int(3),
	}))

	svc := cognitoidentityprovider.New(sess)

	confirmUser, err := svc.ConfirmSignUp(&cognitoidentityprovider.ConfirmSignUpInput{
		ClientId:         aws.String("7lvbeb74tee3ovdi2c1b42a0pr"),
		Username:         aws.String("oscar@mphclub.com"),
		ConfirmationCode: aws.String(confirmCode),
	})

	if err != nil {
		log.Println(err)
	} else {
		log.Println("success!")
		log.Println(confirmUser)
	}
}
