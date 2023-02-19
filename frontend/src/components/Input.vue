<script setup lang="ts">
import InputLabel from '@/components/InputLabel.vue'

defineProps<{
  label: string
  modelValue: string | number
}>()

defineEmits(['update:modelValue'])
</script>

<template>
  <InputLabel :label="label">
    <div class="mt-1">
      <input v-bind="$attrs" :value="modelValue" 
        @input="$emit('update:modelValue', ($event.target as HTMLInputElement).value)"
        class="block w-full rounded-md 
          border border-gray-300 shadow-sm sm:text-sm p-2
          focus:border-indigo-500 focus:ring-indigo-500" 
      />
    </div>  
  </InputLabel>  
</template>