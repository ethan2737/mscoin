<template>
  <div>
    <el-carousel autoplay loop :interval="3000" arrow="always" id="carousel">
      <el-carousel-item v-for="item in 3" :key="item">
        <div class="demo-carousel">
          <img src="../../assets/img/banner.jpg" alt="">
        </div>
      </el-carousel-item>
    </el-carousel>
  </div>
</template>

<script setup>
/**
 * Vue 3 迁移 - OTC 轮播组件
 */
import { ElCarousel, ElCarouselItem } from 'element-plus'
</script>

<style scoped>
#carousel img {
  width: 100%;
  height: 500px;
  object-fit: cover;
}

.demo-carousel {
  width: 100%;
}

:deep(.el-carousel__arrow) {
  background-color: rgba(240, 167, 10, 0.8);
}

:deep(.el-carousel__arrow:hover) {
  background-color: rgba(240, 167, 10, 1);
}
</style>
