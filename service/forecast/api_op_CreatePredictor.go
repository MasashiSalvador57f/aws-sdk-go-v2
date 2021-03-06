// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package forecast

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type CreatePredictorInput struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Name (ARN) of the algorithm to use for model training.
	// Required if PerformAutoML is not set to true.
	//
	// Supported algorithms:
	//
	//    * arn:aws:forecast:::algorithm/ARIMA
	//
	//    * arn:aws:forecast:::algorithm/Deep_AR_Plus Supports hyperparameter optimization
	//    (HPO)
	//
	//    * arn:aws:forecast:::algorithm/ETS
	//
	//    * arn:aws:forecast:::algorithm/NPTS
	//
	//    * arn:aws:forecast:::algorithm/Prophet
	AlgorithmArn *string `type:"string"`

	// An AWS Key Management Service (KMS) key and the AWS Identity and Access Management
	// (IAM) role that Amazon Forecast can assume to access the key.
	EncryptionConfig *EncryptionConfig `type:"structure"`

	// Used to override the default evaluation parameters of the specified algorithm.
	// Amazon Forecast evaluates a predictor by splitting a dataset into training
	// data and testing data. The evaluation parameters define how to perform the
	// split and the number of iterations.
	EvaluationParameters *EvaluationParameters `type:"structure"`

	// The featurization configuration.
	//
	// FeaturizationConfig is a required field
	FeaturizationConfig *FeaturizationConfig `type:"structure" required:"true"`

	// Specifies the number of time-steps that the model is trained to predict.
	// The forecast horizon is also called the prediction length.
	//
	// For example, if you configure a dataset for daily data collection (using
	// the DataFrequency parameter of the CreateDataset operation) and set the forecast
	// horizon to 10, the model returns predictions for 10 days.
	//
	// The maximum forecast horizon is the lesser of 500 time-steps or 1/3 of the
	// TARGET_TIME_SERIES dataset length.
	//
	// ForecastHorizon is a required field
	ForecastHorizon *int64 `type:"integer" required:"true"`

	// Provides hyperparameter override values for the algorithm. If you don't provide
	// this parameter, Amazon Forecast uses default values. The individual algorithms
	// specify which hyperparameters support hyperparameter optimization (HPO).
	// For more information, see aws-forecast-choosing-recipes.
	//
	// If you included the HPOConfig object, you must set PerformHPO to true.
	HPOConfig *HyperParameterTuningJobConfig `type:"structure"`

	// Describes the dataset group that contains the data to use to train the predictor.
	//
	// InputDataConfig is a required field
	InputDataConfig *InputDataConfig `type:"structure" required:"true"`

	// Whether to perform AutoML. When Amazon Forecast performs AutoML, it evaluates
	// the algorithms it provides and chooses the best algorithm and configuration
	// for your training dataset.
	//
	// The default value is false. In this case, you are required to specify an
	// algorithm.
	//
	// Set PerformAutoML to true to have Amazon Forecast perform AutoML. This is
	// a good option if you aren't sure which algorithm is suitable for your training
	// data. In this case, PerformHPO must be false.
	PerformAutoML *bool `type:"boolean"`

	// Whether to perform hyperparameter optimization (HPO). HPO finds optimal hyperparameter
	// values for your training data. The process of performing HPO is known as
	// running a hyperparameter tuning job.
	//
	// The default value is false. In this case, Amazon Forecast uses default hyperparameter
	// values from the chosen algorithm.
	//
	// To override the default values, set PerformHPO to true and, optionally, supply
	// the HyperParameterTuningJobConfig object. The tuning job specifies a metric
	// to optimize, which hyperparameters participate in tuning, and the valid range
	// for each tunable hyperparameter. In this case, you are required to specify
	// an algorithm and PerformAutoML must be false.
	//
	// The following algorithm supports HPO:
	//
	//    * DeepAR+
	PerformHPO *bool `type:"boolean"`

	// A name for the predictor.
	//
	// PredictorName is a required field
	PredictorName *string `min:"1" type:"string" required:"true"`

	// The optional metadata that you apply to the predictor to help you categorize
	// and organize them. Each tag consists of a key and an optional value, both
	// of which you define.
	//
	// The following basic restrictions apply to tags:
	//
	//    * Maximum number of tags per resource - 50.
	//
	//    * For each resource, each tag key must be unique, and each tag key can
	//    have only one value.
	//
	//    * Maximum key length - 128 Unicode characters in UTF-8.
	//
	//    * Maximum value length - 256 Unicode characters in UTF-8.
	//
	//    * If your tagging schema is used across multiple services and resources,
	//    remember that other services may have restrictions on allowed characters.
	//    Generally allowed characters are: letters, numbers, and spaces representable
	//    in UTF-8, and the following characters: + - = . _ : / @.
	//
	//    * Tag keys and values are case sensitive.
	//
	//    * Do not use aws:, AWS:, or any upper or lowercase combination of such
	//    as a prefix for keys as it is reserved for AWS use. You cannot edit or
	//    delete tag keys with this prefix. Values can have this prefix. If a tag
	//    value has aws as its prefix but the key does not, then Forecast considers
	//    it to be a user tag and will count against the limit of 50 tags. Tags
	//    with only the key prefix of aws do not count against your tags per resource
	//    limit.
	Tags []Tag `type:"list"`

	// The hyperparameters to override for model training. The hyperparameters that
	// you can override are listed in the individual algorithms. For the list of
	// supported algorithms, see aws-forecast-choosing-recipes.
	TrainingParameters map[string]string `type:"map"`
}

