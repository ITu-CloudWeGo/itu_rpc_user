// Code generated by Fastpb v0.0.2. DO NOT EDIT.

package user_service

import (
	fmt "fmt"
	fastpb "github.com/cloudwego/fastpb"
)

var (
	_ = fmt.Errorf
	_ = fastpb.Skip
)

func (x *Response) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
	switch number {
	case 1:
		offset, err = x.fastReadField1(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 2:
		offset, err = x.fastReadField2(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	default:
		offset, err = fastpb.Skip(buf, _type, number)
		if err != nil {
			goto SkipFieldError
		}
	}
	return offset, nil
SkipFieldError:
	return offset, fmt.Errorf("%T cannot parse invalid wire-format data, error: %s", x, err)
ReadFieldError:
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_Response[number], err)
}

func (x *Response) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.Status, offset, err = fastpb.ReadInt64(buf, _type)
	return offset, err
}

func (x *Response) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	x.Msg, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *SignRequest) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
	switch number {
	case 1:
		offset, err = x.fastReadField1(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 2:
		offset, err = x.fastReadField2(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 3:
		offset, err = x.fastReadField3(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 4:
		offset, err = x.fastReadField4(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 5:
		offset, err = x.fastReadField5(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	default:
		offset, err = fastpb.Skip(buf, _type, number)
		if err != nil {
			goto SkipFieldError
		}
	}
	return offset, nil
SkipFieldError:
	return offset, fmt.Errorf("%T cannot parse invalid wire-format data, error: %s", x, err)
ReadFieldError:
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_SignRequest[number], err)
}

func (x *SignRequest) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.Email, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *SignRequest) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	x.Username, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *SignRequest) fastReadField3(buf []byte, _type int8) (offset int, err error) {
	x.Password, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *SignRequest) fastReadField4(buf []byte, _type int8) (offset int, err error) {
	x.Nickname, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *SignRequest) fastReadField5(buf []byte, _type int8) (offset int, err error) {
	x.Icon, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *SignResponse) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
	switch number {
	case 1:
		offset, err = x.fastReadField1(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 2:
		offset, err = x.fastReadField2(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 3:
		offset, err = x.fastReadField3(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 4:
		offset, err = x.fastReadField4(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 5:
		offset, err = x.fastReadField5(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	default:
		offset, err = fastpb.Skip(buf, _type, number)
		if err != nil {
			goto SkipFieldError
		}
	}
	return offset, nil
SkipFieldError:
	return offset, fmt.Errorf("%T cannot parse invalid wire-format data, error: %s", x, err)
ReadFieldError:
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_SignResponse[number], err)
}

func (x *SignResponse) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	var v Response
	offset, err = fastpb.ReadMessage(buf, _type, &v)
	if err != nil {
		return offset, err
	}
	x.Response = &v
	return offset, nil
}

func (x *SignResponse) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	x.Email, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *SignResponse) fastReadField3(buf []byte, _type int8) (offset int, err error) {
	x.Username, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *SignResponse) fastReadField4(buf []byte, _type int8) (offset int, err error) {
	x.Nickname, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *SignResponse) fastReadField5(buf []byte, _type int8) (offset int, err error) {
	x.Icon, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *LoginRequest) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
	switch number {
	case 1:
		offset, err = x.fastReadField1(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 2:
		offset, err = x.fastReadField2(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 3:
		offset, err = x.fastReadField3(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	default:
		offset, err = fastpb.Skip(buf, _type, number)
		if err != nil {
			goto SkipFieldError
		}
	}
	return offset, nil
SkipFieldError:
	return offset, fmt.Errorf("%T cannot parse invalid wire-format data, error: %s", x, err)
ReadFieldError:
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_LoginRequest[number], err)
}

func (x *LoginRequest) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.Username, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *LoginRequest) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	x.Email, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *LoginRequest) fastReadField3(buf []byte, _type int8) (offset int, err error) {
	x.Password, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *LoginResponse) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
	switch number {
	case 1:
		offset, err = x.fastReadField1(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 2:
		offset, err = x.fastReadField2(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 3:
		offset, err = x.fastReadField3(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 4:
		offset, err = x.fastReadField4(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 5:
		offset, err = x.fastReadField5(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	default:
		offset, err = fastpb.Skip(buf, _type, number)
		if err != nil {
			goto SkipFieldError
		}
	}
	return offset, nil
SkipFieldError:
	return offset, fmt.Errorf("%T cannot parse invalid wire-format data, error: %s", x, err)
ReadFieldError:
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_LoginResponse[number], err)
}

func (x *LoginResponse) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	var v Response
	offset, err = fastpb.ReadMessage(buf, _type, &v)
	if err != nil {
		return offset, err
	}
	x.Response = &v
	return offset, nil
}

func (x *LoginResponse) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	x.Email, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *LoginResponse) fastReadField3(buf []byte, _type int8) (offset int, err error) {
	x.Username, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *LoginResponse) fastReadField4(buf []byte, _type int8) (offset int, err error) {
	x.Nickname, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *LoginResponse) fastReadField5(buf []byte, _type int8) (offset int, err error) {
	x.Icon, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *GetUserRequest) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
	switch number {
	case 1:
		offset, err = x.fastReadField1(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 2:
		offset, err = x.fastReadField2(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	default:
		offset, err = fastpb.Skip(buf, _type, number)
		if err != nil {
			goto SkipFieldError
		}
	}
	return offset, nil
SkipFieldError:
	return offset, fmt.Errorf("%T cannot parse invalid wire-format data, error: %s", x, err)
ReadFieldError:
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_GetUserRequest[number], err)
}

func (x *GetUserRequest) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.Uid, offset, err = fastpb.ReadUint64(buf, _type)
	return offset, err
}

func (x *GetUserRequest) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	x.Username, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *GetUserResponse) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
	switch number {
	case 1:
		offset, err = x.fastReadField1(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 2:
		offset, err = x.fastReadField2(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 3:
		offset, err = x.fastReadField3(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 4:
		offset, err = x.fastReadField4(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	default:
		offset, err = fastpb.Skip(buf, _type, number)
		if err != nil {
			goto SkipFieldError
		}
	}
	return offset, nil
SkipFieldError:
	return offset, fmt.Errorf("%T cannot parse invalid wire-format data, error: %s", x, err)
ReadFieldError:
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_GetUserResponse[number], err)
}

func (x *GetUserResponse) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	var v Response
	offset, err = fastpb.ReadMessage(buf, _type, &v)
	if err != nil {
		return offset, err
	}
	x.Response = &v
	return offset, nil
}

func (x *GetUserResponse) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	x.Username, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *GetUserResponse) fastReadField3(buf []byte, _type int8) (offset int, err error) {
	x.Nickname, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *GetUserResponse) fastReadField4(buf []byte, _type int8) (offset int, err error) {
	x.Icon, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *UpdateUserRequest) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
	switch number {
	case 1:
		offset, err = x.fastReadField1(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 2:
		offset, err = x.fastReadField2(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 3:
		offset, err = x.fastReadField3(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 4:
		offset, err = x.fastReadField4(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 5:
		offset, err = x.fastReadField5(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 6:
		offset, err = x.fastReadField6(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 7:
		offset, err = x.fastReadField7(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	default:
		offset, err = fastpb.Skip(buf, _type, number)
		if err != nil {
			goto SkipFieldError
		}
	}
	return offset, nil
SkipFieldError:
	return offset, fmt.Errorf("%T cannot parse invalid wire-format data, error: %s", x, err)
ReadFieldError:
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_UpdateUserRequest[number], err)
}

func (x *UpdateUserRequest) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.Uid, offset, err = fastpb.ReadUint64(buf, _type)
	return offset, err
}

func (x *UpdateUserRequest) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	x.NewEmail, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *UpdateUserRequest) fastReadField3(buf []byte, _type int8) (offset int, err error) {
	x.NewNickname, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *UpdateUserRequest) fastReadField4(buf []byte, _type int8) (offset int, err error) {
	x.NewIcon, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *UpdateUserRequest) fastReadField5(buf []byte, _type int8) (offset int, err error) {
	x.Password, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *UpdateUserRequest) fastReadField6(buf []byte, _type int8) (offset int, err error) {
	x.NewPassword, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *UpdateUserRequest) fastReadField7(buf []byte, _type int8) (offset int, err error) {
	x.ConfirmPasswd, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *UpdateUserResponse) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
	switch number {
	case 1:
		offset, err = x.fastReadField1(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 2:
		offset, err = x.fastReadField2(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 3:
		offset, err = x.fastReadField3(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 4:
		offset, err = x.fastReadField4(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 5:
		offset, err = x.fastReadField5(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	default:
		offset, err = fastpb.Skip(buf, _type, number)
		if err != nil {
			goto SkipFieldError
		}
	}
	return offset, nil
SkipFieldError:
	return offset, fmt.Errorf("%T cannot parse invalid wire-format data, error: %s", x, err)
ReadFieldError:
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_UpdateUserResponse[number], err)
}

func (x *UpdateUserResponse) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	var v Response
	offset, err = fastpb.ReadMessage(buf, _type, &v)
	if err != nil {
		return offset, err
	}
	x.Response = &v
	return offset, nil
}

func (x *UpdateUserResponse) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	x.Username, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *UpdateUserResponse) fastReadField3(buf []byte, _type int8) (offset int, err error) {
	x.Email, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *UpdateUserResponse) fastReadField4(buf []byte, _type int8) (offset int, err error) {
	x.Nickname, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *UpdateUserResponse) fastReadField5(buf []byte, _type int8) (offset int, err error) {
	x.Icon, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *Response) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	return offset
}

func (x *Response) fastWriteField1(buf []byte) (offset int) {
	if x.Status == 0 {
		return offset
	}
	offset += fastpb.WriteInt64(buf[offset:], 1, x.GetStatus())
	return offset
}

func (x *Response) fastWriteField2(buf []byte) (offset int) {
	if x.Msg == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 2, x.GetMsg())
	return offset
}

func (x *SignRequest) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	offset += x.fastWriteField3(buf[offset:])
	offset += x.fastWriteField4(buf[offset:])
	offset += x.fastWriteField5(buf[offset:])
	return offset
}

func (x *SignRequest) fastWriteField1(buf []byte) (offset int) {
	if x.Email == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 1, x.GetEmail())
	return offset
}

