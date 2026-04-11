import $ from 'jquery';

const pickDefinedNumber = function() {
    for (var i = 0; i < arguments.length; i++) {
        var value = arguments[i];
        if (typeof value === 'number' && !Number.isNaN(value)) {
            return value;
        }
    }
    return null;
};

export var resolvePriceScale = function(scale, coin) {
    var precision = pickDefinedNumber(
        scale,
        coin && coin.baseCoinScale,
        coin && coin.priceScale,
        2
    );
    return Math.pow(10, precision);
};

export var resolvePricePrecision = function(scale, coin) {
    return pickDefinedNumber(
        scale,
        coin && coin.baseCoinScale,
        coin && coin.priceScale,
        2
    );
};

export var normalizeRealtimeBarPayload = function(payload) {
    return {
        time: pickDefinedNumber(payload && payload.time, payload && payload.Time),
        open: pickDefinedNumber(payload && payload.openPrice, payload && payload.OpenPrice),
        high: pickDefinedNumber(payload && payload.highestPrice, payload && payload.HighestPrice),
        low: pickDefinedNumber(payload && payload.lowestPrice, payload && payload.LowestPrice),
        close: pickDefinedNumber(payload && payload.closePrice, payload && payload.ClosePrice),
        volume: pickDefinedNumber(payload && payload.volume, payload && payload.Volume, 0)
    };
};

export var resolutionToMilliseconds = function(resolution) {
    if (/^\d+$/.test(resolution)) {
        return parseInt(resolution, 10) * 60 * 1000;
    }
    if (resolution === '1D' || resolution === 'D') {
        return 24 * 60 * 60 * 1000;
    }
    if (resolution === '1W' || resolution === 'W') {
        return 7 * 24 * 60 * 60 * 1000;
    }
    if (resolution === '1M' || resolution === 'M') {
        return 30 * 24 * 60 * 60 * 1000;
    }
    return 60 * 60 * 1000;
};

export var alignTimestampToResolution = function(timestamp, resolution) {
    var interval = resolutionToMilliseconds(resolution);
    return Math.floor(timestamp / interval) * interval;
};

export var normalizeTradePayload = function(payload) {
    if (Array.isArray(payload)) {
        return payload.length > 0 ? payload[payload.length - 1] : null;
    }
    return payload || null;
};

export var updateBarFromTrade = function(lastBar, tradePayload, resolution) {
    var trade = normalizeTradePayload(tradePayload);
    var price = pickDefinedNumber(trade && trade.price, trade && trade.Price);
    if (price === null) {
        return lastBar;
    }

    var tradeTime = pickDefinedNumber(trade && trade.time, trade && trade.Time, Date.now());
    var tradeVolume = pickDefinedNumber(
        trade && trade.amount,
        trade && trade.Amount,
        trade && trade.volume,
        trade && trade.Volume,
        0
    );
    var nextBarTime = alignTimestampToResolution(tradeTime, resolution);

    if (!lastBar || pickDefinedNumber(lastBar.time) === null) {
        return {
            time: nextBarTime,
            open: price,
            high: price,
            low: price,
            close: price,
            volume: tradeVolume
        };
    }

    var lastBarTime = alignTimestampToResolution(lastBar.time, resolution);
    if (nextBarTime > lastBarTime) {
        var previousClose = pickDefinedNumber(lastBar.close, price);
        return {
            time: nextBarTime,
            open: previousClose,
            high: Math.max(previousClose, price),
            low: Math.min(previousClose, price),
            close: price,
            volume: tradeVolume
        };
    }

    if (nextBarTime < lastBarTime) {
        return lastBar;
    }

    return {
        time: lastBar.time,
        open: lastBar.open,
        high: Math.max(pickDefinedNumber(lastBar.high, price), price),
        low: Math.min(pickDefinedNumber(lastBar.low, price), price),
        close: price,
        volume: pickDefinedNumber(lastBar.volume, 0) + tradeVolume
    };
};

var WebsockFeed = function(url,coin,socket,scale){
    this._datafeedURL = url;
    this.coin = coin;
    this.socket = socket;
    this.lastBar = null;
    this.currentBar = null;
    this.subscribe = true;
    this.scale = scale;
    this.subscriptions = {};
};

WebsockFeed.prototype.onReady=function(callback){
    var config = {};
    config.exchanges = [];
    config.supported_resolutions = ["1","5","15","30","60","240","1D","1W","1M"];
    config.supports_group_request = false;
    config.supports_marks = false;
    config.supports_search = false;
    config.supports_time = true;
    config.supports_timescale_marks = false;

    $("#"+window.tvWidget.id).contents().on("click",".date-range-list>a",function(){
      if (window.tvWidget) {
        if ($(this).html() == "分时") {
          $(this).parent().addClass("real-op").removeClass("common-op");
          window.tvWidget.chart().setChartType(3);
        }else {
          $(this).parent().addClass("common-op").removeClass("real-op");
          window.tvWidget.chart().setChartType(1);
        }
      }
    });

    setTimeout(function() {
        callback(config);
    }, 0);
};

