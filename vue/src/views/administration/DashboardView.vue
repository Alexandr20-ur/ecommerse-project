<script setup>
import { onMounted, ref } from 'vue'
import Chart from 'chart.js/auto'

const chartCanvas = ref(null)

onMounted(() => {
  new Chart(chartCanvas.value, {
    type: 'bar',
    data: {
      labels: ['Yan', 'Fev', 'Mar', 'Apr', 'May', 'Iyun', 'Iyu', 'Avg', 'Sen', 'Okt', 'Noy', 'Dek'],
      datasets: [{ label: 'Sotuv', data: [12, 19, 18, 17, 15, 10, 12, 2, 3, 8, 19, 3], borderWidth: 1 }]
    },
    options: { plugins: { legend: { display: false } }, scales: { y: { beginAtZero: true } } }
  })
})
</script>

<template>
  <div>
    <!-- Секция с кнопками Export data -->
    <section class="d-flex gap-2 mb-5">...</section>

    <!-- Секция с карточками (Total orders, Gross Sale...) -->
    <section class="mb-5 d-grid ...">...</section>

    <!-- Секция с графиком -->
    <article class="card p-5 shadow-sm h-full">
      <canvas ref="chartCanvas" style="display: block; max-height:350px"></canvas>
    </article>

    <!-- Таблица Latest orders -->
    <section>...</section>
  </div>
</template>