<template>
  <div class="flex w-full p-4 gap-4">
    <div class="flex flex-col gap-4 w-3/10">
      <div class="grid grid-cols-1 gap-4">
        <div
          class="backdrop-filter backdrop-blur-sm bg-opacity-20 bg-gray-800 border border-gray-700 rounded-lg p-4 shadow-lg"
        >
          <h2 class="text-xl font-bold mb-4 text-white">Index Price</h2>
          <p class="text-gray-300 text-xl">{{ indexPrice }}</p>
        </div>
      </div>
      <div class="grid grid-cols-1 gap-4">
        <div
          class="backdrop-filter backdrop-blur-sm bg-opacity-20 bg-gray-800 border border-gray-700 rounded-lg p-4 shadow-lg"
        >
          <h2 class="text-xl font-bold mb-4 text-white">Put To Call</h2>
          <p class="text-gray-300 text-xl">{{ putCallRatio }}</p>
        </div>
      </div>
      <div class="grid grid-cols-1 gap-4">
        <div
          class="backdrop-filter backdrop-blur-sm bg-opacity-20 bg-gray-800 border border-gray-700 rounded-lg p-4 shadow-lg"
        >
          <h2 class="text-xl font-bold mb-4 text-white">Call Volume</h2>
          <p class="text-gray-300 text-xl">{{ callVolume }}</p>
        </div>
      </div>
      <div class="grid grid-cols-1 gap-4">
        <div
          class="backdrop-filter backdrop-blur-sm bg-opacity-20 bg-gray-800 border border-gray-700 rounded-lg p-4 shadow-lg"
        >
          <h2 class="text-xl font-bold mb-4 text-white">Put Volume</h2>
          <p class="text-gray-300 text-xl">{{ putVolume }}</p>
        </div>
      </div>
    </div>
    <div class="flex-1 w-7/10">
      <div class="flex gap-4 mb-2">
        <button
          @click="toggleSidebar"
          class="backdrop-filter backdrop-blur-sm bg-opacity-20 bg-gray-800 hover:bg-opacity-30 text-white font-semibold py-2 px-4 border border-gray-700 rounded shadow flex items-center justify-center"
        >
          <i class="fas fa-filter"></i>
        </button>
        <button
          @click="toggleAudio"
          class="backdrop-filter backdrop-blur-sm bg-opacity-20 bg-gray-800 hover:bg-opacity-30 text-white font-semibold py-2 px-4 border border-gray-700 rounded shadow flex items-center justify-center"
        >
          <i class="fas fa-volume-up"></i>
        </button>
      </div>

      <div class="grid grid-cols-1 gap-4">
        <div class="overflow-x-auto relative shadow-md sm:rounded-lg">
          <div
            class="backdrop-filter backdrop-blur-sm bg-opacity-20 bg-gray-800 border border-gray-700 rounded-lg overflow-hidden"
          >
            <table class="w-full text-sm text-left text-gray-300">
              <thead class="text-xs text-gray-500 uppercase bg-gray-800">
                <tr>
                  <th scope="col" class="py-3 px-6">TIME</th>
                  <th scope="col" class="py-3 px-6">DIR</th>
                  <th scope="col" class="py-3 px-6">C/P</th>
                  <th scope="col" class="py-3 px-6">EXPIRY</th>
                  <th scope="col" class="py-3 px-6">STRIKE</th>
                  <th scope="col" class="py-3 px-6">SPOT</th>
                  <th scope="col" class="py-3 px-6">PRICE</th>
                  <th scope="col" class="py-3 px-6">SIZE</th>
                  <th scope="col" class="py-3 px-6">PREM</th>
                  <th scope="col" class="py-3 px-6">IV</th>
                </tr>
              </thead>
              <tbody>
                <tr
                  v-for="item in optionChain"
                  :key="item.id"
                  :class="{
                    'glass-yellow-highlight': item.highlight,
                    'bg-gray-800': !item.highlight,
                    'border-b border-gray-700 hover:bg-gray-600': true,
                  }"
                  class="bg-gray-800 border-b border-gray-700 hover:bg-gray-600"
                >
                  <td class="py-4 px-6">
                    {{ item.time }}
                  </td>
                  <td
                    class="py-4 px-6"
                    :class="{
                      'text-green-500': item.dir === 'BUY',
                      'text-red-500': item.dir === 'SELL',
                    }"
                  >
                    {{ item.dir }}
                  </td>
                  <td
                    class="py-4 px-6"
                    :class="{
                      'text-blue-500': item.cp === 'CALL',
                      'text-orange-500': item.cp === 'PUT',
                    }"
                  >
                    {{ item.cp }}
                  </td>
                  <td class="py-4 px-6">
                    {{ item.expiry }}
                  </td>
                  <td class="py-4 px-6">
                    {{ item.strike }}
                  </td>
                  <td class="py-4 px-6">
                    {{ item.spot }}
                  </td>
                  <td class="py-4 px-6">
                    {{ item.price }}
                  </td>
                  <td class="py-4 px-6">
                    {{ item.size }}
                  </td>
                  <td class="py-4 px-6">
                    {{ item.prem }}
                  </td>
                  <td class="py-4 px-6">
                    {{ item.iv }}
                  </td>
                </tr>
              </tbody>
            </table>
          </div>
        </div>
      </div>
    </div>
  </div>
  <!-- Sidebar with transition for sliding effect -->
  <div
    :class="{
      'translate-x-full': !isSidebarOpen,
      'translate-x-0': isSidebarOpen,
    }"
    class="fixed inset-y-0 right-0 transform transition-transform duration-300 ease-in-out w-64 bg-gray-800 bg-opacity-20 backdrop-filter backdrop-blur-sm border-l border-gray-700 shadow-xl"
  >
    <Sidebar
      @closeSidebar="toggleSidebar"
      @filtersSelected="handleFiltersSelected"
    />
  </div>
