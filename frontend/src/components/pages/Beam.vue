<script lang="ts" setup>
import { CopyToClipboard } from '@wails/go/main/App'
import { useToast } from "primevue/usetoast"
import Button from 'primevue/button'
import Card from 'primevue/card'

defineProps<{
  messages: Array<BeamMessage>,
}>()

const port = '3333'
const toast = useToast();

function copyLaravelString() {
  const data = `Http::post("http://localhost:${port}/beam", ["payload"=> json_encode(["foo" => "bar"])]);`
  CopyToClipboard(data)
  toast.add({ severity: 'info', detail: 'Copied to clipboard.', life: 3000 });
}

function copyCurlString() {
  const data = `curl -X POST -d '{"payload": "{ \\\"foo\\\": \\\"bar\\\" }"}' http://localhost:${port}/beam`
  CopyToClipboard(data)
  toast.add({ severity: 'info', detail: 'Copied to clipboard.', life: 3000 });
}

function getPayload(payload: string): string {
  try {
    return JSON.stringify(JSON.parse(payload), null, 2)
  } catch (e) {
    return payload
  }
}
</script>

<template>
  <main>
    <div>
      Beam (running on port :{{ port }})
      <Button label="Laravel" @click="copyLaravelString" class="mr-2 p-button-sm p-button-outlined" icon="pi pi-copy" iconPos="right" />
      <Button label="CURL" @click="copyCurlString" class="p-button-sm p-button-outlined" icon="pi pi-copy" iconPos="right" />

      <div style="float: right;">
        <Button label="Clear" @click="$emit('clearBeamMessages')" class="p-button-sm" icon="pi pi-delete-left" iconPos="right" />
      </div>
    </div>

    <Card id="empty-state" v-if="messages.length === 0">
      <template #content>
        <div>No messages received yet…</div>
        <i class="large-icon pi pi-times-circle"></i>
      </template>
    </Card>

    <ul class="log_box">
      <li v-for="(msg, index) in messages" :key="index">
        <Card class="mt-2">
          <template #content>
            <div class="message">
              <div class="message__time">{{ msg.timestamp }}</div>
              <pre class="message__payload">{{ getPayload(msg.payload) }}</pre>
            </div>
          </template>
        </Card>
      </li>
    </ul>
  </main>
</template>

<style scoped>
main {
  display: flex;
  height: auto;
  flex-direction: column;
}

#empty-state {
  text-align: center;
  margin-top: 2rem;
}

.large-icon {
  margin-top: 2rem;
  margin-bottom: 1rem;
  font-size: 5rem;
}

.message {
  display: flex;
  flex-direction: row;
}

.message__time {
  width: 5rem;
  flex-shrink: 0;
}

.message__payload {
  width: 100%;
}

.log_box {
  list-style: none;
  padding: 0;
}

.log_box li:first-child .p-card {
  border: solid 1px #999;
  background-color: white;
}
</style>
