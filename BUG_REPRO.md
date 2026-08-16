# BUG_REPRO

## Bug 是什么
配置读取返回空；积分判断只要金额为正就放行；消费写成加钱且忽略余额校验；秒杀扣库存不写回；过期扫描统计多 1。

## 如何触发
`go test ./...`

## 错误信息
- config.TestLoad 失败（AppName 为空）。
- model.TestCanSpend 失败（余额不足也放行）。
- service.TestSpendRejectsInsufficient 失败（不足也扣）。
- store.TestFlashStockDeduct 失败（库存未扣减）。
- worker.TestRunExpiresOnce 失败（统计多 1）。
