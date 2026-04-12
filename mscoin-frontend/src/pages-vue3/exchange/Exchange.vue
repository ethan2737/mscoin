<template>
  <div class="container exchange" :class="skin">
    <!-- 主交易区 -->
    <div class="main">
      <!-- 右侧：币种列表 -->
      <div class="right">
        <div class="coin-menu">
          <div style="padding: 8px 10px;height:48px;">
            <el-input
              v-model="searchKey"
              :placeholder="$t('common.searchplaceholder')"
              @input="seachInputChange"
              clearable
            />
          </div>
          <div class="sc_filter">
            <span v-show="isLogin" @click="changeBaseCion('favor')" :class="{active:basecion==='favor'}">{{$t('service.CUSTOM')}}</span>
            <span @click="changeBaseCion('usdt')" :class="{active:basecion==='usdt'}">USDT</span>
            <span @click="changeBaseCion('btc')" :class="{active:basecion==='btc'}">BTC</span>
            <span @click="changeBaseCion('eth')" :class="{active:basecion==='eth'}">ETH</span>
          </div>
          <el-table
            v-show="basecion==='usdt'"
            :data="dataIndex"
            highlight-row
            @current-change="gohref"
          >
            <el-table-column prop="coin" :label="$t('exchange.coin')" width="120">
              <template #default="{ row }">
                <div style="display: flex; align-items: center;">
                  <el-icon
                    style="margin-right: 5px; cursor: pointer;"
                    :class="{ 'favor-icon-active': row.isFavor }"
                    @click.stop="toggleFavor(row)"
                  >
                    <StarFilled v-if="row.isFavor" />
                    <Star v-else />
                  </el-icon>
                  <span>{{ row.symbol }}</span>
                  <span v-if="row.coin === 'BZB'" style="font-size: 8px; padding: 2px 5px; color: #FFF; background: #F30; border-radius: 4px; margin-left: 5px;">HOT</span>
                </div>
              </template>
            </el-table-column>
            <el-table-column prop="close" :label="$t('exchange.lastprice')" sortable />
            <el-table-column prop="rose" :label="$t('exchange.daychange')" sortable>
              <template #default="{ row }">
                <span :class="parseFloat(row.rose) < 0 ? 'sell' : 'buy'">{{ row.rose }}</span>
              </template>
            </el-table-column>
          </el-table>

          <el-table
            v-show="basecion==='btc'"
            :data="dataIndex"
            highlight-row
            @current-change="gohref"
          >
            <el-table-column prop="coin" :label="$t('exchange.coin')" width="120">
              <template #default="{ row }">
                <div style="display: flex; align-items: center;">
                  <el-icon
                    style="margin-right: 5px; cursor: pointer;"
                    :class="{ 'favor-icon-active': row.isFavor }"
                    @click.stop="toggleFavor(row)"
                  >
                    <StarFilled v-if="row.isFavor" />
                    <Star v-else />
                  </el-icon>
                  <span>{{ row.symbol }}</span>
                </div>
              </template>
            </el-table-column>
            <el-table-column prop="close" :label="$t('exchange.lastprice')" sortable />
            <el-table-column prop="rose" :label="$t('exchange.daychange')" sortable>
              <template #default="{ row }">
                <span :class="parseFloat(row.rose) < 0 ? 'sell' : 'buy'">{{ row.rose }}</span>
              </template>
            </el-table-column>
          </el-table>

          <el-table
            v-show="basecion==='eth'"
            :data="dataIndex"
            highlight-row
            @current-change="gohref"
          >
            <el-table-column prop="coin" :label="$t('exchange.coin')" width="120">
              <template #default="{ row }">
                <div style="display: flex; align-items: center;">
                  <el-icon
                    style="margin-right: 5px; cursor: pointer;"
                    :class="{ 'favor-icon-active': row.isFavor }"
                    @click.stop="toggleFavor(row)"
                  >
                    <StarFilled v-if="row.isFavor" />
                    <Star v-else />
                  </el-icon>
                  <span>{{ row.symbol }}</span>
                </div>
              </template>
            </el-table-column>
            <el-table-column prop="close" :label="$t('exchange.lastprice')" sortable />
            <el-table-column prop="rose" :label="$t('exchange.daychange')" sortable>
              <template #default="{ row }">
                <span :class="parseFloat(row.rose) < 0 ? 'sell' : 'buy'">{{ row.rose }}</span>
              </template>
            </el-table-column>
          </el-table>

          <el-table
            v-show="basecion==='favor'"
            :data="dataIndex"
            highlight-row
            @current-change="gohref"
            :no-data-text="$t('common.nodata')"
          >
            <el-table-column prop="coin" :label="$t('exchange.coin')" width="120" />
            <el-table-column prop="close" :label="$t('exchange.lastprice')" sortable />
            <el-table-column prop="rose" :label="$t('exchange.daychange')" sortable>
              <template #default="{ row }">
                <span :class="parseFloat(row.rose) < 0 ? 'sell' : 'buy'">{{ row.rose }}</span>
              </template>
            </el-table-column>
          </el-table>
        </div>
      </div>

      <!-- 中间：K 线图和交易面板 -->
      <div class="center">
        <!-- 币种信息 -->
        <div class="symbol">
          <div class="item" @click="currentCoinFavorChange">
            <el-icon :size="24" :class="{ 'favor-icon-active': currentCoinIsFavor }">
              <StarFilled v-if="currentCoinIsFavor" />
              <Star v-else />
            </el-icon>
          </div>
          <div class="item" style="margin-left: -40px;">
            <span class="coin">{{ currentCoin.coin }}<small>/{{ currentCoin.base }}</small></span>
            <el-popover placement="bottom-start" :width="300" trigger="hover">
              <template #reference>
                <el-icon style="color:#546886; margin-left:5px;"><InfoFilled /></el-icon>
              </template>
              <div class="coin-info">{{ coinInfo.information }}</div>
              <p style="text-align:right; margin-top: 10px;">
                <a :href="coinInfo.infolink" target="_blank">{{ $t("exchange.moredetail") }}</a>
              </p>
            </el-popover>
          </div>
          <div class="item">
            <span class="text">{{ $t('service.NewPrice') }}</span>
            <span class="num" :class="{buy: currentCoin.change > 0, sell: currentCoin.change < 0}">
              {{ currentCoin.close?.toFixed(baseCoinScale) }}
            </span>
            <span class="price-cny">≈ ￥{{ (currentCoin.usdRate * CNYRate)?.toFixed(2) }}</span>
          </div>
          <div class="item">
            <span class="text">{{ $t('service.Change') }}</span>
            <span class="num" :class="{buy: currentCoin.change > 0, sell: currentCoin.change < 0}">{{ currentCoin.rose }}</span>
          </div>
          <div class="item">
            <span class="text">{{ $t('service.high') }}</span>
            <span class="num">{{ currentCoin.high?.toFixed(baseCoinScale) }}</span>
          </div>
          <div class="item">
            <span class="text">{{ $t('service.low') }}</span>
            <span class="num">{{ currentCoin.low?.toFixed(baseCoinScale) }}</span>
          </div>
          <div class="item">
            <span class="text">{{ $t('service.ExchangeNum') }}</span>
            <span class="num">{{ currentCoin.volume }} {{ currentCoin.coin }}</span>
          </div>
        </div>

        <!-- K 线图和深度图 -->
        <div class="imgtable">
          <div class="chart-panel">
            <div id="kline_container" :class="{hidden: currentImgTable === 's'}"></div>
            <DepthGraph
              ref="depthGraphRef"
              :values="depthPlate"
              :class="{hidden: currentImgTable === 'k'}"
            />
          </div>
          <div class="handler">
            <span @click="changeImgTable('k')" :class="{active: currentImgTable === 'k'}">{{ $t("exchange.kline") }}</span>
            <span @click="changeImgTable('s')" :class="{active: currentImgTable === 's'}">{{ $t("exchange.depth") }}</span>
          </div>
        </div>

        <!-- 交易面板 -->
        <div class="trade_wrap">
          <div class="trade_panel trade_panel_logout">
            <!-- 未登录遮罩 -->
            <div class="mask" v-show="!isLogin">
              <span>{{ $t("common.please") }}
                <router-link to="/login"><span style="color:#f0a70a;">{{ $t("common.login") }}</span></router-link> /
                <router-link to="/register"><span style="color:#00dcff;">{{ $t("common.register") }}</span></router-link>
              </span>
            </div>

            <!-- 发行遮罩 -->
            <div class="publish-mask" v-show="isLogin && showPublish && memberId !== 2">
              <div style="width: 100%; margin-top:17%; text-align:center; letter-spacing:3px;">
                <span v-if="publishState === 0">{{ $t("exchange.publishstate0") }}</span>
                <span v-if="publishState === 1">{{ $t("exchange.publishstate1") }}</span>
                <span v-if="publishState === 2">{{ $t("exchange.publishstate2") }}</span>
                <span v-if="publishState === 3">{{ $t("exchange.publishstate3") }}</span>
              </div>
            </div>

            <!-- 交易菜单 -->
            <div class="trade_menu">
              <span @click="limited_price" :class="{active: !showMarket}">{{ $t("exchange.limited_price") }}</span>
              <span @click="market_price" :class="{active: showMarket}">{{ $t("exchange.market_price") }}</span>
            </div>

            <!-- 交易表单 - 紧凑型布局 -->
            <div class="trade_bd">
              <!-- 限价交易区 -->
              <div class="trade-grid" v-show="!showMarket">
                <!-- 限价买入 -->
                <div class="trade-grid-item trade-grid-buy">
                  <div v-if="isLogin" class="hd hd_login hd-compact">
                    <span class="balance-unit">{{ currentCoin.base }}</span>
                    <span class="balance-value">{{ wallet.base?.toFixed(baseCoinScale) }}</span>
                    <span class="balance-label">{{ $t("exchange.canuse") }}</span>
                    <router-link :to="rechargeUSDTUrl">{{ $t("exchange.recharge") }}</router-link>
                  </div>
                  <el-form :model="form.buy" class="compact-form">
                    <!-- 标题和输入框合并到一行 -->
                    <div class="form-row-title">
                      <span class="form-title buy-title">{{ $t("exchange.buyin") }} {{ currentCoin.coin }}</span>
                    </div>
                    <div class="form-row-inline">
                      <label class="inline-label">{{ $t('exchange.buyprice') }}</label>
                      <el-input v-model="form.buy.limitPrice" :placeholder="$t('exchange.buyprice')" @input="keyEvent" size="small" />
                      <span class="inline-unit">{{ currentCoin.base }}</span>
                    </div>
                    <div class="form-row-inline">
                      <label class="inline-label">{{ $t('exchange.buynum') }}</label>
                      <el-input v-model="form.buy.limitAmount" :placeholder="$t('exchange.buynum')" @input="keyEvent" size="small" />
                      <span class="inline-unit">{{ currentCoin.coin }}</span>
                    </div>
                    <div class="form-row-total">
                      <span class="total-label">{{ $t("exchange.amount") }}:</span>
                      <span class="total-value">{{ form.buy.limitTurnover?.toFixed(baseCoinScale) }} {{ currentCoin.base }}</span>
                    </div>
                    <el-button
                      v-if="exchangeable === 1"
                      type="success"
                      @click="buyWithLimitPrice"
                      :loading="buying"
                      class="btn-compact"
                    >
                      {{ $t("exchange.buyin") }}
                    </el-button>
                  </el-form>
                </div>

                <!-- 限价卖出 -->
                <div class="trade-grid-item trade-grid-sell">
                  <div v-if="isLogin" class="hd hd_login hd-compact">
                    <span class="balance-label">{{ $t("exchange.canuse") }}</span>
                    <span class="balance-value">{{ wallet.coin?.toFixed(coinScale) }}</span>
                    <span class="balance-unit">{{ currentCoin.coin }}</span>
                    <router-link :to="rechargeCoinUrl">{{ $t("exchange.recharge") }}</router-link>
                  </div>
                  <el-form :model="form.sell" class="compact-form">
                    <div class="form-row-title">
                      <span class="form-title sell-title">{{ $t("exchange.sellout") }} {{ currentCoin.coin }}</span>
                    </div>
                    <div class="form-row-inline">
                      <label class="inline-label">{{ $t('exchange.sellprice') }}</label>
                      <el-input v-model="form.sell.limitPrice" :placeholder="$t('exchange.sellprice')" @input="keyEvent" size="small" />
                      <span class="inline-unit">{{ currentCoin.base }}</span>
                    </div>
                    <div class="form-row-inline">
                      <label class="inline-label">{{ $t('exchange.sellnum') }}</label>
                      <el-input v-model="form.sell.limitAmount" :placeholder="$t('exchange.sellnum')" @input="keyEvent" size="small" />
                      <span class="inline-unit">{{ currentCoin.coin }}</span>
                    </div>
                    <div class="form-row-total">
                      <span class="total-label">{{ $t("exchange.amount") }}:</span>
                      <span class="total-value">{{ form.sell.limitTurnover?.toFixed(coinScale) }} {{ currentCoin.base }}</span>
                    </div>
                    <el-button
                      v-if="exchangeable === 1"
                      type="danger"
                      @click="sellLimitPrice"
                      :loading="selling"
                      class="btn-compact"
                    >
                      {{ $t("exchange.sellout") }}
                    </el-button>
                  </el-form>
                </div>
              </div>

              <!-- 市价交易区 -->
              <div class="trade-grid" v-show="showMarket">
                <!-- 市价买入 -->
                <div class="trade-grid-item trade-grid-buy">
                  <div v-if="isLogin" class="hd hd_login hd-compact">
                    <span class="balance-unit">{{ currentCoin.base }}</span>
                    <span class="balance-value">{{ wallet.base?.toFixed(baseCoinScale) }}</span>
                    <span class="balance-label">{{ $t("exchange.canuse") }}</span>
                    <router-link :to="rechargeUSDTUrl">{{ $t("exchange.recharge") }}</router-link>
                  </div>
                  <el-form :model="form.buy" class="compact-form">
                    <div class="form-row-title">
                      <span class="form-title buy-title">{{ $t("exchange.buyin") }} {{ currentCoin.coin }}</span>
                    </div>
                    <div class="form-row-inline">
                      <label class="inline-label">{{ $t('exchange.buyprice') }}</label>
                      <el-input disabled :placeholder="$t('exchange.buytip')" size="small" />
                      <span class="inline-unit">{{ currentCoin.base }}</span>
                    </div>
                    <div class="form-row-inline">
                      <label class="inline-label">{{ $t('exchange.amount') }}</label>
                      <el-input v-model="form.buy.marketAmount" :placeholder="$t('exchange.amount')" @input="keyEvent" size="small" />
                      <span class="inline-unit">{{ currentCoin.base }}</span>
                    </div>
                    <el-button
                      v-if="enableMarketBuy === 1 && exchangeable === 1"
                      type="success"
                      @click="buyWithMarketPrice"
                      :loading="buying"
                      class="btn-compact"
                    >
                      {{ $t("exchange.buyin") }}
                    </el-button>
                  </el-form>
                </div>

                <!-- 市价卖出 -->
                <div class="trade-grid-item trade-grid-sell">
                  <div v-if="isLogin" class="hd hd_login hd-compact">
                    <span class="balance-label">{{ $t("exchange.canuse") }}</span>
                    <span class="balance-value">{{ wallet.coin?.toFixed(coinScale) }}</span>
                    <span class="balance-unit">{{ currentCoin.coin }}</span>
                    <router-link :to="rechargeCoinUrl">{{ $t("exchange.recharge") }}</router-link>
                  </div>
                  <el-form :model="form.sell" class="compact-form">
                    <div class="form-row-title">
                      <span class="form-title sell-title">{{ $t("exchange.sellout") }} {{ currentCoin.coin }}</span>
                    </div>
                    <div class="form-row-inline">
                      <label class="inline-label">{{ $t('exchange.sellprice') }}</label>
                      <el-input disabled :placeholder="$t('exchange.selltip')" size="small" />
                      <span class="inline-unit">{{ currentCoin.base }}</span>
                    </div>
                    <div class="form-row-inline">
                      <label class="inline-label">{{ $t('exchange.num') }}</label>
                      <el-input v-model="form.sell.marketAmount" :placeholder="$t('exchange.sellnum')" @input="keyEvent" size="small" />
                      <span class="inline-unit">{{ currentCoin.coin }}</span>
                    </div>
                    <el-button
                      v-if="enableMarketSell === 1 && exchangeable === 1"
                      type="danger"
                      @click="sellMarketPrice"
                      :loading="selling"
                      class="btn-compact"
                    >
                      {{ $t("exchange.sellout") }}
                    </el-button>
                  </el-form>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>

      <!-- 左侧：买卖盘和成交 -->
      <div class="left plate-wrap" style="position:relative; flex: 1; display: flex; flex-direction: column;">
        <!-- 倒计时面板 -->
        <div class="lightning-panel" v-if="showCountDown" :style="{backgroundColor: countDownBgColor}">
          <img v-if="lang === '简体中文' && publishType === 'FENTAN'" src="../../assets/images/lightning-bg.png" alt="" />
          <img v-if="lang === 'English' && publishType === 'FENTAN'" src="../../assets/images/lightning-bg-en.png" alt="" />
          <img v-if="lang === '简体中文' && publishType === 'QIANGGOU'" src="../../assets/images/lightning-bg2.png" alt="" />
          <img v-if="lang === 'English' && publishType === 'QIANGGOU'" src="../../assets/images/lightning-bg2-en.png" alt="" />
          <div class="l-content">
            <BZCountDown
              style="width:100%; margin-top:5px;"
              :countDownBgColor="countDownBgColor"
              :publishState="publishState"
              :publishType="publishType"
              :currentTime="currentTime"
              :startTime="startTime"
              :clearTime="clearTime"
              :endTime="endTime"
              :showPublishMask="showPublishMask"
              :hidePublishMask="hidePublishMask"
              :hideCountDown="hideCountDown"
              @update:countDownBgColor="val => countDownBgColor = val"
              @update:publishState="val => publishState = val"
            />
            <p class="l-info">
              <span class="l-txt">{{ $t("exchange.publishamount") }}：</span>
              <span class="l-count">{{ publishAmount }}</span>
              <span class="l-unit">{{ currentCoin.coin }}</span>
              &nbsp;&nbsp;&nbsp;&nbsp;
              <span class="l-txt">{{ $t("exchange.publishprice") }}：</span>
              <span class="l-price">{{ publishPrice }}</span>
              <span class="l-unit">{{ currentCoin.base }}</span>
            </p>
            <p class="l-detail">
              <router-link target="_blank" to="/announcement/118930">{{ $t("exchange.publishdetail") }}</router-link>
            </p>
          </div>
        </div>

        <!-- 盘口选择器 -->
        <div class="plate-handlers">
          <span @click="changePlate('all')" :class="{active: selectedPlate === 'all'}">{{ $t('exchange.all') }}</span>
          <span @click="changePlate('buy')" :class="{active: selectedPlate === 'buy'}">{{ $t('exchange.buyplate') }}</span>
          <span @click="changePlate('sell')" :class="{active: selectedPlate === 'sell'}">{{ $t('exchange.sellplate') }}</span>
        </div>

        <div class="plate-book">
          <!-- 卖盘 -->
          <div v-show="selectedPlate !== 'buy'" class="plate-table plate-table-sell">
            <el-table
              :data="plate.askRows"
              highlight-current-row
              @row-click="buyPlate"
              :no-data-text="$t('common.nodata')"
            >
              <el-table-column prop="price" :label="$t('exchange.price')">
                <template #default="{ row }">
                  <span :class="row.direction?.toLowerCase()">
                    {{ row.isPlaceholder ? row.displayPrice : row.price?.toFixed(baseCoinScale) }}
                  </span>
                </template>
              </el-table-column>
              <el-table-column prop="amount" :label="$t('exchange.num')">
                <template #default="{ row }">
                  {{ row.isPlaceholder ? row.displayAmount : row.amount }}
                </template>
              </el-table-column>
              <el-table-column prop="totalAmount" :label="$t('exchange.total')">
                <template #default="{ row }">
                  {{ row.isPlaceholder ? row.displayTotalAmount : row.totalAmount }}
                </template>
              </el-table-column>
            </el-table>
          </div>

          <!-- 当前价格 -->
          <div class="plate-nowprice">
            <div class="plate-nowprice__labels">
              <span class="sell">{{ $t('exchange.sellplate') }}</span>
              <span class="plate-nowprice__label">{{ $t('exchange.lastprice') }}</span>
              <span class="buy">{{ $t('exchange.buyplate') }}</span>
            </div>
            <div class="plate-nowprice__value">
              <span class="price" :class="{buy: currentCoin.change > 0, sell: currentCoin.change < 0}">
                {{ currentCoin.price?.toFixed(baseCoinScale) }}
              </span>
              <span v-if="currentCoin.change > 0" class="buy">↑</span>
              <span v-else-if="currentCoin.change < 0" class="sell">↓</span>
              <span class="price-cny"> ≈ {{ (currentCoin.usdRate * CNYRate)?.toFixed(2) }} CNY</span>
            </div>
          </div>

          <!-- 买盘 -->
          <div v-show="selectedPlate !== 'sell'" class="plate-table plate-table-buy">
            <el-table
              :data="plate.bidRows"
              highlight-current-row
              @row-click="sellPlate"
              :no-data-text="$t('common.nodata')"
            >
              <el-table-column prop="price" :label="$t('exchange.price')">
                <template #default="{ row }">
                  <span :class="row.direction?.toLowerCase()">
                    {{ row.isPlaceholder ? row.displayPrice : row.price?.toFixed(baseCoinScale) }}
                  </span>
                </template>
              </el-table-column>
              <el-table-column prop="amount" :label="$t('exchange.num')">
                <template #default="{ row }">
                  {{ row.isPlaceholder ? row.displayAmount : row.amount }}
                </template>
              </el-table-column>
              <el-table-column prop="totalAmount" :label="$t('exchange.total')">
                <template #default="{ row }">
                  {{ row.isPlaceholder ? row.displayTotalAmount : row.totalAmount }}
                </template>
              </el-table-column>
            </el-table>
          </div>
        </div>

        <!-- 最新成交 -->
        <div class="trade-wrap" v-show="!showCountDown">
          <div class="trade-wrap__header">最新成交</div>
          <el-table :data="trade.rows" :no-data-text="$t('common.nodata')">
            <el-table-column prop="amount" :label="$t('exchange.num')">
              <template #default="{ row }">
                {{ row.amount?.toFixed(coinScale) }}
              </template>
            </el-table-column>
            <el-table-column prop="price" :label="$t('exchange.price')">
              <template #default="{ row }">
                <span :class="row.direction === 'BUY' ? 'buy' : 'sell'">
                  {{ row.price?.toFixed(baseCoinScale) }}
                </span>
              </template>
            </el-table-column>
            <el-table-column prop="time" :label="$t('exchange.time')">
              <template #default="{ row }">
                {{ timeFormat(row.time) }}
              </template>
            </el-table-column>
          </el-table>
        </div>
      </div>
    </div>

    <!-- 委托记录 -->
    <div class="order">
      <div class="order-handler">
        <span @click="changeOrder('current')" :class="{active: selectedOrder === 'current'}">{{ $t('exchange.curdelegation') }}</span>
        <span @click="changeOrder('history')" :class="{active: selectedOrder === 'history'}">{{ $t('exchange.hisdelegation') }}</span>
        <router-link v-show="selectedOrder === 'current'" class="linkmore" to="/uc/entrust/current">{{ $t("common.more") }}</router-link>
        <router-link v-show="selectedOrder === 'history'" class="linkmore" to="/uc/entrust/history">{{ $t("common.more") }}</router-link>
      </div>
      <div class="table">
        <el-table
          v-if="selectedOrder === 'current'"
          :data="currentOrder.rows"
          :no-data-text="$t('common.nodata')"
        >
          <el-table-column type="expand">
            <template #default="{ row }">
              <expandRow :skin="skin" :rows="row.detail" />
            </template>
          </el-table-column>
          <el-table-column prop="time" :label="$t('exchange.time')" />
          <el-table-column prop="symbol" :label="$t('exchange.symbol')" />
          <el-table-column prop="type" :label="$t('exchange.type')" />
          <el-table-column prop="direction" :label="$t('exchange.direction')" />
          <el-table-column prop="price" :label="$t('exchange.price')">
            <template #default="{ row }">
              <span>{{ formatOrderPrice(row) }}</span>
            </template>
          </el-table-column>
          <el-table-column prop="amount" :label="$t('exchange.num')" />
          <el-table-column prop="tradedAmount" :label="$t('exchange.traded')" />
          <el-table-column prop="turnover" :label="$t('exchange.dealamount')" />
          <el-table-column :label="$t('exchange.action')" width="110">
            <template #default="{ row, $index }">
              <el-button size="small" @click="cancel($index)" style="border-color: #f0ac19; color: #f1ac19;">
                {{ $t("exchange.undo") }}
              </el-button>
            </template>
          </el-table-column>
        </el-table>

        <el-table
          v-else
          :data="historyOrder.rows"
          :no-data-text="$t('common.nodata')"
        >
          <el-table-column type="expand">
            <template #default="{ row }">
              <expandRow :skin="skin" :rows="row.detail" />
            </template>
          </el-table-column>
          <el-table-column prop="time" :label="$t('exchange.time')" />
          <el-table-column prop="symbol" :label="$t('exchange.symbol')" />
          <el-table-column prop="type" :label="$t('exchange.type')" />
          <el-table-column prop="direction" :label="$t('exchange.direction')" />
          <el-table-column prop="price" :label="$t('exchange.price')">
            <template #default="{ row }">
              <span>{{ formatOrderPrice(row) }}</span>
            </template>
          </el-table-column>
          <el-table-column prop="amount" :label="$t('exchange.num')" />
          <el-table-column prop="tradedAmount" :label="$t('exchange.done')" />
          <el-table-column prop="turnover" :label="$t('exchange.dealamount')" />
          <el-table-column prop="status" :label="$t('exchange.status')">
            <template #default="{ row }">
              <span :class="statusClass(row.status)" v-if="row.status === 'COMPLETED'">{{ $t("exchange.finished") }}</span>
              <span :class="statusClass(row.status)" v-else-if="row.status === 'CANCELED'">{{ $t("exchange.canceled") }}</span>
              <span v-else>--</span>
            </template>
          </el-table-column>
        </el-table>
      </div>
    </div>
  </div>
