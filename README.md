# Garlic



soc 或者 siem 平台用的规则引擎  使用antlr 解析自定义的脚本语言
部分规则借鉴gengine。


## ✨ 特性
增加 滑动窗口 聚合计算支持。
多规则 联动处理


## 🎁 例子
滑动窗口 聚合计算 规则

rule "1" "2"
begin
count(Rs.Small>11) >10 within 100 second
end
`
