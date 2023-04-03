<script lang="ts" setup>
import { ref } from 'vue'

defineProps<{
  active: boolean,
}>();

const emit = defineEmits<{
  (e: '@tabRenamed', name: string): void
}>()

const renameTab = () => {
  const newName = prompt('New name for tab')
  if(newName) {
    emit('@tabRenamed', newName)
  }
}

</script>
<template>
  <div
    @dblclick.prevent="() => renameTab()"
    :class="[active 
      ? 'bg-stone-900 text-white border-stone-600 underline' 
      : 'text-stone-300 border-transparent', 
      `px-3 py-2 cursor-pointer relative border-y-0 border-x border-transparent
      bg-stone-800 text-xs hover:border-stone-700 select-none`
    ]"
  >
    <slot />
  </div>
</template>