</template>

<script setup>
/**
 * Vue 3 迁移 - 交易所主页面
 *
 * 迁移点:
 * 1. 使用 <script setup> 语法
 * 2. 使用 Composition API 替代 Options API
 * 3. Element Plus 组件替代 iView 组件
 * 4. 使用 inject('store') 和 inject('router') 兼容 Vuex 3.x 和 Vue Router 3.x
 * 5. 使用 provide/inject 传递依赖
 *
 * 子组件:
 * - DepthGraph: 深度图组件
 * - BZCountDown: 倒计时组件
 * - expandRow: 委托展开组件
 */

import { ref, reactive, computed, inject, onMounted, onBeforeUnmount, watch, nextTick } from 'vue'
import { ElMessage, ElNotification, ElMessageBox } from 'element-plus'
import { Star, StarFilled, InfoFilled } from '@element-plus/icons-vue'
import axios from 'axios'
import moment from 'moment'
import io from 'socket.io-client'
import { getAuthHeaders, useRuntimeContract } from '../../config/runtime-vue3'

// 导入子组件
import DepthGraph from './DepthGraph.vue'
import BZCountDown from './BZCountDown.vue'
import expandRow from './expand.vue'

// 导入 TradingView datafeed
import Datafeeds from '../../assets/js/charting_library/datafeed/bitrade.js'
import { shouldUseAreaChartForSymbol } from './chart-preferences'
import { applyFavorState, getFavorSuccessMessage } from './favor'
import { pickWalletBalances } from './wallet'
import {
  PLATE_DISPLAY_ROWS,
  buildBuyPlateRows,
  buildSellPlateRows,
  getPlateRowPrice
} from './plate-utils'

