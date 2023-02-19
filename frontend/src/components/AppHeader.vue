<script lang="ts" setup>
import Badge from '@/components/Badge.vue'

defineProps<{
  currentProject?: Project,
  openProjectManager: () => void,
  badgeColor?: string,
}>()

</script>
<template>
  <div class="flex w-full border-b border-stone-300 select-none">
    <div class="app-title | text-xl pl-6 pr-4 py-3 relative text-brand-red bg-brand-black cursor-default select-none">
      Renfield
    </div>
    <div class="pl-10 pr-5 rounded-r-full flex items-center cursor-pointer relative h-auto hover:bg-white"
      @click="openProjectManager"
    >
      <span class="text-2xl mr-4">{{ currentProject?.Name }}</span>
      <Badge class="text-white" :style="{ backgroundColor: badgeColor }">{{ currentProject?.Tag }}</Badge>
    </div>
  </div>
</template>

<style scoped>
.app-title::after {
  content: "";
  position: absolute;
  display: block;
  right: -20px;
  top: 0;
  border-top: 50px solid var(--brand-black);
  border-right: 20px solid transparent;
  z-index: 100;
}
</style>