func (x *SignRequest) fastWriteField2(buf []byte) (offset int) {
	if x.Username == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 2, x.GetUsername())
	return offset
}

func (x *SignRequest) fastWriteField3(buf []byte) (offset int) {
	if x.Password == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 3, x.GetPassword())
	return offset
}

func (x *SignRequest) fastWriteField4(buf []byte) (offset int) {
	if x.Nickname == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 4, x.GetNickname())
	return offset
}

func (x *SignRequest) fastWriteField5(buf []byte) (offset int) {
	if x.Icon == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 5, x.GetIcon())
	return offset
}

func (x *SignResponse) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	offset += x.fastWriteField3(buf[offset:])
	offset += x.fastWriteField4(buf[offset:])
	offset += x.fastWriteField5(buf[offset:])
	return offset
}

func (x *SignResponse) fastWriteField1(buf []byte) (offset int) {
	if x.Response == nil {
		return offset
	}
	offset += fastpb.WriteMessage(buf[offset:], 1, x.GetResponse())
	return offset
}

func (x *SignResponse) fastWriteField2(buf []byte) (offset int) {
	if x.Email == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 2, x.GetEmail())
	return offset
}

func (x *SignResponse) fastWriteField3(buf []byte) (offset int) {
	if x.Username == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 3, x.GetUsername())
	return offset
}