// Vuex 3.x 和 Vue Router 3.x 兼容
const store = inject('store')
const router = inject('router')

// 常量
const runtime = useRuntimeContract()
const host = runtime.host
const api = runtime.api
const isProd = import.meta.env.PROD
const chartLibraryScript = isProd
  ? '/assets/charting_library/charting_library.min.js'
  : '/src/assets/js/charting_library/charting_library.min.js'
const chartLibraryPath = isProd
  ? '/assets/charting_library/'
  : '/src/assets/js/charting_library/'

// 响应式数据
const skin = ref('night')
const currentImgTable = ref('k')
const selectedOrder = ref('current')
const selectedPlate = ref('all')
const CNYRate = ref(null)
const datafeed = ref(null)
const defaultPath = ref('btc_usdt')
const basecion = ref('usdt')
const coinScale = ref(6)
const baseCoinScale = ref(6)
const symbolFee = ref(0.001)
const memberId = ref(0)
const dataIndex = ref([])
const dataIndex2 = ref([])
const searchKey = ref('')
const coinInfo = reactive({})
const isLogin = ref(false)
const lang = ref('zh')
const member = ref(null)

// 当前币种
const currentCoin = reactive({
  base: '',
  coin: '',
  symbol: '',
  close: 0,
  price: 0,
  rose: '',
  high: 0,
  low: 0,
  volume: 0,
  turnover: 0,
  usdRate: 0,
  change: 0,
  isFavor: false
})

// 交易控制
const enableMarketBuy = ref(1)
const enableMarketSell = ref(1)
const exchangeable = ref(1)

// 发行相关
const publishType = ref('NONE')
const currentTime = ref(0)
const publishAmount = ref(0)
const publishPrice = ref(0)
const startTime = ref('2000-01-01 00:00:00')
const endTime = ref('2000-01-01 00:00:00')
const clearTime = ref('2000-01-01 00:00:00')
const showPublish = ref(false)
const showCountDown = ref(false)
const countDownBgColor = ref('#003478')
const publishState = ref(0)

