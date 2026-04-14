<template>
  <div class="container swap" :class="skin">
    <div class="main">
      <div class="trade-workspace">
        <!-- 左侧列：币种列表 + 最新交易 -->
        <div class="left-column">
          <!-- 币种列表 -->
          <div class="market-rail">
            <div class="coin-menu panel-shell">
              <div class="panel-header panel-header-compact market-rail-header">
                <div>
                  <span class="panel-kicker">Perpetual</span>
                  <strong>{{ $t('service.ExchangeNum') }}</strong>
                </div>
                <span class="panel-meta">{{ dataIndex.length }}</span>
              </div>
              <div class="panel-toolbar market-search">
                <el-input
                  v-model="searchKey"
                  :placeholder="$t('common.searchplaceholder')"
                  @input="seachInputChange"
                  clearable
                >
                  <template #suffix>
                    <el-icon><Search /></el-icon>
                  </template>
                </el-input>
              </div>
              <div class="sc_filter" style="display: none">
                <span @click="changeBaseCion('usdt')" :class="{ active: basecion === 'usdt' }">USDT</span>
              </div>
              <el-table
                v-show="basecion === 'usdt'"
                :data="dataIndex"
                height="463"
                highlight-row
                @current-change="gohref"
              >
                <el-table-column label="合约/成交量" min-width="40" class-name="coin-menu-symbol swap-coin-menu-symbol">
                  <template #default="{ row }">
                    <div style="padding: 5px 0">
                      <span style="font-size: 14px">{{ row.name }}</span>
                      <br />
                      <span style="color: #61688A; font-size: 14px">{{ row.volume?.toFixed(4) }}</span>
                    </div>
                  </template>
                </el-table-column>
                <el-table-column label="最新价/涨跌幅" min-width="40" class-name="swap-coin-menu-lastprice">
                  <template #default="{ row }">
                    <div style="padding: 5px 10px; text-align: right; font-size: 14px">
                      <span>{{ row.close }}</span>
                      <br />
                      <span :class="parseFloat(row.rose) < 0 ? 'sell' : 'buy'">{{ row.rose }}</span>
                    </div>
                  </template>
                </el-table-column>
              </el-table>
            </div>
          </div>

          <!-- 最新交易 -->
          <div class="latest-trade-panel">
            <div class="trade-wrap panel-shell">
              <div class="panel-header panel-header-compact">
                <div>
                  <span class="panel-kicker">Tape</span>
                  <strong>{{ $t('swap.latestdeal') }}</strong>
                </div>
                <span class="panel-meta">{{ trade.rows.length }}</span>
              </div>
              <el-table :data="trade.rows" height="430" :no-data-text="$t('common.nodata')">
                <el-table-column label="价格">
                  <template #default="{ row }">
                    <span :class="row.direction === 'BUY' ? 'buy' : 'sell'">{{ row.price }}</span>
                  </template>
                </el-table-column>
                <el-table-column label="数量">
                  <template #default="{ row }">
                    {{ row.amount?.toFixed(4) }}
                  </template>
                </el-table-column>
                <el-table-column label="时间">
                  <template #default="{ row }">
                    {{ timeFormat(row.time) }}
                  </template>
                </el-table-column>
              </el-table>
            </div>
          </div>
        </div>

        <!-- 中间：K 线图、账户信息和交易表单 -->
        <div class="center">
          <!-- 币种信息 -->
          <div class="symbol panel-shell market-overview">
            <div class="item symbol-primary">
              <span class="panel-kicker">Contract</span>
              <span class="coin">{{ currentCoin.name }}</span>
              <el-popover trigger="hover" :width="300" placement="bottom-start">
                <template #reference>
                  <el-icon style="color: #546886; margin-left: 5px"><InfoFilled /></el-icon>
                </template>
                <div class="api">
                  <div class="coin-info">{{ coinInfo.information }}</div>
                  <p style="text-align: right; margin-top: 10px">
                    <a :href="coinInfo.infolink" target="_blank">{{ $t('swap.moredetail') }}</a>
                  </p>
                </div>
              </el-popover>
            </div>
            <div class="item">
              <span class="text">{{ $t('service.NewPrice') }}</span>
              <span class="num" :class="{ buy: currentCoin.change > 0, sell: currentCoin.change < 0 }">
                {{ currentCoin.close?.toFixed(4) }}
              </span>
              <span class="price-cny">≈ ￥{{ (currentCoin.usdRate * CNYRate)?.toFixed(2) }}</span>
            </div>
            <div class="item">
              <span class="text">{{ $t('service.Change') }}</span>
              <span class="num" :class="{ buy: currentCoin.change > 0, sell: currentCoin.change < 0 }">
                {{ currentCoin.rose }}
              </span>
            </div>
            <div class="item">
              <span class="text">{{ $t('service.high') }}</span>
              <span class="num">{{ currentCoin.high?.toFixed(4) }}</span>
            </div>
            <div class="item">
              <span class="text">{{ $t('service.low') }}</span>
              <span class="num">{{ currentCoin.low?.toFixed(4) }}</span>
            </div>
            <div class="item">
              <span class="text">{{ $t('service.ExchangeNum') }}</span>
              <span class="num">{{ currentCoin.volume?.toFixed(4) }} {{ currentCoin.coin }}</span>
            </div>
          </div>

          <!-- K 线图和深度图 -->
          <div class="imgtable panel-shell chart-shell">
            <div class="handler">
              <span @click="changeImgTable('k')" :class="{ active: currentImgTable === 'k' }">{{ $t('swap.kline') }}</span>
              <span @click="changeImgTable('s')" :class="{ active: currentImgTable === 's' }">{{ $t('swap.depth') }}</span>
            </div>
            <div id="swap_kline_container" :class="{ hidden: currentImgTable === 's' }"></div>
            <DepthGraph :class="{ hidden: currentImgTable === 'k' }" ref="depthGraphRef" />
          </div>

          <!-- 我的合约账户 -->
          <div class="account-panel-wrapper">
            <div class="order panel-shell account-panel">
              <div class="panel-header panel-header-compact account-panel-header">
                <div>
                  <span class="panel-kicker">Assets</span>
                  <strong>{{ $t('swap.myswapaccount') }}</strong>
                </div>
                <router-link class="linkmore" to="/uc/contract-money">
                  {{ $t('swap.zijinhuazhuan') }}
                </router-link>
              </div>
              <div class="table swap-my-account">
                <div class="account-item">
                  <div>{{ $t('swap.accountquanyi') }}</div>
                  <div style="font-size: 12px">
                    <span>{{ (wallet.base + wallet.frozen)?.toFixed(baseCoinScale) }}(USD)</span>
                  </div>
                </div>
                <div class="account-item">
                  <div>{{ $t('swap.profitandloss') }}</div>
                  <div style="font-size: 12px">
                    <span>{{ wallet.profit?.toFixed(baseCoinScale) }}</span>
                  </div>
                </div>
                <div class="account-item">
                  <div>{{ $t('swap.principalAmount') }}</div>
                  <div style="font-size: 12px">
                    <span>{{ wallet.base?.toFixed(baseCoinScale) }}(USD)</span>
                  </div>
                </div>
                <div class="account-item">
                  <div>{{ $t('swap.positionAmount') }}</div>
                  <div style="font-size: 12px">
                    <span>{{ positionNum }}</span>
                  </div>
                </div>
                <div class="account-item">
                  <div>{{ $t('swap.frozenAmount') }}</div>
                  <div style="font-size: 12px">
                    <span>{{ wallet.frozen?.toFixed(baseCoinScale) }}(USD)</span>
                  </div>
                </div>
              </div>
            </div>
          </div>

          <!-- 开仓/平仓表单 -->
          <div class="trade-form-panel-wrapper">
            <div class="order panel-shell trade-form-panel">
              <div v-if="currentCoin.type === 'ALWAYS'" class="order-handler trade-form-header">
                <span
                  :style="{ width: currentCoin.type === 'ALWAYS' ? '50%' : '100%', textAlign: 'center' }"
                  @click="entrustChange(1)"
                  :class="{ active: form.entrustType === 1 }"
                >
                  {{ $t('swap.open') }}
                </span>
                <span
                  v-if="currentCoin.type === 'ALWAYS'"
                  style="width: 50%; text-align: center"
                  @click="entrustChange(2)"
                  :class="{ active: form.entrustType === 2 }"
                >
                  {{ $t('swap.close') }}
                </span>
              </div>
              <div class="table open-close">
                <div class="open trade-form-body">
                  <!-- 委托类型 -->
                  <div class="form-row-with-label">
                    <span class="form-label">委托类型：</span>
                    <el-radio-group v-model="form.type" type="button" size="default">
                      <el-radio label="2">{{ $t('swap.limited_price') }}</el-radio>
                      <el-radio label="3">{{ $t('swap.trigger_price') }}</el-radio>
                    </el-radio-group>
                  </div>

                  <!-- 账户保证金 -->
                  <div class="form-row">
                    <span class="form-label">{{ $t('swap.accountmargin') }}</span>
                    <el-button
                      @click="showAdjustLeverage(1)"
                      size="small"
                      class="leverage-btn"
                    >
                      {{ form.leverage }}X
                    </el-button>
                  </div>

                  <!-- 时间选择（秒合约） -->
                  <div
                    v-if="currentCoin.type === 'SECOND'"
                    class="form-row"
                  >
                    <span class="form-label">{{ $t('swap.time') }}</span>
                    <el-button
                      @click="showAdjustTime(1)"
                      size="small"
                      class="leverage-btn"
                    >
                      {{ form.time }}S
                    </el-button>
                  </div>

                  <el-form class="trade-form">
                    <!-- 触发价格 -->
                    <el-form-item v-if="form.type === '3'" class="trade-form-item">
                      <label class="form-item-label">
                        {{ $t('swap.triggerPrice') }}
                      </label>
                      <div class="form-input-wrap">
                        <el-input-number
                          v-model="form.triggerPrice"
                          :min="0"
                          :precision="baseCoinScale"
                        />
                      </div>
                    </el-form-item>

                    <!-- 价格 -->
                    <el-form-item class="trade-form-item">
                      <label class="form-item-label">
                        {{ $t('swap.price') }}
                      </label>
                      <div class="form-input-wrap">
                        <el-button
                          v-if="form.market"
                          @click="form.market = !form.market"
                          type="primary"
                          class="market-toggle-btn"
                          shape="circle"
                        >
                          {{ $t('swap.price') }}
                        </el-button>
                        <el-button
                          v-else
                          @click="form.market = !form.market"
                          type="primary"
                          class="market-toggle-btn"
                          shape="circle"
                        >
                          {{ $t('swap.marketPrice') }}
                        </el-button>
                        <el-input-number
                          v-if="!form.market"
                          v-model="form.limitPrice"
                          :min="0"
                          :precision="baseCoinScale"
                        />
                        <el-input
                          v-else
                          :value="$t('swap.marketPrice')"
                          disabled
                        />
                      </div>
                    </el-form-item>

                    <!-- 数量 -->
                    <el-form-item class="trade-form-item">
                      <label class="form-item-label">
                        {{ $t('swap.num') }}
                      </label>
                      <div class="form-input-wrap">
                        <el-input-number
                          v-model="form.limitAmount"
                          :min="1"
                          :precision="0"
                        />
                      </div>
                    </el-form-item>
                  </el-form>

                  <!-- 登录后可交易 -->
                  <template v-if="isLogin">
                    <template v-if="form.entrustType === 1">
                      <div class="trade-actions">
                        <div class="trade-action-item buy-action">
                          <div class="trade-action-info">
                            <span class="green">{{ $t('swap.canup') }}:</span>
                            <span class="num">{{ maxOpen }}</span>
                          </div>
                          <el-button
                            @click="openOrder(1)"
                            type="primary"
                            shape="circle"
                            class="open-up operate_btn"
                          >
                            {{ $t('swap.openup') }}
                          </el-button>
                        </div>
                        <div class="trade-action-item sell-action">
                          <div class="trade-action-info">
                            <span class="red">{{ $t('swap.candown') }}:</span>
                            <span class="num">{{ maxOpen }}</span>
                          </div>
                          <el-button
                            @click="openOrder(2)"
                            type="primary"
                            shape="circle"
                            class="open-down operate_btn"
                          >
                            {{ $t('swap.opendown') }}
                          </el-button>
                        </div>
                      </div>
                    </template>
                    <template v-else>
                      <div class="trade-actions">
                        <div class="trade-action-item sell-action">
                          <div class="trade-action-info">
                            <span class="red">{{ $t('swap.canclosedown') }}:</span>
                            <span class="num">{{ maxCloseSellAmount }}</span>
                          </div>
                          <el-button
                            @click="closeOrder(2)"
                            type="primary"
                            shape="circle"
                            class="open-down operate_btn"
                          >
                            {{ $t('swap.closedown') }}
                          </el-button>
                        </div>
                        <div class="trade-action-item buy-action">
                          <div class="trade-action-info">
                            <span class="green">{{ $t('swap.cancloseup') }}:</span>
                            <span class="num">{{ maxCloseBuyAmount }}</span>
                          </div>
                          <el-button
                            @click="closeOrder(1)"
                            type="primary"
                            shape="circle"
                            class="open-up operate_btn"
                          >
                            {{ $t('swap.closeup') }}
                          </el-button>
                        </div>
                      </div>
                    </template>
                  </template>

                  <!-- 未登录提示 -->
                  <div v-if="!isLogin" class="operate-login">
                    <div class="login-tip">
                      {{ $t('common.please') }}
                      <router-link to="/login"><span style="color: #f0a70a">{{ $t('common.login') }}</span></router-link> /
                      <router-link to="/register"><span style="color: #00dcff">{{ $t('common.register') }}</span></router-link>
                    </div>
                  </div>
                </div>
              </div>
            </div>
          </div>
        </div>

        <!-- 底部：订单记录 -->
        <div class="workspace-orders">
          <div class="order panel-shell ledger-panel">
            <div class="order-handler">
              <span @click="changeOrder('currentPositions')" :class="{ active: selectedOrder === 'currentPositions' }">
                {{ $t('swap.currentposition') }}
              </span>
              <span @click="changeOrder('current')" :class="{ active: selectedOrder === 'current' }">
                {{ $t('swap.curdelegation') }}
              </span>
              <span @click="changeOrder('history')" :class="{ active: selectedOrder === 'history' }">
                {{ $t('swap.hisdelegation') }}
              </span>
              <router-link v-show="selectedOrder === 'current'" class="linkmore" to="/uc/contract/entrust/current">
                {{ $t('common.more') }}
              </router-link>
              <router-link v-show="selectedOrder === 'history'" class="linkmore" to="/uc/contract/entrust/history">
                {{ $t('common.more') }}
              </router-link>
            </div>
            <div class="table">
              <!-- 当前持仓 -->
              <el-table
                v-if="selectedOrder === 'currentPositions'"
                :data="currentPosition.rows"
                height="320"
                :no-data-text="$t('common.nodata')"
              >
                <el-table-column label="#" width="180" height="40">
                  <template #default="{ row }">
                    <span :class="row.direction?.toLowerCase()">
                      {{ row.contractCoinName }}{{ row.contractCoinType === 'SECOND' ? '.' + row.holdTime + 'SEC' : '' }}.{{ row.multiple }}X.{{ row.direction === 'BUY' ? $t('swap.up') : $t('swap.down') }}
                    </span>
                  </template>
                </el-table-column>
                <el-table-column label="开仓均价" key="openPrice">
                  <template #default="{ row }">
                    {{ row.openPrice?.toFixed(2) }}
                  </template>
                </el-table-column>
                <el-table-column label="盈亏" key="profit">
                  <template #default="{ row }">
                    <span :class="row.profit > 0 ? 'buy' : 'sell'">{{ row.profit }}</span>
                  </template>
                </el-table-column>
                <el-table-column label="保证金" key="principalAmount">
                  <template #default="{ row }">
                    {{ toFloor(row.principalAmount, baseCoinScale) }}
                  </template>
                </el-table-column>
                <el-table-column label="持仓数量" key="position" />
                <el-table-column label="可平数量" key="availablePosition" />
                <el-table-column label="强平价格" key="liquidationPrice">
                  <template #default="{ row }">
                    {{ toFloor(row.liquidationPrice, 2) }}
                  </template>
                </el-table-column>
                <el-table-column label="操作" key="operate" width="110">
                  <template #default="{ row }">
                    <div>
                      <el-button
                        v-if="row.contractCoinType !== 'SECOND'"
                        size="small"
                        :class="row.direction === 'BUY' ? 'bg-red' : 'bg-green'"
                        @click="quickClose(row)"
                      >
                        {{ $t('swap.quickclose') }}
                      </el-button>
                    </div>
                  </template>
                </el-table-column>
              </el-table>

              <!-- 当前委托 -->
              <div v-if="selectedOrder === 'current'">
                <div class="order-type-btn-wrap" style="height: 40px">
                  <div class="order-type-btn-wrap-left" style="padding-top: 5px">
                    <button @click="setSelectedType(2)" :class="{ 'btn-selected': selectedType == 2 }">
                      {{ $t('swap.limited_price') }}
                    </button>
                    <button @click="setSelectedType(3)" :class="{ 'btn-selected': selectedType == 3 }">
                      {{ $t('swap.trigger_price') }}
                    </button>
                  </div>
                </div>
                <el-table
                  v-show="selectedType == 2"
                  :data="currentOrder.rows"
                  height="280"
                  :no-data-text="$t('common.nodata')"
                >
                  <el-table-column label="#" width="180">
                    <template #default="{ row }">
                      <span :class="row.direction?.toLowerCase()">
                        {{ row.contractCoinName }}{{ row.contractCoinType === 'SECOND' ? '.' + row.holdTime + 'SEC' : '' }}.{{ row.leverage }}X.{{ row.direction === 'BUY' ? (row.entrustType === 'OPEN' ? $t('swap.openLong') : $t('swap.closeShort')) : (row.entrustType === 'OPEN' ? $t('swap.openShort') : $t('swap.closeLong')) }}
                      </span>
                    </template>
                  </el-table-column>
                  <el-table-column label="委托类型">
                    <template #default="{ row }">
                      {{ row.type === 'LIMIT_PRICE' ? '限价委托' : '计划委托' }}
                    </template>
                  </el-table-column>
                  <el-table-column label="价格">
                    <template #default="{ row }">
                      {{ row.marketPrice ? '对手价' : toFloor(row.entrustPrice, baseCoinScale) }}
                    </template>
                  </el-table-column>
                  <el-table-column label="数量">
                    <template #default="{ row }">
                      {{ row.share }}
                    </template>
                  </el-table-column>
                  <el-table-column label="打仓价格">
                    <template #default="{ row }">
                      {{ row.strikePrice > 0 ? toFloor(row.strikePrice, baseCoinScale) : '-' }}
                    </template>
                  </el-table-column>
                  <el-table-column label="保证金">
                    <template #default="{ row }">
                      {{ toFloor(row.principalAmount, baseCoinScale) }}
                    </template>
                  </el-table-column>
                  <el-table-column label="操作" key="operate" width="110">
                    <template #default="{ row, $index }">
                      <el-button
                        size="small"
                        style="border: 1px solid #f0ac19; color: #f1ac19; line-height: 1.2; border-radius: 10px"
                        @click="cancel($index)"
                      >
                        {{ $t('swap.undo') }}
                      </el-button>
                    </template>
                  </el-table-column>
                </el-table>
                <el-table
                  v-show="selectedType == 3"
                  :data="currentOrder.rows"
                  height="280"
                  :no-data-text="$t('common.nodata')"
                >
                  <el-table-column label="#" width="180">
                    <template #default="{ row }">
                      <span :class="row.direction?.toLowerCase()">
                        {{ row.contractCoinName }}{{ row.contractCoinType === 'SECOND' ? '.' + row.holdTime + 'SEC' : '' }}.{{ row.leverage }}X.{{ row.direction === 'BUY' ? (row.entrustType === 'OPEN' ? $t('swap.openLong') : $t('swap.closeShort')) : (row.entrustType === 'OPEN' ? $t('swap.openShort') : $t('swap.closeLong')) }}
                      </span>
                    </template>
                  </el-table-column>
                  <el-table-column label="价格">
                    <template #default="{ row }">
                      {{ row.marketPrice ? '对手价' : toFloor(row.entrustPrice, baseCoinScale) }}
                    </template>
                  </el-table-column>
                  <el-table-column label="数量">
                    <template #default="{ row }">
                      {{ row.share }}
                    </template>
                  </el-table-column>
                  <el-table-column label="触发价格">
                    <template #default="{ row }">
                      {{ row.triggerPrice > 0 ? toFloor(row.triggerPrice, baseCoinScale) : '--' }}
                    </template>
                  </el-table-column>
                  <el-table-column label="打仓价格">
                    <template #default="{ row }">
                      {{ row.strikePrice > 0 ? toFloor(row.strikePrice, baseCoinScale) : '-' }}
                    </template>
                  </el-table-column>
                  <el-table-column label="保证金">
                    <template #default="{ row }">
                      {{ toFloor(row.principalAmount, baseCoinScale) }}
                    </template>
                  </el-table-column>
                  <el-table-column label="操作" key="operate" width="110">
                    <template #default="{ row, $index }">
                      <el-button
                        size="small"
                        style="border: 1px solid #f0ac19; color: #f1ac19; line-height: 1.2; border-radius: 10px"
                        @click="cancel($index)"
                      >
                        {{ $t('swap.undo') }}
                      </el-button>
                    </template>
                  </el-table-column>
                </el-table>
              </div>

              <!-- 历史委托 -->
              <div v-if="selectedOrder === 'history'">
                <div class="order-type-btn-wrap" style="height: 40px">
                  <div class="order-type-btn-wrap-left" style="padding-top: 5px">
                    <button @click="setSelectedType(2)" :class="{ 'btn-selected': selectedType == 2 }">
                      {{ $t('swap.limited_price') }}
                    </button>
                    <button @click="setSelectedType(3)" :class="{ 'btn-selected': selectedType == 3 }">
                      {{ $t('swap.trigger_price') }}
                    </button>
                  </div>
                </div>
                <el-table
                  v-show="selectedType == 2"
                  :data="historyOrder.rows"
                  height="280"
                  :no-data-text="$t('common.nodata')"
                >
                  <el-table-column label="#" width="180">
                    <template #default="{ row }">
                      <span :class="row.direction?.toLowerCase()">
                        {{ row.contractCoinName }}{{ row.contractCoinType === 'SECOND' ? '.' + row.holdTime + 'SEC' : '' }}.{{ row.leverage }}X.{{ row.direction === 'BUY' ? (row.entrustType === 'OPEN' ? $t('swap.openLong') : $t('swap.closeShort')) : (row.entrustType === 'OPEN' ? $t('swap.openShort') : $t('swap.closeLong')) }}
                      </span>
                    </template>
                  </el-table-column>
                  <el-table-column label="委托类型">
                    <template #default="{ row }">
                      {{ row.type === 'LIMIT_PRICE' ? '限价委托' : '计划委托' }}
                    </template>
                  </el-table-column>
                  <el-table-column label="价格">
                    <template #default="{ row }">
                      {{ row.marketPrice ? '对手价' : toFloor(row.entrustPrice, baseCoinScale) }}
                    </template>
                  </el-table-column>
                  <el-table-column label="数量">
                    <template #default="{ row }">
                      {{ row.share }}
                    </template>
                  </el-table-column>
                  <el-table-column label="打仓价格">
                    <template #default="{ row }">
                      {{ row.strikePrice > 0 ? toFloor(row.strikePrice, baseCoinScale) : '--' }}
                    </template>
                  </el-table-column>
                  <el-table-column label="手续费">
                    <template #default="{ row }">
                      <span class="sell">{{ !row.fee ? '--' : '-' + toFloor(row.fee, baseCoinScale) }}</span>
                    </template>
                  </el-table-column>
                  <el-table-column label="收益">
                    <template #default="{ row }">
                      <span :class="row.profit >= 0 ? 'buy' : 'sell'">
                        {{ (row.entrustType === 'CLOSE' && row.contractCoinType === 'ALWAYS') || (row.entrustType === 'OPEN' && row.contractCoinType === 'SECOND') ? (row.profit > 0 ? '+' : '') + toFloor(row.profit, baseCoinScale) : '--' }}
                      </span>
                    </template>
                  </el-table-column>
                  <el-table-column label="时间" key="createTime">
                    <template #default="{ row }">
                      {{ dateFormat(row.createTime) }}
                    </template>
                  </el-table-column>
                  <el-table-column label="状态" key="status">
                    <template #default="{ row }">
                      <span v-if="row.status === 'ENTRUST_SUCCESS'" style="color: #f0a70a">{{ $t('swap._success') }}</span>
                      <span v-else-if="row.status === 'ENTRUST_FAILURE'" style="color: #f15057">{{ $t('swap.failure') }}</span>
                      <span v-else-if="row.status === 'ENTRUST_CANCEL'" style="color: #7c7f82">{{ $t('swap.canceled') }}</span>
                      <span v-else>--</span>
                    </template>
                  </el-table-column>
                </el-table>
                <el-table
                  v-show="selectedType == 3"
                  :data="historyOrder.rows"
                  height="280"
                  :no-data-text="$t('common.nodata')"
                >
                  <el-table-column label="#" width="180">
                    <template #default="{ row }">
                      <span :class="row.direction?.toLowerCase()">
                        {{ row.contractCoinName }}{{ row.contractCoinType === 'SECOND' ? '.' + row.holdTime + 'SEC' : '' }}.{{ row.leverage }}X.{{ row.direction === 'BUY' ? (row.entrustType === 'OPEN' ? $t('swap.openLong') : $t('swap.closeShort')) : (row.entrustType === 'OPEN' ? $t('swap.openShort') : $t('swap.closeLong')) }}
                      </span>
                    </template>
                  </el-table-column>
                  <el-table-column label="价格">
                    <template #default="{ row }">
                      {{ row.marketPrice ? '对手价' : toFloor(row.entrustPrice, baseCoinScale) }}
                    </template>
                  </el-table-column>
                  <el-table-column label="数量">
                    <template #default="{ row }">
                      {{ row.share }}
                    </template>
                  </el-table-column>
                  <el-table-column label="触发价格">
                    <template #default="{ row }">
                      {{ row.triggerPrice > 0 ? toFloor(row.triggerPrice, baseCoinScale) : '--' }}
                    </template>
                  </el-table-column>
                  <el-table-column label="触发时间" key="triggeringTime">
                    <template #default="{ row }">
                      {{ row.triggeringTime ? dateFormat(row.triggeringTime) : '--' }}
                    </template>
                  </el-table-column>
                  <el-table-column label="打仓价格">
                    <template #default="{ row }">
                      {{ row.strikePrice > 0 ? toFloor(row.strikePrice, baseCoinScale) : '--' }}
                    </template>
                  </el-table-column>
                  <el-table-column label="手续费">
                    <template #default="{ row }">
                      <span class="sell">{{ !row.fee ? '--' : '-' + toFloor(row.fee, baseCoinScale) }}</span>
                    </template>
                  </el-table-column>
                  <el-table-column label="收益">
                    <template #default="{ row }">
                      <span :class="row.profit >= 0 ? 'buy' : 'sell'">
                        {{ row.entrustType === 'OPEN' ? '--' : (row.profit > 0 ? '+' : '') + toFloor(row.profit, baseCoinScale) }}
                      </span>
                    </template>
                  </el-table-column>
                  <el-table-column label="时间" key="createTime">
                    <template #default="{ row }">
                      {{ dateFormat(row.createTime) }}
                    </template>
                  </el-table-column>
                  <el-table-column label="状态" key="status">
                    <template #default="{ row }">
                      <span v-if="row.status === 'ENTRUST_SUCCESS'" style="color: #f0a70a">{{ $t('swap._success') }}</span>
                      <span v-else-if="row.status === 'ENTRUST_FAILURE'" style="color: #f15057">{{ $t('swap.failure') }}</span>
                      <span v-else-if="row.status === 'ENTRUST_CANCEL'" style="color: #7c7f82">{{ $t('swap.canceled') }}</span>
                      <span v-else>--</span>
                    </template>
                  </el-table-column>
                </el-table>
              </div>
            </div>
          </div>
        </div>
      </div>

      <!-- 右侧：盘口数据 -->
      <div class="orderbook-area panel-shell">
        <div class="panel-header panel-header-compact">
          <div>
            <span class="panel-kicker">Order Book</span>
            <strong>{{ $t('swap.depth') }}</strong>
          </div>
        </div>
        <!-- 盘口选择器 -->
        <div class="handlers">
          <span @click="changePlate('all')" class="handler handler-all" :class="{ active: selectedPlate === 'all' }"></span>
          <span @click="changePlate('buy')" class="handler handler-green" :class="{ active: selectedPlate === 'buy' }"></span>
          <span @click="changePlate('sell')" class="handler handler-red" :class="{ active: selectedPlate === 'sell' }"></span>
        </div>

        <!-- 卖盘 -->
        <el-table
          v-show="selectedPlate !== 'buy'"
          :data="plate.askRows"
          highlight-row
          @current-change="buyPlate"
          :no-data-text="$t('common.nodata')"
        >
          <el-table-column prop="price" label="价格">
            <template #default="{ row }">
              <span v-if="row.price === 0">--</span>
              <span v-else :class="row.direction?.toLowerCase()">{{ row.price?.toFixed(4) }}</span>
            </template>
          </el-table-column>
          <el-table-column prop="amount" label="数量">
            <template #default="{ row }">
              <span v-if="row.amount === 0">--</span>
              <span v-else>{{ row.amount?.toFixed(4) }}</span>
            </template>
          </el-table-column>
          <el-table-column prop="totalAmount" label="累计">
            <template #default="{ row }">
              <span v-if="row.price === 0 || row.totalAmount === 0">--</span>
              <span v-else>{{ row.totalAmount?.toFixed(4) }}</span>
            </template>
          </el-table-column>
          <el-table-column class-name="percenttd" width="1">
            <template #default="{ row }">
              <div
                v-if="row.totalAmount"
                class="percentdiv"
                :style="{
                  width: (row.totalAmount / (row.direction === 'BUY' ? plate.bidTotle : plate.askTotle)) * 100 + '%',
                  backgroundColor: row.direction === 'BUY' ? '#00b275' : '#f15057'
                }"
              />
            </template>
          </el-table-column>
        </el-table>

        <!-- 当前价格 -->
        <div class="plate-nowprice">
          <span class="price" :class="{ buy: currentCoin.change > 0, sell: currentCoin.change < 0 }">
            {{ currentCoin.price?.toFixed(baseCoinScale) }}
          </span>
          <span v-if="currentCoin.change > 0" class="buy">↑</span>
          <span v-else-if="currentCoin.change < 0" class="sell">↓</span>
          <span class="price-cny">≈ {{ (currentCoin.usdRate * CNYRate)?.toFixed(2) }} CNY</span>
        </div>

        <!-- 买盘 -->
        <el-table
          v-show="selectedPlate !== 'sell'"
          :data="plate.bidRows"
          highlight-row
          @current-change="sellPlate"
          :no-data-text="$t('common.nodata')"
        >
          <el-table-column prop="price" label="价格">
            <template #default="{ row }">
              <span v-if="row.price === 0">--</span>
              <span v-else :class="row.direction?.toLowerCase()">{{ row.price?.toFixed(4) }}</span>
            </template>
          </el-table-column>
          <el-table-column prop="amount" label="数量">
            <template #default="{ row }">
              <span v-if="row.amount === 0">--</span>
              <span v-else>{{ row.amount?.toFixed(4) }}</span>
            </template>
          </el-table-column>
          <el-table-column prop="totalAmount" label="累计">
            <template #default="{ row }">
              <span v-if="row.price === 0 || row.totalAmount === 0">--</span>
              <span v-else>{{ row.totalAmount?.toFixed(4) }}</span>
            </template>
          </el-table-column>
          <el-table-column class-name="percenttd" width="1">
            <template #default="{ row }">
              <div
                v-if="row.totalAmount"
                class="percentdiv"
                :style="{
                  width: (row.totalAmount / (row.direction === 'BUY' ? plate.bidTotle : plate.askTotle)) * 100 + '%',
                  backgroundColor: row.direction === 'BUY' ? '#00b275' : '#f15057'
                }"
              />
            </template>
          </el-table-column>
        </el-table>
      </div>
    </div>

    <!-- 弹窗：调整杠杆 -->
    <el-dialog v-model="form.leverageModal" :title="$t('swap.modifyLeverage')" width="20%" class="vertical-center-modal">
      <div class="leverage">
        <el-slider
          v-model="form.leverageAdjustVal2"
          @input="xxlever"
          @change="xxlever"
          :marks="leverages"
          :min="1"
          :max="5"
        />
      </div>
      <template #footer>
        <el-button type="default" size="large" @click="form.leverageModal = false">{{ $t('common.close') }}</el-button>
        <el-button type="primary" size="large" @click="adjustLeverage">{{ $t('common.ok') }}</el-button>
      </template>
    </el-dialog>

    <!-- 弹窗：调整时间 -->
    <el-dialog v-model="form.timeModal" :title="$t('swap.time')" width="20%" class="vertical-center-modal">
      <div class="leverage">
        <el-slider
          v-model="form.timeAdjustVal2"
          @input="xxlever2"
          @change="xxlever2"
          :marks="times"
          :min="1"
          :max="6"
        />
      </div>
      <template #footer>
        <el-button type="default" size="large" @click="form.timeModal = false">{{ $t('common.close') }}</el-button>
        <el-button type="primary" size="large" @click="adjustTime">{{ $t('common.ok') }}</el-button>
      </template>
    </el-dialog>
  </div>
