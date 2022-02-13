package vaws

import (
	"bytes"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/olekukonko/tablewriter"
	"testing"
)

func Test_showSubnets(t *testing.T) {
	type args struct {
		outputs      []*ec2.DescribeSubnetsOutput
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
				outputs: []*ec2.DescribeSubnetsOutput{
					{
						Subnets: []types.Subnet{
							{
								AvailabilityZone:        aws.String("ap-northeast-1a"),
								AvailabilityZoneId:      aws.String("apne1-az4"),
								AvailableIpAddressCount: aws.Int32(250),
								CidrBlock:               aws.String("10.3.0.0/24"),
								MapPublicIpOnLaunch:     aws.Bool(false),
								SubnetId:                aws.String("subnet-zzzzzzzz"),
								Tags: []types.Tag{
									{
										Key:   aws.String("Name"),
										Value: aws.String("test-subnet-03"),
									},
								},
								VpcId: aws.String("vpc-12345678"),
							},
						},
					},
					{
						Subnets: []types.Subnet{
							{
								AvailabilityZone:        aws.String("ap-northeast-1a"),
								AvailabilityZoneId:      aws.String("apne1-az4"),
								AvailableIpAddressCount: aws.Int32(250),
								CidrBlock:               aws.String("10.1.0.0/24"),
								MapPublicIpOnLaunch:     aws.Bool(false),
								SubnetId:                aws.String("subnet-yyyyyyyy"),
								Tags: []types.Tag{
									{
										Key:   aws.String("Name"),
										Value: aws.String("test-subnet-01"),
									},
								},
								VpcId: aws.String("vpc-12345678"),
							},
							{
								AvailabilityZone:        aws.String("ap-northeast-1a"),
								AvailabilityZoneId:      aws.String("apne1-az4"),
								AvailableIpAddressCount: aws.Int32(250),
								CidrBlock:               aws.String("10.2.0.0/24"),
								MapPublicIpOnLaunch:     aws.Bool(true),
								SubnetId:                aws.String("subnet-xxxxxxxx"),
								Tags: []types.Tag{
									{
										Key:   aws.String("Name"),
										Value: aws.String("test-subnet-02"),
									},
								},
								VpcId: aws.String("vpc-12345678"),
							},
						},
					},
				},
				sortPosition: 1,
			},
			want: `+----------------+-----------------+-------------+--------------+-----------------+-----------+---------------+--------------------+
|      NAME      |    SUBNET ID    |    CIDR     |     VPC      |       AZ        |   AZ ID   | MAP PUBLIC IP | AVAILABLE IP COUNT |
+----------------+-----------------+-------------+--------------+-----------------+-----------+---------------+--------------------+
| test-subnet-01 | subnet-yyyyyyyy | 10.1.0.0/24 | vpc-12345678 | ap-northeast-1a | apne1-az4 | false         |                250 |
| test-subnet-02 | subnet-xxxxxxxx | 10.2.0.0/24 | vpc-12345678 | ap-northeast-1a | apne1-az4 | true          |                250 |
| test-subnet-03 | subnet-zzzzzzzz | 10.3.0.0/24 | vpc-12345678 | ap-northeast-1a | apne1-az4 | false         |                250 |
+----------------+-----------------+-------------+--------------+-----------------+-----------+---------------+--------------------+
`,
		},
	}
	for _, tt := range tests {
		var buf bytes.Buffer
		t.Run(tt.name, func(t *testing.T) {
			showSubnets(tt.args.outputs, tablewriter.NewWriter(&buf), tt.args.sortPosition)
			if buf.String() != tt.want {
				t.Errorf("failed to test: %s\n", tt.name)
				t.Errorf("\nwant:\n%s\n", tt.want)
				t.Errorf("\ninput:\n%s\n", buf.String())
			}
		})
	}
}