// Slider 控制
const sliderStep = [25, 50, 75, 100]
const sliderBuyLimitPercent = ref(0)
const sliderSellLimitPercent = ref(0)
const sliderBuyMarketPercent = ref(0)
const sliderSellMarketPercent = ref(0)
const sliderBuyDisabled = computed(() => {
  const min = 1 / Math.pow(10, baseCoinScale.value)
  return wallet.base < min
})
const sliderSellDisabled = computed(() => {
  const min = 1 / Math.pow(10, coinScale.value)
  return wallet.coin < min
})

// 钱包
const wallet = reactive({
  base: 0,
  coin: 0,
  kick: 0
})

// 市价/限价
const showMarket = ref(false)

// 表单数据
const form = reactive({
  buy: {
    limitPrice: 0,
    limitAmount: 0,
    marketAmount: 0,
    limitTurnover: 0
  },
  sell: {
    limitPrice: 0,
    limitAmount: 0,
    marketAmount: 0,
    limitTurnover: 0
  }
})

// 交易数据
const trade = reactive({
  rows: [],
  columns: []
})

// 买卖盘数据
const plate = reactive({
  maxPostion: PLATE_DISPLAY_ROWS,
  columns: [],
  askRows: [],
  bidRows: [],
  askTotle: 0,
  bidTotle: 0
})

const depthPlate = reactive({
  bid: { items: [], highestPrice: 0, lowestPrice: 0 },
  ask: { items: [], highestPrice: 0, lowestPrice: 0 },
  skin: 'night',
  priceUnit: '',
  amountUnit: '',
  priceScale: 4,
  amountScale: 4,
  symbol: ''
})

// 当前委托
const currentOrder = reactive({
  columns: [],
  rows: []
})

// 历史委托
const historyOrder = reactive({
  pageSize: 10,
  total: 10,
  page: 0,
  columns: [],
  rows: []
})

// 币种列表
const coins = reactive({
  _map: [],
  USDT: [],
  BTC: [],
  ETH: [],
  USDT2: [],
  BTC2: [],
  ETH2: [],
  favor: [],
  columns: []
})

// 收藏状态
const collecRequesting = ref(false)
const currentCoinIsFavor = ref(false)

// 交易状态
const buying = ref(false)
const selling = ref(false)

// 深度图引用
const depthGraphRef = ref(null)
const socketRef = ref(null)
let tvWidget = null
let tradingViewLoader = null
let depthPlateRequestSeq = 0

// 计算属性
const rechargeUSDTUrl = computed(() => `/uc/recharge?name=${currentCoin.base}`)
const rechargeCoinUrl = computed(() => `/uc/recharge?name=${currentCoin.coin}`)
const preferAreaChart = computed(() => shouldUseAreaChartForSymbol({
  base: currentCoin.base,
  baseCoinScale: baseCoinScale.value,
  close: currentCoin.close
}))
const buildAuthConfig = (config = {}) => ({
  ...config,
  headers: {
    ...getAuthHeaders(),
    ...(config.headers || {})
  }
})
const post = (url, payload = {}, config = {}) => axios.post(url, payload, buildAuthConfig(config))
const silentPost = (url, payload = {}, config = {}) => post(url, payload, config).catch(() => null)
const unwrapPayload = (response) => {
  const payload = response?.data
  if (!payload) {
    return null
  }
  if (typeof payload === 'object' && payload !== null && Object.prototype.hasOwnProperty.call(payload, 'code')) {
    return payload.code === 0 ? payload.data : null
  }
  return payload
}

// 方法实现
const tipFormat = (val) => `${val}%`

const cloneDepthItems = (items) => (
  Array.isArray(items)
    ? items
      .filter(item => item && item.price !== undefined && item.amount !== undefined)
      .map(item => ({
        price: Number(item.price) || 0,
        amount: Number(item.amount) || 0
      }))
    : []
)

const resetDepthPlate = (symbol = currentCoin.symbol) => {
  depthPlate.bid = { items: [], highestPrice: 0, lowestPrice: 0 }
  depthPlate.ask = { items: [], highestPrice: 0, lowestPrice: 0 }
  depthPlate.skin = skin.value
  depthPlate.priceUnit = currentCoin.base || ''
  depthPlate.amountUnit = currentCoin.coin || ''
  depthPlate.priceScale = baseCoinScale.value
  depthPlate.amountScale = coinScale.value
  depthPlate.symbol = symbol || ''
}

const syncDepthPlate = (payload, symbol = currentCoin.symbol) => {
  depthPlate.bid = {
    highestPrice: Number(payload?.bid?.highestPrice) || 0,
    lowestPrice: Number(payload?.bid?.lowestPrice) || 0,
    items: cloneDepthItems(payload?.bid?.items)
  }
  depthPlate.ask = {
    highestPrice: Number(payload?.ask?.highestPrice) || 0,
    lowestPrice: Number(payload?.ask?.lowestPrice) || 0,
    items: cloneDepthItems(payload?.ask?.items)
  }
  depthPlate.skin = skin.value
  depthPlate.priceUnit = currentCoin.base || ''
  depthPlate.amountUnit = currentCoin.coin || ''
  depthPlate.priceScale = baseCoinScale.value
  depthPlate.amountScale = coinScale.value
  depthPlate.symbol = symbol || currentCoin.symbol
}

const keyEvent = () => {
  // 键盘事件处理
}

const seachInputChange = () => {
  searchKey.value = searchKey.value.toUpperCase()
  filterCoinList()
}

const filterCoinList = () => {
  let sourceList = []
  if (basecion.value === 'favor') {
    sourceList = coins.favor
  } else if (basecion.value === 'usdt') {
    sourceList = coins.USDT
  } else if (basecion.value === 'btc') {
    sourceList = coins.BTC
  } else if (basecion.value === 'eth') {
    sourceList = coins.ETH
  }

  dataIndex.value = sourceList.filter(item => item.coin.indexOf(searchKey.value) === 0)
}

const changeBaseCion = (str) => {
  basecion.value = str
  if (str === 'usdt') {
    dataIndex.value = coins.USDT
    dataIndex2.value = coins.USDT2
  } else if (str === 'btc') {
    dataIndex.value = coins.BTC
    dataIndex2.value = coins.BTC2
  } else if (str === 'eth') {
    dataIndex.value = coins.ETH
    dataIndex2.value = coins.ETH2
  } else if (str === 'favor') {
    dataIndex.value = coins.favor
  }
}

const changePlate = (str) => {
  plate.maxPostion = PLATE_DISPLAY_ROWS
  getPlate(str)
}

const changeImgTable = (str) => {
  currentImgTable.value = str
  if (str === 's') {
    getPlateFull()
    return
  }
  if (!tvWidget) {
    initKline()
  }
}

const changeOrder = (str) => {
  selectedOrder.value = str
}

const limited_price = () => {
  showMarket.value = false
}

const market_price = () => {
  showMarket.value = true
}

const gohref = (currentRow) => {
  router.push({
    name: 'ExchangePair',
    params: {
      pair: currentRow.href
    }
  })
}

const toggleFavor = (row) => {
  if (!isLogin.value) {
    ElMessage.warning('请先登录')
    return
  }
  if (collecRequesting.value) return

  collecRequesting.value = true
  if (row.isFavor) {
    cancelCollect(null, row)
  } else {
    collect(null, row)
  }
}

const currentCoinFavorChange = () => {
  if (!isLogin.value) {
    ElMessage.warning('请先登录')
    return
  }
  if (collecRequesting.value) return

  collecRequesting.value = true
  const nextFavorState = !currentCoinIsFavor.value
  const requestUrl = nextFavorState ? api.exchange.favorAdd : api.exchange.favorDelete

  post(host + requestUrl, { symbol: currentCoin.symbol })
    .then(response => {
      const resp = response.data
      if (resp.code === 0) {
        const updated = applyFavorState({
          coins,
          symbol: currentCoin.symbol,
          isFavor: nextFavorState,
          currentCoinSymbol: currentCoin.symbol
        })

        currentCoin.isFavor = nextFavorState
        currentCoinIsFavor.value = updated.currentCoinIsFavor ?? currentCoinIsFavor.value
        ElMessage.success(getFavorSuccessMessage(nextFavorState))
      } else {
        ElMessage.error(resp.message || '自选操作失败')
      }
    })
    .catch(() => {
      ElMessage.error('自选操作失败')
    })
    .finally(() => {
      collecRequesting.value = false
    })
}

const collect = (index, row) => {
  if (!isLogin.value) {
    ElMessage.info('请先登录')
    return
  }
  post(host + api.exchange.favorAdd, { symbol: row.symbol })
    .then(response => {
      const resp = response.data
      if (resp.code === 0) {
        const updated = applyFavorState({
          coins,
          row,
          symbol: row.symbol,
          isFavor: true,
          currentCoinSymbol: currentCoin.symbol
        })

        if (currentCoin.symbol === row.symbol) {
          currentCoin.isFavor = true
          currentCoinIsFavor.value = updated.currentCoinIsFavor ?? true
        }
        ElMessage.success(getFavorSuccessMessage(true))
      } else {
        ElMessage.error(resp.message || '自选操作失败')
      }
    })
    .catch(() => {
      ElMessage.error('自选操作失败')
    })
    .finally(() => {
      collecRequesting.value = false
    })
}

const cancelCollect = (index, row) => {
  if (!isLogin.value) {
    ElMessage.info('请先登录')
    return
  }
  post(host + api.exchange.favorDelete, { symbol: row.symbol })
    .then(response => {
      const resp = response.data
      if (resp.code === 0) {
        const updated = applyFavorState({
          coins,
          row,
          symbol: row.symbol,
          isFavor: false,
          currentCoinSymbol: currentCoin.symbol
        })

        if (currentCoin.symbol === row.symbol) {
          currentCoin.isFavor = false
          currentCoinIsFavor.value = updated.currentCoinIsFavor ?? false
        }
        ElMessage.success(getFavorSuccessMessage(false))
      } else {
        ElMessage.error(resp.message || '自选操作失败')
      }
    })
    .catch(() => {
      ElMessage.error('自选操作失败')
    })
    .finally(() => {
      collecRequesting.value = false
    })
}

