package vaws

import (
	"bytes"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/olekukonko/tablewriter"
	"testing"
)

func Test_showVpc(t *testing.T) {
	type args struct {
		outputs      []*ec2.DescribeVpcsOutput
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
				outputs: []*ec2.DescribeVpcsOutput{
					{
						NextToken: nil,
						Vpcs: []types.Vpc{
							{
								CidrBlock: aws.String("10.1.0.0/16"),
								Tags: []types.Tag{
									{
										Key:   aws.String("Name"),
										Value: aws.String("default vpc"),
									},
								},
								VpcId: aws.String("vpc-123XXXXX"),
							},
						},
					},
					{
						NextToken: nil,
						Vpcs: []types.Vpc{
							{
								CidrBlock: aws.String("10.0.0.0/16"),
								Tags: []types.Tag{
									{
										Key:   aws.String("Name"),
										Value: aws.String("hoge service"),
									},
								},
								VpcId: aws.String("vpc-123ZZZZZ"),
							},
						},
					},
				},
				sortPosition: 3,
			},
			want: `+--------------+--------------+-------------+
|     NAME     |      ID      |    CIDR     |
+--------------+--------------+-------------+
| hoge service | vpc-123ZZZZZ | 10.0.0.0/16 |
| default vpc  | vpc-123XXXXX | 10.1.0.0/16 |
+--------------+--------------+-------------+
`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf bytes.Buffer
			showVpc(tt.args.outputs, tablewriter.NewWriter(&buf), tt.args.sortPosition)
			if buf.String() != tt.want {
				t.Errorf("failed to test: %s\n", tt.name)
				t.Errorf("\nwant:\n%s\n", tt.want)
				t.Errorf("\ninput:\n%s\n", buf.String())
			}
		})
	}
}