</template>

<script setup>
/**
 * Vue 3 迁移 - 合约交易主页面 (Swapindex)
 * 迁移点:
 * 1. 使用 <script setup> 语法和 Composition API
 * 2. Element Plus 组件替代 iView 组件
 * 3. 使用 inject('store') 和 inject('router') 兼容 Vuex 3.x 和 Vue Router 3.x
 * 4. 使用 axios 替代 vue-resource
 * 5. 使用辅助函数替代 Vue 2 filters
 * 6. 使用 table column slots 替代 render functions
 */
import { ref, reactive, computed, inject, onMounted, onBeforeUnmount, watch } from 'vue'
import { useRoute } from 'vue-router'
import { ElMessage, ElNotification, ElMessageBox } from 'element-plus'
import { InfoFilled, Search } from '@element-plus/icons-vue'
import axios from 'axios'
import moment from 'moment'
import io from 'socket.io-client'
import { useRuntimeContract } from '../../config/runtime-vue3'

// 导入子组件
import DepthGraph from '../../components/exchange/DepthGraph.vue'

// 导入 TradingView datafeed
import Datafeeds from '../../assets/js/charting_library/datafeed/swaptrade.js'

// Vuex 3.x 和 Vue Router 3.x 兼容
const store = inject('store')
const router = inject('router')
const route = useRoute()