const getCoin = (symbol) => {
  return coins._map[symbol]
}

const getLocale = () => (lang.value === 'English' ? 'en' : 'zh')

const destroyKline = () => {
  if (tvWidget) {
    tvWidget.remove()
    tvWidget = null
    window.tvWidget = null
  }
  const container = document.getElementById('kline_container')
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

const createPeriodButtons = (widget) => {
  const periods = [
    { title: 'realtime', text: '分时', resolution: '1', chartType: 3 },
    { title: 'M1', text: 'M1', resolution: '1', chartType: preferAreaChart.value ? 3 : 1 },
    { title: 'M5', text: 'M5', resolution: '5', chartType: preferAreaChart.value ? 3 : 1 },
    { title: 'M15', text: 'M15', resolution: '15', chartType: preferAreaChart.value ? 3 : 1 },
    { title: 'M30', text: 'M30', resolution: '30', chartType: preferAreaChart.value ? 3 : 1 },
    { title: 'H1', text: 'H1', resolution: '60', chartType: preferAreaChart.value ? 3 : 1 },
    { title: 'D1', text: 'D1', resolution: '1D', chartType: preferAreaChart.value ? 3 : 1 },
    { title: 'W1', text: 'W1', resolution: '1W', chartType: preferAreaChart.value ? 3 : 1 }
  ]

  periods.forEach((period) => {
    widget.createButton()
      .attr('title', period.title)
      .on('click', function() {
        if ($(this).hasClass('selected')) return
        $(this)
          .addClass('selected')
          .parent('.group')
          .siblings('.group')
          .find('.button.selected')
          .removeClass('selected')
        widget.chart().setChartType(period.chartType)
        widget.setSymbol(currentCoin.symbol, period.resolution)
      })
      .append(`<span>${period.text}</span>`)
  })
}

const initKline = async () => {
  if (!datafeed.value) {
    return
  }
  destroyKline()
  try {
    await loadTradingView()
  } catch (error) {
    ElMessage.error('K线图加载失败')
    return
  }

  const config = {
    autosize: true,
    height: 350,
    fullscreen: false,
    symbol: currentCoin.symbol,
    interval: '60',
    timezone: 'Asia/Shanghai',
    toolbar_bg: '#192330',
    container_id: 'kline_container',
    datafeed: datafeed.value,
    library_path: chartLibraryPath,
    locale: getLocale(),
    debug: false,
    drawings_access: {
      type: 'black',
      tools: [{ name: 'Regression Trend' }]
    },
    disabled_features: [
      'header_resolutions',
      'timeframes_toolbar',
      'header_symbol_search',
      'header_chart_type',
      'header_compare',
      'header_undo_redo',
      'header_screenshot',
      'header_saveload',
      'use_localstorage_for_settings',
      'left_toolbar',
      'volume_force_overlay'
    ],
    enabled_features: [
      'hide_last_na_study_output',
      'move_logo_to_main_pane'
    ],
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
    }
  }

  if (skin.value === 'day') {
    config.toolbar_bg = '#fff'
    config.custom_css_url = 'bundles/common_day.css'
    config.overrides['paneProperties.background'] = '#fff'
    config.overrides['mainSeriesProperties.candleStyle.upColor'] = '#a6d3a5'
    config.overrides['mainSeriesProperties.candleStyle.downColor'] = '#ffa5a6'
  }

  await nextTick()
  tvWidget = window.tvWidget = new window.TradingView.widget(config)
  tvWidget.onChartReady(() => {
    if (preferAreaChart.value) {
      tvWidget.chart().setChartType(3)
    }
    tvWidget.chart().executeActionById('drawingToolbarAction')
    tvWidget.chart().createStudy('Moving Average', false, false, [5], null, { 'plot.color': '#EDEDED' })
    tvWidget.chart().createStudy('Moving Average', false, false, [10], null, { 'plot.color': '#ffe000' })
    tvWidget.chart().createStudy('Moving Average', false, false, [30], null, { 'plot.color': '#ce00ff' })
    tvWidget.chart().createStudy('Moving Average', false, false, [60], null, { 'plot.color': '#00adff' })
    createPeriodButtons(tvWidget)
  })
}

const getSymbol = () => {
  return silentPost(host + api.market.thumb, {})
    .then(response => {
      const resp = unwrapPayload(response)
      if (!Array.isArray(resp)) return
      // 清空已有数据
      coins.USDT = []
      coins.BTC = []
      coins.ETH = []
      coins.USDT2 = []
      coins.BTC2 = []
      coins.ETH2 = []
      coins._map = []
      coins.favor = []

      for (let i = 0; i < resp.length; i++) {
        const coin = resp[i]
        coin.price = coin.close
        coin.rose = coin.chg > 0 ? '+' + (coin.chg * 100).toFixed(2) + '%' : (coin.chg * 100).toFixed(2) + '%'
        coin.coin = coin.symbol.split('/')[0]
        coin.base = coin.symbol.split('/')[1]
        coin.href = (coin.coin + '_' + coin.base).toLowerCase()
        coin.isFavor = false
        coins._map[coin.symbol] = coin

        if (coin.zone === 0) {
          coins[coin.base].push(coin)
        } else {
          coins[coin.base + '2'].push(coin)
        }

        if (coin.symbol === currentCoin.symbol) {
          currentCoin.close = coin.close
          currentCoin.price = coin.price
          currentCoin.rose = coin.rose
          currentCoin.change = coin.change
          currentCoin.high = coin.high
          currentCoin.low = coin.low
          currentCoin.volume = coin.volume
          currentCoin.turnover = coin.turnover
          currentCoin.usdRate = coin.usdRate
          form.buy.limitPrice = form.sell.limitPrice = coin.price
        }
      }

      if (isLogin.value) {
        getFavor()
      }

      startWebsock()
      initKline()
      changeBaseCion(basecion.value)
    })
}

const getFavor = () => {
  post(host + api.exchange.favorFind, {})
    .then(response => {
      coins.favor = []
      currentCoinIsFavor.value = false
      const resp = unwrapPayload(response) || []
      for (let i = 0; i < resp.length; i++) {
        const coin = getCoin(resp[i].symbol)
        if (coin) {
          coin.isFavor = true
          coins.favor.push(coin)
        }
        if (currentCoin.symbol === resp[i].symbol) {
          currentCoinIsFavor.value = true
        }
      }
    })
}

const getPlate = (str = '') => {
  silentPost(host + api.market.platemini, { symbol: currentCoin.symbol })
    .then(response => {
      const resp = unwrapPayload(response)
      if (!resp) return
      plate.askRows = buildSellPlateRows(resp.ask?.items, plate.maxPostion)
      plate.bidRows = buildBuyPlateRows(resp.bid?.items, plate.maxPostion)
      plate.askTotle = (resp.ask?.items || []).reduce((total, item) => total + (Number(item?.amount) || 0), 0)
      plate.bidTotle = (resp.bid?.items || []).reduce((total, item) => total + (Number(item?.amount) || 0), 0)

      if (str !== '') {
        selectedPlate.value = str
      }
    })
}

const getPlateFull = () => {
  const requestSymbol = currentCoin.symbol
  const requestSeq = ++depthPlateRequestSeq

  silentPost(host + api.market.platefull, { symbol: requestSymbol })
    .then(response => {
      const resp = unwrapPayload(response)
      if (!resp) return
      if (requestSeq !== depthPlateRequestSeq || requestSymbol !== currentCoin.symbol) return
      syncDepthPlate(resp, requestSymbol)
    })
}

const getTrade = () => {
  silentPost(host + api.market.trade, { symbol: currentCoin.symbol, size: 20 })
    .then(response => {
      const resp = unwrapPayload(response)
      if (!Array.isArray(resp)) return
      trade.rows = []
      for (let i = 0; i < resp.length; i++) {
        trade.rows.push(resp[i])
      }
    })
}

