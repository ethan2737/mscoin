export const toOrderNumber = (value) => {
  const numeric = Number(value)
  return Number.isFinite(numeric) ? numeric : 0
}

export const buildOrderPayload = ({
  symbol,
  direction,
  type,
  price,
  amount,
  useDiscount
}) => {
  const payload = {
    symbol,
    direction,
    type,
    amount: toOrderNumber(amount)
  }

  if (type === 'LIMIT_PRICE') {
    payload.price = toOrderNumber(price)
  }

  if (useDiscount !== undefined) {
    payload.useDiscount = useDiscount
  }

  return payload
}
