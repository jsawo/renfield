<script setup lang="ts">
import { onMounted, onBeforeUnmount } from 'vue';
import EditorTab from '@/components/EditorTab.vue'
import { NewTab, CloseTab, SetActiveTab, RenameTab } from '@wails/go/main/App'

const props = withDefaults(defineProps<{
  module: string,
  tabs?: Tab[],
  activeTab?: string,
}>(), {
  activeTab: "",
})

const emit = defineEmits<{
  (e: '@newTab'): void,
  (e: '@closeTab', tabId: string): void
  (e: '@selectTab', tabId: string): void
  (e: '@tabRenamed', tabId: string, newName: string): void
}>()

function closeTab(tabId: string) {
  CloseTab(props.module, tabId);
  emit('@closeTab', tabId)
}

function switchTab(tabId: string) {
  if(tabId === props.activeTab) return
  SetActiveTab(props.module, tabId);
  emit('@selectTab', tabId)
}

function newTab() {
  NewTab(props.module);
  emit('@newTab')
}

function tabRenamed(tabId: string, newName: string) {
  RenameTab(props.module, tabId, newName)
  emit('@tabRenamed', tabId, newName)
}

function handleKeyboardShortcuts(e: KeyboardEvent): void {
  if (e.ctrlKey && e.code === 'KeyW') {
    closeTab(props.activeTab)
  } else if (e.ctrlKey && e.code === 'KeyT') {
    newTab()
  } 
}

onMounted(() => {
  document.addEventListener('keydown', handleKeyboardShortcuts)
})

onBeforeUnmount(() => {
  document.removeEventListener('keydown', handleKeyboardShortcuts)
})
</script>

<template>
  <div class="flex bg-stone-800">
    <slot name="left" />
    <div class="flex grow">
      <EditorTab v-for="tab in tabs" 
        :active="tab.Id === activeTab" 
        @click.middle.prevent="() => closeTab(tab.Id)"
        @click.left.prevent="() => switchTab(tab.Id)"
        @@tabRenamed="(newName) => tabRenamed(tab.Id, newName)"
      >
        {{ tab.Name }}
      </EditorTab>
      <EditorTab :active="false" @click="newTab" >+</EditorTab>
    </div>
    <slot name="right" />
  </div>
</template>