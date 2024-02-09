<template>
  <div class="flex w-full p-4 gap-4">
    <div class="flex flex-col gap-4 w-3/10">
      <div class="grid grid-cols-1 gap-4">
        <div class="">
          <h2 class="text-xl font-bold mb-4">Index Price</h2>
          {{ indexPrice }}
        </div>
      </div>
      <div class="grid grid-cols-1 gap-4">
        <div class="">
          <h2 class="text-xl font-bold mb-4">Put to Call</h2>
          {{ putCallRatio }}
        </div>
      </div>
      <div class="grid grid-cols-1 gap-4">
        <div class="">
          <h2 class="text-xl font-bold mb-4">Call Volume</h2>
          {{ callVolume }}
        </div>
      </div>
      <div class="grid grid-cols-1 gap-4">
        <div class="">
          <h2 class="text-xl font-bold mb-4">Put Volume</h2>
          {{ putVolume }}
        </div>
      </div>
    </div>
    <div class="flex-1 w-7/10">
      <div class="flex justify-between">
        <button @click="toggleAudio"
          class="w-1/8 bg-white hover:bg-gray-100 text-gray-800 font-semibold py-2 px-4 border border-gray-400 rounded shadow">
          Audio
        </button>
        <div class="w-3/4"></div>
        <button @click="toggleSidebar"
          class="w-1/8 bg-white hover:bg-gray-100 text-gray-800 font-semibold py-2 px-4 border border-gray-400 rounded shadow">
          Filter
        </button>
      </div>
      
      <div class="grid grid-cols-1 gap-4">
        <div>
          <table class="min-w-full table-auto">
            <thead>
              <tr>
                <th class="px-4 py-2">TIME</th>
                <th class="px-4 py-2">TICK</th>
                <th class="px-4 py-2">DIR</th>
                <th class="px-4 py-2">C/P</th>
                <th class="px-4 py-2">EXPIRY</th>
                <th class="px-4 py-2">STRIKE</th>
                <th class="px-4 py-2">SPOT</th>
                <th class="px-4 py-2">PRICE</th>
                <th class="px-4 py-2">SIZE</th>
                <th class="px-4 py-2">PREM</th>
                <th class="px-4 py-2">IV</th>
              </tr>
            </thead>
            <tbody>
              <tr v-for="item in optionChain" :key="item.id">
                <td class="border px-4 py-2">{{ item.id }}</td>
                <td class="border px-4 py-2">{{ item.time }}</td>
                <td class="border px-4 py-2">{{ item.dir }}</td>
                <td class="border px-4 py-2">{{ item.cp }}</td>
                <td class="border px-4 py-2">{{ item.expiry }}</td>
                <td class="border px-4 py-2">{{ item.strike }}</td>
                <td class="border px-4 py-2">{{ item.spot }}</td>
                <td class="border px-4 py-2">{{ item.price }}</td>
                <td class="border px-4 py-2">{{ item.size }}</td>
                <td class="border px-4 py-2">{{ item.prem }}</td>
                <td class="border px-4 py-2">{{ item.iv }}</td>
              </tr>
            </tbody>
          </table>
        </div>
      </div>
    </div>
  </div>
  <div :class="{ 'translate-x-0': isSidebarOpen, 'translate-x-full': !isSidebarOpen }"
        class="absolute inset-y-0 right-0 transform transition-transform duration-300 ease-in-out w-64 bg-gray-100 shadow-xl">
        <Sidebar class="sidebar"/>
      </div>
</template>

<script>
import Sidebar from './Sidebar.vue';
export default {
  components: { Sidebar },
  data() {
    return {
      indexPrice: 100,
      putCallRatio: 100,
      callVolume: 100,
      putVolume: 100,
      optionChain: [
        { id: 1, time: "time", tick: "tick", dir: "dir", cp: "cp", expiry: "expiry", strike: "strike", spot: "spot", price: "price", size: "size", prem: "prem", iv: "iv" },
        { id: 2, time: "time", tick: "tick", dir: "dir", cp: "cp", expiry: "expiry", strike: "strike", spot: "spot", price: "price", size: "size", prem: "prem", iv: "iv" },
        { id: 3, time: "time", tick: "tick", dir: "dir", cp: "cp", expiry: "expiry", strike: "strike", spot: "spot", price: "price", size: "size", prem: "prem", iv: "iv" },
        { id: 4, time: "time", tick: "tick", dir: "dir", cp: "cp", expiry: "expiry", strike: "strike", spot: "spot", price: "price", size: "size", prem: "prem", iv: "iv" },
        { id: 5, time: "time", tick: "tick", dir: "dir", cp: "cp", expiry: "expiry", strike: "strike", spot: "spot", price: "price", size: "size", prem: "prem", iv: "iv" },
      ],
      isSidebarOpen: false,
      isAudioEnabled: false
    };
  },
  mounted() {
    // Initialize WebSocket and update grid data here
  },
  methods: {
    toggleSidebar() {
      this.isSidebarOpen = !this.isSidebarOpen;
    },
    toggleAudio() {
      this.isSidebarOpen = !this.isSidebarOpen;
    }
  },
};
</script>
