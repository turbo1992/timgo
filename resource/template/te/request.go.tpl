package request

import "nong-xin-bao/model"

type {{.StructName}}Search struct{
    model.{{.StructName}}
    PageInfo
}