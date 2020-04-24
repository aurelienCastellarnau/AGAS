resource "aws_instance" "tfer--i-002D-02be90f34ce0c67c2_" {
  ami                         = "ami-00077e3fed5089981"
  associate_public_ip_address = "true"
  availability_zone           = "eu-west-3c"
  cpu_core_count              = "1"
  cpu_threads_per_core        = "1"

  credit_specification {
    cpu_credits = "standard"
  }

  disable_api_termination = "false"
  ebs_optimized           = "false"
  get_password_data       = "false"
  hibernation             = "false"
  instance_type           = "t2.micro"
  ipv6_address_count      = "0"
  monitoring              = "false"
  private_ip              = "172.31.35.43"

  root_block_device {
    delete_on_termination = "true"
    encrypted             = "false"
    iops                  = "100"
    volume_size           = "8"
    volume_type           = "gp2"
  }

  security_groups        = ["default"]
  source_dest_check      = "true"
  subnet_id              = "subnet-1aa53f57"
  tenancy                = "default"
  vpc_security_group_ids = ["sg-8f5953e0"]
}

resource "aws_instance" "tfer--i-002D-0327c6bf44b6cbe0f_" {
  ami                         = "ami-00077e3fed5089981"
  associate_public_ip_address = "true"
  availability_zone           = "eu-west-3c"
  cpu_core_count              = "1"
  cpu_threads_per_core        = "1"

  credit_specification {
    cpu_credits = "standard"
  }

  disable_api_termination = "false"
  ebs_optimized           = "false"
  get_password_data       = "false"
  hibernation             = "false"
  instance_type           = "t2.micro"
  ipv6_address_count      = "0"
  monitoring              = "false"
  private_ip              = "172.31.47.80"

  root_block_device {
    delete_on_termination = "true"
    encrypted             = "false"
    iops                  = "100"
    volume_size           = "8"
    volume_type           = "gp2"
  }

  security_groups        = ["default"]
  source_dest_check      = "true"
  subnet_id              = "subnet-1aa53f57"
  tenancy                = "default"
  vpc_security_group_ids = ["sg-8f5953e0"]
}

resource "aws_instance" "tfer--i-002D-03eabb70d388f1379_" {
  ami                         = "ami-00077e3fed5089981"
  associate_public_ip_address = "true"
  availability_zone           = "eu-west-3c"
  cpu_core_count              = "1"
  cpu_threads_per_core        = "1"

  credit_specification {
    cpu_credits = "standard"
  }

  disable_api_termination = "false"
  ebs_optimized           = "false"
  get_password_data       = "false"
  hibernation             = "false"
  instance_type           = "t2.micro"
  ipv6_address_count      = "0"
  monitoring              = "false"
  private_ip              = "172.31.45.23"

  root_block_device {
    delete_on_termination = "true"
    encrypted             = "false"
    iops                  = "100"
    volume_size           = "8"
    volume_type           = "gp2"
  }

  security_groups        = ["default"]
  source_dest_check      = "true"
  subnet_id              = "subnet-1aa53f57"
  tenancy                = "default"
  vpc_security_group_ids = ["sg-8f5953e0"]
}

resource "aws_instance" "tfer--i-002D-06799f8cb392ff374_" {
  ami                         = "ami-00077e3fed5089981"
  associate_public_ip_address = "true"
  availability_zone           = "eu-west-3c"
  cpu_core_count              = "1"
  cpu_threads_per_core        = "1"

  credit_specification {
    cpu_credits = "standard"
  }

  disable_api_termination = "false"
  ebs_optimized           = "false"
  get_password_data       = "false"
  hibernation             = "false"
  instance_type           = "t2.micro"
  ipv6_address_count      = "0"
  monitoring              = "false"
  private_ip              = "172.31.47.92"

  root_block_device {
    delete_on_termination = "true"
    encrypted             = "false"
    iops                  = "100"
    volume_size           = "8"
    volume_type           = "gp2"
  }

  security_groups        = ["default"]
  source_dest_check      = "true"
  subnet_id              = "subnet-1aa53f57"
  tenancy                = "default"
  vpc_security_group_ids = ["sg-8f5953e0"]
}

resource "aws_instance" "tfer--i-002D-06b06652916be3aa0_" {
  ami                         = "ami-00077e3fed5089981"
  associate_public_ip_address = "true"
  availability_zone           = "eu-west-3c"
  cpu_core_count              = "1"
  cpu_threads_per_core        = "1"

  credit_specification {
    cpu_credits = "standard"
  }

  disable_api_termination = "false"
  ebs_optimized           = "false"
  get_password_data       = "false"
  hibernation             = "false"
  instance_type           = "t2.micro"
  ipv6_address_count      = "0"
  monitoring              = "false"
  private_ip              = "172.31.47.230"

  root_block_device {
    delete_on_termination = "true"
    encrypted             = "false"
    iops                  = "100"
    volume_size           = "8"
    volume_type           = "gp2"
  }

  security_groups        = ["default"]
  source_dest_check      = "true"
  subnet_id              = "subnet-1aa53f57"
  tenancy                = "default"
  vpc_security_group_ids = ["sg-8f5953e0"]
}

