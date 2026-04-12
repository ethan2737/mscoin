export const PLATE_DISPLAY_ROWS = 10

const toFiniteNumber = (value) => {
  const numericValue = Number(value)
  return Number.isFinite(numericValue) ? numericValue : null
}

const normalizePlateItems = (items = []) => (
  Array.isArray(items)
    ? items
      .map((item) => {
        const price = toFiniteNumber(item?.price)
        const amount = toFiniteNumber(item?.amount)
        if (price === null || amount === null || price <= 0 || amount < 0) {
          return null
        }
        return {
          price,
          amount
        }
      })
      .filter(Boolean)
    : []
)

const withCumulativeTotals = (items) => {
  let totalAmount = 0
  return items.map((item) => {
    totalAmount += item.amount
    return {
      ...item,
      totalAmount
    }
  })
}

export const createPlaceholderRow = (direction, position) => ({
  direction,
  position,
  price: null,
  amount: null,
  totalAmount: null,
  displayPrice: '--',
  displayAmount: '--',
  displayTotalAmount: '--',
  isPlaceholder: true
})

const createDisplayRow = (item, direction, position) => ({
  direction,
  position,
  price: item.price,
  amount: item.amount,
  totalAmount: item.totalAmount,
  displayPrice: item.price,
  displayAmount: item.amount,
  displayTotalAmount: item.totalAmount,
  isPlaceholder: false
})

export const buildSellPlateRows = (items = [], maxRows = PLATE_DISPLAY_ROWS) => {
  const rows = []
  const normalizedItems = withCumulativeTotals(normalizePlateItems(items))
  const visibleItems = normalizedItems.slice(0, maxRows).reverse()
  const placeholderCount = Math.max(maxRows - visibleItems.length, 0)

  for (let index = 0; index < placeholderCount; index++) {
    rows.push(createPlaceholderRow('SELL', maxRows - index))
  }

  for (let index = 0; index < visibleItems.length; index++) {
    rows.push(createDisplayRow(visibleItems[index], 'SELL', visibleItems.length - index))
  }

  return rows
}

export const buildBuyPlateRows = (items = [], maxRows = PLATE_DISPLAY_ROWS) => {
  const rows = []
  const normalizedItems = withCumulativeTotals(normalizePlateItems(items)).slice(0, maxRows)

  for (let index = 0; index < normalizedItems.length; index++) {
    rows.push(createDisplayRow(normalizedItems[index], 'BUY', index + 1))
  }

  for (let index = normalizedItems.length; index < maxRows; index++) {
    rows.push(createPlaceholderRow('BUY', index + 1))
  }

  return rows
}

export const getPlateRowPrice = (row) => {
  const price = toFiniteNumber(row?.price)
  if (price === null || price <= 0 || row?.isPlaceholder) {
    return null
  }
  return price
}
