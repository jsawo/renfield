<script lang="ts" setup>
defineProps<{
  active: boolean,
}>();
</script>
<template>
  <div
    :class="[active 
      ? 'bg-stone-900 text-stone-200 border-stone-600' 
      : 'text-stone-300 border-transparent', 
      'px-3 py-2 cursor-pointer relative border-y-0 border-x border-transparent bg-stone-800 text-xs hover:border-stone-700'
    ]"
  >
    <slot />
  </div>
</template>
