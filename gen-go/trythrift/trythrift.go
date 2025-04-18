// Code generated by Thrift Compiler (0.21.0). DO NOT EDIT.

package trythrift

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"log/slog"
	"time"
	thrift "github.com/apache/thrift/lib/go/thrift"
	"strings"
	"regexp"
)

// (needed to ensure safety because of naive import list construction.)
var _ = bytes.Equal
var _ = context.Background
var _ = errors.New
var _ = fmt.Printf
var _ = slog.Log
var _ = time.Now
var _ = thrift.ZERO
// (needed by validator.)
var _ = strings.Contains
var _ = regexp.MatchString

type Calculator interface {
  // Parameters:
  //  - Num1
  //  - Num2
  Add(ctx context.Context, num1 int32, num2 int32) (_r int32, _err error)
  // Parameters:
  //  - Num1
  //  - Num2
  Multiply(ctx context.Context, num1 int32, num2 int32) (_r int32, _err error)
}

type CalculatorClient struct {
  c thrift.TClient
  meta thrift.ResponseMeta
}

func NewCalculatorClientFactory(t thrift.TTransport, f thrift.TProtocolFactory) *CalculatorClient {
  return &CalculatorClient{
    c: thrift.NewTStandardClient(f.GetProtocol(t), f.GetProtocol(t)),
  }
}

func NewCalculatorClientProtocol(t thrift.TTransport, iprot thrift.TProtocol, oprot thrift.TProtocol) *CalculatorClient {
  return &CalculatorClient{
    c: thrift.NewTStandardClient(iprot, oprot),
  }
}

func NewCalculatorClient(c thrift.TClient) *CalculatorClient {
  return &CalculatorClient{
    c: c,
  }
}

func (p *CalculatorClient) Client_() thrift.TClient {
  return p.c
}

func (p *CalculatorClient) LastResponseMeta_() thrift.ResponseMeta {
  return p.meta
}

func (p *CalculatorClient) SetLastResponseMeta_(meta thrift.ResponseMeta) {
  p.meta = meta
}

// Parameters:
//  - Num1
//  - Num2
func (p *CalculatorClient) Add(ctx context.Context, num1 int32, num2 int32) (_r int32, _err error) {
  var _args0 CalculatorAddArgs
  _args0.Num1 = num1
  _args0.Num2 = num2
  var _result2 CalculatorAddResult
  var _meta1 thrift.ResponseMeta
  _meta1, _err = p.Client_().Call(ctx, "add", &_args0, &_result2)
  p.SetLastResponseMeta_(_meta1)
  if _err != nil {
    return
  }
  return _result2.GetSuccess(), nil
}

// Parameters:
//  - Num1
//  - Num2
func (p *CalculatorClient) Multiply(ctx context.Context, num1 int32, num2 int32) (_r int32, _err error) {
  var _args3 CalculatorMultiplyArgs
  _args3.Num1 = num1
  _args3.Num2 = num2
  var _result5 CalculatorMultiplyResult
  var _meta4 thrift.ResponseMeta
  _meta4, _err = p.Client_().Call(ctx, "multiply", &_args3, &_result5)
  p.SetLastResponseMeta_(_meta4)
  if _err != nil {
    return
  }
  return _result5.GetSuccess(), nil
}

type CalculatorProcessor struct {
  processorMap map[string]thrift.TProcessorFunction
  handler Calculator
}

func (p *CalculatorProcessor) AddToProcessorMap(key string, processor thrift.TProcessorFunction) {
  p.processorMap[key] = processor
}

func (p *CalculatorProcessor) GetProcessorFunction(key string) (processor thrift.TProcessorFunction, ok bool) {
  processor, ok = p.processorMap[key]
  return processor, ok
}

func (p *CalculatorProcessor) ProcessorMap() map[string]thrift.TProcessorFunction {
  return p.processorMap
}

func NewCalculatorProcessor(handler Calculator) *CalculatorProcessor {

  self6 := &CalculatorProcessor{handler:handler, processorMap:make(map[string]thrift.TProcessorFunction)}
  self6.processorMap["add"] = &calculatorProcessorAdd{handler:handler}
  self6.processorMap["multiply"] = &calculatorProcessorMultiply{handler:handler}
return self6
}