resource "aws_instance" "tfer--i-002D-08d275addcba84649_" {
  ami                         = "ami-00077e3fed5089981"
  associate_public_ip_address = "true"
  availability_zone           = "eu-west-3c"
  cpu_core_count              = "1"
  cpu_threads_per_core        = "1"

  credit_specification {
    cpu_credits = "standard"
  }

  disable_api_termination = "false"
  ebs_optimized           = "false"
  get_password_data       = "false"
  hibernation             = "false"
  instance_type           = "t2.micro"
  ipv6_address_count      = "0"
  monitoring              = "false"
  private_ip              = "172.31.35.74"

  root_block_device {
    delete_on_termination = "true"
    encrypted             = "false"
    iops                  = "100"
    volume_size           = "8"
    volume_type           = "gp2"
  }

  security_groups        = ["default"]
  source_dest_check      = "true"
  subnet_id              = "subnet-1aa53f57"
  tenancy                = "default"
  vpc_security_group_ids = ["sg-8f5953e0"]
}

resource "aws_instance" "tfer--i-002D-0949ce3b016ffea26_" {
  ami                         = "ami-00077e3fed5089981"
  associate_public_ip_address = "true"
  availability_zone           = "eu-west-3c"
  cpu_core_count              = "1"
  cpu_threads_per_core        = "1"

  credit_specification {
    cpu_credits = "standard"
  }

  disable_api_termination = "false"
  ebs_optimized           = "false"
  get_password_data       = "false"
  hibernation             = "false"
  instance_type           = "t2.micro"
  ipv6_address_count      = "0"
  monitoring              = "false"
  private_ip              = "172.31.37.19"

  root_block_device {
    delete_on_termination = "true"
    encrypted             = "false"
    iops                  = "100"
    volume_size           = "8"
    volume_type           = "gp2"
  }

  security_groups        = ["default"]
  source_dest_check      = "true"
  subnet_id              = "subnet-1aa53f57"
  tenancy                = "default"
  vpc_security_group_ids = ["sg-8f5953e0"]
}

resource "aws_instance" "tfer--i-002D-0c0225562643183a2_" {
  ami                         = "ami-00077e3fed5089981"
  associate_public_ip_address = "true"
  availability_zone           = "eu-west-3c"
  cpu_core_count              = "1"
  cpu_threads_per_core        = "1"

  credit_specification {
    cpu_credits = "standard"
  }

  disable_api_termination = "false"
  ebs_optimized           = "false"
  get_password_data       = "false"
  hibernation             = "false"
  instance_type           = "t2.micro"
  ipv6_address_count      = "0"
  monitoring              = "false"
  private_ip              = "172.31.36.106"

  root_block_device {
    delete_on_termination = "true"
    encrypted             = "false"
    iops                  = "100"
    volume_size           = "8"
    volume_type           = "gp2"
  }

  security_groups        = ["default"]
  source_dest_check      = "true"
  subnet_id              = "subnet-1aa53f57"
  tenancy                = "default"
  vpc_security_group_ids = ["sg-8f5953e0"]
}

resource "aws_instance" "tfer--i-002D-0f419db15a6995e9b_" {
  ami                         = "ami-00077e3fed5089981"
  associate_public_ip_address = "true"
  availability_zone           = "eu-west-3b"
  cpu_core_count              = "1"
  cpu_threads_per_core        = "1"

  credit_specification {
    cpu_credits = "standard"
  }

  disable_api_termination = "false"
  ebs_optimized           = "false"
  get_password_data       = "false"
  hibernation             = "false"
  instance_type           = "t2.micro"
  ipv6_address_count      = "0"
  monitoring              = "false"
  private_ip              = "172.31.27.159"

  root_block_device {
    delete_on_termination = "true"
    encrypted             = "false"
    iops                  = "100"
    volume_size           = "8"
    volume_type           = "gp2"
  }

  security_groups        = ["default"]
  source_dest_check      = "true"
  subnet_id              = "subnet-a5b1e0de"
  tenancy                = "default"
  vpc_security_group_ids = ["sg-8f5953e0"]
}

resource "aws_instance" "tfer--i-002D-0f919bb8c4dfedaed_" {
  ami                         = "ami-00077e3fed5089981"
  associate_public_ip_address = "true"
  availability_zone           = "eu-west-3b"
  cpu_core_count              = "1"
  cpu_threads_per_core        = "1"

  credit_specification {
    cpu_credits = "standard"
  }

  disable_api_termination = "false"
  ebs_optimized           = "false"
  get_password_data       = "false"
  hibernation             = "false"
  instance_type           = "t2.micro"
  ipv6_address_count      = "0"
  monitoring              = "false"
  private_ip              = "172.31.24.50"

  root_block_device {
    delete_on_termination = "true"
    encrypted             = "false"
    iops                  = "100"
    volume_size           = "8"
    volume_type           = "gp2"
  }

  security_groups        = ["default"]
  source_dest_check      = "true"
  subnet_id              = "subnet-a5b1e0de"
  tenancy                = "default"
  vpc_security_group_ids = ["sg-8f5953e0"]
}