func (x *SignResponse) fastWriteField4(buf []byte) (offset int) {
	if x.Nickname == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 4, x.GetNickname())
	return offset
}

func (x *SignResponse) fastWriteField5(buf []byte) (offset int) {
	if x.Icon == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 5, x.GetIcon())
	return offset
}

func (x *LoginRequest) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	offset += x.fastWriteField3(buf[offset:])
	return offset
}

func (x *LoginRequest) fastWriteField1(buf []byte) (offset int) {
	if x.Username == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 1, x.GetUsername())
	return offset
}

func (x *LoginRequest) fastWriteField2(buf []byte) (offset int) {
	if x.Email == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 2, x.GetEmail())
	return offset
}

func (x *LoginRequest) fastWriteField3(buf []byte) (offset int) {
	if x.Password == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 3, x.GetPassword())
	return offset
}

func (x *LoginResponse) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	offset += x.fastWriteField3(buf[offset:])
	offset += x.fastWriteField4(buf[offset:])
	offset += x.fastWriteField5(buf[offset:])
	return offset
}

func (x *LoginResponse) fastWriteField1(buf []byte) (offset int) {
	if x.Response == nil {
		return offset
	}
	offset += fastpb.WriteMessage(buf[offset:], 1, x.GetResponse())
	return offset
}