func (p *CalculatorProcessor) Process(ctx context.Context, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
  name, _, seqId, err2 := iprot.ReadMessageBegin(ctx)
  if err2 != nil { return false, thrift.WrapTException(err2) }
  if processor, ok := p.GetProcessorFunction(name); ok {
    return processor.Process(ctx, seqId, iprot, oprot)
  }
  iprot.Skip(ctx, thrift.STRUCT)
  iprot.ReadMessageEnd(ctx)
  x7 := thrift.NewTApplicationException(thrift.UNKNOWN_METHOD, "Unknown function " + name)
  oprot.WriteMessageBegin(ctx, name, thrift.EXCEPTION, seqId)
  x7.Write(ctx, oprot)
  oprot.WriteMessageEnd(ctx)
  oprot.Flush(ctx)
  return false, x7

}

type calculatorProcessorAdd struct {
  handler Calculator
}

func (p *calculatorProcessorAdd) Process(ctx context.Context, seqId int32, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
  var _write_err8 error
  args := CalculatorAddArgs{}
  if err2 := args.Read(ctx, iprot); err2 != nil {
    iprot.ReadMessageEnd(ctx)
    x := thrift.NewTApplicationException(thrift.PROTOCOL_ERROR, err2.Error())
    oprot.WriteMessageBegin(ctx, "add", thrift.EXCEPTION, seqId)
    x.Write(ctx, oprot)
    oprot.WriteMessageEnd(ctx)
    oprot.Flush(ctx)
    return false, thrift.WrapTException(err2)
  }
  iprot.ReadMessageEnd(ctx)

  tickerCancel := func() {}
  // Start a goroutine to do server side connectivity check.
  if thrift.ServerConnectivityCheckInterval > 0 {
    var cancel context.CancelCauseFunc
    ctx, cancel = context.WithCancelCause(ctx)
    defer cancel(nil)
    var tickerCtx context.Context
    tickerCtx, tickerCancel = context.WithCancel(context.Background())
    defer tickerCancel()
    go func(ctx context.Context, cancel context.CancelCauseFunc) {
      ticker := time.NewTicker(thrift.ServerConnectivityCheckInterval)
      defer ticker.Stop()
      for {
        select {
        case <-ctx.Done():
          return
        case <-ticker.C:
          if !iprot.Transport().IsOpen() {
            cancel(thrift.ErrAbandonRequest)
            return
          }
        }
      }
    }(tickerCtx, cancel)
  }

  result := CalculatorAddResult{}
  if retval, err2 := p.handler.Add(ctx, args.Num1, args.Num2); err2 != nil {
    tickerCancel()
    err = thrift.WrapTException(err2)
    if errors.Is(err2, thrift.ErrAbandonRequest) {
      return false, thrift.WrapTException(err2)
    }
    if errors.Is(err2, context.Canceled) {
      if err := context.Cause(ctx); errors.Is(err, thrift.ErrAbandonRequest) {
        return false, thrift.WrapTException(err)
      }
    }
    _exc9 := thrift.NewTApplicationException(thrift.INTERNAL_ERROR, "Internal error processing add: " + err2.Error())
    if err2 := oprot.WriteMessageBegin(ctx, "add", thrift.EXCEPTION, seqId); err2 != nil {
      _write_err8 = thrift.WrapTException(err2)
    }
    if err2 := _exc9.Write(ctx, oprot); _write_err8 == nil && err2 != nil {
      _write_err8 = thrift.WrapTException(err2)
    }
    if err2 := oprot.WriteMessageEnd(ctx); _write_err8 == nil && err2 != nil {
      _write_err8 = thrift.WrapTException(err2)
    }
    if err2 := oprot.Flush(ctx); _write_err8 == nil && err2 != nil {
      _write_err8 = thrift.WrapTException(err2)
    }
    if _write_err8 != nil {
      return false, thrift.WrapTException(_write_err8)
    }
    return true, err
  } else {
    result.Success = &retval
  }
  tickerCancel()
  if err2 := oprot.WriteMessageBegin(ctx, "add", thrift.REPLY, seqId); err2 != nil {
    _write_err8 = thrift.WrapTException(err2)
  }
  if err2 := result.Write(ctx, oprot); _write_err8 == nil && err2 != nil {
    _write_err8 = thrift.WrapTException(err2)
  }
  if err2 := oprot.WriteMessageEnd(ctx); _write_err8 == nil && err2 != nil {
    _write_err8 = thrift.WrapTException(err2)
  }
  if err2 := oprot.Flush(ctx); _write_err8 == nil && err2 != nil {
    _write_err8 = thrift.WrapTException(err2)
  }
  if _write_err8 != nil {
    return false, thrift.WrapTException(_write_err8)
  }
  return true, err
}