// 常量
const runtime = useRuntimeContract()
const host = runtime.host
const api = runtime.api
const silentPost = (url, payload = {}) => axios.post(url, payload).catch(() => null)
const silentGet = (url) => axios.get(url).catch(() => null)
const unwrapResult = (response) => response?.data?.data ?? response?.data ?? {}
const unwrapList = (response) => {
  const payload = unwrapResult(response)
  if (Array.isArray(payload)) return payload
  if (Array.isArray(payload.list)) return payload.list
  if (Array.isArray(payload.content)) return payload.content
  return []
}
const normalizeTradePlate = (payload) => {
  const source = payload?.data ?? payload ?? {}
  const askItems = source.ask?.items || source.asks || []
  const bidItems = source.bid?.items || source.bids || []
  const ask = source.ask || {
    direction: 'SELL',
    items: askItems,
    highestPrice: askItems.length > 0 ? askItems[askItems.length - 1].price : 0,
    lowestPrice: askItems.length > 0 ? askItems[0].price : 0
  }
  const bid = source.bid || {
    direction: 'BUY',
    items: bidItems,
    highestPrice: bidItems.length > 0 ? bidItems[0].price : 0,
    lowestPrice: bidItems.length > 0 ? bidItems[bidItems.length - 1].price : 0
  }
  return { ask, bid }
}

