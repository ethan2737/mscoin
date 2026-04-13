# Swap API 接口自动化测试

## 测试环境

- Base URL: `http://localhost:8086`
- 服务：swap-api (永续合约 HTTP API)

## 测试用例

### 1. 获取合约信息接口

**接口**: `GET /swap/contract-info/info`

**请求参数**:
```json
{
  "symbol": "BTC/USDT"
}
```

**预期响应**:
```json
{
  "code": 200,
  "data": {
    "id": 1,
    "symbol": "BTC/USDT",
    "name": "比特币永续合约",
    "status": 1
  }
}
```

### 2. 获取合约列表接口

**接口**: `GET /swap/contract-info/list`

**预期响应**:
```json
{
  "code": 200,
  "data": {
    "list": [
      {
        "id": 1,
        "symbol": "BTC/USDT",
        "name": "比特币永续合约"
      }
    ]
  }
}
```

### 3. 委托下单接口

**接口**: `POST /swap/order-entrust/add`

**请求参数**:
```json
{
  "memberId": 1001,
  "contractCoinId": 1,
  "entrustType": 1,
  "type": 1,
  "direction": 1,
  "leverage": 10,
  "price": 50000,
  "amount": 0.1
}
```

**预期响应**:
```json
{
  "code": 200,
  "data": {
    "orderId": 1
  }
}
```

### 4. 撤销订单接口

**接口**: `POST /swap/order-entrust/cancel/:id`

**请求参数**:
```json
{
  "memberId": 1001
}
```

**预期响应**:
```json
{
  "code": 200,
  "data": {
    "success": true
  }
}
```

### 5. 当前委托接口

**接口**: `POST /swap/order-entrust/current`

**请求参数**:
```json
{
  "memberId": 1001
}
```

**预期响应**:
```json
{
  "code": 200,
  "data": {
    "list": []
  }
}
```

### 6. 历史委托接口

**接口**: `POST /swap/order-entrust/history`

**请求参数**:
```json
{
  "memberId": 1001
}
```

**预期响应**:
```json
{
  "code": 200,
  "data": {
    "list": []
  }
}
```

### 7. 持仓信息接口

**接口**: `POST /swap/order-entrust/position`

**请求参数**:
```json
{
  "memberId": 1001
}
```

**预期响应**:
```json
{
  "code": 200,
  "data": {
    "list": []
  }
}
```

### 8. 闪电平仓接口

**接口**: `POST /swap/order-entrust/quick-close/:id`

**请求参数**:
```json
{
  "memberId": 1001,
  "price": 50000
}
```

**预期响应**:
```json
{
  "code": 200,
  "data": {
    "success": true
  }
}
```

### 9. 杠杆倍数接口

**接口**: `POST /swap/order-entrust/leverage`

**请求参数**:
```json
{
  "memberId": 1001,
  "contractCoinId": 1
}
```

**预期响应**:
```json
{
  "code": 200,
  "data": {
    "leverage": 10
  }
}
```

### 10. 合约钱包信息接口

**接口**: `POST /swap/wallet/info`

**请求参数**:
```json
{
  "memberId": 1001,
  "unit": "USDT"
}
```

**预期响应**:
```json
{
  "code": 200,
  "data": {
    "balance": 10000,
    "frozen": 0,
    "mainBalance": 10000
  }
}
```

### 11. 合约转账接口

**接口**: `POST /swap/wallet/transfer`

**请求参数**:
```json
{
  "memberId": 1001,
  "unit": "USDT",
  "amount": 1000,
  "type": 1
}
```

**预期响应**:
```json
{
  "code": 200,
  "data": {
    "success": true
  }
}
```

### 12. 交易流水接口

**接口**: `POST /swap/wallet/transaction`

**请求参数**:
```json
{
  "memberId": 1001,
  "page": 1,
  "pageSize": 10
}
```

**预期响应**:
```json
{
  "code": 200,
  "data": {
    "list": []
  }
}
```

## 测试步骤

1. 启动 Docker Compose 服务
2. 等待服务启动完成
3. 按顺序执行上述 API 测试
4. 验证所有接口返回正确响应

## 测试通过标准

- 所有接口返回 HTTP 200 状态码
- 所有接口返回正确的 JSON 格式
- 业务逻辑验证通过（如下单后查询当前委托应包含该订单）
