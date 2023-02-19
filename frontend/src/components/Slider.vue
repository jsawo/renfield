<script lang="ts" setup>

defineProps<{
  min: number,
  max: number,
}>()

</script>
<template>
  <input type="range" :min="min" :max="max"
    class="w-full h-2 bg-blue-400 accent-blue-500 rounded-lg appearance-none cursor-pointer dark:bg-gray-700"
    >
</template>

