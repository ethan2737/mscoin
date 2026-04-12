const normalizeSymbol = (value) => String(value || '').trim().toUpperCase()

const readWalletUnit = (wallet) => (
  normalizeSymbol(
    wallet?.coin?.unit ||
    wallet?.coin?.name ||
    wallet?.coinName ||
    wallet?.unit
  )
)

const readWalletBalance = (wallet) => {
  const balance = Number(wallet?.balance)
  return Number.isFinite(balance) ? balance : 0
}

const findWalletBySymbol = (wallets, symbol) => {
  const normalizedSymbol = normalizeSymbol(symbol)
  if (!normalizedSymbol || !Array.isArray(wallets)) {
    return null
  }

  return wallets.find((wallet) => readWalletUnit(wallet) === normalizedSymbol) || null
}

export const pickWalletBalances = ({
  wallets,
  baseSymbol,
  coinSymbol
} = {}) => {
  const baseWallet = findWalletBySymbol(wallets, baseSymbol)
  const coinWallet = findWalletBySymbol(wallets, coinSymbol)

  return {
    base: readWalletBalance(baseWallet),
    coin: readWalletBalance(coinWallet)
  }
}
