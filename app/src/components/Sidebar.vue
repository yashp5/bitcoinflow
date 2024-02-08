<template>
  <aside class="min-h-screen p-4 space-y-6 bg-gray-800 text-gray-100">
    <!-- Currency Selector -->
    <div class="space-y-2">
      <label for="currency" class="block text-gray-300">Currency</label>
      <select id="currency" class="w-full bg-gray-700 p-2 rounded">
        <option>Ethereum</option>
        <!-- Other options -->
      </select>
    </div>

    <!-- Minimum Premium Size Buttons -->
    <div class="space-y-2">
      <label class="block text-gray-300">Minimum Premium Size</label>
      <div class="flex flex-wrap gap-2">
        <button class="bg-gray-600 p-2 rounded">0.1 BTC</button>
        <button class="bg-gray-600 p-2 rounded">0.5 BTC</button>
        <button class="bg-gray-600 p-2 rounded">1 BTC</button>
        <button class="bg-gray-600 p-2 rounded">5 BTC</button>
        <button class="bg-gray-600 p-2 rounded">10 BTC</button>
        <button class="bg-gray-600 p-2 rounded">25 BTC</button>
      </div>
    </div>

    <!-- Minimum Unit Size Input -->
    <div class="space-y-2">
      <label for="unit-size" class="block text-gray-300">Minimum Unit Size</label>
      <input id="unit-size" type="text" class="w-full bg-gray-700 p-2 rounded" />
    </div>

    <!-- Call or Put Buttons -->
    <div class="flex gap-2">
      <label for="unit-size" class="block text-gray-300">Minimum Unit Size</label>
      <button class="bg-gray-600 p-2 rounded flex-1">Call</button>
      <button class="bg-gray-600 p-2 rounded flex-1">Put</button>
    </div>

    <!-- Side Buttons -->
    <div class="flex gap-2">
      <label for="unit-size" class="block text-gray-300">Minimum Unit Size</label>
      <button class="bg-gray-600 p-2 rounded flex-1">Buy</button>
      <button class="bg-gray-600 p-2 rounded flex-1">Sell</button>
    </div>

    <!-- Expiration Buttons -->
    <div class="flex gap-2">
      <button class="bg-gray-600 p-2 rounded flex-1">Clear</button>
      <button class="bg-gray-900 p-2 rounded flex-1">Go</button>
    </div>
  </aside>
</template>

<script>
export default {
  // Component logic here (if needed)
};
</script>

