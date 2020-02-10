# tflint-naming-convention

`tflint-naming-convention` is a [tflint](https://github.com/terraform-linters/tflint) plugin which check whether an attribute of resource match the given naming convention.


## installation

```bash
$ git clone https://github.com/RossyWhite/tflint-naming-convention
$ make install
```

## Usage

### 1. Create `conventions.json`

`conventions.json` defines a set of naming rules which is applied to your terraform resources.
The rules can be written by using regex.
The config file must be placed in `~/.tflint.d/configs/` or `./.tflint.d/configs/`  
(working directory when `tflint` command is executed)  

example is available [here](https://github.com/RossyWhite/tflint-naming-convention/blob/master/example/.tflint.d/configs/conventions.json). 

### 2. Edit `.tflint.hcl`

Add following snippet to `.tflint.hcl`

```hcl
plugin "naming_convention" {
  enabled = true
}
```

## 3. Run tflint !

Just run `tflint`. an example output is shown below

```bash
Error: aws_sns_topic.name does not match pattern `.*-topic$` (one_name)

  on resource.tf line 7:
   7:   name = "my-topik" # <= mistake

Reference: https://github.com/RossyWhite/tflint-naming-convention
```