// 响应式数据
const skin = ref('night')
const currentImgTable = ref('k')
const selectedOrder = ref('currentPositions')
const selectedPlate = ref('all')
const selectedType = ref(2)
const CNYRate = ref(null)
const datafeed = ref(null)
const defaultPath = ref(1)
const basecion = ref('usdt')
const coinScale = ref(6)
const baseCoinScale = ref(4)
const takerFee = ref(0.001)
const makerFee = ref(0.001)
const memberId = ref(0)
const dataIndex = ref([])
const dataIndex2 = ref([])
const searchKey = ref('')
const isLogin = ref(false)
const lang = ref('zh')

// 当前币种
const currentCoin = reactive({
  id: '',
  base: '',
  coin: '',
  symbol: '',
  name: '',
  perUsdt: 100000,
  close: 0,
  price: 0,
  rose: '',
  high: 0,
  low: 0,
  volume: 0,
  turnover: 0,
  usdRate: 0,
  change: 0,
  type: 'ALWAYS',
  coinScale: 6,
  baseCoinScale: 4
})

const coinInfo = reactive({
  name: '',
  information: '',
  infolink: ''
})

// 杠杆和时间选项
const leverages = { 1: '1x', 2: '2x', 3: '5x', 4: '10x', 5: '20x' }
const times = { 1: '10S', 2: '30S', 3: '60S', 4: '120S', 5: '360S', 6: '3000S' }

// 交易控制
const enableMarketBuy = ref(1)
const enableMarketSell = ref(1)
const exchangeable = ref(1)

// 交易表单
const form = reactive({
  type: '2',
  time: 10,
  patterns: '1',
  entrustType: 1,
  limitPrice: 0.0,
  triggerPrice: 0,
  limitAmount: 0,
  marketAmount: 0.0,
  limitTurnover: 0.0,
  limitPercent: 0,
  leverageModal: false,
  leverage: 10,
  timeAdjustVal: 10,
  timeAdjustVal2: 1,
  leverageAdjustVal: 10,
  leverageAdjustVal2: 4,
  market: false,
  timeModal: false
})

// 钱包
const wallet = reactive({
  base: 0.0,
  frozen: 0.0,
  profit: 0.0,
  kick: 0.0
})

// 买卖盘数据
const plate = reactive({
  maxPostion: 10,
  columns: [],
  askRows: [],
  bidRows: [],
  askTotle: 0,
  bidTotle: 0
})

// 成交记录
const trade = reactive({
  rows: [],
  columns: []
})

// 当前持仓
const currentPosition = reactive({
  columns: [],
  rows: []
})

// 当前委托
const currentOrder = reactive({
  columns: [],
  columns2: [],
  rows: []
})

// 历史委托
const historyOrder = reactive({
  pageSize: 10,
  total: 10,
  page: 0,
  columns: [],
  columns2: [],
  rows: []
})

// 币种列表
const coins = reactive({
  _map: [],
  USDT: [],
  BTC: [],
  ETH: [],
  USDT2: [],
  favor: [],
  columns: []
})

// 深度图引用
const depthGraphRef = ref(null)
const socketRef = ref(null)

// 计算属性
const positionNum = computed(() => {
  let num = 0
  let buyNum = 0
  let sellNum = 0
  currentPosition.rows.forEach(e => {
    if (e.direction === 'BUY') {
      buyNum += e.availablePosition
    } else {
      sellNum += e.availablePosition
    }
    num += e.position
  })
  maxCloseSellAmount.value = buyNum
  maxCloseBuyAmount.value = sellNum
  return num
})

const maxOpen = computed(() => {
  const balance = wallet.base
  const perUsdt = currentCoin.perUsdt
  const leverage = form.leverage
  if (balance <= 0 || !perUsdt || !leverage) {
    return 0
  }
  return parseInt((balance * leverage) / perUsdt)
})

const maxCloseBuyAmount = ref(0)
const maxCloseSellAmount = ref(0)

const sliderBuyDisabled = computed(() => {
  const min = toFloor(1 / Math.pow(10, baseCoinScale.value))
  return wallet.base < min
})

const sliderSellDisabled = computed(() => {
  const min = toFloor(1 / Math.pow(10, coinScale.value))
  return wallet.base < min
})

const rechargeUSDTUrl = computed(() => `/uc/recharge?name=${currentCoin.base}`)
const rechargeCoinUrl = computed(() => `/uc/recharge?name=${currentCoin.coin}`)

// 工具函数
const toFixed = (value, scale = 4) => {
  if (value === null || value === undefined) return 0
  return Number(value).toFixed(scale)
}

const toFloor = (number, scale = 8) => {
  if (Number(number) === 0) return 0
  const str = number + ''
  if (str.indexOf('e') > -1 || str.indexOf('E') > -1) {
    const num = Number(number).toFixed(scale + 1)
    return (num + '').substring(0, (num + '').length - 1)
  } else if (str.indexOf('.') > -1) {
    if (scale === 0) return str.substring(0, str.indexOf('.'))
    return str.substring(0, str.indexOf('.') + scale + 1)
  }
  return str
}

const timeFormat = (tick) => {
  return moment(tick).format('HH:mm:ss')
}

const dateFormat = (tick) => {
  return moment(tick).format('YYYY-MM-DD HH:mm:ss')
}

// 方法实现
const xxlever = (value) => {
  console.log(value)
  if (value === 1) {
    form.leverageAdjustVal = 1
  } else if (value === 2) {
    form.leverageAdjustVal = 2
  } else if (value === 3) {
    form.leverageAdjustVal = 5
  } else if (value === 4) {
    form.leverageAdjustVal = 10
  } else if (value === 5) {
    form.leverageAdjustVal = 20
  }
}

const xxlever2 = (value) => {
  console.log(value)
  if (value === 1) {
    form.timeAdjustVal = 10
  } else if (value === 2) {
    form.timeAdjustVal = 30
  } else if (value === 3) {
    form.timeAdjustVal = 60
  } else if (value === 4) {
    form.timeAdjustVal = 120
  } else if (value === 5) {
    form.timeAdjustVal = 360
  } else if (value === 6) {
    form.timeAdjustVal = 3000
  }
}

const getLeverage = () => {
  axios.post(`${host}${api.swap.leverage}`, {
    memberId: memberId.value,
    contractCoinId: currentCoin.id
  })
    .then(response => {
      const resp = unwrapResult(response)
      form.leverage = resp.leverage
      if (resp.leverage === 1) {
        form.leverageAdjustVal2 = 1
      } else if (resp.leverage === 2) {
        form.leverageAdjustVal2 = 2
      } else if (resp.leverage === 5) {
        form.leverageAdjustVal2 = 3
      } else if (resp.leverage === 10) {
        form.leverageAdjustVal2 = 4
      } else if (resp.leverage === 20) {
        form.leverageAdjustVal2 = 5
      }
    })
}

const quickClose = (order) => {
  if (order.availablePosition <= 0) {
    return
  }
  ElMessageBox.confirm(
    `${$t('swap.closetip')}${order.availablePosition}`,
    $t('swap.tip'),
    {
      confirmButtonText: $t('common.ok'),
      cancelButtonText: $t('common.cancel'),
      type: 'warning'
    }
  ).then(() => {
    axios.post(`${host}${api.swap.quickClose}${order.contractCoinId || currentCoin.id}`, {
      memberId: memberId.value,
      contractCoinId: order.contractCoinId || currentCoin.id,
      price: currentCoin.close
    })
      .then(response => {
        const resp = response.data
        if (resp.code === 0) {
          ElNotification.success({
            title: $t('swap.tip'),
            message: $t('swap.success')
          })
          refreshAccount()
        } else {
          ElNotification.error({
            title: $t('swap.tip'),
            message: resp.message
          })
        }
      })
  }).catch(() => {})
}

const setSelectedType = (e) => {
  selectedType.value = e
  getCurrentOrder()
  getHistoryOrder()
}

const closeOrder = (direction) => {
  if (direction === 1 && form.limitAmount > maxCloseBuyAmount.value) {
    ElNotification.error({
      title: $t('swap.tip'),
      message: $t('swap.buyamounttipwarning') + maxCloseBuyAmount.value
    })
    return
  }
  if (direction === 2 && form.limitAmount > maxCloseSellAmount.value) {
    ElNotification.error({
      title: $t('swap.tip'),
      message: $t('swap.sellamounttipwarning') + maxCloseSellAmount.value
    })
    return
  }
  const triggerType = form.triggerPrice <= currentCoin.close ? 1 : 2
  axios.post(`${host}${api.swap.addOrderEntrust}`, {
    memberId: memberId.value,
    contractCoinId: currentCoin.id,
    entrustType: form.type === '2' ? 1 : 2,
    type: form.entrustType,
    leverage: form.leverage,
    price: form.market ? currentCoin.close : Number(form.limitPrice || 0),
    amount: Number(form.limitAmount || 0),
    direction
  }).then(response => {
    const resp = response.data
    if (resp.code === 0) {
      ElNotification.success({
        title: $t('swap.tip'),
        message: $t('swap.success')
      })
      refreshAccount()
      form.limitPrice = 0
    } else {
      ElNotification.error({
        title: $t('swap.tip'),
        message: resp.message
      })
    }
  })
}

const openOrder = (direction) => {
  if (form.limitAmount > maxOpen.value) {
    ElNotification.error({
      title: $t('swap.tip'),
      message: (direction === 1 ? $t('swap.buyamounttipwarning') : $t('swap.sellamounttipwarning')) + maxOpen.value
    })
    return
  }
  let feeEnable = true
  if (direction === 1 && form.entrustType === 1 && parseInt(form.limitAmount) * currentCoin.perUsdt * takerFee.value > wallet.kick) {
    feeEnable = false
  } else if (direction === 2 && form.entrustType === 1 && parseInt(form.limitAmount) * currentCoin.perUsdt * makerFee.value > wallet.kick) {
    feeEnable = false
  }

  if (!feeEnable) {
    ElMessageBox.confirm(
      $t('common.feetipdetail'),
      $t('common.feetip'),
      {
        confirmButtonText: $t('common.ok'),
        cancelButtonText: $t('common.feeclosetip'),
        type: 'warning'
      }
    ).then(() => {
      router.push('/crowdfunding')
    }).catch(() => {
      doOpen(direction)
    })
  } else {
    doOpen(direction)
  }
}