type calculatorProcessorMultiply struct {
  handler Calculator
}

func (p *calculatorProcessorMultiply) Process(ctx context.Context, seqId int32, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
  var _write_err10 error
  args := CalculatorMultiplyArgs{}
  if err2 := args.Read(ctx, iprot); err2 != nil {
    iprot.ReadMessageEnd(ctx)
    x := thrift.NewTApplicationException(thrift.PROTOCOL_ERROR, err2.Error())
    oprot.WriteMessageBegin(ctx, "multiply", thrift.EXCEPTION, seqId)
    x.Write(ctx, oprot)
    oprot.WriteMessageEnd(ctx)
    oprot.Flush(ctx)
    return false, thrift.WrapTException(err2)
  }
  iprot.ReadMessageEnd(ctx)

  tickerCancel := func() {}
  // Start a goroutine to do server side connectivity check.
  if thrift.ServerConnectivityCheckInterval > 0 {
    var cancel context.CancelCauseFunc
    ctx, cancel = context.WithCancelCause(ctx)
    defer cancel(nil)
    var tickerCtx context.Context
    tickerCtx, tickerCancel = context.WithCancel(context.Background())
    defer tickerCancel()
    go func(ctx context.Context, cancel context.CancelCauseFunc) {
      ticker := time.NewTicker(thrift.ServerConnectivityCheckInterval)
      defer ticker.Stop()
      for {
        select {
        case <-ctx.Done():
          return
        case <-ticker.C:
          if !iprot.Transport().IsOpen() {
            cancel(thrift.ErrAbandonRequest)
            return
          }
        }
      }
    }(tickerCtx, cancel)
  }

  result := CalculatorMultiplyResult{}
  if retval, err2 := p.handler.Multiply(ctx, args.Num1, args.Num2); err2 != nil {
    tickerCancel()
    err = thrift.WrapTException(err2)
    if errors.Is(err2, thrift.ErrAbandonRequest) {
      return false, thrift.WrapTException(err2)
    }
    if errors.Is(err2, context.Canceled) {
      if err := context.Cause(ctx); errors.Is(err, thrift.ErrAbandonRequest) {
        return false, thrift.WrapTException(err)
      }
    }
    _exc11 := thrift.NewTApplicationException(thrift.INTERNAL_ERROR, "Internal error processing multiply: " + err2.Error())
    if err2 := oprot.WriteMessageBegin(ctx, "multiply", thrift.EXCEPTION, seqId); err2 != nil {
      _write_err10 = thrift.WrapTException(err2)
    }
    if err2 := _exc11.Write(ctx, oprot); _write_err10 == nil && err2 != nil {
      _write_err10 = thrift.WrapTException(err2)
    }
    if err2 := oprot.WriteMessageEnd(ctx); _write_err10 == nil && err2 != nil {
      _write_err10 = thrift.WrapTException(err2)
    }
    if err2 := oprot.Flush(ctx); _write_err10 == nil && err2 != nil {
      _write_err10 = thrift.WrapTException(err2)
    }
    if _write_err10 != nil {
      return false, thrift.WrapTException(_write_err10)
    }
    return true, err
  } else {
    result.Success = &retval
  }
  tickerCancel()
  if err2 := oprot.WriteMessageBegin(ctx, "multiply", thrift.REPLY, seqId); err2 != nil {
    _write_err10 = thrift.WrapTException(err2)
  }
  if err2 := result.Write(ctx, oprot); _write_err10 == nil && err2 != nil {
    _write_err10 = thrift.WrapTException(err2)
  }
  if err2 := oprot.WriteMessageEnd(ctx); _write_err10 == nil && err2 != nil {
    _write_err10 = thrift.WrapTException(err2)
  }
  if err2 := oprot.Flush(ctx); _write_err10 == nil && err2 != nil {
    _write_err10 = thrift.WrapTException(err2)
  }
  if _write_err10 != nil {
    return false, thrift.WrapTException(_write_err10)
  }
  return true, err
}


