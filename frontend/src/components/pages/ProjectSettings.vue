<script lang="ts" setup>
import { reactive } from 'vue'
import {
  OpenDirectoryDialog,
  SetCurrentProject,
  UpdateProjectSettings,
  RemoveProject,
  CreateProject,
} from '@wails/go/main/App'

const props = defineProps<{
  appConfig: AppConfig,
  currentProject: Project,
}>()

const emit = defineEmits(['close'])

const data = reactive({
  showNotification: false
})

const form = reactive({
  name: '',
  path: '',
})

const setFormDefaults = () => {
  form.name = props.currentProject.Name
  form.path = props.currentProject.Path
}

setFormDefaults()

const setCurrentProject = (id: string) => {
  SetCurrentProject(id).then(() => setFormDefaults())
}

const selectProjectDir = (id: string) => {
  OpenDirectoryDialog().then((value: string) => {
    if (value.trim() == '') return
    form.path = value
  })
}

const deleteProject = (id: string) => {
  RemoveProject(id).then(() => {
    setFormDefaults()
    data.showNotification = true
  })
}

const createProject = () => {
  CreateProject().then(() => {})
}

const saveSettings = () => {
  UpdateProjectSettings(props.currentProject.Id, {
    ...props.currentProject,
    Path: form.path,
    Name: form.name,
  })
  emit('close')
}
</script>

<template>
  <main>
    <w-notification v-model="data.showNotification" :timeout="1000" success plain round shadow absolute>
      Project removed.
    </w-notification>

    <div class="title2">
      Project Manager
    </div>

    <div class="settings mt3">
      <div class="settings__projects">
        <div class="settings__new-project">
          <w-button class="ma1" bg-color="primary" lg @click="createProject">
            New
            <w-icon class="ml1">wi-plus</w-icon>
          </w-button>
        </div>
        <div v-for="project in appConfig.Projects"
             class="project" :class="currentProject?.Id === project.Id ? 'active' : ''"
             @click="() => setCurrentProject(project.Id)"
             :key="project.Id"
          >
          {{ project.Name }}
        </div>
      </div>

      <div class="settings__form">
        <w-input class="mb4 mt4" label="Name" placeholder="Unnamed" v-model="form.name" />

        <div class="path-selector">
          <w-input class="mb4" label="Path" placeholder="/path/to/project/" v-model="form.path" />
          <w-button class="ml5" color="primary" @click="() => selectProjectDir(currentProject?.Id)" outline lg>Select</w-button>
        </div>

        <div style="text-align: right;">
          <w-button v-if="Object.keys(appConfig.Projects).length > 1" class="ma1" bg-color="error" lg @click="() => deleteProject(currentProject.Id)">
            <w-icon class="mr1">wi-cross</w-icon>
            Delete
          </w-button>
          <w-button class="ma1" outline lg @click="$emit('close')">
            Cancel
          </w-button>
          <w-button class="ma1" bg-color="primary" lg @click="saveSettings">
            <w-icon class="mr1">wi-check</w-icon>
            Save
          </w-button>
        </div>
      </div>
    </div>
  </main>
</template>

<style scoped>
.settings {
  display: flex;
}
.settings__projects {
  flex-grow: 1;
  max-width: 15rem;
  overflow: hidden;
}

.settings__new-project {
  text-align: right;
}

.settings__form {
  padding: .5rem 1.5rem;
  flex-grow: 1;
  background: #fff;
}

.project {
  /* margin-bottom: .5rem; */
  padding: .7rem 1rem;
  cursor: pointer;
  /* border: solid 1px red; */
}

.project.active {
  background: #fff;
}

.project:hover {
  background: #fff;
}

.path-selector {
  display: flex;
}
</style>