const startWebsock = () => {
  disconnectSocket()
  const socket = io(runtime.wshost || undefined)
  socketRef.value = socket
  datafeed.value = new Datafeeds.WebsockFeed(
    host + '/market',
    {
      ...currentCoin,
      baseCoinScale: baseCoinScale.value,
      coinScale: coinScale.value
    },
    socket,
    baseCoinScale.value
  )

  socket.on('connect', () => {
    console.log('connect', socket.id)
  })

  socket.on('disconnect', () => {
    socket.io.on('reconnect', () => {
      console.log('reconnect', socket.id)
    })
  })

  // 订阅价格变化
  socket.on('/topic/market/thumb', (msg) => {
    const resp = JSON.parse(msg)
    const coin = getCoin(resp.symbol)
    if (coin) {
      coin.price = resp.close
      coin.rose = resp.chg > 0 ? '+' + (resp.chg * 100).toFixed(2) + '%' : (resp.chg * 100).toFixed(2) + '%'
      coin.close = resp.close
      coin.high = resp.high
      coin.low = resp.low
      coin.turnover = parseInt(resp.volume)
      coin.volume = resp.volume
      coin.usdRate = resp.usdRate
    }
    if (resp.symbol === currentCoin.symbol) {
      currentCoin.price = resp.close
      currentCoin.close = resp.close
      currentCoin.rose = resp.chg > 0 ? `+${(resp.chg * 100).toFixed(2)}%` : `${(resp.chg * 100).toFixed(2)}%`
      currentCoin.change = resp.change || 0
      currentCoin.high = resp.high
      currentCoin.low = resp.low
      currentCoin.volume = resp.volume
      currentCoin.turnover = resp.turnover || parseInt(resp.volume)
      currentCoin.usdRate = resp.usdRate
    }
  })

  // 订阅实时成交
  socket.on(`/topic/market/trade/${currentCoin.symbol}`, (msg) => {
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

  // 订阅盘口消息
  socket.on(`/topic/market/trade-plate/${currentCoin.symbol}`, (msg) => {
    const resp = JSON.parse(msg)
    if (resp.direction === 'SELL') {
      plate.askRows = buildSellPlateRows(resp.items, plate.maxPostion)
      plate.askTotle = (resp.items || []).reduce((total, item) => total + (Number(item?.amount) || 0), 0)
    } else {
      plate.bidRows = buildBuyPlateRows(resp.items, plate.maxPostion)
      plate.bidTotle = (resp.items || []).reduce((total, item) => total + (Number(item?.amount) || 0), 0)
    }
    if (currentImgTable.value === 's') {
      getPlateFull()
    }
  })
}

const buyPlate = (currentRow) => {
  const price = getPlateRowPrice(currentRow)
  if (price !== null) {
    form.buy.limitPrice = price
  }
}

const sellPlate = (currentRow) => {
  const price = getPlateRowPrice(currentRow)
  if (price !== null) {
    form.sell.limitPrice = price
  }
}

const timeFormat = (time) => {
  return moment(time).format('HH:mm:ss')
}

const formatOrderPrice = (row) => {
  if (row?.type === 'MARKET_PRICE') {
    return '市价'
  }
  if (row?.price === null || row?.price === undefined || row?.price === '') {
    return '--'
  }
  return row.price
}

const statusClass = (status) => {
  if (status === 'COMPLETED') {
    return 'order-status order-status-completed'
  }
  if (status === 'CANCELED') {
    return 'order-status order-status-canceled'
  }
  return 'order-status'
}

const dateFormat = (time) => {
  return moment(time).format('YYYY-MM-DD HH:mm:ss')
}

const toFloor = (value, scale = 6) => {
  if (value === null || value === undefined) return 0
  return Math.floor(value * Math.pow(10, scale)) / Math.pow(10, scale)
}

const init = async () => {
  disconnectSocket()
  destroyKline()
  const params = router.currentRoute.value.params.pair || defaultPath.value
  if (!router.currentRoute.value.params.pair) {
    router.push(`/exchange/${defaultPath.value}`)
  }

  const base = params.split('_')[1]
  if (base) {
    basecion.value = base
  }

  const coin = params.toUpperCase().split('_')[0]
  const baseCoin = params.toUpperCase().split('_')[1]
  currentCoin.symbol = `${coin}/${baseCoin}`
  currentCoin.coin = coin
  currentCoin.base = baseCoin
  resetDepthPlate(currentCoin.symbol)

  store.commit('navigate', '/exchange')
  store.commit('setSkin', skin.value)

  getCNYRate()
  await getSymbolScale()
  getCoinInfo()
  await getSymbol()
  getPlate()
  getPlateFull()
  getTrade()

  if (isLogin.value) {
    getWallet()
    getCurrentOrder()
    getHistoryOrder()
  }

  sliderBuyLimitPercent.value = 0
  sliderSellLimitPercent.value = 0
  sliderBuyMarketPercent.value = 0
}

const getCNYRate = () => {
  silentPost(host + '/market/exchange-rate/usd/cny')
    .then(response => {
      const resp = unwrapPayload(response)
      if (resp === null || resp === undefined) return
      CNYRate.value = resp
    })
}

const getSymbolScale = () => {
  return silentPost(host + api.market.symbolInfo, { symbol: currentCoin.symbol })
    .then(response => {
      const resp = unwrapPayload(response)
      if (resp) {
        baseCoinScale.value = resp.baseCoinScale
        coinScale.value = resp.coinScale
        symbolFee.value = resp.fee
        enableMarketBuy.value = resp.enableMarketBuy
        enableMarketSell.value = resp.enableMarketSell
        exchangeable.value = resp.exchangeable

        publishType.value = resp.publishType
        startTime.value = resp.startTime
        endTime.value = resp.endTime
        clearTime.value = resp.clearTime
        currentTime.value = parseInt(resp.currentTime / 1000)
        publishAmount.value = resp.publishAmount
        publishPrice.value = resp.publishPrice

        const temCurT = moment(resp.currentTime).format('YYYY-MM-DD HH:mm:ss')
        if (temCurT < clearTime.value) {
          if (publishType.value === 'QIANGGOU' || publishType.value === 'FENTAN') {
            showPublish.value = true
            showCountDown.value = true
          }
        }
      }
    })
}

const getCoinInfo = () => {
  silentPost(host + api.market.coinInfo, { unit: currentCoin.coin })
    .then(response => {
      const resp = unwrapPayload(response)
      if (resp) {
        Object.assign(coinInfo, resp || {})
      }
    })
}

const getWallet = () => {
  silentPost(host + api.ucenter.wallet, {})
    .then(response => {
      const account = unwrapPayload(response)
      if (account) {
        const balances = pickWalletBalances({
          wallets: account,
          baseSymbol: currentCoin.base,
          coinSymbol: currentCoin.coin
        })
        wallet.base = balances.base
        wallet.coin = balances.coin
      }
    })
}

const getCurrentOrder = () => {
  silentPost(host + api.exchange.currentOrder, {
    symbol: currentCoin.symbol,
    pageNo: 1,
    pageSize: 20
  })
    .then(response => {
      const resp = unwrapPayload(response)
      currentOrder.rows = resp?.content || []
    })
}

const getHistoryOrder = () => {
  silentPost(host + api.exchange.historyOrder, {
    symbol: currentCoin.symbol,
    pageNo: historyOrder.page,
    pageSize: historyOrder.pageSize
  })
    .then(response => {
      const resp = unwrapPayload(response)
      historyOrder.rows = resp?.content || []
    })
}

const cancel = (index) => {
  const order = currentOrder.rows[index]
  post(host + api.exchange.cancel, { orderId: order.orderId || order.id })
    .then(response => {
      const resp = response.data
      if (resp.code === 0) {
        ElMessage.success('撤单成功')
        getWallet()
        getCurrentOrder()
        getHistoryOrder()
      } else {
        ElMessage.error(resp.message || '撤单失败')
      }
    })
}

const buyWithLimitPrice = () => {
  if (!form.buy.limitAmount) {
    ElMessage.error('请输入买入数量')
    return
  }

  const maxAmount = wallet.base / form.buy.limitPrice
  if (form.buy.limitAmount > maxAmount) {
    ElMessage.error('买入数量超过可用余额')
    return
  }

  buying.value = true

  const params = {
    symbol: currentCoin.symbol,
    price: form.buy.limitPrice,
    amount: form.buy.limitAmount,
    direction: 'BUY',
    type: 'LIMIT_PRICE',
    useDiscount: 0
  }

  post(host + api.exchange.order, params)
    .then(response => {
      const resp = response.data
      if (resp.code === 0) {
        ElMessage.success('买入委托成功')
        getWallet()
        getCurrentOrder()
        getHistoryOrder()
        form.buy.limitAmount = 0
        sliderBuyLimitPercent.value = 0
      } else {
        ElMessage.error(resp.message || '买入失败')
      }
    })
    .catch(() => {
      ElMessage.error('买入失败')
    })
    .finally(() => {
      buying.value = false
    })
}

const buyWithMarketPrice = () => {
  if (!form.buy.marketAmount) {
    ElMessage.error('请输入买入金额')
    return
  }

  buying.value = true

  const params = {
    symbol: currentCoin.symbol,
    amount: form.buy.marketAmount,
    direction: 'BUY',
    type: 'MARKET_PRICE'
  }

  post(host + api.exchange.order, params)
    .then(response => {
      const resp = response.data
      if (resp.code === 0) {
        ElMessage.success('市价买入成功')
        getWallet()
        getCurrentOrder()
        getHistoryOrder()
        form.buy.marketAmount = 0
        sliderBuyMarketPercent.value = 0
      } else {
        ElMessage.error(resp.message || '买入失败')
      }
    })
    .catch(() => {
      ElMessage.error('买入失败')
    })
    .finally(() => {
      buying.value = false
    })
}

const sellLimitPrice = () => {
  if (!form.sell.limitAmount) {
    ElMessage.error('请输入卖出数量')
    return
  }

  selling.value = true

  const params = {
    symbol: currentCoin.symbol,
    price: form.sell.limitPrice,
    amount: form.sell.limitAmount,
    direction: 'SELL',
    type: 'LIMIT_PRICE'
  }

  post(host + api.exchange.order, params)
    .then(response => {
      const resp = response.data
      if (resp.code === 0) {
        ElMessage.success('卖出委托成功')
        getWallet()
        getCurrentOrder()
        getHistoryOrder()
        form.sell.limitAmount = 0
        sliderSellLimitPercent.value = 0
      } else {
        ElMessage.error(resp.message || '卖出失败')
      }
    })
    .catch(() => {
      ElMessage.error('卖出失败')
    })
    .finally(() => {
      selling.value = false
    })
}

const sellMarketPrice = () => {
  if (!form.sell.marketAmount) {
    ElMessage.error('请输入卖出数量')
    return
  }

  selling.value = true

  const params = {
    symbol: currentCoin.symbol,
    amount: form.sell.marketAmount,
    direction: 'SELL',
    type: 'MARKET_PRICE'
  }

  post(host + api.exchange.order, params)
    .then(response => {
      const resp = response.data
      if (resp.code === 0) {
        ElMessage.success('市价卖出成功')
        getWallet()
        getCurrentOrder()
        getHistoryOrder()
        form.sell.marketAmount = 0
        sliderSellMarketPercent.value = 0
      } else {
        ElMessage.error(resp.message || '卖出失败')
      }
    })
    .catch(() => {
      ElMessage.error('卖出失败')
    })
    .finally(() => {
      selling.value = false
    })
}

const showPublishMask = () => {}
const hidePublishMask = () => {}
const hideCountDown = () => {}

// 监听表单价格变化
watch(() => form.buy.limitPrice, (val) => {
  if (val > 0) {
    const amount = toFloor(
      wallet.base / val * sliderBuyLimitPercent.value * 0.01,
      coinScale.value
    )
    form.buy.limitAmount = amount
    form.buy.limitTurnover = toFloor(val * amount, baseCoinScale.value)
  }
})

watch(() => form.buy.limitAmount, () => {
  form.buy.limitTurnover = toFloor(
    form.buy.limitAmount * form.buy.limitPrice,
    baseCoinScale.value
  )
})

watch(() => form.sell.limitPrice, () => {
  form.sell.limitTurnover = toFloor(
    form.sell.limitPrice * form.sell.limitAmount,
    coinScale.value
  )
})

watch(() => form.sell.limitAmount, () => {
  form.sell.limitTurnover = toFloor(
    form.sell.limitAmount * form.sell.limitPrice,
    coinScale.value
  )
})

// 监听 Slider 变化
watch(sliderBuyLimitPercent, () => {
  if (form.buy.limitPrice > 0) {
    form.buy.limitAmount = toFloor(
      wallet.base / form.buy.limitPrice * sliderBuyLimitPercent.value * 0.01,
      coinScale.value
    )
  }
})

watch(sliderSellLimitPercent, () => {
  form.sell.limitAmount = toFloor(
    wallet.coin * sliderSellLimitPercent.value * 0.01,
    coinScale.value
  )
})

watch(sliderBuyMarketPercent, () => {
  form.buy.marketAmount = toFloor(
    wallet.base * sliderBuyMarketPercent.value * 0.01,
    baseCoinScale.value
  )
})

watch(sliderSellMarketPercent, () => {
  form.sell.marketAmount = toFloor(
    wallet.coin * sliderSellMarketPercent.value * 0.01,
    coinScale.value
  )
})

// 监听语言变化
watch(() => store.state.lang, (newLang) => {
  lang.value = newLang
  if (tvWidget) {
    initKline()
  }
})

watch(() => router.currentRoute.value.params.pair, (pair, previousPair) => {
  if (!pair || pair === previousPair) {
    return
  }
  init()
})

// 生命周期
onMounted(() => {
  store.commit('navigate', '/exchange')

  // 检查登录状态
  isLogin.value = store.getters?.isLogin || !!localStorage.getItem('TOKEN')
  member.value = store.getters?.member || JSON.parse(localStorage.getItem('MEMBER') || '{}')
  memberId.value = member.value?.id || 0
  lang.value = store.state.lang || 'zh'

  init()
})

onBeforeUnmount(() => {
  disconnectSocket()
  destroyKline()
})
</script>

<style scoped lang="scss">
@import "../../assets/css/exchange.css";

$night-bg: #0b1520;
$night-headerbg: #27313e;
$night-contentbg: #192330;
$night-color: #fff;

.exchange {
  color: #fff;
  background-color: #0b1520;
  min-height: calc(100vh - 200px);
  --el-bg-color: #192330;
  --el-bg-color-overlay: #192330;
  --el-fill-color-blank: #192330;
  --el-fill-color-light: #27313e;
  --el-fill-color-lighter: #243041;
  --el-border-color: #324158;
  --el-border-color-lighter: #243041;
  --el-text-color-regular: #d7deea;
  --el-text-color-primary: #ffffff;
  --el-text-color-placeholder: #8f9ca5;
  --el-table-bg-color: #192330;
  --el-table-tr-bg-color: #192330;
  --el-table-row-hover-bg-color: #243041;
  --el-table-header-bg-color: #27313e;
  --el-table-border-color: #324158;
  --el-table-text-color: #d7deea;
  --el-table-header-text-color: #aeb9d0;
  --el-input-bg-color: #27313e;
  --el-input-border-color: #324158;
  --el-input-hover-border-color: #4a5a73;
  --el-input-focus-border-color: #f0a70a;

  .main {
    width: 99%;
    margin-left: 0.5%;
    display: flex;
    align-items: stretch;
    margin-top: 5px;
    max-width: 100%;
    overflow-x: hidden;
    gap: 5px;
    height: calc(100vh - 200px);

    .left {
      flex: 0 0 20%;
      border-radius: 0px;
      margin-right: 0;
      min-width: 0;
      overflow: hidden;
      display: flex;
      flex-direction: column;
      height: 100%;

      .plate-handlers {
        position: relative;
        background-color: #192330;
        border-bottom: 1px solid #27313e;
        font-size: 0;
        height: 40px;
        line-height: 40px;
        flex-shrink: 0;

        > span {
          font-size: 14px;
          padding: 0 20px;
          cursor: pointer;
          display: inline-block;
          color: #8f9ca5;

          &.active {
            color: #fff;
            border-bottom: 2px solid #f0a70a;
          }
        }
      }

      .plate-nowprice {
        text-align: center;
        background-color: #27313e;
        display: flex;
        flex-direction: column;
        justify-content: center;
        min-height: 48px;
        padding: 4px 10px;
        font-size: 14px;
        font-weight: 500;
        gap: 2px;

        .plate-nowprice__labels,
        .plate-nowprice__value {
          display: flex;
          align-items: center;
          justify-content: space-between;
          width: 100%;
        }

        .plate-nowprice__label {
          color: #c7cce6;
          font-size: 12px;
          letter-spacing: 0.3px;
        }

        .plate-nowprice__labels {
          font-size: 12px;
        }

        .plate-nowprice__value {
          justify-content: center;
          line-height: 1;
        }

        .price {
          font-size: 18px;
          margin-right: 10px;
        }

        .price-cny {
          font-size: 12px;
          margin-left: 10px;
          font-weight: 400;
          color: rgba(219, 222, 246, 0.3);
        }
      }
    }

    .center {
      flex: 0 0 60%;
      margin-right: 5px;
      min-width: 0;
      overflow: hidden;
      display: flex;
      flex-direction: column;
      height: 100%;

      .symbol {
        flex-shrink: 0;
      }

      .imgtable {
        height: 350px;
        position: relative;
        overflow: hidden;
        margin-bottom: 5px;
        flex-shrink: 0;

        .chart-panel {
          height: calc(100% - 40px);
        }

        .handler {
          height: 40px;
          display: flex;
          justify-content: flex-end;
          align-items: center;
          padding: 6px 16px;
          background-color: #192330;
          border-top: 1px solid #223047;
          position: relative;
          z-index: 2;

          > span {
            background-color: #2c3b59;
            color: #c7cce6;
            padding: 5px 10px;
            cursor: pointer;
            font-size: 13px;
            opacity: 1;
            border: 1px solid transparent;
            border-radius: 4px;
            line-height: 1;

            & + span {
              margin-left: 8px;
            }

            &.active {
              color: #fff;
              background-color: #3a4a6a;
              border-color: #66789d;
            }
          }
        }
      }

      .trade_wrap {
        background-color: #192330;
        flex: 1;
        min-height: 0;
        display: flex;
        flex-direction: column;

        .trade_menu {
          position: relative;
          background-color: #192330;
          border-bottom: 1px solid #27313e;
          font-size: 0;
          height: 40px;
          line-height: 40px;
          flex-shrink: 0;
        }

        .trade_bd {
          flex: 1;
          min-height: 0;
          display: flex;
          flex-direction: column;
          justify-content: flex-start;
          padding-bottom: 8px;

          // 限价/市价网格容器
          .trade-grid {
            padding: 8px;
            margin-bottom: 0;
          }
        }

        .trade_menu {
          > span {
            font-size: 16px;
            padding: 11px 20px;
            cursor: pointer;

            &.active {
              color: #fff;
              border-bottom: 2px solid #f0a70a;
            }
          }
        }

        .trade_panel {
          position: relative;

          .trade-form-title {
            margin-bottom: 12px;
            font-size: 15px;
            font-weight: 600;
            letter-spacing: 0.3px;
          }

          .buy-title {
            color: #5ec66b;
          }

          .sell-title {
            color: #ff7a7a;
          }

          .trade-hint {
            margin: -4px 0 10px;
            font-size: 12px;
            color: #7d8aa6;
          }

          .mask {
            position: absolute;
            width: 100%;
            height: 100%;
            display: flex;
            background-color: rgba(0, 52, 77, 0.8);
            justify-content: center;
            align-items: center;
            z-index: 100;
            font-size: 24px;
          }

          .publish-mask {
            position: absolute;
            width: 100%;
            height: 100%;
            background-color: rgba(0, 37, 64, 0.93);
            display: flex;
            justify-content: center;
            align-items: center;
            z-index: 101;
            font-size: 24px;
          }
        }
      }
    }

    .right {
      flex: 0 0 19%;
      margin-right: 0;
      min-width: 0;
      overflow: hidden;
      display: flex;
      flex-direction: column;
      height: 100%;

      .coin-menu {
        background-color: #192330;
        flex: 1;
        min-height: 0;
        overflow: hidden;
        display: flex;
        flex-direction: column;

        > div:first-child {
          flex-shrink: 0;
        }

        .sc_filter {
          flex-shrink: 0;
        }

        .el-table {
          flex: 1;
          min-height: 0;
          overflow: hidden;

          :deep(.el-table__body-wrapper) {
            flex: 1;
            min-height: 0;
            overflow-y: auto;
          }
        }
      }
    }
  }

  .symbol {
    display: flex;
    justify-content: space-between;
    padding: 15px 30px;
    margin-bottom: 5px;
    align-items: center;
    background-color: #192330;
    line-height: 1;

    .item {
      .price-cny {
        font-size: 12px;
        color: #546886;
      }

      .coin {
        font-size: 20px;
      }

      .text {
        width: 100%;
        display: block;
        font-size: 12px;
        color: #999;
        margin-bottom: 5px;
      }

      .num {
        font-size: 12px;
        color: #fff;
      }
    }
  }

  // 委托订单区域
  .order {
    width: 98.6%;
    margin-left: 0.5%;
    margin-bottom: 10px;
    margin-top: 0;
    flex-shrink: 0;
    border-top: 1px solid #27313e;

    .order-handler {
      background-color: #192330;
      font-size: 0;

      > span {
        padding: 0 20px;
        font-size: 14px;
        display: inline-block;
        color: #fff;
        cursor: pointer;
        line-height: 36px;
        background-color: #192330;

        &.active {
          color: #f0a70a;
          background-color: #27313e;
        }
      }
    }

    .table {
      background-color: #192330;

      :deep(.el-table__body-wrapper) {
        max-height: 200px;
        overflow-y: auto;
      }
    }
  }

  .balance-label {
    color: #7d8aa6;
  }

  .balance-value {
    color: #fff;
    font-weight: 600;
  }

  .balance-unit {
    color: #c8d3ea;
  }

  .order-status {
    display: inline-flex;
    align-items: center;
    padding: 2px 8px;
    border-radius: 999px;
    font-size: 12px;
    line-height: 1.5;
  }

  .order-status-completed {
    color: #f0a70a;
    background: rgba(240, 167, 10, 0.12);
  }

  .order-status-canceled {
    color: #9aa3b2;
    background: rgba(154, 163, 178, 0.14);
  }
}

.exchange.day {
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
    }

    .right {
      .coin-menu {
        background-color: #fff;
        box-shadow: 0 2px 2px #ccc;
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

  // 日间模式 - 并排交易面板
  .trade_bd {
    .trade-grid-item {
      background: #fff;
      border: 1px solid #e0e0e0;
    }

    .hd-compact {
      &.hd_login {
        background: #f5f5f5;
        color: #333;
      }

      .balance-label {
        color: #7d8aa6;
      }

      .balance-value {
        color: #333;
        font-weight: 600;
      }

      .balance-unit {
        color: #666;
      }
    }

    .form-row-inline {
      .inline-label {
        color: #7d8aa6;
      }

      .inline-unit {
        color: #666;
      }
    }

    .form-row-total {
      border-top: 1px solid #e8e8e8;

      .total-label {
        color: #7d8aa6;
      }

      .total-value {
        color: #333;
      }
    }
  }
}

#kline_container {
  background: #192330;
  height: 100%;
}

.hidden {
  display: none !important;
}

.exchange :deep(.el-input__wrapper) {
  background-color: #27313e;
  box-shadow: 0 0 0 1px #324158 inset !important;
}

.exchange :deep(.el-input__inner) {
  color: #fff;
}

.exchange :deep(.el-input.is-disabled .el-input__wrapper) {
  background-color: #243041 !important;
  box-shadow: 0 0 0 1px #324158 inset !important;
}

.exchange :deep(.el-table) {
  background-color: #192330 !important;
  color: #d7deea !important;
}

.exchange :deep(.el-table__inner-wrapper),
.exchange :deep(.el-table__header-wrapper),
.exchange :deep(.el-table__body-wrapper),
.exchange :deep(.el-scrollbar__view),
.exchange :deep(.el-scrollbar__wrap) {
  background-color: #192330 !important;
}

.exchange :deep(.el-table__header),
.exchange :deep(.el-table__body) {
  background-color: #192330 !important;
}

.exchange :deep(.el-table th.el-table__cell),
.exchange :deep(.el-table tr),
.exchange :deep(.el-table td.el-table__cell),
.exchange :deep(.el-table__inner-wrapper::before) {
  background-color: #192330 !important;
}

.exchange :deep(.el-table th.el-table__cell) {
  background-color: #27313e !important;
  color: #aeb9d0 !important;
  border-bottom-color: #324158 !important;
}

.exchange :deep(.el-table td.el-table__cell) {
  color: #d7deea !important;
  border-bottom-color: #243041 !important;
}

.exchange :deep(.el-table__row:hover > td.el-table__cell) {
  background-color: #243041 !important;
}

.exchange :deep(.el-table__expanded-cell) {
  background-color: #192330 !important;
  color: #d7deea !important;
}

.exchange :deep(.el-table__empty-block),
.exchange :deep(.el-table__empty-text) {
  background-color: #192330 !important;
  color: #8f9ca5;
}

.exchange :deep(.el-button) {
  border-radius: 4px;
}

.exchange :deep(.plate-wrap .el-table__body-wrapper) {
  flex: 1;
  min-height: 0;
  overflow-y: auto;
  overflow-x: hidden;
}

// 左侧区域 flex 布局
.plate-wrap {
  .plate-handlers {
    flex-shrink: 0;
  }

  .plate-book {
    flex: 1;
    min-height: 0;
    display: flex;
    flex-direction: column;
    gap: 0;
  }

  .plate-table {
    flex: 1;
    min-height: 0;
    display: flex;
    flex-direction: column;

    .el-table {
      flex: 1;
      min-height: 0;
    }
  }

  .plate-nowprice {
    flex-shrink: 0;
  }

  .trade-wrap {
    flex: 0 0 176px;
    min-height: 0;
    display: flex;
    flex-direction: column;
    overflow: hidden;
    margin-top: 10px;

    .trade-wrap__header {
      height: 40px;
      line-height: 40px;
      padding: 0 20px;
      color: #fff;
      font-size: 14px;
      background-color: #192330;
      border-bottom: 1px solid #27313e;
      flex-shrink: 0;
    }

    :deep(.el-table__body-wrapper) {
      flex: 1;
      min-height: 0;
    }
  }
}

.exchange :deep(.right .el-table__body-wrapper) {
  max-height: 585px;
  overflow-y: auto;
  overflow-x: hidden;
}

.exchange :deep(.order .el-table__body-wrapper) {
  max-height: 240px;
  overflow-y: auto;
}

.coin-info {
  color: #8f9ca5;
  display: -webkit-box;
  -webkit-box-orient: vertical;
  -webkit-line-clamp: 5;
  overflow: hidden;
  padding-top: 15px;
}

.favor-icon-active {
  color: #f0a70a;
}

// 买卖盘样式
.buy {
  color: #00b275;
}

.sell {
  color: #f15057;
}

// 交易表单样式 - 紧凑型布局
.trade_bd {
  // 并排网格布局
  .trade-grid {
    display: grid;
    grid-template-columns: 1fr 1fr;
    gap: 8px;
    padding: 8px;
    flex-shrink: 0;
  }

  .trade-grid-item {
    background: #1f2a38;
    border-radius: 4px;
    padding: 10px;
    // 固定高度确保左右一致
    height: 165px;

    &.trade-grid-buy,
    &.trade-grid-sell {
      border: none;
    }
  }

  // 紧凑头部
  .hd-compact {
    margin: -10px -10px 6px;
    padding: 6px 10px;
    border-radius: 4px 4px 0 0;
    display: flex;
    align-items: center;
    gap: 6px;
    font-size: 10px;

    &.hd_login {
      background: #27313e;
      color: #fff;
    }

    .balance-label {
      color: #7d8aa6;
    }

    .balance-value {
      color: #fff;
      font-weight: 600;
    }

    .balance-unit {
      color: #c8d3ea;
    }

    a {
      color: #f0a70a;
      font-size: 10px;
    }
  }

  // 标题行
  .form-row-title {
    margin-bottom: 6px;

    .form-title {
      font-size: 12px;
      font-weight: 600;

      &.buy-title {
        color: #5ec66b;
      }

      &.sell-title {
        color: #ff7a7a;
      }
    }
  }

  // 行内表单
  .form-row-inline {
    display: flex;
    align-items: center;
    margin-bottom: 5px;
    position: relative;

    .inline-label {
      width: 40px;
      font-size: 11px;
      color: #546886;
      flex-shrink: 0;
    }

    .el-input {
      flex: 1;
      margin: 0 6px;

      :deep(.el-input__wrapper) {
        padding: 0 8px;
        height: 26px;
        border-radius: 3px;
        box-shadow: 0 0 0 1px #324158 inset !important;
      }

      :deep(.el-input__inner) {
        font-size: 12px;
        height: 26px;
        line-height: 26px;
        color: #fff;
      }
    }

    .inline-unit {
      width: 35px;
      font-size: 11px;
      color: #fff;
      text-align: right;
      flex-shrink: 0;
    }
  }

  // 合计行
  .form-row-total {
    display: flex;
    justify-content: space-between;
    align-items: center;
    padding: 5px 0;
    margin: 5px 0;
    border-top: 1px solid #2a384d;

    .total-label {
      font-size: 11px;
      color: #546886;
    }

    .total-value {
      font-size: 11px;
      color: #fff;
      font-weight: 500;
    }
  }

  // 紧凑按钮
  .btn-compact {
    width: 100%;
    height: 28px;
    font-size: 13px;
    font-weight: 500;
    margin-top: 4px;
    border-radius: 4px;

    &.el-button--success {
      background: linear-gradient(135deg, #00b275 0%, #009966 100%);
      border: none;
    }

    &.el-button--danger {
      background: linear-gradient(135deg, #f15057 0%, #dd444a 100%);
      border: none;
    }
  }

  // 旧版 panel 样式保留
  .panel {
    padding: 10px;

    .hd {
      display: flex;
      justify-content: space-between;
      align-items: center;
      padding: 10px;
      background-color: #27313e;
      margin-bottom: 10px;

      &.hd_login {
        color: #fff;
      }

      a {
        color: #f0a70a;
      }
    }

    .bd {
      .before {
        color: #546886;
        font-size: 12px;
      }

      .after {
        color: #fff;
        font-size: 12px;
        position: absolute;
        right: 10px;
        top: 50%;
        transform: translateY(-50%);
      }

      .math_price {
        color: #546886;
        font-size: 12px;
        margin-top: 5px;
      }

      .slider-wrap {
        margin: 15px 0;
      }

      .total {
        font-size: 12px;
        color: #546886;
        margin: 10px 0;

        span {
          color: #fff;
        }
      }
    }

    .bg-green {
      background-color: #00b275;
      border-color: #00b275;
      width: 100%;
      margin-top: 10px;
    }

    .bg-red {
      background-color: #f15057;
      border-color: #f15057;
      width: 100%;
      margin-top: 10px;
    }

    .bg-gray {
      background-color: #546886;
      border-color: #546886;
      width: 100%;
      margin-top: 10px;
    }
  }
}

// 倒计时面板
.lightning-panel {
  position: relative;
  padding: 20px;
  margin-bottom: 10px;

  img {
    position: absolute;
    top: 0;
    left: 0;
    width: 100%;
    height: 100%;
    object-fit: cover;
  }

  .l-content {
    position: relative;
    z-index: 10;
    color: #fff;
    text-align: center;

    .l-info {
      margin: 15px 0;
      font-size: 14px;
    }

    .l-detail {
      a {
        color: #f0a70a;
      }
    }
  }
}
</style>