// HELPER FUNCTIONS AND STRUCTURES

// Attributes:
//  - Num1
//  - Num2
type CalculatorAddArgs struct {
  Num1 int32 `thrift:"num1,1" db:"num1" json:"num1"`
  Num2 int32 `thrift:"num2,2" db:"num2" json:"num2"`
}

func NewCalculatorAddArgs() *CalculatorAddArgs {
  return &CalculatorAddArgs{}
}


func (p *CalculatorAddArgs) GetNum1() int32 {
  return p.Num1
}

func (p *CalculatorAddArgs) GetNum2() int32 {
  return p.Num2
}
func (p *CalculatorAddArgs) Read(ctx context.Context, iprot thrift.TProtocol) error {
  if _, err := iprot.ReadStructBegin(ctx); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
  }


  for {
    _, fieldTypeId, fieldId, err := iprot.ReadFieldBegin(ctx)
    if err != nil {
      return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
    }
    if fieldTypeId == thrift.STOP { break; }
    switch fieldId {
    case 1:
      if fieldTypeId == thrift.I32 {
        if err := p.ReadField1(ctx, iprot); err != nil {
          return err
        }
      } else {
        if err := iprot.Skip(ctx, fieldTypeId); err != nil {
          return err
        }
      }
    case 2:
      if fieldTypeId == thrift.I32 {
        if err := p.ReadField2(ctx, iprot); err != nil {
          return err
        }
      } else {
        if err := iprot.Skip(ctx, fieldTypeId); err != nil {
          return err
        }
      }
    default:
      if err := iprot.Skip(ctx, fieldTypeId); err != nil {
        return err
      }
    }
    if err := iprot.ReadFieldEnd(ctx); err != nil {
      return err
    }
  }
  if err := iprot.ReadStructEnd(ctx); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
  }
  return nil
}

func (p *CalculatorAddArgs)  ReadField1(ctx context.Context, iprot thrift.TProtocol) error {
  if v, err := iprot.ReadI32(ctx); err != nil {
  return thrift.PrependError("error reading field 1: ", err)
} else {
  p.Num1 = v
}
  return nil
}

func (p *CalculatorAddArgs)  ReadField2(ctx context.Context, iprot thrift.TProtocol) error {
  if v, err := iprot.ReadI32(ctx); err != nil {
  return thrift.PrependError("error reading field 2: ", err)
} else {
  p.Num2 = v
}
  return nil
}

func (p *CalculatorAddArgs) Write(ctx context.Context, oprot thrift.TProtocol) error {
  if err := oprot.WriteStructBegin(ctx, "add_args"); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err) }
  if p != nil {
    if err := p.writeField1(ctx, oprot); err != nil { return err }
    if err := p.writeField2(ctx, oprot); err != nil { return err }
  }
  if err := oprot.WriteFieldStop(ctx); err != nil {
    return thrift.PrependError("write field stop error: ", err) }
  if err := oprot.WriteStructEnd(ctx); err != nil {
    return thrift.PrependError("write struct stop error: ", err) }
  return nil
}

func (p *CalculatorAddArgs) writeField1(ctx context.Context, oprot thrift.TProtocol) (err error) {
  if err := oprot.WriteFieldBegin(ctx, "num1", thrift.I32, 1); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:num1: ", p), err) }
  if err := oprot.WriteI32(ctx, int32(p.Num1)); err != nil {
  return thrift.PrependError(fmt.Sprintf("%T.num1 (1) field write error: ", p), err) }
  if err := oprot.WriteFieldEnd(ctx); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field end error 1:num1: ", p), err) }
  return err
}

func (p *CalculatorAddArgs) writeField2(ctx context.Context, oprot thrift.TProtocol) (err error) {
  if err := oprot.WriteFieldBegin(ctx, "num2", thrift.I32, 2); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field begin error 2:num2: ", p), err) }
  if err := oprot.WriteI32(ctx, int32(p.Num2)); err != nil {
  return thrift.PrependError(fmt.Sprintf("%T.num2 (2) field write error: ", p), err) }
  if err := oprot.WriteFieldEnd(ctx); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field end error 2:num2: ", p), err) }
  return err
}

