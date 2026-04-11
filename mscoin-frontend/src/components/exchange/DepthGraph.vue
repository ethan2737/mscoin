<template>
  <div id="depth_chart" style="position: relative;width: 100%;height: 100%;">
    <canvas id="depthGraph" style="position:absolute;left:0"></canvas>
    <canvas id="depthGraph_cover" style="position:absolute;left:0"></canvas>
  </div>
</template>

<script>
export default {
  name: "depth-graph",
  data: function() {
    return {
      xObj: {},
      yObj: {},
      canvas: null,
      context: null,
      canvas_cover: null,
      context_cover: null,
      pWidth: 900,
      pHeight: 300,
      buyPoints: [],
      sellPoints: [],
      values: null
    };
  },
  // props:{
  //     values:{
  //       type:Object,
  //       required:true
  //     }
  // },
  // computed:{
  //     xMin:function () {
  //       return buyArr[0].price;
  //     },
  //     xMax:function () {
  //       return sellArr[sellArr.length - 1].price;
  //     },
  //     yMin:function () {
  //       return 0;
  //     },
  //     yMax:function () {
  //      return  buyArr[0].amount * 1.2;
  //     }
  // },
  // watch:{
  //   values:{
  //     handler(newValue, oldValue) {
  //       this.initPoints();
  //       this.draw();
  //     },
  //     deep: true
  //   }
  // },
  mounted: function() {
    this.initCanvas();
    // this.initPoints();
    // this.draw();
  },
  methods: {
    initCanvas() {
      this.canvas = document.getElementById("depthGraph");
      this.canvas_cover = document.getElementById("depthGraph_cover");
      this.canvas.width = this.canvas_cover.width = this.pWidth = Math.floor(
        document.getElementById("depth_chart").offsetWidth
      );
      this.canvas.height = this.canvas_cover.height = this.pHeight = Math.floor(
        document.getElementById("depth_chart").offsetHeight
      );
      this.context = this.canvas.getContext("2d");
      this.context_cover = this.canvas_cover.getContext("2d");
      this.initHoverEvent();
    },
    initPoints() {
      var buy = this.values.bid,
        sell = this.values.ask;
      if (buy != undefined && sell != undefined) {
        this.xObj.min = buy.lowestPrice;
        this.xObj.max = sell.highestPrice;
        this.yObj.min = 0;
        let bidValue = null;
        this.values.bid.items.length >= 1 &&
          (bidValue = this.values.bid.items[this.values.bid.items.length - 1]
            .amount);
        this.values.bid.items.length == 0 && (bidValue = 0);
        let askValue = null;
        this.values.ask.items.length >= 1 &&
          (askValue = this.values.ask.items[this.values.ask.items.length - 1]
            .amount);
        this.values.ask.items.length == 0 && (askValue = 0);
        this.yObj.max = Math.max(bidValue, askValue);
        this.yObj.max = this.yObj.max * 1.2; //y轴最大值略大于最大委托量
        var buyArr = this.values.bid.items;

        var xScale = Math.floor((this.pWidth - 50) / (2 * (buyArr.length - 1))); //买入x轴等分
        var yScale = parseFloat(
          (this.pHeight - 50) / (this.yObj.max - this.yObj.min)
        ).toFixed(4);
        this.buyPoints = [];
        for (var i = buyArr.length - 1; i >= 0; i--) {
          var x = (buyArr.length - 1 - i) * xScale;
          var y = Math.floor(
            this.pHeight - 50 - (buyArr[i].amount - this.yObj.min) * yScale
          );

          this.buyPoints[this.buyPoints.length] = {
            x: x,
            y: y,
            amount: buyArr[i].amount,
            price: buyArr[i].price
          };
        }

        this.sellPoints = [];
        var sellArr = this.values.ask.items;
        sellArr.unshift({ price: this.values.bid.highestPrice, amount: 0 });

        xScale = Math.floor((this.pWidth - 50) / (2 * (sellArr.length - 1)));
        for (var i = 0; i < sellArr.length; i++) {
          var x = i * xScale + Math.floor((this.pWidth - 50) / 2);
          var y = Math.floor(
            this.pHeight - 50 - (sellArr[i].amount - this.yObj.min) * yScale
          );
          this.sellPoints[this.sellPoints.length] = {
            x: x,
            y: y,
            amount: sellArr[i].amount,
            price: sellArr[i].price
          };
        }
      }
    },
    convertTotal() {
      for (var i = 1; i < this.values.ask.items.length; i++) {
        this.values.ask.items[i].amount += this.values.ask.items[i - 1].amount;
      }

      for (var i = 1; i < this.values.bid.items.length; i++) {
        this.values.bid.items[i].amount += this.values.bid.items[i - 1].amount;
      }
    },
    draw(plate) {
      this.values = plate;
      if(this.values.bid.items.length > 100){
        this.values.bid.items = this.values.bid.items.slice(0, 100);
      }
      if(this.values.ask.items.length > 100){
        this.values.ask.items = this.values.ask.items.slice(0, 100);
      }
      this.convertTotal();
      this.initPoints();
      this.canvas.height = this.pHeight;
      this.drawAxis();
      this.drawBuy();
      this.drawSell();
      this.drawLabels();
    },
    drawBuy() {
      this.context.beginPath();
      this.context.moveTo(0, this.pHeight - 0);
      for (var i = 0; i < this.buyPoints.length; i++) {
        this.context.lineTo(this.buyPoints[i].x, this.buyPoints[i].y);
      }
      this.context.lineTo((this.pWidth - 50) / 2, this.pHeight - 0);
      if (this.values.skin === "day") {
        this.context.fillStyle = "rgba(0,178,117,.5)";
      } else {
        // 深色模式使用明显绿色
        this.context.fillStyle = "rgba(0, 255, 136, 0.6)";
      }
      this.context.fill();
    },
    drawSell() {
      this.context.beginPath();
      this.context.moveTo((this.pWidth - 50) / 2, this.pHeight - 0);
      for (var i = 0; i < this.sellPoints.length; i++) {
        this.context.lineTo(this.sellPoints[i].x, this.sellPoints[i].y);
      }
      this.context.lineTo(
        this.pWidth - 5,
        this.sellPoints[this.sellPoints.length - 1].y
      );
      this.context.lineTo(this.pWidth - 5, this.pHeight - 0);
      // 深色模式使用明显红色
      if (this.values.skin === "day") {
        this.context.fillStyle = "rgba(241,80,87,.5)";
      } else {
        this.context.fillStyle = "rgba(255, 95, 81, 0.6)";
      }
      this.context.fill();
    },
    drawAxis() {
      const centerX = (this.pWidth - 50) / 2;
      const axisColor = "#61688a";

      // 绘制中心垂直分界线
      this.context.beginPath();
      this.context.moveTo(centerX, 0);
      this.context.lineTo(centerX, this.pHeight - 25);
      this.context.strokeStyle = axisColor;
      this.context.setLineDash([5, 5]);
      this.context.stroke();
      this.context.setLineDash([]);
      this.context.closePath();

      // 绘制 X 轴水平线
      this.context.beginPath();
      this.context.moveTo(0, this.pHeight - 25);
      this.context.lineTo(this.pWidth - 50, this.pHeight - 25);
      this.context.strokeStyle = axisColor;
      this.context.stroke();
      this.context.closePath();

      // 绘制 Y 轴垂直线
      this.context.beginPath();
      this.context.moveTo(this.pWidth - 50, 0);
      this.context.lineTo(this.pWidth - 50, this.pHeight - 25);
      this.context.strokeStyle = axisColor;
      this.context.stroke();
      this.context.closePath();

      // X 轴价格刻度 - 买入侧
      this.context.fillStyle = axisColor;
      this.context.font = "11px Microsoft YaHei";
      const buyStep = Math.floor(this.buyPoints.length / 4);
      for (let i = 0; i < this.buyPoints.length; i += buyStep || 1) {
        const point = this.buyPoints[i];
        if (!point) continue;

        // 绘制刻度线
        this.context.beginPath();
        this.context.moveTo(point.x, this.pHeight - 25);
        this.context.lineTo(point.x, this.pHeight - 20);
        this.context.strokeStyle = axisColor;
        this.context.stroke();
        this.context.closePath();

        // 绘制价格标签
        const priceLabel = point.price?.toFixed(4) ?? "";
        this.context.fillText(priceLabel, point.x - 15, this.pHeight - 10);
      }

      // X 轴价格刻度 - 卖出侧
      const sellStep = Math.floor(this.sellPoints.length / 4);
      for (let i = 0; i < this.sellPoints.length; i += sellStep || 1) {
        const point = this.sellPoints[i];
        if (!point) continue;

        // 绘制刻度线
        this.context.beginPath();
        this.context.moveTo(point.x, this.pHeight - 25);
        this.context.lineTo(point.x, this.pHeight - 20);
        this.context.strokeStyle = axisColor;
        this.context.stroke();
        this.context.closePath();

        // 绘制价格标签
        const priceLabel = point.price?.toFixed(4) ?? "";
        this.context.fillText(priceLabel, point.x - 15, this.pHeight - 10);
      }

      // Y 轴数量刻度
      const yInterval = Math.floor((this.pHeight - 50) / 5);
      for (let i = 1; i <= 4; i++) {
        const yPos = this.pHeight - 25 - i * yInterval;
        const yValue = (this.yObj.max * i / 5) || 0;

        // 绘制刻度线
        this.context.beginPath();
        this.context.moveTo(this.pWidth - 50, yPos);
        this.context.lineTo(this.pWidth - 45, yPos);
        this.context.strokeStyle = axisColor;
        this.context.stroke();
        this.context.closePath();

        // 绘制数量标签
        this.context.fillStyle = axisColor;
        this.context.fillText(yValue.toFixed(2), this.pWidth - 48, yPos + 3);
      }
    },
    // 绘制买卖文字标签
    drawLabels() {
      const centerX = (this.pWidth - 50) / 2;

      // 买单标签 (左侧)
      this.context.fillStyle = "rgba(0, 255, 136, 0.9)";
      this.context.font = "bold 14px Microsoft YaHei";
      this.context.fillText("买单 Bid", 10, this.pHeight - 40);

      // 卖单标签 (右侧)
      this.context.fillStyle = "rgba(255, 95, 81, 0.9)";
      this.context.font = "bold 14px Microsoft YaHei";
      this.context.fillText("卖单 Ask", this.pWidth - 100, this.pHeight - 40);
    },
    initHoverEvent() {
      var opts = this;
      this.canvas_cover.addEventListener(
        "mousemove",
        function(e) {
          var arr = 0;
          if (e.offsetX <= (opts.pWidth - 50) / 2)
            arr = opts.getPoint(opts.buyPoints, e.offsetX);
          else arr = opts.getPoint(opts.sellPoints, e.offsetX);

          if (arr != null) {
            /*清空画板*/
            opts.canvas_cover.height = 500;
            /*画圆*/
            opts.context_cover.beginPath();
            opts.context_cover.arc(e.offsetX, arr.y, 10, 0, 2 * Math.PI);
            opts.context_cover.fillStyle = "#354067";
            opts.context_cover.fill();
            opts.context_cover.closePath();
            opts.context_cover.beginPath();
            opts.context_cover.arc(e.offsetX, arr.y, 5, 0, 2 * Math.PI);
            opts.context_cover.fillStyle = "#7a98f7";
            opts.context_cover.fill();
            opts.context_cover.closePath();

            /*画矩形框*/
            var rx = 0;
            if (e.offsetX <= 155) {
              rx = 5;
            } else {
              rx = e.offsetX - 150;
            }
            var ry = 0;
            if (arr.y <= 100) {
              ry = arr.y + 20;
            } else if (arr.y + 100 > opts.pHeight - 50) {
              ry = arr.y - 100;
            } else {
              ry = arr.y - 100;
            }
            opts.context_cover.beginPath();
            opts.context_cover.rect(rx, ry, 180, 80);
            opts.context_cover.fillStyle = "#262a42";
            opts.context_cover.fill();
            opts.context_cover.closePath();

            /*填充文本*/
            opts.context_cover.beginPath();
            opts.context_cover.fillStyle = "#ddd";
            opts.context_cover.font = "14px Microsoft YaHei";
            opts.context_cover.fillText("委托价", rx + 20, ry + 30);
            opts.context_cover.fillText(arr.price, rx + 70, ry + 30);
            opts.context_cover.fillText("累计", rx + 20, ry + 60);
            opts.context_cover.fillText(arr.amount.toFixed(8), rx + 70, ry + 60);
            opts.context_cover.closePath();
          }
        },
        false
      );
    },
    getPoint(arr, x) {
      var obj = null;
      for (var i = 0; i < arr.length; i++) {
        if (x == arr[i].x) {
          return arr[i];
        }
      }
      return obj;
    }
  }
};
</script>

<style scoped>
</style>
