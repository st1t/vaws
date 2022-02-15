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

func Test_showSecurityGroup(t *testing.T) {
	type args struct {
		outputs      []*ec2.DescribeSecurityGroupsOutput
		table        *tablewriter.Table
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
				outputs: []*ec2.DescribeSecurityGroupsOutput{
					{
						SecurityGroups: []types.SecurityGroup{
							{
								GroupName: aws.String("default"),
								GroupId:   aws.String("sg-0d642190887707fd0"),
								VpcId:     aws.String("vpc-0f9999c7db8c44b21"),
								IpPermissions: []types.IpPermission{
									{
										IpRanges: []types.IpRange{
											{
												CidrIp: aws.String("0.0.0.0/0"),
											},
										},
										ToPort: nil,
									},
								},
							},
							{
								GroupName: aws.String("launch-wizard-1"),
								GroupId:   aws.String("sg-0d642190887707fd0"),
								VpcId:     aws.String("vpc-0f9999c7db8c44b21"),
								IpPermissions: []types.IpPermission{
									{
										IpRanges: []types.IpRange{
											{
												CidrIp: aws.String("0.0.0.0/0"),
											},
										},
										ToPort: aws.Int32(22),
									},
								},
							},
							{
								GroupName: aws.String("launch-wizard-2"),
								GroupId:   aws.String("sg-08d35fef29987e75e"),
								VpcId:     aws.String("vpc-0f9999c7db8c44b21"),
								IpPermissions: []types.IpPermission{
									{
										IpRanges: []types.IpRange{
											{
												CidrIp: aws.String("8.8.8.8/32"),
											},
										},
										ToPort: aws.Int32(22),
									},
								},
							},
							{
								GroupName: aws.String("launch-wizard-2"),
								GroupId:   aws.String("sg-08d35fef29987e75e"),
								VpcId:     aws.String("vpc-0f9999c7db8c44b21"),
								IpPermissions: []types.IpPermission{
									{
										IpRanges: []types.IpRange{
											{
												CidrIp: aws.String("0.0.0.0/0"),
											},
										},
										ToPort: aws.Int32(22),
										UserIdGroupPairs: []types.UserIdGroupPair{
											{
												GroupId: aws.String("sg-0d642190887707fd0"),
											},
										},
									},
								},
							},
							{
								GroupName: aws.String("launch-wizard-2"),
								GroupId:   aws.String("sg-08d35fef29987e75e"),
								VpcId:     aws.String("vpc-0f9999c7db8c44b21"),
								IpPermissions: []types.IpPermission{
									{
										ToPort: aws.Int32(53),
										PrefixListIds: []types.PrefixListId{
											{
												PrefixListId: aws.String("pl-61a12345"),
											},
										},
									},
								},
							},
						},
						ResultMetadata: middleware.Metadata{},
					},
				},
				table:        nil,
				sortPosition: 1,
			},
			want: `+-----------------+---------+----------------------+------+----------------------+-----------------------+
|      NAME       |  TYPE   |          ID          | PORT |        SOURCE        |          VPC          |
+-----------------+---------+----------------------+------+----------------------+-----------------------+
| default         | inbound | sg-0d642190887707fd0 |   -1 | 0.0.0.0/0            | vpc-0f9999c7db8c44b21 |
| launch-wizard-1 | inbound | sg-0d642190887707fd0 |   22 | 0.0.0.0/0            | vpc-0f9999c7db8c44b21 |
| launch-wizard-2 | inbound | sg-08d35fef29987e75e |   22 | 8.8.8.8/32           | vpc-0f9999c7db8c44b21 |
| launch-wizard-2 | inbound | sg-08d35fef29987e75e |   22 | 0.0.0.0/0            | vpc-0f9999c7db8c44b21 |
| launch-wizard-2 | inbound | sg-08d35fef29987e75e |   22 | sg-0d642190887707fd0 | vpc-0f9999c7db8c44b21 |
| launch-wizard-2 | inbound | sg-08d35fef29987e75e |   53 | pl-61a12345          | vpc-0f9999c7db8c44b21 |
+-----------------+---------+----------------------+------+----------------------+-----------------------+
`,
		},
		{
			name: "classic ec2",
			args: args{
				outputs: []*ec2.DescribeSecurityGroupsOutput{
					{
						SecurityGroups: []types.SecurityGroup{
							{
								GroupName: aws.String("default"),
								GroupId:   aws.String("sg-0d642190887707fd0"),
								VpcId:     nil,
								IpPermissions: []types.IpPermission{
									{
										IpRanges: []types.IpRange{
											{
												CidrIp: aws.String("0.0.0.0/0"),
											},
										},
										ToPort: nil,
									},
								},
							},
							{
								GroupName: aws.String("launch-wizard-1"),
								GroupId:   aws.String("sg-0d642190887707fd0"),
								VpcId:     aws.String("vpc-0f9999c7db8c44b21"),
								IpPermissions: []types.IpPermission{
									{
										IpRanges: []types.IpRange{
											{
												CidrIp: aws.String("0.0.0.0/0"),
											},
										},
										ToPort: aws.Int32(22),
									},
								},
							},
						},
						ResultMetadata: middleware.Metadata{},
					},
				},
				table:        nil,
				sortPosition: 1,
			},
			want: `+-----------------+---------+----------------------+------+-----------+-----------------------+
|      NAME       |  TYPE   |          ID          | PORT |  SOURCE   |          VPC          |
+-----------------+---------+----------------------+------+-----------+-----------------------+
| default         | inbound | sg-0d642190887707fd0 |   -1 | 0.0.0.0/0 | none                  |
| launch-wizard-1 | inbound | sg-0d642190887707fd0 |   22 | 0.0.0.0/0 | vpc-0f9999c7db8c44b21 |
+-----------------+---------+----------------------+------+-----------+-----------------------+
`,
		},
		{
			name: "sort_port",
			args: args{
				outputs: []*ec2.DescribeSecurityGroupsOutput{
					{
						SecurityGroups: []types.SecurityGroup{
							{
								GroupName: aws.String("default"),
								GroupId:   aws.String("sg-0d642190887707fd0"),
								VpcId:     aws.String("vpc-0f9999c7db8c44b21"),
								IpPermissions: []types.IpPermission{
									{
										IpRanges: []types.IpRange{
											{
												CidrIp: aws.String("0.0.0.0/0"),
											},
										},
										ToPort: nil,
									},
								},
							},
							{
								GroupName: aws.String("launch-wizard-1"),
								GroupId:   aws.String("sg-0d642190887707fd0"),
								VpcId:     aws.String("vpc-0f9999c7db8c44b21"),
								IpPermissions: []types.IpPermission{
									{
										IpRanges: []types.IpRange{
											{
												CidrIp: aws.String("0.0.0.0/0"),
											},
										},
										ToPort: aws.Int32(22),
									},
								},
							},
							{
								GroupName: aws.String("launch-wizard-2"),
								GroupId:   aws.String("sg-08d35fef29987e75e"),
								VpcId:     aws.String("vpc-0f9999c7db8c44b21"),
								IpPermissions: []types.IpPermission{
									{
										IpRanges: []types.IpRange{
											{
												CidrIp: aws.String("8.8.8.8/32"),
											},
										},
										ToPort: aws.Int32(22),
									},
								},
							},
							{
								GroupName: aws.String("launch-wizard-2"),
								GroupId:   aws.String("sg-08d35fef29987e75e"),
								VpcId:     aws.String("vpc-0f9999c7db8c44b21"),
								IpPermissions: []types.IpPermission{
									{
										IpRanges: []types.IpRange{
											{
												CidrIp: aws.String("0.0.0.0/0"),
											},
										},
										ToPort: aws.Int32(22),
										UserIdGroupPairs: []types.UserIdGroupPair{
											{
												GroupId: aws.String("sg-0d642190887707fd0"),
											},
										},
									},
								},
							},
							{
								GroupName: aws.String("launch-wizard-2"),
								GroupId:   aws.String("sg-08d35fef29987e75e"),
								VpcId:     aws.String("vpc-0f9999c7db8c44b21"),
								IpPermissions: []types.IpPermission{
									{
										ToPort: aws.Int32(53),
										PrefixListIds: []types.PrefixListId{
											{
												PrefixListId: aws.String("pl-61a12345"),
											},
										},
									},
								},
							},
						},
						ResultMetadata: middleware.Metadata{},
					},
				},
				table:        nil,
				sortPosition: 5,
			},
			want: `+-----------------+---------+----------------------+------+----------------------+-----------------------+
|      NAME       |  TYPE   |          ID          | PORT |        SOURCE        |          VPC          |
+-----------------+---------+----------------------+------+----------------------+-----------------------+
| default         | inbound | sg-0d642190887707fd0 |   -1 | 0.0.0.0/0            | vpc-0f9999c7db8c44b21 |
| launch-wizard-1 | inbound | sg-0d642190887707fd0 |   22 | 0.0.0.0/0            | vpc-0f9999c7db8c44b21 |
| launch-wizard-2 | inbound | sg-08d35fef29987e75e |   22 | 0.0.0.0/0            | vpc-0f9999c7db8c44b21 |
| launch-wizard-2 | inbound | sg-08d35fef29987e75e |   22 | 8.8.8.8/32           | vpc-0f9999c7db8c44b21 |
| launch-wizard-2 | inbound | sg-08d35fef29987e75e |   53 | pl-61a12345          | vpc-0f9999c7db8c44b21 |
| launch-wizard-2 | inbound | sg-08d35fef29987e75e |   22 | sg-0d642190887707fd0 | vpc-0f9999c7db8c44b21 |
+-----------------+---------+----------------------+------+----------------------+-----------------------+
`,
		},
		{
			name: "two request",
			args: args{
				outputs: []*ec2.DescribeSecurityGroupsOutput{
					{
						SecurityGroups: []types.SecurityGroup{
							{
								GroupName: aws.String("default"),
								GroupId:   aws.String("sg-0d642190887707fd0"),
								VpcId:     aws.String("vpc-0f9999c7db8c44b21"),
								IpPermissions: []types.IpPermission{
									{
										IpRanges: []types.IpRange{
											{
												CidrIp: aws.String("0.0.0.0/0"),
											},
										},
										ToPort: nil,
									},
								},
							},
							{
								GroupName: aws.String("launch-wizard-1"),
								GroupId:   aws.String("sg-0d642190887707fd0"),
								VpcId:     aws.String("vpc-0f9999c7db8c44b21"),
								IpPermissions: []types.IpPermission{
									{
										IpRanges: []types.IpRange{
											{
												CidrIp: aws.String("0.0.0.0/0"),
											},
										},
										ToPort: aws.Int32(22),
									},
								},
							},
							{
								GroupName: aws.String("launch-wizard-2"),
								GroupId:   aws.String("sg-08d35fef29987e75e"),
								VpcId:     aws.String("vpc-0f9999c7db8c44b21"),
								IpPermissions: []types.IpPermission{
									{
										IpRanges: []types.IpRange{
											{
												CidrIp: aws.String("8.8.8.8/32"),
											},
										},
										ToPort: aws.Int32(22),
									},
								},
							},
							{
								GroupName: aws.String("launch-wizard-2"),
								GroupId:   aws.String("sg-08d35fef29987e75e"),
								VpcId:     aws.String("vpc-0f9999c7db8c44b21"),
								IpPermissions: []types.IpPermission{
									{
										IpRanges: []types.IpRange{
											{
												CidrIp: aws.String("0.0.0.0/0"),
											},
										},
										ToPort: aws.Int32(22),
										UserIdGroupPairs: []types.UserIdGroupPair{
											{
												GroupId: aws.String("sg-0d642190887707fd0"),
											},
										},
									},
								},
							},
							{
								GroupName: aws.String("launch-wizard-2"),
								GroupId:   aws.String("sg-08d35fef29987e75e"),
								VpcId:     aws.String("vpc-0f9999c7db8c44b21"),
								IpPermissions: []types.IpPermission{
									{
										ToPort: aws.Int32(53),
										PrefixListIds: []types.PrefixListId{
											{
												PrefixListId: aws.String("pl-61a12345"),
											},
										},
									},
								},
							},
						},
						ResultMetadata: middleware.Metadata{},
					},
					{
						SecurityGroups: []types.SecurityGroup{
							{
								GroupName: aws.String("default"),
								GroupId:   aws.String("sg-0d642190887707fd0"),
								VpcId:     aws.String("vpc-0f9999c7db8c44b21"),
								IpPermissions: []types.IpPermission{
									{
										IpRanges: []types.IpRange{
											{
												CidrIp: aws.String("0.0.0.0/0"),
											},
										},
										ToPort: nil,
									},
								},
							},
							{
								GroupName: aws.String("launch-wizard-1"),
								GroupId:   aws.String("sg-0d642190887707fd0"),
								VpcId:     aws.String("vpc-0f9999c7db8c44b21"),
								IpPermissions: []types.IpPermission{
									{
										IpRanges: []types.IpRange{
											{
												CidrIp: aws.String("0.0.0.0/0"),
											},
										},
										ToPort: aws.Int32(22),
									},
								},
							},
							{
								GroupName: aws.String("launch-wizard-2"),
								GroupId:   aws.String("sg-08d35fef29987e75e"),
								VpcId:     aws.String("vpc-0f9999c7db8c44b21"),
								IpPermissions: []types.IpPermission{
									{
										IpRanges: []types.IpRange{
											{
												CidrIp: aws.String("8.8.8.8/32"),
											},
										},
										ToPort: aws.Int32(22),
									},
								},
							},
							{
								GroupName: aws.String("launch-wizard-2"),
								GroupId:   aws.String("sg-08d35fef29987e75e"),
								VpcId:     aws.String("vpc-0f9999c7db8c44b21"),
								IpPermissions: []types.IpPermission{
									{
										IpRanges: []types.IpRange{
											{
												CidrIp: aws.String("0.0.0.0/0"),
											},
										},
										ToPort: aws.Int32(22),
										UserIdGroupPairs: []types.UserIdGroupPair{
											{
												GroupId: aws.String("sg-0d642190887707fd0"),
											},
										},
									},
								},
							},
							{
								GroupName: aws.String("launch-wizard-2"),
								GroupId:   aws.String("sg-08d35fef29987e75e"),
								VpcId:     aws.String("vpc-0f9999c7db8c44b21"),
								IpPermissions: []types.IpPermission{
									{
										ToPort: aws.Int32(53),
										PrefixListIds: []types.PrefixListId{
											{
												PrefixListId: aws.String("pl-61a12345"),
											},
										},
									},
								},
							},
						},
						ResultMetadata: middleware.Metadata{},
					},
				},
				table:        nil,
				sortPosition: 1,
			},
			want: `+-----------------+---------+----------------------+------+----------------------+-----------------------+
|      NAME       |  TYPE   |          ID          | PORT |        SOURCE        |          VPC          |
+-----------------+---------+----------------------+------+----------------------+-----------------------+
| default         | inbound | sg-0d642190887707fd0 |   -1 | 0.0.0.0/0            | vpc-0f9999c7db8c44b21 |
| default         | inbound | sg-0d642190887707fd0 |   -1 | 0.0.0.0/0            | vpc-0f9999c7db8c44b21 |
| launch-wizard-1 | inbound | sg-0d642190887707fd0 |   22 | 0.0.0.0/0            | vpc-0f9999c7db8c44b21 |
| launch-wizard-1 | inbound | sg-0d642190887707fd0 |   22 | 0.0.0.0/0            | vpc-0f9999c7db8c44b21 |
| launch-wizard-2 | inbound | sg-08d35fef29987e75e |   22 | 8.8.8.8/32           | vpc-0f9999c7db8c44b21 |
| launch-wizard-2 | inbound | sg-08d35fef29987e75e |   22 | 0.0.0.0/0            | vpc-0f9999c7db8c44b21 |
| launch-wizard-2 | inbound | sg-08d35fef29987e75e |   22 | sg-0d642190887707fd0 | vpc-0f9999c7db8c44b21 |
| launch-wizard-2 | inbound | sg-08d35fef29987e75e |   53 | pl-61a12345          | vpc-0f9999c7db8c44b21 |
| launch-wizard-2 | inbound | sg-08d35fef29987e75e |   22 | 8.8.8.8/32           | vpc-0f9999c7db8c44b21 |
| launch-wizard-2 | inbound | sg-08d35fef29987e75e |   22 | 0.0.0.0/0            | vpc-0f9999c7db8c44b21 |
| launch-wizard-2 | inbound | sg-08d35fef29987e75e |   22 | sg-0d642190887707fd0 | vpc-0f9999c7db8c44b21 |
| launch-wizard-2 | inbound | sg-08d35fef29987e75e |   53 | pl-61a12345          | vpc-0f9999c7db8c44b21 |
+-----------------+---------+----------------------+------+----------------------+-----------------------+
`,
		},
	}
	for _, tt := range tests {
		var buf bytes.Buffer
		t.Run(tt.name, func(t *testing.T) {
			showSecurityGroup(tt.args.outputs, tablewriter.NewWriter(&buf), tt.args.sortPosition)
			if buf.String() != tt.want {
				t.Errorf("failed to test: %s\n", tt.name)
				t.Errorf("\nwant:\n%s\n", tt.want)
				t.Errorf("\ninput:\n%s\n", buf.String())
			}
		})
	}
}
