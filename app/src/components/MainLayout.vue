<template>
    <div class="flex">
      <Grid/>
    </div>
  </template>

  <script>
  import Grid from './Grid.vue';
  import Sidebar from './Sidebar.vue';

  export default {
    components: { Grid, Sidebar },
    data() {
      return {
        isSidebarOpen: false,
      };
    },
    methods: {
      toggleSidebar() {
        this.isSidebarOpen = !this.isSidebarOpen;
      },
    },
  };
  </script>
  