<template>
  <svg :style="{ height: opts.height + 'px', width: opts.width + 'px' }">
    <polygon :fill="pColor" :points="polygonPoints"></polygon>
    <polyline
      fill="none"
      :points="polylinePoints"
      :stroke="sColor"
      :stroke-width="opts.strokeWidth"
      stroke-linecap="square"
    ></polyline>
  </svg>
</template>

<script setup>
import { ref, computed, onMounted } from 'vue'

const props = defineProps({
  values: {
    type: Array,
    required: true
  },
  width: {
    type: Number,
    required: false
  },
  height: {
    type: Number,
    required: false
  },
  rose: {
    type: String,
    required: false
  }
})

const coords = ref([])
const opts = ref({ strokeWidth: 1 })
const pColor = ref('#c6d9fd')
const sColor = ref('#4d89f9')

const x = ref(function(input) {
  return input * (opts.value.width / (props.values.length - 1))
})

const y = ref(function(input) {
  let y = opts.value.height - opts.value.strokeWidth
  const diff = opts.value.max - opts.value.min
  if (diff) {
    y -= ((input - opts.value.min) / diff) * (opts.value.height - opts.value.strokeWidth)
  }
  return y + opts.value.strokeWidth / 2
})

const polylinePoints = computed(() => {
  return coords.value.slice(2, coords.value.length - 2).join(' ')
})

const polygonPoints = computed(() => {
  return coords.value.join(' ')
})

const draw = () => {
  const currentOpts = opts.value
  const values = props.values

  if (values.length === 1) {
    values.push(values[0])
  }

  const max = Math.max.apply(Math, currentOpts.max === undefined ? values : values.concat(currentOpts.max))
  const min = Math.min.apply(Math, currentOpts.min === undefined ? values : values.concat(currentOpts.min))

  const strokeWidth = currentOpts.strokeWidth
  const width = currentOpts.width
  const height = currentOpts.height - strokeWidth
  const diff = max - min

  const xScale = (input) => {
    return input * (width / (values.length - 1))
  }

  const yScale = (input) => {
    let y = height
    if (diff) {
      y -= ((input - min) / diff) * height
    }
    return y + strokeWidth / 2
  }

  const zero = yScale(Math.max(min, 0))
  coords.value = [0, zero]

  for (let i = 0; i < values.length; i++) {
    coords.value.push(xScale(i), yScale(values[i]))
  }

  coords.value.push(width, zero)
}

onMounted(() => {
  opts.value.width = props.width || 120
  opts.value.height = props.height || 50
  opts.value.rose = props.rose || 0

  if (parseFloat(opts.value.rose) < 0) {
    pColor.value = '#f39494'
    sColor.value = '#e67f7f'
  } else {
    pColor.value = '#91baa7'
    sColor.value = '#66a488'
  }

  opts.value.max = undefined
  opts.value.min = undefined

  draw()
})
</script>
