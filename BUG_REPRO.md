# BUG_REPRO

## Bug 是什么
存储初始化未建 accounts/items map，写入时 panic；积分余额判断反向；余额查询多 1；过期扫描统计多 1。

## 如何触发
`go test ./...`

## 错误信息
- model.TestCanSpend 失败（余额判断反向）。
- store.TestFlashStockDeduct 触发 assignment to entry in nil map panic。
- service.TestSpendRejectsInsufficient 触发 nil map panic。
- worker.TestRunExpiresOnce 失败（first run=2 期望 1）。
