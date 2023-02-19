<script setup lang="ts">
import ExclamationDiscIcon from './icons/ExclamationDiscIcon.vue'
</script>
<template>
  <NotificationGroup group="top-right">
    <div class="fixed inset-0 flex items-start justify-end p-6 px-4 py-6 pointer-events-none">
      <div class="w-full max-w-sm">
        <Notification v-slot="{ notifications }"
          enter="transform ease-out duration-300 transition"
          enter-from="translate-y-2 opacity-0 sm:translate-y-0 sm:translate-x-4"
          enter-to="translate-y-0 opacity-100 sm:translate-x-0"
          leave="transition ease-in duration-500"
          leave-from="opacity-100"
          leave-to="opacity-0"
          move="transition duration-500"
          move-delay="delay-300"
        >
          <div v-for="notification in notifications" :key="notification.id">
            <div class="w-full max-w-sm mx-auto mt-4 overflow-hidden bg-white rounded-lg shadow-md">
              <div v-if="notification.type === 'info'" class="flex">
                <div class="flex items-center justify-center w-12 bg-blue-500">
                  <ExclamationDiscIcon />
                </div>
                <div class="px-4 py-2 -mx-3">
                  <div class="mx-3">
                    <span class="font-semibold text-blue-500">{{ notification.title }}</span>
                    <p class="text-sm text-gray-600">{{ notification.text }}</p>
                  </div>
                </div>
              </div>

              <div v-if="notification.type === 'warning'" class="flex">
                <div class="flex items-center justify-center w-12 bg-yellow-500">
                  <ExclamationDiscIcon />
                </div>
                <div class="px-4 py-2 -mx-3">
                  <div class="mx-3">
                    <span class="font-semibold text-yellow-500">{{ notification.title }}</span>
                    <p class="text-sm text-gray-600">{{ notification.text }}</p>
                  </div>
                </div>
              </div>

              <div v-if="notification.type === 'error'" class="flex">
                <div class="flex items-center justify-center w-12 bg-red-500">
                  <ExclamationDiscIcon />
                </div>
                <div class="px-4 py-2 -mx-3">
                  <div class="mx-3">
                    <span class="font-semibold text-red-500">{{ notification.title }}</span>
                    <p class="text-sm text-gray-600">{{ notification.text }}</p>
                  </div>
                </div>
              </div>

              <div v-if="notification.type === 'success'" class="flex">
                <div class="flex items-center justify-center w-12 bg-green-500">
                  <ExclamationDiscIcon />
                </div>
                <div class="px-4 py-2 -mx-3">
                  <div class="mx-3">
                    <span class="font-semibold text-green-500">{{ notification.title }}</span>
                    <p class="text-sm text-gray-600">{{ notification.text }}</p>
                  </div>
                </div>
              </div>
            </div>
          </div>
        </Notification>
      </div>
    </div>
  </NotificationGroup>
</template>