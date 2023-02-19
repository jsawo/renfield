<script setup lang="ts">
defineProps<{
  label: string
}>();
</script>

<template>
  <label class="block text-sm font-medium text-gray-700">
    {{ label }}
    <slot />
  </label>  
</template>