func (p *CalculatorAddArgs) String() string {
  if p == nil {
    return "<nil>"
  }
  return fmt.Sprintf("CalculatorAddArgs(%+v)", *p)
}

func (p *CalculatorAddArgs) LogValue() slog.Value {
  if p == nil {
    return slog.AnyValue(nil)
  }
  v := thrift.SlogTStructWrapper{
    Type: "*trythrift.CalculatorAddArgs",
    Value: p,
  }
  return slog.AnyValue(v)
}

var _ slog.LogValuer = (*CalculatorAddArgs)(nil)

// Attributes:
//  - Success
type CalculatorAddResult struct {
  Success *int32 `thrift:"success,0" db:"success" json:"success,omitempty"`
}

func NewCalculatorAddResult() *CalculatorAddResult {
  return &CalculatorAddResult{}
}

var CalculatorAddResult_Success_DEFAULT int32
func (p *CalculatorAddResult) GetSuccess() int32 {
  if !p.IsSetSuccess() {
    return CalculatorAddResult_Success_DEFAULT
  }
  return *p.Success
}
func (p *CalculatorAddResult) IsSetSuccess() bool {
  return p.Success != nil
}

func (p *CalculatorAddResult) Read(ctx context.Context, iprot thrift.TProtocol) error {
  if _, err := iprot.ReadStructBegin(ctx); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
  }


  for {
    _, fieldTypeId, fieldId, err := iprot.ReadFieldBegin(ctx)
    if err != nil {
      return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
    }
    if fieldTypeId == thrift.STOP { break; }
    switch fieldId {
    case 0:
      if fieldTypeId == thrift.I32 {
        if err := p.ReadField0(ctx, iprot); err != nil {
          return err
        }
      } else {
        if err := iprot.Skip(ctx, fieldTypeId); err != nil {
          return err
        }
      }
    default:
      if err := iprot.Skip(ctx, fieldTypeId); err != nil {
        return err
      }
    }
    if err := iprot.ReadFieldEnd(ctx); err != nil {
      return err
    }
  }
  if err := iprot.ReadStructEnd(ctx); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
  }
  return nil
}

func (p *CalculatorAddResult)  ReadField0(ctx context.Context, iprot thrift.TProtocol) error {
  if v, err := iprot.ReadI32(ctx); err != nil {
  return thrift.PrependError("error reading field 0: ", err)
} else {
  p.Success = &v
}
  return nil
}

func (p *CalculatorAddResult) Write(ctx context.Context, oprot thrift.TProtocol) error {
  if err := oprot.WriteStructBegin(ctx, "add_result"); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err) }
  if p != nil {
    if err := p.writeField0(ctx, oprot); err != nil { return err }
  }
  if err := oprot.WriteFieldStop(ctx); err != nil {
    return thrift.PrependError("write field stop error: ", err) }
  if err := oprot.WriteStructEnd(ctx); err != nil {
    return thrift.PrependError("write struct stop error: ", err) }
  return nil
}

func (p *CalculatorAddResult) writeField0(ctx context.Context, oprot thrift.TProtocol) (err error) {
  if p.IsSetSuccess() {
    if err := oprot.WriteFieldBegin(ctx, "success", thrift.I32, 0); err != nil {
      return thrift.PrependError(fmt.Sprintf("%T write field begin error 0:success: ", p), err) }
    if err := oprot.WriteI32(ctx, int32(*p.Success)); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T.success (0) field write error: ", p), err) }
    if err := oprot.WriteFieldEnd(ctx); err != nil {
      return thrift.PrependError(fmt.Sprintf("%T write field end error 0:success: ", p), err) }
  }
  return err
}

func (p *CalculatorAddResult) String() string {
  if p == nil {
    return "<nil>"
  }
  return fmt.Sprintf("CalculatorAddResult(%+v)", *p)
}

func (p *CalculatorAddResult) LogValue() slog.Value {
  if p == nil {
    return slog.AnyValue(nil)
  }
  v := thrift.SlogTStructWrapper{
    Type: "*trythrift.CalculatorAddResult",
    Value: p,
  }
  return slog.AnyValue(v)
}

var _ slog.LogValuer = (*CalculatorAddResult)(nil)

