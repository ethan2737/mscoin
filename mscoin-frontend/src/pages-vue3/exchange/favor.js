export const applyFavorState = ({
  coins,
  row,
  symbol,
  isFavor,
  currentCoinSymbol
}) => {
  const trackedCoin = coins._map[symbol]

  if (trackedCoin) {
    trackedCoin.isFavor = isFavor
  }

  if (row && row.symbol === symbol) {
    row.isFavor = isFavor
  }

  if (isFavor) {
    if (trackedCoin && !coins.favor.some(item => item.symbol === symbol)) {
      coins.favor.push(trackedCoin)
    }
  } else {
    const favorIndex = coins.favor.findIndex(item => item.symbol === symbol)
    if (favorIndex !== -1) {
      coins.favor.splice(favorIndex, 1)
    }
  }

  return {
    currentCoinIsFavor: currentCoinSymbol === symbol ? isFavor : undefined
  }
}

export const getFavorSuccessMessage = (isFavor) => {
  return isFavor ? '已添加自选' : '已取消自选'
}
