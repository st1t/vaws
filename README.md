# Vaws

The vaws command was created to simplify the display of AWS resources.  
This repository is a Go version of the command that was created in the following repository.  
https://github.com/st1t/vaws

## Usage

```bash
$ vaws -h
The vaws command was created to simplify the display of AWS resources.

Usage:
  vaws [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  ec2         Show EC2 instances.
  help        Help about any command

Flags:
  -p, --aws-profile string   -p my-aws
  -h, --help                 help for vaws
  -s, --sort-position int    -s 1 (default 1)
  -t, --toggle               Help message for toggle
  -v, --version              version for vaws

Use "vaws [command] --help" for more information about a command.
$
```

## License

The gem is available as open source under the terms of the [MIT License](https://opensource.org/licenses/MIT).