const doOpen = (direction) => {
  axios.post(`${host}${api.swap.addOrderEntrust}`, {
    memberId: memberId.value,
    contractCoinId: currentCoin.id,
    entrustType: form.type === '2' ? 1 : 2,
    type: form.entrustType,
    leverage: form.leverage,
    price: form.market ? currentCoin.close : Number(form.limitPrice || 0),
    amount: Number(form.limitAmount || 0),
    direction
  }).then(response => {
    const resp = response.data
    if (resp.code === 0) {
      ElNotification.success({
        title: $t('swap.tip'),
        message: $t('swap.success')
      })
      refreshAccount()
      form.limitPrice = 0
    } else {
      ElNotification.error({
        title: $t('swap.tip'),
        message: resp.message
      })
    }
  })
}

const seachInputChange = () => {
  searchKey.value = searchKey.value.toUpperCase()
  dataIndex.value = coins.USDT.filter(item => item.coin.indexOf(searchKey.value) === 0)
}

const seachInputChange2 = () => {
  searchKey.value = searchKey.value.toUpperCase()
  dataIndex2.value = coins.USDT2.filter(item => item.coin.indexOf(searchKey.value) === 0)
}

const tipFormat = (val) => `${val}%`

const changeBaseCion = (str) => {
  basecion.value = str
  dataIndex.value = coins.USDT
  dataIndex2.value = coins.USDT2
}

const changePlate = (str) => {
  if (str !== 'all') {
    plate.maxPostion = 20
  } else {
    plate.maxPostion = 10
  }
  getPlate(str)
}

const changeImgTable = (str) => {
  currentImgTable.value = str
}

const changeOrder = (str) => {
  selectedOrder.value = str
}

const entrustChange = (val) => {
  form.entrustType = val
}

const gohref = (currentRow) => {
  router.push({
    name: 'SwapPair',
    params: { pair: currentRow.href }
  })
}

const buyPlate = (currentRow) => {
  form.limitPrice = parseFloat(currentRow.price)
  if (form.type === '3') {
    form.triggerPrice = parseFloat(currentRow.price)
  }
}

const sellPlate = (currentRow) => {
  form.limitPrice = parseFloat(currentRow.price)
  if (form.type === '3') {
    form.triggerPrice = parseFloat(currentRow.price)
  }
}

const getWallet = () => {
  axios.post(`${host}/uc/contract-wallet`, {
    memberId: memberId.value
  })
    .then(response => {
      const rows = unwrapList(response)
      const usdtWallet = rows.find(item => item?.coin?.unit === 'USDT') || rows[0]
      if (usdtWallet) {
        wallet.base = Number(usdtWallet.balance || 0)
        wallet.frozen = Number(usdtWallet.frozenBalance || 0)
      }
    })
  axios.post(`${host}${api.uc.wallet}KICK`, {})
    .then(response => {
      const resp = response.data
      wallet.kick = resp.data.balance
    })
}

const getCurrentPosition = () => {
  const params = { memberId: memberId.value, contractCoinId: currentCoin.id }
  currentPosition.rows = []
  axios.post(`${host}${api.swap.position}`, params)
    .then(response => {
      const resp = unwrapResult(response)
      const rows = resp.list || resp
      if (rows && rows.length > 0) {
        currentPosition.rows = rows.map(e => ({ ...e, profit: 0.0 }))
      }
    })
}

const getCurrentOrder = () => {
  const params = {
    memberId: memberId.value,
      pageNo: 0,
      pageSize: 100,
      contractCoinId: currentCoin.id,
      type: Number(selectedType.value)
    }
  currentOrder.rows = []
  axios.post(`${host}${api.swap.current}`, params)
    .then(response => {
      const resp = unwrapResult(response)
      const rows = resp.content || resp.list || []
      if (rows.length > 0) {
        currentOrder.rows = rows
      }
    })
}

const getHistoryOrder = (pageNo) => {
  if (pageNo === undefined) {
    pageNo = historyOrder.page
  } else {
    pageNo = pageNo - 1
  }
  historyOrder.rows = []
  const params = {
    memberId: memberId.value,
    pageNo,
      pageSize: historyOrder.pageSize,
      contractCoinId: currentCoin.id,
      type: Number(selectedType.value)
    }
  axios.post(`${host}${api.swap.history}`, params)
    .then(response => {
      const resp = unwrapResult(response)
      const rows = resp.content || resp.list || []
      if (rows.length > 0) {
        historyOrder.total = resp.totalElements ?? rows.length
        historyOrder.page = resp.number ?? pageNo
        historyOrder.rows = rows
      }
    })
}

const cancel = (index) => {
  const order = currentOrder.rows[index]
  ElMessageBox.confirm(
    $t('swap.undotip'),
    $t('swap.tip'),
    {
      confirmButtonText: $t('common.ok'),
      cancelButtonText: $t('common.cancel'),
      type: 'warning'
    }
  ).then(() => {
    axios.post(`${host}${api.swap.entrustCancel}/${order.id}`, {})
      .then(response => {
        const resp = response.data
        if (resp.code === 0) {
          refreshAccount()
          ElNotification.success({
            title: $t('swap.tip'),
            message: $t('swap.cancelsuccess')
          })
        } else {
          ElNotification.error({
            title: $t('swap.tip'),
            message: resp.message
          })
        }
      })
  }).catch(() => {})
}

const refreshAccount = () => {
  getCurrentOrder()
  getHistoryOrder()
  getCurrentPosition()
  getWallet()
}

const keyEvent = () => {
  const re1 = new RegExp(`([0-9]+.[0-9]{${baseCoinScale.value}})[0-9]*`, '')
  form.limitPrice = form.limitPrice.toString().replace(re1, '$1')
  form.limitPrice = form.limitPrice.toString().replace(re1, '$1')
  form.marketAmount = form.marketAmount.toString().replace(re1, '$1')

  const re2 = new RegExp(`([0-9]+.[0-9]{${coinScale.value}})[0-9]*`, '')
  form.limitAmount = form.limitAmount.toString().replace(re2, '$1')
  form.limitAmount = form.limitAmount.toString().replace(re2, '$1')
  form.marketAmount = form.marketAmount.toString().replace(re2, '$1')
}

const showAdjustLeverage = (type) => {
  form.leverageModal = true
  form.leverageAdjustVal = form.leverage
}

const showAdjustTime = () => {
  form.timeModal = true
  form.timeAdjustVal = form.time
  if (form.time === 10) {
    form.timeAdjustVal2 = 1
  } else if (form.time === 30) {
    form.timeAdjustVal2 = 2
  } else if (form.time === 60) {
    form.timeAdjustVal2 = 3
  } else if (form.time === 120) {
    form.timeAdjustVal2 = 4
  } else if (form.time === 360) {
    form.timeAdjustVal2 = 5
  } else if (form.time === 3000) {
    form.timeAdjustVal2 = 6
  }
}

const adjustTime = () => {
  form.time = form.timeAdjustVal
  form.timeModal = false
}

const adjustLeverage = () => {
  if (!isLogin.value) {
    ElNotification.error({
      title: $t('common.tip'),
      message: $t('common.logintip')
    })
    return
  }
  axios.post(`${host}${api.swap.leverage}`, {
    memberId: memberId.value,
    contractCoinId: currentCoin.id,
    leverage: form.leverageAdjustVal
  }).then(response => {
    const resp = response.data
    if (resp.code === 0) {
      form.leverage = form.leverageAdjustVal
      form.leverageModal = false
    } else {
      ElNotification.error({
        title: $t('swap.tip'),
        message: $t('swap.modifyLeverageWrong')
      })
      form.leverageModal = false
    }
  })
}

const getCNYRate = () => {
  silentPost(`${host}/market/exchange-rate/usd/cny`)
    .then(response => {
      if (!response?.data) return
      const resp = response.data
      CNYRate.value = resp.data
    })
}

const getCoin = (symbol) => {
  return coins._map[symbol]
}

const getSymbol = () => {
  silentGet(`${host}${api.swap.thumb}?type=${currentCoin.type}`)
    .then(response => {
      // 响应结构是 {code: 0, message: 'success', data: [...]}
      const apiResponse = response?.data?.data || response?.data
      if (!apiResponse || !Array.isArray(apiResponse)) return
      const resp = apiResponse
      // 先清空已有数据
      coins.USDT = []
      coins.USDT2 = []
      coins._map = {}
      for (let i = 0; i < resp.length; i++) {
        const coin = resp[i]
        coins[coin.base] = []
        coins[coin.base + '2'] = []
      }
      for (let i = 0; i < resp.length; i++) {
        const coin = resp[i]
        coin.price = coin.close = parseFloat(coin.close).toFixed(baseCoinScale.value)
        coin.rose = coin.chg > 0 ? '+' + (coin.chg * 100).toFixed(2) + '%' : (coin.chg * 100).toFixed(2) + '%'
        coin.coin = resp[i].symbol.split('/')[0]
        coin.base = resp[i].symbol.split('/')[1]
        coin.href = coin.id
        coin.isFavor = false
        if (!coins._map[coin.symbol]) {
          coins._map[coin.symbol] = [coin]
        } else {
          coins._map[coin.symbol].push(coin)
        }
        if (coin.zone === 0) {
          coins[coin.base].push(coin)
        } else {
          coins[coin.base + '2'].push(coin)
        }
        if (coin.id === currentCoin.id) {
          coin.perUsdt = currentCoin.perUsdt
          coin.type = currentCoin.type
          currentCoin.name = coin.name || coin.coin
          currentCoin.close = coin.close
          currentCoin.price = coin.price
          Object.assign(currentCoin, coin)
          form.limitPrice = parseFloat(coin.price)
        }
      }
      startWebsock()
      changeBaseCion(basecion.value)
    })
}

const getCoinInfo = () => {
  silentPost(`${host}${api.market.coinInfo}`, { unit: currentCoin.coin })
    .then(response => {
      if (!response?.data) return
      const resp = response.data
      if (resp) {
        coinInfo.name = resp.name
        coinInfo.information = resp.introduction
        coinInfo.infolink = resp.website
      }
    })
}