WebsockFeed.prototype.subscribeBars = function(symbolInfo, resolution, onRealtimeCallback, listenerGUID, onResetCacheNeededCallback) {
	var that = this;
    var tradeTopic = '/topic/market/trade/' + symbolInfo.name;
    var klineTopic = '/topic/market/kline/' + symbolInfo.name;
    this.unsubscribeBars(listenerGUID);

    var tradeHandler = function(msg) {
        var resp = JSON.parse(msg);
        var nextBar = updateBarFromTrade(that.lastBar, resp, resolution);
        if (nextBar) {
            that.lastBar = nextBar;
            that.currentBar = nextBar;
            onRealtimeCallback(nextBar);
        }
    };
    var klineHandler = function(msg) {
        if (resolution != "1") return;
        var resp = JSON.parse(msg);
        that.lastBar = normalizeRealtimeBarPayload(resp);
        that.currentBar = that.lastBar;
        onRealtimeCallback(that.lastBar);
    };

    this.subscriptions[listenerGUID] = {
        tradeTopic: tradeTopic,
        tradeHandler: tradeHandler,
        klineTopic: klineTopic,
        klineHandler: klineHandler
    };

    this.socket.on(tradeTopic, tradeHandler);
    this.socket.on(klineTopic, klineHandler);
};

WebsockFeed.prototype.unsubscribeBars = function(subscriberUID){
    this.subscribe = false;
    var subscription = this.subscriptions[subscriberUID];
    if (!subscription || !this.socket) {
        return;
    }
    if (typeof this.socket.off === 'function') {
        this.socket.off(subscription.tradeTopic, subscription.tradeHandler);
        this.socket.off(subscription.klineTopic, subscription.klineHandler);
    } else if (typeof this.socket.removeListener === 'function') {
        this.socket.removeListener(subscription.tradeTopic, subscription.tradeHandler);
        this.socket.removeListener(subscription.klineTopic, subscription.klineHandler);
    }
    delete this.subscriptions[subscriberUID];
}

WebsockFeed.prototype.resolveSymbol = function(symbolName, onSymbolResolvedCallback, onResolveErrorCallback){
    // var data = {"name":this.coin.symbol,"exchange-traded":"","exchange-listed":"","minmov":1,"minmov2":0,"pointvalue":1,"has_intraday":true,"has_no_volume":false,"description":this.coin.symbol,"type":"bitcoin","session":"24x7","supported_resolutions":["1","5","15","30","60","240","1D","1W","1M"],"pricescale":500,"ticker":"","timezone":"Asia/Shanghai"};
  // var data = {"name":this.coin.symbol,"exchange-traded":"","exchange-listed":"","minmov":1,"volumescale":10000,"has_daily":true,"has_weekly_and_monthly":true,"has_intraday":true,"description":this.coin.symbol,"type":"bitcoin","session":"24x7","supported_resolutions":["1","5","15","30","60","240","1D","1W","1M"],"pricescale":100,"ticker":"","timezone":"Asia/Shanghai"};
  var precision = resolvePricePrecision(this.scale, this.coin);
  var data = {"name":this.coin.symbol,"exchange-traded":"","exchange-listed":"","minmov":1,"minmove2":0,"fractional":false,"format":"price","volumescale":10000,"volume_precision":precision,"has_daily":true,"has_weekly_and_monthly":true,"has_intraday":true,"description":this.coin.symbol,"type":"bitcoin","session":"24x7","supported_resolutions":["1","5","15","30","60","1D","1W","1M"],"pricescale":resolvePriceScale(this.scale, this.coin),"ticker":"","timezone":"Asia/Shanghai"};
    setTimeout(function() {
        onSymbolResolvedCallback(data);
    }, 0);
};

WebsockFeed.prototype._send = function(url, params) {
    var request = url;
    if (params) {
        for (var i = 0; i < Object.keys(params).length; ++i) {
            var key = Object.keys(params)[i];
            var value = encodeURIComponent(params[key]);
            request += (i === 0 ? '?' : '&') + key + '=' + value;
        }
    }

    return $.ajax({
        type: 'GET',
        url: request,
        dataType: 'json'
    });
};

WebsockFeed.prototype.getBars = function(symbolInfo, resolution, from, to, onHistoryCallback, onErrorCallback, firstDataRequest){
    var bars = [];
    var that = this;
    this._send(this._datafeedURL+'/history',{
        symbol: symbolInfo.name,
        from: from*1000,
        to: firstDataRequest ? new Date().getTime():to*1000,
        resolution: resolution
    })
    .done(function(response) {
		if (response.code == "0") {
			var data = response.data;
			for(var i = 0;i<data.length;i++){
				var item = data[i];
				bars.push({time:item[0],open:item[1],high:item[2],low:item[3],close:item[4],volume:item[5]})
			}
			that.lastBar = bars.length > 0 ? bars[bars.length-1]:null;
			that.currentBar = that.lastBar;
			var noData = bars.length == 0;
			var retBars = onHistoryCallback(bars,{noData:noData});
			
		}else{
			onErrorCallback(response.message);
		}
    })
    .fail(function(reason) {
        onErrorCallback(reason);
    });
};
WebsockFeed.prototype.periodLengthSeconds = function(resolution, requiredPeriodsCount) {
    var daysCount = 0;
    if (resolution === 'D') {
        daysCount = requiredPeriodsCount;
    } else if (resolution === 'M') {
        daysCount = 31 * requiredPeriodsCount;
    } else if (resolution === 'W') {
        daysCount = 7 * requiredPeriodsCount;
    }
    else if(resolution === 'H'){
        daysCount = requiredPeriodsCount * resolution / 24;
    }
    else {
        daysCount = requiredPeriodsCount * resolution / (24 * 60);
    }

    return daysCount * 24 * 60 * 60;
};

export default {WebsockFeed}
