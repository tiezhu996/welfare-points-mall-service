# BUG_REPRO

## Bug 是什么
积分消费把“余额恰好够”判为不足，扣减写成加钱；秒杀库存校验反向导致超卖；过期扫描 worker 返回的是总流水数而不是本次新过期数，重复统计。

## 如何触发
`go test ./...`

## 错误信息
- model.TestCanSpend 失败（余额等于金额被拒）。
- service.TestSpendRejectsInsufficient / TestFlashPurchaseRejectsOversell 失败（余额反向、超卖）。
- store.TestFlashStockDeduct 失败（充足库存被判为不足）。
- worker.TestRunExpiresOnce 失败（第二次仍统计为 1）。
