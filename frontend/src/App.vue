<script lang="ts" setup>
import 'wave-ui/dist/wave-ui.css'
import { reactive, ref } from 'vue'
import { EventsOn } from '../wailsjs/runtime'
import JsonFormatter from './components/pages/JsonFormatter.vue'
import Beam from './components/pages/Beam.vue'
import Tinker from './components/pages/Tinker.vue'

const tabs = [
  { id: "tinker", title: 'Tinker', content: Tinker },
  { id: "jsonformatter", title: 'JSON Formatter', content: JsonFormatter },
  { id: "beam", title: 'Beam', content: Beam },
]

const data = reactive({
  tabs: tabs,
  value: tabs[0].title,
})

const messages = ref([])

EventsOn("beamMessage", function (messageData: BeamMessage) {
  // @ts-ignore
  messages.value.unshift(messageData)
})
</script>

<template>
  <w-app style="height: 100%;">
    <div class="main_stack">
      <w-tabs :items="tabs" transition="none" card>
        <template #item-content="{ item }">
          <Beam v-if="item.id === 'beam'" :messages="messages" class="component" />
          <component v-else class="component" :is="item.content" />
        </template>
      </w-tabs>
    </div>
  </w-app>
</template>

<style scoped>
.main_stack {
  height: 100%;
}
</style>
