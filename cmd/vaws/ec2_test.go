package vaws

import (
	"bytes"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/aws/smithy-go/middleware"
	"github.com/olekukonko/tablewriter"
	"testing"
)

func Test_printEc2Instances(t *testing.T) {
	type args struct {
		outputs      []*ec2.DescribeInstancesOutput
		sortPosition int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "default",
			args: args{
				outputs: []*ec2.DescribeInstancesOutput{
					{
						NextToken: nil,
						Reservations: []types.Reservation{
							{
								Instances: []types.Instance{
									{
										InstanceId:       aws.String("i-0abee92626b0a28a7"),
										InstanceType:     "t3.nano",
										PrivateIpAddress: aws.String("172.31.18.8"),
										PublicIpAddress:  aws.String("35.73.127.100"),
										SecurityGroups: []types.GroupIdentifier{
											{
												GroupId:   aws.String("sg-0d642190887707fd0"),
												GroupName: aws.String("launch-wizard-1"),
											},
										},
										State: &types.InstanceState{
											Name: "running",
										},
										Tags: []types.Tag{
											{
												Key:   aws.String("Name"),
												Value: aws.String("web01"),
											},
										},
									},
								},
							},
							{
								Instances: []types.Instance{
									{
										InstanceId:       aws.String("i-06723a6629e542c50"),
										InstanceType:     "t3.small",
										PrivateIpAddress: aws.String("172.31.21.33"),
										PublicIpAddress:  aws.String("18.179.33.182"),
										SecurityGroups: []types.GroupIdentifier{
											{
												GroupId:   aws.String("sg-08d35fef29987e75e"),
												GroupName: aws.String("launch-wizard-3"),
											},
										},
										State: &types.InstanceState{
											Name: "running",
										},
										Tags: []types.Tag{
											{
												Key:   aws.String("Name"),
												Value: aws.String("app02"),
											},
										},
									},
								},
							},
							{
								Instances: []types.Instance{
									{
										InstanceId:       aws.String("i-06d4c29e4e5ccadc4"),
										InstanceType:     "t2.micro",
										PrivateIpAddress: aws.String("172.31.35.175"),
										PublicIpAddress:  aws.String("54.238.30.226"),
										SecurityGroups: []types.GroupIdentifier{
											{
												GroupId:   aws.String("sg-0f0b4c4642ffb5ef2"),
												GroupName: aws.String("launch-wizard-2"),
											},
										},
										State: &types.InstanceState{
											Name: "running",
										},
										Tags: []types.Tag{
											{
												Key:   aws.String("Name"),
												Value: aws.String("app01"),
											},
										},
									},
								},
							},
						},
						ResultMetadata: middleware.Metadata{},
					},
				},
				sortPosition: 1,
			},
			want: `+-------+---------------------+----------+---------------+---------------+---------+---------------------------------------+
| NAME  |         ID          |   TYPE   |  PRIVATE IP   |   PUBLIC IP   |  STATE  |            SECURITY GROUP             |
+-------+---------------------+----------+---------------+---------------+---------+---------------------------------------+
| app01 | i-06d4c29e4e5ccadc4 | t2.micro | 172.31.35.175 | 54.238.30.226 | running | launch-wizard-2(sg-0f0b4c4642ffb5ef2) |
| app02 | i-06723a6629e542c50 | t3.small | 172.31.21.33  | 18.179.33.182 | running | launch-wizard-3(sg-08d35fef29987e75e) |
| web01 | i-0abee92626b0a28a7 | t3.nano  | 172.31.18.8   | 35.73.127.100 | running | launch-wizard-1(sg-0d642190887707fd0) |
+-------+---------------------+----------+---------------+---------------+---------+---------------------------------------+
`,
		},
		{
			name: "sort request",
			args: args{
				outputs: []*ec2.DescribeInstancesOutput{
					{
						NextToken: nil,
						Reservations: []types.Reservation{
							{
								Instances: []types.Instance{
									{
										InstanceId:       aws.String("i-0abee92626b0a28a7"),
										InstanceType:     "t3.nano",
										PrivateIpAddress: aws.String("172.31.18.8"),
										PublicIpAddress:  aws.String("35.73.127.100"),
										SecurityGroups: []types.GroupIdentifier{
											{
												GroupId:   aws.String("sg-0d642190887707fd0"),
												GroupName: aws.String("launch-wizard-1"),
											},
										},
										State: &types.InstanceState{
											Name: "running",
										},
										Tags: []types.Tag{
											{
												Key:   aws.String("Name"),
												Value: aws.String("web01"),
											},
										},
									},
								},
							},
							{
								Instances: []types.Instance{
									{
										InstanceId:       aws.String("i-06723a6629e542c50"),
										InstanceType:     "t3.small",
										PrivateIpAddress: aws.String("172.31.21.33"),
										PublicIpAddress:  aws.String("18.179.33.182"),
										SecurityGroups: []types.GroupIdentifier{
											{
												GroupId:   aws.String("sg-08d35fef29987e75e"),
												GroupName: aws.String("launch-wizard-3"),
											},
										},
										State: &types.InstanceState{
											Name: "running",
										},
										Tags: []types.Tag{
											{
												Key:   aws.String("Name"),
												Value: aws.String("app02"),
											},
										},
									},
								},
							},
							{
								Instances: []types.Instance{
									{
										InstanceId:       aws.String("i-06d4c29e4e5ccadc4"),
										InstanceType:     "t2.micro",
										PrivateIpAddress: aws.String("172.31.35.175"),
										PublicIpAddress:  aws.String("54.238.30.226"),
										SecurityGroups: []types.GroupIdentifier{
											{
												GroupId:   aws.String("sg-0f0b4c4642ffb5ef2"),
												GroupName: aws.String("launch-wizard-2"),
											},
										},
										State: &types.InstanceState{
											Name: "running",
										},
										Tags: []types.Tag{
											{
												Key:   aws.String("Name"),
												Value: aws.String("app01"),
											},
										},
									},
								},
							},
						},
						ResultMetadata: middleware.Metadata{},
					},
				},
				sortPosition: 7,
			},
			want: `+-------+---------------------+----------+---------------+---------------+---------+---------------------------------------+
| NAME  |         ID          |   TYPE   |  PRIVATE IP   |   PUBLIC IP   |  STATE  |            SECURITY GROUP             |
+-------+---------------------+----------+---------------+---------------+---------+---------------------------------------+
| web01 | i-0abee92626b0a28a7 | t3.nano  | 172.31.18.8   | 35.73.127.100 | running | launch-wizard-1(sg-0d642190887707fd0) |
| app01 | i-06d4c29e4e5ccadc4 | t2.micro | 172.31.35.175 | 54.238.30.226 | running | launch-wizard-2(sg-0f0b4c4642ffb5ef2) |
| app02 | i-06723a6629e542c50 | t3.small | 172.31.21.33  | 18.179.33.182 | running | launch-wizard-3(sg-08d35fef29987e75e) |
+-------+---------------------+----------+---------------+---------------+---------+---------------------------------------+
`,
		},
		{
			name: "two request",
			args: args{
				outputs: []*ec2.DescribeInstancesOutput{
					{
						NextToken: nil,
						Reservations: []types.Reservation{
							{
								Instances: []types.Instance{
									{
										InstanceId:       aws.String("i-0abee92626b0a28a7"),
										InstanceType:     "t3.nano",
										PrivateIpAddress: aws.String("172.31.18.8"),
										PublicIpAddress:  aws.String("35.73.127.100"),
										SecurityGroups: []types.GroupIdentifier{
											{
												GroupId:   aws.String("sg-0d642190887707fd0"),
												GroupName: aws.String("launch-wizard-1"),
											},
										},
										State: &types.InstanceState{
											Name: "running",
										},
										Tags: []types.Tag{
											{
												Key:   aws.String("Name"),
												Value: aws.String("web01"),
											},
										},
									},
								},
							},
							{
								Instances: []types.Instance{
									{
										InstanceId:       aws.String("i-06723a6629e542c50"),
										InstanceType:     "t3.small",
										PrivateIpAddress: aws.String("172.31.21.33"),
										PublicIpAddress:  aws.String("18.179.33.182"),
										SecurityGroups: []types.GroupIdentifier{
											{
												GroupId:   aws.String("sg-08d35fef29987e75e"),
												GroupName: aws.String("launch-wizard-3"),
											},
										},
										State: &types.InstanceState{
											Name: "running",
										},
										Tags: []types.Tag{
											{
												Key:   aws.String("Name"),
												Value: aws.String("app02"),
											},
										},
									},
								},
							},
							{
								Instances: []types.Instance{
									{
										InstanceId:       aws.String("i-06d4c29e4e5ccadc4"),
										InstanceType:     "t2.micro",
										PrivateIpAddress: aws.String("172.31.35.175"),
										PublicIpAddress:  aws.String("54.238.30.226"),
										SecurityGroups: []types.GroupIdentifier{
											{
												GroupId:   aws.String("sg-0f0b4c4642ffb5ef2"),
												GroupName: aws.String("launch-wizard-2"),
											},
										},
										State: &types.InstanceState{
											Name: "running",
										},
										Tags: []types.Tag{
											{
												Key:   aws.String("Name"),
												Value: aws.String("app01"),
											},
										},
									},
								},
							},
						},
						ResultMetadata: middleware.Metadata{},
					},
					{
						NextToken: nil,
						Reservations: []types.Reservation{
							{
								Instances: []types.Instance{
									{
										InstanceId:       aws.String("i-06d4c29e4e5edx5tf"),
										InstanceType:     "t2.micro",
										PrivateIpAddress: aws.String("172.31.35.130"),
										PublicIpAddress:  aws.String("54.238.30.211"),
										SecurityGroups: []types.GroupIdentifier{
											{
												GroupId:   aws.String("sg-0f0b4c4642ffb5ef2"),
												GroupName: aws.String("launch-wizard-2"),
											},
										},
										State: &types.InstanceState{
											Name: "running",
										},
										Tags: []types.Tag{
											{
												Key:   aws.String("Name"),
												Value: aws.String("web02"),
											},
										},
									},
								},
							},
						},
						ResultMetadata: middleware.Metadata{},
					},
				},
				sortPosition: 1,
			},
			want: `+-------+---------------------+----------+---------------+---------------+---------+---------------------------------------+
| NAME  |         ID          |   TYPE   |  PRIVATE IP   |   PUBLIC IP   |  STATE  |            SECURITY GROUP             |
+-------+---------------------+----------+---------------+---------------+---------+---------------------------------------+
| app01 | i-06d4c29e4e5ccadc4 | t2.micro | 172.31.35.175 | 54.238.30.226 | running | launch-wizard-2(sg-0f0b4c4642ffb5ef2) |
| app02 | i-06723a6629e542c50 | t3.small | 172.31.21.33  | 18.179.33.182 | running | launch-wizard-3(sg-08d35fef29987e75e) |
| web01 | i-0abee92626b0a28a7 | t3.nano  | 172.31.18.8   | 35.73.127.100 | running | launch-wizard-1(sg-0d642190887707fd0) |
| web02 | i-06d4c29e4e5edx5tf | t2.micro | 172.31.35.130 | 54.238.30.211 | running | launch-wizard-2(sg-0f0b4c4642ffb5ef2) |
+-------+---------------------+----------+---------------+---------------+---------+---------------------------------------+
`,
		},
	}
	for _, tt := range tests {
		var buf bytes.Buffer
		t.Run(tt.name, func(t *testing.T) {
			showEc2Instances(tt.args.outputs, tablewriter.NewWriter(&buf), tt.args.sortPosition)
			if buf.String() != tt.want {
				t.Errorf("failed to test: %s\n", tt.name)
				t.Errorf("\nwant:\n%s\n", tt.want)
				t.Errorf("\ninput:\n%s\n", buf.String())
			}
		})
	}
}