// String returns the string representation
func (s CreatePredictorInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreatePredictorInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreatePredictorInput"}

	if s.FeaturizationConfig == nil {
		invalidParams.Add(aws.NewErrParamRequired("FeaturizationConfig"))
	}

	if s.ForecastHorizon == nil {
		invalidParams.Add(aws.NewErrParamRequired("ForecastHorizon"))
	}

	if s.InputDataConfig == nil {
		invalidParams.Add(aws.NewErrParamRequired("InputDataConfig"))
	}

	if s.PredictorName == nil {
		invalidParams.Add(aws.NewErrParamRequired("PredictorName"))
	}
	if s.PredictorName != nil && len(*s.PredictorName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("PredictorName", 1))
	}
	if s.EncryptionConfig != nil {
		if err := s.EncryptionConfig.Validate(); err != nil {
			invalidParams.AddNested("EncryptionConfig", err.(aws.ErrInvalidParams))
		}
	}
	if s.FeaturizationConfig != nil {
		if err := s.FeaturizationConfig.Validate(); err != nil {
			invalidParams.AddNested("FeaturizationConfig", err.(aws.ErrInvalidParams))
		}
	}
	if s.HPOConfig != nil {
		if err := s.HPOConfig.Validate(); err != nil {
			invalidParams.AddNested("HPOConfig", err.(aws.ErrInvalidParams))
		}
	}
	if s.InputDataConfig != nil {
		if err := s.InputDataConfig.Validate(); err != nil {
			invalidParams.AddNested("InputDataConfig", err.(aws.ErrInvalidParams))
		}
	}
	if s.Tags != nil {
		for i, v := range s.Tags {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "Tags", i), err.(aws.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type CreatePredictorOutput struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Name (ARN) of the predictor.
	PredictorArn *string `type:"string"`
}

// String returns the string representation
func (s CreatePredictorOutput) String() string {
	return awsutil.Prettify(s)
}

const opCreatePredictor = "CreatePredictor"

// CreatePredictorRequest returns a request value for making API operation for
// Amazon Forecast Service.
//
// Creates an Amazon Forecast predictor.
//
// In the request, you provide a dataset group and either specify an algorithm
// or let Amazon Forecast choose the algorithm for you using AutoML. If you
// specify an algorithm, you also can override algorithm-specific hyperparameters.
//
// Amazon Forecast uses the chosen algorithm to train a model using the latest
// version of the datasets in the specified dataset group. The result is called
// a predictor. You then generate a forecast using the CreateForecast operation.
//
// After training a model, the CreatePredictor operation also evaluates it.
// To see the evaluation metrics, use the GetAccuracyMetrics operation. Always
// review the evaluation metrics before deciding to use the predictor to generate
// a forecast.
//
// Optionally, you can specify a featurization configuration to fill and aggregate
// the data fields in the TARGET_TIME_SERIES dataset to improve model training.
// For more information, see FeaturizationConfig.
//
// For RELATED_TIME_SERIES datasets, CreatePredictor verifies that the DataFrequency
// specified when the dataset was created matches the ForecastFrequency. TARGET_TIME_SERIES
// datasets don't have this restriction. Amazon Forecast also verifies the delimiter
// and timestamp format. For more information, see howitworks-datasets-groups.
//
// AutoML
//
// If you want Amazon Forecast to evaluate each algorithm and choose the one
// that minimizes the objective function, set PerformAutoML to true. The objective
// function is defined as the mean of the weighted p10, p50, and p90 quantile
// losses. For more information, see EvaluationResult.
//
// When AutoML is enabled, the following properties are disallowed:
//
//    * AlgorithmArn
//
//    * HPOConfig
//
//    * PerformHPO
//
//    * TrainingParameters
//
// To get a list of all of your predictors, use the ListPredictors operation.
//
// Before you can use the predictor to create a forecast, the Status of the
// predictor must be ACTIVE, signifying that training has completed. To get
// the status, use the DescribePredictor operation.
//
//    // Example sending a request using CreatePredictorRequest.
//    req := client.CreatePredictorRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/forecast-2018-06-26/CreatePredictor
func (c *Client) CreatePredictorRequest(input *CreatePredictorInput) CreatePredictorRequest {
	op := &aws.Operation{
		Name:       opCreatePredictor,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CreatePredictorInput{}
	}

	req := c.newRequest(op, input, &CreatePredictorOutput{})

	return CreatePredictorRequest{Request: req, Input: input, Copy: c.CreatePredictorRequest}
}

// CreatePredictorRequest is the request type for the
// CreatePredictor API operation.
type CreatePredictorRequest struct {
	*aws.Request
	Input *CreatePredictorInput
	Copy  func(*CreatePredictorInput) CreatePredictorRequest
}

// Send marshals and sends the CreatePredictor API request.
func (r CreatePredictorRequest) Send(ctx context.Context) (*CreatePredictorResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreatePredictorResponse{
		CreatePredictorOutput: r.Request.Data.(*CreatePredictorOutput),
		response:              &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreatePredictorResponse is the response type for the
// CreatePredictor API operation.
type CreatePredictorResponse struct {
	*CreatePredictorOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreatePredictor request.
func (r *CreatePredictorResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