</template>

<style>
.fade-slide-enter-active {
  transition: all 0.3s ease;
}
.fade-slide-enter, .fade-slide-leave-to /* .fade-slide-leave-active below version 2.1.8 */ {
  opacity: 0;
  transform: translateY(-20px);
}
.glass-yellow-highlight {
  background-color: rgba(
    253,
    253,
    150,
    0.2
  ); /* Soft yellow with transparency */
  backdrop-filter: blur(4px);
  border: 1px solid rgba(253, 253, 150, 0.4); /* Slightly more opaque border for definition */
}
</style>

<script>
import Sidebar from "./Sidebar.vue";
import webSocketService from "@/services/websocketService";
import audioService from "@/services/audioService";
import dateService from "../services/dateService";

function formatTimestamp(timestamp) {
  const date = new Date(timestamp);
  return new Intl.DateTimeFormat("en-US", {
    hour: "numeric",
    minute: "2-digit",
    second: "2-digit",
    hour12: true,
  }).format(date);
}

export default {
  components: { Sidebar },
  data() {
    return {
      indexPrice: 0,
      putCallRatio: 0,
      callVolume: 0,
      putVolume: 0,
      optionChain: [],
      isSidebarOpen: false,
      isAudioEnabled: false,
      filters: { currency: "Bitcoin" },
    };
  },
  created() {
    webSocketService.connect("ws://localhost:9000");
    webSocketService.addListener(this.handleMessage);
  },
  beforeUnmount() {
    webSocketService.removeListener(this.handleMessage);
  },
  methods: {
    handleMessage(message) {
      if (this.filters.currency == "Bitcoin" && message.index == "ETH") return;
      if (this.filters.currency == "Ethereum" && message.index == "BTC") return;
      switch (message.feedType) {
        case "INDEX_PRICE": {
          this.indexPrice = message.data.indexPrice;
          break;
        }
        case "VOLUME": {
          this.callVolume = message.data.callVolume;
          this.putVolume = message.data.putVolume;
          this.putCallRatio = message.data.putCallRatio;
          break;
        }
        case "OPTION": {
          console.log(message.data);
          const MIN_PREM_HIGHLIGHT = 1;
          const newItem = {
            ...message.data,
            time: formatTimestamp(message.data.time),
            highlight: message.data.prem >= MIN_PREM_HIGHLIGHT,
          };
          if (
            this.filters &&
            this.filters.minPremSize &&
            !(newItem.prem >= this.filters.minPremSize)
          )
            break;
          if (
            this.filters &&
            this.filters.minUnitSize &&
            !(newItem.size >= this.filters.minUnitSize)
          )
            break;
          if (
            this.filters &&
            this.filters.typeOption &&
            !(newItem.cp == this.filters.typeOption)
          )
            break;
          if (
            this.filters &&
            this.filters.dirOption &&
            !(newItem.dir == this.filters.dirOption)
          )
            break;
          if (this.filters && this.filters.expiry) {
            const dateString = message.data.expiry;
            const day = dateString.substring(0, 2); // Extract two characters for the day
            const month = dateString.substring(2, 5);
            const year = "20" + dateString.substring(5, 7);
            const date = new Date(`${month} ${day}, ${year}`);
            const timestamp = date.getTime();
            let flag = 1;
            switch (this.filters.expiry) {
              case "WEEK": {
                if (dateService.isInThisWeek(timestamp)) flag = 0;
                break;
              }
              case "MONTH": {
                if (dateService.isInThisMonth(timestamp)) flag = 0;
                break;
              }
              case "QUARTER": {
                if (dateService.isInThisQuarter(timestamp)) flag = 0;
                break;
              }
              case "HALF_YEAR": {
                if (dateService.isInThisHalfYear(timestamp)) flag = 0;
                break;
              }
              case "YEAR": {
                if (dateService.isInThisYear(timestamp)) flag = 0;
                break;
              }
            }
            if (flag) {
              break;
            }
          }
          this.optionChain.unshift(newItem);
          this.triggerHighlight(newItem);
          const MAX_OPTIONS = 30;
          if (this.optionChain.length > MAX_OPTIONS) {
            this.optionChain.pop();
          }
          if (this.isAudioEnabled) {
            audioService.handleOption(message.data);
          }
          break;
        }
      }
    },
    toggleSidebar() {
      this.isSidebarOpen = !this.isSidebarOpen;
    },
    toggleAudio() {
      this.isAudioEnabled = !this.isAudioEnabled;
    },
    handleFiltersSelected(filters) {
      this.filters = filters;
      this.optionChain = [];
    },
    triggerHighlight(item) {
      if (item.highlight) {
        setTimeout(() => {
          const itemIndex = this.optionChain.findIndex((i) => i.id === item.id);
          if (itemIndex !== -1) {
            this.optionChain[itemIndex].highlight = false;
          }
        }, 2000);
      }
    },
  },
};
</script>