const getPlate = (str = '') => {
  const params = { symbol: currentCoin.symbol }
  silentPost(`${host}${api.swap.platemini}`, params)
    .then(response => {
      if (!response?.data) return
      plate.askRows = []
      plate.bidRows = []
      const resp = normalizeTradePlate(response.data)
      if (resp.ask && resp.ask.items) {
        for (let i = 0; i < resp.ask.items.length; i++) {
          if (i === 0) {
            resp.ask.items[i].totalAmount = resp.ask.items[i].amount
          } else {
            resp.ask.items[i].totalAmount = resp.ask.items[i - 1].totalAmount + resp.ask.items[i].amount
          }
        }
        if (resp.ask.items.length >= plate.maxPostion) {
          for (let i = plate.maxPostion; i > 0; i--) {
            const ask = resp.ask.items[i - 1]
            ask.direction = 'SELL'
            ask.position = i
            plate.askRows.push(ask)
          }
          const rows = plate.askRows
          const totle = rows[0].totalAmount
          plate.askTotle = totle
        } else {
          for (let i = plate.maxPostion; i > resp.ask.items.length; i--) {
            const ask = { price: 0, amount: 0 }
            ask.direction = 'SELL'
            ask.position = i
            ask.totalAmount = ask.amount
            plate.askRows.push(ask)
          }
          for (let i = resp.ask.items.length; i > 0; i--) {
            const ask = resp.ask.items[i - 1]
            ask.direction = 'SELL'
            ask.position = i
            plate.askRows.push(ask)
          }
          const askItemIndex = Math.max(resp.ask.items.length - 1, 0)
          const rows = plate.askRows
          const totle = rows[askItemIndex].totalAmount
          plate.askTotle = totle
        }
      }
      if (resp.bid && resp.bid.items) {
        for (let i = 0; i < resp.bid.items.length; i++) {
          if (i === 0) {
            resp.bid.items[i].totalAmount = resp.bid.items[i].amount
          } else {
            resp.bid.items[i].totalAmount = resp.bid.items[i - 1].totalAmount + resp.bid.items[i].amount
          }
        }
        for (let i = 0; i < resp.bid.items.length; i++) {
          const bid = resp.bid.items[i]
          bid.direction = 'BUY'
          bid.position = i + 1
          plate.bidRows.push(bid)
          if (i === plate.maxPostion - 1) break
        }
        if (resp.bid.items.length < plate.maxPostion) {
          for (let i = resp.bid.items.length; i < plate.maxPostion; i++) {
            const bid = { price: 0, amount: 0 }
            bid.direction = 'BUY'
            bid.position = i + 1
            bid.totalAmount = 0
            plate.bidRows.push(bid)
          }
          const bidItemIndex = Math.max(resp.bid.items.length - 1, 0)
          const rows = plate.bidRows
          const totle = rows[bidItemIndex].totalAmount
          plate.bidTotle = totle
        } else {
          const rows = plate.bidRows
          const totle = rows[rows.length - 1].totalAmount
          plate.bidTotle = totle
        }
      }
      if (str !== '') {
        selectedPlate.value = str
      }
    })
}

const getPlateFull = () => {
  const params = { symbol: currentCoin.symbol }
  silentPost(`${host}${api.swap.platefull}`, params)
    .then(response => {
      if (!response?.data) return
      const resp = normalizeTradePlate(response.data)
      resp.skin = skin.value
      if (depthGraphRef.value && depthGraphRef.value.draw) {
        depthGraphRef.value.draw(resp)
      }
    })
}

const getTrade = () => {
  const params = { symbol: currentCoin.symbol, size: 20 }
  silentPost(`${host}${api.swap.trade}`, params)
    .then(response => {
      if (!response?.data) return
      trade.rows = []
      const resp = unwrapList(response)
      for (let i = 0; i < resp.length; i++) {
        trade.rows.push(resp[i])
      }
    })
}

const startWebsock = () => {
  disconnectSocket()
  const socket = io(runtime.wshost || undefined)
  socketRef.value = socket
  socket.on('connect', () => {
    console.log('connect', socket.id)
  })
  socket.on('disconnect', () => {
    socket.io.on('reconnect', () => {
      console.log('reconnect', socket.id)
    })
  })

  datafeed.value = new Datafeeds.WebsockFeed(`${host}/swap`, currentCoin, socket)
  getKline()

  // 订阅价格变化
  socket.on('/topic/swap/thumb', (msg) => {
    const resp = JSON.parse(msg)
    const coinsList = getCoin(resp.symbol)
    if (coinsList && coinsList.length > 0) {
      coinsList.forEach(coin => {
        coin.price = resp.close
        coin.rose = resp.chg > 0 ? '+' + (resp.chg * 100).toFixed(2) + '%' : (resp.chg * 100).toFixed(2) + '%'
        coin.close = resp.close
        coin.high = resp.high
        coin.low = resp.low
        coin.turnover = parseInt(resp.volume)
        coin.volume = resp.volume
        coin.usdRate = resp.usdRate
      })
    }
  })

  // 订阅实时成交
  socket.on(`/topic/swap/trade/${currentCoin.symbol}`, (msg) => {
    const resp = JSON.parse(msg)
    if (resp.length > 0) {
      for (let i = 0; i < resp.length; i++) {
        trade.rows.unshift(resp[i])
      }
    }
    if (trade.rows.length > 30) {
      trade.rows = trade.rows.slice(0, 30)
    }
  })

  if (isLogin.value) {
    // 订阅委托取消
    socket.on(`/topic/swap/order-canceled/${currentCoin.symbol}/${memberId.value}`, (msg) => {
      refreshAccount()
    })

    // 订阅委托完成
    socket.on(`/topic/swap/order-completed/${currentCoin.symbol}/${memberId.value}`, (msg) => {
      refreshAccount()
    })

    // 订阅委托部分成交
    socket.on(`/topic/swap/order-trade/${currentCoin.symbol}/${memberId.value}`, (msg) => {
      refreshAccount()
    })

    socket.on(`/topic/swap/refresh/${currentCoin.symbol}/${memberId.value}`, (msg) => {
      refreshAccount()
    })
  }

  // 订阅盘口消息
  socket.on(`/topic/swap/trade-plate/${currentCoin.symbol}`, (msg) => {
    const resp = JSON.parse(msg)
    if (resp.direction === 'SELL') {
      const asks = resp.items
      plate.askRows = []
      let totle = 0
      for (let i = plate.maxPostion - 1; i >= 0; i--) {
        const ask = i < asks.length ? asks[i] : { price: 0, amount: 0 }
        ask.direction = 'SELL'
        ask.position = i + 1
        plate.askRows.push(ask)
      }
      for (let i = plate.askRows.length - 1; i >= 0; i--) {
        if (i === plate.askRows.length - 1 || plate.askRows[i].price === 0) {
          plate.askRows[i].totalAmount = plate.askRows[i].amount
        } else {
          plate.askRows[i].totalAmount = plate.askRows[i + 1].totalAmount + plate.askRows[i].amount
        }
        totle += plate.askRows[i].amount
      }
      plate.askTotle = totle
    } else {
      const bids = resp.items
      plate.bidRows = []
      let totle = 0
      for (let i = 0; i < plate.maxPostion; i++) {
        const bid = i < bids.length ? bids[i] : { price: 0, amount: 0 }
        bid.direction = 'BUY'
        bid.position = i + 1
        plate.bidRows.push(bid)
      }
      for (let i = 0; i < plate.bidRows.length; i++) {
        if (i === 0 || plate.bidRows[i].amount === 0) {
          plate.bidRows[i].totalAmount = plate.bidRows[i].amount
        } else {
          plate.bidRows[i].totalAmount = plate.bidRows[i - 1].totalAmount + plate.bidRows[i].amount
        }
        totle += plate.bidRows[i].amount
      }
      plate.bidTotle = totle
    }
    if (currentImgTable.value === 's') {
      getPlateFull()
    }
  })
}

// K 线图表初始化
let tvWidget = null
const isLocalAcceptanceHost = typeof window !== 'undefined' && ['127.0.0.1', 'localhost'].includes(window.location.hostname)
const chartLibraryScript = process.env.NODE_ENV === 'production'
  ? '/assets/charting_library/charting_library.min.js'
  : '/src/assets/js/charting_library/charting_library.min.js'
let tradingViewLoader = null

const destroyKline = () => {
  if (tvWidget) {
    tvWidget.remove()
    tvWidget = null
    window.tvWidget = null
  }

  const container = document.getElementById('swap_kline_container')
  if (container) {
    container.innerHTML = ''
  }
}

const disconnectSocket = () => {
  if (socketRef.value) {
    socketRef.value.removeAllListeners()
    socketRef.value.disconnect()
    socketRef.value = null
  }
}

const loadTradingView = () => {
  if (window.TradingView) {
    return Promise.resolve(window.TradingView)
  }

  if (tradingViewLoader) {
    return tradingViewLoader
  }

  tradingViewLoader = new Promise((resolve, reject) => {
    const existing = document.querySelector(`script[data-charting-library="${chartLibraryScript}"]`)
    if (existing) {
      existing.addEventListener('load', () => resolve(window.TradingView), { once: true })
      existing.addEventListener('error', reject, { once: true })
      return
    }

    const script = document.createElement('script')
    script.src = chartLibraryScript
    script.async = true
    script.dataset.chartingLibrary = chartLibraryScript
    script.onload = () => resolve(window.TradingView)
    script.onerror = reject
    document.head.appendChild(script)
  }).catch((error) => {
    tradingViewLoader = null
    throw error
  })

  return tradingViewLoader
}

const getKline = async () => {
  // 本地环境也显示真实 K 线图，用于测试
  // if (isLocalAcceptanceHost) {
  //   const container = document.getElementById('swap_kline_container')
  //   if (container) {
  //     container.innerHTML = '<div style="height:500px;display:flex;align-items:center;justify-content:center;color:#8f9ca5;background:#192330;">本地开发环境使用静态图表占位，合约行情联调已降级为页面级验收。</div>'
  //   }
  //   return
  // }

  const config = {
    autosize: true,
    height: 500,
    fullscreen: true,
    symbol: currentCoin.symbol,
    interval: '60',
    timezone: 'Asia/Shanghai',
    toolbar_bg: '#192330',
    container_id: 'swap_kline_container',
    datafeed: datafeed.value,
    library_path: process.env.NODE_ENV === 'production' ? '/assets/charting_library/' : '/src/assets/js/charting_library/',
    locale: 'zh',
    debug: false,
    drawings_access: { type: 'black', tools: [{ name: 'Regression Trend' }] },
    disabled_features: [
      'header_resolutions', 'timeframes_toolbar', 'header_symbol_search',
      'header_chart_type', 'header_compare', 'header_undo_redo',
      'header_screenshot', 'header_saveload', 'use_localstorage_for_settings',
      'left_toolbar', 'volume_force_overlay'
    ],
    enabled_features: ['hide_last_na_study_output', 'move_logo_to_main_pane'],
    custom_css_url: 'bundles/common.css',
    supported_resolutions: ['1', '5', '15', '30', '60', '1D', '1W', '1M'],
    charts_storage_url: 'http://saveload.tradingview.com',
    charts_storage_api_version: '1.1',
    client_id: 'tradingview.com',
    user_id: 'public_user_id',
    overrides: {
      'paneProperties.background': '#192330',
      'paneProperties.vertGridProperties.color': 'rgba(0,0,0,.1)',
      'paneProperties.horzGridProperties.color': 'rgba(0,0,0,.1)',
      'scalesProperties.textColor': '#61688A',
      'mainSeriesProperties.candleStyle.upColor': '#00b275',
      'mainSeriesProperties.candleStyle.downColor': '#f15057',
      'mainSeriesProperties.candleStyle.drawBorder': false,
      'mainSeriesProperties.candleStyle.wickUpColor': '#589065',
      'mainSeriesProperties.candleStyle.wickDownColor': '#AE4E54',
      'paneProperties.legendProperties.showLegend': false,
      'mainSeriesProperties.areaStyle.color1': 'rgba(71, 78, 112, 0.5)',
      'mainSeriesProperties.areaStyle.color2': 'rgba(71, 78, 112, 0.5)',
      'mainSeriesProperties.areaStyle.linecolor': '#9194a4',
      volumePaneSize: 'small'
    },
    time_frames: [
      { text: '1min', resolution: '1', description: 'realtime', title: '实时' },
      { text: '1min', resolution: '1', description: '1min' },
      { text: '5min', resolution: '5', description: '5min' },
      { text: '15min', resolution: '15', description: '15min' },
      { text: '30min', resolution: '30', description: '30min' },
      { text: '1hour', resolution: '60', description: '1hour' },
      { text: '4hour', resolution: '240', description: '4hour' },
      { text: '1day', resolution: '1D', description: '1day' },
      { text: '1week', resolution: '1W', description: '1week' },
      { text: '1mon', resolution: '1M', description: '1mon' }
    ]
  }

  if (skin.value === 'day') {
    config.toolbar_bg = '#fff'
    config.custom_css_url = 'bundles/common_day.css'
    config.overrides['paneProperties.background'] = '#fff'
    config.overrides['mainSeriesProperties.candleStyle.upColor'] = '#a6d3a5'
    config.overrides['mainSeriesProperties.candleStyle.downColor'] = '#ffa5a6'
  }

  try {
    const TradingView = await loadTradingView()
    destroyKline()
    tvWidget = window.tvWidget = new TradingView.widget(config)
    tvWidget.onChartReady(() => {
      tvWidget.chart().executeActionById('drawingToolbarAction')
      tvWidget.chart().createStudy('Moving Average', false, false, [5], null, { 'plot.color': '#EDEDED' })
      tvWidget.chart().createStudy('Moving Average', false, false, [10], null, { 'plot.color': '#ffe000' })
      tvWidget.chart().createStudy('Moving Average', false, false, [30], null, { 'plot.color': '#ce00ff' })
      tvWidget.chart().createStudy('Moving Average', false, false, [60], null, { 'plot.color': '#00adff' })
      createPeriodButtons(tvWidget)
    })
  } catch (error) {
    console.error('TradingView load failed:', error)
  }
}

