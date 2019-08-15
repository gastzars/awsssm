package awsssm

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ssm"
)

var singletonSSMService = &SSMService{}

//SSMService for SIX
type SSMService struct {
	service ssm.SSM
}

// GetSSMService singleton
func GetSSMService() *SSMService {
	return singletonSSMService
}

// InitSSMService function
func InitSSMService(session *session.Session) *SSMService {
	singletonSSMService.service = *ssm.New(session)
	return singletonSSMService
}

// PutStringParameter function
func (c *SSMService) PutStringParameter(paramName string, paramValue string) error {
	_, err := c.service.PutParameter(&ssm.PutParameterInput{
		Overwrite: aws.Bool(true),
		Name:      aws.String(paramName),
		Value:     aws.String(paramValue),
		Type:      aws.String("String"),
	})
	return err
}

// GetStringParameter function
func (c *SSMService) GetStringParameter(paramName string) (*string, error) {
	strValue, err := c.service.GetParameter(&ssm.GetParameterInput{
		Name: aws.String(paramName),
	})
	if err != nil {
		return nil, err
	}

	return strValue.Parameter.Value, err
}

// GetStringParameters function
func (c *SSMService) GetStringParameters(paramNames []*string) (map[string]string, error) {
	output, err := c.service.GetParameters(&ssm.GetParametersInput{
		Names:          paramNames,
		WithDecryption: aws.Bool(false),
	})
	if err != nil {
		return nil, err
	}
	fmt.Printf("Output param: %v\n", output)
	paramMap := map[string]string{}
	for _, param := range output.Parameters {
		paramMap[*param.Name] = *param.Value
	}
	return paramMap, nil
}
