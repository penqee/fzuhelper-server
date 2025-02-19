/*
Copyright 2024 The west2-online Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by thriftgo (0.3.18). DO NOT EDIT.

package course

import (
	"context"
	"fmt"
	"strings"

	thrift "github.com/cloudwego/kitex/pkg/protocol/bthrift/apache"

	"github.com/west2-online/fzuhelper-server/kitex_gen/model"
)

type CourseListRequest struct {
	LoginData *model.LoginData `thrift:"loginData,1,required" frugal:"1,required,model.LoginData" json:"loginData"`
	Term      string           `thrift:"term,2,required" frugal:"2,required,string" json:"term"`
}

func NewCourseListRequest() *CourseListRequest {
	return &CourseListRequest{}
}

func (p *CourseListRequest) InitDefault() {
}

var CourseListRequest_LoginData_DEFAULT *model.LoginData

func (p *CourseListRequest) GetLoginData() (v *model.LoginData) {
	if !p.IsSetLoginData() {
		return CourseListRequest_LoginData_DEFAULT
	}
	return p.LoginData
}

func (p *CourseListRequest) GetTerm() (v string) {
	return p.Term
}
func (p *CourseListRequest) SetLoginData(val *model.LoginData) {
	p.LoginData = val
}
func (p *CourseListRequest) SetTerm(val string) {
	p.Term = val
}

var fieldIDToName_CourseListRequest = map[int16]string{
	1: "loginData",
	2: "term",
}

func (p *CourseListRequest) IsSetLoginData() bool {
	return p.LoginData != nil
}

func (p *CourseListRequest) Read(iprot thrift.TProtocol) (err error) {

	var fieldTypeId thrift.TType
	var fieldId int16
	var issetLoginData bool = false
	var issetTerm bool = false

	if _, err = iprot.ReadStructBegin(); err != nil {
		goto ReadStructBeginError
	}

	for {
		_, fieldTypeId, fieldId, err = iprot.ReadFieldBegin()
		if err != nil {
			goto ReadFieldBeginError
		}
		if fieldTypeId == thrift.STOP {
			break
		}

		switch fieldId {
		case 1:
			if fieldTypeId == thrift.STRUCT {
				if err = p.ReadField1(iprot); err != nil {
					goto ReadFieldError
				}
				issetLoginData = true
			} else if err = iprot.Skip(fieldTypeId); err != nil {
				goto SkipFieldError
			}
		case 2:
			if fieldTypeId == thrift.STRING {
				if err = p.ReadField2(iprot); err != nil {
					goto ReadFieldError
				}
				issetTerm = true
			} else if err = iprot.Skip(fieldTypeId); err != nil {
				goto SkipFieldError
			}
		default:
			if err = iprot.Skip(fieldTypeId); err != nil {
				goto SkipFieldError
			}
		}
		if err = iprot.ReadFieldEnd(); err != nil {
			goto ReadFieldEndError
		}
	}
	if err = iprot.ReadStructEnd(); err != nil {
		goto ReadStructEndError
	}

	if !issetLoginData {
		fieldId = 1
		goto RequiredFieldNotSetError
	}

	if !issetTerm {
		fieldId = 2
		goto RequiredFieldNotSetError
	}
	return nil
ReadStructBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read struct begin error: ", p), err)
ReadFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read field %d begin error: ", p, fieldId), err)
ReadFieldError:
	return thrift.PrependError(fmt.Sprintf("%T read field %d '%s' error: ", p, fieldId, fieldIDToName_CourseListRequest[fieldId]), err)
SkipFieldError:
	return thrift.PrependError(fmt.Sprintf("%T field %d skip type %d error: ", p, fieldId, fieldTypeId), err)

ReadFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T read field end error", p), err)
ReadStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
RequiredFieldNotSetError:
	return thrift.NewTProtocolExceptionWithType(thrift.INVALID_DATA, fmt.Errorf("required field %s is not set", fieldIDToName_CourseListRequest[fieldId]))
}

func (p *CourseListRequest) ReadField1(iprot thrift.TProtocol) error {
	_field := model.NewLoginData()
	if err := _field.Read(iprot); err != nil {
		return err
	}
	p.LoginData = _field
	return nil
}
func (p *CourseListRequest) ReadField2(iprot thrift.TProtocol) error {

	var _field string
	if v, err := iprot.ReadString(); err != nil {
		return err
	} else {
		_field = v
	}
	p.Term = _field
	return nil
}

func (p *CourseListRequest) Write(oprot thrift.TProtocol) (err error) {

	var fieldId int16
	if err = oprot.WriteStructBegin("CourseListRequest"); err != nil {
		goto WriteStructBeginError
	}
	if p != nil {
		if err = p.writeField1(oprot); err != nil {
			fieldId = 1
			goto WriteFieldError
		}
		if err = p.writeField2(oprot); err != nil {
			fieldId = 2
			goto WriteFieldError
		}
	}
	if err = oprot.WriteFieldStop(); err != nil {
		goto WriteFieldStopError
	}
	if err = oprot.WriteStructEnd(); err != nil {
		goto WriteStructEndError
	}
	return nil
WriteStructBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
WriteFieldError:
	return thrift.PrependError(fmt.Sprintf("%T write field %d error: ", p, fieldId), err)
WriteFieldStopError:
	return thrift.PrependError(fmt.Sprintf("%T write field stop error: ", p), err)
WriteStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T write struct end error: ", p), err)
}

func (p *CourseListRequest) writeField1(oprot thrift.TProtocol) (err error) {
	if err = oprot.WriteFieldBegin("loginData", thrift.STRUCT, 1); err != nil {
		goto WriteFieldBeginError
	}
	if err := p.LoginData.Write(oprot); err != nil {
		return err
	}
	if err = oprot.WriteFieldEnd(); err != nil {
		goto WriteFieldEndError
	}
	return nil
WriteFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write field 1 begin error: ", p), err)
WriteFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T write field 1 end error: ", p), err)
}

func (p *CourseListRequest) writeField2(oprot thrift.TProtocol) (err error) {
	if err = oprot.WriteFieldBegin("term", thrift.STRING, 2); err != nil {
		goto WriteFieldBeginError
	}
	if err := oprot.WriteString(p.Term); err != nil {
		return err
	}
	if err = oprot.WriteFieldEnd(); err != nil {
		goto WriteFieldEndError
	}
	return nil
WriteFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write field 2 begin error: ", p), err)
WriteFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T write field 2 end error: ", p), err)
}

func (p *CourseListRequest) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("CourseListRequest(%+v)", *p)

}

func (p *CourseListRequest) DeepEqual(ano *CourseListRequest) bool {
	if p == ano {
		return true
	} else if p == nil || ano == nil {
		return false
	}
	if !p.Field1DeepEqual(ano.LoginData) {
		return false
	}
	if !p.Field2DeepEqual(ano.Term) {
		return false
	}
	return true
}

func (p *CourseListRequest) Field1DeepEqual(src *model.LoginData) bool {

	if !p.LoginData.DeepEqual(src) {
		return false
	}
	return true
}
func (p *CourseListRequest) Field2DeepEqual(src string) bool {

	if strings.Compare(p.Term, src) != 0 {
		return false
	}
	return true
}

type CourseListResponse struct {
	Base *model.BaseResp `thrift:"base,1,required" frugal:"1,required,model.BaseResp" json:"base"`
	Data []*model.Course `thrift:"data,2,required" frugal:"2,required,list<model.Course>" json:"data"`
}

func NewCourseListResponse() *CourseListResponse {
	return &CourseListResponse{}
}

func (p *CourseListResponse) InitDefault() {
}

var CourseListResponse_Base_DEFAULT *model.BaseResp

func (p *CourseListResponse) GetBase() (v *model.BaseResp) {
	if !p.IsSetBase() {
		return CourseListResponse_Base_DEFAULT
	}
	return p.Base
}

func (p *CourseListResponse) GetData() (v []*model.Course) {
	return p.Data
}
func (p *CourseListResponse) SetBase(val *model.BaseResp) {
	p.Base = val
}
func (p *CourseListResponse) SetData(val []*model.Course) {
	p.Data = val
}

var fieldIDToName_CourseListResponse = map[int16]string{
	1: "base",
	2: "data",
}

func (p *CourseListResponse) IsSetBase() bool {
	return p.Base != nil
}

func (p *CourseListResponse) Read(iprot thrift.TProtocol) (err error) {

	var fieldTypeId thrift.TType
	var fieldId int16
	var issetBase bool = false
	var issetData bool = false

	if _, err = iprot.ReadStructBegin(); err != nil {
		goto ReadStructBeginError
	}

	for {
		_, fieldTypeId, fieldId, err = iprot.ReadFieldBegin()
		if err != nil {
			goto ReadFieldBeginError
		}
		if fieldTypeId == thrift.STOP {
			break
		}

		switch fieldId {
		case 1:
			if fieldTypeId == thrift.STRUCT {
				if err = p.ReadField1(iprot); err != nil {
					goto ReadFieldError
				}
				issetBase = true
			} else if err = iprot.Skip(fieldTypeId); err != nil {
				goto SkipFieldError
			}
		case 2:
			if fieldTypeId == thrift.LIST {
				if err = p.ReadField2(iprot); err != nil {
					goto ReadFieldError
				}
				issetData = true
			} else if err = iprot.Skip(fieldTypeId); err != nil {
				goto SkipFieldError
			}
		default:
			if err = iprot.Skip(fieldTypeId); err != nil {
				goto SkipFieldError
			}
		}
		if err = iprot.ReadFieldEnd(); err != nil {
			goto ReadFieldEndError
		}
	}
	if err = iprot.ReadStructEnd(); err != nil {
		goto ReadStructEndError
	}

	if !issetBase {
		fieldId = 1
		goto RequiredFieldNotSetError
	}

	if !issetData {
		fieldId = 2
		goto RequiredFieldNotSetError
	}
	return nil
ReadStructBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read struct begin error: ", p), err)
ReadFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read field %d begin error: ", p, fieldId), err)
ReadFieldError:
	return thrift.PrependError(fmt.Sprintf("%T read field %d '%s' error: ", p, fieldId, fieldIDToName_CourseListResponse[fieldId]), err)
SkipFieldError:
	return thrift.PrependError(fmt.Sprintf("%T field %d skip type %d error: ", p, fieldId, fieldTypeId), err)

ReadFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T read field end error", p), err)
ReadStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
RequiredFieldNotSetError:
	return thrift.NewTProtocolExceptionWithType(thrift.INVALID_DATA, fmt.Errorf("required field %s is not set", fieldIDToName_CourseListResponse[fieldId]))
}

func (p *CourseListResponse) ReadField1(iprot thrift.TProtocol) error {
	_field := model.NewBaseResp()
	if err := _field.Read(iprot); err != nil {
		return err
	}
	p.Base = _field
	return nil
}
func (p *CourseListResponse) ReadField2(iprot thrift.TProtocol) error {
	_, size, err := iprot.ReadListBegin()
	if err != nil {
		return err
	}
	_field := make([]*model.Course, 0, size)
	values := make([]model.Course, size)
	for i := 0; i < size; i++ {
		_elem := &values[i]
		_elem.InitDefault()

		if err := _elem.Read(iprot); err != nil {
			return err
		}

		_field = append(_field, _elem)
	}
	if err := iprot.ReadListEnd(); err != nil {
		return err
	}
	p.Data = _field
	return nil
}

func (p *CourseListResponse) Write(oprot thrift.TProtocol) (err error) {

	var fieldId int16
	if err = oprot.WriteStructBegin("CourseListResponse"); err != nil {
		goto WriteStructBeginError
	}
	if p != nil {
		if err = p.writeField1(oprot); err != nil {
			fieldId = 1
			goto WriteFieldError
		}
		if err = p.writeField2(oprot); err != nil {
			fieldId = 2
			goto WriteFieldError
		}
	}
	if err = oprot.WriteFieldStop(); err != nil {
		goto WriteFieldStopError
	}
	if err = oprot.WriteStructEnd(); err != nil {
		goto WriteStructEndError
	}
	return nil
WriteStructBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
WriteFieldError:
	return thrift.PrependError(fmt.Sprintf("%T write field %d error: ", p, fieldId), err)
WriteFieldStopError:
	return thrift.PrependError(fmt.Sprintf("%T write field stop error: ", p), err)
WriteStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T write struct end error: ", p), err)
}

func (p *CourseListResponse) writeField1(oprot thrift.TProtocol) (err error) {
	if err = oprot.WriteFieldBegin("base", thrift.STRUCT, 1); err != nil {
		goto WriteFieldBeginError
	}
	if err := p.Base.Write(oprot); err != nil {
		return err
	}
	if err = oprot.WriteFieldEnd(); err != nil {
		goto WriteFieldEndError
	}
	return nil
WriteFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write field 1 begin error: ", p), err)
WriteFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T write field 1 end error: ", p), err)
}

func (p *CourseListResponse) writeField2(oprot thrift.TProtocol) (err error) {
	if err = oprot.WriteFieldBegin("data", thrift.LIST, 2); err != nil {
		goto WriteFieldBeginError
	}
	if err := oprot.WriteListBegin(thrift.STRUCT, len(p.Data)); err != nil {
		return err
	}
	for _, v := range p.Data {
		if err := v.Write(oprot); err != nil {
			return err
		}
	}
	if err := oprot.WriteListEnd(); err != nil {
		return err
	}
	if err = oprot.WriteFieldEnd(); err != nil {
		goto WriteFieldEndError
	}
	return nil
WriteFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write field 2 begin error: ", p), err)
WriteFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T write field 2 end error: ", p), err)
}

func (p *CourseListResponse) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("CourseListResponse(%+v)", *p)

}

func (p *CourseListResponse) DeepEqual(ano *CourseListResponse) bool {
	if p == ano {
		return true
	} else if p == nil || ano == nil {
		return false
	}
	if !p.Field1DeepEqual(ano.Base) {
		return false
	}
	if !p.Field2DeepEqual(ano.Data) {
		return false
	}
	return true
}

func (p *CourseListResponse) Field1DeepEqual(src *model.BaseResp) bool {

	if !p.Base.DeepEqual(src) {
		return false
	}
	return true
}
func (p *CourseListResponse) Field2DeepEqual(src []*model.Course) bool {

	if len(p.Data) != len(src) {
		return false
	}
	for i, v := range p.Data {
		_src := src[i]
		if !v.DeepEqual(_src) {
			return false
		}
	}
	return true
}

type CourseService interface {
	GetCourseList(ctx context.Context, req *CourseListRequest) (r *CourseListResponse, err error)
}

type CourseServiceGetCourseListArgs struct {
	Req *CourseListRequest `thrift:"req,1" frugal:"1,default,CourseListRequest" json:"req"`
}

func NewCourseServiceGetCourseListArgs() *CourseServiceGetCourseListArgs {
	return &CourseServiceGetCourseListArgs{}
}

func (p *CourseServiceGetCourseListArgs) InitDefault() {
}

var CourseServiceGetCourseListArgs_Req_DEFAULT *CourseListRequest

func (p *CourseServiceGetCourseListArgs) GetReq() (v *CourseListRequest) {
	if !p.IsSetReq() {
		return CourseServiceGetCourseListArgs_Req_DEFAULT
	}
	return p.Req
}
func (p *CourseServiceGetCourseListArgs) SetReq(val *CourseListRequest) {
	p.Req = val
}

var fieldIDToName_CourseServiceGetCourseListArgs = map[int16]string{
	1: "req",
}

func (p *CourseServiceGetCourseListArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *CourseServiceGetCourseListArgs) Read(iprot thrift.TProtocol) (err error) {

	var fieldTypeId thrift.TType
	var fieldId int16

	if _, err = iprot.ReadStructBegin(); err != nil {
		goto ReadStructBeginError
	}

	for {
		_, fieldTypeId, fieldId, err = iprot.ReadFieldBegin()
		if err != nil {
			goto ReadFieldBeginError
		}
		if fieldTypeId == thrift.STOP {
			break
		}

		switch fieldId {
		case 1:
			if fieldTypeId == thrift.STRUCT {
				if err = p.ReadField1(iprot); err != nil {
					goto ReadFieldError
				}
			} else if err = iprot.Skip(fieldTypeId); err != nil {
				goto SkipFieldError
			}
		default:
			if err = iprot.Skip(fieldTypeId); err != nil {
				goto SkipFieldError
			}
		}
		if err = iprot.ReadFieldEnd(); err != nil {
			goto ReadFieldEndError
		}
	}
	if err = iprot.ReadStructEnd(); err != nil {
		goto ReadStructEndError
	}

	return nil
ReadStructBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read struct begin error: ", p), err)
ReadFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read field %d begin error: ", p, fieldId), err)
ReadFieldError:
	return thrift.PrependError(fmt.Sprintf("%T read field %d '%s' error: ", p, fieldId, fieldIDToName_CourseServiceGetCourseListArgs[fieldId]), err)
SkipFieldError:
	return thrift.PrependError(fmt.Sprintf("%T field %d skip type %d error: ", p, fieldId, fieldTypeId), err)

ReadFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T read field end error", p), err)
ReadStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
}

func (p *CourseServiceGetCourseListArgs) ReadField1(iprot thrift.TProtocol) error {
	_field := NewCourseListRequest()
	if err := _field.Read(iprot); err != nil {
		return err
	}
	p.Req = _field
	return nil
}

func (p *CourseServiceGetCourseListArgs) Write(oprot thrift.TProtocol) (err error) {

	var fieldId int16
	if err = oprot.WriteStructBegin("GetCourseList_args"); err != nil {
		goto WriteStructBeginError
	}
	if p != nil {
		if err = p.writeField1(oprot); err != nil {
			fieldId = 1
			goto WriteFieldError
		}
	}
	if err = oprot.WriteFieldStop(); err != nil {
		goto WriteFieldStopError
	}
	if err = oprot.WriteStructEnd(); err != nil {
		goto WriteStructEndError
	}
	return nil
WriteStructBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
WriteFieldError:
	return thrift.PrependError(fmt.Sprintf("%T write field %d error: ", p, fieldId), err)
WriteFieldStopError:
	return thrift.PrependError(fmt.Sprintf("%T write field stop error: ", p), err)
WriteStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T write struct end error: ", p), err)
}

func (p *CourseServiceGetCourseListArgs) writeField1(oprot thrift.TProtocol) (err error) {
	if err = oprot.WriteFieldBegin("req", thrift.STRUCT, 1); err != nil {
		goto WriteFieldBeginError
	}
	if err := p.Req.Write(oprot); err != nil {
		return err
	}
	if err = oprot.WriteFieldEnd(); err != nil {
		goto WriteFieldEndError
	}
	return nil
WriteFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write field 1 begin error: ", p), err)
WriteFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T write field 1 end error: ", p), err)
}

func (p *CourseServiceGetCourseListArgs) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("CourseServiceGetCourseListArgs(%+v)", *p)

}

func (p *CourseServiceGetCourseListArgs) DeepEqual(ano *CourseServiceGetCourseListArgs) bool {
	if p == ano {
		return true
	} else if p == nil || ano == nil {
		return false
	}
	if !p.Field1DeepEqual(ano.Req) {
		return false
	}
	return true
}

func (p *CourseServiceGetCourseListArgs) Field1DeepEqual(src *CourseListRequest) bool {

	if !p.Req.DeepEqual(src) {
		return false
	}
	return true
}

type CourseServiceGetCourseListResult struct {
	Success *CourseListResponse `thrift:"success,0,optional" frugal:"0,optional,CourseListResponse" json:"success,omitempty"`
}

func NewCourseServiceGetCourseListResult() *CourseServiceGetCourseListResult {
	return &CourseServiceGetCourseListResult{}
}

func (p *CourseServiceGetCourseListResult) InitDefault() {
}

var CourseServiceGetCourseListResult_Success_DEFAULT *CourseListResponse

func (p *CourseServiceGetCourseListResult) GetSuccess() (v *CourseListResponse) {
	if !p.IsSetSuccess() {
		return CourseServiceGetCourseListResult_Success_DEFAULT
	}
	return p.Success
}
func (p *CourseServiceGetCourseListResult) SetSuccess(x interface{}) {
	p.Success = x.(*CourseListResponse)
}

var fieldIDToName_CourseServiceGetCourseListResult = map[int16]string{
	0: "success",
}

func (p *CourseServiceGetCourseListResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *CourseServiceGetCourseListResult) Read(iprot thrift.TProtocol) (err error) {

	var fieldTypeId thrift.TType
	var fieldId int16

	if _, err = iprot.ReadStructBegin(); err != nil {
		goto ReadStructBeginError
	}

	for {
		_, fieldTypeId, fieldId, err = iprot.ReadFieldBegin()
		if err != nil {
			goto ReadFieldBeginError
		}
		if fieldTypeId == thrift.STOP {
			break
		}

		switch fieldId {
		case 0:
			if fieldTypeId == thrift.STRUCT {
				if err = p.ReadField0(iprot); err != nil {
					goto ReadFieldError
				}
			} else if err = iprot.Skip(fieldTypeId); err != nil {
				goto SkipFieldError
			}
		default:
			if err = iprot.Skip(fieldTypeId); err != nil {
				goto SkipFieldError
			}
		}
		if err = iprot.ReadFieldEnd(); err != nil {
			goto ReadFieldEndError
		}
	}
	if err = iprot.ReadStructEnd(); err != nil {
		goto ReadStructEndError
	}

	return nil
ReadStructBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read struct begin error: ", p), err)
ReadFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read field %d begin error: ", p, fieldId), err)
ReadFieldError:
	return thrift.PrependError(fmt.Sprintf("%T read field %d '%s' error: ", p, fieldId, fieldIDToName_CourseServiceGetCourseListResult[fieldId]), err)
SkipFieldError:
	return thrift.PrependError(fmt.Sprintf("%T field %d skip type %d error: ", p, fieldId, fieldTypeId), err)

ReadFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T read field end error", p), err)
ReadStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
}

func (p *CourseServiceGetCourseListResult) ReadField0(iprot thrift.TProtocol) error {
	_field := NewCourseListResponse()
	if err := _field.Read(iprot); err != nil {
		return err
	}
	p.Success = _field
	return nil
}

func (p *CourseServiceGetCourseListResult) Write(oprot thrift.TProtocol) (err error) {

	var fieldId int16
	if err = oprot.WriteStructBegin("GetCourseList_result"); err != nil {
		goto WriteStructBeginError
	}
	if p != nil {
		if err = p.writeField0(oprot); err != nil {
			fieldId = 0
			goto WriteFieldError
		}
	}
	if err = oprot.WriteFieldStop(); err != nil {
		goto WriteFieldStopError
	}
	if err = oprot.WriteStructEnd(); err != nil {
		goto WriteStructEndError
	}
	return nil
WriteStructBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
WriteFieldError:
	return thrift.PrependError(fmt.Sprintf("%T write field %d error: ", p, fieldId), err)
WriteFieldStopError:
	return thrift.PrependError(fmt.Sprintf("%T write field stop error: ", p), err)
WriteStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T write struct end error: ", p), err)
}

func (p *CourseServiceGetCourseListResult) writeField0(oprot thrift.TProtocol) (err error) {
	if p.IsSetSuccess() {
		if err = oprot.WriteFieldBegin("success", thrift.STRUCT, 0); err != nil {
			goto WriteFieldBeginError
		}
		if err := p.Success.Write(oprot); err != nil {
			return err
		}
		if err = oprot.WriteFieldEnd(); err != nil {
			goto WriteFieldEndError
		}
	}
	return nil
WriteFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write field 0 begin error: ", p), err)
WriteFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T write field 0 end error: ", p), err)
}

func (p *CourseServiceGetCourseListResult) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("CourseServiceGetCourseListResult(%+v)", *p)

}

func (p *CourseServiceGetCourseListResult) DeepEqual(ano *CourseServiceGetCourseListResult) bool {
	if p == ano {
		return true
	} else if p == nil || ano == nil {
		return false
	}
	if !p.Field0DeepEqual(ano.Success) {
		return false
	}
	return true
}

func (p *CourseServiceGetCourseListResult) Field0DeepEqual(src *CourseListResponse) bool {

	if !p.Success.DeepEqual(src) {
		return false
	}
	return true
}
