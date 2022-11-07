<script lang="ts" setup>
import { reactive } from 'vue'
import Button from 'primevue/button'
import InputText from 'primevue/inputtext'
import Dropdown from 'primevue/dropdown'
import PTag from 'primevue/tag'
import { useToast } from "primevue/usetoast"
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
const toast = useToast();

const data = reactive({
  tags: props.appConfig.Tags,
  selectedTag: null,
})

const form = reactive({
  name: '',
  path: '',
  tag: '',
})

const setFormDefaults = () => {
  form.name = props.currentProject.Name
  form.path = props.currentProject.Path
  form.tag = props.currentProject.Tag
  data.selectedTag = props.appConfig.Tags.find((tag) => tag.Label === form.tag)
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
    toast.add({ severity: 'success', detail: 'Project removed.', life: 3000 });
  })
}

const createProject = () => {
  CreateProject()
}

const saveSettings = () => {
  UpdateProjectSettings(props.currentProject.Id, {
    ...props.currentProject,
    Path: form.path,
    Name: form.name,
    Tag: form.tag,
  })
  emit('close')
}

const tagChanged = (event) => {
  form.tag = event.value.Label
}
</script>

<template>
  <main>
    Project Manager

    <div class="flex mt-3">
      <div class="settings__projects">
        <div class="text-right">
          <Button label="New" class="p-button-sm" icon="pi pi-plus" @click="createProject" />
        </div>
        <div v-for="project in appConfig.Projects"
             class="project" :class="currentProject?.Id === project.Id ? 'active' : ''"
             @click="() => setCurrentProject(project.Id)"
             :key="project.Id"
          >
          {{ project.Name }}
          <PTag class="project-tag mx-2"
            :style="{ backgroundColor: props.appConfig.Tags.find((tag) => tag.Label === project.Tag).Color }"
            :value="project.Tag" />
        </div>
      </div>

      <div class="settings__form">
        <span class="p-float-label flex-grow-1 my-5">
          <InputText id="project_path" class="w-full" type="text" placeholder="Unnamed" v-model="form.name" />
          <label for="project_path">Name</label>
        </span>

        <div class="flex">
          <span class="p-float-label flex-grow-1">
            <InputText id="project_path" class="w-full" type="text" placeholder="/path/to/project/" v-model="form.path" />
            <label for="project_path">Path</label>
          </span>
          <Button label="Select" class="p-button-sm p-button-outlined ml-5" icon="pi pi-plus" @click="() => selectProjectDir(currentProject?.Id)" />
        </div>

        <div class="mt-4">
          <label for="tag_select" class="mr-2">Tag</label>
          <Dropdown id="tag_select" @change="tagChanged" v-model="data.selectedTag" :options="data.tags" optionLabel="Label" placeholder="Select Tag" />
        </div>

        <div class="text-right mt-4 mb-2">
          <Button label="Delete" class="p-button-danger p-button-sm ml-5" icon="pi pi-minus-circle"
            @click="() => deleteProject(currentProject.Id)"
            v-if="Object.keys(appConfig.Projects).length > 1"
          />

          <Button label="Cancel" class="p-button-warning p-button-sm ml-2" icon="pi pi-times"
            @click="$emit('close')"
          />

          <Button label="Save" class="p-button-primary p-button-sm ml-2" icon="pi pi-check"
            @click="saveSettings"
          />
        </div>
      </div>
    </div>
  </main>
</template>

<style scoped>
.settings__projects {
  flex-grow: 1;
  max-width: 15rem;
  overflow: hidden;
}

.settings__form {
  padding: .5rem 1.5rem;
  flex-grow: 1;
  background: #fff;
}

.project {
  padding: .7rem 1rem;
  cursor: pointer;
}

.project.active, .project:hover {
  background: #fff;
}
</style>
