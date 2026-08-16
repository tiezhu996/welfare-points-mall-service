# welfaremall

企业福利积分商城核心服务（纯 Go、内存存储），覆盖员工积分账户、积分兑换与秒杀库存。

## 目录结构
```
cmd/welfaremall/      程序入口
internal/config/      环境配置
internal/model/       积分账户与秒杀模型
internal/store/       内存账户与秒杀库存
internal/service/     积分与兑换业务
internal/worker/      积分过期扫描 worker
```

## 运行与测试
```bash
go build ./...
go test ./...
```
