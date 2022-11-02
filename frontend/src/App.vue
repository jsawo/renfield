<script lang="ts" setup>
import 'wave-ui/dist/wave-ui.css'
import dayjs from 'dayjs'
import timezone from 'dayjs/plugin/timezone'
import utc from 'dayjs/plugin/utc'
import { reactive, onMounted, computed } from 'vue'
import { EventsOn } from '@wails/runtime'
import { GetConfig } from '@wails/go/main/App'
import JsonFormatter from '@/components/pages/JsonFormatter.vue'
import Beam from '@/components/pages/Beam.vue'
import Tinker from '@/components/pages/Tinker.vue'
import ProjectSettings from '@/components/pages/ProjectSettings.vue'
import '@mdi/font/css/materialdesignicons.min.css'

dayjs.extend(utc)
dayjs.extend(timezone)
dayjs.locale('pl')
dayjs.tz.setDefault("Europe/Warsaw")

const tabs = [
  { id: "tinker", title: 'Tinker', content: Tinker },
  { id: "jsonformatter", title: 'JSON Formatter', content: JsonFormatter },
  { id: "beam", title: 'Beam', content: Beam },
]

const data = reactive({
  currentSection: "app", // app / projectSettings
  tabs: tabs,
  value: tabs[0].title,
  messages: [] as Array<BeamMessage>,
  appConfig: {} as AppConfig,
  showNotification: false,
})

EventsOn("beamMessage", function (messageData: RawBeamMessage) {
  data.messages.unshift({
    timestamp: dayjs().toDate().toLocaleTimeString(),
    payload: messageData.Payload
  })
})

EventsOn("configUpdated", function (appConfig: AppConfig) {
  console.log("EVENT: received configUpdated message - updating appConfig")
  data.appConfig = appConfig
})


const clearMessages = () => data.messages = []

const currentProject = computed<Project>(function () {
  return data.appConfig.Projects
    ? data.appConfig.Projects[data.appConfig.Currentproject]
    : {}
})

const currentColor = computed<string>(function() {
  return getTagColor(currentProject.value?.Tag)
})

const getTagColor = (tag: string): string => {
  const foundTag = data.appConfig.Tags?.find(function(tagConfig) {
    if (tag === tagConfig.Label) {
      return tagConfig.Color
    }
  })

  return foundTag ? foundTag.Color : "black"
}

const openProjectManager = () => {
  data.currentSection = data.currentSection === 'projectSettings' ? 'app' : 'projectSettings'
}

const refreshAppConfig = () => {
  GetConfig().then((config) => {
    data.appConfig = config
  })
}

onMounted(() => refreshAppConfig())
</script>

<template>
  <w-app style="height: 100%;">
    <w-notification v-model="data.showNotification" :timeout="1000" success plain round shadow absolute>
      Project saved.
    </w-notification>

    <div class="main-stack">
      <w-toolbar shadow>
        <div class="app-title">Renfield <div class="triangle"></div></div>
        <div class="project-name" @click="openProjectManager">
          {{ currentProject?.Name }}
          <w-tag class="ml2 mr2" outline bg-color="white" :color="currentColor">{{ currentProject?.Tag }}</w-tag>
        </div>
        <div class="spacer"></div>
        <w-icon class="mr1" xl>mdi mdi-cog</w-icon>
      </w-toolbar>

      <w-tabs v-if="data.currentSection === 'app'" :items="tabs" transition="none" style="flex-grow: 1;" card>
        <template #item-content="{ item }">
          <div class="component-wrapper">
            <Beam v-if="item.id === 'beam'" :messages="data.messages" class="component" @clear-beam-messages="clearMessages"/>
            <component v-else class="component" :is="item.content" />
          </div>
        </template>
      </w-tabs>
      <div v-else-if="data.currentSection === 'projectSettings'" class="component-wrapper ml3 mr3 mt3 mb3">
        <ProjectSettings
          :app-config="data.appConfig"
          :current-project="currentProject"
          @close="data.currentSection = 'app'"
          />
      </div>
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

.app-title {
  font-size: 1.3rem;
  background: rgb(37, 37, 37);
  color: #eb5f5a;
  margin: -.5rem 0 -.5rem -.8rem;
  padding: .6rem 1rem .6rem 1.5rem;
  position: relative;
  -webkit-touch-callout: none;
  -webkit-user-select: none;
  -khtml-user-select: none;
  -moz-user-select: none;
  -ms-user-select: none;
  user-select: none;
  cursor: default;
}

.app-title::after {
  content: "";
  position: absolute;
  display: block;
  right: -20px;
  top: 0;
  border-top: 50px solid rgb(37, 37, 37);
  border-right: 20px solid transparent;
}

.project-name {
  font-size: 1.3rem;
  margin: -.5rem 1rem -.5rem 0rem;
  padding: .6rem .5rem .6rem 2.5rem;
  cursor: pointer;
}

.project-name:hover {
  background: #eee;
}
</style>
