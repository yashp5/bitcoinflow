<template>
  <aside class="min-h-screen p-4 space-y-6">
    <div class="space-y-2">
      <label for="currency" class="block">Currency</label>
      <select id="currency" class="w-full p-2">
        <option>Bitcoin</option>
        <option>Ethereum</option>
      </select>
    </div>

    <div class="space-y-2">
      <label class="block">Minimum Premium Size</label>
      <div class="flex flex-wrap gap-2">
        <button
          class="p-2 flex-1 bg-white hover:bg-gray-100 text-gray-800 font-semibold py-2 px-4 border border-gray-400 rounded shadow">0.1
          BTC</button>
        <button
          class="p-2 flex-1 bg-white hover:bg-gray-100 text-gray-800 font-semibold py-2 px-4 border border-gray-400 rounded shadow">0.5
          BTC</button>
        <button
          class="p-2 flex-1 bg-white hover:bg-gray-100 text-gray-800 font-semibold py-2 px-4 border border-gray-400 rounded shadow">1
          BTC</button>
        <button
          class="p-2 flex-1 bg-white hover:bg-gray-100 text-gray-800 font-semibold py-2 px-4 border border-gray-400 rounded shadow">5
          BTC</button>
        <button
          class="p-2 flex-1 bg-white hover:bg-gray-100 text-gray-800 font-semibold py-2 px-4 border border-gray-400 rounded shadow">10
          BTC</button>
        <button
          class="p-2 flex-1 bg-white hover:bg-gray-100 text-gray-800 font-semibold py-2 px-4 border border-gray-400 rounded shadow">25
          BTC</button>
      </div>
    </div>

    <div class="space-y-2">
      <label for="unit-size" class="block">Minimum Unit Size</label>
      <input id="unit-size" type="text" class="w-full p-2" />
    </div>

    <div class="space-y-2">
      <label for="unit-size" class="block">Call or Put</label>
      <div class="flex flex-wrap gap-2">
        <button
          class="p-2 flex-1 bg-white hover:bg-gray-100 text-gray-800 font-semibold py-2 px-4 border border-gray-400 rounded shadow">Call</button>
        <button
          class="p-2 flex-1 bg-white hover:bg-gray-100 text-gray-800 font-semibold py-2 px-4 border border-gray-400 rounded shadow">Put</button>
      </div>
    </div>

    <div class="space-y-2">
      <label for="unit-size" class="block">Side</label>
      <div class="flex flex-wrap gap-2">
        <button
          class="p-2 flex-1 bg-white hover:bg-gray-100 text-gray-800 font-semibold py-2 px-4 border border-gray-400 rounded shadow">Buy</button>
        <button
          class="p-2 flex-1 bg-white hover:bg-gray-100 text-gray-800 font-semibold py-2 px-4 border border-gray-400 rounded shadow">Sell</button>
      </div>
    </div>

    <div class="flex gap-2">
      <button
        class="p-2 flex-1 bg-white hover:bg-gray-100 text-gray-800 font-semibold py-2 px-4 border border-gray-400 rounded shadow">Clear</button>
      <button
        class="p-2 flex-1 bg-white hover:bg-gray-100 text-gray-800 font-semibold py-2 px-4 border border-gray-400 rounded shadow">Go</button>
  </div>
</aside></template>

<script>
export default {
  // Component logic here (if needed)
};
</script>

