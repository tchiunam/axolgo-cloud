package ec2

import "github.com/tchiunam/axolgo-lib/types"

var SecurityGroupIDTemplateString string = types.AxolVarTemplateStringFieldWithKeyValue.New("SecurityGroups", "GroupId", "GroupName")

var TagsTemplateString string = types.AxolVarTemplateStringFieldWithKeyValue.New("Tags", "Key", "Value")
