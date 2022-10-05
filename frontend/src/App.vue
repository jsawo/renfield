<script lang="ts" setup>
import 'wave-ui/dist/wave-ui.css'
import { reactive, ref } from 'vue'
import { EventsOn } from '@wails/runtime'
import JsonFormatter from '@/components/pages/JsonFormatter.vue'
import Beam from '@/components/pages/Beam.vue'
import Tinker from '@/components/pages/Tinker.vue'
import '@mdi/font/css/materialdesignicons.min.css'

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
    <div class="main-stack">
      <w-toolbar shadow>
        <div class="title2">Renfield</div>
        <div class="spacer"></div>
        <w-icon class="mr1" xl>mdi mdi-cog</w-icon>
      </w-toolbar>

      <w-tabs :items="tabs" transition="none" style="flex-grow: 1;" card>
        <template #item-content="{ item }">
          <div class="component-wrapper">
            <Beam v-if="item.id === 'beam'" :messages="messages" class="component" />
            <component v-else class="component" :is="item.content" />
          </div>
        </template>
      </w-tabs>
    </div>
  </w-app>
</template>

<style scoped>
.main-stack {
  height: 100%;
  display: flex;
  flex-direction: column;
}

.component-wrapper {
  height: 100%;
  padding: 0 0 2em 0;
}
</style>
