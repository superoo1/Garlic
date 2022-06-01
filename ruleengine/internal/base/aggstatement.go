package base

import (
	"errors"
	"github.com/bilibili/gengine/context"
	"github.com/bilibili/gengine/util"
	"reflect"
	"time"
)

type AggStatement struct {
	SourceCode SourceCode
	Sss        string
	Expression *Expression
	Countopert string
	// todo 支持泛型
	ExtreValue int64
	ExpireTime int
	TimeForm   string
}

// expression
func (agg *AggStatement) Evaluate(dc *context.DataContext, Vars map[string]reflect.Value) (value reflect.Value, err error, b bool) {
	it, err := agg.Expression.Evaluate(dc, Vars)
	flag := it.Bool()
	if flag {
		// 原始rule代码+ 数据源作为键值对
		ruleExist := util.RedisDB.Exists("siem:" + agg.Sss).Val()
		if ruleExist >= 0 {
			util.RedisDB.Incr("siem:" + agg.Sss)
			aggValue, _ := util.RedisDB.Get("siem:" + agg.Sss).Int64()
			switch agg.Countopert {
			case ">":
				b = aggValue > agg.ExtreValue
				value = reflect.ValueOf(b)
				return
			case "<":
				b = aggValue > agg.ExtreValue
				value = reflect.ValueOf(b)
				return
			default:
			}
		} else {
			util.RedisDB.Incr("siem:" + agg.Sss)
			// 暂时 写死成分钟
			util.RedisDB.Expire("siem:"+agg.Sss, time.Minute*time.Duration(agg.ExpireTime))
		}
	}
	return
}

//  listener 模式太繁琐了
func (ef *AggStatement) AcceptExpression(expr *Expression) error {
	if ef.Expression == nil {
		ef.Expression = expr
		return nil
	}
	return errors.New("AggStatement's Expression set twice! ")
}

func (ef *AggStatement) AcceptCountoper(expr string) error {
	if ef.Countopert == "" {
		ef.Countopert = expr
		return nil
	}
	return errors.New("AggStatement's Expression set twice! ")
}