const createPeriodButtons = (widget) => {
  // 分时按钮
  widget.createButton()
    .attr('title', 'realtime')
    .on('click', function() {
      if ($(this).hasClass('selected')) return
      $(this).addClass('selected').parent('.group').siblings('.group').find('.button.selected').removeClass('selected')
      widget.chart().setChartType(3)
      widget.setSymbol('', '1')
    })
    .append('<span>分时</span>')

  // M1
  widget.createButton()
    .attr('title', 'M1')
    .on('click', function() {
      if ($(this).hasClass('selected')) return
      $(this).addClass('selected').parent('.group').siblings('.group').find('.button.selected').removeClass('selected')
      widget.chart().setChartType(1)
      widget.setSymbol('', '1')
    })
    .append('<span>M1</span>')

  // M5
  widget.createButton()
    .attr('title', 'M5')
    .on('click', function() {
      if ($(this).hasClass('selected')) return
      $(this).addClass('selected').parent('.group').siblings('.group').find('.button.selected').removeClass('selected')
      widget.chart().setChartType(1)
      widget.setSymbol('', '5')
    })
    .append('<span>M5</span>')

  // M15
  widget.createButton()
    .attr('title', 'M15')
    .on('click', function() {
      if ($(this).hasClass('selected')) return
      $(this).addClass('selected').parent('.group').siblings('.group').find('.button.selected').removeClass('selected')
      widget.chart().setChartType(1)
      widget.setSymbol('', '15')
    })
    .append('<span>M15</span>')

  // M30
  widget.createButton()
    .attr('title', 'M30')
    .on('click', function() {
      if ($(this).hasClass('selected')) return
      $(this).addClass('selected').parent('.group').siblings('.group').find('.button.selected').removeClass('selected')
      widget.chart().setChartType(1)
      widget.setSymbol('', '30')
    })
    .append('<span>M30</span>')

  // H1 (默认选中)
  widget.createButton()
    .attr('title', 'H1')
    .on('click', function() {
      if ($(this).hasClass('selected')) return
      $(this).addClass('selected').parent('.group').siblings('.group').find('.button.selected').removeClass('selected')
      widget.chart().setChartType(1)
      widget.setSymbol('', '60')
    })
    .append('<span>H1</span>')
    .addClass('selected')

  // D1
  widget.createButton()
    .attr('title', 'D1')
    .on('click', function() {
      if ($(this).hasClass('selected')) return
      $(this).addClass('selected').parent('.group').siblings('.group').find('.button.selected').removeClass('selected')
      widget.chart().setChartType(1)
      widget.setSymbol('', '1D')
    })
    .append('<span>D1</span>')

  // W1
  widget.createButton()
    .attr('title', 'W1')
    .on('click', function() {
      if ($(this).hasClass('selected')) return
      $(this).addClass('selected').parent('.group').siblings('.group').find('.button.selected').removeClass('selected')
      widget.chart().setChartType(1)
      widget.setSymbol('', '1W')
    })
    .append('<span>W1</span>')
}

const init = () => {
  disconnectSocket()
  destroyKline()

  const id = route.params.pair
  if (id === undefined) {
    router.push('/swapindex/' + defaultPath.value)
  }

  silentGet(`${host}${api.swap.info}${id}`)
    .then(response => {
      const resp = response?.data?.data ?? response?.data
      if (!resp?.baseSymbol || !resp?.coinSymbol) return

      basecion.value = resp.baseSymbol.toLowerCase()
      currentCoin.symbol = resp.coinSymbol + '/' + resp.baseSymbol
      currentCoin.coin = resp.coinSymbol
      currentCoin.id = id
      currentCoin.type = resp.type
      currentCoin.base = resp.baseSymbol
      currentCoin.name = resp.coinSymbol

        if (currentCoin.type === 'ALWAYS') {
          store.commit('navigate', '/swapindex/1')
        } else {
          store.commit('navigate', '/swapindex/2')
        }
      store.commit('setSkin', skin.value)

      currentCoin.coinScale = resp.coinScale
      currentCoin.baseCoinScale = resp.baseCoinScale
      currentCoin.perUsdt = resp.shareNumber
      baseCoinScale.value = resp.baseCoinScale
      coinScale.value = resp.coinScale
      takerFee.value = resp.takerFee
      makerFee.value = resp.makerFee
      enableMarketBuy.value = resp.enableMarketBuy
      enableMarketSell.value = resp.enableMarketSell
      exchangeable.value = resp.exchangeable

      getCNYRate()
      getCoinInfo()
      getSymbol()
      getPlate()
      getPlateFull()
      getTrade()

      if (isLogin.value) {
        getWallet()
        getCurrentPosition()
        getCurrentOrder()
        getHistoryOrder()
        getLeverage()
      }

      form.entrustType = 1
    })
}

// 监听语言变化
watch(() => store.state.lang, () => {
  // 更新语言数据
})

watch(() => route.params.pair, (pair, previousPair) => {
  if (!pair || pair === previousPair) {
    return
  }
  init()
})

// 监听当前币种价格变化
watch(() => currentCoin.close, (p) => {
  let profit = 0.0
  currentPosition.rows.forEach(e => {
    if (e.direction === 'BUY') {
      e.profit = parseFloat((1 / e.openPrice - 1 / p) * e.multiple * e.principalAmount * p).toFixed(baseCoinScale.value)
    } else {
      e.profit = parseFloat((1 / p - 1 / e.openPrice) * e.multiple * e.principalAmount * p).toFixed(baseCoinScale.value)
    }
    profit += parseFloat(e.profit)
  })
  wallet.profit = parseFloat(profit).toFixed(baseCoinScale.value)
})

// 生命周期
onMounted(() => {
  isLogin.value = store.getters.isLogin
  memberId.value = store.getters.member?.id || 0
  lang.value = store.state.lang || 'zh'
  init()
})

onBeforeUnmount(() => {
  disconnectSocket()
  destroyKline()
})
</script>

<style scoped lang="scss">
@import '../../assets/css/swap.css';