// Attributes:
//  - Num1
//  - Num2
type CalculatorMultiplyArgs struct {
  Num1 int32 `thrift:"num1,1" db:"num1" json:"num1"`
  Num2 int32 `thrift:"num2,2" db:"num2" json:"num2"`
}

func NewCalculatorMultiplyArgs() *CalculatorMultiplyArgs {
  return &CalculatorMultiplyArgs{}
}


func (p *CalculatorMultiplyArgs) GetNum1() int32 {
  return p.Num1
}

func (p *CalculatorMultiplyArgs) GetNum2() int32 {
  return p.Num2
}
func (p *CalculatorMultiplyArgs) Read(ctx context.Context, iprot thrift.TProtocol) error {
  if _, err := iprot.ReadStructBegin(ctx); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
  }


  for {
    _, fieldTypeId, fieldId, err := iprot.ReadFieldBegin(ctx)
    if err != nil {
      return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
    }
    if fieldTypeId == thrift.STOP { break; }
    switch fieldId {
    case 1:
      if fieldTypeId == thrift.I32 {
        if err := p.ReadField1(ctx, iprot); err != nil {
          return err
        }
      } else {
        if err := iprot.Skip(ctx, fieldTypeId); err != nil {
          return err
        }
      }
    case 2:
      if fieldTypeId == thrift.I32 {
        if err := p.ReadField2(ctx, iprot); err != nil {
          return err
        }
      } else {
        if err := iprot.Skip(ctx, fieldTypeId); err != nil {
          return err
        }
      }
    default:
      if err := iprot.Skip(ctx, fieldTypeId); err != nil {
        return err
      }
    }
    if err := iprot.ReadFieldEnd(ctx); err != nil {
      return err
    }
  }
  if err := iprot.ReadStructEnd(ctx); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
  }
  return nil
}

func (p *CalculatorMultiplyArgs)  ReadField1(ctx context.Context, iprot thrift.TProtocol) error {
  if v, err := iprot.ReadI32(ctx); err != nil {
  return thrift.PrependError("error reading field 1: ", err)
} else {
  p.Num1 = v
}
  return nil
}

func (p *CalculatorMultiplyArgs)  ReadField2(ctx context.Context, iprot thrift.TProtocol) error {
  if v, err := iprot.ReadI32(ctx); err != nil {
  return thrift.PrependError("error reading field 2: ", err)
} else {
  p.Num2 = v
}
  return nil
}

func (p *CalculatorMultiplyArgs) Write(ctx context.Context, oprot thrift.TProtocol) error {
  if err := oprot.WriteStructBegin(ctx, "multiply_args"); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err) }
  if p != nil {
    if err := p.writeField1(ctx, oprot); err != nil { return err }
    if err := p.writeField2(ctx, oprot); err != nil { return err }
  }
  if err := oprot.WriteFieldStop(ctx); err != nil {
    return thrift.PrependError("write field stop error: ", err) }
  if err := oprot.WriteStructEnd(ctx); err != nil {
    return thrift.PrependError("write struct stop error: ", err) }
  return nil
}

func (p *CalculatorMultiplyArgs) writeField1(ctx context.Context, oprot thrift.TProtocol) (err error) {
  if err := oprot.WriteFieldBegin(ctx, "num1", thrift.I32, 1); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:num1: ", p), err) }
  if err := oprot.WriteI32(ctx, int32(p.Num1)); err != nil {
  return thrift.PrependError(fmt.Sprintf("%T.num1 (1) field write error: ", p), err) }
  if err := oprot.WriteFieldEnd(ctx); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field end error 1:num1: ", p), err) }
  return err
}

func (p *CalculatorMultiplyArgs) writeField2(ctx context.Context, oprot thrift.TProtocol) (err error) {
  if err := oprot.WriteFieldBegin(ctx, "num2", thrift.I32, 2); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field begin error 2:num2: ", p), err) }
  if err := oprot.WriteI32(ctx, int32(p.Num2)); err != nil {
  return thrift.PrependError(fmt.Sprintf("%T.num2 (2) field write error: ", p), err) }
  if err := oprot.WriteFieldEnd(ctx); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field end error 2:num2: ", p), err) }
  return err
}

