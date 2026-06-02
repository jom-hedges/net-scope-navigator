# net-scope
A lightweight network visibility tool built in Go and provisioned with Terraform.

net-scope is the observability plane for your internal developer platform. In net-scope's
initial stage, the user can inspect subnets, routing tables, security groups, and flow logs directly from the terminal without opening the AWS Management Console.

## Why build net-scope
When building on AWS, understanding events at the network layer is one of the hardest parts. net-scope makes your VPC layout and traffic behavior observable from the command-line, closing the gap between the infrastructure provisioned and the network behavior.

## ROADMAP
- [ ] VPC layout display
- [ ] Route table inspection
- [ ] Security group rule querying
- [ ] Flow log filterning and display
- [ ] Export to JSON for use in other tools