.container.swap {
  --swap-surface: #16202b;
  --swap-surface-strong: #192330;
  --swap-surface-soft: #223041;
  --swap-border: rgba(115, 138, 166, 0.16);
  --swap-text-main: #f4f7fb;
  --swap-text-muted: #8a98ad;
  --swap-accent: #f0a70a;
  --swap-shadow: 0 18px 48px rgba(3, 10, 18, 0.32);
  color: #fff;
  background:
    radial-gradient(circle at top left, rgba(240, 167, 10, 0.08), transparent 24%),
    radial-gradient(circle at top right, rgba(0, 178, 117, 0.08), transparent 22%),
    linear-gradient(180deg, #0b1520 0%, #0d1722 38%, #101b27 100%);
  padding: 72px 16px 24px;

  .main {
    width: min(100%, 1680px);
    margin: 0 auto;
    display: grid;
    grid-template-columns: 320px minmax(0, 1fr) 340px;
    grid-template-rows: auto 1fr;
    grid-template-areas:
      'left chart orderbook'
      'orders orders orders';
    gap: 14px;
    align-items: start;

    .trade-workspace {
      display: contents;
    }

    /* 左侧列：币种列表 + 最新交易 */
    .left-column {
      grid-area: left;
      display: flex;
      flex-direction: column;
      gap: 14px;
      min-width: 0;
    }

    /* 左侧：币种列表 */
    .market-rail {
      flex: 0 0 auto;
      min-width: 0;
    }

    /* 左侧中部：最新交易 */
    .latest-trade-panel {
      flex: 1 1 auto;
      min-width: 0;
      display: flex;
      flex-direction: column;
    }

    .latest-trade-panel .trade-wrap {
      flex: 1;
      overflow: hidden;
    }

    /* 中间：K 线图、账户信息和交易表单 */
    .center {
      grid-area: chart;
      min-width: 0;
      display: flex;
      flex-direction: column;
      gap: 14px;

      /* K 线图和深度图 */
      .imgtable {
        flex: 0 0 518px;
        position: relative;
        overflow: hidden;

        .handler {
          position: absolute;
          top: 18px;
          right: 22px;
          z-index: 1000;
          display: flex;
          gap: 8px;

          > span {
            background-color: rgba(44, 59, 89, 0.72);
            color: #c7cce6;
            padding: 6px 10px;
            cursor: pointer;
            font-size: 13px;
            opacity: 0.72;
            border: 1px solid transparent;
            border-radius: 999px;
            transition: all 0.2s ease;

            &.active {
              opacity: 1;
              border-color: rgba(240, 167, 10, 0.35);
              background: rgba(240, 167, 10, 0.16);
              color: #fff2ca;
            }
          }
        }
      }
    }

    /* 右侧：盘口数据 */
    .orderbook-area {
      grid-area: orderbook;
      min-width: 0;
    }
  }

  .symbol {
    display: flex;
    justify-content: space-between;
    padding: 18px 22px;
    margin-bottom: 0;
    align-items: center;
    background-color: transparent;
    line-height: 1;
    gap: 18px;
    flex-wrap: wrap;

    .item {
      min-width: 0;

      .price-cny {
        font-size: 12px;
        color: var(--swap-text-muted);
      }

      .coin {
        font-size: 26px;
        font-weight: 700;
        letter-spacing: 0.02em;
      }

      .text {
        width: 100%;
        display: block;
        font-size: 11px;
        color: var(--swap-text-muted);
        margin-bottom: 8px;
        text-transform: uppercase;
        letter-spacing: 0.08em;
      }

      .num {
        font-size: 16px;
        color: #fff;
        font-weight: 600;
      }

      > img {
        display: block;
        width: 18px;
        height: 18px;
        cursor: pointer;
      }
    }

    .symbol-primary {
      min-width: 180px;
    }
  }

  .order {
    min-height: 227px;
    margin-bottom: 0;
    overflow: hidden;

    .order-handler {
      background-color: transparent;
      font-size: 0;
      border-bottom: 1px solid rgba(255, 255, 255, 0.06);
      display: flex;

      > span {
        flex: 1;
        font-size: 14px;
        display: flex;
        align-items: center;
        justify-content: center;
        color: var(--swap-text-muted);
        cursor: pointer;
        line-height: 42px;
        background-color: transparent;
        transition: color 0.2s ease, border-color 0.2s ease;

        &.active {
          color: #f7fafc;
          background-color: transparent;
          border-bottom: 2px solid var(--swap-accent);
        }
      }
    }
  }

  .panel-shell {
    background: linear-gradient(180deg, rgba(24, 32, 43, 0.96), rgba(20, 28, 39, 0.96));
    border: 1px solid var(--swap-border);
    border-radius: 18px;
    box-shadow: var(--swap-shadow);
    backdrop-filter: blur(10px);
  }

  .panel-header {
    display: flex;
    align-items: center;
    justify-content: space-between;
    gap: 12px;
    padding: 16px 18px;
    border-bottom: 1px solid rgba(255, 255, 255, 0.05);

    strong {
      display: block;
      font-size: 14px;
      color: var(--swap-text-main);
      font-weight: 600;
    }
  }

  .panel-header-compact {
    padding: 14px 16px;
  }

  .panel-toolbar {
    padding: 12px 14px;
  }

  .panel-kicker {
    display: inline-block;
    margin-bottom: 4px;
    font-size: 10px;
    letter-spacing: 0.12em;
    text-transform: uppercase;
    color: var(--swap-accent);
  }

  .panel-meta {
    font-size: 12px;
    color: var(--swap-text-muted);
  }

  .market-search {
    padding-bottom: 8px;
  }

  .chart-shell {
    padding-top: 0;
  }

  .workspace-orders {
    grid-area: orders;
    min-width: 0;
  }

  .ledger-panel {
    min-height: 352px;
  }

  .orderbook-area {
    overflow: hidden;
  }

  .trade-form-panel {
    min-height: 352px;
  }

  .latest-trade-panel {
    grid-area: trade;
    min-height: 352px;
  }

  .account-panel-wrapper {
    min-height: 120px;
    flex-shrink: 0;
  }

  .trade-form-panel-wrapper {
    min-height: 352px;
    flex-shrink: 0;
  }

  /* 中间容器内的账户和表单自动在 chart 区域内垂直排列 */
  .center {
    display: flex;
    flex-direction: column;
    gap: 5px;
  }

  .trade-form-header {
    border-bottom: 1px solid rgba(255, 255, 255, 0.06);
  }

  .trade-form-body {
    padding: 16px;
    min-height: 420px; /* 固定高度，防止切换时抖动 */

    .form-row-with-label {
      display: flex;
      align-items: center;
      justify-content: flex-start;
      gap: 12px;
      margin-bottom: 16px;

      .form-label {
        color: var(--swap-text-muted);
        font-size: 13px;
        white-space: nowrap;
        min-width: 70px;
        text-align: left;
      }

      .el-radio-group {
        display: flex;
        gap: 12px;

        .el-radio {
          margin-right: 0;
          color: #b0b8db;
          font-size: 13px;
        }
      }
    }

    .form-row {
      display: flex;
      align-items: center;
      justify-content: space-between;
      margin-bottom: 16px;

      .form-label {
        color: var(--swap-text-muted);
        font-size: 13px;
        min-width: 70px;
        text-align: left;
      }

      .leverage-btn {
        flex: 1;
        max-width: 200px;
        border: none;
        background: rgb(240 172 25);
        color: #000;
        padding: 6px 10px;
        font-size: 13px;
        height: 32px;
      }
    }

    .trade-form {
      width: 100%;

      .trade-form-item {
        margin-bottom: 16px;
        display: flex;
        align-items: center;
        min-height: 40px;

        .form-item-label {
          color: #b0b8db;
          font-size: 13px;
          min-width: 70px;
          text-align: left;
          white-space: nowrap;
        }

        .form-input-wrap {
          flex: 1;
          display: flex;
          align-items: center;
          gap: 8px;

          .market-toggle-btn {
            flex-shrink: 0;
            padding: 4px 10px;
            min-width: auto;
            font-size: 12px;
          }

          .el-input-number {
            flex: 1;
            max-width: 180px;
          }

          .el-input {
            flex: 1;
            max-width: 180px;
          }
        }
      }
    }

    /* 交易操作区域 */
    .trade-actions {
      display: flex;
      justify-content: space-between;
      gap: 12px;
      margin-top: 20px;

      .trade-action-item {
        flex: 1;
        display: flex;
        flex-direction: column;
        align-items: center;
        padding: 12px;
        background: rgba(255, 255, 255, 0.03);
        border-radius: 8px;
        gap: 8px;

        .trade-action-info {
          display: flex;
          align-items: center;
          gap: 6px;
          font-size: 12px;

          .num {
            font-weight: 600;
          }
        }

        .operate_btn {
          margin-top: 0;
          width: 100%;
          padding: 8px 16px;
          font-size: 13px;
        }
      }

      .buy-action {
        border: 1px solid rgba(65, 179, 113, 0.3);
      }

      .sell-action {
        border: 1px solid rgba(215, 78, 90, 0.3);
      }
    }

    /* 未登录提示 */
    .operate-login {
      margin-top: 20px;

      .login-tip {
        display: flex;
        align-items: center;
        justify-content: center;
        padding: 12px 16px;
        background: rgba(255, 255, 255, 0.03);
        border: 1px solid #232d3a;
        border-radius: 5px;
        font-size: 14px;
      }
    }
  }

  .account-panel {
    min-height: 120px;
    color: var(--swap-text-muted);
    order: 1;
  }

  .account-panel-header {
    .linkmore {
      margin-right: 0;
      line-height: 1.4;
    }
  }

  .latest-trade-panel {
    min-height: 352px;
    order: 2;
  }

  .latest-trade-table {
    overflow: hidden;
  }

  /* 左侧列高度对齐 - 以 .trade-form-panel 为基准 */
  .left-column {
    .market-rail,
    .latest-trade-panel {
      min-height: 494px; /* 与 .trade-form-panel 总高度一致 */
    }

    .coin-menu.panel-shell,
    .trade-wrap.panel-shell {
      min-height: 494px;
    }
  }

  /* 调整币种列表表格高度 */
  .coin-menu .el-table {
    height: 420px !important;
  }

  /* 调整最新交易表格高度 */
  .trade-wrap .el-table {
    height: 420px !important;
  }

  .market-overview {
    padding-bottom: 8px;
  }

  .summary-market-table {
    padding: 0 12px 12px;
  }

  .market-list-cell,
  .market-list-price {
    display: flex;
    flex-direction: column;
    gap: 4px;
    padding: 4px 0;
  }

  .market-list-cell {
    align-items: flex-start;
  }

  .market-list-price {
    align-items: flex-end;
  }

  .market-list-name {
    font-size: 14px;
    color: #f2f5fa;
  }

  .market-list-volume {
    font-size: 12px;
    color: var(--swap-text-muted);
  }

  .chart-shell {
    overflow: hidden;
  }
}

// 白天模式
.container.swap.day {
  color: #333;
  background-color: #fff;

  .main {
    .left {
      background-color: #fff;
      box-shadow: 0 0 2px #ccc;

      .handlers {
        background-color: #fff;
      }

      .plate-nowprice {
        background-color: #fff;
        border-top: 1px solid #f0f0f0;
        border-bottom: 1px solid #f0f0f0;
      }
    }

    .imgtable {
      border-radius: 6px;
      box-shadow: 0 0 2px #ccc;

      .handler {
        > span {
          border: 1px solid #333;
        }
      }
    }
  }

  .symbol {
    background-color: #fff;
    box-shadow: 0 0 2px #ccc;

    .item {
      .text {
        color: #999;
      }

      .num {
        color: #333;
      }
    }
  }

  .order {
    box-shadow: 0 2px 2px #ccc;

    .order-handler {
      margin-right: -2px;
      background-color: #fff;

      > span {
        color: #333;
        background-color: #fafafa;
        box-shadow: 0 0 2px #ccc;

        &.active {
          color: #f0a70a;
          background-color: #fff;
        }
      }
    }
  }
}

.container.swap {
  :deep(.el-table) {
    --el-table-border-color: transparent;
    --el-table-header-bg-color: transparent;
    --el-table-row-hover-bg-color: rgba(255, 255, 255, 0.03);
    --el-table-tr-bg-color: transparent;
    --el-table-bg-color: transparent;
    --el-table-text-color: #d5dceb;
    --el-table-header-text-color: #7f8ea3;
  }

  :deep(.el-table th.el-table__cell) {
    height: 42px;
    font-size: 12px;
    font-weight: 500;
    letter-spacing: 0.04em;
    border-bottom: 1px solid rgba(255, 255, 255, 0.05);
  }

  :deep(.el-table td.el-table__cell) {
    height: 42px;
    border-bottom: 1px solid rgba(255, 255, 255, 0.04);
  }

  :deep(.el-input__wrapper),
  :deep(.el-input-number__decrease),
  :deep(.el-input-number__increase),
  :deep(.el-input-number .el-input__wrapper) {
    background: rgba(8, 14, 22, 0.86);
    box-shadow: inset 0 0 0 1px rgba(115, 138, 166, 0.18);
    color: #fff;
  }

  :deep(.el-radio-button__inner) {
    background: rgba(10, 16, 24, 0.88);
    border-color: rgba(115, 138, 166, 0.18);
    color: var(--swap-text-muted);
  }

  :deep(.el-radio-button__original-radio:checked + .el-radio-button__inner) {
    background: rgba(240, 167, 10, 0.16);
    border-color: rgba(240, 167, 10, 0.45);
    color: #fff3d0;
    box-shadow: none;
  }
}

#swap_kline_container {
  background: transparent;
}

@media (max-width: 1440px) {
  .container.swap {
    .main {
      grid-template-columns: 280px minmax(0, 1fr) 300px;
    }
  }
}

@media (max-width: 1200px) {
  .container.swap {
    .main {
      width: 100%;
      grid-template-areas:
        'market chart'
        'trade chart'
        'orders orders';
      grid-template-columns: 280px minmax(0, 1fr);
    }

    .orderbook-area {
      display: none;
    }
  }
}

@media (max-width: 768px) {
  .container.swap {
    padding: 68px 10px 20px;

    .main,
    .main {
      gap: 10px;
      grid-template-columns: minmax(0, 1fr);
      grid-template-areas:
        'market'
        'chart'
        'trade'
        'orders';
    }

    .symbol {
      padding: 16px;
    }

    .main .center .imgtable {
      height: 420px;
    }

    .panel-header,
    .panel-toolbar {
      padding-left: 12px;
      padding-right: 12px;
    }
  }
}

.coin-info {
  color: #8f9ca5;
  display: -webkit-box;
  -webkit-box-orient: vertical;
  -webkit-line-clamp: 5;
  overflow: hidden;
  padding-top: 15px;
}

// 账户项
.account-item {
  width: 90%;
  margin: 0 auto;

  > div:nth-child(1) {
    float: left;
    line-height: 38px;
    font-size: 12px;
    color: #b0b8db;
  }

  > div:nth-child(2) {
    float: right;
    line-height: 38px;
    color: #fff;
  }
}

// 委托类型按钮
.order-type-btn-wrap {
  height: 30px;
  padding: 0 16px;
  width: 100%;
  background: #192330;

  button {
    width: auto;
    border-radius: 5px;
    height: auto;
    background: #192330;
    border: 1px solid #fff;
    padding: 3px 15px;
    text-align: center;
    font-size: 12px;
    margin-right: 10px;
    color: #fff;
  }

  .btn-selected {
    color: #000 !important;
    background: #f0ac19;
    border: 1px solid #f0ac19 !important;
  }
}

// 交易按钮颜色
.trade-actions .buy-action .operate_btn {
  background: #41b371;
}

.trade-actions .sell-action .operate_btn {
  background: #d74e5a;
}

.swap .order {
  margin-bottom: 5px !important;
}

// 按钮样式
.bg-green {
  background: #41b371 !important;
  border-color: #41b371 !important;
}

.bg-red {
  background: #d74e5a !important;
  border-color: #d74e5a !important;
}

// 百分比背景条
.percentdiv {
  position: absolute;
  left: 0;
  top: 0;
  height: 100%;
  opacity: 0.3;
}

.percenttd {
  position: relative;
}

// 买卖颜色
.buy {
  color: #00b275;
}

.sell {
  color: #f15057;
}

// 隐藏类
.hidden {
  display: none;
}
</style>