func (p *CalculatorMultiplyArgs) String() string {
  if p == nil {
    return "<nil>"
  }
  return fmt.Sprintf("CalculatorMultiplyArgs(%+v)", *p)
}

func (p *CalculatorMultiplyArgs) LogValue() slog.Value {
  if p == nil {
    return slog.AnyValue(nil)
  }
  v := thrift.SlogTStructWrapper{
    Type: "*trythrift.CalculatorMultiplyArgs",
    Value: p,
  }
  return slog.AnyValue(v)
}

var _ slog.LogValuer = (*CalculatorMultiplyArgs)(nil)

// Attributes:
//  - Success
type CalculatorMultiplyResult struct {
  Success *int32 `thrift:"success,0" db:"success" json:"success,omitempty"`
}

func NewCalculatorMultiplyResult() *CalculatorMultiplyResult {
  return &CalculatorMultiplyResult{}
}

var CalculatorMultiplyResult_Success_DEFAULT int32
func (p *CalculatorMultiplyResult) GetSuccess() int32 {
  if !p.IsSetSuccess() {
    return CalculatorMultiplyResult_Success_DEFAULT
  }
  return *p.Success
}
func (p *CalculatorMultiplyResult) IsSetSuccess() bool {
  return p.Success != nil
}

func (p *CalculatorMultiplyResult) Read(ctx context.Context, iprot thrift.TProtocol) error {
  if _, err := iprot.ReadStructBegin(ctx); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
  }


  for {
    _, fieldTypeId, fieldId, err := iprot.ReadFieldBegin(ctx)
    if err != nil {
      return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
    }
    if fieldTypeId == thrift.STOP { break; }
    switch fieldId {
    case 0:
      if fieldTypeId == thrift.I32 {
        if err := p.ReadField0(ctx, iprot); err != nil {
          return err
        }
      } else {
        if err := iprot.Skip(ctx, fieldTypeId); err != nil {
          return err
        }
      }
    default:
      if err := iprot.Skip(ctx, fieldTypeId); err != nil {
        return err
      }
    }
    if err := iprot.ReadFieldEnd(ctx); err != nil {
      return err
    }
  }
  if err := iprot.ReadStructEnd(ctx); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
  }
  return nil
}

func (p *CalculatorMultiplyResult)  ReadField0(ctx context.Context, iprot thrift.TProtocol) error {
  if v, err := iprot.ReadI32(ctx); err != nil {
  return thrift.PrependError("error reading field 0: ", err)
} else {
  p.Success = &v
}
  return nil
}

func (p *CalculatorMultiplyResult) Write(ctx context.Context, oprot thrift.TProtocol) error {
  if err := oprot.WriteStructBegin(ctx, "multiply_result"); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err) }
  if p != nil {
    if err := p.writeField0(ctx, oprot); err != nil { return err }
  }
  if err := oprot.WriteFieldStop(ctx); err != nil {
    return thrift.PrependError("write field stop error: ", err) }
  if err := oprot.WriteStructEnd(ctx); err != nil {
    return thrift.PrependError("write struct stop error: ", err) }
  return nil
}

func (p *CalculatorMultiplyResult) writeField0(ctx context.Context, oprot thrift.TProtocol) (err error) {
  if p.IsSetSuccess() {
    if err := oprot.WriteFieldBegin(ctx, "success", thrift.I32, 0); err != nil {
      return thrift.PrependError(fmt.Sprintf("%T write field begin error 0:success: ", p), err) }
    if err := oprot.WriteI32(ctx, int32(*p.Success)); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T.success (0) field write error: ", p), err) }
    if err := oprot.WriteFieldEnd(ctx); err != nil {
      return thrift.PrependError(fmt.Sprintf("%T write field end error 0:success: ", p), err) }
  }
  return err
}

func (p *CalculatorMultiplyResult) String() string {
  if p == nil {
    return "<nil>"
  }
  return fmt.Sprintf("CalculatorMultiplyResult(%+v)", *p)
}

func (p *CalculatorMultiplyResult) LogValue() slog.Value {
  if p == nil {
    return slog.AnyValue(nil)
  }
  v := thrift.SlogTStructWrapper{
    Type: "*trythrift.CalculatorMultiplyResult",
    Value: p,
  }
  return slog.AnyValue(v)
}

var _ slog.LogValuer = (*CalculatorMultiplyResult)(nil)


