package libs

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/evidently"
	"github.com/aws/aws-sdk-go-v2/service/evidently/types"
)

type EvidentlyClient struct {
	client *evidently.Client
}

type NewClientInput struct {
	Context context.Context
	Region  string
}

func NewEvidentlyClient(in *NewClientInput) (*EvidentlyClient, error) {
	cfg, err := config.LoadDefaultConfig(in.Context,
		config.WithRegion(in.Region),
	)
	if err != nil {
		return nil, err
	}

	c := evidently.NewFromConfig(cfg)
	return &EvidentlyClient{client: c}, nil
}

type EvaluateFeatureInput struct {
	Project  string
	Feature  string
	EntityID string
}

type EvaluateFeatureOutput struct {
	StringValue  string
	BoolValue    bool
	Int64Value   int64
	Float64Value float64
	Reason       string
}

func (e *EvidentlyClient) EvaluateFeature(ctx context.Context, in *EvaluateFeatureInput) (*EvaluateFeatureOutput, error) {
	// TODO: validate EvaluateFeatureInput
	efIn := &evidently.EvaluateFeatureInput{
		Project:  aws.String(in.Project),
		Feature:  aws.String(in.Feature),
		EntityId: aws.String(in.EntityID),
	}

	out, err := e.client.EvaluateFeature(ctx, efIn)
	if err != nil {
		return nil, err
	}

	res := EvaluateFeatureOutput{
		Reason: aws.ToString(out.Reason),
	}
	switch out.Value.(type) {
	case *types.VariableValueMemberStringValue:
		res.StringValue = out.Value.(*types.VariableValueMemberStringValue).Value
	case *types.VariableValueMemberBoolValue:
		res.BoolValue = out.Value.(*types.VariableValueMemberBoolValue).Value
	case *types.VariableValueMemberLongValue:
		res.Int64Value = out.Value.(*types.VariableValueMemberLongValue).Value
	case *types.VariableValueMemberDoubleValue:
		res.Float64Value = out.Value.(*types.VariableValueMemberDoubleValue).Value
	default:
		// noop
	}

	return &res, nil
}
