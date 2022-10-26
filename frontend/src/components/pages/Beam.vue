<script lang="ts" setup>
import { ref } from 'vue'
import { CopyToClipboard } from '@wails/go/main/App'

defineProps<{
  messages: Array<BeamMessage>,
}>()

const port = '3333'

function copyLaravelString() {
  const data = `Http::post("http://localhost:${port}/beam", ["payload"=> json_encode(["foo" => "bar"])]);`
  CopyToClipboard(data)
  showNotification.value = true
}

function copyCurlString() {
  const data = `curl -X POST -d '{"payload": "{ \\\"foo\\\": \\\"bar\\\" }"}' http://localhost:${port}/beam`
  CopyToClipboard(data)
  showNotification.value = true
}

function getPayload(payload: string): string {
  try {
    return JSON.stringify(JSON.parse(payload), null, 2)
  } catch (e) {
    return payload
  }
}

const showNotification = ref(false)
</script>

<template>
  <main>
    <w-notification v-model="showNotification" :timeout="1000" success plain round shadow absolute>
      Copied to clipboard.
    </w-notification>

    <div>
      Beam (running on port :{{ port }})
      <w-button class="ml2" outline @click="copyLaravelString">
        Laravel
        <w-icon sm class="ml1" xl>mdi mdi-clipboard-multiple-outline</w-icon>
      </w-button>
      <w-button class="ml2" outline @click="copyCurlString">
        CURL
        <w-icon sm class="ml1" xl>mdi mdi-clipboard-multiple-outline</w-icon>
      </w-button>

      <div style="float: right;">
        <w-button outline @click="$emit('clearBeamMessages')">
          <w-icon sm class="mr1" xl>mdi mdi-delete-outline</w-icon>
          Clear
        </w-button>
      </div>
    </div>

    <w-card class="mt2" id="empty-state" v-if="messages.length === 0">
      <div>No messages received yet…</div>
      <w-icon class="large-icon" xl>mdi mdi-message-minus-outline</w-icon>
    </w-card>

    <ul class="log_box">
      <li v-for="(msg, index) in messages" :key="index">
        <w-card class="mt2">
          <div class="message">
            <div class="message__time">{{ msg.timestamp }}</div>
            <pre class="message__payload">{{ getPayload(msg.payload) }}</pre>
          </div>
        </w-card>
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

.log_box li:first-child .w-card {
  border: solid 1px #999;
}
</style>
