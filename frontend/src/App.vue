<script lang="ts" setup>
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
import Toast from 'primevue/toast'
import Toolbar from 'primevue/toolbar'
import TabView from 'primevue/tabview'
import TabPanel from 'primevue/tabpanel'
import Tag from 'primevue/tag'
import 'primevue/resources/themes/saga-blue/theme.css'
import 'primevue/resources/primevue.min.css'
import 'primeflex/primeflex.css'
import 'primeicons/primeicons.css'
import './style.css'

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
  <div class="h-full">
    <Toast />

    <div class="main-stack h-full flex flex-column">
      <Toolbar class="toolbar">
        <template #start>
          <div class="app-title">Renfield <div class="triangle"></div></div>
          <div class="project-name" @click="openProjectManager">
            {{ currentProject?.Name }}
            <Tag class="project-tag mx-2" :style="{backgroundColor: currentColor}" :value="currentProject?.Tag"></Tag>
          </div>
        </template>
        <template #end>
          <!-- <i class="pi pi-cog text-3xl"></i> -->
        </template>
      </Toolbar>

      <TabView v-if="data.currentSection === 'app'">
        <TabPanel v-for="tab in tabs" :header="tab.title" :key="tab.id">
          <div class="component-wrapper h-full">
            <Beam v-if="tab.id === 'beam'" :messages="data.messages" class="component" @clear-beam-messages="clearMessages" />
            <component v-else class="component" :is="tab.content" />
          </div>
        </TabPanel>
      </TabView>
      <div v-else-if="data.currentSection === 'projectSettings'" class="component-wrapper pb-5 h-full m-3">
        <ProjectSettings
          :app-config="data.appConfig"
          :current-project="currentProject"
          @close="data.currentSection = 'app'"
          />
      </div>
    </div>
  </div>
</template>

<style scoped>
.toolbar {
  padding: 0rem 1rem 0rem 0;
  box-shadow: 0 0 3px 0 black;
  z-index: 10;
}

.app-title {
  font-size: 1.3rem;
  background: rgb(37, 37, 37);
  color: #eb5f5a;
  padding: 1rem 1rem .6rem 1.5rem;
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
  z-index: 100;
}

.project-name {
  font-size: 1.3rem;
  margin: -.5rem 1rem -.5rem 0rem;
  padding: 1rem 1.5rem .6rem 2.5rem;
  cursor: pointer;
  position: relative;
}

.project-name:hover {
  background: #eee;
}

.project-name .project-tag {
  top: -3px;
  position: relative;
}
</style>
