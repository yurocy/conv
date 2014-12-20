// Copyright 2014 by caixw, All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

// conv 提供了基础的类型转换功能。会尽最大可能地
// 将当前的值转换成指定类型的值。
//  conv.MustInt("123", 0)  // 返回123的数值
//  conv.MustString(123, "")// 返回字符串123
//  conv.Int("123", 0)      // 返回123数值和nil的error接口
//  conv.ToByte("12M")      // 返回12*1024*1024的计算结果值.
//  v := 5
//  conv.To("3", reflect.ValueOf(v)) // 将3转换成数值,并写入v中.
package conv

const Version = "0.1.0.141220"