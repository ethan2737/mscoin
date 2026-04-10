import { mapThumbToCoinViewModel, mergeRealtimeThumb } from '@/pages-vue3/index/market'

describe('index market mapping', () => {
  test('maps rise percent and 24h amount change from different fields', () => {
    const coin = mapThumbToCoinViewModel({
      symbol: 'BTC/USDT',
      close: 71558.4,
      usdRate: 71558.4,
      chg: 0.0123,
      change: 874.16,
      high: 73130,
      low: 71558.4,
      volume: 278310676.25523883
    }, ['BTC/USDT'])

    expect(coin.rise).toBe(1.23)
    expect(coin.change24h).toBe(874.16)
    expect(coin.isFavor).toBe(true)
  })

  test('falls back to legacy change24h/rise fields for local placeholder data', () => {
    const coin = mapThumbToCoinViewModel({
      symbol: 'SOL/USDT',
      price: 82.92,
      priceCNY: 603.12,
      rise: -0.85,
      change24h: -0.71,
      volume: 52746423.17244762
    }, [])

    expect(coin.rise).toBe(-0.85)
    expect(coin.change24h).toBe(-0.71)
    expect(coin.isFavor).toBe(false)
  })

  test('preserves existing 24h metrics when realtime payload carries empty summary fields', () => {
    const merged = mergeRealtimeThumb({
      symbol: 'BTC/USDT',
      close: 71629.9,
      change: 350,
      chg: 0.00491022,
      volume: 410167440.15642124,
      lastDayClose: 71279.9,
      high: 72350,
      low: 70531.6
    }, {
      symbol: 'BTC/USDT',
      close: 71680.1,
      change: 0,
      chg: 0,
      volume: 0,
      lastDayClose: 0
    })

    expect(merged.close).toBe(71680.1)
    expect(merged.change).toBe(350)
    expect(merged.chg).toBe(0.00491022)
    expect(merged.volume).toBe(410167440.15642124)
    expect(merged.lastDayClose).toBe(71279.9)
  })
})
