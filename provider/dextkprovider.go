package provider

import (
	"log"

	"github.com/csnewman/dextk"
	"github.com/gookit/goutil/dump"
	"golang.org/x/exp/mmap"
)

type DexTkProvider struct {
	IProvider
	r *dextk.Reader
}

func (d *DexTkProvider) ListClasses() *[]ClassDef {
	arr := make([]ClassDef, 0)
	ci := d.r.ClassIter()
	for ci.HasNext() {
		node, err := ci.Next()
		if err != nil {
			log.Panicln(err)
		}
		item := ClassDef{
			Id:    node.Id,
			Name:  node.Name.Parsed,
			Flags: node.AccessFlags,
		}
		dump.Println(item)
		arr = append(arr, item)
	}

	return &arr
}

func (d *DexTkProvider) ListMethods(id ClassDefId) *[]MethodDef {
	arr := make([]MethodDef, 0)
	ci := d.r.ClassIter()
	for ci.HasNext() {
		node, err := ci.Next()
		if err != nil {
			log.Panicln(err)
		}
		for _, method := range node.DirectMethods {
			params := make([]MethodParamDef, len(method.Params))
			for i, param := range method.Params {
				params[i] = MethodParamDef{
					Name: param.ClassName.Parsed,
					Type: param.Type,
				}
			}
			item := MethodDef{
				Id:     method.Id,
				Name:   method.Name.Parsed,
				Flags:  method.AccessFlags,
				Params: params,
				Ret:    method.ReturnType.ClassName.Parsed,
			}
			arr = append(arr, item)
		}
		for _, method := range node.VirtualMethods {
			params := make([]MethodParamDef, len(method.Params))
			for i, param := range method.Params {
				params[i] = MethodParamDef{
					Name: param.ClassName.Parsed,
					Type: param.Type,
				}
			}
			item := MethodDef{
				Id:     method.Id,
				Name:   method.Name.Parsed,
				Flags:  method.AccessFlags,
				Params: params,
				Ret:    method.ReturnType.ClassName.Parsed,
			}
			arr = append(arr, item)
		}
	}
	return &arr
}

func NewDexTkProvider(file string) (*DexTkProvider, error) {
	fd, err := mmap.Open(file)
	if err != nil {
		return nil, err
	}
	r, err := dextk.Read(fd)
	if err != nil {
		return nil, err
	}

	return &DexTkProvider{
		r: r,
	}, nil
}