func (x *LoginResponse) fastWriteField2(buf []byte) (offset int) {
	if x.Email == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 2, x.GetEmail())
	return offset
}

func (x *LoginResponse) fastWriteField3(buf []byte) (offset int) {
	if x.Username == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 3, x.GetUsername())
	return offset
}

func (x *LoginResponse) fastWriteField4(buf []byte) (offset int) {
	if x.Nickname == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 4, x.GetNickname())
	return offset
}

func (x *LoginResponse) fastWriteField5(buf []byte) (offset int) {
	if x.Icon == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 5, x.GetIcon())
	return offset
}

func (x *GetUserRequest) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	return offset
}

func (x *GetUserRequest) fastWriteField1(buf []byte) (offset int) {
	if x.Uid == 0 {
		return offset
	}
	offset += fastpb.WriteUint64(buf[offset:], 1, x.GetUid())
	return offset
}

func (x *GetUserRequest) fastWriteField2(buf []byte) (offset int) {
	if x.Username == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 2, x.GetUsername())
	return offset
}

func (x *GetUserResponse) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	offset += x.fastWriteField3(buf[offset:])
	offset += x.fastWriteField4(buf[offset:])
	return offset
}

func (x *GetUserResponse) fastWriteField1(buf []byte) (offset int) {
	if x.Response == nil {
		return offset
	}
	offset += fastpb.WriteMessage(buf[offset:], 1, x.GetResponse())
	return offset
}

func (x *GetUserResponse) fastWriteField2(buf []byte) (offset int) {
	if x.Username == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 2, x.GetUsername())
	return offset
}

func (x *GetUserResponse) fastWriteField3(buf []byte) (offset int) {
	if x.Nickname == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 3, x.GetNickname())
	return offset
}

func (x *GetUserResponse) fastWriteField4(buf []byte) (offset int) {
	if x.Icon == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 4, x.GetIcon())
	return offset
}

func (x *UpdateUserRequest) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	offset += x.fastWriteField3(buf[offset:])
	offset += x.fastWriteField4(buf[offset:])
	offset += x.fastWriteField5(buf[offset:])
	offset += x.fastWriteField6(buf[offset:])
	offset += x.fastWriteField7(buf[offset:])
	return offset
}

func (x *UpdateUserRequest) fastWriteField1(buf []byte) (offset int) {
	if x.Uid == 0 {
		return offset
	}
	offset += fastpb.WriteUint64(buf[offset:], 1, x.GetUid())
	return offset
}

