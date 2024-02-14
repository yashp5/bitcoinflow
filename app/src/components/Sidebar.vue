<template>
  <aside
    class="min-h-screen p-4 space-y-6 bg-gray-800 bg-opacity-40 backdrop-filter backdrop-blur-lg rounded-lg shadow-lg"
  >
    <div class="space-y-2">
      <label for="currency" class="block text-gray-300">Currency</label>
      <select
        id="currency"
        class="w-full p-2 bg-gray-700 text-white rounded-md"
      >
        <option>Bitcoin</option>
        <option>Ethereum</option>
      </select>
    </div>

    <div class="space-y-2">
      <label class="block text-gray-300">Minimum Premium Size</label>
      <div class="flex flex-wrap gap-2">
        <button
          @click="selectMinPremSize(0.1)"
          :class="{'bg-blue-500': selectedValue === '0.1 BTC'}"
          class="p-2 flex-1 bg-gray-700 hover:bg-gray-600 text-white font-semibold py-2 px-4 rounded shadow"
        >
          0.1 BTC
        </button>
        <button
          class="p-2 flex-1 bg-gray-700 hover:bg-gray-600 text-white font-semibold py-2 px-4 rounded shadow"
        >
          0.5 BTC
        </button>
        <button
          class="p-2 flex-1 bg-gray-700 hover:bg-gray-600 text-white font-semibold py-2 px-4 rounded shadow"
        >
          1 BTC
        </button>
        <button
          class="p-2 flex-1 bg-gray-700 hover:bg-gray-600 text-white font-semibold py-2 px-4 rounded shadow"
        >
          5 BTC
        </button>
        <button
          class="p-2 flex-1 bg-gray-700 hover:bg-gray-600 text-white font-semibold py-2 px-4 rounded shadow"
        >
          10 BTC
        </button>
        <button
          class="p-2 flex-1 bg-gray-700 hover:bg-gray-600 text-white font-semibold py-2 px-4 rounded shadow"
        >
          25 BTC
        </button>
      </div>
    </div>

    <div class="space-y-2">
      <label for="unit-size" class="block text-gray-300"
        >Minimum Unit Size</label
      >
      <input
        id="unit-size"
        type="text"
        class="w-full p-2 bg-gray-700 text-white rounded-md"
      />
    </div>

    <div class="space-y-2">
      <label for="unit-size" class="block text-gray-300">Call or Put</label>
      <div class="flex flex-wrap gap-2">
        <button
          class="p-2 flex-1 bg-gray-700 hover:bg-gray-600 text-white font-semibold py-2 px-4 rounded shadow"
        >
          Call
        </button>
        <button
          class="p-2 flex-1 bg-gray-700 hover:bg-gray-600 text-white font-semibold py-2 px-4 rounded shadow"
        >
          Put
        </button>
      </div>
    </div>

    <div class="space-y-2">
      <label for="unit-size" class="block text-gray-300">Side</label>
      <div class="flex flex-wrap gap-2">
        <button
          class="p-2 flex-1 bg-gray-700 hover:bg-gray-600 text-white font-semibold py-2 px-4 rounded shadow"
        >
          Buy
        </button>
        <button
          class="p-2 flex-1 bg-gray-700 hover:bg-gray-600 text-white font-semibold py-2 px-4 rounded shadow"
        >
          Sell
        </button>
      </div>
    </div>

    <div class="flex gap-2">
      <button
        class="p-2 flex-1 bg-gray-700 hover:bg-gray-600 text-white font-semibold py-2 px-4 rounded shadow"
      >
        Clear
      </button>
      <button
        class="p-2 flex-1 bg-gray-700 hover:bg-gray-600 text-white font-semibold py-2 px-4 rounded shadow"
      >
        Go
      </button>
    </div>
  </aside>
</template>

<script>
export default {
  data() {
    return {
      minPremSize: null,
    };
  },
  methods: {
    selectMinPremSize(value) {
      this.minPremSize = value;
      this.$emit('minPremSizeSelected', this.minPremSize);
    },
  },
};
</script>
