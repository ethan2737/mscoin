# MongoDB 集合设计

## swap_klines (K 线数据)

```javascript
{
    _id: ObjectId,
    symbol: String,        // 合约符号
    period: String,        // K 线周期：1m, 5m, 15m, 30m, 1h, 4h, 1d
    open_time: Date,       // 开盘时间
    close_time: Date,      // 收盘时间
    open: Number,          // 开盘价
    high: Number,          // 最高价
    low: Number,           // 最低价
    close: Number,         // 收盘价
    volume: Number,        // 成交量（张）
    amount: Number,        // 成交额（USDT）
    created_at: Date
}
```

**索引：**
```javascript
db.swap_klines.createIndex({symbol: 1, period: 1, open_time: -1})
```

## swap_trades (成交记录)

```javascript
{
    _id: ObjectId,
    symbol: String,        // 合约符号
    price: Number,         // 成交价
    amount: Number,        // 成交量（张）
    direction: String,     // 方向：buy, sell
    created_at: Date
}
```

**索引：**
```javascript
db.swap_trades.createIndex({symbol: 1, created_at: -1})
```

## swap_symbols (行情快照)

```javascript
{
    _id: ObjectId,
    symbol: String,        // 合约符号
    price: Number,         // 最新价
    change_24h: Number,    // 24h 涨跌幅
    high_24h: Number,      // 24h 最高价
    low_24h: Number,       // 24h 最低价
    volume_24h: Number,    // 24h 成交量
    amount_24h: Number,    // 24h 成交额
    updated_at: Date
}
```

**索引：**
```javascript
db.swap_symbols.createIndex({symbol: 1})
```
