package errors

import "github.com/pkg/errors"

// 信息 code 额外errs
func New(s string,code int ,extras ...error) error {
	return &StackError{
		code:      code,
		err:       errors.New(s),
		extraErrs: extras,
	}
}

// 错误堆栈设置+额外errs
func WithStack(err error,code int, extras ...error) error {
	return &StackError{
		code:      code,
		err:       errors.WithStack(err),
		extraErrs: extras,
	}
}

// 错误+信息+额外errs
func WithMessage(err error,msg string,code int , extras ...error) error {
	err = errors.WithStack(err)
	return &StackError{
		code:      code,
		err:       errors.Wrap(err, msg),
		extraErrs: extras,
	}
}

// 错误+格式+信息
func WithMessageF(err error, format string, code int,args ...interface{}) error {
	return &StackError{
		code: code,
		err:  errors.Wrapf(err, format, args...),
	}
}
