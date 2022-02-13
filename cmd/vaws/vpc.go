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
)

const vpcMaxResult = 50

// vpcCmd represents the vpc command
var vpcCmd = &cobra.Command{
	Use:   "vpc",
	Short: "Show VPC",
	Long:  `Show VPC`,
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
		output, err := getVpc(newAwsConfig())
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		sortPosition, err := cmd.Flags().GetInt("sort-position")
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		showVpc(output, tablewriter.NewWriter(os.Stdout), sortPosition)
	},
}

func init() {
	rootCmd.AddCommand(vpcCmd)
}

func getVpc(cfg aws.Config) ([]*ec2.DescribeVpcsOutput, error) {
	var outputs []*ec2.DescribeVpcsOutput
	var err error
	client := ec2.NewFromConfig(cfg)
	// The DescribeVpcs API executes the API once at the beginning because the NextToken "" is disallowed
	output, err := client.DescribeVpcs(context.TODO(), &ec2.DescribeVpcsInput{
		MaxResults: aws.Int32(vpcMaxResult),
	})
	if err != nil {
		return nil, err
	}
	outputs = append(outputs, output)
	for output.NextToken != nil {
		output, err = client.DescribeVpcs(context.TODO(), &ec2.DescribeVpcsInput{
			MaxResults: aws.Int32(vpcMaxResult),
			NextToken:  output.NextToken,
		})
		if err != nil {
			return nil, err
		}
		outputs = append(outputs, output)
	}
	return outputs, nil
}

func showVpc(outputs []*ec2.DescribeVpcsOutput, table *tablewriter.Table, sortPosition int) error {
	header := []string{"NAME", "ID", "CIDR"}
	if sortPosition > len(header) || 1 > sortPosition {
		return fmt.Errorf("out of sort range number when using --sort option")
	}
	recordIndex := sortPosition - 1
	table.SetHeader(header)
	var records [][]string
	for _, o := range outputs {
		for _, v := range o.Vpcs {
			id := v.VpcId
			name := ""
			for _, tag := range v.Tags {
				if *tag.Key == "Name" {
					name = *tag.Value
				}
			}
			cidr := v.CidrBlock
			records = append(records, []string{name, *id, *cidr})
		}
	}
	sort.Slice(records, func(i, j int) bool { return records[i][recordIndex] < records[j][recordIndex] })
	table.AppendBulk(records)
	table.Render()
	return nil
}
