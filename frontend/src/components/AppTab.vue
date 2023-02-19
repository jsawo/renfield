<script lang="ts" setup>
defineProps<{
  active: boolean,
}>();
</script>
<template>
  <div
    :class="[active 
      ? 'text-blue-600 border-b-2' 
      : '', 
      'px-3 py-3 font-medium cursor-pointer relative top-[3px] border-blue-400 hover:border-b-2'
    ]"
  >
    <slot />
  </div>
</template>
