// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package pi

import (
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

var _ aws.Config
var _ = awsutil.Prettify

// A timestamp, and a single numerical value, which together represent a measurement
// at a particular point in time.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/pi-2018-02-27/DataPoint
type DataPoint struct {
	_ struct{} `type:"structure"`

	// The time, in epoch format, associated with a particular Value.
	//
	// Timestamp is a required field
	Timestamp *time.Time `type:"timestamp" timestampFormat:"unix" required:"true"`

	// The actual value associated with a particular Timestamp.
	//
	// Value is a required field
	Value *float64 `type:"double" required:"true"`
}

// String returns the string representation
func (s DataPoint) String() string {
	return awsutil.Prettify(s)
}

// A logical grouping of Performance Insights metrics for a related subject
// area. For example, the db.sql dimension group consists of the following dimensions:
// db.sql.id, db.sql.db_id, db.sql.statement, and db.sql.tokenized_id.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/pi-2018-02-27/DimensionGroup
type DimensionGroup struct {
	_ struct{} `type:"structure"`

	// A list of specific dimensions from a dimension group. If this parameter is
	// not present, then it signifies that all of the dimensions in the group were
	// requested, or are present in the response.
	//
	// Valid values for elements in the Dimensions array are:
	//
	//    * db.user.id
	//
	//    * db.user.name
	//
	//    * db.host.id
	//
	//    * db.host.name
	//
	//    * db.sql.id
	//
	//    * db.sql.db_id
	//
	//    * db.sql.statement
	//
	//    * db.sql.tokenized_id
	//
	//    * db.sql_tokenized.id
	//
	//    * db.sql_tokenized.db_id
	//
	//    * db.sql_tokenized.statement
	//
	//    * db.wait_event.name
	//
	//    * db.wait_event.type
	//
	//    * db.wait_event_type.name
	Dimensions []string `min:"1" type:"list"`

	// The name of the dimension group. Valid values are:
	//
	//    * db.user
	//
	//    * db.host
	//
	//    * db.sql
	//
	//    * db.sql_tokenized
	//
	//    * db.wait_event
	//
	//    * db.wait_event_type
	//
	// Group is a required field
	Group *string `type:"string" required:"true"`

	// The maximum number of items to fetch for this dimension group.
	Limit *int64 `min:"1" type:"integer"`
}

// String returns the string representation
func (s DimensionGroup) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DimensionGroup) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DimensionGroup"}
	if s.Dimensions != nil && len(s.Dimensions) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Dimensions", 1))
	}

	if s.Group == nil {
		invalidParams.Add(aws.NewErrParamRequired("Group"))
	}
	if s.Limit != nil && *s.Limit < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("Limit", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// An array of descriptions and aggregated values for each dimension within
// a dimension group.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/pi-2018-02-27/DimensionKeyDescription
type DimensionKeyDescription struct {
	_ struct{} `type:"structure"`

	// A map of name-value pairs for the dimensions in the group.
	Dimensions map[string]string `type:"map"`

	// If PartitionBy was specified, PartitionKeys contains the dimensions that
	// were.
	Partitions []float64 `type:"list"`

	// The aggregated metric value for the dimension(s), over the requested time
	// range.
	Total *float64 `type:"double"`
}

// String returns the string representation
func (s DimensionKeyDescription) String() string {
	return awsutil.Prettify(s)
}

// A time-ordered series of data points, correpsonding to a dimension of a Performance
// Insights metric.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/pi-2018-02-27/MetricKeyDataPoints
type MetricKeyDataPoints struct {
	_ struct{} `type:"structure"`

	// An array of timestamp-value pairs, representing measurements over a period
	// of time.
	DataPoints []DataPoint `type:"list"`

	// The dimension(s) to which the data points apply.
	Key *ResponseResourceMetricKey `type:"structure"`
}

// String returns the string representation
func (s MetricKeyDataPoints) String() string {
	return awsutil.Prettify(s)
}

// A single query to be processed. You must provide the metric to query. If
// no other parameters are specified, Performance Insights returns all of the
// data points for that metric. You can optionally request that the data points
// be aggregated by dimension group ( GroupBy), and return only those data points
// that match your criteria (Filter).
// Please also see https://docs.aws.amazon.com/goto/WebAPI/pi-2018-02-27/MetricQuery
type MetricQuery struct {
	_ struct{} `type:"structure"`

	// One or more filters to apply in the request. Restrictions:
	//
	//    * Any number of filters by the same dimension, as specified in the GroupBy
	//    parameter.
	//
	//    * A single filter for any other dimension in this dimension group.
	Filter map[string]string `type:"map"`

	// A specification for how to aggregate the data points from a query result.
	// You must specify a valid dimension group. Performance Insights will return
	// all of the dimensions within that group, unless you provide the names of
	// specific dimensions within that group. You can also request that Performance
	// Insights return a limited number of values for a dimension.
	GroupBy *DimensionGroup `type:"structure"`

	// The name of a Performance Insights metric to be measured.
	//
	// Valid values for Metric are:
	//
	//    * db.load.avg - a scaled representation of the number of active sessions
	//    for the database engine.
	//
	//    * db.sampledload.avg - the raw number of active sessions for the database
	//    engine.
	//
	// Metric is a required field
	Metric *string `type:"string" required:"true"`
}

// String returns the string representation
func (s MetricQuery) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *MetricQuery) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "MetricQuery"}

	if s.Metric == nil {
		invalidParams.Add(aws.NewErrParamRequired("Metric"))
	}
	if s.GroupBy != nil {
		if err := s.GroupBy.Validate(); err != nil {
			invalidParams.AddNested("GroupBy", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// If PartitionBy was specified in a DescribeDimensionKeys request, the dimensions
// are returned in an array. Each element in the array specifies one dimension.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/pi-2018-02-27/ResponsePartitionKey
type ResponsePartitionKey struct {
	_ struct{} `type:"structure"`

	// A dimension map that contains the dimension(s) for this partition.
	//
	// Dimensions is a required field
	Dimensions map[string]string `type:"map" required:"true"`
}

// String returns the string representation
func (s ResponsePartitionKey) String() string {
	return awsutil.Prettify(s)
}

// An object describing a Performance Insights metric and one or more dimensions
// for that metric.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/pi-2018-02-27/ResponseResourceMetricKey
type ResponseResourceMetricKey struct {
	_ struct{} `type:"structure"`

	// The valid dimensions for the metric.
	Dimensions map[string]string `type:"map"`

	// The name of a Performance Insights metric to be measured.
	//
	// Valid values for Metric are:
	//
	//    * db.load.avg - a scaled representation of the number of active sessions
	//    for the database engine.
	//
	//    * db.sampledload.avg - the raw number of active sessions for the database
	//    engine.
	//
	// Metric is a required field
	Metric *string `type:"string" required:"true"`
}

// String returns the string representation
func (s ResponseResourceMetricKey) String() string {
	return awsutil.Prettify(s)
}