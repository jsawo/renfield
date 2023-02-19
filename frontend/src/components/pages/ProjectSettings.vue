<script lang="ts" setup>
import { reactive, computed } from 'vue'
import Button from '@/components/Button.vue'
import Input from '@/components/Input.vue'
import InputLabel from '@/components/InputLabel.vue'
import Badge from '@/components/Badge.vue'
import FolderIcon from '@/components/icons/FolderIcon.vue'
import PlusSmallIcon from '@/components/icons/PlusSmallIcon.vue'
import { notify } from "notiwind"
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
  tags: props.appConfig.Tags,
  selectedTag: null,
  commandPresets: [
    "php artisan tinker",
    "docker-compose exec -T php php artisan tinker",
  ]
})

const form = reactive({
  name: '',
  path: '',
  tag: '',
  command: '',
})

const setFormDefaults = () => {
  form.name = props.currentProject.Name
  form.path = props.currentProject.Path
  form.tag = props.currentProject.Tag
  form.command = props.currentProject.Command
  data.selectedTag = props.currentProject.Tag
}

Object.keys(props.appConfig.Projects).forEach((projectId) => {
  data.commandPresets.push(props.appConfig.Projects[projectId].Command)
})

data.commandPresets = [...new Set(data.commandPresets)]

setFormDefaults()

const sortedProjects = computed(function () {
  const keys = Object.keys(props.appConfig.Projects).sort((a, b) => {
    return props.appConfig.Projects[a].Name.localeCompare(props.appConfig.Projects[b].Name)
  })

  const sorted: Record<string, Project> = {}

  keys.forEach(key => {
    sorted[key] = props.appConfig.Projects[key]
  })

  return sorted
})

function getTag(tag: string): Tag {
  return props.appConfig.Tags.find((candidate) => candidate.Label === tag)
}

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
    notify({
      title: '',
      text: 'Project removed.',
      type: 'success',
      group: 'top-right',
    }, 3000)
  })
}

const createProject = () => {
  CreateProject().then(() => setFormDefaults())
}

const saveSettings = () => {
  UpdateProjectSettings(props.currentProject.Id, {
    ...props.currentProject,
    Path: form.path,
    Name: form.name,
    Tag: form.tag,
    Command: form.command,
  })
  emit('close')
}

const tagChanged = (event) => {
  form.tag = event.target.value
}

const commandPresetSelected = (event) => {
  form.command = event.target.value
}
</script>

<template>
  <main>
    <div class="flex mt-3">
      <div class="max-w-xs grow-1 overflow-hidden">
        <div class="text-right">
          <Button class="mt-2 mr-2 h-7 pl-1" @click="createProject">
            <PlusSmallIcon class="w-5 h-5 mr-1" />New
          </Button>
        </div>
        <div v-for="project in sortedProjects"
             class="cursor-pointer px-4 py-3" :class="currentProject?.Id === project.Id ? 'bg-white' : ''"
             @click="() => setCurrentProject(project.Id)"
             :key="project.Id"
          >
          {{ project.Name }}
          <Badge class="text-white ml-1" 
            :style="{ backgroundColor: getTag(project.Tag).Color }">
              {{ project.Tag }}
          </Badge> 
        </div>
      </div>
      <div class="bg-white grow p-4">
        <Input label="Name" v-model="form.name" />

        <div class="flex items-end mt-5">
          <Input label="Path" v-model="form.path" placeholder="/path/to/project/" class="grow" />
          <Button class="ml-3 h-9 w-28" @click="() => selectProjectDir(currentProject?.Id)">
            <FolderIcon class="mr-3 w-5" /> Select
          </Button>
        </div>

        <div class="flex items-end mt-5">
          <Input label="Command" v-model="form.command" class="grow" /> 

          <div class="ml-3">
            <select @change="commandPresetSelected($event)" placeholder="Preset"
              class="block w-28 max-w-lg rounded-md border-gray-300 shadow-sm h-9
                focus:border-indigo-500 focus:ring-indigo-500 sm:max-w-xs sm:text-sm"
            >
              <option></option>
              <option v-for="(preset, id) in data.commandPresets" :key="id">{{ preset }}</option>
            </select>
          </div>
        </div>

        <div class="mt-3">
          <InputLabel label="Label" />
          <select v-model="data.selectedTag" @change="tagChanged($event)"
            class="block w-28 max-w-lg rounded-md border-gray-300 shadow-sm h-9
              focus:border-indigo-500 focus:ring-indigo-500 sm:max-w-xs sm:text-sm"
          >
            <option v-for="(tag, id) in data.tags" :key="id" :value="tag.Label">{{ tag.Label }}</option>
          </select>
        </div>

        <div class="text-right mt-4 mb-2">
          <Button class="!bg-red-500 ml-5"
            @click="() => deleteProject(currentProject.Id)"
            v-if="Object.keys(appConfig.Projects).length > 1"
          >
            Delete
          </Button>

          <Button class="!bg-orange-500 ml-2" icon="pi pi-times" @click="$emit('close')">
            Cancel
          </Button>

          <Button class="p-button-primary p-button-sm ml-2" icon="pi pi-check" @click="saveSettings">
            Save
          </Button>
        </div>
      </div>
    </div>
  </main>
</template>

<style scoped>
</style>
