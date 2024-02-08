<template>
  <div class="flex w-full p-4 gap-4">
    <div class="flex flex-col gap-4 w-3/10">
      <div class="grid grid-cols-1 gap-4">
        <div class="bg-gray-200 p-4 shadow rounded-lg">
          <h2 class="text-xl font-bold mb-4">Index Price</h2>
          {{ indexPrice }}
        </div>
      </div>
      <div class="grid grid-cols-1 gap-4">
        <div class="bg-gray-200 p-4 shadow rounded-lg">
          <h2 class="text-xl font-bold mb-4">Put to Call</h2>
          {{ putCallRatio }}
        </div>
      </div>
      <div class="grid grid-cols-1 gap-4">
        <div class="bg-gray-200 p-4 shadow rounded-lg">
          <h2 class="text-xl font-bold mb-4">Call Volume</h2>
          {{ callVolume }}
        </div>
      </div>
      <div class="grid grid-cols-1 gap-4">
        <div class="bg-gray-200 p-4 shadow rounded-lg">
          <h2 class="text-xl font-bold mb-4">Put Volume</h2>
          {{ putVolume }}
        </div>
      </div>
    </div>

    <div class="flex-1 w-7/10">
      <div class="grid grid-cols-2 gap-4">
        <div>audio</div>
        <div>filter</div>
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
                <td class="border px-4 py-2">{{ item.time}}</td>
                <td class="border px-4 py-2">{{ item.dir}}</td>
                <td class="border px-4 py-2">{{ item.cp }}</td>
                <td class="border px-4 py-2">{{ item.expiry}}</td>
                <td class="border px-4 py-2">{{ item.strike }}</td>
                <td class="border px-4 py-2">{{ item.spot}}</td>
                <td class="border px-4 py-2">{{ item.price}}</td>
                <td class="border px-4 py-2">{{ item.size}}</td>
                <td class="border px-4 py-2">{{ item.prem}}</td>
                <td class="border px-4 py-2">{{ item.iv}}</td>
              </tr>
            </tbody>
          </table>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
export default {
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
    };
  },
  mounted() {
    // Initialize WebSocket and update grid data here
  },
  methods: {
      toggleSidebar() {
        this.isSidebarOpen = !this.isSidebarOpen;
      },
    },
};
</script>