func (x *UpdateUserRequest) fastWriteField2(buf []byte) (offset int) {
	if x.NewEmail == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 2, x.GetNewEmail())
	return offset
}

func (x *UpdateUserRequest) fastWriteField3(buf []byte) (offset int) {
	if x.NewNickname == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 3, x.GetNewNickname())
	return offset
}

func (x *UpdateUserRequest) fastWriteField4(buf []byte) (offset int) {
	if x.NewIcon == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 4, x.GetNewIcon())
	return offset
}

func (x *UpdateUserRequest) fastWriteField5(buf []byte) (offset int) {
	if x.Password == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 5, x.GetPassword())
	return offset
}

func (x *UpdateUserRequest) fastWriteField6(buf []byte) (offset int) {
	if x.NewPassword == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 6, x.GetNewPassword())
	return offset
}

func (x *UpdateUserRequest) fastWriteField7(buf []byte) (offset int) {
	if x.ConfirmPasswd == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 7, x.GetConfirmPasswd())
	return offset
}

func (x *UpdateUserResponse) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	offset += x.fastWriteField3(buf[offset:])
	offset += x.fastWriteField4(buf[offset:])
	offset += x.fastWriteField5(buf[offset:])
	return offset
}

func (x *UpdateUserResponse) fastWriteField1(buf []byte) (offset int) {
	if x.Response == nil {
		return offset
	}
	offset += fastpb.WriteMessage(buf[offset:], 1, x.GetResponse())
	return offset
}

func (x *UpdateUserResponse) fastWriteField2(buf []byte) (offset int) {
	if x.Username == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 2, x.GetUsername())
	return offset
}

func (x *UpdateUserResponse) fastWriteField3(buf []byte) (offset int) {
	if x.Email == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 3, x.GetEmail())
	return offset
}

func (x *UpdateUserResponse) fastWriteField4(buf []byte) (offset int) {
	if x.Nickname == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 4, x.GetNickname())
	return offset
}

func (x *UpdateUserResponse) fastWriteField5(buf []byte) (offset int) {
	if x.Icon == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 5, x.GetIcon())
	return offset
}

func (x *Response) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	return n
}

func (x *Response) sizeField1() (n int) {
	if x.Status == 0 {
		return n
	}
	n += fastpb.SizeInt64(1, x.GetStatus())
	return n
}

func (x *Response) sizeField2() (n int) {
	if x.Msg == "" {
		return n
	}
	n += fastpb.SizeString(2, x.GetMsg())
	return n
}

func (x *SignRequest) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	n += x.sizeField3()
	n += x.sizeField4()
	n += x.sizeField5()
	return n
}

func (x *SignRequest) sizeField1() (n int) {
	if x.Email == "" {
		return n
	}
	n += fastpb.SizeString(1, x.GetEmail())
	return n
}

func (x *SignRequest) sizeField2() (n int) {
	if x.Username == "" {
		return n
	}
	n += fastpb.SizeString(2, x.GetUsername())
	return n
}

func (x *SignRequest) sizeField3() (n int) {
	if x.Password == "" {
		return n
	}
	n += fastpb.SizeString(3, x.GetPassword())
	return n
}

func (x *SignRequest) sizeField4() (n int) {
	if x.Nickname == "" {
		return n
	}
	n += fastpb.SizeString(4, x.GetNickname())
	return n
}

func (x *SignRequest) sizeField5() (n int) {
	if x.Icon == "" {
		return n
	}
	n += fastpb.SizeString(5, x.GetIcon())
	return n
}

func (x *SignResponse) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	n += x.sizeField3()
	n += x.sizeField4()
	n += x.sizeField5()
	return n
}

func (x *SignResponse) sizeField1() (n int) {
	if x.Response == nil {
		return n
	}
	n += fastpb.SizeMessage(1, x.GetResponse())
	return n
}

