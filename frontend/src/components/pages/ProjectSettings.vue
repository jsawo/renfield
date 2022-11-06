<script lang="ts" setup>
import { reactive } from 'vue'
import Button from 'primevue/button'
import InputText from 'primevue/inputtext'
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
    toast.add({ severity: 'success', detail: 'Project removed.', life: 3000 });
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
    <div class="title2">
      Project Manager
    </div>

    <div class="settings mt3">
      <div class="settings__projects">
        <div class="settings__new-project">
          <Button label="New" class="p-button-sm" icon="pi pi-plus" @click="createProject" />
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
</style>
