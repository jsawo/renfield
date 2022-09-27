<script lang="ts" setup>
import { computed } from '@vue/reactivity';
import { reactive, ref } from 'vue'
import { EventsOn } from '../wailsjs/runtime'
import JsonFormatter from './components/JsonFormatter.vue'
import Beam from './components/Beam.vue'

const tabs = [
  { id: "beam", icon: 'lightbulb', title: 'Beam', content: Beam },
  { id: "jsonformatter", icon: 'auto_awesome', title: 'JSON Formatter', content: JsonFormatter },
]

const data = reactive({
  tabs: tabs,
  value: tabs[0].title,
})

const messages = ref([])

EventsOn("beamMessage", function (messageData: BeamMessage) {
  console.log("GOT AN EVENT", messageData)
  // @ts-ignore
  messages.value.unshift(messageData)
})

const currentTab = computed(() => tabs.find(({ title }) => title === data.value) || tabs[0])
</script>

<template>
  <div class="main_stack">
    <va-tabs class="tabs" v-model="data.value">
      <template #tabs>
        <va-tab v-for="tab in data.tabs" :key="tab.id" :name="tab.title" class="px-2 py-2">
          <va-icon :name="tab.icon" size="small" class="mr-2" />
          {{tab.title}}
        </va-tab>
      </template>
    </va-tabs>

    <Beam v-if="currentTab.id === 'beam'" :messages="messages" class="component" />
    <component v-else class="component" :is="currentTab.content" />
  </div>
</template>

<style scoped>
.main_stack {
  display: flex;
  flex-direction: column;
  height: 100%;
}

.tabs {
  flex-grow: 0;
  font-size: 1.5rem;
  padding-left: 1rem;
}

.component {
  flex-grow: 1;
}
</style>