func (x *SignResponse) sizeField2() (n int) {
	if x.Email == "" {
		return n
	}
	n += fastpb.SizeString(2, x.GetEmail())
	return n
}

func (x *SignResponse) sizeField3() (n int) {
	if x.Username == "" {
		return n
	}
	n += fastpb.SizeString(3, x.GetUsername())
	return n
}

func (x *SignResponse) sizeField4() (n int) {
	if x.Nickname == "" {
		return n
	}
	n += fastpb.SizeString(4, x.GetNickname())
	return n
}

func (x *SignResponse) sizeField5() (n int) {
	if x.Icon == "" {
		return n
	}
	n += fastpb.SizeString(5, x.GetIcon())
	return n
}

func (x *LoginRequest) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	n += x.sizeField3()
	return n
}

func (x *LoginRequest) sizeField1() (n int) {
	if x.Username == "" {
		return n
	}
	n += fastpb.SizeString(1, x.GetUsername())
	return n
}

func (x *LoginRequest) sizeField2() (n int) {
	if x.Email == "" {
		return n
	}
	n += fastpb.SizeString(2, x.GetEmail())
	return n
}

func (x *LoginRequest) sizeField3() (n int) {
	if x.Password == "" {
		return n
	}
	n += fastpb.SizeString(3, x.GetPassword())
	return n
}

func (x *LoginResponse) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	n += x.sizeField3()
	n += x.sizeField4()
	n += x.sizeField5()
	return n
}

func (x *LoginResponse) sizeField1() (n int) {
	if x.Response == nil {
		return n
	}
	n += fastpb.SizeMessage(1, x.GetResponse())
	return n
}

func (x *LoginResponse) sizeField2() (n int) {
	if x.Email == "" {
		return n
	}
	n += fastpb.SizeString(2, x.GetEmail())
	return n
}

func (x *LoginResponse) sizeField3() (n int) {
	if x.Username == "" {
		return n
	}
	n += fastpb.SizeString(3, x.GetUsername())
	return n
}

func (x *LoginResponse) sizeField4() (n int) {
	if x.Nickname == "" {
		return n
	}
	n += fastpb.SizeString(4, x.GetNickname())
	return n
}

func (x *LoginResponse) sizeField5() (n int) {
	if x.Icon == "" {
		return n
	}
	n += fastpb.SizeString(5, x.GetIcon())
	return n
}

func (x *GetUserRequest) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	return n
}

func (x *GetUserRequest) sizeField1() (n int) {
	if x.Uid == 0 {
		return n
	}
	n += fastpb.SizeUint64(1, x.GetUid())
	return n
}

func (x *GetUserRequest) sizeField2() (n int) {
	if x.Username == "" {
		return n
	}
	n += fastpb.SizeString(2, x.GetUsername())
	return n
}

func (x *GetUserResponse) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	n += x.sizeField3()
	n += x.sizeField4()
	return n
}

func (x *GetUserResponse) sizeField1() (n int) {
	if x.Response == nil {
		return n
	}
	n += fastpb.SizeMessage(1, x.GetResponse())
	return n
}

func (x *GetUserResponse) sizeField2() (n int) {
	if x.Username == "" {
		return n
	}
	n += fastpb.SizeString(2, x.GetUsername())
	return n
}

func (x *GetUserResponse) sizeField3() (n int) {
	if x.Nickname == "" {
		return n
	}
	n += fastpb.SizeString(3, x.GetNickname())
	return n
}

func (x *GetUserResponse) sizeField4() (n int) {
	if x.Icon == "" {
		return n
	}
	n += fastpb.SizeString(4, x.GetIcon())
	return n
}

func (x *UpdateUserRequest) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	n += x.sizeField3()
	n += x.sizeField4()
	n += x.sizeField5()
	n += x.sizeField6()
	n += x.sizeField7()
	return n
}

