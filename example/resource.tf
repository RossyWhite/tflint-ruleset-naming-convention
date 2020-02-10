resource "aws_instance" "web" {
  ami = "ami-07296175bc6b826a5"
  instance_type = "c5.large"
}

resource "aws_sns_topic" "web" {
  name = "my-topik" # <= mistake
}
