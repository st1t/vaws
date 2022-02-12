package vaws

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
	"os"
	"sort"
	"strconv"
)

const sgMaxResult = 1000

// securityGroupCmd represents the securityGroup command
var securityGroupCmd = &cobra.Command{
	Use:   "sg",
	Short: "Show Security Group",
	Long:  `Show Security Group`,
	Run: func(cmd *cobra.Command, args []string) {
		profile, err := cmd.Flags().GetString("aws-profile")
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		if profile != "" {
			err := os.Setenv("AWS_PROFILE", profile)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
		}
		output, err := getSecurityGroups(newAwsConfig())
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		sortPosition, err := cmd.Flags().GetInt("sort-position")
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		err = showSecurityGroup(output, tablewriter.NewWriter(os.Stdout), sortPosition)
	},
}

func init() {
	rootCmd.AddCommand(securityGroupCmd)
}

func getSecurityGroups(cfg aws.Config) ([]*ec2.DescribeSecurityGroupsOutput, error) {
	var outputs []*ec2.DescribeSecurityGroupsOutput
	var err error
	client := ec2.NewFromConfig(cfg)
	output := &ec2.DescribeSecurityGroupsOutput{
		NextToken: aws.String(""),
	}
	for output.NextToken != nil {
		output, err = client.DescribeSecurityGroups(context.TODO(), &ec2.DescribeSecurityGroupsInput{
			MaxResults: aws.Int32(sgMaxResult),
			NextToken:  output.NextToken,
		})
		if err != nil {
			return nil, err
		}
		outputs = append(outputs, output)
	}
	return outputs, nil
}

func showSecurityGroup(outputs []*ec2.DescribeSecurityGroupsOutput, table *tablewriter.Table, sortPosition int) error {
	header := []string{"NAME", "TYPE", "ID", "PORT", "SOURCE", "VPC"}
	if sortPosition > len(header) || 1 > sortPosition {
		return fmt.Errorf("out of sort range number when using --sort option")
	}
	recordIndex := sortPosition - 1
	table.SetHeader(header)
	var records [][]string
	var allowPort int32
	for _, o := range outputs {
		for _, sg := range o.SecurityGroups {
			for _, in := range sg.IpPermissions {
				allowType := "inbound"
				if in.ToPort != nil {
					allowPort = *in.ToPort
				} else {
					allowPort = -1
				}
				if in.IpRanges != nil {
					for _, v := range in.IpRanges {
						records = append(records, []string{
							*sg.GroupName,
							allowType,
							*sg.GroupId,
							strconv.Itoa(int(allowPort)),
							*v.CidrIp,
							*sg.VpcId,
						})
					}
				}
				if in.PrefixListIds != nil {
					for _, prefix := range in.PrefixListIds {
						records = append(records, []string{
							*sg.GroupName,
							allowType,
							*sg.GroupId,
							strconv.Itoa(int(allowPort)),
							*prefix.PrefixListId,
							*sg.VpcId,
						})
					}
				}
				if in.UserIdGroupPairs != nil {
					for _, v := range in.UserIdGroupPairs {
						records = append(records, []string{
							*sg.GroupName,
							allowType,
							*sg.GroupId,
							strconv.Itoa(int(allowPort)),
							*v.GroupId,
							*sg.VpcId,
						})
					}
				}
			}
		}
	}
	sort.Slice(records, func(i, j int) bool { return records[i][recordIndex] < records[j][recordIndex] })
	table.AppendBulk(records)
	table.Render()
	return nil
}