func (x *UpdateUserRequest) sizeField1() (n int) {
	if x.Uid == 0 {
		return n
	}
	n += fastpb.SizeUint64(1, x.GetUid())
	return n
}

func (x *UpdateUserRequest) sizeField2() (n int) {
	if x.NewEmail == "" {
		return n
	}
	n += fastpb.SizeString(2, x.GetNewEmail())
	return n
}

func (x *UpdateUserRequest) sizeField3() (n int) {
	if x.NewNickname == "" {
		return n
	}
	n += fastpb.SizeString(3, x.GetNewNickname())
	return n
}

func (x *UpdateUserRequest) sizeField4() (n int) {
	if x.NewIcon == "" {
		return n
	}
	n += fastpb.SizeString(4, x.GetNewIcon())
	return n
}

func (x *UpdateUserRequest) sizeField5() (n int) {
	if x.Password == "" {
		return n
	}
	n += fastpb.SizeString(5, x.GetPassword())
	return n
}

func (x *UpdateUserRequest) sizeField6() (n int) {
	if x.NewPassword == "" {
		return n
	}
	n += fastpb.SizeString(6, x.GetNewPassword())
	return n
}

func (x *UpdateUserRequest) sizeField7() (n int) {
	if x.ConfirmPasswd == "" {
		return n
	}
	n += fastpb.SizeString(7, x.GetConfirmPasswd())
	return n
}

func (x *UpdateUserResponse) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	n += x.sizeField3()
	n += x.sizeField4()
	n += x.sizeField5()
	return n
}

func (x *UpdateUserResponse) sizeField1() (n int) {
	if x.Response == nil {
		return n
	}
	n += fastpb.SizeMessage(1, x.GetResponse())
	return n
}

func (x *UpdateUserResponse) sizeField2() (n int) {
	if x.Username == "" {
		return n
	}
	n += fastpb.SizeString(2, x.GetUsername())
	return n
}

func (x *UpdateUserResponse) sizeField3() (n int) {
	if x.Email == "" {
		return n
	}
	n += fastpb.SizeString(3, x.GetEmail())
	return n
}

func (x *UpdateUserResponse) sizeField4() (n int) {
	if x.Nickname == "" {
		return n
	}
	n += fastpb.SizeString(4, x.GetNickname())
	return n
}

func (x *UpdateUserResponse) sizeField5() (n int) {
	if x.Icon == "" {
		return n
	}
	n += fastpb.SizeString(5, x.GetIcon())
	return n
}

var fieldIDToName_Response = map[int32]string{
	1: "Status",
	2: "Msg",
}

var fieldIDToName_SignRequest = map[int32]string{
	1: "Email",
	2: "Username",
	3: "Password",
	4: "Nickname",
	5: "Icon",
}

var fieldIDToName_SignResponse = map[int32]string{
	1: "Response",
	2: "Email",
	3: "Username",
	4: "Nickname",
	5: "Icon",
}

var fieldIDToName_LoginRequest = map[int32]string{
	1: "Username",
	2: "Email",
	3: "Password",
}

var fieldIDToName_LoginResponse = map[int32]string{
	1: "Response",
	2: "Email",
	3: "Username",
	4: "Nickname",
	5: "Icon",
}

var fieldIDToName_GetUserRequest = map[int32]string{
	1: "Uid",
	2: "Username",
}

var fieldIDToName_GetUserResponse = map[int32]string{
	1: "Response",
	2: "Username",
	3: "Nickname",
	4: "Icon",
}

var fieldIDToName_UpdateUserRequest = map[int32]string{
	1: "Uid",
	2: "NewEmail",
	3: "NewNickname",
	4: "NewIcon",
	5: "Password",
	6: "NewPassword",
	7: "ConfirmPasswd",
}

var fieldIDToName_UpdateUserResponse = map[int32]string{
	1: "Response",
	2: "Username",
	3: "Email",
	4: "Nickname",
	5: "Icon",
}
